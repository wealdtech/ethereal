// Copyright © 2019 Weald Technology Trading
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

package funcparser

//nogo:generate java -jar /home/jgm/tools/antlr/antlr-4.13.0-complete.jar -Dlanguage=Go -o parser Func.g4
//go:generate abigen --abi Tester.abi --pkg funcparser --type tester --out tester.go
//go:generate java -Xmx500M -jar /home/jgm/tools/antlr/antlr-4.13.0-complete.jar -Dlanguage=Go -package parser -o parser Func.g4
//go:generate rm -f parser/Func.interp parser/Func.tokens parser/FuncLexer.interp parser/FuncLexer.tokens
