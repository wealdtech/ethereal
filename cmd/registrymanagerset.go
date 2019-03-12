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

var registryManagerSetManagerStr string

// registryManagerSetCmd represents the registry manager set command
var registryManagerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the address of an ERC-1820 address manager",
	Long: `Set the manager of an address in the ERC-1820 registry.  For example:

    ethereal registry manager set --address=0x1234...5678 --manager=0x9abc...def0

In quiet mode this will return 0 if  the transaction to set the manager is sent successfully, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(registryManagerAddressStr != "", quiet, "--address is required")
		address, err := ens.Resolve(client, registryManagerAddressStr)
		cli.ErrCheck(err, quiet, "failed to resolve address")

		cli.Assert(registryManagerSetManagerStr != "", quiet, "--manager is required")
		manager, err := ens.Resolve(client, registryManagerSetManagerStr)
		if err != nil {
			if err.Error() == "could not parse address" {
				cli.Err(quiet, "Invalid manager address; if you are trying to clear an existing entry use \"registry manager clear\"")
			}
		}
		cli.ErrCheck(err, quiet, "failed to resolve manager")

		registry, err := erc1820.NewRegistry(client)
		cli.ErrCheck(err, quiet, "failed to obtain ERC-1820 registry")

		existingManager, err := registry.Manager(&address)
		cli.ErrCheck(err, quiet, "failed to obtain existing manager")

		opts, err := generateTxOpts(*existingManager)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetManager(opts, &address, &manager)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		setupLogging()
		log.WithFields(log.Fields{
			"group":         "registry/manager",
			"command":       "set",
			"address":       address.Hex(),
			"manager":       manager.Hex(),
			"networkid":     chainID,
			"gas":           signedTx.Gas(),
			"gasprice":      signedTx.GasPrice().String(),
			"transactionid": signedTx.Hash().Hex(),
		}).Info("success")

		if quiet {
			os.Exit(0)
		}

		fmt.Println(signedTx.Hash().Hex())
		os.Exit(0)
	},
}

func init() {
	registryManagerFlags(registryManagerSetCmd)
	registryManagerSetCmd.Flags().StringVar(&registryManagerSetManagerStr, "manager", "", "manager for the address")
	registryManagerCmd.AddCommand(registryManagerSetCmd)
	addTransactionFlags(registryManagerSetCmd, "passphrase for the address")
}
