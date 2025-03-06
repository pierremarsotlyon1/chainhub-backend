// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vmTokenFactory

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
)

// VmTokenFactoryMetaData contains all meta data concerning the VmTokenFactory contract.
var VmTokenFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mainChainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MinterAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrappedTokenDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VmTokenFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VmTokenFactoryMetaData.ABI instead.
var VmTokenFactoryABI = VmTokenFactoryMetaData.ABI

// VmTokenFactory is an auto generated Go binding around an Ethereum contract.
type VmTokenFactory struct {
	VmTokenFactoryCaller     // Read-only binding to the contract
	VmTokenFactoryTransactor // Write-only binding to the contract
	VmTokenFactoryFilterer   // Log filterer for contract events
}

// VmTokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmTokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmTokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmTokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmTokenFactorySession struct {
	Contract     *VmTokenFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmTokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmTokenFactoryCallerSession struct {
	Contract *VmTokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VmTokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmTokenFactoryTransactorSession struct {
	Contract     *VmTokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VmTokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmTokenFactoryRaw struct {
	Contract *VmTokenFactory // Generic contract binding to access the raw methods on
}

// VmTokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmTokenFactoryCallerRaw struct {
	Contract *VmTokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VmTokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmTokenFactoryTransactorRaw struct {
	Contract *VmTokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVmTokenFactory creates a new instance of VmTokenFactory, bound to a specific deployed contract.
func NewVmTokenFactory(address common.Address, backend bind.ContractBackend) (*VmTokenFactory, error) {
	contract, err := bindVmTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VmTokenFactory{VmTokenFactoryCaller: VmTokenFactoryCaller{contract: contract}, VmTokenFactoryTransactor: VmTokenFactoryTransactor{contract: contract}, VmTokenFactoryFilterer: VmTokenFactoryFilterer{contract: contract}}, nil
}

// NewVmTokenFactoryCaller creates a new read-only instance of VmTokenFactory, bound to a specific deployed contract.
func NewVmTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*VmTokenFactoryCaller, error) {
	contract, err := bindVmTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmTokenFactoryCaller{contract: contract}, nil
}

// NewVmTokenFactoryTransactor creates a new write-only instance of VmTokenFactory, bound to a specific deployed contract.
func NewVmTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VmTokenFactoryTransactor, error) {
	contract, err := bindVmTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmTokenFactoryTransactor{contract: contract}, nil
}

// NewVmTokenFactoryFilterer creates a new log filterer instance of VmTokenFactory, bound to a specific deployed contract.
func NewVmTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VmTokenFactoryFilterer, error) {
	contract, err := bindVmTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmTokenFactoryFilterer{contract: contract}, nil
}

// bindVmTokenFactory binds a generic wrapper to an already deployed contract.
func bindVmTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VmTokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmTokenFactory *VmTokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmTokenFactory.Contract.VmTokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmTokenFactory *VmTokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.VmTokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmTokenFactory *VmTokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.VmTokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmTokenFactory *VmTokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmTokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmTokenFactory *VmTokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmTokenFactory *VmTokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_VmTokenFactory *VmTokenFactoryCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_VmTokenFactory *VmTokenFactorySession) CHAINID() (*big.Int, error) {
	return _VmTokenFactory.Contract.CHAINID(&_VmTokenFactory.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_VmTokenFactory *VmTokenFactoryCallerSession) CHAINID() (*big.Int, error) {
	return _VmTokenFactory.Contract.CHAINID(&_VmTokenFactory.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(string name, string symbol, uint8 decimals)
func (_VmTokenFactory *VmTokenFactoryCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) (struct {
	Name     string
	Symbol   string
	Decimals uint8
}, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "getTokenMetadata", token)

	outstruct := new(struct {
		Name     string
		Symbol   string
		Decimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(string name, string symbol, uint8 decimals)
func (_VmTokenFactory *VmTokenFactorySession) GetTokenMetadata(token common.Address) (struct {
	Name     string
	Symbol   string
	Decimals uint8
}, error) {
	return _VmTokenFactory.Contract.GetTokenMetadata(&_VmTokenFactory.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(string name, string symbol, uint8 decimals)
func (_VmTokenFactory *VmTokenFactoryCallerSession) GetTokenMetadata(token common.Address) (struct {
	Name     string
	Symbol   string
	Decimals uint8
}, error) {
	return _VmTokenFactory.Contract.GetTokenMetadata(&_VmTokenFactory.CallOpts, token)
}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address ) view returns(bool)
func (_VmTokenFactory *VmTokenFactoryCaller) IsWrapped(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "isWrapped", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address ) view returns(bool)
func (_VmTokenFactory *VmTokenFactorySession) IsWrapped(arg0 common.Address) (bool, error) {
	return _VmTokenFactory.Contract.IsWrapped(&_VmTokenFactory.CallOpts, arg0)
}

// IsWrapped is a free data retrieval call binding the contract method 0x495ee13e.
//
// Solidity: function isWrapped(address ) view returns(bool)
func (_VmTokenFactory *VmTokenFactoryCallerSession) IsWrapped(arg0 common.Address) (bool, error) {
	return _VmTokenFactory.Contract.IsWrapped(&_VmTokenFactory.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_VmTokenFactory *VmTokenFactoryCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_VmTokenFactory *VmTokenFactorySession) Minter() (common.Address, error) {
	return _VmTokenFactory.Contract.Minter(&_VmTokenFactory.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_VmTokenFactory *VmTokenFactoryCallerSession) Minter() (common.Address, error) {
	return _VmTokenFactory.Contract.Minter(&_VmTokenFactory.CallOpts)
}

// NativeTokens is a free data retrieval call binding the contract method 0xc86726f6.
//
// Solidity: function nativeTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactoryCaller) NativeTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "nativeTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokens is a free data retrieval call binding the contract method 0xc86726f6.
//
// Solidity: function nativeTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactorySession) NativeTokens(arg0 common.Address) (common.Address, error) {
	return _VmTokenFactory.Contract.NativeTokens(&_VmTokenFactory.CallOpts, arg0)
}

// NativeTokens is a free data retrieval call binding the contract method 0xc86726f6.
//
// Solidity: function nativeTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactoryCallerSession) NativeTokens(arg0 common.Address) (common.Address, error) {
	return _VmTokenFactory.Contract.NativeTokens(&_VmTokenFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmTokenFactory *VmTokenFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmTokenFactory *VmTokenFactorySession) Owner() (common.Address, error) {
	return _VmTokenFactory.Contract.Owner(&_VmTokenFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmTokenFactory *VmTokenFactoryCallerSession) Owner() (common.Address, error) {
	return _VmTokenFactory.Contract.Owner(&_VmTokenFactory.CallOpts)
}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactoryCaller) WrappedTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VmTokenFactory.contract.Call(opts, &out, "wrappedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactorySession) WrappedTokens(arg0 common.Address) (common.Address, error) {
	return _VmTokenFactory.Contract.WrappedTokens(&_VmTokenFactory.CallOpts, arg0)
}

// WrappedTokens is a free data retrieval call binding the contract method 0xd5c6b504.
//
// Solidity: function wrappedTokens(address ) view returns(address)
func (_VmTokenFactory *VmTokenFactoryCallerSession) WrappedTokens(arg0 common.Address) (common.Address, error) {
	return _VmTokenFactory.Contract.WrappedTokens(&_VmTokenFactory.CallOpts, arg0)
}

// Burn is a paid mutator transaction binding the contract method 0xf6b911bc.
//
// Solidity: function burn(address nativeToken, address from, uint256 amount) returns()
func (_VmTokenFactory *VmTokenFactoryTransactor) Burn(opts *bind.TransactOpts, nativeToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmTokenFactory.contract.Transact(opts, "burn", nativeToken, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf6b911bc.
//
// Solidity: function burn(address nativeToken, address from, uint256 amount) returns()
func (_VmTokenFactory *VmTokenFactorySession) Burn(nativeToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.Burn(&_VmTokenFactory.TransactOpts, nativeToken, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf6b911bc.
//
// Solidity: function burn(address nativeToken, address from, uint256 amount) returns()
func (_VmTokenFactory *VmTokenFactoryTransactorSession) Burn(nativeToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.Burn(&_VmTokenFactory.TransactOpts, nativeToken, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x8b608899.
//
// Solidity: function mint(address nativeToken, address to, uint256 amount, string name, string symbol, uint8 decimals) returns()
func (_VmTokenFactory *VmTokenFactoryTransactor) Mint(opts *bind.TransactOpts, nativeToken common.Address, to common.Address, amount *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _VmTokenFactory.contract.Transact(opts, "mint", nativeToken, to, amount, name, symbol, decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x8b608899.
//
// Solidity: function mint(address nativeToken, address to, uint256 amount, string name, string symbol, uint8 decimals) returns()
func (_VmTokenFactory *VmTokenFactorySession) Mint(nativeToken common.Address, to common.Address, amount *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.Mint(&_VmTokenFactory.TransactOpts, nativeToken, to, amount, name, symbol, decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x8b608899.
//
// Solidity: function mint(address nativeToken, address to, uint256 amount, string name, string symbol, uint8 decimals) returns()
func (_VmTokenFactory *VmTokenFactoryTransactorSession) Mint(nativeToken common.Address, to common.Address, amount *big.Int, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.Mint(&_VmTokenFactory.TransactOpts, nativeToken, to, amount, name, symbol, decimals)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmTokenFactory *VmTokenFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmTokenFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmTokenFactory *VmTokenFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _VmTokenFactory.Contract.RenounceOwnership(&_VmTokenFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmTokenFactory *VmTokenFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VmTokenFactory.Contract.RenounceOwnership(&_VmTokenFactory.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns()
func (_VmTokenFactory *VmTokenFactoryTransactor) SetMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.contract.Transact(opts, "setMinter", newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns()
func (_VmTokenFactory *VmTokenFactorySession) SetMinter(newMinter common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.SetMinter(&_VmTokenFactory.TransactOpts, newMinter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address newMinter) returns()
func (_VmTokenFactory *VmTokenFactoryTransactorSession) SetMinter(newMinter common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.SetMinter(&_VmTokenFactory.TransactOpts, newMinter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmTokenFactory *VmTokenFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmTokenFactory *VmTokenFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.TransferOwnership(&_VmTokenFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmTokenFactory *VmTokenFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmTokenFactory.Contract.TransferOwnership(&_VmTokenFactory.TransactOpts, newOwner)
}

// VmTokenFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VmTokenFactory contract.
type VmTokenFactoryOwnershipTransferredIterator struct {
	Event *VmTokenFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VmTokenFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmTokenFactoryOwnershipTransferred)
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
		it.Event = new(VmTokenFactoryOwnershipTransferred)
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
func (it *VmTokenFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmTokenFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmTokenFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the VmTokenFactory contract.
type VmTokenFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VmTokenFactory *VmTokenFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VmTokenFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmTokenFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VmTokenFactoryOwnershipTransferredIterator{contract: _VmTokenFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VmTokenFactory *VmTokenFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VmTokenFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmTokenFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmTokenFactoryOwnershipTransferred)
				if err := _VmTokenFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VmTokenFactory *VmTokenFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VmTokenFactoryOwnershipTransferred, error) {
	event := new(VmTokenFactoryOwnershipTransferred)
	if err := _VmTokenFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
