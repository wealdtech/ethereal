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
	ens "github.com/wealdtech/go-ens/v2"
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
		resolver, err := ens.NewResolver(client, ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		bytes, err := resolver.Contenthash()
		cli.ErrCheck(err, quiet, "Failed to obtain content hash for that domain")
		cli.Assert(len(bytes) > 0, quiet, "No content hash for that domain")

		if ensContenthashGetRaw {
			if !quiet {
				fmt.Printf("%x\n", bytes)
			}
			os.Exit(_exit_success)
		}
		outputIf(debug, fmt.Sprintf("data is %x", bytes))

		res, err := contenthashBytesToString(bytes)
		cli.ErrCheck(err, quiet, "Invalid content hash data")

		if !quiet {
			fmt.Printf("%s\n", res)
		}
		os.Exit(_exit_success)
	},
}

func contenthashBytesToString(bytes []byte) (string, error) {
	data, codec, err := multicodec.RemoveCodec(bytes)
	if err != nil {
		return "", err
	}
	codecName, err := multicodec.Name(codec)
	if err != nil {
		return "", err
	}
	id, offset := binary.Uvarint(data)
	if id == 0 {
		return "", fmt.Errorf("unknown CID")
	}
	data, subCodec, err := multicodec.RemoveCodec(data[offset:])
	if err != nil {
		return "", err
	}
	_, err = multicodec.Name(subCodec)
	if err != nil {
		return "", err
	}

	switch codecName {
	case "ipfs-ns":
		mHash := multihash.Multihash(data)
		return fmt.Sprintf("/ipfs/%s", mHash.B58String()), nil
	case "swarm-ns":
		hash, err := multihash.Decode(data)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("/swarm/%x", hash.Digest), nil
	default:
		return "", fmt.Errorf("unknown codec %s", codecName)
	}
}

func init() {
	ensContenthashFlags(ensContenthashGetCmd)
	ensContenthashGetCmd.Flags().BoolVar(&ensContenthashGetRaw, "raw", false, "output raw content hash bytes")
	ensContenthashCmd.AddCommand(ensContenthashGetCmd)
}
