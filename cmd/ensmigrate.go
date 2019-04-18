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
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var ensMigrateDomains string

// ensMigrateCmd represents the migrate command
var ensMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate an ENS domain from the temporary to permanent registrar",
	Long: `Migrate an Ethereum Name Service (ENS) domain to the permanent registrar.  For example:

    ethereal ens migrate --domain=enstest.eth --passphrase="my secret passphrase"

Multiple domains can be migrated with a single command (note this still creates one transaction for each name).  For example:

    ethereal ens migrate --domains=mydomain1.eth&&mydomain2.eth --passphrase="my secret passphrase"

The keystore for the domain(s) owner must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "" || ensMigrateDomains != "", quiet, "--domain or --domains is required")

		var domains []string
		if ensMigrateDomains != "" {
			domains = strings.Split(ensMigrateDomains, "&&")
		} else {
			domains = make([]string, 1)
			domains[0] = ensDomain
		}

		// Obtain the required registrars
		registrar, err := ens.NewBaseRegistrar(client, ens.Tld(domains[0]))
		cli.ErrCheck(err, quiet, "Cannot obtain ENS base registrar contract")
		auctionRegistrar, err := registrar.PriorAuctionContract()
		cli.ErrCheck(err, quiet, "Cannot obtain ENS auction registrar contract")

		for _, domain := range domains {
			cli.Assert(len(domain) > 10, quiet, fmt.Sprintf("Domain %s must be at least 7 characters long", domain))
			cli.Assert(len(strings.Split(domain, ".")) == 2, quiet, fmt.Sprintf("Domain %s must not contain . (except for ending in .eth)", domain))

			location, err := registrar.RegisteredWith(domain)
			cli.ErrCheck(err, quiet, "Cannot obtain ENS registration information")
			switch location {
			case "temporary":
				// Good
			case "permanent":
				cli.Err(quiet, fmt.Sprintf("Domain %s already migrated", domain))
			case "none":
				cli.Err(quiet, fmt.Sprintf("Domain %s not registered", domain))
			}

			name, err := ens.DomainPart(domain, 1)

			// Ensure the domain is in a suitable state to be migrated
			entry, err := auctionRegistrar.Entry(name)
			cli.ErrCheck(err, quiet, "Cannot obtain domain details")
			cli.Assert(entry.State == "Won" || entry.State == "Owned", quiet, fmt.Sprintf("domain not in a suitable state to be transferred; please run \"ethereal ens info --domain=%s\" to obtain more information about the state of the domain", domain))

			owner, err := auctionRegistrar.Owner(name)
			cli.ErrCheck(err, quiet, "Failed to obtain domain owner")

			outputIf(verbose, fmt.Sprintf("Domain %s owner is %s", domain, ens.Format(client, owner)))

			opts, err := generateTxOpts(owner)
			cli.ErrCheck(err, quiet, "Failed to generate transaction options")
			signedTx, err := auctionRegistrar.Migrate(opts, name)
			cli.ErrCheck(err, quiet, "Failed to send transaction")
			nextNonce(owner)

			handleSubmittedTransaction(signedTx, log.Fields{
				"group":     "ens",
				"command":   "migrate",
				"ensdomain": domain,
			})
		}
	},
}

func init() {
	ensCmd.AddCommand(ensMigrateCmd)
	ensFlags(ensMigrateCmd)
	ensMigrateCmd.Flags().StringVar(&ensMigrateDomains, "domains", "", "multiple ENS domains to migrate at the same time; separate with \"&&\" e.g. --domains=mydomain1.eth&&mydomain2.eth")
	addTransactionFlags(ensMigrateCmd, "passphrase for the account that owns the domain")
}
