// Copyright © 2017 - 2023 Weald Technology Trading.
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

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
	"github.com/wealdtech/ethereal/v2/cli"
	"golang.org/x/text/unicode/norm"
)

var (
	hdKeysPath     string
	hdKeysSecret   string
	hdKeysMnemonic string
)

// hdKeysCmd represents the hd keys command.
var hdKeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Display keys for a given seed and path",
	Long: `Display public and private keys for a given seed and path.  For example:

    ethereal hd keys --seed="correct horse battery staple" --path="m/44'/60'/0'/0/0"

In quiet mode this will return 0 if the keys were successfully obtained, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(hdKeysMnemonic != "", quiet, "seed is required")

		components := strings.Split(hdKeysPath, "/")
		cli.Assert(len(components) > 3, quiet, "Invalid path")
		path := make([]uint32, len(components)-1)
		for i := 0; i < len(components)-1; i++ {
			if strings.HasSuffix(components[i+1], "'") {
				unhardened, err := strconv.ParseInt(components[i+1][:len(components[i+1])-1], 10, 0)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid path component %v", components[i+1]))
				path[i] = 0x80000000 + uint32(unhardened)
			} else {
				unhardened, err := strconv.ParseInt(components[i+1], 10, 0)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Invalid path component %v", components[i+1]))
				path[i] = uint32(unhardened)
			}
		}

		seed, err := bip39.NewSeedWithErrorChecking(expandMnemonic(hdKeysMnemonic), hdKeysSecret)
		cli.ErrCheck(err, quiet, "Failed to obtain seed from mnemonic")

		masterKey, err := bip32.NewMasterKey(seed)
		cli.ErrCheck(err, quiet, "Failed to obtain master key from seed")

		childKey := masterKey
		for i := range path {
			childKey, err = childKey.NewChildKey(path[i])
			cli.ErrCheck(err, quiet, fmt.Sprintf("failed to derive child key from path component %d (%x)", i, path[i]))
		}

		key, err := crypto.ToECDSA(childKey.Key)
		cli.ErrCheck(err, quiet, "Failed to obtain private key from master key")

		outputIf(!quiet, fmt.Sprintf("Private key:\t\t0x%032x", key.D))
		outputIf(!quiet, fmt.Sprintf("Public key:\t\t0x%s", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey))))
		outputIf(!quiet, fmt.Sprintf("Ethereum address:\t%s", crypto.PubkeyToAddress(key.PublicKey).Hex()))

		os.Exit(exitSuccess)
	},
}

// expandMnmenonic expands mnemonics from their 4-letter versions.
func expandMnemonic(input string) string {
	wordList := bip39.GetWordList()
	truncatedWords := make(map[string]string, len(wordList))
	for _, word := range wordList {
		if len(word) > 4 {
			truncatedWords[firstFour(word)] = word
		}
	}
	mnemonicWords := strings.Split(input, " ")
	for i := range mnemonicWords {
		if fullWord, exists := truncatedWords[norm.NFKC.String(mnemonicWords[i])]; exists {
			mnemonicWords[i] = fullWord
		}
	}
	return strings.Join(mnemonicWords, " ")
}

// firstFour provides the first four letters for a potentially longer word.
func firstFour(s string) string {
	// Use NFKC here for composition, to avoid accents counting as their own characters.
	s = norm.NFKC.String(s)
	r := []rune(s)
	if len(r) > 4 {
		return string(r[:4])
	}
	return s
}

func init() {
	offlineCmds["hd:keys"] = true
	hdCmd.AddCommand(hdKeysCmd)
	hdKeysCmd.Flags().StringVar(&hdKeysMnemonic, "seed", "", "12- or 24-word BIP-39 seed phrase")
	hdKeysCmd.Flags().StringVar(&hdKeysSecret, "secret", "", "optional secret to add to seed")
	hdKeysCmd.Flags().StringVar(&hdKeysPath, "path", "", "path for keys (e.g. m/44'/60'/0'/0/0)")
}
