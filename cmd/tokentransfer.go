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
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
)

var (
	tokenTransferAmount      string
	tokenTransferFromAddress string
	tokenTransferToAddress   string
	tokenTransferDecimals    string
)

// tokenTransferCmd represents the token transfer command.
var tokenTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer tokens to a given address",
	Long: `Transfer token from one address to another.  For example:

    ethereal token transfer --token=omg --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=10 --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(tokenTransferFromAddress != "", quiet, "--from is required")
		fromAddress, err := c.Resolve(tokenTransferFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", tokenTransferFromAddress))

		cli.Assert(tokenTransferToAddress != "", quiet, "--to is required")
		toAddress, err := c.Resolve(tokenTransferToAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", tokenTransferToAddress))

		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		var decimals uint8
		if offline {
			gasLimit := viper.GetInt64("gaslimit")
			cli.Assert(gasLimit > 0, quiet, "--gaslimit is required if offline")
			cli.Assert(tokenTransferDecimals != "", quiet, "--decimals is required if offline")
			tmpDecimals, err := strconv.Atoi(tokenTransferDecimals)
			cli.ErrCheck(err, quiet, "Failed to obtain token decimals")
			decimals = uint8(tmpDecimals)
		} else {
			decimals, err = token.Decimals(nil)
			cli.ErrCheck(err, quiet, "Failed to obtain token decimals")
		}

		cli.Assert(tokenTransferAmount != "", quiet, "--amount is required")
		amount, err := util.StringToTokenValue(tokenTransferAmount, decimals)
		cli.ErrCheck(err, quiet, "Invalid amount")

		// Obtain the balance of the address (if online).
		if !offline {
			balance, err := token.BalanceOf(nil, fromAddress)
			cli.ErrCheck(err, quiet, "Failed to obtain balance of address from which to send funds")
			cli.Assert(balance.Cmp(amount) >= 0, quiet, fmt.Sprintf("Balance of %s insufficient for transfer", util.TokenValueToString(balance, decimals, false)))
		}

		opts, err := generateTxOpts(fromAddress)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")

		signedTx, err := token.Transfer(opts, toAddress, amount)
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
			"command":        "transfer",
			"token":          tokenStr,
			"tokenholder":    fromAddress.Hex(),
			"tokenrecipient": toAddress.Hex(),
			"tokenamount":    amount.String(),
		}, true)
	},
}

func init() {
	tokenCmd.AddCommand(tokenTransferCmd)
	tokenFlags(tokenTransferCmd)
	tokenTransferCmd.Flags().StringVar(&tokenTransferAmount, "amount", "", "Amount to transfer")
	tokenTransferCmd.Flags().StringVar(&tokenTransferFromAddress, "from", "", "Address from which to transfer tokens")
	tokenTransferCmd.Flags().StringVar(&tokenTransferToAddress, "to", "", "Address to which to transfer tokens")
	tokenTransferCmd.Flags().StringVar(&tokenTransferDecimals, "decimals", "18", "Number of decimals for the transfer (only required if offline)")
	addTransactionFlags(tokenTransferCmd, "the address from which to transfer tokens")
}
