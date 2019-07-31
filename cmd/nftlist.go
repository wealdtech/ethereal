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
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
)

var nftListOwner string

// nftListCmd represents the nft list command
var nftListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all non-fungible tokens owned by an address",
	Long: `List all non-fungible tokens owned by an address.  For example:

    ethereal nft list --token=0xBd13e53255eF917DA7557db1B7D2d5C38a2EFe24 --owner=0x388Ea662EF2c223eC0B047D41Bf3c0f362142ad6

In quiet mode this will return 0 if the owner holds any tokens, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(nftStr != "", quiet, "--token is required")
		nft, err := nftContract(nftStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		cli.Assert(nft.isEnumerable, quiet, "Token contract is not enumerable")

		cli.Assert(nftListOwner != "", quiet, "--owner is required")
		owner, err := ens.Resolve(client, nftListOwner)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve owner address %s", nftListOwner))

		for index := big.NewInt(0); ; index.Add(index, big.NewInt(1)) {
			id, err := nft.contract.TokenOfOwnerByIndex(nil, owner, index)
			if err != nil {
				if index.Cmp(big.NewInt(0)) == 0 {
					os.Exit(_exit_failure)
				}
				break
			}
			if quiet {
				os.Exit(_exit_success)
			}
			fmt.Printf("%v\n", id)
		}
	},
}

func init() {
	nftFlags(nftListCmd)
	nftListCmd.Flags().StringVar(&nftListOwner, "owner", "", "Owner of tokens")
	nftCmd.AddCommand(nftListCmd)
}
