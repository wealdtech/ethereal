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
	"bytes"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/orinocopay/go-etherutils"
	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens"
)

var zero = big.NewInt(0)

// ensInfoCmd represents the ens info command
var ensInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Obtain information about an ENS domain",
	Long: `Obtain information about a domain registered with the Ethereum Name Service (ENS).  For example:

    ens info enstest.eth

In quiet mode this will return 0 if the domain is owned, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		cli.Assert(ensDomain != "", quiet, "--domain is required")

		ensDomain = ens.NormaliseDomain(ensDomain)
		outputIf(verbose, fmt.Sprintf("Normalised domain is %s", ensDomain))

		outputIf(verbose, fmt.Sprintf("Top-level domain is %s", ens.Tld(ensDomain)))
		registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(ensDomain))
		cli.ErrCheck(err, quiet, "Failed to obtain ENS registrar contract")

		outputIf(verbose, fmt.Sprintf("Name hash is 0x%x", ens.NameHash(ensDomain)))
		registry, err := ens.RegistryContract(client)
		cli.ErrCheck(err, quiet, "Failed to obtain registry contract")
		domainOwnerAddress, err := registry.Owner(nil, ens.NameHash(ensDomain))
		cli.ErrCheck(err, quiet, "Failed to obtain domain owner")

		if ens.DomainLevel(ensDomain) == 1 {
			state, err := ens.State(registrarContract, client, ensDomain)
			if err == nil {
				if quiet {
					if state == "Owned" {
						os.Exit(_exit_success)
					} else {
						os.Exit(_exit_failure)
					}
				} else {
					switch state {
					case "Available":
						availableInfo(ensDomain)
					case "Bidding":
						biddingInfo(ensDomain)
					case "Revealing":
						revealingInfo(ensDomain)
					case "Won":
						wonInfo(ensDomain)
					case "Owned":
						domainOwnerName, _ := ens.ReverseResolve(client, &domainOwnerAddress)
						if domainOwnerName == "" {
							fmt.Printf("Domain owner is %s\n", domainOwnerAddress.Hex())
						} else {
							fmt.Printf("Domain owner is %s (%s)\n", domainOwnerName, domainOwnerAddress.Hex())
						}
						ownedInfo(ensDomain)
					default:
						fmt.Println(state)
					}
				}
			} else {
				ownedInfo(ensDomain)
			}
		} else {
			domainOwnerName, _ := ens.ReverseResolve(client, &domainOwnerAddress)
			if domainOwnerName == "" {
				fmt.Printf("Domain owner is %s\n", domainOwnerAddress.Hex())
			} else {
				fmt.Printf("Domain owner is %s (%s)\n", domainOwnerName, domainOwnerAddress.Hex())
			}
			genericInfo(ensDomain)
		}
	},
}

func init() {
	ensCmd.AddCommand(ensInfoCmd)
	ensFlags(ensInfoCmd)
}

func availableInfo(name string) {
	if len(name) < 11 { // 7 + 4 for '.eth'
		fmt.Println("Unavailable due to name length restrictions")
	} else {
		fmt.Println("Available")
	}
}

func biddingInfo(name string) {
	registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(name))
	cli.ErrCheck(err, quiet, "Failed to obtain ENS registrar contract")
	_, _, registrationDate, _, _, err := ens.Entry(registrarContract, client, name)
	cli.ErrCheck(err, quiet, "Cannot obtain auction status")
	twoDaysAgo := time.Duration(-48) * time.Hour
	fmt.Println("Bidding until", registrationDate.Add(twoDaysAgo))
}

func revealingInfo(name string) {
	registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(name))
	cli.ErrCheck(err, quiet, "Failed to obtain ENS registrar contract")
	_, _, registrationDate, value, highestBid, err := ens.Entry(registrarContract, client, name)
	cli.ErrCheck(err, quiet, "Cannot obtain information for that name")
	fmt.Println("Revealing until", registrationDate)
	// If the value is 0 then it is is minvalue instead
	if value.Cmp(zero) == 0 {
		value, _ = etherutils.StringToWei("0.01 ether")
	}
	fmt.Println("Locked value is", etherutils.WeiToString(value, true))
	fmt.Println("Highest bid is", etherutils.WeiToString(highestBid, true))
}

func wonInfo(name string) {
	registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(name))
	cli.ErrCheck(err, quiet, "Failed to obtain ENS registrar contract")
	_, deedAddress, registrationDate, value, highestBid, err := ens.Entry(registrarContract, client, name)
	cli.ErrCheck(err, quiet, "Cannot obtain information for that name")
	fmt.Println("Won since", registrationDate)
	if value.Cmp(zero) == 0 {
		value, _ = etherutils.StringToWei("0.01 ether")
	}
	fmt.Println("Locked value is", etherutils.WeiToString(value, true))
	fmt.Println("Highest bid was", etherutils.WeiToString(highestBid, true))

	// Deed
	deedContract, err := ens.DeedContract(client, &deedAddress)
	cli.ErrCheck(err, quiet, "Failed to obtain deed contract")
	// Deed owner
	deedOwner, err := ens.Owner(deedContract)
	cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
	deedOwnerName, _ := ens.ReverseResolve(client, &deedOwner)
	if deedOwnerName == "" {
		fmt.Println("Deed owner is", deedOwner.Hex())
	} else {
		fmt.Printf("Deed owner is %s (%s)\n", deedOwnerName, deedOwner.Hex())
	}
}

func ownedInfo(name string) {
	registrarContract, err := ens.AuctionRegistrarContract(client, ens.Tld(name))
	cli.ErrCheck(err, quiet, "Failed to obtain ENS registrar contract")
	_, deedAddress, registrationDate, value, highestBid, err := ens.Entry(registrarContract, client, name)
	if err == nil {
		fmt.Println("Owned since", registrationDate)
		fmt.Println("Locked value is", etherutils.WeiToString(value, true))
		fmt.Println("Highest bid was", etherutils.WeiToString(highestBid, true))

		// Deed
		deedContract, err := ens.DeedContract(client, &deedAddress)
		cli.ErrCheck(err, quiet, "Failed to obtain deed contract")
		// Deed owner
		deedOwner, err := deedContract.Owner(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
		deedOwnerName, _ := ens.ReverseResolve(client, &deedOwner)
		if deedOwnerName == "" {
			fmt.Println("Deed owner is", deedOwner.Hex())
		} else {
			fmt.Printf("Deed owner is %s (%s)\n", deedOwnerName, deedOwner.Hex())
		}

		previousDeedOwner, err := deedContract.PreviousOwner(nil)
		cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
		if bytes.Compare(previousDeedOwner.Bytes(), ens.UnknownAddress.Bytes()) != 0 {
			previousDeedOwnerName, _ := ens.ReverseResolve(client, &previousDeedOwner)
			if previousDeedOwnerName == "" {
				fmt.Println("Previous deed owner is", previousDeedOwner.Hex())
			} else {
				fmt.Printf("Previous deed owner is %s (%s)\n", previousDeedOwnerName, previousDeedOwner.Hex())
			}
		}
	}

	genericInfo(name)
}

func genericInfo(name string) {
	// Domain owner
	registry, err := ens.RegistryContract(client)
	cli.ErrCheck(err, quiet, "Failed to obtain registry contract")
	domainOwnerAddress, err := registry.Owner(nil, ens.NameHash(name))
	cli.ErrCheck(err, quiet, "Failed to obtain domain owner")
	if domainOwnerAddress == ens.UnknownAddress {
		fmt.Println("Domain owner not set")
		return
	}
	domainOwnerName, _ := ens.ReverseResolve(client, &domainOwnerAddress)
	if domainOwnerName == "" {
		fmt.Printf("Domain owner is %s\n", domainOwnerAddress.Hex())
	} else {
		fmt.Printf("Domain owner is %s (%s)\n", domainOwnerName, domainOwnerAddress.Hex())
	}

	// Resolver
	resolverAddress, err := ens.Resolver(registry, name)
	if err != nil {
		fmt.Println("Resolver not configured")
		return
	}
	resolverName, _ := ens.ReverseResolve(client, &resolverAddress)
	if resolverName == "" {
		fmt.Printf("Resolver is %s\n", resolverAddress.Hex())
	} else {
		fmt.Printf("Resolver is %s (%s)\n", resolverName, resolverAddress.Hex())
	}

	// Address
	address, err := ens.Resolve(client, name)
	if err == nil && address != ens.UnknownAddress {
		fmt.Printf("Domain resolves to %s\n", address.Hex())
		// Reverse resolution
		reverseDomain, err := ens.ReverseResolve(client, &address)
		if err == nil && reverseDomain != "" {
			fmt.Printf("Address resolves to %s\n", reverseDomain)
		}
	}

	// Content hash
	resolverContract, err := ens.ResolverContractByAddress(client, resolverAddress)
	if err == nil {
		bytes, err := resolverContract.Contenthash(nil, ens.NameHash(ensDomain))
		if err == nil && len(bytes) > 0 {
			contentHash, err := contenthashBytesToString(bytes)
			if err == nil {
				fmt.Printf("Content hash is %v\n", contentHash)
			}
		}
	}
}
