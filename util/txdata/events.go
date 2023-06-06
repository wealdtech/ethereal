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

package txdata

// initEventMap initialises the event map with event-specific signatures.
func initEventMap() {
	AddEventSignature("ABIChanged(bytes32,uint256)")
	AddEventSignature("AddrChanged(bytes32,address)")
	AddEventSignature("Allowance(address,address,address,uint256)")
	AddEventSignature("AllowanceSet(address,address,address,uint256)")
	AddEventSignature("Approval(address,address,uint256)")
	AddEventSignature("AuthorizedOperator(address,address)")
	AddEventSignature("Burn(address,uint256)")
	AddEventSignature("Burned(address,address,uint256,bytes,bytes)")
	AddEventSignature("Cleared(bytes32)")
	AddEventSignature("ContentChanged(bytes32,bytes32)")
	AddEventSignature("CounterSignatoryCleared(address)")
	AddEventSignature("CounterSignatorySet(address,address)")
	AddEventSignature("DeedClosed()")
	AddEventSignature("Deleted(bytes32,bytes,uint16)")
	AddEventSignature("DepositEvent(bytes,bytes,bytes,bytes,bytes)")
	AddEventSignature("ForwardingAddressCleared(address)")
	AddEventSignature("ForwardingAddressSet(address,address)")
	AddEventSignature("InterfaceImplementerSet(address,bytes32,address)")
	AddEventSignature("LockupExpires(address,address,uint256)")
	AddEventSignature("ManagerChanged(address,address)")
	AddEventSignature("Mark(bool)")
	AddEventSignature("Message(address,address,string)")
	AddEventSignature("MessageCleared(address,address)")
	AddEventSignature("MessageSet(address,address,string)")
	AddEventSignature("MinimumAmountCleared(address,address)")
	AddEventSignature("MinimumAmountSet(address,address,uint256)")
	AddEventSignature("Mint(address,uint256)")
	AddEventSignature("Minted(address,address,uint256,bytes,bytes)")
	AddEventSignature("NameChanged(bytes32,string)")
	AddEventSignature("NewInstance(uint256,uint256)")
	AddEventSignature("NewOwner(bytes32,bytes32,address)")
	AddEventSignature("NewResolver(bytes32,address)")
	AddEventSignature("NewTTL(bytes32,uint64")
	AddEventSignature("OwnerChanged(address)")
	AddEventSignature("Pause()")
	AddEventSignature("PausedUntil(uint256)")
	AddEventSignature("PermissionChanged(address,bytes32,bool)")
	AddEventSignature("PricePerToken(address,address,uint256)")
	AddEventSignature("PubkeyChanged(bytes32,bytes32,bytes32)")
	AddEventSignature("RecipientCleared(address,address)")
	AddEventSignature("RecipientSet(address,address)")
	AddEventSignature("Redirect(address)")
	AddEventSignature("ReleaseTimestamp(address,address,uint256)")
	AddEventSignature("RevokedOperator(address,address)")
	AddEventSignature("Root(address,address,bytes32)")
	AddEventSignature("Sent(address,address,address,uint256,bytes,bytes)")
	AddEventSignature("SupercededBy(address)")
	AddEventSignature("SupplementRemoved(address)")
	AddEventSignature("SupplementSet(address,address,uint16)")
	AddEventSignature("TextChanged(bytes32,string,string)")
	AddEventSignature("TokenAdded(address,address)")
	AddEventSignature("TokenRemoved(address,address)")
	AddEventSignature("Transfer(address,address,uint256)")
	AddEventSignature("Transfer(bytes32,address")
	AddEventSignature("Unpause()")
	AddEventSignature("Updated(bytes32,bytes,uint16)")
	AddEventSignature("Updated(bytes32,bytes,uint16,uint256)")
	AddEventSignature("Vault(address,address,address)")
}
