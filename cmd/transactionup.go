// Copyright © 2017 Weald Technology Trading
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
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	etherutils "github.com/orinocopay/go-etherutils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/cli"
)

// transactionUpCmd represents the transaction up command
var transactionUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Increase the gas cost for a pending transaction",
	Long: `Increase the gas cost for a pending transaction.  For example:

    ethereal transaction up --gasprice=20gwei --passphrase=secret --transaction=0x454d2274155cce506359de6358785ce5366f6c13e825263674c272eec8532c0c

If no gas price is supplied then it will default to 10% higher than the current gas price for the transaction.

In quiet mode this will return 0 if the transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(transactionStr != "", quiet, "--transaction is required")
		txHash := common.HexToHash(transactionStr)
		ctx, cancel := localContext()
		defer cancel()
		tx, pending, err := client.TransactionByHash(ctx, txHash)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain transaction %s", txHash.Hex()))
		cli.Assert(pending, quiet, fmt.Sprintf("Transaction %s has already been mined", txHash.Hex()))

		minGasPrice := big.NewInt(0).Add(big.NewInt(0).Add(tx.GasPrice(), big.NewInt(0).Div(tx.GasPrice(), big.NewInt(10))), big.NewInt(10))
		if viper.GetString("gasprice") == "" {
			// No gas price supplied; use the calculated minimum
			gasPrice = minGasPrice
		} else {
			// Gas price supplied; ensure it is at least 10% more than the current gas price
			cli.Assert(gasPrice.Cmp(minGasPrice) >= 0, quiet, fmt.Sprintf("Gas price must be at least %s", etherutils.WeiToString(minGasPrice, true)))
		}

		// Create and sign the transaction
		fromAddress, err := txFrom(tx)
		cli.ErrCheck(err, quiet, "Failed to obtain from address")

		nonce = int64(tx.Nonce())
		signedTx, err := createSignedTransaction(fromAddress, tx.To(), tx.Value(), tx.Gas(), tx.Data())
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			ctx, cancel := localContext()
			defer cancel()
			err = client.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")

			if tx.To() == nil {
				setupLogging()
				log.WithFields(log.Fields{
					"group":         "transaction",
					"command":       "up",
					"from":          fromAddress.Hex(),
					"amount":        tx.Value().String(),
					"data":          hex.EncodeToString(tx.Data()),
					"networkid":     chainID,
					"gas":           signedTx.Gas(),
					"gasprice":      signedTx.GasPrice().String(),
					"transactionid": signedTx.Hash().Hex(),
				}).Info("success")
			} else {
				setupLogging()
				log.WithFields(log.Fields{
					"group":         "transaction",
					"command":       "up",
					"from":          fromAddress.Hex(),
					"to":            tx.To().Hex(),
					"amount":        tx.Value().String(),
					"data":          hex.EncodeToString(tx.Data()),
					"networkid":     chainID,
					"gas":           signedTx.Gas(),
					"gasprice":      signedTx.GasPrice().String(),
					"transactionid": signedTx.Hash().Hex(),
				}).Info("success")
			}

			if quiet {
				os.Exit(0)
			}
			fmt.Println(signedTx.Hash().Hex())
		}
	},
}

func init() {
	transactionCmd.AddCommand(transactionUpCmd)
	transactionFlags(transactionUpCmd)
	addTransactionFlags(transactionUpCmd, "the address that holds the funds")
}
