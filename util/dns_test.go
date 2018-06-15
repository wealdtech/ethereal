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
package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWireFormat(t *testing.T) {
	tests := []struct {
		input  string
		output []byte
	}{
		{input: "", output: []byte{0x00}},
		{input: ".", output: []byte{0x00}},
		{input: "eth.", output: []byte{0x03, 0x65, 0x74, 0x68, 0x00}},
		{input: ".eth.", output: []byte{0x03, 0x65, 0x74, 0x68, 0x00}},
		{input: "test.eth.", output: []byte{0x04, 0x74, 0x65, 0x73, 0x74, 0x03, 0x65, 0x74, 0x68, 0x00}},
	}
	for _, tt := range tests {
		output := DNSWireFormat(tt.input)
		assert.Equal(t, tt.output, output)
	}
}
