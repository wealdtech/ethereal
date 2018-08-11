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
	"github.com/wealdtech/ethereal/ens"
)

var ensResolverSetResolverStr string

// ensResolverSetCmd represents the ens resolver set command
var ensResolverSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the resolver of an ENS domain",
	Long: `Set the resolver of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens resolver set --domain=enstest.eth --resolver=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase="my secret passphrase"

If the resolver is not supplied then the public resolver for the network will be used.

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to set the resolver is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "cannot obtain ENS registry contract")

		// Fetch the owner of the name
		owner, err := registryContract.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "cannot obtain owner")
		cli.Assert(bytes.Compare(owner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		// Set the resolver from either command-line or default
		resolverAddress, err := ens.Resolve(client, ensResolverSetResolverStr)
		if err != nil {
			resolverAddress, err = ens.PublicResolver(client)
			cli.ErrCheck(err, quiet, fmt.Sprintf("no public resolver for network id %v", chainID))
		}

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		tx, err := registryContract.SetResolver(opts, ens.NameHash(ensDomain), resolverAddress)
		cli.ErrCheck(err, quiet, "failed to send transaction")
		if !quiet {
			fmt.Println("Transaction ID is", tx.Hash().Hex())
		}
	},
}

func init() {
	ensResolverCmd.AddCommand(ensResolverSetCmd)
	ensResolverFlags(ensResolverSetCmd)
	ensResolverSetCmd.Flags().StringVarP(&ensResolverSetResolverStr, "resolver", "r", "", "The resolver's name or address")
	addTransactionFlags(ensResolverSetCmd, "passphrase for the account that owns the domain")
}
