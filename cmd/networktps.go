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
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

var networkTPSBlocks int64

// networkTPSCmd represents the network tps command.
var networkTPSCmd = &cobra.Command{
	Use:   "tps",
	Short: "Obtain transactions-per-second",
	Long: `Obtain information about the number of transactions processed by the network for a given number of blocks.  For example:

    ethereal network tps

In quiet mode this will return 0 if the network is processing transactions, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		lastBlockTime := now
		lastBlockTransactions := 0
		transactions := 0
		duration := float64(0)
		var blockNumber *big.Int
		for i := networkTPSBlocks + 1; i > 0; i-- {
			ctx, cancel := localContext()
			defer cancel()
			block, err := c.Client().BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about block")

			blockTime := time.Unix(int64(block.Time()), 0)

			if blockNumber != nil {
				transactions += lastBlockTransactions
				lastBlockNumber := big.NewInt(0).Set(blockNumber)
				lastBlockNumber = lastBlockNumber.Add(lastBlockNumber, big.NewInt(1))
				blockDuration := lastBlockTime.Sub(blockTime).Seconds()
				duration += blockDuration
				if verbose {
					fmt.Printf("Block %v processed %v transactions in %v seconds\n", lastBlockNumber, lastBlockTransactions, blockDuration)
				}
			}

			blockNumber = big.NewInt(0).Set(block.Number())
			blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))

			lastBlockTime = blockTime
			lastBlockTransactions = len(block.Transactions())
		}

		if quiet {
			if transactions == 0 {
				os.Exit(exitFailure)
			}
			os.Exit(exitSuccess)
		}

		transactionsPerSecond := float64(transactions) / duration
		fmt.Printf("%.2f\n", transactionsPerSecond)
	},
}

func init() {
	networkCmd.AddCommand(networkTPSCmd)
	networkTPSCmd.Flags().Int64Var(&networkTPSBlocks, "blocks", 5, "Number of blocks to use")

	networkFlags(networkTPSCmd)
}
