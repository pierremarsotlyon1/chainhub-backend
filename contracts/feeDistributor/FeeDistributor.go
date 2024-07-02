// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feeDistributor

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

// FeeDistributorMetaData contains all meta data concerning the FeeDistributor contract.
var FeeDistributorMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ToggleAllowCheckpointToken\",\"inputs\":[{\"name\":\"toggle_flag\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CheckpointToken\",\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Claimed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"claim_epoch\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"max_epoch\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_voting_escrow\",\"type\":\"address\"},{\"name\":\"_start_time\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_emergency_return\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"checkpoint_token\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ve_for_at\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"checkpoint_total_supply\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_many\",\"inputs\":[{\"name\":\"_receivers\",\"type\":\"address[20]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_admin\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_admin\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"toggle_allow_checkpoint_token\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"kill_me\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"recover_balance\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"start_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"time_cursor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"time_cursor_of\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_epoch_of\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_token_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens_per_week\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"voting_escrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_received\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token_last_balance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ve_supply\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"can_checkpoint_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"emergency_return\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]}]",
}

// FeeDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeDistributorMetaData.ABI instead.
var FeeDistributorABI = FeeDistributorMetaData.ABI

// FeeDistributor is an auto generated Go binding around an Ethereum contract.
type FeeDistributor struct {
	FeeDistributorCaller     // Read-only binding to the contract
	FeeDistributorTransactor // Write-only binding to the contract
	FeeDistributorFilterer   // Log filterer for contract events
}

// FeeDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeDistributorSession struct {
	Contract     *FeeDistributor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeDistributorCallerSession struct {
	Contract *FeeDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FeeDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeDistributorTransactorSession struct {
	Contract     *FeeDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FeeDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeDistributorRaw struct {
	Contract *FeeDistributor // Generic contract binding to access the raw methods on
}

// FeeDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeDistributorCallerRaw struct {
	Contract *FeeDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// FeeDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeDistributorTransactorRaw struct {
	Contract *FeeDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeDistributor creates a new instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributor(address common.Address, backend bind.ContractBackend) (*FeeDistributor, error) {
	contract, err := bindFeeDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeDistributor{FeeDistributorCaller: FeeDistributorCaller{contract: contract}, FeeDistributorTransactor: FeeDistributorTransactor{contract: contract}, FeeDistributorFilterer: FeeDistributorFilterer{contract: contract}}, nil
}

// NewFeeDistributorCaller creates a new read-only instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorCaller(address common.Address, caller bind.ContractCaller) (*FeeDistributorCaller, error) {
	contract, err := bindFeeDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorCaller{contract: contract}, nil
}

// NewFeeDistributorTransactor creates a new write-only instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeDistributorTransactor, error) {
	contract, err := bindFeeDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorTransactor{contract: contract}, nil
}

// NewFeeDistributorFilterer creates a new log filterer instance of FeeDistributor, bound to a specific deployed contract.
func NewFeeDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeDistributorFilterer, error) {
	contract, err := bindFeeDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorFilterer{contract: contract}, nil
}

// bindFeeDistributor binds a generic wrapper to an already deployed contract.
func bindFeeDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeDistributor *FeeDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeDistributor.Contract.FeeDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeDistributor *FeeDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.Contract.FeeDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeDistributor *FeeDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeDistributor.Contract.FeeDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeDistributor *FeeDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeDistributor *FeeDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeDistributor *FeeDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeDistributor.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FeeDistributor *FeeDistributorSession) Admin() (common.Address, error) {
	return _FeeDistributor.Contract.Admin(&_FeeDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) Admin() (common.Address, error) {
	return _FeeDistributor.Contract.Admin(&_FeeDistributor.CallOpts)
}

// CanCheckpointToken is a free data retrieval call binding the contract method 0xaeba4737.
//
// Solidity: function can_checkpoint_token() view returns(bool)
func (_FeeDistributor *FeeDistributorCaller) CanCheckpointToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "can_checkpoint_token")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCheckpointToken is a free data retrieval call binding the contract method 0xaeba4737.
//
// Solidity: function can_checkpoint_token() view returns(bool)
func (_FeeDistributor *FeeDistributorSession) CanCheckpointToken() (bool, error) {
	return _FeeDistributor.Contract.CanCheckpointToken(&_FeeDistributor.CallOpts)
}

// CanCheckpointToken is a free data retrieval call binding the contract method 0xaeba4737.
//
// Solidity: function can_checkpoint_token() view returns(bool)
func (_FeeDistributor *FeeDistributorCallerSession) CanCheckpointToken() (bool, error) {
	return _FeeDistributor.Contract.CanCheckpointToken(&_FeeDistributor.CallOpts)
}

// EmergencyReturn is a free data retrieval call binding the contract method 0x2c3f531e.
//
// Solidity: function emergency_return() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) EmergencyReturn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "emergency_return")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyReturn is a free data retrieval call binding the contract method 0x2c3f531e.
//
// Solidity: function emergency_return() view returns(address)
func (_FeeDistributor *FeeDistributorSession) EmergencyReturn() (common.Address, error) {
	return _FeeDistributor.Contract.EmergencyReturn(&_FeeDistributor.CallOpts)
}

// EmergencyReturn is a free data retrieval call binding the contract method 0x2c3f531e.
//
// Solidity: function emergency_return() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) EmergencyReturn() (common.Address, error) {
	return _FeeDistributor.Contract.EmergencyReturn(&_FeeDistributor.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FeeDistributor *FeeDistributorSession) FutureAdmin() (common.Address, error) {
	return _FeeDistributor.Contract.FutureAdmin(&_FeeDistributor.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) FutureAdmin() (common.Address, error) {
	return _FeeDistributor.Contract.FutureAdmin(&_FeeDistributor.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_FeeDistributor *FeeDistributorCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_FeeDistributor *FeeDistributorSession) IsKilled() (bool, error) {
	return _FeeDistributor.Contract.IsKilled(&_FeeDistributor.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_FeeDistributor *FeeDistributorCallerSession) IsKilled() (bool, error) {
	return _FeeDistributor.Contract.IsKilled(&_FeeDistributor.CallOpts)
}

// LastTokenTime is a free data retrieval call binding the contract method 0x7f58e8f8.
//
// Solidity: function last_token_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) LastTokenTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "last_token_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenTime is a free data retrieval call binding the contract method 0x7f58e8f8.
//
// Solidity: function last_token_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) LastTokenTime() (*big.Int, error) {
	return _FeeDistributor.Contract.LastTokenTime(&_FeeDistributor.CallOpts)
}

// LastTokenTime is a free data retrieval call binding the contract method 0x7f58e8f8.
//
// Solidity: function last_token_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) LastTokenTime() (*big.Int, error) {
	return _FeeDistributor.Contract.LastTokenTime(&_FeeDistributor.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "start_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) StartTime() (*big.Int, error) {
	return _FeeDistributor.Contract.StartTime(&_FeeDistributor.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) StartTime() (*big.Int, error) {
	return _FeeDistributor.Contract.StartTime(&_FeeDistributor.CallOpts)
}

// TimeCursor is a free data retrieval call binding the contract method 0x127dcbd3.
//
// Solidity: function time_cursor() view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) TimeCursor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "time_cursor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeCursor is a free data retrieval call binding the contract method 0x127dcbd3.
//
// Solidity: function time_cursor() view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) TimeCursor() (*big.Int, error) {
	return _FeeDistributor.Contract.TimeCursor(&_FeeDistributor.CallOpts)
}

// TimeCursor is a free data retrieval call binding the contract method 0x127dcbd3.
//
// Solidity: function time_cursor() view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) TimeCursor() (*big.Int, error) {
	return _FeeDistributor.Contract.TimeCursor(&_FeeDistributor.CallOpts)
}

// TimeCursorOf is a free data retrieval call binding the contract method 0x2a2a314b.
//
// Solidity: function time_cursor_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) TimeCursorOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "time_cursor_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeCursorOf is a free data retrieval call binding the contract method 0x2a2a314b.
//
// Solidity: function time_cursor_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) TimeCursorOf(arg0 common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.TimeCursorOf(&_FeeDistributor.CallOpts, arg0)
}

// TimeCursorOf is a free data retrieval call binding the contract method 0x2a2a314b.
//
// Solidity: function time_cursor_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) TimeCursorOf(arg0 common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.TimeCursorOf(&_FeeDistributor.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FeeDistributor *FeeDistributorSession) Token() (common.Address, error) {
	return _FeeDistributor.Contract.Token(&_FeeDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) Token() (common.Address, error) {
	return _FeeDistributor.Contract.Token(&_FeeDistributor.CallOpts)
}

// TokenLastBalance is a free data retrieval call binding the contract method 0x22b04bfc.
//
// Solidity: function token_last_balance() view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) TokenLastBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "token_last_balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenLastBalance is a free data retrieval call binding the contract method 0x22b04bfc.
//
// Solidity: function token_last_balance() view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) TokenLastBalance() (*big.Int, error) {
	return _FeeDistributor.Contract.TokenLastBalance(&_FeeDistributor.CallOpts)
}

// TokenLastBalance is a free data retrieval call binding the contract method 0x22b04bfc.
//
// Solidity: function token_last_balance() view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) TokenLastBalance() (*big.Int, error) {
	return _FeeDistributor.Contract.TokenLastBalance(&_FeeDistributor.CallOpts)
}

// TokensPerWeek is a free data retrieval call binding the contract method 0xedf59997.
//
// Solidity: function tokens_per_week(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) TokensPerWeek(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "tokens_per_week", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerWeek is a free data retrieval call binding the contract method 0xedf59997.
//
// Solidity: function tokens_per_week(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) TokensPerWeek(arg0 *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.TokensPerWeek(&_FeeDistributor.CallOpts, arg0)
}

// TokensPerWeek is a free data retrieval call binding the contract method 0xedf59997.
//
// Solidity: function tokens_per_week(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) TokensPerWeek(arg0 *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.TokensPerWeek(&_FeeDistributor.CallOpts, arg0)
}

// TotalReceived is a free data retrieval call binding the contract method 0x2f0c222e.
//
// Solidity: function total_received() view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) TotalReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "total_received")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReceived is a free data retrieval call binding the contract method 0x2f0c222e.
//
// Solidity: function total_received() view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) TotalReceived() (*big.Int, error) {
	return _FeeDistributor.Contract.TotalReceived(&_FeeDistributor.CallOpts)
}

// TotalReceived is a free data retrieval call binding the contract method 0x2f0c222e.
//
// Solidity: function total_received() view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) TotalReceived() (*big.Int, error) {
	return _FeeDistributor.Contract.TotalReceived(&_FeeDistributor.CallOpts)
}

// UserEpochOf is a free data retrieval call binding the contract method 0xd5d46e88.
//
// Solidity: function user_epoch_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) UserEpochOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "user_epoch_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserEpochOf is a free data retrieval call binding the contract method 0xd5d46e88.
//
// Solidity: function user_epoch_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) UserEpochOf(arg0 common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.UserEpochOf(&_FeeDistributor.CallOpts, arg0)
}

// UserEpochOf is a free data retrieval call binding the contract method 0xd5d46e88.
//
// Solidity: function user_epoch_of(address arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) UserEpochOf(arg0 common.Address) (*big.Int, error) {
	return _FeeDistributor.Contract.UserEpochOf(&_FeeDistributor.CallOpts, arg0)
}

// VeForAt is a free data retrieval call binding the contract method 0xace296fb.
//
// Solidity: function ve_for_at(address _user, uint256 _timestamp) view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) VeForAt(opts *bind.CallOpts, _user common.Address, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "ve_for_at", _user, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeForAt is a free data retrieval call binding the contract method 0xace296fb.
//
// Solidity: function ve_for_at(address _user, uint256 _timestamp) view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) VeForAt(_user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.VeForAt(&_FeeDistributor.CallOpts, _user, _timestamp)
}

// VeForAt is a free data retrieval call binding the contract method 0xace296fb.
//
// Solidity: function ve_for_at(address _user, uint256 _timestamp) view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) VeForAt(_user common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.VeForAt(&_FeeDistributor.CallOpts, _user, _timestamp)
}

// VeSupply is a free data retrieval call binding the contract method 0xd4dafba8.
//
// Solidity: function ve_supply(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCaller) VeSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "ve_supply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSupply is a free data retrieval call binding the contract method 0xd4dafba8.
//
// Solidity: function ve_supply(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorSession) VeSupply(arg0 *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.VeSupply(&_FeeDistributor.CallOpts, arg0)
}

// VeSupply is a free data retrieval call binding the contract method 0xd4dafba8.
//
// Solidity: function ve_supply(uint256 arg0) view returns(uint256)
func (_FeeDistributor *FeeDistributorCallerSession) VeSupply(arg0 *big.Int) (*big.Int, error) {
	return _FeeDistributor.Contract.VeSupply(&_FeeDistributor.CallOpts, arg0)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_FeeDistributor *FeeDistributorCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeDistributor.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_FeeDistributor *FeeDistributorSession) VotingEscrow() (common.Address, error) {
	return _FeeDistributor.Contract.VotingEscrow(&_FeeDistributor.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_FeeDistributor *FeeDistributorCallerSession) VotingEscrow() (common.Address, error) {
	return _FeeDistributor.Contract.VotingEscrow(&_FeeDistributor.CallOpts)
}

// ApplyAdmin is a paid mutator transaction binding the contract method 0xc0e991a6.
//
// Solidity: function apply_admin() returns()
func (_FeeDistributor *FeeDistributorTransactor) ApplyAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "apply_admin")
}

// ApplyAdmin is a paid mutator transaction binding the contract method 0xc0e991a6.
//
// Solidity: function apply_admin() returns()
func (_FeeDistributor *FeeDistributorSession) ApplyAdmin() (*types.Transaction, error) {
	return _FeeDistributor.Contract.ApplyAdmin(&_FeeDistributor.TransactOpts)
}

// ApplyAdmin is a paid mutator transaction binding the contract method 0xc0e991a6.
//
// Solidity: function apply_admin() returns()
func (_FeeDistributor *FeeDistributorTransactorSession) ApplyAdmin() (*types.Transaction, error) {
	return _FeeDistributor.Contract.ApplyAdmin(&_FeeDistributor.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorTransactor) Burn(opts *bind.TransactOpts, _coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "burn", _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.Burn(&_FeeDistributor.TransactOpts, _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorTransactorSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.Burn(&_FeeDistributor.TransactOpts, _coin)
}

// CheckpointToken is a paid mutator transaction binding the contract method 0x811a40fe.
//
// Solidity: function checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorTransactor) CheckpointToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "checkpoint_token")
}

// CheckpointToken is a paid mutator transaction binding the contract method 0x811a40fe.
//
// Solidity: function checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorSession) CheckpointToken() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CheckpointToken(&_FeeDistributor.TransactOpts)
}

// CheckpointToken is a paid mutator transaction binding the contract method 0x811a40fe.
//
// Solidity: function checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorTransactorSession) CheckpointToken() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CheckpointToken(&_FeeDistributor.TransactOpts)
}

// CheckpointTotalSupply is a paid mutator transaction binding the contract method 0xb21ed502.
//
// Solidity: function checkpoint_total_supply() returns()
func (_FeeDistributor *FeeDistributorTransactor) CheckpointTotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "checkpoint_total_supply")
}

// CheckpointTotalSupply is a paid mutator transaction binding the contract method 0xb21ed502.
//
// Solidity: function checkpoint_total_supply() returns()
func (_FeeDistributor *FeeDistributorSession) CheckpointTotalSupply() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CheckpointTotalSupply(&_FeeDistributor.TransactOpts)
}

// CheckpointTotalSupply is a paid mutator transaction binding the contract method 0xb21ed502.
//
// Solidity: function checkpoint_total_supply() returns()
func (_FeeDistributor *FeeDistributorTransactorSession) CheckpointTotalSupply() (*types.Transaction, error) {
	return _FeeDistributor.Contract.CheckpointTotalSupply(&_FeeDistributor.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_FeeDistributor *FeeDistributorTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_FeeDistributor *FeeDistributorSession) Claim() (*types.Transaction, error) {
	return _FeeDistributor.Contract.Claim(&_FeeDistributor.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_FeeDistributor *FeeDistributorTransactorSession) Claim() (*types.Transaction, error) {
	return _FeeDistributor.Contract.Claim(&_FeeDistributor.TransactOpts)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _addr) returns(uint256)
func (_FeeDistributor *FeeDistributorTransactor) Claim0(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "claim0", _addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _addr) returns(uint256)
func (_FeeDistributor *FeeDistributorSession) Claim0(_addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.Claim0(&_FeeDistributor.TransactOpts, _addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _addr) returns(uint256)
func (_FeeDistributor *FeeDistributorTransactorSession) Claim0(_addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.Claim0(&_FeeDistributor.TransactOpts, _addr)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x7b935a23.
//
// Solidity: function claim_many(address[20] _receivers) returns(bool)
func (_FeeDistributor *FeeDistributorTransactor) ClaimMany(opts *bind.TransactOpts, _receivers [20]common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "claim_many", _receivers)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x7b935a23.
//
// Solidity: function claim_many(address[20] _receivers) returns(bool)
func (_FeeDistributor *FeeDistributorSession) ClaimMany(_receivers [20]common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.ClaimMany(&_FeeDistributor.TransactOpts, _receivers)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x7b935a23.
//
// Solidity: function claim_many(address[20] _receivers) returns(bool)
func (_FeeDistributor *FeeDistributorTransactorSession) ClaimMany(_receivers [20]common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.ClaimMany(&_FeeDistributor.TransactOpts, _receivers)
}

// CommitAdmin is a paid mutator transaction binding the contract method 0xb1d3db75.
//
// Solidity: function commit_admin(address _addr) returns()
func (_FeeDistributor *FeeDistributorTransactor) CommitAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "commit_admin", _addr)
}

// CommitAdmin is a paid mutator transaction binding the contract method 0xb1d3db75.
//
// Solidity: function commit_admin(address _addr) returns()
func (_FeeDistributor *FeeDistributorSession) CommitAdmin(_addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.CommitAdmin(&_FeeDistributor.TransactOpts, _addr)
}

// CommitAdmin is a paid mutator transaction binding the contract method 0xb1d3db75.
//
// Solidity: function commit_admin(address _addr) returns()
func (_FeeDistributor *FeeDistributorTransactorSession) CommitAdmin(_addr common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.CommitAdmin(&_FeeDistributor.TransactOpts, _addr)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_FeeDistributor *FeeDistributorTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_FeeDistributor *FeeDistributorSession) KillMe() (*types.Transaction, error) {
	return _FeeDistributor.Contract.KillMe(&_FeeDistributor.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_FeeDistributor *FeeDistributorTransactorSession) KillMe() (*types.Transaction, error) {
	return _FeeDistributor.Contract.KillMe(&_FeeDistributor.TransactOpts)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0xdb2f5f79.
//
// Solidity: function recover_balance(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorTransactor) RecoverBalance(opts *bind.TransactOpts, _coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "recover_balance", _coin)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0xdb2f5f79.
//
// Solidity: function recover_balance(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorSession) RecoverBalance(_coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.RecoverBalance(&_FeeDistributor.TransactOpts, _coin)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0xdb2f5f79.
//
// Solidity: function recover_balance(address _coin) returns(bool)
func (_FeeDistributor *FeeDistributorTransactorSession) RecoverBalance(_coin common.Address) (*types.Transaction, error) {
	return _FeeDistributor.Contract.RecoverBalance(&_FeeDistributor.TransactOpts, _coin)
}

// ToggleAllowCheckpointToken is a paid mutator transaction binding the contract method 0x2121bfc3.
//
// Solidity: function toggle_allow_checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorTransactor) ToggleAllowCheckpointToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeDistributor.contract.Transact(opts, "toggle_allow_checkpoint_token")
}

// ToggleAllowCheckpointToken is a paid mutator transaction binding the contract method 0x2121bfc3.
//
// Solidity: function toggle_allow_checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorSession) ToggleAllowCheckpointToken() (*types.Transaction, error) {
	return _FeeDistributor.Contract.ToggleAllowCheckpointToken(&_FeeDistributor.TransactOpts)
}

// ToggleAllowCheckpointToken is a paid mutator transaction binding the contract method 0x2121bfc3.
//
// Solidity: function toggle_allow_checkpoint_token() returns()
func (_FeeDistributor *FeeDistributorTransactorSession) ToggleAllowCheckpointToken() (*types.Transaction, error) {
	return _FeeDistributor.Contract.ToggleAllowCheckpointToken(&_FeeDistributor.TransactOpts)
}

// FeeDistributorApplyAdminIterator is returned from FilterApplyAdmin and is used to iterate over the raw logs and unpacked data for ApplyAdmin events raised by the FeeDistributor contract.
type FeeDistributorApplyAdminIterator struct {
	Event *FeeDistributorApplyAdmin // Event containing the contract specifics and raw log

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
func (it *FeeDistributorApplyAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorApplyAdmin)
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
		it.Event = new(FeeDistributorApplyAdmin)
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
func (it *FeeDistributorApplyAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorApplyAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorApplyAdmin represents a ApplyAdmin event raised by the FeeDistributor contract.
type FeeDistributorApplyAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyAdmin is a free log retrieval operation binding the contract event 0x756f845176805c8ebf249854e909627308157f63c96e470e44a9e8549ba6fb1e.
//
// Solidity: event ApplyAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) FilterApplyAdmin(opts *bind.FilterOpts) (*FeeDistributorApplyAdminIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "ApplyAdmin")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorApplyAdminIterator{contract: _FeeDistributor.contract, event: "ApplyAdmin", logs: logs, sub: sub}, nil
}

// WatchApplyAdmin is a free log subscription operation binding the contract event 0x756f845176805c8ebf249854e909627308157f63c96e470e44a9e8549ba6fb1e.
//
// Solidity: event ApplyAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) WatchApplyAdmin(opts *bind.WatchOpts, sink chan<- *FeeDistributorApplyAdmin) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "ApplyAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorApplyAdmin)
				if err := _FeeDistributor.contract.UnpackLog(event, "ApplyAdmin", log); err != nil {
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

// ParseApplyAdmin is a log parse operation binding the contract event 0x756f845176805c8ebf249854e909627308157f63c96e470e44a9e8549ba6fb1e.
//
// Solidity: event ApplyAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) ParseApplyAdmin(log types.Log) (*FeeDistributorApplyAdmin, error) {
	event := new(FeeDistributorApplyAdmin)
	if err := _FeeDistributor.contract.UnpackLog(event, "ApplyAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorCheckpointTokenIterator is returned from FilterCheckpointToken and is used to iterate over the raw logs and unpacked data for CheckpointToken events raised by the FeeDistributor contract.
type FeeDistributorCheckpointTokenIterator struct {
	Event *FeeDistributorCheckpointToken // Event containing the contract specifics and raw log

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
func (it *FeeDistributorCheckpointTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorCheckpointToken)
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
		it.Event = new(FeeDistributorCheckpointToken)
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
func (it *FeeDistributorCheckpointTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorCheckpointTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorCheckpointToken represents a CheckpointToken event raised by the FeeDistributor contract.
type FeeDistributorCheckpointToken struct {
	Time   *big.Int
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCheckpointToken is a free log retrieval operation binding the contract event 0xce749457b74e10f393f2c6b1ce4261b78791376db5a3f501477a809f03f500d6.
//
// Solidity: event CheckpointToken(uint256 time, uint256 tokens)
func (_FeeDistributor *FeeDistributorFilterer) FilterCheckpointToken(opts *bind.FilterOpts) (*FeeDistributorCheckpointTokenIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "CheckpointToken")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorCheckpointTokenIterator{contract: _FeeDistributor.contract, event: "CheckpointToken", logs: logs, sub: sub}, nil
}

// WatchCheckpointToken is a free log subscription operation binding the contract event 0xce749457b74e10f393f2c6b1ce4261b78791376db5a3f501477a809f03f500d6.
//
// Solidity: event CheckpointToken(uint256 time, uint256 tokens)
func (_FeeDistributor *FeeDistributorFilterer) WatchCheckpointToken(opts *bind.WatchOpts, sink chan<- *FeeDistributorCheckpointToken) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "CheckpointToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorCheckpointToken)
				if err := _FeeDistributor.contract.UnpackLog(event, "CheckpointToken", log); err != nil {
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

// ParseCheckpointToken is a log parse operation binding the contract event 0xce749457b74e10f393f2c6b1ce4261b78791376db5a3f501477a809f03f500d6.
//
// Solidity: event CheckpointToken(uint256 time, uint256 tokens)
func (_FeeDistributor *FeeDistributorFilterer) ParseCheckpointToken(log types.Log) (*FeeDistributorCheckpointToken, error) {
	event := new(FeeDistributorCheckpointToken)
	if err := _FeeDistributor.contract.UnpackLog(event, "CheckpointToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the FeeDistributor contract.
type FeeDistributorClaimedIterator struct {
	Event *FeeDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *FeeDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorClaimed)
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
		it.Event = new(FeeDistributorClaimed)
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
func (it *FeeDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorClaimed represents a Claimed event raised by the FeeDistributor contract.
type FeeDistributorClaimed struct {
	Recipient  common.Address
	Amount     *big.Int
	ClaimEpoch *big.Int
	MaxEpoch   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed recipient, uint256 amount, uint256 claim_epoch, uint256 max_epoch)
func (_FeeDistributor *FeeDistributorFilterer) FilterClaimed(opts *bind.FilterOpts, recipient []common.Address) (*FeeDistributorClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "Claimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &FeeDistributorClaimedIterator{contract: _FeeDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed recipient, uint256 amount, uint256 claim_epoch, uint256 max_epoch)
func (_FeeDistributor *FeeDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *FeeDistributorClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "Claimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorClaimed)
				if err := _FeeDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e.
//
// Solidity: event Claimed(address indexed recipient, uint256 amount, uint256 claim_epoch, uint256 max_epoch)
func (_FeeDistributor *FeeDistributorFilterer) ParseClaimed(log types.Log) (*FeeDistributorClaimed, error) {
	event := new(FeeDistributorClaimed)
	if err := _FeeDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorCommitAdminIterator is returned from FilterCommitAdmin and is used to iterate over the raw logs and unpacked data for CommitAdmin events raised by the FeeDistributor contract.
type FeeDistributorCommitAdminIterator struct {
	Event *FeeDistributorCommitAdmin // Event containing the contract specifics and raw log

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
func (it *FeeDistributorCommitAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorCommitAdmin)
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
		it.Event = new(FeeDistributorCommitAdmin)
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
func (it *FeeDistributorCommitAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorCommitAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorCommitAdmin represents a CommitAdmin event raised by the FeeDistributor contract.
type FeeDistributorCommitAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitAdmin is a free log retrieval operation binding the contract event 0x59a407284ae6e2986675fa1400d6498af928ed01f4fd2dd6be4a2a8b4fc35b34.
//
// Solidity: event CommitAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) FilterCommitAdmin(opts *bind.FilterOpts) (*FeeDistributorCommitAdminIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "CommitAdmin")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorCommitAdminIterator{contract: _FeeDistributor.contract, event: "CommitAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitAdmin is a free log subscription operation binding the contract event 0x59a407284ae6e2986675fa1400d6498af928ed01f4fd2dd6be4a2a8b4fc35b34.
//
// Solidity: event CommitAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) WatchCommitAdmin(opts *bind.WatchOpts, sink chan<- *FeeDistributorCommitAdmin) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "CommitAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorCommitAdmin)
				if err := _FeeDistributor.contract.UnpackLog(event, "CommitAdmin", log); err != nil {
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

// ParseCommitAdmin is a log parse operation binding the contract event 0x59a407284ae6e2986675fa1400d6498af928ed01f4fd2dd6be4a2a8b4fc35b34.
//
// Solidity: event CommitAdmin(address admin)
func (_FeeDistributor *FeeDistributorFilterer) ParseCommitAdmin(log types.Log) (*FeeDistributorCommitAdmin, error) {
	event := new(FeeDistributorCommitAdmin)
	if err := _FeeDistributor.contract.UnpackLog(event, "CommitAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeDistributorToggleAllowCheckpointTokenIterator is returned from FilterToggleAllowCheckpointToken and is used to iterate over the raw logs and unpacked data for ToggleAllowCheckpointToken events raised by the FeeDistributor contract.
type FeeDistributorToggleAllowCheckpointTokenIterator struct {
	Event *FeeDistributorToggleAllowCheckpointToken // Event containing the contract specifics and raw log

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
func (it *FeeDistributorToggleAllowCheckpointTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeDistributorToggleAllowCheckpointToken)
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
		it.Event = new(FeeDistributorToggleAllowCheckpointToken)
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
func (it *FeeDistributorToggleAllowCheckpointTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeDistributorToggleAllowCheckpointTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeDistributorToggleAllowCheckpointToken represents a ToggleAllowCheckpointToken event raised by the FeeDistributor contract.
type FeeDistributorToggleAllowCheckpointToken struct {
	ToggleFlag bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterToggleAllowCheckpointToken is a free log retrieval operation binding the contract event 0xdbe6ac1081ebd8e648718341126659456f4009fcadfe1c23f66f5e61522610b2.
//
// Solidity: event ToggleAllowCheckpointToken(bool toggle_flag)
func (_FeeDistributor *FeeDistributorFilterer) FilterToggleAllowCheckpointToken(opts *bind.FilterOpts) (*FeeDistributorToggleAllowCheckpointTokenIterator, error) {

	logs, sub, err := _FeeDistributor.contract.FilterLogs(opts, "ToggleAllowCheckpointToken")
	if err != nil {
		return nil, err
	}
	return &FeeDistributorToggleAllowCheckpointTokenIterator{contract: _FeeDistributor.contract, event: "ToggleAllowCheckpointToken", logs: logs, sub: sub}, nil
}

// WatchToggleAllowCheckpointToken is a free log subscription operation binding the contract event 0xdbe6ac1081ebd8e648718341126659456f4009fcadfe1c23f66f5e61522610b2.
//
// Solidity: event ToggleAllowCheckpointToken(bool toggle_flag)
func (_FeeDistributor *FeeDistributorFilterer) WatchToggleAllowCheckpointToken(opts *bind.WatchOpts, sink chan<- *FeeDistributorToggleAllowCheckpointToken) (event.Subscription, error) {

	logs, sub, err := _FeeDistributor.contract.WatchLogs(opts, "ToggleAllowCheckpointToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeDistributorToggleAllowCheckpointToken)
				if err := _FeeDistributor.contract.UnpackLog(event, "ToggleAllowCheckpointToken", log); err != nil {
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

// ParseToggleAllowCheckpointToken is a log parse operation binding the contract event 0xdbe6ac1081ebd8e648718341126659456f4009fcadfe1c23f66f5e61522610b2.
//
// Solidity: event ToggleAllowCheckpointToken(bool toggle_flag)
func (_FeeDistributor *FeeDistributorFilterer) ParseToggleAllowCheckpointToken(log types.Log) (*FeeDistributorToggleAllowCheckpointToken, error) {
	event := new(FeeDistributorToggleAllowCheckpointToken)
	if err := _FeeDistributor.contract.UnpackLog(event, "ToggleAllowCheckpointToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
