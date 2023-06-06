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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

var contractStorageKey string

// contractStorageCmd represents the contract storage command.
var contractStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Obtain details about a contract's storage",
	Long: `Obtain the value of a contract's storage key.  For example:

   ethereal contract storage --contract=0xd26114cd6EE289AccF82350c8d8487fedB8A0C07 --key=0x01

In quiet mode this will return 0 if the storage contains a non-zero value, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(contractStr != "", quiet, "--contract is required")
		contractAddress, err := c.Resolve(contractStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve contract address %s", contractStr))

		cli.Assert(contractStorageKey != "", quiet, "--key is required")
		hash := common.HexToHash(strings.TrimPrefix(contractStorageKey, "0x"))
		ctx, cancel := localContext()
		defer cancel()
		value, err := c.Client().StorageAt(ctx, contractAddress, hash, nil)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain storage for contract %s", contractStr))

		if quiet {
			for _, b := range value {
				if b != 0 {
					os.Exit(exitSuccess)
				}
			}
			os.Exit(exitFailure)
		}

		// Output the result.
		fmt.Printf("0x%x\n", value)
	},
}

func init() {
	contractCmd.AddCommand(contractStorageCmd)
	contractFlags(contractStorageCmd)
	contractStorageCmd.Flags().StringVar(&contractStorageKey, "key", "", "Storage key")
}
