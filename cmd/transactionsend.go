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
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	transactionSendAmount      string
	transactionSendFromAddress string
	transactionSendToAddress   string
	transactionSendData        string
	transactionSendRaw         string
	transactionSendRepeat      int
)

// transactionSendCmd represents the transaction send command.
var transactionSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a transaction",
	Long: `Send a transaction.  For example:

    ethereal transaction send --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x2ab7150Bba7D5F181b3aF5623e52b15bB1054845	 --amount=1ether --passphrase=secret --data=0x12345

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		if transactionSendRaw != "" {
			// Send raw transactions.
			signedTxs := make([]*types.Transaction, 0)

			if !strings.HasPrefix(transactionSendRaw, "0x") {
				// Data is a file.
				data, err := os.ReadFile(transactionSendRaw)
				cli.ErrCheck(err, quiet, "Failed to read raw transaction from filesystem")
				lines := bytes.Split(bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n")), []byte("\n"))
				for i := range lines {
					if len(lines[i]) > 2 {
						data, err := hex.DecodeString(strings.TrimPrefix(string(lines[i]), "0x"))
						cli.ErrCheck(err, quiet, "Failed to decode transaction")
						signedTx := &types.Transaction{}
						stream := rlp.NewStream(bytes.NewReader(data), 0)
						err = signedTx.DecodeRLP(stream)
						cli.ErrCheck(err, quiet, "Failed to decode transaction")
						signedTxs = append(signedTxs, signedTx)
					}
				}
			} else {
				// Data is a direct transaction.
				data, err := hex.DecodeString(strings.TrimPrefix(transactionSendRaw, "0x"))
				cli.ErrCheck(err, quiet, "Failed to decode data")
				// Decode the raw transaction.
				signedTx := &types.Transaction{}
				stream := rlp.NewStream(bytes.NewReader(data), 0)
				err = signedTx.DecodeRLP(stream)
				cli.ErrCheck(err, quiet, "Failed to decode transaction")
				signedTxs = append(signedTxs, signedTx)
			}

			for i := range signedTxs {
				if offline {
					if !quiet {
						buf := new(bytes.Buffer)
						cli.ErrCheck(signedTxs[i].EncodeRLP(buf), quiet, "failed to encode transaction")
						fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
					}
				} else {
					err = c.SendTransaction(context.Background(), signedTxs[i])
					cli.ErrCheck(err, quiet, "Failed to send transaction")

					logTransaction(signedTxs[i], log.Fields{
						"group":   "transaction",
						"command": "send",
					})

					if !quiet {
						fmt.Println(signedTxs[i].Hash().Hex())
					}
				}
			}
			os.Exit(exitSuccess)
		}

		cli.Assert(transactionSendFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(transactionSendFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", transactionSendFromAddress))

		var toAddress *common.Address
		if transactionSendToAddress == "" {
			// This is valid because it can be a contract creation, but only if there is data as well.
			cli.Assert(transactionSendData != "", quiet, "Transactions without a to address are contract creations and must have data")
		} else {
			tmp, err := c.Resolve(transactionSendToAddress)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", transactionSendToAddress))
			toAddress = &tmp
		}

		var amount *big.Int
		if transactionSendAmount == "" {
			amount = big.NewInt(0)
		} else {
			amount, err = string2eth.StringToWei(transactionSendAmount)
			cli.ErrCheck(err, quiet, "Invalid amount")
		}

		// Obtain the balance of the address.
		if c.Client() != nil {
			ctx, cancel := localContext()
			defer cancel()
			balance, err := c.Client().BalanceAt(ctx, fromAddress, nil)
			cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
			cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", string2eth.WeiToString(balance, true)))
		}

		var gasLimit *uint64
		limit := uint64(viper.GetInt64("gaslimit"))
		if limit > 0 {
			gasLimit = &limit
		}

		// Turn the data string in to hex.
		transactionSendData = strings.TrimPrefix(transactionSendData, "0x")
		if len(transactionSendData)%2 == 1 {
			// Doesn't like odd numbers.
			transactionSendData = "0" + transactionSendData
		}
		data, err := hex.DecodeString(transactionSendData)
		cli.ErrCheck(err, quiet, "Failed to parse data")

		for i := 0; i < transactionSendRepeat; i++ {
			// Create and sign the transaction.
			signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
				From:     fromAddress,
				To:       toAddress,
				Value:    amount,
				GasLimit: gasLimit,
				Data:     data,
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
					"group":   "transaction",
					"command": "send",
				}, false)
			}
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
	transactionSendCmd.Flags().IntVar(&transactionSendRepeat, "repeat", 1, "Number of times to repeat sending the transaction (incrementing the nonce each time)")
	addTransactionFlags(transactionSendCmd, "the address from which to transfer Ether")
}
