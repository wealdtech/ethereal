// Copyright © 2017-2025 Weald Technology Trading.
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

package cmd

import (
	"fmt"
	"os"
	dbg "runtime/debug"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of Ethereal",
	Long: `Obtain the version of Ethereal.  For example:

    ethereal version.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("2.11.5")
		if viper.GetBool("verbose") {
			buildInfo, ok := dbg.ReadBuildInfo()
			if ok {
				fmt.Printf("Package: %s\n", buildInfo.Path)
				fmt.Println("Dependencies:")
				for _, dep := range buildInfo.Deps {
					for dep.Replace != nil {
						dep = dep.Replace
					}
					fmt.Printf("\t%v %v\n", dep.Path, dep.Version)
				}
			}
		}
		os.Exit(exitSuccess)
	},
}

func init() {
	offlineCmds["version"] = true
	RootCmd.AddCommand(versionCmd)
}
