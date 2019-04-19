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

	"github.com/spf13/cobra"
	"github.com/wealdtech/ethereal/cli"
	ens "github.com/wealdtech/go-ens/v2"
	string2eth "github.com/wealdtech/go-string2eth"
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
		registrar, err := ens.NewBaseRegistrar(client, ens.Tld(ensDomain))
		cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain ENS registrar contract for %s", ens.Tld(ensDomain)))

		outputIf(verbose, fmt.Sprintf("Domain level is %v", ens.DomainLevel(ensDomain)))
		outputIf(verbose, fmt.Sprintf("Name hash is 0x%x", ens.NameHash(ensDomain)))

		registry, err := ens.NewRegistry(client)
		cli.ErrCheck(err, quiet, "Failed to obtain registry contract")

		// Work out if this is on the old or new registrar
		location, err := registrar.RegisteredWith(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to obtain domain location")
		switch location {
		case "none":
			outputIf(!quiet, "Domain not registered")
			os.Exit(_exit_failure)
		case "temporary":
			outputIf(!quiet, "Domain registered with temporary registrar")
		case "permanent":
			outputIf(verbose, "Domain registered on permanent registrar")
		default:
			cli.Err(quiet, fmt.Sprintf("Unexpected domain location %s", location))
		}

		domainOwnerAddress, err := registry.Owner(ensDomain)
		cli.ErrCheck(err, quiet, "Failed to obtain domain owner")
		if domainOwnerAddress != ens.UnknownAddress {
			domainOwnerName, _ := ens.ReverseResolve(client, domainOwnerAddress)
			if domainOwnerName == "" {
				fmt.Printf("Owner is %s\n", domainOwnerAddress.Hex())
			} else {
				fmt.Printf("Owner is %s (%s)\n", domainOwnerName, domainOwnerAddress.Hex())
			}
		}

		if location == "permanent" {
			domain, err := ens.DomainPart(ensDomain, 1)
			registrant, err := registrar.Owner(domain)
			cli.ErrCheck(err, quiet, "Failed to obtain registrant")
			registrantName, _ := ens.ReverseResolve(client, registrant)
			if registrantName == "" {
				fmt.Printf("Registrant is %s\n", registrant.Hex())
			} else {
				fmt.Printf("Registrant is %s (%s)\n", registrantName, registrant.Hex())
			}
			genericInfo(ensDomain)
			expiry, err := registrar.Expiry(domain)
			cli.ErrCheck(err, quiet, "Failed to obtain expiry")
			fmt.Printf("Registration expires at %v\n", time.Unix(int64(expiry.Uint64()), 0))
		}
		if location == "temporary" {
			auctionRegistrar, err := registrar.PriorAuctionContract()
			cli.ErrCheck(err, quiet, fmt.Sprintf("Failed to obtain auction registrar contract for %s", ens.Tld(ensDomain)))
			if ens.DomainLevel(ensDomain) == 1 {
				state, err := auctionRegistrar.State(ensDomain)
				cli.ErrCheck(err, quiet, "Failed to obtain domain state")

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
							biddingInfo(auctionRegistrar, ensDomain)
						case "Revealing":
							revealingInfo(auctionRegistrar, ensDomain)
						case "Won":
							wonInfo(auctionRegistrar, ensDomain)
						case "Owned":
							ownedInfo(auctionRegistrar, ensDomain)
						default:
							fmt.Println(state)
						}
					}
				} else {
					ownedInfo(auctionRegistrar, ensDomain)
				}
			} else {
				genericInfo(ensDomain)
			}
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

func biddingInfo(registrar *ens.AuctionRegistrar, name string) {
	entry, err := registrar.Entry(name)
	cli.ErrCheck(err, quiet, "Cannot obtain information for that name")

	twoDaysAgo := time.Duration(-48) * time.Hour
	fmt.Println("Bidding until", entry.Registration.Add(twoDaysAgo))
}

func revealingInfo(registrar *ens.AuctionRegistrar, name string) {
	entry, err := registrar.Entry(name)
	cli.ErrCheck(err, quiet, "Cannot obtain information for that name")

	fmt.Println("Revealing until", entry.Registration)
	// If the value is 0 then it is is minvalue instead
	if entry.Value.Cmp(zero) == 0 {
		entry.Value, _ = string2eth.StringToWei("0.01 ether")
	}
	fmt.Println("Locked value is", string2eth.WeiToString(entry.Value, true))
	fmt.Println("Highest bid is", string2eth.WeiToString(entry.HighestBid, true))
}

func wonInfo(registrar *ens.AuctionRegistrar, name string) {
	entry, err := registrar.Entry(name)
	cli.ErrCheck(err, quiet, "Cannot obtain information for that name")

	fmt.Println("Won since", entry.Registration)
	if entry.Value.Cmp(zero) == 0 {
		entry.Value, _ = string2eth.StringToWei("0.01 ether")
	}
	fmt.Println("Locked value is", string2eth.WeiToString(entry.Value, true))
	fmt.Println("Highest bid was", string2eth.WeiToString(entry.HighestBid, true))

	// Deed
	deed, err := ens.NewDeedAt(client, entry.Deed)
	cli.ErrCheck(err, quiet, "Failed to obtain deed contract")
	// Deed owner
	deedOwner, err := deed.Owner()
	cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
	deedOwnerName, _ := ens.ReverseResolve(client, deedOwner)
	if deedOwnerName == "" {
		fmt.Println("Deed owner is", deedOwner.Hex())
	} else {
		fmt.Printf("Deed owner is %s (%s)\n", deedOwnerName, deedOwner.Hex())
	}
}

func ownedInfo(registrar *ens.AuctionRegistrar, name string) {
	entry, err := registrar.Entry(name)
	if err == nil {
		fmt.Println("Owned since", entry.Registration)
		fmt.Println("Locked value is", string2eth.WeiToString(entry.Value, true))
		fmt.Println("Highest bid was", string2eth.WeiToString(entry.HighestBid, true))

		// Deed
		deed, err := ens.NewDeedAt(client, entry.Deed)
		cli.ErrCheck(err, quiet, "Failed to obtain deed contract")
		// Deed owner
		deedOwner, err := deed.Owner()
		cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
		deedOwnerName, _ := ens.ReverseResolve(client, deedOwner)
		if deedOwnerName == "" {
			fmt.Println("Deed owner is", deedOwner.Hex())
		} else {
			fmt.Printf("Deed owner is %s (%s)\n", deedOwnerName, deedOwner.Hex())
		}

		previousDeedOwner, err := deed.PreviousOwner()
		cli.ErrCheck(err, quiet, "Failed to obtain deed owner")
		if bytes.Compare(previousDeedOwner.Bytes(), ens.UnknownAddress.Bytes()) != 0 {
			previousDeedOwnerName, _ := ens.ReverseResolve(client, previousDeedOwner)
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
	registry, err := ens.NewRegistry(client)
	cli.ErrCheck(err, quiet, "Failed to obtain registry contract")
	domainOwnerAddress, err := registry.Owner(name)
	cli.ErrCheck(err, quiet, "Failed to obtain domain owner")
	if domainOwnerAddress == ens.UnknownAddress {
		fmt.Println("Owner not set")
		return
	}

	// Resolver
	resolverAddress, err := registry.ResolverAddress(name)
	if err != nil || resolverAddress == ens.UnknownAddress {
		fmt.Println("Resolver not configured")
		return
	}
	resolverName, _ := ens.ReverseResolve(client, resolverAddress)
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
		reverseDomain, err := ens.ReverseResolve(client, address)
		if err == nil && reverseDomain != "" {
			fmt.Printf("Address resolves to %s\n", reverseDomain)
		}
	}

	// Content hash
	resolver, err := ens.NewResolverAt(client, name, resolverAddress)
	if err == nil {
		bytes, err := resolver.Contenthash()
		if err == nil && len(bytes) > 0 {
			contentHash, err := contenthashBytesToString(bytes)
			if err == nil {
				fmt.Printf("Content hash is %v\n", contentHash)
			}
		}
	}
}
