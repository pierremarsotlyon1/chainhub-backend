// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voter

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

// VoterMetaData contains all meta data concerning the Voter contract.
var VoterMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"},{\"name\":\"_yeaPct\",\"type\":\"uint256\"},{\"name\":\"_nayPct\",\"type\":\"uint256\"},{\"name\":\"_executesIfDecided\",\"type\":\"bool\"}],\"name\":\"votePct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"canCreateNewVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MODIFY_QUORUM_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"name\":\"open\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"},{\"name\":\"startDate\",\"type\":\"uint64\"},{\"name\":\"snapshotBlock\",\"type\":\"uint64\"},{\"name\":\"supportRequired\",\"type\":\"uint64\"},{\"name\":\"minAcceptQuorum\",\"type\":\"uint64\"},{\"name\":\"yea\",\"type\":\"uint256\"},{\"name\":\"nay\",\"type\":\"uint256\"},{\"name\":\"votingPower\",\"type\":\"uint256\"},{\"name\":\"script\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minAcceptQuorumPct\",\"type\":\"uint64\"}],\"name\":\"changeMinAcceptQuorumPct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MODIFY_SUPPORT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastCreateVoteTimes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_supportRequiredPct\",\"type\":\"uint64\"}],\"name\":\"changeSupportRequiredPct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ENABLE_VOTE_CREATION\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBalanceUpperLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minTime\",\"type\":\"uint256\"}],\"name\":\"setMinTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_supportRequiredPct\",\"type\":\"uint64\"},{\"name\":\"_minAcceptQuorumPct\",\"type\":\"uint64\"},{\"name\":\"_voteTime\",\"type\":\"uint64\"},{\"name\":\"_minBalance\",\"type\":\"uint256\"},{\"name\":\"_minTime\",\"type\":\"uint256\"},{\"name\":\"_minBalanceLowerLimit\",\"type\":\"uint256\"},{\"name\":\"_minBalanceUpperLimit\",\"type\":\"uint256\"},{\"name\":\"_minTimeLowerLimit\",\"type\":\"uint256\"},{\"name\":\"_minTimeUpperLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_MIN_BALANCE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minTimeLowerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CREATE_VOTES_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"canForward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_MIN_TIME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minBalance\",\"type\":\"uint256\"}],\"name\":\"setMinBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"}],\"name\":\"canExecute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"canVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_executionScript\",\"type\":\"bytes\"},{\"name\":\"_metadata\",\"type\":\"string\"}],\"name\":\"newVote\",\"outputs\":[{\"name\":\"voteId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evmScript\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableVoteCreationOnce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minAcceptQuorumPct\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteData\",\"type\":\"uint256\"},{\"name\":\"_supports\",\"type\":\"bool\"},{\"name\":\"_executesIfDecided\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableVoteCreationOnce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minTimeUpperLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBalanceLowerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enableVoteCreation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_executionScript\",\"type\":\"bytes\"},{\"name\":\"_metadata\",\"type\":\"string\"},{\"name\":\"_castVote\",\"type\":\"bool\"},{\"name\":\"_executesIfDecided\",\"type\":\"bool\"}],\"name\":\"newVote\",\"outputs\":[{\"name\":\"voteId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteId\",\"type\":\"uint256\"}],\"name\":\"executeVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supportRequiredPct\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PCT_BASE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForwarder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DISABLE_VOTE_CREATION\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"voteId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"minBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creatorVotingPower\",\"type\":\"uint256\"}],\"name\":\"StartVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"voteId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"supports\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"CastVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"voteId\",\"type\":\"uint256\"}],\"name\":\"ExecuteVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"supportRequiredPct\",\"type\":\"uint64\"}],\"name\":\"ChangeSupportRequired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minAcceptQuorumPct\",\"type\":\"uint64\"}],\"name\":\"ChangeMinQuorum\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minBalance\",\"type\":\"uint256\"}],\"name\":\"MinimumBalanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minTime\",\"type\":\"uint256\"}],\"name\":\"MinimumTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"}]",
}

// VoterABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterMetaData.ABI instead.
var VoterABI = VoterMetaData.ABI

// Voter is an auto generated Go binding around an Ethereum contract.
type Voter struct {
	VoterCaller     // Read-only binding to the contract
	VoterTransactor // Write-only binding to the contract
	VoterFilterer   // Log filterer for contract events
}

// VoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterSession struct {
	Contract     *Voter            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterCallerSession struct {
	Contract *VoterCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterTransactorSession struct {
	Contract     *VoterTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterRaw struct {
	Contract *Voter // Generic contract binding to access the raw methods on
}

// VoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterCallerRaw struct {
	Contract *VoterCaller // Generic read-only contract binding to access the raw methods on
}

// VoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterTransactorRaw struct {
	Contract *VoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoter creates a new instance of Voter, bound to a specific deployed contract.
func NewVoter(address common.Address, backend bind.ContractBackend) (*Voter, error) {
	contract, err := bindVoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voter{VoterCaller: VoterCaller{contract: contract}, VoterTransactor: VoterTransactor{contract: contract}, VoterFilterer: VoterFilterer{contract: contract}}, nil
}

// NewVoterCaller creates a new read-only instance of Voter, bound to a specific deployed contract.
func NewVoterCaller(address common.Address, caller bind.ContractCaller) (*VoterCaller, error) {
	contract, err := bindVoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterCaller{contract: contract}, nil
}

// NewVoterTransactor creates a new write-only instance of Voter, bound to a specific deployed contract.
func NewVoterTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterTransactor, error) {
	contract, err := bindVoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterTransactor{contract: contract}, nil
}

// NewVoterFilterer creates a new log filterer instance of Voter, bound to a specific deployed contract.
func NewVoterFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterFilterer, error) {
	contract, err := bindVoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterFilterer{contract: contract}, nil
}

// bindVoter binds a generic wrapper to an already deployed contract.
func bindVoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voter *VoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voter.Contract.VoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voter *VoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.Contract.VoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voter *VoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voter.Contract.VoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voter *VoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voter *VoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voter *VoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voter.Contract.contract.Transact(opts, method, params...)
}

// CREATEVOTESROLE is a free data retrieval call binding the contract method 0xbe2c64d4.
//
// Solidity: function CREATE_VOTES_ROLE() view returns(bytes32)
func (_Voter *VoterCaller) CREATEVOTESROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "CREATE_VOTES_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATEVOTESROLE is a free data retrieval call binding the contract method 0xbe2c64d4.
//
// Solidity: function CREATE_VOTES_ROLE() view returns(bytes32)
func (_Voter *VoterSession) CREATEVOTESROLE() ([32]byte, error) {
	return _Voter.Contract.CREATEVOTESROLE(&_Voter.CallOpts)
}

// CREATEVOTESROLE is a free data retrieval call binding the contract method 0xbe2c64d4.
//
// Solidity: function CREATE_VOTES_ROLE() view returns(bytes32)
func (_Voter *VoterCallerSession) CREATEVOTESROLE() ([32]byte, error) {
	return _Voter.Contract.CREATEVOTESROLE(&_Voter.CallOpts)
}

// DISABLEVOTECREATION is a free data retrieval call binding the contract method 0xff311838.
//
// Solidity: function DISABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterCaller) DISABLEVOTECREATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "DISABLE_VOTE_CREATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISABLEVOTECREATION is a free data retrieval call binding the contract method 0xff311838.
//
// Solidity: function DISABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterSession) DISABLEVOTECREATION() ([32]byte, error) {
	return _Voter.Contract.DISABLEVOTECREATION(&_Voter.CallOpts)
}

// DISABLEVOTECREATION is a free data retrieval call binding the contract method 0xff311838.
//
// Solidity: function DISABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterCallerSession) DISABLEVOTECREATION() ([32]byte, error) {
	return _Voter.Contract.DISABLEVOTECREATION(&_Voter.CallOpts)
}

// ENABLEVOTECREATION is a free data retrieval call binding the contract method 0x868ce31a.
//
// Solidity: function ENABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterCaller) ENABLEVOTECREATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "ENABLE_VOTE_CREATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENABLEVOTECREATION is a free data retrieval call binding the contract method 0x868ce31a.
//
// Solidity: function ENABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterSession) ENABLEVOTECREATION() ([32]byte, error) {
	return _Voter.Contract.ENABLEVOTECREATION(&_Voter.CallOpts)
}

// ENABLEVOTECREATION is a free data retrieval call binding the contract method 0x868ce31a.
//
// Solidity: function ENABLE_VOTE_CREATION() view returns(bytes32)
func (_Voter *VoterCallerSession) ENABLEVOTECREATION() ([32]byte, error) {
	return _Voter.Contract.ENABLEVOTECREATION(&_Voter.CallOpts)
}

// MODIFYQUORUMROLE is a free data retrieval call binding the contract method 0x3c624c75.
//
// Solidity: function MODIFY_QUORUM_ROLE() view returns(bytes32)
func (_Voter *VoterCaller) MODIFYQUORUMROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "MODIFY_QUORUM_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODIFYQUORUMROLE is a free data retrieval call binding the contract method 0x3c624c75.
//
// Solidity: function MODIFY_QUORUM_ROLE() view returns(bytes32)
func (_Voter *VoterSession) MODIFYQUORUMROLE() ([32]byte, error) {
	return _Voter.Contract.MODIFYQUORUMROLE(&_Voter.CallOpts)
}

// MODIFYQUORUMROLE is a free data retrieval call binding the contract method 0x3c624c75.
//
// Solidity: function MODIFY_QUORUM_ROLE() view returns(bytes32)
func (_Voter *VoterCallerSession) MODIFYQUORUMROLE() ([32]byte, error) {
	return _Voter.Contract.MODIFYQUORUMROLE(&_Voter.CallOpts)
}

// MODIFYSUPPORTROLE is a free data retrieval call binding the contract method 0x62de7e5a.
//
// Solidity: function MODIFY_SUPPORT_ROLE() view returns(bytes32)
func (_Voter *VoterCaller) MODIFYSUPPORTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "MODIFY_SUPPORT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODIFYSUPPORTROLE is a free data retrieval call binding the contract method 0x62de7e5a.
//
// Solidity: function MODIFY_SUPPORT_ROLE() view returns(bytes32)
func (_Voter *VoterSession) MODIFYSUPPORTROLE() ([32]byte, error) {
	return _Voter.Contract.MODIFYSUPPORTROLE(&_Voter.CallOpts)
}

// MODIFYSUPPORTROLE is a free data retrieval call binding the contract method 0x62de7e5a.
//
// Solidity: function MODIFY_SUPPORT_ROLE() view returns(bytes32)
func (_Voter *VoterCallerSession) MODIFYSUPPORTROLE() ([32]byte, error) {
	return _Voter.Contract.MODIFYSUPPORTROLE(&_Voter.CallOpts)
}

// PCTBASE is a free data retrieval call binding the contract method 0xfc157cb4.
//
// Solidity: function PCT_BASE() view returns(uint64)
func (_Voter *VoterCaller) PCTBASE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "PCT_BASE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PCTBASE is a free data retrieval call binding the contract method 0xfc157cb4.
//
// Solidity: function PCT_BASE() view returns(uint64)
func (_Voter *VoterSession) PCTBASE() (uint64, error) {
	return _Voter.Contract.PCTBASE(&_Voter.CallOpts)
}

// PCTBASE is a free data retrieval call binding the contract method 0xfc157cb4.
//
// Solidity: function PCT_BASE() view returns(uint64)
func (_Voter *VoterCallerSession) PCTBASE() (uint64, error) {
	return _Voter.Contract.PCTBASE(&_Voter.CallOpts)
}

// SETMINBALANCEROLE is a free data retrieval call binding the contract method 0xb6f68254.
//
// Solidity: function SET_MIN_BALANCE_ROLE() view returns(bytes32)
func (_Voter *VoterCaller) SETMINBALANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "SET_MIN_BALANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETMINBALANCEROLE is a free data retrieval call binding the contract method 0xb6f68254.
//
// Solidity: function SET_MIN_BALANCE_ROLE() view returns(bytes32)
func (_Voter *VoterSession) SETMINBALANCEROLE() ([32]byte, error) {
	return _Voter.Contract.SETMINBALANCEROLE(&_Voter.CallOpts)
}

// SETMINBALANCEROLE is a free data retrieval call binding the contract method 0xb6f68254.
//
// Solidity: function SET_MIN_BALANCE_ROLE() view returns(bytes32)
func (_Voter *VoterCallerSession) SETMINBALANCEROLE() ([32]byte, error) {
	return _Voter.Contract.SETMINBALANCEROLE(&_Voter.CallOpts)
}

// SETMINTIMEROLE is a free data retrieval call binding the contract method 0xc54c24b5.
//
// Solidity: function SET_MIN_TIME_ROLE() view returns(bytes32)
func (_Voter *VoterCaller) SETMINTIMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "SET_MIN_TIME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETMINTIMEROLE is a free data retrieval call binding the contract method 0xc54c24b5.
//
// Solidity: function SET_MIN_TIME_ROLE() view returns(bytes32)
func (_Voter *VoterSession) SETMINTIMEROLE() ([32]byte, error) {
	return _Voter.Contract.SETMINTIMEROLE(&_Voter.CallOpts)
}

// SETMINTIMEROLE is a free data retrieval call binding the contract method 0xc54c24b5.
//
// Solidity: function SET_MIN_TIME_ROLE() view returns(bytes32)
func (_Voter *VoterCallerSession) SETMINTIMEROLE() ([32]byte, error) {
	return _Voter.Contract.SETMINTIMEROLE(&_Voter.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Voter *VoterCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Voter *VoterSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Voter.Contract.AllowRecoverability(&_Voter.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Voter *VoterCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Voter.Contract.AllowRecoverability(&_Voter.CallOpts, token)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Voter *VoterCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Voter *VoterSession) AppId() ([32]byte, error) {
	return _Voter.Contract.AppId(&_Voter.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Voter *VoterCallerSession) AppId() ([32]byte, error) {
	return _Voter.Contract.AppId(&_Voter.CallOpts)
}

// CanCreateNewVote is a free data retrieval call binding the contract method 0x2d573501.
//
// Solidity: function canCreateNewVote(address _sender) view returns(bool)
func (_Voter *VoterCaller) CanCreateNewVote(opts *bind.CallOpts, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "canCreateNewVote", _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCreateNewVote is a free data retrieval call binding the contract method 0x2d573501.
//
// Solidity: function canCreateNewVote(address _sender) view returns(bool)
func (_Voter *VoterSession) CanCreateNewVote(_sender common.Address) (bool, error) {
	return _Voter.Contract.CanCreateNewVote(&_Voter.CallOpts, _sender)
}

// CanCreateNewVote is a free data retrieval call binding the contract method 0x2d573501.
//
// Solidity: function canCreateNewVote(address _sender) view returns(bool)
func (_Voter *VoterCallerSession) CanCreateNewVote(_sender common.Address) (bool, error) {
	return _Voter.Contract.CanCreateNewVote(&_Voter.CallOpts, _sender)
}

// CanExecute is a free data retrieval call binding the contract method 0xcc63604a.
//
// Solidity: function canExecute(uint256 _voteId) view returns(bool)
func (_Voter *VoterCaller) CanExecute(opts *bind.CallOpts, _voteId *big.Int) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "canExecute", _voteId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExecute is a free data retrieval call binding the contract method 0xcc63604a.
//
// Solidity: function canExecute(uint256 _voteId) view returns(bool)
func (_Voter *VoterSession) CanExecute(_voteId *big.Int) (bool, error) {
	return _Voter.Contract.CanExecute(&_Voter.CallOpts, _voteId)
}

// CanExecute is a free data retrieval call binding the contract method 0xcc63604a.
//
// Solidity: function canExecute(uint256 _voteId) view returns(bool)
func (_Voter *VoterCallerSession) CanExecute(_voteId *big.Int) (bool, error) {
	return _Voter.Contract.CanExecute(&_Voter.CallOpts, _voteId)
}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes ) view returns(bool)
func (_Voter *VoterCaller) CanForward(opts *bind.CallOpts, _sender common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "canForward", _sender, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes ) view returns(bool)
func (_Voter *VoterSession) CanForward(_sender common.Address, arg1 []byte) (bool, error) {
	return _Voter.Contract.CanForward(&_Voter.CallOpts, _sender, arg1)
}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes ) view returns(bool)
func (_Voter *VoterCallerSession) CanForward(_sender common.Address, arg1 []byte) (bool, error) {
	return _Voter.Contract.CanForward(&_Voter.CallOpts, _sender, arg1)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Voter *VoterCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Voter *VoterSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Voter.Contract.CanPerform(&_Voter.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Voter *VoterCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Voter.Contract.CanPerform(&_Voter.CallOpts, _sender, _role, _params)
}

// CanVote is a free data retrieval call binding the contract method 0xcdb2867b.
//
// Solidity: function canVote(uint256 _voteId, address _voter) view returns(bool)
func (_Voter *VoterCaller) CanVote(opts *bind.CallOpts, _voteId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "canVote", _voteId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanVote is a free data retrieval call binding the contract method 0xcdb2867b.
//
// Solidity: function canVote(uint256 _voteId, address _voter) view returns(bool)
func (_Voter *VoterSession) CanVote(_voteId *big.Int, _voter common.Address) (bool, error) {
	return _Voter.Contract.CanVote(&_Voter.CallOpts, _voteId, _voter)
}

// CanVote is a free data retrieval call binding the contract method 0xcdb2867b.
//
// Solidity: function canVote(uint256 _voteId, address _voter) view returns(bool)
func (_Voter *VoterCallerSession) CanVote(_voteId *big.Int, _voter common.Address) (bool, error) {
	return _Voter.Contract.CanVote(&_Voter.CallOpts, _voteId, _voter)
}

// EnableVoteCreation is a free data retrieval call binding the contract method 0xee0f160e.
//
// Solidity: function enableVoteCreation() view returns(bool)
func (_Voter *VoterCaller) EnableVoteCreation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "enableVoteCreation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnableVoteCreation is a free data retrieval call binding the contract method 0xee0f160e.
//
// Solidity: function enableVoteCreation() view returns(bool)
func (_Voter *VoterSession) EnableVoteCreation() (bool, error) {
	return _Voter.Contract.EnableVoteCreation(&_Voter.CallOpts)
}

// EnableVoteCreation is a free data retrieval call binding the contract method 0xee0f160e.
//
// Solidity: function enableVoteCreation() view returns(bool)
func (_Voter *VoterCallerSession) EnableVoteCreation() (bool, error) {
	return _Voter.Contract.EnableVoteCreation(&_Voter.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Voter *VoterCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Voter *VoterSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Voter.Contract.GetEVMScriptExecutor(&_Voter.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Voter *VoterCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Voter.Contract.GetEVMScriptExecutor(&_Voter.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Voter *VoterCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Voter *VoterSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Voter.Contract.GetEVMScriptRegistry(&_Voter.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Voter *VoterCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Voter.Contract.GetEVMScriptRegistry(&_Voter.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Voter *VoterCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Voter *VoterSession) GetInitializationBlock() (*big.Int, error) {
	return _Voter.Contract.GetInitializationBlock(&_Voter.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Voter *VoterCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _Voter.Contract.GetInitializationBlock(&_Voter.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Voter *VoterCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Voter *VoterSession) GetRecoveryVault() (common.Address, error) {
	return _Voter.Contract.GetRecoveryVault(&_Voter.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Voter *VoterCallerSession) GetRecoveryVault() (common.Address, error) {
	return _Voter.Contract.GetRecoveryVault(&_Voter.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(bool open, bool executed, uint64 startDate, uint64 snapshotBlock, uint64 supportRequired, uint64 minAcceptQuorum, uint256 yea, uint256 nay, uint256 votingPower, bytes script)
func (_Voter *VoterCaller) GetVote(opts *bind.CallOpts, _voteId *big.Int) (struct {
	Open            bool
	Executed        bool
	StartDate       uint64
	SnapshotBlock   uint64
	SupportRequired uint64
	MinAcceptQuorum uint64
	Yea             *big.Int
	Nay             *big.Int
	VotingPower     *big.Int
	Script          []byte
}, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getVote", _voteId)

	outstruct := new(struct {
		Open            bool
		Executed        bool
		StartDate       uint64
		SnapshotBlock   uint64
		SupportRequired uint64
		MinAcceptQuorum uint64
		Yea             *big.Int
		Nay             *big.Int
		VotingPower     *big.Int
		Script          []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Open = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Executed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.StartDate = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.SnapshotBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.SupportRequired = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.MinAcceptQuorum = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Yea = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Nay = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VotingPower = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Script = *abi.ConvertType(out[9], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(bool open, bool executed, uint64 startDate, uint64 snapshotBlock, uint64 supportRequired, uint64 minAcceptQuorum, uint256 yea, uint256 nay, uint256 votingPower, bytes script)
func (_Voter *VoterSession) GetVote(_voteId *big.Int) (struct {
	Open            bool
	Executed        bool
	StartDate       uint64
	SnapshotBlock   uint64
	SupportRequired uint64
	MinAcceptQuorum uint64
	Yea             *big.Int
	Nay             *big.Int
	VotingPower     *big.Int
	Script          []byte
}, error) {
	return _Voter.Contract.GetVote(&_Voter.CallOpts, _voteId)
}

// GetVote is a free data retrieval call binding the contract method 0x5a55c1f0.
//
// Solidity: function getVote(uint256 _voteId) view returns(bool open, bool executed, uint64 startDate, uint64 snapshotBlock, uint64 supportRequired, uint64 minAcceptQuorum, uint256 yea, uint256 nay, uint256 votingPower, bytes script)
func (_Voter *VoterCallerSession) GetVote(_voteId *big.Int) (struct {
	Open            bool
	Executed        bool
	StartDate       uint64
	SnapshotBlock   uint64
	SupportRequired uint64
	MinAcceptQuorum uint64
	Yea             *big.Int
	Nay             *big.Int
	VotingPower     *big.Int
	Script          []byte
}, error) {
	return _Voter.Contract.GetVote(&_Voter.CallOpts, _voteId)
}

// GetVoterState is a free data retrieval call binding the contract method 0x4b12311c.
//
// Solidity: function getVoterState(uint256 _voteId, address _voter) view returns(uint8)
func (_Voter *VoterCaller) GetVoterState(opts *bind.CallOpts, _voteId *big.Int, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "getVoterState", _voteId, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetVoterState is a free data retrieval call binding the contract method 0x4b12311c.
//
// Solidity: function getVoterState(uint256 _voteId, address _voter) view returns(uint8)
func (_Voter *VoterSession) GetVoterState(_voteId *big.Int, _voter common.Address) (uint8, error) {
	return _Voter.Contract.GetVoterState(&_Voter.CallOpts, _voteId, _voter)
}

// GetVoterState is a free data retrieval call binding the contract method 0x4b12311c.
//
// Solidity: function getVoterState(uint256 _voteId, address _voter) view returns(uint8)
func (_Voter *VoterCallerSession) GetVoterState(_voteId *big.Int, _voter common.Address) (uint8, error) {
	return _Voter.Contract.GetVoterState(&_Voter.CallOpts, _voteId, _voter)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Voter *VoterCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Voter *VoterSession) HasInitialized() (bool, error) {
	return _Voter.Contract.HasInitialized(&_Voter.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Voter *VoterCallerSession) HasInitialized() (bool, error) {
	return _Voter.Contract.HasInitialized(&_Voter.CallOpts)
}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Voter *VoterCaller) IsForwarder(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "isForwarder")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Voter *VoterSession) IsForwarder() (bool, error) {
	return _Voter.Contract.IsForwarder(&_Voter.CallOpts)
}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Voter *VoterCallerSession) IsForwarder() (bool, error) {
	return _Voter.Contract.IsForwarder(&_Voter.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Voter *VoterCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Voter *VoterSession) IsPetrified() (bool, error) {
	return _Voter.Contract.IsPetrified(&_Voter.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Voter *VoterCallerSession) IsPetrified() (bool, error) {
	return _Voter.Contract.IsPetrified(&_Voter.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Voter *VoterCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Voter *VoterSession) Kernel() (common.Address, error) {
	return _Voter.Contract.Kernel(&_Voter.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Voter *VoterCallerSession) Kernel() (common.Address, error) {
	return _Voter.Contract.Kernel(&_Voter.CallOpts)
}

// LastCreateVoteTimes is a free data retrieval call binding the contract method 0x7bbddcfc.
//
// Solidity: function lastCreateVoteTimes(address ) view returns(uint256)
func (_Voter *VoterCaller) LastCreateVoteTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "lastCreateVoteTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCreateVoteTimes is a free data retrieval call binding the contract method 0x7bbddcfc.
//
// Solidity: function lastCreateVoteTimes(address ) view returns(uint256)
func (_Voter *VoterSession) LastCreateVoteTimes(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.LastCreateVoteTimes(&_Voter.CallOpts, arg0)
}

// LastCreateVoteTimes is a free data retrieval call binding the contract method 0x7bbddcfc.
//
// Solidity: function lastCreateVoteTimes(address ) view returns(uint256)
func (_Voter *VoterCallerSession) LastCreateVoteTimes(arg0 common.Address) (*big.Int, error) {
	return _Voter.Contract.LastCreateVoteTimes(&_Voter.CallOpts, arg0)
}

// MinAcceptQuorumPct is a free data retrieval call binding the contract method 0xdc474b1a.
//
// Solidity: function minAcceptQuorumPct() view returns(uint64)
func (_Voter *VoterCaller) MinAcceptQuorumPct(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minAcceptQuorumPct")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinAcceptQuorumPct is a free data retrieval call binding the contract method 0xdc474b1a.
//
// Solidity: function minAcceptQuorumPct() view returns(uint64)
func (_Voter *VoterSession) MinAcceptQuorumPct() (uint64, error) {
	return _Voter.Contract.MinAcceptQuorumPct(&_Voter.CallOpts)
}

// MinAcceptQuorumPct is a free data retrieval call binding the contract method 0xdc474b1a.
//
// Solidity: function minAcceptQuorumPct() view returns(uint64)
func (_Voter *VoterCallerSession) MinAcceptQuorumPct() (uint64, error) {
	return _Voter.Contract.MinAcceptQuorumPct(&_Voter.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Voter *VoterCaller) MinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Voter *VoterSession) MinBalance() (*big.Int, error) {
	return _Voter.Contract.MinBalance(&_Voter.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Voter *VoterCallerSession) MinBalance() (*big.Int, error) {
	return _Voter.Contract.MinBalance(&_Voter.CallOpts)
}

// MinBalanceLowerLimit is a free data retrieval call binding the contract method 0xeb1e3b77.
//
// Solidity: function minBalanceLowerLimit() view returns(uint256)
func (_Voter *VoterCaller) MinBalanceLowerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minBalanceLowerLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBalanceLowerLimit is a free data retrieval call binding the contract method 0xeb1e3b77.
//
// Solidity: function minBalanceLowerLimit() view returns(uint256)
func (_Voter *VoterSession) MinBalanceLowerLimit() (*big.Int, error) {
	return _Voter.Contract.MinBalanceLowerLimit(&_Voter.CallOpts)
}

// MinBalanceLowerLimit is a free data retrieval call binding the contract method 0xeb1e3b77.
//
// Solidity: function minBalanceLowerLimit() view returns(uint256)
func (_Voter *VoterCallerSession) MinBalanceLowerLimit() (*big.Int, error) {
	return _Voter.Contract.MinBalanceLowerLimit(&_Voter.CallOpts)
}

// MinBalanceUpperLimit is a free data retrieval call binding the contract method 0x8b568327.
//
// Solidity: function minBalanceUpperLimit() view returns(uint256)
func (_Voter *VoterCaller) MinBalanceUpperLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minBalanceUpperLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBalanceUpperLimit is a free data retrieval call binding the contract method 0x8b568327.
//
// Solidity: function minBalanceUpperLimit() view returns(uint256)
func (_Voter *VoterSession) MinBalanceUpperLimit() (*big.Int, error) {
	return _Voter.Contract.MinBalanceUpperLimit(&_Voter.CallOpts)
}

// MinBalanceUpperLimit is a free data retrieval call binding the contract method 0x8b568327.
//
// Solidity: function minBalanceUpperLimit() view returns(uint256)
func (_Voter *VoterCallerSession) MinBalanceUpperLimit() (*big.Int, error) {
	return _Voter.Contract.MinBalanceUpperLimit(&_Voter.CallOpts)
}

// MinTime is a free data retrieval call binding the contract method 0x1aa43078.
//
// Solidity: function minTime() view returns(uint256)
func (_Voter *VoterCaller) MinTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTime is a free data retrieval call binding the contract method 0x1aa43078.
//
// Solidity: function minTime() view returns(uint256)
func (_Voter *VoterSession) MinTime() (*big.Int, error) {
	return _Voter.Contract.MinTime(&_Voter.CallOpts)
}

// MinTime is a free data retrieval call binding the contract method 0x1aa43078.
//
// Solidity: function minTime() view returns(uint256)
func (_Voter *VoterCallerSession) MinTime() (*big.Int, error) {
	return _Voter.Contract.MinTime(&_Voter.CallOpts)
}

// MinTimeLowerLimit is a free data retrieval call binding the contract method 0xb8795725.
//
// Solidity: function minTimeLowerLimit() view returns(uint256)
func (_Voter *VoterCaller) MinTimeLowerLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minTimeLowerLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimeLowerLimit is a free data retrieval call binding the contract method 0xb8795725.
//
// Solidity: function minTimeLowerLimit() view returns(uint256)
func (_Voter *VoterSession) MinTimeLowerLimit() (*big.Int, error) {
	return _Voter.Contract.MinTimeLowerLimit(&_Voter.CallOpts)
}

// MinTimeLowerLimit is a free data retrieval call binding the contract method 0xb8795725.
//
// Solidity: function minTimeLowerLimit() view returns(uint256)
func (_Voter *VoterCallerSession) MinTimeLowerLimit() (*big.Int, error) {
	return _Voter.Contract.MinTimeLowerLimit(&_Voter.CallOpts)
}

// MinTimeUpperLimit is a free data retrieval call binding the contract method 0xdf352197.
//
// Solidity: function minTimeUpperLimit() view returns(uint256)
func (_Voter *VoterCaller) MinTimeUpperLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "minTimeUpperLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimeUpperLimit is a free data retrieval call binding the contract method 0xdf352197.
//
// Solidity: function minTimeUpperLimit() view returns(uint256)
func (_Voter *VoterSession) MinTimeUpperLimit() (*big.Int, error) {
	return _Voter.Contract.MinTimeUpperLimit(&_Voter.CallOpts)
}

// MinTimeUpperLimit is a free data retrieval call binding the contract method 0xdf352197.
//
// Solidity: function minTimeUpperLimit() view returns(uint256)
func (_Voter *VoterCallerSession) MinTimeUpperLimit() (*big.Int, error) {
	return _Voter.Contract.MinTimeUpperLimit(&_Voter.CallOpts)
}

// SupportRequiredPct is a free data retrieval call binding the contract method 0xfad167ab.
//
// Solidity: function supportRequiredPct() view returns(uint64)
func (_Voter *VoterCaller) SupportRequiredPct(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "supportRequiredPct")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SupportRequiredPct is a free data retrieval call binding the contract method 0xfad167ab.
//
// Solidity: function supportRequiredPct() view returns(uint64)
func (_Voter *VoterSession) SupportRequiredPct() (uint64, error) {
	return _Voter.Contract.SupportRequiredPct(&_Voter.CallOpts)
}

// SupportRequiredPct is a free data retrieval call binding the contract method 0xfad167ab.
//
// Solidity: function supportRequiredPct() view returns(uint64)
func (_Voter *VoterCallerSession) SupportRequiredPct() (uint64, error) {
	return _Voter.Contract.SupportRequiredPct(&_Voter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Voter *VoterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Voter *VoterSession) Token() (common.Address, error) {
	return _Voter.Contract.Token(&_Voter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Voter *VoterCallerSession) Token() (common.Address, error) {
	return _Voter.Contract.Token(&_Voter.CallOpts)
}

// VoteTime is a free data retrieval call binding the contract method 0xbcf93dd6.
//
// Solidity: function voteTime() view returns(uint64)
func (_Voter *VoterCaller) VoteTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "voteTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoteTime is a free data retrieval call binding the contract method 0xbcf93dd6.
//
// Solidity: function voteTime() view returns(uint64)
func (_Voter *VoterSession) VoteTime() (uint64, error) {
	return _Voter.Contract.VoteTime(&_Voter.CallOpts)
}

// VoteTime is a free data retrieval call binding the contract method 0xbcf93dd6.
//
// Solidity: function voteTime() view returns(uint64)
func (_Voter *VoterCallerSession) VoteTime() (uint64, error) {
	return _Voter.Contract.VoteTime(&_Voter.CallOpts)
}

// VotesLength is a free data retrieval call binding the contract method 0xde4f6347.
//
// Solidity: function votesLength() view returns(uint256)
func (_Voter *VoterCaller) VotesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voter.contract.Call(opts, &out, "votesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesLength is a free data retrieval call binding the contract method 0xde4f6347.
//
// Solidity: function votesLength() view returns(uint256)
func (_Voter *VoterSession) VotesLength() (*big.Int, error) {
	return _Voter.Contract.VotesLength(&_Voter.CallOpts)
}

// VotesLength is a free data retrieval call binding the contract method 0xde4f6347.
//
// Solidity: function votesLength() view returns(uint256)
func (_Voter *VoterCallerSession) VotesLength() (*big.Int, error) {
	return _Voter.Contract.VotesLength(&_Voter.CallOpts)
}

// ChangeMinAcceptQuorumPct is a paid mutator transaction binding the contract method 0x5eb24332.
//
// Solidity: function changeMinAcceptQuorumPct(uint64 _minAcceptQuorumPct) returns()
func (_Voter *VoterTransactor) ChangeMinAcceptQuorumPct(opts *bind.TransactOpts, _minAcceptQuorumPct uint64) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "changeMinAcceptQuorumPct", _minAcceptQuorumPct)
}

// ChangeMinAcceptQuorumPct is a paid mutator transaction binding the contract method 0x5eb24332.
//
// Solidity: function changeMinAcceptQuorumPct(uint64 _minAcceptQuorumPct) returns()
func (_Voter *VoterSession) ChangeMinAcceptQuorumPct(_minAcceptQuorumPct uint64) (*types.Transaction, error) {
	return _Voter.Contract.ChangeMinAcceptQuorumPct(&_Voter.TransactOpts, _minAcceptQuorumPct)
}

// ChangeMinAcceptQuorumPct is a paid mutator transaction binding the contract method 0x5eb24332.
//
// Solidity: function changeMinAcceptQuorumPct(uint64 _minAcceptQuorumPct) returns()
func (_Voter *VoterTransactorSession) ChangeMinAcceptQuorumPct(_minAcceptQuorumPct uint64) (*types.Transaction, error) {
	return _Voter.Contract.ChangeMinAcceptQuorumPct(&_Voter.TransactOpts, _minAcceptQuorumPct)
}

// ChangeSupportRequiredPct is a paid mutator transaction binding the contract method 0x7c1d0b87.
//
// Solidity: function changeSupportRequiredPct(uint64 _supportRequiredPct) returns()
func (_Voter *VoterTransactor) ChangeSupportRequiredPct(opts *bind.TransactOpts, _supportRequiredPct uint64) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "changeSupportRequiredPct", _supportRequiredPct)
}

// ChangeSupportRequiredPct is a paid mutator transaction binding the contract method 0x7c1d0b87.
//
// Solidity: function changeSupportRequiredPct(uint64 _supportRequiredPct) returns()
func (_Voter *VoterSession) ChangeSupportRequiredPct(_supportRequiredPct uint64) (*types.Transaction, error) {
	return _Voter.Contract.ChangeSupportRequiredPct(&_Voter.TransactOpts, _supportRequiredPct)
}

// ChangeSupportRequiredPct is a paid mutator transaction binding the contract method 0x7c1d0b87.
//
// Solidity: function changeSupportRequiredPct(uint64 _supportRequiredPct) returns()
func (_Voter *VoterTransactorSession) ChangeSupportRequiredPct(_supportRequiredPct uint64) (*types.Transaction, error) {
	return _Voter.Contract.ChangeSupportRequiredPct(&_Voter.TransactOpts, _supportRequiredPct)
}

// DisableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdc177eec.
//
// Solidity: function disableVoteCreationOnce() returns()
func (_Voter *VoterTransactor) DisableVoteCreationOnce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "disableVoteCreationOnce")
}

// DisableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdc177eec.
//
// Solidity: function disableVoteCreationOnce() returns()
func (_Voter *VoterSession) DisableVoteCreationOnce() (*types.Transaction, error) {
	return _Voter.Contract.DisableVoteCreationOnce(&_Voter.TransactOpts)
}

// DisableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdc177eec.
//
// Solidity: function disableVoteCreationOnce() returns()
func (_Voter *VoterTransactorSession) DisableVoteCreationOnce() (*types.Transaction, error) {
	return _Voter.Contract.DisableVoteCreationOnce(&_Voter.TransactOpts)
}

// EnableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdf136602.
//
// Solidity: function enableVoteCreationOnce() returns()
func (_Voter *VoterTransactor) EnableVoteCreationOnce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "enableVoteCreationOnce")
}

// EnableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdf136602.
//
// Solidity: function enableVoteCreationOnce() returns()
func (_Voter *VoterSession) EnableVoteCreationOnce() (*types.Transaction, error) {
	return _Voter.Contract.EnableVoteCreationOnce(&_Voter.TransactOpts)
}

// EnableVoteCreationOnce is a paid mutator transaction binding the contract method 0xdf136602.
//
// Solidity: function enableVoteCreationOnce() returns()
func (_Voter *VoterTransactorSession) EnableVoteCreationOnce() (*types.Transaction, error) {
	return _Voter.Contract.EnableVoteCreationOnce(&_Voter.TransactOpts)
}

// ExecuteVote is a paid mutator transaction binding the contract method 0xf98a4eca.
//
// Solidity: function executeVote(uint256 _voteId) returns()
func (_Voter *VoterTransactor) ExecuteVote(opts *bind.TransactOpts, _voteId *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "executeVote", _voteId)
}

// ExecuteVote is a paid mutator transaction binding the contract method 0xf98a4eca.
//
// Solidity: function executeVote(uint256 _voteId) returns()
func (_Voter *VoterSession) ExecuteVote(_voteId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ExecuteVote(&_Voter.TransactOpts, _voteId)
}

// ExecuteVote is a paid mutator transaction binding the contract method 0xf98a4eca.
//
// Solidity: function executeVote(uint256 _voteId) returns()
func (_Voter *VoterTransactorSession) ExecuteVote(_voteId *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.ExecuteVote(&_Voter.TransactOpts, _voteId)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Voter *VoterTransactor) Forward(opts *bind.TransactOpts, _evmScript []byte) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "forward", _evmScript)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Voter *VoterSession) Forward(_evmScript []byte) (*types.Transaction, error) {
	return _Voter.Contract.Forward(&_Voter.TransactOpts, _evmScript)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Voter *VoterTransactorSession) Forward(_evmScript []byte) (*types.Transaction, error) {
	return _Voter.Contract.Forward(&_Voter.TransactOpts, _evmScript)
}

// Initialize is a paid mutator transaction binding the contract method 0xa87f8624.
//
// Solidity: function initialize(address _token, uint64 _supportRequiredPct, uint64 _minAcceptQuorumPct, uint64 _voteTime, uint256 _minBalance, uint256 _minTime, uint256 _minBalanceLowerLimit, uint256 _minBalanceUpperLimit, uint256 _minTimeLowerLimit, uint256 _minTimeUpperLimit) returns()
func (_Voter *VoterTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _supportRequiredPct uint64, _minAcceptQuorumPct uint64, _voteTime uint64, _minBalance *big.Int, _minTime *big.Int, _minBalanceLowerLimit *big.Int, _minBalanceUpperLimit *big.Int, _minTimeLowerLimit *big.Int, _minTimeUpperLimit *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "initialize", _token, _supportRequiredPct, _minAcceptQuorumPct, _voteTime, _minBalance, _minTime, _minBalanceLowerLimit, _minBalanceUpperLimit, _minTimeLowerLimit, _minTimeUpperLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xa87f8624.
//
// Solidity: function initialize(address _token, uint64 _supportRequiredPct, uint64 _minAcceptQuorumPct, uint64 _voteTime, uint256 _minBalance, uint256 _minTime, uint256 _minBalanceLowerLimit, uint256 _minBalanceUpperLimit, uint256 _minTimeLowerLimit, uint256 _minTimeUpperLimit) returns()
func (_Voter *VoterSession) Initialize(_token common.Address, _supportRequiredPct uint64, _minAcceptQuorumPct uint64, _voteTime uint64, _minBalance *big.Int, _minTime *big.Int, _minBalanceLowerLimit *big.Int, _minBalanceUpperLimit *big.Int, _minTimeLowerLimit *big.Int, _minTimeUpperLimit *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Initialize(&_Voter.TransactOpts, _token, _supportRequiredPct, _minAcceptQuorumPct, _voteTime, _minBalance, _minTime, _minBalanceLowerLimit, _minBalanceUpperLimit, _minTimeLowerLimit, _minTimeUpperLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xa87f8624.
//
// Solidity: function initialize(address _token, uint64 _supportRequiredPct, uint64 _minAcceptQuorumPct, uint64 _voteTime, uint256 _minBalance, uint256 _minTime, uint256 _minBalanceLowerLimit, uint256 _minBalanceUpperLimit, uint256 _minTimeLowerLimit, uint256 _minTimeUpperLimit) returns()
func (_Voter *VoterTransactorSession) Initialize(_token common.Address, _supportRequiredPct uint64, _minAcceptQuorumPct uint64, _voteTime uint64, _minBalance *big.Int, _minTime *big.Int, _minBalanceLowerLimit *big.Int, _minBalanceUpperLimit *big.Int, _minTimeLowerLimit *big.Int, _minTimeUpperLimit *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.Initialize(&_Voter.TransactOpts, _token, _supportRequiredPct, _minAcceptQuorumPct, _voteTime, _minBalance, _minTime, _minBalanceLowerLimit, _minBalanceUpperLimit, _minTimeLowerLimit, _minTimeUpperLimit)
}

// NewVote is a paid mutator transaction binding the contract method 0xd5db2c80.
//
// Solidity: function newVote(bytes _executionScript, string _metadata) returns(uint256 voteId)
func (_Voter *VoterTransactor) NewVote(opts *bind.TransactOpts, _executionScript []byte, _metadata string) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "newVote", _executionScript, _metadata)
}

// NewVote is a paid mutator transaction binding the contract method 0xd5db2c80.
//
// Solidity: function newVote(bytes _executionScript, string _metadata) returns(uint256 voteId)
func (_Voter *VoterSession) NewVote(_executionScript []byte, _metadata string) (*types.Transaction, error) {
	return _Voter.Contract.NewVote(&_Voter.TransactOpts, _executionScript, _metadata)
}

// NewVote is a paid mutator transaction binding the contract method 0xd5db2c80.
//
// Solidity: function newVote(bytes _executionScript, string _metadata) returns(uint256 voteId)
func (_Voter *VoterTransactorSession) NewVote(_executionScript []byte, _metadata string) (*types.Transaction, error) {
	return _Voter.Contract.NewVote(&_Voter.TransactOpts, _executionScript, _metadata)
}

// NewVote0 is a paid mutator transaction binding the contract method 0xf4b00513.
//
// Solidity: function newVote(bytes _executionScript, string _metadata, bool _castVote, bool _executesIfDecided) returns(uint256 voteId)
func (_Voter *VoterTransactor) NewVote0(opts *bind.TransactOpts, _executionScript []byte, _metadata string, _castVote bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "newVote0", _executionScript, _metadata, _castVote, _executesIfDecided)
}

// NewVote0 is a paid mutator transaction binding the contract method 0xf4b00513.
//
// Solidity: function newVote(bytes _executionScript, string _metadata, bool _castVote, bool _executesIfDecided) returns(uint256 voteId)
func (_Voter *VoterSession) NewVote0(_executionScript []byte, _metadata string, _castVote bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.NewVote0(&_Voter.TransactOpts, _executionScript, _metadata, _castVote, _executesIfDecided)
}

// NewVote0 is a paid mutator transaction binding the contract method 0xf4b00513.
//
// Solidity: function newVote(bytes _executionScript, string _metadata, bool _castVote, bool _executesIfDecided) returns(uint256 voteId)
func (_Voter *VoterTransactorSession) NewVote0(_executionScript []byte, _metadata string, _castVote bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.NewVote0(&_Voter.TransactOpts, _executionScript, _metadata, _castVote, _executesIfDecided)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(uint256 _minBalance) returns()
func (_Voter *VoterTransactor) SetMinBalance(opts *bind.TransactOpts, _minBalance *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "setMinBalance", _minBalance)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(uint256 _minBalance) returns()
func (_Voter *VoterSession) SetMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.SetMinBalance(&_Voter.TransactOpts, _minBalance)
}

// SetMinBalance is a paid mutator transaction binding the contract method 0xc91d956c.
//
// Solidity: function setMinBalance(uint256 _minBalance) returns()
func (_Voter *VoterTransactorSession) SetMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.SetMinBalance(&_Voter.TransactOpts, _minBalance)
}

// SetMinTime is a paid mutator transaction binding the contract method 0x93f89a80.
//
// Solidity: function setMinTime(uint256 _minTime) returns()
func (_Voter *VoterTransactor) SetMinTime(opts *bind.TransactOpts, _minTime *big.Int) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "setMinTime", _minTime)
}

// SetMinTime is a paid mutator transaction binding the contract method 0x93f89a80.
//
// Solidity: function setMinTime(uint256 _minTime) returns()
func (_Voter *VoterSession) SetMinTime(_minTime *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.SetMinTime(&_Voter.TransactOpts, _minTime)
}

// SetMinTime is a paid mutator transaction binding the contract method 0x93f89a80.
//
// Solidity: function setMinTime(uint256 _minTime) returns()
func (_Voter *VoterTransactorSession) SetMinTime(_minTime *big.Int) (*types.Transaction, error) {
	return _Voter.Contract.SetMinTime(&_Voter.TransactOpts, _minTime)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Voter *VoterTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Voter *VoterSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Voter.Contract.TransferToVault(&_Voter.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Voter *VoterTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Voter.Contract.TransferToVault(&_Voter.TransactOpts, _token)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _voteData, bool _supports, bool _executesIfDecided) returns()
func (_Voter *VoterTransactor) Vote(opts *bind.TransactOpts, _voteData *big.Int, _supports bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "vote", _voteData, _supports, _executesIfDecided)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _voteData, bool _supports, bool _executesIfDecided) returns()
func (_Voter *VoterSession) Vote(_voteData *big.Int, _supports bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.Vote(&_Voter.TransactOpts, _voteData, _supports, _executesIfDecided)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _voteData, bool _supports, bool _executesIfDecided) returns()
func (_Voter *VoterTransactorSession) Vote(_voteData *big.Int, _supports bool, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.Vote(&_Voter.TransactOpts, _voteData, _supports, _executesIfDecided)
}

// VotePct is a paid mutator transaction binding the contract method 0x1e526610.
//
// Solidity: function votePct(uint256 _voteId, uint256 _yeaPct, uint256 _nayPct, bool _executesIfDecided) returns()
func (_Voter *VoterTransactor) VotePct(opts *bind.TransactOpts, _voteId *big.Int, _yeaPct *big.Int, _nayPct *big.Int, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.contract.Transact(opts, "votePct", _voteId, _yeaPct, _nayPct, _executesIfDecided)
}

// VotePct is a paid mutator transaction binding the contract method 0x1e526610.
//
// Solidity: function votePct(uint256 _voteId, uint256 _yeaPct, uint256 _nayPct, bool _executesIfDecided) returns()
func (_Voter *VoterSession) VotePct(_voteId *big.Int, _yeaPct *big.Int, _nayPct *big.Int, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.VotePct(&_Voter.TransactOpts, _voteId, _yeaPct, _nayPct, _executesIfDecided)
}

// VotePct is a paid mutator transaction binding the contract method 0x1e526610.
//
// Solidity: function votePct(uint256 _voteId, uint256 _yeaPct, uint256 _nayPct, bool _executesIfDecided) returns()
func (_Voter *VoterTransactorSession) VotePct(_voteId *big.Int, _yeaPct *big.Int, _nayPct *big.Int, _executesIfDecided bool) (*types.Transaction, error) {
	return _Voter.Contract.VotePct(&_Voter.TransactOpts, _voteId, _yeaPct, _nayPct, _executesIfDecided)
}

// VoterCastVoteIterator is returned from FilterCastVote and is used to iterate over the raw logs and unpacked data for CastVote events raised by the Voter contract.
type VoterCastVoteIterator struct {
	Event *VoterCastVote // Event containing the contract specifics and raw log

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
func (it *VoterCastVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterCastVote)
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
		it.Event = new(VoterCastVote)
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
func (it *VoterCastVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterCastVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterCastVote represents a CastVote event raised by the Voter contract.
type VoterCastVote struct {
	VoteId   *big.Int
	Voter    common.Address
	Supports bool
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCastVote is a free log retrieval operation binding the contract event 0xb34ee265e3d4f5ec4e8b52d59b2a9be8fceca2f274ebc080d8fba797fea9391f.
//
// Solidity: event CastVote(uint256 indexed voteId, address indexed voter, bool supports, uint256 stake)
func (_Voter *VoterFilterer) FilterCastVote(opts *bind.FilterOpts, voteId []*big.Int, voter []common.Address) (*VoterCastVoteIterator, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "CastVote", voteIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &VoterCastVoteIterator{contract: _Voter.contract, event: "CastVote", logs: logs, sub: sub}, nil
}

// WatchCastVote is a free log subscription operation binding the contract event 0xb34ee265e3d4f5ec4e8b52d59b2a9be8fceca2f274ebc080d8fba797fea9391f.
//
// Solidity: event CastVote(uint256 indexed voteId, address indexed voter, bool supports, uint256 stake)
func (_Voter *VoterFilterer) WatchCastVote(opts *bind.WatchOpts, sink chan<- *VoterCastVote, voteId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "CastVote", voteIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterCastVote)
				if err := _Voter.contract.UnpackLog(event, "CastVote", log); err != nil {
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

// ParseCastVote is a log parse operation binding the contract event 0xb34ee265e3d4f5ec4e8b52d59b2a9be8fceca2f274ebc080d8fba797fea9391f.
//
// Solidity: event CastVote(uint256 indexed voteId, address indexed voter, bool supports, uint256 stake)
func (_Voter *VoterFilterer) ParseCastVote(log types.Log) (*VoterCastVote, error) {
	event := new(VoterCastVote)
	if err := _Voter.contract.UnpackLog(event, "CastVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterChangeMinQuorumIterator is returned from FilterChangeMinQuorum and is used to iterate over the raw logs and unpacked data for ChangeMinQuorum events raised by the Voter contract.
type VoterChangeMinQuorumIterator struct {
	Event *VoterChangeMinQuorum // Event containing the contract specifics and raw log

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
func (it *VoterChangeMinQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterChangeMinQuorum)
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
		it.Event = new(VoterChangeMinQuorum)
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
func (it *VoterChangeMinQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterChangeMinQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterChangeMinQuorum represents a ChangeMinQuorum event raised by the Voter contract.
type VoterChangeMinQuorum struct {
	MinAcceptQuorumPct uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangeMinQuorum is a free log retrieval operation binding the contract event 0x3172f2e9273c729c2a47cc8bf7e7f18506e3e3035126d562602bd2155bc78a50.
//
// Solidity: event ChangeMinQuorum(uint64 minAcceptQuorumPct)
func (_Voter *VoterFilterer) FilterChangeMinQuorum(opts *bind.FilterOpts) (*VoterChangeMinQuorumIterator, error) {

	logs, sub, err := _Voter.contract.FilterLogs(opts, "ChangeMinQuorum")
	if err != nil {
		return nil, err
	}
	return &VoterChangeMinQuorumIterator{contract: _Voter.contract, event: "ChangeMinQuorum", logs: logs, sub: sub}, nil
}

// WatchChangeMinQuorum is a free log subscription operation binding the contract event 0x3172f2e9273c729c2a47cc8bf7e7f18506e3e3035126d562602bd2155bc78a50.
//
// Solidity: event ChangeMinQuorum(uint64 minAcceptQuorumPct)
func (_Voter *VoterFilterer) WatchChangeMinQuorum(opts *bind.WatchOpts, sink chan<- *VoterChangeMinQuorum) (event.Subscription, error) {

	logs, sub, err := _Voter.contract.WatchLogs(opts, "ChangeMinQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterChangeMinQuorum)
				if err := _Voter.contract.UnpackLog(event, "ChangeMinQuorum", log); err != nil {
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

// ParseChangeMinQuorum is a log parse operation binding the contract event 0x3172f2e9273c729c2a47cc8bf7e7f18506e3e3035126d562602bd2155bc78a50.
//
// Solidity: event ChangeMinQuorum(uint64 minAcceptQuorumPct)
func (_Voter *VoterFilterer) ParseChangeMinQuorum(log types.Log) (*VoterChangeMinQuorum, error) {
	event := new(VoterChangeMinQuorum)
	if err := _Voter.contract.UnpackLog(event, "ChangeMinQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterChangeSupportRequiredIterator is returned from FilterChangeSupportRequired and is used to iterate over the raw logs and unpacked data for ChangeSupportRequired events raised by the Voter contract.
type VoterChangeSupportRequiredIterator struct {
	Event *VoterChangeSupportRequired // Event containing the contract specifics and raw log

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
func (it *VoterChangeSupportRequiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterChangeSupportRequired)
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
		it.Event = new(VoterChangeSupportRequired)
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
func (it *VoterChangeSupportRequiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterChangeSupportRequiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterChangeSupportRequired represents a ChangeSupportRequired event raised by the Voter contract.
type VoterChangeSupportRequired struct {
	SupportRequiredPct uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangeSupportRequired is a free log retrieval operation binding the contract event 0x903b617f7f36eb047a29b89d1bf7885fdae31d250c3320fccf11d045c11b396e.
//
// Solidity: event ChangeSupportRequired(uint64 supportRequiredPct)
func (_Voter *VoterFilterer) FilterChangeSupportRequired(opts *bind.FilterOpts) (*VoterChangeSupportRequiredIterator, error) {

	logs, sub, err := _Voter.contract.FilterLogs(opts, "ChangeSupportRequired")
	if err != nil {
		return nil, err
	}
	return &VoterChangeSupportRequiredIterator{contract: _Voter.contract, event: "ChangeSupportRequired", logs: logs, sub: sub}, nil
}

// WatchChangeSupportRequired is a free log subscription operation binding the contract event 0x903b617f7f36eb047a29b89d1bf7885fdae31d250c3320fccf11d045c11b396e.
//
// Solidity: event ChangeSupportRequired(uint64 supportRequiredPct)
func (_Voter *VoterFilterer) WatchChangeSupportRequired(opts *bind.WatchOpts, sink chan<- *VoterChangeSupportRequired) (event.Subscription, error) {

	logs, sub, err := _Voter.contract.WatchLogs(opts, "ChangeSupportRequired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterChangeSupportRequired)
				if err := _Voter.contract.UnpackLog(event, "ChangeSupportRequired", log); err != nil {
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

// ParseChangeSupportRequired is a log parse operation binding the contract event 0x903b617f7f36eb047a29b89d1bf7885fdae31d250c3320fccf11d045c11b396e.
//
// Solidity: event ChangeSupportRequired(uint64 supportRequiredPct)
func (_Voter *VoterFilterer) ParseChangeSupportRequired(log types.Log) (*VoterChangeSupportRequired, error) {
	event := new(VoterChangeSupportRequired)
	if err := _Voter.contract.UnpackLog(event, "ChangeSupportRequired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterExecuteVoteIterator is returned from FilterExecuteVote and is used to iterate over the raw logs and unpacked data for ExecuteVote events raised by the Voter contract.
type VoterExecuteVoteIterator struct {
	Event *VoterExecuteVote // Event containing the contract specifics and raw log

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
func (it *VoterExecuteVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterExecuteVote)
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
		it.Event = new(VoterExecuteVote)
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
func (it *VoterExecuteVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterExecuteVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterExecuteVote represents a ExecuteVote event raised by the Voter contract.
type VoterExecuteVote struct {
	VoteId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecuteVote is a free log retrieval operation binding the contract event 0xbf8e2b108bb7c980e08903a8a46527699d5e84905a082d56dacb4150725c8cab.
//
// Solidity: event ExecuteVote(uint256 indexed voteId)
func (_Voter *VoterFilterer) FilterExecuteVote(opts *bind.FilterOpts, voteId []*big.Int) (*VoterExecuteVoteIterator, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "ExecuteVote", voteIdRule)
	if err != nil {
		return nil, err
	}
	return &VoterExecuteVoteIterator{contract: _Voter.contract, event: "ExecuteVote", logs: logs, sub: sub}, nil
}

// WatchExecuteVote is a free log subscription operation binding the contract event 0xbf8e2b108bb7c980e08903a8a46527699d5e84905a082d56dacb4150725c8cab.
//
// Solidity: event ExecuteVote(uint256 indexed voteId)
func (_Voter *VoterFilterer) WatchExecuteVote(opts *bind.WatchOpts, sink chan<- *VoterExecuteVote, voteId []*big.Int) (event.Subscription, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "ExecuteVote", voteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterExecuteVote)
				if err := _Voter.contract.UnpackLog(event, "ExecuteVote", log); err != nil {
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

// ParseExecuteVote is a log parse operation binding the contract event 0xbf8e2b108bb7c980e08903a8a46527699d5e84905a082d56dacb4150725c8cab.
//
// Solidity: event ExecuteVote(uint256 indexed voteId)
func (_Voter *VoterFilterer) ParseExecuteVote(log types.Log) (*VoterExecuteVote, error) {
	event := new(VoterExecuteVote)
	if err := _Voter.contract.UnpackLog(event, "ExecuteVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterMinimumBalanceSetIterator is returned from FilterMinimumBalanceSet and is used to iterate over the raw logs and unpacked data for MinimumBalanceSet events raised by the Voter contract.
type VoterMinimumBalanceSetIterator struct {
	Event *VoterMinimumBalanceSet // Event containing the contract specifics and raw log

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
func (it *VoterMinimumBalanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterMinimumBalanceSet)
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
		it.Event = new(VoterMinimumBalanceSet)
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
func (it *VoterMinimumBalanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterMinimumBalanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterMinimumBalanceSet represents a MinimumBalanceSet event raised by the Voter contract.
type VoterMinimumBalanceSet struct {
	MinBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinimumBalanceSet is a free log retrieval operation binding the contract event 0xbd5318adc778260bd213cc16f3ef06cfb4f729adb9309495fed0e10abc61c375.
//
// Solidity: event MinimumBalanceSet(uint256 minBalance)
func (_Voter *VoterFilterer) FilterMinimumBalanceSet(opts *bind.FilterOpts) (*VoterMinimumBalanceSetIterator, error) {

	logs, sub, err := _Voter.contract.FilterLogs(opts, "MinimumBalanceSet")
	if err != nil {
		return nil, err
	}
	return &VoterMinimumBalanceSetIterator{contract: _Voter.contract, event: "MinimumBalanceSet", logs: logs, sub: sub}, nil
}

// WatchMinimumBalanceSet is a free log subscription operation binding the contract event 0xbd5318adc778260bd213cc16f3ef06cfb4f729adb9309495fed0e10abc61c375.
//
// Solidity: event MinimumBalanceSet(uint256 minBalance)
func (_Voter *VoterFilterer) WatchMinimumBalanceSet(opts *bind.WatchOpts, sink chan<- *VoterMinimumBalanceSet) (event.Subscription, error) {

	logs, sub, err := _Voter.contract.WatchLogs(opts, "MinimumBalanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterMinimumBalanceSet)
				if err := _Voter.contract.UnpackLog(event, "MinimumBalanceSet", log); err != nil {
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

// ParseMinimumBalanceSet is a log parse operation binding the contract event 0xbd5318adc778260bd213cc16f3ef06cfb4f729adb9309495fed0e10abc61c375.
//
// Solidity: event MinimumBalanceSet(uint256 minBalance)
func (_Voter *VoterFilterer) ParseMinimumBalanceSet(log types.Log) (*VoterMinimumBalanceSet, error) {
	event := new(VoterMinimumBalanceSet)
	if err := _Voter.contract.UnpackLog(event, "MinimumBalanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterMinimumTimeSetIterator is returned from FilterMinimumTimeSet and is used to iterate over the raw logs and unpacked data for MinimumTimeSet events raised by the Voter contract.
type VoterMinimumTimeSetIterator struct {
	Event *VoterMinimumTimeSet // Event containing the contract specifics and raw log

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
func (it *VoterMinimumTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterMinimumTimeSet)
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
		it.Event = new(VoterMinimumTimeSet)
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
func (it *VoterMinimumTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterMinimumTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterMinimumTimeSet represents a MinimumTimeSet event raised by the Voter contract.
type VoterMinimumTimeSet struct {
	MinTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinimumTimeSet is a free log retrieval operation binding the contract event 0xcb34d0577abe5b2f96b65ea1bb5de2209ba6309c6909438c6d50dd277ca3b580.
//
// Solidity: event MinimumTimeSet(uint256 minTime)
func (_Voter *VoterFilterer) FilterMinimumTimeSet(opts *bind.FilterOpts) (*VoterMinimumTimeSetIterator, error) {

	logs, sub, err := _Voter.contract.FilterLogs(opts, "MinimumTimeSet")
	if err != nil {
		return nil, err
	}
	return &VoterMinimumTimeSetIterator{contract: _Voter.contract, event: "MinimumTimeSet", logs: logs, sub: sub}, nil
}

// WatchMinimumTimeSet is a free log subscription operation binding the contract event 0xcb34d0577abe5b2f96b65ea1bb5de2209ba6309c6909438c6d50dd277ca3b580.
//
// Solidity: event MinimumTimeSet(uint256 minTime)
func (_Voter *VoterFilterer) WatchMinimumTimeSet(opts *bind.WatchOpts, sink chan<- *VoterMinimumTimeSet) (event.Subscription, error) {

	logs, sub, err := _Voter.contract.WatchLogs(opts, "MinimumTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterMinimumTimeSet)
				if err := _Voter.contract.UnpackLog(event, "MinimumTimeSet", log); err != nil {
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

// ParseMinimumTimeSet is a log parse operation binding the contract event 0xcb34d0577abe5b2f96b65ea1bb5de2209ba6309c6909438c6d50dd277ca3b580.
//
// Solidity: event MinimumTimeSet(uint256 minTime)
func (_Voter *VoterFilterer) ParseMinimumTimeSet(log types.Log) (*VoterMinimumTimeSet, error) {
	event := new(VoterMinimumTimeSet)
	if err := _Voter.contract.UnpackLog(event, "MinimumTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the Voter contract.
type VoterRecoverToVaultIterator struct {
	Event *VoterRecoverToVault // Event containing the contract specifics and raw log

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
func (it *VoterRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRecoverToVault)
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
		it.Event = new(VoterRecoverToVault)
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
func (it *VoterRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRecoverToVault represents a RecoverToVault event raised by the Voter contract.
type VoterRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Voter *VoterFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*VoterRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VoterRecoverToVaultIterator{contract: _Voter.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Voter *VoterFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *VoterRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRecoverToVault)
				if err := _Voter.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Voter *VoterFilterer) ParseRecoverToVault(log types.Log) (*VoterRecoverToVault, error) {
	event := new(VoterRecoverToVault)
	if err := _Voter.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the Voter contract.
type VoterScriptResultIterator struct {
	Event *VoterScriptResult // Event containing the contract specifics and raw log

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
func (it *VoterScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterScriptResult)
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
		it.Event = new(VoterScriptResult)
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
func (it *VoterScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterScriptResult represents a ScriptResult event raised by the Voter contract.
type VoterScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Voter *VoterFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*VoterScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &VoterScriptResultIterator{contract: _Voter.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Voter *VoterFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *VoterScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterScriptResult)
				if err := _Voter.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Voter *VoterFilterer) ParseScriptResult(log types.Log) (*VoterScriptResult, error) {
	event := new(VoterScriptResult)
	if err := _Voter.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterStartVoteIterator is returned from FilterStartVote and is used to iterate over the raw logs and unpacked data for StartVote events raised by the Voter contract.
type VoterStartVoteIterator struct {
	Event *VoterStartVote // Event containing the contract specifics and raw log

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
func (it *VoterStartVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterStartVote)
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
		it.Event = new(VoterStartVote)
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
func (it *VoterStartVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterStartVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterStartVote represents a StartVote event raised by the Voter contract.
type VoterStartVote struct {
	VoteId             *big.Int
	Creator            common.Address
	Metadata           string
	MinBalance         *big.Int
	MinTime            *big.Int
	TotalSupply        *big.Int
	CreatorVotingPower *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStartVote is a free log retrieval operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed voteId, address indexed creator, string metadata, uint256 minBalance, uint256 minTime, uint256 totalSupply, uint256 creatorVotingPower)
func (_Voter *VoterFilterer) FilterStartVote(opts *bind.FilterOpts, voteId []*big.Int, creator []common.Address) (*VoterStartVoteIterator, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Voter.contract.FilterLogs(opts, "StartVote", voteIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &VoterStartVoteIterator{contract: _Voter.contract, event: "StartVote", logs: logs, sub: sub}, nil
}

// WatchStartVote is a free log subscription operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed voteId, address indexed creator, string metadata, uint256 minBalance, uint256 minTime, uint256 totalSupply, uint256 creatorVotingPower)
func (_Voter *VoterFilterer) WatchStartVote(opts *bind.WatchOpts, sink chan<- *VoterStartVote, voteId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var voteIdRule []interface{}
	for _, voteIdItem := range voteId {
		voteIdRule = append(voteIdRule, voteIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Voter.contract.WatchLogs(opts, "StartVote", voteIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterStartVote)
				if err := _Voter.contract.UnpackLog(event, "StartVote", log); err != nil {
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

// ParseStartVote is a log parse operation binding the contract event 0x0730610a5322c6584fb6f5bb760719e650c888cfd965a2beb2d598bcd425e394.
//
// Solidity: event StartVote(uint256 indexed voteId, address indexed creator, string metadata, uint256 minBalance, uint256 minTime, uint256 totalSupply, uint256 creatorVotingPower)
func (_Voter *VoterFilterer) ParseStartVote(log types.Log) (*VoterStartVote, error) {
	event := new(VoterStartVote)
	if err := _Voter.contract.UnpackLog(event, "StartVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
