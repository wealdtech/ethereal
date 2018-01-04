// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnsresolvercontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DnsResolverContractABI is the input ABI used to generate the binding from.
const DnsResolverContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"rr\",\"type\":\"uint16\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDns\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dnsRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"rr\",\"type\":\"uint16\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"dns\",\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

// DnsResolverContract is an auto generated Go binding around an Ethereum contract.
type DnsResolverContract struct {
	DnsResolverContractCaller     // Read-only binding to the contract
	DnsResolverContractTransactor // Write-only binding to the contract
}

// DnsResolverContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsResolverContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsResolverContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsResolverContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsResolverContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsResolverContractSession struct {
	Contract     *DnsResolverContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DnsResolverContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsResolverContractCallerSession struct {
	Contract *DnsResolverContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DnsResolverContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsResolverContractTransactorSession struct {
	Contract     *DnsResolverContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DnsResolverContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsResolverContractRaw struct {
	Contract *DnsResolverContract // Generic contract binding to access the raw methods on
}

// DnsResolverContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsResolverContractCallerRaw struct {
	Contract *DnsResolverContractCaller // Generic read-only contract binding to access the raw methods on
}

// DnsResolverContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsResolverContractTransactorRaw struct {
	Contract *DnsResolverContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsResolverContract creates a new instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContract(address common.Address, backend bind.ContractBackend) (*DnsResolverContract, error) {
	contract, err := bindDnsResolverContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContract{DnsResolverContractCaller: DnsResolverContractCaller{contract: contract}, DnsResolverContractTransactor: DnsResolverContractTransactor{contract: contract}}, nil
}

// NewDnsResolverContractCaller creates a new read-only instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContractCaller(address common.Address, caller bind.ContractCaller) (*DnsResolverContractCaller, error) {
	contract, err := bindDnsResolverContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractCaller{contract: contract}, nil
}

// NewDnsResolverContractTransactor creates a new write-only instance of DnsResolverContract, bound to a specific deployed contract.
func NewDnsResolverContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsResolverContractTransactor, error) {
	contract, err := bindDnsResolverContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DnsResolverContractTransactor{contract: contract}, nil
}

// bindDnsResolverContract binds a generic wrapper to an already deployed contract.
func bindDnsResolverContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsResolverContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsResolverContract *DnsResolverContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DnsResolverContract.Contract.DnsResolverContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsResolverContract *DnsResolverContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.DnsResolverContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsResolverContract *DnsResolverContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.DnsResolverContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsResolverContract *DnsResolverContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DnsResolverContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsResolverContract *DnsResolverContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsResolverContract *DnsResolverContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _DnsResolverContract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DnsResolverContract.Contract.ABI(&_DnsResolverContract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DnsResolverContract.Contract.ABI(&_DnsResolverContract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractSession) Addr(node [32]byte) (common.Address, error) {
	return _DnsResolverContract.Contract.Addr(&_DnsResolverContract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DnsResolverContract *DnsResolverContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _DnsResolverContract.Contract.Addr(&_DnsResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractSession) Content(node [32]byte) ([32]byte, error) {
	return _DnsResolverContract.Contract.Content(&_DnsResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DnsResolverContract *DnsResolverContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _DnsResolverContract.Contract.Content(&_DnsResolverContract.CallOpts, node)
}

// Dns is a free data retrieval call binding the contract method 0xaf6e6e9e.
//
// Solidity: function dns(node bytes32, rr uint16, key string) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCaller) Dns(opts *bind.CallOpts, node [32]byte, rr uint16, key string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "dns", node, rr, key)
	return *ret0, err
}

// Dns is a free data retrieval call binding the contract method 0xaf6e6e9e.
//
// Solidity: function dns(node bytes32, rr uint16, key string) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractSession) Dns(node [32]byte, rr uint16, key string) ([]byte, error) {
	return _DnsResolverContract.Contract.Dns(&_DnsResolverContract.CallOpts, node, rr, key)
}

// Dns is a free data retrieval call binding the contract method 0xaf6e6e9e.
//
// Solidity: function dns(node bytes32, rr uint16, key string) constant returns(data bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) Dns(node [32]byte, rr uint16, key string) ([]byte, error) {
	return _DnsResolverContract.Contract.Dns(&_DnsResolverContract.CallOpts, node, rr, key)
}

// DnsRecords is a free data retrieval call binding the contract method 0x75333d28.
//
// Solidity: function dnsRecords( bytes32,  uint16,  bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCaller) DnsRecords(opts *bind.CallOpts, arg0 [32]byte, arg1 uint16, arg2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "dnsRecords", arg0, arg1, arg2)
	return *ret0, err
}

// DnsRecords is a free data retrieval call binding the contract method 0x75333d28.
//
// Solidity: function dnsRecords( bytes32,  uint16,  bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractSession) DnsRecords(arg0 [32]byte, arg1 uint16, arg2 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsRecords(&_DnsResolverContract.CallOpts, arg0, arg1, arg2)
}

// DnsRecords is a free data retrieval call binding the contract method 0x75333d28.
//
// Solidity: function dnsRecords( bytes32,  uint16,  bytes32) constant returns(bytes)
func (_DnsResolverContract *DnsResolverContractCallerSession) DnsRecords(arg0 [32]byte, arg1 uint16, arg2 [32]byte) ([]byte, error) {
	return _DnsResolverContract.Contract.DnsRecords(&_DnsResolverContract.CallOpts, arg0, arg1, arg2)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractSession) Name(node [32]byte) (string, error) {
	return _DnsResolverContract.Contract.Name(&_DnsResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCallerSession) Name(node [32]byte) (string, error) {
	return _DnsResolverContract.Contract.Name(&_DnsResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _DnsResolverContract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DnsResolverContract.Contract.Pubkey(&_DnsResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DnsResolverContract *DnsResolverContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DnsResolverContract.Contract.Pubkey(&_DnsResolverContract.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsResolverContract.Contract.SupportsInterface(&_DnsResolverContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DnsResolverContract *DnsResolverContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsResolverContract.Contract.SupportsInterface(&_DnsResolverContract.CallOpts, interfaceId)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DnsResolverContract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractSession) Text(node [32]byte, key string) (string, error) {
	return _DnsResolverContract.Contract.Text(&_DnsResolverContract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DnsResolverContract *DnsResolverContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _DnsResolverContract.Contract.Text(&_DnsResolverContract.CallOpts, node, key)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetABI(&_DnsResolverContract.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetABI(&_DnsResolverContract.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetAddr(&_DnsResolverContract.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetAddr(&_DnsResolverContract.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetContent(&_DnsResolverContract.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetContent(&_DnsResolverContract.TransactOpts, node, hash)
}

// SetDns is a paid mutator transaction binding the contract method 0x6cf433e7.
//
// Solidity: function setDns(node bytes32, rr uint16, key string, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetDns(opts *bind.TransactOpts, node [32]byte, rr uint16, key string, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setDns", node, rr, key, data)
}

// SetDns is a paid mutator transaction binding the contract method 0x6cf433e7.
//
// Solidity: function setDns(node bytes32, rr uint16, key string, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetDns(node [32]byte, rr uint16, key string, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDns(&_DnsResolverContract.TransactOpts, node, rr, key, data)
}

// SetDns is a paid mutator transaction binding the contract method 0x6cf433e7.
//
// Solidity: function setDns(node bytes32, rr uint16, key string, data bytes) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetDns(node [32]byte, rr uint16, key string, data []byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetDns(&_DnsResolverContract.TransactOpts, node, rr, key, data)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetName(&_DnsResolverContract.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetName(&_DnsResolverContract.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetPubkey(&_DnsResolverContract.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetPubkey(&_DnsResolverContract.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetText(&_DnsResolverContract.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_DnsResolverContract *DnsResolverContractTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _DnsResolverContract.Contract.SetText(&_DnsResolverContract.TransactOpts, node, key, value)
}
