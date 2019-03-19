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
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var blockInfoTransactions bool

var blockInfoNumberRegexp = regexp.MustCompile("^[0-9]+$")

// blockInfoCmd represents the block info command
var blockInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a block",
	Long: `Obtain information about a block.  For example:

    ethereal block info --block=0xfdf173c82f1e3e393166719ddc580c161b622fa504fa4b2ddd55f174af554fb7

In quiet mode this will return 0 if the block exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(blockStr != "", quiet, "--block is required")
		var block *types.Block
		ctx, cancel := localContext()
		defer cancel()
		if blockStr == "latest" {
			block, err = client.BlockByNumber(ctx, nil)
		} else if blockInfoNumberRegexp.MatchString(blockStr) {
			blockNum, succeeded := big.NewInt(0).SetString(blockStr, 10)
			cli.Assert(succeeded, quiet, fmt.Sprintf("Failed to parse block number %s", blockStr))
			block, err = client.BlockByNumber(ctx, blockNum)
		} else {
			blockHash := common.HexToHash(blockStr)
			block, err = client.BlockByHash(ctx, blockHash)
		}
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain block %s", blockStr))

		if quiet {
			os.Exit(_exit_success)
		}

		fmt.Printf("Number:\t\t\t%v\n", block.Number())
		fmt.Printf("Hash:\t\t\t%v\n", block.Hash().Hex())
		fmt.Printf("Block time:\t\t%v (%v)\n", block.Time(), time.Unix(block.Time().Int64(), 0))
		if verbose {
			coinbase := block.Coinbase()
			fmt.Printf("Mined by:\t\t%s\n", ens.Format(client, &coinbase))
		}
		outputIf(verbose, fmt.Sprintf("Extra:\t\t\t%s", block.Extra()))
		outputIf(verbose, fmt.Sprintf("Difficulty:\t\t%v", block.Difficulty()))
		fmt.Printf("Gas limit:\t\t%v\n", block.GasLimit())
		gasPct := big.NewFloat(0).Quo(big.NewFloat(0).Mul(big.NewFloat(100), big.NewFloat(0).SetInt(big.NewInt(int64(block.GasUsed())))), big.NewFloat(0).SetInt(big.NewInt(int64(block.GasLimit()))))
		fmt.Printf("Gas used:\t\t%v (%s%%)\n", block.GasUsed(), gasPct.Text('f', 2))
		if verbose {
			if len(block.Uncles()) > 0 {
				fmt.Println("Uncles:")
				for i, uncle := range block.Uncles() {
					fmt.Printf("\t%d:\tblock:\t%v (%v)\n", i, uncle.Number, big.NewInt(0).Sub(uncle.Number, block.Number()))
					fmt.Printf("\t\thash:\t0x%x\n", uncle.Hash())
				}
			}
		} else {
			fmt.Printf("Uncles:\t\t\t%v\n", len(block.Uncles()))
		}
		if blockInfoTransactions {
			if block.Transactions().Len() > 0 {
				fmt.Println("Transactions:")
				for i, tx := range block.Transactions() {
					fmt.Printf("\t%4d: %v\n", i, tx.Hash().Hex())
				}
			}
		} else {
			fmt.Printf("Transactions:\t\t%v\n", block.Transactions().Len())
		}
	},
}

func init() {
	blockCmd.AddCommand(blockInfoCmd)
	blockInfoCmd.Flags().BoolVar(&blockInfoTransactions, "transactions", false, "Display hashes of all block transactions")
	blockFlags(blockInfoCmd)
}
