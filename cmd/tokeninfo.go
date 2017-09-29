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
	"fmt"
	"os"

	"github.com/orinocopay/go-etherutils/cli"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/util"
)

// tokenInfoCmd represents the token info command
var tokenInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a token",
	Long: `Obtain information about a token.  For example:

    ethereal token info --token=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the token exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		if quiet {
			os.Exit(0)
		}

		name, err := token.Name(nil)
		if err == nil {
			fmt.Printf("Name: %s\n", name)
		}

		symbol, err := token.Symbol(nil)
		if err == nil {
			fmt.Printf("Symbol: %s\n", symbol)
		}

		decimals, err := token.Decimals(nil)
		if err == nil {
			fmt.Printf("Decimals: %d\n", decimals)
		}

		totalSupply, err := token.TotalSupply(nil)
		if err == nil {
			fmt.Printf("Total supply: %s\n", util.TokenValueToString(totalSupply, decimals, true))
		}
	},
}

func init() {
	tokenFlags(tokenInfoCmd)
	tokenCmd.AddCommand(tokenInfoCmd)
}
