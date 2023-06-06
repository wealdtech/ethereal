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

var (
	transactionCancelAmount    string
	transactionCancelToAddress string
)

// transactionCancelCmd represents the transaction up command.
var transactionCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a pending transaction",
	Long: `Cancel a pending transaction.  For example:

    ethereal transaction cancel --transaction=0x454d2274155cce506359de6358785ce5366f6c13e825263674c272eec8532c0c

Note that in reality Ethereum has no notion of cancelling transactions so instead the transaction is replaced with a new transaction that does nothing.  For this command to succeed transaction's maximum base fee and priority fee must both be increased by 10% over that of the existing transaction; this will happen automatically.

The cancellation transaction will cost 21000 gas.

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
		cli.ErrCheck(err, quiet, "Failed to obtain sender")

		nonce := int64(tx.Nonce())
		signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
			From:                 fromAddress,
			To:                   &fromAddress,
			Nonce:                &nonce,
			MaxFeePerGas:         feePerGas,
			MaxPriorityFeePerGas: priorityFeePerGas,
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
				"group":            "transaction",
				"command":          "cancel",
				"oldtransactionid": txHash.Hex(),
			}, true)
		}
	},
}

func init() {
	transactionCmd.AddCommand(transactionCancelCmd)
	transactionFlags(transactionCancelCmd)
	transactionCancelCmd.Flags().StringVar(&transactionCancelAmount, "amount", "", "Amount of Ether to transfer")
	transactionCancelCmd.Flags().StringVar(&transactionCancelToAddress, "to", "", "Address to which to transfer Ether")
	addTransactionFlags(transactionCancelCmd, "the address that holds the funds")
}
