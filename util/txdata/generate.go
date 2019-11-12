// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

func main() {
	file, err := os.Create("signatures.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, header)
	if err != nil {
		panic(err)
	}

	err = addFunctionSignatures(file)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(file, footer)
	if err != nil {
		panic(err)
	}
}

func addFunctionSignatures(file *os.File) error {
	// Set up the key and value builders
	var keyBuilder strings.Builder
	keyBuilder.WriteString("var functionsKeys = [][4]byte{\n")
	var valBuilder strings.Builder
	valBuilder.WriteString("var functionsValues = []function{\n")

	// Fetch data from 4byte one page at a time
	resp, err := http.Get("https://www.4byte.directory/api/v1/signatures/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := make(map[string]interface{})
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	// Loop over each page (100 items per page)
	count := int(res["count"].(float64))
	for page := 0; page < count/100; page++ {
		for _, result := range res["results"].([]interface{}) {
			signature := result.(map[string]interface{})["text_signature"].(string)
			sigBits := strings.Split(strings.TrimSuffix(signature, ")"), "(")
			name := sigBits[0]
			params := strings.Split(sigBits[1], ",")
			if params[0] == "" {
				params = make([]string, 0)
			}
			for i := range params {
				params[i] = strings.TrimSpace(params[i])
				params[i] = strings.Split(params[i], " ")[0]
			}

			signature = fmt.Sprintf("%s(%s)", name, strings.Join(params, ","))
			var hash [32]byte
			sha := sha3.NewLegacyKeccak256()
			sha.Write([]byte(signature))
			sha.Sum(hash[:0])
			var sig [4]byte
			copy(sig[:], hash[:4])

			addFunctionSignature(&keyBuilder, &valBuilder, sig, name, params)
		}

		resp, err := http.Get(fmt.Sprintf("https://www.4byte.directory/api/v1/signatures/?page=%d", page+1))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		res = make(map[string]interface{})
		err = json.Unmarshal(data, &res)
		if err != nil {
			return err
		}
	}

	keyBuilder.WriteString("}\n")
	_, err = io.WriteString(file, keyBuilder.String())
	if err != nil {
		return err
	}
	valBuilder.WriteString("}\n")
	_, err = io.WriteString(file, valBuilder.String())
	return err
}

func addFunctionSignature(keyBuilder *strings.Builder, valBuilder *strings.Builder, sig [4]byte, name string, params []string) {
	keyBuilder.WriteString(fmt.Sprintf("[4]byte{0x%02x,0x%02x,0x%02x,0x%02x},\n", sig[0], sig[1], sig[2], sig[3]))

	valBuilder.WriteString(fmt.Sprintf("function{name:\"%s\",params:[]string{", name))
	for i := 0; i < len(params)-1; i++ {
		valBuilder.WriteString(fmt.Sprintf("\"%s\",", params[i]))
	}
	if len(params) > 0 {
		valBuilder.WriteString(fmt.Sprintf("\"%s\"", params[len(params)-1]))
	}
	valBuilder.WriteString("}},\n")
}

const header = `
// Copyright © 2018, 2019 Weald Technology Trading
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

package txdata

`

const footer = `
    func init() {
        functions = make(map[[4]byte]function)
        for i := 0; i < len(functionsKeys); i++ {
            functions[functionsKeys[i]] = functionsValues[i]
		}
    }
`
