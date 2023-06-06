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
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/conn"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	etherTransferAmount      string
	etherTransferFromAddress string
	etherTransferToAddress   string
	etherTransferData        string
)

// etherTransferCmd represents the ether transfer command.
var etherTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer funds to a given address",
	Long: `Transfer Ether funds from one address to another.  For example:

    ethereal ether transfer --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=1.5ether --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Aliases: []string{"send"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherTransferFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(etherTransferFromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain from address for transfer")

		cli.Assert(etherTransferToAddress != "", quiet, "--to is required")
		toAddress, err := c.Resolve(etherTransferToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain to address for transfer")

		cli.Assert(etherTransferAmount != "", quiet, "--amount is required")
		amount, err := string2eth.StringToWei(etherTransferAmount)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address..
		if !offline {
			ctx, cancel := localContext()
			defer cancel()
			balance, err := c.Client().BalanceAt(ctx, fromAddress, nil)
			cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
			cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", string2eth.WeiToString(balance, true)))
		}

		// Turn the data string in to hex.
		etherTransferData = strings.TrimPrefix(etherTransferData, "0x")
		if len(etherTransferData)%2 == 1 {
			// Doesn't like odd numbers.
			etherTransferData = "0" + etherTransferData
		}
		data, err := hex.DecodeString(etherTransferData)
		cli.ErrCheck(err, quiet, "Failed to parse data")

		var gasLimit *uint64
		limit := uint64(viper.GetInt64("gaslimit"))
		if limit > 0 {
			gasLimit = &limit
		}

		// Create and sign the transaction.
		signedTx, err := c.CreateSignedTransaction(context.Background(), &conn.TransactionData{
			From:     fromAddress,
			To:       &toAddress,
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
				"group":   "ether",
				"command": "transfer",
			}, true)
		}
	},
}

func init() {
	etherCmd.AddCommand(etherTransferCmd)
	etherTransferCmd.Flags().StringVar(&etherTransferAmount, "amount", "", "Amount of Ether to transfer")
	etherTransferCmd.Flags().StringVar(&etherTransferFromAddress, "from", "", "Address from which to transfer Ether")
	etherTransferCmd.Flags().StringVar(&etherTransferToAddress, "to", "", "Address to which to transfer Ether")
	etherTransferCmd.Flags().StringVar(&etherTransferData, "data", "", "data to send with transaction (as a hex string)")
	addTransactionFlags(etherTransferCmd, "the address from which to transfer Ether")
}
