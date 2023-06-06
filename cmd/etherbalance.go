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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	etherBalanceAddress string
	etherBalanceBlock   string
	etherBalanceWei     bool
)

// etherBalanceCmd represents the ether balance command.
var etherBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Obtain the balance for an address",
	Long: `Obtain the Ether balance for an address.  For example:

    ethereal ether balance --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the balance is greater than 0, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(etherBalanceAddress != "", quiet, "--address is required")
		address, err := c.Resolve(etherBalanceAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain address")

		var blockNumber *big.Int
		if etherBalanceBlock != "" {
			if blockInfoNumberRegexp.MatchString(etherBalanceBlock) {
				var succeeded bool
				blockNumber, succeeded = big.NewInt(0).SetString(etherBalanceBlock, 10)
				cli.Assert(succeeded, quiet, fmt.Sprintf("Failed to parse block number %s", etherBalanceBlock))
			} else {
				blockHash := common.HexToHash(etherBalanceBlock)
				ctx, cancel := localContext()
				defer cancel()
				block, err := c.Client().BlockByHash(ctx, blockHash)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain block %s", etherBalanceBlock))
				blockNumber = block.Number()
			}
		}

		ctx, cancel := localContext()
		defer cancel()
		balance, err := c.Client().BalanceAt(ctx, address, blockNumber)
		cli.Assert(err == nil || !strings.HasPrefix(err.Error(), "missing trie node"), quiet, "Connection does not have information on that block, please change the connection parameter to point to a full synced node")
		cli.ErrCheck(err, quiet, "Failed to obtain balance")

		if balance.Cmp(big.NewInt(0)) == 0 {
			outputIf(!quiet, "0")
			os.Exit(exitFailure)
		}

		if !quiet {
			if etherBalanceWei {
				fmt.Printf("%s\n", balance.String())
			} else {
				fmt.Printf("%s\n", string2eth.WeiToString(balance, true))
			}
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	etherCmd.AddCommand(etherBalanceCmd)
	etherBalanceCmd.Flags().BoolVar(&etherBalanceWei, "wei", false, "Display output in number of Wei")
	etherBalanceCmd.Flags().StringVar(&etherBalanceAddress, "address", "", "Address to show Ether balance")
	etherBalanceCmd.Flags().StringVar(&etherBalanceBlock, "block", "", "block hash or number at which to show Ether balance (must be run against an archive node)")
}
