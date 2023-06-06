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
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var dnsZonehashSetZonehashStr string

// dnsZonehashSetCmd represents the dns zonehash set command.
var dnsZonehashSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the zone hash of a DNS domain held in ENS",
	Long: `Set the zone hash of a DNS domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal dns zone set --domain=enstest.eth --zonehash=/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162 --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(dnsDomain != "", quiet, "--domain is required")
		if !strings.HasSuffix(dnsDomain, ".") {
			dnsDomain += "."
		}
		dnsDomain, err := ens.NormaliseDomain(dnsDomain)
		cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")
		outputIf(verbose, fmt.Sprintf("DNS domain is %s", dnsDomain))
		ensDomain := strings.TrimSuffix(dnsDomain, ".")
		outputIf(verbose, fmt.Sprintf("ENS domain is %s", ensDomain))
		domainHash, err := ens.NameHash(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to obtain name hash of ENS domain")
		outputIf(verbose, fmt.Sprintf("ENS domain hash is 0x%x", domainHash))

		cli.Assert(dnsZonehashSetZonehashStr != "", quiet, "--zonehash is required; if you are trying to clear an existing zone use \"dns zonehash clear\"")
		data, err := ens.StringToContenthash(dnsZonehashSetZonehashStr)
		cli.ErrCheck(err, quiet, "Invalid zone")
		outputIf(verbose, fmt.Sprintf("Zonehash is %#x", data))

		// Obtain the registry contract.
		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Obtain owner for the domain.
		domainOwner, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")

		cli.Assert(!bytes.Equal(domainOwner.Bytes(), ens.UnknownAddress.Bytes()), quiet, "Owner is not set")
		outputIf(verbose, fmt.Sprintf("Domain owner is %s", ens.Format(c.Client(), domainOwner)))

		// Obtain DNS resolver for the domain.
		resolver, err := ens.NewDNSResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", dnsDomain))

		opts, err := generateTxOpts(domainOwner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")

		signedTx, err := resolver.SetZonehash(opts, data)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":     "dns/zone",
			"command":   "set",
			"ensdomain": ensDomain,
			"zonehash":  dnsZonehashSetZonehashStr,
		}, true)
	},
}

func init() {
	dnsZonehashCmd.AddCommand(dnsZonehashSetCmd)
	dnsZonehashFlags(dnsZonehashSetCmd)
	dnsZonehashSetCmd.Flags().StringVar(&dnsZonehashSetZonehashStr, "zonehash", "", "The path to the zonefile e.g. /ipfs/QmdTEBPdNxJFFsH1wRE3YeWHREWDiSex8xhgTnqknyxWgu")
	addTransactionFlags(dnsZonehashSetCmd, "passphrase for the account that owns the domain")
}
