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
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
	"github.com/wealdtech/ethereal/util"
)

var dnsSetTtl time.Duration
var dnsSetValue string
var dnsSetNoSoa bool

// dnsSetCmd represents the dns set command
var dnsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value for a DNS record",
	Long: `Set a value for a DNS record.  For example to set the A record for www.wealdtech.eth to 193.62.81.1:

    ethereal dns record set --domain=wealdtech.eth --ttl=3600 --resource=A --name=www --value=193.62.81.1 --passphrase=secret

In quiet mode this will return 0 if the set transaction is successfully sent, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(dnsDomain != "", quiet, "--domain is required")
		if !strings.HasSuffix(dnsDomain, ".") {
			dnsDomain = dnsDomain + "."
		}
		dnsDomain = ens.NormaliseDomain(dnsDomain)
		outputIf(verbose, fmt.Sprintf("DNS domain is %s", dnsDomain))
		ensDomain := strings.TrimSuffix(dnsDomain, ".")
		outputIf(verbose, fmt.Sprintf("ENS domain is %s", ensDomain))
		domainHash := ens.NameHash(ensDomain)

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

		var signedTx *types.Transaction
		data := make([]byte, 16384)
		if dnsZonefile != "" {
			// Zone-based
			offset := 0
			file, err := os.Open(dnsZonefile)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to open zone file %s", dnsZonefile))
			for rec := range dns.ParseZone(file, dnsDomain, "") {
				cli.Assert(rec.Error == nil, quiet, fmt.Sprintf("Failed to parse zone file %s: %v", dnsZonefile, rec.Error))
				offset, err = dns.PackRR(rec.RR, data, offset, nil, false)
			}
			var b bytes.Buffer
			w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
			cli.ErrCheck(err, quiet, "Failed to compress zone file")
			w.Write(data[0:offset])
			w.Close()
			data = b.Bytes()
			fmt.Printf("Data size is %d\n", len(data))

			// Build the transaction
			opts, err := generateTxOpts(domainOwner)
			cli.ErrCheck(err, quiet, "Failed to generate transaction options")
			signedTx, err = resolverContract.SetDnsZone(opts, domainHash, data)
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
					"domain":        dnsDomain,
					"owner":         domainOwner,
					"networkid":     chainID,
					"gas":           signedTx.Gas(),
					"gasprice":      signedTx.GasPrice().String(),
					"transactionid": signedTx.Hash().Hex(),
				}).Info("success")

				if quiet {
					os.Exit(0)
				}
				fmt.Println(signedTx.Hash().Hex())
			}

		} else {
			// Record-based
			dnsName = strings.ToLower(dnsName)
			if dnsName == "" {
				dnsName = dnsDomain
			} else {
				if !strings.HasSuffix(dnsName, ".") {
					dnsName = dnsName + "." + dnsDomain
				}
			}
			outputIf(verbose, fmt.Sprintf("DNS name is %s", dnsName))
			nameHash := util.DnsDomainHash(dnsName)

			cli.Assert(dnsSetTtl != time.Duration(0), quiet, "--ttl is required")

			cli.Assert(dnsResource != "", quiet, "--resource is required")
			dnsResource := strings.ToUpper(dnsResource)
			resourceNum, exists := stringToType[dnsResource]
			cli.Assert(exists, quiet, fmt.Sprintf("Unknown resource %s", dnsResource))
			outputIf(verbose, fmt.Sprintf("Resource record is %s (%d)", dnsResource, resourceNum))

			cli.Assert(dnsSetValue != "", quiet, "--value is required")

			// Create the data resource record(s)
			offset := 0
			values := strings.Split(dnsSetValue, "&&")
			for _, value := range values {
				source := fmt.Sprintf("%s %d %s %s", dnsName, int(dnsSetTtl.Seconds()), dnsResource, value)
				resource, err := dns.NewRR(source)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to generate resource record from source %s", source))
				offset, err = dns.PackRR(resource, data, offset, nil, false)
			}
			data = data[0:offset]

			var soaData []byte
			if !dnsSetNoSoa {
				// Obtain the current SOA
				curSoaData, err := resolverContract.DnsRecord(nil, domainHash, util.DnsDomainHash(dnsDomain), dns.TypeSOA)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain SOA resource for %s", dnsDomain))
				if len(curSoaData) > 0 {
					// We have an SOA so increment the serial
					soaRr, _, err := dns.UnpackRR(curSoaData, 0)
					cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to unpack SOA resource for %s", dnsDomain))
					outputIf(verbose, fmt.Sprintf("Current SOA record is %v", soaRr))
					soaRr.(*dns.SOA).Serial += 1
					outputIf(verbose, fmt.Sprintf("New SOA record is %v", soaRr))
					soaData = make([]byte, 16384)
					offset, err := dns.PackRR(soaRr, soaData, 0, nil, false)
					soaData = soaData[0:offset]
				}
			}

			// Build the transaction
			opts, err := generateTxOpts(domainOwner)
			cli.ErrCheck(err, quiet, "Failed to generate transaction options")
			signedTx, err = resolverContract.SetDnsRecord(opts, domainHash, nameHash, resourceNum, data, soaData)
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
					"resource":      dnsResource,
					"domain":        dnsDomain,
					"name":          dnsName,
					"value":         dnsSetValue,
					"ttl":           dnsSetTtl,
					"owner":         domainOwner,
					"networkid":     chainID,
					"gas":           signedTx.Gas(),
					"gasprice":      signedTx.GasPrice().String(),
					"transactionid": signedTx.Hash().Hex(),
				}).Info("success")

				if quiet {
					os.Exit(0)
				}

				fmt.Println(signedTx.Hash().Hex())
			}
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsSetCmd)
	dnsFlags(dnsSetCmd)
	dnsSetCmd.Flags().DurationVar(&dnsSetTtl, "ttl", time.Duration(0), "The time-to-live for the record")
	dnsSetCmd.Flags().StringVar(&dnsSetValue, "value", "", "The value for the resource (separate multiple items with &&)")
	dnsSetCmd.Flags().BoolVar(&dnsSetNoSoa, "nosoa", false, "Do not update the zone's SOA record")
	addTransactionFlags(dnsSetCmd, "the owner of the domain")
}
