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
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	multihash "github.com/multiformats/go-multihash"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
	multicodec "github.com/wealdtech/go-multicodec"
)

var ensContenthashSetHashStr string

// ensContenthashSetCmd represents the ens content hash set command
var ensContenthashSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the content hash of an ENS domain",
	Long: `Set the content hash of a name registered with the Ethereum Name Service (ENS).  For example:

    ethereal ens contenthash set --domain=enstest.eth --multiaddr=/ip4/1.2.3.4 --passphrase="my secret passphrase"

The keystore for the account that owns the name must be local (i.e. listed with 'get accounts list') and unlockable with the supplied passphrase.

In quiet mode this will return 0 if the transaction to set the content hash is sent successfully, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "Offline mode not supported at current with this command")
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		registryContract, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Cannot obtain ENS registry contract")

		// Fetch the owner of the name
		owner, err := registryContract.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "Cannot obtain owner")
		cli.Assert(bytes.Compare(owner.Bytes(), ens.UnknownAddress.Bytes()) != 0, quiet, fmt.Sprintf("owner of %s is not set", ensDomain))

		cli.Assert(ensContenthashSetHashStr != "", quiet, "--content is required")
		// Break apart the content
		hashBits := strings.Split(ensContenthashSetHashStr, "/")
		cli.Assert(len(hashBits) == 3, quiet, "Invalid content string")

		data := make([]byte, 0)
		switch hashBits[1] {
		case "ipfs":
			// Codec
			ipfsNum, err := multicodec.ID("ipfs-ns")
			cli.ErrCheck(err, quiet, "Failed to obtain IPFS codec value")
			buf := make([]byte, binary.MaxVarintLen64)
			size := binary.PutUvarint(buf, ipfsNum)
			data = append(data, buf[0:size]...)
			// CID
			size = binary.PutUvarint(buf, 1)
			data = append(data, buf[0:size]...)
			// Subcodec
			dagNum, err := multicodec.ID("dag-pb")
			cli.ErrCheck(err, quiet, "Failed to obtain IPFS codec value")
			size = binary.PutUvarint(buf, dagNum)
			data = append(data, buf[0:size]...)
			// Hash
			hash, err := multihash.FromB58String(hashBits[2])
			cli.ErrCheck(err, quiet, "Failed to obtain IPFS content hash")
			data = append(data, []byte(hash)...)
		case "swarm":
			// Codec
			ipfsNum, err := multicodec.ID("swarm-ns")
			cli.ErrCheck(err, quiet, "Failed to obtain swarm codec value")
			buf := make([]byte, binary.MaxVarintLen64)
			size := binary.PutUvarint(buf, ipfsNum)
			data = append(data, buf[0:size]...)
			// CID
			size = binary.PutUvarint(buf, 1)
			data = append(data, buf[0:size]...)
			// Subcodec
			dagNum, err := multicodec.ID("dag-pb")
			cli.ErrCheck(err, quiet, "Failed to obtain swarm codec value")
			size = binary.PutUvarint(buf, dagNum)
			data = append(data, buf[0:size]...)
			// Hash
			hashBit, err := hex.DecodeString(hashBits[2])
			cli.ErrCheck(err, quiet, "Failed to decode swarm content hash")
			hash, err := multihash.Encode(hashBit, multihash.KECCAK_256)
			cli.ErrCheck(err, quiet, "Failed to obtain swarm content hash")
			data = append(data, []byte(hash)...)
		default:
			cli.Err(quiet, fmt.Sprintf("Unknown codec %s", hashBits[1]))
		}

		// Obtain the resolver for this name
		resolver, err := ens.ResolverContract(client, ensDomain)
		cli.ErrCheck(err, quiet, "No resolver for that name")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "failed to generate transaction options")

		signedTx, err := resolver.SetContenthash(opts, ens.NameHash(ensDomain), data)
		cli.ErrCheck(err, quiet, "failed to send transaction")

		logTransaction(signedTx, log.Fields{
			"group":       "ens/contenthash",
			"command":     "set",
			"ensdomain":   ensDomain,
			"contenthash": ensContenthashSetHashStr,
		})

		if !quiet {
			fmt.Printf("%s\n", signedTx.Hash().Hex())
		}
		os.Exit(0)
	},
}

func init() {
	ensContenthashCmd.AddCommand(ensContenthashSetCmd)
	ensContenthashFlags(ensContenthashSetCmd)
	ensContenthashSetCmd.Flags().StringVar(&ensContenthashSetHashStr, "hash", "", "The address to set e.g. /ipfs/QmdTEBPdNxJFFsH1wRE3YeWHREWDiSex8xhgTnqknyxWgu")
	addTransactionFlags(ensContenthashSetCmd, "passphrase for the account that owns the domain")
}
