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

// registryImplementerClearCmd represents the registry implementer clear command.
var registryImplementerClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the address of an ERC-1820 interface implementer",
	Long: `Clear the address of an implementer registered with the ERC-1820 registry for a given interface.  For example:

    ethereal registry implementer clear --interface=ERC777Token --address=0x1234...5678

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryImplementerInterface != "", quiet, "--interface is required")

		address, err := c.Resolve(registryImplementerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		registry, err := erc1820.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		opts, err := generateTxOpts(address)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetInterfaceImplementer(opts, registryImplementerInterface, &address, &ens.UnknownAddress)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":           "registry/implementer",
			"command":         "clear",
			"registryaddress": address.Hex(),
		}, true)
	},
}

func init() {
	registryImplementerFlags(registryImplementerClearCmd)
	registryImplementerCmd.AddCommand(registryImplementerClearCmd)
	addTransactionFlags(registryImplementerClearCmd, "passphrase for the address")
}
