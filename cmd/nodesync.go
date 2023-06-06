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
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

// nodeSyncCmd represents the node sync command.
var nodeSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Obtain sync information",
	Long: `Obtain information about the synchronisation state of the node.  For example:

    ethereal node sync

In quiet mode this will return 0 if the node is synchronised, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := localContext()
		defer cancel()

		syncProgress, err := c.Client().SyncProgress(ctx)

		cli.ErrCheck(err, quiet, "Failed to obtain node sync status")

		if quiet {
			if syncProgress == nil {
				os.Exit(exitSuccess)
			}
			os.Exit(exitFailure)
		}

		if syncProgress == nil {
			fmt.Printf("Node is synchronised\n")
		} else {
			fmt.Printf("Node is at block %v, syncing to block %v\n", syncProgress.CurrentBlock, syncProgress.HighestBlock)
			outputIf(verbose, fmt.Sprintf("Pulled states is %v, known states is %v", syncProgress.PulledStates, syncProgress.KnownStates))
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	nodeCmd.AddCommand(nodeSyncCmd)

	nodeFlags(nodeSyncCmd)
}
