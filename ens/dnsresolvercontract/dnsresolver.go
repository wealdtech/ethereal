// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnsresolvercontract

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

// DNSResolverContractABI is the input ABI used to generate the binding from.
const DNSResolverContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSRecords\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"records\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_x\",\"type\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"}],\"name\":\"hasDNSRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setDNSZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_contentType\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"zones\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameEntriesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"dnsRecord\",\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"clearDNSZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"dnsZone\",\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"resource\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"resource\",\"type\":\"uint16\"}],\"name\":\"Deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

// DNSResolverContract is an auto generated Go binding around an Ethereum contract.
type DNSResolverContract struct {
	DNSResolverContractCaller     // Read-only binding to the contract
	DNSResolverContractTransactor // Write-only binding to the contract
	DNSResolverContractFilterer   // Log filterer for contract events
}

// DNSResolverContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNSResolverContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNSResolverContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNSResolverContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSResolverContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNSResolverContractSession struct {
	Contract     *DNSResolverContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DNSResolverContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNSResolverContractCallerSession struct {
	Contract *DNSResolverContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DNSResolverContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNSResolverContractTransactorSession struct {
	Contract     *DNSResolverContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DNSResolverContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DNSResolverContractRaw struct {
	Contract *DNSResolverContract // Generic contract binding to access the raw methods on
}

// DNSResolverContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNSResolverContractCallerRaw struct {
	Contract *DNSResolverContractCaller // Generic read-only contract binding to access the raw methods on
}

// DNSResolverContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNSResolverContractTransactorRaw struct {
	Contract *DNSResolverContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDNSResolverContract creates a new instance of DNSResolverContract, bound to a specific deployed contract.
func NewDNSResolverContract(address common.Address, backend bind.ContractBackend) (*DNSResolverContract, error) {
	contract, err := bindDNSResolverContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContract{DNSResolverContractCaller: DNSResolverContractCaller{contract: contract}, DNSResolverContractTransactor: DNSResolverContractTransactor{contract: contract}, DNSResolverContractFilterer: DNSResolverContractFilterer{contract: contract}}, nil
}

// NewDNSResolverContractCaller creates a new read-only instance of DNSResolverContract, bound to a specific deployed contract.
func NewDNSResolverContractCaller(address common.Address, caller bind.ContractCaller) (*DNSResolverContractCaller, error) {
	contract, err := bindDNSResolverContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractCaller{contract: contract}, nil
}

// NewDNSResolverContractTransactor creates a new write-only instance of DNSResolverContract, bound to a specific deployed contract.
func NewDNSResolverContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DNSResolverContractTransactor, error) {
	contract, err := bindDNSResolverContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractTransactor{contract: contract}, nil
}

// NewDNSResolverContractFilterer creates a new log filterer instance of DNSResolverContract, bound to a specific deployed contract.
func NewDNSResolverContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DNSResolverContractFilterer, error) {
	contract, err := bindDNSResolverContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractFilterer{contract: contract}, nil
}

// bindDNSResolverContract binds a generic wrapper to an already deployed contract.
func bindDNSResolverContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSResolverContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSResolverContract *DNSResolverContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSResolverContract.Contract.DNSResolverContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSResolverContract *DNSResolverContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.DNSResolverContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSResolverContract *DNSResolverContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.DNSResolverContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSResolverContract *DNSResolverContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSResolverContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSResolverContract *DNSResolverContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSResolverContract *DNSResolverContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DNSResolverContract *DNSResolverContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _DNSResolverContract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DNSResolverContract *DNSResolverContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DNSResolverContract.Contract.ABI(&_DNSResolverContract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_DNSResolverContract *DNSResolverContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _DNSResolverContract.Contract.ABI(&_DNSResolverContract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DNSResolverContract *DNSResolverContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DNSResolverContract *DNSResolverContractSession) Addr(node [32]byte) (common.Address, error) {
	return _DNSResolverContract.Contract.Addr(&_DNSResolverContract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_DNSResolverContract *DNSResolverContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _DNSResolverContract.Contract.Addr(&_DNSResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DNSResolverContract *DNSResolverContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DNSResolverContract *DNSResolverContractSession) Content(node [32]byte) ([32]byte, error) {
	return _DNSResolverContract.Contract.Content(&_DNSResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_DNSResolverContract *DNSResolverContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _DNSResolverContract.Contract.Content(&_DNSResolverContract.CallOpts, node)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractCaller) DnsRecord(opts *bind.CallOpts, node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "dnsRecord", node, name, resource)
	return *ret0, err
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _DNSResolverContract.Contract.DnsRecord(&_DNSResolverContract.CallOpts, node, name, resource)
}

// DnsRecord is a free data retrieval call binding the contract method 0xa8fa5682.
//
// Solidity: function dnsRecord(node bytes32, name bytes32, resource uint16) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractCallerSession) DnsRecord(node [32]byte, name [32]byte, resource uint16) ([]byte, error) {
	return _DNSResolverContract.Contract.DnsRecord(&_DNSResolverContract.CallOpts, node, name, resource)
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractCaller) DnsZone(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "dnsZone", node)
	return *ret0, err
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractSession) DnsZone(node [32]byte) ([]byte, error) {
	return _DNSResolverContract.Contract.DnsZone(&_DNSResolverContract.CallOpts, node)
}

// DnsZone is a free data retrieval call binding the contract method 0xdbfc5d00.
//
// Solidity: function dnsZone(node bytes32) constant returns(data bytes)
func (_DNSResolverContract *DNSResolverContractCallerSession) DnsZone(node [32]byte) ([]byte, error) {
	return _DNSResolverContract.Contract.DnsZone(&_DNSResolverContract.CallOpts, node)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(_node bytes32, x bytes32) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractCaller) HasDNSRecords(opts *bind.CallOpts, _node [32]byte, x [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "hasDNSRecords", _node, x)
	return *ret0, err
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(_node bytes32, x bytes32) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractSession) HasDNSRecords(_node [32]byte, x [32]byte) (bool, error) {
	return _DNSResolverContract.Contract.HasDNSRecords(&_DNSResolverContract.CallOpts, _node, x)
}

// HasDNSRecords is a free data retrieval call binding the contract method 0x4cbf6ba4.
//
// Solidity: function hasDNSRecords(_node bytes32, x bytes32) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractCallerSession) HasDNSRecords(_node [32]byte, x [32]byte) (bool, error) {
	return _DNSResolverContract.Contract.HasDNSRecords(&_DNSResolverContract.CallOpts, _node, x)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractSession) Name(node [32]byte) (string, error) {
	return _DNSResolverContract.Contract.Name(&_DNSResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractCallerSession) Name(node [32]byte) (string, error) {
	return _DNSResolverContract.Contract.Name(&_DNSResolverContract.CallOpts, node)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DNSResolverContract *DNSResolverContractCaller) NameEntriesCount(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "nameEntriesCount", arg0, arg1)
	return *ret0, err
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DNSResolverContract *DNSResolverContractSession) NameEntriesCount(arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	return _DNSResolverContract.Contract.NameEntriesCount(&_DNSResolverContract.CallOpts, arg0, arg1)
}

// NameEntriesCount is a free data retrieval call binding the contract method 0x932fc0e0.
//
// Solidity: function nameEntriesCount( bytes32,  bytes32) constant returns(uint16)
func (_DNSResolverContract *DNSResolverContractCallerSession) NameEntriesCount(arg0 [32]byte, arg1 [32]byte) (uint16, error) {
	return _DNSResolverContract.Contract.NameEntriesCount(&_DNSResolverContract.CallOpts, arg0, arg1)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DNSResolverContract *DNSResolverContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _DNSResolverContract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DNSResolverContract *DNSResolverContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DNSResolverContract.Contract.Pubkey(&_DNSResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_DNSResolverContract *DNSResolverContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _DNSResolverContract.Contract.Pubkey(&_DNSResolverContract.CallOpts, node)
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractCaller) Records(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "records", arg0, arg1, arg2)
	return *ret0, err
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractSession) Records(arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	return _DNSResolverContract.Contract.Records(&_DNSResolverContract.CallOpts, arg0, arg1, arg2)
}

// Records is a free data retrieval call binding the contract method 0x107de931.
//
// Solidity: function records( bytes32,  bytes32,  uint16) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractCallerSession) Records(arg0 [32]byte, arg1 [32]byte, arg2 uint16) ([]byte, error) {
	return _DNSResolverContract.Contract.Records(&_DNSResolverContract.CallOpts, arg0, arg1, arg2)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DNSResolverContract.Contract.SupportsInterface(&_DNSResolverContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceId bytes4) constant returns(bool)
func (_DNSResolverContract *DNSResolverContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DNSResolverContract.Contract.SupportsInterface(&_DNSResolverContract.CallOpts, interfaceId)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractSession) Text(node [32]byte, key string) (string, error) {
	return _DNSResolverContract.Contract.Text(&_DNSResolverContract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_DNSResolverContract *DNSResolverContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _DNSResolverContract.Contract.Text(&_DNSResolverContract.CallOpts, node, key)
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractCaller) Zones(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSResolverContract.contract.Call(opts, out, "zones", arg0)
	return *ret0, err
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractSession) Zones(arg0 [32]byte) ([]byte, error) {
	return _DNSResolverContract.Contract.Zones(&_DNSResolverContract.CallOpts, arg0)
}

// Zones is a free data retrieval call binding the contract method 0x6de75ff8.
//
// Solidity: function zones( bytes32) constant returns(bytes)
func (_DNSResolverContract *DNSResolverContractCallerSession) Zones(arg0 [32]byte) ([]byte, error) {
	return _DNSResolverContract.Contract.Zones(&_DNSResolverContract.CallOpts, arg0)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(_node bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) ClearDNSZone(opts *bind.TransactOpts, _node [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "clearDNSZone", _node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(_node bytes32) returns()
func (_DNSResolverContract *DNSResolverContractSession) ClearDNSZone(_node [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.ClearDNSZone(&_DNSResolverContract.TransactOpts, _node)
}

// ClearDNSZone is a paid mutator transaction binding the contract method 0xad5780af.
//
// Solidity: function clearDNSZone(_node bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) ClearDNSZone(_node [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.ClearDNSZone(&_DNSResolverContract.TransactOpts, _node)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(_node bytes32, _contentType uint256, _data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetABI(opts *bind.TransactOpts, _node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setABI", _node, _contentType, _data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(_node bytes32, _contentType uint256, _data bytes) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetABI(_node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetABI(&_DNSResolverContract.TransactOpts, _node, _contentType, _data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(_node bytes32, _contentType uint256, _data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetABI(_node [32]byte, _contentType *big.Int, _data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetABI(&_DNSResolverContract.TransactOpts, _node, _contentType, _data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(_node bytes32, _addr address) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetAddr(opts *bind.TransactOpts, _node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setAddr", _node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(_node bytes32, _addr address) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetAddr(_node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetAddr(&_DNSResolverContract.TransactOpts, _node, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(_node bytes32, _addr address) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetAddr(_node [32]byte, _addr common.Address) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetAddr(&_DNSResolverContract.TransactOpts, _node, _addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(_node bytes32, _hash bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetContent(opts *bind.TransactOpts, _node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setContent", _node, _hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(_node bytes32, _hash bytes32) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetContent(_node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetContent(&_DNSResolverContract.TransactOpts, _node, _hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(_node bytes32, _hash bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetContent(_node [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetContent(&_DNSResolverContract.TransactOpts, _node, _hash)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetDNSRecords(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setDNSRecords", node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetDNSRecords(&_DNSResolverContract.TransactOpts, node, data)
}

// SetDNSRecords is a paid mutator transaction binding the contract method 0x0af179d7.
//
// Solidity: function setDNSRecords(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetDNSRecords(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetDNSRecords(&_DNSResolverContract.TransactOpts, node, data)
}

// SetDNSZone is a paid mutator transaction binding the contract method 0x5d21719d.
//
// Solidity: function setDNSZone(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetDNSZone(opts *bind.TransactOpts, node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setDNSZone", node, data)
}

// SetDNSZone is a paid mutator transaction binding the contract method 0x5d21719d.
//
// Solidity: function setDNSZone(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetDNSZone(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetDNSZone(&_DNSResolverContract.TransactOpts, node, data)
}

// SetDNSZone is a paid mutator transaction binding the contract method 0x5d21719d.
//
// Solidity: function setDNSZone(node bytes32, data bytes) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetDNSZone(node [32]byte, data []byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetDNSZone(&_DNSResolverContract.TransactOpts, node, data)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(_node bytes32, _name string) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetName(opts *bind.TransactOpts, _node [32]byte, _name string) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setName", _node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(_node bytes32, _name string) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetName(_node [32]byte, _name string) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetName(&_DNSResolverContract.TransactOpts, _node, _name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(_node bytes32, _name string) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetName(_node [32]byte, _name string) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetName(&_DNSResolverContract.TransactOpts, _node, _name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(_node bytes32, _x bytes32, _y bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetPubkey(opts *bind.TransactOpts, _node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setPubkey", _node, _x, _y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(_node bytes32, _x bytes32, _y bytes32) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetPubkey(_node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetPubkey(&_DNSResolverContract.TransactOpts, _node, _x, _y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(_node bytes32, _x bytes32, _y bytes32) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetPubkey(_node [32]byte, _x [32]byte, _y [32]byte) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetPubkey(&_DNSResolverContract.TransactOpts, _node, _x, _y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(_node bytes32, _key string, _value string) returns()
func (_DNSResolverContract *DNSResolverContractTransactor) SetText(opts *bind.TransactOpts, _node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _DNSResolverContract.contract.Transact(opts, "setText", _node, _key, _value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(_node bytes32, _key string, _value string) returns()
func (_DNSResolverContract *DNSResolverContractSession) SetText(_node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetText(&_DNSResolverContract.TransactOpts, _node, _key, _value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(_node bytes32, _key string, _value string) returns()
func (_DNSResolverContract *DNSResolverContractTransactorSession) SetText(_node [32]byte, _key string, _value string) (*types.Transaction, error) {
	return _DNSResolverContract.Contract.SetText(&_DNSResolverContract.TransactOpts, _node, _key, _value)
}

// DNSResolverContractABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the DNSResolverContract contract.
type DNSResolverContractABIChangedIterator struct {
	Event *DNSResolverContractABIChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractABIChanged)
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
		it.Event = new(DNSResolverContractABIChanged)
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
func (it *DNSResolverContractABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractABIChanged represents a ABIChanged event raised by the DNSResolverContract contract.
type DNSResolverContractABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*DNSResolverContractABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractABIChangedIterator{contract: _DNSResolverContract.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: e ABIChanged(node indexed bytes32, contentType indexed uint256)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractABIChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// DNSResolverContractAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the DNSResolverContract contract.
type DNSResolverContractAddrChangedIterator struct {
	Event *DNSResolverContractAddrChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractAddrChanged)
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
		it.Event = new(DNSResolverContractAddrChanged)
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
func (it *DNSResolverContractAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractAddrChanged represents a AddrChanged event raised by the DNSResolverContract contract.
type DNSResolverContractAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*DNSResolverContractAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractAddrChangedIterator{contract: _DNSResolverContract.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractAddrChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// DNSResolverContractContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the DNSResolverContract contract.
type DNSResolverContractContentChangedIterator struct {
	Event *DNSResolverContractContentChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractContentChanged)
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
		it.Event = new(DNSResolverContractContentChanged)
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
func (it *DNSResolverContractContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractContentChanged represents a ContentChanged event raised by the DNSResolverContract contract.
type DNSResolverContractContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*DNSResolverContractContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractContentChangedIterator{contract: _DNSResolverContract.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractContentChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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

// DNSResolverContractDeletedIterator is returned from FilterDeleted and is used to iterate over the raw logs and unpacked data for Deleted events raised by the DNSResolverContract contract.
type DNSResolverContractDeletedIterator struct {
	Event *DNSResolverContractDeleted // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractDeleted)
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
		it.Event = new(DNSResolverContractDeleted)
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
func (it *DNSResolverContractDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractDeleted represents a Deleted event raised by the DNSResolverContract contract.
type DNSResolverContractDeleted struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleted is a free log retrieval operation binding the contract event 0x133052c72ea386f24d31f74751f618e877370038e43ae5a1571abd4e7039a10b.
//
// Solidity: e Deleted(node bytes32, name bytes, resource uint16)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterDeleted(opts *bind.FilterOpts) (*DNSResolverContractDeletedIterator, error) {

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "Deleted")
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractDeletedIterator{contract: _DNSResolverContract.contract, event: "Deleted", logs: logs, sub: sub}, nil
}

// WatchDeleted is a free log subscription operation binding the contract event 0x133052c72ea386f24d31f74751f618e877370038e43ae5a1571abd4e7039a10b.
//
// Solidity: e Deleted(node bytes32, name bytes, resource uint16)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchDeleted(opts *bind.WatchOpts, sink chan<- *DNSResolverContractDeleted) (event.Subscription, error) {

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "Deleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractDeleted)
				if err := _DNSResolverContract.contract.UnpackLog(event, "Deleted", log); err != nil {
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

// DNSResolverContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the DNSResolverContract contract.
type DNSResolverContractNameChangedIterator struct {
	Event *DNSResolverContractNameChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractNameChanged)
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
		it.Event = new(DNSResolverContractNameChanged)
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
func (it *DNSResolverContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractNameChanged represents a NameChanged event raised by the DNSResolverContract contract.
type DNSResolverContractNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*DNSResolverContractNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractNameChangedIterator{contract: _DNSResolverContract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: e NameChanged(node indexed bytes32, name string)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractNameChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// DNSResolverContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the DNSResolverContract contract.
type DNSResolverContractPubkeyChangedIterator struct {
	Event *DNSResolverContractPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractPubkeyChanged)
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
		it.Event = new(DNSResolverContractPubkeyChanged)
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
func (it *DNSResolverContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractPubkeyChanged represents a PubkeyChanged event raised by the DNSResolverContract contract.
type DNSResolverContractPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*DNSResolverContractPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractPubkeyChangedIterator{contract: _DNSResolverContract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: e PubkeyChanged(node indexed bytes32, x bytes32, y bytes32)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractPubkeyChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// DNSResolverContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the DNSResolverContract contract.
type DNSResolverContractTextChangedIterator struct {
	Event *DNSResolverContractTextChanged // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractTextChanged)
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
		it.Event = new(DNSResolverContractTextChanged)
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
func (it *DNSResolverContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractTextChanged represents a TextChanged event raised by the DNSResolverContract contract.
type DNSResolverContractTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*DNSResolverContractTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractTextChangedIterator{contract: _DNSResolverContract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: e TextChanged(node indexed bytes32, indexedKey indexed string, key string)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *DNSResolverContractTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractTextChanged)
				if err := _DNSResolverContract.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// DNSResolverContractUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the DNSResolverContract contract.
type DNSResolverContractUpdatedIterator struct {
	Event *DNSResolverContractUpdated // Event containing the contract specifics and raw log

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
func (it *DNSResolverContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSResolverContractUpdated)
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
		it.Event = new(DNSResolverContractUpdated)
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
func (it *DNSResolverContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSResolverContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSResolverContractUpdated represents a Updated event raised by the DNSResolverContract contract.
type DNSResolverContractUpdated struct {
	Node     [32]byte
	Name     []byte
	Resource uint16
	Length   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0x54c176550009b22baa8f09052da16ee02e1df250881ba30e944948669fe8482f.
//
// Solidity: e Updated(node bytes32, name bytes, resource uint16, length uint256)
func (_DNSResolverContract *DNSResolverContractFilterer) FilterUpdated(opts *bind.FilterOpts) (*DNSResolverContractUpdatedIterator, error) {

	logs, sub, err := _DNSResolverContract.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &DNSResolverContractUpdatedIterator{contract: _DNSResolverContract.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0x54c176550009b22baa8f09052da16ee02e1df250881ba30e944948669fe8482f.
//
// Solidity: e Updated(node bytes32, name bytes, resource uint16, length uint256)
func (_DNSResolverContract *DNSResolverContractFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *DNSResolverContractUpdated) (event.Subscription, error) {

	logs, sub, err := _DNSResolverContract.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSResolverContractUpdated)
				if err := _DNSResolverContract.contract.UnpackLog(event, "Updated", log); err != nil {
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
