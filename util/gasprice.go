// Copyright © 2018 Weald Technology Trading
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

package util

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	string2eth "github.com/wealdtech/go-string2eth"
)

// GasPriceForBlocks attempts to calculate a suitable gas price for a transaction given the gas used and the number of blocks
// within which the transaction is desired to be included in.  It does this by looking back over a number of blocks to calculate
// the gas price that would have been required.
// This generally assumes that miners are rational in their selection of transactions to include in blocks, specifically that they
// select the highest gas price transactions available to them.  An obvious exception to this assumption is transactions generated
// by the miners themselves, for example by a mining pool to pay out their miners.  Due to this, blocks that contain self-mined
// transactions are excluded.
// As with any algorithm that uses historic information to calculate future values there are no guarantees that the resultant value
// will provide the desired result; significant changes to gas price between transactions in prior blocks and those in the
// transaction pool can result in an over- or under-estimation of the required gas.
func GasPriceForBlocks(client *ethclient.Client, blocks int64, gasRequired uint64, verbose bool) (*big.Int, error) {
	lowestGasPrice := big.NewInt(0)
	var blockNumber *big.Int

	// We cap blocks at 40 to avoid requesting too many blocks and hammering the server
	if blocks > 40 {
		blocks = 40
	}

	// Fetch the chain ID
	ctx := context.Background()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	for i := blocks; i > 0; i-- {
		ctx := context.Background()
		block, err := client.BlockByNumber(ctx, blockNumber)
		if err != nil {
			return nil, err
		}
		blockNumber = big.NewInt(0).Set(block.Number())
		blockTime := time.Unix(int64(block.Time()), 0)

		// We check the block to see if it contains transactions from the miner, and if so ignore it.  This is because self-mined
		// transactions do not represent market value for gas price
		if BlockHasMinerTransactions(block, chainID) {
			if verbose {
				fmt.Printf("Block %v contains self-mined transactions; ignoring\n", blockNumber)
			}
			blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
			i++
			continue
		}

		txs := block.Transactions()
		if len(txs) > 0 {
			// Order transactions by gas price
			sort.Slice(txs, func(i, j int) bool {
				return txs[i].GasPrice().Cmp(txs[j].GasPrice()) > 0
			})
			// Remove any 0 gas-price transactions
			validTxs := txs[:0]
			for _, tx := range txs {
				if tx.GasPrice().Cmp(zero) != 0 {
					validTxs = append(validTxs, tx)
				}
			}

			gasRemaining := block.GasLimit()
			blockGasPrice := big.NewInt(0)
			for _, tx := range validTxs {
				blockGasPrice = tx.GasPrice()
				if gasRequired > gasRemaining-tx.Gas() {
					break
				}
				gasRemaining -= tx.Gas()
			}
			if verbose {
				fmt.Printf("Inclusion price for block %v (%s) is %s\n", blockNumber, blockTime.Format("06/01/02 15:04:05"), string2eth.WeiToString(blockGasPrice, true))
			}
			if lowestGasPrice.Cmp(zero) == 0 || blockGasPrice.Cmp(lowestGasPrice) < 0 {
				lowestGasPrice = blockGasPrice
			}
		}
		blockNumber = blockNumber.Sub(blockNumber, big.NewInt(1))
	}

	// Increase the gas price by 0.1Gwei to put it above the historic price
	addition, err := string2eth.StringToWei("0.1 gwei")
	if err != nil {
		return nil, err
	}
	lowestGasPrice = lowestGasPrice.Add(lowestGasPrice, addition)

	return lowestGasPrice, nil
}

// BlockHasMinerTransactions returns true if the block contains any transactions signed by the same account that mined the block.
func BlockHasMinerTransactions(block *types.Block, chainID *big.Int) bool {
	signer := types.NewLondonSigner(chainID)
	for _, tx := range block.Transactions() {
		sender, err := types.Sender(signer, tx)
		if err == nil && sender == block.Coinbase() {
			return true
		}
	}
	return false
}
