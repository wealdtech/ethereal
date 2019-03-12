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
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var ensOwnerSetOwnerStr string

// ensOwnerSetCmd represents the ens owner set command
var ensOwnerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the owner of an ENS domain",
	Long: `Set the owner of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens owner set --domain=enstest.eth --owner=0x1234...5678 --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to set the owner is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "cannot obtain ENS registry contract")

		// Fetch the owner of the name
		owner, err := registryContract.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "cannot obtain owner")
		cli.Assert(bytes.Compare(owner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		cli.Assert(ensOwnerSetOwnerStr != "", quiet, "--owner is required")
		newOwnerAddress, err := ens.Resolve(client, ensOwnerSetOwnerStr)
		cli.ErrCheck(err, quiet, "Failed to obtain new owner address")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		tx, err := registryContract.SetOwner(opts, ens.NameHash(ensDomain), newOwnerAddress)
		cli.ErrCheck(err, quiet, "failed to send transaction")
		if !quiet {
			fmt.Println("Transaction ID is", tx.Hash().Hex())
		}
	},
}

func init() {
	initAliases(ensOwnerSetCmd)
	ensOwnerCmd.AddCommand(ensOwnerSetCmd)
	ensOwnerFlags(ensOwnerSetCmd)
	ensOwnerSetCmd.Flags().StringVar(&ensOwnerSetOwnerStr, "owner", "", "The owner's name or address")
	addTransactionFlags(ensOwnerSetCmd, "passphrase for the account that owns the domain")
}
