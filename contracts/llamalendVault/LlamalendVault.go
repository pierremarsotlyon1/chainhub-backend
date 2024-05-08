// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package llamalendVault

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

// LlamalendVaultMetaData contains all meta data concerning the LlamalendVault contract.
var LlamalendVaultMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"amm_impl\",\"type\":\"address\"},{\"name\":\"controller_impl\",\"type\":\"address\"},{\"name\":\"borrowed_token\",\"type\":\"address\"},{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"price_oracle\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"borrow_apr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lend_apr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pricePerShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pricePerShare\",\"inputs\":[{\"name\":\"is_floor\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_add_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_sub_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"borrowed_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"collateral_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// LlamalendVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use LlamalendVaultMetaData.ABI instead.
var LlamalendVaultABI = LlamalendVaultMetaData.ABI

// LlamalendVault is an auto generated Go binding around an Ethereum contract.
type LlamalendVault struct {
	LlamalendVaultCaller     // Read-only binding to the contract
	LlamalendVaultTransactor // Write-only binding to the contract
	LlamalendVaultFilterer   // Log filterer for contract events
}

// LlamalendVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type LlamalendVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LlamalendVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LlamalendVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LlamalendVaultSession struct {
	Contract     *LlamalendVault   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LlamalendVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LlamalendVaultCallerSession struct {
	Contract *LlamalendVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LlamalendVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LlamalendVaultTransactorSession struct {
	Contract     *LlamalendVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LlamalendVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type LlamalendVaultRaw struct {
	Contract *LlamalendVault // Generic contract binding to access the raw methods on
}

// LlamalendVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LlamalendVaultCallerRaw struct {
	Contract *LlamalendVaultCaller // Generic read-only contract binding to access the raw methods on
}

// LlamalendVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LlamalendVaultTransactorRaw struct {
	Contract *LlamalendVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLlamalendVault creates a new instance of LlamalendVault, bound to a specific deployed contract.
func NewLlamalendVault(address common.Address, backend bind.ContractBackend) (*LlamalendVault, error) {
	contract, err := bindLlamalendVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LlamalendVault{LlamalendVaultCaller: LlamalendVaultCaller{contract: contract}, LlamalendVaultTransactor: LlamalendVaultTransactor{contract: contract}, LlamalendVaultFilterer: LlamalendVaultFilterer{contract: contract}}, nil
}

// NewLlamalendVaultCaller creates a new read-only instance of LlamalendVault, bound to a specific deployed contract.
func NewLlamalendVaultCaller(address common.Address, caller bind.ContractCaller) (*LlamalendVaultCaller, error) {
	contract, err := bindLlamalendVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultCaller{contract: contract}, nil
}

// NewLlamalendVaultTransactor creates a new write-only instance of LlamalendVault, bound to a specific deployed contract.
func NewLlamalendVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*LlamalendVaultTransactor, error) {
	contract, err := bindLlamalendVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultTransactor{contract: contract}, nil
}

// NewLlamalendVaultFilterer creates a new log filterer instance of LlamalendVault, bound to a specific deployed contract.
func NewLlamalendVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*LlamalendVaultFilterer, error) {
	contract, err := bindLlamalendVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultFilterer{contract: contract}, nil
}

// bindLlamalendVault binds a generic wrapper to an already deployed contract.
func bindLlamalendVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LlamalendVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendVault *LlamalendVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendVault.Contract.LlamalendVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendVault *LlamalendVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendVault.Contract.LlamalendVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendVault *LlamalendVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendVault.Contract.LlamalendVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendVault *LlamalendVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendVault *LlamalendVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendVault *LlamalendVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendVault.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) Admin() (common.Address, error) {
	return _LlamalendVault.Contract.Admin(&_LlamalendVault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) Admin() (common.Address, error) {
	return _LlamalendVault.Contract.Admin(&_LlamalendVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.Allowance(&_LlamalendVault.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.Allowance(&_LlamalendVault.CallOpts, arg0, arg1)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) Amm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "amm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) Amm() (common.Address, error) {
	return _LlamalendVault.Contract.Amm(&_LlamalendVault.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) Amm() (common.Address, error) {
	return _LlamalendVault.Contract.Amm(&_LlamalendVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) Asset() (common.Address, error) {
	return _LlamalendVault.Contract.Asset(&_LlamalendVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) Asset() (common.Address, error) {
	return _LlamalendVault.Contract.Asset(&_LlamalendVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.BalanceOf(&_LlamalendVault.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.BalanceOf(&_LlamalendVault.CallOpts, arg0)
}

// BorrowApr is a free data retrieval call binding the contract method 0xbac4daa2.
//
// Solidity: function borrow_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) BorrowApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "borrow_apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowApr is a free data retrieval call binding the contract method 0xbac4daa2.
//
// Solidity: function borrow_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) BorrowApr() (*big.Int, error) {
	return _LlamalendVault.Contract.BorrowApr(&_LlamalendVault.CallOpts)
}

// BorrowApr is a free data retrieval call binding the contract method 0xbac4daa2.
//
// Solidity: function borrow_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) BorrowApr() (*big.Int, error) {
	return _LlamalendVault.Contract.BorrowApr(&_LlamalendVault.CallOpts)
}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) BorrowedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "borrowed_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) BorrowedToken() (common.Address, error) {
	return _LlamalendVault.Contract.BorrowedToken(&_LlamalendVault.CallOpts)
}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) BorrowedToken() (common.Address, error) {
	return _LlamalendVault.Contract.BorrowedToken(&_LlamalendVault.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "collateral_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) CollateralToken() (common.Address, error) {
	return _LlamalendVault.Contract.CollateralToken(&_LlamalendVault.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) CollateralToken() (common.Address, error) {
	return _LlamalendVault.Contract.CollateralToken(&_LlamalendVault.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) Controller() (common.Address, error) {
	return _LlamalendVault.Contract.Controller(&_LlamalendVault.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) Controller() (common.Address, error) {
	return _LlamalendVault.Contract.Controller(&_LlamalendVault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.ConvertToAssets(&_LlamalendVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.ConvertToAssets(&_LlamalendVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.ConvertToShares(&_LlamalendVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.ConvertToShares(&_LlamalendVault.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LlamalendVault *LlamalendVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LlamalendVault *LlamalendVaultSession) Decimals() (uint8, error) {
	return _LlamalendVault.Contract.Decimals(&_LlamalendVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LlamalendVault *LlamalendVaultCallerSession) Decimals() (uint8, error) {
	return _LlamalendVault.Contract.Decimals(&_LlamalendVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) Factory() (common.Address, error) {
	return _LlamalendVault.Contract.Factory(&_LlamalendVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) Factory() (common.Address, error) {
	return _LlamalendVault.Contract.Factory(&_LlamalendVault.CallOpts)
}

// LendApr is a free data retrieval call binding the contract method 0xa980497e.
//
// Solidity: function lend_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) LendApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "lend_apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendApr is a free data retrieval call binding the contract method 0xa980497e.
//
// Solidity: function lend_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) LendApr() (*big.Int, error) {
	return _LlamalendVault.Contract.LendApr(&_LlamalendVault.CallOpts)
}

// LendApr is a free data retrieval call binding the contract method 0xa980497e.
//
// Solidity: function lend_apr() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) LendApr() (*big.Int, error) {
	return _LlamalendVault.Contract.LendApr(&_LlamalendVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxDeposit(&_LlamalendVault.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxDeposit(&_LlamalendVault.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxMint(&_LlamalendVault.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxMint(&_LlamalendVault.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxRedeem(&_LlamalendVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxRedeem(&_LlamalendVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxWithdraw(&_LlamalendVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _LlamalendVault.Contract.MaxWithdraw(&_LlamalendVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LlamalendVault *LlamalendVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LlamalendVault *LlamalendVaultSession) Name() (string, error) {
	return _LlamalendVault.Contract.Name(&_LlamalendVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LlamalendVault *LlamalendVaultCallerSession) Name() (string, error) {
	return _LlamalendVault.Contract.Name(&_LlamalendVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewDeposit(&_LlamalendVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewDeposit(&_LlamalendVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewMint(&_LlamalendVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewMint(&_LlamalendVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewRedeem(&_LlamalendVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewRedeem(&_LlamalendVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewWithdraw(&_LlamalendVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _LlamalendVault.Contract.PreviewWithdraw(&_LlamalendVault.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PricePerShare() (*big.Int, error) {
	return _LlamalendVault.Contract.PricePerShare(&_LlamalendVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PricePerShare() (*big.Int, error) {
	return _LlamalendVault.Contract.PricePerShare(&_LlamalendVault.CallOpts)
}

// PricePerShare0 is a free data retrieval call binding the contract method 0x28c06203.
//
// Solidity: function pricePerShare(bool is_floor) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) PricePerShare0(opts *bind.CallOpts, is_floor bool) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "pricePerShare0", is_floor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare0 is a free data retrieval call binding the contract method 0x28c06203.
//
// Solidity: function pricePerShare(bool is_floor) view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) PricePerShare0(is_floor bool) (*big.Int, error) {
	return _LlamalendVault.Contract.PricePerShare0(&_LlamalendVault.CallOpts, is_floor)
}

// PricePerShare0 is a free data retrieval call binding the contract method 0x28c06203.
//
// Solidity: function pricePerShare(bool is_floor) view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) PricePerShare0(is_floor bool) (*big.Int, error) {
	return _LlamalendVault.Contract.PricePerShare0(&_LlamalendVault.CallOpts, is_floor)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(address)
func (_LlamalendVault *LlamalendVaultCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(address)
func (_LlamalendVault *LlamalendVaultSession) PriceOracle() (common.Address, error) {
	return _LlamalendVault.Contract.PriceOracle(&_LlamalendVault.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(address)
func (_LlamalendVault *LlamalendVaultCallerSession) PriceOracle() (common.Address, error) {
	return _LlamalendVault.Contract.PriceOracle(&_LlamalendVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LlamalendVault *LlamalendVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LlamalendVault *LlamalendVaultSession) Symbol() (string, error) {
	return _LlamalendVault.Contract.Symbol(&_LlamalendVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LlamalendVault *LlamalendVaultCallerSession) Symbol() (string, error) {
	return _LlamalendVault.Contract.Symbol(&_LlamalendVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) TotalAssets() (*big.Int, error) {
	return _LlamalendVault.Contract.TotalAssets(&_LlamalendVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _LlamalendVault.Contract.TotalAssets(&_LlamalendVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) TotalSupply() (*big.Int, error) {
	return _LlamalendVault.Contract.TotalSupply(&_LlamalendVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LlamalendVault *LlamalendVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _LlamalendVault.Contract.TotalSupply(&_LlamalendVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Approve(&_LlamalendVault.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Approve(&_LlamalendVault.TransactOpts, _spender, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "decreaseAllowance", _spender, _sub_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_LlamalendVault *LlamalendVaultSession) DecreaseAllowance(_spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.DecreaseAllowance(&_LlamalendVault.TransactOpts, _spender, _sub_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactorSession) DecreaseAllowance(_spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.DecreaseAllowance(&_LlamalendVault.TransactOpts, _spender, _sub_value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "deposit", assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Deposit(&_LlamalendVault.TransactOpts, assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Deposit(&_LlamalendVault.TransactOpts, assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Deposit0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "deposit0", assets, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Deposit0(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Deposit0(&_LlamalendVault.TransactOpts, assets, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Deposit0(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Deposit0(&_LlamalendVault.TransactOpts, assets, receiver)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "increaseAllowance", _spender, _add_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_LlamalendVault *LlamalendVaultSession) IncreaseAllowance(_spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.IncreaseAllowance(&_LlamalendVault.TransactOpts, _spender, _add_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactorSession) IncreaseAllowance(_spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.IncreaseAllowance(&_LlamalendVault.TransactOpts, _spender, _add_value)
}

// Initialize is a paid mutator transaction binding the contract method 0x43687bba.
//
// Solidity: function initialize(address amm_impl, address controller_impl, address borrowed_token, address collateral_token, uint256 A, uint256 fee, address price_oracle, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount) returns(address, address)
func (_LlamalendVault *LlamalendVaultTransactor) Initialize(opts *bind.TransactOpts, amm_impl common.Address, controller_impl common.Address, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, price_oracle common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "initialize", amm_impl, controller_impl, borrowed_token, collateral_token, A, fee, price_oracle, monetary_policy, loan_discount, liquidation_discount)
}

// Initialize is a paid mutator transaction binding the contract method 0x43687bba.
//
// Solidity: function initialize(address amm_impl, address controller_impl, address borrowed_token, address collateral_token, uint256 A, uint256 fee, address price_oracle, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount) returns(address, address)
func (_LlamalendVault *LlamalendVaultSession) Initialize(amm_impl common.Address, controller_impl common.Address, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, price_oracle common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Initialize(&_LlamalendVault.TransactOpts, amm_impl, controller_impl, borrowed_token, collateral_token, A, fee, price_oracle, monetary_policy, loan_discount, liquidation_discount)
}

// Initialize is a paid mutator transaction binding the contract method 0x43687bba.
//
// Solidity: function initialize(address amm_impl, address controller_impl, address borrowed_token, address collateral_token, uint256 A, uint256 fee, address price_oracle, address monetary_policy, uint256 loan_discount, uint256 liquidation_discount) returns(address, address)
func (_LlamalendVault *LlamalendVaultTransactorSession) Initialize(amm_impl common.Address, controller_impl common.Address, borrowed_token common.Address, collateral_token common.Address, A *big.Int, fee *big.Int, price_oracle common.Address, monetary_policy common.Address, loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Initialize(&_LlamalendVault.TransactOpts, amm_impl, controller_impl, borrowed_token, collateral_token, A, fee, price_oracle, monetary_policy, loan_discount, liquidation_discount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "mint", shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Mint(shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Mint(&_LlamalendVault.TransactOpts, shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Mint(shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Mint(&_LlamalendVault.TransactOpts, shares)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Mint0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "mint0", shares, receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Mint0(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Mint0(&_LlamalendVault.TransactOpts, shares, receiver)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Mint0(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Mint0(&_LlamalendVault.TransactOpts, shares, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "redeem", shares)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Redeem(shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem(&_LlamalendVault.TransactOpts, shares)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Redeem(shares *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem(&_LlamalendVault.TransactOpts, shares)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "redeem0", shares, receiver)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Redeem0(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem0(&_LlamalendVault.TransactOpts, shares, receiver)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x7bde82f2.
//
// Solidity: function redeem(uint256 shares, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Redeem0(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem0(&_LlamalendVault.TransactOpts, shares, receiver)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Redeem1(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "redeem1", shares, receiver, owner)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem1(&_LlamalendVault.TransactOpts, shares, receiver, owner)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Redeem1(&_LlamalendVault.TransactOpts, shares, receiver, owner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Transfer(&_LlamalendVault.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Transfer(&_LlamalendVault.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.TransferFrom(&_LlamalendVault.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LlamalendVault *LlamalendVaultTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.TransferFrom(&_LlamalendVault.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "withdraw", assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Withdraw(assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw(&_LlamalendVault.TransactOpts, assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Withdraw(assets *big.Int) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw(&_LlamalendVault.TransactOpts, assets)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "withdraw0", assets, receiver)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Withdraw0(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw0(&_LlamalendVault.TransactOpts, assets, receiver)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Withdraw0(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw0(&_LlamalendVault.TransactOpts, assets, receiver)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactor) Withdraw1(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.contract.Transact(opts, "withdraw1", assets, receiver, owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw1(&_LlamalendVault.TransactOpts, assets, receiver, owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_LlamalendVault *LlamalendVaultTransactorSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _LlamalendVault.Contract.Withdraw1(&_LlamalendVault.TransactOpts, assets, receiver, owner)
}

// LlamalendVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LlamalendVault contract.
type LlamalendVaultApprovalIterator struct {
	Event *LlamalendVaultApproval // Event containing the contract specifics and raw log

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
func (it *LlamalendVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendVaultApproval)
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
		it.Event = new(LlamalendVaultApproval)
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
func (it *LlamalendVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendVaultApproval represents a Approval event raised by the LlamalendVault contract.
type LlamalendVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LlamalendVault *LlamalendVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LlamalendVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LlamalendVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultApprovalIterator{contract: _LlamalendVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LlamalendVault *LlamalendVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LlamalendVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LlamalendVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendVaultApproval)
				if err := _LlamalendVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LlamalendVault *LlamalendVaultFilterer) ParseApproval(log types.Log) (*LlamalendVaultApproval, error) {
	event := new(LlamalendVaultApproval)
	if err := _LlamalendVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LlamalendVault contract.
type LlamalendVaultDepositIterator struct {
	Event *LlamalendVaultDeposit // Event containing the contract specifics and raw log

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
func (it *LlamalendVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendVaultDeposit)
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
		it.Event = new(LlamalendVaultDeposit)
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
func (it *LlamalendVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendVaultDeposit represents a Deposit event raised by the LlamalendVault contract.
type LlamalendVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*LlamalendVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LlamalendVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultDepositIterator{contract: _LlamalendVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LlamalendVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LlamalendVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendVaultDeposit)
				if err := _LlamalendVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) ParseDeposit(log types.Log) (*LlamalendVaultDeposit, error) {
	event := new(LlamalendVaultDeposit)
	if err := _LlamalendVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LlamalendVault contract.
type LlamalendVaultTransferIterator struct {
	Event *LlamalendVaultTransfer // Event containing the contract specifics and raw log

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
func (it *LlamalendVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendVaultTransfer)
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
		it.Event = new(LlamalendVaultTransfer)
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
func (it *LlamalendVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendVaultTransfer represents a Transfer event raised by the LlamalendVault contract.
type LlamalendVaultTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_LlamalendVault *LlamalendVaultFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*LlamalendVaultTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LlamalendVault.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultTransferIterator{contract: _LlamalendVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_LlamalendVault *LlamalendVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LlamalendVaultTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _LlamalendVault.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendVaultTransfer)
				if err := _LlamalendVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LlamalendVault *LlamalendVaultFilterer) ParseTransfer(log types.Log) (*LlamalendVaultTransfer, error) {
	event := new(LlamalendVaultTransfer)
	if err := _LlamalendVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LlamalendVault contract.
type LlamalendVaultWithdrawIterator struct {
	Event *LlamalendVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *LlamalendVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendVaultWithdraw)
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
		it.Event = new(LlamalendVaultWithdraw)
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
func (it *LlamalendVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendVaultWithdraw represents a Withdraw event raised by the LlamalendVault contract.
type LlamalendVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*LlamalendVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LlamalendVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendVaultWithdrawIterator{contract: _LlamalendVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LlamalendVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _LlamalendVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendVaultWithdraw)
				if err := _LlamalendVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_LlamalendVault *LlamalendVaultFilterer) ParseWithdraw(log types.Log) (*LlamalendVaultWithdraw, error) {
	event := new(LlamalendVaultWithdraw)
	if err := _LlamalendVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
