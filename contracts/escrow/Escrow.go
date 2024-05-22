// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package escrow

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

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"prevSupply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"string\",\"name\":\"_version\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37597},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38497},{\"name\":\"commit_smart_wallet_checker\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36307},{\"name\":\"apply_smart_wallet_checker\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37095},{\"name\":\"get_last_user_slope\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2569},{\"name\":\"user_point_history__ts\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_idx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1672},{\"name\":\"locked__end\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1593},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37052342},{\"name\":\"deposit_for\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74279891},{\"name\":\"create_lock\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281465},{\"name\":\"increase_amount\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280830},{\"name\":\"increase_unlock_time\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281578},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37223566},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOfAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":514333},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":812560},{\"name\":\"changeController\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_newController\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36907},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"locked\",\"outputs\":[{\"type\":\"int128\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3359},{\"name\":\"epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5550},{\"name\":\"user_point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6079},{\"name\":\"user_point_epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2175},{\"name\":\"slope_changes\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2166},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2081},{\"name\":\"transfersEnabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8543},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7596},{\"name\":\"version\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7626},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2231},{\"name\":\"future_smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2261},{\"name\":\"smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351}]",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Escrow *EscrowCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Escrow *EscrowSession) Admin() (common.Address, error) {
	return _Escrow.Contract.Admin(&_Escrow.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Escrow *EscrowCallerSession) Admin() (common.Address, error) {
	return _Escrow.Contract.Admin(&_Escrow.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Escrow *EscrowCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Escrow *EscrowSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.BalanceOf(&_Escrow.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Escrow *EscrowCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.BalanceOf(&_Escrow.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Escrow *EscrowCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "balanceOf0", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Escrow *EscrowSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Escrow.Contract.BalanceOf0(&_Escrow.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Escrow *EscrowCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Escrow.Contract.BalanceOf0(&_Escrow.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Escrow *EscrowCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Escrow *EscrowSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Escrow.Contract.BalanceOfAt(&_Escrow.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Escrow *EscrowCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Escrow.Contract.BalanceOfAt(&_Escrow.CallOpts, addr, _block)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Escrow *EscrowCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Escrow *EscrowSession) Controller() (common.Address, error) {
	return _Escrow.Contract.Controller(&_Escrow.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Escrow *EscrowCallerSession) Controller() (common.Address, error) {
	return _Escrow.Contract.Controller(&_Escrow.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Escrow *EscrowCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Escrow *EscrowSession) Decimals() (*big.Int, error) {
	return _Escrow.Contract.Decimals(&_Escrow.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Escrow *EscrowCallerSession) Decimals() (*big.Int, error) {
	return _Escrow.Contract.Decimals(&_Escrow.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Escrow *EscrowCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Escrow *EscrowSession) Epoch() (*big.Int, error) {
	return _Escrow.Contract.Epoch(&_Escrow.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Escrow *EscrowCallerSession) Epoch() (*big.Int, error) {
	return _Escrow.Contract.Epoch(&_Escrow.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Escrow *EscrowCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Escrow *EscrowSession) FutureAdmin() (common.Address, error) {
	return _Escrow.Contract.FutureAdmin(&_Escrow.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Escrow *EscrowCallerSession) FutureAdmin() (common.Address, error) {
	return _Escrow.Contract.FutureAdmin(&_Escrow.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Escrow *EscrowCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "future_smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Escrow *EscrowSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Escrow.Contract.FutureSmartWalletChecker(&_Escrow.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Escrow *EscrowCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Escrow.Contract.FutureSmartWalletChecker(&_Escrow.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Escrow *EscrowCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Escrow *EscrowSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.GetLastUserSlope(&_Escrow.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Escrow *EscrowCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.GetLastUserSlope(&_Escrow.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Escrow *EscrowCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "locked", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Escrow *EscrowSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Escrow.Contract.Locked(&_Escrow.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_Escrow *EscrowCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _Escrow.Contract.Locked(&_Escrow.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Escrow *EscrowCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Escrow *EscrowSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.LockedEnd(&_Escrow.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Escrow *EscrowCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Escrow.Contract.LockedEnd(&_Escrow.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Escrow *EscrowCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Escrow *EscrowSession) Name() (string, error) {
	return _Escrow.Contract.Name(&_Escrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Escrow *EscrowCallerSession) Name() (string, error) {
	return _Escrow.Contract.Name(&_Escrow.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Escrow.Contract.PointHistory(&_Escrow.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Escrow.Contract.PointHistory(&_Escrow.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Escrow *EscrowCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Escrow *EscrowSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.SlopeChanges(&_Escrow.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Escrow *EscrowCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Escrow.Contract.SlopeChanges(&_Escrow.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Escrow *EscrowCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Escrow *EscrowSession) SmartWalletChecker() (common.Address, error) {
	return _Escrow.Contract.SmartWalletChecker(&_Escrow.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Escrow *EscrowCallerSession) SmartWalletChecker() (common.Address, error) {
	return _Escrow.Contract.SmartWalletChecker(&_Escrow.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Escrow *EscrowCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Escrow *EscrowSession) Supply() (*big.Int, error) {
	return _Escrow.Contract.Supply(&_Escrow.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Escrow *EscrowCallerSession) Supply() (*big.Int, error) {
	return _Escrow.Contract.Supply(&_Escrow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Escrow *EscrowCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Escrow *EscrowSession) Symbol() (string, error) {
	return _Escrow.Contract.Symbol(&_Escrow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Escrow *EscrowCallerSession) Symbol() (string, error) {
	return _Escrow.Contract.Symbol(&_Escrow.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Escrow *EscrowCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Escrow *EscrowSession) Token() (common.Address, error) {
	return _Escrow.Contract.Token(&_Escrow.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Escrow *EscrowCallerSession) Token() (common.Address, error) {
	return _Escrow.Contract.Token(&_Escrow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Escrow *EscrowCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Escrow *EscrowSession) TotalSupply() (*big.Int, error) {
	return _Escrow.Contract.TotalSupply(&_Escrow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Escrow *EscrowCallerSession) TotalSupply() (*big.Int, error) {
	return _Escrow.Contract.TotalSupply(&_Escrow.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Escrow *EscrowCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "totalSupply0", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Escrow *EscrowSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Escrow.Contract.TotalSupply0(&_Escrow.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Escrow *EscrowCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Escrow.Contract.TotalSupply0(&_Escrow.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Escrow *EscrowCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Escrow *EscrowSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Escrow.Contract.TotalSupplyAt(&_Escrow.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Escrow *EscrowCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Escrow.Contract.TotalSupplyAt(&_Escrow.CallOpts, _block)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Escrow *EscrowCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "transfersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Escrow *EscrowSession) TransfersEnabled() (bool, error) {
	return _Escrow.Contract.TransfersEnabled(&_Escrow.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_Escrow *EscrowCallerSession) TransfersEnabled() (bool, error) {
	return _Escrow.Contract.TransfersEnabled(&_Escrow.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Escrow *EscrowCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Escrow *EscrowSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Escrow.Contract.UserPointEpoch(&_Escrow.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Escrow *EscrowCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Escrow.Contract.UserPointEpoch(&_Escrow.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Escrow.Contract.UserPointHistory(&_Escrow.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_Escrow *EscrowCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _Escrow.Contract.UserPointHistory(&_Escrow.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Escrow *EscrowCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Escrow *EscrowSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Escrow.Contract.UserPointHistoryTs(&_Escrow.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Escrow *EscrowCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Escrow.Contract.UserPointHistoryTs(&_Escrow.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Escrow *EscrowCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Escrow *EscrowSession) Version() (string, error) {
	return _Escrow.Contract.Version(&_Escrow.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Escrow *EscrowCallerSession) Version() (string, error) {
	return _Escrow.Contract.Version(&_Escrow.CallOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Escrow *EscrowTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Escrow *EscrowSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Escrow.Contract.ApplySmartWalletChecker(&_Escrow.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Escrow *EscrowTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Escrow.Contract.ApplySmartWalletChecker(&_Escrow.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Escrow *EscrowTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Escrow *EscrowSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.ApplyTransferOwnership(&_Escrow.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Escrow *EscrowTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.ApplyTransferOwnership(&_Escrow.TransactOpts)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Escrow *EscrowTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Escrow *EscrowSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.ChangeController(&_Escrow.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_Escrow *EscrowTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.ChangeController(&_Escrow.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Escrow *EscrowTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Escrow *EscrowSession) Checkpoint() (*types.Transaction, error) {
	return _Escrow.Contract.Checkpoint(&_Escrow.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Escrow *EscrowTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Escrow.Contract.Checkpoint(&_Escrow.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Escrow *EscrowTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Escrow *EscrowSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.CommitSmartWalletChecker(&_Escrow.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Escrow *EscrowTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.CommitSmartWalletChecker(&_Escrow.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Escrow *EscrowTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Escrow *EscrowSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.CommitTransferOwnership(&_Escrow.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Escrow *EscrowTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.CommitTransferOwnership(&_Escrow.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Escrow *EscrowTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Escrow *EscrowSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CreateLock(&_Escrow.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Escrow *EscrowTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.CreateLock(&_Escrow.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Escrow *EscrowTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Escrow *EscrowSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.DepositFor(&_Escrow.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Escrow *EscrowTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.DepositFor(&_Escrow.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Escrow *EscrowTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Escrow *EscrowSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.IncreaseAmount(&_Escrow.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Escrow *EscrowTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.IncreaseAmount(&_Escrow.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Escrow *EscrowTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Escrow *EscrowSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.IncreaseUnlockTime(&_Escrow.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Escrow *EscrowTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.IncreaseUnlockTime(&_Escrow.TransactOpts, _unlock_time)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Escrow *EscrowTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Escrow *EscrowSession) Withdraw() (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Escrow *EscrowTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts)
}

// EscrowApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the Escrow contract.
type EscrowApplyOwnershipIterator struct {
	Event *EscrowApplyOwnership // Event containing the contract specifics and raw log

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
func (it *EscrowApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowApplyOwnership)
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
		it.Event = new(EscrowApplyOwnership)
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
func (it *EscrowApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowApplyOwnership represents a ApplyOwnership event raised by the Escrow contract.
type EscrowApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Escrow *EscrowFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*EscrowApplyOwnershipIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &EscrowApplyOwnershipIterator{contract: _Escrow.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Escrow *EscrowFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *EscrowApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowApplyOwnership)
				if err := _Escrow.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Escrow *EscrowFilterer) ParseApplyOwnership(log types.Log) (*EscrowApplyOwnership, error) {
	event := new(EscrowApplyOwnership)
	if err := _Escrow.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the Escrow contract.
type EscrowCommitOwnershipIterator struct {
	Event *EscrowCommitOwnership // Event containing the contract specifics and raw log

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
func (it *EscrowCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowCommitOwnership)
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
		it.Event = new(EscrowCommitOwnership)
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
func (it *EscrowCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowCommitOwnership represents a CommitOwnership event raised by the Escrow contract.
type EscrowCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Escrow *EscrowFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*EscrowCommitOwnershipIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &EscrowCommitOwnershipIterator{contract: _Escrow.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Escrow *EscrowFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *EscrowCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowCommitOwnership)
				if err := _Escrow.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Escrow *EscrowFilterer) ParseCommitOwnership(log types.Log) (*EscrowCommitOwnership, error) {
	event := new(EscrowCommitOwnership)
	if err := _Escrow.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Escrow contract.
type EscrowDepositIterator struct {
	Event *EscrowDeposit // Event containing the contract specifics and raw log

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
func (it *EscrowDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowDeposit)
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
		it.Event = new(EscrowDeposit)
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
func (it *EscrowDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowDeposit represents a Deposit event raised by the Escrow contract.
type EscrowDeposit struct {
	Provider common.Address
	Value    *big.Int
	Locktime *big.Int
	Arg3     *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Escrow *EscrowFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*EscrowDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &EscrowDepositIterator{contract: _Escrow.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Escrow *EscrowFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EscrowDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowDeposit)
				if err := _Escrow.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Escrow *EscrowFilterer) ParseDeposit(log types.Log) (*EscrowDeposit, error) {
	event := new(EscrowDeposit)
	if err := _Escrow.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Escrow contract.
type EscrowSupplyIterator struct {
	Event *EscrowSupply // Event containing the contract specifics and raw log

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
func (it *EscrowSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowSupply)
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
		it.Event = new(EscrowSupply)
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
func (it *EscrowSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowSupply represents a Supply event raised by the Escrow contract.
type EscrowSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Escrow *EscrowFilterer) FilterSupply(opts *bind.FilterOpts) (*EscrowSupplyIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &EscrowSupplyIterator{contract: _Escrow.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Escrow *EscrowFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *EscrowSupply) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowSupply)
				if err := _Escrow.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Escrow *EscrowFilterer) ParseSupply(log types.Log) (*EscrowSupply, error) {
	event := new(EscrowSupply)
	if err := _Escrow.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Escrow contract.
type EscrowWithdrawIterator struct {
	Event *EscrowWithdraw // Event containing the contract specifics and raw log

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
func (it *EscrowWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowWithdraw)
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
		it.Event = new(EscrowWithdraw)
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
func (it *EscrowWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowWithdraw represents a Withdraw event raised by the Escrow contract.
type EscrowWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Escrow *EscrowFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*EscrowWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &EscrowWithdrawIterator{contract: _Escrow.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Escrow *EscrowFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EscrowWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowWithdraw)
				if err := _Escrow.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Escrow *EscrowFilterer) ParseWithdraw(log types.Log) (*EscrowWithdraw, error) {
	event := new(EscrowWithdraw)
	if err := _Escrow.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
