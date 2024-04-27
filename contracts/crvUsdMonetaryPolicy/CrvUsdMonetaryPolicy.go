// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdMonetaryPolicy

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Candle0   *big.Int
	Candle1   *big.Int
	Timestamp *big.Int
}

// CrvUsdMonetaryPolicyMetaData contains all meta data concerning the CrvUsdMonetaryPolicy contract.
var CrvUsdMonetaryPolicyMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddPegKeeper\",\"inputs\":[{\"name\":\"peg_keeper\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemovePegKeeper\",\"inputs\":[{\"name\":\"peg_keeper\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetRate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetSigma\",\"inputs\":[{\"name\":\"sigma\",\"type\":\"int256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetTargetDebtFraction\",\"inputs\":[{\"name\":\"target_debt_fraction\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"price_oracle\",\"type\":\"address\"},{\"name\":\"controller_factory\",\"type\":\"address\"},{\"name\":\"peg_keepers\",\"type\":\"address[5]\"},{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"sigma\",\"type\":\"int256\"},{\"name\":\"target_debt_fraction\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_peg_keeper\",\"inputs\":[{\"name\":\"pk\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_peg_keeper\",\"inputs\":[{\"name\":\"pk\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rate\",\"inputs\":[{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"rate_write\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"rate_write\",\"inputs\":[{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_rate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_sigma\",\"inputs\":[{\"name\":\"sigma\",\"type\":\"int256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_target_debt_fraction\",\"inputs\":[{\"name\":\"target_debt_fraction\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rate0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"sigma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"target_debt_fraction\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"peg_keepers\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"PRICE_ORACLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"CONTROLLER_FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"n_controllers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controllers\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_debt_candles\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"candle0\",\"type\":\"uint256\"},{\"name\":\"candle1\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}]}]}]",
}

// CrvUsdMonetaryPolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdMonetaryPolicyMetaData.ABI instead.
var CrvUsdMonetaryPolicyABI = CrvUsdMonetaryPolicyMetaData.ABI

// CrvUsdMonetaryPolicy is an auto generated Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicy struct {
	CrvUsdMonetaryPolicyCaller     // Read-only binding to the contract
	CrvUsdMonetaryPolicyTransactor // Write-only binding to the contract
	CrvUsdMonetaryPolicyFilterer   // Log filterer for contract events
}

// CrvUsdMonetaryPolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdMonetaryPolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdMonetaryPolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdMonetaryPolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdMonetaryPolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdMonetaryPolicySession struct {
	Contract     *CrvUsdMonetaryPolicy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrvUsdMonetaryPolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdMonetaryPolicyCallerSession struct {
	Contract *CrvUsdMonetaryPolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CrvUsdMonetaryPolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdMonetaryPolicyTransactorSession struct {
	Contract     *CrvUsdMonetaryPolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CrvUsdMonetaryPolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicyRaw struct {
	Contract *CrvUsdMonetaryPolicy // Generic contract binding to access the raw methods on
}

// CrvUsdMonetaryPolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicyCallerRaw struct {
	Contract *CrvUsdMonetaryPolicyCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdMonetaryPolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdMonetaryPolicyTransactorRaw struct {
	Contract *CrvUsdMonetaryPolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdMonetaryPolicy creates a new instance of CrvUsdMonetaryPolicy, bound to a specific deployed contract.
func NewCrvUsdMonetaryPolicy(address common.Address, backend bind.ContractBackend) (*CrvUsdMonetaryPolicy, error) {
	contract, err := bindCrvUsdMonetaryPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicy{CrvUsdMonetaryPolicyCaller: CrvUsdMonetaryPolicyCaller{contract: contract}, CrvUsdMonetaryPolicyTransactor: CrvUsdMonetaryPolicyTransactor{contract: contract}, CrvUsdMonetaryPolicyFilterer: CrvUsdMonetaryPolicyFilterer{contract: contract}}, nil
}

// NewCrvUsdMonetaryPolicyCaller creates a new read-only instance of CrvUsdMonetaryPolicy, bound to a specific deployed contract.
func NewCrvUsdMonetaryPolicyCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdMonetaryPolicyCaller, error) {
	contract, err := bindCrvUsdMonetaryPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicyCaller{contract: contract}, nil
}

// NewCrvUsdMonetaryPolicyTransactor creates a new write-only instance of CrvUsdMonetaryPolicy, bound to a specific deployed contract.
func NewCrvUsdMonetaryPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdMonetaryPolicyTransactor, error) {
	contract, err := bindCrvUsdMonetaryPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicyTransactor{contract: contract}, nil
}

// NewCrvUsdMonetaryPolicyFilterer creates a new log filterer instance of CrvUsdMonetaryPolicy, bound to a specific deployed contract.
func NewCrvUsdMonetaryPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdMonetaryPolicyFilterer, error) {
	contract, err := bindCrvUsdMonetaryPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicyFilterer{contract: contract}, nil
}

// bindCrvUsdMonetaryPolicy binds a generic wrapper to an already deployed contract.
func bindCrvUsdMonetaryPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdMonetaryPolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdMonetaryPolicy.Contract.CrvUsdMonetaryPolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.CrvUsdMonetaryPolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.CrvUsdMonetaryPolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdMonetaryPolicy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.contract.Transact(opts, method, params...)
}

// CONTROLLERFACTORY is a free data retrieval call binding the contract method 0xa38d1b21.
//
// Solidity: function CONTROLLER_FACTORY() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) CONTROLLERFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "CONTROLLER_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROLLERFACTORY is a free data retrieval call binding the contract method 0xa38d1b21.
//
// Solidity: function CONTROLLER_FACTORY() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) CONTROLLERFACTORY() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.CONTROLLERFACTORY(&_CrvUsdMonetaryPolicy.CallOpts)
}

// CONTROLLERFACTORY is a free data retrieval call binding the contract method 0xa38d1b21.
//
// Solidity: function CONTROLLER_FACTORY() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) CONTROLLERFACTORY() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.CONTROLLERFACTORY(&_CrvUsdMonetaryPolicy.CallOpts)
}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) PRICEORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "PRICE_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) PRICEORACLE() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.PRICEORACLE(&_CrvUsdMonetaryPolicy.CallOpts)
}

// PRICEORACLE is a free data retrieval call binding the contract method 0x0a19399a.
//
// Solidity: function PRICE_ORACLE() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) PRICEORACLE() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.PRICEORACLE(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Admin() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.Admin(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Admin() (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.Admin(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Controllers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.Controllers(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.Controllers(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// MinDebtCandles is a free data retrieval call binding the contract method 0x5db441ed.
//
// Solidity: function min_debt_candles(address arg0) view returns((uint256,uint256,uint256))
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) MinDebtCandles(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "min_debt_candles", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// MinDebtCandles is a free data retrieval call binding the contract method 0x5db441ed.
//
// Solidity: function min_debt_candles(address arg0) view returns((uint256,uint256,uint256))
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) MinDebtCandles(arg0 common.Address) (Struct0, error) {
	return _CrvUsdMonetaryPolicy.Contract.MinDebtCandles(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// MinDebtCandles is a free data retrieval call binding the contract method 0x5db441ed.
//
// Solidity: function min_debt_candles(address arg0) view returns((uint256,uint256,uint256))
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) MinDebtCandles(arg0 common.Address) (Struct0, error) {
	return _CrvUsdMonetaryPolicy.Contract.MinDebtCandles(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) NControllers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "n_controllers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) NControllers() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.NControllers(&_CrvUsdMonetaryPolicy.CallOpts)
}

// NControllers is a free data retrieval call binding the contract method 0x1d741825.
//
// Solidity: function n_controllers() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) NControllers() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.NControllers(&_CrvUsdMonetaryPolicy.CallOpts)
}

// PegKeepers is a free data retrieval call binding the contract method 0xf6235138.
//
// Solidity: function peg_keepers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) PegKeepers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "peg_keepers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegKeepers is a free data retrieval call binding the contract method 0xf6235138.
//
// Solidity: function peg_keepers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) PegKeepers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.PegKeepers(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// PegKeepers is a free data retrieval call binding the contract method 0xf6235138.
//
// Solidity: function peg_keepers(uint256 arg0) view returns(address)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) PegKeepers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdMonetaryPolicy.Contract.PegKeepers(&_CrvUsdMonetaryPolicy.CallOpts, arg0)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Rate() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Rate() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Rate0 is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate(address _for) view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Rate0(opts *bind.CallOpts, _for common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "rate0", _for)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate0 is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate(address _for) view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Rate0(_for common.Address) (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate0(&_CrvUsdMonetaryPolicy.CallOpts, _for)
}

// Rate0 is a free data retrieval call binding the contract method 0x0ba9d8ca.
//
// Solidity: function rate(address _for) view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Rate0(_for common.Address) (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate0(&_CrvUsdMonetaryPolicy.CallOpts, _for)
}

// Rate00 is a free data retrieval call binding the contract method 0x93c19e18.
//
// Solidity: function rate0() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Rate00(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "rate00")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate00 is a free data retrieval call binding the contract method 0x93c19e18.
//
// Solidity: function rate0() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Rate00() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate00(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Rate00 is a free data retrieval call binding the contract method 0x93c19e18.
//
// Solidity: function rate0() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Rate00() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Rate00(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(int256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) Sigma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "sigma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(int256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) Sigma() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Sigma(&_CrvUsdMonetaryPolicy.CallOpts)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(int256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) Sigma() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.Sigma(&_CrvUsdMonetaryPolicy.CallOpts)
}

// TargetDebtFraction is a free data retrieval call binding the contract method 0xa155b53a.
//
// Solidity: function target_debt_fraction() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCaller) TargetDebtFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdMonetaryPolicy.contract.Call(opts, &out, "target_debt_fraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetDebtFraction is a free data retrieval call binding the contract method 0xa155b53a.
//
// Solidity: function target_debt_fraction() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) TargetDebtFraction() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.TargetDebtFraction(&_CrvUsdMonetaryPolicy.CallOpts)
}

// TargetDebtFraction is a free data retrieval call binding the contract method 0xa155b53a.
//
// Solidity: function target_debt_fraction() view returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyCallerSession) TargetDebtFraction() (*big.Int, error) {
	return _CrvUsdMonetaryPolicy.Contract.TargetDebtFraction(&_CrvUsdMonetaryPolicy.CallOpts)
}

// AddPegKeeper is a paid mutator transaction binding the contract method 0x0eb57e3c.
//
// Solidity: function add_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) AddPegKeeper(opts *bind.TransactOpts, pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "add_peg_keeper", pk)
}

// AddPegKeeper is a paid mutator transaction binding the contract method 0x0eb57e3c.
//
// Solidity: function add_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) AddPegKeeper(pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.AddPegKeeper(&_CrvUsdMonetaryPolicy.TransactOpts, pk)
}

// AddPegKeeper is a paid mutator transaction binding the contract method 0x0eb57e3c.
//
// Solidity: function add_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) AddPegKeeper(pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.AddPegKeeper(&_CrvUsdMonetaryPolicy.TransactOpts, pk)
}

// RateWrite is a paid mutator transaction binding the contract method 0xe91f2f4c.
//
// Solidity: function rate_write() returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) RateWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "rate_write")
}

// RateWrite is a paid mutator transaction binding the contract method 0xe91f2f4c.
//
// Solidity: function rate_write() returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) RateWrite() (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RateWrite(&_CrvUsdMonetaryPolicy.TransactOpts)
}

// RateWrite is a paid mutator transaction binding the contract method 0xe91f2f4c.
//
// Solidity: function rate_write() returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) RateWrite() (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RateWrite(&_CrvUsdMonetaryPolicy.TransactOpts)
}

// RateWrite0 is a paid mutator transaction binding the contract method 0xbdb09f2e.
//
// Solidity: function rate_write(address _for) returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) RateWrite0(opts *bind.TransactOpts, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "rate_write0", _for)
}

// RateWrite0 is a paid mutator transaction binding the contract method 0xbdb09f2e.
//
// Solidity: function rate_write(address _for) returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) RateWrite0(_for common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RateWrite0(&_CrvUsdMonetaryPolicy.TransactOpts, _for)
}

// RateWrite0 is a paid mutator transaction binding the contract method 0xbdb09f2e.
//
// Solidity: function rate_write(address _for) returns(uint256)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) RateWrite0(_for common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RateWrite0(&_CrvUsdMonetaryPolicy.TransactOpts, _for)
}

// RemovePegKeeper is a paid mutator transaction binding the contract method 0x55aab44a.
//
// Solidity: function remove_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) RemovePegKeeper(opts *bind.TransactOpts, pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "remove_peg_keeper", pk)
}

// RemovePegKeeper is a paid mutator transaction binding the contract method 0x55aab44a.
//
// Solidity: function remove_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) RemovePegKeeper(pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RemovePegKeeper(&_CrvUsdMonetaryPolicy.TransactOpts, pk)
}

// RemovePegKeeper is a paid mutator transaction binding the contract method 0x55aab44a.
//
// Solidity: function remove_peg_keeper(address pk) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) RemovePegKeeper(pk common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.RemovePegKeeper(&_CrvUsdMonetaryPolicy.TransactOpts, pk)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) SetAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "set_admin", admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetAdmin(&_CrvUsdMonetaryPolicy.TransactOpts, admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetAdmin(&_CrvUsdMonetaryPolicy.TransactOpts, admin)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) SetRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "set_rate", rate)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) SetRate(rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetRate(&_CrvUsdMonetaryPolicy.TransactOpts, rate)
}

// SetRate is a paid mutator transaction binding the contract method 0xd4387a99.
//
// Solidity: function set_rate(uint256 rate) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) SetRate(rate *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetRate(&_CrvUsdMonetaryPolicy.TransactOpts, rate)
}

// SetSigma is a paid mutator transaction binding the contract method 0xbc8f1050.
//
// Solidity: function set_sigma(int256 sigma) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) SetSigma(opts *bind.TransactOpts, sigma *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "set_sigma", sigma)
}

// SetSigma is a paid mutator transaction binding the contract method 0xbc8f1050.
//
// Solidity: function set_sigma(int256 sigma) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) SetSigma(sigma *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetSigma(&_CrvUsdMonetaryPolicy.TransactOpts, sigma)
}

// SetSigma is a paid mutator transaction binding the contract method 0xbc8f1050.
//
// Solidity: function set_sigma(int256 sigma) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) SetSigma(sigma *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetSigma(&_CrvUsdMonetaryPolicy.TransactOpts, sigma)
}

// SetTargetDebtFraction is a paid mutator transaction binding the contract method 0xc18b9396.
//
// Solidity: function set_target_debt_fraction(uint256 target_debt_fraction) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactor) SetTargetDebtFraction(opts *bind.TransactOpts, target_debt_fraction *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.contract.Transact(opts, "set_target_debt_fraction", target_debt_fraction)
}

// SetTargetDebtFraction is a paid mutator transaction binding the contract method 0xc18b9396.
//
// Solidity: function set_target_debt_fraction(uint256 target_debt_fraction) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicySession) SetTargetDebtFraction(target_debt_fraction *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetTargetDebtFraction(&_CrvUsdMonetaryPolicy.TransactOpts, target_debt_fraction)
}

// SetTargetDebtFraction is a paid mutator transaction binding the contract method 0xc18b9396.
//
// Solidity: function set_target_debt_fraction(uint256 target_debt_fraction) returns()
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyTransactorSession) SetTargetDebtFraction(target_debt_fraction *big.Int) (*types.Transaction, error) {
	return _CrvUsdMonetaryPolicy.Contract.SetTargetDebtFraction(&_CrvUsdMonetaryPolicy.TransactOpts, target_debt_fraction)
}

// CrvUsdMonetaryPolicyAddPegKeeperIterator is returned from FilterAddPegKeeper and is used to iterate over the raw logs and unpacked data for AddPegKeeper events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicyAddPegKeeperIterator struct {
	Event *CrvUsdMonetaryPolicyAddPegKeeper // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicyAddPegKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicyAddPegKeeper)
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
		it.Event = new(CrvUsdMonetaryPolicyAddPegKeeper)
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
func (it *CrvUsdMonetaryPolicyAddPegKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicyAddPegKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicyAddPegKeeper represents a AddPegKeeper event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicyAddPegKeeper struct {
	PegKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPegKeeper is a free log retrieval operation binding the contract event 0xf395c3706a8194522b942d1992143a7b60a92a83f99ec30e3833c7630e3c1331.
//
// Solidity: event AddPegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterAddPegKeeper(opts *bind.FilterOpts, peg_keeper []common.Address) (*CrvUsdMonetaryPolicyAddPegKeeperIterator, error) {

	var peg_keeperRule []interface{}
	for _, peg_keeperItem := range peg_keeper {
		peg_keeperRule = append(peg_keeperRule, peg_keeperItem)
	}

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "AddPegKeeper", peg_keeperRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicyAddPegKeeperIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "AddPegKeeper", logs: logs, sub: sub}, nil
}

// WatchAddPegKeeper is a free log subscription operation binding the contract event 0xf395c3706a8194522b942d1992143a7b60a92a83f99ec30e3833c7630e3c1331.
//
// Solidity: event AddPegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchAddPegKeeper(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicyAddPegKeeper, peg_keeper []common.Address) (event.Subscription, error) {

	var peg_keeperRule []interface{}
	for _, peg_keeperItem := range peg_keeper {
		peg_keeperRule = append(peg_keeperRule, peg_keeperItem)
	}

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "AddPegKeeper", peg_keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicyAddPegKeeper)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "AddPegKeeper", log); err != nil {
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

// ParseAddPegKeeper is a log parse operation binding the contract event 0xf395c3706a8194522b942d1992143a7b60a92a83f99ec30e3833c7630e3c1331.
//
// Solidity: event AddPegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseAddPegKeeper(log types.Log) (*CrvUsdMonetaryPolicyAddPegKeeper, error) {
	event := new(CrvUsdMonetaryPolicyAddPegKeeper)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "AddPegKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdMonetaryPolicyRemovePegKeeperIterator is returned from FilterRemovePegKeeper and is used to iterate over the raw logs and unpacked data for RemovePegKeeper events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicyRemovePegKeeperIterator struct {
	Event *CrvUsdMonetaryPolicyRemovePegKeeper // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicyRemovePegKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicyRemovePegKeeper)
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
		it.Event = new(CrvUsdMonetaryPolicyRemovePegKeeper)
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
func (it *CrvUsdMonetaryPolicyRemovePegKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicyRemovePegKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicyRemovePegKeeper represents a RemovePegKeeper event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicyRemovePegKeeper struct {
	PegKeeper common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovePegKeeper is a free log retrieval operation binding the contract event 0x52182c3057b74a074adcacf89ba9ff9860a1265c89cfecd998a111e06bc80267.
//
// Solidity: event RemovePegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterRemovePegKeeper(opts *bind.FilterOpts, peg_keeper []common.Address) (*CrvUsdMonetaryPolicyRemovePegKeeperIterator, error) {

	var peg_keeperRule []interface{}
	for _, peg_keeperItem := range peg_keeper {
		peg_keeperRule = append(peg_keeperRule, peg_keeperItem)
	}

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "RemovePegKeeper", peg_keeperRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicyRemovePegKeeperIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "RemovePegKeeper", logs: logs, sub: sub}, nil
}

// WatchRemovePegKeeper is a free log subscription operation binding the contract event 0x52182c3057b74a074adcacf89ba9ff9860a1265c89cfecd998a111e06bc80267.
//
// Solidity: event RemovePegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchRemovePegKeeper(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicyRemovePegKeeper, peg_keeper []common.Address) (event.Subscription, error) {

	var peg_keeperRule []interface{}
	for _, peg_keeperItem := range peg_keeper {
		peg_keeperRule = append(peg_keeperRule, peg_keeperItem)
	}

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "RemovePegKeeper", peg_keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicyRemovePegKeeper)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "RemovePegKeeper", log); err != nil {
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

// ParseRemovePegKeeper is a log parse operation binding the contract event 0x52182c3057b74a074adcacf89ba9ff9860a1265c89cfecd998a111e06bc80267.
//
// Solidity: event RemovePegKeeper(address indexed peg_keeper)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseRemovePegKeeper(log types.Log) (*CrvUsdMonetaryPolicyRemovePegKeeper, error) {
	event := new(CrvUsdMonetaryPolicyRemovePegKeeper)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "RemovePegKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdMonetaryPolicySetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetAdminIterator struct {
	Event *CrvUsdMonetaryPolicySetAdmin // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicySetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicySetAdmin)
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
		it.Event = new(CrvUsdMonetaryPolicySetAdmin)
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
func (it *CrvUsdMonetaryPolicySetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicySetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicySetAdmin represents a SetAdmin event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*CrvUsdMonetaryPolicySetAdminIterator, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicySetAdminIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicySetAdmin) (event.Subscription, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicySetAdmin)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseSetAdmin(log types.Log) (*CrvUsdMonetaryPolicySetAdmin, error) {
	event := new(CrvUsdMonetaryPolicySetAdmin)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdMonetaryPolicySetRateIterator is returned from FilterSetRate and is used to iterate over the raw logs and unpacked data for SetRate events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetRateIterator struct {
	Event *CrvUsdMonetaryPolicySetRate // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicySetRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicySetRate)
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
		it.Event = new(CrvUsdMonetaryPolicySetRate)
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
func (it *CrvUsdMonetaryPolicySetRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicySetRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicySetRate represents a SetRate event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetRate struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRate is a free log retrieval operation binding the contract event 0x2640b4015d3473fd09bf2b30939e17deb4068cdacf3892136e737e166ceb3210.
//
// Solidity: event SetRate(uint256 rate)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterSetRate(opts *bind.FilterOpts) (*CrvUsdMonetaryPolicySetRateIterator, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicySetRateIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "SetRate", logs: logs, sub: sub}, nil
}

// WatchSetRate is a free log subscription operation binding the contract event 0x2640b4015d3473fd09bf2b30939e17deb4068cdacf3892136e737e166ceb3210.
//
// Solidity: event SetRate(uint256 rate)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchSetRate(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicySetRate) (event.Subscription, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "SetRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicySetRate)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetRate", log); err != nil {
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

// ParseSetRate is a log parse operation binding the contract event 0x2640b4015d3473fd09bf2b30939e17deb4068cdacf3892136e737e166ceb3210.
//
// Solidity: event SetRate(uint256 rate)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseSetRate(log types.Log) (*CrvUsdMonetaryPolicySetRate, error) {
	event := new(CrvUsdMonetaryPolicySetRate)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdMonetaryPolicySetSigmaIterator is returned from FilterSetSigma and is used to iterate over the raw logs and unpacked data for SetSigma events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetSigmaIterator struct {
	Event *CrvUsdMonetaryPolicySetSigma // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicySetSigmaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicySetSigma)
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
		it.Event = new(CrvUsdMonetaryPolicySetSigma)
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
func (it *CrvUsdMonetaryPolicySetSigmaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicySetSigmaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicySetSigma represents a SetSigma event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetSigma struct {
	Sigma *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetSigma is a free log retrieval operation binding the contract event 0xdd7b7d5ad5206db1e2995b71e974a52e324ffefe428b26c79f323b814de1d155.
//
// Solidity: event SetSigma(int256 sigma)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterSetSigma(opts *bind.FilterOpts) (*CrvUsdMonetaryPolicySetSigmaIterator, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "SetSigma")
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicySetSigmaIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "SetSigma", logs: logs, sub: sub}, nil
}

// WatchSetSigma is a free log subscription operation binding the contract event 0xdd7b7d5ad5206db1e2995b71e974a52e324ffefe428b26c79f323b814de1d155.
//
// Solidity: event SetSigma(int256 sigma)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchSetSigma(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicySetSigma) (event.Subscription, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "SetSigma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicySetSigma)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetSigma", log); err != nil {
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

// ParseSetSigma is a log parse operation binding the contract event 0xdd7b7d5ad5206db1e2995b71e974a52e324ffefe428b26c79f323b814de1d155.
//
// Solidity: event SetSigma(int256 sigma)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseSetSigma(log types.Log) (*CrvUsdMonetaryPolicySetSigma, error) {
	event := new(CrvUsdMonetaryPolicySetSigma)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetSigma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdMonetaryPolicySetTargetDebtFractionIterator is returned from FilterSetTargetDebtFraction and is used to iterate over the raw logs and unpacked data for SetTargetDebtFraction events raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetTargetDebtFractionIterator struct {
	Event *CrvUsdMonetaryPolicySetTargetDebtFraction // Event containing the contract specifics and raw log

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
func (it *CrvUsdMonetaryPolicySetTargetDebtFractionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdMonetaryPolicySetTargetDebtFraction)
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
		it.Event = new(CrvUsdMonetaryPolicySetTargetDebtFraction)
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
func (it *CrvUsdMonetaryPolicySetTargetDebtFractionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdMonetaryPolicySetTargetDebtFractionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdMonetaryPolicySetTargetDebtFraction represents a SetTargetDebtFraction event raised by the CrvUsdMonetaryPolicy contract.
type CrvUsdMonetaryPolicySetTargetDebtFraction struct {
	TargetDebtFraction *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetTargetDebtFraction is a free log retrieval operation binding the contract event 0x282d6b198c8375838e9528471dafbdbbee8d421fb7fa793395872d42da32fa2b.
//
// Solidity: event SetTargetDebtFraction(uint256 target_debt_fraction)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) FilterSetTargetDebtFraction(opts *bind.FilterOpts) (*CrvUsdMonetaryPolicySetTargetDebtFractionIterator, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.FilterLogs(opts, "SetTargetDebtFraction")
	if err != nil {
		return nil, err
	}
	return &CrvUsdMonetaryPolicySetTargetDebtFractionIterator{contract: _CrvUsdMonetaryPolicy.contract, event: "SetTargetDebtFraction", logs: logs, sub: sub}, nil
}

// WatchSetTargetDebtFraction is a free log subscription operation binding the contract event 0x282d6b198c8375838e9528471dafbdbbee8d421fb7fa793395872d42da32fa2b.
//
// Solidity: event SetTargetDebtFraction(uint256 target_debt_fraction)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) WatchSetTargetDebtFraction(opts *bind.WatchOpts, sink chan<- *CrvUsdMonetaryPolicySetTargetDebtFraction) (event.Subscription, error) {

	logs, sub, err := _CrvUsdMonetaryPolicy.contract.WatchLogs(opts, "SetTargetDebtFraction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdMonetaryPolicySetTargetDebtFraction)
				if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetTargetDebtFraction", log); err != nil {
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

// ParseSetTargetDebtFraction is a log parse operation binding the contract event 0x282d6b198c8375838e9528471dafbdbbee8d421fb7fa793395872d42da32fa2b.
//
// Solidity: event SetTargetDebtFraction(uint256 target_debt_fraction)
func (_CrvUsdMonetaryPolicy *CrvUsdMonetaryPolicyFilterer) ParseSetTargetDebtFraction(log types.Log) (*CrvUsdMonetaryPolicySetTargetDebtFraction, error) {
	event := new(CrvUsdMonetaryPolicySetTargetDebtFraction)
	if err := _CrvUsdMonetaryPolicy.contract.UnpackLog(event, "SetTargetDebtFraction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
