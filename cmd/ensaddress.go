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
	"github.com/spf13/cobra"
)

var ensAddressCoinType uint64

// ensAddressCmd represents the ens address command.
var ensAddressCmd = &cobra.Command{
	Use:     "address",
	Aliases: []string{"addr"},
	Short:   "Manage ENS addresses",
	Long:    `Set and obtain Ethereum Name Service address information`,
}

func init() {
	ensCmd.AddCommand(ensAddressCmd)
}

func ensAddressFlags(cmd *cobra.Command) {
	cmd.Flags().Uint64Var(&ensAddressCoinType, "cointype", 60, "The coin type of the address (default 60 for Ethereum)")
	ensFlags(cmd)
}
