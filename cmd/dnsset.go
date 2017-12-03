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
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
)

var dnsSetTtl time.Duration
var dnsSetResource string
var dnsSetName string
var dnsSetValue string

// dnsSetCmd represents the dns set command
var dnsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value for a DNS record",
	Long: `Set a value for a DNS resource record.  For example to set the A record for www.wealdtech.eth to 193.62.81.1:

    ethereal dns set --domain=wealdtech.eth --ttl=3600 --resource=A --name=www --value=193.62.81.1 --passphrase=secret

In quiet mode this will return 0 if the set transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")

		cli.Assert(dnsDomain != "", quiet, "--domain is required")
		if !strings.HasSuffix(dnsDomain, ".") {
			dnsDomain = dnsDomain + "."
		}
		dnsDomain = ens.NormaliseDomain(dnsDomain)
		outputIf(verbose, fmt.Sprintf("DNS domain is %s", dnsDomain))
		ensDomain := strings.TrimSuffix(dnsDomain, ".")
		outputIf(verbose, fmt.Sprintf("ENS domain is %s", ensDomain))

		dnsSetName = strings.ToLower(dnsSetName)
		if dnsSetName == "" {
			dnsSetName = dnsDomain
		} else {
			if !strings.HasSuffix(dnsSetName, ".") {
				dnsSetName = dnsSetName + "." + dnsDomain
			}
		}
		outputIf(verbose, fmt.Sprintf("DNS name is %s", dnsSetName))

		domainHash := ens.NameHash(ensDomain)

		cli.Assert(dnsSetTtl != time.Duration(0), quiet, "--ttl is required")

		cli.Assert(dnsSetResource != "", quiet, "--resource is required")
		dnsSetResource := strings.ToUpper(dnsSetResource)
		resourceNum, exists := stringToType[dnsSetResource]
		cli.Assert(exists, quiet, fmt.Sprintf("Unknown resource %s", dnsSetResource))
		outputIf(verbose, fmt.Sprintf("Resource record is %s (%d)", dnsSetResource, resourceNum))

		cli.Assert(dnsSetName != "", quiet, "--key is required")
		dnsSetName = strings.ToLower(dnsSetName)

		cli.Assert(dnsSetValue != "", quiet, "--value is required")

		// Obtain the registry contract
		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Obtain owner for the domain
		domainOwner, err := registryContract.Owner(nil, domainHash)
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(bytes.Compare(domainOwner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, "Owner is not set")
		outputIf(verbose, fmt.Sprintf("Domain owner is %s", domainOwner.Hex()))

		// Obtain resolver for the domain
		resolverAddress, err := ens.Resolver(registryContract, ensDomain)
		cli.ErrCheck(err, quiet, fmt.Sprintf("No resolver registered for %s", dnsDomain))
		resolverContract, err := ens.DnsResolverContractByAddress(client, resolverAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", dnsDomain))
		outputIf(verbose, fmt.Sprintf("Resolver contract is at %s", resolverAddress.Hex()))

		// Create the data resource record(s)
		data := make([]byte, 16384)
		offset := 0
		values := strings.Split(dnsSetValue, "&&")
		for _, value := range values {
			source := fmt.Sprintf("%s %d %s %s", dnsSetName, int(dnsSetTtl.Seconds()), dnsSetResource, value)
			resource, err := dns.NewRR(source)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to generate resource record from source %s", source))
			offset, err = dns.PackRR(resource, data, offset, nil, false)
		}
		data = data[0:offset]

		// Send the transaction
		opts, err := generateTxOpts(domainOwner)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err := resolverContract.SetDns(opts, domainHash, resourceNum, dnsSetName, data)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
		} else {
			log.WithFields(log.Fields{
				"group":         "dns",
				"command":       "set",
				"resource":      dnsSetResource,
				"domain":        dnsDomain,
				"name":          dnsSetName,
				"value":         dnsSetValue,
				"ttl":           dnsSetTtl,
				"owner":         domainOwner,
				"networkid":     chainID,
				"gas":           signedTx.Gas().String(),
				"gasprice":      signedTx.GasPrice().String(),
				"transactionid": signedTx.Hash().Hex(),
			}).Info("success")

			if quiet {
				os.Exit(0)
			}

			fmt.Println(signedTx.Hash().Hex())
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsSetCmd)
	dnsFlags(dnsSetCmd)
	dnsSetCmd.Flags().DurationVar(&dnsSetTtl, "ttl", time.Duration(0), "The time-to-live for the record")
	dnsSetCmd.Flags().StringVar(&dnsSetResource, "resource", "", "The resource (A, NS, CNAME etc.)")
	dnsSetCmd.Flags().StringVar(&dnsSetName, "name", "", "The name for the resource (end with \".\" for fully-qualified domain, otherwise zone will be added)")
	dnsSetCmd.Flags().StringVar(&dnsSetValue, "value", "", "The value for the resource (separate multiple items with &&)")
	addTransactionFlags(dnsSetCmd, "the owner of the domain")
}
