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
	"sort"

	etherutils "github.com/orinocopay/go-etherutils"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var gasPriceBlocks int64
var gasPriceWei bool

// gasPriceCmd represents the gas price command
var gasPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "Calculate an expected gas price",
	Long: `Calculate an expected gas price for inclusion based on prior blocks.  For example:

    ethereal gas price --blocks=5

The per-block expected inclusion price is based on the average of the gas price of the transactions in the 9th decile when ordering the blocks' transactions by descending gas price.  The overall inclusion price is the average of the gas price of the transactions over all blocks.

In quiet mode this will return 0 if it can calculate a gas price, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(gasPriceBlocks > 0, quiet, "--blocks must be greater than 0")
		// Always fetch the latest block

		totalGasPrice := big.NewInt(0)
		totalTxs := int64(0)

		var blockNumber *big.Int
		for i := gasPriceBlocks; i > 0; i-- {
			ctx, cancel := localContext()
			defer cancel()
			block, err := client.BlockByNumber(ctx, blockNumber)
			cli.ErrCheck(err, quiet, "Failed to obtain information about latest block")
			blockNumber = big.NewInt(0).Set(block.Number())
			txs := block.Transactions()
			if len(txs) > 0 {
				// Order transactions by gas price
				sort.Slice(txs, func(i, j int) bool {
					return txs[i].GasPrice().Cmp(txs[j].GasPrice()) < 0
				})
				// Remove any 0 gas-price transactions
				validTxs := txs[:0]
				for _, tx := range txs {
					if tx.GasPrice().Cmp(zero) != 0 {
						validTxs = append(validTxs, tx)
					}
				}
				if len(validTxs) > 0 {
					// Take the average gas price of the 9th decile of transactions
					blockGasPrice := big.NewInt(0)
					blockTxs := int64(0)
					for _, tx := range validTxs[(len(validTxs)*8)/10 : (len(validTxs)*9)/10+1] {
						blockGasPrice = blockGasPrice.Add(blockGasPrice, tx.GasPrice())
						blockTxs++
					}
					totalGasPrice = totalGasPrice.Add(totalGasPrice, blockGasPrice)
					totalTxs += blockTxs
					blockGasPrice = blockGasPrice.Div(blockGasPrice, big.NewInt(blockTxs))
					outputIf(verbose, fmt.Sprintf("Expected inclusion price for block %v over %d transactions is %s", blockNumber, blockTxs, etherutils.WeiToString(blockGasPrice, true)))
				}
			}

			blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
		}

		// Obtain final value
		avgPrice := totalGasPrice.Div(totalGasPrice, big.NewInt(totalTxs))

		if quiet {
			os.Exit(0)
		}

		if gasPriceWei {
			fmt.Printf("%s\n", avgPrice.String())
		} else {
			fmt.Printf("%s\n", etherutils.WeiToString(avgPrice, true))
		}
	},
}

func init() {
	gasCmd.AddCommand(gasPriceCmd)
	gasPriceCmd.Flags().BoolVar(&gasPriceWei, "wei", false, "Display output in number of Wei")
	gasPriceCmd.Flags().Int64Var(&gasPriceBlocks, "blocks", 5, "Number of blocks to go back to average gas price")
}
