// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feeCollector

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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Coin   common.Address
	To     common.Address
	Amount *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Coin   common.Address
	Amount *big.Int
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	HookId uint8
	Value  *big.Int
	Data   []byte
}

// FeeCollectorMetaData contains all meta data concerning the FeeCollector contract.
var FeeCollectorMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"SetMaxFee\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"max_fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetBurner\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetHooker\",\"inputs\":[{\"name\":\"hooker\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetKilled\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"epoch_mask\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetEmergencyOwner\",\"inputs\":[{\"name\":\"emergency_owner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_target_coin\",\"type\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_emergency_owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_many\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[{\"name\":\"ts\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epoch_time_frame\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epoch_time_frame\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_ts\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_transfers\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"coin\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"collect\",\"inputs\":[{\"name\":\"_coins\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"collect\",\"inputs\":[{\"name\":\"_coins\",\"type\":\"address[]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"can_exchange\",\"inputs\":[{\"name\":\"_coins\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"_hook_inputs\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"hook_id\",\"type\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"_hook_inputs\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"hook_id\",\"type\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}]},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"recover\",\"inputs\":[{\"name\":\"_recovers\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"coin\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}]},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_max_fee\",\"inputs\":[{\"name\":\"_epoch\",\"type\":\"uint256\"},{\"name\":\"_max_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_burner\",\"inputs\":[{\"name\":\"_new_burner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_hooker\",\"inputs\":[{\"name\":\"_new_hooker\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_target\",\"inputs\":[{\"name\":\"_new_target\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_killed\",\"inputs\":[{\"name\":\"_input\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"coin\",\"type\":\"address\"},{\"name\":\"killed\",\"type\":\"uint256\"}]}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_owner\",\"inputs\":[{\"name\":\"_new_owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_emergency_owner\",\"inputs\":[{\"name\":\"_new_owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"target\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_fee\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"burner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"hooker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"emergency_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// FeeCollectorABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeCollectorMetaData.ABI instead.
var FeeCollectorABI = FeeCollectorMetaData.ABI

// FeeCollector is an auto generated Go binding around an Ethereum contract.
type FeeCollector struct {
	FeeCollectorCaller     // Read-only binding to the contract
	FeeCollectorTransactor // Write-only binding to the contract
	FeeCollectorFilterer   // Log filterer for contract events
}

// FeeCollectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeCollectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeCollectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeCollectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeCollectorSession struct {
	Contract     *FeeCollector     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeCollectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeCollectorCallerSession struct {
	Contract *FeeCollectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FeeCollectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeCollectorTransactorSession struct {
	Contract     *FeeCollectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FeeCollectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeCollectorRaw struct {
	Contract *FeeCollector // Generic contract binding to access the raw methods on
}

// FeeCollectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeCollectorCallerRaw struct {
	Contract *FeeCollectorCaller // Generic read-only contract binding to access the raw methods on
}

// FeeCollectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeCollectorTransactorRaw struct {
	Contract *FeeCollectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeCollector creates a new instance of FeeCollector, bound to a specific deployed contract.
func NewFeeCollector(address common.Address, backend bind.ContractBackend) (*FeeCollector, error) {
	contract, err := bindFeeCollector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeCollector{FeeCollectorCaller: FeeCollectorCaller{contract: contract}, FeeCollectorTransactor: FeeCollectorTransactor{contract: contract}, FeeCollectorFilterer: FeeCollectorFilterer{contract: contract}}, nil
}

// NewFeeCollectorCaller creates a new read-only instance of FeeCollector, bound to a specific deployed contract.
func NewFeeCollectorCaller(address common.Address, caller bind.ContractCaller) (*FeeCollectorCaller, error) {
	contract, err := bindFeeCollector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorCaller{contract: contract}, nil
}

// NewFeeCollectorTransactor creates a new write-only instance of FeeCollector, bound to a specific deployed contract.
func NewFeeCollectorTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeCollectorTransactor, error) {
	contract, err := bindFeeCollector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorTransactor{contract: contract}, nil
}

// NewFeeCollectorFilterer creates a new log filterer instance of FeeCollector, bound to a specific deployed contract.
func NewFeeCollectorFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeCollectorFilterer, error) {
	contract, err := bindFeeCollector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorFilterer{contract: contract}, nil
}

// bindFeeCollector binds a generic wrapper to an already deployed contract.
func bindFeeCollector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeCollectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCollector *FeeCollectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCollector.Contract.FeeCollectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCollector *FeeCollectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCollector.Contract.FeeCollectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCollector *FeeCollectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCollector.Contract.FeeCollectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCollector *FeeCollectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCollector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCollector *FeeCollectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCollector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCollector *FeeCollectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCollector.Contract.contract.Transact(opts, method, params...)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_FeeCollector *FeeCollectorCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "burner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_FeeCollector *FeeCollectorSession) Burner() (common.Address, error) {
	return _FeeCollector.Contract.Burner(&_FeeCollector.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_FeeCollector *FeeCollectorCallerSession) Burner() (common.Address, error) {
	return _FeeCollector.Contract.Burner(&_FeeCollector.CallOpts)
}

// CanExchange is a free data retrieval call binding the contract method 0xdc7168e9.
//
// Solidity: function can_exchange(address[] _coins) view returns(bool)
func (_FeeCollector *FeeCollectorCaller) CanExchange(opts *bind.CallOpts, _coins []common.Address) (bool, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "can_exchange", _coins)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExchange is a free data retrieval call binding the contract method 0xdc7168e9.
//
// Solidity: function can_exchange(address[] _coins) view returns(bool)
func (_FeeCollector *FeeCollectorSession) CanExchange(_coins []common.Address) (bool, error) {
	return _FeeCollector.Contract.CanExchange(&_FeeCollector.CallOpts, _coins)
}

// CanExchange is a free data retrieval call binding the contract method 0xdc7168e9.
//
// Solidity: function can_exchange(address[] _coins) view returns(bool)
func (_FeeCollector *FeeCollectorCallerSession) CanExchange(_coins []common.Address) (bool, error) {
	return _FeeCollector.Contract.CanExchange(&_FeeCollector.CallOpts, _coins)
}

// EmergencyOwner is a free data retrieval call binding the contract method 0x63a4042a.
//
// Solidity: function emergency_owner() view returns(address)
func (_FeeCollector *FeeCollectorCaller) EmergencyOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "emergency_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyOwner is a free data retrieval call binding the contract method 0x63a4042a.
//
// Solidity: function emergency_owner() view returns(address)
func (_FeeCollector *FeeCollectorSession) EmergencyOwner() (common.Address, error) {
	return _FeeCollector.Contract.EmergencyOwner(&_FeeCollector.CallOpts)
}

// EmergencyOwner is a free data retrieval call binding the contract method 0x63a4042a.
//
// Solidity: function emergency_owner() view returns(address)
func (_FeeCollector *FeeCollectorCallerSession) EmergencyOwner() (common.Address, error) {
	return _FeeCollector.Contract.EmergencyOwner(&_FeeCollector.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FeeCollector *FeeCollectorSession) Epoch() (*big.Int, error) {
	return _FeeCollector.Contract.Epoch(&_FeeCollector.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) Epoch() (*big.Int, error) {
	return _FeeCollector.Contract.Epoch(&_FeeCollector.CallOpts)
}

// Epoch0 is a free data retrieval call binding the contract method 0x5487c577.
//
// Solidity: function epoch(uint256 ts) view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) Epoch0(opts *bind.CallOpts, ts *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "epoch0", ts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch0 is a free data retrieval call binding the contract method 0x5487c577.
//
// Solidity: function epoch(uint256 ts) view returns(uint256)
func (_FeeCollector *FeeCollectorSession) Epoch0(ts *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Epoch0(&_FeeCollector.CallOpts, ts)
}

// Epoch0 is a free data retrieval call binding the contract method 0x5487c577.
//
// Solidity: function epoch(uint256 ts) view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) Epoch0(ts *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Epoch0(&_FeeCollector.CallOpts, ts)
}

// EpochTimeFrame is a free data retrieval call binding the contract method 0xdffe059c.
//
// Solidity: function epoch_time_frame(uint256 _epoch) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorCaller) EpochTimeFrame(opts *bind.CallOpts, _epoch *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "epoch_time_frame", _epoch)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// EpochTimeFrame is a free data retrieval call binding the contract method 0xdffe059c.
//
// Solidity: function epoch_time_frame(uint256 _epoch) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorSession) EpochTimeFrame(_epoch *big.Int) (*big.Int, *big.Int, error) {
	return _FeeCollector.Contract.EpochTimeFrame(&_FeeCollector.CallOpts, _epoch)
}

// EpochTimeFrame is a free data retrieval call binding the contract method 0xdffe059c.
//
// Solidity: function epoch_time_frame(uint256 _epoch) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorCallerSession) EpochTimeFrame(_epoch *big.Int) (*big.Int, *big.Int, error) {
	return _FeeCollector.Contract.EpochTimeFrame(&_FeeCollector.CallOpts, _epoch)
}

// EpochTimeFrame0 is a free data retrieval call binding the contract method 0xb676f748.
//
// Solidity: function epoch_time_frame(uint256 _epoch, uint256 _ts) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorCaller) EpochTimeFrame0(opts *bind.CallOpts, _epoch *big.Int, _ts *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "epoch_time_frame0", _epoch, _ts)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// EpochTimeFrame0 is a free data retrieval call binding the contract method 0xb676f748.
//
// Solidity: function epoch_time_frame(uint256 _epoch, uint256 _ts) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorSession) EpochTimeFrame0(_epoch *big.Int, _ts *big.Int) (*big.Int, *big.Int, error) {
	return _FeeCollector.Contract.EpochTimeFrame0(&_FeeCollector.CallOpts, _epoch, _ts)
}

// EpochTimeFrame0 is a free data retrieval call binding the contract method 0xb676f748.
//
// Solidity: function epoch_time_frame(uint256 _epoch, uint256 _ts) view returns(uint256, uint256)
func (_FeeCollector *FeeCollectorCallerSession) EpochTimeFrame0(_epoch *big.Int, _ts *big.Int) (*big.Int, *big.Int, error) {
	return _FeeCollector.Contract.EpochTimeFrame0(&_FeeCollector.CallOpts, _epoch, _ts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeCollector *FeeCollectorSession) Fee() (*big.Int, error) {
	return _FeeCollector.Contract.Fee(&_FeeCollector.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) Fee() (*big.Int, error) {
	return _FeeCollector.Contract.Fee(&_FeeCollector.CallOpts)
}

// Fee0 is a free data retrieval call binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 _epoch) view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) Fee0(opts *bind.CallOpts, _epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "fee0", _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee0 is a free data retrieval call binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 _epoch) view returns(uint256)
func (_FeeCollector *FeeCollectorSession) Fee0(_epoch *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Fee0(&_FeeCollector.CallOpts, _epoch)
}

// Fee0 is a free data retrieval call binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 _epoch) view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) Fee0(_epoch *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Fee0(&_FeeCollector.CallOpts, _epoch)
}

// Fee1 is a free data retrieval call binding the contract method 0x939f5ea4.
//
// Solidity: function fee(uint256 _epoch, uint256 _ts) view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) Fee1(opts *bind.CallOpts, _epoch *big.Int, _ts *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "fee1", _epoch, _ts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee1 is a free data retrieval call binding the contract method 0x939f5ea4.
//
// Solidity: function fee(uint256 _epoch, uint256 _ts) view returns(uint256)
func (_FeeCollector *FeeCollectorSession) Fee1(_epoch *big.Int, _ts *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Fee1(&_FeeCollector.CallOpts, _epoch, _ts)
}

// Fee1 is a free data retrieval call binding the contract method 0x939f5ea4.
//
// Solidity: function fee(uint256 _epoch, uint256 _ts) view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) Fee1(_epoch *big.Int, _ts *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.Fee1(&_FeeCollector.CallOpts, _epoch, _ts)
}

// Hooker is a free data retrieval call binding the contract method 0x6dfaf772.
//
// Solidity: function hooker() view returns(address)
func (_FeeCollector *FeeCollectorCaller) Hooker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "hooker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hooker is a free data retrieval call binding the contract method 0x6dfaf772.
//
// Solidity: function hooker() view returns(address)
func (_FeeCollector *FeeCollectorSession) Hooker() (common.Address, error) {
	return _FeeCollector.Contract.Hooker(&_FeeCollector.CallOpts)
}

// Hooker is a free data retrieval call binding the contract method 0x6dfaf772.
//
// Solidity: function hooker() view returns(address)
func (_FeeCollector *FeeCollectorCallerSession) Hooker() (common.Address, error) {
	return _FeeCollector.Contract.Hooker(&_FeeCollector.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x59e63166.
//
// Solidity: function is_killed(address arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) IsKilled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "is_killed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x59e63166.
//
// Solidity: function is_killed(address arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorSession) IsKilled(arg0 common.Address) (*big.Int, error) {
	return _FeeCollector.Contract.IsKilled(&_FeeCollector.CallOpts, arg0)
}

// IsKilled is a free data retrieval call binding the contract method 0x59e63166.
//
// Solidity: function is_killed(address arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) IsKilled(arg0 common.Address) (*big.Int, error) {
	return _FeeCollector.Contract.IsKilled(&_FeeCollector.CallOpts, arg0)
}

// MaxFee is a free data retrieval call binding the contract method 0xc47339c2.
//
// Solidity: function max_fee(uint256 arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorCaller) MaxFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "max_fee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFee is a free data retrieval call binding the contract method 0xc47339c2.
//
// Solidity: function max_fee(uint256 arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorSession) MaxFee(arg0 *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.MaxFee(&_FeeCollector.CallOpts, arg0)
}

// MaxFee is a free data retrieval call binding the contract method 0xc47339c2.
//
// Solidity: function max_fee(uint256 arg0) view returns(uint256)
func (_FeeCollector *FeeCollectorCallerSession) MaxFee(arg0 *big.Int) (*big.Int, error) {
	return _FeeCollector.Contract.MaxFee(&_FeeCollector.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCollector *FeeCollectorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCollector *FeeCollectorSession) Owner() (common.Address, error) {
	return _FeeCollector.Contract.Owner(&_FeeCollector.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCollector *FeeCollectorCallerSession) Owner() (common.Address, error) {
	return _FeeCollector.Contract.Owner(&_FeeCollector.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_FeeCollector *FeeCollectorCaller) Target(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCollector.contract.Call(opts, &out, "target")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_FeeCollector *FeeCollectorSession) Target() (common.Address, error) {
	return _FeeCollector.Contract.Target(&_FeeCollector.CallOpts)
}

// Target is a free data retrieval call binding the contract method 0xd4b83992.
//
// Solidity: function target() view returns(address)
func (_FeeCollector *FeeCollectorCallerSession) Target() (common.Address, error) {
	return _FeeCollector.Contract.Target(&_FeeCollector.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) payable returns(bool)
func (_FeeCollector *FeeCollectorTransactor) Burn(opts *bind.TransactOpts, _coin common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "burn", _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) payable returns(bool)
func (_FeeCollector *FeeCollectorSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Burn(&_FeeCollector.TransactOpts, _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) payable returns(bool)
func (_FeeCollector *FeeCollectorTransactorSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Burn(&_FeeCollector.TransactOpts, _coin)
}

// Collect is a paid mutator transaction binding the contract method 0xa4520aee.
//
// Solidity: function collect(address[] _coins) returns()
func (_FeeCollector *FeeCollectorTransactor) Collect(opts *bind.TransactOpts, _coins []common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "collect", _coins)
}

// Collect is a paid mutator transaction binding the contract method 0xa4520aee.
//
// Solidity: function collect(address[] _coins) returns()
func (_FeeCollector *FeeCollectorSession) Collect(_coins []common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Collect(&_FeeCollector.TransactOpts, _coins)
}

// Collect is a paid mutator transaction binding the contract method 0xa4520aee.
//
// Solidity: function collect(address[] _coins) returns()
func (_FeeCollector *FeeCollectorTransactorSession) Collect(_coins []common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Collect(&_FeeCollector.TransactOpts, _coins)
}

// Collect0 is a paid mutator transaction binding the contract method 0x42b1689d.
//
// Solidity: function collect(address[] _coins, address _receiver) returns()
func (_FeeCollector *FeeCollectorTransactor) Collect0(opts *bind.TransactOpts, _coins []common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "collect0", _coins, _receiver)
}

// Collect0 is a paid mutator transaction binding the contract method 0x42b1689d.
//
// Solidity: function collect(address[] _coins, address _receiver) returns()
func (_FeeCollector *FeeCollectorSession) Collect0(_coins []common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Collect0(&_FeeCollector.TransactOpts, _coins, _receiver)
}

// Collect0 is a paid mutator transaction binding the contract method 0x42b1689d.
//
// Solidity: function collect(address[] _coins, address _receiver) returns()
func (_FeeCollector *FeeCollectorTransactorSession) Collect0(_coins []common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Collect0(&_FeeCollector.TransactOpts, _coins, _receiver)
}

// Forward is a paid mutator transaction binding the contract method 0xd11d83ab.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs) payable returns(uint256)
func (_FeeCollector *FeeCollectorTransactor) Forward(opts *bind.TransactOpts, _hook_inputs []Struct0) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "forward", _hook_inputs)
}

// Forward is a paid mutator transaction binding the contract method 0xd11d83ab.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs) payable returns(uint256)
func (_FeeCollector *FeeCollectorSession) Forward(_hook_inputs []Struct0) (*types.Transaction, error) {
	return _FeeCollector.Contract.Forward(&_FeeCollector.TransactOpts, _hook_inputs)
}

// Forward is a paid mutator transaction binding the contract method 0xd11d83ab.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs) payable returns(uint256)
func (_FeeCollector *FeeCollectorTransactorSession) Forward(_hook_inputs []Struct0) (*types.Transaction, error) {
	return _FeeCollector.Contract.Forward(&_FeeCollector.TransactOpts, _hook_inputs)
}

// Forward0 is a paid mutator transaction binding the contract method 0x167f83f4.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs, address _receiver) payable returns(uint256)
func (_FeeCollector *FeeCollectorTransactor) Forward0(opts *bind.TransactOpts, _hook_inputs []Struct0, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "forward0", _hook_inputs, _receiver)
}

// Forward0 is a paid mutator transaction binding the contract method 0x167f83f4.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs, address _receiver) payable returns(uint256)
func (_FeeCollector *FeeCollectorSession) Forward0(_hook_inputs []Struct0, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Forward0(&_FeeCollector.TransactOpts, _hook_inputs, _receiver)
}

// Forward0 is a paid mutator transaction binding the contract method 0x167f83f4.
//
// Solidity: function forward((uint8,uint256,bytes)[] _hook_inputs, address _receiver) payable returns(uint256)
func (_FeeCollector *FeeCollectorTransactorSession) Forward0(_hook_inputs []Struct0, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Forward0(&_FeeCollector.TransactOpts, _hook_inputs, _receiver)
}

// Recover is a paid mutator transaction binding the contract method 0x4e1602df.
//
// Solidity: function recover((address,uint256)[] _recovers, address _receiver) returns()
func (_FeeCollector *FeeCollectorTransactor) Recover(opts *bind.TransactOpts, _recovers []Struct1, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "recover", _recovers, _receiver)
}

// Recover is a paid mutator transaction binding the contract method 0x4e1602df.
//
// Solidity: function recover((address,uint256)[] _recovers, address _receiver) returns()
func (_FeeCollector *FeeCollectorSession) Recover(_recovers []Struct1, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Recover(&_FeeCollector.TransactOpts, _recovers, _receiver)
}

// Recover is a paid mutator transaction binding the contract method 0x4e1602df.
//
// Solidity: function recover((address,uint256)[] _recovers, address _receiver) returns()
func (_FeeCollector *FeeCollectorTransactorSession) Recover(_recovers []Struct1, _receiver common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.Recover(&_FeeCollector.TransactOpts, _recovers, _receiver)
}

// SetBurner is a paid mutator transaction binding the contract method 0xe0774862.
//
// Solidity: function set_burner(address _new_burner) returns()
func (_FeeCollector *FeeCollectorTransactor) SetBurner(opts *bind.TransactOpts, _new_burner common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_burner", _new_burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xe0774862.
//
// Solidity: function set_burner(address _new_burner) returns()
func (_FeeCollector *FeeCollectorSession) SetBurner(_new_burner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetBurner(&_FeeCollector.TransactOpts, _new_burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xe0774862.
//
// Solidity: function set_burner(address _new_burner) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetBurner(_new_burner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetBurner(&_FeeCollector.TransactOpts, _new_burner)
}

// SetEmergencyOwner is a paid mutator transaction binding the contract method 0x7e1b1e4b.
//
// Solidity: function set_emergency_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorTransactor) SetEmergencyOwner(opts *bind.TransactOpts, _new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_emergency_owner", _new_owner)
}

// SetEmergencyOwner is a paid mutator transaction binding the contract method 0x7e1b1e4b.
//
// Solidity: function set_emergency_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorSession) SetEmergencyOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetEmergencyOwner(&_FeeCollector.TransactOpts, _new_owner)
}

// SetEmergencyOwner is a paid mutator transaction binding the contract method 0x7e1b1e4b.
//
// Solidity: function set_emergency_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetEmergencyOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetEmergencyOwner(&_FeeCollector.TransactOpts, _new_owner)
}

// SetHooker is a paid mutator transaction binding the contract method 0xa35eba8d.
//
// Solidity: function set_hooker(address _new_hooker) returns()
func (_FeeCollector *FeeCollectorTransactor) SetHooker(opts *bind.TransactOpts, _new_hooker common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_hooker", _new_hooker)
}

// SetHooker is a paid mutator transaction binding the contract method 0xa35eba8d.
//
// Solidity: function set_hooker(address _new_hooker) returns()
func (_FeeCollector *FeeCollectorSession) SetHooker(_new_hooker common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetHooker(&_FeeCollector.TransactOpts, _new_hooker)
}

// SetHooker is a paid mutator transaction binding the contract method 0xa35eba8d.
//
// Solidity: function set_hooker(address _new_hooker) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetHooker(_new_hooker common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetHooker(&_FeeCollector.TransactOpts, _new_hooker)
}

// SetKilled is a paid mutator transaction binding the contract method 0x3f4822d6.
//
// Solidity: function set_killed((address,uint256)[] _input) returns()
func (_FeeCollector *FeeCollectorTransactor) SetKilled(opts *bind.TransactOpts, _input []Struct1) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_killed", _input)
}

// SetKilled is a paid mutator transaction binding the contract method 0x3f4822d6.
//
// Solidity: function set_killed((address,uint256)[] _input) returns()
func (_FeeCollector *FeeCollectorSession) SetKilled(_input []Struct1) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetKilled(&_FeeCollector.TransactOpts, _input)
}

// SetKilled is a paid mutator transaction binding the contract method 0x3f4822d6.
//
// Solidity: function set_killed((address,uint256)[] _input) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetKilled(_input []Struct1) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetKilled(&_FeeCollector.TransactOpts, _input)
}

// SetMaxFee is a paid mutator transaction binding the contract method 0x97160c5b.
//
// Solidity: function set_max_fee(uint256 _epoch, uint256 _max_fee) returns()
func (_FeeCollector *FeeCollectorTransactor) SetMaxFee(opts *bind.TransactOpts, _epoch *big.Int, _max_fee *big.Int) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_max_fee", _epoch, _max_fee)
}

// SetMaxFee is a paid mutator transaction binding the contract method 0x97160c5b.
//
// Solidity: function set_max_fee(uint256 _epoch, uint256 _max_fee) returns()
func (_FeeCollector *FeeCollectorSession) SetMaxFee(_epoch *big.Int, _max_fee *big.Int) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetMaxFee(&_FeeCollector.TransactOpts, _epoch, _max_fee)
}

// SetMaxFee is a paid mutator transaction binding the contract method 0x97160c5b.
//
// Solidity: function set_max_fee(uint256 _epoch, uint256 _max_fee) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetMaxFee(_epoch *big.Int, _max_fee *big.Int) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetMaxFee(&_FeeCollector.TransactOpts, _epoch, _max_fee)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorTransactor) SetOwner(opts *bind.TransactOpts, _new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_owner", _new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorSession) SetOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetOwner(&_FeeCollector.TransactOpts, _new_owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x7cb97b2b.
//
// Solidity: function set_owner(address _new_owner) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetOwner(_new_owner common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetOwner(&_FeeCollector.TransactOpts, _new_owner)
}

// SetTarget is a paid mutator transaction binding the contract method 0x427c9a95.
//
// Solidity: function set_target(address _new_target) returns()
func (_FeeCollector *FeeCollectorTransactor) SetTarget(opts *bind.TransactOpts, _new_target common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "set_target", _new_target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x427c9a95.
//
// Solidity: function set_target(address _new_target) returns()
func (_FeeCollector *FeeCollectorSession) SetTarget(_new_target common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetTarget(&_FeeCollector.TransactOpts, _new_target)
}

// SetTarget is a paid mutator transaction binding the contract method 0x427c9a95.
//
// Solidity: function set_target(address _new_target) returns()
func (_FeeCollector *FeeCollectorTransactorSession) SetTarget(_new_target common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.SetTarget(&_FeeCollector.TransactOpts, _new_target)
}

// Transfer is a paid mutator transaction binding the contract method 0x0e67c6bc.
//
// Solidity: function transfer((address,address,uint256)[] _transfers) returns()
func (_FeeCollector *FeeCollectorTransactor) Transfer(opts *bind.TransactOpts, _transfers []Struct2) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "transfer", _transfers)
}

// Transfer is a paid mutator transaction binding the contract method 0x0e67c6bc.
//
// Solidity: function transfer((address,address,uint256)[] _transfers) returns()
func (_FeeCollector *FeeCollectorSession) Transfer(_transfers []Struct2) (*types.Transaction, error) {
	return _FeeCollector.Contract.Transfer(&_FeeCollector.TransactOpts, _transfers)
}

// Transfer is a paid mutator transaction binding the contract method 0x0e67c6bc.
//
// Solidity: function transfer((address,address,uint256)[] _transfers) returns()
func (_FeeCollector *FeeCollectorTransactorSession) Transfer(_transfers []Struct2) (*types.Transaction, error) {
	return _FeeCollector.Contract.Transfer(&_FeeCollector.TransactOpts, _transfers)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0x755da811.
//
// Solidity: function withdraw_many(address[] _pools) returns()
func (_FeeCollector *FeeCollectorTransactor) WithdrawMany(opts *bind.TransactOpts, _pools []common.Address) (*types.Transaction, error) {
	return _FeeCollector.contract.Transact(opts, "withdraw_many", _pools)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0x755da811.
//
// Solidity: function withdraw_many(address[] _pools) returns()
func (_FeeCollector *FeeCollectorSession) WithdrawMany(_pools []common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.WithdrawMany(&_FeeCollector.TransactOpts, _pools)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0x755da811.
//
// Solidity: function withdraw_many(address[] _pools) returns()
func (_FeeCollector *FeeCollectorTransactorSession) WithdrawMany(_pools []common.Address) (*types.Transaction, error) {
	return _FeeCollector.Contract.WithdrawMany(&_FeeCollector.TransactOpts, _pools)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeCollector *FeeCollectorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FeeCollector.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeCollector *FeeCollectorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeCollector.Contract.Fallback(&_FeeCollector.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeCollector *FeeCollectorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeCollector.Contract.Fallback(&_FeeCollector.TransactOpts, calldata)
}

// FeeCollectorSetBurnerIterator is returned from FilterSetBurner and is used to iterate over the raw logs and unpacked data for SetBurner events raised by the FeeCollector contract.
type FeeCollectorSetBurnerIterator struct {
	Event *FeeCollectorSetBurner // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetBurnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetBurner)
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
		it.Event = new(FeeCollectorSetBurner)
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
func (it *FeeCollectorSetBurnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetBurnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetBurner represents a SetBurner event raised by the FeeCollector contract.
type FeeCollectorSetBurner struct {
	Burner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetBurner is a free log retrieval operation binding the contract event 0x5d02513563a6890385b0d6684a867cbf8032b19adbd10bca1f78f430db041e37.
//
// Solidity: event SetBurner(address indexed burner)
func (_FeeCollector *FeeCollectorFilterer) FilterSetBurner(opts *bind.FilterOpts, burner []common.Address) (*FeeCollectorSetBurnerIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetBurner", burnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetBurnerIterator{contract: _FeeCollector.contract, event: "SetBurner", logs: logs, sub: sub}, nil
}

// WatchSetBurner is a free log subscription operation binding the contract event 0x5d02513563a6890385b0d6684a867cbf8032b19adbd10bca1f78f430db041e37.
//
// Solidity: event SetBurner(address indexed burner)
func (_FeeCollector *FeeCollectorFilterer) WatchSetBurner(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetBurner, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetBurner", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetBurner)
				if err := _FeeCollector.contract.UnpackLog(event, "SetBurner", log); err != nil {
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

// ParseSetBurner is a log parse operation binding the contract event 0x5d02513563a6890385b0d6684a867cbf8032b19adbd10bca1f78f430db041e37.
//
// Solidity: event SetBurner(address indexed burner)
func (_FeeCollector *FeeCollectorFilterer) ParseSetBurner(log types.Log) (*FeeCollectorSetBurner, error) {
	event := new(FeeCollectorSetBurner)
	if err := _FeeCollector.contract.UnpackLog(event, "SetBurner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetEmergencyOwnerIterator is returned from FilterSetEmergencyOwner and is used to iterate over the raw logs and unpacked data for SetEmergencyOwner events raised by the FeeCollector contract.
type FeeCollectorSetEmergencyOwnerIterator struct {
	Event *FeeCollectorSetEmergencyOwner // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetEmergencyOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetEmergencyOwner)
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
		it.Event = new(FeeCollectorSetEmergencyOwner)
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
func (it *FeeCollectorSetEmergencyOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetEmergencyOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetEmergencyOwner represents a SetEmergencyOwner event raised by the FeeCollector contract.
type FeeCollectorSetEmergencyOwner struct {
	EmergencyOwner common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetEmergencyOwner is a free log retrieval operation binding the contract event 0x8d244bde4451f1c7643d677f668fcf551617d4046657105728f1e0c6d1f41abd.
//
// Solidity: event SetEmergencyOwner(address indexed emergency_owner)
func (_FeeCollector *FeeCollectorFilterer) FilterSetEmergencyOwner(opts *bind.FilterOpts, emergency_owner []common.Address) (*FeeCollectorSetEmergencyOwnerIterator, error) {

	var emergency_ownerRule []interface{}
	for _, emergency_ownerItem := range emergency_owner {
		emergency_ownerRule = append(emergency_ownerRule, emergency_ownerItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetEmergencyOwner", emergency_ownerRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetEmergencyOwnerIterator{contract: _FeeCollector.contract, event: "SetEmergencyOwner", logs: logs, sub: sub}, nil
}

// WatchSetEmergencyOwner is a free log subscription operation binding the contract event 0x8d244bde4451f1c7643d677f668fcf551617d4046657105728f1e0c6d1f41abd.
//
// Solidity: event SetEmergencyOwner(address indexed emergency_owner)
func (_FeeCollector *FeeCollectorFilterer) WatchSetEmergencyOwner(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetEmergencyOwner, emergency_owner []common.Address) (event.Subscription, error) {

	var emergency_ownerRule []interface{}
	for _, emergency_ownerItem := range emergency_owner {
		emergency_ownerRule = append(emergency_ownerRule, emergency_ownerItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetEmergencyOwner", emergency_ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetEmergencyOwner)
				if err := _FeeCollector.contract.UnpackLog(event, "SetEmergencyOwner", log); err != nil {
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

// ParseSetEmergencyOwner is a log parse operation binding the contract event 0x8d244bde4451f1c7643d677f668fcf551617d4046657105728f1e0c6d1f41abd.
//
// Solidity: event SetEmergencyOwner(address indexed emergency_owner)
func (_FeeCollector *FeeCollectorFilterer) ParseSetEmergencyOwner(log types.Log) (*FeeCollectorSetEmergencyOwner, error) {
	event := new(FeeCollectorSetEmergencyOwner)
	if err := _FeeCollector.contract.UnpackLog(event, "SetEmergencyOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetHookerIterator is returned from FilterSetHooker and is used to iterate over the raw logs and unpacked data for SetHooker events raised by the FeeCollector contract.
type FeeCollectorSetHookerIterator struct {
	Event *FeeCollectorSetHooker // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetHookerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetHooker)
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
		it.Event = new(FeeCollectorSetHooker)
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
func (it *FeeCollectorSetHookerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetHookerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetHooker represents a SetHooker event raised by the FeeCollector contract.
type FeeCollectorSetHooker struct {
	Hooker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetHooker is a free log retrieval operation binding the contract event 0x557aa1ccd69dcfb5220a0ee57e91034c3e2ae382692ffe000c16b2865fcca969.
//
// Solidity: event SetHooker(address indexed hooker)
func (_FeeCollector *FeeCollectorFilterer) FilterSetHooker(opts *bind.FilterOpts, hooker []common.Address) (*FeeCollectorSetHookerIterator, error) {

	var hookerRule []interface{}
	for _, hookerItem := range hooker {
		hookerRule = append(hookerRule, hookerItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetHooker", hookerRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetHookerIterator{contract: _FeeCollector.contract, event: "SetHooker", logs: logs, sub: sub}, nil
}

// WatchSetHooker is a free log subscription operation binding the contract event 0x557aa1ccd69dcfb5220a0ee57e91034c3e2ae382692ffe000c16b2865fcca969.
//
// Solidity: event SetHooker(address indexed hooker)
func (_FeeCollector *FeeCollectorFilterer) WatchSetHooker(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetHooker, hooker []common.Address) (event.Subscription, error) {

	var hookerRule []interface{}
	for _, hookerItem := range hooker {
		hookerRule = append(hookerRule, hookerItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetHooker", hookerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetHooker)
				if err := _FeeCollector.contract.UnpackLog(event, "SetHooker", log); err != nil {
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

// ParseSetHooker is a log parse operation binding the contract event 0x557aa1ccd69dcfb5220a0ee57e91034c3e2ae382692ffe000c16b2865fcca969.
//
// Solidity: event SetHooker(address indexed hooker)
func (_FeeCollector *FeeCollectorFilterer) ParseSetHooker(log types.Log) (*FeeCollectorSetHooker, error) {
	event := new(FeeCollectorSetHooker)
	if err := _FeeCollector.contract.UnpackLog(event, "SetHooker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetKilledIterator is returned from FilterSetKilled and is used to iterate over the raw logs and unpacked data for SetKilled events raised by the FeeCollector contract.
type FeeCollectorSetKilledIterator struct {
	Event *FeeCollectorSetKilled // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetKilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetKilled)
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
		it.Event = new(FeeCollectorSetKilled)
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
func (it *FeeCollectorSetKilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetKilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetKilled represents a SetKilled event raised by the FeeCollector contract.
type FeeCollectorSetKilled struct {
	Coin      common.Address
	EpochMask *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetKilled is a free log retrieval operation binding the contract event 0xbe9ca79c5858b597d131478211339014c63cf8066fc364a3a616e5635de82e6a.
//
// Solidity: event SetKilled(address indexed coin, uint256 epoch_mask)
func (_FeeCollector *FeeCollectorFilterer) FilterSetKilled(opts *bind.FilterOpts, coin []common.Address) (*FeeCollectorSetKilledIterator, error) {

	var coinRule []interface{}
	for _, coinItem := range coin {
		coinRule = append(coinRule, coinItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetKilled", coinRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetKilledIterator{contract: _FeeCollector.contract, event: "SetKilled", logs: logs, sub: sub}, nil
}

// WatchSetKilled is a free log subscription operation binding the contract event 0xbe9ca79c5858b597d131478211339014c63cf8066fc364a3a616e5635de82e6a.
//
// Solidity: event SetKilled(address indexed coin, uint256 epoch_mask)
func (_FeeCollector *FeeCollectorFilterer) WatchSetKilled(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetKilled, coin []common.Address) (event.Subscription, error) {

	var coinRule []interface{}
	for _, coinItem := range coin {
		coinRule = append(coinRule, coinItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetKilled", coinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetKilled)
				if err := _FeeCollector.contract.UnpackLog(event, "SetKilled", log); err != nil {
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

// ParseSetKilled is a log parse operation binding the contract event 0xbe9ca79c5858b597d131478211339014c63cf8066fc364a3a616e5635de82e6a.
//
// Solidity: event SetKilled(address indexed coin, uint256 epoch_mask)
func (_FeeCollector *FeeCollectorFilterer) ParseSetKilled(log types.Log) (*FeeCollectorSetKilled, error) {
	event := new(FeeCollectorSetKilled)
	if err := _FeeCollector.contract.UnpackLog(event, "SetKilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetMaxFeeIterator is returned from FilterSetMaxFee and is used to iterate over the raw logs and unpacked data for SetMaxFee events raised by the FeeCollector contract.
type FeeCollectorSetMaxFeeIterator struct {
	Event *FeeCollectorSetMaxFee // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetMaxFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetMaxFee)
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
		it.Event = new(FeeCollectorSetMaxFee)
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
func (it *FeeCollectorSetMaxFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetMaxFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetMaxFee represents a SetMaxFee event raised by the FeeCollector contract.
type FeeCollectorSetMaxFee struct {
	Epoch  *big.Int
	MaxFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxFee is a free log retrieval operation binding the contract event 0x62c711266aa88195f0a4d039bff4072f30cd1171b5473007a6ed72225cf8f329.
//
// Solidity: event SetMaxFee(uint256 indexed epoch, uint256 max_fee)
func (_FeeCollector *FeeCollectorFilterer) FilterSetMaxFee(opts *bind.FilterOpts, epoch []*big.Int) (*FeeCollectorSetMaxFeeIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetMaxFee", epochRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetMaxFeeIterator{contract: _FeeCollector.contract, event: "SetMaxFee", logs: logs, sub: sub}, nil
}

// WatchSetMaxFee is a free log subscription operation binding the contract event 0x62c711266aa88195f0a4d039bff4072f30cd1171b5473007a6ed72225cf8f329.
//
// Solidity: event SetMaxFee(uint256 indexed epoch, uint256 max_fee)
func (_FeeCollector *FeeCollectorFilterer) WatchSetMaxFee(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetMaxFee, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetMaxFee", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetMaxFee)
				if err := _FeeCollector.contract.UnpackLog(event, "SetMaxFee", log); err != nil {
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

// ParseSetMaxFee is a log parse operation binding the contract event 0x62c711266aa88195f0a4d039bff4072f30cd1171b5473007a6ed72225cf8f329.
//
// Solidity: event SetMaxFee(uint256 indexed epoch, uint256 max_fee)
func (_FeeCollector *FeeCollectorFilterer) ParseSetMaxFee(log types.Log) (*FeeCollectorSetMaxFee, error) {
	event := new(FeeCollectorSetMaxFee)
	if err := _FeeCollector.contract.UnpackLog(event, "SetMaxFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetOwnerIterator is returned from FilterSetOwner and is used to iterate over the raw logs and unpacked data for SetOwner events raised by the FeeCollector contract.
type FeeCollectorSetOwnerIterator struct {
	Event *FeeCollectorSetOwner // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetOwner)
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
		it.Event = new(FeeCollectorSetOwner)
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
func (it *FeeCollectorSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetOwner represents a SetOwner event raised by the FeeCollector contract.
type FeeCollectorSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetOwner is a free log retrieval operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed owner)
func (_FeeCollector *FeeCollectorFilterer) FilterSetOwner(opts *bind.FilterOpts, owner []common.Address) (*FeeCollectorSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetOwnerIterator{contract: _FeeCollector.contract, event: "SetOwner", logs: logs, sub: sub}, nil
}

// WatchSetOwner is a free log subscription operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed owner)
func (_FeeCollector *FeeCollectorFilterer) WatchSetOwner(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetOwner)
				if err := _FeeCollector.contract.UnpackLog(event, "SetOwner", log); err != nil {
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

// ParseSetOwner is a log parse operation binding the contract event 0x167d3e9c1016ab80e58802ca9da10ce5c6a0f4debc46a2e7a2cd9e56899a4fb5.
//
// Solidity: event SetOwner(address indexed owner)
func (_FeeCollector *FeeCollectorFilterer) ParseSetOwner(log types.Log) (*FeeCollectorSetOwner, error) {
	event := new(FeeCollectorSetOwner)
	if err := _FeeCollector.contract.UnpackLog(event, "SetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCollectorSetTargetIterator is returned from FilterSetTarget and is used to iterate over the raw logs and unpacked data for SetTarget events raised by the FeeCollector contract.
type FeeCollectorSetTargetIterator struct {
	Event *FeeCollectorSetTarget // Event containing the contract specifics and raw log

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
func (it *FeeCollectorSetTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectorSetTarget)
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
		it.Event = new(FeeCollectorSetTarget)
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
func (it *FeeCollectorSetTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectorSetTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectorSetTarget represents a SetTarget event raised by the FeeCollector contract.
type FeeCollectorSetTarget struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTarget is a free log retrieval operation binding the contract event 0x529334a1c6836c58fdfaae38f1bc5b62d6a69e91fc4181371b8cd8db71831ea7.
//
// Solidity: event SetTarget(address indexed target)
func (_FeeCollector *FeeCollectorFilterer) FilterSetTarget(opts *bind.FilterOpts, target []common.Address) (*FeeCollectorSetTargetIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeCollector.contract.FilterLogs(opts, "SetTarget", targetRule)
	if err != nil {
		return nil, err
	}
	return &FeeCollectorSetTargetIterator{contract: _FeeCollector.contract, event: "SetTarget", logs: logs, sub: sub}, nil
}

// WatchSetTarget is a free log subscription operation binding the contract event 0x529334a1c6836c58fdfaae38f1bc5b62d6a69e91fc4181371b8cd8db71831ea7.
//
// Solidity: event SetTarget(address indexed target)
func (_FeeCollector *FeeCollectorFilterer) WatchSetTarget(opts *bind.WatchOpts, sink chan<- *FeeCollectorSetTarget, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _FeeCollector.contract.WatchLogs(opts, "SetTarget", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectorSetTarget)
				if err := _FeeCollector.contract.UnpackLog(event, "SetTarget", log); err != nil {
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

// ParseSetTarget is a log parse operation binding the contract event 0x529334a1c6836c58fdfaae38f1bc5b62d6a69e91fc4181371b8cd8db71831ea7.
//
// Solidity: event SetTarget(address indexed target)
func (_FeeCollector *FeeCollectorFilterer) ParseSetTarget(log types.Log) (*FeeCollectorSetTarget, error) {
	event := new(FeeCollectorSetTarget)
	if err := _FeeCollector.contract.UnpackLog(event, "SetTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
