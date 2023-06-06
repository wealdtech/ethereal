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
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

var (
	networkBlocktimeBlocks int64
	networkBlocktimeTime   time.Duration
)

// networkBlocktimeCmd represents the network blocktime command.
var networkBlocktimeCmd = &cobra.Command{
	Use:   "blocktime",
	Short: "Obtain the time between recent blocks",
	Long: `Obtain the time between recent blocks.  For example:

    ethereal network blocktime

Or to find the blocktime over a given time period (in this case the last hour):

    ethereal network blocktime --time=1h

In quiet mode this will return 0 if the blocks exist, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch current block.
		var lastBlockTime time.Time
		var lastBlockNumber *big.Int
		{
			ctx, cancel := localContext()
			defer cancel()
			block, err := c.Client().BlockByNumber(ctx, lastBlockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about latest block")
			lastBlockNumber = new(big.Int).Set(block.Number())
			lastBlockTime = time.Unix(int64(block.Time()), 0)
			outputIf(verbose, fmt.Sprintf("Block %v mined at %v", lastBlockNumber, lastBlockTime))
		}

		var oldBlockTime time.Time
		var oldBlockNumber *big.Int
		if networkBlocktimeTime > time.Duration(0) {
			// Time.
			requiredBlockTime := time.Now().Add(-networkBlocktimeTime)

			// Start off by guessing 15s per block.
			guessBlockNumber := new(big.Int).Sub(lastBlockNumber, big.NewInt(int64(networkBlocktimeTime.Seconds()/15)))

			// Loop until we find a suitable block.
			interval := 14
			checkedBlocks := make(map[uint64]bool)
			checkedBlocks[lastBlockNumber.Uint64()] = true
			oldBlockNumber = lastBlockNumber
			huntBlockNumber := lastBlockNumber
			huntBlockTime := lastBlockTime
			for {
				ctx, cancel := localContext()
				defer cancel()
				block, err := c.Client().BlockByNumber(ctx, guessBlockNumber)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain information about block %v", guessBlockNumber))
				huntBlockNumber = new(big.Int).Set(block.Number())
				huntBlockTime = time.Unix(int64(block.Time()), 0)
				checkedBlocks[huntBlockNumber.Uint64()] = true
				// If this block is next to (or equal to) the previous block we're done.
				blockDiff := new(big.Int).Abs(new(big.Int).Sub(huntBlockNumber, oldBlockNumber))
				if blockDiff.Cmp(big.NewInt(0)) == 0 || blockDiff.Cmp(big.NewInt(1)) == 0 {
					break
				}

				// Guess a new block given our diff.
				delta := huntBlockTime.Sub(requiredBlockTime)
				guessBlockNumber.Sub(guessBlockNumber, big.NewInt(int64(delta.Seconds()/float64(interval))))
				// If we have already seen this block increase the interval.
				if checkedBlocks[guessBlockNumber.Uint64()] {
					interval *= 2
					guessBlockNumber.Add(guessBlockNumber, big.NewInt(int64(delta.Seconds()/float64(interval))))
				}

				oldBlockTime = huntBlockTime
				oldBlockNumber = huntBlockNumber

				// If our next guess is the same block as the last one we're done.
				if guessBlockNumber.Cmp(oldBlockNumber) == 0 {
					break
				}
			}
			outputIf(verbose, fmt.Sprintf("Block %v mined at %v", oldBlockNumber, oldBlockTime))
		} else {
			// Number of blocks.
			oldBlockNumber = new(big.Int).Sub(lastBlockNumber, big.NewInt(networkBlocktimeBlocks))
			{
				ctx, cancel := localContext()
				defer cancel()
				block, err := c.Client().BlockByNumber(ctx, oldBlockNumber)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain information about block %v", oldBlockNumber))
				oldBlockTime = time.Unix(int64(block.Time()), 0)
				outputIf(verbose, fmt.Sprintf("Block %v mined at %v", oldBlockNumber, oldBlockTime))
			}
		}

		if quiet {
			os.Exit(exitSuccess)
		}

		gap := lastBlockTime.Sub(oldBlockTime) / time.Duration(new(big.Int).Sub(lastBlockNumber, oldBlockNumber).Int64())
		fmt.Printf("%v\n", (gap/10000000)*10000000)
	},
}

func init() {
	networkCmd.AddCommand(networkBlocktimeCmd)
	networkBlocktimeCmd.Flags().Int64Var(&networkBlocktimeBlocks, "blocks", 72, "Number of blocks over which to calculate blocktime")
	networkBlocktimeCmd.Flags().DurationVar(&networkBlocktimeTime, "time", 0, "Time over which to calculate blocktime")

	networkFlags(networkBlocktimeCmd)
}
