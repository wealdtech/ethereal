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
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	ens "github.com/wealdtech/go-ens/v3"
)

var signatureSignerSignature string

// signatureSignerCmd represents the signature signer command.
var signatureSignerCmd = &cobra.Command{
	Use:   "signer",
	Short: "Signer of a signature",
	Long: `Obtain the signer of a presented signature.  For example:

    ethereal signature signer --data="false,2,0x5FfC014343cd971B7eb70732021E26C35B744cc4" --types="bool,uint256,address" --signature=0xcefd09e935b867a231086f41d98644655081a6e4e87c43e05fbbf621dfda69ea305c64fcf73907e09ce242c8ab8bcb953c4b45dd78262d8e34b22a8e4309734f00

In quiet mode this will return 0 if the signature provides a valid signer, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(signatureDataStr != "", quiet, "--data is required")

		dataHash := generateDataHash()

		signature, err := hex.DecodeString(strings.TrimPrefix(signatureSignerSignature, "0x"))
		cli.ErrCheck(err, quiet, "Invalid signature")

		key, err := crypto.SigToPub(dataHash, signature)
		cli.ErrCheck(err, quiet, "Failed to signer signature")
		// nolint:staticcheck
		address := crypto.PubkeyToAddress(*key)

		if quiet {
			os.Exit(exitSuccess)
		}

		fmt.Printf("%s\n", ens.Format(c.Client(), address))
	},
}

func init() {
	offlineCmds["signature:signer"] = true
	signatureCmd.AddCommand(signatureSignerCmd)
	signatureFlags(signatureSignerCmd)
	signatureSignerCmd.Flags().StringVar(&signatureSignerSignature, "signature", "", "Hex string signature from which to obtain the signer")
}
