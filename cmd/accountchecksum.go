// Copyright © 2018, 2019 Weald Technology Trading
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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var (
	accountChecksumAddress string
	accountChecksumCheck   bool
)

// accountChecksumCmd represents the account checksum command.
var accountChecksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Generate or verify the checksum for an account",
	Long: `Generate or verify the checksum for a provided account address.  For example:

    ethereal account checksum --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --check

In quiet mode this will return 0 if the provided address is correctly checksummed, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(accountChecksumAddress != "", quiet, "--address is required")
		// We don't use the ususal resolution process as we want to ensure that the address is well-formed.
		if !strings.HasPrefix(accountChecksumAddress, "0x") {
			cli.Err(quiet, "address does not start with 0x")
		}
		if len(accountChecksumAddress) != 42 {
			cli.Err(quiet, "address of incorrect length")
		}
		address := common.HexToAddress(accountChecksumAddress)
		if address == ens.UnknownAddress {
			cli.Err(quiet, "could not parse address")
		}
		checksummedAddress := address.String()

		if accountChecksumCheck || quiet {
			if accountChecksumAddress != checksummedAddress {
				cli.Err(quiet, "checksum is incorrect")
			}
			outputIf(!quiet, "Checksum is correct")
			os.Exit(exitSuccess)
		}
		fmt.Printf("%s\n", checksummedAddress)
		os.Exit(exitSuccess)
	},
}

func init() {
	offlineCmds["account:checksum"] = true
	accountCmd.AddCommand(accountChecksumCmd)
	accountChecksumCmd.Flags().StringVar(&accountChecksumAddress, "address", "", "Address of the account for which to verify the checksum")
	accountChecksumCmd.Flags().BoolVar(&accountChecksumCheck, "check", false, "Check only; do not print the correctly-checksummed address")
}
