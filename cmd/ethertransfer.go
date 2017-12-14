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
	"os"
	"strings"

	"github.com/orinocopay/go-etherutils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
)

var etherTransferAmount string
var etherTransferFromAddress string
var etherTransferToAddress string
var etherTransferData string

// etherTransferCmd represents the ether transfer command
var etherTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer funds to a given address",
	Long: `Transfer Ether funds from one address to another.  For example:

    ethereal ether transfer --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=1.5ether --passphrase=secret

In quiet mode this will return 0 if the transfer transaction is successfully sent, otherwise 1.`,
	Aliases: []string{"send"},
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherTransferFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, etherTransferFromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain from address for transfer")

		cli.Assert(etherTransferToAddress != "", quiet, "--to is required")
		toAddress, err := ens.Resolve(client, etherTransferToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain to address for transfer")

		cli.Assert(etherTransferAmount != "", quiet, "--amount is required")
		amount, err := etherutils.StringToWei(etherTransferAmount)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address
		ctx, cancel := localContext()
		defer cancel()
		balance, err := client.BalanceAt(ctx, fromAddress, nil)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", etherutils.WeiToString(balance, true)))

		// Turn the data string in to hex
		etherTransferData = strings.TrimPrefix(etherTransferData, "0x")
		if len(etherTransferData)%2 == 1 {
			// Doesn't like odd numbers
			etherTransferData = "0" + etherTransferData
		}
		data, err := hex.DecodeString(etherTransferData)
		cli.ErrCheck(err, quiet, "Failed to parse data")

		// Create and sign the transaction
		signedTx, err := createSignedTransaction(fromAddress, &toAddress, amount, gasLimit, data)
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

			log.WithFields(log.Fields{
				"group":         "ether",
				"command":       "transfer",
				"from":          fromAddress.Hex(),
				"to":            toAddress.Hex(),
				"amount":        amount.String(),
				"networkid":     chainID,
				"gas":           signedTx.Gas().String(),
				"gasprice":      signedTx.GasPrice().String(),
				"transactionid": signedTx.Hash().Hex(),
			}).Info("success")

			if quiet {
				os.Exit(0)
			}
			fmt.Println(signedTx.Hash().Hex())
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
