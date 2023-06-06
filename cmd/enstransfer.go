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
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensTransferNewRegistrantStr string

// ensTransferCmd represents the transfer command.
var ensTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer an ENS name",
	Long: `Transfer an Ethereum Name Service (ENS) name's registration to another address.  For example:

    ethereal ens transfer --domain=enstest.eth --newregistrant=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase="my secret passphrase"

The keystore for the address must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")
		cli.Assert(ensTransferNewRegistrantStr != "", quiet, "--newregistrant is required")
		cli.Assert(len(ensDomain) > 10, quiet, "Domain must be at least 7 characters long")
		cli.Assert(len(strings.Split(ensDomain, ".")) == 2, quiet, "Name must not contain . (except for ending in .eth)")

		registrar, err := ens.NewBaseRegistrar(c.Client(), ens.Tld(ensDomain))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain ENS registrar contract for %s", ens.Tld(ensDomain)))

		// Obtain the registrant.
		domain, err := ens.DomainPart(ensDomain, 1)
		cli.ErrCheck(err, quiet, fmt.Sprintf("failed to obtain domain for %s", ensDomain))
		// Work out if this is on the old or new registrar.
		location, err := registrar.RegisteredWith(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to obtain domain location")
		var registrant common.Address
		var auctionRegistrar *ens.AuctionRegistrar
		switch location {
		case "none":
			outputIf(!quiet, "Domain not registered")
			os.Exit(exitFailure)
		case "temporary":
			auctionRegistrar, err = registrar.PriorAuctionContract()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain auction registrar contract for %s", ens.Tld(ensDomain)))
			registrant, err = auctionRegistrar.Owner(domain)
			cli.ErrCheck(err, quiet, "Failed to obtain domain registrant")
		case "permanent":
			registrant, err = registrar.Owner(domain)
			cli.ErrCheck(err, quiet, "Failed to obtain domain registrant")
		default:
			cli.Err(quiet, fmt.Sprintf("Unexpected domain location %s", location))
		}
		cli.Assert(registrant != ens.UnknownAddress, quiet, "Failed to obtain registrant")

		outputIf(verbose, fmt.Sprintf("Current registrant is %s", ens.Format(c.Client(), registrant)))

		// Transfer the registration.
		newRegistrantAddress, err := c.Resolve(ensTransferNewRegistrantStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("unknown new registrant %s", ensTransferNewRegistrantStr))
		opts, err := generateTxOpts(registrant)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		cli.ErrCheck(err, quiet, fmt.Sprintf("failed to parse domain %s", ensDomain))

		var signedTx *types.Transaction
		switch location {
		case "permanent":
			signedTx, err = registrar.SetOwner(opts, domain, newRegistrantAddress)
		case "temporary":
			signedTx, err = auctionRegistrar.SetOwner(opts, domain, newRegistrantAddress)
		}
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":            "ens",
			"command":          "transfer",
			"ensdomain":        ensDomain,
			"ensnewregistrant": newRegistrantAddress.Hex(),
		}, true)
	},
}

func init() {
	ensCmd.AddCommand(ensTransferCmd)
	ensFlags(ensTransferCmd)
	ensTransferCmd.Flags().StringVar(&ensTransferNewRegistrantStr, "newregistrant", "", "The new registrant of the domain")
	addTransactionFlags(ensTransferCmd, "passphrase for the account that owns the domain")
}
