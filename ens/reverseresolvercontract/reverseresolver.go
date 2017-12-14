// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package reverseresolvercontract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ReverseResolverABI is the input ABI used to generate the binding from.
const ReverseResolverABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// ReverseResolver is an auto generated Go binding around an Ethereum contract.
type ReverseResolver struct {
	ReverseResolverCaller     // Read-only binding to the contract
	ReverseResolverTransactor // Write-only binding to the contract
}

// ReverseResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverseResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverseResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverseResolverSession struct {
	Contract     *ReverseResolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReverseResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverseResolverCallerSession struct {
	Contract *ReverseResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReverseResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverseResolverTransactorSession struct {
	Contract     *ReverseResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReverseResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverseResolverRaw struct {
	Contract *ReverseResolver // Generic contract binding to access the raw methods on
}

// ReverseResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverseResolverCallerRaw struct {
	Contract *ReverseResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ReverseResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverseResolverTransactorRaw struct {
	Contract *ReverseResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverseResolver creates a new instance of ReverseResolver, bound to a specific deployed contract.
func NewReverseResolver(address common.Address, backend bind.ContractBackend) (*ReverseResolver, error) {
	contract, err := bindReverseResolver(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReverseResolver{ReverseResolverCaller: ReverseResolverCaller{contract: contract}, ReverseResolverTransactor: ReverseResolverTransactor{contract: contract}}, nil
}

// NewReverseResolverCaller creates a new read-only instance of ReverseResolver, bound to a specific deployed contract.
func NewReverseResolverCaller(address common.Address, caller bind.ContractCaller) (*ReverseResolverCaller, error) {
	contract, err := bindReverseResolver(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseResolverCaller{contract: contract}, nil
}

// NewReverseResolverTransactor creates a new write-only instance of ReverseResolver, bound to a specific deployed contract.
func NewReverseResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverseResolverTransactor, error) {
	contract, err := bindReverseResolver(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReverseResolverTransactor{contract: contract}, nil
}

// bindReverseResolver binds a generic wrapper to an already deployed contract.
func bindReverseResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverseResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseResolver *ReverseResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseResolver.Contract.ReverseResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseResolver *ReverseResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseResolver.Contract.ReverseResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseResolver *ReverseResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseResolver.Contract.ReverseResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReverseResolver *ReverseResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReverseResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReverseResolver *ReverseResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReverseResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReverseResolver *ReverseResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReverseResolver.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseResolver *ReverseResolverCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReverseResolver.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseResolver *ReverseResolverSession) Ens() (common.Address, error) {
	return _ReverseResolver.Contract.Ens(&_ReverseResolver.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ReverseResolver *ReverseResolverCallerSession) Ens() (common.Address, error) {
	return _ReverseResolver.Contract.Ens(&_ReverseResolver.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name( bytes32) constant returns(string)
func (_ReverseResolver *ReverseResolverCaller) Name(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReverseResolver.contract.Call(opts, out, "name", arg0)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name( bytes32) constant returns(string)
func (_ReverseResolver *ReverseResolverSession) Name(arg0 [32]byte) (string, error) {
	return _ReverseResolver.Contract.Name(&_ReverseResolver.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name( bytes32) constant returns(string)
func (_ReverseResolver *ReverseResolverCallerSession) Name(arg0 [32]byte) (string, error) {
	return _ReverseResolver.Contract.Name(&_ReverseResolver.CallOpts, arg0)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, _name string) returns()
func (_ReverseResolver *ReverseResolverTransactor) SetName(opts *bind.TransactOpts, node [32]byte, _name string) (*types.Transaction, error) {
	return _ReverseResolver.contract.Transact(opts, "setName", node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, _name string) returns()
func (_ReverseResolver *ReverseResolverSession) SetName(node [32]byte, _name string) (*types.Transaction, error) {
	return _ReverseResolver.Contract.SetName(&_ReverseResolver.TransactOpts, node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, _name string) returns()
func (_ReverseResolver *ReverseResolverTransactorSession) SetName(node [32]byte, _name string) (*types.Transaction, error) {
	return _ReverseResolver.Contract.SetName(&_ReverseResolver.TransactOpts, node, _name)
}
