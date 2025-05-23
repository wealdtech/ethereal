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

// ensControllerGetCmd represents the controller get command.
var ensControllerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the controller of an ENS domain",
	Long: `Obtain the controller of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens controller get --domain=enstest.eth

In quiet mode this will return 0 if the name has a controller, otherwise 1.`,

	Run: func(_ *cobra.Command, _ []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registry, err := ens.NewRegistry(c.Client())
		cli.ErrCheck(err, quiet, "failed to obtain registry contract")
		controller, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "failed to obtain controller")

		if !quiet {
			fmt.Printf("%s\n", ens.Format(c.Client(), controller))
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	ensControllerFlags(ensControllerGetCmd)
	ensControllerCmd.AddCommand(ensControllerGetCmd)
}
