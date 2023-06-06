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
	"time"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var ensExpiryTimestamp bool

// ensExpiryCmd represents the ens expiry command.
var ensExpiryCmd = &cobra.Command{
	Use:   "expiry",
	Short: "Obtain the expiry date of an ENS domain",
	Long: `Obtain the expiry date of a domain registered with the Ethereum Name Service (ENS).  For example:

    ens expiry enstest.eth

In quiet mode this will return 0 if the domain has an expiry date in the future, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		ensDomain, err := ens.NormaliseDomain(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to normalise ENS domain")

		registrar, err := ens.NewBaseRegistrar(c.Client(), ens.Tld(ensDomain))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain ENS registrar contract for %s", ens.Tld(ensDomain)))

		expiryTS, err := registrar.Expiry(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to obtain expiry")

		if expiryTS.Uint64() == uint64(0) {
			// No expiry.
			outputIf(!quiet, "Domain is not registered")
			os.Exit(exitFailure)
		}

		expiry := time.Unix(int64(expiryTS.Uint64()), 0)

		if !quiet {
			if ensExpiryTimestamp {
				fmt.Printf("%v\n", expiryTS)
			} else {
				fmt.Printf("%v\n", expiry)
			}
		}

		if time.Until(expiry) < 0 {
			os.Exit(exitFailure)
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	ensCmd.AddCommand(ensExpiryCmd)
	ensExpiryCmd.Flags().BoolVar(&ensExpiryTimestamp, "timestamp", false, "Output the expiry as a Unix timestamp")
	ensFlags(ensExpiryCmd)
}
