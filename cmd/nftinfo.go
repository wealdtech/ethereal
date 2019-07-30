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
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
)

var nftInfoTokenIDStr string

// nftInfoCmd represents the nft info command
var nftInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about a non-fungible token",
	Long: `Obtain information about a non-fungible token.  For example:

    ethereal nft info --token=0xBd13e53255eF917DA7557db1B7D2d5C38a2EFe24

In quiet mode this will return 0 if the token exists, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(nftStr != "", quiet, "--token is required")
		nft, err := nftContract(nftStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		if nftInfoTokenIDStr == "" {
			if quiet {
				os.Exit(_exit_success)
			}

			if nft.hasMetadata {
				name, err := nft.contract.Name(nil)
				if err == nil {
					fmt.Printf("Name:\t\t%s\n", name)
				}
			}

			if verbose {
				address, err := nftContractAddress(nftStr)
				if err == nil {
					fmt.Printf("Address:\t%s\n", ens.Format(client, address))
				}
			}

			if nft.hasMetadata {
				symbol, err := nft.contract.Symbol(nil)
				if err == nil {
					fmt.Printf("Symbol:\t\t%s\n", symbol)
				}
			}

			if nft.isEnumerable {
				totalSupply, err := nft.contract.TotalSupply(nil)
				if err == nil {
					fmt.Printf("Total supply:\t%v\n", totalSupply)
				}
			}

			if nft.hasMetadata {
				fmt.Printf("Supports the metadata extension\n")
			}

			if nft.hasMint {
				fmt.Printf("Supports the mint extension\n")
			}

			if nft.isEnumerable {
				fmt.Printf("Supports the enumerable extension\n")
			}
		} else {
			cli.Assert(nft.hasMetadata, quiet, "token does not supply metadata")

			err = initIPFSProvider()
			cli.ErrCheck(err, quiet, "Failed to access IPFS provider")

			tokenID, err := nftTokenID(nftInfoTokenIDStr)
			cli.ErrCheck(err, quiet, "Invalid token ID")

			uri, err := nft.contract.TokenURI(nil, tokenID)
			cli.ErrCheck(err, quiet, "Failed to obtain token URI")
			cli.Assert(uri != "", quiet, "no  URI for that token")
			outputIf(verbose, fmt.Sprintf("token URI is %s", uri))

			if strings.Contains(uri, "ipfs://") || strings.Contains(uri, "/ipfs/") {
				uri, err = ipfsProvider.GatewayURL(uri)
				cli.ErrCheck(err, quiet, "Failed to obtain token URI")
			}
			response, err := http.Get(uri)
			cli.ErrCheck(err, quiet, "Failed to obtain token metadata")

			if quiet {
				os.Exit(_exit_success)
			}

			meta, err := ioutil.ReadAll(response.Body)
			cli.ErrCheck(err, quiet, "Failed to read metadata")
			cli.Assert(strings.HasPrefix(string(meta), "{"), quiet, "Metadata unavailable or not JSON")

			fmt.Printf("%s\n", string(meta))
		}
	},
}

func init() {
	nftFlags(nftInfoCmd)
	nftInfoCmd.Flags().StringVar(&nftInfoTokenIDStr, "tokenid", "", "ID of token")
	nftCmd.AddCommand(nftInfoCmd)
}
