// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votiumVLCVXOld

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

// VotiumVLCVXOldMetaData contains all meta data concerning the VotiumVLCVXOld contract.
var VotiumVLCVXOldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_proposal\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_choiceIndex\",\"type\":\"uint256\"}],\"name\":\"Bribed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_proposal\",\"type\":\"bytes32\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"ModifiedTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"UpdatedDistributor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_requireWhitelist\",\"type\":\"bool\"}],\"name\":\"WhitelistRequirement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"approveDelegationVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTeam\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delegationHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_proposal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_choiceIndex\",\"type\":\"uint256\"}],\"name\":\"depositBribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxIndex\",\"type\":\"uint256\"}],\"name\":\"initiateProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isWinningSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"modifyTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposalInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_requireWhitelist\",\"type\":\"bool\"}],\"name\":\"setWhitelistRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"updateDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"updateFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"updateFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"whitelistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"whitelistTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VotiumVLCVXOldABI is the input ABI used to generate the binding from.
// Deprecated: Use VotiumVLCVXOldMetaData.ABI instead.
var VotiumVLCVXOldABI = VotiumVLCVXOldMetaData.ABI

// VotiumVLCVXOld is an auto generated Go binding around an Ethereum contract.
type VotiumVLCVXOld struct {
	VotiumVLCVXOldCaller     // Read-only binding to the contract
	VotiumVLCVXOldTransactor // Write-only binding to the contract
	VotiumVLCVXOldFilterer   // Log filterer for contract events
}

// VotiumVLCVXOldCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotiumVLCVXOldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXOldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotiumVLCVXOldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXOldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotiumVLCVXOldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXOldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotiumVLCVXOldSession struct {
	Contract     *VotiumVLCVXOld   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotiumVLCVXOldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotiumVLCVXOldCallerSession struct {
	Contract *VotiumVLCVXOldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VotiumVLCVXOldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotiumVLCVXOldTransactorSession struct {
	Contract     *VotiumVLCVXOldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VotiumVLCVXOldRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotiumVLCVXOldRaw struct {
	Contract *VotiumVLCVXOld // Generic contract binding to access the raw methods on
}

// VotiumVLCVXOldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotiumVLCVXOldCallerRaw struct {
	Contract *VotiumVLCVXOldCaller // Generic read-only contract binding to access the raw methods on
}

// VotiumVLCVXOldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotiumVLCVXOldTransactorRaw struct {
	Contract *VotiumVLCVXOldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotiumVLCVXOld creates a new instance of VotiumVLCVXOld, bound to a specific deployed contract.
func NewVotiumVLCVXOld(address common.Address, backend bind.ContractBackend) (*VotiumVLCVXOld, error) {
	contract, err := bindVotiumVLCVXOld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOld{VotiumVLCVXOldCaller: VotiumVLCVXOldCaller{contract: contract}, VotiumVLCVXOldTransactor: VotiumVLCVXOldTransactor{contract: contract}, VotiumVLCVXOldFilterer: VotiumVLCVXOldFilterer{contract: contract}}, nil
}

// NewVotiumVLCVXOldCaller creates a new read-only instance of VotiumVLCVXOld, bound to a specific deployed contract.
func NewVotiumVLCVXOldCaller(address common.Address, caller bind.ContractCaller) (*VotiumVLCVXOldCaller, error) {
	contract, err := bindVotiumVLCVXOld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldCaller{contract: contract}, nil
}

// NewVotiumVLCVXOldTransactor creates a new write-only instance of VotiumVLCVXOld, bound to a specific deployed contract.
func NewVotiumVLCVXOldTransactor(address common.Address, transactor bind.ContractTransactor) (*VotiumVLCVXOldTransactor, error) {
	contract, err := bindVotiumVLCVXOld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldTransactor{contract: contract}, nil
}

// NewVotiumVLCVXOldFilterer creates a new log filterer instance of VotiumVLCVXOld, bound to a specific deployed contract.
func NewVotiumVLCVXOldFilterer(address common.Address, filterer bind.ContractFilterer) (*VotiumVLCVXOldFilterer, error) {
	contract, err := bindVotiumVLCVXOld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldFilterer{contract: contract}, nil
}

// bindVotiumVLCVXOld binds a generic wrapper to an already deployed contract.
func bindVotiumVLCVXOld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotiumVLCVXOldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVLCVXOld *VotiumVLCVXOldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVLCVXOld.Contract.VotiumVLCVXOldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVLCVXOld *VotiumVLCVXOldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.VotiumVLCVXOldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVLCVXOld *VotiumVLCVXOldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.VotiumVLCVXOldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVLCVXOld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.contract.Transact(opts, method, params...)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) DENOMINATOR() (*big.Int, error) {
	return _VotiumVLCVXOld.Contract.DENOMINATOR(&_VotiumVLCVXOld.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) DENOMINATOR() (*big.Int, error) {
	return _VotiumVLCVXOld.Contract.DENOMINATOR(&_VotiumVLCVXOld.CallOpts)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) ApprovedTeam(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "approvedTeam", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXOld.Contract.ApprovedTeam(&_VotiumVLCVXOld.CallOpts, arg0)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXOld.Contract.ApprovedTeam(&_VotiumVLCVXOld.CallOpts, arg0)
}

// DelegationHash is a free data retrieval call binding the contract method 0x14ee0a5e.
//
// Solidity: function delegationHash(bytes32 ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) DelegationHash(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "delegationHash", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationHash is a free data retrieval call binding the contract method 0x14ee0a5e.
//
// Solidity: function delegationHash(bytes32 ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) DelegationHash(arg0 [32]byte) (bool, error) {
	return _VotiumVLCVXOld.Contract.DelegationHash(&_VotiumVLCVXOld.CallOpts, arg0)
}

// DelegationHash is a free data retrieval call binding the contract method 0x14ee0a5e.
//
// Solidity: function delegationHash(bytes32 ) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) DelegationHash(arg0 [32]byte) (bool, error) {
	return _VotiumVLCVXOld.Contract.DelegationHash(&_VotiumVLCVXOld.CallOpts, arg0)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) FeeAddress() (common.Address, error) {
	return _VotiumVLCVXOld.Contract.FeeAddress(&_VotiumVLCVXOld.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) FeeAddress() (common.Address, error) {
	return _VotiumVLCVXOld.Contract.FeeAddress(&_VotiumVLCVXOld.CallOpts)
}

// IsWinningSignature is a free data retrieval call binding the contract method 0x5d3a8999.
//
// Solidity: function isWinningSignature(bytes32 _hash, bytes _signature) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) IsWinningSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "isWinningSignature", _hash, _signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWinningSignature is a free data retrieval call binding the contract method 0x5d3a8999.
//
// Solidity: function isWinningSignature(bytes32 _hash, bytes _signature) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) IsWinningSignature(_hash [32]byte, _signature []byte) (bool, error) {
	return _VotiumVLCVXOld.Contract.IsWinningSignature(&_VotiumVLCVXOld.CallOpts, _hash, _signature)
}

// IsWinningSignature is a free data retrieval call binding the contract method 0x5d3a8999.
//
// Solidity: function isWinningSignature(bytes32 _hash, bytes _signature) view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) IsWinningSignature(_hash [32]byte, _signature []byte) (bool, error) {
	return _VotiumVLCVXOld.Contract.IsWinningSignature(&_VotiumVLCVXOld.CallOpts, _hash, _signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) Owner() (common.Address, error) {
	return _VotiumVLCVXOld.Contract.Owner(&_VotiumVLCVXOld.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) Owner() (common.Address, error) {
	return _VotiumVLCVXOld.Contract.Owner(&_VotiumVLCVXOld.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) PlatformFee() (*big.Int, error) {
	return _VotiumVLCVXOld.Contract.PlatformFee(&_VotiumVLCVXOld.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) PlatformFee() (*big.Int, error) {
	return _VotiumVLCVXOld.Contract.PlatformFee(&_VotiumVLCVXOld.CallOpts)
}

// ProposalInfo is a free data retrieval call binding the contract method 0x53da67f6.
//
// Solidity: function proposalInfo(bytes32 ) view returns(uint256 deadline, uint256 maxIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) ProposalInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Deadline *big.Int
	MaxIndex *big.Int
}, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "proposalInfo", arg0)

	outstruct := new(struct {
		Deadline *big.Int
		MaxIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deadline = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalInfo is a free data retrieval call binding the contract method 0x53da67f6.
//
// Solidity: function proposalInfo(bytes32 ) view returns(uint256 deadline, uint256 maxIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) ProposalInfo(arg0 [32]byte) (struct {
	Deadline *big.Int
	MaxIndex *big.Int
}, error) {
	return _VotiumVLCVXOld.Contract.ProposalInfo(&_VotiumVLCVXOld.CallOpts, arg0)
}

// ProposalInfo is a free data retrieval call binding the contract method 0x53da67f6.
//
// Solidity: function proposalInfo(bytes32 ) view returns(uint256 deadline, uint256 maxIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) ProposalInfo(arg0 [32]byte) (struct {
	Deadline *big.Int
	MaxIndex *big.Int
}, error) {
	return _VotiumVLCVXOld.Contract.ProposalInfo(&_VotiumVLCVXOld.CallOpts, arg0)
}

// RequireWhitelist is a free data retrieval call binding the contract method 0x856734c4.
//
// Solidity: function requireWhitelist() view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) RequireWhitelist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "requireWhitelist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireWhitelist is a free data retrieval call binding the contract method 0x856734c4.
//
// Solidity: function requireWhitelist() view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) RequireWhitelist() (bool, error) {
	return _VotiumVLCVXOld.Contract.RequireWhitelist(&_VotiumVLCVXOld.CallOpts)
}

// RequireWhitelist is a free data retrieval call binding the contract method 0x856734c4.
//
// Solidity: function requireWhitelist() view returns(bool)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) RequireWhitelist() (bool, error) {
	return _VotiumVLCVXOld.Contract.RequireWhitelist(&_VotiumVLCVXOld.CallOpts)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(bool whitelist, address distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldCaller) TokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Whitelist   bool
	Distributor common.Address
}, error) {
	var out []interface{}
	err := _VotiumVLCVXOld.contract.Call(opts, &out, "tokenInfo", arg0)

	outstruct := new(struct {
		Whitelist   bool
		Distributor common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Whitelist = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Distributor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(bool whitelist, address distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) TokenInfo(arg0 common.Address) (struct {
	Whitelist   bool
	Distributor common.Address
}, error) {
	return _VotiumVLCVXOld.Contract.TokenInfo(&_VotiumVLCVXOld.CallOpts, arg0)
}

// TokenInfo is a free data retrieval call binding the contract method 0xf5dab711.
//
// Solidity: function tokenInfo(address ) view returns(bool whitelist, address distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldCallerSession) TokenInfo(arg0 common.Address) (struct {
	Whitelist   bool
	Distributor common.Address
}, error) {
	return _VotiumVLCVXOld.Contract.TokenInfo(&_VotiumVLCVXOld.CallOpts, arg0)
}

// ApproveDelegationVote is a paid mutator transaction binding the contract method 0x92800f86.
//
// Solidity: function approveDelegationVote(bytes32 _hash) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) ApproveDelegationVote(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "approveDelegationVote", _hash)
}

// ApproveDelegationVote is a paid mutator transaction binding the contract method 0x92800f86.
//
// Solidity: function approveDelegationVote(bytes32 _hash) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) ApproveDelegationVote(_hash [32]byte) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.ApproveDelegationVote(&_VotiumVLCVXOld.TransactOpts, _hash)
}

// ApproveDelegationVote is a paid mutator transaction binding the contract method 0x92800f86.
//
// Solidity: function approveDelegationVote(bytes32 _hash) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) ApproveDelegationVote(_hash [32]byte) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.ApproveDelegationVote(&_VotiumVLCVXOld.TransactOpts, _hash)
}

// DepositBribe is a paid mutator transaction binding the contract method 0xf7b5d4a0.
//
// Solidity: function depositBribe(address _token, uint256 _amount, bytes32 _proposal, uint256 _choiceIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) DepositBribe(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _proposal [32]byte, _choiceIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "depositBribe", _token, _amount, _proposal, _choiceIndex)
}

// DepositBribe is a paid mutator transaction binding the contract method 0xf7b5d4a0.
//
// Solidity: function depositBribe(address _token, uint256 _amount, bytes32 _proposal, uint256 _choiceIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) DepositBribe(_token common.Address, _amount *big.Int, _proposal [32]byte, _choiceIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.DepositBribe(&_VotiumVLCVXOld.TransactOpts, _token, _amount, _proposal, _choiceIndex)
}

// DepositBribe is a paid mutator transaction binding the contract method 0xf7b5d4a0.
//
// Solidity: function depositBribe(address _token, uint256 _amount, bytes32 _proposal, uint256 _choiceIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) DepositBribe(_token common.Address, _amount *big.Int, _proposal [32]byte, _choiceIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.DepositBribe(&_VotiumVLCVXOld.TransactOpts, _token, _amount, _proposal, _choiceIndex)
}

// InitiateProposal is a paid mutator transaction binding the contract method 0x986ec135.
//
// Solidity: function initiateProposal(bytes32 _proposal, uint256 _deadline, uint256 _maxIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) InitiateProposal(opts *bind.TransactOpts, _proposal [32]byte, _deadline *big.Int, _maxIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "initiateProposal", _proposal, _deadline, _maxIndex)
}

// InitiateProposal is a paid mutator transaction binding the contract method 0x986ec135.
//
// Solidity: function initiateProposal(bytes32 _proposal, uint256 _deadline, uint256 _maxIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) InitiateProposal(_proposal [32]byte, _deadline *big.Int, _maxIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.InitiateProposal(&_VotiumVLCVXOld.TransactOpts, _proposal, _deadline, _maxIndex)
}

// InitiateProposal is a paid mutator transaction binding the contract method 0x986ec135.
//
// Solidity: function initiateProposal(bytes32 _proposal, uint256 _deadline, uint256 _maxIndex) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) InitiateProposal(_proposal [32]byte, _deadline *big.Int, _maxIndex *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.InitiateProposal(&_VotiumVLCVXOld.TransactOpts, _proposal, _deadline, _maxIndex)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) ModifyTeam(opts *bind.TransactOpts, _member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "modifyTeam", _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.ModifyTeam(&_VotiumVLCVXOld.TransactOpts, _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.ModifyTeam(&_VotiumVLCVXOld.TransactOpts, _member, _approval)
}

// SetWhitelistRequired is a paid mutator transaction binding the contract method 0xb5c12f4d.
//
// Solidity: function setWhitelistRequired(bool _requireWhitelist) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) SetWhitelistRequired(opts *bind.TransactOpts, _requireWhitelist bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "setWhitelistRequired", _requireWhitelist)
}

// SetWhitelistRequired is a paid mutator transaction binding the contract method 0xb5c12f4d.
//
// Solidity: function setWhitelistRequired(bool _requireWhitelist) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) SetWhitelistRequired(_requireWhitelist bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.SetWhitelistRequired(&_VotiumVLCVXOld.TransactOpts, _requireWhitelist)
}

// SetWhitelistRequired is a paid mutator transaction binding the contract method 0xb5c12f4d.
//
// Solidity: function setWhitelistRequired(bool _requireWhitelist) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) SetWhitelistRequired(_requireWhitelist bool) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.SetWhitelistRequired(&_VotiumVLCVXOld.TransactOpts, _requireWhitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.TransferOwnership(&_VotiumVLCVXOld.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.TransferOwnership(&_VotiumVLCVXOld.TransactOpts, newOwner)
}

// TransferToDistributor is a paid mutator transaction binding the contract method 0xb4c3ed94.
//
// Solidity: function transferToDistributor(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) TransferToDistributor(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "transferToDistributor", _token)
}

// TransferToDistributor is a paid mutator transaction binding the contract method 0xb4c3ed94.
//
// Solidity: function transferToDistributor(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) TransferToDistributor(_token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.TransferToDistributor(&_VotiumVLCVXOld.TransactOpts, _token)
}

// TransferToDistributor is a paid mutator transaction binding the contract method 0xb4c3ed94.
//
// Solidity: function transferToDistributor(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) TransferToDistributor(_token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.TransferToDistributor(&_VotiumVLCVXOld.TransactOpts, _token)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0x1fac1379.
//
// Solidity: function updateDistributor(address _token, address _distributor) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) UpdateDistributor(opts *bind.TransactOpts, _token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "updateDistributor", _token, _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0x1fac1379.
//
// Solidity: function updateDistributor(address _token, address _distributor) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) UpdateDistributor(_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateDistributor(&_VotiumVLCVXOld.TransactOpts, _token, _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0x1fac1379.
//
// Solidity: function updateDistributor(address _token, address _distributor) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) UpdateDistributor(_token common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateDistributor(&_VotiumVLCVXOld.TransactOpts, _token, _distributor)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) UpdateFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "updateFeeAddress", _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateFeeAddress(&_VotiumVLCVXOld.TransactOpts, _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateFeeAddress(&_VotiumVLCVXOld.TransactOpts, _feeAddress)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) UpdateFeeAmount(opts *bind.TransactOpts, _feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "updateFeeAmount", _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateFeeAmount(&_VotiumVLCVXOld.TransactOpts, _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.UpdateFeeAmount(&_VotiumVLCVXOld.TransactOpts, _feeAmount)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) WhitelistToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "whitelistToken", _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.WhitelistToken(&_VotiumVLCVXOld.TransactOpts, _token)
}

// WhitelistToken is a paid mutator transaction binding the contract method 0x6247f6f2.
//
// Solidity: function whitelistToken(address _token) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) WhitelistToken(_token common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.WhitelistToken(&_VotiumVLCVXOld.TransactOpts, _token)
}

// WhitelistTokens is a paid mutator transaction binding the contract method 0x4cb1e9e3.
//
// Solidity: function whitelistTokens(address[] _tokens) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactor) WhitelistTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.contract.Transact(opts, "whitelistTokens", _tokens)
}

// WhitelistTokens is a paid mutator transaction binding the contract method 0x4cb1e9e3.
//
// Solidity: function whitelistTokens(address[] _tokens) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldSession) WhitelistTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.WhitelistTokens(&_VotiumVLCVXOld.TransactOpts, _tokens)
}

// WhitelistTokens is a paid mutator transaction binding the contract method 0x4cb1e9e3.
//
// Solidity: function whitelistTokens(address[] _tokens) returns()
func (_VotiumVLCVXOld *VotiumVLCVXOldTransactorSession) WhitelistTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXOld.Contract.WhitelistTokens(&_VotiumVLCVXOld.TransactOpts, _tokens)
}

// VotiumVLCVXOldBribedIterator is returned from FilterBribed and is used to iterate over the raw logs and unpacked data for Bribed events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldBribedIterator struct {
	Event *VotiumVLCVXOldBribed // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldBribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldBribed)
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
		it.Event = new(VotiumVLCVXOldBribed)
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
func (it *VotiumVLCVXOldBribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldBribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldBribed represents a Bribed event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldBribed struct {
	Token       common.Address
	Amount      *big.Int
	Proposal    [32]byte
	ChoiceIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBribed is a free log retrieval operation binding the contract event 0x74bd0b58587a15767427910140bcf99db1ef7f905cb0a2983a72cd2033954227.
//
// Solidity: event Bribed(address _token, uint256 _amount, bytes32 indexed _proposal, uint256 _choiceIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterBribed(opts *bind.FilterOpts, _proposal [][32]byte) (*VotiumVLCVXOldBribedIterator, error) {

	var _proposalRule []interface{}
	for _, _proposalItem := range _proposal {
		_proposalRule = append(_proposalRule, _proposalItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "Bribed", _proposalRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldBribedIterator{contract: _VotiumVLCVXOld.contract, event: "Bribed", logs: logs, sub: sub}, nil
}

// WatchBribed is a free log subscription operation binding the contract event 0x74bd0b58587a15767427910140bcf99db1ef7f905cb0a2983a72cd2033954227.
//
// Solidity: event Bribed(address _token, uint256 _amount, bytes32 indexed _proposal, uint256 _choiceIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchBribed(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldBribed, _proposal [][32]byte) (event.Subscription, error) {

	var _proposalRule []interface{}
	for _, _proposalItem := range _proposal {
		_proposalRule = append(_proposalRule, _proposalItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "Bribed", _proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldBribed)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Bribed", log); err != nil {
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

// ParseBribed is a log parse operation binding the contract event 0x74bd0b58587a15767427910140bcf99db1ef7f905cb0a2983a72cd2033954227.
//
// Solidity: event Bribed(address _token, uint256 _amount, bytes32 indexed _proposal, uint256 _choiceIndex)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseBribed(log types.Log) (*VotiumVLCVXOldBribed, error) {
	event := new(VotiumVLCVXOldBribed)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Bribed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldInitiatedIterator struct {
	Event *VotiumVLCVXOldInitiated // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldInitiated)
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
		it.Event = new(VotiumVLCVXOldInitiated)
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
func (it *VotiumVLCVXOldInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldInitiated represents a Initiated event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldInitiated struct {
	Proposal [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0x08007a3b331cd9bd3d1d3667a3724ba04d1b2799b75845215f1944debbdf844f.
//
// Solidity: event Initiated(bytes32 _proposal)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterInitiated(opts *bind.FilterOpts) (*VotiumVLCVXOldInitiatedIterator, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "Initiated")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldInitiatedIterator{contract: _VotiumVLCVXOld.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0x08007a3b331cd9bd3d1d3667a3724ba04d1b2799b75845215f1944debbdf844f.
//
// Solidity: event Initiated(bytes32 _proposal)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldInitiated) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "Initiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldInitiated)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Initiated", log); err != nil {
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

// ParseInitiated is a log parse operation binding the contract event 0x08007a3b331cd9bd3d1d3667a3724ba04d1b2799b75845215f1944debbdf844f.
//
// Solidity: event Initiated(bytes32 _proposal)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseInitiated(log types.Log) (*VotiumVLCVXOldInitiated, error) {
	event := new(VotiumVLCVXOldInitiated)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Initiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldModifiedTeamIterator is returned from FilterModifiedTeam and is used to iterate over the raw logs and unpacked data for ModifiedTeam events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldModifiedTeamIterator struct {
	Event *VotiumVLCVXOldModifiedTeam // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldModifiedTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldModifiedTeam)
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
		it.Event = new(VotiumVLCVXOldModifiedTeam)
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
func (it *VotiumVLCVXOldModifiedTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldModifiedTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldModifiedTeam represents a ModifiedTeam event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldModifiedTeam struct {
	Member   common.Address
	Approval bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifiedTeam is a free log retrieval operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterModifiedTeam(opts *bind.FilterOpts) (*VotiumVLCVXOldModifiedTeamIterator, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldModifiedTeamIterator{contract: _VotiumVLCVXOld.contract, event: "ModifiedTeam", logs: logs, sub: sub}, nil
}

// WatchModifiedTeam is a free log subscription operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchModifiedTeam(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldModifiedTeam) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldModifiedTeam)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
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
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseModifiedTeam(log types.Log) (*VotiumVLCVXOldModifiedTeam, error) {
	event := new(VotiumVLCVXOldModifiedTeam)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldOwnershipTransferredIterator struct {
	Event *VotiumVLCVXOldOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldOwnershipTransferred)
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
		it.Event = new(VotiumVLCVXOldOwnershipTransferred)
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
func (it *VotiumVLCVXOldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldOwnershipTransferred represents a OwnershipTransferred event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotiumVLCVXOldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldOwnershipTransferredIterator{contract: _VotiumVLCVXOld.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldOwnershipTransferred)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseOwnershipTransferred(log types.Log) (*VotiumVLCVXOldOwnershipTransferred, error) {
	event := new(VotiumVLCVXOldOwnershipTransferred)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldUpdatedDistributorIterator is returned from FilterUpdatedDistributor and is used to iterate over the raw logs and unpacked data for UpdatedDistributor events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldUpdatedDistributorIterator struct {
	Event *VotiumVLCVXOldUpdatedDistributor // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldUpdatedDistributorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldUpdatedDistributor)
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
		it.Event = new(VotiumVLCVXOldUpdatedDistributor)
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
func (it *VotiumVLCVXOldUpdatedDistributorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldUpdatedDistributorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldUpdatedDistributor represents a UpdatedDistributor event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldUpdatedDistributor struct {
	Token       common.Address
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDistributor is a free log retrieval operation binding the contract event 0xfbeb7d6e3dcb0b2bee779c64bea292263c335a0ad3f39c459aec2a70b0745653.
//
// Solidity: event UpdatedDistributor(address indexed _token, address _distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterUpdatedDistributor(opts *bind.FilterOpts, _token []common.Address) (*VotiumVLCVXOldUpdatedDistributorIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "UpdatedDistributor", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldUpdatedDistributorIterator{contract: _VotiumVLCVXOld.contract, event: "UpdatedDistributor", logs: logs, sub: sub}, nil
}

// WatchUpdatedDistributor is a free log subscription operation binding the contract event 0xfbeb7d6e3dcb0b2bee779c64bea292263c335a0ad3f39c459aec2a70b0745653.
//
// Solidity: event UpdatedDistributor(address indexed _token, address _distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchUpdatedDistributor(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldUpdatedDistributor, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "UpdatedDistributor", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldUpdatedDistributor)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
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

// ParseUpdatedDistributor is a log parse operation binding the contract event 0xfbeb7d6e3dcb0b2bee779c64bea292263c335a0ad3f39c459aec2a70b0745653.
//
// Solidity: event UpdatedDistributor(address indexed _token, address _distributor)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseUpdatedDistributor(log types.Log) (*VotiumVLCVXOldUpdatedDistributor, error) {
	event := new(VotiumVLCVXOldUpdatedDistributor)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldUpdatedFeeIterator is returned from FilterUpdatedFee and is used to iterate over the raw logs and unpacked data for UpdatedFee events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldUpdatedFeeIterator struct {
	Event *VotiumVLCVXOldUpdatedFee // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldUpdatedFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldUpdatedFee)
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
		it.Event = new(VotiumVLCVXOldUpdatedFee)
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
func (it *VotiumVLCVXOldUpdatedFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldUpdatedFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldUpdatedFee represents a UpdatedFee event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldUpdatedFee struct {
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFee is a free log retrieval operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterUpdatedFee(opts *bind.FilterOpts) (*VotiumVLCVXOldUpdatedFeeIterator, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldUpdatedFeeIterator{contract: _VotiumVLCVXOld.contract, event: "UpdatedFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedFee is a free log subscription operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchUpdatedFee(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldUpdatedFee) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldUpdatedFee)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
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
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseUpdatedFee(log types.Log) (*VotiumVLCVXOldUpdatedFee, error) {
	event := new(VotiumVLCVXOldUpdatedFee)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldWhitelistRequirementIterator is returned from FilterWhitelistRequirement and is used to iterate over the raw logs and unpacked data for WhitelistRequirement events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldWhitelistRequirementIterator struct {
	Event *VotiumVLCVXOldWhitelistRequirement // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldWhitelistRequirementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldWhitelistRequirement)
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
		it.Event = new(VotiumVLCVXOldWhitelistRequirement)
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
func (it *VotiumVLCVXOldWhitelistRequirementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldWhitelistRequirementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldWhitelistRequirement represents a WhitelistRequirement event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldWhitelistRequirement struct {
	RequireWhitelist bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRequirement is a free log retrieval operation binding the contract event 0x1fc0ad5b128ce828ca9ef82c0a5b404f173ff52a5e9a3607c182f38a6bbdb1b7.
//
// Solidity: event WhitelistRequirement(bool _requireWhitelist)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterWhitelistRequirement(opts *bind.FilterOpts) (*VotiumVLCVXOldWhitelistRequirementIterator, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "WhitelistRequirement")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldWhitelistRequirementIterator{contract: _VotiumVLCVXOld.contract, event: "WhitelistRequirement", logs: logs, sub: sub}, nil
}

// WatchWhitelistRequirement is a free log subscription operation binding the contract event 0x1fc0ad5b128ce828ca9ef82c0a5b404f173ff52a5e9a3607c182f38a6bbdb1b7.
//
// Solidity: event WhitelistRequirement(bool _requireWhitelist)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchWhitelistRequirement(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldWhitelistRequirement) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "WhitelistRequirement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldWhitelistRequirement)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "WhitelistRequirement", log); err != nil {
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

// ParseWhitelistRequirement is a log parse operation binding the contract event 0x1fc0ad5b128ce828ca9ef82c0a5b404f173ff52a5e9a3607c182f38a6bbdb1b7.
//
// Solidity: event WhitelistRequirement(bool _requireWhitelist)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseWhitelistRequirement(log types.Log) (*VotiumVLCVXOldWhitelistRequirement, error) {
	event := new(VotiumVLCVXOldWhitelistRequirement)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "WhitelistRequirement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXOldWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldWhitelistedIterator struct {
	Event *VotiumVLCVXOldWhitelisted // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXOldWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXOldWhitelisted)
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
		it.Event = new(VotiumVLCVXOldWhitelisted)
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
func (it *VotiumVLCVXOldWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXOldWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXOldWhitelisted represents a Whitelisted event raised by the VotiumVLCVXOld contract.
type VotiumVLCVXOldWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address _token)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) FilterWhitelisted(opts *bind.FilterOpts) (*VotiumVLCVXOldWhitelistedIterator, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.FilterLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXOldWhitelistedIterator{contract: _VotiumVLCVXOld.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address _token)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXOldWhitelisted) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXOld.contract.WatchLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXOldWhitelisted)
				if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address _token)
func (_VotiumVLCVXOld *VotiumVLCVXOldFilterer) ParseWhitelisted(log types.Log) (*VotiumVLCVXOldWhitelisted, error) {
	event := new(VotiumVLCVXOldWhitelisted)
	if err := _VotiumVLCVXOld.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
