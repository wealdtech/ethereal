// Copyright © 2017-2019 Weald Technology Trading
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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
)

var ipfsUploadFilename string

// ipfsUploadCmd represents the ipfs upload command
var ipfsUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file on to an IPFS gateway",
	Long: `Set the data of a name registered with the Ethereum Name Service (ENS) for a given name.  For example:

    ethereal ipfs upload --file=/path/to/file

This will return an exit status of 0 if the file is successfully stored on the gateway, 1 if the transaction is not successfully submitted, and 2 if the transaction is successfully submitted but not mined within the supplied time limit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ipfsUploadFilename != "", quiet, "--file is required")

		err := initIPFSProvider()
		cli.ErrCheck(err, quiet, "Failed to set up IPFS provider")

		file, err := os.Open(ipfsUploadFilename)
		cli.ErrCheck(err, quiet, "Failed to access file")

		hash, err := ipfsProvider.PinContent(filepath.Base(ipfsUploadFilename), file, nil)
		cli.ErrCheck(err, quiet, "Failed to upload file")

		outputIf(!quiet, hash)
		os.Exit(_exit_success)
	},
}

func init() {
	offlineCmds["ipfs:upload"] = true
	ipfsCmd.AddCommand(ipfsUploadCmd)
	ipfsFlags(ipfsUploadCmd)
	ipfsUploadCmd.Flags().StringVar(&ipfsUploadFilename, "file", "", "Path of the file to upload")
}
