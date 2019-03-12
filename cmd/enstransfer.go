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
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var ensTransferNewOwnerStr string

// ensTransferCmd represents the transfer command
var ensTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer an ENS name",
	Long: `Transfer an Ethereum Name Service (ENS) name's ownership to another address.  For example:

    ethereal ens transfer --domain=enstest.eth --newowner=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase="my secret passphrase"

The keystore for the address must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to transfer the name is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")
		cli.Assert(ensTransferNewOwnerStr != "", quiet, "--newowner is required")
		cli.Assert(len(ensDomain) > 10, quiet, "Domain must be at least 7 characters long")
		cli.Assert(len(strings.Split(ensDomain, ".")) == 2, quiet, "Name must not contain . (except for ending in .eth)")

		// Ensure that the name is in a suitable state
		registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(ensDomain))
		cli.ErrCheck(err, quiet, "cannot obtain ENS registrar contract")

		// Obtain the registry contract
		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "cannot obtain ENS registry contract")

		// Fetch the owner of the name
		owner, err := registryContract.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "cannot obtain owner")
		cli.Assert(bytes.Compare(owner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, fmt.Sprintf("owner of %s is not set", ensDomain))
		outputIf(verbose, fmt.Sprintf("Current owner of %s is %s", ensDomain, owner.Hex()))

		// Transfer the deed
		newOwnerAddress, err := ens.Resolve(client, ensTransferNewOwnerStr)
		cli.ErrCheck(err, quiet, fmt.Sprintf("unknown new owner %s", ensTransferNewOwnerStr))
		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")
		domain, err := ens.DomainPart(ensDomain, 1)
		cli.ErrCheck(err, quiet, fmt.Sprintf("failed to parse domain %s", ensDomain))
		tx, err := registrarContract.Transfer(opts, ens.LabelHash(domain), newOwnerAddress)
		cli.ErrCheck(err, quiet, "failed to send transaction")
		if !quiet {
			fmt.Println(tx.Hash().Hex())
		}
		setupLogging()
		log.WithFields(log.Fields{"transactionid": tx.Hash().Hex(),
			"domain":    ensDomain,
			"networkid": chainID,
			"newowner":  newOwnerAddress.Hex()}).Info("ENS transfer")
	},
}

func init() {
	ensCmd.AddCommand(ensTransferCmd)
	ensFlags(ensTransferCmd)
	ensTransferCmd.Flags().StringVar(&ensTransferNewOwnerStr, "newowner", "", "The new owner of the domain")
	addTransactionFlags(ensTransferCmd, "passphrase for the account that owns the domain")
}
