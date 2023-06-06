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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

// dnsZonehashGetCmd represents the zonehash get command.
var dnsZonehashGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the zonehash of a DNS domain on ENS",
	Long: `Obtain the zonehash of a DNS domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal dns zonehash get --domain=enstest.eth

In quiet mode this will return 0 if the name has a valid zone hash, otherwise 1.`,

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

		bytes, err := resolver.Zonehash()
		cli.ErrCheck(err, quiet, "Failed to obtain zonehash for that domain")
		cli.Assert(len(bytes) > 0, quiet, "No zonehash for that domain")

		outputIf(debug, fmt.Sprintf("data is %x", bytes))
		res, err := ens.ContenthashToString(bytes)
		cli.ErrCheck(err, quiet, "Invalid content hash data")

		if !quiet {
			fmt.Printf("%s\n", res)
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	dnsZonehashFlags(dnsZonehashGetCmd)
	dnsZonehashCmd.AddCommand(dnsZonehashGetCmd)
}
