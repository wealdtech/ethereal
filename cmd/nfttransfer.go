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
	"encoding/hex"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
)

var nftTransferTokenIDStr string
var nftTransferAmount string
var nftTransferToAddress string
var nftTransferData string

// nftTransferCmd represents the nft transfer command
var nftTransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer a non-fungible token to a given address",
	Long: `Transfer a non-fungible token from one address to another.  For example:

    ethereal nft transfer --token=omg --tokenid=TODO --from=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --to=0x52f1A3027d3aA514F17E454C93ae1F79b3B12d5d --amount=10 --passphrase=secret

This will return an exit status of 0 if the transaction is successfully submitted (and mined if --wait is supplied), 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(!offline, quiet, "this command cannot be run offline")

		cli.Assert(nftTransferToAddress != "", quiet, "--to is required")
		toAddress, err := ens.Resolve(client, nftTransferToAddress)
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to resolve to address %s", nftTransferToAddress))

		cli.Assert(nftStr != "", quiet, "--token is required")
		nft, err := nftContract(nftStr)
		cli.ErrCheck(err, quiet, "Failed to obtain token contract")

		tokenID, err := nftTokenID(nftTransferTokenIDStr)
		cli.ErrCheck(err, quiet, "Invalid token ID")

		// Turn the data string in to hex
		nftTransferData = strings.TrimPrefix(nftTransferData, "0x")
		if len(nftTransferData)%2 == 1 {
			// Doesn't like odd numbers
			nftTransferData = "0" + nftTransferData
		}
		data, err := hex.DecodeString(nftTransferData)
		cli.ErrCheck(err, quiet, "Failed to parse data")

		owner, err := nft.contract.OwnerOf(nil, tokenID)
		cli.ErrCheck(err, quiet, "Failed to obtain current owner")

		opts, err := generateTxOpts(owner)
		cli.ErrCheck(err, quiet, "Failed to generate transaction options")

		signedTx, err := nft.contract.SafeTransferFrom0(opts, owner, toAddress, tokenID, data)
		cli.ErrCheck(err, quiet, "Failed to create transaction")

		handleSubmittedTransaction(signedTx, log.Fields{
			"group":          "nft",
			"command":        "transfer",
			"token":          nftStr,
			"tokenid":        tokenID,
			"tokenholder":    owner.Hex(),
			"tokenrecipient": toAddress.Hex(),
			"tokendata":      data,
		}, true)
	},
}

func init() {
	nftCmd.AddCommand(nftTransferCmd)
	nftFlags(nftTransferCmd)
	nftTransferCmd.Flags().StringVar(&nftTransferTokenIDStr, "tokenid", "", "ID of the token")
	nftTransferCmd.Flags().StringVar(&nftTransferToAddress, "to", "", "Address to which to transfer the token")
	nftTransferCmd.Flags().StringVar(&nftTransferData, "data", "", "data to send with the transfer (as a hex string)")
	addTransactionFlags(nftTransferCmd, "the address from which to transfer the token")
}
