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
	"os"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

// ensAddressGetCmd represents the address get command
var ensAddressGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the address of an ENS domain",
	Long: `Obtain the address of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens address get --domain=enstest.eth

In quiet mode this will return 0 if the name has an address, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		address, err := ens.Resolve(client, ensDomain)
		cli.ErrCheck(err, quiet, "failure")
		if !quiet {
			fmt.Println(address.Hex())
		}

		if address == ens.UnknownAddress {
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	ensAddressFlags(ensAddressGetCmd)
	ensAddressCmd.AddCommand(ensAddressGetCmd)
}
