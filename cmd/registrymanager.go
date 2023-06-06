// Copyright © 2019 Weald Technology Trading
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

var registryManagerAddressStr string

// registryManagerCmd represents the registry manager command.
var registryManagerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Manage ERC-1820 registry managers",
	Long:  `Set and obtain ERC-1820 registry manager information`,
}

func init() {
	registryCmd.AddCommand(registryManagerCmd)
}

func registryManagerFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&registryManagerAddressStr, "address", "", "address against which to operate (e.g. wealdtech.eth)")
}
