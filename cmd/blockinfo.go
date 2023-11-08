// Copyright © 2017-2023 Weald Technology Trading.
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
	"encoding/json"
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
	"github.com/wealdtech/ethereal/v2/cli"
	string2eth "github.com/wealdtech/go-string2eth"
)

var (
	blockInfoTransactions bool
	blockInfoJSON         bool
	blockInfoNumberRegexp = regexp.MustCompile("^[0-9]+$")
	gWeiToWei             = big.NewInt(1e9)
)

// blockInfoCmd represents the block info command.
var blockInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a block",
	Long: `Obtain information about a block.  For example:

    ethereal block info --block=0xfdf173c82f1e3e393166719ddc580c161b622fa504fa4b2ddd55f174af554fb7

In quiet mode this will return 0 if the block exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli.Assert(blockStr != "", quiet, "--block is required")

		connectionAddress, err := connectionAddress(ctx)
		cli.ErrCheck(err, quiet, "Failed to obtain connection address")

		execClient, err := jsonrpc.New(context.Background(),
			jsonrpc.WithLogLevel(zerolog.Disabled),
			jsonrpc.WithAddress(connectionAddress),
		)
		cli.ErrCheck(err, quiet, "Failed to access client")
		block, err := execClient.(execclient.BlocksProvider).Block(ctx, blockStr)
		cli.ErrCheck(err, quiet, "Failed to access block")

		if quiet {
			os.Exit(exitSuccess)
		}

		var res string
		if blockInfoJSON {
			res = outputBlockInfoJSON(ctx, block)
		} else {
			res = outputBlockInfoText(ctx, block)
		}

		fmt.Println(strings.TrimSuffix(res, "\n"))
	},
}

func outputBlockInfoJSON(_ context.Context, block *spec.Block) string {
	var res []byte
	var err error

	switch block.Fork {
	case spec.ForkBerlin:
		res, err = json.Marshal(block.Berlin)
		cli.ErrCheck(err, quiet, "failed to generate berlin block info")
	case spec.ForkLondon:
		res, err = json.Marshal(block.London)
		cli.ErrCheck(err, quiet, "failed to generate london block info")
	case spec.ForkShanghai:
		res, err = json.Marshal(block.Shanghai)
		cli.ErrCheck(err, quiet, "failed to generate shanghai block info")
	case spec.ForkCancun:
		res, err = json.Marshal(block.Cancun)
		cli.ErrCheck(err, quiet, "failed to generate cancun block info")
	default:
		res = []byte(fmt.Sprintf("Unhandled block fork %v", block.Fork))
	}

	return string(res)
}

func outputBlockInfoText(ctx context.Context, block *spec.Block) string {
	var res string
	var err error

	switch block.Fork {
	case spec.ForkBerlin:
		res, err = outputBerlinText(ctx, block.Berlin)
		cli.ErrCheck(err, quiet, "failed to generate berlin block info")
	case spec.ForkLondon:
		res, err = outputLondonText(ctx, block.London)
		cli.ErrCheck(err, quiet, "failed to generate london block info")
	case spec.ForkShanghai:
		res, err = outputShanghaiText(ctx, block.Shanghai)
		cli.ErrCheck(err, quiet, "failed to generate shanghai block info")
	case spec.ForkCancun:
		res, err = outputCancunText(ctx, block.Cancun)
		cli.ErrCheck(err, quiet, "failed to generate cancun block info")
	default:
		res = fmt.Sprintf("Unhandled block fork %v", block.Fork)
	}

	return res
}

func outputBerlinText(_ context.Context, block *spec.BerlinBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		if block.Difficulty != 0 {
			outputDifficulty(builder, block.Difficulty)
			outputTotalDifficulty(builder, block.TotalDifficulty)
		}
	}
	outputUncles(builder, block.Uncles, verbose)
	outputTransactions(builder, block.Transactions, verbose)

	return builder.String(), nil
}

func outputLondonText(_ context.Context, block *spec.LondonBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputBaseFee(builder, block.BaseFeePerGas)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		if block.Difficulty != 0 {
			outputDifficulty(builder, block.Difficulty)
			outputTotalDifficulty(builder, block.TotalDifficulty)
		}
	}
	outputTransactions(builder, block.Transactions, verbose)

	return builder.String(), nil
}

func outputShanghaiText(_ context.Context, block *spec.ShanghaiBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputBaseFee(builder, block.BaseFeePerGas)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		if block.Difficulty != 0 {
			outputDifficulty(builder, block.Difficulty)
			outputTotalDifficulty(builder, block.TotalDifficulty)
		}
	}
	outputTransactions(builder, block.Transactions, verbose)
	outputWithdrawals(builder, block.Withdrawals, verbose)

	return builder.String(), nil
}

func outputCancunText(_ context.Context, block *spec.CancunBlock) (string, error) {
	builder := new(strings.Builder)
	outputNumber(builder, block.Number)
	outputHash(builder, block.Hash)
	outputTimestamp(builder, block.Timestamp)
	outputBaseFee(builder, block.BaseFeePerGas)
	outputGas(builder, block.GasUsed, block.GasLimit)
	if verbose {
		outputCoinbase(builder, block.Miner)
		outputExtraData(builder, block.ExtraData)
		if block.Difficulty != 0 {
			outputDifficulty(builder, block.Difficulty)
			outputTotalDifficulty(builder, block.TotalDifficulty)
		}
	}
	outputParentBeaconBlockRoot(builder, block.ParentBeaconBlockRoot)
	outputTransactions(builder, block.Transactions, verbose)
	outputWithdrawals(builder, block.Withdrawals, verbose)

	return builder.String(), nil
}

func outputNumber(builder *strings.Builder, number uint32) {
	builder.WriteString(fmt.Sprintf("Number: %d\n", number))
}

func outputHash(builder *strings.Builder, hash types.Hash) {
	builder.WriteString(fmt.Sprintf("Hash: %#x\n", hash))
}

func outputCoinbase(builder *strings.Builder, coinbase types.Address) {
	builder.WriteString(fmt.Sprintf("Fee recipient: %#x\n", coinbase))
}

func outputTimestamp(builder *strings.Builder, timestamp time.Time) {
	builder.WriteString(fmt.Sprintf("Timestamp: %s (%d)\n", timestamp, timestamp.Unix()))
}

func outputBaseFee(builder *strings.Builder, baseFee uint64) {
	builder.WriteString(fmt.Sprintf("Base fee: %s\n", string2eth.WeiToGWeiString(big.NewInt(int64(baseFee)))))
}

func outputGas(builder *strings.Builder, gasUsed uint32, gasLimit uint32) {
	builder.WriteString(fmt.Sprintf("Gas used: %d/%d (%0.2f%%)\n", gasUsed, gasLimit, float64(gasUsed)*100.0/float64(gasLimit)))
}

func outputParentBeaconBlockRoot(builder *strings.Builder, root types.Root) {
	builder.WriteString(fmt.Sprintf("Parent beacon block root: %s\n", root))
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
			builder.WriteString(fmt.Sprintf("Uncles: %d", len(uncles)))
		}
	}
}

func outputTransactions(builder *strings.Builder, transactions []*spec.Transaction, _ bool) {
	builder.WriteString("Transactions: ")
	builder.WriteString(fmt.Sprintf("%d\n", len(transactions)))
}

func outputWithdrawals(builder *strings.Builder, withdrawals []*spec.Withdrawal, verbose bool) {
	builder.WriteString("Withdrawals: ")
	if !verbose {
		builder.WriteString(fmt.Sprintf("%d\n", len(withdrawals)))
	} else {
		builder.WriteString("\n")
		for _, withdrawal := range withdrawals {
			builder.WriteString(fmt.Sprintf("  %s from %d to %#x\n", string2eth.WeiToString(new(big.Int).Mul(big.NewInt(int64(withdrawal.Amount.Uint64())), gWeiToWei), true), withdrawal.ValidatorIndex, withdrawal.Address))
		}
	}
}

func init() {
	blockCmd.AddCommand(blockInfoCmd)
	blockInfoCmd.Flags().BoolVar(&blockInfoTransactions, "transactions", false, "Display hashes of all block transactions")
	blockInfoCmd.Flags().BoolVar(&blockInfoJSON, "json", false, "Display JSON output")
	blockFlags(blockInfoCmd)
}
