// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pegKeeperPool

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

// PegKeeperPoolMetaData contains all meta data concerning the PegKeeperPool contract.
var PegKeeperPoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewFee\",\"inputs\":[{\"name\":\"new_fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_fee\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_action_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// PegKeeperPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PegKeeperPoolMetaData.ABI instead.
var PegKeeperPoolABI = PegKeeperPoolMetaData.ABI

// PegKeeperPool is an auto generated Go binding around an Ethereum contract.
type PegKeeperPool struct {
	PegKeeperPoolCaller     // Read-only binding to the contract
	PegKeeperPoolTransactor // Write-only binding to the contract
	PegKeeperPoolFilterer   // Log filterer for contract events
}

// PegKeeperPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PegKeeperPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PegKeeperPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PegKeeperPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PegKeeperPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PegKeeperPoolSession struct {
	Contract     *PegKeeperPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PegKeeperPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PegKeeperPoolCallerSession struct {
	Contract *PegKeeperPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PegKeeperPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PegKeeperPoolTransactorSession struct {
	Contract     *PegKeeperPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PegKeeperPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PegKeeperPoolRaw struct {
	Contract *PegKeeperPool // Generic contract binding to access the raw methods on
}

// PegKeeperPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PegKeeperPoolCallerRaw struct {
	Contract *PegKeeperPoolCaller // Generic read-only contract binding to access the raw methods on
}

// PegKeeperPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PegKeeperPoolTransactorRaw struct {
	Contract *PegKeeperPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPegKeeperPool creates a new instance of PegKeeperPool, bound to a specific deployed contract.
func NewPegKeeperPool(address common.Address, backend bind.ContractBackend) (*PegKeeperPool, error) {
	contract, err := bindPegKeeperPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPool{PegKeeperPoolCaller: PegKeeperPoolCaller{contract: contract}, PegKeeperPoolTransactor: PegKeeperPoolTransactor{contract: contract}, PegKeeperPoolFilterer: PegKeeperPoolFilterer{contract: contract}}, nil
}

// NewPegKeeperPoolCaller creates a new read-only instance of PegKeeperPool, bound to a specific deployed contract.
func NewPegKeeperPoolCaller(address common.Address, caller bind.ContractCaller) (*PegKeeperPoolCaller, error) {
	contract, err := bindPegKeeperPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolCaller{contract: contract}, nil
}

// NewPegKeeperPoolTransactor creates a new write-only instance of PegKeeperPool, bound to a specific deployed contract.
func NewPegKeeperPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PegKeeperPoolTransactor, error) {
	contract, err := bindPegKeeperPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolTransactor{contract: contract}, nil
}

// NewPegKeeperPoolFilterer creates a new log filterer instance of PegKeeperPool, bound to a specific deployed contract.
func NewPegKeeperPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PegKeeperPoolFilterer, error) {
	contract, err := bindPegKeeperPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolFilterer{contract: contract}, nil
}

// bindPegKeeperPool binds a generic wrapper to an already deployed contract.
func bindPegKeeperPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PegKeeperPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PegKeeperPool *PegKeeperPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PegKeeperPool.Contract.PegKeeperPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PegKeeperPool *PegKeeperPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.PegKeeperPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PegKeeperPool *PegKeeperPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.PegKeeperPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PegKeeperPool *PegKeeperPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PegKeeperPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PegKeeperPool *PegKeeperPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PegKeeperPool *PegKeeperPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) A() (*big.Int, error) {
	return _PegKeeperPool.Contract.A(&_PegKeeperPool.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) A() (*big.Int, error) {
	return _PegKeeperPool.Contract.A(&_PegKeeperPool.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) APrecise() (*big.Int, error) {
	return _PegKeeperPool.Contract.APrecise(&_PegKeeperPool.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) APrecise() (*big.Int, error) {
	return _PegKeeperPool.Contract.APrecise(&_PegKeeperPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PegKeeperPool *PegKeeperPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PegKeeperPool *PegKeeperPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PegKeeperPool.Contract.DOMAINSEPARATOR(&_PegKeeperPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PegKeeperPool *PegKeeperPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PegKeeperPool.Contract.DOMAINSEPARATOR(&_PegKeeperPool.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) AdminActionDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "admin_action_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) AdminActionDeadline() (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminActionDeadline(&_PegKeeperPool.CallOpts)
}

// AdminActionDeadline is a free data retrieval call binding the contract method 0xe66f43f5.
//
// Solidity: function admin_action_deadline() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) AdminActionDeadline() (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminActionDeadline(&_PegKeeperPool.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminBalances(&_PegKeeperPool.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminBalances(&_PegKeeperPool.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) AdminFee() (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminFee(&_PegKeeperPool.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) AdminFee() (*big.Int, error) {
	return _PegKeeperPool.Contract.AdminFee(&_PegKeeperPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.Allowance(&_PegKeeperPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.Allowance(&_PegKeeperPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.BalanceOf(&_PegKeeperPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.BalanceOf(&_PegKeeperPool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.Balances(&_PegKeeperPool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.Balances(&_PegKeeperPool.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _PegKeeperPool.Contract.CalcTokenAmount(&_PegKeeperPool.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _PegKeeperPool.Contract.CalcTokenAmount(&_PegKeeperPool.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.CalcWithdrawOneCoin(&_PegKeeperPool.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.CalcWithdrawOneCoin(&_PegKeeperPool.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_PegKeeperPool *PegKeeperPoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_PegKeeperPool *PegKeeperPoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _PegKeeperPool.Contract.Coins(&_PegKeeperPool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _PegKeeperPool.Contract.Coins(&_PegKeeperPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PegKeeperPool *PegKeeperPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PegKeeperPool *PegKeeperPoolSession) Decimals() (uint8, error) {
	return _PegKeeperPool.Contract.Decimals(&_PegKeeperPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Decimals() (uint8, error) {
	return _PegKeeperPool.Contract.Decimals(&_PegKeeperPool.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) EmaPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "ema_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) EmaPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.EmaPrice(&_PegKeeperPool.CallOpts)
}

// EmaPrice is a free data retrieval call binding the contract method 0xc24c7c29.
//
// Solidity: function ema_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) EmaPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.EmaPrice(&_PegKeeperPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PegKeeperPool *PegKeeperPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PegKeeperPool *PegKeeperPoolSession) Factory() (common.Address, error) {
	return _PegKeeperPool.Contract.Factory(&_PegKeeperPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Factory() (common.Address, error) {
	return _PegKeeperPool.Contract.Factory(&_PegKeeperPool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Fee() (*big.Int, error) {
	return _PegKeeperPool.Contract.Fee(&_PegKeeperPool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Fee() (*big.Int, error) {
	return _PegKeeperPool.Contract.Fee(&_PegKeeperPool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) FutureA() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureA(&_PegKeeperPool.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) FutureA() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureA(&_PegKeeperPool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) FutureATime() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureATime(&_PegKeeperPool.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) FutureATime() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureATime(&_PegKeeperPool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) FutureFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "future_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) FutureFee() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureFee(&_PegKeeperPool.CallOpts)
}

// FutureFee is a free data retrieval call binding the contract method 0x58680d0b.
//
// Solidity: function future_fee() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) FutureFee() (*big.Int, error) {
	return _PegKeeperPool.Contract.FutureFee(&_PegKeeperPool.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolCaller) GetBalances(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolSession) GetBalances() ([2]*big.Int, error) {
	return _PegKeeperPool.Contract.GetBalances(&_PegKeeperPool.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolCallerSession) GetBalances() ([2]*big.Int, error) {
	return _PegKeeperPool.Contract.GetBalances(&_PegKeeperPool.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.GetDx(&_PegKeeperPool.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.GetDx(&_PegKeeperPool.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.GetDy(&_PegKeeperPool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _PegKeeperPool.Contract.GetDy(&_PegKeeperPool.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) GetP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "get_p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) GetP() (*big.Int, error) {
	return _PegKeeperPool.Contract.GetP(&_PegKeeperPool.CallOpts)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) GetP() (*big.Int, error) {
	return _PegKeeperPool.Contract.GetP(&_PegKeeperPool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) GetVirtualPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.GetVirtualPrice(&_PegKeeperPool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.GetVirtualPrice(&_PegKeeperPool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) InitialA() (*big.Int, error) {
	return _PegKeeperPool.Contract.InitialA(&_PegKeeperPool.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) InitialA() (*big.Int, error) {
	return _PegKeeperPool.Contract.InitialA(&_PegKeeperPool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) InitialATime() (*big.Int, error) {
	return _PegKeeperPool.Contract.InitialATime(&_PegKeeperPool.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) InitialATime() (*big.Int, error) {
	return _PegKeeperPool.Contract.InitialATime(&_PegKeeperPool.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "last_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) LastPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.LastPrice(&_PegKeeperPool.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) LastPrice() (*big.Int, error) {
	return _PegKeeperPool.Contract.LastPrice(&_PegKeeperPool.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) MaExpTime() (*big.Int, error) {
	return _PegKeeperPool.Contract.MaExpTime(&_PegKeeperPool.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) MaExpTime() (*big.Int, error) {
	return _PegKeeperPool.Contract.MaExpTime(&_PegKeeperPool.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) MaLastTime() (*big.Int, error) {
	return _PegKeeperPool.Contract.MaLastTime(&_PegKeeperPool.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) MaLastTime() (*big.Int, error) {
	return _PegKeeperPool.Contract.MaLastTime(&_PegKeeperPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PegKeeperPool *PegKeeperPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PegKeeperPool *PegKeeperPoolSession) Name() (string, error) {
	return _PegKeeperPool.Contract.Name(&_PegKeeperPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Name() (string, error) {
	return _PegKeeperPool.Contract.Name(&_PegKeeperPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.Nonces(&_PegKeeperPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PegKeeperPool.Contract.Nonces(&_PegKeeperPool.CallOpts, arg0)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) PriceOracle() (*big.Int, error) {
	return _PegKeeperPool.Contract.PriceOracle(&_PegKeeperPool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) PriceOracle() (*big.Int, error) {
	return _PegKeeperPool.Contract.PriceOracle(&_PegKeeperPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PegKeeperPool *PegKeeperPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PegKeeperPool *PegKeeperPoolSession) Symbol() (string, error) {
	return _PegKeeperPool.Contract.Symbol(&_PegKeeperPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Symbol() (string, error) {
	return _PegKeeperPool.Contract.Symbol(&_PegKeeperPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) TotalSupply() (*big.Int, error) {
	return _PegKeeperPool.Contract.TotalSupply(&_PegKeeperPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PegKeeperPool *PegKeeperPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _PegKeeperPool.Contract.TotalSupply(&_PegKeeperPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PegKeeperPool *PegKeeperPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PegKeeperPool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PegKeeperPool *PegKeeperPoolSession) Version() (string, error) {
	return _PegKeeperPool.Contract.Version(&_PegKeeperPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PegKeeperPool *PegKeeperPoolCallerSession) Version() (string, error) {
	return _PegKeeperPool.Contract.Version(&_PegKeeperPool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.AddLiquidity(&_PegKeeperPool.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) AddLiquidity(_amounts [2]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.AddLiquidity(&_PegKeeperPool.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.AddLiquidity0(&_PegKeeperPool.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x0c3e4b54.
//
// Solidity: function add_liquidity(uint256[2] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) AddLiquidity0(_amounts [2]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.AddLiquidity0(&_PegKeeperPool.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) ApplyNewFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "apply_new_fee")
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_PegKeeperPool *PegKeeperPoolSession) ApplyNewFee() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.ApplyNewFee(&_PegKeeperPool.TransactOpts)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0x4f12fe97.
//
// Solidity: function apply_new_fee() returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) ApplyNewFee() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.ApplyNewFee(&_PegKeeperPool.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Approve(&_PegKeeperPool.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Approve(&_PegKeeperPool.TransactOpts, _spender, _value)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) CommitNewFee(opts *bind.TransactOpts, _new_fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "commit_new_fee", _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_PegKeeperPool *PegKeeperPoolSession) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.CommitNewFee(&_PegKeeperPool.TransactOpts, _new_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xa48eac9d.
//
// Solidity: function commit_new_fee(uint256 _new_fee) returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) CommitNewFee(_new_fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.CommitNewFee(&_PegKeeperPool.TransactOpts, _new_fee)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Exchange(&_PegKeeperPool.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Exchange(&_PegKeeperPool.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Exchange0(&_PegKeeperPool.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Exchange0(&_PegKeeperPool.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "initialize", _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_PegKeeperPool *PegKeeperPoolSession) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Initialize(&_PegKeeperPool.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xa461b3c8.
//
// Solidity: function initialize(string _name, string _symbol, address[4] _coins, uint256[4] _rate_multipliers, uint256 _A, uint256 _fee) returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Initialize(_name string, _symbol string, _coins [4]common.Address, _rate_multipliers [4]*big.Int, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Initialize(&_PegKeeperPool.TransactOpts, _name, _symbol, _coins, _rate_multipliers, _A, _fee)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_PegKeeperPool *PegKeeperPoolSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Permit(&_PegKeeperPool.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Permit(&_PegKeeperPool.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_PegKeeperPool *PegKeeperPoolSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RampA(&_PegKeeperPool.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RampA(&_PegKeeperPool.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidity(&_PegKeeperPool.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidity(&_PegKeeperPool.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidity0(&_PegKeeperPool.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x3eb1719f.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[2] _min_amounts, address _receiver) returns(uint256[2])
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts [2]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidity0(&_PegKeeperPool.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityImbalance(&_PegKeeperPool.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidityImbalance(_amounts [2]*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityImbalance(&_PegKeeperPool.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityImbalance0(&_PegKeeperPool.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x52d2cfdd.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidityImbalance0(_amounts [2]*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityImbalance0(&_PegKeeperPool.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityOneCoin(&_PegKeeperPool.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityOneCoin(&_PegKeeperPool.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityOneCoin0(&_PegKeeperPool.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.RemoveLiquidityOneCoin0(&_PegKeeperPool.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_PegKeeperPool *PegKeeperPoolSession) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.SetMaExpTime(&_PegKeeperPool.TransactOpts, _ma_exp_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x7f3e17cb.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time) returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) SetMaExpTime(_ma_exp_time *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.SetMaExpTime(&_PegKeeperPool.TransactOpts, _ma_exp_time)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_PegKeeperPool *PegKeeperPoolSession) StopRampA() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.StopRampA(&_PegKeeperPool.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) StopRampA() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.StopRampA(&_PegKeeperPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Transfer(&_PegKeeperPool.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.Transfer(&_PegKeeperPool.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.TransferFrom(&_PegKeeperPool.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_PegKeeperPool *PegKeeperPoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PegKeeperPool.Contract.TransferFrom(&_PegKeeperPool.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_PegKeeperPool *PegKeeperPoolTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PegKeeperPool.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_PegKeeperPool *PegKeeperPoolSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.WithdrawAdminFees(&_PegKeeperPool.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_PegKeeperPool *PegKeeperPoolTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _PegKeeperPool.Contract.WithdrawAdminFees(&_PegKeeperPool.TransactOpts)
}

// PegKeeperPoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the PegKeeperPool contract.
type PegKeeperPoolAddLiquidityIterator struct {
	Event *PegKeeperPoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolAddLiquidity)
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
		it.Event = new(PegKeeperPoolAddLiquidity)
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
func (it *PegKeeperPoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolAddLiquidity represents a AddLiquidity event raised by the PegKeeperPool contract.
type PegKeeperPoolAddLiquidity struct {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*PegKeeperPoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolAddLiquidityIterator{contract: _PegKeeperPool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolAddLiquidity)
				if err := _PegKeeperPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseAddLiquidity(log types.Log) (*PegKeeperPoolAddLiquidity, error) {
	event := new(PegKeeperPoolAddLiquidity)
	if err := _PegKeeperPool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the PegKeeperPool contract.
type PegKeeperPoolApplyNewFeeIterator struct {
	Event *PegKeeperPoolApplyNewFee // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolApplyNewFee)
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
		it.Event = new(PegKeeperPoolApplyNewFee)
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
func (it *PegKeeperPoolApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolApplyNewFee represents a ApplyNewFee event raised by the PegKeeperPool contract.
type PegKeeperPoolApplyNewFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*PegKeeperPoolApplyNewFeeIterator, error) {

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolApplyNewFeeIterator{contract: _PegKeeperPool.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0xa8715770654f54603947addf38c689adbd7182e21673b28bcf306a957aaba215.
//
// Solidity: event ApplyNewFee(uint256 fee)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolApplyNewFee)
				if err := _PegKeeperPool.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseApplyNewFee(log types.Log) (*PegKeeperPoolApplyNewFee, error) {
	event := new(PegKeeperPoolApplyNewFee)
	if err := _PegKeeperPool.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PegKeeperPool contract.
type PegKeeperPoolApprovalIterator struct {
	Event *PegKeeperPoolApproval // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolApproval)
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
		it.Event = new(PegKeeperPoolApproval)
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
func (it *PegKeeperPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolApproval represents a Approval event raised by the PegKeeperPool contract.
type PegKeeperPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PegKeeperPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolApprovalIterator{contract: _PegKeeperPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolApproval)
				if err := _PegKeeperPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseApproval(log types.Log) (*PegKeeperPoolApproval, error) {
	event := new(PegKeeperPoolApproval)
	if err := _PegKeeperPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolCommitNewFeeIterator is returned from FilterCommitNewFee and is used to iterate over the raw logs and unpacked data for CommitNewFee events raised by the PegKeeperPool contract.
type PegKeeperPoolCommitNewFeeIterator struct {
	Event *PegKeeperPoolCommitNewFee // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolCommitNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolCommitNewFee)
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
		it.Event = new(PegKeeperPoolCommitNewFee)
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
func (it *PegKeeperPoolCommitNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolCommitNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolCommitNewFee represents a CommitNewFee event raised by the PegKeeperPool contract.
type PegKeeperPoolCommitNewFee struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommitNewFee is a free log retrieval operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterCommitNewFee(opts *bind.FilterOpts) (*PegKeeperPoolCommitNewFeeIterator, error) {

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolCommitNewFeeIterator{contract: _PegKeeperPool.contract, event: "CommitNewFee", logs: logs, sub: sub}, nil
}

// WatchCommitNewFee is a free log subscription operation binding the contract event 0x878eb36b3f197f05821c06953d9bc8f14b332a227b1e26df06a4215bbfe5d73f.
//
// Solidity: event CommitNewFee(uint256 new_fee)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchCommitNewFee(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolCommitNewFee) (event.Subscription, error) {

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "CommitNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolCommitNewFee)
				if err := _PegKeeperPool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseCommitNewFee(log types.Log) (*PegKeeperPoolCommitNewFee, error) {
	event := new(PegKeeperPoolCommitNewFee)
	if err := _PegKeeperPool.contract.UnpackLog(event, "CommitNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the PegKeeperPool contract.
type PegKeeperPoolRampAIterator struct {
	Event *PegKeeperPoolRampA // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolRampA)
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
		it.Event = new(PegKeeperPoolRampA)
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
func (it *PegKeeperPoolRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolRampA represents a RampA event raised by the PegKeeperPool contract.
type PegKeeperPoolRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterRampA(opts *bind.FilterOpts) (*PegKeeperPoolRampAIterator, error) {

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolRampAIterator{contract: _PegKeeperPool.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolRampA) (event.Subscription, error) {

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolRampA)
				if err := _PegKeeperPool.contract.UnpackLog(event, "RampA", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseRampA(log types.Log) (*PegKeeperPoolRampA, error) {
	event := new(PegKeeperPoolRampA)
	if err := _PegKeeperPool.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidityIterator struct {
	Event *PegKeeperPoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolRemoveLiquidity)
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
		it.Event = new(PegKeeperPoolRemoveLiquidity)
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
func (it *PegKeeperPoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolRemoveLiquidity represents a RemoveLiquidity event raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*PegKeeperPoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolRemoveLiquidityIterator{contract: _PegKeeperPool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolRemoveLiquidity)
				if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseRemoveLiquidity(log types.Log) (*PegKeeperPoolRemoveLiquidity, error) {
	event := new(PegKeeperPoolRemoveLiquidity)
	if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidityImbalanceIterator struct {
	Event *PegKeeperPoolRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolRemoveLiquidityImbalance)
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
		it.Event = new(PegKeeperPoolRemoveLiquidityImbalance)
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
func (it *PegKeeperPoolRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidityImbalance struct {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*PegKeeperPoolRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolRemoveLiquidityImbalanceIterator{contract: _PegKeeperPool.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolRemoveLiquidityImbalance)
				if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*PegKeeperPoolRemoveLiquidityImbalance, error) {
	event := new(PegKeeperPoolRemoveLiquidityImbalance)
	if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidityOneIterator struct {
	Event *PegKeeperPoolRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolRemoveLiquidityOne)
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
		it.Event = new(PegKeeperPoolRemoveLiquidityOne)
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
func (it *PegKeeperPoolRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the PegKeeperPool contract.
type PegKeeperPoolRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinAmount  *big.Int
	TokenSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*PegKeeperPoolRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolRemoveLiquidityOneIterator{contract: _PegKeeperPool.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolRemoveLiquidityOne)
				if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseRemoveLiquidityOne(log types.Log) (*PegKeeperPoolRemoveLiquidityOne, error) {
	event := new(PegKeeperPoolRemoveLiquidityOne)
	if err := _PegKeeperPool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the PegKeeperPool contract.
type PegKeeperPoolStopRampAIterator struct {
	Event *PegKeeperPoolStopRampA // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolStopRampA)
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
		it.Event = new(PegKeeperPoolStopRampA)
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
func (it *PegKeeperPoolStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolStopRampA represents a StopRampA event raised by the PegKeeperPool contract.
type PegKeeperPoolStopRampA struct {
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterStopRampA(opts *bind.FilterOpts) (*PegKeeperPoolStopRampAIterator, error) {

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolStopRampAIterator{contract: _PegKeeperPool.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolStopRampA) (event.Subscription, error) {

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolStopRampA)
				if err := _PegKeeperPool.contract.UnpackLog(event, "StopRampA", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseStopRampA(log types.Log) (*PegKeeperPoolStopRampA, error) {
	event := new(PegKeeperPoolStopRampA)
	if err := _PegKeeperPool.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the PegKeeperPool contract.
type PegKeeperPoolTokenExchangeIterator struct {
	Event *PegKeeperPoolTokenExchange // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolTokenExchange)
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
		it.Event = new(PegKeeperPoolTokenExchange)
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
func (it *PegKeeperPoolTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolTokenExchange represents a TokenExchange event raised by the PegKeeperPool contract.
type PegKeeperPoolTokenExchange struct {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*PegKeeperPoolTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolTokenExchangeIterator{contract: _PegKeeperPool.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolTokenExchange)
				if err := _PegKeeperPool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseTokenExchange(log types.Log) (*PegKeeperPoolTokenExchange, error) {
	event := new(PegKeeperPoolTokenExchange)
	if err := _PegKeeperPool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PegKeeperPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PegKeeperPool contract.
type PegKeeperPoolTransferIterator struct {
	Event *PegKeeperPoolTransfer // Event containing the contract specifics and raw log

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
func (it *PegKeeperPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PegKeeperPoolTransfer)
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
		it.Event = new(PegKeeperPoolTransfer)
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
func (it *PegKeeperPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PegKeeperPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PegKeeperPoolTransfer represents a Transfer event raised by the PegKeeperPool contract.
type PegKeeperPoolTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_PegKeeperPool *PegKeeperPoolFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*PegKeeperPoolTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PegKeeperPool.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PegKeeperPoolTransferIterator{contract: _PegKeeperPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_PegKeeperPool *PegKeeperPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PegKeeperPoolTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _PegKeeperPool.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PegKeeperPoolTransfer)
				if err := _PegKeeperPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PegKeeperPool *PegKeeperPoolFilterer) ParseTransfer(log types.Log) (*PegKeeperPoolTransfer, error) {
	event := new(PegKeeperPoolTransfer)
	if err := _PegKeeperPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
