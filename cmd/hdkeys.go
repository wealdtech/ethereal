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
	"strconv"
	"strings"

	bip32 "github.com/FactomProject/go-bip32"
	bip39 "github.com/FactomProject/go-bip39"
	bip44 "github.com/FactomProject/go-bip44"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var hdKeysPath string
var hdKeysSecret string
var hdKeysSeed string

// hdKeysCmd represents the hd keys command
var hdKeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Display keys for a given seed and path",
	Long: `Display public and private keys for a given seed and path.  For example:

    ethereal hd keys --seed="correct horse battery staple" --path="m/44'/60'/0'/0/0"

In quiet mode this will return 0 if the keys were successfully obtained, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(hdKeysSeed != "", quiet, "seed is required")

		components := strings.Split(hdKeysPath, "/")
		cli.Assert(len(components) == 6, quiet, "Invalid path")
		path := make([]uint32, 4)
		for i := 0; i < 4; i++ {
			if strings.HasSuffix(components[i+2], "'") {
				unhardened, err := strconv.ParseInt(components[i+2][:len(components[i+2])-1], 10, 0)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid path component %v", components[i+2]))
				path[i] = 0x80000000 + uint32(unhardened)
			} else {
				unhardened, err := strconv.ParseInt(components[i+2], 10, 0)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid path component %v", components[i+2]))
				path[i] = uint32(unhardened)
			}
		}

		seed, err := bip39.NewSeedWithErrorChecking(hdKeysSeed, hdKeysSecret)
		cli.ErrCheck(err, quiet, "Failed to obtain seed from passphrase")

		masterKey, err := bip32.NewMasterKey(seed)
		cli.ErrCheck(err, quiet, "Failed to obtain master key from seed")

		childKey, err := bip44.NewKeyFromMasterKey(masterKey, path[0], path[1], path[2], path[3])
		cli.ErrCheck(err, quiet, "Failed to obtain child key from master key")

		key, err := crypto.ToECDSA(childKey.Key)
		cli.ErrCheck(err, quiet, "Failed to obtain private key from master key")

		outputIf(!quiet, fmt.Sprintf("Private key:\t\t0x%032x", key.D))
		outputIf(!quiet, fmt.Sprintf("Public key:\t\t0x%s", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey))))
		outputIf(!quiet, fmt.Sprintf("Ethereum address:\t%s", crypto.PubkeyToAddress(key.PublicKey).Hex()))

		os.Exit(exitSuccess)

		//		cli.Assert((hdKeysAddress != "" && hdKeysPassphrase != "") || hdKeysPrivateKey != "", quiet, "--privatekey or both of --address and --passphrase are required")
		//
		//		var key *ecdsa.PrivateKey
		//		if hdKeysPrivateKey != "" {
		//			key, err = crypto.HexToECDSA(strings.TrimPrefix(hdKeysPrivateKey, "0x"))
		//			cli.ErrCheck(err, quiet, "Invalid private key")
		//		} else {
		//			address := common.HexToAddress(hdKeysAddress)
		//			key, err = util.PrivateKeyForAccount(chainID, address, hdKeysPassphrase)
		//		}
		//		cli.ErrCheck(err, quiet, "Failed to access account")
		//		if quiet {
		//			os.Exit(_exit_success)
		//		}
		//
		//		fmt.Printf("Private key:\t\t0x%032x\n", key.D)
		//		fmt.Printf("Public key:\t\t0x%s\n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
		//		fmt.Printf("Ethereum address:\t%s\n", crypto.PubkeyToAddress(key.PublicKey).Hex())
	},
}

func init() {
	offlineCmds["hd:keys"] = true
	hdCmd.AddCommand(hdKeysCmd)
	hdKeysCmd.Flags().StringVar(&hdKeysSeed, "seed", "", "12- or 24-word BIP-39 seed phrase")
	hdKeysCmd.Flags().StringVar(&hdKeysSecret, "secret", "", "optional secret to add to seed")
	hdKeysCmd.Flags().StringVar(&hdKeysPath, "path", "", "path for keys (e.g. m/44'/60'/0'/0/0)")
}
