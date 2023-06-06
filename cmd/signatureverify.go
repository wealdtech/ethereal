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
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
)

var (
	signatureVerifySignature string
	signatureVerifySigner    string
)

// signatureVerifyCmd represents the signature verify command.
var signatureVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature",
	Long: `Verify a presented signature.  For example:

    ethereal data verify --data="false,2,0x5FfC014343cd971B7eb70732021E26C35B744cc4" --types="bool,uint256,address" --signature=0xcefd09e935b867a231086f41d98644655081a6e4e87c43e05fbbf621dfda69ea305c64fcf73907e09ce242c8ab8bcb953c4b45dd78262d8e34b22a8e4309734f00 --signer=0x0x5FfC014343cd971B7eb70732021E26C35B744cc4

In quiet mode this will return 0 if the signature is valid, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(signatureDataStr != "", quiet, "--data is required")
		cli.Assert(signatureVerifySigner != "", quiet, "--signer is required")

		dataHash := generateDataHash()

		signature, err := hex.DecodeString(strings.TrimPrefix(signatureVerifySignature, "0x"))
		cli.ErrCheck(err, quiet, "Invalid signature")

		key, err := crypto.SigToPub(dataHash, signature)
		cli.ErrCheck(err, quiet, "Failed to signer signature")
		// nolint:staticcheck
		signer := crypto.PubkeyToAddress(*key)

		verifySigner := common.HexToAddress(signatureVerifySigner)

		if bytes.Equal(signer.Bytes(), verifySigner.Bytes()) {
			outputIf(!quiet, "Verified")
			os.Exit(exitSuccess)
		}
		outputIf(!quiet, "Not verified")
		os.Exit(exitFailure)
	},
}

func init() {
	offlineCmds["signature:verify"] = true
	signatureCmd.AddCommand(signatureVerifyCmd)
	signatureFlags(signatureVerifyCmd)
	signatureVerifyCmd.Flags().StringVar(&signatureVerifySignature, "signature", "", "Hex string signature from which to verify the signer")
	signatureVerifyCmd.Flags().StringVar(&signatureVerifySigner, "signer", "", "Address of the signer")
}
