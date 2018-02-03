// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package reverseregistrarcontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReverseRegistrarContractABI is the input ABI used to generate the binding from.
const ReverseRegistrarContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithResolver\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"},{\"name\":\"resolverAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ReverseRegistrarContract is an auto generated Go binding around an Ethereum contract.
type ReverseRegistrarContract struct {
	ReverseRegistrarContractCaller     // Read-only binding to the contract
	ReverseRegistrarContractTransactor // Write-only binding to the contract
}

// ReverseRegistrarContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverseRegistrarContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseRegistrarContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverseRegistrarContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseRegistrarContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverseRegistrarContractSession struct {
	Contract     *ReverseRegistrarContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReverseRegistrarContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverseRegistrarContractCallerSession struct {
	Contract *ReverseRegistrarContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ReverseRegistrarContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverseRegistrarContractTransactorSession struct {
	Contract     *ReverseRegistrarContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ReverseRegistrarContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverseRegistrarContractRaw struct {
	Contract *ReverseRegistrarContract // Generic contract binding to access the raw methods on
}

// ReverseRegistrarContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverseRegistrarContractCallerRaw struct {
	Contract *ReverseRegistrarContractCaller // Generic read-only contract binding to access the raw methods on
}

// ReverseRegistrarContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverseRegistrarContractTransactorRaw struct {
	Contract *ReverseRegistrarContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverseRegistrarContract creates a new instance of ReverseRegistrarContract, bound to a specific deployed contract.
func NewReverseRegistrarContract(address common.Address, backend bind.ContractBackend) (*ReverseRegistrarContract, error) {
	contract, err := bindReverseRegistrarContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarContract{ReverseRegistrarContractCaller: ReverseRegistrarContractCaller{contract: contract}, ReverseRegistrarContractTransactor: ReverseRegistrarContractTransactor{contract: contract}}, nil
}

// NewReverseRegistrarContractCaller creates a new read-only instance of ReverseRegistrarContract, bound to a specific deployed contract.
func NewReverseRegistrarContractCaller(address common.Address, caller bind.ContractCaller) (*ReverseRegistrarContractCaller, error) {
	contract, err := bindReverseRegistrarContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarContractCaller{contract: contract}, nil
}

// NewReverseRegistrarContractTransactor creates a new write-only instance of ReverseRegistrarContract, bound to a specific deployed contract.
func NewReverseRegistrarContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverseRegistrarContractTransactor, error) {
	contract, err := bindReverseRegistrarContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReverseRegistrarContractTransactor{contract: contract}, nil
}

// bindReverseRegistrarContract binds a generic wrapper to an already deployed contract.
func bindReverseRegistrarContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverseRegistrarContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseRegistrarContract *ReverseRegistrarContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseRegistrarContract.Contract.ReverseRegistrarContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseRegistrarContract *ReverseRegistrarContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.ReverseRegistrarContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseRegistrarContract *ReverseRegistrarContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.ReverseRegistrarContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseRegistrarContract *ReverseRegistrarContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseRegistrarContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.contract.Transact(opts, method, params...)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractCaller) DefaultResolver(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReverseRegistrarContract.contract.Call(opts, out, "defaultResolver")
	return *ret0, err
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) DefaultResolver() (common.Address, error) {
	return _ReverseRegistrarContract.Contract.DefaultResolver(&_ReverseRegistrarContract.CallOpts)
}

// DefaultResolver is a free data retrieval call binding the contract method 0x828eab0e.
//
// Solidity: function defaultResolver() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractCallerSession) DefaultResolver() (common.Address, error) {
	return _ReverseRegistrarContract.Contract.DefaultResolver(&_ReverseRegistrarContract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReverseRegistrarContract.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) Ens() (common.Address, error) {
	return _ReverseRegistrarContract.Contract.Ens(&_ReverseRegistrarContract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseRegistrarContract *ReverseRegistrarContractCallerSession) Ens() (common.Address, error) {
	return _ReverseRegistrarContract.Contract.Ens(&_ReverseRegistrarContract.CallOpts)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(ret bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ReverseRegistrarContract.contract.Call(opts, out, "node", addr)
	return *ret0, err
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(ret bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) Node(addr common.Address) ([32]byte, error) {
	return _ReverseRegistrarContract.Contract.Node(&_ReverseRegistrarContract.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(addr address) constant returns(ret bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _ReverseRegistrarContract.Contract.Node(&_ReverseRegistrarContract.CallOpts, addr)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactor) Claim(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.contract.Transact(opts, "claim", owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.Claim(&_ReverseRegistrarContract.TransactOpts, owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(owner address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactorSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.Claim(&_ReverseRegistrarContract.TransactOpts, owner)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactor) ClaimWithResolver(opts *bind.TransactOpts, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.contract.Transact(opts, "claimWithResolver", owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.ClaimWithResolver(&_ReverseRegistrarContract.TransactOpts, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(owner address, resolver address) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactorSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.ClaimWithResolver(&_ReverseRegistrarContract.TransactOpts, owner, resolver)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ReverseRegistrarContract.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractSession) SetName(name string) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.SetName(&_ReverseRegistrarContract.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name string) returns(node bytes32)
func (_ReverseRegistrarContract *ReverseRegistrarContractTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _ReverseRegistrarContract.Contract.SetName(&_ReverseRegistrarContract.TransactOpts, name)
}
