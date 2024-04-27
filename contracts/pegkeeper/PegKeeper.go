// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pegKeeper

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

// PegKeeperMetaData contains all meta data concerning the PegKeeper contract.
var PegKeeperMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Provide\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Profit\",\"inputs\":[{\"name\":\"lp_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetNewCallerShare\",\"inputs\":[{\"name\":\"caller_share\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_caller_share\",\"type\":\"uint256\"},{\"name\":\"_factory\",\"type\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"pegged\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"estimate_caller_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_new_caller_share\",\"inputs\":[{\"name\":\"_new_caller_share\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_admin\",\"inputs\":[{\"name\":\"_new_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_admin\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_receiver\",\"inputs\":[{\"name\":\"_new_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_receiver\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_options\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_change\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"caller_share\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"new_admin_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"new_receiver_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// PegKeeperABI is the input ABI used to generate the binding from.
// Deprecated: Use PegKeeperMetaData.ABI instead.
var PegKeeperABI = PegKeeperMetaData.ABI

// PegKeeper is an auto generated Go binding around an Ethereum contract.
type PegKeeper struct {
	PegKeeperCaller     // Read-only binding to the contract
	PegKeeperTransactor // Write-only binding to the contract
	PegKeeperFilterer   // Log filterer for contract events
}

// PegKeeperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PegKeeperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PegKeeperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PegKeeperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PegKeeperSession struct {
	Contract     *PegKeeper        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PegKeeperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PegKeeperCallerSession struct {
	Contract *PegKeeperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PegKeeperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PegKeeperTransactorSession struct {
	Contract     *PegKeeperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PegKeeperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PegKeeperRaw struct {
	Contract *PegKeeper // Generic contract binding to access the raw methods on
}

// PegKeeperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PegKeeperCallerRaw struct {
	Contract *PegKeeperCaller // Generic read-only contract binding to access the raw methods on
}

// PegKeeperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PegKeeperTransactorRaw struct {
	Contract *PegKeeperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPegKeeper creates a new instance of PegKeeper, bound to a specific deployed contract.
func NewPegKeeper(address common.Address, backend bind.ContractBackend) (*PegKeeper, error) {
	contract, err := bindPegKeeper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PegKeeper{PegKeeperCaller: PegKeeperCaller{contract: contract}, PegKeeperTransactor: PegKeeperTransactor{contract: contract}, PegKeeperFilterer: PegKeeperFilterer{contract: contract}}, nil
}

// NewPegKeeperCaller creates a new read-only instance of PegKeeper, bound to a specific deployed contract.
func NewPegKeeperCaller(address common.Address, caller bind.ContractCaller) (*PegKeeperCaller, error) {
	contract, err := bindPegKeeper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PegKeeperCaller{contract: contract}, nil
}

// NewPegKeeperTransactor creates a new write-only instance of PegKeeper, bound to a specific deployed contract.
func NewPegKeeperTransactor(address common.Address, transactor bind.ContractTransactor) (*PegKeeperTransactor, error) {
	contract, err := bindPegKeeper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PegKeeperTransactor{contract: contract}, nil
}

// NewPegKeeperFilterer creates a new log filterer instance of PegKeeper, bound to a specific deployed contract.
func NewPegKeeperFilterer(address common.Address, filterer bind.ContractFilterer) (*PegKeeperFilterer, error) {
	contract, err := bindPegKeeper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PegKeeperFilterer{contract: contract}, nil
}

// bindPegKeeper binds a generic wrapper to an already deployed contract.
func bindPegKeeper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PegKeeperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PegKeeper *PegKeeperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PegKeeper.Contract.PegKeeperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PegKeeper *PegKeeperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.Contract.PegKeeperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PegKeeper *PegKeeperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PegKeeper.Contract.PegKeeperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PegKeeper *PegKeeperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PegKeeper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PegKeeper *PegKeeperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PegKeeper *PegKeeperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PegKeeper.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PegKeeper *PegKeeperCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PegKeeper *PegKeeperSession) Admin() (common.Address, error) {
	return _PegKeeper.Contract.Admin(&_PegKeeper.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PegKeeper *PegKeeperCallerSession) Admin() (common.Address, error) {
	return _PegKeeper.Contract.Admin(&_PegKeeper.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() pure returns(address)
func (_PegKeeper *PegKeeperCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() pure returns(address)
func (_PegKeeper *PegKeeperSession) Aggregator() (common.Address, error) {
	return _PegKeeper.Contract.Aggregator(&_PegKeeper.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() pure returns(address)
func (_PegKeeper *PegKeeperCallerSession) Aggregator() (common.Address, error) {
	return _PegKeeper.Contract.Aggregator(&_PegKeeper.CallOpts)
}

// CalcProfit is a free data retrieval call binding the contract method 0x322f15d9.
//
// Solidity: function calc_profit() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) CalcProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "calc_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcProfit is a free data retrieval call binding the contract method 0x322f15d9.
//
// Solidity: function calc_profit() view returns(uint256)
func (_PegKeeper *PegKeeperSession) CalcProfit() (*big.Int, error) {
	return _PegKeeper.Contract.CalcProfit(&_PegKeeper.CallOpts)
}

// CalcProfit is a free data retrieval call binding the contract method 0x322f15d9.
//
// Solidity: function calc_profit() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) CalcProfit() (*big.Int, error) {
	return _PegKeeper.Contract.CalcProfit(&_PegKeeper.CallOpts)
}

// CallerShare is a free data retrieval call binding the contract method 0xdf81cd3f.
//
// Solidity: function caller_share() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) CallerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "caller_share")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallerShare is a free data retrieval call binding the contract method 0xdf81cd3f.
//
// Solidity: function caller_share() view returns(uint256)
func (_PegKeeper *PegKeeperSession) CallerShare() (*big.Int, error) {
	return _PegKeeper.Contract.CallerShare(&_PegKeeper.CallOpts)
}

// CallerShare is a free data retrieval call binding the contract method 0xdf81cd3f.
//
// Solidity: function caller_share() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) CallerShare() (*big.Int, error) {
	return _PegKeeper.Contract.CallerShare(&_PegKeeper.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_PegKeeper *PegKeeperSession) Debt() (*big.Int, error) {
	return _PegKeeper.Contract.Debt(&_PegKeeper.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) Debt() (*big.Int, error) {
	return _PegKeeper.Contract.Debt(&_PegKeeper.CallOpts)
}

// EstimateCallerProfit is a free data retrieval call binding the contract method 0xd0d18491.
//
// Solidity: function estimate_caller_profit() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) EstimateCallerProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "estimate_caller_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCallerProfit is a free data retrieval call binding the contract method 0xd0d18491.
//
// Solidity: function estimate_caller_profit() view returns(uint256)
func (_PegKeeper *PegKeeperSession) EstimateCallerProfit() (*big.Int, error) {
	return _PegKeeper.Contract.EstimateCallerProfit(&_PegKeeper.CallOpts)
}

// EstimateCallerProfit is a free data retrieval call binding the contract method 0xd0d18491.
//
// Solidity: function estimate_caller_profit() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) EstimateCallerProfit() (*big.Int, error) {
	return _PegKeeper.Contract.EstimateCallerProfit(&_PegKeeper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_PegKeeper *PegKeeperCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_PegKeeper *PegKeeperSession) Factory() (common.Address, error) {
	return _PegKeeper.Contract.Factory(&_PegKeeper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_PegKeeper *PegKeeperCallerSession) Factory() (common.Address, error) {
	return _PegKeeper.Contract.Factory(&_PegKeeper.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_PegKeeper *PegKeeperCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_PegKeeper *PegKeeperSession) FutureAdmin() (common.Address, error) {
	return _PegKeeper.Contract.FutureAdmin(&_PegKeeper.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_PegKeeper *PegKeeperCallerSession) FutureAdmin() (common.Address, error) {
	return _PegKeeper.Contract.FutureAdmin(&_PegKeeper.CallOpts)
}

// FutureReceiver is a free data retrieval call binding the contract method 0x3bea9ddd.
//
// Solidity: function future_receiver() view returns(address)
func (_PegKeeper *PegKeeperCaller) FutureReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "future_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureReceiver is a free data retrieval call binding the contract method 0x3bea9ddd.
//
// Solidity: function future_receiver() view returns(address)
func (_PegKeeper *PegKeeperSession) FutureReceiver() (common.Address, error) {
	return _PegKeeper.Contract.FutureReceiver(&_PegKeeper.CallOpts)
}

// FutureReceiver is a free data retrieval call binding the contract method 0x3bea9ddd.
//
// Solidity: function future_receiver() view returns(address)
func (_PegKeeper *PegKeeperCallerSession) FutureReceiver() (common.Address, error) {
	return _PegKeeper.Contract.FutureReceiver(&_PegKeeper.CallOpts)
}

// LastChange is a free data retrieval call binding the contract method 0x394b146d.
//
// Solidity: function last_change() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) LastChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "last_change")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastChange is a free data retrieval call binding the contract method 0x394b146d.
//
// Solidity: function last_change() view returns(uint256)
func (_PegKeeper *PegKeeperSession) LastChange() (*big.Int, error) {
	return _PegKeeper.Contract.LastChange(&_PegKeeper.CallOpts)
}

// LastChange is a free data retrieval call binding the contract method 0x394b146d.
//
// Solidity: function last_change() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) LastChange() (*big.Int, error) {
	return _PegKeeper.Contract.LastChange(&_PegKeeper.CallOpts)
}

// NewAdminDeadline is a free data retrieval call binding the contract method 0x553b549c.
//
// Solidity: function new_admin_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) NewAdminDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "new_admin_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewAdminDeadline is a free data retrieval call binding the contract method 0x553b549c.
//
// Solidity: function new_admin_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperSession) NewAdminDeadline() (*big.Int, error) {
	return _PegKeeper.Contract.NewAdminDeadline(&_PegKeeper.CallOpts)
}

// NewAdminDeadline is a free data retrieval call binding the contract method 0x553b549c.
//
// Solidity: function new_admin_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) NewAdminDeadline() (*big.Int, error) {
	return _PegKeeper.Contract.NewAdminDeadline(&_PegKeeper.CallOpts)
}

// NewReceiverDeadline is a free data retrieval call binding the contract method 0x62396066.
//
// Solidity: function new_receiver_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperCaller) NewReceiverDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "new_receiver_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewReceiverDeadline is a free data retrieval call binding the contract method 0x62396066.
//
// Solidity: function new_receiver_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperSession) NewReceiverDeadline() (*big.Int, error) {
	return _PegKeeper.Contract.NewReceiverDeadline(&_PegKeeper.CallOpts)
}

// NewReceiverDeadline is a free data retrieval call binding the contract method 0x62396066.
//
// Solidity: function new_receiver_deadline() view returns(uint256)
func (_PegKeeper *PegKeeperCallerSession) NewReceiverDeadline() (*big.Int, error) {
	return _PegKeeper.Contract.NewReceiverDeadline(&_PegKeeper.CallOpts)
}

// Pegged is a free data retrieval call binding the contract method 0x09078613.
//
// Solidity: function pegged() pure returns(address)
func (_PegKeeper *PegKeeperCaller) Pegged(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "pegged")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pegged is a free data retrieval call binding the contract method 0x09078613.
//
// Solidity: function pegged() pure returns(address)
func (_PegKeeper *PegKeeperSession) Pegged() (common.Address, error) {
	return _PegKeeper.Contract.Pegged(&_PegKeeper.CallOpts)
}

// Pegged is a free data retrieval call binding the contract method 0x09078613.
//
// Solidity: function pegged() pure returns(address)
func (_PegKeeper *PegKeeperCallerSession) Pegged() (common.Address, error) {
	return _PegKeeper.Contract.Pegged(&_PegKeeper.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() pure returns(address)
func (_PegKeeper *PegKeeperCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() pure returns(address)
func (_PegKeeper *PegKeeperSession) Pool() (common.Address, error) {
	return _PegKeeper.Contract.Pool(&_PegKeeper.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() pure returns(address)
func (_PegKeeper *PegKeeperCallerSession) Pool() (common.Address, error) {
	return _PegKeeper.Contract.Pool(&_PegKeeper.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_PegKeeper *PegKeeperCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeper.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_PegKeeper *PegKeeperSession) Receiver() (common.Address, error) {
	return _PegKeeper.Contract.Receiver(&_PegKeeper.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_PegKeeper *PegKeeperCallerSession) Receiver() (common.Address, error) {
	return _PegKeeper.Contract.Receiver(&_PegKeeper.CallOpts)
}

// ApplyNewAdmin is a paid mutator transaction binding the contract method 0xc336fbd5.
//
// Solidity: function apply_new_admin() returns()
func (_PegKeeper *PegKeeperTransactor) ApplyNewAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "apply_new_admin")
}

// ApplyNewAdmin is a paid mutator transaction binding the contract method 0xc336fbd5.
//
// Solidity: function apply_new_admin() returns()
func (_PegKeeper *PegKeeperSession) ApplyNewAdmin() (*types.Transaction, error) {
	return _PegKeeper.Contract.ApplyNewAdmin(&_PegKeeper.TransactOpts)
}

// ApplyNewAdmin is a paid mutator transaction binding the contract method 0xc336fbd5.
//
// Solidity: function apply_new_admin() returns()
func (_PegKeeper *PegKeeperTransactorSession) ApplyNewAdmin() (*types.Transaction, error) {
	return _PegKeeper.Contract.ApplyNewAdmin(&_PegKeeper.TransactOpts)
}

// ApplyNewReceiver is a paid mutator transaction binding the contract method 0xceb667c5.
//
// Solidity: function apply_new_receiver() returns()
func (_PegKeeper *PegKeeperTransactor) ApplyNewReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "apply_new_receiver")
}

// ApplyNewReceiver is a paid mutator transaction binding the contract method 0xceb667c5.
//
// Solidity: function apply_new_receiver() returns()
func (_PegKeeper *PegKeeperSession) ApplyNewReceiver() (*types.Transaction, error) {
	return _PegKeeper.Contract.ApplyNewReceiver(&_PegKeeper.TransactOpts)
}

// ApplyNewReceiver is a paid mutator transaction binding the contract method 0xceb667c5.
//
// Solidity: function apply_new_receiver() returns()
func (_PegKeeper *PegKeeperTransactorSession) ApplyNewReceiver() (*types.Transaction, error) {
	return _PegKeeper.Contract.ApplyNewReceiver(&_PegKeeper.TransactOpts)
}

// CommitNewAdmin is a paid mutator transaction binding the contract method 0xf636b05f.
//
// Solidity: function commit_new_admin(address _new_admin) returns()
func (_PegKeeper *PegKeeperTransactor) CommitNewAdmin(opts *bind.TransactOpts, _new_admin common.Address) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "commit_new_admin", _new_admin)
}

// CommitNewAdmin is a paid mutator transaction binding the contract method 0xf636b05f.
//
// Solidity: function commit_new_admin(address _new_admin) returns()
func (_PegKeeper *PegKeeperSession) CommitNewAdmin(_new_admin common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.CommitNewAdmin(&_PegKeeper.TransactOpts, _new_admin)
}

// CommitNewAdmin is a paid mutator transaction binding the contract method 0xf636b05f.
//
// Solidity: function commit_new_admin(address _new_admin) returns()
func (_PegKeeper *PegKeeperTransactorSession) CommitNewAdmin(_new_admin common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.CommitNewAdmin(&_PegKeeper.TransactOpts, _new_admin)
}

// CommitNewReceiver is a paid mutator transaction binding the contract method 0x561a4147.
//
// Solidity: function commit_new_receiver(address _new_receiver) returns()
func (_PegKeeper *PegKeeperTransactor) CommitNewReceiver(opts *bind.TransactOpts, _new_receiver common.Address) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "commit_new_receiver", _new_receiver)
}

// CommitNewReceiver is a paid mutator transaction binding the contract method 0x561a4147.
//
// Solidity: function commit_new_receiver(address _new_receiver) returns()
func (_PegKeeper *PegKeeperSession) CommitNewReceiver(_new_receiver common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.CommitNewReceiver(&_PegKeeper.TransactOpts, _new_receiver)
}

// CommitNewReceiver is a paid mutator transaction binding the contract method 0x561a4147.
//
// Solidity: function commit_new_receiver(address _new_receiver) returns()
func (_PegKeeper *PegKeeperTransactorSession) CommitNewReceiver(_new_receiver common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.CommitNewReceiver(&_PegKeeper.TransactOpts, _new_receiver)
}

// RevertNewOptions is a paid mutator transaction binding the contract method 0x7bfd8fad.
//
// Solidity: function revert_new_options() returns()
func (_PegKeeper *PegKeeperTransactor) RevertNewOptions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "revert_new_options")
}

// RevertNewOptions is a paid mutator transaction binding the contract method 0x7bfd8fad.
//
// Solidity: function revert_new_options() returns()
func (_PegKeeper *PegKeeperSession) RevertNewOptions() (*types.Transaction, error) {
	return _PegKeeper.Contract.RevertNewOptions(&_PegKeeper.TransactOpts)
}

// RevertNewOptions is a paid mutator transaction binding the contract method 0x7bfd8fad.
//
// Solidity: function revert_new_options() returns()
func (_PegKeeper *PegKeeperTransactorSession) RevertNewOptions() (*types.Transaction, error) {
	return _PegKeeper.Contract.RevertNewOptions(&_PegKeeper.TransactOpts)
}

// SetNewCallerShare is a paid mutator transaction binding the contract method 0x9f5ef86f.
//
// Solidity: function set_new_caller_share(uint256 _new_caller_share) returns()
func (_PegKeeper *PegKeeperTransactor) SetNewCallerShare(opts *bind.TransactOpts, _new_caller_share *big.Int) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "set_new_caller_share", _new_caller_share)
}

// SetNewCallerShare is a paid mutator transaction binding the contract method 0x9f5ef86f.
//
// Solidity: function set_new_caller_share(uint256 _new_caller_share) returns()
func (_PegKeeper *PegKeeperSession) SetNewCallerShare(_new_caller_share *big.Int) (*types.Transaction, error) {
	return _PegKeeper.Contract.SetNewCallerShare(&_PegKeeper.TransactOpts, _new_caller_share)
}

// SetNewCallerShare is a paid mutator transaction binding the contract method 0x9f5ef86f.
//
// Solidity: function set_new_caller_share(uint256 _new_caller_share) returns()
func (_PegKeeper *PegKeeperTransactorSession) SetNewCallerShare(_new_caller_share *big.Int) (*types.Transaction, error) {
	return _PegKeeper.Contract.SetNewCallerShare(&_PegKeeper.TransactOpts, _new_caller_share)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(uint256)
func (_PegKeeper *PegKeeperTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(uint256)
func (_PegKeeper *PegKeeperSession) Update() (*types.Transaction, error) {
	return _PegKeeper.Contract.Update(&_PegKeeper.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns(uint256)
func (_PegKeeper *PegKeeperTransactorSession) Update() (*types.Transaction, error) {
	return _PegKeeper.Contract.Update(&_PegKeeper.TransactOpts)
}

// Update0 is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _beneficiary) returns(uint256)
func (_PegKeeper *PegKeeperTransactor) Update0(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "update0", _beneficiary)
}

// Update0 is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _beneficiary) returns(uint256)
func (_PegKeeper *PegKeeperSession) Update0(_beneficiary common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.Update0(&_PegKeeper.TransactOpts, _beneficiary)
}

// Update0 is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _beneficiary) returns(uint256)
func (_PegKeeper *PegKeeperTransactorSession) Update0(_beneficiary common.Address) (*types.Transaction, error) {
	return _PegKeeper.Contract.Update0(&_PegKeeper.TransactOpts, _beneficiary)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x2c9f7f92.
//
// Solidity: function withdraw_profit() returns(uint256)
func (_PegKeeper *PegKeeperTransactor) WithdrawProfit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeper.contract.Transact(opts, "withdraw_profit")
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x2c9f7f92.
//
// Solidity: function withdraw_profit() returns(uint256)
func (_PegKeeper *PegKeeperSession) WithdrawProfit() (*types.Transaction, error) {
	return _PegKeeper.Contract.WithdrawProfit(&_PegKeeper.TransactOpts)
}

// WithdrawProfit is a paid mutator transaction binding the contract method 0x2c9f7f92.
//
// Solidity: function withdraw_profit() returns(uint256)
func (_PegKeeper *PegKeeperTransactorSession) WithdrawProfit() (*types.Transaction, error) {
	return _PegKeeper.Contract.WithdrawProfit(&_PegKeeper.TransactOpts)
}

// PegKeeperApplyNewAdminIterator is returned from FilterApplyNewAdmin and is used to iterate over the raw logs and unpacked data for ApplyNewAdmin events raised by the PegKeeper contract.
type PegKeeperApplyNewAdminIterator struct {
	Event *PegKeeperApplyNewAdmin // Event containing the contract specifics and raw log

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
func (it *PegKeeperApplyNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperApplyNewAdmin)
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
		it.Event = new(PegKeeperApplyNewAdmin)
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
func (it *PegKeeperApplyNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperApplyNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperApplyNewAdmin represents a ApplyNewAdmin event raised by the PegKeeper contract.
type PegKeeperApplyNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyNewAdmin is a free log retrieval operation binding the contract event 0xc0a6e80c08b9d10c650a9f3e8ff16d7ee0f62bbd5c8ac3b5eba62dae04ec995a.
//
// Solidity: event ApplyNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) FilterApplyNewAdmin(opts *bind.FilterOpts) (*PegKeeperApplyNewAdminIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "ApplyNewAdmin")
	if err != nil {
		return nil, err
	}
	return &PegKeeperApplyNewAdminIterator{contract: _PegKeeper.contract, event: "ApplyNewAdmin", logs: logs, sub: sub}, nil
}

// WatchApplyNewAdmin is a free log subscription operation binding the contract event 0xc0a6e80c08b9d10c650a9f3e8ff16d7ee0f62bbd5c8ac3b5eba62dae04ec995a.
//
// Solidity: event ApplyNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) WatchApplyNewAdmin(opts *bind.WatchOpts, sink chan<- *PegKeeperApplyNewAdmin) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "ApplyNewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperApplyNewAdmin)
				if err := _PegKeeper.contract.UnpackLog(event, "ApplyNewAdmin", log); err != nil {
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

// ParseApplyNewAdmin is a log parse operation binding the contract event 0xc0a6e80c08b9d10c650a9f3e8ff16d7ee0f62bbd5c8ac3b5eba62dae04ec995a.
//
// Solidity: event ApplyNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) ParseApplyNewAdmin(log types.Log) (*PegKeeperApplyNewAdmin, error) {
	event := new(PegKeeperApplyNewAdmin)
	if err := _PegKeeper.contract.UnpackLog(event, "ApplyNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperApplyNewReceiverIterator is returned from FilterApplyNewReceiver and is used to iterate over the raw logs and unpacked data for ApplyNewReceiver events raised by the PegKeeper contract.
type PegKeeperApplyNewReceiverIterator struct {
	Event *PegKeeperApplyNewReceiver // Event containing the contract specifics and raw log

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
func (it *PegKeeperApplyNewReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperApplyNewReceiver)
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
		it.Event = new(PegKeeperApplyNewReceiver)
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
func (it *PegKeeperApplyNewReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperApplyNewReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperApplyNewReceiver represents a ApplyNewReceiver event raised by the PegKeeper contract.
type PegKeeperApplyNewReceiver struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApplyNewReceiver is a free log retrieval operation binding the contract event 0xd9338c71e91ed720d2bd77976c87a0cd6f01ee443a6410481ffdebb9f442db67.
//
// Solidity: event ApplyNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) FilterApplyNewReceiver(opts *bind.FilterOpts) (*PegKeeperApplyNewReceiverIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "ApplyNewReceiver")
	if err != nil {
		return nil, err
	}
	return &PegKeeperApplyNewReceiverIterator{contract: _PegKeeper.contract, event: "ApplyNewReceiver", logs: logs, sub: sub}, nil
}

// WatchApplyNewReceiver is a free log subscription operation binding the contract event 0xd9338c71e91ed720d2bd77976c87a0cd6f01ee443a6410481ffdebb9f442db67.
//
// Solidity: event ApplyNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) WatchApplyNewReceiver(opts *bind.WatchOpts, sink chan<- *PegKeeperApplyNewReceiver) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "ApplyNewReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperApplyNewReceiver)
				if err := _PegKeeper.contract.UnpackLog(event, "ApplyNewReceiver", log); err != nil {
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

// ParseApplyNewReceiver is a log parse operation binding the contract event 0xd9338c71e91ed720d2bd77976c87a0cd6f01ee443a6410481ffdebb9f442db67.
//
// Solidity: event ApplyNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) ParseApplyNewReceiver(log types.Log) (*PegKeeperApplyNewReceiver, error) {
	event := new(PegKeeperApplyNewReceiver)
	if err := _PegKeeper.contract.UnpackLog(event, "ApplyNewReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the PegKeeper contract.
type PegKeeperCommitNewAdminIterator struct {
	Event *PegKeeperCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *PegKeeperCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperCommitNewAdmin)
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
		it.Event = new(PegKeeperCommitNewAdmin)
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
func (it *PegKeeperCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperCommitNewAdmin represents a CommitNewAdmin event raised by the PegKeeper contract.
type PegKeeperCommitNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x0305c49816a5a9a099d81e90d76421c9a4a529e0640cd15297b6fc8f1c9ac6ff.
//
// Solidity: event CommitNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts) (*PegKeeperCommitNewAdminIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "CommitNewAdmin")
	if err != nil {
		return nil, err
	}
	return &PegKeeperCommitNewAdminIterator{contract: _PegKeeper.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x0305c49816a5a9a099d81e90d76421c9a4a529e0640cd15297b6fc8f1c9ac6ff.
//
// Solidity: event CommitNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *PegKeeperCommitNewAdmin) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "CommitNewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperCommitNewAdmin)
				if err := _PegKeeper.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x0305c49816a5a9a099d81e90d76421c9a4a529e0640cd15297b6fc8f1c9ac6ff.
//
// Solidity: event CommitNewAdmin(address admin)
func (_PegKeeper *PegKeeperFilterer) ParseCommitNewAdmin(log types.Log) (*PegKeeperCommitNewAdmin, error) {
	event := new(PegKeeperCommitNewAdmin)
	if err := _PegKeeper.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperCommitNewReceiverIterator is returned from FilterCommitNewReceiver and is used to iterate over the raw logs and unpacked data for CommitNewReceiver events raised by the PegKeeper contract.
type PegKeeperCommitNewReceiverIterator struct {
	Event *PegKeeperCommitNewReceiver // Event containing the contract specifics and raw log

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
func (it *PegKeeperCommitNewReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperCommitNewReceiver)
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
		it.Event = new(PegKeeperCommitNewReceiver)
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
func (it *PegKeeperCommitNewReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperCommitNewReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperCommitNewReceiver represents a CommitNewReceiver event raised by the PegKeeper contract.
type PegKeeperCommitNewReceiver struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewReceiver is a free log retrieval operation binding the contract event 0xf918f289f5185efe74d8bdf25364dbbc9c83a2bc438ede64b939cf6751ccf27c.
//
// Solidity: event CommitNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) FilterCommitNewReceiver(opts *bind.FilterOpts) (*PegKeeperCommitNewReceiverIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "CommitNewReceiver")
	if err != nil {
		return nil, err
	}
	return &PegKeeperCommitNewReceiverIterator{contract: _PegKeeper.contract, event: "CommitNewReceiver", logs: logs, sub: sub}, nil
}

// WatchCommitNewReceiver is a free log subscription operation binding the contract event 0xf918f289f5185efe74d8bdf25364dbbc9c83a2bc438ede64b939cf6751ccf27c.
//
// Solidity: event CommitNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) WatchCommitNewReceiver(opts *bind.WatchOpts, sink chan<- *PegKeeperCommitNewReceiver) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "CommitNewReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperCommitNewReceiver)
				if err := _PegKeeper.contract.UnpackLog(event, "CommitNewReceiver", log); err != nil {
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

// ParseCommitNewReceiver is a log parse operation binding the contract event 0xf918f289f5185efe74d8bdf25364dbbc9c83a2bc438ede64b939cf6751ccf27c.
//
// Solidity: event CommitNewReceiver(address receiver)
func (_PegKeeper *PegKeeperFilterer) ParseCommitNewReceiver(log types.Log) (*PegKeeperCommitNewReceiver, error) {
	event := new(PegKeeperCommitNewReceiver)
	if err := _PegKeeper.contract.UnpackLog(event, "CommitNewReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperProfitIterator is returned from FilterProfit and is used to iterate over the raw logs and unpacked data for Profit events raised by the PegKeeper contract.
type PegKeeperProfitIterator struct {
	Event *PegKeeperProfit // Event containing the contract specifics and raw log

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
func (it *PegKeeperProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperProfit)
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
		it.Event = new(PegKeeperProfit)
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
func (it *PegKeeperProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperProfit represents a Profit event raised by the PegKeeper contract.
type PegKeeperProfit struct {
	LpAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProfit is a free log retrieval operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 lp_amount)
func (_PegKeeper *PegKeeperFilterer) FilterProfit(opts *bind.FilterOpts) (*PegKeeperProfitIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "Profit")
	if err != nil {
		return nil, err
	}
	return &PegKeeperProfitIterator{contract: _PegKeeper.contract, event: "Profit", logs: logs, sub: sub}, nil
}

// WatchProfit is a free log subscription operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 lp_amount)
func (_PegKeeper *PegKeeperFilterer) WatchProfit(opts *bind.WatchOpts, sink chan<- *PegKeeperProfit) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "Profit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperProfit)
				if err := _PegKeeper.contract.UnpackLog(event, "Profit", log); err != nil {
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

// ParseProfit is a log parse operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 lp_amount)
func (_PegKeeper *PegKeeperFilterer) ParseProfit(log types.Log) (*PegKeeperProfit, error) {
	event := new(PegKeeperProfit)
	if err := _PegKeeper.contract.UnpackLog(event, "Profit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperProvideIterator is returned from FilterProvide and is used to iterate over the raw logs and unpacked data for Provide events raised by the PegKeeper contract.
type PegKeeperProvideIterator struct {
	Event *PegKeeperProvide // Event containing the contract specifics and raw log

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
func (it *PegKeeperProvideIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperProvide)
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
		it.Event = new(PegKeeperProvide)
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
func (it *PegKeeperProvideIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperProvideIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperProvide represents a Provide event raised by the PegKeeper contract.
type PegKeeperProvide struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProvide is a free log retrieval operation binding the contract event 0x8d685bd3f45d861c759ed7a46ea3d30eb5cc6ce9fe06c526931f94c963bca7d2.
//
// Solidity: event Provide(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) FilterProvide(opts *bind.FilterOpts) (*PegKeeperProvideIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "Provide")
	if err != nil {
		return nil, err
	}
	return &PegKeeperProvideIterator{contract: _PegKeeper.contract, event: "Provide", logs: logs, sub: sub}, nil
}

// WatchProvide is a free log subscription operation binding the contract event 0x8d685bd3f45d861c759ed7a46ea3d30eb5cc6ce9fe06c526931f94c963bca7d2.
//
// Solidity: event Provide(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) WatchProvide(opts *bind.WatchOpts, sink chan<- *PegKeeperProvide) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "Provide")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperProvide)
				if err := _PegKeeper.contract.UnpackLog(event, "Provide", log); err != nil {
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

// ParseProvide is a log parse operation binding the contract event 0x8d685bd3f45d861c759ed7a46ea3d30eb5cc6ce9fe06c526931f94c963bca7d2.
//
// Solidity: event Provide(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) ParseProvide(log types.Log) (*PegKeeperProvide, error) {
	event := new(PegKeeperProvide)
	if err := _PegKeeper.contract.UnpackLog(event, "Provide", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperSetNewCallerShareIterator is returned from FilterSetNewCallerShare and is used to iterate over the raw logs and unpacked data for SetNewCallerShare events raised by the PegKeeper contract.
type PegKeeperSetNewCallerShareIterator struct {
	Event *PegKeeperSetNewCallerShare // Event containing the contract specifics and raw log

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
func (it *PegKeeperSetNewCallerShareIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperSetNewCallerShare)
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
		it.Event = new(PegKeeperSetNewCallerShare)
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
func (it *PegKeeperSetNewCallerShareIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperSetNewCallerShareIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperSetNewCallerShare represents a SetNewCallerShare event raised by the PegKeeper contract.
type PegKeeperSetNewCallerShare struct {
	CallerShare *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNewCallerShare is a free log retrieval operation binding the contract event 0x43abfd568a12e5f6ea230e56ba11b83e29e9921019211df9cc2e216ba9f225dc.
//
// Solidity: event SetNewCallerShare(uint256 caller_share)
func (_PegKeeper *PegKeeperFilterer) FilterSetNewCallerShare(opts *bind.FilterOpts) (*PegKeeperSetNewCallerShareIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "SetNewCallerShare")
	if err != nil {
		return nil, err
	}
	return &PegKeeperSetNewCallerShareIterator{contract: _PegKeeper.contract, event: "SetNewCallerShare", logs: logs, sub: sub}, nil
}

// WatchSetNewCallerShare is a free log subscription operation binding the contract event 0x43abfd568a12e5f6ea230e56ba11b83e29e9921019211df9cc2e216ba9f225dc.
//
// Solidity: event SetNewCallerShare(uint256 caller_share)
func (_PegKeeper *PegKeeperFilterer) WatchSetNewCallerShare(opts *bind.WatchOpts, sink chan<- *PegKeeperSetNewCallerShare) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "SetNewCallerShare")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperSetNewCallerShare)
				if err := _PegKeeper.contract.UnpackLog(event, "SetNewCallerShare", log); err != nil {
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

// ParseSetNewCallerShare is a log parse operation binding the contract event 0x43abfd568a12e5f6ea230e56ba11b83e29e9921019211df9cc2e216ba9f225dc.
//
// Solidity: event SetNewCallerShare(uint256 caller_share)
func (_PegKeeper *PegKeeperFilterer) ParseSetNewCallerShare(log types.Log) (*PegKeeperSetNewCallerShare, error) {
	event := new(PegKeeperSetNewCallerShare)
	if err := _PegKeeper.contract.UnpackLog(event, "SetNewCallerShare", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PegKeeper contract.
type PegKeeperWithdrawIterator struct {
	Event *PegKeeperWithdraw // Event containing the contract specifics and raw log

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
func (it *PegKeeperWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperWithdraw)
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
		it.Event = new(PegKeeperWithdraw)
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
func (it *PegKeeperWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperWithdraw represents a Withdraw event raised by the PegKeeper contract.
type PegKeeperWithdraw struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) FilterWithdraw(opts *bind.FilterOpts) (*PegKeeperWithdrawIterator, error) {

	logs, sub, err := _PegKeeper.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &PegKeeperWithdrawIterator{contract: _PegKeeper.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PegKeeperWithdraw) (event.Subscription, error) {

	logs, sub, err := _PegKeeper.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperWithdraw)
				if err := _PegKeeper.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 amount)
func (_PegKeeper *PegKeeperFilterer) ParseWithdraw(log types.Log) (*PegKeeperWithdraw, error) {
	event := new(PegKeeperWithdraw)
	if err := _PegKeeper.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
