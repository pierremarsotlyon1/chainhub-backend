// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votiumVECrv

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

// VotiumVECrvMetaData contains all meta data concerning the VotiumVECrv contract.
var VotiumVECrvMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"Listed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"ModifiedTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_week\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"NewReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"Unlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"UpdatedDistributor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTeam\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_week\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"depositReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gc\",\"outputs\":[{\"internalType\":\"contractGController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"listTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"modifyTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"unlistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"updateDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"updateFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"updateFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gc\",\"type\":\"address\"}],\"name\":\"updateGaugeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"week\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotiumVECrvABI is the input ABI used to generate the binding from.
// Deprecated: Use VotiumVECrvMetaData.ABI instead.
var VotiumVECrvABI = VotiumVECrvMetaData.ABI

// VotiumVECrv is an auto generated Go binding around an Ethereum contract.
type VotiumVECrv struct {
	VotiumVECrvCaller     // Read-only binding to the contract
	VotiumVECrvTransactor // Write-only binding to the contract
	VotiumVECrvFilterer   // Log filterer for contract events
}

// VotiumVECrvCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotiumVECrvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVECrvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotiumVECrvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVECrvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotiumVECrvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVECrvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotiumVECrvSession struct {
	Contract     *VotiumVECrv      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotiumVECrvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotiumVECrvCallerSession struct {
	Contract *VotiumVECrvCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VotiumVECrvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotiumVECrvTransactorSession struct {
	Contract     *VotiumVECrvTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VotiumVECrvRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotiumVECrvRaw struct {
	Contract *VotiumVECrv // Generic contract binding to access the raw methods on
}

// VotiumVECrvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotiumVECrvCallerRaw struct {
	Contract *VotiumVECrvCaller // Generic read-only contract binding to access the raw methods on
}

// VotiumVECrvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotiumVECrvTransactorRaw struct {
	Contract *VotiumVECrvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotiumVECrv creates a new instance of VotiumVECrv, bound to a specific deployed contract.
func NewVotiumVECrv(address common.Address, backend bind.ContractBackend) (*VotiumVECrv, error) {
	contract, err := bindVotiumVECrv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrv{VotiumVECrvCaller: VotiumVECrvCaller{contract: contract}, VotiumVECrvTransactor: VotiumVECrvTransactor{contract: contract}, VotiumVECrvFilterer: VotiumVECrvFilterer{contract: contract}}, nil
}

// NewVotiumVECrvCaller creates a new read-only instance of VotiumVECrv, bound to a specific deployed contract.
func NewVotiumVECrvCaller(address common.Address, caller bind.ContractCaller) (*VotiumVECrvCaller, error) {
	contract, err := bindVotiumVECrv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvCaller{contract: contract}, nil
}

// NewVotiumVECrvTransactor creates a new write-only instance of VotiumVECrv, bound to a specific deployed contract.
func NewVotiumVECrvTransactor(address common.Address, transactor bind.ContractTransactor) (*VotiumVECrvTransactor, error) {
	contract, err := bindVotiumVECrv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvTransactor{contract: contract}, nil
}

// NewVotiumVECrvFilterer creates a new log filterer instance of VotiumVECrv, bound to a specific deployed contract.
func NewVotiumVECrvFilterer(address common.Address, filterer bind.ContractFilterer) (*VotiumVECrvFilterer, error) {
	contract, err := bindVotiumVECrv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvFilterer{contract: contract}, nil
}

// bindVotiumVECrv binds a generic wrapper to an already deployed contract.
func bindVotiumVECrv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotiumVECrvABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVECrv *VotiumVECrvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVECrv.Contract.VotiumVECrvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVECrv *VotiumVECrvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.VotiumVECrvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVECrv *VotiumVECrvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.VotiumVECrvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVECrv *VotiumVECrvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVECrv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVECrv *VotiumVECrvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVECrv *VotiumVECrvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.contract.Transact(opts, method, params...)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvSession) DENOMINATOR() (*big.Int, error) {
	return _VotiumVECrv.Contract.DENOMINATOR(&_VotiumVECrv.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCallerSession) DENOMINATOR() (*big.Int, error) {
	return _VotiumVECrv.Contract.DENOMINATOR(&_VotiumVECrv.CallOpts)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvCaller) ApprovedTeam(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "approvedTeam", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvSession) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVECrv.Contract.ApprovedTeam(&_VotiumVECrv.CallOpts, arg0)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvCallerSession) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVECrv.Contract.ApprovedTeam(&_VotiumVECrv.CallOpts, arg0)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVECrv *VotiumVECrvCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVECrv *VotiumVECrvSession) Distributor() (common.Address, error) {
	return _VotiumVECrv.Contract.Distributor(&_VotiumVECrv.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVECrv *VotiumVECrvCallerSession) Distributor() (common.Address, error) {
	return _VotiumVECrv.Contract.Distributor(&_VotiumVECrv.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVECrv *VotiumVECrvCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVECrv *VotiumVECrvSession) FeeAddress() (common.Address, error) {
	return _VotiumVECrv.Contract.FeeAddress(&_VotiumVECrv.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVECrv *VotiumVECrvCallerSession) FeeAddress() (common.Address, error) {
	return _VotiumVECrv.Contract.FeeAddress(&_VotiumVECrv.CallOpts)
}

// Gc is a free data retrieval call binding the contract method 0x1b71d85f.
//
// Solidity: function gc() view returns(address)
func (_VotiumVECrv *VotiumVECrvCaller) Gc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "gc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gc is a free data retrieval call binding the contract method 0x1b71d85f.
//
// Solidity: function gc() view returns(address)
func (_VotiumVECrv *VotiumVECrvSession) Gc() (common.Address, error) {
	return _VotiumVECrv.Contract.Gc(&_VotiumVECrv.CallOpts)
}

// Gc is a free data retrieval call binding the contract method 0x1b71d85f.
//
// Solidity: function gc() view returns(address)
func (_VotiumVECrv *VotiumVECrvCallerSession) Gc() (common.Address, error) {
	return _VotiumVECrv.Contract.Gc(&_VotiumVECrv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVECrv *VotiumVECrvCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVECrv *VotiumVECrvSession) Owner() (common.Address, error) {
	return _VotiumVECrv.Contract.Owner(&_VotiumVECrv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVECrv *VotiumVECrvCallerSession) Owner() (common.Address, error) {
	return _VotiumVECrv.Contract.Owner(&_VotiumVECrv.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvSession) PlatformFee() (*big.Int, error) {
	return _VotiumVECrv.Contract.PlatformFee(&_VotiumVECrv.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCallerSession) PlatformFee() (*big.Int, error) {
	return _VotiumVECrv.Contract.PlatformFee(&_VotiumVECrv.CallOpts)
}

// TokenListed is a free data retrieval call binding the contract method 0xa7ab30e1.
//
// Solidity: function tokenListed(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvCaller) TokenListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "tokenListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenListed is a free data retrieval call binding the contract method 0xa7ab30e1.
//
// Solidity: function tokenListed(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvSession) TokenListed(arg0 common.Address) (bool, error) {
	return _VotiumVECrv.Contract.TokenListed(&_VotiumVECrv.CallOpts, arg0)
}

// TokenListed is a free data retrieval call binding the contract method 0xa7ab30e1.
//
// Solidity: function tokenListed(address ) view returns(bool)
func (_VotiumVECrv *VotiumVECrvCallerSession) TokenListed(arg0 common.Address) (bool, error) {
	return _VotiumVECrv.Contract.TokenListed(&_VotiumVECrv.CallOpts, arg0)
}

// Week is a free data retrieval call binding the contract method 0x4995b458.
//
// Solidity: function week() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCaller) Week(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVECrv.contract.Call(opts, &out, "week")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Week is a free data retrieval call binding the contract method 0x4995b458.
//
// Solidity: function week() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvSession) Week() (*big.Int, error) {
	return _VotiumVECrv.Contract.Week(&_VotiumVECrv.CallOpts)
}

// Week is a free data retrieval call binding the contract method 0x4995b458.
//
// Solidity: function week() view returns(uint256)
func (_VotiumVECrv *VotiumVECrvCallerSession) Week() (*big.Int, error) {
	return _VotiumVECrv.Contract.Week(&_VotiumVECrv.CallOpts)
}

// DepositReward is a paid mutator transaction binding the contract method 0x0d2f4f31.
//
// Solidity: function depositReward(address _token, uint256 _amount, uint256 _week, address _gauge) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) DepositReward(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _week *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "depositReward", _token, _amount, _week, _gauge)
}

// DepositReward is a paid mutator transaction binding the contract method 0x0d2f4f31.
//
// Solidity: function depositReward(address _token, uint256 _amount, uint256 _week, address _gauge) returns()
func (_VotiumVECrv *VotiumVECrvSession) DepositReward(_token common.Address, _amount *big.Int, _week *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.DepositReward(&_VotiumVECrv.TransactOpts, _token, _amount, _week, _gauge)
}

// DepositReward is a paid mutator transaction binding the contract method 0x0d2f4f31.
//
// Solidity: function depositReward(address _token, uint256 _amount, uint256 _week, address _gauge) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) DepositReward(_token common.Address, _amount *big.Int, _week *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.DepositReward(&_VotiumVECrv.TransactOpts, _token, _amount, _week, _gauge)
}

// ListToken is a paid mutator transaction binding the contract method 0x1fc1e25f.
//
// Solidity: function listToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) ListToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "listToken", _token)
}

// ListToken is a paid mutator transaction binding the contract method 0x1fc1e25f.
//
// Solidity: function listToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvSession) ListToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ListToken(&_VotiumVECrv.TransactOpts, _token)
}

// ListToken is a paid mutator transaction binding the contract method 0x1fc1e25f.
//
// Solidity: function listToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) ListToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ListToken(&_VotiumVECrv.TransactOpts, _token)
}

// ListTokens is a paid mutator transaction binding the contract method 0x9fb99fc3.
//
// Solidity: function listTokens(address[] _tokens) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) ListTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "listTokens", _tokens)
}

// ListTokens is a paid mutator transaction binding the contract method 0x9fb99fc3.
//
// Solidity: function listTokens(address[] _tokens) returns()
func (_VotiumVECrv *VotiumVECrvSession) ListTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ListTokens(&_VotiumVECrv.TransactOpts, _tokens)
}

// ListTokens is a paid mutator transaction binding the contract method 0x9fb99fc3.
//
// Solidity: function listTokens(address[] _tokens) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) ListTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ListTokens(&_VotiumVECrv.TransactOpts, _tokens)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) ModifyTeam(opts *bind.TransactOpts, _member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "modifyTeam", _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVECrv *VotiumVECrvSession) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ModifyTeam(&_VotiumVECrv.TransactOpts, _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.ModifyTeam(&_VotiumVECrv.TransactOpts, _member, _approval)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVECrv *VotiumVECrvSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.TransferOwnership(&_VotiumVECrv.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.TransferOwnership(&_VotiumVECrv.TransactOpts, newOwner)
}

// UnlistToken is a paid mutator transaction binding the contract method 0xabcb9934.
//
// Solidity: function unlistToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) UnlistToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "unlistToken", _token)
}

// UnlistToken is a paid mutator transaction binding the contract method 0xabcb9934.
//
// Solidity: function unlistToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvSession) UnlistToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UnlistToken(&_VotiumVECrv.TransactOpts, _token)
}

// UnlistToken is a paid mutator transaction binding the contract method 0xabcb9934.
//
// Solidity: function unlistToken(address _token) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) UnlistToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UnlistToken(&_VotiumVECrv.TransactOpts, _token)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) UpdateDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "updateDistributor", _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVECrv *VotiumVECrvSession) UpdateDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateDistributor(&_VotiumVECrv.TransactOpts, _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) UpdateDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateDistributor(&_VotiumVECrv.TransactOpts, _distributor)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) UpdateFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "updateFeeAddress", _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVECrv *VotiumVECrvSession) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateFeeAddress(&_VotiumVECrv.TransactOpts, _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateFeeAddress(&_VotiumVECrv.TransactOpts, _feeAddress)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) UpdateFeeAmount(opts *bind.TransactOpts, _feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "updateFeeAmount", _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVECrv *VotiumVECrvSession) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateFeeAmount(&_VotiumVECrv.TransactOpts, _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateFeeAmount(&_VotiumVECrv.TransactOpts, _feeAmount)
}

// UpdateGaugeController is a paid mutator transaction binding the contract method 0xc325c0d2.
//
// Solidity: function updateGaugeController(address _gc) returns()
func (_VotiumVECrv *VotiumVECrvTransactor) UpdateGaugeController(opts *bind.TransactOpts, _gc common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.contract.Transact(opts, "updateGaugeController", _gc)
}

// UpdateGaugeController is a paid mutator transaction binding the contract method 0xc325c0d2.
//
// Solidity: function updateGaugeController(address _gc) returns()
func (_VotiumVECrv *VotiumVECrvSession) UpdateGaugeController(_gc common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateGaugeController(&_VotiumVECrv.TransactOpts, _gc)
}

// UpdateGaugeController is a paid mutator transaction binding the contract method 0xc325c0d2.
//
// Solidity: function updateGaugeController(address _gc) returns()
func (_VotiumVECrv *VotiumVECrvTransactorSession) UpdateGaugeController(_gc common.Address) (*types.Transaction, error) {
	return _VotiumVECrv.Contract.UpdateGaugeController(&_VotiumVECrv.TransactOpts, _gc)
}

// VotiumVECrvListedIterator is returned from FilterListed and is used to iterate over the raw logs and unpacked data for Listed events raised by the VotiumVECrv contract.
type VotiumVECrvListedIterator struct {
	Event *VotiumVECrvListed // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvListed)
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
		it.Event = new(VotiumVECrvListed)
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
func (it *VotiumVECrvListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvListed represents a Listed event raised by the VotiumVECrv contract.
type VotiumVECrvListed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterListed is a free log retrieval operation binding the contract event 0xfb3fdb4942f7aa0b8ecdf8508875e7f6c8142bb7870f0455b87a9f093608bc82.
//
// Solidity: event Listed(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterListed(opts *bind.FilterOpts) (*VotiumVECrvListedIterator, error) {

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "Listed")
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvListedIterator{contract: _VotiumVECrv.contract, event: "Listed", logs: logs, sub: sub}, nil
}

// WatchListed is a free log subscription operation binding the contract event 0xfb3fdb4942f7aa0b8ecdf8508875e7f6c8142bb7870f0455b87a9f093608bc82.
//
// Solidity: event Listed(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchListed(opts *bind.WatchOpts, sink chan<- *VotiumVECrvListed) (event.Subscription, error) {

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "Listed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvListed)
				if err := _VotiumVECrv.contract.UnpackLog(event, "Listed", log); err != nil {
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

// ParseListed is a log parse operation binding the contract event 0xfb3fdb4942f7aa0b8ecdf8508875e7f6c8142bb7870f0455b87a9f093608bc82.
//
// Solidity: event Listed(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseListed(log types.Log) (*VotiumVECrvListed, error) {
	event := new(VotiumVECrvListed)
	if err := _VotiumVECrv.contract.UnpackLog(event, "Listed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvModifiedTeamIterator is returned from FilterModifiedTeam and is used to iterate over the raw logs and unpacked data for ModifiedTeam events raised by the VotiumVECrv contract.
type VotiumVECrvModifiedTeamIterator struct {
	Event *VotiumVECrvModifiedTeam // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvModifiedTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvModifiedTeam)
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
		it.Event = new(VotiumVECrvModifiedTeam)
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
func (it *VotiumVECrvModifiedTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvModifiedTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvModifiedTeam represents a ModifiedTeam event raised by the VotiumVECrv contract.
type VotiumVECrvModifiedTeam struct {
	Member   common.Address
	Approval bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifiedTeam is a free log retrieval operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterModifiedTeam(opts *bind.FilterOpts) (*VotiumVECrvModifiedTeamIterator, error) {

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvModifiedTeamIterator{contract: _VotiumVECrv.contract, event: "ModifiedTeam", logs: logs, sub: sub}, nil
}

// WatchModifiedTeam is a free log subscription operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchModifiedTeam(opts *bind.WatchOpts, sink chan<- *VotiumVECrvModifiedTeam) (event.Subscription, error) {

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvModifiedTeam)
				if err := _VotiumVECrv.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
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

// ParseModifiedTeam is a log parse operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseModifiedTeam(log types.Log) (*VotiumVECrvModifiedTeam, error) {
	event := new(VotiumVECrvModifiedTeam)
	if err := _VotiumVECrv.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvNewRewardIterator is returned from FilterNewReward and is used to iterate over the raw logs and unpacked data for NewReward events raised by the VotiumVECrv contract.
type VotiumVECrvNewRewardIterator struct {
	Event *VotiumVECrvNewReward // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvNewRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvNewReward)
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
		it.Event = new(VotiumVECrvNewReward)
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
func (it *VotiumVECrvNewRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvNewRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvNewReward represents a NewReward event raised by the VotiumVECrv contract.
type VotiumVECrvNewReward struct {
	Token  common.Address
	Amount *big.Int
	Week   *big.Int
	Gauge  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewReward is a free log retrieval operation binding the contract event 0x51c8cd367a987b8c2f652c101ea7076ec8e4dfd33c4c77bb80e018e7143b6512.
//
// Solidity: event NewReward(address indexed _token, uint256 _amount, uint256 indexed _week, address indexed _gauge)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterNewReward(opts *bind.FilterOpts, _token []common.Address, _week []*big.Int, _gauge []common.Address) (*VotiumVECrvNewRewardIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	var _weekRule []interface{}
	for _, _weekItem := range _week {
		_weekRule = append(_weekRule, _weekItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "NewReward", _tokenRule, _weekRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvNewRewardIterator{contract: _VotiumVECrv.contract, event: "NewReward", logs: logs, sub: sub}, nil
}

// WatchNewReward is a free log subscription operation binding the contract event 0x51c8cd367a987b8c2f652c101ea7076ec8e4dfd33c4c77bb80e018e7143b6512.
//
// Solidity: event NewReward(address indexed _token, uint256 _amount, uint256 indexed _week, address indexed _gauge)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchNewReward(opts *bind.WatchOpts, sink chan<- *VotiumVECrvNewReward, _token []common.Address, _week []*big.Int, _gauge []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	var _weekRule []interface{}
	for _, _weekItem := range _week {
		_weekRule = append(_weekRule, _weekItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "NewReward", _tokenRule, _weekRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvNewReward)
				if err := _VotiumVECrv.contract.UnpackLog(event, "NewReward", log); err != nil {
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

// ParseNewReward is a log parse operation binding the contract event 0x51c8cd367a987b8c2f652c101ea7076ec8e4dfd33c4c77bb80e018e7143b6512.
//
// Solidity: event NewReward(address indexed _token, uint256 _amount, uint256 indexed _week, address indexed _gauge)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseNewReward(log types.Log) (*VotiumVECrvNewReward, error) {
	event := new(VotiumVECrvNewReward)
	if err := _VotiumVECrv.contract.UnpackLog(event, "NewReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotiumVECrv contract.
type VotiumVECrvOwnershipTransferredIterator struct {
	Event *VotiumVECrvOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvOwnershipTransferred)
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
		it.Event = new(VotiumVECrvOwnershipTransferred)
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
func (it *VotiumVECrvOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvOwnershipTransferred represents a OwnershipTransferred event raised by the VotiumVECrv contract.
type VotiumVECrvOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotiumVECrvOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvOwnershipTransferredIterator{contract: _VotiumVECrv.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotiumVECrvOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvOwnershipTransferred)
				if err := _VotiumVECrv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseOwnershipTransferred(log types.Log) (*VotiumVECrvOwnershipTransferred, error) {
	event := new(VotiumVECrvOwnershipTransferred)
	if err := _VotiumVECrv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvUnlistedIterator is returned from FilterUnlisted and is used to iterate over the raw logs and unpacked data for Unlisted events raised by the VotiumVECrv contract.
type VotiumVECrvUnlistedIterator struct {
	Event *VotiumVECrvUnlisted // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvUnlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvUnlisted)
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
		it.Event = new(VotiumVECrvUnlisted)
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
func (it *VotiumVECrvUnlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvUnlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvUnlisted represents a Unlisted event raised by the VotiumVECrv contract.
type VotiumVECrvUnlisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlisted is a free log retrieval operation binding the contract event 0xc48322c2e165c864176005884983dd81c87a9fc8258ca6a1967d17b26cd8d0e3.
//
// Solidity: event Unlisted(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterUnlisted(opts *bind.FilterOpts) (*VotiumVECrvUnlistedIterator, error) {

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "Unlisted")
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvUnlistedIterator{contract: _VotiumVECrv.contract, event: "Unlisted", logs: logs, sub: sub}, nil
}

// WatchUnlisted is a free log subscription operation binding the contract event 0xc48322c2e165c864176005884983dd81c87a9fc8258ca6a1967d17b26cd8d0e3.
//
// Solidity: event Unlisted(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchUnlisted(opts *bind.WatchOpts, sink chan<- *VotiumVECrvUnlisted) (event.Subscription, error) {

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "Unlisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvUnlisted)
				if err := _VotiumVECrv.contract.UnpackLog(event, "Unlisted", log); err != nil {
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

// ParseUnlisted is a log parse operation binding the contract event 0xc48322c2e165c864176005884983dd81c87a9fc8258ca6a1967d17b26cd8d0e3.
//
// Solidity: event Unlisted(address _token)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseUnlisted(log types.Log) (*VotiumVECrvUnlisted, error) {
	event := new(VotiumVECrvUnlisted)
	if err := _VotiumVECrv.contract.UnpackLog(event, "Unlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvUpdatedDistributorIterator is returned from FilterUpdatedDistributor and is used to iterate over the raw logs and unpacked data for UpdatedDistributor events raised by the VotiumVECrv contract.
type VotiumVECrvUpdatedDistributorIterator struct {
	Event *VotiumVECrvUpdatedDistributor // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvUpdatedDistributorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvUpdatedDistributor)
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
		it.Event = new(VotiumVECrvUpdatedDistributor)
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
func (it *VotiumVECrvUpdatedDistributorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvUpdatedDistributorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvUpdatedDistributor represents a UpdatedDistributor event raised by the VotiumVECrv contract.
type VotiumVECrvUpdatedDistributor struct {
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDistributor is a free log retrieval operation binding the contract event 0x9941240ae138e71bb3f963c4f6fd96e99fb386a8e63563343f9294902bdbf744.
//
// Solidity: event UpdatedDistributor(address _distributor)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterUpdatedDistributor(opts *bind.FilterOpts) (*VotiumVECrvUpdatedDistributorIterator, error) {

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "UpdatedDistributor")
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvUpdatedDistributorIterator{contract: _VotiumVECrv.contract, event: "UpdatedDistributor", logs: logs, sub: sub}, nil
}

// WatchUpdatedDistributor is a free log subscription operation binding the contract event 0x9941240ae138e71bb3f963c4f6fd96e99fb386a8e63563343f9294902bdbf744.
//
// Solidity: event UpdatedDistributor(address _distributor)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchUpdatedDistributor(opts *bind.WatchOpts, sink chan<- *VotiumVECrvUpdatedDistributor) (event.Subscription, error) {

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "UpdatedDistributor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvUpdatedDistributor)
				if err := _VotiumVECrv.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
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

// ParseUpdatedDistributor is a log parse operation binding the contract event 0x9941240ae138e71bb3f963c4f6fd96e99fb386a8e63563343f9294902bdbf744.
//
// Solidity: event UpdatedDistributor(address _distributor)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseUpdatedDistributor(log types.Log) (*VotiumVECrvUpdatedDistributor, error) {
	event := new(VotiumVECrvUpdatedDistributor)
	if err := _VotiumVECrv.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVECrvUpdatedFeeIterator is returned from FilterUpdatedFee and is used to iterate over the raw logs and unpacked data for UpdatedFee events raised by the VotiumVECrv contract.
type VotiumVECrvUpdatedFeeIterator struct {
	Event *VotiumVECrvUpdatedFee // Event containing the contract specifics and raw log

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
func (it *VotiumVECrvUpdatedFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVECrvUpdatedFee)
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
		it.Event = new(VotiumVECrvUpdatedFee)
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
func (it *VotiumVECrvUpdatedFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVECrvUpdatedFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVECrvUpdatedFee represents a UpdatedFee event raised by the VotiumVECrv contract.
type VotiumVECrvUpdatedFee struct {
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFee is a free log retrieval operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVECrv *VotiumVECrvFilterer) FilterUpdatedFee(opts *bind.FilterOpts) (*VotiumVECrvUpdatedFeeIterator, error) {

	logs, sub, err := _VotiumVECrv.contract.FilterLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return &VotiumVECrvUpdatedFeeIterator{contract: _VotiumVECrv.contract, event: "UpdatedFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedFee is a free log subscription operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVECrv *VotiumVECrvFilterer) WatchUpdatedFee(opts *bind.WatchOpts, sink chan<- *VotiumVECrvUpdatedFee) (event.Subscription, error) {

	logs, sub, err := _VotiumVECrv.contract.WatchLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVECrvUpdatedFee)
				if err := _VotiumVECrv.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
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

// ParseUpdatedFee is a log parse operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVECrv *VotiumVECrvFilterer) ParseUpdatedFee(log types.Log) (*VotiumVECrvUpdatedFee, error) {
	event := new(VotiumVECrvUpdatedFee)
	if err := _VotiumVECrv.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
