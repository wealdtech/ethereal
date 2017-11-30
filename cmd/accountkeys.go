// Copyright © 2017 Weald Technology Trading
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
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

var accountKeysAddress string
var accountKeysPassphrase string
var accountKeysPrivateKey string

const (
	keyHeaderKDF = "scrypt"

	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptN = 1 << 18

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptP = 1

	// LightScryptN is the N parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptN = 1 << 12

	// LightScryptP is the P parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptP = 6

	scryptR     = 8
	scryptDKLen = 32

	version = 3
)

var (
	ErrLocked  = accounts.NewAuthNeededError("password or unlock")
	ErrNoMatch = errors.New("no key for given address or file")
	ErrDecrypt = errors.New("could not decrypt key with given passphrase")
)

// accountKeysCmd represents the account keys command
var accountKeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Display keys for a given address",
	Long: `Display public and private keys for a given address.  For example:

    ethereal account keys --address=0x5FfC014343cd971B7eb70732021E26C35B744cc4 --passphrase=secret

Note that this will only work for filesystem-based keystores.  Hardware wallets never reveal their keys so these cannot be obtained.

In quiet mode this will return 0 if the account was successfully decoded, otherwise 1.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert((accountKeysAddress != "" && accountKeysPassphrase != "") || accountKeysPrivateKey != "", quiet, "--privatekey or both of --address and --passphrase are required")

		var key *ecdsa.PrivateKey
		if accountKeysPrivateKey != "" {
			if strings.HasPrefix(accountKeysPrivateKey, "0x") {
				key, err = crypto.HexToECDSA(accountKeysPrivateKey[2:])
			} else {
				key, err = crypto.HexToECDSA(accountKeysPrivateKey)
			}
			cli.ErrCheck(err, quiet, "Invalid private key")
			if quiet {
				os.Exit(0)
			}
		} else {
			var account *accounts.Account
			address := common.HexToAddress(accountKeysAddress)
			_, account, err = obtainWalletAndAccount(address)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to find account %s", accountKeysAddress))
			ks := keystore.NewKeyStore(filepath.Dir(account.URL.Path), keystore.StandardScryptN, keystore.StandardScryptP)
			exported, err := ks.Export(*account, accountKeysPassphrase, accountKeysPassphrase)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to export keystore for %s", accountKeysAddress))
			m := make(map[string]interface{})
			err = json.Unmarshal(exported, &m)
			cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to unmarshal keystore for %s", accountKeysAddress))
			var keyBytes []byte
			if version, ok := m["version"].(string); ok && version == "1" {
				k := new(encryptedKeyJSONV1)
				err = json.Unmarshal(exported, k)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to unmarshal keystore for %s", accountKeysAddress))
				keyBytes, _, err = decryptKeyV1(k, accountKeysPassphrase)
			} else {
				k := new(encryptedKeyJSONV3)
				err = json.Unmarshal(exported, k)
				cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to unmarshal keystore for %s", accountKeysAddress))
				keyBytes, _, err = decryptKeyV3(k, accountKeysPassphrase)
			}
			cli.ErrCheck(err, quiet, fmt.Sprintf("Unable to unmarshal keystore for %s", accountKeysAddress))
			key = crypto.ToECDSAUnsafe(keyBytes)

			if quiet {
				os.Exit(0)
			}
		}
		fmt.Printf("Private key:\t\t0x%032x\n", key.D)
		fmt.Printf("Public key:\t\t0x%s\n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
		fmt.Printf("Ethereum address:\t%s\n", crypto.PubkeyToAddress(key.PublicKey).Hex())
	},
}

func init() {
	accountCmd.AddCommand(accountKeysCmd)
	accountKeysCmd.Flags().StringVar(&accountKeysAddress, "address", "", "address for account keys")
	accountKeysCmd.Flags().StringVar(&accountKeysPassphrase, "passphrase", "", "passphrase for account keys")
	accountKeysCmd.Flags().StringVar(&accountKeysPrivateKey, "privatekey", "", "private key for account keys")
}

func decryptKeyV3(keyProtected *encryptedKeyJSONV3, auth string) (keyBytes []byte, keyId []byte, err error) {
	if keyProtected.Version != version {
		return nil, nil, fmt.Errorf("Version not supported: %v", keyProtected.Version)
	}

	if keyProtected.Crypto.Cipher != "aes-128-ctr" {
		return nil, nil, fmt.Errorf("Cipher not supported: %v", keyProtected.Crypto.Cipher)
	}

	keyId = uuid.Parse(keyProtected.Id)
	mac, err := hex.DecodeString(keyProtected.Crypto.MAC)
	if err != nil {
		return nil, nil, err
	}

	iv, err := hex.DecodeString(keyProtected.Crypto.CipherParams.IV)
	if err != nil {
		return nil, nil, err
	}

	cipherText, err := hex.DecodeString(keyProtected.Crypto.CipherText)
	if err != nil {
		return nil, nil, err
	}

	derivedKey, err := getKDFKey(keyProtected.Crypto, auth)
	if err != nil {
		return nil, nil, err
	}

	calculatedMAC := crypto.Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, nil, ErrDecrypt
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, nil, err
	}
	return plainText, keyId, err
}

func decryptKeyV1(keyProtected *encryptedKeyJSONV1, auth string) (keyBytes []byte, keyId []byte, err error) {
	keyId = uuid.Parse(keyProtected.Id)
	mac, err := hex.DecodeString(keyProtected.Crypto.MAC)
	if err != nil {
		return nil, nil, err
	}

	iv, err := hex.DecodeString(keyProtected.Crypto.CipherParams.IV)
	if err != nil {
		return nil, nil, err
	}

	cipherText, err := hex.DecodeString(keyProtected.Crypto.CipherText)
	if err != nil {
		return nil, nil, err
	}

	derivedKey, err := getKDFKey(keyProtected.Crypto, auth)
	if err != nil {
		return nil, nil, err
	}

	calculatedMAC := crypto.Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, nil, ErrDecrypt
	}

	plainText, err := aesCBCDecrypt(crypto.Keccak256(derivedKey[:16])[:16], cipherText, iv)
	if err != nil {
		return nil, nil, err
	}
	return plainText, keyId, err
}

type encryptedKeyJSONV3 struct {
	Address string     `json:"address"`
	Crypto  cryptoJSON `json:"crypto"`
	Id      string     `json:"id"`
	Version int        `json:"version"`
}

type encryptedKeyJSONV1 struct {
	Address string     `json:"address"`
	Crypto  cryptoJSON `json:"crypto"`
	Id      string     `json:"id"`
	Version string     `json:"version"`
}
type cryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherparamsJSON       `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
	IV string `json:"iv"`
}

func getKDFKey(cryptoJSON cryptoJSON, auth string) ([]byte, error) {
	authArray := []byte(auth)
	salt, err := hex.DecodeString(cryptoJSON.KDFParams["salt"].(string))
	if err != nil {
		return nil, err
	}
	dkLen := ensureInt(cryptoJSON.KDFParams["dklen"])

	if cryptoJSON.KDF == keyHeaderKDF {
		n := ensureInt(cryptoJSON.KDFParams["n"])
		r := ensureInt(cryptoJSON.KDFParams["r"])
		p := ensureInt(cryptoJSON.KDFParams["p"])
		return scrypt.Key(authArray, salt, n, r, p, dkLen)

	} else if cryptoJSON.KDF == "pbkdf2" {
		c := ensureInt(cryptoJSON.KDFParams["c"])
		prf := cryptoJSON.KDFParams["prf"].(string)
		if prf != "hmac-sha256" {
			return nil, fmt.Errorf("Unsupported PBKDF2 PRF: %s", prf)
		}
		key := pbkdf2.Key(authArray, salt, c, dkLen, sha256.New)
		return key, nil
	}

	return nil, fmt.Errorf("Unsupported KDF: %s", cryptoJSON.KDF)
}

func ensureInt(x interface{}) int {
	res, ok := x.(int)
	if !ok {
		res = int(x.(float64))
	}
	return res
}
func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	// AES-128 is selected due to size of encryptKey.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}

func aesCBCDecrypt(key, cipherText, iv []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypter := cipher.NewCBCDecrypter(aesBlock, iv)
	paddedPlaintext := make([]byte, len(cipherText))
	decrypter.CryptBlocks(paddedPlaintext, cipherText)
	plaintext := pkcs7Unpad(paddedPlaintext)
	if plaintext == nil {
		return nil, ErrDecrypt
	}
	return plaintext, err
}
func pkcs7Unpad(in []byte) []byte {
	if len(in) == 0 {
		return nil
	}

	padding := in[len(in)-1]
	if int(padding) > len(in) || padding > aes.BlockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
		if in[i] != padding {
			return nil
		}
	}
	return in[:len(in)-int(padding)]
}
