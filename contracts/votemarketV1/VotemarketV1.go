// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votemarketV1

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

// PlatformBribe is an auto generated low-level Go binding around an user-defined struct.
type PlatformBribe struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	Blacklist         []common.Address
}

// PlatformPeriod is an auto generated low-level Go binding around an user-defined struct.
type PlatformPeriod struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}

// PlatformUpgrade is an auto generated low-level Go binding around an user-defined struct.
type PlatformUpgrade struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
	Blacklist         []common.Address
}

// VotemarketV1MetaData contains all meta data concerning the VotemarketV1 contract.
var VotemarketV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gaugeController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ALREADY_INCREASED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER_ONLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_NUMBER_OF_PERIODS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KILLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_ALLOWED_OPERATION\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_UPGRADEABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_PERIODS_LEFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WRONG_INPUT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingReward\",\"type\":\"uint256\"}],\"name\":\"BribeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isUpgradeable\",\"type\":\"bool\"}],\"name\":\"BribeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"name\":\"BribeDurationIncreaseQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"name\":\"BribeDurationIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"PeriodRolledOver\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMUM_PERIOD\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activePeriod\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"amountClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bribes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"claimAllFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claimFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"closeBribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"upgradeable\",\"type\":\"bool\"}],\"name\":\"createBribe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBribeID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugeController\",\"outputs\":[{\"internalType\":\"contractGaugeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getActivePeriod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structPlatform.Period\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getActivePeriodPerBribe\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getBlacklistedAddressesForBribe\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getBribe\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"}],\"internalType\":\"structPlatform.Bribe\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getPeriodsLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodsLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getUpgradedBribeQueued\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"}],\"internalType\":\"structPlatform.Upgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bribeId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_additionnalPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_increasedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newMaxPricePerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addressesBlacklisted\",\"type\":\"address[]\"}],\"name\":\"increaseBribeDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isKilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isUpgradeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastUserClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardPerVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"updateBribePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"updateBribePeriods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"updateManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"upgradeBribeQueue\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotemarketV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use VotemarketV1MetaData.ABI instead.
var VotemarketV1ABI = VotemarketV1MetaData.ABI

// VotemarketV1 is an auto generated Go binding around an Ethereum contract.
type VotemarketV1 struct {
	VotemarketV1Caller     // Read-only binding to the contract
	VotemarketV1Transactor // Write-only binding to the contract
	VotemarketV1Filterer   // Log filterer for contract events
}

// VotemarketV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type VotemarketV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VotemarketV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotemarketV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotemarketV1Session struct {
	Contract     *VotemarketV1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotemarketV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotemarketV1CallerSession struct {
	Contract *VotemarketV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VotemarketV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotemarketV1TransactorSession struct {
	Contract     *VotemarketV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VotemarketV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type VotemarketV1Raw struct {
	Contract *VotemarketV1 // Generic contract binding to access the raw methods on
}

// VotemarketV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotemarketV1CallerRaw struct {
	Contract *VotemarketV1Caller // Generic read-only contract binding to access the raw methods on
}

// VotemarketV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotemarketV1TransactorRaw struct {
	Contract *VotemarketV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVotemarketV1 creates a new instance of VotemarketV1, bound to a specific deployed contract.
func NewVotemarketV1(address common.Address, backend bind.ContractBackend) (*VotemarketV1, error) {
	contract, err := bindVotemarketV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1{VotemarketV1Caller: VotemarketV1Caller{contract: contract}, VotemarketV1Transactor: VotemarketV1Transactor{contract: contract}, VotemarketV1Filterer: VotemarketV1Filterer{contract: contract}}, nil
}

// NewVotemarketV1Caller creates a new read-only instance of VotemarketV1, bound to a specific deployed contract.
func NewVotemarketV1Caller(address common.Address, caller bind.ContractCaller) (*VotemarketV1Caller, error) {
	contract, err := bindVotemarketV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1Caller{contract: contract}, nil
}

// NewVotemarketV1Transactor creates a new write-only instance of VotemarketV1, bound to a specific deployed contract.
func NewVotemarketV1Transactor(address common.Address, transactor bind.ContractTransactor) (*VotemarketV1Transactor, error) {
	contract, err := bindVotemarketV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1Transactor{contract: contract}, nil
}

// NewVotemarketV1Filterer creates a new log filterer instance of VotemarketV1, bound to a specific deployed contract.
func NewVotemarketV1Filterer(address common.Address, filterer bind.ContractFilterer) (*VotemarketV1Filterer, error) {
	contract, err := bindVotemarketV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1Filterer{contract: contract}, nil
}

// bindVotemarketV1 binds a generic wrapper to an already deployed contract.
func bindVotemarketV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotemarketV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotemarketV1 *VotemarketV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotemarketV1.Contract.VotemarketV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotemarketV1 *VotemarketV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV1.Contract.VotemarketV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotemarketV1 *VotemarketV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotemarketV1.Contract.VotemarketV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotemarketV1 *VotemarketV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotemarketV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotemarketV1 *VotemarketV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotemarketV1 *VotemarketV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotemarketV1.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV1 *VotemarketV1Caller) MINIMUMPERIOD(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "MINIMUM_PERIOD")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV1 *VotemarketV1Session) MINIMUMPERIOD() (uint8, error) {
	return _VotemarketV1.Contract.MINIMUMPERIOD(&_VotemarketV1.CallOpts)
}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV1 *VotemarketV1CallerSession) MINIMUMPERIOD() (uint8, error) {
	return _VotemarketV1.Contract.MINIMUMPERIOD(&_VotemarketV1.CallOpts)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x21bf936a.
//
// Solidity: function activePeriod(uint256 ) view returns(uint8 id, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1Caller) ActivePeriod(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "activePeriod", arg0)

	outstruct := new(struct {
		Id              uint8
		Timestamp       *big.Int
		RewardPerPeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardPerPeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActivePeriod is a free data retrieval call binding the contract method 0x21bf936a.
//
// Solidity: function activePeriod(uint256 ) view returns(uint8 id, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1Session) ActivePeriod(arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	return _VotemarketV1.Contract.ActivePeriod(&_VotemarketV1.CallOpts, arg0)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x21bf936a.
//
// Solidity: function activePeriod(uint256 ) view returns(uint8 id, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1CallerSession) ActivePeriod(arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	return _VotemarketV1.Contract.ActivePeriod(&_VotemarketV1.CallOpts, arg0)
}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Caller) AmountClaimed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "amountClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) AmountClaimed(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.AmountClaimed(&_VotemarketV1.CallOpts, arg0)
}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1CallerSession) AmountClaimed(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.AmountClaimed(&_VotemarketV1.CallOpts, arg0)
}

// Bribes is a free data retrieval call binding the contract method 0xa9d46e5b.
//
// Solidity: function bribes(uint256 ) view returns(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 endTimestamp, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VotemarketV1 *VotemarketV1Caller) Bribes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "bribes", arg0)

	outstruct := new(struct {
		Gauge             common.Address
		Manager           common.Address
		RewardToken       common.Address
		NumberOfPeriods   uint8
		EndTimestamp      *big.Int
		MaxRewardPerVote  *big.Int
		TotalRewardAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gauge = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Manager = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RewardToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NumberOfPeriods = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.EndTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxRewardPerVote = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalRewardAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bribes is a free data retrieval call binding the contract method 0xa9d46e5b.
//
// Solidity: function bribes(uint256 ) view returns(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 endTimestamp, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VotemarketV1 *VotemarketV1Session) Bribes(arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	return _VotemarketV1.Contract.Bribes(&_VotemarketV1.CallOpts, arg0)
}

// Bribes is a free data retrieval call binding the contract method 0xa9d46e5b.
//
// Solidity: function bribes(uint256 ) view returns(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 endTimestamp, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VotemarketV1 *VotemarketV1CallerSession) Bribes(arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	return _VotemarketV1.Contract.Bribes(&_VotemarketV1.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV1 *VotemarketV1Caller) Claimable(opts *bind.CallOpts, user common.Address, bribeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "claimable", user, bribeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV1 *VotemarketV1Session) Claimable(user common.Address, bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.Claimable(&_VotemarketV1.CallOpts, user, bribeId)
}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV1 *VotemarketV1CallerSession) Claimable(user common.Address, bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.Claimable(&_VotemarketV1.CallOpts, user, bribeId)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VotemarketV1 *VotemarketV1Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VotemarketV1 *VotemarketV1Session) Factory() (common.Address, error) {
	return _VotemarketV1.Contract.Factory(&_VotemarketV1.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VotemarketV1 *VotemarketV1CallerSession) Factory() (common.Address, error) {
	return _VotemarketV1.Contract.Factory(&_VotemarketV1.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV1 *VotemarketV1Caller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "gaugeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV1 *VotemarketV1Session) GaugeController() (common.Address, error) {
	return _VotemarketV1.Contract.GaugeController(&_VotemarketV1.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV1 *VotemarketV1CallerSession) GaugeController() (common.Address, error) {
	return _VotemarketV1.Contract.GaugeController(&_VotemarketV1.CallOpts)
}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV1 *VotemarketV1Caller) GetActivePeriod(opts *bind.CallOpts, bribeId *big.Int) (PlatformPeriod, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getActivePeriod", bribeId)

	if err != nil {
		return *new(PlatformPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformPeriod)).(*PlatformPeriod)

	return out0, err

}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV1 *VotemarketV1Session) GetActivePeriod(bribeId *big.Int) (PlatformPeriod, error) {
	return _VotemarketV1.Contract.GetActivePeriod(&_VotemarketV1.CallOpts, bribeId)
}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV1 *VotemarketV1CallerSession) GetActivePeriod(bribeId *big.Int) (PlatformPeriod, error) {
	return _VotemarketV1.Contract.GetActivePeriod(&_VotemarketV1.CallOpts, bribeId)
}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV1 *VotemarketV1Caller) GetActivePeriodPerBribe(opts *bind.CallOpts, bribeId *big.Int) (uint8, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getActivePeriodPerBribe", bribeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV1 *VotemarketV1Session) GetActivePeriodPerBribe(bribeId *big.Int) (uint8, error) {
	return _VotemarketV1.Contract.GetActivePeriodPerBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV1 *VotemarketV1CallerSession) GetActivePeriodPerBribe(bribeId *big.Int) (uint8, error) {
	return _VotemarketV1.Contract.GetActivePeriodPerBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV1 *VotemarketV1Caller) GetBlacklistedAddressesForBribe(opts *bind.CallOpts, bribeId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getBlacklistedAddressesForBribe", bribeId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV1 *VotemarketV1Session) GetBlacklistedAddressesForBribe(bribeId *big.Int) ([]common.Address, error) {
	return _VotemarketV1.Contract.GetBlacklistedAddressesForBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV1 *VotemarketV1CallerSession) GetBlacklistedAddressesForBribe(bribeId *big.Int) ([]common.Address, error) {
	return _VotemarketV1.Contract.GetBlacklistedAddressesForBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1Caller) GetBribe(opts *bind.CallOpts, bribeId *big.Int) (PlatformBribe, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getBribe", bribeId)

	if err != nil {
		return *new(PlatformBribe), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformBribe)).(*PlatformBribe)

	return out0, err

}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1Session) GetBribe(bribeId *big.Int) (PlatformBribe, error) {
	return _VotemarketV1.Contract.GetBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1CallerSession) GetBribe(bribeId *big.Int) (PlatformBribe, error) {
	return _VotemarketV1.Contract.GetBribe(&_VotemarketV1.CallOpts, bribeId)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV1 *VotemarketV1Caller) GetCurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getCurrentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) GetCurrentPeriod() (*big.Int, error) {
	return _VotemarketV1.Contract.GetCurrentPeriod(&_VotemarketV1.CallOpts)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV1 *VotemarketV1CallerSession) GetCurrentPeriod() (*big.Int, error) {
	return _VotemarketV1.Contract.GetCurrentPeriod(&_VotemarketV1.CallOpts)
}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV1 *VotemarketV1Caller) GetPeriodsLeft(opts *bind.CallOpts, bribeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getPeriodsLeft", bribeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV1 *VotemarketV1Session) GetPeriodsLeft(bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.GetPeriodsLeft(&_VotemarketV1.CallOpts, bribeId)
}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV1 *VotemarketV1CallerSession) GetPeriodsLeft(bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.GetPeriodsLeft(&_VotemarketV1.CallOpts, bribeId)
}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1Caller) GetUpgradedBribeQueued(opts *bind.CallOpts, bribeId *big.Int) (PlatformUpgrade, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "getUpgradedBribeQueued", bribeId)

	if err != nil {
		return *new(PlatformUpgrade), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformUpgrade)).(*PlatformUpgrade)

	return out0, err

}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1Session) GetUpgradedBribeQueued(bribeId *big.Int) (PlatformUpgrade, error) {
	return _VotemarketV1.Contract.GetUpgradedBribeQueued(&_VotemarketV1.CallOpts, bribeId)
}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV1 *VotemarketV1CallerSession) GetUpgradedBribeQueued(bribeId *big.Int) (PlatformUpgrade, error) {
	return _VotemarketV1.Contract.GetUpgradedBribeQueued(&_VotemarketV1.CallOpts, bribeId)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV1 *VotemarketV1Caller) IsBlacklisted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "isBlacklisted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV1 *VotemarketV1Session) IsBlacklisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotemarketV1.Contract.IsBlacklisted(&_VotemarketV1.CallOpts, arg0, arg1)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV1 *VotemarketV1CallerSession) IsBlacklisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotemarketV1.Contract.IsBlacklisted(&_VotemarketV1.CallOpts, arg0, arg1)
}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV1 *VotemarketV1Caller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "isKilled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV1 *VotemarketV1Session) IsKilled() (bool, error) {
	return _VotemarketV1.Contract.IsKilled(&_VotemarketV1.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV1 *VotemarketV1CallerSession) IsKilled() (bool, error) {
	return _VotemarketV1.Contract.IsKilled(&_VotemarketV1.CallOpts)
}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV1 *VotemarketV1Caller) IsUpgradeable(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "isUpgradeable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV1 *VotemarketV1Session) IsUpgradeable(arg0 *big.Int) (bool, error) {
	return _VotemarketV1.Contract.IsUpgradeable(&_VotemarketV1.CallOpts, arg0)
}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV1 *VotemarketV1CallerSession) IsUpgradeable(arg0 *big.Int) (bool, error) {
	return _VotemarketV1.Contract.IsUpgradeable(&_VotemarketV1.CallOpts, arg0)
}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Caller) LastUserClaim(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "lastUserClaim", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) LastUserClaim(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.LastUserClaim(&_VotemarketV1.CallOpts, arg0, arg1)
}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1CallerSession) LastUserClaim(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.LastUserClaim(&_VotemarketV1.CallOpts, arg0, arg1)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV1 *VotemarketV1Caller) NextID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "nextID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) NextID() (*big.Int, error) {
	return _VotemarketV1.Contract.NextID(&_VotemarketV1.CallOpts)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV1 *VotemarketV1CallerSession) NextID() (*big.Int, error) {
	return _VotemarketV1.Contract.NextID(&_VotemarketV1.CallOpts)
}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Caller) RewardPerVote(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "rewardPerVote", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) RewardPerVote(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.RewardPerVote(&_VotemarketV1.CallOpts, arg0)
}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV1 *VotemarketV1CallerSession) RewardPerVote(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV1.Contract.RewardPerVote(&_VotemarketV1.CallOpts, arg0)
}

// UpgradeBribeQueue is a free data retrieval call binding the contract method 0x83287273.
//
// Solidity: function upgradeBribeQueue(uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VotemarketV1 *VotemarketV1Caller) UpgradeBribeQueue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV1.contract.Call(opts, &out, "upgradeBribeQueue", arg0)

	outstruct := new(struct {
		NumberOfPeriods   uint8
		TotalRewardAmount *big.Int
		MaxRewardPerVote  *big.Int
		EndTimestamp      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumberOfPeriods = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TotalRewardAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxRewardPerVote = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UpgradeBribeQueue is a free data retrieval call binding the contract method 0x83287273.
//
// Solidity: function upgradeBribeQueue(uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VotemarketV1 *VotemarketV1Session) UpgradeBribeQueue(arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VotemarketV1.Contract.UpgradeBribeQueue(&_VotemarketV1.CallOpts, arg0)
}

// UpgradeBribeQueue is a free data retrieval call binding the contract method 0x83287273.
//
// Solidity: function upgradeBribeQueue(uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VotemarketV1 *VotemarketV1CallerSession) UpgradeBribeQueue(arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VotemarketV1.Contract.UpgradeBribeQueue(&_VotemarketV1.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1Transactor) Claim(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "claim", bribeId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) Claim(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.Claim(&_VotemarketV1.TransactOpts, bribeId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1TransactorSession) Claim(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.Claim(&_VotemarketV1.TransactOpts, bribeId)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Transactor) ClaimAll(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "claimAll", ids)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Session) ClaimAll(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimAll(&_VotemarketV1.TransactOpts, ids)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) ClaimAll(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimAll(&_VotemarketV1.TransactOpts, ids)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address user, uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Transactor) ClaimAllFor(opts *bind.TransactOpts, user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "claimAllFor", user, ids)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address user, uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Session) ClaimAllFor(user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimAllFor(&_VotemarketV1.TransactOpts, user, ids)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address user, uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) ClaimAllFor(user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimAllFor(&_VotemarketV1.TransactOpts, user, ids)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1Transactor) ClaimFor(opts *bind.TransactOpts, user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "claimFor", user, bribeId)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1Session) ClaimFor(user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimFor(&_VotemarketV1.TransactOpts, user, bribeId)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV1 *VotemarketV1TransactorSession) ClaimFor(user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.ClaimFor(&_VotemarketV1.TransactOpts, user, bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1Transactor) CloseBribe(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "closeBribe", bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1Session) CloseBribe(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.CloseBribe(&_VotemarketV1.TransactOpts, bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) CloseBribe(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.CloseBribe(&_VotemarketV1.TransactOpts, bribeId)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV1 *VotemarketV1Transactor) CreateBribe(opts *bind.TransactOpts, gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "createBribe", gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV1 *VotemarketV1Session) CreateBribe(gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV1.Contract.CreateBribe(&_VotemarketV1.TransactOpts, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV1 *VotemarketV1TransactorSession) CreateBribe(gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV1.Contract.CreateBribe(&_VotemarketV1.TransactOpts, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV1 *VotemarketV1Transactor) IncreaseBribeDuration(opts *bind.TransactOpts, _bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "increaseBribeDuration", _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV1 *VotemarketV1Session) IncreaseBribeDuration(_bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV1.Contract.IncreaseBribeDuration(&_VotemarketV1.TransactOpts, _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) IncreaseBribeDuration(_bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV1.Contract.IncreaseBribeDuration(&_VotemarketV1.TransactOpts, _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV1 *VotemarketV1Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV1 *VotemarketV1Session) Kill() (*types.Transaction, error) {
	return _VotemarketV1.Contract.Kill(&_VotemarketV1.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) Kill() (*types.Transaction, error) {
	return _VotemarketV1.Contract.Kill(&_VotemarketV1.TransactOpts)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1Transactor) UpdateBribePeriod(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "updateBribePeriod", bribeId)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1Session) UpdateBribePeriod(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateBribePeriod(&_VotemarketV1.TransactOpts, bribeId)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) UpdateBribePeriod(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateBribePeriod(&_VotemarketV1.TransactOpts, bribeId)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Transactor) UpdateBribePeriods(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "updateBribePeriods", ids)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1Session) UpdateBribePeriods(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateBribePeriods(&_VotemarketV1.TransactOpts, ids)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) UpdateBribePeriods(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateBribePeriods(&_VotemarketV1.TransactOpts, ids)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV1 *VotemarketV1Transactor) UpdateManager(opts *bind.TransactOpts, bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV1.contract.Transact(opts, "updateManager", bribeId, newManager)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV1 *VotemarketV1Session) UpdateManager(bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateManager(&_VotemarketV1.TransactOpts, bribeId, newManager)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV1 *VotemarketV1TransactorSession) UpdateManager(bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV1.Contract.UpdateManager(&_VotemarketV1.TransactOpts, bribeId, newManager)
}

// VotemarketV1BribeClosedIterator is returned from FilterBribeClosed and is used to iterate over the raw logs and unpacked data for BribeClosed events raised by the VotemarketV1 contract.
type VotemarketV1BribeClosedIterator struct {
	Event *VotemarketV1BribeClosed // Event containing the contract specifics and raw log

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
func (it *VotemarketV1BribeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1BribeClosed)
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
		it.Event = new(VotemarketV1BribeClosed)
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
func (it *VotemarketV1BribeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1BribeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1BribeClosed represents a BribeClosed event raised by the VotemarketV1 contract.
type VotemarketV1BribeClosed struct {
	Id              *big.Int
	RemainingReward *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBribeClosed is a free log retrieval operation binding the contract event 0x046a3d8b38f161b27b53792783d95179ee019fafaf2aedf362807cb52fc2ab46.
//
// Solidity: event BribeClosed(uint256 id, uint256 remainingReward)
func (_VotemarketV1 *VotemarketV1Filterer) FilterBribeClosed(opts *bind.FilterOpts) (*VotemarketV1BribeClosedIterator, error) {

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "BribeClosed")
	if err != nil {
		return nil, err
	}
	return &VotemarketV1BribeClosedIterator{contract: _VotemarketV1.contract, event: "BribeClosed", logs: logs, sub: sub}, nil
}

// WatchBribeClosed is a free log subscription operation binding the contract event 0x046a3d8b38f161b27b53792783d95179ee019fafaf2aedf362807cb52fc2ab46.
//
// Solidity: event BribeClosed(uint256 id, uint256 remainingReward)
func (_VotemarketV1 *VotemarketV1Filterer) WatchBribeClosed(opts *bind.WatchOpts, sink chan<- *VotemarketV1BribeClosed) (event.Subscription, error) {

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "BribeClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1BribeClosed)
				if err := _VotemarketV1.contract.UnpackLog(event, "BribeClosed", log); err != nil {
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

// ParseBribeClosed is a log parse operation binding the contract event 0x046a3d8b38f161b27b53792783d95179ee019fafaf2aedf362807cb52fc2ab46.
//
// Solidity: event BribeClosed(uint256 id, uint256 remainingReward)
func (_VotemarketV1 *VotemarketV1Filterer) ParseBribeClosed(log types.Log) (*VotemarketV1BribeClosed, error) {
	event := new(VotemarketV1BribeClosed)
	if err := _VotemarketV1.contract.UnpackLog(event, "BribeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1BribeCreatedIterator is returned from FilterBribeCreated and is used to iterate over the raw logs and unpacked data for BribeCreated events raised by the VotemarketV1 contract.
type VotemarketV1BribeCreatedIterator struct {
	Event *VotemarketV1BribeCreated // Event containing the contract specifics and raw log

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
func (it *VotemarketV1BribeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1BribeCreated)
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
		it.Event = new(VotemarketV1BribeCreated)
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
func (it *VotemarketV1BribeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1BribeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1BribeCreated represents a BribeCreated event raised by the VotemarketV1 contract.
type VotemarketV1BribeCreated struct {
	Id                *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	RewardPerPeriod   *big.Int
	TotalRewardAmount *big.Int
	IsUpgradeable     bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBribeCreated is a free log retrieval operation binding the contract event 0x4ed4160f5ef12a0abd9d6134687dd7da5b8274bf240f997a1f377a76c52ccaf5.
//
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address indexed rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV1 *VotemarketV1Filterer) FilterBribeCreated(opts *bind.FilterOpts, id []*big.Int, gauge []common.Address, rewardToken []common.Address) (*VotemarketV1BribeCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "BribeCreated", idRule, gaugeRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1BribeCreatedIterator{contract: _VotemarketV1.contract, event: "BribeCreated", logs: logs, sub: sub}, nil
}

// WatchBribeCreated is a free log subscription operation binding the contract event 0x4ed4160f5ef12a0abd9d6134687dd7da5b8274bf240f997a1f377a76c52ccaf5.
//
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address indexed rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV1 *VotemarketV1Filterer) WatchBribeCreated(opts *bind.WatchOpts, sink chan<- *VotemarketV1BribeCreated, id []*big.Int, gauge []common.Address, rewardToken []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "BribeCreated", idRule, gaugeRule, rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1BribeCreated)
				if err := _VotemarketV1.contract.UnpackLog(event, "BribeCreated", log); err != nil {
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

// ParseBribeCreated is a log parse operation binding the contract event 0x4ed4160f5ef12a0abd9d6134687dd7da5b8274bf240f997a1f377a76c52ccaf5.
//
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address indexed rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV1 *VotemarketV1Filterer) ParseBribeCreated(log types.Log) (*VotemarketV1BribeCreated, error) {
	event := new(VotemarketV1BribeCreated)
	if err := _VotemarketV1.contract.UnpackLog(event, "BribeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1BribeDurationIncreaseQueuedIterator is returned from FilterBribeDurationIncreaseQueued and is used to iterate over the raw logs and unpacked data for BribeDurationIncreaseQueued events raised by the VotemarketV1 contract.
type VotemarketV1BribeDurationIncreaseQueuedIterator struct {
	Event *VotemarketV1BribeDurationIncreaseQueued // Event containing the contract specifics and raw log

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
func (it *VotemarketV1BribeDurationIncreaseQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1BribeDurationIncreaseQueued)
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
		it.Event = new(VotemarketV1BribeDurationIncreaseQueued)
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
func (it *VotemarketV1BribeDurationIncreaseQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1BribeDurationIncreaseQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1BribeDurationIncreaseQueued represents a BribeDurationIncreaseQueued event raised by the VotemarketV1 contract.
type VotemarketV1BribeDurationIncreaseQueued struct {
	Id                *big.Int
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBribeDurationIncreaseQueued is a free log retrieval operation binding the contract event 0x0c841045cbcf87e9cc7521ce9e85cf523d731f87fd5b45feea376cff34067263.
//
// Solidity: event BribeDurationIncreaseQueued(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) FilterBribeDurationIncreaseQueued(opts *bind.FilterOpts) (*VotemarketV1BribeDurationIncreaseQueuedIterator, error) {

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "BribeDurationIncreaseQueued")
	if err != nil {
		return nil, err
	}
	return &VotemarketV1BribeDurationIncreaseQueuedIterator{contract: _VotemarketV1.contract, event: "BribeDurationIncreaseQueued", logs: logs, sub: sub}, nil
}

// WatchBribeDurationIncreaseQueued is a free log subscription operation binding the contract event 0x0c841045cbcf87e9cc7521ce9e85cf523d731f87fd5b45feea376cff34067263.
//
// Solidity: event BribeDurationIncreaseQueued(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) WatchBribeDurationIncreaseQueued(opts *bind.WatchOpts, sink chan<- *VotemarketV1BribeDurationIncreaseQueued) (event.Subscription, error) {

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "BribeDurationIncreaseQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1BribeDurationIncreaseQueued)
				if err := _VotemarketV1.contract.UnpackLog(event, "BribeDurationIncreaseQueued", log); err != nil {
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

// ParseBribeDurationIncreaseQueued is a log parse operation binding the contract event 0x0c841045cbcf87e9cc7521ce9e85cf523d731f87fd5b45feea376cff34067263.
//
// Solidity: event BribeDurationIncreaseQueued(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) ParseBribeDurationIncreaseQueued(log types.Log) (*VotemarketV1BribeDurationIncreaseQueued, error) {
	event := new(VotemarketV1BribeDurationIncreaseQueued)
	if err := _VotemarketV1.contract.UnpackLog(event, "BribeDurationIncreaseQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1BribeDurationIncreasedIterator is returned from FilterBribeDurationIncreased and is used to iterate over the raw logs and unpacked data for BribeDurationIncreased events raised by the VotemarketV1 contract.
type VotemarketV1BribeDurationIncreasedIterator struct {
	Event *VotemarketV1BribeDurationIncreased // Event containing the contract specifics and raw log

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
func (it *VotemarketV1BribeDurationIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1BribeDurationIncreased)
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
		it.Event = new(VotemarketV1BribeDurationIncreased)
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
func (it *VotemarketV1BribeDurationIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1BribeDurationIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1BribeDurationIncreased represents a BribeDurationIncreased event raised by the VotemarketV1 contract.
type VotemarketV1BribeDurationIncreased struct {
	Id                *big.Int
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBribeDurationIncreased is a free log retrieval operation binding the contract event 0xe90b0f7fffa9942eb28c4453083b14e929e4f6d39de19dd5f8cef36148b9c63a.
//
// Solidity: event BribeDurationIncreased(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) FilterBribeDurationIncreased(opts *bind.FilterOpts) (*VotemarketV1BribeDurationIncreasedIterator, error) {

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "BribeDurationIncreased")
	if err != nil {
		return nil, err
	}
	return &VotemarketV1BribeDurationIncreasedIterator{contract: _VotemarketV1.contract, event: "BribeDurationIncreased", logs: logs, sub: sub}, nil
}

// WatchBribeDurationIncreased is a free log subscription operation binding the contract event 0xe90b0f7fffa9942eb28c4453083b14e929e4f6d39de19dd5f8cef36148b9c63a.
//
// Solidity: event BribeDurationIncreased(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) WatchBribeDurationIncreased(opts *bind.WatchOpts, sink chan<- *VotemarketV1BribeDurationIncreased) (event.Subscription, error) {

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "BribeDurationIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1BribeDurationIncreased)
				if err := _VotemarketV1.contract.UnpackLog(event, "BribeDurationIncreased", log); err != nil {
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

// ParseBribeDurationIncreased is a log parse operation binding the contract event 0xe90b0f7fffa9942eb28c4453083b14e929e4f6d39de19dd5f8cef36148b9c63a.
//
// Solidity: event BribeDurationIncreased(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV1 *VotemarketV1Filterer) ParseBribeDurationIncreased(log types.Log) (*VotemarketV1BribeDurationIncreased, error) {
	event := new(VotemarketV1BribeDurationIncreased)
	if err := _VotemarketV1.contract.UnpackLog(event, "BribeDurationIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1ClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the VotemarketV1 contract.
type VotemarketV1ClaimedIterator struct {
	Event *VotemarketV1Claimed // Event containing the contract specifics and raw log

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
func (it *VotemarketV1ClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1Claimed)
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
		it.Event = new(VotemarketV1Claimed)
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
func (it *VotemarketV1ClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1ClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1Claimed represents a Claimed event raised by the VotemarketV1 contract.
type VotemarketV1Claimed struct {
	User         common.Address
	RewardToken  common.Address
	BribeId      *big.Int
	Amount       *big.Int
	ProtocolFees *big.Int
	Period       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76.
//
// Solidity: event Claimed(address indexed user, address indexed rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV1 *VotemarketV1Filterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address, bribeId []*big.Int) (*VotemarketV1ClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var bribeIdRule []interface{}
	for _, bribeIdItem := range bribeId {
		bribeIdRule = append(bribeIdRule, bribeIdItem)
	}

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "Claimed", userRule, rewardTokenRule, bribeIdRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1ClaimedIterator{contract: _VotemarketV1.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76.
//
// Solidity: event Claimed(address indexed user, address indexed rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV1 *VotemarketV1Filterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *VotemarketV1Claimed, user []common.Address, rewardToken []common.Address, bribeId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var bribeIdRule []interface{}
	for _, bribeIdItem := range bribeId {
		bribeIdRule = append(bribeIdRule, bribeIdItem)
	}

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "Claimed", userRule, rewardTokenRule, bribeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1Claimed)
				if err := _VotemarketV1.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76.
//
// Solidity: event Claimed(address indexed user, address indexed rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV1 *VotemarketV1Filterer) ParseClaimed(log types.Log) (*VotemarketV1Claimed, error) {
	event := new(VotemarketV1Claimed)
	if err := _VotemarketV1.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1ManagerUpdatedIterator is returned from FilterManagerUpdated and is used to iterate over the raw logs and unpacked data for ManagerUpdated events raised by the VotemarketV1 contract.
type VotemarketV1ManagerUpdatedIterator struct {
	Event *VotemarketV1ManagerUpdated // Event containing the contract specifics and raw log

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
func (it *VotemarketV1ManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1ManagerUpdated)
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
		it.Event = new(VotemarketV1ManagerUpdated)
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
func (it *VotemarketV1ManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1ManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1ManagerUpdated represents a ManagerUpdated event raised by the VotemarketV1 contract.
type VotemarketV1ManagerUpdated struct {
	Id      *big.Int
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerUpdated is a free log retrieval operation binding the contract event 0x5baaf19ee8739d1720c1401554ff4abe5682505ec43d3f2eb61b9dfc0abd9745.
//
// Solidity: event ManagerUpdated(uint256 id, address indexed manager)
func (_VotemarketV1 *VotemarketV1Filterer) FilterManagerUpdated(opts *bind.FilterOpts, manager []common.Address) (*VotemarketV1ManagerUpdatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV1ManagerUpdatedIterator{contract: _VotemarketV1.contract, event: "ManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchManagerUpdated is a free log subscription operation binding the contract event 0x5baaf19ee8739d1720c1401554ff4abe5682505ec43d3f2eb61b9dfc0abd9745.
//
// Solidity: event ManagerUpdated(uint256 id, address indexed manager)
func (_VotemarketV1 *VotemarketV1Filterer) WatchManagerUpdated(opts *bind.WatchOpts, sink chan<- *VotemarketV1ManagerUpdated, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1ManagerUpdated)
				if err := _VotemarketV1.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
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

// ParseManagerUpdated is a log parse operation binding the contract event 0x5baaf19ee8739d1720c1401554ff4abe5682505ec43d3f2eb61b9dfc0abd9745.
//
// Solidity: event ManagerUpdated(uint256 id, address indexed manager)
func (_VotemarketV1 *VotemarketV1Filterer) ParseManagerUpdated(log types.Log) (*VotemarketV1ManagerUpdated, error) {
	event := new(VotemarketV1ManagerUpdated)
	if err := _VotemarketV1.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV1PeriodRolledOverIterator is returned from FilterPeriodRolledOver and is used to iterate over the raw logs and unpacked data for PeriodRolledOver events raised by the VotemarketV1 contract.
type VotemarketV1PeriodRolledOverIterator struct {
	Event *VotemarketV1PeriodRolledOver // Event containing the contract specifics and raw log

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
func (it *VotemarketV1PeriodRolledOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV1PeriodRolledOver)
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
		it.Event = new(VotemarketV1PeriodRolledOver)
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
func (it *VotemarketV1PeriodRolledOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV1PeriodRolledOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV1PeriodRolledOver represents a PeriodRolledOver event raised by the VotemarketV1 contract.
type VotemarketV1PeriodRolledOver struct {
	Id              *big.Int
	PeriodId        *big.Int
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPeriodRolledOver is a free log retrieval operation binding the contract event 0xb77c22cd311809931524bcc8d4a33a61a392e2304c8a7d476b64018e11ed6cb1.
//
// Solidity: event PeriodRolledOver(uint256 id, uint256 periodId, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1Filterer) FilterPeriodRolledOver(opts *bind.FilterOpts) (*VotemarketV1PeriodRolledOverIterator, error) {

	logs, sub, err := _VotemarketV1.contract.FilterLogs(opts, "PeriodRolledOver")
	if err != nil {
		return nil, err
	}
	return &VotemarketV1PeriodRolledOverIterator{contract: _VotemarketV1.contract, event: "PeriodRolledOver", logs: logs, sub: sub}, nil
}

// WatchPeriodRolledOver is a free log subscription operation binding the contract event 0xb77c22cd311809931524bcc8d4a33a61a392e2304c8a7d476b64018e11ed6cb1.
//
// Solidity: event PeriodRolledOver(uint256 id, uint256 periodId, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1Filterer) WatchPeriodRolledOver(opts *bind.WatchOpts, sink chan<- *VotemarketV1PeriodRolledOver) (event.Subscription, error) {

	logs, sub, err := _VotemarketV1.contract.WatchLogs(opts, "PeriodRolledOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV1PeriodRolledOver)
				if err := _VotemarketV1.contract.UnpackLog(event, "PeriodRolledOver", log); err != nil {
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

// ParsePeriodRolledOver is a log parse operation binding the contract event 0xb77c22cd311809931524bcc8d4a33a61a392e2304c8a7d476b64018e11ed6cb1.
//
// Solidity: event PeriodRolledOver(uint256 id, uint256 periodId, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV1 *VotemarketV1Filterer) ParsePeriodRolledOver(log types.Log) (*VotemarketV1PeriodRolledOver, error) {
	event := new(VotemarketV1PeriodRolledOver)
	if err := _VotemarketV1.contract.UnpackLog(event, "PeriodRolledOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
