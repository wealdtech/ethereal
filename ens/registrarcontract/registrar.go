// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrarcontract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RegistrarContractABI is the input ABI used to generate the binding from.
const RegistrarContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"releaseDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getAllowedTime\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"unhashedName\",\"type\":\"string\"}],\"name\":\"invalidateName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"shaBid\",\"outputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"seal\",\"type\":\"bytes32\"}],\"name\":\"cancelBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"entries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"unsealBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"transferRegistrars\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sealedBids\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"finalizeAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registryStarted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"launchLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"newBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"labels\",\"type\":\"bytes32[]\"}],\"name\":\"eraseNode\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"startAuctions\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"deed\",\"type\":\"address\"},{\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"acceptRegistrarTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"name\":\"sealedBid\",\"type\":\"bytes32\"}],\"name\":\"startAuctionsAndBid\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"},{\"name\":\"_startDate\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"HashReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"registrationDate\",\"type\":\"uint256\"}],\"name\":\"HashInvalidated\",\"type\":\"event\"}]"

// RegistrarContract is an auto generated Go binding around an Ethereum contract.
type RegistrarContract struct {
	RegistrarContractCaller     // Read-only binding to the contract
	RegistrarContractTransactor // Write-only binding to the contract
	RegistrarContractFilterer   // Log filterer for contract events
}

// RegistrarContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarContractSession struct {
	Contract     *RegistrarContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RegistrarContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarContractCallerSession struct {
	Contract *RegistrarContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RegistrarContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarContractTransactorSession struct {
	Contract     *RegistrarContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RegistrarContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarContractRaw struct {
	Contract *RegistrarContract // Generic contract binding to access the raw methods on
}

// RegistrarContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarContractCallerRaw struct {
	Contract *RegistrarContractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarContractTransactorRaw struct {
	Contract *RegistrarContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrarContract creates a new instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContract(address common.Address, backend bind.ContractBackend) (*RegistrarContract, error) {
	contract, err := bindRegistrarContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistrarContract{RegistrarContractCaller: RegistrarContractCaller{contract: contract}, RegistrarContractTransactor: RegistrarContractTransactor{contract: contract}, RegistrarContractFilterer: RegistrarContractFilterer{contract: contract}}, nil
}

// NewRegistrarContractCaller creates a new read-only instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractCaller(address common.Address, caller bind.ContractCaller) (*RegistrarContractCaller, error) {
	contract, err := bindRegistrarContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractCaller{contract: contract}, nil
}

// NewRegistrarContractTransactor creates a new write-only instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarContractTransactor, error) {
	contract, err := bindRegistrarContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractTransactor{contract: contract}, nil
}

// NewRegistrarContractFilterer creates a new log filterer instance of RegistrarContract, bound to a specific deployed contract.
func NewRegistrarContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarContractFilterer, error) {
	contract, err := bindRegistrarContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractFilterer{contract: contract}, nil
}

// bindRegistrarContract binds a generic wrapper to an already deployed contract.
func bindRegistrarContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarContract *RegistrarContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RegistrarContract.Contract.RegistrarContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarContract *RegistrarContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegistrarContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarContract *RegistrarContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarContract.Contract.RegistrarContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrarContract *RegistrarContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RegistrarContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrarContract *RegistrarContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrarContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrarContract *RegistrarContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrarContract.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegistrarContract *RegistrarContractCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegistrarContract *RegistrarContractSession) Ens() (common.Address, error) {
	return _RegistrarContract.Contract.Ens(&_RegistrarContract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) Ens() (common.Address, error) {
	return _RegistrarContract.Contract.Ens(&_RegistrarContract.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_RegistrarContract *RegistrarContractCaller) Entries(opts *bind.CallOpts, _hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _RegistrarContract.contract.Call(opts, out, "entries", _hash)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_RegistrarContract *RegistrarContractSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _RegistrarContract.Contract.Entries(&_RegistrarContract.CallOpts, _hash)
}

// Entries is a free data retrieval call binding the contract method 0x267b6922.
//
// Solidity: function entries(_hash bytes32) constant returns(uint8, address, uint256, uint256, uint256)
func (_RegistrarContract *RegistrarContractCallerSession) Entries(_hash [32]byte) (uint8, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _RegistrarContract.Contract.Entries(&_RegistrarContract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_RegistrarContract *RegistrarContractCaller) GetAllowedTime(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "getAllowedTime", _hash)
	return *ret0, err
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_RegistrarContract *RegistrarContractSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _RegistrarContract.Contract.GetAllowedTime(&_RegistrarContract.CallOpts, _hash)
}

// GetAllowedTime is a free data retrieval call binding the contract method 0x13c89a8f.
//
// Solidity: function getAllowedTime(_hash bytes32) constant returns(timestamp uint256)
func (_RegistrarContract *RegistrarContractCallerSession) GetAllowedTime(_hash [32]byte) (*big.Int, error) {
	return _RegistrarContract.Contract.GetAllowedTime(&_RegistrarContract.CallOpts, _hash)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_RegistrarContract *RegistrarContractCaller) IsAllowed(opts *bind.CallOpts, _hash [32]byte, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "isAllowed", _hash, _timestamp)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_RegistrarContract *RegistrarContractSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _RegistrarContract.Contract.IsAllowed(&_RegistrarContract.CallOpts, _hash, _timestamp)
}

// IsAllowed is a free data retrieval call binding the contract method 0x93503337.
//
// Solidity: function isAllowed(_hash bytes32, _timestamp uint256) constant returns(allowed bool)
func (_RegistrarContract *RegistrarContractCallerSession) IsAllowed(_hash [32]byte, _timestamp *big.Int) (bool, error) {
	return _RegistrarContract.Contract.IsAllowed(&_RegistrarContract.CallOpts, _hash, _timestamp)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_RegistrarContract *RegistrarContractCaller) LaunchLength(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "launchLength")
	return *ret0, err
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_RegistrarContract *RegistrarContractSession) LaunchLength() (uint32, error) {
	return _RegistrarContract.Contract.LaunchLength(&_RegistrarContract.CallOpts)
}

// LaunchLength is a free data retrieval call binding the contract method 0xae1a0b0c.
//
// Solidity: function launchLength() constant returns(uint32)
func (_RegistrarContract *RegistrarContractCallerSession) LaunchLength() (uint32, error) {
	return _RegistrarContract.Contract.LaunchLength(&_RegistrarContract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_RegistrarContract *RegistrarContractCaller) RegistryStarted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "registryStarted")
	return *ret0, err
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_RegistrarContract *RegistrarContractSession) RegistryStarted() (*big.Int, error) {
	return _RegistrarContract.Contract.RegistryStarted(&_RegistrarContract.CallOpts)
}

// RegistryStarted is a free data retrieval call binding the contract method 0x9c67f06f.
//
// Solidity: function registryStarted() constant returns(uint256)
func (_RegistrarContract *RegistrarContractCallerSession) RegistryStarted() (*big.Int, error) {
	return _RegistrarContract.Contract.RegistryStarted(&_RegistrarContract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegistrarContract *RegistrarContractCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegistrarContract *RegistrarContractSession) RootNode() ([32]byte, error) {
	return _RegistrarContract.Contract.RootNode(&_RegistrarContract.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_RegistrarContract *RegistrarContractCallerSession) RootNode() ([32]byte, error) {
	return _RegistrarContract.Contract.RootNode(&_RegistrarContract.CallOpts)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_RegistrarContract *RegistrarContractCaller) SealedBids(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "sealedBids", arg0, arg1)
	return *ret0, err
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_RegistrarContract *RegistrarContractSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _RegistrarContract.Contract.SealedBids(&_RegistrarContract.CallOpts, arg0, arg1)
}

// SealedBids is a free data retrieval call binding the contract method 0x5e431709.
//
// Solidity: function sealedBids( address,  bytes32) constant returns(address)
func (_RegistrarContract *RegistrarContractCallerSession) SealedBids(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _RegistrarContract.Contract.SealedBids(&_RegistrarContract.CallOpts, arg0, arg1)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_RegistrarContract *RegistrarContractCaller) ShaBid(opts *bind.CallOpts, hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "shaBid", hash, owner, value, salt)
	return *ret0, err
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_RegistrarContract *RegistrarContractSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _RegistrarContract.Contract.ShaBid(&_RegistrarContract.CallOpts, hash, owner, value, salt)
}

// ShaBid is a free data retrieval call binding the contract method 0x22ec1244.
//
// Solidity: function shaBid(hash bytes32, owner address, value uint256, salt bytes32) constant returns(sealedBid bytes32)
func (_RegistrarContract *RegistrarContractCallerSession) ShaBid(hash [32]byte, owner common.Address, value *big.Int, salt [32]byte) ([32]byte, error) {
	return _RegistrarContract.Contract.ShaBid(&_RegistrarContract.CallOpts, hash, owner, value, salt)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_RegistrarContract *RegistrarContractCaller) State(opts *bind.CallOpts, _hash [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RegistrarContract.contract.Call(opts, out, "state", _hash)
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_RegistrarContract *RegistrarContractSession) State(_hash [32]byte) (uint8, error) {
	return _RegistrarContract.Contract.State(&_RegistrarContract.CallOpts, _hash)
}

// State is a free data retrieval call binding the contract method 0x61d585da.
//
// Solidity: function state(_hash bytes32) constant returns(uint8)
func (_RegistrarContract *RegistrarContractCallerSession) State(_hash [32]byte) (uint8, error) {
	return _RegistrarContract.Contract.State(&_RegistrarContract.CallOpts, _hash)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_RegistrarContract *RegistrarContractTransactor) AcceptRegistrarTransfer(opts *bind.TransactOpts, hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "acceptRegistrarTransfer", hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_RegistrarContract *RegistrarContractSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptRegistrarTransfer(&_RegistrarContract.TransactOpts, hash, deed, registrationDate)
}

// AcceptRegistrarTransfer is a paid mutator transaction binding the contract method 0xea9e107a.
//
// Solidity: function acceptRegistrarTransfer(hash bytes32, deed address, registrationDate uint256) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) AcceptRegistrarTransfer(hash [32]byte, deed common.Address, registrationDate *big.Int) (*types.Transaction, error) {
	return _RegistrarContract.Contract.AcceptRegistrarTransfer(&_RegistrarContract.TransactOpts, hash, deed, registrationDate)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) CancelBid(opts *bind.TransactOpts, bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "cancelBid", bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CancelBid(&_RegistrarContract.TransactOpts, bidder, seal)
}

// CancelBid is a paid mutator transaction binding the contract method 0x2525f5c1.
//
// Solidity: function cancelBid(bidder address, seal bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) CancelBid(bidder common.Address, seal [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.CancelBid(&_RegistrarContract.TransactOpts, bidder, seal)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_RegistrarContract *RegistrarContractTransactor) EraseNode(opts *bind.TransactOpts, labels [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "eraseNode", labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_RegistrarContract *RegistrarContractSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.EraseNode(&_RegistrarContract.TransactOpts, labels)
}

// EraseNode is a paid mutator transaction binding the contract method 0xde10f04b.
//
// Solidity: function eraseNode(labels bytes32[]) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) EraseNode(labels [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.EraseNode(&_RegistrarContract.TransactOpts, labels)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) FinalizeAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "finalizeAuction", _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.FinalizeAuction(&_RegistrarContract.TransactOpts, _hash)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x983b94fb.
//
// Solidity: function finalizeAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) FinalizeAuction(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.FinalizeAuction(&_RegistrarContract.TransactOpts, _hash)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_RegistrarContract *RegistrarContractTransactor) InvalidateName(opts *bind.TransactOpts, unhashedName string) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "invalidateName", unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_RegistrarContract *RegistrarContractSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _RegistrarContract.Contract.InvalidateName(&_RegistrarContract.TransactOpts, unhashedName)
}

// InvalidateName is a paid mutator transaction binding the contract method 0x15f73331.
//
// Solidity: function invalidateName(unhashedName string) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) InvalidateName(unhashedName string) (*types.Transaction, error) {
	return _RegistrarContract.Contract.InvalidateName(&_RegistrarContract.TransactOpts, unhashedName)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) NewBid(opts *bind.TransactOpts, sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "newBid", sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.NewBid(&_RegistrarContract.TransactOpts, sealedBid)
}

// NewBid is a paid mutator transaction binding the contract method 0xce92dced.
//
// Solidity: function newBid(sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) NewBid(sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.NewBid(&_RegistrarContract.TransactOpts, sealedBid)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) ReleaseDeed(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "releaseDeed", _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.ReleaseDeed(&_RegistrarContract.TransactOpts, _hash)
}

// ReleaseDeed is a paid mutator transaction binding the contract method 0x0230a07c.
//
// Solidity: function releaseDeed(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) ReleaseDeed(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.ReleaseDeed(&_RegistrarContract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) StartAuction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "startAuction", _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuction(&_RegistrarContract.TransactOpts, _hash)
}

// StartAuction is a paid mutator transaction binding the contract method 0xede8acdb.
//
// Solidity: function startAuction(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) StartAuction(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuction(&_RegistrarContract.TransactOpts, _hash)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_RegistrarContract *RegistrarContractTransactor) StartAuctions(opts *bind.TransactOpts, _hashes [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "startAuctions", _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_RegistrarContract *RegistrarContractSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuctions(&_RegistrarContract.TransactOpts, _hashes)
}

// StartAuctions is a paid mutator transaction binding the contract method 0xe27fe50f.
//
// Solidity: function startAuctions(_hashes bytes32[]) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) StartAuctions(_hashes [][32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuctions(&_RegistrarContract.TransactOpts, _hashes)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) StartAuctionsAndBid(opts *bind.TransactOpts, hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "startAuctionsAndBid", hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuctionsAndBid(&_RegistrarContract.TransactOpts, hashes, sealedBid)
}

// StartAuctionsAndBid is a paid mutator transaction binding the contract method 0xfebefd61.
//
// Solidity: function startAuctionsAndBid(hashes bytes32[], sealedBid bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) StartAuctionsAndBid(hashes [][32]byte, sealedBid [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.StartAuctionsAndBid(&_RegistrarContract.TransactOpts, hashes, sealedBid)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_RegistrarContract *RegistrarContractTransactor) Transfer(opts *bind.TransactOpts, _hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transfer", _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_RegistrarContract *RegistrarContractSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Transfer(&_RegistrarContract.TransactOpts, _hash, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(_hash bytes32, newOwner address) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) Transfer(_hash [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrarContract.Contract.Transfer(&_RegistrarContract.TransactOpts, _hash, newOwner)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) TransferRegistrars(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "transferRegistrars", _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferRegistrars(&_RegistrarContract.TransactOpts, _hash)
}

// TransferRegistrars is a paid mutator transaction binding the contract method 0x5ddae283.
//
// Solidity: function transferRegistrars(_hash bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) TransferRegistrars(_hash [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.TransferRegistrars(&_RegistrarContract.TransactOpts, _hash)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactor) UnsealBid(opts *bind.TransactOpts, _hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.contract.Transact(opts, "unsealBid", _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_RegistrarContract *RegistrarContractSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UnsealBid(&_RegistrarContract.TransactOpts, _hash, _value, _salt)
}

// UnsealBid is a paid mutator transaction binding the contract method 0x47872b42.
//
// Solidity: function unsealBid(_hash bytes32, _value uint256, _salt bytes32) returns()
func (_RegistrarContract *RegistrarContractTransactorSession) UnsealBid(_hash [32]byte, _value *big.Int, _salt [32]byte) (*types.Transaction, error) {
	return _RegistrarContract.Contract.UnsealBid(&_RegistrarContract.TransactOpts, _hash, _value, _salt)
}

// RegistrarContractAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the RegistrarContract contract.
type RegistrarContractAuctionStartedIterator struct {
	Event *RegistrarContractAuctionStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractAuctionStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractAuctionStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractAuctionStarted represents a AuctionStarted event raised by the RegistrarContract contract.
type RegistrarContractAuctionStarted struct {
	Hash             [32]byte
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x87e97e825a1d1fa0c54e1d36c7506c1dea8b1efd451fe68b000cf96f7cf40003.
//
// Solidity: e AuctionStarted(hash indexed bytes32, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) FilterAuctionStarted(opts *bind.FilterOpts, hash [][32]byte) (*RegistrarContractAuctionStartedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "AuctionStarted", hashRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractAuctionStartedIterator{contract: _RegistrarContract.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x87e97e825a1d1fa0c54e1d36c7506c1dea8b1efd451fe68b000cf96f7cf40003.
//
// Solidity: e AuctionStarted(hash indexed bytes32, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *RegistrarContractAuctionStarted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "AuctionStarted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractAuctionStarted)
				if err := _RegistrarContract.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistrarContractBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the RegistrarContract contract.
type RegistrarContractBidRevealedIterator struct {
	Event *RegistrarContractBidRevealed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractBidRevealed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractBidRevealed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractBidRevealed represents a BidRevealed event raised by the RegistrarContract contract.
type RegistrarContractBidRevealed struct {
	Hash   [32]byte
	Owner  common.Address
	Value  *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x7b6c4b278d165a6b33958f8ea5dfb00c8c9d4d0acf1985bef5d10786898bc3e7.
//
// Solidity: e BidRevealed(hash indexed bytes32, owner indexed address, value uint256, status uint8)
func (_RegistrarContract *RegistrarContractFilterer) FilterBidRevealed(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*RegistrarContractBidRevealedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "BidRevealed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractBidRevealedIterator{contract: _RegistrarContract.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x7b6c4b278d165a6b33958f8ea5dfb00c8c9d4d0acf1985bef5d10786898bc3e7.
//
// Solidity: e BidRevealed(hash indexed bytes32, owner indexed address, value uint256, status uint8)
func (_RegistrarContract *RegistrarContractFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *RegistrarContractBidRevealed, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "BidRevealed", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractBidRevealed)
				if err := _RegistrarContract.contract.UnpackLog(event, "BidRevealed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistrarContractHashInvalidatedIterator is returned from FilterHashInvalidated and is used to iterate over the raw logs and unpacked data for HashInvalidated events raised by the RegistrarContract contract.
type RegistrarContractHashInvalidatedIterator struct {
	Event *RegistrarContractHashInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractHashInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractHashInvalidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractHashInvalidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractHashInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractHashInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractHashInvalidated represents a HashInvalidated event raised by the RegistrarContract contract.
type RegistrarContractHashInvalidated struct {
	Hash             [32]byte
	Name             common.Hash
	Value            *big.Int
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHashInvalidated is a free log retrieval operation binding the contract event 0x1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad.
//
// Solidity: e HashInvalidated(hash indexed bytes32, name indexed string, value uint256, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) FilterHashInvalidated(opts *bind.FilterOpts, hash [][32]byte, name []string) (*RegistrarContractHashInvalidatedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "HashInvalidated", hashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractHashInvalidatedIterator{contract: _RegistrarContract.contract, event: "HashInvalidated", logs: logs, sub: sub}, nil
}

// WatchHashInvalidated is a free log subscription operation binding the contract event 0x1f9c649fe47e58bb60f4e52f0d90e4c47a526c9f90c5113df842c025970b66ad.
//
// Solidity: e HashInvalidated(hash indexed bytes32, name indexed string, value uint256, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) WatchHashInvalidated(opts *bind.WatchOpts, sink chan<- *RegistrarContractHashInvalidated, hash [][32]byte, name []string) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "HashInvalidated", hashRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractHashInvalidated)
				if err := _RegistrarContract.contract.UnpackLog(event, "HashInvalidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistrarContractHashRegisteredIterator is returned from FilterHashRegistered and is used to iterate over the raw logs and unpacked data for HashRegistered events raised by the RegistrarContract contract.
type RegistrarContractHashRegisteredIterator struct {
	Event *RegistrarContractHashRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractHashRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractHashRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractHashRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractHashRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractHashRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractHashRegistered represents a HashRegistered event raised by the RegistrarContract contract.
type RegistrarContractHashRegistered struct {
	Hash             [32]byte
	Owner            common.Address
	Value            *big.Int
	RegistrationDate *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHashRegistered is a free log retrieval operation binding the contract event 0x0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670.
//
// Solidity: e HashRegistered(hash indexed bytes32, owner indexed address, value uint256, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) FilterHashRegistered(opts *bind.FilterOpts, hash [][32]byte, owner []common.Address) (*RegistrarContractHashRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "HashRegistered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractHashRegisteredIterator{contract: _RegistrarContract.contract, event: "HashRegistered", logs: logs, sub: sub}, nil
}

// WatchHashRegistered is a free log subscription operation binding the contract event 0x0f0c27adfd84b60b6f456b0e87cdccb1e5fb9603991588d87fa99f5b6b61e670.
//
// Solidity: e HashRegistered(hash indexed bytes32, owner indexed address, value uint256, registrationDate uint256)
func (_RegistrarContract *RegistrarContractFilterer) WatchHashRegistered(opts *bind.WatchOpts, sink chan<- *RegistrarContractHashRegistered, hash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "HashRegistered", hashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractHashRegistered)
				if err := _RegistrarContract.contract.UnpackLog(event, "HashRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistrarContractHashReleasedIterator is returned from FilterHashReleased and is used to iterate over the raw logs and unpacked data for HashReleased events raised by the RegistrarContract contract.
type RegistrarContractHashReleasedIterator struct {
	Event *RegistrarContractHashReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractHashReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractHashReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractHashReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractHashReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractHashReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractHashReleased represents a HashReleased event raised by the RegistrarContract contract.
type RegistrarContractHashReleased struct {
	Hash  [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHashReleased is a free log retrieval operation binding the contract event 0x292b79b9246fa2c8e77d3fe195b251f9cb839d7d038e667c069ee7708c631e16.
//
// Solidity: e HashReleased(hash indexed bytes32, value uint256)
func (_RegistrarContract *RegistrarContractFilterer) FilterHashReleased(opts *bind.FilterOpts, hash [][32]byte) (*RegistrarContractHashReleasedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "HashReleased", hashRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractHashReleasedIterator{contract: _RegistrarContract.contract, event: "HashReleased", logs: logs, sub: sub}, nil
}

// WatchHashReleased is a free log subscription operation binding the contract event 0x292b79b9246fa2c8e77d3fe195b251f9cb839d7d038e667c069ee7708c631e16.
//
// Solidity: e HashReleased(hash indexed bytes32, value uint256)
func (_RegistrarContract *RegistrarContractFilterer) WatchHashReleased(opts *bind.WatchOpts, sink chan<- *RegistrarContractHashReleased, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "HashReleased", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractHashReleased)
				if err := _RegistrarContract.contract.UnpackLog(event, "HashReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RegistrarContractNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the RegistrarContract contract.
type RegistrarContractNewBidIterator struct {
	Event *RegistrarContractNewBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistrarContractNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarContractNewBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistrarContractNewBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistrarContractNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarContractNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarContractNewBid represents a NewBid event raised by the RegistrarContract contract.
type RegistrarContractNewBid struct {
	Hash    [32]byte
	Bidder  common.Address
	Deposit *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xb556ff269c1b6714f432c36431e2041d28436a73b6c3f19c021827bbdc6bfc29.
//
// Solidity: e NewBid(hash indexed bytes32, bidder indexed address, deposit uint256)
func (_RegistrarContract *RegistrarContractFilterer) FilterNewBid(opts *bind.FilterOpts, hash [][32]byte, bidder []common.Address) (*RegistrarContractNewBidIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _RegistrarContract.contract.FilterLogs(opts, "NewBid", hashRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarContractNewBidIterator{contract: _RegistrarContract.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xb556ff269c1b6714f432c36431e2041d28436a73b6c3f19c021827bbdc6bfc29.
//
// Solidity: e NewBid(hash indexed bytes32, bidder indexed address, deposit uint256)
func (_RegistrarContract *RegistrarContractFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *RegistrarContractNewBid, hash [][32]byte, bidder []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _RegistrarContract.contract.WatchLogs(opts, "NewBid", hashRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarContractNewBid)
				if err := _RegistrarContract.contract.UnpackLog(event, "NewBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
