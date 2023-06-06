// Copyright © 2017-2020 Weald Technology Trading
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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var (
	ensSubdomainCreateSubdomain string
	ensSubdomainCreateOwnerStr  string
)

// ensSubdomainCreateCmd represents the ens subdomain create command.
var ensSubdomainCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a subdomain of an ENS domain",
	Long: `Create a subdomain of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens subdomain create --domain=enstest.eth --subdomain=sub --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		cli.Assert(ensSubdomainCreateSubdomain != "", quiet, "--subdomain is required")
		cli.Assert(!strings.Contains(ensSubdomainCreateSubdomain, "."), quiet, "subdomain should not contain the '.' character")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "cannot obtain ENS registry contract")

		// Fetch the controller of the name.
		controller, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "cannot obtain owner")
		cli.Assert(!bytes.Equal(controller.Bytes(), ens.UnknownAddress.Bytes()), quiet, fmt.Sprintf("controller of %s is not set", ensDomain))
		outputIf(debug, fmt.Sprintf("Controller is %s", controller.Hex()))

		// Work out the owner of the subdomain.
		var subdomainOwner common.Address
		if ensSubdomainCreateOwnerStr == "" {
			// Subdomain owner == controller.
			subdomainOwner = controller
		} else {
			subdomainOwner, err = c.Resolve(ensSubdomainCreateOwnerStr)
			cli.ErrCheck(err, quiet, fmt.Sprintf("invalid subdomain name/address %s", ensSubdomainCreateOwnerStr))
		}
		outputIf(debug, fmt.Sprintf("Controller of subdomain will be %s", subdomainOwner.Hex()))

		// Create the subdomain.
		opts, err := generateTxOpts(controller)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		signedTx, err := registry.SetSubdomainOwner(opts, ensDomain, ensSubdomainCreateSubdomain, subdomainOwner)
		cli.ErrCheck(err, quiet, "failed to broadcast transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":             "ens/subdomain",
			"command":           "create",
			"ensdomain":         ensDomain,
			"enssubdomain":      ensSubdomainCreateSubdomain,
			"enssubdomainowner": subdomainOwner.Hex(),
		}, true)
	},
}

func init() {
	ensSubdomainCmd.AddCommand(ensSubdomainCreateCmd)
	ensSubdomainFlags(ensSubdomainCreateCmd)
	ensSubdomainCreateCmd.Flags().StringVar(&ensSubdomainCreateSubdomain, "subdomain", "", "The name of the subdomain")
	ensSubdomainCreateCmd.Flags().StringVar(&ensSubdomainCreateOwnerStr, "owner", "", "The owner of the subdomain (defaults to the owner of the domain)")
	addTransactionFlags(ensSubdomainCreateCmd, "passphrase for the account that owns the domain")
}
