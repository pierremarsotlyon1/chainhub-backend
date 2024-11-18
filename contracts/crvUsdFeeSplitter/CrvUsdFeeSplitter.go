// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdFeeSplitter

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Addr   common.Address
	Weight *big.Int
}

// CrvUsdFeeSplitterMetaData contains all meta data concerning the CrvUsdFeeSplitter contract.
var CrvUsdFeeSplitterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"SetReceivers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LivenessProtectionTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"FeeDispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previous_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"name\":\"new_owner\",\"type\":\"address\"}],\"name\":\"transfer_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounce_ownership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update_controllers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n_controllers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"allowed_controllers\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"controllers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dispatch_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"controllers\",\"type\":\"address[]\"}],\"name\":\"dispatch_fees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"receivers\",\"type\":\"tuple[]\"}],\"name\":\"set_receivers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excess_receiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n_receivers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"receivers\",\"outputs\":[{\"components\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_crvusd\",\"type\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\"},{\"components\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"receivers\",\"type\":\"tuple[]\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// CrvUsdFeeSplitterABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdFeeSplitterMetaData.ABI instead.
var CrvUsdFeeSplitterABI = CrvUsdFeeSplitterMetaData.ABI

// CrvUsdFeeSplitter is an auto generated Go binding around an Ethereum contract.
type CrvUsdFeeSplitter struct {
	CrvUsdFeeSplitterCaller     // Read-only binding to the contract
	CrvUsdFeeSplitterTransactor // Write-only binding to the contract
	CrvUsdFeeSplitterFilterer   // Log filterer for contract events
}

// CrvUsdFeeSplitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdFeeSplitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdFeeSplitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdFeeSplitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdFeeSplitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdFeeSplitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdFeeSplitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdFeeSplitterSession struct {
	Contract     *CrvUsdFeeSplitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrvUsdFeeSplitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdFeeSplitterCallerSession struct {
	Contract *CrvUsdFeeSplitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CrvUsdFeeSplitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdFeeSplitterTransactorSession struct {
	Contract     *CrvUsdFeeSplitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CrvUsdFeeSplitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdFeeSplitterRaw struct {
	Contract *CrvUsdFeeSplitter // Generic contract binding to access the raw methods on
}

// CrvUsdFeeSplitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdFeeSplitterCallerRaw struct {
	Contract *CrvUsdFeeSplitterCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdFeeSplitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdFeeSplitterTransactorRaw struct {
	Contract *CrvUsdFeeSplitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdFeeSplitter creates a new instance of CrvUsdFeeSplitter, bound to a specific deployed contract.
func NewCrvUsdFeeSplitter(address common.Address, backend bind.ContractBackend) (*CrvUsdFeeSplitter, error) {
	contract, err := bindCrvUsdFeeSplitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitter{CrvUsdFeeSplitterCaller: CrvUsdFeeSplitterCaller{contract: contract}, CrvUsdFeeSplitterTransactor: CrvUsdFeeSplitterTransactor{contract: contract}, CrvUsdFeeSplitterFilterer: CrvUsdFeeSplitterFilterer{contract: contract}}, nil
}

// NewCrvUsdFeeSplitterCaller creates a new read-only instance of CrvUsdFeeSplitter, bound to a specific deployed contract.
func NewCrvUsdFeeSplitterCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdFeeSplitterCaller, error) {
	contract, err := bindCrvUsdFeeSplitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterCaller{contract: contract}, nil
}

// NewCrvUsdFeeSplitterTransactor creates a new write-only instance of CrvUsdFeeSplitter, bound to a specific deployed contract.
func NewCrvUsdFeeSplitterTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdFeeSplitterTransactor, error) {
	contract, err := bindCrvUsdFeeSplitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterTransactor{contract: contract}, nil
}

// NewCrvUsdFeeSplitterFilterer creates a new log filterer instance of CrvUsdFeeSplitter, bound to a specific deployed contract.
func NewCrvUsdFeeSplitterFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdFeeSplitterFilterer, error) {
	contract, err := bindCrvUsdFeeSplitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterFilterer{contract: contract}, nil
}

// bindCrvUsdFeeSplitter binds a generic wrapper to an already deployed contract.
func bindCrvUsdFeeSplitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdFeeSplitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdFeeSplitter.Contract.CrvUsdFeeSplitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.CrvUsdFeeSplitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.CrvUsdFeeSplitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdFeeSplitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.contract.Transact(opts, method, params...)
}

// AllowedControllers is a free data retrieval call binding the contract method 0x04783ee4.
//
// Solidity: function allowed_controllers(address arg0) view returns(bool)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) AllowedControllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "allowed_controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedControllers is a free data retrieval call binding the contract method 0x04783ee4.
//
// Solidity: function allowed_controllers(address arg0) view returns(bool)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) AllowedControllers(arg0 common.Address) (bool, error) {
	return _CrvUsdFeeSplitter.Contract.AllowedControllers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// AllowedControllers is a free data retrieval call binding the contract method 0x04783ee4.
//
// Solidity: function allowed_controllers(address arg0) view returns(bool)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) AllowedControllers(arg0 common.Address) (bool, error) {
	return _CrvUsdFeeSplitter.Contract.AllowedControllers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) Controllers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.Controllers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.Controllers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// ExcessReceiver is a free data retrieval call binding the contract method 0x2e33f751.
//
// Solidity: function excess_receiver() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) ExcessReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "excess_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExcessReceiver is a free data retrieval call binding the contract method 0x2e33f751.
//
// Solidity: function excess_receiver() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) ExcessReceiver() (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.ExcessReceiver(&_CrvUsdFeeSplitter.CallOpts)
}

// ExcessReceiver is a free data retrieval call binding the contract method 0x2e33f751.
//
// Solidity: function excess_receiver() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) ExcessReceiver() (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.ExcessReceiver(&_CrvUsdFeeSplitter.CallOpts)
}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) NControllers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "n_controllers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) NControllers() (*big.Int, error) {
	return _CrvUsdFeeSplitter.Contract.NControllers(&_CrvUsdFeeSplitter.CallOpts)
}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) NControllers() (*big.Int, error) {
	return _CrvUsdFeeSplitter.Contract.NControllers(&_CrvUsdFeeSplitter.CallOpts)
}

// NReceivers is a free data retrieval call binding the contract method 0x632838bc.
//
// Solidity: function n_receivers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) NReceivers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "n_receivers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NReceivers is a free data retrieval call binding the contract method 0x632838bc.
//
// Solidity: function n_receivers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) NReceivers() (*big.Int, error) {
	return _CrvUsdFeeSplitter.Contract.NReceivers(&_CrvUsdFeeSplitter.CallOpts)
}

// NReceivers is a free data retrieval call binding the contract method 0x632838bc.
//
// Solidity: function n_receivers() view returns(uint256)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) NReceivers() (*big.Int, error) {
	return _CrvUsdFeeSplitter.Contract.NReceivers(&_CrvUsdFeeSplitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) Owner() (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.Owner(&_CrvUsdFeeSplitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) Owner() (common.Address, error) {
	return _CrvUsdFeeSplitter.Contract.Owner(&_CrvUsdFeeSplitter.CallOpts)
}

// Receivers is a free data retrieval call binding the contract method 0xbfd772fc.
//
// Solidity: function receivers(uint256 arg0) view returns((address,uint256))
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) Receivers(opts *bind.CallOpts, arg0 *big.Int) (Struct0, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "receivers", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// Receivers is a free data retrieval call binding the contract method 0xbfd772fc.
//
// Solidity: function receivers(uint256 arg0) view returns((address,uint256))
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) Receivers(arg0 *big.Int) (Struct0, error) {
	return _CrvUsdFeeSplitter.Contract.Receivers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// Receivers is a free data retrieval call binding the contract method 0xbfd772fc.
//
// Solidity: function receivers(uint256 arg0) view returns((address,uint256))
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) Receivers(arg0 *big.Int) (Struct0, error) {
	return _CrvUsdFeeSplitter.Contract.Receivers(&_CrvUsdFeeSplitter.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrvUsdFeeSplitter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) Version() (string, error) {
	return _CrvUsdFeeSplitter.Contract.Version(&_CrvUsdFeeSplitter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterCallerSession) Version() (string, error) {
	return _CrvUsdFeeSplitter.Contract.Version(&_CrvUsdFeeSplitter.CallOpts)
}

// DispatchFees is a paid mutator transaction binding the contract method 0xb8295111.
//
// Solidity: function dispatch_fees() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) DispatchFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "dispatch_fees")
}

// DispatchFees is a paid mutator transaction binding the contract method 0xb8295111.
//
// Solidity: function dispatch_fees() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) DispatchFees() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.DispatchFees(&_CrvUsdFeeSplitter.TransactOpts)
}

// DispatchFees is a paid mutator transaction binding the contract method 0xb8295111.
//
// Solidity: function dispatch_fees() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) DispatchFees() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.DispatchFees(&_CrvUsdFeeSplitter.TransactOpts)
}

// DispatchFees0 is a paid mutator transaction binding the contract method 0x1501e4d3.
//
// Solidity: function dispatch_fees(address[] controllers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) DispatchFees0(opts *bind.TransactOpts, controllers []common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "dispatch_fees0", controllers)
}

// DispatchFees0 is a paid mutator transaction binding the contract method 0x1501e4d3.
//
// Solidity: function dispatch_fees(address[] controllers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) DispatchFees0(controllers []common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.DispatchFees0(&_CrvUsdFeeSplitter.TransactOpts, controllers)
}

// DispatchFees0 is a paid mutator transaction binding the contract method 0x1501e4d3.
//
// Solidity: function dispatch_fees(address[] controllers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) DispatchFees0(controllers []common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.DispatchFees0(&_CrvUsdFeeSplitter.TransactOpts, controllers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "renounce_ownership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.RenounceOwnership(&_CrvUsdFeeSplitter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0xb15e13ee.
//
// Solidity: function renounce_ownership() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.RenounceOwnership(&_CrvUsdFeeSplitter.TransactOpts)
}

// SetReceivers is a paid mutator transaction binding the contract method 0x30ffa04e.
//
// Solidity: function set_receivers((address,uint256)[] receivers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) SetReceivers(opts *bind.TransactOpts, receivers []Struct0) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "set_receivers", receivers)
}

// SetReceivers is a paid mutator transaction binding the contract method 0x30ffa04e.
//
// Solidity: function set_receivers((address,uint256)[] receivers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) SetReceivers(receivers []Struct0) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.SetReceivers(&_CrvUsdFeeSplitter.TransactOpts, receivers)
}

// SetReceivers is a paid mutator transaction binding the contract method 0x30ffa04e.
//
// Solidity: function set_receivers((address,uint256)[] receivers) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) SetReceivers(receivers []Struct0) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.SetReceivers(&_CrvUsdFeeSplitter.TransactOpts, receivers)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) TransferOwnership(opts *bind.TransactOpts, new_owner common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "transfer_ownership", new_owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) TransferOwnership(new_owner common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.TransferOwnership(&_CrvUsdFeeSplitter.TransactOpts, new_owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf0350c04.
//
// Solidity: function transfer_ownership(address new_owner) returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) TransferOwnership(new_owner common.Address) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.TransferOwnership(&_CrvUsdFeeSplitter.TransactOpts, new_owner)
}

// UpdateControllers is a paid mutator transaction binding the contract method 0x44ba5664.
//
// Solidity: function update_controllers() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactor) UpdateControllers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.contract.Transact(opts, "update_controllers")
}

// UpdateControllers is a paid mutator transaction binding the contract method 0x44ba5664.
//
// Solidity: function update_controllers() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterSession) UpdateControllers() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.UpdateControllers(&_CrvUsdFeeSplitter.TransactOpts)
}

// UpdateControllers is a paid mutator transaction binding the contract method 0x44ba5664.
//
// Solidity: function update_controllers() returns()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterTransactorSession) UpdateControllers() (*types.Transaction, error) {
	return _CrvUsdFeeSplitter.Contract.UpdateControllers(&_CrvUsdFeeSplitter.TransactOpts)
}

// CrvUsdFeeSplitterFeeDispatchedIterator is returned from FilterFeeDispatched and is used to iterate over the raw logs and unpacked data for FeeDispatched events raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterFeeDispatchedIterator struct {
	Event *CrvUsdFeeSplitterFeeDispatched // Event containing the contract specifics and raw log

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
func (it *CrvUsdFeeSplitterFeeDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdFeeSplitterFeeDispatched)
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
		it.Event = new(CrvUsdFeeSplitterFeeDispatched)
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
func (it *CrvUsdFeeSplitterFeeDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdFeeSplitterFeeDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdFeeSplitterFeeDispatched represents a FeeDispatched event raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterFeeDispatched struct {
	Receiver common.Address
	Weight   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeDispatched is a free log retrieval operation binding the contract event 0x3ec7c36ff485aa9a27938503e3094604652d1f7262464127fb79577970abe12a.
//
// Solidity: event FeeDispatched(address indexed receiver, uint256 weight)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) FilterFeeDispatched(opts *bind.FilterOpts, receiver []common.Address) (*CrvUsdFeeSplitterFeeDispatchedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CrvUsdFeeSplitter.contract.FilterLogs(opts, "FeeDispatched", receiverRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterFeeDispatchedIterator{contract: _CrvUsdFeeSplitter.contract, event: "FeeDispatched", logs: logs, sub: sub}, nil
}

// WatchFeeDispatched is a free log subscription operation binding the contract event 0x3ec7c36ff485aa9a27938503e3094604652d1f7262464127fb79577970abe12a.
//
// Solidity: event FeeDispatched(address indexed receiver, uint256 weight)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) WatchFeeDispatched(opts *bind.WatchOpts, sink chan<- *CrvUsdFeeSplitterFeeDispatched, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CrvUsdFeeSplitter.contract.WatchLogs(opts, "FeeDispatched", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdFeeSplitterFeeDispatched)
				if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "FeeDispatched", log); err != nil {
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

// ParseFeeDispatched is a log parse operation binding the contract event 0x3ec7c36ff485aa9a27938503e3094604652d1f7262464127fb79577970abe12a.
//
// Solidity: event FeeDispatched(address indexed receiver, uint256 weight)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) ParseFeeDispatched(log types.Log) (*CrvUsdFeeSplitterFeeDispatched, error) {
	event := new(CrvUsdFeeSplitterFeeDispatched)
	if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "FeeDispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdFeeSplitterLivenessProtectionTriggeredIterator is returned from FilterLivenessProtectionTriggered and is used to iterate over the raw logs and unpacked data for LivenessProtectionTriggered events raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterLivenessProtectionTriggeredIterator struct {
	Event *CrvUsdFeeSplitterLivenessProtectionTriggered // Event containing the contract specifics and raw log

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
func (it *CrvUsdFeeSplitterLivenessProtectionTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdFeeSplitterLivenessProtectionTriggered)
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
		it.Event = new(CrvUsdFeeSplitterLivenessProtectionTriggered)
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
func (it *CrvUsdFeeSplitterLivenessProtectionTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdFeeSplitterLivenessProtectionTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdFeeSplitterLivenessProtectionTriggered represents a LivenessProtectionTriggered event raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterLivenessProtectionTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLivenessProtectionTriggered is a free log retrieval operation binding the contract event 0x7544693205d94fae4fc2b20449536d0486b497d16e6d7dcbaf967b8fc277d02c.
//
// Solidity: event LivenessProtectionTriggered()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) FilterLivenessProtectionTriggered(opts *bind.FilterOpts) (*CrvUsdFeeSplitterLivenessProtectionTriggeredIterator, error) {

	logs, sub, err := _CrvUsdFeeSplitter.contract.FilterLogs(opts, "LivenessProtectionTriggered")
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterLivenessProtectionTriggeredIterator{contract: _CrvUsdFeeSplitter.contract, event: "LivenessProtectionTriggered", logs: logs, sub: sub}, nil
}

// WatchLivenessProtectionTriggered is a free log subscription operation binding the contract event 0x7544693205d94fae4fc2b20449536d0486b497d16e6d7dcbaf967b8fc277d02c.
//
// Solidity: event LivenessProtectionTriggered()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) WatchLivenessProtectionTriggered(opts *bind.WatchOpts, sink chan<- *CrvUsdFeeSplitterLivenessProtectionTriggered) (event.Subscription, error) {

	logs, sub, err := _CrvUsdFeeSplitter.contract.WatchLogs(opts, "LivenessProtectionTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdFeeSplitterLivenessProtectionTriggered)
				if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "LivenessProtectionTriggered", log); err != nil {
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

// ParseLivenessProtectionTriggered is a log parse operation binding the contract event 0x7544693205d94fae4fc2b20449536d0486b497d16e6d7dcbaf967b8fc277d02c.
//
// Solidity: event LivenessProtectionTriggered()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) ParseLivenessProtectionTriggered(log types.Log) (*CrvUsdFeeSplitterLivenessProtectionTriggered, error) {
	event := new(CrvUsdFeeSplitterLivenessProtectionTriggered)
	if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "LivenessProtectionTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdFeeSplitterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterOwnershipTransferredIterator struct {
	Event *CrvUsdFeeSplitterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrvUsdFeeSplitterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdFeeSplitterOwnershipTransferred)
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
		it.Event = new(CrvUsdFeeSplitterOwnershipTransferred)
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
func (it *CrvUsdFeeSplitterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdFeeSplitterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdFeeSplitterOwnershipTransferred represents a OwnershipTransferred event raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previous_owner []common.Address, new_owner []common.Address) (*CrvUsdFeeSplitterOwnershipTransferredIterator, error) {

	var previous_ownerRule []interface{}
	for _, previous_ownerItem := range previous_owner {
		previous_ownerRule = append(previous_ownerRule, previous_ownerItem)
	}
	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _CrvUsdFeeSplitter.contract.FilterLogs(opts, "OwnershipTransferred", previous_ownerRule, new_ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterOwnershipTransferredIterator{contract: _CrvUsdFeeSplitter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrvUsdFeeSplitterOwnershipTransferred, previous_owner []common.Address, new_owner []common.Address) (event.Subscription, error) {

	var previous_ownerRule []interface{}
	for _, previous_ownerItem := range previous_owner {
		previous_ownerRule = append(previous_ownerRule, previous_ownerItem)
	}
	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _CrvUsdFeeSplitter.contract.WatchLogs(opts, "OwnershipTransferred", previous_ownerRule, new_ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdFeeSplitterOwnershipTransferred)
				if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previous_owner, address indexed new_owner)
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) ParseOwnershipTransferred(log types.Log) (*CrvUsdFeeSplitterOwnershipTransferred, error) {
	event := new(CrvUsdFeeSplitterOwnershipTransferred)
	if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdFeeSplitterSetReceiversIterator is returned from FilterSetReceivers and is used to iterate over the raw logs and unpacked data for SetReceivers events raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterSetReceiversIterator struct {
	Event *CrvUsdFeeSplitterSetReceivers // Event containing the contract specifics and raw log

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
func (it *CrvUsdFeeSplitterSetReceiversIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdFeeSplitterSetReceivers)
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
		it.Event = new(CrvUsdFeeSplitterSetReceivers)
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
func (it *CrvUsdFeeSplitterSetReceiversIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdFeeSplitterSetReceiversIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdFeeSplitterSetReceivers represents a SetReceivers event raised by the CrvUsdFeeSplitter contract.
type CrvUsdFeeSplitterSetReceivers struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetReceivers is a free log retrieval operation binding the contract event 0xcef5c0f74b8e26cac8a85442177d7e9b9792cc2b2627efce6a5c3d764fc34df1.
//
// Solidity: event SetReceivers()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) FilterSetReceivers(opts *bind.FilterOpts) (*CrvUsdFeeSplitterSetReceiversIterator, error) {

	logs, sub, err := _CrvUsdFeeSplitter.contract.FilterLogs(opts, "SetReceivers")
	if err != nil {
		return nil, err
	}
	return &CrvUsdFeeSplitterSetReceiversIterator{contract: _CrvUsdFeeSplitter.contract, event: "SetReceivers", logs: logs, sub: sub}, nil
}

// WatchSetReceivers is a free log subscription operation binding the contract event 0xcef5c0f74b8e26cac8a85442177d7e9b9792cc2b2627efce6a5c3d764fc34df1.
//
// Solidity: event SetReceivers()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) WatchSetReceivers(opts *bind.WatchOpts, sink chan<- *CrvUsdFeeSplitterSetReceivers) (event.Subscription, error) {

	logs, sub, err := _CrvUsdFeeSplitter.contract.WatchLogs(opts, "SetReceivers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdFeeSplitterSetReceivers)
				if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "SetReceivers", log); err != nil {
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

// ParseSetReceivers is a log parse operation binding the contract event 0xcef5c0f74b8e26cac8a85442177d7e9b9792cc2b2627efce6a5c3d764fc34df1.
//
// Solidity: event SetReceivers()
func (_CrvUsdFeeSplitter *CrvUsdFeeSplitterFilterer) ParseSetReceivers(log types.Log) (*CrvUsdFeeSplitterSetReceivers, error) {
	event := new(CrvUsdFeeSplitterSetReceivers)
	if err := _CrvUsdFeeSplitter.contract.UnpackLog(event, "SetReceivers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
