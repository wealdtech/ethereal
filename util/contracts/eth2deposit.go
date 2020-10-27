// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Eth2DepositABI is the input ABI used to generate the binding from.
const Eth2DepositABI = "[{\"name\":\"DepositEvent\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"pubkey\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"withdrawal_credentials\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"amount\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"signature\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"index\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_drain_address\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"get_deposit_root\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":95389},{\"name\":\"get_deposit_count\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":17683},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"pubkey\"},{\"type\":\"bytes\",\"name\":\"withdrawal_credentials\"},{\"type\":\"bytes\",\"name\":\"signature\"},{\"type\":\"bytes32\",\"name\":\"deposit_data_root\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":1754607},{\"name\":\"drain\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35793},{\"name\":\"drain_address\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]"

// Eth2Deposit is an auto generated Go binding around an Ethereum contract.
type Eth2Deposit struct {
	Eth2DepositCaller     // Read-only binding to the contract
	Eth2DepositTransactor // Write-only binding to the contract
	Eth2DepositFilterer   // Log filterer for contract events
}

// Eth2DepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type Eth2DepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2DepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Eth2DepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2DepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Eth2DepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2DepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Eth2DepositSession struct {
	Contract     *Eth2Deposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eth2DepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Eth2DepositCallerSession struct {
	Contract *Eth2DepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Eth2DepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Eth2DepositTransactorSession struct {
	Contract     *Eth2DepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Eth2DepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type Eth2DepositRaw struct {
	Contract *Eth2Deposit // Generic contract binding to access the raw methods on
}

// Eth2DepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Eth2DepositCallerRaw struct {
	Contract *Eth2DepositCaller // Generic read-only contract binding to access the raw methods on
}

// Eth2DepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Eth2DepositTransactorRaw struct {
	Contract *Eth2DepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEth2Deposit creates a new instance of Eth2Deposit, bound to a specific deployed contract.
func NewEth2Deposit(address common.Address, backend bind.ContractBackend) (*Eth2Deposit, error) {
	contract, err := bindEth2Deposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth2Deposit{Eth2DepositCaller: Eth2DepositCaller{contract: contract}, Eth2DepositTransactor: Eth2DepositTransactor{contract: contract}, Eth2DepositFilterer: Eth2DepositFilterer{contract: contract}}, nil
}

// NewEth2DepositCaller creates a new read-only instance of Eth2Deposit, bound to a specific deployed contract.
func NewEth2DepositCaller(address common.Address, caller bind.ContractCaller) (*Eth2DepositCaller, error) {
	contract, err := bindEth2Deposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Eth2DepositCaller{contract: contract}, nil
}

// NewEth2DepositTransactor creates a new write-only instance of Eth2Deposit, bound to a specific deployed contract.
func NewEth2DepositTransactor(address common.Address, transactor bind.ContractTransactor) (*Eth2DepositTransactor, error) {
	contract, err := bindEth2Deposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Eth2DepositTransactor{contract: contract}, nil
}

// NewEth2DepositFilterer creates a new log filterer instance of Eth2Deposit, bound to a specific deployed contract.
func NewEth2DepositFilterer(address common.Address, filterer bind.ContractFilterer) (*Eth2DepositFilterer, error) {
	contract, err := bindEth2Deposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Eth2DepositFilterer{contract: contract}, nil
}

// bindEth2Deposit binds a generic wrapper to an already deployed contract.
func bindEth2Deposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Eth2DepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth2Deposit *Eth2DepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth2Deposit.Contract.Eth2DepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth2Deposit *Eth2DepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Eth2DepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth2Deposit *Eth2DepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Eth2DepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth2Deposit *Eth2DepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth2Deposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth2Deposit *Eth2DepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth2Deposit *Eth2DepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.contract.Transact(opts, method, params...)
}

// DrainAddress is a free data retrieval call binding the contract method 0x8ba35cdf.
//
// Solidity: function drain_address() returns(address out)
func (_Eth2Deposit *Eth2DepositCaller) DrainAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Eth2Deposit.contract.Call(opts, &out, "drain_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DrainAddress is a free data retrieval call binding the contract method 0x8ba35cdf.
//
// Solidity: function drain_address() returns(address out)
func (_Eth2Deposit *Eth2DepositSession) DrainAddress() (common.Address, error) {
	return _Eth2Deposit.Contract.DrainAddress(&_Eth2Deposit.CallOpts)
}

// DrainAddress is a free data retrieval call binding the contract method 0x8ba35cdf.
//
// Solidity: function drain_address() returns(address out)
func (_Eth2Deposit *Eth2DepositCallerSession) DrainAddress() (common.Address, error) {
	return _Eth2Deposit.Contract.DrainAddress(&_Eth2Deposit.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() returns(bytes out)
func (_Eth2Deposit *Eth2DepositCaller) GetDepositCount(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Eth2Deposit.contract.Call(opts, &out, "get_deposit_count")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() returns(bytes out)
func (_Eth2Deposit *Eth2DepositSession) GetDepositCount() ([]byte, error) {
	return _Eth2Deposit.Contract.GetDepositCount(&_Eth2Deposit.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() returns(bytes out)
func (_Eth2Deposit *Eth2DepositCallerSession) GetDepositCount() ([]byte, error) {
	return _Eth2Deposit.Contract.GetDepositCount(&_Eth2Deposit.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() returns(bytes32 out)
func (_Eth2Deposit *Eth2DepositCaller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Eth2Deposit.contract.Call(opts, &out, "get_deposit_root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() returns(bytes32 out)
func (_Eth2Deposit *Eth2DepositSession) GetDepositRoot() ([32]byte, error) {
	return _Eth2Deposit.Contract.GetDepositRoot(&_Eth2Deposit.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() returns(bytes32 out)
func (_Eth2Deposit *Eth2DepositCallerSession) GetDepositRoot() ([32]byte, error) {
	return _Eth2Deposit.Contract.GetDepositRoot(&_Eth2Deposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) returns()
func (_Eth2Deposit *Eth2DepositTransactor) Deposit(opts *bind.TransactOpts, pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2Deposit.contract.Transact(opts, "deposit", pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) returns()
func (_Eth2Deposit *Eth2DepositSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Deposit(&_Eth2Deposit.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) returns()
func (_Eth2Deposit *Eth2DepositTransactorSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Deposit(&_Eth2Deposit.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Eth2Deposit *Eth2DepositTransactor) Drain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth2Deposit.contract.Transact(opts, "drain")
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Eth2Deposit *Eth2DepositSession) Drain() (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Drain(&_Eth2Deposit.TransactOpts)
}

// Drain is a paid mutator transaction binding the contract method 0x9890220b.
//
// Solidity: function drain() returns()
func (_Eth2Deposit *Eth2DepositTransactorSession) Drain() (*types.Transaction, error) {
	return _Eth2Deposit.Contract.Drain(&_Eth2Deposit.TransactOpts)
}

// Eth2DepositDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Eth2Deposit contract.
type Eth2DepositDepositEventIterator struct {
	Event *Eth2DepositDepositEvent // Event containing the contract specifics and raw log

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
func (it *Eth2DepositDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Eth2DepositDepositEvent)
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
		it.Event = new(Eth2DepositDepositEvent)
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
func (it *Eth2DepositDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Eth2DepositDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Eth2DepositDepositEvent represents a DepositEvent event raised by the Eth2Deposit contract.
type Eth2DepositDepositEvent struct {
	Pubkey                []byte
	WithdrawalCredentials []byte
	Amount                []byte
	Signature             []byte
	Index                 []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2Deposit *Eth2DepositFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*Eth2DepositDepositEventIterator, error) {

	logs, sub, err := _Eth2Deposit.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &Eth2DepositDepositEventIterator{contract: _Eth2Deposit.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2Deposit *Eth2DepositFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *Eth2DepositDepositEvent) (event.Subscription, error) {

	logs, sub, err := _Eth2Deposit.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Eth2DepositDepositEvent)
				if err := _Eth2Deposit.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2Deposit *Eth2DepositFilterer) ParseDepositEvent(log types.Log) (*Eth2DepositDepositEvent, error) {
	event := new(Eth2DepositDepositEvent)
	if err := _Eth2Deposit.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
