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
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var dnsGetWire bool

// dnsGetCmd represents the dns get command.
var dnsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value for a DNS record",
	Long: `Get a value for a DNS resource record.  For example:

    ethereal dns get --domain=wealdtech.eth --name=www --resource=A

In quiet mode this will return 0 if the resource exists, otherwise 1.`,
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

		dnsName = strings.ToLower(dnsName)
		if dnsName == "" {
			dnsName = dnsDomain
		} else {
			if !strings.HasSuffix(dnsName, ".") {
				dnsName = dnsName + "." + dnsDomain
			}
		}
		outputIf(verbose, fmt.Sprintf("DNS name is %s", dnsName))

		// Obtain DNS resolver for the domain.
		resolver, err := ens.NewDNSResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", dnsDomain))

		var data []byte
		// Attempt to fetch record.
		dnsResource := strings.ToUpper(dnsResource)
		resourceNum, exists := stringToType[dnsResource]
		cli.Assert(exists, quiet, fmt.Sprintf("Unknown resource %s", dnsResource))
		outputIf(verbose, fmt.Sprintf("Resource record is %s (%d)", dnsResource, resourceNum))
		data, err = resolver.Record(dnsName, resourceNum)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain %s resource %s for %s", dnsResource, dnsName, dnsDomain))
		cli.Assert(len(data) > 0, quiet, fmt.Sprintf("No value of %s resource %s for %s", dnsResource, dnsName, dnsDomain))

		if quiet {
			os.Exit(exitSuccess)
		}

		if dnsGetWire {
			fmt.Println(hex.EncodeToString(data))
		} else {
			// Decode the data resource record(s).
			offset := 0
			var result dns.RR
			for offset < len(data) {
				result, offset, err = dns.UnpackRR(data, offset)
				if err == nil {
					fmt.Println(result)
				}
			}
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsGetCmd)
	dnsFlags(dnsGetCmd)
	dnsGetCmd.Flags().BoolVar(&dnsGetWire, "wire", false, "Display the output as hex in wire format")
}
