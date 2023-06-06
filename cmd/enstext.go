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

var ensTextKey string

// ensTextCmd represents the ens text command.
var ensTextCmd = &cobra.Command{
	Use:   "text",
	Short: "Manage ENS text entries",
	Long:  `Set and obtain Ethereum Name Service text information`,
}

func init() {
	ensCmd.AddCommand(ensTextCmd)
}

func ensTextFlags(cmd *cobra.Command) {
	ensFlags(cmd)
	cmd.Flags().StringVar(&ensTextKey, "key", "", "The key of the text value")
}
