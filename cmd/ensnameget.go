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
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var ensNameGetAddress string

// ensNameGetCmd represents the name get command
var ensNameGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the ENS reverse resolution of an address",
	Long: `Obtain the Ethereum Name Service (ENS) reverse resolution of an address.  For example:

    ethereal ens name get --address=0x217d2707d6CDA43C4807F343a5f5d93a57d86321

In quiet mode this will return 0 if the address has a reverse resolution, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensNameGetAddress != "", quiet, "--address is required")
		address, err := ens.Resolve(client, ensNameGetAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain address for lookup")

		name, err := ens.ReverseResolve(client, &address)
		if err != nil {
			if err.Error() == "No resolution" {
				os.Exit(1)
			} else {
				cli.ErrCheck(err, quiet, "Failed to check reverse resolution")
			}
		}
		fmt.Println(name)
	},
}

func init() {
	ensNameCmd.AddCommand(ensNameGetCmd)
	ensNameFlags(ensNameGetCmd)
	ensNameGetCmd.Flags().StringVar(&ensNameGetAddress, "address", "", "Address for which to obtain reverse resolution")
}
