// Copyright © 2019 Weald Technology Trading
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
	erc1820 "github.com/wealdtech/go-erc1820"
)

var (
	registryImplementsInterface  string
	registryImplementsAddressStr string
)

// registryImplementsCmd represents the registry implements command.
var registryImplementsCmd = &cobra.Command{
	Use:   "implements",
	Short: "Check if an address implements an interface according to ERC-1820",
	Long: `Obtain the address of an implementer registered with the ERC-1820 registry for a given interface.  For example:

    ethereal registry implements --interface=ERC777Token --address=0x1234...5678

In quiet mode this will return 0 if the address implements the interface, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryImplementsInterface != "", quiet, "--interface is required")

		cli.Assert(registryImplementsAddressStr != "", quiet, "--address is required")
		address, err := c.Resolve(registryImplementsAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve name")

		implementer, err := erc1820.NewImplementer(c.Client(), &address)
		cli.ErrCheck(err, quiet, "failed to obtain contract")

		anyone := common.HexToAddress("00")
		implementsInterface, err := implementer.ImplementsInterface(registryImplementsInterface, &anyone)
		cli.ErrCheck(err, quiet, "failed to obtain implementation status")

		if !quiet {
			if implementsInterface {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

		if implementsInterface {
			os.Exit(exitSuccess)
		}
		os.Exit(exitFailure)
	},
}

func init() {
	registryFlags(registryImplementsCmd)
	registryImplementsCmd.Flags().StringVar(&registryImplementsInterface, "interface", "", "interface against which to operate (e.g. ERC777TokensRecipient)")
	registryImplementsCmd.Flags().StringVar(&registryImplementsAddressStr, "address", "", "address against which to operate (e.g. wealdtech.eth)")

	registryCmd.AddCommand(registryImplementsCmd)
}
