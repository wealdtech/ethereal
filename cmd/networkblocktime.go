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
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var networkBlocktimeBlocks int64

// networkBlocktimeCmd represents the network blocktime command
var networkBlocktimeCmd = &cobra.Command{
	Use:   "blocktime",
	Short: "Obtain the time between recent blocks",
	Long: `Obtain the time between recent blocks.  For example:

    ethereal network blocktime

In quiet mode this will return 0 if the blocks exist, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		var blockNumber *big.Int

		// Fetch current and current-blocks blocks
		var lastBlockTime time.Time
		{
			ctx, cancel := localContext()
			defer cancel()
			block, err := client.BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about latest block")
			blockNumber = big.NewInt(0).Set(block.Number())
			lastBlockTime = time.Unix(block.Time().Int64(), 0)
			outputIf(verbose, fmt.Sprintf("Block %v mined at %v", blockNumber, lastBlockTime))
		}

		var oldBlockTime time.Time
		blockNumber.Sub(blockNumber, big.NewInt(networkBlocktimeBlocks))
		{
			ctx, cancel := localContext()
			defer cancel()
			block, err := client.BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain information about block %v", blockNumber))
			blockNumber = big.NewInt(0).Set(block.Number())
			oldBlockTime = time.Unix(block.Time().Int64(), 0)
			outputIf(verbose, fmt.Sprintf("Block %v mined at %v", blockNumber, oldBlockTime))
		}

		if quiet {
			os.Exit(0)
		}

		gap := lastBlockTime.Sub(oldBlockTime) / time.Duration(networkBlocktimeBlocks)
		fmt.Printf("%v\n", gap)
	},
}

func init() {
	networkCmd.AddCommand(networkBlocktimeCmd)
	networkBlocktimeCmd.Flags().Int64Var(&networkBlocktimeBlocks, "blocks", 5, "Number of blocks over which to calculate blocktime")

	networkFlags(networkBlocktimeCmd)
}
