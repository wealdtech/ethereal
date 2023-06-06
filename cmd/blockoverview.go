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
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var blockOverviewBlocks int64

// blockOverviewCmd represents the block overview command.
var blockOverviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Obtain overview about recent blocks",
	Long: `Obtain overview about the latest blocks.  For example:

    ethereal block overview

In quiet mode this will return 0 if the blocks exist, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		var blockNumber *big.Int
		var lastBlockTime *time.Time
		if verbose {
			fmt.Printf("Block\t Gas used/Gas limit\tBlock time\t\tGap\tCoinbase\n")
		}
		for i := blockOverviewBlocks; i > 0; i-- {
			ctx, cancel := localContext()
			defer cancel()
			block, err := c.Client().BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about latest block")
			blockNumber = big.NewInt(0).Set(block.Number())
			blockTime := time.Unix(int64(block.Time()), 0)

			if !quiet {
				fmt.Printf("%v\t%9d/%9d\t", blockNumber, block.GasUsed(), block.GasLimit())
				fmt.Printf("%s\t", blockTime.Format("06/01/02 15:04:05"))
				if lastBlockTime != nil {
					gap := lastBlockTime.Sub(blockTime)
					fmt.Printf("%v", gap)
				}
				coinbase := block.Coinbase()
				fmt.Printf("\t%s\n", ens.Format(c.Client(), coinbase))
				lastBlockTime = &blockTime
			}
			blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
		}
	},
}

func init() {
	blockCmd.AddCommand(blockOverviewCmd)
	blockOverviewCmd.Flags().Int64Var(&blockOverviewBlocks, "blocks", 5, "Number of blocks to show")

	blockFlags(blockOverviewCmd)
}
