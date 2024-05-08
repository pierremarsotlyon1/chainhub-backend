// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package llamalendFactory

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

// LlamalendFactoryMetaData contains all meta data concerning the LlamalendFactory contract.
var LlamalendFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"SetImplementations\",\"inputs\":[{\"name\":\"amm\",\"type\":\"address\",\"indexed\":false},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"price_oracle\",\"type\":\"address\",\"indexed\":false},{\"name\":\"monetary_policy\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetDefaultRates\",\"inputs\":[{\"name\":\"min_rate\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"max_rate\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewVault\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"collateral_token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"borrowed_token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"controller\",\"type\":\"address\",\"indexed\":false},{\"name\":\"amm\",\"type\":\"address\",\"indexed\":false},{\"name\":\"price_oracle\",\"type\":\"address\",\"indexed\":false},{\"name\":\"monetary_policy\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\"},{\"name\":\"amm\",\"type\":\"address\"},{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"pool_price_oracle\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"gauge\",\"type\":\"address\"},{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"price_oracle\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"price_oracle\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"min_borrow_rate\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"price_oracle\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"min_borrow_rate\",\"type\":\"uint256\"},{\"name\":\"max_borrow_rate\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_from_pool\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_from_pool\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"min_borrow_rate\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_from_pool\",\"inputs\":[{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"min_borrow_rate\",\"type\":\"uint256\"},{\"name\":\"max_borrow_rate\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controllers\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"borrowed_tokens\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"collateral_tokens\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracles\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"monetary_policies\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vaults_index\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_for_vault\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_implementations\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\"},{\"name\":\"amm\",\"type\":\"address\"},{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"pool_price_oracle\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"gauge\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_default_rates\",\"inputs\":[{\"name\":\"min_rate\",\"type\":\"uint256\"},{\"name\":\"max_rate\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"vault_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"STABLECOIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MIN_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MAX_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controller_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vault_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_price_oracle_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"monetary_policy_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_impl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_default_borrow_rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_default_borrow_rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"vaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amms\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"market_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token_to_vaults\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token_market_count\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauges\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"names\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]}]",
}

// LlamalendFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LlamalendFactoryMetaData.ABI instead.
var LlamalendFactoryABI = LlamalendFactoryMetaData.ABI

// LlamalendFactory is an auto generated Go binding around an Ethereum contract.
type LlamalendFactory struct {
	LlamalendFactoryCaller     // Read-only binding to the contract
	LlamalendFactoryTransactor // Write-only binding to the contract
	LlamalendFactoryFilterer   // Log filterer for contract events
}

// LlamalendFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LlamalendFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LlamalendFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LlamalendFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LlamalendFactorySession struct {
	Contract     *LlamalendFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LlamalendFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LlamalendFactoryCallerSession struct {
	Contract *LlamalendFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LlamalendFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LlamalendFactoryTransactorSession struct {
	Contract     *LlamalendFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LlamalendFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LlamalendFactoryRaw struct {
	Contract *LlamalendFactory // Generic contract binding to access the raw methods on
}

// LlamalendFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LlamalendFactoryCallerRaw struct {
	Contract *LlamalendFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LlamalendFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LlamalendFactoryTransactorRaw struct {
	Contract *LlamalendFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLlamalendFactory creates a new instance of LlamalendFactory, bound to a specific deployed contract.
func NewLlamalendFactory(address common.Address, backend bind.ContractBackend) (*LlamalendFactory, error) {
	contract, err := bindLlamalendFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LlamalendFactory{LlamalendFactoryCaller: LlamalendFactoryCaller{contract: contract}, LlamalendFactoryTransactor: LlamalendFactoryTransactor{contract: contract}, LlamalendFactoryFilterer: LlamalendFactoryFilterer{contract: contract}}, nil
}

// NewLlamalendFactoryCaller creates a new read-only instance of LlamalendFactory, bound to a specific deployed contract.
func NewLlamalendFactoryCaller(address common.Address, caller bind.ContractCaller) (*LlamalendFactoryCaller, error) {
	contract, err := bindLlamalendFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendFactoryCaller{contract: contract}, nil
}

// NewLlamalendFactoryTransactor creates a new write-only instance of LlamalendFactory, bound to a specific deployed contract.
func NewLlamalendFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LlamalendFactoryTransactor, error) {
	contract, err := bindLlamalendFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendFactoryTransactor{contract: contract}, nil
}

// NewLlamalendFactoryFilterer creates a new log filterer instance of LlamalendFactory, bound to a specific deployed contract.
func NewLlamalendFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LlamalendFactoryFilterer, error) {
	contract, err := bindLlamalendFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LlamalendFactoryFilterer{contract: contract}, nil
}

// bindLlamalendFactory binds a generic wrapper to an already deployed contract.
func bindLlamalendFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LlamalendFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendFactory *LlamalendFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendFactory.Contract.LlamalendFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendFactory *LlamalendFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.LlamalendFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendFactory *LlamalendFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.LlamalendFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendFactory *LlamalendFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendFactory *LlamalendFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendFactory *LlamalendFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.contract.Transact(opts, method, params...)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) MAXRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "MAX_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) MAXRATE() (*big.Int, error) {
	return _LlamalendFactory.Contract.MAXRATE(&_LlamalendFactory.CallOpts)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MAXRATE() (*big.Int, error) {
	return _LlamalendFactory.Contract.MAXRATE(&_LlamalendFactory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) MINRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "MIN_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) MINRATE() (*big.Int, error) {
	return _LlamalendFactory.Contract.MINRATE(&_LlamalendFactory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MINRATE() (*big.Int, error) {
	return _LlamalendFactory.Contract.MINRATE(&_LlamalendFactory.CallOpts)
}

// STABLECOIN is a free data retrieval call binding the contract method 0x93a39776.
//
// Solidity: function STABLECOIN() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) STABLECOIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "STABLECOIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STABLECOIN is a free data retrieval call binding the contract method 0x93a39776.
//
// Solidity: function STABLECOIN() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) STABLECOIN() (common.Address, error) {
	return _LlamalendFactory.Contract.STABLECOIN(&_LlamalendFactory.CallOpts)
}

// STABLECOIN is a free data retrieval call binding the contract method 0x93a39776.
//
// Solidity: function STABLECOIN() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) STABLECOIN() (common.Address, error) {
	return _LlamalendFactory.Contract.STABLECOIN(&_LlamalendFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Admin() (common.Address, error) {
	return _LlamalendFactory.Contract.Admin(&_LlamalendFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Admin() (common.Address, error) {
	return _LlamalendFactory.Contract.Admin(&_LlamalendFactory.CallOpts)
}

// AmmImpl is a free data retrieval call binding the contract method 0x041622b5.
//
// Solidity: function amm_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) AmmImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "amm_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmImpl is a free data retrieval call binding the contract method 0x041622b5.
//
// Solidity: function amm_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) AmmImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.AmmImpl(&_LlamalendFactory.CallOpts)
}

// AmmImpl is a free data retrieval call binding the contract method 0x041622b5.
//
// Solidity: function amm_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) AmmImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.AmmImpl(&_LlamalendFactory.CallOpts)
}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) Amms(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "amms", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Amms(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Amms(&_LlamalendFactory.CallOpts, arg0)
}

// Amms is a free data retrieval call binding the contract method 0x86a8cdbc.
//
// Solidity: function amms(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Amms(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Amms(&_LlamalendFactory.CallOpts, arg0)
}

// BorrowedTokens is a free data retrieval call binding the contract method 0x6fe4501f.
//
// Solidity: function borrowed_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) BorrowedTokens(opts *bind.CallOpts, n *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "borrowed_tokens", n)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowedTokens is a free data retrieval call binding the contract method 0x6fe4501f.
//
// Solidity: function borrowed_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) BorrowedTokens(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.BorrowedTokens(&_LlamalendFactory.CallOpts, n)
}

// BorrowedTokens is a free data retrieval call binding the contract method 0x6fe4501f.
//
// Solidity: function borrowed_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) BorrowedTokens(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.BorrowedTokens(&_LlamalendFactory.CallOpts, n)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 vault_id) view returns(address[2])
func (_LlamalendFactory *LlamalendFactoryCaller) Coins(opts *bind.CallOpts, vault_id *big.Int) ([2]common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "coins", vault_id)

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 vault_id) view returns(address[2])
func (_LlamalendFactory *LlamalendFactorySession) Coins(vault_id *big.Int) ([2]common.Address, error) {
	return _LlamalendFactory.Contract.Coins(&_LlamalendFactory.CallOpts, vault_id)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 vault_id) view returns(address[2])
func (_LlamalendFactory *LlamalendFactoryCallerSession) Coins(vault_id *big.Int) ([2]common.Address, error) {
	return _LlamalendFactory.Contract.Coins(&_LlamalendFactory.CallOpts, vault_id)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x49b89984.
//
// Solidity: function collateral_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) CollateralTokens(opts *bind.CallOpts, n *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "collateral_tokens", n)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralTokens is a free data retrieval call binding the contract method 0x49b89984.
//
// Solidity: function collateral_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) CollateralTokens(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.CollateralTokens(&_LlamalendFactory.CallOpts, n)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x49b89984.
//
// Solidity: function collateral_tokens(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) CollateralTokens(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.CollateralTokens(&_LlamalendFactory.CallOpts, n)
}

// ControllerImpl is a free data retrieval call binding the contract method 0x168819c0.
//
// Solidity: function controller_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) ControllerImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "controller_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ControllerImpl is a free data retrieval call binding the contract method 0x168819c0.
//
// Solidity: function controller_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) ControllerImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.ControllerImpl(&_LlamalendFactory.CallOpts)
}

// ControllerImpl is a free data retrieval call binding the contract method 0x168819c0.
//
// Solidity: function controller_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) ControllerImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.ControllerImpl(&_LlamalendFactory.CallOpts)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) Controllers(opts *bind.CallOpts, n *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "controllers", n)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Controllers(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Controllers(&_LlamalendFactory.CallOpts, n)
}

// Controllers is a free data retrieval call binding the contract method 0xe94b0dd2.
//
// Solidity: function controllers(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Controllers(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Controllers(&_LlamalendFactory.CallOpts, n)
}

// GaugeForVault is a free data retrieval call binding the contract method 0x50c1163a.
//
// Solidity: function gauge_for_vault(address _vault) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) GaugeForVault(opts *bind.CallOpts, _vault common.Address) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "gauge_for_vault", _vault)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeForVault is a free data retrieval call binding the contract method 0x50c1163a.
//
// Solidity: function gauge_for_vault(address _vault) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) GaugeForVault(_vault common.Address) (common.Address, error) {
	return _LlamalendFactory.Contract.GaugeForVault(&_LlamalendFactory.CallOpts, _vault)
}

// GaugeForVault is a free data retrieval call binding the contract method 0x50c1163a.
//
// Solidity: function gauge_for_vault(address _vault) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) GaugeForVault(_vault common.Address) (common.Address, error) {
	return _LlamalendFactory.Contract.GaugeForVault(&_LlamalendFactory.CallOpts, _vault)
}

// GaugeImpl is a free data retrieval call binding the contract method 0x485227a7.
//
// Solidity: function gauge_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) GaugeImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "gauge_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImpl is a free data retrieval call binding the contract method 0x485227a7.
//
// Solidity: function gauge_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) GaugeImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.GaugeImpl(&_LlamalendFactory.CallOpts)
}

// GaugeImpl is a free data retrieval call binding the contract method 0x485227a7.
//
// Solidity: function gauge_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) GaugeImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.GaugeImpl(&_LlamalendFactory.CallOpts)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) Gauges(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "gauges", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Gauges(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Gauges(&_LlamalendFactory.CallOpts, arg0)
}

// Gauges is a free data retrieval call binding the contract method 0xb0539187.
//
// Solidity: function gauges(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Gauges(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Gauges(&_LlamalendFactory.CallOpts, arg0)
}

// MarketCount is a free data retrieval call binding the contract method 0xfd775c78.
//
// Solidity: function market_count() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "market_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCount is a free data retrieval call binding the contract method 0xfd775c78.
//
// Solidity: function market_count() view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) MarketCount() (*big.Int, error) {
	return _LlamalendFactory.Contract.MarketCount(&_LlamalendFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xfd775c78.
//
// Solidity: function market_count() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MarketCount() (*big.Int, error) {
	return _LlamalendFactory.Contract.MarketCount(&_LlamalendFactory.CallOpts)
}

// MaxDefaultBorrowRate is a free data retrieval call binding the contract method 0x99d0b1ba.
//
// Solidity: function max_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) MaxDefaultBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "max_default_borrow_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDefaultBorrowRate is a free data retrieval call binding the contract method 0x99d0b1ba.
//
// Solidity: function max_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) MaxDefaultBorrowRate() (*big.Int, error) {
	return _LlamalendFactory.Contract.MaxDefaultBorrowRate(&_LlamalendFactory.CallOpts)
}

// MaxDefaultBorrowRate is a free data retrieval call binding the contract method 0x99d0b1ba.
//
// Solidity: function max_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MaxDefaultBorrowRate() (*big.Int, error) {
	return _LlamalendFactory.Contract.MaxDefaultBorrowRate(&_LlamalendFactory.CallOpts)
}

// MinDefaultBorrowRate is a free data retrieval call binding the contract method 0x3cfd3d8c.
//
// Solidity: function min_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) MinDefaultBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "min_default_borrow_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDefaultBorrowRate is a free data retrieval call binding the contract method 0x3cfd3d8c.
//
// Solidity: function min_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) MinDefaultBorrowRate() (*big.Int, error) {
	return _LlamalendFactory.Contract.MinDefaultBorrowRate(&_LlamalendFactory.CallOpts)
}

// MinDefaultBorrowRate is a free data retrieval call binding the contract method 0x3cfd3d8c.
//
// Solidity: function min_default_borrow_rate() view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MinDefaultBorrowRate() (*big.Int, error) {
	return _LlamalendFactory.Contract.MinDefaultBorrowRate(&_LlamalendFactory.CallOpts)
}

// MonetaryPolicies is a free data retrieval call binding the contract method 0x762e7b92.
//
// Solidity: function monetary_policies(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) MonetaryPolicies(opts *bind.CallOpts, n *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "monetary_policies", n)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MonetaryPolicies is a free data retrieval call binding the contract method 0x762e7b92.
//
// Solidity: function monetary_policies(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) MonetaryPolicies(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.MonetaryPolicies(&_LlamalendFactory.CallOpts, n)
}

// MonetaryPolicies is a free data retrieval call binding the contract method 0x762e7b92.
//
// Solidity: function monetary_policies(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MonetaryPolicies(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.MonetaryPolicies(&_LlamalendFactory.CallOpts, n)
}

// MonetaryPolicyImpl is a free data retrieval call binding the contract method 0x247ec878.
//
// Solidity: function monetary_policy_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) MonetaryPolicyImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "monetary_policy_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MonetaryPolicyImpl is a free data retrieval call binding the contract method 0x247ec878.
//
// Solidity: function monetary_policy_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) MonetaryPolicyImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.MonetaryPolicyImpl(&_LlamalendFactory.CallOpts)
}

// MonetaryPolicyImpl is a free data retrieval call binding the contract method 0x247ec878.
//
// Solidity: function monetary_policy_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) MonetaryPolicyImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.MonetaryPolicyImpl(&_LlamalendFactory.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 arg0) view returns(string)
func (_LlamalendFactory *LlamalendFactoryCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 arg0) view returns(string)
func (_LlamalendFactory *LlamalendFactorySession) Names(arg0 *big.Int) (string, error) {
	return _LlamalendFactory.Contract.Names(&_LlamalendFactory.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 arg0) view returns(string)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Names(arg0 *big.Int) (string, error) {
	return _LlamalendFactory.Contract.Names(&_LlamalendFactory.CallOpts, arg0)
}

// PoolPriceOracleImpl is a free data retrieval call binding the contract method 0x0b91bd27.
//
// Solidity: function pool_price_oracle_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) PoolPriceOracleImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "pool_price_oracle_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolPriceOracleImpl is a free data retrieval call binding the contract method 0x0b91bd27.
//
// Solidity: function pool_price_oracle_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) PoolPriceOracleImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.PoolPriceOracleImpl(&_LlamalendFactory.CallOpts)
}

// PoolPriceOracleImpl is a free data retrieval call binding the contract method 0x0b91bd27.
//
// Solidity: function pool_price_oracle_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) PoolPriceOracleImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.PoolPriceOracleImpl(&_LlamalendFactory.CallOpts)
}

// PriceOracles is a free data retrieval call binding the contract method 0xc6f2a81d.
//
// Solidity: function price_oracles(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) PriceOracles(opts *bind.CallOpts, n *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "price_oracles", n)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracles is a free data retrieval call binding the contract method 0xc6f2a81d.
//
// Solidity: function price_oracles(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) PriceOracles(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.PriceOracles(&_LlamalendFactory.CallOpts, n)
}

// PriceOracles is a free data retrieval call binding the contract method 0xc6f2a81d.
//
// Solidity: function price_oracles(uint256 n) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) PriceOracles(n *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.PriceOracles(&_LlamalendFactory.CallOpts, n)
}

// TokenMarketCount is a free data retrieval call binding the contract method 0xe5f260ba.
//
// Solidity: function token_market_count(address arg0) view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) TokenMarketCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "token_market_count", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMarketCount is a free data retrieval call binding the contract method 0xe5f260ba.
//
// Solidity: function token_market_count(address arg0) view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) TokenMarketCount(arg0 common.Address) (*big.Int, error) {
	return _LlamalendFactory.Contract.TokenMarketCount(&_LlamalendFactory.CallOpts, arg0)
}

// TokenMarketCount is a free data retrieval call binding the contract method 0xe5f260ba.
//
// Solidity: function token_market_count(address arg0) view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) TokenMarketCount(arg0 common.Address) (*big.Int, error) {
	return _LlamalendFactory.Contract.TokenMarketCount(&_LlamalendFactory.CallOpts, arg0)
}

// TokenToVaults is a free data retrieval call binding the contract method 0xa8acf8df.
//
// Solidity: function token_to_vaults(address arg0, uint256 arg1) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) TokenToVaults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "token_to_vaults", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToVaults is a free data retrieval call binding the contract method 0xa8acf8df.
//
// Solidity: function token_to_vaults(address arg0, uint256 arg1) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) TokenToVaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.TokenToVaults(&_LlamalendFactory.CallOpts, arg0, arg1)
}

// TokenToVaults is a free data retrieval call binding the contract method 0xa8acf8df.
//
// Solidity: function token_to_vaults(address arg0, uint256 arg1) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) TokenToVaults(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.TokenToVaults(&_LlamalendFactory.CallOpts, arg0, arg1)
}

// VaultImpl is a free data retrieval call binding the contract method 0x6677b287.
//
// Solidity: function vault_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) VaultImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "vault_impl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultImpl is a free data retrieval call binding the contract method 0x6677b287.
//
// Solidity: function vault_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) VaultImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.VaultImpl(&_LlamalendFactory.CallOpts)
}

// VaultImpl is a free data retrieval call binding the contract method 0x6677b287.
//
// Solidity: function vault_impl() view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) VaultImpl() (common.Address, error) {
	return _LlamalendFactory.Contract.VaultImpl(&_LlamalendFactory.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "vaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Vaults(&_LlamalendFactory.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 arg0) view returns(address)
func (_LlamalendFactory *LlamalendFactoryCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _LlamalendFactory.Contract.Vaults(&_LlamalendFactory.CallOpts, arg0)
}

// VaultsIndex is a free data retrieval call binding the contract method 0xbcf75a8f.
//
// Solidity: function vaults_index(address vault) view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCaller) VaultsIndex(opts *bind.CallOpts, vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendFactory.contract.Call(opts, &out, "vaults_index", vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultsIndex is a free data retrieval call binding the contract method 0xbcf75a8f.
//
// Solidity: function vaults_index(address vault) view returns(uint256)
func (_LlamalendFactory *LlamalendFactorySession) VaultsIndex(vault common.Address) (*big.Int, error) {
	return _LlamalendFactory.Contract.VaultsIndex(&_LlamalendFactory.CallOpts, vault)
}

// VaultsIndex is a free data retrieval call binding the contract method 0xbcf75a8f.
//
// Solidity: function vaults_index(address vault) view returns(uint256)
func (_LlamalendFactory *LlamalendFactoryCallerSession) VaultsIndex(vault common.Address) (*big.Int, error) {
	return _LlamalendFactory.Contract.VaultsIndex(&_LlamalendFactory.CallOpts, vault)
}

// Create is a paid mutator transaction binding the contract method 0x9c3b70a6.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) Create(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name)
}

// Create is a paid mutator transaction binding the contract method 0x9c3b70a6.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Create(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name)
}

// Create is a paid mutator transaction binding the contract method 0x9c3b70a6.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) Create(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name)
}

// Create0 is a paid mutator transaction binding the contract method 0x5904e8dc.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) Create0(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create0", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate)
}

// Create0 is a paid mutator transaction binding the contract method 0x5904e8dc.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Create0(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create0(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate)
}

// Create0 is a paid mutator transaction binding the contract method 0x5904e8dc.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) Create0(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create0(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate)
}

// Create1 is a paid mutator transaction binding the contract method 0x5673683f.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) Create1(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create1", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate, max_borrow_rate)
}

// Create1 is a paid mutator transaction binding the contract method 0x5673683f.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) Create1(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create1(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate, max_borrow_rate)
}

// Create1 is a paid mutator transaction binding the contract method 0x5673683f.
//
// Solidity: function create(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address price_oracle, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) Create1(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, price_oracle common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.Create1(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, price_oracle, name, min_borrow_rate, max_borrow_rate)
}

// CreateFromPool is a paid mutator transaction binding the contract method 0xd984d2a1.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) CreateFromPool(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create_from_pool", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name)
}

// CreateFromPool is a paid mutator transaction binding the contract method 0xd984d2a1.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) CreateFromPool(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name)
}

// CreateFromPool is a paid mutator transaction binding the contract method 0xd984d2a1.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) CreateFromPool(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name)
}

// CreateFromPool0 is a paid mutator transaction binding the contract method 0x73edeb30.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) CreateFromPool0(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create_from_pool0", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate)
}

// CreateFromPool0 is a paid mutator transaction binding the contract method 0x73edeb30.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) CreateFromPool0(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool0(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate)
}

// CreateFromPool0 is a paid mutator transaction binding the contract method 0x73edeb30.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) CreateFromPool0(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool0(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate)
}

// CreateFromPool1 is a paid mutator transaction binding the contract method 0xbccc056c.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) CreateFromPool1(opts *bind.TransactOpts, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "create_from_pool1", borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate, max_borrow_rate)
}

// CreateFromPool1 is a paid mutator transaction binding the contract method 0xbccc056c.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) CreateFromPool1(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool1(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate, max_borrow_rate)
}

// CreateFromPool1 is a paid mutator transaction binding the contract method 0xbccc056c.
//
// Solidity: function create_from_pool(address borrowed_token, address collateral_token, uint256 A, uint256 fee, uint256 loan_discount, uint256 liquidation_discount, address pool, string name, uint256 min_borrow_rate, uint256 max_borrow_rate) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) CreateFromPool1(borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, loan_discount *big.Int, liquidation_discount *big.Int, pool common.Address, name string, min_borrow_rate *big.Int, max_borrow_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.CreateFromPool1(&_LlamalendFactory.TransactOpts, borrowed_token, collateral_token, A, fee, loan_discount, liquidation_discount, pool, name, min_borrow_rate, max_borrow_rate)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _vault) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "deploy_gauge", _vault)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _vault) returns(address)
func (_LlamalendFactory *LlamalendFactorySession) DeployGauge(_vault common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.DeployGauge(&_LlamalendFactory.TransactOpts, _vault)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _vault) returns(address)
func (_LlamalendFactory *LlamalendFactoryTransactorSession) DeployGauge(_vault common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.DeployGauge(&_LlamalendFactory.TransactOpts, _vault)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_LlamalendFactory *LlamalendFactoryTransactor) SetAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "set_admin", admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_LlamalendFactory *LlamalendFactorySession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetAdmin(&_LlamalendFactory.TransactOpts, admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address admin) returns()
func (_LlamalendFactory *LlamalendFactoryTransactorSession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetAdmin(&_LlamalendFactory.TransactOpts, admin)
}

// SetDefaultRates is a paid mutator transaction binding the contract method 0xa74fcf90.
//
// Solidity: function set_default_rates(uint256 min_rate, uint256 max_rate) returns()
func (_LlamalendFactory *LlamalendFactoryTransactor) SetDefaultRates(opts *bind.TransactOpts, min_rate *big.Int, max_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "set_default_rates", min_rate, max_rate)
}

// SetDefaultRates is a paid mutator transaction binding the contract method 0xa74fcf90.
//
// Solidity: function set_default_rates(uint256 min_rate, uint256 max_rate) returns()
func (_LlamalendFactory *LlamalendFactorySession) SetDefaultRates(min_rate *big.Int, max_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetDefaultRates(&_LlamalendFactory.TransactOpts, min_rate, max_rate)
}

// SetDefaultRates is a paid mutator transaction binding the contract method 0xa74fcf90.
//
// Solidity: function set_default_rates(uint256 min_rate, uint256 max_rate) returns()
func (_LlamalendFactory *LlamalendFactoryTransactorSession) SetDefaultRates(min_rate *big.Int, max_rate *big.Int) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetDefaultRates(&_LlamalendFactory.TransactOpts, min_rate, max_rate)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x6c15eadf.
//
// Solidity: function set_implementations(address controller, address amm, address vault, address pool_price_oracle, address monetary_policy, address gauge) returns()
func (_LlamalendFactory *LlamalendFactoryTransactor) SetImplementations(opts *bind.TransactOpts, controller common.Address, amm common.Address, vault common.Address, pool_price_oracle common.Address, monetary_policy common.Address, gauge common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.contract.Transact(opts, "set_implementations", controller, amm, vault, pool_price_oracle, monetary_policy, gauge)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x6c15eadf.
//
// Solidity: function set_implementations(address controller, address amm, address vault, address pool_price_oracle, address monetary_policy, address gauge) returns()
func (_LlamalendFactory *LlamalendFactorySession) SetImplementations(controller common.Address, amm common.Address, vault common.Address, pool_price_oracle common.Address, monetary_policy common.Address, gauge common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetImplementations(&_LlamalendFactory.TransactOpts, controller, amm, vault, pool_price_oracle, monetary_policy, gauge)
}

// SetImplementations is a paid mutator transaction binding the contract method 0x6c15eadf.
//
// Solidity: function set_implementations(address controller, address amm, address vault, address pool_price_oracle, address monetary_policy, address gauge) returns()
func (_LlamalendFactory *LlamalendFactoryTransactorSession) SetImplementations(controller common.Address, amm common.Address, vault common.Address, pool_price_oracle common.Address, monetary_policy common.Address, gauge common.Address) (*types.Transaction, error) {
	return _LlamalendFactory.Contract.SetImplementations(&_LlamalendFactory.TransactOpts, controller, amm, vault, pool_price_oracle, monetary_policy, gauge)
}

// LlamalendFactoryLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the LlamalendFactory contract.
type LlamalendFactoryLiquidityGaugeDeployedIterator struct {
	Event *LlamalendFactoryLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *LlamalendFactoryLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendFactoryLiquidityGaugeDeployed)
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
		it.Event = new(LlamalendFactoryLiquidityGaugeDeployed)
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
func (it *LlamalendFactoryLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendFactoryLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendFactoryLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the LlamalendFactory contract.
type LlamalendFactoryLiquidityGaugeDeployed struct {
	Vault common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address vault, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*LlamalendFactoryLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _LlamalendFactory.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &LlamalendFactoryLiquidityGaugeDeployedIterator{contract: _LlamalendFactory.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address vault, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *LlamalendFactoryLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _LlamalendFactory.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendFactoryLiquidityGaugeDeployed)
				if err := _LlamalendFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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

// ParseLiquidityGaugeDeployed is a log parse operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address vault, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*LlamalendFactoryLiquidityGaugeDeployed, error) {
	event := new(LlamalendFactoryLiquidityGaugeDeployed)
	if err := _LlamalendFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendFactoryNewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the LlamalendFactory contract.
type LlamalendFactoryNewVaultIterator struct {
	Event *LlamalendFactoryNewVault // Event containing the contract specifics and raw log

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
func (it *LlamalendFactoryNewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendFactoryNewVault)
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
		it.Event = new(LlamalendFactoryNewVault)
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
func (it *LlamalendFactoryNewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendFactoryNewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendFactoryNewVault represents a NewVault event raised by the LlamalendFactory contract.
type LlamalendFactoryNewVault struct {
	Id              *big.Int
	CollateralToken common.Address
	BorrowedToken   common.Address
	Vault           common.Address
	Controller      common.Address
	Amm             common.Address
	PriceOracle     common.Address
	MonetaryPolicy  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0x2a854a597908740dff5f0846840f167547ea0d7614c43bde3ea49be2e68c07ec.
//
// Solidity: event NewVault(uint256 indexed id, address indexed collateral_token, address indexed borrowed_token, address vault, address controller, address amm, address price_oracle, address monetary_policy)
func (_LlamalendFactory *LlamalendFactoryFilterer) FilterNewVault(opts *bind.FilterOpts, id []*big.Int, collateral_token []common.Address, borrowed_token []common.Address) (*LlamalendFactoryNewVaultIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var collateral_tokenRule []interface{}
	for _, collateral_tokenItem := range collateral_token {
		collateral_tokenRule = append(collateral_tokenRule, collateral_tokenItem)
	}
	var borrowed_tokenRule []interface{}
	for _, borrowed_tokenItem := range borrowed_token {
		borrowed_tokenRule = append(borrowed_tokenRule, borrowed_tokenItem)
	}

	logs, sub, err := _LlamalendFactory.contract.FilterLogs(opts, "NewVault", idRule, collateral_tokenRule, borrowed_tokenRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendFactoryNewVaultIterator{contract: _LlamalendFactory.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0x2a854a597908740dff5f0846840f167547ea0d7614c43bde3ea49be2e68c07ec.
//
// Solidity: event NewVault(uint256 indexed id, address indexed collateral_token, address indexed borrowed_token, address vault, address controller, address amm, address price_oracle, address monetary_policy)
func (_LlamalendFactory *LlamalendFactoryFilterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *LlamalendFactoryNewVault, id []*big.Int, collateral_token []common.Address, borrowed_token []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var collateral_tokenRule []interface{}
	for _, collateral_tokenItem := range collateral_token {
		collateral_tokenRule = append(collateral_tokenRule, collateral_tokenItem)
	}
	var borrowed_tokenRule []interface{}
	for _, borrowed_tokenItem := range borrowed_token {
		borrowed_tokenRule = append(borrowed_tokenRule, borrowed_tokenItem)
	}

	logs, sub, err := _LlamalendFactory.contract.WatchLogs(opts, "NewVault", idRule, collateral_tokenRule, borrowed_tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendFactoryNewVault)
				if err := _LlamalendFactory.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0x2a854a597908740dff5f0846840f167547ea0d7614c43bde3ea49be2e68c07ec.
//
// Solidity: event NewVault(uint256 indexed id, address indexed collateral_token, address indexed borrowed_token, address vault, address controller, address amm, address price_oracle, address monetary_policy)
func (_LlamalendFactory *LlamalendFactoryFilterer) ParseNewVault(log types.Log) (*LlamalendFactoryNewVault, error) {
	event := new(LlamalendFactoryNewVault)
	if err := _LlamalendFactory.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendFactorySetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the LlamalendFactory contract.
type LlamalendFactorySetAdminIterator struct {
	Event *LlamalendFactorySetAdmin // Event containing the contract specifics and raw log

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
func (it *LlamalendFactorySetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendFactorySetAdmin)
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
		it.Event = new(LlamalendFactorySetAdmin)
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
func (it *LlamalendFactorySetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendFactorySetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendFactorySetAdmin represents a SetAdmin event raised by the LlamalendFactory contract.
type LlamalendFactorySetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_LlamalendFactory *LlamalendFactoryFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*LlamalendFactorySetAdminIterator, error) {

	logs, sub, err := _LlamalendFactory.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &LlamalendFactorySetAdminIterator{contract: _LlamalendFactory.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_LlamalendFactory *LlamalendFactoryFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *LlamalendFactorySetAdmin) (event.Subscription, error) {

	logs, sub, err := _LlamalendFactory.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendFactorySetAdmin)
				if err := _LlamalendFactory.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_LlamalendFactory *LlamalendFactoryFilterer) ParseSetAdmin(log types.Log) (*LlamalendFactorySetAdmin, error) {
	event := new(LlamalendFactorySetAdmin)
	if err := _LlamalendFactory.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendFactorySetDefaultRatesIterator is returned from FilterSetDefaultRates and is used to iterate over the raw logs and unpacked data for SetDefaultRates events raised by the LlamalendFactory contract.
type LlamalendFactorySetDefaultRatesIterator struct {
	Event *LlamalendFactorySetDefaultRates // Event containing the contract specifics and raw log

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
func (it *LlamalendFactorySetDefaultRatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendFactorySetDefaultRates)
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
		it.Event = new(LlamalendFactorySetDefaultRates)
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
func (it *LlamalendFactorySetDefaultRatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendFactorySetDefaultRatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendFactorySetDefaultRates represents a SetDefaultRates event raised by the LlamalendFactory contract.
type LlamalendFactorySetDefaultRates struct {
	MinRate *big.Int
	MaxRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultRates is a free log retrieval operation binding the contract event 0x279f1fe0f91b15d983792d0305a146961875690054db0d81bec8d1582461fc65.
//
// Solidity: event SetDefaultRates(uint256 min_rate, uint256 max_rate)
func (_LlamalendFactory *LlamalendFactoryFilterer) FilterSetDefaultRates(opts *bind.FilterOpts) (*LlamalendFactorySetDefaultRatesIterator, error) {

	logs, sub, err := _LlamalendFactory.contract.FilterLogs(opts, "SetDefaultRates")
	if err != nil {
		return nil, err
	}
	return &LlamalendFactorySetDefaultRatesIterator{contract: _LlamalendFactory.contract, event: "SetDefaultRates", logs: logs, sub: sub}, nil
}

// WatchSetDefaultRates is a free log subscription operation binding the contract event 0x279f1fe0f91b15d983792d0305a146961875690054db0d81bec8d1582461fc65.
//
// Solidity: event SetDefaultRates(uint256 min_rate, uint256 max_rate)
func (_LlamalendFactory *LlamalendFactoryFilterer) WatchSetDefaultRates(opts *bind.WatchOpts, sink chan<- *LlamalendFactorySetDefaultRates) (event.Subscription, error) {

	logs, sub, err := _LlamalendFactory.contract.WatchLogs(opts, "SetDefaultRates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendFactorySetDefaultRates)
				if err := _LlamalendFactory.contract.UnpackLog(event, "SetDefaultRates", log); err != nil {
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

// ParseSetDefaultRates is a log parse operation binding the contract event 0x279f1fe0f91b15d983792d0305a146961875690054db0d81bec8d1582461fc65.
//
// Solidity: event SetDefaultRates(uint256 min_rate, uint256 max_rate)
func (_LlamalendFactory *LlamalendFactoryFilterer) ParseSetDefaultRates(log types.Log) (*LlamalendFactorySetDefaultRates, error) {
	event := new(LlamalendFactorySetDefaultRates)
	if err := _LlamalendFactory.contract.UnpackLog(event, "SetDefaultRates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendFactorySetImplementationsIterator is returned from FilterSetImplementations and is used to iterate over the raw logs and unpacked data for SetImplementations events raised by the LlamalendFactory contract.
type LlamalendFactorySetImplementationsIterator struct {
	Event *LlamalendFactorySetImplementations // Event containing the contract specifics and raw log

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
func (it *LlamalendFactorySetImplementationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendFactorySetImplementations)
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
		it.Event = new(LlamalendFactorySetImplementations)
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
func (it *LlamalendFactorySetImplementationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendFactorySetImplementationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendFactorySetImplementations represents a SetImplementations event raised by the LlamalendFactory contract.
type LlamalendFactorySetImplementations struct {
	Amm            common.Address
	Controller     common.Address
	Vault          common.Address
	PriceOracle    common.Address
	MonetaryPolicy common.Address
	Gauge          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetImplementations is a free log retrieval operation binding the contract event 0x91d63b24386eae580bbbe65f3f50fd736c41031f36d85641bc13e74ac0cb95bb.
//
// Solidity: event SetImplementations(address amm, address controller, address vault, address price_oracle, address monetary_policy, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) FilterSetImplementations(opts *bind.FilterOpts) (*LlamalendFactorySetImplementationsIterator, error) {

	logs, sub, err := _LlamalendFactory.contract.FilterLogs(opts, "SetImplementations")
	if err != nil {
		return nil, err
	}
	return &LlamalendFactorySetImplementationsIterator{contract: _LlamalendFactory.contract, event: "SetImplementations", logs: logs, sub: sub}, nil
}

// WatchSetImplementations is a free log subscription operation binding the contract event 0x91d63b24386eae580bbbe65f3f50fd736c41031f36d85641bc13e74ac0cb95bb.
//
// Solidity: event SetImplementations(address amm, address controller, address vault, address price_oracle, address monetary_policy, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) WatchSetImplementations(opts *bind.WatchOpts, sink chan<- *LlamalendFactorySetImplementations) (event.Subscription, error) {

	logs, sub, err := _LlamalendFactory.contract.WatchLogs(opts, "SetImplementations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendFactorySetImplementations)
				if err := _LlamalendFactory.contract.UnpackLog(event, "SetImplementations", log); err != nil {
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

// ParseSetImplementations is a log parse operation binding the contract event 0x91d63b24386eae580bbbe65f3f50fd736c41031f36d85641bc13e74ac0cb95bb.
//
// Solidity: event SetImplementations(address amm, address controller, address vault, address price_oracle, address monetary_policy, address gauge)
func (_LlamalendFactory *LlamalendFactoryFilterer) ParseSetImplementations(log types.Log) (*LlamalendFactorySetImplementations, error) {
	event := new(LlamalendFactorySetImplementations)
	if err := _LlamalendFactory.contract.UnpackLog(event, "SetImplementations", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
