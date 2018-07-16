// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deedcontract

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

// DeedContractABI is the input ABI used to generate the binding from.
const DeedContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creationDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroyDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint256\"},{\"name\":\"throwOnFailure\",\"type\":\"bool\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"closeDeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRegistrar\",\"type\":\"address\"}],\"name\":\"setRegistrar\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":true,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DeedClosed\",\"type\":\"event\"}]"

// DeedContract is an auto generated Go binding around an Ethereum contract.
type DeedContract struct {
	DeedContractCaller     // Read-only binding to the contract
	DeedContractTransactor // Write-only binding to the contract
	DeedContractFilterer   // Log filterer for contract events
}

// DeedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeedContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeedContractFilterer struct {
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
	contract, err := bindDeedContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeedContract{DeedContractCaller: DeedContractCaller{contract: contract}, DeedContractTransactor: DeedContractTransactor{contract: contract}, DeedContractFilterer: DeedContractFilterer{contract: contract}}, nil
}

// NewDeedContractCaller creates a new read-only instance of DeedContract, bound to a specific deployed contract.
func NewDeedContractCaller(address common.Address, caller bind.ContractCaller) (*DeedContractCaller, error) {
	contract, err := bindDeedContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeedContractCaller{contract: contract}, nil
}

// NewDeedContractTransactor creates a new write-only instance of DeedContract, bound to a specific deployed contract.
func NewDeedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DeedContractTransactor, error) {
	contract, err := bindDeedContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeedContractTransactor{contract: contract}, nil
}

// NewDeedContractFilterer creates a new log filterer instance of DeedContract, bound to a specific deployed contract.
func NewDeedContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DeedContractFilterer, error) {
	contract, err := bindDeedContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeedContractFilterer{contract: contract}, nil
}

// bindDeedContract binds a generic wrapper to an already deployed contract.
func bindDeedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeedContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DeedContractDeedClosedIterator is returned from FilterDeedClosed and is used to iterate over the raw logs and unpacked data for DeedClosed events raised by the DeedContract contract.
type DeedContractDeedClosedIterator struct {
	Event *DeedContractDeedClosed // Event containing the contract specifics and raw log

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
func (it *DeedContractDeedClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeedContractDeedClosed)
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
		it.Event = new(DeedContractDeedClosed)
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
func (it *DeedContractDeedClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeedContractDeedClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeedContractDeedClosed represents a DeedClosed event raised by the DeedContract contract.
type DeedContractDeedClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeedClosed is a free log retrieval operation binding the contract event 0xbb2ce2f51803bba16bc85282b47deeea9a5c6223eabea1077be696b3f265cf13.
//
// Solidity: e DeedClosed()
func (_DeedContract *DeedContractFilterer) FilterDeedClosed(opts *bind.FilterOpts) (*DeedContractDeedClosedIterator, error) {

	logs, sub, err := _DeedContract.contract.FilterLogs(opts, "DeedClosed")
	if err != nil {
		return nil, err
	}
	return &DeedContractDeedClosedIterator{contract: _DeedContract.contract, event: "DeedClosed", logs: logs, sub: sub}, nil
}

// WatchDeedClosed is a free log subscription operation binding the contract event 0xbb2ce2f51803bba16bc85282b47deeea9a5c6223eabea1077be696b3f265cf13.
//
// Solidity: e DeedClosed()
func (_DeedContract *DeedContractFilterer) WatchDeedClosed(opts *bind.WatchOpts, sink chan<- *DeedContractDeedClosed) (event.Subscription, error) {

	logs, sub, err := _DeedContract.contract.WatchLogs(opts, "DeedClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeedContractDeedClosed)
				if err := _DeedContract.contract.UnpackLog(event, "DeedClosed", log); err != nil {
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

// DeedContractOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the DeedContract contract.
type DeedContractOwnerChangedIterator struct {
	Event *DeedContractOwnerChanged // Event containing the contract specifics and raw log

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
func (it *DeedContractOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeedContractOwnerChanged)
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
		it.Event = new(DeedContractOwnerChanged)
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
func (it *DeedContractOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeedContractOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeedContractOwnerChanged represents a OwnerChanged event raised by the DeedContract contract.
type DeedContractOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(newOwner address)
func (_DeedContract *DeedContractFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*DeedContractOwnerChangedIterator, error) {

	logs, sub, err := _DeedContract.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &DeedContractOwnerChangedIterator{contract: _DeedContract.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: e OwnerChanged(newOwner address)
func (_DeedContract *DeedContractFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *DeedContractOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _DeedContract.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeedContractOwnerChanged)
				if err := _DeedContract.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
