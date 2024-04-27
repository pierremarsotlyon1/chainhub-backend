// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package questDistributor

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

// MultiMerkleDistributorClaimParams is an auto generated low-level Go binding around an user-defined struct.
type MultiMerkleDistributorClaimParams struct {
	QuestID     *big.Int
	Period      *big.Int
	Index       *big.Int
	Amount      *big.Int
	MerkleProof [][32]byte
}

// QuestDistributorMetaData contains all meta data concerning the QuestDistributor contract.
var QuestDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_questBoard\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBeOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRecoverToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyMerkleRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectQuestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRewardAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleRootNotUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodNotClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodNotListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestAlreadyListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestNotListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"name\":\"NewQuest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"QuestPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addQuest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"}],\"name\":\"addQuestPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMultiMerkleDistributor.ClaimParams[]\",\"name\":\"claims\",\"type\":\"tuple[]\"}],\"name\":\"claimQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"emergencyUpdateQuestPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newTotalRewardAmount\",\"type\":\"uint256\"}],\"name\":\"fixQuestPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"}],\"name\":\"getClosedPeriodsByQuests\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMultiMerkleDistributor.ClaimParams[]\",\"name\":\"claims\",\"type\":\"tuple[]\"}],\"name\":\"multiClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"questBoard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questClosedPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questMerkleRootPerPeriod\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questRewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questRewardsPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"updateQuestPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestDistributorMetaData.ABI instead.
var QuestDistributorABI = QuestDistributorMetaData.ABI

// QuestDistributor is an auto generated Go binding around an Ethereum contract.
type QuestDistributor struct {
	QuestDistributorCaller     // Read-only binding to the contract
	QuestDistributorTransactor // Write-only binding to the contract
	QuestDistributorFilterer   // Log filterer for contract events
}

// QuestDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuestDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestDistributorSession struct {
	Contract     *QuestDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestDistributorCallerSession struct {
	Contract *QuestDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// QuestDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestDistributorTransactorSession struct {
	Contract     *QuestDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// QuestDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuestDistributorRaw struct {
	Contract *QuestDistributor // Generic contract binding to access the raw methods on
}

// QuestDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestDistributorCallerRaw struct {
	Contract *QuestDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// QuestDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestDistributorTransactorRaw struct {
	Contract *QuestDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuestDistributor creates a new instance of QuestDistributor, bound to a specific deployed contract.
func NewQuestDistributor(address common.Address, backend bind.ContractBackend) (*QuestDistributor, error) {
	contract, err := bindQuestDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuestDistributor{QuestDistributorCaller: QuestDistributorCaller{contract: contract}, QuestDistributorTransactor: QuestDistributorTransactor{contract: contract}, QuestDistributorFilterer: QuestDistributorFilterer{contract: contract}}, nil
}

// NewQuestDistributorCaller creates a new read-only instance of QuestDistributor, bound to a specific deployed contract.
func NewQuestDistributorCaller(address common.Address, caller bind.ContractCaller) (*QuestDistributorCaller, error) {
	contract, err := bindQuestDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorCaller{contract: contract}, nil
}

// NewQuestDistributorTransactor creates a new write-only instance of QuestDistributor, bound to a specific deployed contract.
func NewQuestDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*QuestDistributorTransactor, error) {
	contract, err := bindQuestDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorTransactor{contract: contract}, nil
}

// NewQuestDistributorFilterer creates a new log filterer instance of QuestDistributor, bound to a specific deployed contract.
func NewQuestDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*QuestDistributorFilterer, error) {
	contract, err := bindQuestDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorFilterer{contract: contract}, nil
}

// bindQuestDistributor binds a generic wrapper to an already deployed contract.
func bindQuestDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuestDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestDistributor *QuestDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestDistributor.Contract.QuestDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestDistributor *QuestDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestDistributor.Contract.QuestDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestDistributor *QuestDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestDistributor.Contract.QuestDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestDistributor *QuestDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestDistributor *QuestDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestDistributor *QuestDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestDistributor.Contract.contract.Transact(opts, method, params...)
}

// GetClosedPeriodsByQuests is a free data retrieval call binding the contract method 0x46fe84e9.
//
// Solidity: function getClosedPeriodsByQuests(uint256 questID) view returns(uint256[])
func (_QuestDistributor *QuestDistributorCaller) GetClosedPeriodsByQuests(opts *bind.CallOpts, questID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "getClosedPeriodsByQuests", questID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetClosedPeriodsByQuests is a free data retrieval call binding the contract method 0x46fe84e9.
//
// Solidity: function getClosedPeriodsByQuests(uint256 questID) view returns(uint256[])
func (_QuestDistributor *QuestDistributorSession) GetClosedPeriodsByQuests(questID *big.Int) ([]*big.Int, error) {
	return _QuestDistributor.Contract.GetClosedPeriodsByQuests(&_QuestDistributor.CallOpts, questID)
}

// GetClosedPeriodsByQuests is a free data retrieval call binding the contract method 0x46fe84e9.
//
// Solidity: function getClosedPeriodsByQuests(uint256 questID) view returns(uint256[])
func (_QuestDistributor *QuestDistributorCallerSession) GetClosedPeriodsByQuests(questID *big.Int) ([]*big.Int, error) {
	return _QuestDistributor.Contract.GetClosedPeriodsByQuests(&_QuestDistributor.CallOpts, questID)
}

// IsClaimed is a free data retrieval call binding the contract method 0x627614ac.
//
// Solidity: function isClaimed(uint256 questID, uint256 period, uint256 index) view returns(bool)
func (_QuestDistributor *QuestDistributorCaller) IsClaimed(opts *bind.CallOpts, questID *big.Int, period *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "isClaimed", questID, period, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x627614ac.
//
// Solidity: function isClaimed(uint256 questID, uint256 period, uint256 index) view returns(bool)
func (_QuestDistributor *QuestDistributorSession) IsClaimed(questID *big.Int, period *big.Int, index *big.Int) (bool, error) {
	return _QuestDistributor.Contract.IsClaimed(&_QuestDistributor.CallOpts, questID, period, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x627614ac.
//
// Solidity: function isClaimed(uint256 questID, uint256 period, uint256 index) view returns(bool)
func (_QuestDistributor *QuestDistributorCallerSession) IsClaimed(questID *big.Int, period *big.Int, index *big.Int) (bool, error) {
	return _QuestDistributor.Contract.IsClaimed(&_QuestDistributor.CallOpts, questID, period, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestDistributor *QuestDistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestDistributor *QuestDistributorSession) Owner() (common.Address, error) {
	return _QuestDistributor.Contract.Owner(&_QuestDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestDistributor *QuestDistributorCallerSession) Owner() (common.Address, error) {
	return _QuestDistributor.Contract.Owner(&_QuestDistributor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_QuestDistributor *QuestDistributorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_QuestDistributor *QuestDistributorSession) PendingOwner() (common.Address, error) {
	return _QuestDistributor.Contract.PendingOwner(&_QuestDistributor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_QuestDistributor *QuestDistributorCallerSession) PendingOwner() (common.Address, error) {
	return _QuestDistributor.Contract.PendingOwner(&_QuestDistributor.CallOpts)
}

// QuestBoard is a free data retrieval call binding the contract method 0x13122798.
//
// Solidity: function questBoard() view returns(address)
func (_QuestDistributor *QuestDistributorCaller) QuestBoard(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "questBoard")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuestBoard is a free data retrieval call binding the contract method 0x13122798.
//
// Solidity: function questBoard() view returns(address)
func (_QuestDistributor *QuestDistributorSession) QuestBoard() (common.Address, error) {
	return _QuestDistributor.Contract.QuestBoard(&_QuestDistributor.CallOpts)
}

// QuestBoard is a free data retrieval call binding the contract method 0x13122798.
//
// Solidity: function questBoard() view returns(address)
func (_QuestDistributor *QuestDistributorCallerSession) QuestBoard() (common.Address, error) {
	return _QuestDistributor.Contract.QuestBoard(&_QuestDistributor.CallOpts)
}

// QuestClosedPeriods is a free data retrieval call binding the contract method 0x6ec803bb.
//
// Solidity: function questClosedPeriods(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorCaller) QuestClosedPeriods(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "questClosedPeriods", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuestClosedPeriods is a free data retrieval call binding the contract method 0x6ec803bb.
//
// Solidity: function questClosedPeriods(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorSession) QuestClosedPeriods(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _QuestDistributor.Contract.QuestClosedPeriods(&_QuestDistributor.CallOpts, arg0, arg1)
}

// QuestClosedPeriods is a free data retrieval call binding the contract method 0x6ec803bb.
//
// Solidity: function questClosedPeriods(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorCallerSession) QuestClosedPeriods(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _QuestDistributor.Contract.QuestClosedPeriods(&_QuestDistributor.CallOpts, arg0, arg1)
}

// QuestMerkleRootPerPeriod is a free data retrieval call binding the contract method 0x93380358.
//
// Solidity: function questMerkleRootPerPeriod(uint256 , uint256 ) view returns(bytes32)
func (_QuestDistributor *QuestDistributorCaller) QuestMerkleRootPerPeriod(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "questMerkleRootPerPeriod", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QuestMerkleRootPerPeriod is a free data retrieval call binding the contract method 0x93380358.
//
// Solidity: function questMerkleRootPerPeriod(uint256 , uint256 ) view returns(bytes32)
func (_QuestDistributor *QuestDistributorSession) QuestMerkleRootPerPeriod(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _QuestDistributor.Contract.QuestMerkleRootPerPeriod(&_QuestDistributor.CallOpts, arg0, arg1)
}

// QuestMerkleRootPerPeriod is a free data retrieval call binding the contract method 0x93380358.
//
// Solidity: function questMerkleRootPerPeriod(uint256 , uint256 ) view returns(bytes32)
func (_QuestDistributor *QuestDistributorCallerSession) QuestMerkleRootPerPeriod(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _QuestDistributor.Contract.QuestMerkleRootPerPeriod(&_QuestDistributor.CallOpts, arg0, arg1)
}

// QuestRewardToken is a free data retrieval call binding the contract method 0x95861d5c.
//
// Solidity: function questRewardToken(uint256 ) view returns(address)
func (_QuestDistributor *QuestDistributorCaller) QuestRewardToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "questRewardToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuestRewardToken is a free data retrieval call binding the contract method 0x95861d5c.
//
// Solidity: function questRewardToken(uint256 ) view returns(address)
func (_QuestDistributor *QuestDistributorSession) QuestRewardToken(arg0 *big.Int) (common.Address, error) {
	return _QuestDistributor.Contract.QuestRewardToken(&_QuestDistributor.CallOpts, arg0)
}

// QuestRewardToken is a free data retrieval call binding the contract method 0x95861d5c.
//
// Solidity: function questRewardToken(uint256 ) view returns(address)
func (_QuestDistributor *QuestDistributorCallerSession) QuestRewardToken(arg0 *big.Int) (common.Address, error) {
	return _QuestDistributor.Contract.QuestRewardToken(&_QuestDistributor.CallOpts, arg0)
}

// QuestRewardsPerPeriod is a free data retrieval call binding the contract method 0x493b952c.
//
// Solidity: function questRewardsPerPeriod(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorCaller) QuestRewardsPerPeriod(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "questRewardsPerPeriod", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuestRewardsPerPeriod is a free data retrieval call binding the contract method 0x493b952c.
//
// Solidity: function questRewardsPerPeriod(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorSession) QuestRewardsPerPeriod(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _QuestDistributor.Contract.QuestRewardsPerPeriod(&_QuestDistributor.CallOpts, arg0, arg1)
}

// QuestRewardsPerPeriod is a free data retrieval call binding the contract method 0x493b952c.
//
// Solidity: function questRewardsPerPeriod(uint256 , uint256 ) view returns(uint256)
func (_QuestDistributor *QuestDistributorCallerSession) QuestRewardsPerPeriod(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _QuestDistributor.Contract.QuestRewardsPerPeriod(&_QuestDistributor.CallOpts, arg0, arg1)
}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(bool)
func (_QuestDistributor *QuestDistributorCaller) RewardTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _QuestDistributor.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(bool)
func (_QuestDistributor *QuestDistributorSession) RewardTokens(arg0 common.Address) (bool, error) {
	return _QuestDistributor.Contract.RewardTokens(&_QuestDistributor.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(bool)
func (_QuestDistributor *QuestDistributorCallerSession) RewardTokens(arg0 common.Address) (bool, error) {
	return _QuestDistributor.Contract.RewardTokens(&_QuestDistributor.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_QuestDistributor *QuestDistributorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_QuestDistributor *QuestDistributorSession) AcceptOwnership() (*types.Transaction, error) {
	return _QuestDistributor.Contract.AcceptOwnership(&_QuestDistributor.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_QuestDistributor *QuestDistributorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _QuestDistributor.Contract.AcceptOwnership(&_QuestDistributor.TransactOpts)
}

// AddQuest is a paid mutator transaction binding the contract method 0x58ad4064.
//
// Solidity: function addQuest(uint256 questID, address token) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) AddQuest(opts *bind.TransactOpts, questID *big.Int, token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "addQuest", questID, token)
}

// AddQuest is a paid mutator transaction binding the contract method 0x58ad4064.
//
// Solidity: function addQuest(uint256 questID, address token) returns(bool)
func (_QuestDistributor *QuestDistributorSession) AddQuest(questID *big.Int, token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.AddQuest(&_QuestDistributor.TransactOpts, questID, token)
}

// AddQuest is a paid mutator transaction binding the contract method 0x58ad4064.
//
// Solidity: function addQuest(uint256 questID, address token) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) AddQuest(questID *big.Int, token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.AddQuest(&_QuestDistributor.TransactOpts, questID, token)
}

// AddQuestPeriod is a paid mutator transaction binding the contract method 0xc25a4d9e.
//
// Solidity: function addQuestPeriod(uint256 questID, uint256 period, uint256 totalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) AddQuestPeriod(opts *bind.TransactOpts, questID *big.Int, period *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "addQuestPeriod", questID, period, totalRewardAmount)
}

// AddQuestPeriod is a paid mutator transaction binding the contract method 0xc25a4d9e.
//
// Solidity: function addQuestPeriod(uint256 questID, uint256 period, uint256 totalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorSession) AddQuestPeriod(questID *big.Int, period *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.Contract.AddQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, totalRewardAmount)
}

// AddQuestPeriod is a paid mutator transaction binding the contract method 0xc25a4d9e.
//
// Solidity: function addQuestPeriod(uint256 questID, uint256 period, uint256 totalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) AddQuestPeriod(questID *big.Int, period *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.Contract.AddQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, totalRewardAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x0c3a0fff.
//
// Solidity: function claim(uint256 questID, uint256 period, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_QuestDistributor *QuestDistributorTransactor) Claim(opts *bind.TransactOpts, questID *big.Int, period *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "claim", questID, period, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x0c3a0fff.
//
// Solidity: function claim(uint256 questID, uint256 period, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_QuestDistributor *QuestDistributorSession) Claim(questID *big.Int, period *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.Claim(&_QuestDistributor.TransactOpts, questID, period, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x0c3a0fff.
//
// Solidity: function claim(uint256 questID, uint256 period, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_QuestDistributor *QuestDistributorTransactorSession) Claim(questID *big.Int, period *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.Claim(&_QuestDistributor.TransactOpts, questID, period, index, account, amount, merkleProof)
}

// ClaimQuest is a paid mutator transaction binding the contract method 0x3a3e6a77.
//
// Solidity: function claimQuest(address account, uint256 questID, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorTransactor) ClaimQuest(opts *bind.TransactOpts, account common.Address, questID *big.Int, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "claimQuest", account, questID, claims)
}

// ClaimQuest is a paid mutator transaction binding the contract method 0x3a3e6a77.
//
// Solidity: function claimQuest(address account, uint256 questID, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorSession) ClaimQuest(account common.Address, questID *big.Int, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.Contract.ClaimQuest(&_QuestDistributor.TransactOpts, account, questID, claims)
}

// ClaimQuest is a paid mutator transaction binding the contract method 0x3a3e6a77.
//
// Solidity: function claimQuest(address account, uint256 questID, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorTransactorSession) ClaimQuest(account common.Address, questID *big.Int, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.Contract.ClaimQuest(&_QuestDistributor.TransactOpts, account, questID, claims)
}

// EmergencyUpdateQuestPeriod is a paid mutator transaction binding the contract method 0x17fea75c.
//
// Solidity: function emergencyUpdateQuestPeriod(uint256 questID, uint256 period, uint256 addedRewardAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) EmergencyUpdateQuestPeriod(opts *bind.TransactOpts, questID *big.Int, period *big.Int, addedRewardAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "emergencyUpdateQuestPeriod", questID, period, addedRewardAmount, merkleRoot)
}

// EmergencyUpdateQuestPeriod is a paid mutator transaction binding the contract method 0x17fea75c.
//
// Solidity: function emergencyUpdateQuestPeriod(uint256 questID, uint256 period, uint256 addedRewardAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorSession) EmergencyUpdateQuestPeriod(questID *big.Int, period *big.Int, addedRewardAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.EmergencyUpdateQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, addedRewardAmount, merkleRoot)
}

// EmergencyUpdateQuestPeriod is a paid mutator transaction binding the contract method 0x17fea75c.
//
// Solidity: function emergencyUpdateQuestPeriod(uint256 questID, uint256 period, uint256 addedRewardAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) EmergencyUpdateQuestPeriod(questID *big.Int, period *big.Int, addedRewardAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.EmergencyUpdateQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, addedRewardAmount, merkleRoot)
}

// FixQuestPeriod is a paid mutator transaction binding the contract method 0x657c5e8c.
//
// Solidity: function fixQuestPeriod(uint256 questID, uint256 period, uint256 newTotalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) FixQuestPeriod(opts *bind.TransactOpts, questID *big.Int, period *big.Int, newTotalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "fixQuestPeriod", questID, period, newTotalRewardAmount)
}

// FixQuestPeriod is a paid mutator transaction binding the contract method 0x657c5e8c.
//
// Solidity: function fixQuestPeriod(uint256 questID, uint256 period, uint256 newTotalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorSession) FixQuestPeriod(questID *big.Int, period *big.Int, newTotalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.Contract.FixQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, newTotalRewardAmount)
}

// FixQuestPeriod is a paid mutator transaction binding the contract method 0x657c5e8c.
//
// Solidity: function fixQuestPeriod(uint256 questID, uint256 period, uint256 newTotalRewardAmount) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) FixQuestPeriod(questID *big.Int, period *big.Int, newTotalRewardAmount *big.Int) (*types.Transaction, error) {
	return _QuestDistributor.Contract.FixQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, newTotalRewardAmount)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x6a9b8ce5.
//
// Solidity: function multiClaim(address account, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorTransactor) MultiClaim(opts *bind.TransactOpts, account common.Address, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "multiClaim", account, claims)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x6a9b8ce5.
//
// Solidity: function multiClaim(address account, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorSession) MultiClaim(account common.Address, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.Contract.MultiClaim(&_QuestDistributor.TransactOpts, account, claims)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x6a9b8ce5.
//
// Solidity: function multiClaim(address account, (uint256,uint256,uint256,uint256,bytes32[])[] claims) returns()
func (_QuestDistributor *QuestDistributorTransactorSession) MultiClaim(account common.Address, claims []MultiMerkleDistributorClaimParams) (*types.Transaction, error) {
	return _QuestDistributor.Contract.MultiClaim(&_QuestDistributor.TransactOpts, account, claims)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x9e8c708e.
//
// Solidity: function recoverERC20(address token) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "recoverERC20", token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x9e8c708e.
//
// Solidity: function recoverERC20(address token) returns(bool)
func (_QuestDistributor *QuestDistributorSession) RecoverERC20(token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.RecoverERC20(&_QuestDistributor.TransactOpts, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x9e8c708e.
//
// Solidity: function recoverERC20(address token) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) RecoverERC20(token common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.RecoverERC20(&_QuestDistributor.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestDistributor *QuestDistributorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestDistributor *QuestDistributorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestDistributor.Contract.RenounceOwnership(&_QuestDistributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestDistributor *QuestDistributorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestDistributor.Contract.RenounceOwnership(&_QuestDistributor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestDistributor *QuestDistributorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestDistributor *QuestDistributorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.TransferOwnership(&_QuestDistributor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestDistributor *QuestDistributorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestDistributor.Contract.TransferOwnership(&_QuestDistributor.TransactOpts, newOwner)
}

// UpdateQuestPeriod is a paid mutator transaction binding the contract method 0x75d0d380.
//
// Solidity: function updateQuestPeriod(uint256 questID, uint256 period, uint256 totalAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorTransactor) UpdateQuestPeriod(opts *bind.TransactOpts, questID *big.Int, period *big.Int, totalAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.contract.Transact(opts, "updateQuestPeriod", questID, period, totalAmount, merkleRoot)
}

// UpdateQuestPeriod is a paid mutator transaction binding the contract method 0x75d0d380.
//
// Solidity: function updateQuestPeriod(uint256 questID, uint256 period, uint256 totalAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorSession) UpdateQuestPeriod(questID *big.Int, period *big.Int, totalAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.UpdateQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, totalAmount, merkleRoot)
}

// UpdateQuestPeriod is a paid mutator transaction binding the contract method 0x75d0d380.
//
// Solidity: function updateQuestPeriod(uint256 questID, uint256 period, uint256 totalAmount, bytes32 merkleRoot) returns(bool)
func (_QuestDistributor *QuestDistributorTransactorSession) UpdateQuestPeriod(questID *big.Int, period *big.Int, totalAmount *big.Int, merkleRoot [32]byte) (*types.Transaction, error) {
	return _QuestDistributor.Contract.UpdateQuestPeriod(&_QuestDistributor.TransactOpts, questID, period, totalAmount, merkleRoot)
}

// QuestDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the QuestDistributor contract.
type QuestDistributorClaimedIterator struct {
	Event *QuestDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *QuestDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestDistributorClaimed)
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
		it.Event = new(QuestDistributorClaimed)
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
func (it *QuestDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestDistributorClaimed represents a Claimed event raised by the QuestDistributor contract.
type QuestDistributorClaimed struct {
	QuestID     *big.Int
	Period      *big.Int
	Index       *big.Int
	Amount      *big.Int
	RewardToken common.Address
	Account     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9a5376f7dcf8631c2b6249c9bec3d715cb97bdd4c82d92e55d147f6b4eea4197.
//
// Solidity: event Claimed(uint256 indexed questID, uint256 indexed period, uint256 index, uint256 amount, address rewardToken, address indexed account)
func (_QuestDistributor *QuestDistributorFilterer) FilterClaimed(opts *bind.FilterOpts, questID []*big.Int, period []*big.Int, account []common.Address) (*QuestDistributorClaimedIterator, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _QuestDistributor.contract.FilterLogs(opts, "Claimed", questIDRule, periodRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorClaimedIterator{contract: _QuestDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9a5376f7dcf8631c2b6249c9bec3d715cb97bdd4c82d92e55d147f6b4eea4197.
//
// Solidity: event Claimed(uint256 indexed questID, uint256 indexed period, uint256 index, uint256 amount, address rewardToken, address indexed account)
func (_QuestDistributor *QuestDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *QuestDistributorClaimed, questID []*big.Int, period []*big.Int, account []common.Address) (event.Subscription, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _QuestDistributor.contract.WatchLogs(opts, "Claimed", questIDRule, periodRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestDistributorClaimed)
				if err := _QuestDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9a5376f7dcf8631c2b6249c9bec3d715cb97bdd4c82d92e55d147f6b4eea4197.
//
// Solidity: event Claimed(uint256 indexed questID, uint256 indexed period, uint256 index, uint256 amount, address rewardToken, address indexed account)
func (_QuestDistributor *QuestDistributorFilterer) ParseClaimed(log types.Log) (*QuestDistributorClaimed, error) {
	event := new(QuestDistributorClaimed)
	if err := _QuestDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestDistributorNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the QuestDistributor contract.
type QuestDistributorNewPendingOwnerIterator struct {
	Event *QuestDistributorNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *QuestDistributorNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestDistributorNewPendingOwner)
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
		it.Event = new(QuestDistributorNewPendingOwner)
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
func (it *QuestDistributorNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestDistributorNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestDistributorNewPendingOwner represents a NewPendingOwner event raised by the QuestDistributor contract.
type QuestDistributorNewPendingOwner struct {
	PreviousPendingOwner common.Address
	NewPendingOwner      common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed previousPendingOwner, address indexed newPendingOwner)
func (_QuestDistributor *QuestDistributorFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, previousPendingOwner []common.Address, newPendingOwner []common.Address) (*QuestDistributorNewPendingOwnerIterator, error) {

	var previousPendingOwnerRule []interface{}
	for _, previousPendingOwnerItem := range previousPendingOwner {
		previousPendingOwnerRule = append(previousPendingOwnerRule, previousPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _QuestDistributor.contract.FilterLogs(opts, "NewPendingOwner", previousPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorNewPendingOwnerIterator{contract: _QuestDistributor.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed previousPendingOwner, address indexed newPendingOwner)
func (_QuestDistributor *QuestDistributorFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *QuestDistributorNewPendingOwner, previousPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var previousPendingOwnerRule []interface{}
	for _, previousPendingOwnerItem := range previousPendingOwner {
		previousPendingOwnerRule = append(previousPendingOwnerRule, previousPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _QuestDistributor.contract.WatchLogs(opts, "NewPendingOwner", previousPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestDistributorNewPendingOwner)
				if err := _QuestDistributor.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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

// ParseNewPendingOwner is a log parse operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed previousPendingOwner, address indexed newPendingOwner)
func (_QuestDistributor *QuestDistributorFilterer) ParseNewPendingOwner(log types.Log) (*QuestDistributorNewPendingOwner, error) {
	event := new(QuestDistributorNewPendingOwner)
	if err := _QuestDistributor.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestDistributorNewQuestIterator is returned from FilterNewQuest and is used to iterate over the raw logs and unpacked data for NewQuest events raised by the QuestDistributor contract.
type QuestDistributorNewQuestIterator struct {
	Event *QuestDistributorNewQuest // Event containing the contract specifics and raw log

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
func (it *QuestDistributorNewQuestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestDistributorNewQuest)
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
		it.Event = new(QuestDistributorNewQuest)
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
func (it *QuestDistributorNewQuestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestDistributorNewQuestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestDistributorNewQuest represents a NewQuest event raised by the QuestDistributor contract.
type QuestDistributorNewQuest struct {
	QuestID     *big.Int
	RewardToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewQuest is a free log retrieval operation binding the contract event 0x6746344851206b5fc2dc76bcb3512131fd07c324969cd3ee2278192805477961.
//
// Solidity: event NewQuest(uint256 indexed questID, address rewardToken)
func (_QuestDistributor *QuestDistributorFilterer) FilterNewQuest(opts *bind.FilterOpts, questID []*big.Int) (*QuestDistributorNewQuestIterator, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}

	logs, sub, err := _QuestDistributor.contract.FilterLogs(opts, "NewQuest", questIDRule)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorNewQuestIterator{contract: _QuestDistributor.contract, event: "NewQuest", logs: logs, sub: sub}, nil
}

// WatchNewQuest is a free log subscription operation binding the contract event 0x6746344851206b5fc2dc76bcb3512131fd07c324969cd3ee2278192805477961.
//
// Solidity: event NewQuest(uint256 indexed questID, address rewardToken)
func (_QuestDistributor *QuestDistributorFilterer) WatchNewQuest(opts *bind.WatchOpts, sink chan<- *QuestDistributorNewQuest, questID []*big.Int) (event.Subscription, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}

	logs, sub, err := _QuestDistributor.contract.WatchLogs(opts, "NewQuest", questIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestDistributorNewQuest)
				if err := _QuestDistributor.contract.UnpackLog(event, "NewQuest", log); err != nil {
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

// ParseNewQuest is a log parse operation binding the contract event 0x6746344851206b5fc2dc76bcb3512131fd07c324969cd3ee2278192805477961.
//
// Solidity: event NewQuest(uint256 indexed questID, address rewardToken)
func (_QuestDistributor *QuestDistributorFilterer) ParseNewQuest(log types.Log) (*QuestDistributorNewQuest, error) {
	event := new(QuestDistributorNewQuest)
	if err := _QuestDistributor.contract.UnpackLog(event, "NewQuest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestDistributorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuestDistributor contract.
type QuestDistributorOwnershipTransferredIterator struct {
	Event *QuestDistributorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuestDistributorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestDistributorOwnershipTransferred)
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
		it.Event = new(QuestDistributorOwnershipTransferred)
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
func (it *QuestDistributorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestDistributorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestDistributorOwnershipTransferred represents a OwnershipTransferred event raised by the QuestDistributor contract.
type QuestDistributorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestDistributor *QuestDistributorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuestDistributorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestDistributor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorOwnershipTransferredIterator{contract: _QuestDistributor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestDistributor *QuestDistributorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuestDistributorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestDistributor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestDistributorOwnershipTransferred)
				if err := _QuestDistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuestDistributor *QuestDistributorFilterer) ParseOwnershipTransferred(log types.Log) (*QuestDistributorOwnershipTransferred, error) {
	event := new(QuestDistributorOwnershipTransferred)
	if err := _QuestDistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestDistributorQuestPeriodUpdatedIterator is returned from FilterQuestPeriodUpdated and is used to iterate over the raw logs and unpacked data for QuestPeriodUpdated events raised by the QuestDistributor contract.
type QuestDistributorQuestPeriodUpdatedIterator struct {
	Event *QuestDistributorQuestPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *QuestDistributorQuestPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestDistributorQuestPeriodUpdated)
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
		it.Event = new(QuestDistributorQuestPeriodUpdated)
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
func (it *QuestDistributorQuestPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestDistributorQuestPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestDistributorQuestPeriodUpdated represents a QuestPeriodUpdated event raised by the QuestDistributor contract.
type QuestDistributorQuestPeriodUpdated struct {
	QuestID    *big.Int
	Period     *big.Int
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestPeriodUpdated is a free log retrieval operation binding the contract event 0x1d3b13bb6ce947c806a76a3476eafd33ec751e6ed2cbbdf7ab8503a4fb76a1ee.
//
// Solidity: event QuestPeriodUpdated(uint256 indexed questID, uint256 indexed period, bytes32 merkleRoot)
func (_QuestDistributor *QuestDistributorFilterer) FilterQuestPeriodUpdated(opts *bind.FilterOpts, questID []*big.Int, period []*big.Int) (*QuestDistributorQuestPeriodUpdatedIterator, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _QuestDistributor.contract.FilterLogs(opts, "QuestPeriodUpdated", questIDRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &QuestDistributorQuestPeriodUpdatedIterator{contract: _QuestDistributor.contract, event: "QuestPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchQuestPeriodUpdated is a free log subscription operation binding the contract event 0x1d3b13bb6ce947c806a76a3476eafd33ec751e6ed2cbbdf7ab8503a4fb76a1ee.
//
// Solidity: event QuestPeriodUpdated(uint256 indexed questID, uint256 indexed period, bytes32 merkleRoot)
func (_QuestDistributor *QuestDistributorFilterer) WatchQuestPeriodUpdated(opts *bind.WatchOpts, sink chan<- *QuestDistributorQuestPeriodUpdated, questID []*big.Int, period []*big.Int) (event.Subscription, error) {

	var questIDRule []interface{}
	for _, questIDItem := range questID {
		questIDRule = append(questIDRule, questIDItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _QuestDistributor.contract.WatchLogs(opts, "QuestPeriodUpdated", questIDRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestDistributorQuestPeriodUpdated)
				if err := _QuestDistributor.contract.UnpackLog(event, "QuestPeriodUpdated", log); err != nil {
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

// ParseQuestPeriodUpdated is a log parse operation binding the contract event 0x1d3b13bb6ce947c806a76a3476eafd33ec751e6ed2cbbdf7ab8503a4fb76a1ee.
//
// Solidity: event QuestPeriodUpdated(uint256 indexed questID, uint256 indexed period, bytes32 merkleRoot)
func (_QuestDistributor *QuestDistributorFilterer) ParseQuestPeriodUpdated(log types.Log) (*QuestDistributorQuestPeriodUpdated, error) {
	event := new(QuestDistributorQuestPeriodUpdated)
	if err := _QuestDistributor.contract.UnpackLog(event, "QuestPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
