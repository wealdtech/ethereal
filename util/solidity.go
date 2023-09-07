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

package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

// Contract contains some basic information about a contract.
type Contract struct {
	Name   string
	Abi    abi.ABI
	Binary []byte
}

// ParseCombinedJSON parses a combined JSON output of solc for a specific contract.
func ParseCombinedJSON(input string, name string) (*Contract, error) {
	var err error
	var data []byte
	if strings.HasPrefix(input, "{") {
		data = []byte(input)
	} else {
		// Input is a filename.
		data, err = os.ReadFile(input)
		if err != nil {
			return nil, err
		}
	}

	var contractsJSON map[string]any
	err = json.Unmarshal(data, &contractsJSON)
	if err != nil {
		return nil, err
	}
	contracts, exists := contractsJSON["contracts"]
	if !exists {
		return nil, errors.New("JSON does not contain contracts element")
	}
	contractsMap := contracts.(map[string]any)
	// See if this is our name.
	for contractKey, contractValue := range contractsMap {
		if strings.HasSuffix(contractKey, fmt.Sprintf(":%s", name)) {
			// Found our contract.
			contract := &Contract{Name: name}

			contractJSON := contractValue.(map[string]any)

			// Obtain ABI.
			abiJSON, exists := contractJSON["abi"]
			if exists {
				bytes, err := json.Marshal(abiJSON)
				if err != nil {
					return nil, err
				}
				var abi abi.ABI
				if err := json.Unmarshal(bytes, &abi); err != nil {
					return nil, errors.Wrap(err, "failed to decode json")
				}
				contract.Abi = abi
			}

			// Obtain binary.
			binStr, exists := contractJSON["bin"]
			if exists {
				bin, err := hex.DecodeString(binStr.(string))
				if err != nil {
					return nil, err
				}
				contract.Binary = bin
			}

			return contract, nil
		}
	}
	return nil, fmt.Errorf("no contract \"%s\" in JSON; use --name to provide the name of the contract", name)
}
