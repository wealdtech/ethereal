// Copyright 2017 Orinoco Payments
//
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

package ens

import (
	"errors"
	"strings"

	"golang.org/x/net/idna"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

var p = idna.New(idna.MapForLookup(), idna.StrictDomainName(true), idna.Transitional(false))

// Normalize normalizes a name according to the ENS rules
func Normalize(input string) (output string) {
	output, err := p.ToUnicode(input)
	if err != nil {
		// TODO find out why ToUnicode() might fail and handle it here
		output = input
	}
	// If the name started with a period then ToUnicode() removes it, but we want to keep it
	if strings.HasPrefix(input, ".") && !strings.HasPrefix(output, ".") {
		output = "." + output
	}
	return
}

// LabelHash generates a simple hash for a piece of a name.
func LabelHash(label string) (hash [32]byte) {
	if label == "" {
		return
	}
	normalizedLabel := Normalize(label)

	sha := sha3.NewKeccak256()
	sha.Write([]byte(normalizedLabel))
	sha.Sum(hash[:0])
	return
}

// NameHash generates a hash from a name that can be used to
// look up the name in ENS
func NameHash(name string) (hash [32]byte) {
	if name == "" {
		return
	}
	normalizedName := Normalize(name)
	parts := strings.Split(normalizedName, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		hash = nameHashPart(hash, parts[i])
	}
	return
}

// Domain returns the domain directly before the '.eth' in a name
func Domain(name string) (domain string, err error) {
	normalizedName := Normalize(name)
	nameBits := strings.Split(normalizedName, ".")
	if len(nameBits) < 2 {
		err = errors.New("invalid name")
		return
	}

	domain = nameBits[len(nameBits)-2]
	return
}

func nameHashPart(currentHash [32]byte, name string) (hash [32]byte) {
	sha := sha3.NewKeccak256()
	sha.Write(currentHash[:])
	nameSha := sha3.NewKeccak256()
	nameSha.Write([]byte(name))
	nameHash := nameSha.Sum(nil)
	sha.Write(nameHash)
	sha.Sum(hash[:0])
	return
}
