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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/util"
	ens "github.com/wealdtech/go-ens"
)

var tokenApproveAmount string
var tokenApproveHolderAddress string
var tokenApproveSpenderAddress string
var tokenApproveData string

// tokenApproveCmd represents the token approve command
var tokenApproveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve an address to transfer tokens",
	Long: `Approve one address to spend tokens on behalf of another.  For example:

    ethereal token approve --token=omg --holder=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --spender=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=10 --passphrase=secret

In quiet mode this will return 0 if the approval transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(tokenApproveHolderAddress != "", quiet, "--holder is required")
		holderAddress, err := ens.Resolve(client, tokenApproveHolderAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve holder address %s", tokenApproveHolderAddress))

		cli.Assert(tokenApproveSpenderAddress != "", quiet, "--spender is required")
		spenderAddress, err := ens.Resolve(client, tokenApproveSpenderAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve spender address %s", tokenApproveSpenderAddress))

		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		decimals, err := token.Decimals(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain token decimals")

		cli.Assert(tokenApproveAmount != "", quiet, "--amount is required")
		amount, err := util.StringToTokenValue(tokenApproveAmount, decimals)
		cli.ErrCheck(err, quiet, "Invalid amount")

		allowance, err := token.Allowance(nil, holderAddress, spenderAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain allowance")

		cli.Assert(allowance.Cmp(big.NewInt(0)) == 0 || amount.Cmp(big.NewInt(0)) == 0, quiet, fmt.Sprintf("Allowance is currently %s; it must be set to zero before being changed to avoid a potential double spend", util.TokenValueToString(allowance, decimals, false)))

		opts, err := generateTxOpts(holderAddress)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")

		signedTx, err := token.Approve(opts, spenderAddress, amount)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			logTransaction(signedTx, log.Fields{
				"group":   "token",
				"command": "approve",
				"token":   tokenStr,
				"holder":  holderAddress.Hex(),
				"spender": spenderAddress.Hex(),
				"amount":  amount.String(),
			})

			if quiet {
				os.Exit(0)
			}

			fmt.Println(signedTx.Hash().Hex())
		}
	},
}

func init() {
	tokenCmd.AddCommand(tokenApproveCmd)
	tokenFlags(tokenApproveCmd)
	tokenApproveCmd.Flags().StringVar(&tokenApproveAmount, "amount", "", "Amount to approve")
	tokenApproveCmd.Flags().StringVar(&tokenApproveHolderAddress, "holder", "", "Address that holds tokens")
	tokenApproveCmd.Flags().StringVar(&tokenApproveSpenderAddress, "spender", "", "Address that can spend tokens")
	addTransactionFlags(tokenApproveCmd, "the address from which to approve tokens")
}
