// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdAmm

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

// CrvUsdAmmMetaData contains all meta data concerning the CrvUsdAmm contract.
var CrvUsdAmmMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"n1\",\"type\":\"int256\",\"indexed\":false},{\"name\":\"n2\",\"type\":\"int256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount_borrowed\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"amount_collateral\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetRate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"rate_mul\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdminFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_borrowed_token\",\"type\":\"address\"},{\"name\":\"_borrowed_precision\",\"type\":\"uint256\"},{\"name\":\"_collateral_token\",\"type\":\"address\"},{\"name\":\"_collateral_precision\",\"type\":\"uint256\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_sqrt_band_ratio\",\"type\":\"uint256\"},{\"name\":\"_log_A_ratio\",\"type\":\"int256\"},{\"name\":\"_base_price\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"_price_oracle_contract\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"dynamic_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_rate_mul\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"p_current_up\",\"inputs\":[{\"name\":\"n\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"p_current_down\",\"inputs\":[{\"name\":\"n\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"p_oracle_up\",\"inputs\":[{\"name\":\"n\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"p_oracle_down\",\"inputs\":[{\"name\":\"n\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"read_user_tick_numbers\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"can_skip_bands\",\"inputs\":[{\"name\":\"n_end\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"active_band_with_skip\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"has_liquidity\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit_range\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"n1\",\"type\":\"int256\"},{\"name\":\"n2\",\"type\":\"int256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"frac\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"in_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dxdy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"in_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"out_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dydx\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"out_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"in_amount\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"in_amount\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"out_amount\",\"type\":\"uint256\"},{\"name\":\"max_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"out_amount\",\"type\":\"uint256\"},{\"name\":\"max_amount\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_y_up\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_x_down\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_sum_xy\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_xy\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_amount_for_price\",\"inputs\":[{\"name\":\"p\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_rate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"reset_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_callback\",\"inputs\":[{\"name\":\"liquidity_mining_callback\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"active_band\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_band\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_band\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fees_x\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fees_y\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle_contract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bands_x\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bands_y\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"liquidity_mining_callback\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CrvUsdAmmABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdAmmMetaData.ABI instead.
var CrvUsdAmmABI = CrvUsdAmmMetaData.ABI

// CrvUsdAmm is an auto generated Go binding around an Ethereum contract.
type CrvUsdAmm struct {
	CrvUsdAmmCaller     // Read-only binding to the contract
	CrvUsdAmmTransactor // Write-only binding to the contract
	CrvUsdAmmFilterer   // Log filterer for contract events
}

// CrvUsdAmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdAmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdAmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdAmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdAmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdAmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdAmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdAmmSession struct {
	Contract     *CrvUsdAmm        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvUsdAmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdAmmCallerSession struct {
	Contract *CrvUsdAmmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrvUsdAmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdAmmTransactorSession struct {
	Contract     *CrvUsdAmmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrvUsdAmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdAmmRaw struct {
	Contract *CrvUsdAmm // Generic contract binding to access the raw methods on
}

// CrvUsdAmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdAmmCallerRaw struct {
	Contract *CrvUsdAmmCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdAmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdAmmTransactorRaw struct {
	Contract *CrvUsdAmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdAmm creates a new instance of CrvUsdAmm, bound to a specific deployed contract.
func NewCrvUsdAmm(address common.Address, backend bind.ContractBackend) (*CrvUsdAmm, error) {
	contract, err := bindCrvUsdAmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmm{CrvUsdAmmCaller: CrvUsdAmmCaller{contract: contract}, CrvUsdAmmTransactor: CrvUsdAmmTransactor{contract: contract}, CrvUsdAmmFilterer: CrvUsdAmmFilterer{contract: contract}}, nil
}

// NewCrvUsdAmmCaller creates a new read-only instance of CrvUsdAmm, bound to a specific deployed contract.
func NewCrvUsdAmmCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdAmmCaller, error) {
	contract, err := bindCrvUsdAmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmCaller{contract: contract}, nil
}

// NewCrvUsdAmmTransactor creates a new write-only instance of CrvUsdAmm, bound to a specific deployed contract.
func NewCrvUsdAmmTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdAmmTransactor, error) {
	contract, err := bindCrvUsdAmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmTransactor{contract: contract}, nil
}

// NewCrvUsdAmmFilterer creates a new log filterer instance of CrvUsdAmm, bound to a specific deployed contract.
func NewCrvUsdAmmFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdAmmFilterer, error) {
	contract, err := bindCrvUsdAmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmFilterer{contract: contract}, nil
}

// bindCrvUsdAmm binds a generic wrapper to an already deployed contract.
func bindCrvUsdAmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdAmmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdAmm *CrvUsdAmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdAmm.Contract.CrvUsdAmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdAmm *CrvUsdAmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.CrvUsdAmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdAmm *CrvUsdAmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.CrvUsdAmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdAmm *CrvUsdAmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdAmm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdAmm *CrvUsdAmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdAmm *CrvUsdAmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) A() (*big.Int, error) {
	return _CrvUsdAmm.Contract.A(&_CrvUsdAmm.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) A() (*big.Int, error) {
	return _CrvUsdAmm.Contract.A(&_CrvUsdAmm.CallOpts)
}

// ActiveBand is a free data retrieval call binding the contract method 0x8f8654c5.
//
// Solidity: function active_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCaller) ActiveBand(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "active_band")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveBand is a free data retrieval call binding the contract method 0x8f8654c5.
//
// Solidity: function active_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmSession) ActiveBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.ActiveBand(&_CrvUsdAmm.CallOpts)
}

// ActiveBand is a free data retrieval call binding the contract method 0x8f8654c5.
//
// Solidity: function active_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) ActiveBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.ActiveBand(&_CrvUsdAmm.CallOpts)
}

// ActiveBandWithSkip is a free data retrieval call binding the contract method 0xc16ef264.
//
// Solidity: function active_band_with_skip() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCaller) ActiveBandWithSkip(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "active_band_with_skip")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveBandWithSkip is a free data retrieval call binding the contract method 0xc16ef264.
//
// Solidity: function active_band_with_skip() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmSession) ActiveBandWithSkip() (*big.Int, error) {
	return _CrvUsdAmm.Contract.ActiveBandWithSkip(&_CrvUsdAmm.CallOpts)
}

// ActiveBandWithSkip is a free data retrieval call binding the contract method 0xc16ef264.
//
// Solidity: function active_band_with_skip() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) ActiveBandWithSkip() (*big.Int, error) {
	return _CrvUsdAmm.Contract.ActiveBandWithSkip(&_CrvUsdAmm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmSession) Admin() (common.Address, error) {
	return _CrvUsdAmm.Contract.Admin(&_CrvUsdAmm.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) Admin() (common.Address, error) {
	return _CrvUsdAmm.Contract.Admin(&_CrvUsdAmm.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) AdminFee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFee(&_CrvUsdAmm.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) AdminFee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFee(&_CrvUsdAmm.CallOpts)
}

// AdminFeesX is a free data retrieval call binding the contract method 0xd1fea733.
//
// Solidity: function admin_fees_x() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) AdminFeesX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "admin_fees_x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFeesX is a free data retrieval call binding the contract method 0xd1fea733.
//
// Solidity: function admin_fees_x() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) AdminFeesX() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFeesX(&_CrvUsdAmm.CallOpts)
}

// AdminFeesX is a free data retrieval call binding the contract method 0xd1fea733.
//
// Solidity: function admin_fees_x() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) AdminFeesX() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFeesX(&_CrvUsdAmm.CallOpts)
}

// AdminFeesY is a free data retrieval call binding the contract method 0x89960ba7.
//
// Solidity: function admin_fees_y() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) AdminFeesY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "admin_fees_y")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFeesY is a free data retrieval call binding the contract method 0x89960ba7.
//
// Solidity: function admin_fees_y() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) AdminFeesY() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFeesY(&_CrvUsdAmm.CallOpts)
}

// AdminFeesY is a free data retrieval call binding the contract method 0x89960ba7.
//
// Solidity: function admin_fees_y() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) AdminFeesY() (*big.Int, error) {
	return _CrvUsdAmm.Contract.AdminFeesY(&_CrvUsdAmm.CallOpts)
}

// BandsX is a free data retrieval call binding the contract method 0xebcb0067.
//
// Solidity: function bands_x(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) BandsX(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "bands_x", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BandsX is a free data retrieval call binding the contract method 0xebcb0067.
//
// Solidity: function bands_x(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) BandsX(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.BandsX(&_CrvUsdAmm.CallOpts, arg0)
}

// BandsX is a free data retrieval call binding the contract method 0xebcb0067.
//
// Solidity: function bands_x(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) BandsX(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.BandsX(&_CrvUsdAmm.CallOpts, arg0)
}

// BandsY is a free data retrieval call binding the contract method 0x31f7e306.
//
// Solidity: function bands_y(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) BandsY(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "bands_y", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BandsY is a free data retrieval call binding the contract method 0x31f7e306.
//
// Solidity: function bands_y(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) BandsY(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.BandsY(&_CrvUsdAmm.CallOpts, arg0)
}

// BandsY is a free data retrieval call binding the contract method 0x31f7e306.
//
// Solidity: function bands_y(int256 arg0) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) BandsY(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.BandsY(&_CrvUsdAmm.CallOpts, arg0)
}

// CanSkipBands is a free data retrieval call binding the contract method 0xec654706.
//
// Solidity: function can_skip_bands(int256 n_end) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmCaller) CanSkipBands(opts *bind.CallOpts, n_end *big.Int) (bool, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "can_skip_bands", n_end)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSkipBands is a free data retrieval call binding the contract method 0xec654706.
//
// Solidity: function can_skip_bands(int256 n_end) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmSession) CanSkipBands(n_end *big.Int) (bool, error) {
	return _CrvUsdAmm.Contract.CanSkipBands(&_CrvUsdAmm.CallOpts, n_end)
}

// CanSkipBands is a free data retrieval call binding the contract method 0xec654706.
//
// Solidity: function can_skip_bands(int256 n_end) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) CanSkipBands(n_end *big.Int) (bool, error) {
	return _CrvUsdAmm.Contract.CanSkipBands(&_CrvUsdAmm.CallOpts, n_end)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) pure returns(address)
func (_CrvUsdAmm *CrvUsdAmmCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) pure returns(address)
func (_CrvUsdAmm *CrvUsdAmmSession) Coins(i *big.Int) (common.Address, error) {
	return _CrvUsdAmm.Contract.Coins(&_CrvUsdAmm.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) pure returns(address)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _CrvUsdAmm.Contract.Coins(&_CrvUsdAmm.CallOpts, i)
}

// DynamicFee is a free data retrieval call binding the contract method 0x77c34594.
//
// Solidity: function dynamic_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) DynamicFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "dynamic_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x77c34594.
//
// Solidity: function dynamic_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) DynamicFee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.DynamicFee(&_CrvUsdAmm.CallOpts)
}

// DynamicFee is a free data retrieval call binding the contract method 0x77c34594.
//
// Solidity: function dynamic_fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) DynamicFee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.DynamicFee(&_CrvUsdAmm.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) Fee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.Fee(&_CrvUsdAmm.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) Fee() (*big.Int, error) {
	return _CrvUsdAmm.Contract.Fee(&_CrvUsdAmm.CallOpts)
}

// GetAmountForPrice is a free data retrieval call binding the contract method 0x48e995f9.
//
// Solidity: function get_amount_for_price(uint256 p) view returns(uint256, bool)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetAmountForPrice(opts *bind.CallOpts, p *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_amount_for_price", p)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetAmountForPrice is a free data retrieval call binding the contract method 0x48e995f9.
//
// Solidity: function get_amount_for_price(uint256 p) view returns(uint256, bool)
func (_CrvUsdAmm *CrvUsdAmmSession) GetAmountForPrice(p *big.Int) (*big.Int, bool, error) {
	return _CrvUsdAmm.Contract.GetAmountForPrice(&_CrvUsdAmm.CallOpts, p)
}

// GetAmountForPrice is a free data retrieval call binding the contract method 0x48e995f9.
//
// Solidity: function get_amount_for_price(uint256 p) view returns(uint256, bool)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetAmountForPrice(p *big.Int) (*big.Int, bool, error) {
	return _CrvUsdAmm.Contract.GetAmountForPrice(&_CrvUsdAmm.CallOpts, p)
}

// GetBasePrice is a free data retrieval call binding the contract method 0xa7db79a5.
//
// Solidity: function get_base_price() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetBasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_base_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasePrice is a free data retrieval call binding the contract method 0xa7db79a5.
//
// Solidity: function get_base_price() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetBasePrice() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetBasePrice(&_CrvUsdAmm.CallOpts)
}

// GetBasePrice is a free data retrieval call binding the contract method 0xa7db79a5.
//
// Solidity: function get_base_price() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetBasePrice() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetBasePrice(&_CrvUsdAmm.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_dx", i, j, out_amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetDx(i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetDx(&_CrvUsdAmm.CallOpts, i, j, out_amount)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetDx(i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetDx(&_CrvUsdAmm.CallOpts, i, j, out_amount)
}

// GetDxdy is a free data retrieval call binding the contract method 0xc49202e7.
//
// Solidity: function get_dxdy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetDxdy(opts *bind.CallOpts, i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_dxdy", i, j, in_amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDxdy is a free data retrieval call binding the contract method 0xc49202e7.
//
// Solidity: function get_dxdy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetDxdy(i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CrvUsdAmm.Contract.GetDxdy(&_CrvUsdAmm.CallOpts, i, j, in_amount)
}

// GetDxdy is a free data retrieval call binding the contract method 0xc49202e7.
//
// Solidity: function get_dxdy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetDxdy(i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CrvUsdAmm.Contract.GetDxdy(&_CrvUsdAmm.CallOpts, i, j, in_amount)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_dy", i, j, in_amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetDy(i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetDy(&_CrvUsdAmm.CallOpts, i, j, in_amount)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 in_amount) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetDy(i *big.Int, j *big.Int, in_amount *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetDy(&_CrvUsdAmm.CallOpts, i, j, in_amount)
}

// GetDydx is a free data retrieval call binding the contract method 0xed7110cf.
//
// Solidity: function get_dydx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetDydx(opts *bind.CallOpts, i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_dydx", i, j, out_amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDydx is a free data retrieval call binding the contract method 0xed7110cf.
//
// Solidity: function get_dydx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetDydx(i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CrvUsdAmm.Contract.GetDydx(&_CrvUsdAmm.CallOpts, i, j, out_amount)
}

// GetDydx is a free data retrieval call binding the contract method 0xed7110cf.
//
// Solidity: function get_dydx(uint256 i, uint256 j, uint256 out_amount) view returns(uint256, uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetDydx(i *big.Int, j *big.Int, out_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CrvUsdAmm.Contract.GetDydx(&_CrvUsdAmm.CallOpts, i, j, out_amount)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_p")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetP() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetP(&_CrvUsdAmm.CallOpts)
}

// GetP is a free data retrieval call binding the contract method 0xf2388acb.
//
// Solidity: function get_p() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetP() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetP(&_CrvUsdAmm.CallOpts)
}

// GetRateMul is a free data retrieval call binding the contract method 0x095a0fc6.
//
// Solidity: function get_rate_mul() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetRateMul(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_rate_mul")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateMul is a free data retrieval call binding the contract method 0x095a0fc6.
//
// Solidity: function get_rate_mul() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetRateMul() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetRateMul(&_CrvUsdAmm.CallOpts)
}

// GetRateMul is a free data retrieval call binding the contract method 0x095a0fc6.
//
// Solidity: function get_rate_mul() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetRateMul() (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetRateMul(&_CrvUsdAmm.CallOpts)
}

// GetSumXy is a free data retrieval call binding the contract method 0x544fb5c1.
//
// Solidity: function get_sum_xy(address user) view returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmCaller) GetSumXy(opts *bind.CallOpts, user common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_sum_xy", user)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetSumXy is a free data retrieval call binding the contract method 0x544fb5c1.
//
// Solidity: function get_sum_xy(address user) view returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) GetSumXy(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdAmm.Contract.GetSumXy(&_CrvUsdAmm.CallOpts, user)
}

// GetSumXy is a free data retrieval call binding the contract method 0x544fb5c1.
//
// Solidity: function get_sum_xy(address user) view returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetSumXy(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdAmm.Contract.GetSumXy(&_CrvUsdAmm.CallOpts, user)
}

// GetXDown is a free data retrieval call binding the contract method 0x62ca4b18.
//
// Solidity: function get_x_down(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetXDown(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_x_down", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetXDown is a free data retrieval call binding the contract method 0x62ca4b18.
//
// Solidity: function get_x_down(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetXDown(user common.Address) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetXDown(&_CrvUsdAmm.CallOpts, user)
}

// GetXDown is a free data retrieval call binding the contract method 0x62ca4b18.
//
// Solidity: function get_x_down(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetXDown(user common.Address) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetXDown(&_CrvUsdAmm.CallOpts, user)
}

// GetXy is a free data retrieval call binding the contract method 0x84738380.
//
// Solidity: function get_xy(address user) view returns(uint256[][2])
func (_CrvUsdAmm *CrvUsdAmmCaller) GetXy(opts *bind.CallOpts, user common.Address) ([2][]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_xy", user)

	if err != nil {
		return *new([2][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2][]*big.Int)).(*[2][]*big.Int)

	return out0, err

}

// GetXy is a free data retrieval call binding the contract method 0x84738380.
//
// Solidity: function get_xy(address user) view returns(uint256[][2])
func (_CrvUsdAmm *CrvUsdAmmSession) GetXy(user common.Address) ([2][]*big.Int, error) {
	return _CrvUsdAmm.Contract.GetXy(&_CrvUsdAmm.CallOpts, user)
}

// GetXy is a free data retrieval call binding the contract method 0x84738380.
//
// Solidity: function get_xy(address user) view returns(uint256[][2])
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetXy(user common.Address) ([2][]*big.Int, error) {
	return _CrvUsdAmm.Contract.GetXy(&_CrvUsdAmm.CallOpts, user)
}

// GetYUp is a free data retrieval call binding the contract method 0xee4c32ee.
//
// Solidity: function get_y_up(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) GetYUp(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "get_y_up", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYUp is a free data retrieval call binding the contract method 0xee4c32ee.
//
// Solidity: function get_y_up(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) GetYUp(user common.Address) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetYUp(&_CrvUsdAmm.CallOpts, user)
}

// GetYUp is a free data retrieval call binding the contract method 0xee4c32ee.
//
// Solidity: function get_y_up(address user) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) GetYUp(user common.Address) (*big.Int, error) {
	return _CrvUsdAmm.Contract.GetYUp(&_CrvUsdAmm.CallOpts, user)
}

// HasLiquidity is a free data retrieval call binding the contract method 0xe8dd1ef1.
//
// Solidity: function has_liquidity(address user) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmCaller) HasLiquidity(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "has_liquidity", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLiquidity is a free data retrieval call binding the contract method 0xe8dd1ef1.
//
// Solidity: function has_liquidity(address user) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmSession) HasLiquidity(user common.Address) (bool, error) {
	return _CrvUsdAmm.Contract.HasLiquidity(&_CrvUsdAmm.CallOpts, user)
}

// HasLiquidity is a free data retrieval call binding the contract method 0xe8dd1ef1.
//
// Solidity: function has_liquidity(address user) view returns(bool)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) HasLiquidity(user common.Address) (bool, error) {
	return _CrvUsdAmm.Contract.HasLiquidity(&_CrvUsdAmm.CallOpts, user)
}

// LiquidityMiningCallback is a free data retrieval call binding the contract method 0x611105d3.
//
// Solidity: function liquidity_mining_callback() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCaller) LiquidityMiningCallback(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "liquidity_mining_callback")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityMiningCallback is a free data retrieval call binding the contract method 0x611105d3.
//
// Solidity: function liquidity_mining_callback() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmSession) LiquidityMiningCallback() (common.Address, error) {
	return _CrvUsdAmm.Contract.LiquidityMiningCallback(&_CrvUsdAmm.CallOpts)
}

// LiquidityMiningCallback is a free data retrieval call binding the contract method 0x611105d3.
//
// Solidity: function liquidity_mining_callback() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) LiquidityMiningCallback() (common.Address, error) {
	return _CrvUsdAmm.Contract.LiquidityMiningCallback(&_CrvUsdAmm.CallOpts)
}

// MaxBand is a free data retrieval call binding the contract method 0xaaa615fc.
//
// Solidity: function max_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCaller) MaxBand(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "max_band")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBand is a free data retrieval call binding the contract method 0xaaa615fc.
//
// Solidity: function max_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmSession) MaxBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.MaxBand(&_CrvUsdAmm.CallOpts)
}

// MaxBand is a free data retrieval call binding the contract method 0xaaa615fc.
//
// Solidity: function max_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) MaxBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.MaxBand(&_CrvUsdAmm.CallOpts)
}

// MinBand is a free data retrieval call binding the contract method 0xca72a821.
//
// Solidity: function min_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCaller) MinBand(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "min_band")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBand is a free data retrieval call binding the contract method 0xca72a821.
//
// Solidity: function min_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmSession) MinBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.MinBand(&_CrvUsdAmm.CallOpts)
}

// MinBand is a free data retrieval call binding the contract method 0xca72a821.
//
// Solidity: function min_band() view returns(int256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) MinBand() (*big.Int, error) {
	return _CrvUsdAmm.Contract.MinBand(&_CrvUsdAmm.CallOpts)
}

// PCurrentDown is a free data retrieval call binding the contract method 0xc32bd03c.
//
// Solidity: function p_current_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) PCurrentDown(opts *bind.CallOpts, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "p_current_down", n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PCurrentDown is a free data retrieval call binding the contract method 0xc32bd03c.
//
// Solidity: function p_current_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) PCurrentDown(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.PCurrentDown(&_CrvUsdAmm.CallOpts, n)
}

// PCurrentDown is a free data retrieval call binding the contract method 0xc32bd03c.
//
// Solidity: function p_current_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) PCurrentDown(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.PCurrentDown(&_CrvUsdAmm.CallOpts, n)
}

// PCurrentUp is a free data retrieval call binding the contract method 0x7c1bbd83.
//
// Solidity: function p_current_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) PCurrentUp(opts *bind.CallOpts, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "p_current_up", n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PCurrentUp is a free data retrieval call binding the contract method 0x7c1bbd83.
//
// Solidity: function p_current_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) PCurrentUp(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.PCurrentUp(&_CrvUsdAmm.CallOpts, n)
}

// PCurrentUp is a free data retrieval call binding the contract method 0x7c1bbd83.
//
// Solidity: function p_current_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) PCurrentUp(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.PCurrentUp(&_CrvUsdAmm.CallOpts, n)
}

// POracleDown is a free data retrieval call binding the contract method 0x24299b7a.
//
// Solidity: function p_oracle_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) POracleDown(opts *bind.CallOpts, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "p_oracle_down", n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POracleDown is a free data retrieval call binding the contract method 0x24299b7a.
//
// Solidity: function p_oracle_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) POracleDown(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.POracleDown(&_CrvUsdAmm.CallOpts, n)
}

// POracleDown is a free data retrieval call binding the contract method 0x24299b7a.
//
// Solidity: function p_oracle_down(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) POracleDown(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.POracleDown(&_CrvUsdAmm.CallOpts, n)
}

// POracleUp is a free data retrieval call binding the contract method 0x2eb858e7.
//
// Solidity: function p_oracle_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) POracleUp(opts *bind.CallOpts, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "p_oracle_up", n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POracleUp is a free data retrieval call binding the contract method 0x2eb858e7.
//
// Solidity: function p_oracle_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) POracleUp(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.POracleUp(&_CrvUsdAmm.CallOpts, n)
}

// POracleUp is a free data retrieval call binding the contract method 0x2eb858e7.
//
// Solidity: function p_oracle_up(int256 n) view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) POracleUp(n *big.Int) (*big.Int, error) {
	return _CrvUsdAmm.Contract.POracleUp(&_CrvUsdAmm.CallOpts, n)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) PriceOracle() (*big.Int, error) {
	return _CrvUsdAmm.Contract.PriceOracle(&_CrvUsdAmm.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) PriceOracle() (*big.Int, error) {
	return _CrvUsdAmm.Contract.PriceOracle(&_CrvUsdAmm.CallOpts)
}

// PriceOracleContract is a free data retrieval call binding the contract method 0x5ea0e01b.
//
// Solidity: function price_oracle_contract() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCaller) PriceOracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "price_oracle_contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracleContract is a free data retrieval call binding the contract method 0x5ea0e01b.
//
// Solidity: function price_oracle_contract() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmSession) PriceOracleContract() (common.Address, error) {
	return _CrvUsdAmm.Contract.PriceOracleContract(&_CrvUsdAmm.CallOpts)
}

// PriceOracleContract is a free data retrieval call binding the contract method 0x5ea0e01b.
//
// Solidity: function price_oracle_contract() view returns(address)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) PriceOracleContract() (common.Address, error) {
	return _CrvUsdAmm.Contract.PriceOracleContract(&_CrvUsdAmm.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) Rate() (*big.Int, error) {
	return _CrvUsdAmm.Contract.Rate(&_CrvUsdAmm.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmCallerSession) Rate() (*big.Int, error) {
	return _CrvUsdAmm.Contract.Rate(&_CrvUsdAmm.CallOpts)
}

// ReadUserTickNumbers is a free data retrieval call binding the contract method 0xb461100d.
//
// Solidity: function read_user_tick_numbers(address user) view returns(int256[2])
func (_CrvUsdAmm *CrvUsdAmmCaller) ReadUserTickNumbers(opts *bind.CallOpts, user common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdAmm.contract.Call(opts, &out, "read_user_tick_numbers", user)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// ReadUserTickNumbers is a free data retrieval call binding the contract method 0xb461100d.
//
// Solidity: function read_user_tick_numbers(address user) view returns(int256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) ReadUserTickNumbers(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdAmm.Contract.ReadUserTickNumbers(&_CrvUsdAmm.CallOpts, user)
}

// ReadUserTickNumbers is a free data retrieval call binding the contract method 0xb461100d.
//
// Solidity: function read_user_tick_numbers(address user) view returns(int256[2])
func (_CrvUsdAmm *CrvUsdAmmCallerSession) ReadUserTickNumbers(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdAmm.Contract.ReadUserTickNumbers(&_CrvUsdAmm.CallOpts, user)
}

// DepositRange is a paid mutator transaction binding the contract method 0xab047e00.
//
// Solidity: function deposit_range(address user, uint256 amount, int256 n1, int256 n2) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) DepositRange(opts *bind.TransactOpts, user common.Address, amount *big.Int, n1 *big.Int, n2 *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "deposit_range", user, amount, n1, n2)
}

// DepositRange is a paid mutator transaction binding the contract method 0xab047e00.
//
// Solidity: function deposit_range(address user, uint256 amount, int256 n1, int256 n2) returns()
func (_CrvUsdAmm *CrvUsdAmmSession) DepositRange(user common.Address, amount *big.Int, n1 *big.Int, n2 *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.DepositRange(&_CrvUsdAmm.TransactOpts, user, amount, n1, n2)
}

// DepositRange is a paid mutator transaction binding the contract method 0xab047e00.
//
// Solidity: function deposit_range(address user, uint256 amount, int256 n1, int256 n2) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) DepositRange(user common.Address, amount *big.Int, n1 *big.Int, n2 *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.DepositRange(&_CrvUsdAmm.TransactOpts, user, amount, n1, n2)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "exchange", i, j, in_amount, min_amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) Exchange(i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Exchange(&_CrvUsdAmm.TransactOpts, i, j, in_amount, min_amount)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) Exchange(i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Exchange(&_CrvUsdAmm.TransactOpts, i, j, in_amount, min_amount)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "exchange0", i, j, in_amount, min_amount, _for)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) Exchange0(i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Exchange0(&_CrvUsdAmm.TransactOpts, i, j, in_amount, min_amount, _for)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xa64833a0.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 in_amount, uint256 min_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) Exchange0(i *big.Int, j *big.Int, in_amount *big.Int, min_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Exchange0(&_CrvUsdAmm.TransactOpts, i, j, in_amount, min_amount, _for)
}

// ExchangeDy is a paid mutator transaction binding the contract method 0xa3e346ec.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactor) ExchangeDy(opts *bind.TransactOpts, i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "exchange_dy", i, j, out_amount, max_amount)
}

// ExchangeDy is a paid mutator transaction binding the contract method 0xa3e346ec.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) ExchangeDy(i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ExchangeDy(&_CrvUsdAmm.TransactOpts, i, j, out_amount, max_amount)
}

// ExchangeDy is a paid mutator transaction binding the contract method 0xa3e346ec.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) ExchangeDy(i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ExchangeDy(&_CrvUsdAmm.TransactOpts, i, j, out_amount, max_amount)
}

// ExchangeDy0 is a paid mutator transaction binding the contract method 0x3c10269a.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactor) ExchangeDy0(opts *bind.TransactOpts, i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "exchange_dy0", i, j, out_amount, max_amount, _for)
}

// ExchangeDy0 is a paid mutator transaction binding the contract method 0x3c10269a.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) ExchangeDy0(i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ExchangeDy0(&_CrvUsdAmm.TransactOpts, i, j, out_amount, max_amount, _for)
}

// ExchangeDy0 is a paid mutator transaction binding the contract method 0x3c10269a.
//
// Solidity: function exchange_dy(uint256 i, uint256 j, uint256 out_amount, uint256 max_amount, address _for) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) ExchangeDy0(i *big.Int, j *big.Int, out_amount *big.Int, max_amount *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ExchangeDy0(&_CrvUsdAmm.TransactOpts, i, j, out_amount, max_amount, _for)
}

// ResetAdminFees is a paid mutator transaction binding the contract method 0x822fe507.
//
// Solidity: function reset_admin_fees() returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) ResetAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "reset_admin_fees")
}

// ResetAdminFees is a paid mutator transaction binding the contract method 0x822fe507.
//
// Solidity: function reset_admin_fees() returns()
func (_CrvUsdAmm *CrvUsdAmmSession) ResetAdminFees() (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ResetAdminFees(&_CrvUsdAmm.TransactOpts)
}

// ResetAdminFees is a paid mutator transaction binding the contract method 0x822fe507.
//
// Solidity: function reset_admin_fees() returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) ResetAdminFees() (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.ResetAdminFees(&_CrvUsdAmm.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdAmm *CrvUsdAmmSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetAdmin(&_CrvUsdAmm.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetAdmin(&_CrvUsdAmm.TransactOpts, _admin)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x3217902f.
//
// Solidity: function set_admin_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) SetAdminFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "set_admin_fee", fee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x3217902f.
//
// Solidity: function set_admin_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmSession) SetAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetAdminFee(&_CrvUsdAmm.TransactOpts, fee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x3217902f.
//
// Solidity: function set_admin_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) SetAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetAdminFee(&_CrvUsdAmm.TransactOpts, fee)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address liquidity_mining_callback) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) SetCallback(opts *bind.TransactOpts, liquidity_mining_callback common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "set_callback", liquidity_mining_callback)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address liquidity_mining_callback) returns()
func (_CrvUsdAmm *CrvUsdAmmSession) SetCallback(liquidity_mining_callback common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetCallback(&_CrvUsdAmm.TransactOpts, liquidity_mining_callback)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address liquidity_mining_callback) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) SetCallback(liquidity_mining_callback common.Address) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetCallback(&_CrvUsdAmm.TransactOpts, liquidity_mining_callback)
}

// SetFee is a paid mutator transaction binding the contract method 0x1aa02d59.
//
// Solidity: function set_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactor) SetFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "set_fee", fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1aa02d59.
//
// Solidity: function set_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetFee(&_CrvUsdAmm.TransactOpts, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1aa02d59.
//
// Solidity: function set_fee(uint256 fee) returns()
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) SetFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetFee(&_CrvUsdAmm.TransactOpts, fee)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmTransactor) SetRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "set_rate", rate)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmSession) SetRate(rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetRate(&_CrvUsdAmm.TransactOpts, rate)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns(uint256)
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) SetRate(rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.SetRate(&_CrvUsdAmm.TransactOpts, rate)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address user, uint256 frac) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactor) Withdraw(opts *bind.TransactOpts, user common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.contract.Transact(opts, "withdraw", user, frac)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address user, uint256 frac) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmSession) Withdraw(user common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Withdraw(&_CrvUsdAmm.TransactOpts, user, frac)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address user, uint256 frac) returns(uint256[2])
func (_CrvUsdAmm *CrvUsdAmmTransactorSession) Withdraw(user common.Address, frac *big.Int) (*types.Transaction, error) {
	return _CrvUsdAmm.Contract.Withdraw(&_CrvUsdAmm.TransactOpts, user, frac)
}

// CrvUsdAmmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CrvUsdAmm contract.
type CrvUsdAmmDepositIterator struct {
	Event *CrvUsdAmmDeposit // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmDeposit)
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
		it.Event = new(CrvUsdAmmDeposit)
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
func (it *CrvUsdAmmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmDeposit represents a Deposit event raised by the CrvUsdAmm contract.
type CrvUsdAmmDeposit struct {
	Provider common.Address
	Amount   *big.Int
	N1       *big.Int
	N2       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7e4f5fadb3361b33669433b392d1a203b7a236710eb272650052592e6ce62f09.
//
// Solidity: event Deposit(address indexed provider, uint256 amount, int256 n1, int256 n2)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address) (*CrvUsdAmmDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmDepositIterator{contract: _CrvUsdAmm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7e4f5fadb3361b33669433b392d1a203b7a236710eb272650052592e6ce62f09.
//
// Solidity: event Deposit(address indexed provider, uint256 amount, int256 n1, int256 n2)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmDeposit, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmDeposit)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7e4f5fadb3361b33669433b392d1a203b7a236710eb272650052592e6ce62f09.
//
// Solidity: event Deposit(address indexed provider, uint256 amount, int256 n1, int256 n2)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseDeposit(log types.Log) (*CrvUsdAmmDeposit, error) {
	event := new(CrvUsdAmmDeposit)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdAmmSetAdminFeeIterator is returned from FilterSetAdminFee and is used to iterate over the raw logs and unpacked data for SetAdminFee events raised by the CrvUsdAmm contract.
type CrvUsdAmmSetAdminFeeIterator struct {
	Event *CrvUsdAmmSetAdminFee // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmSetAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmSetAdminFee)
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
		it.Event = new(CrvUsdAmmSetAdminFee)
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
func (it *CrvUsdAmmSetAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmSetAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmSetAdminFee represents a SetAdminFee event raised by the CrvUsdAmm contract.
type CrvUsdAmmSetAdminFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetAdminFee is a free log retrieval operation binding the contract event 0x2f0d0ace1d699b471d7b39522b5c8aae053bce1b422b7a4fe8f09bd6562a4b74.
//
// Solidity: event SetAdminFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterSetAdminFee(opts *bind.FilterOpts) (*CrvUsdAmmSetAdminFeeIterator, error) {

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "SetAdminFee")
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmSetAdminFeeIterator{contract: _CrvUsdAmm.contract, event: "SetAdminFee", logs: logs, sub: sub}, nil
}

// WatchSetAdminFee is a free log subscription operation binding the contract event 0x2f0d0ace1d699b471d7b39522b5c8aae053bce1b422b7a4fe8f09bd6562a4b74.
//
// Solidity: event SetAdminFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchSetAdminFee(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmSetAdminFee) (event.Subscription, error) {

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "SetAdminFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmSetAdminFee)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "SetAdminFee", log); err != nil {
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

// ParseSetAdminFee is a log parse operation binding the contract event 0x2f0d0ace1d699b471d7b39522b5c8aae053bce1b422b7a4fe8f09bd6562a4b74.
//
// Solidity: event SetAdminFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseSetAdminFee(log types.Log) (*CrvUsdAmmSetAdminFee, error) {
	event := new(CrvUsdAmmSetAdminFee)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "SetAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdAmmSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the CrvUsdAmm contract.
type CrvUsdAmmSetFeeIterator struct {
	Event *CrvUsdAmmSetFee // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmSetFee)
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
		it.Event = new(CrvUsdAmmSetFee)
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
func (it *CrvUsdAmmSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmSetFee represents a SetFee event raised by the CrvUsdAmm contract.
type CrvUsdAmmSetFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x00172ddfc5ae88d08b3de01a5a187667c37a5a53989e8c175055cb6c993792a7.
//
// Solidity: event SetFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterSetFee(opts *bind.FilterOpts) (*CrvUsdAmmSetFeeIterator, error) {

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmSetFeeIterator{contract: _CrvUsdAmm.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x00172ddfc5ae88d08b3de01a5a187667c37a5a53989e8c175055cb6c993792a7.
//
// Solidity: event SetFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmSetFee) (event.Subscription, error) {

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmSetFee)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x00172ddfc5ae88d08b3de01a5a187667c37a5a53989e8c175055cb6c993792a7.
//
// Solidity: event SetFee(uint256 fee)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseSetFee(log types.Log) (*CrvUsdAmmSetFee, error) {
	event := new(CrvUsdAmmSetFee)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdAmmSetRateIterator is returned from FilterSetRate and is used to iterate over the raw logs and unpacked data for SetRate events raised by the CrvUsdAmm contract.
type CrvUsdAmmSetRateIterator struct {
	Event *CrvUsdAmmSetRate // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmSetRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmSetRate)
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
		it.Event = new(CrvUsdAmmSetRate)
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
func (it *CrvUsdAmmSetRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmSetRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmSetRate represents a SetRate event raised by the CrvUsdAmm contract.
type CrvUsdAmmSetRate struct {
	Rate    *big.Int
	RateMul *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetRate is a free log retrieval operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 rate, uint256 rate_mul, uint256 time)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterSetRate(opts *bind.FilterOpts) (*CrvUsdAmmSetRateIterator, error) {

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmSetRateIterator{contract: _CrvUsdAmm.contract, event: "SetRate", logs: logs, sub: sub}, nil
}

// WatchSetRate is a free log subscription operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 rate, uint256 rate_mul, uint256 time)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchSetRate(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmSetRate) (event.Subscription, error) {

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmSetRate)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "SetRate", log); err != nil {
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

// ParseSetRate is a log parse operation binding the contract event 0x52543716810f73c3fa9bca74622aecb6d3614ca4991472f3e999d531c2f6afb8.
//
// Solidity: event SetRate(uint256 rate, uint256 rate_mul, uint256 time)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseSetRate(log types.Log) (*CrvUsdAmmSetRate, error) {
	event := new(CrvUsdAmmSetRate)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "SetRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdAmmTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CrvUsdAmm contract.
type CrvUsdAmmTokenExchangeIterator struct {
	Event *CrvUsdAmmTokenExchange // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmTokenExchange)
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
		it.Event = new(CrvUsdAmmTokenExchange)
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
func (it *CrvUsdAmmTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmTokenExchange represents a TokenExchange event raised by the CrvUsdAmm contract.
type CrvUsdAmmTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CrvUsdAmmTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmTokenExchangeIterator{contract: _CrvUsdAmm.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmTokenExchange)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseTokenExchange(log types.Log) (*CrvUsdAmmTokenExchange, error) {
	event := new(CrvUsdAmmTokenExchange)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdAmmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CrvUsdAmm contract.
type CrvUsdAmmWithdrawIterator struct {
	Event *CrvUsdAmmWithdraw // Event containing the contract specifics and raw log

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
func (it *CrvUsdAmmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdAmmWithdraw)
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
		it.Event = new(CrvUsdAmmWithdraw)
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
func (it *CrvUsdAmmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdAmmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdAmmWithdraw represents a Withdraw event raised by the CrvUsdAmm contract.
type CrvUsdAmmWithdraw struct {
	Provider         common.Address
	AmountBorrowed   *big.Int
	AmountCollateral *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 amount_borrowed, uint256 amount_collateral)
func (_CrvUsdAmm *CrvUsdAmmFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*CrvUsdAmmWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdAmmWithdrawIterator{contract: _CrvUsdAmm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 amount_borrowed, uint256 amount_collateral)
func (_CrvUsdAmm *CrvUsdAmmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CrvUsdAmmWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CrvUsdAmm.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdAmmWithdraw)
				if err := _CrvUsdAmm.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 amount_borrowed, uint256 amount_collateral)
func (_CrvUsdAmm *CrvUsdAmmFilterer) ParseWithdraw(log types.Log) (*CrvUsdAmmWithdraw, error) {
	event := new(CrvUsdAmmWithdraw)
	if err := _CrvUsdAmm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
