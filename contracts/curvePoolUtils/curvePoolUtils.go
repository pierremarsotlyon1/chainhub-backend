// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvePoolUtils

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

// CurvePoolUtilsMetaData contains all meta data concerning the CurvePoolUtils contract.
var CurvePoolUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"claimable_one\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"claimable_two\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurvePoolUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolUtilsMetaData.ABI instead.
var CurvePoolUtilsABI = CurvePoolUtilsMetaData.ABI

// CurvePoolUtils is an auto generated Go binding around an Ethereum contract.
type CurvePoolUtils struct {
	CurvePoolUtilsCaller     // Read-only binding to the contract
	CurvePoolUtilsTransactor // Write-only binding to the contract
	CurvePoolUtilsFilterer   // Log filterer for contract events
}

// CurvePoolUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolUtilsSession struct {
	Contract     *CurvePoolUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolUtilsCallerSession struct {
	Contract *CurvePoolUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CurvePoolUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolUtilsTransactorSession struct {
	Contract     *CurvePoolUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CurvePoolUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolUtilsRaw struct {
	Contract *CurvePoolUtils // Generic contract binding to access the raw methods on
}

// CurvePoolUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolUtilsCallerRaw struct {
	Contract *CurvePoolUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolUtilsTransactorRaw struct {
	Contract *CurvePoolUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolUtils creates a new instance of CurvePoolUtils, bound to a specific deployed contract.
func NewCurvePoolUtils(address common.Address, backend bind.ContractBackend) (*CurvePoolUtils, error) {
	contract, err := bindCurvePoolUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolUtils{CurvePoolUtilsCaller: CurvePoolUtilsCaller{contract: contract}, CurvePoolUtilsTransactor: CurvePoolUtilsTransactor{contract: contract}, CurvePoolUtilsFilterer: CurvePoolUtilsFilterer{contract: contract}}, nil
}

// NewCurvePoolUtilsCaller creates a new read-only instance of CurvePoolUtils, bound to a specific deployed contract.
func NewCurvePoolUtilsCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolUtilsCaller, error) {
	contract, err := bindCurvePoolUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolUtilsCaller{contract: contract}, nil
}

// NewCurvePoolUtilsTransactor creates a new write-only instance of CurvePoolUtils, bound to a specific deployed contract.
func NewCurvePoolUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolUtilsTransactor, error) {
	contract, err := bindCurvePoolUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolUtilsTransactor{contract: contract}, nil
}

// NewCurvePoolUtilsFilterer creates a new log filterer instance of CurvePoolUtils, bound to a specific deployed contract.
func NewCurvePoolUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolUtilsFilterer, error) {
	contract, err := bindCurvePoolUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolUtilsFilterer{contract: contract}, nil
}

// bindCurvePoolUtils binds a generic wrapper to an already deployed contract.
func bindCurvePoolUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolUtils *CurvePoolUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolUtils.Contract.CurvePoolUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolUtils *CurvePoolUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolUtils.Contract.CurvePoolUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolUtils *CurvePoolUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolUtils.Contract.CurvePoolUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolUtils *CurvePoolUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolUtils *CurvePoolUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolUtils *CurvePoolUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolUtils.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address to) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCaller) BalanceOf(opts *bind.CallOpts, token common.Address, to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolUtils.contract.Call(opts, &out, "balanceOf", token, to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address to) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsSession) BalanceOf(token common.Address, to common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.BalanceOf(&_CurvePoolUtils.CallOpts, token, to)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address to) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCallerSession) BalanceOf(token common.Address, to common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.BalanceOf(&_CurvePoolUtils.CallOpts, token, to)
}

// Balances is a free data retrieval call binding the contract method 0x5f469221.
//
// Solidity: function balances(address to, int128 i) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCaller) Balances(opts *bind.CallOpts, to common.Address, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolUtils.contract.Call(opts, &out, "balances", to, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x5f469221.
//
// Solidity: function balances(address to, int128 i) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsSession) Balances(to common.Address, i *big.Int) (*big.Int, error) {
	return _CurvePoolUtils.Contract.Balances(&_CurvePoolUtils.CallOpts, to, i)
}

// Balances is a free data retrieval call binding the contract method 0x5f469221.
//
// Solidity: function balances(address to, int128 i) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCallerSession) Balances(to common.Address, i *big.Int) (*big.Int, error) {
	return _CurvePoolUtils.Contract.Balances(&_CurvePoolUtils.CallOpts, to, i)
}

// Claimable is a free data retrieval call binding the contract method 0xd4570c1c.
//
// Solidity: function claimable(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCaller) Claimable(opts *bind.CallOpts, poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolUtils.contract.Call(opts, &out, "claimable", poolAddress, feeCollector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xd4570c1c.
//
// Solidity: function claimable(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsSession) Claimable(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.Claimable(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}

// Claimable is a free data retrieval call binding the contract method 0xd4570c1c.
//
// Solidity: function claimable(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCallerSession) Claimable(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.Claimable(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}

// ClaimableOne is a free data retrieval call binding the contract method 0x433218dc.
//
// Solidity: function claimable_one(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCaller) ClaimableOne(opts *bind.CallOpts, poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolUtils.contract.Call(opts, &out, "claimable_one", poolAddress, feeCollector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableOne is a free data retrieval call binding the contract method 0x433218dc.
//
// Solidity: function claimable_one(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsSession) ClaimableOne(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.ClaimableOne(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}

// ClaimableOne is a free data retrieval call binding the contract method 0x433218dc.
//
// Solidity: function claimable_one(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCallerSession) ClaimableOne(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.ClaimableOne(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}

// ClaimableTwo is a free data retrieval call binding the contract method 0x82af8ebd.
//
// Solidity: function claimable_two(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCaller) ClaimableTwo(opts *bind.CallOpts, poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolUtils.contract.Call(opts, &out, "claimable_two", poolAddress, feeCollector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTwo is a free data retrieval call binding the contract method 0x82af8ebd.
//
// Solidity: function claimable_two(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsSession) ClaimableTwo(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.ClaimableTwo(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}

// ClaimableTwo is a free data retrieval call binding the contract method 0x82af8ebd.
//
// Solidity: function claimable_two(address poolAddress, address feeCollector) view returns(uint256)
func (_CurvePoolUtils *CurvePoolUtilsCallerSession) ClaimableTwo(poolAddress common.Address, feeCollector common.Address) (*big.Int, error) {
	return _CurvePoolUtils.Contract.ClaimableTwo(&_CurvePoolUtils.CallOpts, poolAddress, feeCollector)
}
