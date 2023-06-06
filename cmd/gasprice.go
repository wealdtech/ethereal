// Copyright © 2017-2021 Weald Technology Trading
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
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	gasPriceBlocks int64
	gasPriceWei    bool
	gasPriceLowest bool
	gas            uint64
)

// gasPriceCmd represents the gas price command.
var gasPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "Calculate an expected gas price",
	Long: `Calculate an expected gas price for inclusion based on prior blocks.  For example:

    ethereal gas price --blocks=5

The per-block expected inclusion price is based on the average of the gas price of the transactions in the 9th decile when ordering the blocks' transactions by descending gas price.  The overall inclusion price is the average of the gas price of the transactions over all blocks.

If the optional --gas parameter is supplied the price will be based on a transaction with the given supplied gas rather than the 9th decile transactions.  This should provide a more accurate indication of gas price for transactions that require the supplied gas.

In quiet mode this will return 0 if it can calculate a gas price, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(gasPriceBlocks > 0, quiet, "--blocks must be greater than 0")

		lowestGasPrice := big.NewInt(0)
		var err error

		totalGasPrice := big.NewInt(0)
		totalTxs := int64(0)

		if gas > 0 {
			lowestGasPrice, err = util.GasPriceForBlocks(c.Client(), gasPriceBlocks, gas, verbose)
			cli.ErrCheck(err, quiet, "Failed to obtain gas price")
		} else {
			var blockNumber *big.Int
			for blocks := gasPriceBlocks; blocks > 0; blocks-- {
				ctx, cancel := localContext()
				defer cancel()
				block, err := c.Client().BlockByNumber(ctx, blockNumber)
				cli.ErrCheck(err, quiet, "Failed to obtain information about latest block")
				blockNumber = big.NewInt(0).Set(block.Number())
				blockTime := time.Unix(int64(block.Time()), 0)

				if util.BlockHasMinerTransactions(block, c.ChainID()) {
					outputIf(verbose, fmt.Sprintf("Block %v contains self-mined transactions; ignoring", blockNumber))
					blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
					blocks++
					continue
				}

				txs := block.Transactions()
				if len(txs) > 0 {
					// Order transactions by gas price.
					sort.Slice(txs, func(i, j int) bool {
						return txs[i].GasPrice().Cmp(txs[j].GasPrice()) > 0
					})
					// Remove any 0 gas-price transactions.
					validTxs := txs[:0]
					for _, tx := range txs {
						if tx.GasPrice().Cmp(zero) != 0 {
							validTxs = append(validTxs, tx)
						}
					}
					if len(validTxs) > 0 {
						if gasPriceLowest {
							blockLowestGasPrice := validTxs[len(validTxs)-1].GasPrice()
							outputIf(verbose, fmt.Sprintf("Lowest inclusion price for block %v (%s) is %s", blockNumber, blockTime.Format("06/01/02 15:04:05"), string2eth.WeiToString(blockLowestGasPrice, true)))
							if lowestGasPrice.Cmp(zero) == 0 || blockLowestGasPrice.Cmp(lowestGasPrice) < 0 {
								lowestGasPrice = blockLowestGasPrice
							}
						} else {
							// Take the average gas price of the 9th decile of transactions.
							blockGasPrice := big.NewInt(0)
							blockTxs := int64(0)
							for _, tx := range validTxs[(len(validTxs)*8)/10 : (len(validTxs)*9)/10+1] {
								blockGasPrice = blockGasPrice.Add(blockGasPrice, tx.GasPrice())
								blockTxs++
							}
							totalGasPrice = totalGasPrice.Add(totalGasPrice, blockGasPrice)
							totalTxs += blockTxs
							blockGasPrice = blockGasPrice.Div(blockGasPrice, big.NewInt(blockTxs))
							outputIf(verbose, fmt.Sprintf("Expected inclusion price for block %v (%s) over %d transactions is %s", blockNumber, blockTime.Format("06/01/02 15:04:05"), blockTxs, string2eth.WeiToString(blockGasPrice, true)))
						}
					}
				}

				blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
				if blockNumber.Cmp(zero) < 0 {
					// We've reached the beginning of the chain; stop.
					break
				}
			}
		}

		// Obtain final value.
		var finalGasPrice *big.Int
		if gasPriceLowest || gas > 0 {
			finalGasPrice = lowestGasPrice
		} else {
			if totalTxs == 0 {
				cli.Err(quiet, "No transactions obtained")
			}
			finalGasPrice = totalGasPrice.Div(totalGasPrice, big.NewInt(totalTxs))
		}

		if quiet {
			os.Exit(exitSuccess)
		}

		if gasPriceWei {
			fmt.Printf("%s\n", finalGasPrice.String())
		} else {
			fmt.Printf("%s\n", string2eth.WeiToString(finalGasPrice, true))
		}
	},
}

func init() {
	gasCmd.AddCommand(gasPriceCmd)
	gasPriceCmd.Flags().BoolVar(&gasPriceWei, "wei", false, "Display output in number of Wei")
	gasPriceCmd.Flags().Int64Var(&gasPriceBlocks, "blocks", 5, "Number of blocks to go back to average gas price")
	gasPriceCmd.Flags().Uint64Var(&gas, "gas", 0, "Provide gas price based on the amount of gas used by the transaction")
	gasPriceCmd.Flags().BoolVar(&gasPriceLowest, "lowest", false, "Lowest inclusion price over the blocks")
}
