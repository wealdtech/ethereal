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
	tokenAllowanceRaw            bool
	tokenAllowanceHolderAddress  string
	tokenAllowanceSpenderAddress string
)

// tokenAllowanceCmd represents the token allowance command.
var tokenAllowanceCmd = &cobra.Command{
	Use:   "allowance",
	Short: "Obtain the token allowance for a holder and spender",
	Long: `Obtain the token allowance for a holder and spender.  For example:

    ethereal token allowance --token=omg --holder=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --spender=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d

In quiet mode this will return 0 if the allowance is greater than 0, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(tokenAllowanceHolderAddress != "", quiet, "--holder is required")
		holderAddress, err := c.Resolve(tokenAllowanceHolderAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve holder address %s", tokenAllowanceHolderAddress))

		cli.Assert(tokenAllowanceSpenderAddress != "", quiet, "--spender is required")
		spenderAddress, err := c.Resolve(tokenAllowanceSpenderAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain spender address")

		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		decimals, err := token.Decimals(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain token decimals")

		allowance, err := token.Allowance(nil, holderAddress, spenderAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain allowance")

		if quiet {
			if allowance.Cmp(big.NewInt(0)) == 0 {
				os.Exit(exitFailure)
			}
			os.Exit(exitSuccess)
		}

		if tokenAllowanceRaw {
			fmt.Printf("%s\n", allowance.String())
		} else {
			fmt.Printf("%s\n", util.TokenValueToString(allowance, decimals, false))
		}
	},
}

func init() {
	tokenCmd.AddCommand(tokenAllowanceCmd)
	tokenFlags(tokenAllowanceCmd)
	tokenAllowanceCmd.Flags().BoolVar(&tokenAllowanceRaw, "raw", false, "Display raw output (no decimals)")
	tokenAllowanceCmd.Flags().StringVar(&tokenAllowanceHolderAddress, "holder", "", "Address that holds tokens")
	tokenAllowanceCmd.Flags().StringVar(&tokenAllowanceSpenderAddress, "spender", "", "Address that can spend tokens")
}
