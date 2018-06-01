package util

import (
	"strings"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// DNSDomainHash hashes a domain name
func DNSDomainHash(domain string) (hash [32]byte) {
	lower := strings.ToLower(domain)
	sha := sha3.NewKeccak256()
	sha.Write([]byte(lower))
	sha.Sum(hash[:0])
	return
}
