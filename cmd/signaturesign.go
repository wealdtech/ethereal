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
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/v2/cli"
	"github.com/wealdtech/ethereal/v2/util"
)

var (
	signatureSignSigner     string
	signatureSignPrivateKey string
	signatureSignPassphrase string
)

// signatureSignCmd represents the signature sign command.
var signatureSignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign data",
	Long: `Sign presented data.  For example:

    ethereal signature sign --data="false,2,0x5FfC014343cd971B7eb70732021E26C35B744cc4" --types="bool,uint256,address" --signer=0x1234...5678 --passphrase=secret

In quiet mode this will return 0 if the data can be signed, otherwise 1.

Signing data in Ethereum is complex, so details of exactly how this operates are
provided below:
  - data is turned in to a suitable representation:
    - if data is a hex string it is kept as-is
    - if types is provided data is assumed to be a set of comma-separated values
      corresponding to the types provided
    - otherwise data is treated as a simple string
  - data is potentially hashed:
    - if the data is a hex string or a set of values then the data will be
	  hashed by default
    - if data is a simple string it will not be hashed by default
	- hashing can be forced on or off with '--hash=true' or '--hash=false'
    provide a simple 32-byte value, otherwise it is left as-is
  - the message is created as the data prepended with the standard Ethereum
    signing message of "\\x19Ethereum Signed Message:\n" followed by the
	number of bytes in the data and finally the data itself, for example
    "\\x19Ethereum Signed Message:\n11Hello world"
  - the message is signed with the provided account or private key
`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(signatureDataStr != "", quiet, "--data is required")

		dataHash := generateDataHash()

		// Sign the hash.
		var signature []byte
		var key *ecdsa.PrivateKey
		var err error
		if signatureSignPassphrase != "" {
			signer := common.HexToAddress(signatureSignSigner)
			key, err = util.PrivateKeyForAccount(c.ChainID(), signer, signatureSignPassphrase)
			cli.ErrCheck(err, quiet, "Invalid account or passphrse")
		} else if signatureSignPrivateKey != "" {
			key, err = crypto.HexToECDSA(strings.TrimPrefix(signatureSignPrivateKey, "0x"))
			cli.ErrCheck(err, quiet, "Invalid private key")
		} else {
			cli.Err(quiet, "no passphrase or private key; cannot sign")
		}
		signature, err = crypto.Sign(dataHash, key)
		cli.ErrCheck(err, quiet, "Failed to sign data")

		if quiet {
			os.Exit(exitSuccess)
		}

		fmt.Printf("%x\n", signature)
	},
}

func init() {
	offlineCmds["signature:sign"] = true
	signatureCmd.AddCommand(signatureSignCmd)
	signatureFlags(signatureSignCmd)
	signatureSignCmd.Flags().StringVar(&signatureSignSigner, "signer", "", "Address of the account to sign the data")
	signatureSignCmd.Flags().StringVar(&signatureSignPassphrase, "passphrase", "", "Passphrase of the account to sign the data")
	signatureSignCmd.Flags().StringVar(&signatureSignPrivateKey, "privatekey", "", "Private key to sign the data")
}
