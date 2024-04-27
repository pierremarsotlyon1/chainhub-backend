// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveGCComplete

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

// CurveGCCompleteMetaData contains all meta data concerning the CurveGCComplete contract.
var CurveGCCompleteMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddType\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"type_id\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewTypeWeight\",\"inputs\":[{\"type\":\"int128\",\"name\":\"type_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"weight\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"total_weight\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewGaugeWeight\",\"inputs\":[{\"type\":\"address\",\"name\":\"gauge_address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"weight\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"total_weight\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"VoteForGauge\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"address\",\"name\":\"user\",\"indexed\":false},{\"type\":\"address\",\"name\":\"gauge_addr\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"weight\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewGauge\",\"inputs\":[{\"type\":\"address\",\"name\":\"addr\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"gauge_type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"weight\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_voting_escrow\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37597},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38497},{\"name\":\"gauge_types\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1625},{\"name\":\"add_gauge\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"int128\",\"name\":\"gauge_type\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"add_gauge\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"int128\",\"name\":\"gauge_type\"},{\"type\":\"uint256\",\"name\":\"weight\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":18033784416},{\"name\":\"checkpoint_gauge\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":18087678795},{\"name\":\"gauge_relative_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"gauge_relative_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"time\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"gauge_relative_weight_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"gauge_relative_weight_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"add_type\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"add_type\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"uint256\",\"name\":\"weight\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"change_type_weight\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"type_id\"},{\"type\":\"uint256\",\"name\":\"weight\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36246310050},{\"name\":\"change_gauge_weight\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"weight\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36354170809},{\"name\":\"vote_for_gauge_weights\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_gauge_addr\"},{\"type\":\"uint256\",\"name\":\"_user_weight\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":18142052127},{\"name\":\"get_gauge_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2974},{\"name\":\"get_type_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"type_id\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2977},{\"name\":\"get_total_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2693},{\"name\":\"get_weights_sum_per_type\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"type_id\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3109},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"voting_escrow\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"n_gauge_types\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"n_gauges\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1991},{\"name\":\"gauge_type_names\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8628},{\"name\":\"gauges\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2160},{\"name\":\"vote_user_slopes\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"power\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5020},{\"name\":\"vote_user_power\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2265},{\"name\":\"last_user_vote\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2449},{\"name\":\"points_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"bias\"},{\"type\":\"uint256\",\"name\":\"slope\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3859},{\"name\":\"time_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2355},{\"name\":\"points_sum\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"bias\"},{\"type\":\"uint256\",\"name\":\"slope\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3970},{\"name\":\"time_sum\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2370},{\"name\":\"points_total\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2406},{\"name\":\"time_total\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"points_type_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2671},{\"name\":\"time_type_weight\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2490}]",
}

// CurveGCCompleteABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGCCompleteMetaData.ABI instead.
var CurveGCCompleteABI = CurveGCCompleteMetaData.ABI

// CurveGCComplete is an auto generated Go binding around an Ethereum contract.
type CurveGCComplete struct {
	CurveGCCompleteCaller     // Read-only binding to the contract
	CurveGCCompleteTransactor // Write-only binding to the contract
	CurveGCCompleteFilterer   // Log filterer for contract events
}

// CurveGCCompleteCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGCCompleteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCCompleteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGCCompleteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCCompleteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGCCompleteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCCompleteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGCCompleteSession struct {
	Contract     *CurveGCComplete  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveGCCompleteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGCCompleteCallerSession struct {
	Contract *CurveGCCompleteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CurveGCCompleteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGCCompleteTransactorSession struct {
	Contract     *CurveGCCompleteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CurveGCCompleteRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGCCompleteRaw struct {
	Contract *CurveGCComplete // Generic contract binding to access the raw methods on
}

// CurveGCCompleteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGCCompleteCallerRaw struct {
	Contract *CurveGCCompleteCaller // Generic read-only contract binding to access the raw methods on
}

// CurveGCCompleteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGCCompleteTransactorRaw struct {
	Contract *CurveGCCompleteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGCComplete creates a new instance of CurveGCComplete, bound to a specific deployed contract.
func NewCurveGCComplete(address common.Address, backend bind.ContractBackend) (*CurveGCComplete, error) {
	contract, err := bindCurveGCComplete(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGCComplete{CurveGCCompleteCaller: CurveGCCompleteCaller{contract: contract}, CurveGCCompleteTransactor: CurveGCCompleteTransactor{contract: contract}, CurveGCCompleteFilterer: CurveGCCompleteFilterer{contract: contract}}, nil
}

// NewCurveGCCompleteCaller creates a new read-only instance of CurveGCComplete, bound to a specific deployed contract.
func NewCurveGCCompleteCaller(address common.Address, caller bind.ContractCaller) (*CurveGCCompleteCaller, error) {
	contract, err := bindCurveGCComplete(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteCaller{contract: contract}, nil
}

// NewCurveGCCompleteTransactor creates a new write-only instance of CurveGCComplete, bound to a specific deployed contract.
func NewCurveGCCompleteTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveGCCompleteTransactor, error) {
	contract, err := bindCurveGCComplete(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteTransactor{contract: contract}, nil
}

// NewCurveGCCompleteFilterer creates a new log filterer instance of CurveGCComplete, bound to a specific deployed contract.
func NewCurveGCCompleteFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveGCCompleteFilterer, error) {
	contract, err := bindCurveGCComplete(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteFilterer{contract: contract}, nil
}

// bindCurveGCComplete binds a generic wrapper to an already deployed contract.
func bindCurveGCComplete(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveGCCompleteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGCComplete *CurveGCCompleteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGCComplete.Contract.CurveGCCompleteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGCComplete *CurveGCCompleteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CurveGCCompleteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGCComplete *CurveGCCompleteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CurveGCCompleteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGCComplete *CurveGCCompleteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGCComplete.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGCComplete *CurveGCCompleteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGCComplete *CurveGCCompleteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteSession) Admin() (common.Address, error) {
	return _CurveGCComplete.Contract.Admin(&_CurveGCComplete.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCallerSession) Admin() (common.Address, error) {
	return _CurveGCComplete.Contract.Admin(&_CurveGCComplete.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteSession) FutureAdmin() (common.Address, error) {
	return _CurveGCComplete.Contract.FutureAdmin(&_CurveGCComplete.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveGCComplete.Contract.FutureAdmin(&_CurveGCComplete.CallOpts)
}

// GaugeRelativeWeight is a free data retrieval call binding the contract method 0x6207d866.
//
// Solidity: function gauge_relative_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GaugeRelativeWeight(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "gauge_relative_weight", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaugeRelativeWeight is a free data retrieval call binding the contract method 0x6207d866.
//
// Solidity: function gauge_relative_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeRelativeWeight(addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeight(&_CurveGCComplete.CallOpts, addr)
}

// GaugeRelativeWeight is a free data retrieval call binding the contract method 0x6207d866.
//
// Solidity: function gauge_relative_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GaugeRelativeWeight(addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeight(&_CurveGCComplete.CallOpts, addr)
}

// GaugeRelativeWeight0 is a free data retrieval call binding the contract method 0xd3078c94.
//
// Solidity: function gauge_relative_weight(address addr, uint256 time) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GaugeRelativeWeight0(opts *bind.CallOpts, addr common.Address, time *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "gauge_relative_weight0", addr, time)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaugeRelativeWeight0 is a free data retrieval call binding the contract method 0xd3078c94.
//
// Solidity: function gauge_relative_weight(address addr, uint256 time) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeRelativeWeight0(addr common.Address, time *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeight0(&_CurveGCComplete.CallOpts, addr, time)
}

// GaugeRelativeWeight0 is a free data retrieval call binding the contract method 0xd3078c94.
//
// Solidity: function gauge_relative_weight(address addr, uint256 time) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GaugeRelativeWeight0(addr common.Address, time *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeight0(&_CurveGCComplete.CallOpts, addr, time)
}

// GaugeTypeNames is a free data retrieval call binding the contract method 0xd958a8fc.
//
// Solidity: function gauge_type_names(int128 arg0) view returns(string)
func (_CurveGCComplete *CurveGCCompleteCaller) GaugeTypeNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "gauge_type_names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GaugeTypeNames is a free data retrieval call binding the contract method 0xd958a8fc.
//
// Solidity: function gauge_type_names(int128 arg0) view returns(string)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeTypeNames(arg0 *big.Int) (string, error) {
	return _CurveGCComplete.Contract.GaugeTypeNames(&_CurveGCComplete.CallOpts, arg0)
}

// GaugeTypeNames is a free data retrieval call binding the contract method 0xd958a8fc.
//
// Solidity: function gauge_type_names(int128 arg0) view returns(string)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GaugeTypeNames(arg0 *big.Int) (string, error) {
	return _CurveGCComplete.Contract.GaugeTypeNames(&_CurveGCComplete.CallOpts, arg0)
}

// GaugeTypes is a free data retrieval call binding the contract method 0x3f9095b7.
//
// Solidity: function gauge_types(address _addr) view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCaller) GaugeTypes(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "gauge_types", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaugeTypes is a free data retrieval call binding the contract method 0x3f9095b7.
//
// Solidity: function gauge_types(address _addr) view returns(int128)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeTypes(_addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeTypes(&_CurveGCComplete.CallOpts, _addr)
}

// GaugeTypes is a free data retrieval call binding the contract method 0x3f9095b7.
//
// Solidity: function gauge_types(address _addr) view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GaugeTypes(_addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GaugeTypes(&_CurveGCComplete.CallOpts, _addr)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_CurveGCComplete *CurveGCCompleteCaller) Gauges(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "gauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_CurveGCComplete *CurveGCCompleteSession) Gauges(arg0 *big.Int) (common.Address, error) {
	return _CurveGCComplete.Contract.Gauges(&_CurveGCComplete.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_CurveGCComplete *CurveGCCompleteCallerSession) Gauges(arg0 *big.Int) (common.Address, error) {
	return _CurveGCComplete.Contract.Gauges(&_CurveGCComplete.CallOpts, arg0)
}

// GetGaugeWeight is a free data retrieval call binding the contract method 0x4e791a3a.
//
// Solidity: function get_gauge_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GetGaugeWeight(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "get_gauge_weight", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeWeight is a free data retrieval call binding the contract method 0x4e791a3a.
//
// Solidity: function get_gauge_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GetGaugeWeight(addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetGaugeWeight(&_CurveGCComplete.CallOpts, addr)
}

// GetGaugeWeight is a free data retrieval call binding the contract method 0x4e791a3a.
//
// Solidity: function get_gauge_weight(address addr) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GetGaugeWeight(addr common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetGaugeWeight(&_CurveGCComplete.CallOpts, addr)
}

// GetTotalWeight is a free data retrieval call binding the contract method 0x6977ff92.
//
// Solidity: function get_total_weight() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GetTotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "get_total_weight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWeight is a free data retrieval call binding the contract method 0x6977ff92.
//
// Solidity: function get_total_weight() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GetTotalWeight() (*big.Int, error) {
	return _CurveGCComplete.Contract.GetTotalWeight(&_CurveGCComplete.CallOpts)
}

// GetTotalWeight is a free data retrieval call binding the contract method 0x6977ff92.
//
// Solidity: function get_total_weight() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GetTotalWeight() (*big.Int, error) {
	return _CurveGCComplete.Contract.GetTotalWeight(&_CurveGCComplete.CallOpts)
}

// GetTypeWeight is a free data retrieval call binding the contract method 0x72fdccfa.
//
// Solidity: function get_type_weight(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GetTypeWeight(opts *bind.CallOpts, type_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "get_type_weight", type_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTypeWeight is a free data retrieval call binding the contract method 0x72fdccfa.
//
// Solidity: function get_type_weight(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GetTypeWeight(type_id *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetTypeWeight(&_CurveGCComplete.CallOpts, type_id)
}

// GetTypeWeight is a free data retrieval call binding the contract method 0x72fdccfa.
//
// Solidity: function get_type_weight(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GetTypeWeight(type_id *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetTypeWeight(&_CurveGCComplete.CallOpts, type_id)
}

// GetWeightsSumPerType is a free data retrieval call binding the contract method 0x6f214a6a.
//
// Solidity: function get_weights_sum_per_type(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) GetWeightsSumPerType(opts *bind.CallOpts, type_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "get_weights_sum_per_type", type_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeightsSumPerType is a free data retrieval call binding the contract method 0x6f214a6a.
//
// Solidity: function get_weights_sum_per_type(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GetWeightsSumPerType(type_id *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetWeightsSumPerType(&_CurveGCComplete.CallOpts, type_id)
}

// GetWeightsSumPerType is a free data retrieval call binding the contract method 0x6f214a6a.
//
// Solidity: function get_weights_sum_per_type(int128 type_id) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) GetWeightsSumPerType(type_id *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.GetWeightsSumPerType(&_CurveGCComplete.CallOpts, type_id)
}

// LastUserVote is a free data retrieval call binding the contract method 0x7e418fa0.
//
// Solidity: function last_user_vote(address arg0, address arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) LastUserVote(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "last_user_vote", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUserVote is a free data retrieval call binding the contract method 0x7e418fa0.
//
// Solidity: function last_user_vote(address arg0, address arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) LastUserVote(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.LastUserVote(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// LastUserVote is a free data retrieval call binding the contract method 0x7e418fa0.
//
// Solidity: function last_user_vote(address arg0, address arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) LastUserVote(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.LastUserVote(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// NGaugeTypes is a free data retrieval call binding the contract method 0x9fba03a1.
//
// Solidity: function n_gauge_types() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCaller) NGaugeTypes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "n_gauge_types")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NGaugeTypes is a free data retrieval call binding the contract method 0x9fba03a1.
//
// Solidity: function n_gauge_types() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteSession) NGaugeTypes() (*big.Int, error) {
	return _CurveGCComplete.Contract.NGaugeTypes(&_CurveGCComplete.CallOpts)
}

// NGaugeTypes is a free data retrieval call binding the contract method 0x9fba03a1.
//
// Solidity: function n_gauge_types() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCallerSession) NGaugeTypes() (*big.Int, error) {
	return _CurveGCComplete.Contract.NGaugeTypes(&_CurveGCComplete.CallOpts)
}

// NGauges is a free data retrieval call binding the contract method 0xe93841d0.
//
// Solidity: function n_gauges() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCaller) NGauges(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "n_gauges")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NGauges is a free data retrieval call binding the contract method 0xe93841d0.
//
// Solidity: function n_gauges() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteSession) NGauges() (*big.Int, error) {
	return _CurveGCComplete.Contract.NGauges(&_CurveGCComplete.CallOpts)
}

// NGauges is a free data retrieval call binding the contract method 0xe93841d0.
//
// Solidity: function n_gauges() view returns(int128)
func (_CurveGCComplete *CurveGCCompleteCallerSession) NGauges() (*big.Int, error) {
	return _CurveGCComplete.Contract.NGauges(&_CurveGCComplete.CallOpts)
}

// PointsSum is a free data retrieval call binding the contract method 0xa9b48c01.
//
// Solidity: function points_sum(int128 arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteCaller) PointsSum(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "points_sum", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointsSum is a free data retrieval call binding the contract method 0xa9b48c01.
//
// Solidity: function points_sum(int128 arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteSession) PointsSum(arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	return _CurveGCComplete.Contract.PointsSum(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// PointsSum is a free data retrieval call binding the contract method 0xa9b48c01.
//
// Solidity: function points_sum(int128 arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteCallerSession) PointsSum(arg0 *big.Int, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	return _CurveGCComplete.Contract.PointsSum(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// PointsTotal is a free data retrieval call binding the contract method 0x1142916b.
//
// Solidity: function points_total(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) PointsTotal(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "points_total", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointsTotal is a free data retrieval call binding the contract method 0x1142916b.
//
// Solidity: function points_total(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) PointsTotal(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.PointsTotal(&_CurveGCComplete.CallOpts, arg0)
}

// PointsTotal is a free data retrieval call binding the contract method 0x1142916b.
//
// Solidity: function points_total(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) PointsTotal(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.PointsTotal(&_CurveGCComplete.CallOpts, arg0)
}

// PointsTypeWeight is a free data retrieval call binding the contract method 0xafd2bb49.
//
// Solidity: function points_type_weight(int128 arg0, uint256 arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) PointsTypeWeight(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "points_type_weight", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointsTypeWeight is a free data retrieval call binding the contract method 0xafd2bb49.
//
// Solidity: function points_type_weight(int128 arg0, uint256 arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) PointsTypeWeight(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.PointsTypeWeight(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// PointsTypeWeight is a free data retrieval call binding the contract method 0xafd2bb49.
//
// Solidity: function points_type_weight(int128 arg0, uint256 arg1) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) PointsTypeWeight(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.PointsTypeWeight(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// PointsWeight is a free data retrieval call binding the contract method 0xedba5273.
//
// Solidity: function points_weight(address arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteCaller) PointsWeight(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "points_weight", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointsWeight is a free data retrieval call binding the contract method 0xedba5273.
//
// Solidity: function points_weight(address arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteSession) PointsWeight(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	return _CurveGCComplete.Contract.PointsWeight(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// PointsWeight is a free data retrieval call binding the contract method 0xedba5273.
//
// Solidity: function points_weight(address arg0, uint256 arg1) view returns(uint256 bias, uint256 slope)
func (_CurveGCComplete *CurveGCCompleteCallerSession) PointsWeight(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
}, error) {
	return _CurveGCComplete.Contract.PointsWeight(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// TimeSum is a free data retrieval call binding the contract method 0x5a549158.
//
// Solidity: function time_sum(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) TimeSum(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "time_sum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeSum is a free data retrieval call binding the contract method 0x5a549158.
//
// Solidity: function time_sum(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) TimeSum(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeSum(&_CurveGCComplete.CallOpts, arg0)
}

// TimeSum is a free data retrieval call binding the contract method 0x5a549158.
//
// Solidity: function time_sum(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) TimeSum(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeSum(&_CurveGCComplete.CallOpts, arg0)
}

// TimeTotal is a free data retrieval call binding the contract method 0x513872bd.
//
// Solidity: function time_total() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) TimeTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "time_total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeTotal is a free data retrieval call binding the contract method 0x513872bd.
//
// Solidity: function time_total() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) TimeTotal() (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeTotal(&_CurveGCComplete.CallOpts)
}

// TimeTotal is a free data retrieval call binding the contract method 0x513872bd.
//
// Solidity: function time_total() view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) TimeTotal() (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeTotal(&_CurveGCComplete.CallOpts)
}

// TimeTypeWeight is a free data retrieval call binding the contract method 0x51ce6b59.
//
// Solidity: function time_type_weight(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) TimeTypeWeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "time_type_weight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeTypeWeight is a free data retrieval call binding the contract method 0x51ce6b59.
//
// Solidity: function time_type_weight(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) TimeTypeWeight(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeTypeWeight(&_CurveGCComplete.CallOpts, arg0)
}

// TimeTypeWeight is a free data retrieval call binding the contract method 0x51ce6b59.
//
// Solidity: function time_type_weight(uint256 arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) TimeTypeWeight(arg0 *big.Int) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeTypeWeight(&_CurveGCComplete.CallOpts, arg0)
}

// TimeWeight is a free data retrieval call binding the contract method 0xa4d7a250.
//
// Solidity: function time_weight(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) TimeWeight(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "time_weight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeWeight is a free data retrieval call binding the contract method 0xa4d7a250.
//
// Solidity: function time_weight(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) TimeWeight(arg0 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeWeight(&_CurveGCComplete.CallOpts, arg0)
}

// TimeWeight is a free data retrieval call binding the contract method 0xa4d7a250.
//
// Solidity: function time_weight(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) TimeWeight(arg0 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.TimeWeight(&_CurveGCComplete.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveGCComplete *CurveGCCompleteSession) Token() (common.Address, error) {
	return _CurveGCComplete.Contract.Token(&_CurveGCComplete.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCallerSession) Token() (common.Address, error) {
	return _CurveGCComplete.Contract.Token(&_CurveGCComplete.CallOpts)
}

// VoteUserPower is a free data retrieval call binding the contract method 0x411e74b5.
//
// Solidity: function vote_user_power(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCaller) VoteUserPower(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "vote_user_power", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteUserPower is a free data retrieval call binding the contract method 0x411e74b5.
//
// Solidity: function vote_user_power(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) VoteUserPower(arg0 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.VoteUserPower(&_CurveGCComplete.CallOpts, arg0)
}

// VoteUserPower is a free data retrieval call binding the contract method 0x411e74b5.
//
// Solidity: function vote_user_power(address arg0) view returns(uint256)
func (_CurveGCComplete *CurveGCCompleteCallerSession) VoteUserPower(arg0 common.Address) (*big.Int, error) {
	return _CurveGCComplete.Contract.VoteUserPower(&_CurveGCComplete.CallOpts, arg0)
}

// VoteUserSlopes is a free data retrieval call binding the contract method 0x0f467f98.
//
// Solidity: function vote_user_slopes(address arg0, address arg1) view returns(uint256 slope, uint256 power, uint256 end)
func (_CurveGCComplete *CurveGCCompleteCaller) VoteUserSlopes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Slope *big.Int
	Power *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "vote_user_slopes", arg0, arg1)

	outstruct := new(struct {
		Slope *big.Int
		Power *big.Int
		End   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Slope = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Power = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VoteUserSlopes is a free data retrieval call binding the contract method 0x0f467f98.
//
// Solidity: function vote_user_slopes(address arg0, address arg1) view returns(uint256 slope, uint256 power, uint256 end)
func (_CurveGCComplete *CurveGCCompleteSession) VoteUserSlopes(arg0 common.Address, arg1 common.Address) (struct {
	Slope *big.Int
	Power *big.Int
	End   *big.Int
}, error) {
	return _CurveGCComplete.Contract.VoteUserSlopes(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// VoteUserSlopes is a free data retrieval call binding the contract method 0x0f467f98.
//
// Solidity: function vote_user_slopes(address arg0, address arg1) view returns(uint256 slope, uint256 power, uint256 end)
func (_CurveGCComplete *CurveGCCompleteCallerSession) VoteUserSlopes(arg0 common.Address, arg1 common.Address) (struct {
	Slope *big.Int
	Power *big.Int
	End   *big.Int
}, error) {
	return _CurveGCComplete.Contract.VoteUserSlopes(&_CurveGCComplete.CallOpts, arg0, arg1)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGCComplete.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGCComplete *CurveGCCompleteSession) VotingEscrow() (common.Address, error) {
	return _CurveGCComplete.Contract.VotingEscrow(&_CurveGCComplete.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveGCComplete *CurveGCCompleteCallerSession) VotingEscrow() (common.Address, error) {
	return _CurveGCComplete.Contract.VotingEscrow(&_CurveGCComplete.CallOpts)
}

// AddGauge is a paid mutator transaction binding the contract method 0x3a04f900.
//
// Solidity: function add_gauge(address addr, int128 gauge_type) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) AddGauge(opts *bind.TransactOpts, addr common.Address, gauge_type *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "add_gauge", addr, gauge_type)
}

// AddGauge is a paid mutator transaction binding the contract method 0x3a04f900.
//
// Solidity: function add_gauge(address addr, int128 gauge_type) returns()
func (_CurveGCComplete *CurveGCCompleteSession) AddGauge(addr common.Address, gauge_type *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddGauge(&_CurveGCComplete.TransactOpts, addr, gauge_type)
}

// AddGauge is a paid mutator transaction binding the contract method 0x3a04f900.
//
// Solidity: function add_gauge(address addr, int128 gauge_type) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) AddGauge(addr common.Address, gauge_type *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddGauge(&_CurveGCComplete.TransactOpts, addr, gauge_type)
}

// AddGauge0 is a paid mutator transaction binding the contract method 0x18dfe921.
//
// Solidity: function add_gauge(address addr, int128 gauge_type, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) AddGauge0(opts *bind.TransactOpts, addr common.Address, gauge_type *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "add_gauge0", addr, gauge_type, weight)
}

// AddGauge0 is a paid mutator transaction binding the contract method 0x18dfe921.
//
// Solidity: function add_gauge(address addr, int128 gauge_type, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteSession) AddGauge0(addr common.Address, gauge_type *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddGauge0(&_CurveGCComplete.TransactOpts, addr, gauge_type, weight)
}

// AddGauge0 is a paid mutator transaction binding the contract method 0x18dfe921.
//
// Solidity: function add_gauge(address addr, int128 gauge_type, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) AddGauge0(addr common.Address, gauge_type *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddGauge0(&_CurveGCComplete.TransactOpts, addr, gauge_type, weight)
}

// AddType is a paid mutator transaction binding the contract method 0x26e56d5e.
//
// Solidity: function add_type(string _name) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) AddType(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "add_type", _name)
}

// AddType is a paid mutator transaction binding the contract method 0x26e56d5e.
//
// Solidity: function add_type(string _name) returns()
func (_CurveGCComplete *CurveGCCompleteSession) AddType(_name string) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddType(&_CurveGCComplete.TransactOpts, _name)
}

// AddType is a paid mutator transaction binding the contract method 0x26e56d5e.
//
// Solidity: function add_type(string _name) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) AddType(_name string) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddType(&_CurveGCComplete.TransactOpts, _name)
}

// AddType0 is a paid mutator transaction binding the contract method 0x92d0d232.
//
// Solidity: function add_type(string _name, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) AddType0(opts *bind.TransactOpts, _name string, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "add_type0", _name, weight)
}

// AddType0 is a paid mutator transaction binding the contract method 0x92d0d232.
//
// Solidity: function add_type(string _name, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteSession) AddType0(_name string, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddType0(&_CurveGCComplete.TransactOpts, _name, weight)
}

// AddType0 is a paid mutator transaction binding the contract method 0x92d0d232.
//
// Solidity: function add_type(string _name, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) AddType0(_name string, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.AddType0(&_CurveGCComplete.TransactOpts, _name, weight)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveGCComplete *CurveGCCompleteSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ApplyTransferOwnership(&_CurveGCComplete.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ApplyTransferOwnership(&_CurveGCComplete.TransactOpts)
}

// ChangeGaugeWeight is a paid mutator transaction binding the contract method 0xd4d2646e.
//
// Solidity: function change_gauge_weight(address addr, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) ChangeGaugeWeight(opts *bind.TransactOpts, addr common.Address, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "change_gauge_weight", addr, weight)
}

// ChangeGaugeWeight is a paid mutator transaction binding the contract method 0xd4d2646e.
//
// Solidity: function change_gauge_weight(address addr, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteSession) ChangeGaugeWeight(addr common.Address, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ChangeGaugeWeight(&_CurveGCComplete.TransactOpts, addr, weight)
}

// ChangeGaugeWeight is a paid mutator transaction binding the contract method 0xd4d2646e.
//
// Solidity: function change_gauge_weight(address addr, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) ChangeGaugeWeight(addr common.Address, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ChangeGaugeWeight(&_CurveGCComplete.TransactOpts, addr, weight)
}

// ChangeTypeWeight is a paid mutator transaction binding the contract method 0xdb1ca260.
//
// Solidity: function change_type_weight(int128 type_id, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) ChangeTypeWeight(opts *bind.TransactOpts, type_id *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "change_type_weight", type_id, weight)
}

// ChangeTypeWeight is a paid mutator transaction binding the contract method 0xdb1ca260.
//
// Solidity: function change_type_weight(int128 type_id, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteSession) ChangeTypeWeight(type_id *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ChangeTypeWeight(&_CurveGCComplete.TransactOpts, type_id, weight)
}

// ChangeTypeWeight is a paid mutator transaction binding the contract method 0xdb1ca260.
//
// Solidity: function change_type_weight(int128 type_id, uint256 weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) ChangeTypeWeight(type_id *big.Int, weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.ChangeTypeWeight(&_CurveGCComplete.TransactOpts, type_id, weight)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveGCComplete *CurveGCCompleteSession) Checkpoint() (*types.Transaction, error) {
	return _CurveGCComplete.Contract.Checkpoint(&_CurveGCComplete.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _CurveGCComplete.Contract.Checkpoint(&_CurveGCComplete.TransactOpts)
}

// CheckpointGauge is a paid mutator transaction binding the contract method 0x615e5237.
//
// Solidity: function checkpoint_gauge(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) CheckpointGauge(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "checkpoint_gauge", addr)
}

// CheckpointGauge is a paid mutator transaction binding the contract method 0x615e5237.
//
// Solidity: function checkpoint_gauge(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteSession) CheckpointGauge(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CheckpointGauge(&_CurveGCComplete.TransactOpts, addr)
}

// CheckpointGauge is a paid mutator transaction binding the contract method 0x615e5237.
//
// Solidity: function checkpoint_gauge(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) CheckpointGauge(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CheckpointGauge(&_CurveGCComplete.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CommitTransferOwnership(&_CurveGCComplete.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.CommitTransferOwnership(&_CurveGCComplete.TransactOpts, addr)
}

// GaugeRelativeWeightWrite is a paid mutator transaction binding the contract method 0x95cfcec3.
//
// Solidity: function gauge_relative_weight_write(address addr) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteTransactor) GaugeRelativeWeightWrite(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "gauge_relative_weight_write", addr)
}

// GaugeRelativeWeightWrite is a paid mutator transaction binding the contract method 0x95cfcec3.
//
// Solidity: function gauge_relative_weight_write(address addr) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeRelativeWeightWrite(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeightWrite(&_CurveGCComplete.TransactOpts, addr)
}

// GaugeRelativeWeightWrite is a paid mutator transaction binding the contract method 0x95cfcec3.
//
// Solidity: function gauge_relative_weight_write(address addr) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteTransactorSession) GaugeRelativeWeightWrite(addr common.Address) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeightWrite(&_CurveGCComplete.TransactOpts, addr)
}

// GaugeRelativeWeightWrite0 is a paid mutator transaction binding the contract method 0x6472eee1.
//
// Solidity: function gauge_relative_weight_write(address addr, uint256 time) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteTransactor) GaugeRelativeWeightWrite0(opts *bind.TransactOpts, addr common.Address, time *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "gauge_relative_weight_write0", addr, time)
}

// GaugeRelativeWeightWrite0 is a paid mutator transaction binding the contract method 0x6472eee1.
//
// Solidity: function gauge_relative_weight_write(address addr, uint256 time) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteSession) GaugeRelativeWeightWrite0(addr common.Address, time *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeightWrite0(&_CurveGCComplete.TransactOpts, addr, time)
}

// GaugeRelativeWeightWrite0 is a paid mutator transaction binding the contract method 0x6472eee1.
//
// Solidity: function gauge_relative_weight_write(address addr, uint256 time) returns(uint256)
func (_CurveGCComplete *CurveGCCompleteTransactorSession) GaugeRelativeWeightWrite0(addr common.Address, time *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.GaugeRelativeWeightWrite0(&_CurveGCComplete.TransactOpts, addr, time)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address _gauge_addr, uint256 _user_weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactor) VoteForGaugeWeights(opts *bind.TransactOpts, _gauge_addr common.Address, _user_weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.contract.Transact(opts, "vote_for_gauge_weights", _gauge_addr, _user_weight)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address _gauge_addr, uint256 _user_weight) returns()
func (_CurveGCComplete *CurveGCCompleteSession) VoteForGaugeWeights(_gauge_addr common.Address, _user_weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.VoteForGaugeWeights(&_CurveGCComplete.TransactOpts, _gauge_addr, _user_weight)
}

// VoteForGaugeWeights is a paid mutator transaction binding the contract method 0xd7136328.
//
// Solidity: function vote_for_gauge_weights(address _gauge_addr, uint256 _user_weight) returns()
func (_CurveGCComplete *CurveGCCompleteTransactorSession) VoteForGaugeWeights(_gauge_addr common.Address, _user_weight *big.Int) (*types.Transaction, error) {
	return _CurveGCComplete.Contract.VoteForGaugeWeights(&_CurveGCComplete.TransactOpts, _gauge_addr, _user_weight)
}

// CurveGCCompleteAddTypeIterator is returned from FilterAddType and is used to iterate over the raw logs and unpacked data for AddType events raised by the CurveGCComplete contract.
type CurveGCCompleteAddTypeIterator struct {
	Event *CurveGCCompleteAddType // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteAddTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteAddType)
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
		it.Event = new(CurveGCCompleteAddType)
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
func (it *CurveGCCompleteAddTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteAddTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteAddType represents a AddType event raised by the CurveGCComplete contract.
type CurveGCCompleteAddType struct {
	Name   string
	TypeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddType is a free log retrieval operation binding the contract event 0x6fbe76157c712f16b5a3c44ed48baa04e3450bc3fab0c020e848aca72bbccc84.
//
// Solidity: event AddType(string name, int128 type_id)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterAddType(opts *bind.FilterOpts) (*CurveGCCompleteAddTypeIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "AddType")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteAddTypeIterator{contract: _CurveGCComplete.contract, event: "AddType", logs: logs, sub: sub}, nil
}

// WatchAddType is a free log subscription operation binding the contract event 0x6fbe76157c712f16b5a3c44ed48baa04e3450bc3fab0c020e848aca72bbccc84.
//
// Solidity: event AddType(string name, int128 type_id)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchAddType(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteAddType) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "AddType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteAddType)
				if err := _CurveGCComplete.contract.UnpackLog(event, "AddType", log); err != nil {
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

// ParseAddType is a log parse operation binding the contract event 0x6fbe76157c712f16b5a3c44ed48baa04e3450bc3fab0c020e848aca72bbccc84.
//
// Solidity: event AddType(string name, int128 type_id)
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseAddType(log types.Log) (*CurveGCCompleteAddType, error) {
	event := new(CurveGCCompleteAddType)
	if err := _CurveGCComplete.contract.UnpackLog(event, "AddType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the CurveGCComplete contract.
type CurveGCCompleteApplyOwnershipIterator struct {
	Event *CurveGCCompleteApplyOwnership // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteApplyOwnership)
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
		it.Event = new(CurveGCCompleteApplyOwnership)
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
func (it *CurveGCCompleteApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteApplyOwnership represents a ApplyOwnership event raised by the CurveGCComplete contract.
type CurveGCCompleteApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*CurveGCCompleteApplyOwnershipIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteApplyOwnershipIterator{contract: _CurveGCComplete.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteApplyOwnership)
				if err := _CurveGCComplete.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseApplyOwnership(log types.Log) (*CurveGCCompleteApplyOwnership, error) {
	event := new(CurveGCCompleteApplyOwnership)
	if err := _CurveGCComplete.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the CurveGCComplete contract.
type CurveGCCompleteCommitOwnershipIterator struct {
	Event *CurveGCCompleteCommitOwnership // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteCommitOwnership)
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
		it.Event = new(CurveGCCompleteCommitOwnership)
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
func (it *CurveGCCompleteCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteCommitOwnership represents a CommitOwnership event raised by the CurveGCComplete contract.
type CurveGCCompleteCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*CurveGCCompleteCommitOwnershipIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteCommitOwnershipIterator{contract: _CurveGCComplete.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteCommitOwnership)
				if err := _CurveGCComplete.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseCommitOwnership(log types.Log) (*CurveGCCompleteCommitOwnership, error) {
	event := new(CurveGCCompleteCommitOwnership)
	if err := _CurveGCComplete.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteNewGaugeIterator is returned from FilterNewGauge and is used to iterate over the raw logs and unpacked data for NewGauge events raised by the CurveGCComplete contract.
type CurveGCCompleteNewGaugeIterator struct {
	Event *CurveGCCompleteNewGauge // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteNewGaugeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteNewGauge)
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
		it.Event = new(CurveGCCompleteNewGauge)
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
func (it *CurveGCCompleteNewGaugeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteNewGaugeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteNewGauge represents a NewGauge event raised by the CurveGCComplete contract.
type CurveGCCompleteNewGauge struct {
	Addr      common.Address
	GaugeType *big.Int
	Weight    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewGauge is a free log retrieval operation binding the contract event 0xfd55b3191f9c9dd92f4f134dd700e7d76f6a0c836a08687023d6d38f03ebd877.
//
// Solidity: event NewGauge(address addr, int128 gauge_type, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterNewGauge(opts *bind.FilterOpts) (*CurveGCCompleteNewGaugeIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "NewGauge")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteNewGaugeIterator{contract: _CurveGCComplete.contract, event: "NewGauge", logs: logs, sub: sub}, nil
}

// WatchNewGauge is a free log subscription operation binding the contract event 0xfd55b3191f9c9dd92f4f134dd700e7d76f6a0c836a08687023d6d38f03ebd877.
//
// Solidity: event NewGauge(address addr, int128 gauge_type, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchNewGauge(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteNewGauge) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "NewGauge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteNewGauge)
				if err := _CurveGCComplete.contract.UnpackLog(event, "NewGauge", log); err != nil {
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

// ParseNewGauge is a log parse operation binding the contract event 0xfd55b3191f9c9dd92f4f134dd700e7d76f6a0c836a08687023d6d38f03ebd877.
//
// Solidity: event NewGauge(address addr, int128 gauge_type, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseNewGauge(log types.Log) (*CurveGCCompleteNewGauge, error) {
	event := new(CurveGCCompleteNewGauge)
	if err := _CurveGCComplete.contract.UnpackLog(event, "NewGauge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteNewGaugeWeightIterator is returned from FilterNewGaugeWeight and is used to iterate over the raw logs and unpacked data for NewGaugeWeight events raised by the CurveGCComplete contract.
type CurveGCCompleteNewGaugeWeightIterator struct {
	Event *CurveGCCompleteNewGaugeWeight // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteNewGaugeWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteNewGaugeWeight)
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
		it.Event = new(CurveGCCompleteNewGaugeWeight)
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
func (it *CurveGCCompleteNewGaugeWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteNewGaugeWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteNewGaugeWeight represents a NewGaugeWeight event raised by the CurveGCComplete contract.
type CurveGCCompleteNewGaugeWeight struct {
	GaugeAddress common.Address
	Time         *big.Int
	Weight       *big.Int
	TotalWeight  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewGaugeWeight is a free log retrieval operation binding the contract event 0x54c0cf3647e6cdb2fc0a7876e60ba77563fceedf2e06c01c597f8dccb9e6bd72.
//
// Solidity: event NewGaugeWeight(address gauge_address, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterNewGaugeWeight(opts *bind.FilterOpts) (*CurveGCCompleteNewGaugeWeightIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "NewGaugeWeight")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteNewGaugeWeightIterator{contract: _CurveGCComplete.contract, event: "NewGaugeWeight", logs: logs, sub: sub}, nil
}

// WatchNewGaugeWeight is a free log subscription operation binding the contract event 0x54c0cf3647e6cdb2fc0a7876e60ba77563fceedf2e06c01c597f8dccb9e6bd72.
//
// Solidity: event NewGaugeWeight(address gauge_address, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchNewGaugeWeight(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteNewGaugeWeight) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "NewGaugeWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteNewGaugeWeight)
				if err := _CurveGCComplete.contract.UnpackLog(event, "NewGaugeWeight", log); err != nil {
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

// ParseNewGaugeWeight is a log parse operation binding the contract event 0x54c0cf3647e6cdb2fc0a7876e60ba77563fceedf2e06c01c597f8dccb9e6bd72.
//
// Solidity: event NewGaugeWeight(address gauge_address, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseNewGaugeWeight(log types.Log) (*CurveGCCompleteNewGaugeWeight, error) {
	event := new(CurveGCCompleteNewGaugeWeight)
	if err := _CurveGCComplete.contract.UnpackLog(event, "NewGaugeWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteNewTypeWeightIterator is returned from FilterNewTypeWeight and is used to iterate over the raw logs and unpacked data for NewTypeWeight events raised by the CurveGCComplete contract.
type CurveGCCompleteNewTypeWeightIterator struct {
	Event *CurveGCCompleteNewTypeWeight // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteNewTypeWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteNewTypeWeight)
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
		it.Event = new(CurveGCCompleteNewTypeWeight)
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
func (it *CurveGCCompleteNewTypeWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteNewTypeWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteNewTypeWeight represents a NewTypeWeight event raised by the CurveGCComplete contract.
type CurveGCCompleteNewTypeWeight struct {
	TypeId      *big.Int
	Time        *big.Int
	Weight      *big.Int
	TotalWeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewTypeWeight is a free log retrieval operation binding the contract event 0x00170bcdc909b6ac6e12d020fe8942256312cdcd555fb6d712899eba56d2f901.
//
// Solidity: event NewTypeWeight(int128 type_id, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterNewTypeWeight(opts *bind.FilterOpts) (*CurveGCCompleteNewTypeWeightIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "NewTypeWeight")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteNewTypeWeightIterator{contract: _CurveGCComplete.contract, event: "NewTypeWeight", logs: logs, sub: sub}, nil
}

// WatchNewTypeWeight is a free log subscription operation binding the contract event 0x00170bcdc909b6ac6e12d020fe8942256312cdcd555fb6d712899eba56d2f901.
//
// Solidity: event NewTypeWeight(int128 type_id, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchNewTypeWeight(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteNewTypeWeight) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "NewTypeWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteNewTypeWeight)
				if err := _CurveGCComplete.contract.UnpackLog(event, "NewTypeWeight", log); err != nil {
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

// ParseNewTypeWeight is a log parse operation binding the contract event 0x00170bcdc909b6ac6e12d020fe8942256312cdcd555fb6d712899eba56d2f901.
//
// Solidity: event NewTypeWeight(int128 type_id, uint256 time, uint256 weight, uint256 total_weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseNewTypeWeight(log types.Log) (*CurveGCCompleteNewTypeWeight, error) {
	event := new(CurveGCCompleteNewTypeWeight)
	if err := _CurveGCComplete.contract.UnpackLog(event, "NewTypeWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveGCCompleteVoteForGaugeIterator is returned from FilterVoteForGauge and is used to iterate over the raw logs and unpacked data for VoteForGauge events raised by the CurveGCComplete contract.
type CurveGCCompleteVoteForGaugeIterator struct {
	Event *CurveGCCompleteVoteForGauge // Event containing the contract specifics and raw log

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
func (it *CurveGCCompleteVoteForGaugeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCCompleteVoteForGauge)
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
		it.Event = new(CurveGCCompleteVoteForGauge)
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
func (it *CurveGCCompleteVoteForGaugeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCCompleteVoteForGaugeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCCompleteVoteForGauge represents a VoteForGauge event raised by the CurveGCComplete contract.
type CurveGCCompleteVoteForGauge struct {
	Time      *big.Int
	User      common.Address
	GaugeAddr common.Address
	Weight    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteForGauge is a free log retrieval operation binding the contract event 0x45ca9a4c8d0119eb329e580d28fe689e484e1be230da8037ade9547d2d25cc91.
//
// Solidity: event VoteForGauge(uint256 time, address user, address gauge_addr, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) FilterVoteForGauge(opts *bind.FilterOpts) (*CurveGCCompleteVoteForGaugeIterator, error) {

	logs, sub, err := _CurveGCComplete.contract.FilterLogs(opts, "VoteForGauge")
	if err != nil {
		return nil, err
	}
	return &CurveGCCompleteVoteForGaugeIterator{contract: _CurveGCComplete.contract, event: "VoteForGauge", logs: logs, sub: sub}, nil
}

// WatchVoteForGauge is a free log subscription operation binding the contract event 0x45ca9a4c8d0119eb329e580d28fe689e484e1be230da8037ade9547d2d25cc91.
//
// Solidity: event VoteForGauge(uint256 time, address user, address gauge_addr, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) WatchVoteForGauge(opts *bind.WatchOpts, sink chan<- *CurveGCCompleteVoteForGauge) (event.Subscription, error) {

	logs, sub, err := _CurveGCComplete.contract.WatchLogs(opts, "VoteForGauge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCCompleteVoteForGauge)
				if err := _CurveGCComplete.contract.UnpackLog(event, "VoteForGauge", log); err != nil {
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

// ParseVoteForGauge is a log parse operation binding the contract event 0x45ca9a4c8d0119eb329e580d28fe689e484e1be230da8037ade9547d2d25cc91.
//
// Solidity: event VoteForGauge(uint256 time, address user, address gauge_addr, uint256 weight)
func (_CurveGCComplete *CurveGCCompleteFilterer) ParseVoteForGauge(log types.Log) (*CurveGCCompleteVoteForGauge, error) {
	event := new(CurveGCCompleteVoteForGauge)
	if err := _CurveGCComplete.contract.UnpackLog(event, "VoteForGauge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
