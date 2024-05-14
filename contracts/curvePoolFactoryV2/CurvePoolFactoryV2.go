// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvePoolFactoryV2

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

// CurvePoolFactoryV2MetaData contains all meta data concerning the CurvePoolFactoryV2 contract.
var CurvePoolFactoryV2MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"name\":\"new_fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_fee\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_action_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CurvePoolFactoryV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolFactoryV2MetaData.ABI instead.
var CurvePoolFactoryV2ABI = CurvePoolFactoryV2MetaData.ABI

// CurvePoolFactoryV2 is an auto generated Go binding around an Ethereum contract.
type CurvePoolFactoryV2 struct {
	CurvePoolFactoryV2Caller     // Read-only binding to the contract
	CurvePoolFactoryV2Transactor // Write-only binding to the contract
	CurvePoolFactoryV2Filterer   // Log filterer for contract events
}

// CurvePoolFactoryV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolFactoryV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactoryV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolFactoryV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactoryV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolFactoryV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFactoryV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolFactoryV2Session struct {
	Contract     *CurvePoolFactoryV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CurvePoolFactoryV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolFactoryV2CallerSession struct {
	Contract *CurvePoolFactoryV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CurvePoolFactoryV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolFactoryV2TransactorSession struct {
	Contract     *CurvePoolFactoryV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CurvePoolFactoryV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolFactoryV2Raw struct {
	Contract *CurvePoolFactoryV2 // Generic contract binding to access the raw methods on
}

// CurvePoolFactoryV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolFactoryV2CallerRaw struct {
	Contract *CurvePoolFactoryV2Caller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolFactoryV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolFactoryV2TransactorRaw struct {
	Contract *CurvePoolFactoryV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolFactoryV2 creates a new instance of CurvePoolFactoryV2, bound to a specific deployed contract.
func NewCurvePoolFactoryV2(address common.Address, backend bind.ContractBackend) (*CurvePoolFactoryV2, error) {
	contract, err := bindCurvePoolFactoryV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2{CurvePoolFactoryV2Caller: CurvePoolFactoryV2Caller{contract: contract}, CurvePoolFactoryV2Transactor: CurvePoolFactoryV2Transactor{contract: contract}, CurvePoolFactoryV2Filterer: CurvePoolFactoryV2Filterer{contract: contract}}, nil
}

// NewCurvePoolFactoryV2Caller creates a new read-only instance of CurvePoolFactoryV2, bound to a specific deployed contract.
func NewCurvePoolFactoryV2Caller(address common.Address, caller bind.ContractCaller) (*CurvePoolFactoryV2Caller, error) {
	contract, err := bindCurvePoolFactoryV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2Caller{contract: contract}, nil
}

// NewCurvePoolFactoryV2Transactor creates a new write-only instance of CurvePoolFactoryV2, bound to a specific deployed contract.
func NewCurvePoolFactoryV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolFactoryV2Transactor, error) {
	contract, err := bindCurvePoolFactoryV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2Transactor{contract: contract}, nil
}

// NewCurvePoolFactoryV2Filterer creates a new log filterer instance of CurvePoolFactoryV2, bound to a specific deployed contract.
func NewCurvePoolFactoryV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolFactoryV2Filterer, error) {
	contract, err := bindCurvePoolFactoryV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2Filterer{contract: contract}, nil
}

// bindCurvePoolFactoryV2 binds a generic wrapper to an already deployed contract.
func bindCurvePoolFactoryV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolFactoryV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolFactoryV2.Contract.CurvePoolFactoryV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.CurvePoolFactoryV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.CurvePoolFactoryV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolFactoryV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) A() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.A(&_CurvePoolFactoryV2.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) A() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.A(&_CurvePoolFactoryV2.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) APrecise() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.APrecise(&_CurvePoolFactoryV2.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) APrecise() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.APrecise(&_CurvePoolFactoryV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolFactoryV2.Contract.DOMAINSEPARATOR(&_CurvePoolFactoryV2.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePoolFactoryV2.Contract.DOMAINSEPARATOR(&_CurvePoolFactoryV2.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) AdminActionDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "admin_action_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) AdminActionDeadline() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminActionDeadline(&_CurvePoolFactoryV2.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) AdminActionDeadline() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminActionDeadline(&_CurvePoolFactoryV2.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminBalances(&_CurvePoolFactoryV2.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminBalances(&_CurvePoolFactoryV2.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) AdminFee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminFee(&_CurvePoolFactoryV2.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) AdminFee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.AdminFee(&_CurvePoolFactoryV2.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Allowance(&_CurvePoolFactoryV2.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Allowance(&_CurvePoolFactoryV2.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.BalanceOf(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.BalanceOf(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Balances(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Balances(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) CalcTokenAmount(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.CalcTokenAmount(&_CurvePoolFactoryV2.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.CalcTokenAmount(&_CurvePoolFactoryV2.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.CalcWithdrawOneCoin(&_CurvePoolFactoryV2.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.CalcWithdrawOneCoin(&_CurvePoolFactoryV2.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolFactoryV2.Contract.Coins(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolFactoryV2.Contract.Coins(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Decimals() (uint8, error) {
	return _CurvePoolFactoryV2.Contract.Decimals(&_CurvePoolFactoryV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Decimals() (uint8, error) {
	return _CurvePoolFactoryV2.Contract.Decimals(&_CurvePoolFactoryV2.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) EmaPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "ema_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) EmaPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.EmaPrice(&_CurvePoolFactoryV2.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) EmaPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.EmaPrice(&_CurvePoolFactoryV2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Fee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Fee(&_CurvePoolFactoryV2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Fee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Fee(&_CurvePoolFactoryV2.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) FutureA() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureA(&_CurvePoolFactoryV2.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) FutureA() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureA(&_CurvePoolFactoryV2.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) FutureATime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureATime(&_CurvePoolFactoryV2.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) FutureATime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureATime(&_CurvePoolFactoryV2.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) FutureFee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureFee(&_CurvePoolFactoryV2.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) FutureFee() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.FutureFee(&_CurvePoolFactoryV2.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) GetBalances(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) GetBalances() ([2]*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetBalances(&_CurvePoolFactoryV2.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) GetBalances() ([2]*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetBalances(&_CurvePoolFactoryV2.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetDy(&_CurvePoolFactoryV2.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetDy(&_CurvePoolFactoryV2.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) GetP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "get_p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) GetP() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetP(&_CurvePoolFactoryV2.CallOpts)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) GetP() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetP(&_CurvePoolFactoryV2.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) GetVirtualPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetVirtualPrice(&_CurvePoolFactoryV2.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.GetVirtualPrice(&_CurvePoolFactoryV2.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) InitialA() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.InitialA(&_CurvePoolFactoryV2.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) InitialA() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.InitialA(&_CurvePoolFactoryV2.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) InitialATime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.InitialATime(&_CurvePoolFactoryV2.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) InitialATime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.InitialATime(&_CurvePoolFactoryV2.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "last_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) LastPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.LastPrice(&_CurvePoolFactoryV2.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) LastPrice() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.LastPrice(&_CurvePoolFactoryV2.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) MaExpTime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.MaExpTime(&_CurvePoolFactoryV2.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) MaExpTime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.MaExpTime(&_CurvePoolFactoryV2.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) MaLastTime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.MaLastTime(&_CurvePoolFactoryV2.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) MaLastTime() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.MaLastTime(&_CurvePoolFactoryV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Name() (string, error) {
	return _CurvePoolFactoryV2.Contract.Name(&_CurvePoolFactoryV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Name() (string, error) {
	return _CurvePoolFactoryV2.Contract.Name(&_CurvePoolFactoryV2.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Nonces(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.Nonces(&_CurvePoolFactoryV2.CallOpts, arg0)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) PriceOracle() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.PriceOracle(&_CurvePoolFactoryV2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) PriceOracle() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.PriceOracle(&_CurvePoolFactoryV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Symbol() (string, error) {
	return _CurvePoolFactoryV2.Contract.Symbol(&_CurvePoolFactoryV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Symbol() (string, error) {
	return _CurvePoolFactoryV2.Contract.Symbol(&_CurvePoolFactoryV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) TotalSupply() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.TotalSupply(&_CurvePoolFactoryV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) TotalSupply() (*big.Int, error) {
	return _CurvePoolFactoryV2.Contract.TotalSupply(&_CurvePoolFactoryV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePoolFactoryV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Version() (string, error) {
	return _CurvePoolFactoryV2.Contract.Version(&_CurvePoolFactoryV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2CallerSession) Version() (string, error) {
	return _CurvePoolFactoryV2.Contract.Version(&_CurvePoolFactoryV2.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) AddLiquidity(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.AddLiquidity(&_CurvePoolFactoryV2.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.AddLiquidity(&_CurvePoolFactoryV2.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.AddLiquidity0(&_CurvePoolFactoryV2.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.AddLiquidity0(&_CurvePoolFactoryV2.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.ApplyNewFee(&_CurvePoolFactoryV2.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.ApplyNewFee(&_CurvePoolFactoryV2.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Approve(&_CurvePoolFactoryV2.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Approve(&_CurvePoolFactoryV2.TransactOpts, _spender, _value)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) CommitNewFee(opts *bind.TransactOpts, _new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "commit_new_fee", _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.CommitNewFee(&_CurvePoolFactoryV2.TransactOpts, _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.CommitNewFee(&_CurvePoolFactoryV2.TransactOpts, _new_fee)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Exchange(&_CurvePoolFactoryV2.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Exchange(&_CurvePoolFactoryV2.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Exchange0(&_CurvePoolFactoryV2.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Exchange0(&_CurvePoolFactoryV2.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "initialize", _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Initialize(&_CurvePoolFactoryV2.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Initialize(&_CurvePoolFactoryV2.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Permit(&_CurvePoolFactoryV2.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Permit(&_CurvePoolFactoryV2.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RampA(&_CurvePoolFactoryV2.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RampA(&_CurvePoolFactoryV2.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidity(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidity(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidity0(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidity0(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityImbalance(&_CurvePoolFactoryV2.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityImbalance(&_CurvePoolFactoryV2.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityImbalance0(&_CurvePoolFactoryV2.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityImbalance0(&_CurvePoolFactoryV2.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityOneCoin(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityOneCoin(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityOneCoin0(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.RemoveLiquidityOneCoin0(&_CurvePoolFactoryV2.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.SetMaExpTime(&_CurvePoolFactoryV2.TransactOpts, _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.SetMaExpTime(&_CurvePoolFactoryV2.TransactOpts, _ma_exp_time)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) StopRampA() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.StopRampA(&_CurvePoolFactoryV2.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) StopRampA() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.StopRampA(&_CurvePoolFactoryV2.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Transfer(&_CurvePoolFactoryV2.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.Transfer(&_CurvePoolFactoryV2.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.TransferFrom(&_CurvePoolFactoryV2.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.TransferFrom(&_CurvePoolFactoryV2.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Transactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolFactoryV2.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Session) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.WithdrawAdminFees(&_CurvePoolFactoryV2.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2TransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _CurvePoolFactoryV2.Contract.WithdrawAdminFees(&_CurvePoolFactoryV2.TransactOpts)
}

// CurvePoolFactoryV2AddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2AddLiquidityIterator struct {
	Event *CurvePoolFactoryV2AddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2AddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2AddLiquidity)
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
		it.Event = new(CurvePoolFactoryV2AddLiquidity)
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
func (it *CurvePoolFactoryV2AddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2AddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2AddLiquidity represents a AddLiquidity event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2AddLiquidity struct {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolFactoryV2AddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2AddLiquidityIterator{contract: _CurvePoolFactoryV2.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2AddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2AddLiquidity)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseAddLiquidity(log types.Log) (*CurvePoolFactoryV2AddLiquidity, error) {
	event := new(CurvePoolFactoryV2AddLiquidity)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2ApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2ApplyNewFeeIterator struct {
	Event *CurvePoolFactoryV2ApplyNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2ApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2ApplyNewFee)
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
		it.Event = new(CurvePoolFactoryV2ApplyNewFee)
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
func (it *CurvePoolFactoryV2ApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2ApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2ApplyNewFee represents a ApplyNewFee event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2ApplyNewFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterApplyNewFee(opts *bind.FilterOpts) (*CurvePoolFactoryV2ApplyNewFeeIterator, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2ApplyNewFeeIterator{contract: _CurvePoolFactoryV2.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2ApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2ApplyNewFee)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseApplyNewFee(log types.Log) (*CurvePoolFactoryV2ApplyNewFee, error) {
	event := new(CurvePoolFactoryV2ApplyNewFee)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2ApprovalIterator struct {
	Event *CurvePoolFactoryV2Approval // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2Approval)
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
		it.Event = new(CurvePoolFactoryV2Approval)
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
func (it *CurvePoolFactoryV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2Approval represents a Approval event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CurvePoolFactoryV2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2ApprovalIterator{contract: _CurvePoolFactoryV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2Approval)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseApproval(log types.Log) (*CurvePoolFactoryV2Approval, error) {
	event := new(CurvePoolFactoryV2Approval)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2CommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2CommitNewFeeIterator struct {
	Event *CurvePoolFactoryV2CommitNewFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2CommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2CommitNewFee)
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
		it.Event = new(CurvePoolFactoryV2CommitNewFee)
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
func (it *CurvePoolFactoryV2CommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2CommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2CommitNewFee represents a CommitNewFee event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2CommitNewFee struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterCommitNewFee(opts *bind.FilterOpts) (*CurvePoolFactoryV2CommitNewFeeIterator, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2CommitNewFeeIterator{contract: _CurvePoolFactoryV2.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2CommitNewFee) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2CommitNewFee)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseCommitNewFee(log types.Log) (*CurvePoolFactoryV2CommitNewFee, error) {
	event := new(CurvePoolFactoryV2CommitNewFee)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2RampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RampAIterator struct {
	Event *CurvePoolFactoryV2RampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2RampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2RampA)
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
		it.Event = new(CurvePoolFactoryV2RampA)
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
func (it *CurvePoolFactoryV2RampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2RampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2RampA represents a RampA event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterRampA(opts *bind.FilterOpts) (*CurvePoolFactoryV2RampAIterator, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2RampAIterator{contract: _CurvePoolFactoryV2.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2RampA) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2RampA)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseRampA(log types.Log) (*CurvePoolFactoryV2RampA, error) {
	event := new(CurvePoolFactoryV2RampA)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2RemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidityIterator struct {
	Event *CurvePoolFactoryV2RemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2RemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2RemoveLiquidity)
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
		it.Event = new(CurvePoolFactoryV2RemoveLiquidity)
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
func (it *CurvePoolFactoryV2RemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2RemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2RemoveLiquidity represents a RemoveLiquidity event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolFactoryV2RemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2RemoveLiquidityIterator{contract: _CurvePoolFactoryV2.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2RemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2RemoveLiquidity)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseRemoveLiquidity(log types.Log) (*CurvePoolFactoryV2RemoveLiquidity, error) {
	event := new(CurvePoolFactoryV2RemoveLiquidity)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2RemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidityImbalanceIterator struct {
	Event *CurvePoolFactoryV2RemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2RemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2RemoveLiquidityImbalance)
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
		it.Event = new(CurvePoolFactoryV2RemoveLiquidityImbalance)
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
func (it *CurvePoolFactoryV2RemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2RemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2RemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidityImbalance struct {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolFactoryV2RemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2RemoveLiquidityImbalanceIterator{contract: _CurvePoolFactoryV2.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2RemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2RemoveLiquidityImbalance)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseRemoveLiquidityImbalance(log types.Log) (*CurvePoolFactoryV2RemoveLiquidityImbalance, error) {
	event := new(CurvePoolFactoryV2RemoveLiquidityImbalance)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2RemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidityOneIterator struct {
	Event *CurvePoolFactoryV2RemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2RemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2RemoveLiquidityOne)
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
		it.Event = new(CurvePoolFactoryV2RemoveLiquidityOne)
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
func (it *CurvePoolFactoryV2RemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2RemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2RemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2RemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	TokenSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolFactoryV2RemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2RemoveLiquidityOneIterator{contract: _CurvePoolFactoryV2.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2RemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2RemoveLiquidityOne)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseRemoveLiquidityOne(log types.Log) (*CurvePoolFactoryV2RemoveLiquidityOne, error) {
	event := new(CurvePoolFactoryV2RemoveLiquidityOne)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2StopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2StopRampAIterator struct {
	Event *CurvePoolFactoryV2StopRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2StopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2StopRampA)
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
		it.Event = new(CurvePoolFactoryV2StopRampA)
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
func (it *CurvePoolFactoryV2StopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2StopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2StopRampA represents a StopRampA event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2StopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvePoolFactoryV2StopRampAIterator, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2StopRampAIterator{contract: _CurvePoolFactoryV2.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2StopRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2StopRampA)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseStopRampA(log types.Log) (*CurvePoolFactoryV2StopRampA, error) {
	event := new(CurvePoolFactoryV2StopRampA)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2TokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2TokenExchangeIterator struct {
	Event *CurvePoolFactoryV2TokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2TokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2TokenExchange)
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
		it.Event = new(CurvePoolFactoryV2TokenExchange)
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
func (it *CurvePoolFactoryV2TokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2TokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2TokenExchange represents a TokenExchange event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2TokenExchange struct {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvePoolFactoryV2TokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2TokenExchangeIterator{contract: _CurvePoolFactoryV2.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2TokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2TokenExchange)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseTokenExchange(log types.Log) (*CurvePoolFactoryV2TokenExchange, error) {
	event := new(CurvePoolFactoryV2TokenExchange)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolFactoryV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2TransferIterator struct {
	Event *CurvePoolFactoryV2Transfer // Event containing the contract specifics and raw log

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
func (it *CurvePoolFactoryV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolFactoryV2Transfer)
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
		it.Event = new(CurvePoolFactoryV2Transfer)
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
func (it *CurvePoolFactoryV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolFactoryV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolFactoryV2Transfer represents a Transfer event raised by the CurvePoolFactoryV2 contract.
type CurvePoolFactoryV2Transfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*CurvePoolFactoryV2TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFactoryV2TransferIterator{contract: _CurvePoolFactoryV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurvePoolFactoryV2Transfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePoolFactoryV2.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolFactoryV2Transfer)
				if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CurvePoolFactoryV2 *CurvePoolFactoryV2Filterer) ParseTransfer(log types.Log) (*CurvePoolFactoryV2Transfer, error) {
	event := new(CurvePoolFactoryV2Transfer)
	if err := _CurvePoolFactoryV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
