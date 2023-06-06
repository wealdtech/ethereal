// Copyright © 2017-2022 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	string2eth "github.com/wealdtech/go-string2eth"
)

// transactionUpCmd represents the transaction up command.
var transactionUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Increase the gas cost for a pending transaction",
	Long: `Increase the gas cost for a pending transaction.  For example:

    ethereal transaction up --passphrase=secret --transaction=0x454d2274155cce506359de6358785ce5366f6c13e825263674c272eec8532c0c

If no gas price is supplied then it will default to just over 10% higher than the current gas price for the transaction.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(transactionStr != "", quiet, "--transaction is required")
		txHash := common.HexToHash(transactionStr)
		ctx, cancel := localContext()
		defer cancel()
		tx, pending, err := c.Client().TransactionByHash(ctx, txHash)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain transaction %s", txHash.Hex()))
		cli.Assert(pending, quiet, fmt.Sprintf("Transaction %s has already been mined", txHash.Hex()))

		// Increase priority fee by 10% (+1 wei, to avoid rounding issues).
		feePerGas := new(big.Int).Add(new(big.Int).Add(tx.GasFeeCap(), new(big.Int).Div(tx.GasFeeCap(), big.NewInt(10))), big.NewInt(1))
		// Increase priority fee by 10% (+1 wei, to avoid rounding issues).
		priorityFeePerGas := new(big.Int).Add(new(big.Int).Add(tx.GasTipCap(), new(big.Int).Div(tx.GasTipCap(), big.NewInt(10))), big.NewInt(1))

		// Ensure that the total fee per gas does not exceed the max allowed.
		totalFeePerGas := new(big.Int).Add(feePerGas, priorityFeePerGas)
		if viper.GetString("max-fee-per-gas") == "" {
			viper.Set("max-fee-per-gas", "200gwei")
		}
		maxFeePerGas, err := string2eth.StringToWei(viper.GetString("max-fee-per-gas"))
		cli.ErrCheck(err, quiet, "failed to obtain max fee per gas")
		cli.Assert(totalFeePerGas.Cmp(maxFeePerGas) <= 0, quiet, fmt.Sprintf("increased total fee per gas of %s too high; increase with --max-fee-per-gas if you are sure you want to do this", string2eth.WeiToString(totalFeePerGas, true)))

		// Create and sign the transaction.
		fromAddress, err := types.Sender(signer, tx)
		cli.ErrCheck(err, quiet, "Failed to obtain from address")

		nonce := int64(tx.Nonce())
		gasLimit := tx.Gas()
		signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
			From:                 fromAddress,
			To:                   tx.To(),
			Nonce:                &nonce,
			Value:                tx.Value(),
			GasLimit:             &gasLimit,
			MaxFeePerGas:         feePerGas,
			MaxPriorityFeePerGas: priorityFeePerGas,
			Data:                 tx.Data(),
		})
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			err = c.SendTransaction(context.Background(), signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")
			handleSubmittedTransaction(signedTx, log.Fields{
				"group":                    "transaction",
				"command":                  "up",
				"old-fee-per-gas":          tx.GasFeeCap().String(),
				"old-priority-fee-per-gas": tx.GasTipCap().String(),
			}, true)
		}
	},
}

func init() {
	transactionCmd.AddCommand(transactionUpCmd)
	transactionFlags(transactionUpCmd)
	addTransactionFlags(transactionUpCmd, "the address that holds the funds")
}
