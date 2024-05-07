// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vesting

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

// VestingMetaData contains all meta data concerning the Vesting contract.
var VestingMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Fund\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Claim\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"claimed\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ToggleDisable\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"disabled\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"uint256\",\"name\":\"_start_time\"},{\"type\":\"uint256\",\"name\":\"_end_time\"},{\"type\":\"bool\",\"name\":\"_can_disable\"},{\"type\":\"address[4]\",\"name\":\"_fund_admins\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"add_tokens\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39108},{\"name\":\"fund\",\"outputs\":[],\"inputs\":[{\"type\":\"address[100]\",\"name\":\"_recipients\"},{\"type\":\"uint256[100]\",\"name\":\"_amounts\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":3962646},{\"name\":\"toggle_disable\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40280},{\"name\":\"disable_can_disable\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21295},{\"name\":\"disable_fund_admins\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21325},{\"name\":\"vestedSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4468},{\"name\":\"lockedSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5465},{\"name\":\"vestedOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5163},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6275},{\"name\":\"lockedOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6305},{\"name\":\"claim\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"claim\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38032},{\"name\":\"apply_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38932},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1601},{\"name\":\"start_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1631},{\"name\":\"end_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1661},{\"name\":\"initial_locked\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1845},{\"name\":\"total_claimed\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1875},{\"name\":\"initial_locked_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1751},{\"name\":\"unallocated_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1781},{\"name\":\"can_disable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"disabled_at\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1995},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"fund_admins_enabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"fund_admins\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2115}]",
}

// VestingABI is the input ABI used to generate the binding from.
// Deprecated: Use VestingMetaData.ABI instead.
var VestingABI = VestingMetaData.ABI

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vesting *VestingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vesting *VestingSession) Admin() (common.Address, error) {
	return _Vesting.Contract.Admin(&_Vesting.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Vesting *VestingCallerSession) Admin() (common.Address, error) {
	return _Vesting.Contract.Admin(&_Vesting.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCaller) BalanceOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "balanceOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_Vesting *VestingSession) BalanceOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _recipient)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCallerSession) BalanceOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _recipient)
}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_Vesting *VestingCaller) CanDisable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "can_disable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_Vesting *VestingSession) CanDisable() (bool, error) {
	return _Vesting.Contract.CanDisable(&_Vesting.CallOpts)
}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_Vesting *VestingCallerSession) CanDisable() (bool, error) {
	return _Vesting.Contract.CanDisable(&_Vesting.CallOpts)
}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_Vesting *VestingCaller) DisabledAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "disabled_at", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_Vesting *VestingSession) DisabledAt(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.DisabledAt(&_Vesting.CallOpts, arg0)
}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_Vesting *VestingCallerSession) DisabledAt(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.DisabledAt(&_Vesting.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_Vesting *VestingCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "end_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_Vesting *VestingSession) EndTime() (*big.Int, error) {
	return _Vesting.Contract.EndTime(&_Vesting.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_Vesting *VestingCallerSession) EndTime() (*big.Int, error) {
	return _Vesting.Contract.EndTime(&_Vesting.CallOpts)
}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_Vesting *VestingCaller) FundAdmins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "fund_admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_Vesting *VestingSession) FundAdmins(arg0 common.Address) (bool, error) {
	return _Vesting.Contract.FundAdmins(&_Vesting.CallOpts, arg0)
}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_Vesting *VestingCallerSession) FundAdmins(arg0 common.Address) (bool, error) {
	return _Vesting.Contract.FundAdmins(&_Vesting.CallOpts, arg0)
}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_Vesting *VestingCaller) FundAdminsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "fund_admins_enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_Vesting *VestingSession) FundAdminsEnabled() (bool, error) {
	return _Vesting.Contract.FundAdminsEnabled(&_Vesting.CallOpts)
}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_Vesting *VestingCallerSession) FundAdminsEnabled() (bool, error) {
	return _Vesting.Contract.FundAdminsEnabled(&_Vesting.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Vesting *VestingCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Vesting *VestingSession) FutureAdmin() (common.Address, error) {
	return _Vesting.Contract.FutureAdmin(&_Vesting.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Vesting *VestingCallerSession) FutureAdmin() (common.Address, error) {
	return _Vesting.Contract.FutureAdmin(&_Vesting.CallOpts)
}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_Vesting *VestingCaller) InitialLocked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "initial_locked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_Vesting *VestingSession) InitialLocked(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.InitialLocked(&_Vesting.CallOpts, arg0)
}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_Vesting *VestingCallerSession) InitialLocked(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.InitialLocked(&_Vesting.CallOpts, arg0)
}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_Vesting *VestingCaller) InitialLockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "initial_locked_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_Vesting *VestingSession) InitialLockedSupply() (*big.Int, error) {
	return _Vesting.Contract.InitialLockedSupply(&_Vesting.CallOpts)
}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_Vesting *VestingCallerSession) InitialLockedSupply() (*big.Int, error) {
	return _Vesting.Contract.InitialLockedSupply(&_Vesting.CallOpts)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCaller) LockedOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "lockedOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingSession) LockedOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.LockedOf(&_Vesting.CallOpts, _recipient)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCallerSession) LockedOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.LockedOf(&_Vesting.CallOpts, _recipient)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_Vesting *VestingCaller) LockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "lockedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_Vesting *VestingSession) LockedSupply() (*big.Int, error) {
	return _Vesting.Contract.LockedSupply(&_Vesting.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_Vesting *VestingCallerSession) LockedSupply() (*big.Int, error) {
	return _Vesting.Contract.LockedSupply(&_Vesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Vesting *VestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "start_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Vesting *VestingSession) StartTime() (*big.Int, error) {
	return _Vesting.Contract.StartTime(&_Vesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_Vesting *VestingCallerSession) StartTime() (*big.Int, error) {
	return _Vesting.Contract.StartTime(&_Vesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vesting *VestingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vesting *VestingSession) Token() (common.Address, error) {
	return _Vesting.Contract.Token(&_Vesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vesting *VestingCallerSession) Token() (common.Address, error) {
	return _Vesting.Contract.Token(&_Vesting.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_Vesting *VestingCaller) TotalClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "total_claimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_Vesting *VestingSession) TotalClaimed(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.TotalClaimed(&_Vesting.CallOpts, arg0)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_Vesting *VestingCallerSession) TotalClaimed(arg0 common.Address) (*big.Int, error) {
	return _Vesting.Contract.TotalClaimed(&_Vesting.CallOpts, arg0)
}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_Vesting *VestingCaller) UnallocatedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "unallocated_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_Vesting *VestingSession) UnallocatedSupply() (*big.Int, error) {
	return _Vesting.Contract.UnallocatedSupply(&_Vesting.CallOpts)
}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_Vesting *VestingCallerSession) UnallocatedSupply() (*big.Int, error) {
	return _Vesting.Contract.UnallocatedSupply(&_Vesting.CallOpts)
}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCaller) VestedOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "vestedOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingSession) VestedOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.VestedOf(&_Vesting.CallOpts, _recipient)
}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_Vesting *VestingCallerSession) VestedOf(_recipient common.Address) (*big.Int, error) {
	return _Vesting.Contract.VestedOf(&_Vesting.CallOpts, _recipient)
}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_Vesting *VestingCaller) VestedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "vestedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_Vesting *VestingSession) VestedSupply() (*big.Int, error) {
	return _Vesting.Contract.VestedSupply(&_Vesting.CallOpts)
}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_Vesting *VestingCallerSession) VestedSupply() (*big.Int, error) {
	return _Vesting.Contract.VestedSupply(&_Vesting.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_Vesting *VestingTransactor) AddTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "add_tokens", _amount)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_Vesting *VestingSession) AddTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AddTokens(&_Vesting.TransactOpts, _amount)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) AddTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AddTokens(&_Vesting.TransactOpts, _amount)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_Vesting *VestingTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_Vesting *VestingSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Vesting.Contract.ApplyTransferOwnership(&_Vesting.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_Vesting *VestingTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Vesting.Contract.ApplyTransferOwnership(&_Vesting.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vesting *VestingTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vesting *VestingSession) Claim() (*types.Transaction, error) {
	return _Vesting.Contract.Claim(&_Vesting.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vesting *VestingTransactorSession) Claim() (*types.Transaction, error) {
	return _Vesting.Contract.Claim(&_Vesting.TransactOpts)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_Vesting *VestingTransactor) Claim0(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "claim0", addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_Vesting *VestingSession) Claim0(addr common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Claim0(&_Vesting.TransactOpts, addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_Vesting *VestingTransactorSession) Claim0(addr common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Claim0(&_Vesting.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_Vesting *VestingTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_Vesting *VestingSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.CommitTransferOwnership(&_Vesting.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_Vesting *VestingTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.CommitTransferOwnership(&_Vesting.TransactOpts, addr)
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_Vesting *VestingTransactor) DisableCanDisable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "disable_can_disable")
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_Vesting *VestingSession) DisableCanDisable() (*types.Transaction, error) {
	return _Vesting.Contract.DisableCanDisable(&_Vesting.TransactOpts)
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_Vesting *VestingTransactorSession) DisableCanDisable() (*types.Transaction, error) {
	return _Vesting.Contract.DisableCanDisable(&_Vesting.TransactOpts)
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_Vesting *VestingTransactor) DisableFundAdmins(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "disable_fund_admins")
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_Vesting *VestingSession) DisableFundAdmins() (*types.Transaction, error) {
	return _Vesting.Contract.DisableFundAdmins(&_Vesting.TransactOpts)
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_Vesting *VestingTransactorSession) DisableFundAdmins() (*types.Transaction, error) {
	return _Vesting.Contract.DisableFundAdmins(&_Vesting.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_Vesting *VestingTransactor) Fund(opts *bind.TransactOpts, _recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "fund", _recipients, _amounts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_Vesting *VestingSession) Fund(_recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Fund(&_Vesting.TransactOpts, _recipients, _amounts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_Vesting *VestingTransactorSession) Fund(_recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Fund(&_Vesting.TransactOpts, _recipients, _amounts)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_Vesting *VestingTransactor) ToggleDisable(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "toggle_disable", _recipient)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_Vesting *VestingSession) ToggleDisable(_recipient common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.ToggleDisable(&_Vesting.TransactOpts, _recipient)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_Vesting *VestingTransactorSession) ToggleDisable(_recipient common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.ToggleDisable(&_Vesting.TransactOpts, _recipient)
}

// VestingApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the Vesting contract.
type VestingApplyOwnershipIterator struct {
	Event *VestingApplyOwnership // Event containing the contract specifics and raw log

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
func (it *VestingApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingApplyOwnership)
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
		it.Event = new(VestingApplyOwnership)
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
func (it *VestingApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingApplyOwnership represents a ApplyOwnership event raised by the Vesting contract.
type VestingApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Vesting *VestingFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*VestingApplyOwnershipIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &VestingApplyOwnershipIterator{contract: _Vesting.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Vesting *VestingFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *VestingApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingApplyOwnership)
				if err := _Vesting.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Vesting *VestingFilterer) ParseApplyOwnership(log types.Log) (*VestingApplyOwnership, error) {
	event := new(VestingApplyOwnership)
	if err := _Vesting.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vesting contract.
type VestingClaimIterator struct {
	Event *VestingClaim // Event containing the contract specifics and raw log

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
func (it *VestingClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingClaim)
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
		it.Event = new(VestingClaim)
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
func (it *VestingClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingClaim represents a Claim event raised by the Vesting contract.
type VestingClaim struct {
	Recipient common.Address
	Claimed   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_Vesting *VestingFilterer) FilterClaim(opts *bind.FilterOpts, recipient []common.Address) (*VestingClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &VestingClaimIterator{contract: _Vesting.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_Vesting *VestingFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VestingClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingClaim)
				if err := _Vesting.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_Vesting *VestingFilterer) ParseClaim(log types.Log) (*VestingClaim, error) {
	event := new(VestingClaim)
	if err := _Vesting.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the Vesting contract.
type VestingCommitOwnershipIterator struct {
	Event *VestingCommitOwnership // Event containing the contract specifics and raw log

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
func (it *VestingCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingCommitOwnership)
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
		it.Event = new(VestingCommitOwnership)
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
func (it *VestingCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingCommitOwnership represents a CommitOwnership event raised by the Vesting contract.
type VestingCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Vesting *VestingFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*VestingCommitOwnershipIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &VestingCommitOwnershipIterator{contract: _Vesting.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Vesting *VestingFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *VestingCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingCommitOwnership)
				if err := _Vesting.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Vesting *VestingFilterer) ParseCommitOwnership(log types.Log) (*VestingCommitOwnership, error) {
	event := new(VestingCommitOwnership)
	if err := _Vesting.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingFundIterator is returned from FilterFund and is used to iterate over the raw logs and unpacked data for Fund events raised by the Vesting contract.
type VestingFundIterator struct {
	Event *VestingFund // Event containing the contract specifics and raw log

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
func (it *VestingFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingFund)
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
		it.Event = new(VestingFund)
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
func (it *VestingFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingFund represents a Fund event raised by the Vesting contract.
type VestingFund struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFund is a free log retrieval operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_Vesting *VestingFilterer) FilterFund(opts *bind.FilterOpts, recipient []common.Address) (*VestingFundIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "Fund", recipientRule)
	if err != nil {
		return nil, err
	}
	return &VestingFundIterator{contract: _Vesting.contract, event: "Fund", logs: logs, sub: sub}, nil
}

// WatchFund is a free log subscription operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_Vesting *VestingFilterer) WatchFund(opts *bind.WatchOpts, sink chan<- *VestingFund, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "Fund", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingFund)
				if err := _Vesting.contract.UnpackLog(event, "Fund", log); err != nil {
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

// ParseFund is a log parse operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_Vesting *VestingFilterer) ParseFund(log types.Log) (*VestingFund, error) {
	event := new(VestingFund)
	if err := _Vesting.contract.UnpackLog(event, "Fund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingToggleDisableIterator is returned from FilterToggleDisable and is used to iterate over the raw logs and unpacked data for ToggleDisable events raised by the Vesting contract.
type VestingToggleDisableIterator struct {
	Event *VestingToggleDisable // Event containing the contract specifics and raw log

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
func (it *VestingToggleDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingToggleDisable)
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
		it.Event = new(VestingToggleDisable)
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
func (it *VestingToggleDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingToggleDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingToggleDisable represents a ToggleDisable event raised by the Vesting contract.
type VestingToggleDisable struct {
	Recipient common.Address
	Disabled  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterToggleDisable is a free log retrieval operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_Vesting *VestingFilterer) FilterToggleDisable(opts *bind.FilterOpts) (*VestingToggleDisableIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "ToggleDisable")
	if err != nil {
		return nil, err
	}
	return &VestingToggleDisableIterator{contract: _Vesting.contract, event: "ToggleDisable", logs: logs, sub: sub}, nil
}

// WatchToggleDisable is a free log subscription operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_Vesting *VestingFilterer) WatchToggleDisable(opts *bind.WatchOpts, sink chan<- *VestingToggleDisable) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "ToggleDisable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingToggleDisable)
				if err := _Vesting.contract.UnpackLog(event, "ToggleDisable", log); err != nil {
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

// ParseToggleDisable is a log parse operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_Vesting *VestingFilterer) ParseToggleDisable(log types.Log) (*VestingToggleDisable, error) {
	event := new(VestingToggleDisable)
	if err := _Vesting.contract.UnpackLog(event, "ToggleDisable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
