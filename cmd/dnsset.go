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
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
	ens "github.com/wealdtech/go-ens/v3"
)

var (
	dnsSetTTL    time.Duration
	dnsSetRecord string
	dnsSetNoSoa  bool
)

// dnsSetCmd represents the dns set command.
var dnsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value for a DNS record",
	Long: `Set a value for a DNS record.  For example to set the A record for www.wealdtech.eth to 193.62.81.1:

    ethereal dns set --domain=wealdtech.eth --ttl=3600 --resource=A --name=www --record=193.62.81.1 --passphrase=secret

In This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		var signedTx *types.Transaction
		data := make([]byte, 32768)
		dnsName = strings.ToLower(dnsName)
		if dnsName == "" {
			dnsName = dnsDomain
		} else {
			if !strings.HasSuffix(dnsName, ".") {
				dnsName = dnsName + "." + dnsDomain
			}
		}
		outputIf(verbose, fmt.Sprintf("DNS name is %s", dnsName))
		cli.Assert(dnsSetTTL != time.Duration(0), quiet, "--ttl is required")

		cli.Assert(dnsResource != "", quiet, "--resource is required")
		dnsResource := strings.ToUpper(dnsResource)
		resourceNum, exists := stringToType[dnsResource]
		cli.Assert(exists, quiet, fmt.Sprintf("Unknown resource %s", dnsResource))
		outputIf(verbose, fmt.Sprintf("Resource record is %s (%d)", dnsResource, resourceNum))

		cli.Assert(dnsSetRecord != "", quiet, "--record is required")

		// Create the data resource record(s).
		offset := 0
		values := strings.Split(dnsSetRecord, "&&")
		for _, value := range values {
			source := fmt.Sprintf("%s %d %s %s", dnsName, int(dnsSetTTL.Seconds()), dnsResource, value)
			outputIf(verbose, fmt.Sprintf("Adding record %s", source))
			resource, err := dns.NewRR(source)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to generate resource record from source %s", source))
			offset, err = dns.PackRR(resource, data, offset, nil, false)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to pack resource record %v", resource))
		}
		data = data[0:offset]

		if dnsResource != "SOA" && !dnsSetNoSoa {
			// Obtain the current SOA.
			curSoaData, err := resolver.Record(dnsDomain, dns.TypeSOA)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain SOA resource for %s", dnsDomain))
			if len(curSoaData) > 0 {
				// We have an SOA so increment the serial as per RFC 1912.
				soaRr, _, err := dns.UnpackRR(curSoaData, 0)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to unpack SOA resource for %s", dnsDomain))
				outputIf(verbose, fmt.Sprintf("Current SOA record is %v", soaRr))
				soaRr.(*dns.SOA).Serial = util.IncrementSerial(soaRr.(*dns.SOA).Serial)
				soaRr.(*dns.SOA).Serial++
				outputIf(verbose, fmt.Sprintf("New SOA record is %v", soaRr))
				soaData := make([]byte, 16384)
				offset, err := dns.PackRR(soaRr, soaData, 0, nil, false)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to pack resource record %v", soaRr))
				soaData = soaData[0:offset]
				copy(data[offset:], soaData)
			}
		}
		outputIf(verbose, fmt.Sprintf("DNS data is %x", data))

		// Build the transaction.
		opts, err := generateTxOpts(domainOwner)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")
		signedTx, err = resolver.SetRecords(opts, data)
		cli.ErrCheck(err, quiet, "Failed to create transaction")
		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				cli.ErrCheck(signedTx.EncodeRLP(buf), quiet, "failed to encode transaction")
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
			os.Exit(exitSuccess)
		}

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":       "dns",
			"command":     "set",
			"dnsresource": dnsResource,
			"dnsdomain":   dnsDomain,
			"dnsname":     dnsName,
			"dnsvalue":    dnsSetRecord,
			"dnsttl":      dnsSetTTL,
		}, true)
	},
}

func init() {
	dnsCmd.AddCommand(dnsSetCmd)
	dnsFlags(dnsSetCmd)
	dnsSetCmd.Flags().DurationVar(&dnsSetTTL, "ttl", time.Duration(0), "The time-to-live for the record")
	dnsSetCmd.Flags().StringVar(&dnsSetRecord, "record", "", "The record for the resource (separate multiple items with &&)")
	dnsSetCmd.Flags().BoolVar(&dnsSetNoSoa, "nosoa", false, "Do not update the zone's SOA record")
	addTransactionFlags(dnsSetCmd, "the owner of the domain")
}
