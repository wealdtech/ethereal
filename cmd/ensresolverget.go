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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

// ensResolverGetCmd represents the resolver get command.
var ensResolverGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the resolver of an ENS domain",
	Long: `Obtain the resolver of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens resolver get --domain=enstest.eth

In quiet mode this will return 0 if the name has a resolver, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "Failed to obtain registry contract")
		resolver, err := registry.ResolverAddress(ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")
		if !quiet {
			fmt.Printf("%s\n", ens.Format(c.Client(), resolver))
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	ensResolverFlags(ensResolverGetCmd)
	ensResolverCmd.AddCommand(ensResolverGetCmd)
}
