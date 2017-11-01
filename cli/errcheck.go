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

package cli

import (
	"fmt"
	"os"
)

// ErrCheck checks for an error and quits if it is present
func ErrCheck(err error, quiet bool, msg string) {
	if err != nil {
		if !quiet {
			fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
		}
		os.Exit(1)
	}
}

// ErrAssert checks a condition and quits if it is false
func ErrAssert(condition bool, err error, quiet bool, msg string) {
	if !condition {
		if err != nil {
			if !quiet {
				fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
			}
			os.Exit(1)
		}
	}
}

// Assert checks a condition and quits if it is false
func Assert(condition bool, quiet bool, msg string) {
	if !condition {
		Err(quiet, msg)
	}
}

// Err prints an erro rand quits
func Err(quiet bool, msg string) {
	if !quiet {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
	os.Exit(1)
}
