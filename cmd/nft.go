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
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/util/contracts"
	ens "github.com/wealdtech/go-ens/v2"
)

var nftStr string

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Manage ERC-721 non-fungible tokens",
	Long:  `Obtain information, balances and transfer non-fungible tokens between addresses.`,
}

type nft struct {
	contract     *contracts.ERC721
	isEnumerable bool
	hasMetadata  bool
	hasMint      bool
}

// nftTokenID parses an iput token ID
func nftTokenID(input string) (*big.Int, error) {
	if input == "" {
		return nil, errors.New("no token ID")
	}
	base := 10
	if strings.HasPrefix(input, "0x") ||
		strings.Contains(input, "A") ||
		strings.Contains(input, "B") ||
		strings.Contains(input, "C") ||
		strings.Contains(input, "D") ||
		strings.Contains(input, "E") ||
		strings.Contains(input, "F") ||
		strings.Contains(input, "a") ||
		strings.Contains(input, "b") ||
		strings.Contains(input, "c") ||
		strings.Contains(input, "d") ||
		strings.Contains(input, "e") ||
		strings.Contains(input, "f") {
		base = 16
	}
	tokenID, success := new(big.Int).SetString(strings.TrimPrefix(input, "0x"), base)
	if !success {
		return nil, errors.New("failed to convert input to ID")
	}
	return tokenID, nil
}

func nftContractAddress(input string) (common.Address, error) {
	return ens.Resolve(client, input)
}

// nftContract obtains the NFT contract given its name or address
func nftContract(input string) (*nft, error) {
	address, err := nftContractAddress(input)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewERC721(address, client)
	if err != nil {
		return nil, err
	}

	// See if the contract is really ERC721
	isERC721, err := contract.SupportsInterface(nil, [4]byte{0x80, 0xac, 0x58, 0xcd})
	if err != nil {
		return nil, err
	}
	if !isERC721 {
		return nil, fmt.Errorf("%s is not an ERC-721 token contract", input)
	}

	// See if the contract is enumerable
	isEnumerable, err := contract.SupportsInterface(nil, [4]byte{0x78, 0x0e, 0x9d, 0x63})
	if err != nil {
		isEnumerable = false
	}

	// See if the contract has metadata
	hasMetadata, err := contract.SupportsInterface(nil, [4]byte{0x5b, 0x5e, 0x13, 0x9f})
	if err != nil {
		hasMetadata = false
	}

	// See if the contract has minting
	hasMint, err := contract.SupportsInterface(nil, [4]byte{0xd0, 0xde, 0xf5, 0x21})
	if err != nil {
		hasMint = false
	}

	return &nft{
		contract:     contract,
		isEnumerable: isEnumerable,
		hasMetadata:  hasMetadata,
		hasMint:      hasMint,
	}, nil
}

func init() {
	RootCmd.AddCommand(nftCmd)
}

func nftFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&nftStr, "token", "", "Name or address of the token contract")
}
