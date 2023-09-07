// Copyright © 2017 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package util

import (
	"testing"
)

func TestParseCombinedJSON(t *testing.T) {
	tests := []struct {
		input    string
		inputErr error
		name     string
	}{
		{`{"contracts":{"Simple.sol:Simple":{"abi":[{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}],"bin":"6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550603580605d6000396000f3006080604052600080fd00a165627a7a7230582083d530c10e079e85c2f5030dc4ff81c9c24ad62d6af39470de661d4596b5766e0029"}},"version":"0.4.23+commit.124ca40d.Linux.g++"}`, nil, "Simple"},
	}

	for _, tt := range tests {
		_, err := ParseCombinedJSON(tt.input, tt.name)
		if err != tt.inputErr {
			t.Errorf("Failure: parsing resulted in %v (expected %v)", err, tt.inputErr)
		}
	}
}
