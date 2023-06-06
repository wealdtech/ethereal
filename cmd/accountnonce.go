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
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
)

var accountNonceAddress string

// accountNonceCmd represents the account nonce command.
var accountNonceCmd = &cobra.Command{
	Use:   "nonce",
	Short: "Obtain the current nonce for an account",
	Long: `Obtain the current nonce for an account, taking in to account pending transactions.  For example:

    ethereal account nonce --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the nonce can be obtained, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(accountNonceAddress != "", quiet, "--address is required")
		address, err := c.Resolve(accountNonceAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain address of %s", accountNonceAddress))

		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
		defer cancel()

		nonce, err := c.Client().PendingNonceAt(ctx, address)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain nonce for %s", accountNonceAddress))

		if !quiet {
			fmt.Println(nonce)
		}
	},
}

func init() {
	accountCmd.AddCommand(accountNonceCmd)
	accountNonceCmd.Flags().StringVar(&accountNonceAddress, "address", "", "Address of the account for which to obtain the nonce")
}
