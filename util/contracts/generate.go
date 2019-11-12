package contracts

//go:generate abigen -abi ERC20.abi -out ERC20.go -pkg contracts -type ERC20
//go:generate abigen -abi eth2deposit.abi -out Eth2Deposit.go -pkg contracts -type Eth2Deposit
