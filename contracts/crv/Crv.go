// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crv

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

// CrvMetaData contains all meta data concerning the Crv contract.
var CrvMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateMiningParameters\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"time\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rate\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetMinter\",\"inputs\":[{\"type\":\"address\",\"name\":\"minter\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint256\",\"name\":\"_decimals\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"update_mining_parameters\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":148748},{\"name\":\"start_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149603},{\"name\":\"future_epoch_time_write\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":149806},{\"name\":\"available_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4018},{\"name\":\"mintable_in_timeframe\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2216141},{\"name\":\"set_minter\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38698},{\"name\":\"set_admin\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37837},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1421},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1759},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":75139},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111433},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39288},{\"name\":\"mint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":228030},{\"name\":\"burn\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74999},{\"name\":\"set_name\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":178270},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8063},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7116},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1721},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1905},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1781},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"mining_epoch\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"start_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901}]",
}

// CrvABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvMetaData.ABI instead.
var CrvABI = CrvMetaData.ABI

// Crv is an auto generated Go binding around an Ethereum contract.
type Crv struct {
	CrvCaller     // Read-only binding to the contract
	CrvTransactor // Write-only binding to the contract
	CrvFilterer   // Log filterer for contract events
}

// CrvCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvSession struct {
	Contract     *Crv              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvCallerSession struct {
	Contract *CrvCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvTransactorSession struct {
	Contract     *CrvTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvRaw struct {
	Contract *Crv // Generic contract binding to access the raw methods on
}

// CrvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvCallerRaw struct {
	Contract *CrvCaller // Generic read-only contract binding to access the raw methods on
}

// CrvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvTransactorRaw struct {
	Contract *CrvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrv creates a new instance of Crv, bound to a specific deployed contract.
func NewCrv(address common.Address, backend bind.ContractBackend) (*Crv, error) {
	contract, err := bindCrv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crv{CrvCaller: CrvCaller{contract: contract}, CrvTransactor: CrvTransactor{contract: contract}, CrvFilterer: CrvFilterer{contract: contract}}, nil
}

// NewCrvCaller creates a new read-only instance of Crv, bound to a specific deployed contract.
func NewCrvCaller(address common.Address, caller bind.ContractCaller) (*CrvCaller, error) {
	contract, err := bindCrv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvCaller{contract: contract}, nil
}

// NewCrvTransactor creates a new write-only instance of Crv, bound to a specific deployed contract.
func NewCrvTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvTransactor, error) {
	contract, err := bindCrv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvTransactor{contract: contract}, nil
}

// NewCrvFilterer creates a new log filterer instance of Crv, bound to a specific deployed contract.
func NewCrvFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvFilterer, error) {
	contract, err := bindCrv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvFilterer{contract: contract}, nil
}

// bindCrv binds a generic wrapper to an already deployed contract.
func bindCrv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crv *CrvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crv.Contract.CrvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crv *CrvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crv.Contract.CrvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crv *CrvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crv.Contract.CrvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crv *CrvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crv *CrvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crv *CrvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crv.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Crv *CrvCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Crv *CrvSession) Admin() (common.Address, error) {
	return _Crv.Contract.Admin(&_Crv.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Crv *CrvCallerSession) Admin() (common.Address, error) {
	return _Crv.Contract.Admin(&_Crv.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Crv *CrvCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Crv *CrvSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Crv.Contract.Allowance(&_Crv.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Crv *CrvCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Crv.Contract.Allowance(&_Crv.CallOpts, _owner, _spender)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Crv *CrvCaller) AvailableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "available_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Crv *CrvSession) AvailableSupply() (*big.Int, error) {
	return _Crv.Contract.AvailableSupply(&_Crv.CallOpts)
}

// AvailableSupply is a free data retrieval call binding the contract method 0x24f92a25.
//
// Solidity: function available_supply() view returns(uint256)
func (_Crv *CrvCallerSession) AvailableSupply() (*big.Int, error) {
	return _Crv.Contract.AvailableSupply(&_Crv.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Crv *CrvCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Crv *CrvSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Crv.Contract.BalanceOf(&_Crv.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Crv *CrvCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Crv.Contract.BalanceOf(&_Crv.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Crv *CrvCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Crv *CrvSession) Decimals() (*big.Int, error) {
	return _Crv.Contract.Decimals(&_Crv.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Crv *CrvCallerSession) Decimals() (*big.Int, error) {
	return _Crv.Contract.Decimals(&_Crv.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Crv *CrvCaller) MiningEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "mining_epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Crv *CrvSession) MiningEpoch() (*big.Int, error) {
	return _Crv.Contract.MiningEpoch(&_Crv.CallOpts)
}

// MiningEpoch is a free data retrieval call binding the contract method 0xf9a40bf6.
//
// Solidity: function mining_epoch() view returns(int128)
func (_Crv *CrvCallerSession) MiningEpoch() (*big.Int, error) {
	return _Crv.Contract.MiningEpoch(&_Crv.CallOpts)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Crv *CrvCaller) MintableInTimeframe(opts *bind.CallOpts, start *big.Int, end *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "mintable_in_timeframe", start, end)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Crv *CrvSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _Crv.Contract.MintableInTimeframe(&_Crv.CallOpts, start, end)
}

// MintableInTimeframe is a free data retrieval call binding the contract method 0xd725a9ca.
//
// Solidity: function mintable_in_timeframe(uint256 start, uint256 end) view returns(uint256)
func (_Crv *CrvCallerSession) MintableInTimeframe(start *big.Int, end *big.Int) (*big.Int, error) {
	return _Crv.Contract.MintableInTimeframe(&_Crv.CallOpts, start, end)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Crv *CrvCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Crv *CrvSession) Minter() (common.Address, error) {
	return _Crv.Contract.Minter(&_Crv.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Crv *CrvCallerSession) Minter() (common.Address, error) {
	return _Crv.Contract.Minter(&_Crv.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crv *CrvCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crv *CrvSession) Name() (string, error) {
	return _Crv.Contract.Name(&_Crv.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crv *CrvCallerSession) Name() (string, error) {
	return _Crv.Contract.Name(&_Crv.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crv *CrvCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crv *CrvSession) Rate() (*big.Int, error) {
	return _Crv.Contract.Rate(&_Crv.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crv *CrvCallerSession) Rate() (*big.Int, error) {
	return _Crv.Contract.Rate(&_Crv.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Crv *CrvCaller) StartEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "start_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Crv *CrvSession) StartEpochTime() (*big.Int, error) {
	return _Crv.Contract.StartEpochTime(&_Crv.CallOpts)
}

// StartEpochTime is a free data retrieval call binding the contract method 0x7375be26.
//
// Solidity: function start_epoch_time() view returns(uint256)
func (_Crv *CrvCallerSession) StartEpochTime() (*big.Int, error) {
	return _Crv.Contract.StartEpochTime(&_Crv.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crv *CrvCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crv *CrvSession) Symbol() (string, error) {
	return _Crv.Contract.Symbol(&_Crv.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crv *CrvCallerSession) Symbol() (string, error) {
	return _Crv.Contract.Symbol(&_Crv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crv *CrvCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crv.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crv *CrvSession) TotalSupply() (*big.Int, error) {
	return _Crv.Contract.TotalSupply(&_Crv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crv *CrvCallerSession) TotalSupply() (*big.Int, error) {
	return _Crv.Contract.TotalSupply(&_Crv.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Crv *CrvTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Crv *CrvSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Approve(&_Crv.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Crv *CrvTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Approve(&_Crv.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Crv *CrvTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Crv *CrvSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Burn(&_Crv.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool)
func (_Crv *CrvTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Burn(&_Crv.TransactOpts, _value)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Crv *CrvTransactor) FutureEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "future_epoch_time_write")
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Crv *CrvSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _Crv.Contract.FutureEpochTimeWrite(&_Crv.TransactOpts)
}

// FutureEpochTimeWrite is a paid mutator transaction binding the contract method 0xb26b238e.
//
// Solidity: function future_epoch_time_write() returns(uint256)
func (_Crv *CrvTransactorSession) FutureEpochTimeWrite() (*types.Transaction, error) {
	return _Crv.Contract.FutureEpochTimeWrite(&_Crv.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Crv *CrvSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Mint(&_Crv.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Mint(&_Crv.TransactOpts, _to, _value)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Crv *CrvTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "set_admin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Crv *CrvSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Crv.Contract.SetAdmin(&_Crv.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0xe9333fab.
//
// Solidity: function set_admin(address _admin) returns()
func (_Crv *CrvTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Crv.Contract.SetAdmin(&_Crv.TransactOpts, _admin)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Crv *CrvTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Crv *CrvSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Crv.Contract.SetMinter(&_Crv.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Crv *CrvTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Crv.Contract.SetMinter(&_Crv.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Crv *CrvTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Crv *CrvSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Crv.Contract.SetName(&_Crv.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Crv *CrvTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Crv.Contract.SetName(&_Crv.TransactOpts, _name, _symbol)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Crv *CrvTransactor) StartEpochTimeWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "start_epoch_time_write")
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Crv *CrvSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _Crv.Contract.StartEpochTimeWrite(&_Crv.TransactOpts)
}

// StartEpochTimeWrite is a paid mutator transaction binding the contract method 0xadc4cf43.
//
// Solidity: function start_epoch_time_write() returns(uint256)
func (_Crv *CrvTransactorSession) StartEpochTimeWrite() (*types.Transaction, error) {
	return _Crv.Contract.StartEpochTimeWrite(&_Crv.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Crv *CrvSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Transfer(&_Crv.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.Transfer(&_Crv.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Crv *CrvSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.TransferFrom(&_Crv.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Crv *CrvTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Crv.Contract.TransferFrom(&_Crv.TransactOpts, _from, _to, _value)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Crv *CrvTransactor) UpdateMiningParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crv.contract.Transact(opts, "update_mining_parameters")
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Crv *CrvSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _Crv.Contract.UpdateMiningParameters(&_Crv.TransactOpts)
}

// UpdateMiningParameters is a paid mutator transaction binding the contract method 0xd43b40fa.
//
// Solidity: function update_mining_parameters() returns()
func (_Crv *CrvTransactorSession) UpdateMiningParameters() (*types.Transaction, error) {
	return _Crv.Contract.UpdateMiningParameters(&_Crv.TransactOpts)
}

// CrvApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Crv contract.
type CrvApprovalIterator struct {
	Event *CrvApproval // Event containing the contract specifics and raw log

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
func (it *CrvApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvApproval)
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
		it.Event = new(CrvApproval)
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
func (it *CrvApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvApproval represents a Approval event raised by the Crv contract.
type CrvApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Crv *CrvFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CrvApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Crv.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrvApprovalIterator{contract: _Crv.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Crv *CrvFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrvApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Crv.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvApproval)
				if err := _Crv.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Crv *CrvFilterer) ParseApproval(log types.Log) (*CrvApproval, error) {
	event := new(CrvApproval)
	if err := _Crv.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the Crv contract.
type CrvSetAdminIterator struct {
	Event *CrvSetAdmin // Event containing the contract specifics and raw log

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
func (it *CrvSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvSetAdmin)
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
		it.Event = new(CrvSetAdmin)
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
func (it *CrvSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvSetAdmin represents a SetAdmin event raised by the Crv contract.
type CrvSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_Crv *CrvFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*CrvSetAdminIterator, error) {

	logs, sub, err := _Crv.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &CrvSetAdminIterator{contract: _Crv.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_Crv *CrvFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *CrvSetAdmin) (event.Subscription, error) {

	logs, sub, err := _Crv.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvSetAdmin)
				if err := _Crv.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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
func (_Crv *CrvFilterer) ParseSetAdmin(log types.Log) (*CrvSetAdmin, error) {
	event := new(CrvSetAdmin)
	if err := _Crv.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the Crv contract.
type CrvSetMinterIterator struct {
	Event *CrvSetMinter // Event containing the contract specifics and raw log

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
func (it *CrvSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvSetMinter)
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
		it.Event = new(CrvSetMinter)
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
func (it *CrvSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvSetMinter represents a SetMinter event raised by the Crv contract.
type CrvSetMinter struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Crv *CrvFilterer) FilterSetMinter(opts *bind.FilterOpts) (*CrvSetMinterIterator, error) {

	logs, sub, err := _Crv.contract.FilterLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return &CrvSetMinterIterator{contract: _Crv.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Crv *CrvFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *CrvSetMinter) (event.Subscription, error) {

	logs, sub, err := _Crv.contract.WatchLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvSetMinter)
				if err := _Crv.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0xcec52196e972044edde8689a1b608e459c5946b7f3e5c8cd3d6d8e126d422e1c.
//
// Solidity: event SetMinter(address minter)
func (_Crv *CrvFilterer) ParseSetMinter(log types.Log) (*CrvSetMinter, error) {
	event := new(CrvSetMinter)
	if err := _Crv.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Crv contract.
type CrvTransferIterator struct {
	Event *CrvTransfer // Event containing the contract specifics and raw log

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
func (it *CrvTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvTransfer)
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
		it.Event = new(CrvTransfer)
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
func (it *CrvTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvTransfer represents a Transfer event raised by the Crv contract.
type CrvTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Crv *CrvFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CrvTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Crv.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CrvTransferIterator{contract: _Crv.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Crv *CrvFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrvTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Crv.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvTransfer)
				if err := _Crv.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Crv *CrvFilterer) ParseTransfer(log types.Log) (*CrvTransfer, error) {
	event := new(CrvTransfer)
	if err := _Crv.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUpdateMiningParametersIterator is returned from FilterUpdateMiningParameters and is used to iterate over the raw logs and unpacked data for UpdateMiningParameters events raised by the Crv contract.
type CrvUpdateMiningParametersIterator struct {
	Event *CrvUpdateMiningParameters // Event containing the contract specifics and raw log

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
func (it *CrvUpdateMiningParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUpdateMiningParameters)
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
		it.Event = new(CrvUpdateMiningParameters)
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
func (it *CrvUpdateMiningParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUpdateMiningParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUpdateMiningParameters represents a UpdateMiningParameters event raised by the Crv contract.
type CrvUpdateMiningParameters struct {
	Time   *big.Int
	Rate   *big.Int
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateMiningParameters is a free log retrieval operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Crv *CrvFilterer) FilterUpdateMiningParameters(opts *bind.FilterOpts) (*CrvUpdateMiningParametersIterator, error) {

	logs, sub, err := _Crv.contract.FilterLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return &CrvUpdateMiningParametersIterator{contract: _Crv.contract, event: "UpdateMiningParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateMiningParameters is a free log subscription operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Crv *CrvFilterer) WatchUpdateMiningParameters(opts *bind.WatchOpts, sink chan<- *CrvUpdateMiningParameters) (event.Subscription, error) {

	logs, sub, err := _Crv.contract.WatchLogs(opts, "UpdateMiningParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUpdateMiningParameters)
				if err := _Crv.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
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

// ParseUpdateMiningParameters is a log parse operation binding the contract event 0x27e46362a1e6129b6dd539c984ce739291a97128dfcaeca1255e8ac83abd9441.
//
// Solidity: event UpdateMiningParameters(uint256 time, uint256 rate, uint256 supply)
func (_Crv *CrvFilterer) ParseUpdateMiningParameters(log types.Log) (*CrvUpdateMiningParameters, error) {
	event := new(CrvUpdateMiningParameters)
	if err := _Crv.contract.UnpackLog(event, "UpdateMiningParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
