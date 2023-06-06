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

// ensTextGetCmd represents the text get command.
var ensTextGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the text of an ENS domain",
	Long: `Obtain the text of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens text get --domain=enstest.eth --enskey="My key"

In quiet mode this will return 0 if the key has text, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		// Obtain resolver for the domain.
		resolver, err := ens.NewResolver(c.Client(), ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		value, err := resolver.Text(ensTextKey)
		cli.ErrCheck(err, quiet, "Failed to obtain value for that domain")
		cli.Assert(len(value) > 0, quiet, "No value for that domain")
		if !quiet {
			fmt.Printf("%s\n", value)
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	ensTextFlags(ensTextGetCmd)
	ensTextCmd.AddCommand(ensTextGetCmd)
}
