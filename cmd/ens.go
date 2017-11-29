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
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/ens"
)

var ensDomain string

// ensCmd represents the ens command
var ensCmd = &cobra.Command{
	Use:   "ens",
	Short: "Manage ENS",
	Long:  `Set and obtain Ethereum Name Service information held in Ethereum`,
}

// Ensure that a domain is in a suitable state
func inState(domain string, state string) (inState bool) {
	registrarContract, err := ens.RegistrarContract(client, domain)
	if err == nil {
		inState, err = ens.NameInState(registrarContract, client, domain, state)
		if err != nil {
			inState = false
		}
	}
	return
}

func init() {
	RootCmd.AddCommand(ensCmd)
}

func ensFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&ensDomain, "domain", "", "Domain against which to operate (e.g. wealdtech.eth)")
}
