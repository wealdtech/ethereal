// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package resolvercontract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ResolverContractABI is the input ABI used to generate the binding from.
const ResolverContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"contentType\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"name\":\"ret\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"},{\"name\":\"y\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"}]"

// ResolverContract is an auto generated Go binding around an Ethereum contract.
type ResolverContract struct {
	ResolverContractCaller     // Read-only binding to the contract
	ResolverContractTransactor // Write-only binding to the contract
}

// ResolverContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverContractSession struct {
	Contract     *ResolverContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverContractCallerSession struct {
	Contract *ResolverContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolverContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverContractTransactorSession struct {
	Contract     *ResolverContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolverContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverContractRaw struct {
	Contract *ResolverContract // Generic contract binding to access the raw methods on
}

// ResolverContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverContractCallerRaw struct {
	Contract *ResolverContractCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverContractTransactorRaw struct {
	Contract *ResolverContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolverContract creates a new instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContract(address common.Address, backend bind.ContractBackend) (*ResolverContract, error) {
	contract, err := bindResolverContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolverContract{ResolverContractCaller: ResolverContractCaller{contract: contract}, ResolverContractTransactor: ResolverContractTransactor{contract: contract}}, nil
}

// NewResolverContractCaller creates a new read-only instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContractCaller(address common.Address, caller bind.ContractCaller) (*ResolverContractCaller, error) {
	contract, err := bindResolverContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverContractCaller{contract: contract}, nil
}

// NewResolverContractTransactor creates a new write-only instance of ResolverContract, bound to a specific deployed contract.
func NewResolverContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverContractTransactor, error) {
	contract, err := bindResolverContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ResolverContractTransactor{contract: contract}, nil
}

// bindResolverContract binds a generic wrapper to an already deployed contract.
func bindResolverContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverContract *ResolverContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverContract.Contract.ResolverContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverContract *ResolverContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverContract.Contract.ResolverContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverContract *ResolverContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverContract.Contract.ResolverContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolverContract *ResolverContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ResolverContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolverContract *ResolverContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolverContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolverContract *ResolverContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolverContract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	ret := new(struct {
		ContentType *big.Int
		Data        []byte
	})
	out := ret
	err := _ResolverContract.contract.Call(opts, out, "ABI", node, contentTypes)
	return *ret, err
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _ResolverContract.Contract.ABI(&_ResolverContract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(node bytes32, contentTypes uint256) constant returns(contentType uint256, data bytes)
func (_ResolverContract *ResolverContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (struct {
	ContentType *big.Int
	Data        []byte
}, error) {
	return _ResolverContract.Contract.ABI(&_ResolverContract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverContract.Contract.Addr(&_ResolverContract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_ResolverContract *ResolverContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ResolverContract.Contract.Addr(&_ResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverContract.Contract.Content(&_ResolverContract.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_ResolverContract *ResolverContractCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _ResolverContract.Contract.Content(&_ResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "name", node)
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractSession) Name(node [32]byte) (string, error) {
	return _ResolverContract.Contract.Name(&_ResolverContract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(node bytes32) constant returns(ret string)
func (_ResolverContract *ResolverContractCallerSession) Name(node [32]byte) (string, error) {
	return _ResolverContract.Contract.Name(&_ResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	ret := new(struct {
		X [32]byte
		Y [32]byte
	})
	out := ret
	err := _ResolverContract.contract.Call(opts, out, "pubkey", node)
	return *ret, err
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverContract.Contract.Pubkey(&_ResolverContract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(node bytes32) constant returns(x bytes32, y bytes32)
func (_ResolverContract *ResolverContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _ResolverContract.Contract.Pubkey(&_ResolverContract.CallOpts, node)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverContract.Contract.SupportsInterface(&_ResolverContract.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(interfaceID bytes4) constant returns(bool)
func (_ResolverContract *ResolverContractCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ResolverContract.Contract.SupportsInterface(&_ResolverContract.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ResolverContract.contract.Call(opts, out, "text", node, key)
	return *ret0, err
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverContract.Contract.Text(&_ResolverContract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(node bytes32, key string) constant returns(ret string)
func (_ResolverContract *ResolverContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _ResolverContract.Contract.Text(&_ResolverContract.CallOpts, node, key)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetABI(&_ResolverContract.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(node bytes32, contentType uint256, data bytes) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetABI(&_ResolverContract.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetAddr(&_ResolverContract.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetAddr(&_ResolverContract.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetContent(&_ResolverContract.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetContent(&_ResolverContract.TransactOpts, node, hash)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetName(&_ResolverContract.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(node bytes32, name string) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetName(&_ResolverContract.TransactOpts, node, name)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetPubkey(&_ResolverContract.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(node bytes32, x bytes32, y bytes32) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetPubkey(&_ResolverContract.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetText(&_ResolverContract.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(node bytes32, key string, value string) returns()
func (_ResolverContract *ResolverContractTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _ResolverContract.Contract.SetText(&_ResolverContract.TransactOpts, node, key, value)
}
