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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
	provider "github.com/wealdtech/go-ipfs-provider"
)

var nftMintName string
var nftMintImage string
var nftMintMetadata string
var nftMintFromAddress string
var nftMintToAddress string

// nftMintCmd represents the nft mint command
var nftMintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint a non-fungible token token to a given address",
	Long: `Mint a non-fungible token, creating its metadata, and gifting it to its initial owner.  For example:

	ethereal nft mint --token=0xc5cD09a414d6999A6Fe96c6b909900EB5227e019 --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --name="My collectible #12345" --image="tokenimage.svg" --metadata='{"value":10000}' --passphrase=secret

This will return an exit status of 0 if the mint transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.

If the token is minted successfully and --wait is supplied this will also print out the ID of the created token in addition to the hash of the minting transaction.

Note that this requires the token contract to have a method mint(address,string); if not then this transaction will fail.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(nftMintFromAddress != "", quiet, "--from is required")
		fromAddress, err := ens.Resolve(client, nftMintFromAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve from address %s", nftMintFromAddress))

		var toAddress common.Address
		if nftMintToAddress == "" {
			toAddress = fromAddress
		} else {
			toAddress, err = ens.Resolve(client, nftMintToAddress)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", nftMintToAddress))
		}

		cli.Assert(nftStr != "", quiet, "--token is required")
		token, err := nftContract(nftStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		cli.Assert(nftMintName != "", quiet, "--name is required")

		metadata := make(map[string]interface{})
		if nftMintMetadata != "" {
			err = json.Unmarshal([]byte(nftMintMetadata), &metadata)
			cli.ErrCheck(err, quiet, "Metadata is not valid JSON")
		}
		metadata["name"] = nftMintName

		err = initIPFSProvider()
		cli.ErrCheck(err, quiet, "Failed to access IPFS provider")

		// Upload an image to IPFS if required
		if nftMintImage != "" && !strings.HasPrefix(nftMintImage, "/ipfs/") {
			opts := &provider.ContentOpts{}
			if strings.HasSuffix(nftMintImage, ".svg") {
				// SVG files as content hashes don't work as images browsers can't
				// tell the difference between simple XML and an XML image) so put
				// the file in a directory and it can be referenced as <hash>/file.svg
				opts.StoreInDirectory = true
			}
			file, err := os.Open(nftMintImage)
			cli.ErrCheck(err, quiet, "Failed to find file")
			hash, err := ipfsProvider.PinContent(filepath.Base(nftMintImage), file, opts)
			cli.ErrCheck(err, quiet, "Failed to upload image")
			imageURI := fmt.Sprintf("ipfs://%s", hash)
			if opts.StoreInDirectory {
				imageURI = fmt.Sprintf("%s/%s", imageURI, filepath.Base(nftMintImage))
			}
			metadata["image"] = imageURI
		}

		mdBytes, err := json.Marshal(metadata)
		cli.ErrCheck(err, quiet, "Failed to create metadata")
		outputIf(verbose, fmt.Sprintf("Metadata is %s", string(mdBytes)))

		hash, err := ipfsProvider.PinContent("metadata.json", bytes.NewBuffer(mdBytes), nil)
		cli.ErrCheck(err, quiet, "Failed to upload metadata")
		outputIf(verbose, fmt.Sprintf("Token metadata IPFS hash is %v", hash))

		opts, err := generateTxOpts(fromAddress)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")

		uri := fmt.Sprintf("ipfs://%v", hash)
		signedTx, err := token.contract.Mint(opts, toAddress, uri)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		if offline {
			if !quiet {
				buf := new(bytes.Buffer)
				signedTx.EncodeRLP(buf)
				fmt.Printf("0x%s\n", hex.EncodeToString(buf.Bytes()))
			}
			os.Exit(_exit_success)
		}

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":     "nft",
			"command":   "mint",
			"token":     nftStr,
			"recipient": toAddress.Hex(),
			"metadata":  string(mdBytes),
			"uri":       uri,
		}, false)

		// Fetch the token ID
		if viper.GetBool("wait") {
			ctx, cancel := localContext()
			defer cancel()
			receipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
			cli.ErrCheck(err, quiet, "Failed to obtain transaction receipt")
			for _, log := range receipt.Logs {
				// Check if this is a transfer event, if so the token ID is in the fourth topic
				if bytes.Compare([]byte{0xdd, 0xf2, 0x52, 0xad, 0x1b, 0xe2, 0xc8, 0x9b, 0x69, 0xc2, 0xb0, 0x68, 0xfc, 0x37, 0x8d, 0xaa, 0x95, 0x2b, 0xa7, 0xf1, 0x63, 0xc4, 0xa1, 0x16, 0x28, 0xf5, 0x5a, 0x4d, 0xf5, 0x23, 0xb3, 0xef}, log.Topics[0].Bytes()) == 0 {
					id := new(big.Int).SetBytes(log.Topics[3].Bytes())
					outputIf(!quiet, fmt.Sprintf("%v", id))
					os.Exit(_exit_success)
				}
			}
			os.Exit(_exit_failure)
		}
	},
}

func init() {
	nftCmd.AddCommand(nftMintCmd)
	nftFlags(nftMintCmd)
	nftMintCmd.Flags().StringVar(&nftMintName, "name", "", "Name for the token")
	nftMintCmd.Flags().StringVar(&nftMintImage, "image", "", "Image for the token")
	nftMintCmd.Flags().StringVar(&nftMintMetadata, "metadata", "", "JSON metadata for token")
	nftMintCmd.Flags().StringVar(&nftMintFromAddress, "from", "", "Address from which to mint the token")
	nftMintCmd.Flags().StringVar(&nftMintToAddress, "to", "", "Address which will own the token (defaults to the minter)")
	addTransactionFlags(nftMintCmd, "the address from which to mint tokens")
}
