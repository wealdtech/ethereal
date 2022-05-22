// Copyright © 2017-2022 Weald Technology Trading
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
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	execclient "github.com/attestantio/go-execution-client"
	"github.com/attestantio/go-execution-client/jsonrpc"
	"github.com/attestantio/go-execution-client/spec"
	"github.com/attestantio/go-execution-client/types"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
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
		ctx := context.Background()
		cli.Assert(blockStr != "", quiet, "--block is required")

		execClient, err := jsonrpc.New(context.Background(),
			jsonrpc.WithLogLevel(zerolog.Disabled),
			jsonrpc.WithAddress(viper.GetString("connection")),
		)
		cli.ErrCheck(err, quiet, "Failed to access client")
		block, err := execClient.(execclient.BlocksProvider).Block(ctx, blockStr)
		cli.ErrCheck(err, quiet, "Failed to access block")

		if quiet {
			os.Exit(exitSuccess)
		}

		res := strings.Builder{}

		switch block.Fork {
		case spec.ForkBerlin:
			info, err := outputBerlinText(ctx, block.Berlin)
			cli.ErrCheck(err, quiet, "failed to obtain berlin block info")
			res.WriteString(info)
		case spec.ForkLondon:
			info, err := outputLondonText(ctx, block.London)
			cli.ErrCheck(err, quiet, "failed to obtain london block info")
			res.WriteString(info)
		default:
			res.WriteString("Unhandled block fork ")
			res.WriteString(block.Fork.String())

		}
		fmt.Println(res.String())
	},
}

func outputBerlinText(ctx context.Context, block *spec.BerlinBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		outputDifficulty(builder, block.Difficulty)
		outputTotalDifficulty(builder, block.TotalDifficulty)
	}
	outputUncles(builder, block.Uncles, verbose)
	outputTransactions(builder, block.Transactions, verbose)

	return builder.String(), nil
}

func outputLondonText(ctx context.Context, block *spec.LondonBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputBaseFee(builder, block.BaseFeePerGas)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		outputDifficulty(builder, block.Difficulty)
		outputTotalDifficulty(builder, block.TotalDifficulty)
	}
	outputUncles(builder, block.Uncles, verbose)
	outputTimeToMerge(builder, block.TotalDifficulty, block.Difficulty)
	outputTransactions(builder, block.Transactions, verbose)

	return builder.String(), nil
}

func outputNumber(builder *strings.Builder, number uint32) {
	builder.WriteString(fmt.Sprintf("Number: %d\n", number))
}

func outputHash(builder *strings.Builder, hash types.Hash) {
	builder.WriteString(fmt.Sprintf("Hash: %#x\n", hash))
}

func outputCoinbase(builder *strings.Builder, coinbase types.Address) {
	builder.WriteString(fmt.Sprintf("Coinbase: %#x\n", coinbase))
}

func outputTimestamp(builder *strings.Builder, timestamp time.Time) {
	builder.WriteString(fmt.Sprintf("Timestamp: %s (%d)\n", timestamp, timestamp.Unix()))
}

func outputBaseFee(builder *strings.Builder, baseFee uint64) {
	builder.WriteString(fmt.Sprintf("Base fee: %s\n", string2eth.WeiToString(big.NewInt(int64(baseFee)), true)))
}

func outputGas(builder *strings.Builder, gasUsed uint32, gasLimit uint32) {
	builder.WriteString(fmt.Sprintf("Gas used: %d/%d (%0.2f%%)\n", gasUsed, gasLimit, float64(gasUsed)*100.0/float64(gasLimit)))
}

func outputExtraData(builder *strings.Builder, extraData []byte) {
	extraData = bytes.TrimRight(extraData, "\u0000")
	if len(extraData) > 0 {
		if utf8.Valid(extraData) {
			builder.WriteString(fmt.Sprintf("Extra data: %s\n", string(extraData)))
		} else {
			builder.WriteString(fmt.Sprintf("Extra data: %#x\n", extraData))
		}
	}
}

func outputDifficulty(builder *strings.Builder, difficulty uint64) {
	builder.WriteString(fmt.Sprintf("Difficulty: %d\n", difficulty))
}

func outputTotalDifficulty(builder *strings.Builder, totalDifficulty *big.Int) {
	builder.WriteString(fmt.Sprintf("Total difficulty: %s\n", totalDifficulty.String()))
}

func outputTimeToMerge(builder *strings.Builder, totalDifficulty *big.Int, difficulty uint64) {
	if difficulty > 0 {
		ttd := big.NewInt(20000000000000)
		left := new(big.Int).Sub(ttd, totalDifficulty)
		blocksLeft := new(big.Int).Div(left, big.NewInt(int64(difficulty))).Int64()
		timeLeft := blocksLeft * 14
		when := time.Now().Add(time.Duration(timeLeft) * time.Second)
		builder.WriteString(fmt.Sprintf("Approximate merge time: %s\n", when))
	}
}

func outputUncles(builder *strings.Builder, uncles []types.Hash, verbose bool) {
	if verbose {
		if len(uncles) > 0 {
			builder.WriteString("Uncles:\n")
			for _, uncleHash := range uncles {
				builder.WriteString(fmt.Sprintf("  %#x\n", uncleHash))
			}
		}
	} else {
		if len(uncles) > 0 {
			builder.WriteString(fmt.Sprintf("Uncles %d", len(uncles)))
		}
	}
}

func outputTransactions(builder *strings.Builder, transactions []*spec.Transaction, verbose bool) {
	if verbose {
		if len(transactions) > 0 {
			builder.WriteString("TODO: transactions")
		}
	} else {
		builder.WriteString("Transactions: ")
		builder.WriteString(fmt.Sprintf("%d", len(transactions)))
	}
}

func init() {
	blockCmd.AddCommand(blockInfoCmd)
	blockInfoCmd.Flags().BoolVar(&blockInfoTransactions, "transactions", false, "Display hashes of all block transactions")
	blockFlags(blockInfoCmd)
}
