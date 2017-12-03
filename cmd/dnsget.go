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
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"github.com/wealdtech/ethereal/ens"
)

var dnsGetResource string
var dnsGetName string
var dnsGetWire bool

// dnsGetCmd represents the dns get command
var dnsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value for a DNS record",
	Long: `Get a value for a DNS resource record.  For example:

    ethereal dns get --zone=wealdtech.eth --resource=A

In quiet mode this will return 0 if the resource exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(dnsDomain != "", quiet, "--zone is required")
		if !strings.HasSuffix(dnsDomain, ".") {
			dnsDomain = dnsDomain + "."
		}
		dnsDomain = ens.NormaliseDomain(dnsDomain)

		dnsGetName = strings.ToLower(dnsGetName)
		if dnsGetName == "" {
			dnsGetName = dnsDomain
		} else {
			if !strings.HasSuffix(dnsGetName, ".") {
				dnsGetName = dnsGetName + "." + dnsDomain
			}
		}
		ensZone := strings.TrimSuffix(dnsDomain, ".")

		zoneHash := ens.NameHash(ensZone)
		// nameHash := ens.LabelHash(dnsGetName)

		cli.Assert(dnsGetResource != "", quiet, "--resource is required")
		dnsGetResource := strings.ToUpper(dnsGetResource)
		resourceNum, exists := stringToType[dnsGetResource]
		cli.Assert(exists, quiet, fmt.Sprintf("Unknown resource %s", dnsGetResource))
		outputIf(verbose, fmt.Sprintf("Resource record is %s (%d)", dnsGetResource, resourceNum))

		// Obtain the registry contract
		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Obtain resolver for the domain
		resolverAddress, err := ens.Resolver(registryContract, ensZone)
		cli.ErrCheck(err, quiet, fmt.Sprintf("No resolver registered for %s", dnsDomain))
		resolverContract, err := ens.DnsResolverContractByAddress(client, resolverAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", dnsDomain))
		outputIf(verbose, fmt.Sprintf("Resolver contract is at %s", resolverAddress.Hex()))

		data, err := resolverContract.Dns(nil, zoneHash, resourceNum, dnsGetName)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain %s resource %s for %s", dnsGetResource, dnsGetName, dnsDomain))
		cli.Assert(len(data) > 0, quiet, fmt.Sprintf("No value of %s resource %s for %s", dnsGetResource, dnsGetName, dnsDomain))

		if quiet {
			os.Exit(0)
		}

		if dnsGetWire {
			fmt.Println(hex.EncodeToString(data))
		} else {
			// Decode the data resource record(s)
			offset := 0
			var result dns.RR
			for offset < len(data) {
				result, offset, err = dns.UnpackRR(data, offset)
				fmt.Println(result)
			}
		}

	},
}

func init() {
	dnsCmd.AddCommand(dnsGetCmd)
	dnsFlags(dnsGetCmd)
	dnsGetCmd.Flags().StringVar(&dnsGetResource, "resource", "", "The resource (A, NS, CNAME etc.)")
	dnsGetCmd.Flags().StringVar(&dnsGetName, "name", "", "The name for the resource (end with \".\" for fully-qualified domain, otherwise zone will be added)")
	dnsGetCmd.Flags().BoolVar(&dnsGetWire, "wire", false, "Display the output as hex in wire format")
}
