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

// dnsZonehashCmd represents the dns zonehash command.
var dnsZonehashCmd = &cobra.Command{
	Use:   "zonehash",
	Short: "Manage DNS zonehash entries",
	Long:  `Set and obtain DNS zonehash information in Ethereum Name Service`,
}

func init() {
	dnsCmd.AddCommand(dnsZonehashCmd)
}

func dnsZonehashFlags(cmd *cobra.Command) {
	dnsFlags(cmd)
}
