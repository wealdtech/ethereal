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
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var transactionWaitLimit time.Duration

// transactionWaitCmd represents the transaction info command
var transactionWaitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for a transaction to be mined",
	Long: `Wait for a transaction to be mined.  For example:

    ethereal transaction wait --transaction=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --limit=30s

In quiet mode this will return 0 if the transaction is mined before the time limit is reached, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(transactionStr != "", quiet, "--transaction is required")
		txHash := common.HexToHash(transactionStr)

		start := time.Now()

		first := true
		for transactionWaitLimit == 0 || time.Since(start) < transactionWaitLimit {
			if !first {
				time.Sleep(5 * time.Second)
			} else {
				first = false
			}
			outputIf(verbose, "Checking")
			ctx, cancel := localContext()
			defer cancel()
			_, pending, err := client.TransactionByHash(ctx, txHash)
			if err == nil && !pending {
				outputIf(!quiet, "Transaction mined")
				os.Exit(0)
			}
		}
		cli.Err(quiet, "Transaction not mined")
	},
}

func init() {
	transactionCmd.AddCommand(transactionWaitCmd)
	transactionFlags(transactionWaitCmd)
	transactionWaitCmd.Flags().DurationVar(&transactionWaitLimit, "limit", 0, "maximum time to wait before failing (default forever)")
}
