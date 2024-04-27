// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votiumMerkle

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

// MultiMerkleStashclaimParam is an auto generated low-level Go binding around an user-defined struct.
type MultiMerkleStashclaimParam struct {
	Token       common.Address
	Index       *big.Int
	Amount      *big.Int
	MerkleProof [][32]byte
}

// VotiumMerkleMetaData contains all meta data concerning the VotiumMerkle contract.
var VotiumMerkleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"update\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"update\",\"type\":\"uint256\"}],\"name\":\"MerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMultiMerkleStash.claimParam[]\",\"name\":\"claims\",\"type\":\"tuple[]\"}],\"name\":\"claimMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VotiumMerkleABI is the input ABI used to generate the binding from.
// Deprecated: Use VotiumMerkleMetaData.ABI instead.
var VotiumMerkleABI = VotiumMerkleMetaData.ABI

// VotiumMerkle is an auto generated Go binding around an Ethereum contract.
type VotiumMerkle struct {
	VotiumMerkleCaller     // Read-only binding to the contract
	VotiumMerkleTransactor // Write-only binding to the contract
	VotiumMerkleFilterer   // Log filterer for contract events
}

// VotiumMerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotiumMerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumMerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotiumMerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumMerkleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotiumMerkleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumMerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotiumMerkleSession struct {
	Contract     *VotiumMerkle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotiumMerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotiumMerkleCallerSession struct {
	Contract *VotiumMerkleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VotiumMerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotiumMerkleTransactorSession struct {
	Contract     *VotiumMerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VotiumMerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotiumMerkleRaw struct {
	Contract *VotiumMerkle // Generic contract binding to access the raw methods on
}

// VotiumMerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotiumMerkleCallerRaw struct {
	Contract *VotiumMerkleCaller // Generic read-only contract binding to access the raw methods on
}

// VotiumMerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotiumMerkleTransactorRaw struct {
	Contract *VotiumMerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotiumMerkle creates a new instance of VotiumMerkle, bound to a specific deployed contract.
func NewVotiumMerkle(address common.Address, backend bind.ContractBackend) (*VotiumMerkle, error) {
	contract, err := bindVotiumMerkle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkle{VotiumMerkleCaller: VotiumMerkleCaller{contract: contract}, VotiumMerkleTransactor: VotiumMerkleTransactor{contract: contract}, VotiumMerkleFilterer: VotiumMerkleFilterer{contract: contract}}, nil
}

// NewVotiumMerkleCaller creates a new read-only instance of VotiumMerkle, bound to a specific deployed contract.
func NewVotiumMerkleCaller(address common.Address, caller bind.ContractCaller) (*VotiumMerkleCaller, error) {
	contract, err := bindVotiumMerkle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleCaller{contract: contract}, nil
}

// NewVotiumMerkleTransactor creates a new write-only instance of VotiumMerkle, bound to a specific deployed contract.
func NewVotiumMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*VotiumMerkleTransactor, error) {
	contract, err := bindVotiumMerkle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleTransactor{contract: contract}, nil
}

// NewVotiumMerkleFilterer creates a new log filterer instance of VotiumMerkle, bound to a specific deployed contract.
func NewVotiumMerkleFilterer(address common.Address, filterer bind.ContractFilterer) (*VotiumMerkleFilterer, error) {
	contract, err := bindVotiumMerkle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleFilterer{contract: contract}, nil
}

// bindVotiumMerkle binds a generic wrapper to an already deployed contract.
func bindVotiumMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotiumMerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumMerkle *VotiumMerkleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumMerkle.Contract.VotiumMerkleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumMerkle *VotiumMerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.VotiumMerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumMerkle *VotiumMerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.VotiumMerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumMerkle *VotiumMerkleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumMerkle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumMerkle *VotiumMerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumMerkle *VotiumMerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.contract.Transact(opts, method, params...)
}

// IsClaimed is a free data retrieval call binding the contract method 0x562beba8.
//
// Solidity: function isClaimed(address token, uint256 index) view returns(bool)
func (_VotiumMerkle *VotiumMerkleCaller) IsClaimed(opts *bind.CallOpts, token common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _VotiumMerkle.contract.Call(opts, &out, "isClaimed", token, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x562beba8.
//
// Solidity: function isClaimed(address token, uint256 index) view returns(bool)
func (_VotiumMerkle *VotiumMerkleSession) IsClaimed(token common.Address, index *big.Int) (bool, error) {
	return _VotiumMerkle.Contract.IsClaimed(&_VotiumMerkle.CallOpts, token, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x562beba8.
//
// Solidity: function isClaimed(address token, uint256 index) view returns(bool)
func (_VotiumMerkle *VotiumMerkleCallerSession) IsClaimed(token common.Address, index *big.Int) (bool, error) {
	return _VotiumMerkle.Contract.IsClaimed(&_VotiumMerkle.CallOpts, token, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0xe2111586.
//
// Solidity: function merkleRoot(address ) view returns(bytes32)
func (_VotiumMerkle *VotiumMerkleCaller) MerkleRoot(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _VotiumMerkle.contract.Call(opts, &out, "merkleRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0xe2111586.
//
// Solidity: function merkleRoot(address ) view returns(bytes32)
func (_VotiumMerkle *VotiumMerkleSession) MerkleRoot(arg0 common.Address) ([32]byte, error) {
	return _VotiumMerkle.Contract.MerkleRoot(&_VotiumMerkle.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0xe2111586.
//
// Solidity: function merkleRoot(address ) view returns(bytes32)
func (_VotiumMerkle *VotiumMerkleCallerSession) MerkleRoot(arg0 common.Address) ([32]byte, error) {
	return _VotiumMerkle.Contract.MerkleRoot(&_VotiumMerkle.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumMerkle *VotiumMerkleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumMerkle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumMerkle *VotiumMerkleSession) Owner() (common.Address, error) {
	return _VotiumMerkle.Contract.Owner(&_VotiumMerkle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumMerkle *VotiumMerkleCallerSession) Owner() (common.Address, error) {
	return _VotiumMerkle.Contract.Owner(&_VotiumMerkle.CallOpts)
}

// Update is a free data retrieval call binding the contract method 0x1c1b8772.
//
// Solidity: function update(address ) view returns(uint256)
func (_VotiumMerkle *VotiumMerkleCaller) Update(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotiumMerkle.contract.Call(opts, &out, "update", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Update is a free data retrieval call binding the contract method 0x1c1b8772.
//
// Solidity: function update(address ) view returns(uint256)
func (_VotiumMerkle *VotiumMerkleSession) Update(arg0 common.Address) (*big.Int, error) {
	return _VotiumMerkle.Contract.Update(&_VotiumMerkle.CallOpts, arg0)
}

// Update is a free data retrieval call binding the contract method 0x1c1b8772.
//
// Solidity: function update(address ) view returns(uint256)
func (_VotiumMerkle *VotiumMerkleCallerSession) Update(arg0 common.Address) (*big.Int, error) {
	return _VotiumMerkle.Contract.Update(&_VotiumMerkle.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x12d18ed6.
//
// Solidity: function claim(address token, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_VotiumMerkle *VotiumMerkleTransactor) Claim(opts *bind.TransactOpts, token common.Address, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.contract.Transact(opts, "claim", token, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x12d18ed6.
//
// Solidity: function claim(address token, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_VotiumMerkle *VotiumMerkleSession) Claim(token common.Address, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.Claim(&_VotiumMerkle.TransactOpts, token, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x12d18ed6.
//
// Solidity: function claim(address token, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_VotiumMerkle *VotiumMerkleTransactorSession) Claim(token common.Address, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.Claim(&_VotiumMerkle.TransactOpts, token, index, account, amount, merkleProof)
}

// ClaimMulti is a paid mutator transaction binding the contract method 0x20ce64de.
//
// Solidity: function claimMulti(address account, (address,uint256,uint256,bytes32[])[] claims) returns()
func (_VotiumMerkle *VotiumMerkleTransactor) ClaimMulti(opts *bind.TransactOpts, account common.Address, claims []MultiMerkleStashclaimParam) (*types.Transaction, error) {
	return _VotiumMerkle.contract.Transact(opts, "claimMulti", account, claims)
}

// ClaimMulti is a paid mutator transaction binding the contract method 0x20ce64de.
//
// Solidity: function claimMulti(address account, (address,uint256,uint256,bytes32[])[] claims) returns()
func (_VotiumMerkle *VotiumMerkleSession) ClaimMulti(account common.Address, claims []MultiMerkleStashclaimParam) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.ClaimMulti(&_VotiumMerkle.TransactOpts, account, claims)
}

// ClaimMulti is a paid mutator transaction binding the contract method 0x20ce64de.
//
// Solidity: function claimMulti(address account, (address,uint256,uint256,bytes32[])[] claims) returns()
func (_VotiumMerkle *VotiumMerkleTransactorSession) ClaimMulti(account common.Address, claims []MultiMerkleStashclaimParam) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.ClaimMulti(&_VotiumMerkle.TransactOpts, account, claims)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumMerkle *VotiumMerkleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotiumMerkle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumMerkle *VotiumMerkleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.TransferOwnership(&_VotiumMerkle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumMerkle *VotiumMerkleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.TransferOwnership(&_VotiumMerkle.TransactOpts, newOwner)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0xe786818e.
//
// Solidity: function updateMerkleRoot(address token, bytes32 _merkleRoot) returns()
func (_VotiumMerkle *VotiumMerkleTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, token common.Address, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.contract.Transact(opts, "updateMerkleRoot", token, _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0xe786818e.
//
// Solidity: function updateMerkleRoot(address token, bytes32 _merkleRoot) returns()
func (_VotiumMerkle *VotiumMerkleSession) UpdateMerkleRoot(token common.Address, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.UpdateMerkleRoot(&_VotiumMerkle.TransactOpts, token, _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0xe786818e.
//
// Solidity: function updateMerkleRoot(address token, bytes32 _merkleRoot) returns()
func (_VotiumMerkle *VotiumMerkleTransactorSession) UpdateMerkleRoot(token common.Address, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _VotiumMerkle.Contract.UpdateMerkleRoot(&_VotiumMerkle.TransactOpts, token, _merkleRoot)
}

// VotiumMerkleClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the VotiumMerkle contract.
type VotiumMerkleClaimedIterator struct {
	Event *VotiumMerkleClaimed // Event containing the contract specifics and raw log

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
func (it *VotiumMerkleClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumMerkleClaimed)
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
		it.Event = new(VotiumMerkleClaimed)
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
func (it *VotiumMerkleClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumMerkleClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumMerkleClaimed represents a Claimed event raised by the VotiumMerkle contract.
type VotiumMerkleClaimed struct {
	Token   common.Address
	Index   *big.Int
	Amount  *big.Int
	Account common.Address
	Update  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x4766921f5c59646d22d7d266a29164c8e9623684d8dfdbd931731dfdca025238.
//
// Solidity: event Claimed(address indexed token, uint256 index, uint256 amount, address indexed account, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) FilterClaimed(opts *bind.FilterOpts, token []common.Address, account []common.Address, update []*big.Int) (*VotiumMerkleClaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var updateRule []interface{}
	for _, updateItem := range update {
		updateRule = append(updateRule, updateItem)
	}

	logs, sub, err := _VotiumMerkle.contract.FilterLogs(opts, "Claimed", tokenRule, accountRule, updateRule)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleClaimedIterator{contract: _VotiumMerkle.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x4766921f5c59646d22d7d266a29164c8e9623684d8dfdbd931731dfdca025238.
//
// Solidity: event Claimed(address indexed token, uint256 index, uint256 amount, address indexed account, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *VotiumMerkleClaimed, token []common.Address, account []common.Address, update []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var updateRule []interface{}
	for _, updateItem := range update {
		updateRule = append(updateRule, updateItem)
	}

	logs, sub, err := _VotiumMerkle.contract.WatchLogs(opts, "Claimed", tokenRule, accountRule, updateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumMerkleClaimed)
				if err := _VotiumMerkle.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x4766921f5c59646d22d7d266a29164c8e9623684d8dfdbd931731dfdca025238.
//
// Solidity: event Claimed(address indexed token, uint256 index, uint256 amount, address indexed account, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) ParseClaimed(log types.Log) (*VotiumMerkleClaimed, error) {
	event := new(VotiumMerkleClaimed)
	if err := _VotiumMerkle.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumMerkleMerkleRootUpdatedIterator is returned from FilterMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for MerkleRootUpdated events raised by the VotiumMerkle contract.
type VotiumMerkleMerkleRootUpdatedIterator struct {
	Event *VotiumMerkleMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *VotiumMerkleMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumMerkleMerkleRootUpdated)
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
		it.Event = new(VotiumMerkleMerkleRootUpdated)
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
func (it *VotiumMerkleMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumMerkleMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumMerkleMerkleRootUpdated represents a MerkleRootUpdated event raised by the VotiumMerkle contract.
type VotiumMerkleMerkleRootUpdated struct {
	Token      common.Address
	MerkleRoot [32]byte
	Update     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdated is a free log retrieval operation binding the contract event 0xbdc375de8a677af383c6ee6f8b2027dbd231c3c47453ce81d1ce8f1bcdb722dc.
//
// Solidity: event MerkleRootUpdated(address indexed token, bytes32 indexed merkleRoot, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) FilterMerkleRootUpdated(opts *bind.FilterOpts, token []common.Address, merkleRoot [][32]byte, update []*big.Int) (*VotiumMerkleMerkleRootUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var merkleRootRule []interface{}
	for _, merkleRootItem := range merkleRoot {
		merkleRootRule = append(merkleRootRule, merkleRootItem)
	}
	var updateRule []interface{}
	for _, updateItem := range update {
		updateRule = append(updateRule, updateItem)
	}

	logs, sub, err := _VotiumMerkle.contract.FilterLogs(opts, "MerkleRootUpdated", tokenRule, merkleRootRule, updateRule)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleMerkleRootUpdatedIterator{contract: _VotiumMerkle.contract, event: "MerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdated is a free log subscription operation binding the contract event 0xbdc375de8a677af383c6ee6f8b2027dbd231c3c47453ce81d1ce8f1bcdb722dc.
//
// Solidity: event MerkleRootUpdated(address indexed token, bytes32 indexed merkleRoot, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) WatchMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *VotiumMerkleMerkleRootUpdated, token []common.Address, merkleRoot [][32]byte, update []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var merkleRootRule []interface{}
	for _, merkleRootItem := range merkleRoot {
		merkleRootRule = append(merkleRootRule, merkleRootItem)
	}
	var updateRule []interface{}
	for _, updateItem := range update {
		updateRule = append(updateRule, updateItem)
	}

	logs, sub, err := _VotiumMerkle.contract.WatchLogs(opts, "MerkleRootUpdated", tokenRule, merkleRootRule, updateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumMerkleMerkleRootUpdated)
				if err := _VotiumMerkle.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
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

// ParseMerkleRootUpdated is a log parse operation binding the contract event 0xbdc375de8a677af383c6ee6f8b2027dbd231c3c47453ce81d1ce8f1bcdb722dc.
//
// Solidity: event MerkleRootUpdated(address indexed token, bytes32 indexed merkleRoot, uint256 indexed update)
func (_VotiumMerkle *VotiumMerkleFilterer) ParseMerkleRootUpdated(log types.Log) (*VotiumMerkleMerkleRootUpdated, error) {
	event := new(VotiumMerkleMerkleRootUpdated)
	if err := _VotiumMerkle.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumMerkleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotiumMerkle contract.
type VotiumMerkleOwnershipTransferredIterator struct {
	Event *VotiumMerkleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotiumMerkleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumMerkleOwnershipTransferred)
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
		it.Event = new(VotiumMerkleOwnershipTransferred)
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
func (it *VotiumMerkleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumMerkleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumMerkleOwnershipTransferred represents a OwnershipTransferred event raised by the VotiumMerkle contract.
type VotiumMerkleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumMerkle *VotiumMerkleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotiumMerkleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumMerkle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotiumMerkleOwnershipTransferredIterator{contract: _VotiumMerkle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumMerkle *VotiumMerkleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotiumMerkleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumMerkle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumMerkleOwnershipTransferred)
				if err := _VotiumMerkle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VotiumMerkle *VotiumMerkleFilterer) ParseOwnershipTransferred(log types.Log) (*VotiumMerkleOwnershipTransferred, error) {
	event := new(VotiumMerkleOwnershipTransferred)
	if err := _VotiumMerkle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
