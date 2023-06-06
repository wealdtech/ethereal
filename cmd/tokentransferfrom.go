// Copyright © 2017-2019 Weald Technology Trading
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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
)

var (
	tokenTransferFromAmount      string
	tokenTransferFromFromAddress string
	tokenTransferFromToAddress   string
	tokenTransferFromByAddress   string
)

// tokenTransferFromCmd represents the token transferfrom command.
var tokenTransferFromCmd = &cobra.Command{
	Use:   "transferfrom",
	Short: "Transfer tokens from a third-party address to another address",
	Long: `Transfer tokens from one third-party address to another, as part of a user's allowance .  For example:

    ethereal token transferfrom --token=omg --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --by= --amount=10 --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(tokenTransferFromFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(tokenTransferFromFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", tokenTransferFromFromAddress))

		cli.Assert(tokenTransferFromToAddress != "", quiet, "--to is required")
		toAddress, err := c.Resolve(tokenTransferFromToAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", tokenTransferFromToAddress))

		cli.Assert(tokenTransferFromByAddress != "", quiet, "--by is required")
		byAddress, err := c.Resolve(tokenTransferFromByAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve by address %s", tokenTransferFromByAddress))

		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		decimals, err := token.Decimals(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain token decimals")

		cli.Assert(tokenTransferFromAmount != "", quiet, "--amount is required")
		amount, err := util.StringToTokenValue(tokenTransferFromAmount, decimals)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address.
		balance, err := token.BalanceOf(nil, fromAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
		cli.Assert(balance.Cmp(amount) > 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", util.TokenValueToString(balance, decimals, false)))

		// Obtain the allowance of the address.
		allowance, err := token.Allowance(nil, fromAddress, byAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain allowance of address from which to send funds")
		cli.Assert(allowance.Cmp(amount) > 0, quiet, fmt.Sprintf("Allowance of %s insufficient for transfer", util.TokenValueToString(allowance, decimals, false)))

		opts, err := generateTxOpts(byAddress)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")

		signedTx, err := token.TransferFrom(opts, fromAddress, toAddress, amount)
		cli.ErrCheck(err, quiet, "Failed to create transaction")
		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
			os.Exit(exitSuccess)
		}

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":          "token",
			"command":        "transferfrom",
			"token":          tokenStr,
			"tokenholder":    fromAddress.Hex(),
			"tokenrecipient": toAddress.Hex(),
			"tokenoperator":  byAddress.Hex(),
			"tokenamount":    amount.String(),
		}, true)
	},
}

func init() {
	tokenCmd.AddCommand(tokenTransferFromCmd)
	tokenFlags(tokenTransferFromCmd)
	tokenTransferFromCmd.Flags().StringVar(&tokenTransferFromAmount, "amount", "", "Amount to transfer")
	tokenTransferFromCmd.Flags().StringVar(&tokenTransferFromFromAddress, "from", "", "Address from which to transfer tokens")
	tokenTransferFromCmd.Flags().StringVar(&tokenTransferFromToAddress, "to", "", "Address to which to transfer tokens")
	tokenTransferFromCmd.Flags().StringVar(&tokenTransferFromByAddress, "by", "", "Address allowed to transfer tokens")
	addTransactionFlags(tokenTransferFromCmd, "the address from which to transfer tokens")
}
