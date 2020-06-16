package contracts

//go:generate abigen -abi ERC20.abi -out erc20.go -pkg contracts -type ERC20
//go:generate abigen -abi eth2deposit.abi -out eth2deposit.go -pkg contracts -type Eth2Deposit
