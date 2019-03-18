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
	"bytes"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

// ensContenthashClearCmd represents the ens content hash clear command
var ensContenthashClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the content hash of an ENS domain",
	Long: `Clear the content hash of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens contenthash clear --domain=enstest.eth --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to clear the content hash is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the owner of the name
		owner, err := registryContract.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(bytes.Compare(owner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		// Obtain the resolver for this name
		resolver, err := ens.ResolverContract(client, ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")

		signedTx, err := resolver.SetContenthash(opts, ens.NameHash(ensDomain), []byte{})
		cli.ErrCheck(err, quiet, "failed to send transaction")

		logTransaction(signedTx, log.Fields{
			"group":     "ens/contenthash",
			"command":   "clear",
			"ensdomain": ensDomain,
		})

		if !quiet {
			fmt.Printf("%s\n", signedTx.Hash().Hex())
		}
		os.Exit(0)
	},
}

func init() {
	ensContenthashCmd.AddCommand(ensContenthashClearCmd)
	ensContenthashFlags(ensContenthashClearCmd)
	addTransactionFlags(ensContenthashClearCmd, "passphrase for the account that owns the domain")
}
