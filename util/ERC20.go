package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/ethereal/util/contracts"
)

func ERC20Contract(client *ethclient.Client, address common.Address) (contract *contracts.ERC20, err error) {
	contract, err = contracts.NewERC20(address, client)
	return
}
