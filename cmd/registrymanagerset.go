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
	erc1820 "github.com/wealdtech/go-erc1820"
)

var registryManagerSetManagerStr string

// registryManagerSetCmd represents the registry manager set command.
var registryManagerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the address of an ERC-1820 address manager",
	Long: `Set the manager of an address in the ERC-1820 registry.  For example:

    ethereal registry manager set --address=0x1234...5678 --manager=0x9abc...def0

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryManagerAddressStr != "", quiet, "--address is required")
		address, err := c.Resolve(registryManagerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		cli.Assert(registryManagerSetManagerStr != "", quiet, "--manager is required")
		manager, err := c.Resolve(registryManagerSetManagerStr)
		if err != nil {
			if err.Error() == "could not parse address" {
				cli.Err(quiet, "Invalid manager address; if you are trying to clear an existing entry use \"registry manager clear\"")
			}
		}
		cli.ErrCheck(err, quiet, "failed to resolve manager")

		registry, err := erc1820.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		existingManager, err := registry.Manager(&address)
		cli.ErrCheck(err, quiet, "failed to obtain existing manager")

		opts, err := generateTxOpts(*existingManager)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetManager(opts, &address, &manager)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":           "registry/manager",
			"command":         "set",
			"registryaddress": address.Hex(),
			"registrymanager": manager.Hex(),
		}, true)
	},
}

func init() {
	registryManagerFlags(registryManagerSetCmd)
	registryManagerSetCmd.Flags().StringVar(&registryManagerSetManagerStr, "manager", "", "manager for the address")
	registryManagerCmd.AddCommand(registryManagerSetCmd)
	addTransactionFlags(registryManagerSetCmd, "passphrase for the address")
}
