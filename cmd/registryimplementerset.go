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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
	erc1820 "github.com/wealdtech/go-erc1820"
)

var registryImplementerSetImplementerStr string

// registryImplementerSetCmd represents the registry implementer set command
var registryImplementerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the address of an ERC-1820 interface implementer",
	Long: `Set the address of an implementer registered with the ERC-1820 registry for a given interface.  For example:

    ethereal registry implementer set --interface=ERC777Token --address=0x1234...5678 --implementer=0x9abc...def0

In quiet mode this will return 0 if  the transaction to set the implementer is sent successfully, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryImplementerInterface != "", quiet, "--interface is required")

		address, err := ens.Resolve(client, registryImplementerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		implementer, err := ens.Resolve(client, registryImplementerSetImplementerStr)
		if err != nil {
			if err.Error() == "could not parse address" {
				cli.Err(quiet, "Invalid implementer address; if you are trying to clear an existing entry use \"registry implementer clear\"")
			}
		}
		cli.ErrCheck(err, quiet, "failed to resolve implementer")

		registry, err := erc1820.NewRegistry(client)
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		opts, err := generateTxOpts(address)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetInterfaceImplementer(opts, registryImplementerInterface, &address, &implementer)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		logTransaction(signedTx, log.Fields{
			"group":               "registry/implementer",
			"command":             "set",
			"registryaddress":     address.Hex(),
			"registryimplementer": implementer.Hex(),
		})

		if !quiet {
			fmt.Printf("%s\n", signedTx.Hash().Hex())
		}
		os.Exit(0)
	},
}

func init() {
	registryImplementerFlags(registryImplementerSetCmd)
	registryImplementerSetCmd.Flags().StringVar(&registryImplementerSetImplementerStr, "implementer", "", "address that implements the interface")
	registryImplementerCmd.AddCommand(registryImplementerSetCmd)
	addTransactionFlags(registryImplementerSetCmd, "passphrase for the address")
}
