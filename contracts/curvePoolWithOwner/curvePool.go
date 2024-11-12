// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvePoolWithOwner

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

// CurvePoolMetaData contains all meta data concerning the CurvePool contract.
var CurvePoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"sold_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":false},{\"type\":\"int128\",\"name\":\"bought_id\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"coin_amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"deadline\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewFee\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fee\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"admin_fee\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"old_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"new_A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"initial_time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"future_time\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"A\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"t\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address[3]\",\"name\":\"_coins\"},{\"type\":\"address\",\"name\":\"_pool_token\"},{\"type\":\"uint256\",\"name\":\"_A\"},{\"type\":\"uint256\",\"name\":\"_fee\"},{\"type\":\"uint256\",\"name\":\"_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5227},{\"name\":\"get_virtual_price\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1133537},{\"name\":\"calc_token_amount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"bool\",\"name\":\"deposit\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4508776},{\"name\":\"add_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"min_mint_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":6954858},{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2673791},{\"name\":\"get_dy_underlying\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2673474},{\"name\":\"exchange\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"},{\"type\":\"uint256\",\"name\":\"min_dy\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2818066},{\"name\":\"remove_liquidity\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"},{\"type\":\"uint256[3]\",\"name\":\"min_amounts\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":192846},{\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256[3]\",\"name\":\"amounts\"},{\"type\":\"uint256\",\"name\":\"max_burn_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":6951851},{\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1102},{\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_token_amount\"},{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"uint256\",\"name\":\"min_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":4025523},{\"name\":\"ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_future_A\"},{\"type\":\"uint256\",\"name\":\"_future_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":151919},{\"name\":\"stop_ramp_A\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148637},{\"name\":\"commit_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":110461},{\"name\":\"apply_new_fee\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":97242},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21895},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74572},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60710},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21985},{\"name\":\"admin_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3481},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21502},{\"name\":\"donate_admin_fees\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111389},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37998},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":22135},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2220},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2250},{\"name\":\"fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2201},{\"name\":\"owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2231},{\"name\":\"initial_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2261},{\"name\":\"future_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"initial_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"future_A_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin_actions_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"transfer_ownership_deadline\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411},{\"name\":\"future_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2441},{\"name\":\"future_admin_fee\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2471},{\"name\":\"future_owner\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2501}]",
}

// CurvePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolMetaData.ABI instead.
var CurvePoolABI = CurvePoolMetaData.ABI

// CurvePool is an auto generated Go binding around an Ethereum contract.
type CurvePool struct {
	CurvePoolCaller     // Read-only binding to the contract
	CurvePoolTransactor // Write-only binding to the contract
	CurvePoolFilterer   // Log filterer for contract events
}

// CurvePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolSession struct {
	Contract     *CurvePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolCallerSession struct {
	Contract *CurvePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurvePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolTransactorSession struct {
	Contract     *CurvePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurvePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolRaw struct {
	Contract *CurvePool // Generic contract binding to access the raw methods on
}

// CurvePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolCallerRaw struct {
	Contract *CurvePoolCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolTransactorRaw struct {
	Contract *CurvePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePool creates a new instance of CurvePool, bound to a specific deployed contract.
func NewCurvePool(address common.Address, backend bind.ContractBackend) (*CurvePool, error) {
	contract, err := bindCurvePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePool{CurvePoolCaller: CurvePoolCaller{contract: contract}, CurvePoolTransactor: CurvePoolTransactor{contract: contract}, CurvePoolFilterer: CurvePoolFilterer{contract: contract}}, nil
}

// NewCurvePoolCaller creates a new read-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolCaller, error) {
	contract, err := bindCurvePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCaller{contract: contract}, nil
}

// NewCurvePoolTransactor creates a new write-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolTransactor, error) {
	contract, err := bindCurvePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTransactor{contract: contract}, nil
}

// NewCurvePoolFilterer creates a new log filterer instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolFilterer, error) {
	contract, err := bindCurvePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFilterer{contract: contract}, nil
}

// bindCurvePool binds a generic wrapper to an already deployed contract.
func bindCurvePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.CurvePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolSession) A() (*big.Int, error) {
	return _CurvePool.Contract.A(&_CurvePool.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) A() (*big.Int, error) {
	return _CurvePool.Contract.A(&_CurvePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurvePool.Contract.AdminActionsDeadline(&_CurvePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurvePool.Contract.AdminActionsDeadline(&_CurvePool.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.AdminBalances(&_CurvePool.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.AdminBalances(&_CurvePool.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) AdminFee() (*big.Int, error) {
	return _CurvePool.Contract.AdminFee(&_CurvePool.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AdminFee() (*big.Int, error) {
	return _CurvePool.Contract.AdminFee(&_CurvePool.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenAmount(&_CurvePool.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenAmount(&_CurvePool.CallOpts, amounts, deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurvePool *CurvePoolCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "calc_withdraw_one_coin", _token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurvePool *CurvePoolSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcWithdrawOneCoin(&_CurvePool.CallOpts, _token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _token_amount, int128 i) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) CalcWithdrawOneCoin(_token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcWithdrawOneCoin(&_CurvePool.CallOpts, _token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) Fee() (*big.Int, error) {
	return _CurvePool.Contract.Fee(&_CurvePool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Fee() (*big.Int, error) {
	return _CurvePool.Contract.Fee(&_CurvePool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureA() (*big.Int, error) {
	return _CurvePool.Contract.FutureA(&_CurvePool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureA() (*big.Int, error) {
	return _CurvePool.Contract.FutureA(&_CurvePool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureATime() (*big.Int, error) {
	return _CurvePool.Contract.FutureATime(&_CurvePool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureATime() (*big.Int, error) {
	return _CurvePool.Contract.FutureATime(&_CurvePool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureAdminFee() (*big.Int, error) {
	return _CurvePool.Contract.FutureAdminFee(&_CurvePool.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurvePool.Contract.FutureAdminFee(&_CurvePool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureFee() (*big.Int, error) {
	return _CurvePool.Contract.FutureFee(&_CurvePool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureFee() (*big.Int, error) {
	return _CurvePool.Contract.FutureFee(&_CurvePool.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurvePool *CurvePoolCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurvePool *CurvePoolSession) FutureOwner() (common.Address, error) {
	return _CurvePool.Contract.FutureOwner(&_CurvePool.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurvePool *CurvePoolCallerSession) FutureOwner() (common.Address, error) {
	return _CurvePool.Contract.FutureOwner(&_CurvePool.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDyUnderlying(&_CurvePool.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDyUnderlying(&_CurvePool.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.GetVirtualPrice(&_CurvePool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.GetVirtualPrice(&_CurvePool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePool *CurvePoolCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePool *CurvePoolSession) InitialA() (*big.Int, error) {
	return _CurvePool.Contract.InitialA(&_CurvePool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) InitialA() (*big.Int, error) {
	return _CurvePool.Contract.InitialA(&_CurvePool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePool *CurvePoolCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePool *CurvePoolSession) InitialATime() (*big.Int, error) {
	return _CurvePool.Contract.InitialATime(&_CurvePool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) InitialATime() (*big.Int, error) {
	return _CurvePool.Contract.InitialATime(&_CurvePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurvePool *CurvePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurvePool *CurvePoolSession) Owner() (common.Address, error) {
	return _CurvePool.Contract.Owner(&_CurvePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurvePool *CurvePoolCallerSession) Owner() (common.Address, error) {
	return _CurvePool.Contract.Owner(&_CurvePool.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurvePool *CurvePoolSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurvePool.Contract.TransferOwnershipDeadline(&_CurvePool.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurvePool.Contract.TransferOwnershipDeadline(&_CurvePool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurvePool *CurvePoolTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurvePool *CurvePoolSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity(&_CurvePool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurvePool *CurvePoolTransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity(&_CurvePool.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePool *CurvePoolTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePool *CurvePoolSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyNewFee(&_CurvePool.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePool *CurvePoolTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyNewFee(&_CurvePool.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurvePool *CurvePoolTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurvePool *CurvePoolSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyTransferOwnership(&_CurvePool.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurvePool *CurvePoolTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyTransferOwnership(&_CurvePool.TransactOpts)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurvePool *CurvePoolTransactor) CommitNewFee(opts *bind.TransactOpts, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "commit_new_fee", new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurvePool *CurvePoolSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitNewFee(&_CurvePool.TransactOpts, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0x5b5a1467.
//
// Solidity: function commit_new_fee(uint256 new_fee, uint256 new_admin_fee) returns()
func (_CurvePool *CurvePoolTransactorSession) CommitNewFee(new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitNewFee(&_CurvePool.TransactOpts, new_fee, new_admin_fee)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurvePool *CurvePoolTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurvePool *CurvePoolSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitTransferOwnership(&_CurvePool.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurvePool *CurvePoolTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitTransferOwnership(&_CurvePool.TransactOpts, _owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurvePool *CurvePoolTransactor) DonateAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "donate_admin_fees")
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurvePool *CurvePoolSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.DonateAdminFees(&_CurvePool.TransactOpts)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0x524c3901.
//
// Solidity: function donate_admin_fees() returns()
func (_CurvePool *CurvePoolTransactorSession) DonateAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.DonateAdminFees(&_CurvePool.TransactOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurvePool *CurvePoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurvePool *CurvePoolSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurvePool *CurvePoolTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurvePool *CurvePoolTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurvePool *CurvePoolSession) KillMe() (*types.Transaction, error) {
	return _CurvePool.Contract.KillMe(&_CurvePool.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurvePool *CurvePoolTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurvePool.Contract.KillMe(&_CurvePool.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePool *CurvePoolTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePool *CurvePoolSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RampA(&_CurvePool.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePool *CurvePoolTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RampA(&_CurvePool.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurvePool *CurvePoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurvePool *CurvePoolSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity(&_CurvePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity(&_CurvePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_CurvePool *CurvePoolTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity_imbalance", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_CurvePool *CurvePoolSession) RemoveLiquidityImbalance(amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityImbalance(&_CurvePool.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 max_burn_amount) returns()
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidityImbalance(amounts [3]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityImbalance(&_CurvePool.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurvePool *CurvePoolTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurvePool *CurvePoolSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin(&_CurvePool.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin(&_CurvePool.TransactOpts, _token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertNewParameters(&_CurvePool.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertNewParameters(&_CurvePool.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurvePool *CurvePoolTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurvePool *CurvePoolSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertTransferOwnership(&_CurvePool.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurvePool *CurvePoolTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertTransferOwnership(&_CurvePool.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePool *CurvePoolTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePool *CurvePoolSession) StopRampA() (*types.Transaction, error) {
	return _CurvePool.Contract.StopRampA(&_CurvePool.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePool *CurvePoolTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurvePool.Contract.StopRampA(&_CurvePool.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurvePool *CurvePoolTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurvePool *CurvePoolSession) UnkillMe() (*types.Transaction, error) {
	return _CurvePool.Contract.UnkillMe(&_CurvePool.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurvePool *CurvePoolTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurvePool.Contract.UnkillMe(&_CurvePool.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePool *CurvePoolTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePool *CurvePoolSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.WithdrawAdminFees(&_CurvePool.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePool *CurvePoolTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.WithdrawAdminFees(&_CurvePool.TransactOpts)
}

// CurvePoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurvePool contract.
type CurvePoolAddLiquidityIterator struct {
	Event *CurvePoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolAddLiquidity)
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
		it.Event = new(CurvePoolAddLiquidity)
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
func (it *CurvePoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolAddLiquidity represents a AddLiquidity event raised by the CurvePool contract.
type CurvePoolAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolAddLiquidityIterator{contract: _CurvePool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolAddLiquidity)
				if err := _CurvePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) ParseAddLiquidity(log types.Log) (*CurvePoolAddLiquidity, error) {
	event := new(CurvePoolAddLiquidity)
	if err := _CurvePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the CurvePool contract.
type CurvePoolCommitNewAdminIterator struct {
	Event *CurvePoolCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurvePoolCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolCommitNewAdmin)
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
		it.Event = new(CurvePoolCommitNewAdmin)
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
func (it *CurvePoolCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolCommitNewAdmin represents a CommitNewAdmin event raised by the CurvePool contract.
type CurvePoolCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurvePool *CurvePoolFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*CurvePoolCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCommitNewAdminIterator{contract: _CurvePool.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurvePool *CurvePoolFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *CurvePoolCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolCommitNewAdmin)
				if err := _CurvePool.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurvePool *CurvePoolFilterer) ParseCommitNewAdmin(log types.Log) (*CurvePoolCommitNewAdmin, error) {
	event := new(CurvePoolCommitNewAdmin)
	if err := _CurvePool.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the CurvePool contract.
type CurvePoolCommitNewFeeIterator struct {
	Event *CurvePoolCommitNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolCommitNewFee)
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
		it.Event = new(CurvePoolCommitNewFee)
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
func (it *CurvePoolCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolCommitNewFee represents a CommitNewFee event raised by the CurvePool contract.
type CurvePoolCommitNewFee struct {
	Deadline *big.Int
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) FilterCommitNewFee(opts *bind.FilterOpts, deadline []*big.Int) (*CurvePoolCommitNewFeeIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCommitNewFeeIterator{contract: _CurvePool.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolCommitNewFee, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "CommitNewFee", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolCommitNewFee)
				if err := _CurvePool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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

// ParseCommitNewFee is a log parse operation binding the contract event 0x351fc5da2fbf480f2225debf3664a4bc90fa9923743aad58b4603f648e931fe0.
//
// Solidity: event CommitNewFee(uint256 indexed deadline, uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) ParseCommitNewFee(log types.Log) (*CurvePoolCommitNewFee, error) {
	event := new(CurvePoolCommitNewFee)
	if err := _CurvePool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CurvePool contract.
type CurvePoolNewAdminIterator struct {
	Event *CurvePoolNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurvePoolNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolNewAdmin)
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
		it.Event = new(CurvePoolNewAdmin)
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
func (it *CurvePoolNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolNewAdmin represents a NewAdmin event raised by the CurvePool contract.
type CurvePoolNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurvePool *CurvePoolFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*CurvePoolNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolNewAdminIterator{contract: _CurvePool.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurvePool *CurvePoolFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CurvePoolNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolNewAdmin)
				if err := _CurvePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurvePool *CurvePoolFilterer) ParseNewAdmin(log types.Log) (*CurvePoolNewAdmin, error) {
	event := new(CurvePoolNewAdmin)
	if err := _CurvePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolNewFeeIterator is returned from FilterNewFee and is used to iterate over the raw logs and unpacked data for NewFee events raised by the CurvePool contract.
type CurvePoolNewFeeIterator struct {
	Event *CurvePoolNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolNewFee)
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
		it.Event = new(CurvePoolNewFee)
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
func (it *CurvePoolNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolNewFee represents a NewFee event raised by the CurvePool contract.
type CurvePoolNewFee struct {
	Fee      *big.Int
	AdminFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewFee is a free log retrieval operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) FilterNewFee(opts *bind.FilterOpts) (*CurvePoolNewFeeIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return &CurvePoolNewFeeIterator{contract: _CurvePool.contract, event: "NewFee", logs: logs, sub: sub}, nil
}

// WatchNewFee is a free log subscription operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) WatchNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolNewFee) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "NewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolNewFee)
				if err := _CurvePool.contract.UnpackLog(event, "NewFee", log); err != nil {
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

// ParseNewFee is a log parse operation binding the contract event 0xbe12859b636aed607d5230b2cc2711f68d70e51060e6cca1f575ef5d2fcc95d1.
//
// Solidity: event NewFee(uint256 fee, uint256 admin_fee)
func (_CurvePool *CurvePoolFilterer) ParseNewFee(log types.Log) (*CurvePoolNewFee, error) {
	event := new(CurvePoolNewFee)
	if err := _CurvePool.contract.UnpackLog(event, "NewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the CurvePool contract.
type CurvePoolRampAIterator struct {
	Event *CurvePoolRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRampA)
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
		it.Event = new(CurvePoolRampA)
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
func (it *CurvePoolRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRampA represents a RampA event raised by the CurvePool contract.
type CurvePoolRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) FilterRampA(opts *bind.FilterOpts) (*CurvePoolRampAIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolRampAIterator{contract: _CurvePool.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRampA)
				if err := _CurvePool.contract.UnpackLog(event, "RampA", log); err != nil {
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

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) ParseRampA(log types.Log) (*CurvePoolRampA, error) {
	event := new(CurvePoolRampA)
	if err := _CurvePool.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurvePool contract.
type CurvePoolRemoveLiquidityIterator struct {
	Event *CurvePoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoveLiquidity)
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
		it.Event = new(CurvePoolRemoveLiquidity)
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
func (it *CurvePoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoveLiquidity represents a RemoveLiquidity event raised by the CurvePool contract.
type CurvePoolRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemoveLiquidityIterator{contract: _CurvePool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoveLiquidity)
				if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) ParseRemoveLiquidity(log types.Log) (*CurvePoolRemoveLiquidity, error) {
	event := new(CurvePoolRemoveLiquidity)
	if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the CurvePool contract.
type CurvePoolRemoveLiquidityImbalanceIterator struct {
	Event *CurvePoolRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoveLiquidityImbalance)
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
		it.Event = new(CurvePoolRemoveLiquidityImbalance)
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
func (it *CurvePoolRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the CurvePool contract.
type CurvePoolRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemoveLiquidityImbalanceIterator{contract: _CurvePool.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoveLiquidityImbalance)
				if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*CurvePoolRemoveLiquidityImbalance, error) {
	event := new(CurvePoolRemoveLiquidityImbalance)
	if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurvePool contract.
type CurvePoolRemoveLiquidityOneIterator struct {
	Event *CurvePoolRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoveLiquidityOne)
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
		it.Event = new(CurvePoolRemoveLiquidityOne)
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
func (it *CurvePoolRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurvePool contract.
type CurvePoolRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_CurvePool *CurvePoolFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemoveLiquidityOneIterator{contract: _CurvePool.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_CurvePool *CurvePoolFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoveLiquidityOne)
				if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x9e96dd3b997a2a257eec4df9bb6eaf626e206df5f543bd963682d143300be310.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount)
func (_CurvePool *CurvePoolFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurvePoolRemoveLiquidityOne, error) {
	event := new(CurvePoolRemoveLiquidityOne)
	if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurvePool contract.
type CurvePoolStopRampAIterator struct {
	Event *CurvePoolStopRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolStopRampA)
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
		it.Event = new(CurvePoolStopRampA)
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
func (it *CurvePoolStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolStopRampA represents a StopRampA event raised by the CurvePool contract.
type CurvePoolStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePool *CurvePoolFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvePoolStopRampAIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolStopRampAIterator{contract: _CurvePool.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePool *CurvePoolFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolStopRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolStopRampA)
				if err := _CurvePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePool *CurvePoolFilterer) ParseStopRampA(log types.Log) (*CurvePoolStopRampA, error) {
	event := new(CurvePoolStopRampA)
	if err := _CurvePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurvePool contract.
type CurvePoolTokenExchangeIterator struct {
	Event *CurvePoolTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvePoolTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolTokenExchange)
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
		it.Event = new(CurvePoolTokenExchange)
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
func (it *CurvePoolTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolTokenExchange represents a TokenExchange event raised by the CurvePool contract.
type CurvePoolTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurvePool *CurvePoolFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvePoolTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTokenExchangeIterator{contract: _CurvePool.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurvePool *CurvePoolFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvePoolTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolTokenExchange)
				if err := _CurvePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurvePool *CurvePoolFilterer) ParseTokenExchange(log types.Log) (*CurvePoolTokenExchange, error) {
	event := new(CurvePoolTokenExchange)
	if err := _CurvePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
