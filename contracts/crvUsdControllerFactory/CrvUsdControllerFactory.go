// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdControllerFactory

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

// CrvUsdControllerFactoryMetaData contains all meta data concerning the CrvUsdControllerFactory contract.
var CrvUsdControllerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"AddMarket\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false},{\"name\":\"amm\",\"type\":\"address\",\"indexed\":false},{\"name\":\"monetary_policy\",\"type\":\"address\",\"indexed\":false},{\"name\":\"ix\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetDebtCeiling\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true},{\"name\":\"debt_ceiling\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MintForMarket\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveFromMarket\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetImplementations\",\"inputs\":[{\"name\":\"amm\",\"type\":\"address\",\"indexed\":false},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetFeeReceiver\",\"inputs\":[{\"name\":\"fee_receiver\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\"},{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"fee_receiver\",\"type\":\"address\"},{\"name\":\"weth\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stablecoin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_market\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"_price_oracle_contract\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"debt_ceiling\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_debt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_controller\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_controller\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_amm\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_amm\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_implementations\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"amm\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_debt_ceiling\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"debt_ceiling\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"rug_debt_ceiling\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"collect_fees_above_ceiling\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controllers\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amms\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controller_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"n_collaterals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"collaterals\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"collaterals_index\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debt_ceiling\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debt_ceiling_residual\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"WETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CrvUsdControllerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdControllerFactoryMetaData.ABI instead.
var CrvUsdControllerFactoryABI = CrvUsdControllerFactoryMetaData.ABI

// CrvUsdControllerFactory is an auto generated Go binding around an Ethereum contract.
type CrvUsdControllerFactory struct {
	CrvUsdControllerFactoryCaller     // Read-only binding to the contract
	CrvUsdControllerFactoryTransactor // Write-only binding to the contract
	CrvUsdControllerFactoryFilterer   // Log filterer for contract events
}

// CrvUsdControllerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdControllerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdControllerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdControllerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdControllerFactorySession struct {
	Contract     *CrvUsdControllerFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrvUsdControllerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdControllerFactoryCallerSession struct {
	Contract *CrvUsdControllerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// CrvUsdControllerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdControllerFactoryTransactorSession struct {
	Contract     *CrvUsdControllerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// CrvUsdControllerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdControllerFactoryRaw struct {
	Contract *CrvUsdControllerFactory // Generic contract binding to access the raw methods on
}

// CrvUsdControllerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdControllerFactoryCallerRaw struct {
	Contract *CrvUsdControllerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdControllerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdControllerFactoryTransactorRaw struct {
	Contract *CrvUsdControllerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdControllerFactory creates a new instance of CrvUsdControllerFactory, bound to a specific deployed contract.
func NewCrvUsdControllerFactory(address common.Address, backend bind.ContractBackend) (*CrvUsdControllerFactory, error) {
	contract, err := bindCrvUsdControllerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactory{CrvUsdControllerFactoryCaller: CrvUsdControllerFactoryCaller{contract: contract}, CrvUsdControllerFactoryTransactor: CrvUsdControllerFactoryTransactor{contract: contract}, CrvUsdControllerFactoryFilterer: CrvUsdControllerFactoryFilterer{contract: contract}}, nil
}

// NewCrvUsdControllerFactoryCaller creates a new read-only instance of CrvUsdControllerFactory, bound to a specific deployed contract.
func NewCrvUsdControllerFactoryCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdControllerFactoryCaller, error) {
	contract, err := bindCrvUsdControllerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryCaller{contract: contract}, nil
}

// NewCrvUsdControllerFactoryTransactor creates a new write-only instance of CrvUsdControllerFactory, bound to a specific deployed contract.
func NewCrvUsdControllerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdControllerFactoryTransactor, error) {
	contract, err := bindCrvUsdControllerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryTransactor{contract: contract}, nil
}

// NewCrvUsdControllerFactoryFilterer creates a new log filterer instance of CrvUsdControllerFactory, bound to a specific deployed contract.
func NewCrvUsdControllerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdControllerFactoryFilterer, error) {
	contract, err := bindCrvUsdControllerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryFilterer{contract: contract}, nil
}

// bindCrvUsdControllerFactory binds a generic wrapper to an already deployed contract.
func bindCrvUsdControllerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdControllerFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdControllerFactory.Contract.CrvUsdControllerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.CrvUsdControllerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.CrvUsdControllerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdControllerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) WETH() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.WETH(&_CrvUsdControllerFactory.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) WETH() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.WETH(&_CrvUsdControllerFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) Admin() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Admin(&_CrvUsdControllerFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) Admin() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Admin(&_CrvUsdControllerFactory.CallOpts)
}

// AmmImplementation is a free data retrieval call binding the contract method 0x3bc5799f.
//
// Solidity: function amm_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) AmmImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "amm_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmImplementation is a free data retrieval call binding the contract method 0x3bc5799f.
//
// Solidity: function amm_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) AmmImplementation() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.AmmImplementation(&_CrvUsdControllerFactory.CallOpts)
}

// AmmImplementation is a free data retrieval call binding the contract method 0x3bc5799f.
//
// Solidity: function amm_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) AmmImplementation() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.AmmImplementation(&_CrvUsdControllerFactory.CallOpts)
}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) Amms(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "amms", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) Amms(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Amms(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) Amms(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Amms(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0x24c1173b.
//
// Solidity: function collaterals(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) Collaterals(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "collaterals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collaterals is a free data retrieval call binding the contract method 0x24c1173b.
//
// Solidity: function collaterals(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) Collaterals(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Collaterals(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0x24c1173b.
//
// Solidity: function collaterals(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) Collaterals(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Collaterals(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// CollateralsIndex is a free data retrieval call binding the contract method 0xa7367268.
//
// Solidity: function collaterals_index(address arg0, uint256 arg1) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) CollateralsIndex(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "collaterals_index", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralsIndex is a free data retrieval call binding the contract method 0xa7367268.
//
// Solidity: function collaterals_index(address arg0, uint256 arg1) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) CollateralsIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.CollateralsIndex(&_CrvUsdControllerFactory.CallOpts, arg0, arg1)
}

// CollateralsIndex is a free data retrieval call binding the contract method 0xa7367268.
//
// Solidity: function collaterals_index(address arg0, uint256 arg1) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) CollateralsIndex(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.CollateralsIndex(&_CrvUsdControllerFactory.CallOpts, arg0, arg1)
}

// ControllerImplementation is a free data retrieval call binding the contract method 0x0275417e.
//
// Solidity: function controller_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) ControllerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "controller_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ControllerImplementation is a free data retrieval call binding the contract method 0x0275417e.
//
// Solidity: function controller_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) ControllerImplementation() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.ControllerImplementation(&_CrvUsdControllerFactory.CallOpts)
}

// ControllerImplementation is a free data retrieval call binding the contract method 0x0275417e.
//
// Solidity: function controller_implementation() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) ControllerImplementation() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.ControllerImplementation(&_CrvUsdControllerFactory.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) Controllers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Controllers(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 arg0) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) Controllers(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Controllers(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// DebtCeiling is a free data retrieval call binding the contract method 0x602b62d4.
//
// Solidity: function debt_ceiling(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) DebtCeiling(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "debt_ceiling", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtCeiling is a free data retrieval call binding the contract method 0x602b62d4.
//
// Solidity: function debt_ceiling(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) DebtCeiling(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.DebtCeiling(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// DebtCeiling is a free data retrieval call binding the contract method 0x602b62d4.
//
// Solidity: function debt_ceiling(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) DebtCeiling(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.DebtCeiling(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// DebtCeilingResidual is a free data retrieval call binding the contract method 0xc92696b2.
//
// Solidity: function debt_ceiling_residual(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) DebtCeilingResidual(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "debt_ceiling_residual", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtCeilingResidual is a free data retrieval call binding the contract method 0xc92696b2.
//
// Solidity: function debt_ceiling_residual(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) DebtCeilingResidual(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.DebtCeilingResidual(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// DebtCeilingResidual is a free data retrieval call binding the contract method 0xc92696b2.
//
// Solidity: function debt_ceiling_residual(address arg0) view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) DebtCeilingResidual(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.DebtCeilingResidual(&_CrvUsdControllerFactory.CallOpts, arg0)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) FeeReceiver() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.FeeReceiver(&_CrvUsdControllerFactory.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) FeeReceiver() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.FeeReceiver(&_CrvUsdControllerFactory.CallOpts)
}

// GetAmm is a free data retrieval call binding the contract method 0xb12d5460.
//
// Solidity: function get_amm(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) GetAmm(opts *bind.CallOpts, collateral common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "get_amm", collateral)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAmm is a free data retrieval call binding the contract method 0xb12d5460.
//
// Solidity: function get_amm(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) GetAmm(collateral common.Address) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetAmm(&_CrvUsdControllerFactory.CallOpts, collateral)
}

// GetAmm is a free data retrieval call binding the contract method 0xb12d5460.
//
// Solidity: function get_amm(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) GetAmm(collateral common.Address) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetAmm(&_CrvUsdControllerFactory.CallOpts, collateral)
}

// GetAmm0 is a free data retrieval call binding the contract method 0x1539838f.
//
// Solidity: function get_amm(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) GetAmm0(opts *bind.CallOpts, collateral common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "get_amm0", collateral, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAmm0 is a free data retrieval call binding the contract method 0x1539838f.
//
// Solidity: function get_amm(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) GetAmm0(collateral common.Address, i *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetAmm0(&_CrvUsdControllerFactory.CallOpts, collateral, i)
}

// GetAmm0 is a free data retrieval call binding the contract method 0x1539838f.
//
// Solidity: function get_amm(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) GetAmm0(collateral common.Address, i *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetAmm0(&_CrvUsdControllerFactory.CallOpts, collateral, i)
}

// GetController is a free data retrieval call binding the contract method 0xf4410236.
//
// Solidity: function get_controller(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) GetController(opts *bind.CallOpts, collateral common.Address) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "get_controller", collateral)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0xf4410236.
//
// Solidity: function get_controller(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) GetController(collateral common.Address) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetController(&_CrvUsdControllerFactory.CallOpts, collateral)
}

// GetController is a free data retrieval call binding the contract method 0xf4410236.
//
// Solidity: function get_controller(address collateral) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) GetController(collateral common.Address) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetController(&_CrvUsdControllerFactory.CallOpts, collateral)
}

// GetController0 is a free data retrieval call binding the contract method 0x6bbd5ec7.
//
// Solidity: function get_controller(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) GetController0(opts *bind.CallOpts, collateral common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "get_controller0", collateral, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController0 is a free data retrieval call binding the contract method 0x6bbd5ec7.
//
// Solidity: function get_controller(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) GetController0(collateral common.Address, i *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetController0(&_CrvUsdControllerFactory.CallOpts, collateral, i)
}

// GetController0 is a free data retrieval call binding the contract method 0x6bbd5ec7.
//
// Solidity: function get_controller(address collateral, uint256 i) view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) GetController0(collateral common.Address, i *big.Int) (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.GetController0(&_CrvUsdControllerFactory.CallOpts, collateral, i)
}

// NCollaterals is a free data retrieval call binding the contract method 0x12397fa1.
//
// Solidity: function n_collaterals() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) NCollaterals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "n_collaterals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCollaterals is a free data retrieval call binding the contract method 0x12397fa1.
//
// Solidity: function n_collaterals() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) NCollaterals() (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.NCollaterals(&_CrvUsdControllerFactory.CallOpts)
}

// NCollaterals is a free data retrieval call binding the contract method 0x12397fa1.
//
// Solidity: function n_collaterals() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) NCollaterals() (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.NCollaterals(&_CrvUsdControllerFactory.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "stablecoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) Stablecoin() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Stablecoin(&_CrvUsdControllerFactory.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) Stablecoin() (common.Address, error) {
	return _CrvUsdControllerFactory.Contract.Stablecoin(&_CrvUsdControllerFactory.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdControllerFactory.contract.Call(opts, &out, "total_debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) TotalDebt() (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.TotalDebt(&_CrvUsdControllerFactory.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryCallerSession) TotalDebt() (*big.Int, error) {
	return _CrvUsdControllerFactory.Contract.TotalDebt(&_CrvUsdControllerFactory.CallOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0x8a82f1fe.
//
// Solidity: function add_market(address token, uint256 A, uint256 fee, uint256 admin_fee, address _price_oracle_contract, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount, uint256 debt_ceiling) returns(address[2])
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) AddMarket(opts *bind.TransactOpts, token common.Address, A *big.Int, fee *big.Int, admin_fee *big.Int, _price_oracle_contract common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "add_market", token, A, fee, admin_fee, _price_oracle_contract, monetary_policy, loan_discount, liquidation_discount, debt_ceiling)
}

// AddMarket is a paid mutator transaction binding the contract method 0x8a82f1fe.
//
// Solidity: function add_market(address token, uint256 A, uint256 fee, uint256 admin_fee, address _price_oracle_contract, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount, uint256 debt_ceiling) returns(address[2])
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) AddMarket(token common.Address, A *big.Int, fee *big.Int, admin_fee *big.Int, _price_oracle_contract common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.AddMarket(&_CrvUsdControllerFactory.TransactOpts, token, A, fee, admin_fee, _price_oracle_contract, monetary_policy, loan_discount, liquidation_discount, debt_ceiling)
}

// AddMarket is a paid mutator transaction binding the contract method 0x8a82f1fe.
//
// Solidity: function add_market(address token, uint256 A, uint256 fee, uint256 admin_fee, address _price_oracle_contract, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount, uint256 debt_ceiling) returns(address[2])
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) AddMarket(token common.Address, A *big.Int, fee *big.Int, admin_fee *big.Int, _price_oracle_contract common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.AddMarket(&_CrvUsdControllerFactory.TransactOpts, token, A, fee, admin_fee, _price_oracle_contract, monetary_policy, loan_discount, liquidation_discount, debt_ceiling)
}

// CollectFeesAboveCeiling is a paid mutator transaction binding the contract method 0xdf068aba.
//
// Solidity: function collect_fees_above_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) CollectFeesAboveCeiling(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "collect_fees_above_ceiling", _to)
}

// CollectFeesAboveCeiling is a paid mutator transaction binding the contract method 0xdf068aba.
//
// Solidity: function collect_fees_above_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) CollectFeesAboveCeiling(_to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.CollectFeesAboveCeiling(&_CrvUsdControllerFactory.TransactOpts, _to)
}

// CollectFeesAboveCeiling is a paid mutator transaction binding the contract method 0xdf068aba.
//
// Solidity: function collect_fees_above_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) CollectFeesAboveCeiling(_to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.CollectFeesAboveCeiling(&_CrvUsdControllerFactory.TransactOpts, _to)
}

// RugDebtCeiling is a paid mutator transaction binding the contract method 0x2651990f.
//
// Solidity: function rug_debt_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) RugDebtCeiling(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "rug_debt_ceiling", _to)
}

// RugDebtCeiling is a paid mutator transaction binding the contract method 0x2651990f.
//
// Solidity: function rug_debt_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) RugDebtCeiling(_to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.RugDebtCeiling(&_CrvUsdControllerFactory.TransactOpts, _to)
}

// RugDebtCeiling is a paid mutator transaction binding the contract method 0x2651990f.
//
// Solidity: function rug_debt_ceiling(address _to) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) RugDebtCeiling(_to common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.RugDebtCeiling(&_CrvUsdControllerFactory.TransactOpts, _to)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) SetAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "set_admin", admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetAdmin(&_CrvUsdControllerFactory.TransactOpts, admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetAdmin(&_CrvUsdControllerFactory.TransactOpts, admin)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb933c50a.
//
// Solidity: function set_debt_ceiling(address _to, uint256 debt_ceiling) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) SetDebtCeiling(opts *bind.TransactOpts, _to common.Address, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "set_debt_ceiling", _to, debt_ceiling)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb933c50a.
//
// Solidity: function set_debt_ceiling(address _to, uint256 debt_ceiling) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) SetDebtCeiling(_to common.Address, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetDebtCeiling(&_CrvUsdControllerFactory.TransactOpts, _to, debt_ceiling)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb933c50a.
//
// Solidity: function set_debt_ceiling(address _to, uint256 debt_ceiling) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) SetDebtCeiling(_to common.Address, debt_ceiling *big.Int) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetDebtCeiling(&_CrvUsdControllerFactory.TransactOpts, _to, debt_ceiling)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address fee_receiver) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) SetFeeReceiver(opts *bind.TransactOpts, fee_receiver common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "set_fee_receiver", fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address fee_receiver) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) SetFeeReceiver(fee_receiver common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetFeeReceiver(&_CrvUsdControllerFactory.TransactOpts, fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address fee_receiver) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) SetFeeReceiver(fee_receiver common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetFeeReceiver(&_CrvUsdControllerFactory.TransactOpts, fee_receiver)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x9af8f23c.
//
// Solidity: function set_implementations(address controller, address amm) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactor) SetImplementations(opts *bind.TransactOpts, controller common.Address, amm common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.contract.Transact(opts, "set_implementations", controller, amm)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x9af8f23c.
//
// Solidity: function set_implementations(address controller, address amm) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactorySession) SetImplementations(controller common.Address, amm common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetImplementations(&_CrvUsdControllerFactory.TransactOpts, controller, amm)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x9af8f23c.
//
// Solidity: function set_implementations(address controller, address amm) returns()
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryTransactorSession) SetImplementations(controller common.Address, amm common.Address) (*types.Transaction, error) {
	return _CrvUsdControllerFactory.Contract.SetImplementations(&_CrvUsdControllerFactory.TransactOpts, controller, amm)
}

// CrvUsdControllerFactoryAddMarketIterator is returned from FilterAddMarket and is used to iterate over the raw logs and unpacked data for AddMarket events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryAddMarketIterator struct {
	Event *CrvUsdControllerFactoryAddMarket // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactoryAddMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactoryAddMarket)
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
		it.Event = new(CrvUsdControllerFactoryAddMarket)
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
func (it *CrvUsdControllerFactoryAddMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactoryAddMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactoryAddMarket represents a AddMarket event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryAddMarket struct {
	Collateral     common.Address
	Controller     common.Address
	Amm            common.Address
	MonetaryPolicy common.Address
	Ix             *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddMarket is a free log retrieval operation binding the contract event 0xebbe0dfde9dde641808b7a803882653420f3a5b12bb405d238faed959e1e3aa3.
//
// Solidity: event AddMarket(address indexed collateral, address controller, address amm, address monetary_policy, uint256 ix)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterAddMarket(opts *bind.FilterOpts, collateral []common.Address) (*CrvUsdControllerFactoryAddMarketIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "AddMarket", collateralRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryAddMarketIterator{contract: _CrvUsdControllerFactory.contract, event: "AddMarket", logs: logs, sub: sub}, nil
}

// WatchAddMarket is a free log subscription operation binding the contract event 0xebbe0dfde9dde641808b7a803882653420f3a5b12bb405d238faed959e1e3aa3.
//
// Solidity: event AddMarket(address indexed collateral, address controller, address amm, address monetary_policy, uint256 ix)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchAddMarket(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactoryAddMarket, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "AddMarket", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactoryAddMarket)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "AddMarket", log); err != nil {
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

// ParseAddMarket is a log parse operation binding the contract event 0xebbe0dfde9dde641808b7a803882653420f3a5b12bb405d238faed959e1e3aa3.
//
// Solidity: event AddMarket(address indexed collateral, address controller, address amm, address monetary_policy, uint256 ix)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseAddMarket(log types.Log) (*CrvUsdControllerFactoryAddMarket, error) {
	event := new(CrvUsdControllerFactoryAddMarket)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "AddMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactoryMintForMarketIterator is returned from FilterMintForMarket and is used to iterate over the raw logs and unpacked data for MintForMarket events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryMintForMarketIterator struct {
	Event *CrvUsdControllerFactoryMintForMarket // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactoryMintForMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactoryMintForMarket)
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
		it.Event = new(CrvUsdControllerFactoryMintForMarket)
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
func (it *CrvUsdControllerFactoryMintForMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactoryMintForMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactoryMintForMarket represents a MintForMarket event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryMintForMarket struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintForMarket is a free log retrieval operation binding the contract event 0xad17aca0dc59a6d96350f71e2732094471c65b5a5cecd8f95b376edcd5534cc9.
//
// Solidity: event MintForMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterMintForMarket(opts *bind.FilterOpts, addr []common.Address) (*CrvUsdControllerFactoryMintForMarketIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "MintForMarket", addrRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryMintForMarketIterator{contract: _CrvUsdControllerFactory.contract, event: "MintForMarket", logs: logs, sub: sub}, nil
}

// WatchMintForMarket is a free log subscription operation binding the contract event 0xad17aca0dc59a6d96350f71e2732094471c65b5a5cecd8f95b376edcd5534cc9.
//
// Solidity: event MintForMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchMintForMarket(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactoryMintForMarket, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "MintForMarket", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactoryMintForMarket)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "MintForMarket", log); err != nil {
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

// ParseMintForMarket is a log parse operation binding the contract event 0xad17aca0dc59a6d96350f71e2732094471c65b5a5cecd8f95b376edcd5534cc9.
//
// Solidity: event MintForMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseMintForMarket(log types.Log) (*CrvUsdControllerFactoryMintForMarket, error) {
	event := new(CrvUsdControllerFactoryMintForMarket)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "MintForMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactoryRemoveFromMarketIterator is returned from FilterRemoveFromMarket and is used to iterate over the raw logs and unpacked data for RemoveFromMarket events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryRemoveFromMarketIterator struct {
	Event *CrvUsdControllerFactoryRemoveFromMarket // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactoryRemoveFromMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactoryRemoveFromMarket)
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
		it.Event = new(CrvUsdControllerFactoryRemoveFromMarket)
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
func (it *CrvUsdControllerFactoryRemoveFromMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactoryRemoveFromMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactoryRemoveFromMarket represents a RemoveFromMarket event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactoryRemoveFromMarket struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveFromMarket is a free log retrieval operation binding the contract event 0xb21604369d32a00404a085ea01ab0a3f6b63f8a0ebda770e25695572416d9bcf.
//
// Solidity: event RemoveFromMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterRemoveFromMarket(opts *bind.FilterOpts, addr []common.Address) (*CrvUsdControllerFactoryRemoveFromMarketIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "RemoveFromMarket", addrRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactoryRemoveFromMarketIterator{contract: _CrvUsdControllerFactory.contract, event: "RemoveFromMarket", logs: logs, sub: sub}, nil
}

// WatchRemoveFromMarket is a free log subscription operation binding the contract event 0xb21604369d32a00404a085ea01ab0a3f6b63f8a0ebda770e25695572416d9bcf.
//
// Solidity: event RemoveFromMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchRemoveFromMarket(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactoryRemoveFromMarket, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "RemoveFromMarket", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactoryRemoveFromMarket)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "RemoveFromMarket", log); err != nil {
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

// ParseRemoveFromMarket is a log parse operation binding the contract event 0xb21604369d32a00404a085ea01ab0a3f6b63f8a0ebda770e25695572416d9bcf.
//
// Solidity: event RemoveFromMarket(address indexed addr, uint256 amount)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseRemoveFromMarket(log types.Log) (*CrvUsdControllerFactoryRemoveFromMarket, error) {
	event := new(CrvUsdControllerFactoryRemoveFromMarket)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "RemoveFromMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactorySetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetAdminIterator struct {
	Event *CrvUsdControllerFactorySetAdmin // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactorySetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactorySetAdmin)
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
		it.Event = new(CrvUsdControllerFactorySetAdmin)
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
func (it *CrvUsdControllerFactorySetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactorySetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactorySetAdmin represents a SetAdmin event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*CrvUsdControllerFactorySetAdminIterator, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactorySetAdminIterator{contract: _CrvUsdControllerFactory.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactorySetAdmin) (event.Subscription, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactorySetAdmin)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseSetAdmin(log types.Log) (*CrvUsdControllerFactorySetAdmin, error) {
	event := new(CrvUsdControllerFactorySetAdmin)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactorySetDebtCeilingIterator is returned from FilterSetDebtCeiling and is used to iterate over the raw logs and unpacked data for SetDebtCeiling events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetDebtCeilingIterator struct {
	Event *CrvUsdControllerFactorySetDebtCeiling // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactorySetDebtCeilingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactorySetDebtCeiling)
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
		it.Event = new(CrvUsdControllerFactorySetDebtCeiling)
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
func (it *CrvUsdControllerFactorySetDebtCeilingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactorySetDebtCeilingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactorySetDebtCeiling represents a SetDebtCeiling event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetDebtCeiling struct {
	Addr        common.Address
	DebtCeiling *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetDebtCeiling is a free log retrieval operation binding the contract event 0x22d26e5448456e0d2368bca46b2c824717b39390656f1c6314237e11d691e4f2.
//
// Solidity: event SetDebtCeiling(address indexed addr, uint256 debt_ceiling)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterSetDebtCeiling(opts *bind.FilterOpts, addr []common.Address) (*CrvUsdControllerFactorySetDebtCeilingIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "SetDebtCeiling", addrRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactorySetDebtCeilingIterator{contract: _CrvUsdControllerFactory.contract, event: "SetDebtCeiling", logs: logs, sub: sub}, nil
}

// WatchSetDebtCeiling is a free log subscription operation binding the contract event 0x22d26e5448456e0d2368bca46b2c824717b39390656f1c6314237e11d691e4f2.
//
// Solidity: event SetDebtCeiling(address indexed addr, uint256 debt_ceiling)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchSetDebtCeiling(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactorySetDebtCeiling, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "SetDebtCeiling", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactorySetDebtCeiling)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetDebtCeiling", log); err != nil {
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

// ParseSetDebtCeiling is a log parse operation binding the contract event 0x22d26e5448456e0d2368bca46b2c824717b39390656f1c6314237e11d691e4f2.
//
// Solidity: event SetDebtCeiling(address indexed addr, uint256 debt_ceiling)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseSetDebtCeiling(log types.Log) (*CrvUsdControllerFactorySetDebtCeiling, error) {
	event := new(CrvUsdControllerFactorySetDebtCeiling)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetDebtCeiling", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactorySetFeeReceiverIterator is returned from FilterSetFeeReceiver and is used to iterate over the raw logs and unpacked data for SetFeeReceiver events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetFeeReceiverIterator struct {
	Event *CrvUsdControllerFactorySetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactorySetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactorySetFeeReceiver)
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
		it.Event = new(CrvUsdControllerFactorySetFeeReceiver)
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
func (it *CrvUsdControllerFactorySetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactorySetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactorySetFeeReceiver represents a SetFeeReceiver event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetFeeReceiver struct {
	FeeReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetFeeReceiver is a free log retrieval operation binding the contract event 0xffb40bfdfd246e95f543d08d9713c339f1d90fa9265e39b4f562f9011d7c919f.
//
// Solidity: event SetFeeReceiver(address fee_receiver)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterSetFeeReceiver(opts *bind.FilterOpts) (*CrvUsdControllerFactorySetFeeReceiverIterator, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "SetFeeReceiver")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactorySetFeeReceiverIterator{contract: _CrvUsdControllerFactory.contract, event: "SetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchSetFeeReceiver is a free log subscription operation binding the contract event 0xffb40bfdfd246e95f543d08d9713c339f1d90fa9265e39b4f562f9011d7c919f.
//
// Solidity: event SetFeeReceiver(address fee_receiver)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactorySetFeeReceiver) (event.Subscription, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "SetFeeReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactorySetFeeReceiver)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetFeeReceiver", log); err != nil {
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

// ParseSetFeeReceiver is a log parse operation binding the contract event 0xffb40bfdfd246e95f543d08d9713c339f1d90fa9265e39b4f562f9011d7c919f.
//
// Solidity: event SetFeeReceiver(address fee_receiver)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseSetFeeReceiver(log types.Log) (*CrvUsdControllerFactorySetFeeReceiver, error) {
	event := new(CrvUsdControllerFactorySetFeeReceiver)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerFactorySetImplementationsIterator is returned from FilterSetImplementations and is used to iterate over the raw logs and unpacked data for SetImplementations events raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetImplementationsIterator struct {
	Event *CrvUsdControllerFactorySetImplementations // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerFactorySetImplementationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerFactorySetImplementations)
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
		it.Event = new(CrvUsdControllerFactorySetImplementations)
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
func (it *CrvUsdControllerFactorySetImplementationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerFactorySetImplementationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerFactorySetImplementations represents a SetImplementations event raised by the CrvUsdControllerFactory contract.
type CrvUsdControllerFactorySetImplementations struct {
	Amm        common.Address
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetImplementations is a free log retrieval operation binding the contract event 0x1694b0703640754583177bb0c9e8d97e4d163cd89d08ae426ef8cb3f47109542.
//
// Solidity: event SetImplementations(address amm, address controller)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) FilterSetImplementations(opts *bind.FilterOpts) (*CrvUsdControllerFactorySetImplementationsIterator, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.FilterLogs(opts, "SetImplementations")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFactorySetImplementationsIterator{contract: _CrvUsdControllerFactory.contract, event: "SetImplementations", logs: logs, sub: sub}, nil
}

// WatchSetImplementations is a free log subscription operation binding the contract event 0x1694b0703640754583177bb0c9e8d97e4d163cd89d08ae426ef8cb3f47109542.
//
// Solidity: event SetImplementations(address amm, address controller)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) WatchSetImplementations(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerFactorySetImplementations) (event.Subscription, error) {

	logs, sub, err := _CrvUsdControllerFactory.contract.WatchLogs(opts, "SetImplementations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerFactorySetImplementations)
				if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetImplementations", log); err != nil {
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

// ParseSetImplementations is a log parse operation binding the contract event 0x1694b0703640754583177bb0c9e8d97e4d163cd89d08ae426ef8cb3f47109542.
//
// Solidity: event SetImplementations(address amm, address controller)
func (_CrvUsdControllerFactory *CrvUsdControllerFactoryFilterer) ParseSetImplementations(log types.Log) (*CrvUsdControllerFactorySetImplementations, error) {
	event := new(CrvUsdControllerFactorySetImplementations)
	if err := _CrvUsdControllerFactory.contract.UnpackLog(event, "SetImplementations", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
