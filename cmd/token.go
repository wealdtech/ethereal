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
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/util"
	"github.com/wealdtech/ethereal/util/contracts"
)

var tokenStr string

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage tokens",
	Long:  `Obtain information, balances and transfer tokens between addresses.`,
}

var unknownAddress = common.HexToAddress("00")

// Obtain the token contract given a string
func tokenContract(input string) (contract *contracts.ERC20, err error) {
	var addr common.Address
	// See if this looks like a hex address
	if strings.HasPrefix(input, "0x") || len(input) == 64 {
		addr = common.HexToAddress(input)
	}
	if addr == unknownAddress {
		// Guess 2 - try {input}.thetoken.eth
		addr, err = ens.Resolve(client, input+".thetoken.eth")
		if err != nil {
			// Give up
			err = fmt.Errorf("Unknown token %s", input)
		}
	}
	contract, err = util.ERC20Contract(client, addr)
	return
}

func init() {
	RootCmd.AddCommand(tokenCmd)
}

func tokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&tokenStr, "token", "t", "", "Name or address of the token")
}
