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
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
	ens "github.com/wealdtech/go-ens/v3"
)

// tokenInfoCmd represents the token info command.
var tokenInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a token",
	Long: `Obtain information about a token.  For example:

    ethereal token info --token=omg

In quiet mode this will return 0 if the token exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(tokenStr != "", quiet, "--token is required")
		token, err := tokenContract(tokenStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		if quiet {
			os.Exit(exitSuccess)
		}

		name, err := token.Name(nil)
		if err == nil {
			fmt.Printf("Name:\t\t%s\n", name)
		}

		if verbose {
			address, err := tokenContractAddress(tokenStr)
			if err == nil {
				fmt.Printf("Address:\t%s\n", ens.Format(c.Client(), address))
			}
		}

		symbol, err := token.Symbol(nil)
		if err == nil {
			fmt.Printf("Symbol:\t\t%s\n", symbol)
		}

		decimals, err := token.Decimals(nil)
		if err == nil {
			fmt.Printf("Decimals:\t%d\n", decimals)
		}

		totalSupply, err := token.TotalSupply(nil)
		if err == nil {
			fmt.Printf("Total supply:\t%s\n", util.TokenValueToString(totalSupply, decimals, true))
		}
	},
}

func init() {
	tokenFlags(tokenInfoCmd)
	tokenCmd.AddCommand(tokenInfoCmd)
}
