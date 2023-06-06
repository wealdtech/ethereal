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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
	erc1820 "github.com/wealdtech/go-erc1820"
)

// registryImplementerGetCmd represents the registry implementer get command.
var registryImplementerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the address of an ERC-1820 interface implementer",
	Long: `Obtain the address of an implementer registered with the ERC-1820 registry for a given interface.  For example:

    ethereal registry implementer get --interface=ERC777Token --address=0x1234...5678

In quiet mode this will return 0 if the implementer has an address, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryImplementerInterface != "", quiet, "--interface is required")

		address, err := c.Resolve(registryImplementerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve name")

		registry, err := erc1820.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		implementer, err := registry.InterfaceImplementer(registryImplementerInterface, &address)
		cli.ErrCheck(err, quiet, "failed to obtain implementer")

		if *implementer == ens.UnknownAddress {
			os.Exit(exitFailure)
		}
		if !quiet {
			fmt.Printf("%s\n", ens.Format(c.Client(), *implementer))
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	registryImplementerFlags(registryImplementerGetCmd)
	registryImplementerCmd.AddCommand(registryImplementerGetCmd)
}
