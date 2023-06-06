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

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensResolverSetResolverStr string

// ensResolverSetCmd represents the ens resolver set command.
var ensResolverSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the resolver of an ENS domain",
	Long: `Set the resolver of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens resolver set --domain=enstest.eth --resolver=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase="my secret passphrase"

If the resolver is not supplied then the public resolver for the network will be used.

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the owner of the name.
		outputIf(debug, fmt.Sprintf("ENS domain is %s", ensDomain))
		owner, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(!bytes.Equal(owner.Bytes(), ens.UnknownAddress.Bytes()), quiet, fmt.Sprintf("owner of %s is not set", ensDomain))
		outputIf(debug, fmt.Sprintf("Owner of %s is %#x", ensDomain, owner))

		// Set the resolver from either command-line or default.
		var resolverAddress common.Address
		if ensResolverSetResolverStr == "" {
			resolverAddress, err = ens.PublicResolverAddress(c.Client())
			cli.ErrCheck(err, quiet, fmt.Sprintf("No public resolver for network id %v", c.ChainID()))
		} else {
			resolverAddress, err = c.Resolve(ensResolverSetResolverStr)
			cli.Assert(!bytes.Equal(resolverAddress.Bytes(), ens.UnknownAddress.Bytes()), quiet, "Invalid resolver; if you are trying to clear an existing resolver use \"ens resolver clear\"")
			cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid name/address %s", ensAddressSetAddressStr))
		}

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := registry.SetResolver(opts, ensDomain, resolverAddress)
		cli.ErrCheck(err, quiet, "Failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":       "ens/resolver",
			"command":     "set",
			"ensdomain":   ensDomain,
			"ensresolver": resolverAddress.Hex(),
		}, true)
	},
}

func init() {
	ensResolverCmd.AddCommand(ensResolverSetCmd)
	ensResolverFlags(ensResolverSetCmd)
	ensResolverSetCmd.Flags().StringVar(&ensResolverSetResolverStr, "resolver", "", "The resolver's name or address")
	addTransactionFlags(ensResolverSetCmd, "passphrase for the account that owns the domain")
}
