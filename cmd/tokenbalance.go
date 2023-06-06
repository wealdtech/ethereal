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
	"fmt"
	"math/big"
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
)

var (
	tokenBalanceHolderAddress string
	tokenBalanceRaw           bool
)

// tokenBalanceCmd represents the token balance command.
var tokenBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Obtain the token balance for an address",
	Long: `Obtain the token balance for an address.  For example:

    ethereal token balance --token=omg --holder=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the balance is greater than 0, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(tokenBalanceHolderAddress != "", quiet, "--holder is required")
		address, err := c.Resolve(tokenBalanceHolderAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve holder address %s", tokenBalanceHolderAddress))

		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		decimals, err := token.Decimals(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain token decimals")

		balance, err := token.BalanceOf(nil, address)
		cli.ErrCheck(err, quiet, "Failed to obtain token balance")

		if quiet {
			if balance.Cmp(big.NewInt(0)) == 0 {
				os.Exit(exitFailure)
			}
			os.Exit(exitSuccess)
		}

		if tokenBalanceRaw {
			fmt.Printf("%s\n", balance.String())
		} else {
			fmt.Printf("%s\n", util.TokenValueToString(balance, decimals, false))
		}
	},
}

func init() {
	tokenFlags(tokenBalanceCmd)
	tokenCmd.AddCommand(tokenBalanceCmd)
	tokenBalanceCmd.Flags().BoolVar(&tokenBalanceRaw, "raw", false, "Display raw output (no decimals)")
	tokenBalanceCmd.Flags().StringVar(&tokenBalanceHolderAddress, "holder", "", "Holder of tokens")
}
