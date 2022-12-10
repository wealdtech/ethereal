#!/bin/bash 

BASE='https://www.4byte.directory/api/v1/signatures/'
COUNT=`curl -s "${BASE}" | jq .count`
PAGES=$((1+$COUNT/100))
OUTPUT=signatures.go

cat >${OUTPUT} <<EOSTART
// Copyright © 2018 - 2022 Weald Technology Trading
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

// nolint:misspell
var sigs = []string{
EOSTART
for PAGE in $(seq 1 $PAGES)
do
  for FUNC in `curl -s "${BASE}?page=${PAGE}" | jq '.results[].text_signature'`
  do
    echo ${FUNC}, >>${OUTPUT}
  done
done
cat >>${OUTPUT} <<EOEND 
}

// InitFunctionMap initialises the function (and event) map with known signatures
func InitFunctionMap() {
    functions = make(map[[4]byte]function)
    events = make(map[[32]byte]function)

    for _, f := range sigs {
      AddFunctionSignature(f)
    }

    // Also add events
    initEventMap()
}
EOEND

# Sort the functions in-place
echo 'x' | ex -s -c '17,$-14!sort' ${OUTPUT}

# Tidy
gofmt -w ${OUTPUT}
