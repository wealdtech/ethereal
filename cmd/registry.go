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

// registryCmd represents the registry command.
var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage ERC-1820 registry",
	Long:  `Set and obtain ERC-1820 registry information held in Ethereum`,
}

func init() {
	RootCmd.AddCommand(registryCmd)
}

func registryFlags(_ *cobra.Command) {
}
