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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
	erc1820 "github.com/wealdtech/go-erc1820"
)

var registryImplementerSetImplementerStr string

// registryImplementerSetCmd represents the registry implementer set command.
var registryImplementerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the address of an ERC-1820 interface implementer",
	Long: `Set the address of an implementer registered with the ERC-1820 registry for a given interface.  For example:

    ethereal registry implementer set --interface=ERC777Token --address=0x1234...5678 --implementer=0x9abc...def0

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryImplementerInterface != "", quiet, "--interface is required")

		cli.Assert(registryImplementerAddressStr != "", quiet, "--address is required")
		address, err := c.Resolve(registryImplementerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		cli.Assert(registryImplementerSetImplementerStr != "", quiet, "--implementer is required")
		implementer, err := c.Resolve(registryImplementerSetImplementerStr)
		if err != nil {
			if err.Error() == "could not parse address" {
				cli.Err(quiet, "Invalid implementer address; if you are trying to clear an existing entry use \"registry implementer clear\"")
			}
		}
		cli.ErrCheck(err, quiet, "failed to resolve implementer")

		implementerContract, err := erc1820.NewImplementer(c.Client(), &implementer)
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 implementer contract")
		implementsIface, err := implementerContract.ImplementsInterface(registryImplementerInterface, &address)
		cli.ErrCheck(err, quiet, "failed to check if contract implements ERC-1820")
		cli.Assert(implementsIface, quiet, "implementer does not implement that interface for that address")

		registry, err := erc1820.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		managerAddr, err := registry.Manager(&address)
		cli.ErrCheck(err, quiet, "failed to obtain manager")
		if *managerAddr == ens.UnknownAddress {
			managerAddr = &address
		}

		opts, err := generateTxOpts(*managerAddr)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetInterfaceImplementer(opts, registryImplementerInterface, &address, &implementer)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":               "registry/implementer",
			"command":             "set",
			"registryaddress":     address.Hex(),
			"registryimplementer": implementer.Hex(),
		}, true)
	},
}

func init() {
	registryImplementerFlags(registryImplementerSetCmd)
	registryImplementerSetCmd.Flags().StringVar(&registryImplementerSetImplementerStr, "implementer", "", "address that implements the interface")
	registryImplementerCmd.AddCommand(registryImplementerSetCmd)
	addTransactionFlags(registryImplementerSetCmd, "passphrase for the address")
}
