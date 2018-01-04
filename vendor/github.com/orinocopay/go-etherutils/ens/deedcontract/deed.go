// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package deedcontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DeedContractABI is the input ABI used to generate the binding from.
const DeedContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creationDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroyDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint256\"},{\"name\":\"throwOnFailure\",\"type\":\"bool\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"closeDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRegistrar\",\"type\":\"address\"}],\"name\":\"setRegistrar\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":true,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DeedClosed\",\"type\":\"event\"}]"

// DeedContract is an auto generated Go binding around an Ethereum contract.
type DeedContract struct {
	DeedContractCaller     // Read-only binding to the contract
	DeedContractTransactor // Write-only binding to the contract
}

// DeedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeedContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeedContractSession struct {
	Contract     *DeedContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeedContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeedContractCallerSession struct {
	Contract *DeedContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DeedContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeedContractTransactorSession struct {
	Contract     *DeedContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DeedContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeedContractRaw struct {
	Contract *DeedContract // Generic contract binding to access the raw methods on
}

// DeedContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeedContractCallerRaw struct {
	Contract *DeedContractCaller // Generic read-only contract binding to access the raw methods on
}

// DeedContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeedContractTransactorRaw struct {
	Contract *DeedContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeedContract creates a new instance of DeedContract, bound to a specific deployed contract.
func NewDeedContract(address common.Address, backend bind.ContractBackend) (*DeedContract, error) {
	contract, err := bindDeedContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeedContract{DeedContractCaller: DeedContractCaller{contract: contract}, DeedContractTransactor: DeedContractTransactor{contract: contract}}, nil
}

// NewDeedContractCaller creates a new read-only instance of DeedContract, bound to a specific deployed contract.
func NewDeedContractCaller(address common.Address, caller bind.ContractCaller) (*DeedContractCaller, error) {
	contract, err := bindDeedContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DeedContractCaller{contract: contract}, nil
}

// NewDeedContractTransactor creates a new write-only instance of DeedContract, bound to a specific deployed contract.
func NewDeedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DeedContractTransactor, error) {
	contract, err := bindDeedContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DeedContractTransactor{contract: contract}, nil
}

// bindDeedContract binds a generic wrapper to an already deployed contract.
func bindDeedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeedContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeedContract *DeedContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeedContract.Contract.DeedContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeedContract *DeedContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeedContract.Contract.DeedContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeedContract *DeedContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeedContract.Contract.DeedContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeedContract *DeedContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeedContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeedContract *DeedContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeedContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeedContract *DeedContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeedContract.Contract.contract.Transact(opts, method, params...)
}

// CreationDate is a free data retrieval call binding the contract method 0x05b34410.
//
// Solidity: function creationDate() constant returns(uint256)
func (_DeedContract *DeedContractCaller) CreationDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DeedContract.contract.Call(opts, out, "creationDate")
	return *ret0, err
}

// CreationDate is a free data retrieval call binding the contract method 0x05b34410.
//
// Solidity: function creationDate() constant returns(uint256)
func (_DeedContract *DeedContractSession) CreationDate() (*big.Int, error) {
	return _DeedContract.Contract.CreationDate(&_DeedContract.CallOpts)
}

// CreationDate is a free data retrieval call binding the contract method 0x05b34410.
//
// Solidity: function creationDate() constant returns(uint256)
func (_DeedContract *DeedContractCallerSession) CreationDate() (*big.Int, error) {
	return _DeedContract.Contract.CreationDate(&_DeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeedContract *DeedContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeedContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeedContract *DeedContractSession) Owner() (common.Address, error) {
	return _DeedContract.Contract.Owner(&_DeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeedContract *DeedContractCallerSession) Owner() (common.Address, error) {
	return _DeedContract.Contract.Owner(&_DeedContract.CallOpts)
}

// PreviousOwner is a free data retrieval call binding the contract method 0x674f220f.
//
// Solidity: function previousOwner() constant returns(address)
func (_DeedContract *DeedContractCaller) PreviousOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeedContract.contract.Call(opts, out, "previousOwner")
	return *ret0, err
}

// PreviousOwner is a free data retrieval call binding the contract method 0x674f220f.
//
// Solidity: function previousOwner() constant returns(address)
func (_DeedContract *DeedContractSession) PreviousOwner() (common.Address, error) {
	return _DeedContract.Contract.PreviousOwner(&_DeedContract.CallOpts)
}

// PreviousOwner is a free data retrieval call binding the contract method 0x674f220f.
//
// Solidity: function previousOwner() constant returns(address)
func (_DeedContract *DeedContractCallerSession) PreviousOwner() (common.Address, error) {
	return _DeedContract.Contract.PreviousOwner(&_DeedContract.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_DeedContract *DeedContractCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeedContract.contract.Call(opts, out, "registrar")
	return *ret0, err
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_DeedContract *DeedContractSession) Registrar() (common.Address, error) {
	return _DeedContract.Contract.Registrar(&_DeedContract.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() constant returns(address)
func (_DeedContract *DeedContractCallerSession) Registrar() (common.Address, error) {
	return _DeedContract.Contract.Registrar(&_DeedContract.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_DeedContract *DeedContractCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DeedContract.contract.Call(opts, out, "value")
	return *ret0, err
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_DeedContract *DeedContractSession) Value() (*big.Int, error) {
	return _DeedContract.Contract.Value(&_DeedContract.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(uint256)
func (_DeedContract *DeedContractCallerSession) Value() (*big.Int, error) {
	return _DeedContract.Contract.Value(&_DeedContract.CallOpts)
}

// CloseDeed is a paid mutator transaction binding the contract method 0xbbe42771.
//
// Solidity: function closeDeed(refundRatio uint256) returns()
func (_DeedContract *DeedContractTransactor) CloseDeed(opts *bind.TransactOpts, refundRatio *big.Int) (*types.Transaction, error) {
	return _DeedContract.contract.Transact(opts, "closeDeed", refundRatio)
}

// CloseDeed is a paid mutator transaction binding the contract method 0xbbe42771.
//
// Solidity: function closeDeed(refundRatio uint256) returns()
func (_DeedContract *DeedContractSession) CloseDeed(refundRatio *big.Int) (*types.Transaction, error) {
	return _DeedContract.Contract.CloseDeed(&_DeedContract.TransactOpts, refundRatio)
}

// CloseDeed is a paid mutator transaction binding the contract method 0xbbe42771.
//
// Solidity: function closeDeed(refundRatio uint256) returns()
func (_DeedContract *DeedContractTransactorSession) CloseDeed(refundRatio *big.Int) (*types.Transaction, error) {
	return _DeedContract.Contract.CloseDeed(&_DeedContract.TransactOpts, refundRatio)
}

// DestroyDeed is a paid mutator transaction binding the contract method 0x0b5ab3d5.
//
// Solidity: function destroyDeed() returns()
func (_DeedContract *DeedContractTransactor) DestroyDeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeedContract.contract.Transact(opts, "destroyDeed")
}

// DestroyDeed is a paid mutator transaction binding the contract method 0x0b5ab3d5.
//
// Solidity: function destroyDeed() returns()
func (_DeedContract *DeedContractSession) DestroyDeed() (*types.Transaction, error) {
	return _DeedContract.Contract.DestroyDeed(&_DeedContract.TransactOpts)
}

// DestroyDeed is a paid mutator transaction binding the contract method 0x0b5ab3d5.
//
// Solidity: function destroyDeed() returns()
func (_DeedContract *DeedContractTransactorSession) DestroyDeed() (*types.Transaction, error) {
	return _DeedContract.Contract.DestroyDeed(&_DeedContract.TransactOpts)
}

// SetBalance is a paid mutator transaction binding the contract method 0xb0c80972.
//
// Solidity: function setBalance(newValue uint256, throwOnFailure bool) returns()
func (_DeedContract *DeedContractTransactor) SetBalance(opts *bind.TransactOpts, newValue *big.Int, throwOnFailure bool) (*types.Transaction, error) {
	return _DeedContract.contract.Transact(opts, "setBalance", newValue, throwOnFailure)
}

// SetBalance is a paid mutator transaction binding the contract method 0xb0c80972.
//
// Solidity: function setBalance(newValue uint256, throwOnFailure bool) returns()
func (_DeedContract *DeedContractSession) SetBalance(newValue *big.Int, throwOnFailure bool) (*types.Transaction, error) {
	return _DeedContract.Contract.SetBalance(&_DeedContract.TransactOpts, newValue, throwOnFailure)
}

// SetBalance is a paid mutator transaction binding the contract method 0xb0c80972.
//
// Solidity: function setBalance(newValue uint256, throwOnFailure bool) returns()
func (_DeedContract *DeedContractTransactorSession) SetBalance(newValue *big.Int, throwOnFailure bool) (*types.Transaction, error) {
	return _DeedContract.Contract.SetBalance(&_DeedContract.TransactOpts, newValue, throwOnFailure)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_DeedContract *DeedContractTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DeedContract.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_DeedContract *DeedContractSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DeedContract.Contract.SetOwner(&_DeedContract.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_DeedContract *DeedContractTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DeedContract.Contract.SetOwner(&_DeedContract.TransactOpts, newOwner)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(newRegistrar address) returns()
func (_DeedContract *DeedContractTransactor) SetRegistrar(opts *bind.TransactOpts, newRegistrar common.Address) (*types.Transaction, error) {
	return _DeedContract.contract.Transact(opts, "setRegistrar", newRegistrar)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(newRegistrar address) returns()
func (_DeedContract *DeedContractSession) SetRegistrar(newRegistrar common.Address) (*types.Transaction, error) {
	return _DeedContract.Contract.SetRegistrar(&_DeedContract.TransactOpts, newRegistrar)
}

// SetRegistrar is a paid mutator transaction binding the contract method 0xfaab9d39.
//
// Solidity: function setRegistrar(newRegistrar address) returns()
func (_DeedContract *DeedContractTransactorSession) SetRegistrar(newRegistrar common.Address) (*types.Transaction, error) {
	return _DeedContract.Contract.SetRegistrar(&_DeedContract.TransactOpts, newRegistrar)
}
