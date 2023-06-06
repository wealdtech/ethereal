// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package funcparser

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TesterMetaData contains all meta data concerning the Tester contract.
var TesterMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string[]\"}],\"name\":\"testStringArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint256[][]\"}],\"name\":\"testUint2562DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint256[]\"}],\"name\":\"testUint256Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint8[]\"}],\"name\":\"testUint8Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256[]\"}],\"name\":\"testInt256Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int32[][]\"}],\"name\":\"testInt322DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int32[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool[]\"}],\"name\":\"testBoolArray\",\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256\"}],\"name\":\"testInt256\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint32[][]\"}],\"name\":\"testUint322DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int128[][]\"}],\"name\":\"testInt1282DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int128[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes\"}],\"name\":\"testBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool[][]\"}],\"name\":\"testBool2DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"bool[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int32[]\"}],\"name\":\"testInt32Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int32[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int256[][]\"}],\"name\":\"testInt2562DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int256[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int8[]\"}],\"name\":\"testInt8Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int8[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint64\"}],\"name\":\"testUint64\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"testAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint32\"}],\"name\":\"testUint32\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"testUint256\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes[][]\"}],\"name\":\"testBytes2DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint8\"}],\"name\":\"testUint8\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint128[][]\"}],\"name\":\"testUint1282DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int16\"}],\"name\":\"testInt16\",\"outputs\":[{\"name\":\"\",\"type\":\"int16\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint16\"}],\"name\":\"testUint16\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string\"}],\"name\":\"testString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int16[][]\"}],\"name\":\"testInt162DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int16[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int32\"}],\"name\":\"testInt32\",\"outputs\":[{\"name\":\"\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int64[][]\"}],\"name\":\"testInt642DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int64[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint16[]\"}],\"name\":\"testUint16Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"string[][]\"}],\"name\":\"testString2DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"string[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint64[][]\"}],\"name\":\"testUint642DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int64\"}],\"name\":\"testInt64\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint128\"}],\"name\":\"testUint128\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint128[]\"}],\"name\":\"testUint128Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address[]\"}],\"name\":\"testAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"address[][]\"}],\"name\":\"testAddress2DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bytes[]\"}],\"name\":\"testBytesArray\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint32[]\"}],\"name\":\"testUint32Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int8[][]\"}],\"name\":\"testInt82DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"int8[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int8\"}],\"name\":\"testInt8\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int16[]\"}],\"name\":\"testInt16Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int16[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint64[]\"}],\"name\":\"testUint64Array\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int128\"}],\"name\":\"testInt128\",\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int128[]\"}],\"name\":\"testInt128Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int128[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint16[][]\"}],\"name\":\"testUint162DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"uint8[][]\"}],\"name\":\"testUint82DArray\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8[][]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"bool\"}],\"name\":\"testBool\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arg1\",\"type\":\"int64[]\"}],\"name\":\"testInt64Array\",\"outputs\":[{\"name\":\"\",\"type\":\"int64[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// TesterABI is the input ABI used to generate the binding from.
// Deprecated: Use TesterMetaData.ABI instead.
var TesterABI = TesterMetaData.ABI

// Tester is an auto generated Go binding around an Ethereum contract.
type Tester struct {
	TesterCaller     // Read-only binding to the contract
	TesterTransactor // Write-only binding to the contract
	TesterFilterer   // Log filterer for contract events
}

// TesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TesterSession struct {
	Contract     *Tester           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TesterCallerSession struct {
	Contract *TesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TesterTransactorSession struct {
	Contract     *TesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TesterRaw struct {
	Contract *Tester // Generic contract binding to access the raw methods on
}

// TesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TesterCallerRaw struct {
	Contract *TesterCaller // Generic read-only contract binding to access the raw methods on
}

// TesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TesterTransactorRaw struct {
	Contract *TesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTester creates a new instance of Tester, bound to a specific deployed contract.
func NewTester(address common.Address, backend bind.ContractBackend) (*Tester, error) {
	contract, err := bindTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tester{TesterCaller: TesterCaller{contract: contract}, TesterTransactor: TesterTransactor{contract: contract}, TesterFilterer: TesterFilterer{contract: contract}}, nil
}

// NewTesterCaller creates a new read-only instance of Tester, bound to a specific deployed contract.
func NewTesterCaller(address common.Address, caller bind.ContractCaller) (*TesterCaller, error) {
	contract, err := bindTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TesterCaller{contract: contract}, nil
}

// NewTesterTransactor creates a new write-only instance of Tester, bound to a specific deployed contract.
func NewTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*TesterTransactor, error) {
	contract, err := bindTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TesterTransactor{contract: contract}, nil
}

// NewTesterFilterer creates a new log filterer instance of Tester, bound to a specific deployed contract.
func NewTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*TesterFilterer, error) {
	contract, err := bindTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TesterFilterer{contract: contract}, nil
}

// bindTester binds a generic wrapper to an already deployed contract.
func bindTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tester *TesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tester.Contract.TesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tester *TesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tester.Contract.TesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tester *TesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tester.Contract.TesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tester *TesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tester *TesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tester *TesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tester.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_Tester *TesterCaller) Test(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "test")

	if err != nil {
		return err
	}

	return err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_Tester *TesterSession) Test() error {
	return _Tester.Contract.Test(&_Tester.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_Tester *TesterCallerSession) Test() error {
	return _Tester.Contract.Test(&_Tester.CallOpts)
}

// TestAddress is a free data retrieval call binding the contract method 0x42f45790.
//
// Solidity: function testAddress(address arg1) pure returns(address)
func (_Tester *TesterCaller) TestAddress(opts *bind.CallOpts, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testAddress", arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TestAddress is a free data retrieval call binding the contract method 0x42f45790.
//
// Solidity: function testAddress(address arg1) pure returns(address)
func (_Tester *TesterSession) TestAddress(arg1 common.Address) (common.Address, error) {
	return _Tester.Contract.TestAddress(&_Tester.CallOpts, arg1)
}

// TestAddress is a free data retrieval call binding the contract method 0x42f45790.
//
// Solidity: function testAddress(address arg1) pure returns(address)
func (_Tester *TesterCallerSession) TestAddress(arg1 common.Address) (common.Address, error) {
	return _Tester.Contract.TestAddress(&_Tester.CallOpts, arg1)
}

// TestAddress2DArray is a free data retrieval call binding the contract method 0xd107e80c.
//
// Solidity: function testAddress2DArray(address[][] arg1) pure returns(address[][])
func (_Tester *TesterCaller) TestAddress2DArray(opts *bind.CallOpts, arg1 [][]common.Address) ([][]common.Address, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testAddress2DArray", arg1)

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// TestAddress2DArray is a free data retrieval call binding the contract method 0xd107e80c.
//
// Solidity: function testAddress2DArray(address[][] arg1) pure returns(address[][])
func (_Tester *TesterSession) TestAddress2DArray(arg1 [][]common.Address) ([][]common.Address, error) {
	return _Tester.Contract.TestAddress2DArray(&_Tester.CallOpts, arg1)
}

// TestAddress2DArray is a free data retrieval call binding the contract method 0xd107e80c.
//
// Solidity: function testAddress2DArray(address[][] arg1) pure returns(address[][])
func (_Tester *TesterCallerSession) TestAddress2DArray(arg1 [][]common.Address) ([][]common.Address, error) {
	return _Tester.Contract.TestAddress2DArray(&_Tester.CallOpts, arg1)
}

// TestAddressArray is a free data retrieval call binding the contract method 0xcef93f89.
//
// Solidity: function testAddressArray(address[] arg1) pure returns(address[])
func (_Tester *TesterCaller) TestAddressArray(opts *bind.CallOpts, arg1 []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testAddressArray", arg1)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TestAddressArray is a free data retrieval call binding the contract method 0xcef93f89.
//
// Solidity: function testAddressArray(address[] arg1) pure returns(address[])
func (_Tester *TesterSession) TestAddressArray(arg1 []common.Address) ([]common.Address, error) {
	return _Tester.Contract.TestAddressArray(&_Tester.CallOpts, arg1)
}

// TestAddressArray is a free data retrieval call binding the contract method 0xcef93f89.
//
// Solidity: function testAddressArray(address[] arg1) pure returns(address[])
func (_Tester *TesterCallerSession) TestAddressArray(arg1 []common.Address) ([]common.Address, error) {
	return _Tester.Contract.TestAddressArray(&_Tester.CallOpts, arg1)
}

// TestBool is a free data retrieval call binding the contract method 0xe8dde232.
//
// Solidity: function testBool(bool arg1) pure returns(bool)
func (_Tester *TesterCaller) TestBool(opts *bind.CallOpts, arg1 bool) (bool, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBool", arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestBool is a free data retrieval call binding the contract method 0xe8dde232.
//
// Solidity: function testBool(bool arg1) pure returns(bool)
func (_Tester *TesterSession) TestBool(arg1 bool) (bool, error) {
	return _Tester.Contract.TestBool(&_Tester.CallOpts, arg1)
}

// TestBool is a free data retrieval call binding the contract method 0xe8dde232.
//
// Solidity: function testBool(bool arg1) pure returns(bool)
func (_Tester *TesterCallerSession) TestBool(arg1 bool) (bool, error) {
	return _Tester.Contract.TestBool(&_Tester.CallOpts, arg1)
}

// TestBool2DArray is a free data retrieval call binding the contract method 0x3d2021b3.
//
// Solidity: function testBool2DArray(bool[][] arg1) pure returns(bool[][])
func (_Tester *TesterCaller) TestBool2DArray(opts *bind.CallOpts, arg1 [][]bool) ([][]bool, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBool2DArray", arg1)

	if err != nil {
		return *new([][]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([][]bool)).(*[][]bool)

	return out0, err

}

// TestBool2DArray is a free data retrieval call binding the contract method 0x3d2021b3.
//
// Solidity: function testBool2DArray(bool[][] arg1) pure returns(bool[][])
func (_Tester *TesterSession) TestBool2DArray(arg1 [][]bool) ([][]bool, error) {
	return _Tester.Contract.TestBool2DArray(&_Tester.CallOpts, arg1)
}

// TestBool2DArray is a free data retrieval call binding the contract method 0x3d2021b3.
//
// Solidity: function testBool2DArray(bool[][] arg1) pure returns(bool[][])
func (_Tester *TesterCallerSession) TestBool2DArray(arg1 [][]bool) ([][]bool, error) {
	return _Tester.Contract.TestBool2DArray(&_Tester.CallOpts, arg1)
}

// TestBoolArray is a free data retrieval call binding the contract method 0x1b248a2a.
//
// Solidity: function testBoolArray(bool[] arg1) pure returns(bool[])
func (_Tester *TesterCaller) TestBoolArray(opts *bind.CallOpts, arg1 []bool) ([]bool, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBoolArray", arg1)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// TestBoolArray is a free data retrieval call binding the contract method 0x1b248a2a.
//
// Solidity: function testBoolArray(bool[] arg1) pure returns(bool[])
func (_Tester *TesterSession) TestBoolArray(arg1 []bool) ([]bool, error) {
	return _Tester.Contract.TestBoolArray(&_Tester.CallOpts, arg1)
}

// TestBoolArray is a free data retrieval call binding the contract method 0x1b248a2a.
//
// Solidity: function testBoolArray(bool[] arg1) pure returns(bool[])
func (_Tester *TesterCallerSession) TestBoolArray(arg1 []bool) ([]bool, error) {
	return _Tester.Contract.TestBoolArray(&_Tester.CallOpts, arg1)
}

// TestBytes is a free data retrieval call binding the contract method 0x3ca8b1a7.
//
// Solidity: function testBytes(bytes arg1) pure returns(bytes)
func (_Tester *TesterCaller) TestBytes(opts *bind.CallOpts, arg1 []byte) ([]byte, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBytes", arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBytes is a free data retrieval call binding the contract method 0x3ca8b1a7.
//
// Solidity: function testBytes(bytes arg1) pure returns(bytes)
func (_Tester *TesterSession) TestBytes(arg1 []byte) ([]byte, error) {
	return _Tester.Contract.TestBytes(&_Tester.CallOpts, arg1)
}

// TestBytes is a free data retrieval call binding the contract method 0x3ca8b1a7.
//
// Solidity: function testBytes(bytes arg1) pure returns(bytes)
func (_Tester *TesterCallerSession) TestBytes(arg1 []byte) ([]byte, error) {
	return _Tester.Contract.TestBytes(&_Tester.CallOpts, arg1)
}

// TestBytes2DArray is a free data retrieval call binding the contract method 0x4ec5e44b.
//
// Solidity: function testBytes2DArray(bytes[][] arg1) pure returns(bytes[][])
func (_Tester *TesterCaller) TestBytes2DArray(opts *bind.CallOpts, arg1 [][][]byte) ([][][]byte, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBytes2DArray", arg1)

	if err != nil {
		return *new([][][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][][]byte)).(*[][][]byte)

	return out0, err

}

// TestBytes2DArray is a free data retrieval call binding the contract method 0x4ec5e44b.
//
// Solidity: function testBytes2DArray(bytes[][] arg1) pure returns(bytes[][])
func (_Tester *TesterSession) TestBytes2DArray(arg1 [][][]byte) ([][][]byte, error) {
	return _Tester.Contract.TestBytes2DArray(&_Tester.CallOpts, arg1)
}

// TestBytes2DArray is a free data retrieval call binding the contract method 0x4ec5e44b.
//
// Solidity: function testBytes2DArray(bytes[][] arg1) pure returns(bytes[][])
func (_Tester *TesterCallerSession) TestBytes2DArray(arg1 [][][]byte) ([][][]byte, error) {
	return _Tester.Contract.TestBytes2DArray(&_Tester.CallOpts, arg1)
}

// TestBytesArray is a free data retrieval call binding the contract method 0xd11ab5ef.
//
// Solidity: function testBytesArray(bytes[] arg1) pure returns(bytes[])
func (_Tester *TesterCaller) TestBytesArray(opts *bind.CallOpts, arg1 [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testBytesArray", arg1)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// TestBytesArray is a free data retrieval call binding the contract method 0xd11ab5ef.
//
// Solidity: function testBytesArray(bytes[] arg1) pure returns(bytes[])
func (_Tester *TesterSession) TestBytesArray(arg1 [][]byte) ([][]byte, error) {
	return _Tester.Contract.TestBytesArray(&_Tester.CallOpts, arg1)
}

// TestBytesArray is a free data retrieval call binding the contract method 0xd11ab5ef.
//
// Solidity: function testBytesArray(bytes[] arg1) pure returns(bytes[])
func (_Tester *TesterCallerSession) TestBytesArray(arg1 [][]byte) ([][]byte, error) {
	return _Tester.Contract.TestBytesArray(&_Tester.CallOpts, arg1)
}

// TestInt128 is a free data retrieval call binding the contract method 0xe281869d.
//
// Solidity: function testInt128(int128 arg1) pure returns(int128)
func (_Tester *TesterCaller) TestInt128(opts *bind.CallOpts, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt128", arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestInt128 is a free data retrieval call binding the contract method 0xe281869d.
//
// Solidity: function testInt128(int128 arg1) pure returns(int128)
func (_Tester *TesterSession) TestInt128(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestInt128(&_Tester.CallOpts, arg1)
}

// TestInt128 is a free data retrieval call binding the contract method 0xe281869d.
//
// Solidity: function testInt128(int128 arg1) pure returns(int128)
func (_Tester *TesterCallerSession) TestInt128(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestInt128(&_Tester.CallOpts, arg1)
}

// TestInt1282DArray is a free data retrieval call binding the contract method 0x399df78b.
//
// Solidity: function testInt1282DArray(int128[][] arg1) pure returns(int128[][])
func (_Tester *TesterCaller) TestInt1282DArray(opts *bind.CallOpts, arg1 [][]*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt1282DArray", arg1)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// TestInt1282DArray is a free data retrieval call binding the contract method 0x399df78b.
//
// Solidity: function testInt1282DArray(int128[][] arg1) pure returns(int128[][])
func (_Tester *TesterSession) TestInt1282DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestInt1282DArray(&_Tester.CallOpts, arg1)
}

// TestInt1282DArray is a free data retrieval call binding the contract method 0x399df78b.
//
// Solidity: function testInt1282DArray(int128[][] arg1) pure returns(int128[][])
func (_Tester *TesterCallerSession) TestInt1282DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestInt1282DArray(&_Tester.CallOpts, arg1)
}

// TestInt128Array is a free data retrieval call binding the contract method 0xe50f9f49.
//
// Solidity: function testInt128Array(int128[] arg1) pure returns(int128[])
func (_Tester *TesterCaller) TestInt128Array(opts *bind.CallOpts, arg1 []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt128Array", arg1)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TestInt128Array is a free data retrieval call binding the contract method 0xe50f9f49.
//
// Solidity: function testInt128Array(int128[] arg1) pure returns(int128[])
func (_Tester *TesterSession) TestInt128Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestInt128Array(&_Tester.CallOpts, arg1)
}

// TestInt128Array is a free data retrieval call binding the contract method 0xe50f9f49.
//
// Solidity: function testInt128Array(int128[] arg1) pure returns(int128[])
func (_Tester *TesterCallerSession) TestInt128Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestInt128Array(&_Tester.CallOpts, arg1)
}

// TestInt16 is a free data retrieval call binding the contract method 0x5cadb7f4.
//
// Solidity: function testInt16(int16 arg1) pure returns(int16)
func (_Tester *TesterCaller) TestInt16(opts *bind.CallOpts, arg1 int16) (int16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt16", arg1)

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// TestInt16 is a free data retrieval call binding the contract method 0x5cadb7f4.
//
// Solidity: function testInt16(int16 arg1) pure returns(int16)
func (_Tester *TesterSession) TestInt16(arg1 int16) (int16, error) {
	return _Tester.Contract.TestInt16(&_Tester.CallOpts, arg1)
}

// TestInt16 is a free data retrieval call binding the contract method 0x5cadb7f4.
//
// Solidity: function testInt16(int16 arg1) pure returns(int16)
func (_Tester *TesterCallerSession) TestInt16(arg1 int16) (int16, error) {
	return _Tester.Contract.TestInt16(&_Tester.CallOpts, arg1)
}

// TestInt162DArray is a free data retrieval call binding the contract method 0x7486fca1.
//
// Solidity: function testInt162DArray(int16[][] arg1) pure returns(int16[][])
func (_Tester *TesterCaller) TestInt162DArray(opts *bind.CallOpts, arg1 [][]int16) ([][]int16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt162DArray", arg1)

	if err != nil {
		return *new([][]int16), err
	}

	out0 := *abi.ConvertType(out[0], new([][]int16)).(*[][]int16)

	return out0, err

}

// TestInt162DArray is a free data retrieval call binding the contract method 0x7486fca1.
//
// Solidity: function testInt162DArray(int16[][] arg1) pure returns(int16[][])
func (_Tester *TesterSession) TestInt162DArray(arg1 [][]int16) ([][]int16, error) {
	return _Tester.Contract.TestInt162DArray(&_Tester.CallOpts, arg1)
}

// TestInt162DArray is a free data retrieval call binding the contract method 0x7486fca1.
//
// Solidity: function testInt162DArray(int16[][] arg1) pure returns(int16[][])
func (_Tester *TesterCallerSession) TestInt162DArray(arg1 [][]int16) ([][]int16, error) {
	return _Tester.Contract.TestInt162DArray(&_Tester.CallOpts, arg1)
}

// TestInt16Array is a free data retrieval call binding the contract method 0xdd8cc60a.
//
// Solidity: function testInt16Array(int16[] arg1) pure returns(int16[])
func (_Tester *TesterCaller) TestInt16Array(opts *bind.CallOpts, arg1 []int16) ([]int16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt16Array", arg1)

	if err != nil {
		return *new([]int16), err
	}

	out0 := *abi.ConvertType(out[0], new([]int16)).(*[]int16)

	return out0, err

}

// TestInt16Array is a free data retrieval call binding the contract method 0xdd8cc60a.
//
// Solidity: function testInt16Array(int16[] arg1) pure returns(int16[])
func (_Tester *TesterSession) TestInt16Array(arg1 []int16) ([]int16, error) {
	return _Tester.Contract.TestInt16Array(&_Tester.CallOpts, arg1)
}

// TestInt16Array is a free data retrieval call binding the contract method 0xdd8cc60a.
//
// Solidity: function testInt16Array(int16[] arg1) pure returns(int16[])
func (_Tester *TesterCallerSession) TestInt16Array(arg1 []int16) ([]int16, error) {
	return _Tester.Contract.TestInt16Array(&_Tester.CallOpts, arg1)
}

// TestInt256 is a free data retrieval call binding the contract method 0x24c97c60.
//
// Solidity: function testInt256(int256 arg1) pure returns(int256)
func (_Tester *TesterCaller) TestInt256(opts *bind.CallOpts, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt256", arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestInt256 is a free data retrieval call binding the contract method 0x24c97c60.
//
// Solidity: function testInt256(int256 arg1) pure returns(int256)
func (_Tester *TesterSession) TestInt256(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestInt256(&_Tester.CallOpts, arg1)
}

// TestInt256 is a free data retrieval call binding the contract method 0x24c97c60.
//
// Solidity: function testInt256(int256 arg1) pure returns(int256)
func (_Tester *TesterCallerSession) TestInt256(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestInt256(&_Tester.CallOpts, arg1)
}

// TestInt2562DArray is a free data retrieval call binding the contract method 0x3ddeec9a.
//
// Solidity: function testInt2562DArray(int256[][] arg1) pure returns(int256[][])
func (_Tester *TesterCaller) TestInt2562DArray(opts *bind.CallOpts, arg1 [][]*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt2562DArray", arg1)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// TestInt2562DArray is a free data retrieval call binding the contract method 0x3ddeec9a.
//
// Solidity: function testInt2562DArray(int256[][] arg1) pure returns(int256[][])
func (_Tester *TesterSession) TestInt2562DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestInt2562DArray(&_Tester.CallOpts, arg1)
}

// TestInt2562DArray is a free data retrieval call binding the contract method 0x3ddeec9a.
//
// Solidity: function testInt2562DArray(int256[][] arg1) pure returns(int256[][])
func (_Tester *TesterCallerSession) TestInt2562DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestInt2562DArray(&_Tester.CallOpts, arg1)
}

// TestInt256Array is a free data retrieval call binding the contract method 0x15e7b164.
//
// Solidity: function testInt256Array(int256[] arg1) pure returns(int256[])
func (_Tester *TesterCaller) TestInt256Array(opts *bind.CallOpts, arg1 []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt256Array", arg1)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TestInt256Array is a free data retrieval call binding the contract method 0x15e7b164.
//
// Solidity: function testInt256Array(int256[] arg1) pure returns(int256[])
func (_Tester *TesterSession) TestInt256Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestInt256Array(&_Tester.CallOpts, arg1)
}

// TestInt256Array is a free data retrieval call binding the contract method 0x15e7b164.
//
// Solidity: function testInt256Array(int256[] arg1) pure returns(int256[])
func (_Tester *TesterCallerSession) TestInt256Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestInt256Array(&_Tester.CallOpts, arg1)
}

// TestInt32 is a free data retrieval call binding the contract method 0x75c177bb.
//
// Solidity: function testInt32(int32 arg1) pure returns(int32)
func (_Tester *TesterCaller) TestInt32(opts *bind.CallOpts, arg1 int32) (int32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt32", arg1)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// TestInt32 is a free data retrieval call binding the contract method 0x75c177bb.
//
// Solidity: function testInt32(int32 arg1) pure returns(int32)
func (_Tester *TesterSession) TestInt32(arg1 int32) (int32, error) {
	return _Tester.Contract.TestInt32(&_Tester.CallOpts, arg1)
}

// TestInt32 is a free data retrieval call binding the contract method 0x75c177bb.
//
// Solidity: function testInt32(int32 arg1) pure returns(int32)
func (_Tester *TesterCallerSession) TestInt32(arg1 int32) (int32, error) {
	return _Tester.Contract.TestInt32(&_Tester.CallOpts, arg1)
}

// TestInt322DArray is a free data retrieval call binding the contract method 0x16f8f7a7.
//
// Solidity: function testInt322DArray(int32[][] arg1) pure returns(int32[][])
func (_Tester *TesterCaller) TestInt322DArray(opts *bind.CallOpts, arg1 [][]int32) ([][]int32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt322DArray", arg1)

	if err != nil {
		return *new([][]int32), err
	}

	out0 := *abi.ConvertType(out[0], new([][]int32)).(*[][]int32)

	return out0, err

}

// TestInt322DArray is a free data retrieval call binding the contract method 0x16f8f7a7.
//
// Solidity: function testInt322DArray(int32[][] arg1) pure returns(int32[][])
func (_Tester *TesterSession) TestInt322DArray(arg1 [][]int32) ([][]int32, error) {
	return _Tester.Contract.TestInt322DArray(&_Tester.CallOpts, arg1)
}

// TestInt322DArray is a free data retrieval call binding the contract method 0x16f8f7a7.
//
// Solidity: function testInt322DArray(int32[][] arg1) pure returns(int32[][])
func (_Tester *TesterCallerSession) TestInt322DArray(arg1 [][]int32) ([][]int32, error) {
	return _Tester.Contract.TestInt322DArray(&_Tester.CallOpts, arg1)
}

// TestInt32Array is a free data retrieval call binding the contract method 0x3d6cc7a0.
//
// Solidity: function testInt32Array(int32[] arg1) pure returns(int32[])
func (_Tester *TesterCaller) TestInt32Array(opts *bind.CallOpts, arg1 []int32) ([]int32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt32Array", arg1)

	if err != nil {
		return *new([]int32), err
	}

	out0 := *abi.ConvertType(out[0], new([]int32)).(*[]int32)

	return out0, err

}

// TestInt32Array is a free data retrieval call binding the contract method 0x3d6cc7a0.
//
// Solidity: function testInt32Array(int32[] arg1) pure returns(int32[])
func (_Tester *TesterSession) TestInt32Array(arg1 []int32) ([]int32, error) {
	return _Tester.Contract.TestInt32Array(&_Tester.CallOpts, arg1)
}

// TestInt32Array is a free data retrieval call binding the contract method 0x3d6cc7a0.
//
// Solidity: function testInt32Array(int32[] arg1) pure returns(int32[])
func (_Tester *TesterCallerSession) TestInt32Array(arg1 []int32) ([]int32, error) {
	return _Tester.Contract.TestInt32Array(&_Tester.CallOpts, arg1)
}

// TestInt64 is a free data retrieval call binding the contract method 0x9eedf2af.
//
// Solidity: function testInt64(int64 arg1) pure returns(int64)
func (_Tester *TesterCaller) TestInt64(opts *bind.CallOpts, arg1 int64) (int64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt64", arg1)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// TestInt64 is a free data retrieval call binding the contract method 0x9eedf2af.
//
// Solidity: function testInt64(int64 arg1) pure returns(int64)
func (_Tester *TesterSession) TestInt64(arg1 int64) (int64, error) {
	return _Tester.Contract.TestInt64(&_Tester.CallOpts, arg1)
}

// TestInt64 is a free data retrieval call binding the contract method 0x9eedf2af.
//
// Solidity: function testInt64(int64 arg1) pure returns(int64)
func (_Tester *TesterCallerSession) TestInt64(arg1 int64) (int64, error) {
	return _Tester.Contract.TestInt64(&_Tester.CallOpts, arg1)
}

// TestInt642DArray is a free data retrieval call binding the contract method 0x7e375a80.
//
// Solidity: function testInt642DArray(int64[][] arg1) pure returns(int64[][])
func (_Tester *TesterCaller) TestInt642DArray(opts *bind.CallOpts, arg1 [][]int64) ([][]int64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt642DArray", arg1)

	if err != nil {
		return *new([][]int64), err
	}

	out0 := *abi.ConvertType(out[0], new([][]int64)).(*[][]int64)

	return out0, err

}

// TestInt642DArray is a free data retrieval call binding the contract method 0x7e375a80.
//
// Solidity: function testInt642DArray(int64[][] arg1) pure returns(int64[][])
func (_Tester *TesterSession) TestInt642DArray(arg1 [][]int64) ([][]int64, error) {
	return _Tester.Contract.TestInt642DArray(&_Tester.CallOpts, arg1)
}

// TestInt642DArray is a free data retrieval call binding the contract method 0x7e375a80.
//
// Solidity: function testInt642DArray(int64[][] arg1) pure returns(int64[][])
func (_Tester *TesterCallerSession) TestInt642DArray(arg1 [][]int64) ([][]int64, error) {
	return _Tester.Contract.TestInt642DArray(&_Tester.CallOpts, arg1)
}

// TestInt64Array is a free data retrieval call binding the contract method 0xef3340d9.
//
// Solidity: function testInt64Array(int64[] arg1) pure returns(int64[])
func (_Tester *TesterCaller) TestInt64Array(opts *bind.CallOpts, arg1 []int64) ([]int64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt64Array", arg1)

	if err != nil {
		return *new([]int64), err
	}

	out0 := *abi.ConvertType(out[0], new([]int64)).(*[]int64)

	return out0, err

}

// TestInt64Array is a free data retrieval call binding the contract method 0xef3340d9.
//
// Solidity: function testInt64Array(int64[] arg1) pure returns(int64[])
func (_Tester *TesterSession) TestInt64Array(arg1 []int64) ([]int64, error) {
	return _Tester.Contract.TestInt64Array(&_Tester.CallOpts, arg1)
}

// TestInt64Array is a free data retrieval call binding the contract method 0xef3340d9.
//
// Solidity: function testInt64Array(int64[] arg1) pure returns(int64[])
func (_Tester *TesterCallerSession) TestInt64Array(arg1 []int64) ([]int64, error) {
	return _Tester.Contract.TestInt64Array(&_Tester.CallOpts, arg1)
}

// TestInt8 is a free data retrieval call binding the contract method 0xdaa572d3.
//
// Solidity: function testInt8(int8 arg1) pure returns(int8)
func (_Tester *TesterCaller) TestInt8(opts *bind.CallOpts, arg1 int8) (int8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt8", arg1)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// TestInt8 is a free data retrieval call binding the contract method 0xdaa572d3.
//
// Solidity: function testInt8(int8 arg1) pure returns(int8)
func (_Tester *TesterSession) TestInt8(arg1 int8) (int8, error) {
	return _Tester.Contract.TestInt8(&_Tester.CallOpts, arg1)
}

// TestInt8 is a free data retrieval call binding the contract method 0xdaa572d3.
//
// Solidity: function testInt8(int8 arg1) pure returns(int8)
func (_Tester *TesterCallerSession) TestInt8(arg1 int8) (int8, error) {
	return _Tester.Contract.TestInt8(&_Tester.CallOpts, arg1)
}

// TestInt82DArray is a free data retrieval call binding the contract method 0xd85a24f4.
//
// Solidity: function testInt82DArray(int8[][] arg1) pure returns(int8[][])
func (_Tester *TesterCaller) TestInt82DArray(opts *bind.CallOpts, arg1 [][]int8) ([][]int8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt82DArray", arg1)

	if err != nil {
		return *new([][]int8), err
	}

	out0 := *abi.ConvertType(out[0], new([][]int8)).(*[][]int8)

	return out0, err

}

// TestInt82DArray is a free data retrieval call binding the contract method 0xd85a24f4.
//
// Solidity: function testInt82DArray(int8[][] arg1) pure returns(int8[][])
func (_Tester *TesterSession) TestInt82DArray(arg1 [][]int8) ([][]int8, error) {
	return _Tester.Contract.TestInt82DArray(&_Tester.CallOpts, arg1)
}

// TestInt82DArray is a free data retrieval call binding the contract method 0xd85a24f4.
//
// Solidity: function testInt82DArray(int8[][] arg1) pure returns(int8[][])
func (_Tester *TesterCallerSession) TestInt82DArray(arg1 [][]int8) ([][]int8, error) {
	return _Tester.Contract.TestInt82DArray(&_Tester.CallOpts, arg1)
}

// TestInt8Array is a free data retrieval call binding the contract method 0x402d2301.
//
// Solidity: function testInt8Array(int8[] arg1) pure returns(int8[])
func (_Tester *TesterCaller) TestInt8Array(opts *bind.CallOpts, arg1 []int8) ([]int8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testInt8Array", arg1)

	if err != nil {
		return *new([]int8), err
	}

	out0 := *abi.ConvertType(out[0], new([]int8)).(*[]int8)

	return out0, err

}

// TestInt8Array is a free data retrieval call binding the contract method 0x402d2301.
//
// Solidity: function testInt8Array(int8[] arg1) pure returns(int8[])
func (_Tester *TesterSession) TestInt8Array(arg1 []int8) ([]int8, error) {
	return _Tester.Contract.TestInt8Array(&_Tester.CallOpts, arg1)
}

// TestInt8Array is a free data retrieval call binding the contract method 0x402d2301.
//
// Solidity: function testInt8Array(int8[] arg1) pure returns(int8[])
func (_Tester *TesterCallerSession) TestInt8Array(arg1 []int8) ([]int8, error) {
	return _Tester.Contract.TestInt8Array(&_Tester.CallOpts, arg1)
}

// TestString is a free data retrieval call binding the contract method 0x61cb5a01.
//
// Solidity: function testString(string arg1) pure returns(string)
func (_Tester *TesterCaller) TestString(opts *bind.CallOpts, arg1 string) (string, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testString", arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TestString is a free data retrieval call binding the contract method 0x61cb5a01.
//
// Solidity: function testString(string arg1) pure returns(string)
func (_Tester *TesterSession) TestString(arg1 string) (string, error) {
	return _Tester.Contract.TestString(&_Tester.CallOpts, arg1)
}

// TestString is a free data retrieval call binding the contract method 0x61cb5a01.
//
// Solidity: function testString(string arg1) pure returns(string)
func (_Tester *TesterCallerSession) TestString(arg1 string) (string, error) {
	return _Tester.Contract.TestString(&_Tester.CallOpts, arg1)
}

// TestString2DArray is a free data retrieval call binding the contract method 0x8cf729d6.
//
// Solidity: function testString2DArray(string[][] arg1) pure returns(string[][])
func (_Tester *TesterCaller) TestString2DArray(opts *bind.CallOpts, arg1 [][]string) ([][]string, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testString2DArray", arg1)

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// TestString2DArray is a free data retrieval call binding the contract method 0x8cf729d6.
//
// Solidity: function testString2DArray(string[][] arg1) pure returns(string[][])
func (_Tester *TesterSession) TestString2DArray(arg1 [][]string) ([][]string, error) {
	return _Tester.Contract.TestString2DArray(&_Tester.CallOpts, arg1)
}

// TestString2DArray is a free data retrieval call binding the contract method 0x8cf729d6.
//
// Solidity: function testString2DArray(string[][] arg1) pure returns(string[][])
func (_Tester *TesterCallerSession) TestString2DArray(arg1 [][]string) ([][]string, error) {
	return _Tester.Contract.TestString2DArray(&_Tester.CallOpts, arg1)
}

// TestStringArray is a free data retrieval call binding the contract method 0x0528a2e7.
//
// Solidity: function testStringArray(string[] arg1) pure returns(string[])
func (_Tester *TesterCaller) TestStringArray(opts *bind.CallOpts, arg1 []string) ([]string, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testStringArray", arg1)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TestStringArray is a free data retrieval call binding the contract method 0x0528a2e7.
//
// Solidity: function testStringArray(string[] arg1) pure returns(string[])
func (_Tester *TesterSession) TestStringArray(arg1 []string) ([]string, error) {
	return _Tester.Contract.TestStringArray(&_Tester.CallOpts, arg1)
}

// TestStringArray is a free data retrieval call binding the contract method 0x0528a2e7.
//
// Solidity: function testStringArray(string[] arg1) pure returns(string[])
func (_Tester *TesterCallerSession) TestStringArray(arg1 []string) ([]string, error) {
	return _Tester.Contract.TestStringArray(&_Tester.CallOpts, arg1)
}

// TestUint128 is a free data retrieval call binding the contract method 0xa2ca0245.
//
// Solidity: function testUint128(uint128 arg1) pure returns(uint128)
func (_Tester *TesterCaller) TestUint128(opts *bind.CallOpts, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint128", arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestUint128 is a free data retrieval call binding the contract method 0xa2ca0245.
//
// Solidity: function testUint128(uint128 arg1) pure returns(uint128)
func (_Tester *TesterSession) TestUint128(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestUint128(&_Tester.CallOpts, arg1)
}

// TestUint128 is a free data retrieval call binding the contract method 0xa2ca0245.
//
// Solidity: function testUint128(uint128 arg1) pure returns(uint128)
func (_Tester *TesterCallerSession) TestUint128(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestUint128(&_Tester.CallOpts, arg1)
}

// TestUint1282DArray is a free data retrieval call binding the contract method 0x5b6308a3.
//
// Solidity: function testUint1282DArray(uint128[][] arg1) pure returns(uint128[][])
func (_Tester *TesterCaller) TestUint1282DArray(opts *bind.CallOpts, arg1 [][]*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint1282DArray", arg1)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// TestUint1282DArray is a free data retrieval call binding the contract method 0x5b6308a3.
//
// Solidity: function testUint1282DArray(uint128[][] arg1) pure returns(uint128[][])
func (_Tester *TesterSession) TestUint1282DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestUint1282DArray(&_Tester.CallOpts, arg1)
}

// TestUint1282DArray is a free data retrieval call binding the contract method 0x5b6308a3.
//
// Solidity: function testUint1282DArray(uint128[][] arg1) pure returns(uint128[][])
func (_Tester *TesterCallerSession) TestUint1282DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestUint1282DArray(&_Tester.CallOpts, arg1)
}

// TestUint128Array is a free data retrieval call binding the contract method 0xb1b3e5bb.
//
// Solidity: function testUint128Array(uint128[] arg1) pure returns(uint128[])
func (_Tester *TesterCaller) TestUint128Array(opts *bind.CallOpts, arg1 []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint128Array", arg1)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TestUint128Array is a free data retrieval call binding the contract method 0xb1b3e5bb.
//
// Solidity: function testUint128Array(uint128[] arg1) pure returns(uint128[])
func (_Tester *TesterSession) TestUint128Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestUint128Array(&_Tester.CallOpts, arg1)
}

// TestUint128Array is a free data retrieval call binding the contract method 0xb1b3e5bb.
//
// Solidity: function testUint128Array(uint128[] arg1) pure returns(uint128[])
func (_Tester *TesterCallerSession) TestUint128Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestUint128Array(&_Tester.CallOpts, arg1)
}

// TestUint16 is a free data retrieval call binding the contract method 0x5f74190c.
//
// Solidity: function testUint16(uint16 arg1) pure returns(uint16)
func (_Tester *TesterCaller) TestUint16(opts *bind.CallOpts, arg1 uint16) (uint16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint16", arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TestUint16 is a free data retrieval call binding the contract method 0x5f74190c.
//
// Solidity: function testUint16(uint16 arg1) pure returns(uint16)
func (_Tester *TesterSession) TestUint16(arg1 uint16) (uint16, error) {
	return _Tester.Contract.TestUint16(&_Tester.CallOpts, arg1)
}

// TestUint16 is a free data retrieval call binding the contract method 0x5f74190c.
//
// Solidity: function testUint16(uint16 arg1) pure returns(uint16)
func (_Tester *TesterCallerSession) TestUint16(arg1 uint16) (uint16, error) {
	return _Tester.Contract.TestUint16(&_Tester.CallOpts, arg1)
}

// TestUint162DArray is a free data retrieval call binding the contract method 0xe6a6cd26.
//
// Solidity: function testUint162DArray(uint16[][] arg1) pure returns(uint16[][])
func (_Tester *TesterCaller) TestUint162DArray(opts *bind.CallOpts, arg1 [][]uint16) ([][]uint16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint162DArray", arg1)

	if err != nil {
		return *new([][]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([][]uint16)).(*[][]uint16)

	return out0, err

}

// TestUint162DArray is a free data retrieval call binding the contract method 0xe6a6cd26.
//
// Solidity: function testUint162DArray(uint16[][] arg1) pure returns(uint16[][])
func (_Tester *TesterSession) TestUint162DArray(arg1 [][]uint16) ([][]uint16, error) {
	return _Tester.Contract.TestUint162DArray(&_Tester.CallOpts, arg1)
}

// TestUint162DArray is a free data retrieval call binding the contract method 0xe6a6cd26.
//
// Solidity: function testUint162DArray(uint16[][] arg1) pure returns(uint16[][])
func (_Tester *TesterCallerSession) TestUint162DArray(arg1 [][]uint16) ([][]uint16, error) {
	return _Tester.Contract.TestUint162DArray(&_Tester.CallOpts, arg1)
}

// TestUint16Array is a free data retrieval call binding the contract method 0x86c7734d.
//
// Solidity: function testUint16Array(uint16[] arg1) pure returns(uint16[])
func (_Tester *TesterCaller) TestUint16Array(opts *bind.CallOpts, arg1 []uint16) ([]uint16, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint16Array", arg1)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// TestUint16Array is a free data retrieval call binding the contract method 0x86c7734d.
//
// Solidity: function testUint16Array(uint16[] arg1) pure returns(uint16[])
func (_Tester *TesterSession) TestUint16Array(arg1 []uint16) ([]uint16, error) {
	return _Tester.Contract.TestUint16Array(&_Tester.CallOpts, arg1)
}

// TestUint16Array is a free data retrieval call binding the contract method 0x86c7734d.
//
// Solidity: function testUint16Array(uint16[] arg1) pure returns(uint16[])
func (_Tester *TesterCallerSession) TestUint16Array(arg1 []uint16) ([]uint16, error) {
	return _Tester.Contract.TestUint16Array(&_Tester.CallOpts, arg1)
}

// TestUint256 is a free data retrieval call binding the contract method 0x4b6390b2.
//
// Solidity: function testUint256(uint256 arg1) pure returns(uint256)
func (_Tester *TesterCaller) TestUint256(opts *bind.CallOpts, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint256", arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestUint256 is a free data retrieval call binding the contract method 0x4b6390b2.
//
// Solidity: function testUint256(uint256 arg1) pure returns(uint256)
func (_Tester *TesterSession) TestUint256(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestUint256(&_Tester.CallOpts, arg1)
}

// TestUint256 is a free data retrieval call binding the contract method 0x4b6390b2.
//
// Solidity: function testUint256(uint256 arg1) pure returns(uint256)
func (_Tester *TesterCallerSession) TestUint256(arg1 *big.Int) (*big.Int, error) {
	return _Tester.Contract.TestUint256(&_Tester.CallOpts, arg1)
}

// TestUint2562DArray is a free data retrieval call binding the contract method 0x109662df.
//
// Solidity: function testUint2562DArray(uint256[][] arg1) pure returns(uint256[][])
func (_Tester *TesterCaller) TestUint2562DArray(opts *bind.CallOpts, arg1 [][]*big.Int) ([][]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint2562DArray", arg1)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// TestUint2562DArray is a free data retrieval call binding the contract method 0x109662df.
//
// Solidity: function testUint2562DArray(uint256[][] arg1) pure returns(uint256[][])
func (_Tester *TesterSession) TestUint2562DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestUint2562DArray(&_Tester.CallOpts, arg1)
}

// TestUint2562DArray is a free data retrieval call binding the contract method 0x109662df.
//
// Solidity: function testUint2562DArray(uint256[][] arg1) pure returns(uint256[][])
func (_Tester *TesterCallerSession) TestUint2562DArray(arg1 [][]*big.Int) ([][]*big.Int, error) {
	return _Tester.Contract.TestUint2562DArray(&_Tester.CallOpts, arg1)
}

// TestUint256Array is a free data retrieval call binding the contract method 0x10d26202.
//
// Solidity: function testUint256Array(uint256[] arg1) pure returns(uint256[])
func (_Tester *TesterCaller) TestUint256Array(opts *bind.CallOpts, arg1 []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint256Array", arg1)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TestUint256Array is a free data retrieval call binding the contract method 0x10d26202.
//
// Solidity: function testUint256Array(uint256[] arg1) pure returns(uint256[])
func (_Tester *TesterSession) TestUint256Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestUint256Array(&_Tester.CallOpts, arg1)
}

// TestUint256Array is a free data retrieval call binding the contract method 0x10d26202.
//
// Solidity: function testUint256Array(uint256[] arg1) pure returns(uint256[])
func (_Tester *TesterCallerSession) TestUint256Array(arg1 []*big.Int) ([]*big.Int, error) {
	return _Tester.Contract.TestUint256Array(&_Tester.CallOpts, arg1)
}

// TestUint32 is a free data retrieval call binding the contract method 0x44bcce7b.
//
// Solidity: function testUint32(uint32 arg1) pure returns(uint32)
func (_Tester *TesterCaller) TestUint32(opts *bind.CallOpts, arg1 uint32) (uint32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint32", arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TestUint32 is a free data retrieval call binding the contract method 0x44bcce7b.
//
// Solidity: function testUint32(uint32 arg1) pure returns(uint32)
func (_Tester *TesterSession) TestUint32(arg1 uint32) (uint32, error) {
	return _Tester.Contract.TestUint32(&_Tester.CallOpts, arg1)
}

// TestUint32 is a free data retrieval call binding the contract method 0x44bcce7b.
//
// Solidity: function testUint32(uint32 arg1) pure returns(uint32)
func (_Tester *TesterCallerSession) TestUint32(arg1 uint32) (uint32, error) {
	return _Tester.Contract.TestUint32(&_Tester.CallOpts, arg1)
}

// TestUint322DArray is a free data retrieval call binding the contract method 0x398a7bd2.
//
// Solidity: function testUint322DArray(uint32[][] arg1) pure returns(uint32[][])
func (_Tester *TesterCaller) TestUint322DArray(opts *bind.CallOpts, arg1 [][]uint32) ([][]uint32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint322DArray", arg1)

	if err != nil {
		return *new([][]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([][]uint32)).(*[][]uint32)

	return out0, err

}

// TestUint322DArray is a free data retrieval call binding the contract method 0x398a7bd2.
//
// Solidity: function testUint322DArray(uint32[][] arg1) pure returns(uint32[][])
func (_Tester *TesterSession) TestUint322DArray(arg1 [][]uint32) ([][]uint32, error) {
	return _Tester.Contract.TestUint322DArray(&_Tester.CallOpts, arg1)
}

// TestUint322DArray is a free data retrieval call binding the contract method 0x398a7bd2.
//
// Solidity: function testUint322DArray(uint32[][] arg1) pure returns(uint32[][])
func (_Tester *TesterCallerSession) TestUint322DArray(arg1 [][]uint32) ([][]uint32, error) {
	return _Tester.Contract.TestUint322DArray(&_Tester.CallOpts, arg1)
}

// TestUint32Array is a free data retrieval call binding the contract method 0xd52a52d4.
//
// Solidity: function testUint32Array(uint32[] arg1) pure returns(uint32[])
func (_Tester *TesterCaller) TestUint32Array(opts *bind.CallOpts, arg1 []uint32) ([]uint32, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint32Array", arg1)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// TestUint32Array is a free data retrieval call binding the contract method 0xd52a52d4.
//
// Solidity: function testUint32Array(uint32[] arg1) pure returns(uint32[])
func (_Tester *TesterSession) TestUint32Array(arg1 []uint32) ([]uint32, error) {
	return _Tester.Contract.TestUint32Array(&_Tester.CallOpts, arg1)
}

// TestUint32Array is a free data retrieval call binding the contract method 0xd52a52d4.
//
// Solidity: function testUint32Array(uint32[] arg1) pure returns(uint32[])
func (_Tester *TesterCallerSession) TestUint32Array(arg1 []uint32) ([]uint32, error) {
	return _Tester.Contract.TestUint32Array(&_Tester.CallOpts, arg1)
}

// TestUint64 is a free data retrieval call binding the contract method 0x414a3ba9.
//
// Solidity: function testUint64(uint64 arg1) pure returns(uint64)
func (_Tester *TesterCaller) TestUint64(opts *bind.CallOpts, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint64", arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TestUint64 is a free data retrieval call binding the contract method 0x414a3ba9.
//
// Solidity: function testUint64(uint64 arg1) pure returns(uint64)
func (_Tester *TesterSession) TestUint64(arg1 uint64) (uint64, error) {
	return _Tester.Contract.TestUint64(&_Tester.CallOpts, arg1)
}

// TestUint64 is a free data retrieval call binding the contract method 0x414a3ba9.
//
// Solidity: function testUint64(uint64 arg1) pure returns(uint64)
func (_Tester *TesterCallerSession) TestUint64(arg1 uint64) (uint64, error) {
	return _Tester.Contract.TestUint64(&_Tester.CallOpts, arg1)
}

// TestUint642DArray is a free data retrieval call binding the contract method 0x914ab81d.
//
// Solidity: function testUint642DArray(uint64[][] arg1) pure returns(uint64[][])
func (_Tester *TesterCaller) TestUint642DArray(opts *bind.CallOpts, arg1 [][]uint64) ([][]uint64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint642DArray", arg1)

	if err != nil {
		return *new([][]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([][]uint64)).(*[][]uint64)

	return out0, err

}

// TestUint642DArray is a free data retrieval call binding the contract method 0x914ab81d.
//
// Solidity: function testUint642DArray(uint64[][] arg1) pure returns(uint64[][])
func (_Tester *TesterSession) TestUint642DArray(arg1 [][]uint64) ([][]uint64, error) {
	return _Tester.Contract.TestUint642DArray(&_Tester.CallOpts, arg1)
}

// TestUint642DArray is a free data retrieval call binding the contract method 0x914ab81d.
//
// Solidity: function testUint642DArray(uint64[][] arg1) pure returns(uint64[][])
func (_Tester *TesterCallerSession) TestUint642DArray(arg1 [][]uint64) ([][]uint64, error) {
	return _Tester.Contract.TestUint642DArray(&_Tester.CallOpts, arg1)
}

// TestUint64Array is a free data retrieval call binding the contract method 0xdda3a5c5.
//
// Solidity: function testUint64Array(uint64[] arg1) pure returns(uint64[])
func (_Tester *TesterCaller) TestUint64Array(opts *bind.CallOpts, arg1 []uint64) ([]uint64, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint64Array", arg1)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// TestUint64Array is a free data retrieval call binding the contract method 0xdda3a5c5.
//
// Solidity: function testUint64Array(uint64[] arg1) pure returns(uint64[])
func (_Tester *TesterSession) TestUint64Array(arg1 []uint64) ([]uint64, error) {
	return _Tester.Contract.TestUint64Array(&_Tester.CallOpts, arg1)
}

// TestUint64Array is a free data retrieval call binding the contract method 0xdda3a5c5.
//
// Solidity: function testUint64Array(uint64[] arg1) pure returns(uint64[])
func (_Tester *TesterCallerSession) TestUint64Array(arg1 []uint64) ([]uint64, error) {
	return _Tester.Contract.TestUint64Array(&_Tester.CallOpts, arg1)
}

// TestUint8 is a free data retrieval call binding the contract method 0x4fdba6a0.
//
// Solidity: function testUint8(uint8 arg1) pure returns(uint8)
func (_Tester *TesterCaller) TestUint8(opts *bind.CallOpts, arg1 uint8) (uint8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint8", arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TestUint8 is a free data retrieval call binding the contract method 0x4fdba6a0.
//
// Solidity: function testUint8(uint8 arg1) pure returns(uint8)
func (_Tester *TesterSession) TestUint8(arg1 uint8) (uint8, error) {
	return _Tester.Contract.TestUint8(&_Tester.CallOpts, arg1)
}

// TestUint8 is a free data retrieval call binding the contract method 0x4fdba6a0.
//
// Solidity: function testUint8(uint8 arg1) pure returns(uint8)
func (_Tester *TesterCallerSession) TestUint8(arg1 uint8) (uint8, error) {
	return _Tester.Contract.TestUint8(&_Tester.CallOpts, arg1)
}

// TestUint82DArray is a free data retrieval call binding the contract method 0xe85c0582.
//
// Solidity: function testUint82DArray(uint8[][] arg1) pure returns(uint8[][])
func (_Tester *TesterCaller) TestUint82DArray(opts *bind.CallOpts, arg1 [][]uint8) ([][]uint8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint82DArray", arg1)

	if err != nil {
		return *new([][]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([][]uint8)).(*[][]uint8)

	return out0, err

}

// TestUint82DArray is a free data retrieval call binding the contract method 0xe85c0582.
//
// Solidity: function testUint82DArray(uint8[][] arg1) pure returns(uint8[][])
func (_Tester *TesterSession) TestUint82DArray(arg1 [][]uint8) ([][]uint8, error) {
	return _Tester.Contract.TestUint82DArray(&_Tester.CallOpts, arg1)
}

// TestUint82DArray is a free data retrieval call binding the contract method 0xe85c0582.
//
// Solidity: function testUint82DArray(uint8[][] arg1) pure returns(uint8[][])
func (_Tester *TesterCallerSession) TestUint82DArray(arg1 [][]uint8) ([][]uint8, error) {
	return _Tester.Contract.TestUint82DArray(&_Tester.CallOpts, arg1)
}

// TestUint8Array is a free data retrieval call binding the contract method 0x11324860.
//
// Solidity: function testUint8Array(uint8[] arg1) pure returns(uint8[])
func (_Tester *TesterCaller) TestUint8Array(opts *bind.CallOpts, arg1 []uint8) ([]uint8, error) {
	var out []interface{}
	err := _Tester.contract.Call(opts, &out, "testUint8Array", arg1)

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// TestUint8Array is a free data retrieval call binding the contract method 0x11324860.
//
// Solidity: function testUint8Array(uint8[] arg1) pure returns(uint8[])
func (_Tester *TesterSession) TestUint8Array(arg1 []uint8) ([]uint8, error) {
	return _Tester.Contract.TestUint8Array(&_Tester.CallOpts, arg1)
}

// TestUint8Array is a free data retrieval call binding the contract method 0x11324860.
//
// Solidity: function testUint8Array(uint8[] arg1) pure returns(uint8[])
func (_Tester *TesterCallerSession) TestUint8Array(arg1 []uint8) ([]uint8, error) {
	return _Tester.Contract.TestUint8Array(&_Tester.CallOpts, arg1)
}
