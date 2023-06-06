package util

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
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
	"github.com/wealdtech/ethereal/v2/cli"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

// PrivateKeyForAccount returns the private key for an account.
func PrivateKeyForAccount(chainID *big.Int, address common.Address, passphrase string) (*ecdsa.PrivateKey, error) {
	var account *accounts.Account
	_, account, err := cli.ObtainWalletAndAccount(chainID, address)
	if err != nil {
		return nil, fmt.Errorf("unable to find account %s", address.Hex())
	}
	ks := keystore.NewKeyStore(filepath.Dir(account.URL.Path), keystore.StandardScryptN, keystore.StandardScryptP)
	exported, err := ks.Export(*account, passphrase, passphrase)
	if err != nil {
		return nil, fmt.Errorf("unable to export keystore for %v", address)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(exported, &m)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal keystore for %v", address)
	}
	var keyBytes []byte
	if version, ok := m["version"].(string); ok && version == "1" {
		k := new(encryptedKeyJSONV1)
		err = json.Unmarshal(exported, k)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal keystore for %v", address)
		}
		keyBytes, _, err = decryptKeyV1(k, passphrase)
	} else {
		k := new(encryptedKeyJSONV3)
		err = json.Unmarshal(exported, k)
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal keystore for %v", address)
		}
		keyBytes, _, err = decryptKeyV3(k, passphrase)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal keystore for %v", address)
	}
	return crypto.ToECDSAUnsafe(keyBytes), nil
}

const (
	keyHeaderKDF = "scrypt"

	version = 3
)

var errDecrypt = errors.New("could not decrypt key with given passphrase")

func decryptKeyV3(keyProtected *encryptedKeyJSONV3, auth string) ([]byte, []byte, error) {
	if keyProtected.Version != version {
		return nil, nil, fmt.Errorf("version not supported: %v", keyProtected.Version)
	}

	if keyProtected.Crypto.Cipher != "aes-128-ctr" {
		return nil, nil, fmt.Errorf("cipher not supported: %v", keyProtected.Crypto.Cipher)
	}

	keyID := uuid.Parse(keyProtected.ID)
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
		return nil, nil, errDecrypt
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, nil, err
	}
	return plainText, keyID, err
}

func decryptKeyV1(keyProtected *encryptedKeyJSONV1, auth string) ([]byte, []byte, error) {
	keyID := uuid.Parse(keyProtected.ID)
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
		return nil, nil, errDecrypt
	}

	plainText, err := aesCBCDecrypt(crypto.Keccak256(derivedKey[:16])[:16], cipherText, iv)
	if err != nil {
		return nil, nil, err
	}
	return plainText, keyID, err
}

type encryptedKeyJSONV3 struct {
	Address string     `json:"address"`
	Crypto  cryptoJSON `json:"crypto"`
	ID      string     `json:"id"`
	Version int        `json:"version"`
}

type encryptedKeyJSONV1 struct {
	Address string     `json:"address"`
	Crypto  cryptoJSON `json:"crypto"`
	ID      string     `json:"id"`
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
			return nil, fmt.Errorf("unsupported PBKDF2 PRF: %s", prf)
		}
		key := pbkdf2.Key(authArray, salt, c, dkLen, sha256.New)
		return key, nil
	}

	return nil, fmt.Errorf("unsupported KDF: %s", cryptoJSON.KDF)
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
		return nil, errDecrypt
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
