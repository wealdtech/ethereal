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

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

// ensAddressGetCmd represents the address get command.
var ensAddressGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the address of an ENS domain",
	Long: `Obtain the address of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens address get --domain=enstest.eth

In quiet mode this will return 0 if the name has an address, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		resolver, err := ens.NewResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, "failed to obtain resolver")

		bytes, err := resolver.MultiAddress(ensAddressCoinType)
		cli.ErrCheck(err, quiet, "failed to obtain address")
		if len(bytes) == 0 {
			outputIf(verbose, "no address")
			os.Exit(exitFailure)
		}
		if quiet {
			os.Exit(exitSuccess)
		}

		switch ensAddressCoinType {
		case 60:
			address := common.BytesToAddress(bytes)
			fmt.Printf("%s\n", address.Hex())
		default:
			fmt.Printf("%#x\n", bytes)
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	ensAddressFlags(ensAddressGetCmd)
	ensAddressCmd.AddCommand(ensAddressGetCmd)
}
