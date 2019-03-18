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
	"encoding/binary"
	"fmt"
	"os"

	multihash "github.com/multiformats/go-multihash"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
	multicodec "github.com/wealdtech/go-multicodec"
)

var ensContenthashGetRaw bool

// ensContenthashGetCmd represents the content hash get command
var ensContenthashGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Obtain the content hash of an ENS domain",
	Long: `Obtain the content hash of a domain registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens contenthash get --domain=enstest.eth

In quiet mode this will return 0 if the name has a valid content hash, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		// Obtain resolver for the domain
		resolver, err := ens.ResolverContract(client, ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		bytes, err := resolver.Contenthash(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "Failed to obtain content hash for that domain")
		cli.Assert(len(bytes) > 0, quiet, "No content hash for that domain")

		if ensContenthashGetRaw {
			if !quiet {
				fmt.Printf("%x\n", bytes)
			}
			os.Exit(0)
		}
		outputIf(debug, fmt.Sprintf("data is %x", bytes))

		data, codec, err := multicodec.RemoveCodec(bytes)
		cli.ErrCheck(err, quiet, "Invalid codec")
		codecName, err := multicodec.Name(codec)
		cli.ErrCheck(err, quiet, "Unknown codec")
		id, offset := binary.Uvarint(data)
		cli.Assert(id != 0, quiet, "Unknown CID")
		data, subCodec, err := multicodec.RemoveCodec(data[offset:])
		cli.ErrCheck(err, quiet, "Invalid codec")
		_, err = multicodec.Name(subCodec)
		cli.ErrCheck(err, quiet, "Unknown subcodec")

		switch codecName {
		case "ipfs-ns":
			mHash := multihash.Multihash(data)
			if !quiet {
				fmt.Printf("/ipfs/%s\n", mHash.B58String())
			}
		case "swarm-ns":
			hash, err := multihash.Decode(data)
			cli.ErrCheck(err, quiet, "Failed to decode swarm multihash")
			if !quiet {
				fmt.Printf("/swarm/%x\n", hash.Digest)
			}
		default:
			cli.Err(quiet, fmt.Sprintf("Unknown codec %s", codecName))
		}
		os.Exit(0)
	},
}

func init() {
	ensContenthashFlags(ensContenthashGetCmd)
	ensContenthashGetCmd.Flags().BoolVar(&ensContenthashGetRaw, "raw", false, "output raw content hash bytes")
	ensContenthashCmd.AddCommand(ensContenthashGetCmd)
}
