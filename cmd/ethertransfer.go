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
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

var etherTransferAmount string
var etherTransferToAddress string
var etherTransferData string

// etherTransferCmd represents the ether transfer command
var etherTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer funds to a given address",
	Long: `Transfer Ether funds from one address to another.  For example:

    ethereal ether transfer --to=x --amount=y --passphrase=secret 0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the balance is greater than 0, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(args[0] != "", quiet, "Sender address is required")

		fromAddress, err := ens.Resolve(client, args[0])
		cli.ErrCheck(err, quiet, "Failed to obtain sender for transfer")

		toAddress, err := ens.Resolve(client, etherTransferToAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain recipient for transfer")

		cli.Assert(etherTransferAmount != "", quiet, "Require an amount to transfer with --to")
		amount, err := etherutils.StringToWei(etherTransferAmount)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address
		balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
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
		signedTx, err := createSignedTransaction(fromAddress, toAddress, amount, data)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		err = client.SendTransaction(context.Background(), signedTx)
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		if quiet {
			os.Exit(0)
		}
		fmt.Println(signedTx.Hash().Hex())
	},
}

func init() {
	etherCmd.AddCommand(etherTransferCmd)
	etherTransferCmd.Flags().StringVar(&etherTransferAmount, "amount", "", "Amount of Ether to transfer")
	etherTransferCmd.Flags().StringVar(&etherTransferToAddress, "to", "", "Address to which to transfer Ether")
	etherTransferCmd.Flags().StringVar(&etherTransferData, "data", "", "data to send with transaction (as a hex string)")
	addTransactionFlags(etherTransferCmd, "Passphrase for the address that holds the funds")
}
