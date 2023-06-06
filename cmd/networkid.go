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
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

// networkIDCmd represents the network id command.
var networkIDCmd = &cobra.Command{
	Use:   "id",
	Short: "Obtain the ID of the network",
	Long: `Obtain the ID of the network.  For example:

    ethereal network id

In quiet mode this will return 0 if the network ID is obtained, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Cannot obtain network ID when offline")

		ctx, cancel := localContext()
		defer cancel()
		id, err := c.Client().NetworkID(ctx)
		cli.ErrCheck(err, quiet, "Failed to obtain network ID")
		if !quiet {
			fmt.Printf("%v\n", id)
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	networkCmd.AddCommand(networkIDCmd)
	networkFlags(networkIDCmd)
}
