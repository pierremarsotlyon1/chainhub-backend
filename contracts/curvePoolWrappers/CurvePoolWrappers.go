// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvePoolWrappers

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

// CurvePoolWrappersMetaData contains all meta data concerning the CurvePoolWrappers contract.
var CurvePoolWrappersMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"name\":\"new_fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_fee\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_action_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurvePoolWrappersABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolWrappersMetaData.ABI instead.
var CurvePoolWrappersABI = CurvePoolWrappersMetaData.ABI

// CurvePoolWrappers is an auto generated Go binding around an Ethereum contract.
type CurvePoolWrappers struct {
	CurvePoolWrappersCaller     // Read-only binding to the contract
	CurvePoolWrappersTransactor // Write-only binding to the contract
	CurvePoolWrappersFilterer   // Log filterer for contract events
}

// CurvePoolWrappersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolWrappersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolWrappersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolWrappersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolWrappersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolWrappersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolWrappersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolWrappersSession struct {
	Contract     *CurvePoolWrappers // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurvePoolWrappersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolWrappersCallerSession struct {
	Contract *CurvePoolWrappersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CurvePoolWrappersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolWrappersTransactorSession struct {
	Contract     *CurvePoolWrappersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CurvePoolWrappersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolWrappersRaw struct {
	Contract *CurvePoolWrappers // Generic contract binding to access the raw methods on
}

// CurvePoolWrappersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolWrappersCallerRaw struct {
	Contract *CurvePoolWrappersCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolWrappersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolWrappersTransactorRaw struct {
	Contract *CurvePoolWrappersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolWrappers creates a new instance of CurvePoolWrappers, bound to a specific deployed contract.
func NewCurvePoolWrappers(address common.Address, backend bind.ContractBackend) (*CurvePoolWrappers, error) {
	contract, err := bindCurvePoolWrappers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappers{CurvePoolWrappersCaller: CurvePoolWrappersCaller{contract: contract}, CurvePoolWrappersTransactor: CurvePoolWrappersTransactor{contract: contract}, CurvePoolWrappersFilterer: CurvePoolWrappersFilterer{contract: contract}}, nil
}

// NewCurvePoolWrappersCaller creates a new read-only instance of CurvePoolWrappers, bound to a specific deployed contract.
func NewCurvePoolWrappersCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolWrappersCaller, error) {
	contract, err := bindCurvePoolWrappers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersCaller{contract: contract}, nil
}

// NewCurvePoolWrappersTransactor creates a new write-only instance of CurvePoolWrappers, bound to a specific deployed contract.
func NewCurvePoolWrappersTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolWrappersTransactor, error) {
	contract, err := bindCurvePoolWrappers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersTransactor{contract: contract}, nil
}

// NewCurvePoolWrappersFilterer creates a new log filterer instance of CurvePoolWrappers, bound to a specific deployed contract.
func NewCurvePoolWrappersFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolWrappersFilterer, error) {
	contract, err := bindCurvePoolWrappers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersFilterer{contract: contract}, nil
}

// bindCurvePoolWrappers binds a generic wrapper to an already deployed contract.
func bindCurvePoolWrappers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolWrappersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolWrappers *CurvePoolWrappersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolWrappers.Contract.CurvePoolWrappersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolWrappers *CurvePoolWrappersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.CurvePoolWrappersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolWrappers *CurvePoolWrappersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.CurvePoolWrappersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolWrappers *CurvePoolWrappersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolWrappers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolWrappers *CurvePoolWrappersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolWrappers *CurvePoolWrappersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) A() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.A(&_CurvePoolWrappers.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) A() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.A(&_CurvePoolWrappers.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) APrecise() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.APrecise(&_CurvePoolWrappers.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) APrecise() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.APrecise(&_CurvePoolWrappers.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolWrappers *CurvePoolWrappersSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolWrappers.Contract.DOMAINSEPARATOR(&_CurvePoolWrappers.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolWrappers.Contract.DOMAINSEPARATOR(&_CurvePoolWrappers.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) AdminActionDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "admin_action_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) AdminActionDeadline() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminActionDeadline(&_CurvePoolWrappers.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) AdminActionDeadline() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminActionDeadline(&_CurvePoolWrappers.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminBalances(&_CurvePoolWrappers.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminBalances(&_CurvePoolWrappers.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) AdminFee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminFee(&_CurvePoolWrappers.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) AdminFee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.AdminFee(&_CurvePoolWrappers.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Allowance(&_CurvePoolWrappers.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Allowance(&_CurvePoolWrappers.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.BalanceOf(&_CurvePoolWrappers.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.BalanceOf(&_CurvePoolWrappers.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Balances(&_CurvePoolWrappers.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Balances(&_CurvePoolWrappers.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.CalcTokenAmount(&_CurvePoolWrappers.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.CalcTokenAmount(&_CurvePoolWrappers.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.CalcWithdrawOneCoin(&_CurvePoolWrappers.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.CalcWithdrawOneCoin(&_CurvePoolWrappers.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolWrappers.Contract.Coins(&_CurvePoolWrappers.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolWrappers.Contract.Coins(&_CurvePoolWrappers.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Decimals() (uint8, error) {
	return _CurvePoolWrappers.Contract.Decimals(&_CurvePoolWrappers.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Decimals() (uint8, error) {
	return _CurvePoolWrappers.Contract.Decimals(&_CurvePoolWrappers.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) EmaPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "ema_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) EmaPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.EmaPrice(&_CurvePoolWrappers.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) EmaPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.EmaPrice(&_CurvePoolWrappers.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Fee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Fee(&_CurvePoolWrappers.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Fee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Fee(&_CurvePoolWrappers.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) FutureA() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureA(&_CurvePoolWrappers.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) FutureA() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureA(&_CurvePoolWrappers.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) FutureATime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureATime(&_CurvePoolWrappers.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) FutureATime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureATime(&_CurvePoolWrappers.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) FutureFee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureFee(&_CurvePoolWrappers.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) FutureFee() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.FutureFee(&_CurvePoolWrappers.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersCaller) GetBalances(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersSession) GetBalances() ([2]*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetBalances(&_CurvePoolWrappers.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) GetBalances() ([2]*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetBalances(&_CurvePoolWrappers.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetDy(&_CurvePoolWrappers.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetDy(&_CurvePoolWrappers.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) GetP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "get_p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) GetP() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetP(&_CurvePoolWrappers.CallOpts)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) GetP() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetP(&_CurvePoolWrappers.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetVirtualPrice(&_CurvePoolWrappers.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.GetVirtualPrice(&_CurvePoolWrappers.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) InitialA() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.InitialA(&_CurvePoolWrappers.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) InitialA() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.InitialA(&_CurvePoolWrappers.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) InitialATime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.InitialATime(&_CurvePoolWrappers.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) InitialATime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.InitialATime(&_CurvePoolWrappers.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "last_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) LastPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.LastPrice(&_CurvePoolWrappers.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) LastPrice() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.LastPrice(&_CurvePoolWrappers.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) MaExpTime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.MaExpTime(&_CurvePoolWrappers.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) MaExpTime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.MaExpTime(&_CurvePoolWrappers.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) MaLastTime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.MaLastTime(&_CurvePoolWrappers.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) MaLastTime() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.MaLastTime(&_CurvePoolWrappers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Name() (string, error) {
	return _CurvePoolWrappers.Contract.Name(&_CurvePoolWrappers.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Name() (string, error) {
	return _CurvePoolWrappers.Contract.Name(&_CurvePoolWrappers.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Nonces(&_CurvePoolWrappers.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolWrappers.Contract.Nonces(&_CurvePoolWrappers.CallOpts, arg0)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) PriceOracle() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.PriceOracle(&_CurvePoolWrappers.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) PriceOracle() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.PriceOracle(&_CurvePoolWrappers.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Symbol() (string, error) {
	return _CurvePoolWrappers.Contract.Symbol(&_CurvePoolWrappers.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Symbol() (string, error) {
	return _CurvePoolWrappers.Contract.Symbol(&_CurvePoolWrappers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) TotalSupply() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.TotalSupply(&_CurvePoolWrappers.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) TotalSupply() (*big.Int, error) {
	return _CurvePoolWrappers.Contract.TotalSupply(&_CurvePoolWrappers.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolWrappers.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Version() (string, error) {
	return _CurvePoolWrappers.Contract.Version(&_CurvePoolWrappers.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolWrappers *CurvePoolWrappersCallerSession) Version() (string, error) {
	return _CurvePoolWrappers.Contract.Version(&_CurvePoolWrappers.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.AddLiquidity(&_CurvePoolWrappers.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.AddLiquidity(&_CurvePoolWrappers.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.AddLiquidity0(&_CurvePoolWrappers.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.AddLiquidity0(&_CurvePoolWrappers.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.ApplyNewFee(&_CurvePoolWrappers.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.ApplyNewFee(&_CurvePoolWrappers.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Approve(&_CurvePoolWrappers.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Approve(&_CurvePoolWrappers.TransactOpts, _spender, _value)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) CommitNewFee(opts *bind.TransactOpts, _new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "commit_new_fee", _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.CommitNewFee(&_CurvePoolWrappers.TransactOpts, _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.CommitNewFee(&_CurvePoolWrappers.TransactOpts, _new_fee)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Exchange(&_CurvePoolWrappers.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Exchange(&_CurvePoolWrappers.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Exchange0(&_CurvePoolWrappers.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Exchange0(&_CurvePoolWrappers.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "initialize", _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Initialize(&_CurvePoolWrappers.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Initialize(&_CurvePoolWrappers.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Permit(&_CurvePoolWrappers.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Permit(&_CurvePoolWrappers.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RampA(&_CurvePoolWrappers.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RampA(&_CurvePoolWrappers.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidity(&_CurvePoolWrappers.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidity(&_CurvePoolWrappers.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidity0(&_CurvePoolWrappers.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidity0(&_CurvePoolWrappers.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityImbalance(&_CurvePoolWrappers.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityImbalance(&_CurvePoolWrappers.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityImbalance0(&_CurvePoolWrappers.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityImbalance0(&_CurvePoolWrappers.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityOneCoin(&_CurvePoolWrappers.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityOneCoin(&_CurvePoolWrappers.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityOneCoin0(&_CurvePoolWrappers.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.RemoveLiquidityOneCoin0(&_CurvePoolWrappers.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.SetMaExpTime(&_CurvePoolWrappers.TransactOpts, _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.SetMaExpTime(&_CurvePoolWrappers.TransactOpts, _ma_exp_time)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) StopRampA() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.StopRampA(&_CurvePoolWrappers.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.StopRampA(&_CurvePoolWrappers.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Transfer(&_CurvePoolWrappers.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.Transfer(&_CurvePoolWrappers.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.TransferFrom(&_CurvePoolWrappers.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.TransferFrom(&_CurvePoolWrappers.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolWrappers.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolWrappers *CurvePoolWrappersSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.WithdrawAdminFees(&_CurvePoolWrappers.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolWrappers *CurvePoolWrappersTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePoolWrappers.Contract.WithdrawAdminFees(&_CurvePoolWrappers.TransactOpts)
}

// CurvePoolWrappersAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersAddLiquidityIterator struct {
	Event *CurvePoolWrappersAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersAddLiquidity)
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
		it.Event = new(CurvePoolWrappersAddLiquidity)
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
func (it *CurvePoolWrappersAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersAddLiquidity represents a AddLiquidity event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolWrappersAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersAddLiquidityIterator{contract: _CurvePoolWrappers.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersAddLiquidity)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseAddLiquidity(log types.Log) (*CurvePoolWrappersAddLiquidity, error) {
	event := new(CurvePoolWrappersAddLiquidity)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersApplyNewFeeIterator struct {
	Event *CurvePoolWrappersApplyNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersApplyNewFee)
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
		it.Event = new(CurvePoolWrappersApplyNewFee)
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
func (it *CurvePoolWrappersApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersApplyNewFee represents a ApplyNewFee event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersApplyNewFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*CurvePoolWrappersApplyNewFeeIterator, error) {

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersApplyNewFeeIterator{contract: _CurvePoolWrappers.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersApplyNewFee)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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

// ParseApplyNewFee is a log parse operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseApplyNewFee(log types.Log) (*CurvePoolWrappersApplyNewFee, error) {
	event := new(CurvePoolWrappersApplyNewFee)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersApprovalIterator struct {
	Event *CurvePoolWrappersApproval // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersApproval)
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
		it.Event = new(CurvePoolWrappersApproval)
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
func (it *CurvePoolWrappersApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersApproval represents a Approval event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CurvePoolWrappersApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersApprovalIterator{contract: _CurvePoolWrappers.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersApproval)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseApproval(log types.Log) (*CurvePoolWrappersApproval, error) {
	event := new(CurvePoolWrappersApproval)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersCommitNewFeeIterator struct {
	Event *CurvePoolWrappersCommitNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersCommitNewFee)
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
		it.Event = new(CurvePoolWrappersCommitNewFee)
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
func (it *CurvePoolWrappersCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersCommitNewFee represents a CommitNewFee event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersCommitNewFee struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterCommitNewFee(opts *bind.FilterOpts) (*CurvePoolWrappersCommitNewFeeIterator, error) {

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersCommitNewFeeIterator{contract: _CurvePoolWrappers.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersCommitNewFee) (event.Subscription, error) {

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersCommitNewFee)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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

// ParseCommitNewFee is a log parse operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseCommitNewFee(log types.Log) (*CurvePoolWrappersCommitNewFee, error) {
	event := new(CurvePoolWrappersCommitNewFee)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRampAIterator struct {
	Event *CurvePoolWrappersRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersRampA)
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
		it.Event = new(CurvePoolWrappersRampA)
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
func (it *CurvePoolWrappersRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersRampA represents a RampA event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterRampA(opts *bind.FilterOpts) (*CurvePoolWrappersRampAIterator, error) {

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersRampAIterator{contract: _CurvePoolWrappers.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersRampA)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseRampA(log types.Log) (*CurvePoolWrappersRampA, error) {
	event := new(CurvePoolWrappersRampA)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidityIterator struct {
	Event *CurvePoolWrappersRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersRemoveLiquidity)
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
		it.Event = new(CurvePoolWrappersRemoveLiquidity)
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
func (it *CurvePoolWrappersRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersRemoveLiquidity represents a RemoveLiquidity event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolWrappersRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersRemoveLiquidityIterator{contract: _CurvePoolWrappers.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersRemoveLiquidity)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseRemoveLiquidity(log types.Log) (*CurvePoolWrappersRemoveLiquidity, error) {
	event := new(CurvePoolWrappersRemoveLiquidity)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidityImbalanceIterator struct {
	Event *CurvePoolWrappersRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersRemoveLiquidityImbalance)
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
		it.Event = new(CurvePoolWrappersRemoveLiquidityImbalance)
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
func (it *CurvePoolWrappersRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolWrappersRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersRemoveLiquidityImbalanceIterator{contract: _CurvePoolWrappers.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersRemoveLiquidityImbalance)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*CurvePoolWrappersRemoveLiquidityImbalance, error) {
	event := new(CurvePoolWrappersRemoveLiquidityImbalance)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidityOneIterator struct {
	Event *CurvePoolWrappersRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersRemoveLiquidityOne)
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
		it.Event = new(CurvePoolWrappersRemoveLiquidityOne)
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
func (it *CurvePoolWrappersRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	TokenSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolWrappersRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersRemoveLiquidityOneIterator{contract: _CurvePoolWrappers.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersRemoveLiquidityOne)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurvePoolWrappersRemoveLiquidityOne, error) {
	event := new(CurvePoolWrappersRemoveLiquidityOne)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersStopRampAIterator struct {
	Event *CurvePoolWrappersStopRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersStopRampA)
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
		it.Event = new(CurvePoolWrappersStopRampA)
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
func (it *CurvePoolWrappersStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersStopRampA represents a StopRampA event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvePoolWrappersStopRampAIterator, error) {

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersStopRampAIterator{contract: _CurvePoolWrappers.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersStopRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersStopRampA)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseStopRampA(log types.Log) (*CurvePoolWrappersStopRampA, error) {
	event := new(CurvePoolWrappersStopRampA)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersTokenExchangeIterator struct {
	Event *CurvePoolWrappersTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersTokenExchange)
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
		it.Event = new(CurvePoolWrappersTokenExchange)
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
func (it *CurvePoolWrappersTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersTokenExchange represents a TokenExchange event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersTokenExchange struct {
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
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvePoolWrappersTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersTokenExchangeIterator{contract: _CurvePoolWrappers.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersTokenExchange)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseTokenExchange(log types.Log) (*CurvePoolWrappersTokenExchange, error) {
	event := new(CurvePoolWrappersTokenExchange)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolWrappersTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurvePoolWrappers contract.
type CurvePoolWrappersTransferIterator struct {
	Event *CurvePoolWrappersTransfer // Event containing the contract specifics and raw log

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
func (it *CurvePoolWrappersTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolWrappersTransfer)
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
		it.Event = new(CurvePoolWrappersTransfer)
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
func (it *CurvePoolWrappersTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolWrappersTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolWrappersTransfer represents a Transfer event raised by the CurvePoolWrappers contract.
type CurvePoolWrappersTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*CurvePoolWrappersTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolWrappersTransferIterator{contract: _CurvePoolWrappers.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurvePoolWrappersTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePoolWrappers.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolWrappersTransfer)
				if err := _CurvePoolWrappers.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePoolWrappers *CurvePoolWrappersFilterer) ParseTransfer(log types.Log) (*CurvePoolWrappersTransfer, error) {
	event := new(CurvePoolWrappersTransfer)
	if err := _CurvePoolWrappers.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
