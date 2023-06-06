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
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/util/contracts"
)

var tokenStr string

// tokenCmd represents the token command.
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage tokens",
	Long:  `Obtain information, balances and transfer tokens between addresses.`,
}

var unknownAddress = common.HexToAddress("00")

func tokenContractAddress(input string) (common.Address, error) {
	// Guess 1 - might be an ENS name or a hex string.
	address, err := c.Resolve(input)
	if (address == unknownAddress || err != nil) && !strings.HasSuffix(input, ".eth") {
		// Guess 2 - try {input}.thetoken.eth.
		address, err = c.Resolve(input + ".thetoken.eth")
		if err != nil {
			// Give up.
			err = fmt.Errorf("unknown token %s", input)
		}
	}
	return address, err
}

// Obtain the token contract given a string.
func tokenContract(input string) (*contracts.ERC20, error) {
	address, err := tokenContractAddress(input)
	if err != nil {
		return nil, err
	}

	return contracts.NewERC20(address, c.Client())
}

func init() {
	RootCmd.AddCommand(tokenCmd)
}

func tokenFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&tokenStr, "token", "", "Name (resolved as <name>.thetoken.eth) or address of the token contract")
}
