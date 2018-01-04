package event

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// ReadString reads a string at a given position in the log
func ReadString(log *types.Log, position int) (result string, err error) {
	// Fetch the offset of the string in the data
	offset := binary.BigEndian.Uint64(log.Data[32*position+24 : 32*(position+1)])
	// Go to the offset and find the length of the string
	length := binary.BigEndian.Uint64(log.Data[offset+24 : offset+32])
	// Fetch the string itself
	result = string(log.Data[offset+32 : offset+32+length])

	return result, nil
}

// ReadInt reads an integer at a given position in the log
func ReadInt(log *types.Log, position int) (result *big.Int) {
	result = big.NewInt(0)
	result.SetBytes(log.Data[32*position : 32*(position+1)])
	return
}
