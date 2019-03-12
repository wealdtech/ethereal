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
	"os"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

// ensContenthashGetCmd represents the content hash get command
var ensContenthashGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the content hash of an ENS domain",
	Long: `Obtain the content hash of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens contenthash get --domain=enstest.eth

In quiet mode this will return 0 if the name has a valid content hash, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		// Obtain the registry contract
		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Obtain resolver for the domain
		resolverAddress, err := ens.Resolver(registryContract, ensDomain)
		cli.ErrCheck(err, quiet, fmt.Sprintf("No resolver registered for %s", ensDomain))
		resolverContract, err := ens.ResolverContractByAddress(client, resolverAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain resolver contract for %s", ensDomain))
		outputIf(verbose, fmt.Sprintf("Resolver contract is at %s", resolverAddress.Hex()))

		bytes, err := resolverContract.Contenthash(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "Failed to obtain content hash for that domain")
		cli.Assert(len(bytes) > 0, quiet, "No content hash for that domain")

		contentHash, err := ma.NewMultiaddrBytes(bytes)
		cli.ErrCheck(err, quiet, "Invalid content hash for that domain")
		if quiet {
			os.Exit(0)
		}
		fmt.Printf("%v\n", contentHash)
	},
}

func init() {
	ensContenthashFlags(ensContenthashGetCmd)
	ensContenthashCmd.AddCommand(ensContenthashGetCmd)
}
