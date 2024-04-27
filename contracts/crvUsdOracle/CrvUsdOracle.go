// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdOracle

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
	Pool      common.Address
	IsInverse bool
}

// CrvUsdOracleMetaData contains all meta data concerning the CrvUsdOracle contract.
var CrvUsdOracleMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"AddPricePair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"is_inverse\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemovePricePair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MovePricePair\",\"inputs\":[{\"name\":\"n_from\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"n_to\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"stablecoin\",\"type\":\"address\"},{\"name\":\"sigma\",\"type\":\"uint256\"},{\"name\":\"admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"sigma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stablecoin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_price_pair\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_price_pair\",\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_tvl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"price_w\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_pairs\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"is_inverse\",\"type\":\"bool\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_tvl\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TVL_MA_TIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CrvUsdOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdOracleMetaData.ABI instead.
var CrvUsdOracleABI = CrvUsdOracleMetaData.ABI

// CrvUsdOracle is an auto generated Go binding around an Ethereum contract.
type CrvUsdOracle struct {
	CrvUsdOracleCaller     // Read-only binding to the contract
	CrvUsdOracleTransactor // Write-only binding to the contract
	CrvUsdOracleFilterer   // Log filterer for contract events
}

// CrvUsdOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdOracleSession struct {
	Contract     *CrvUsdOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvUsdOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdOracleCallerSession struct {
	Contract *CrvUsdOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrvUsdOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdOracleTransactorSession struct {
	Contract     *CrvUsdOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrvUsdOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdOracleRaw struct {
	Contract *CrvUsdOracle // Generic contract binding to access the raw methods on
}

// CrvUsdOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdOracleCallerRaw struct {
	Contract *CrvUsdOracleCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdOracleTransactorRaw struct {
	Contract *CrvUsdOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdOracle creates a new instance of CrvUsdOracle, bound to a specific deployed contract.
func NewCrvUsdOracle(address common.Address, backend bind.ContractBackend) (*CrvUsdOracle, error) {
	contract, err := bindCrvUsdOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracle{CrvUsdOracleCaller: CrvUsdOracleCaller{contract: contract}, CrvUsdOracleTransactor: CrvUsdOracleTransactor{contract: contract}, CrvUsdOracleFilterer: CrvUsdOracleFilterer{contract: contract}}, nil
}

// NewCrvUsdOracleCaller creates a new read-only instance of CrvUsdOracle, bound to a specific deployed contract.
func NewCrvUsdOracleCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdOracleCaller, error) {
	contract, err := bindCrvUsdOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleCaller{contract: contract}, nil
}

// NewCrvUsdOracleTransactor creates a new write-only instance of CrvUsdOracle, bound to a specific deployed contract.
func NewCrvUsdOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdOracleTransactor, error) {
	contract, err := bindCrvUsdOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleTransactor{contract: contract}, nil
}

// NewCrvUsdOracleFilterer creates a new log filterer instance of CrvUsdOracle, bound to a specific deployed contract.
func NewCrvUsdOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdOracleFilterer, error) {
	contract, err := bindCrvUsdOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleFilterer{contract: contract}, nil
}

// bindCrvUsdOracle binds a generic wrapper to an already deployed contract.
func bindCrvUsdOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdOracle *CrvUsdOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdOracle.Contract.CrvUsdOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdOracle *CrvUsdOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.CrvUsdOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdOracle *CrvUsdOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.CrvUsdOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdOracle *CrvUsdOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdOracle *CrvUsdOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdOracle *CrvUsdOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.contract.Transact(opts, method, params...)
}

// TVLMATIME is a free data retrieval call binding the contract method 0x8d45972e.
//
// Solidity: function TVL_MA_TIME() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) TVLMATIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "TVL_MA_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TVLMATIME is a free data retrieval call binding the contract method 0x8d45972e.
//
// Solidity: function TVL_MA_TIME() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) TVLMATIME() (*big.Int, error) {
	return _CrvUsdOracle.Contract.TVLMATIME(&_CrvUsdOracle.CallOpts)
}

// TVLMATIME is a free data retrieval call binding the contract method 0x8d45972e.
//
// Solidity: function TVL_MA_TIME() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) TVLMATIME() (*big.Int, error) {
	return _CrvUsdOracle.Contract.TVLMATIME(&_CrvUsdOracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleSession) Admin() (common.Address, error) {
	return _CrvUsdOracle.Contract.Admin(&_CrvUsdOracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) Admin() (common.Address, error) {
	return _CrvUsdOracle.Contract.Admin(&_CrvUsdOracle.CallOpts)
}

// EmaTvl is a free data retrieval call binding the contract method 0x33e3f712.
//
// Solidity: function ema_tvl() view returns(uint256[])
func (_CrvUsdOracle *CrvUsdOracleCaller) EmaTvl(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "ema_tvl")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EmaTvl is a free data retrieval call binding the contract method 0x33e3f712.
//
// Solidity: function ema_tvl() view returns(uint256[])
func (_CrvUsdOracle *CrvUsdOracleSession) EmaTvl() ([]*big.Int, error) {
	return _CrvUsdOracle.Contract.EmaTvl(&_CrvUsdOracle.CallOpts)
}

// EmaTvl is a free data retrieval call binding the contract method 0x33e3f712.
//
// Solidity: function ema_tvl() view returns(uint256[])
func (_CrvUsdOracle *CrvUsdOracleCallerSession) EmaTvl() ([]*big.Int, error) {
	return _CrvUsdOracle.Contract.EmaTvl(&_CrvUsdOracle.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "last_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) LastPrice() (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastPrice(&_CrvUsdOracle.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0xfde625e6.
//
// Solidity: function last_price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) LastPrice() (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastPrice(&_CrvUsdOracle.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) LastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "last_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) LastTimestamp() (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastTimestamp(&_CrvUsdOracle.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x4d23bfa0.
//
// Solidity: function last_timestamp() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) LastTimestamp() (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastTimestamp(&_CrvUsdOracle.CallOpts)
}

// LastTvl is a free data retrieval call binding the contract method 0x42e5a6c8.
//
// Solidity: function last_tvl(uint256 arg0) view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) LastTvl(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "last_tvl", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTvl is a free data retrieval call binding the contract method 0x42e5a6c8.
//
// Solidity: function last_tvl(uint256 arg0) view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) LastTvl(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastTvl(&_CrvUsdOracle.CallOpts, arg0)
}

// LastTvl is a free data retrieval call binding the contract method 0x42e5a6c8.
//
// Solidity: function last_tvl(uint256 arg0) view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) LastTvl(arg0 *big.Int) (*big.Int, error) {
	return _CrvUsdOracle.Contract.LastTvl(&_CrvUsdOracle.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) Price() (*big.Int, error) {
	return _CrvUsdOracle.Contract.Price(&_CrvUsdOracle.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) Price() (*big.Int, error) {
	return _CrvUsdOracle.Contract.Price(&_CrvUsdOracle.CallOpts)
}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUsdOracle *CrvUsdOracleCaller) PricePairs(opts *bind.CallOpts, arg0 *big.Int) (Struct0, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "price_pairs", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUsdOracle *CrvUsdOracleSession) PricePairs(arg0 *big.Int) (Struct0, error) {
	return _CrvUsdOracle.Contract.PricePairs(&_CrvUsdOracle.CallOpts, arg0)
}

// PricePairs is a free data retrieval call binding the contract method 0xba5feb37.
//
// Solidity: function price_pairs(uint256 arg0) view returns((address,bool))
func (_CrvUsdOracle *CrvUsdOracleCallerSession) PricePairs(arg0 *big.Int) (Struct0, error) {
	return _CrvUsdOracle.Contract.PricePairs(&_CrvUsdOracle.CallOpts, arg0)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCaller) Sigma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "sigma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) Sigma() (*big.Int, error) {
	return _CrvUsdOracle.Contract.Sigma(&_CrvUsdOracle.CallOpts)
}

// Sigma is a free data retrieval call binding the contract method 0xafdf31cd.
//
// Solidity: function sigma() view returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) Sigma() (*big.Int, error) {
	return _CrvUsdOracle.Contract.Sigma(&_CrvUsdOracle.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdOracle.contract.Call(opts, &out, "stablecoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleSession) Stablecoin() (common.Address, error) {
	return _CrvUsdOracle.Contract.Stablecoin(&_CrvUsdOracle.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() view returns(address)
func (_CrvUsdOracle *CrvUsdOracleCallerSession) Stablecoin() (common.Address, error) {
	return _CrvUsdOracle.Contract.Stablecoin(&_CrvUsdOracle.CallOpts)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactor) AddPricePair(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.contract.Transact(opts, "add_price_pair", _pool)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUsdOracle *CrvUsdOracleSession) AddPricePair(_pool common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.AddPricePair(&_CrvUsdOracle.TransactOpts, _pool)
}

// AddPricePair is a paid mutator transaction binding the contract method 0xa51ec1dd.
//
// Solidity: function add_price_pair(address _pool) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactorSession) AddPricePair(_pool common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.AddPricePair(&_CrvUsdOracle.TransactOpts, _pool)
}

// PriceW is a paid mutator transaction binding the contract method 0xceb7f759.
//
// Solidity: function price_w() returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleTransactor) PriceW(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdOracle.contract.Transact(opts, "price_w")
}

// PriceW is a paid mutator transaction binding the contract method 0xceb7f759.
//
// Solidity: function price_w() returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleSession) PriceW() (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.PriceW(&_CrvUsdOracle.TransactOpts)
}

// PriceW is a paid mutator transaction binding the contract method 0xceb7f759.
//
// Solidity: function price_w() returns(uint256)
func (_CrvUsdOracle *CrvUsdOracleTransactorSession) PriceW() (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.PriceW(&_CrvUsdOracle.TransactOpts)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactor) RemovePricePair(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CrvUsdOracle.contract.Transact(opts, "remove_price_pair", n)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUsdOracle *CrvUsdOracleSession) RemovePricePair(n *big.Int) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.RemovePricePair(&_CrvUsdOracle.TransactOpts, n)
}

// RemovePricePair is a paid mutator transaction binding the contract method 0xb3162fdb.
//
// Solidity: function remove_price_pair(uint256 n) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactorSession) RemovePricePair(n *big.Int) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.RemovePricePair(&_CrvUsdOracle.TransactOpts, n)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdOracle *CrvUsdOracleSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.SetAdmin(&_CrvUsdOracle.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_CrvUsdOracle *CrvUsdOracleTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CrvUsdOracle.Contract.SetAdmin(&_CrvUsdOracle.TransactOpts, _admin)
}

// CrvUsdOracleAddPricePairIterator is returned from FilterAddPricePair and is used to iterate over the raw logs and unpacked data for AddPricePair events raised by the CrvUsdOracle contract.
type CrvUsdOracleAddPricePairIterator struct {
	Event *CrvUsdOracleAddPricePair // Event containing the contract specifics and raw log

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
func (it *CrvUsdOracleAddPricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdOracleAddPricePair)
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
		it.Event = new(CrvUsdOracleAddPricePair)
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
func (it *CrvUsdOracleAddPricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdOracleAddPricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdOracleAddPricePair represents a AddPricePair event raised by the CrvUsdOracle contract.
type CrvUsdOracleAddPricePair struct {
	N         *big.Int
	Pool      common.Address
	IsInverse bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPricePair is a free log retrieval operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUsdOracle *CrvUsdOracleFilterer) FilterAddPricePair(opts *bind.FilterOpts) (*CrvUsdOracleAddPricePairIterator, error) {

	logs, sub, err := _CrvUsdOracle.contract.FilterLogs(opts, "AddPricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleAddPricePairIterator{contract: _CrvUsdOracle.contract, event: "AddPricePair", logs: logs, sub: sub}, nil
}

// WatchAddPricePair is a free log subscription operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUsdOracle *CrvUsdOracleFilterer) WatchAddPricePair(opts *bind.WatchOpts, sink chan<- *CrvUsdOracleAddPricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUsdOracle.contract.WatchLogs(opts, "AddPricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdOracleAddPricePair)
				if err := _CrvUsdOracle.contract.UnpackLog(event, "AddPricePair", log); err != nil {
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

// ParseAddPricePair is a log parse operation binding the contract event 0xfa41d19543baad0c0257c3430e1b3cecfa57f753cdc21330b9ec0862cde1b258.
//
// Solidity: event AddPricePair(uint256 n, address pool, bool is_inverse)
func (_CrvUsdOracle *CrvUsdOracleFilterer) ParseAddPricePair(log types.Log) (*CrvUsdOracleAddPricePair, error) {
	event := new(CrvUsdOracleAddPricePair)
	if err := _CrvUsdOracle.contract.UnpackLog(event, "AddPricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdOracleMovePricePairIterator is returned from FilterMovePricePair and is used to iterate over the raw logs and unpacked data for MovePricePair events raised by the CrvUsdOracle contract.
type CrvUsdOracleMovePricePairIterator struct {
	Event *CrvUsdOracleMovePricePair // Event containing the contract specifics and raw log

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
func (it *CrvUsdOracleMovePricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdOracleMovePricePair)
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
		it.Event = new(CrvUsdOracleMovePricePair)
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
func (it *CrvUsdOracleMovePricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdOracleMovePricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdOracleMovePricePair represents a MovePricePair event raised by the CrvUsdOracle contract.
type CrvUsdOracleMovePricePair struct {
	NFrom *big.Int
	NTo   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMovePricePair is a free log retrieval operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUsdOracle *CrvUsdOracleFilterer) FilterMovePricePair(opts *bind.FilterOpts) (*CrvUsdOracleMovePricePairIterator, error) {

	logs, sub, err := _CrvUsdOracle.contract.FilterLogs(opts, "MovePricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleMovePricePairIterator{contract: _CrvUsdOracle.contract, event: "MovePricePair", logs: logs, sub: sub}, nil
}

// WatchMovePricePair is a free log subscription operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUsdOracle *CrvUsdOracleFilterer) WatchMovePricePair(opts *bind.WatchOpts, sink chan<- *CrvUsdOracleMovePricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUsdOracle.contract.WatchLogs(opts, "MovePricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdOracleMovePricePair)
				if err := _CrvUsdOracle.contract.UnpackLog(event, "MovePricePair", log); err != nil {
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

// ParseMovePricePair is a log parse operation binding the contract event 0x3be85492dc07bf3888ee5d022674f0eced07fa5845d47b118257a8a0616d97ec.
//
// Solidity: event MovePricePair(uint256 n_from, uint256 n_to)
func (_CrvUsdOracle *CrvUsdOracleFilterer) ParseMovePricePair(log types.Log) (*CrvUsdOracleMovePricePair, error) {
	event := new(CrvUsdOracleMovePricePair)
	if err := _CrvUsdOracle.contract.UnpackLog(event, "MovePricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdOracleRemovePricePairIterator is returned from FilterRemovePricePair and is used to iterate over the raw logs and unpacked data for RemovePricePair events raised by the CrvUsdOracle contract.
type CrvUsdOracleRemovePricePairIterator struct {
	Event *CrvUsdOracleRemovePricePair // Event containing the contract specifics and raw log

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
func (it *CrvUsdOracleRemovePricePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdOracleRemovePricePair)
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
		it.Event = new(CrvUsdOracleRemovePricePair)
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
func (it *CrvUsdOracleRemovePricePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdOracleRemovePricePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdOracleRemovePricePair represents a RemovePricePair event raised by the CrvUsdOracle contract.
type CrvUsdOracleRemovePricePair struct {
	N   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemovePricePair is a free log retrieval operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUsdOracle *CrvUsdOracleFilterer) FilterRemovePricePair(opts *bind.FilterOpts) (*CrvUsdOracleRemovePricePairIterator, error) {

	logs, sub, err := _CrvUsdOracle.contract.FilterLogs(opts, "RemovePricePair")
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleRemovePricePairIterator{contract: _CrvUsdOracle.contract, event: "RemovePricePair", logs: logs, sub: sub}, nil
}

// WatchRemovePricePair is a free log subscription operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUsdOracle *CrvUsdOracleFilterer) WatchRemovePricePair(opts *bind.WatchOpts, sink chan<- *CrvUsdOracleRemovePricePair) (event.Subscription, error) {

	logs, sub, err := _CrvUsdOracle.contract.WatchLogs(opts, "RemovePricePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdOracleRemovePricePair)
				if err := _CrvUsdOracle.contract.UnpackLog(event, "RemovePricePair", log); err != nil {
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

// ParseRemovePricePair is a log parse operation binding the contract event 0x017592f2f16e82cccce60102865c737270289c308f34ff88e754d5e99ea0bae1.
//
// Solidity: event RemovePricePair(uint256 n)
func (_CrvUsdOracle *CrvUsdOracleFilterer) ParseRemovePricePair(log types.Log) (*CrvUsdOracleRemovePricePair, error) {
	event := new(CrvUsdOracleRemovePricePair)
	if err := _CrvUsdOracle.contract.UnpackLog(event, "RemovePricePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdOracleSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the CrvUsdOracle contract.
type CrvUsdOracleSetAdminIterator struct {
	Event *CrvUsdOracleSetAdmin // Event containing the contract specifics and raw log

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
func (it *CrvUsdOracleSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdOracleSetAdmin)
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
		it.Event = new(CrvUsdOracleSetAdmin)
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
func (it *CrvUsdOracleSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdOracleSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdOracleSetAdmin represents a SetAdmin event raised by the CrvUsdOracle contract.
type CrvUsdOracleSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdOracle *CrvUsdOracleFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*CrvUsdOracleSetAdminIterator, error) {

	logs, sub, err := _CrvUsdOracle.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &CrvUsdOracleSetAdminIterator{contract: _CrvUsdOracle.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_CrvUsdOracle *CrvUsdOracleFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *CrvUsdOracleSetAdmin) (event.Subscription, error) {

	logs, sub, err := _CrvUsdOracle.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdOracleSetAdmin)
				if err := _CrvUsdOracle.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_CrvUsdOracle *CrvUsdOracleFilterer) ParseSetAdmin(log types.Log) (*CrvUsdOracleSetAdmin, error) {
	event := new(CrvUsdOracleSetAdmin)
	if err := _CrvUsdOracle.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
