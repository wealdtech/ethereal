package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type logDefinition struct {
	Signature string
}

var logDefinitions = []logDefinition{
	logDefinition{Signature: "Transfer(address,address,uint256)"},
	logDefinition{Signature: "Approval(address,address,uint256)"},
	logDefinition{Signature: "Mint(address,uint256)"},
	logDefinition{Signature: "Burn(address,uint256)"},
}

// Map of common log entries
var LogDefinitions map[common.Hash]*logDefinition

func InitLogDefinitions() {
	LogDefinitions = make(map[common.Hash]*logDefinition, 0)
	for _, logDefinition := range logDefinitions {
		sha := sha3.NewKeccak256()
		sha.Write([]byte(logDefinition.Signature))
		sig := common.BytesToHash(sha.Sum(nil))
		LogDefinitions[sig] = &logDefinition
	}
}
