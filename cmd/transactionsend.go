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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var transactionSendAmount string
var transactionSendFromAddress string
var transactionSendToAddress string
var transactionSendData string
var transactionSendRaw string

// transactionSendCmd represents the transaction send command
var transactionSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a transaction",
	Long: `Send a transaction.  For example:

    ethereal transaction send --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x2ab7150Bba7D5F181b3aF5623e52b15bB1054845	 --amount=1ether --passphrase=secret --data=0x12345

In quiet mode this will return 0 if the transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		if transactionSendRaw != "" {
			// Send a raw transaction

			// Decode the raw transaction
			data, err := hex.DecodeString(strings.TrimPrefix(transactionSendRaw, "0x"))
			cli.ErrCheck(err, quiet, "Failed to decode data")
			signedTx := &types.Transaction{}
			stream := rlp.NewStream(bytes.NewReader(data), 0)
			err = signedTx.DecodeRLP(stream)
			cli.ErrCheck(err, quiet, "Failed to decode transaction")

			ctx, cancel := localContext()
			defer cancel()
			err = client.SendTransaction(ctx, signedTx)
			cli.ErrCheck(err, quiet, "Failed to send transaction")

			fromAddress, err := txFrom(signedTx)
			cli.ErrCheck(err, quiet, "Failed to obtain from address")

			log.WithFields(log.Fields{
				"group":         "transaction",
				"command":       "send",
				"from":          fromAddress.Hex(),
				"to":            signedTx.To().Hex(),
				"amount":        signedTx.Value().String(),
				"data":          hex.EncodeToString(signedTx.Data()),
				"networkid":     chainID,
				"gas":           signedTx.Gas().String(),
				"gasprice":      signedTx.GasPrice().String(),
				"transactionid": signedTx.Hash().Hex(),
			}).Info("success")

			if !quiet {
				fmt.Println(signedTx.Hash().Hex())
			}
			os.Exit(0)
		}

		cli.Assert(transactionSendFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, transactionSendFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", transactionSendFromAddress))

		var toAddress *common.Address
		if transactionSendToAddress == "" {
			// This is valid because it can be a contract creation, but only if there is data as well
			cli.Assert(transactionSendData != "", quiet, "Transactions without a to address are contract creations and must have data")
		} else {
			tmp, err := ens.Resolve(client, transactionSendToAddress)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", transactionSendToAddress))
			toAddress = &tmp
		}

		var amount *big.Int
		if transactionSendAmount == "" {
			amount = big.NewInt(0)
		} else {
			amount, err = etherutils.StringToWei(transactionSendAmount)
			cli.ErrCheck(err, quiet, "Invalid amount")
		}

		// Obtain the balance of the address
		ctx, cancel := localContext()
		defer cancel()
		balance, err := client.BalanceAt(ctx, fromAddress, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", etherutils.WeiToString(balance, true)))

		// Turn the data string in to hex
		transactionSendData = strings.TrimPrefix(transactionSendData, "0x")
		if len(transactionSendData)%2 == 1 {
			// Doesn't like odd numbers
			transactionSendData = "0" + transactionSendData
		}
		data, err := hex.DecodeString(transactionSendData)
		cli.ErrCheck(err, quiet, "Failed to parse data")

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, toAddress, amount, gasLimit, data)
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

			if toAddress == nil {
				log.WithFields(log.Fields{
					"group":         "transaction",
					"command":       "send",
					"from":          fromAddress.Hex(),
					"to":            toAddress.Hex(),
					"amount":        amount.String(),
					"data":          hex.EncodeToString(data),
					"networkid":     chainID,
					"gas":           signedTx.Gas().String(),
					"gasprice":      signedTx.GasPrice().String(),
					"transactionid": signedTx.Hash().Hex(),
				}).Info("success")
			} else {
				log.WithFields(log.Fields{
					"group":         "transaction",
					"command":       "send",
					"from":          fromAddress.Hex(),
					"to":            toAddress.Hex(),
					"amount":        amount.String(),
					"data":          hex.EncodeToString(data),
					"networkid":     chainID,
					"gas":           signedTx.Gas().String(),
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
	transactionCmd.AddCommand(transactionSendCmd)
	transactionSendCmd.Flags().StringVar(&transactionSendAmount, "amount", "", "Amount of Ether to transfer")
	transactionSendCmd.Flags().StringVar(&transactionSendFromAddress, "from", "", "Address from which to transfer Ether")
	transactionSendCmd.Flags().StringVar(&transactionSendToAddress, "to", "", "Address to which to transfer Ether")
	transactionSendCmd.Flags().StringVar(&transactionSendData, "data", "", "data to send with transaction (as a hex string)")
	transactionSendCmd.Flags().StringVar(&transactionSendRaw, "raw", "", "raw transaction (as a hex string).  This overrides all other options")
	addTransactionFlags(transactionSendCmd, "Passphrase for the address from which to transfer Ether")
}
