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
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensReleaseDomains string

// ensReleaseCmd represents the release command.
var ensReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release funds from an auction registrar deed",
	Long: `Release an Ethereum Name Service (ENS) auction registrar deed, returning the funds locked for that auction.  For example:

    ethereal ens release --domain=enstest.eth --passphrase="my secret passphrase"

Multiple domains can be released with a single command (note this still creates one transaction for each name).  For example:

    ethereal ens release --domains=mydomain1.eth&&mydomain2.eth --passphrase="my secret passphrase"

The keystore for the domain(s) owner must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transactions are successfully submitted (and mined if --wait is supplied), 1 if the transactions are not successfully submitted, and 2 if the transactions are successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "" || ensReleaseDomains != "", quiet, "--domain or --domains is required")

		var domains []string
		if ensReleaseDomains != "" {
			domains = strings.Split(ensReleaseDomains, "&&")
		} else {
			domains = make([]string, 1)
			domains[0] = ensDomain
		}

		auctionRegistrarAddress := common.HexToAddress("0x6090A6e47849629b7245Dfa1Ca21D94cd15878Ef")
		auctionRegistrar, err := ens.NewAuctionRegistrarAt(c.Client(), ens.Tld(domains[0]), auctionRegistrarAddress)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS auction registrar contract")

		for _, domain := range domains {
			cli.Assert(len(domain) > 10, quiet, fmt.Sprintf("Domain %s must be at least 7 characters long", domain))
			cli.Assert(len(strings.Split(domain, ".")) == 2, quiet, fmt.Sprintf("Domain %s must not contain . (except for ending in .eth)", domain))

			entry, err := auctionRegistrar.Entry(domain)
			cli.ErrCheck(err, quiet, "Cannot obtain domain details")
			cli.Assert(entry.Deed != ens.UnknownAddress, quiet, fmt.Sprintf("Domain %q has no deed; unknown or already released", domain))

			owner, err := auctionRegistrar.Owner(domain)
			cli.ErrCheck(err, quiet, "Failed to obtain domain owner")

			outputIf(verbose, fmt.Sprintf("Domain %s owner is %s", domain, ens.Format(c.Client(), owner)))

			opts, err := generateTxOpts(owner)
			cli.ErrCheck(err, quiet, "Failed to generate transaction options")
			signedTx, err := auctionRegistrar.Release(opts, domain)
			cli.ErrCheck(err, quiet, "Failed to send transaction")
			_, err = c.NextNonce(context.Background(), owner)
			cli.ErrCheck(err, quiet, "failed to increment nonce")

			handleSubmittedTransaction(signedTx, log.Fields{
				"group":     "ens",
				"command":   "release",
				"ensdomain": domain,
			}, false)
		}
	},
}

func init() {
	ensCmd.AddCommand(ensReleaseCmd)
	ensFlags(ensReleaseCmd)
	ensReleaseCmd.Flags().StringVar(&ensReleaseDomains, "domains", "", "multiple ENS domains to migrate at the same time; separate with \"&&\" e.g. --domains=mydomain1.eth&&mydomain2.eth")
	addTransactionFlags(ensReleaseCmd, "passphrase for the account that owns the domain")
}
