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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

var networkUsageBlocks int64

// networkUsageCmd represents the network usage command.
var networkUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Obtain usage of the network in terms of % of gas capacity",
	Long: `Obtain information about the percentage of available gas used by the network for a given number of blocks.  For example:

    ethereal network usage

In quiet mode this will return 0 if the network is processing transactions, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		gas := uint64(0)
		gasLimit := uint64(0)
		var blockNumber *big.Int
		for i := networkUsageBlocks; i > 0; i-- {
			ctx, cancel := localContext()
			defer cancel()
			block, err := c.Client().BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about block")

			gasPct := big.NewFloat(0).Quo(big.NewFloat(0).Mul(big.NewFloat(100), big.NewFloat(0).SetInt(big.NewInt(int64(block.GasUsed())))), big.NewFloat(0).SetInt(big.NewInt(int64(block.GasLimit()))))
			if verbose {
				fmt.Printf("Block %v used %s%% of gas limit (%v/%v)\n", block.Number(), gasPct.Text('f', 2), block.GasUsed(), block.GasLimit())
			}

			gas += block.GasUsed()
			gasLimit += block.GasLimit()

			if blockNumber == nil {
				blockNumber = new(big.Int).Set(block.Number())
			}
			blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
		}

		if quiet {
			if gas == 0 {
				os.Exit(exitFailure)
			}
			os.Exit(exitSuccess)
		}

		gasPct := big.NewFloat(0).Quo(big.NewFloat(0).Mul(big.NewFloat(100), big.NewFloat(0).SetInt(big.NewInt(int64(gas)))), big.NewFloat(0).SetInt(big.NewInt(int64(gasLimit))))
		fmt.Printf("%s%%\n", gasPct.Text('f', 2))
	},
}

func init() {
	networkCmd.AddCommand(networkUsageCmd)
	networkUsageCmd.Flags().Int64Var(&networkUsageBlocks, "blocks", 5, "Number of blocks to use")

	networkFlags(networkUsageCmd)
}
