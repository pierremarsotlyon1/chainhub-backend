// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votemarketV2

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

// VotemarketV2MetaData contains all meta data concerning the VotemarketV2 contract.
var VotemarketV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gaugeController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER_ONLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_NUMBER_OF_PERIODS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KILLED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_UPGRADEABLE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_PERIODS_LEFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WRONG_INPUT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingReward\",\"type\":\"uint256\"}],\"name\":\"BribeClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isUpgradeable\",\"type\":\"bool\"}],\"name\":\"BribeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"name\":\"BribeDurationIncreaseQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"name\":\"BribeDurationIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"PeriodRolledOver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RecipientSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMUM_PERIOD\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activePeriod\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"amountClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bribes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"claimAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"claimAllFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardTokens\",\"type\":\"address[]\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claimFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"closeBribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"upgradeable\",\"type\":\"bool\"}],\"name\":\"createBribe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBribeID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gaugeController\",\"outputs\":[{\"internalType\":\"contractGaugeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getActivePeriod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structPlatform.Period\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getActivePeriodPerBribe\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getBlacklistedAddressesForBribe\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getBribe\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"}],\"internalType\":\"structPlatform.Bribe\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getPeriodsLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodsLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"getUpgradedBribeQueued\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blacklist\",\"type\":\"address[]\"}],\"internalType\":\"structPlatform.Upgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bribeId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_additionnalPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_increasedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newMaxPricePerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addressesBlacklisted\",\"type\":\"address[]\"}],\"name\":\"increaseBribeDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isKilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isUpgradeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastUserClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardPerVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipientFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"}],\"name\":\"updateBribePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"updateBribePeriods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"updateManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"upgradeBribeQueue\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotemarketV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VotemarketV2MetaData.ABI instead.
var VotemarketV2ABI = VotemarketV2MetaData.ABI

// VotemarketV2 is an auto generated Go binding around an Ethereum contract.
type VotemarketV2 struct {
	VotemarketV2Caller     // Read-only binding to the contract
	VotemarketV2Transactor // Write-only binding to the contract
	VotemarketV2Filterer   // Log filterer for contract events
}

// VotemarketV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VotemarketV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VotemarketV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotemarketV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemarketV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotemarketV2Session struct {
	Contract     *VotemarketV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotemarketV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotemarketV2CallerSession struct {
	Contract *VotemarketV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VotemarketV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotemarketV2TransactorSession struct {
	Contract     *VotemarketV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VotemarketV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VotemarketV2Raw struct {
	Contract *VotemarketV2 // Generic contract binding to access the raw methods on
}

// VotemarketV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotemarketV2CallerRaw struct {
	Contract *VotemarketV2Caller // Generic read-only contract binding to access the raw methods on
}

// VotemarketV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotemarketV2TransactorRaw struct {
	Contract *VotemarketV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVotemarketV2 creates a new instance of VotemarketV2, bound to a specific deployed contract.
func NewVotemarketV2(address common.Address, backend bind.ContractBackend) (*VotemarketV2, error) {
	contract, err := bindVotemarketV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2{VotemarketV2Caller: VotemarketV2Caller{contract: contract}, VotemarketV2Transactor: VotemarketV2Transactor{contract: contract}, VotemarketV2Filterer: VotemarketV2Filterer{contract: contract}}, nil
}

// NewVotemarketV2Caller creates a new read-only instance of VotemarketV2, bound to a specific deployed contract.
func NewVotemarketV2Caller(address common.Address, caller bind.ContractCaller) (*VotemarketV2Caller, error) {
	contract, err := bindVotemarketV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2Caller{contract: contract}, nil
}

// NewVotemarketV2Transactor creates a new write-only instance of VotemarketV2, bound to a specific deployed contract.
func NewVotemarketV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VotemarketV2Transactor, error) {
	contract, err := bindVotemarketV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2Transactor{contract: contract}, nil
}

// NewVotemarketV2Filterer creates a new log filterer instance of VotemarketV2, bound to a specific deployed contract.
func NewVotemarketV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VotemarketV2Filterer, error) {
	contract, err := bindVotemarketV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2Filterer{contract: contract}, nil
}

// bindVotemarketV2 binds a generic wrapper to an already deployed contract.
func bindVotemarketV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotemarketV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotemarketV2 *VotemarketV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotemarketV2.Contract.VotemarketV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotemarketV2 *VotemarketV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV2.Contract.VotemarketV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotemarketV2 *VotemarketV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotemarketV2.Contract.VotemarketV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotemarketV2 *VotemarketV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotemarketV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotemarketV2 *VotemarketV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotemarketV2 *VotemarketV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotemarketV2.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV2 *VotemarketV2Caller) MINIMUMPERIOD(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "MINIMUM_PERIOD")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV2 *VotemarketV2Session) MINIMUMPERIOD() (uint8, error) {
	return _VotemarketV2.Contract.MINIMUMPERIOD(&_VotemarketV2.CallOpts)
}

// MINIMUMPERIOD is a free data retrieval call binding the contract method 0x51cd41e8.
//
// Solidity: function MINIMUM_PERIOD() view returns(uint8)
func (_VotemarketV2 *VotemarketV2CallerSession) MINIMUMPERIOD() (uint8, error) {
	return _VotemarketV2.Contract.MINIMUMPERIOD(&_VotemarketV2.CallOpts)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x21bf936a.
//
// Solidity: function activePeriod(uint256 ) view returns(uint8 id, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV2 *VotemarketV2Caller) ActivePeriod(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "activePeriod", arg0)

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
func (_VotemarketV2 *VotemarketV2Session) ActivePeriod(arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	return _VotemarketV2.Contract.ActivePeriod(&_VotemarketV2.CallOpts, arg0)
}

// ActivePeriod is a free data retrieval call binding the contract method 0x21bf936a.
//
// Solidity: function activePeriod(uint256 ) view returns(uint8 id, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV2 *VotemarketV2CallerSession) ActivePeriod(arg0 *big.Int) (struct {
	Id              uint8
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
}, error) {
	return _VotemarketV2.Contract.ActivePeriod(&_VotemarketV2.CallOpts, arg0)
}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) AmountClaimed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "amountClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) AmountClaimed(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.AmountClaimed(&_VotemarketV2.CallOpts, arg0)
}

// AmountClaimed is a free data retrieval call binding the contract method 0xd1d1bb4f.
//
// Solidity: function amountClaimed(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) AmountClaimed(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.AmountClaimed(&_VotemarketV2.CallOpts, arg0)
}

// Bribes is a free data retrieval call binding the contract method 0xa9d46e5b.
//
// Solidity: function bribes(uint256 ) view returns(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 endTimestamp, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VotemarketV2 *VotemarketV2Caller) Bribes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "bribes", arg0)

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
func (_VotemarketV2 *VotemarketV2Session) Bribes(arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	return _VotemarketV2.Contract.Bribes(&_VotemarketV2.CallOpts, arg0)
}

// Bribes is a free data retrieval call binding the contract method 0xa9d46e5b.
//
// Solidity: function bribes(uint256 ) view returns(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 endTimestamp, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VotemarketV2 *VotemarketV2CallerSession) Bribes(arg0 *big.Int) (struct {
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	EndTimestamp      *big.Int
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
}, error) {
	return _VotemarketV2.Contract.Bribes(&_VotemarketV2.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV2 *VotemarketV2Caller) Claimable(opts *bind.CallOpts, user common.Address, bribeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "claimable", user, bribeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV2 *VotemarketV2Session) Claimable(user common.Address, bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.Claimable(&_VotemarketV2.CallOpts, user, bribeId)
}

// Claimable is a free data retrieval call binding the contract method 0x60efe334.
//
// Solidity: function claimable(address user, uint256 bribeId) view returns(uint256 amount)
func (_VotemarketV2 *VotemarketV2CallerSession) Claimable(user common.Address, bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.Claimable(&_VotemarketV2.CallOpts, user, bribeId)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) Fee() (*big.Int, error) {
	return _VotemarketV2.Contract.Fee(&_VotemarketV2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) Fee() (*big.Int, error) {
	return _VotemarketV2.Contract.Fee(&_VotemarketV2.CallOpts)
}

// FeeAccrued is a free data retrieval call binding the contract method 0xc7da3328.
//
// Solidity: function feeAccrued(address ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) FeeAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "feeAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAccrued is a free data retrieval call binding the contract method 0xc7da3328.
//
// Solidity: function feeAccrued(address ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) FeeAccrued(arg0 common.Address) (*big.Int, error) {
	return _VotemarketV2.Contract.FeeAccrued(&_VotemarketV2.CallOpts, arg0)
}

// FeeAccrued is a free data retrieval call binding the contract method 0xc7da3328.
//
// Solidity: function feeAccrued(address ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) FeeAccrued(arg0 common.Address) (*big.Int, error) {
	return _VotemarketV2.Contract.FeeAccrued(&_VotemarketV2.CallOpts, arg0)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VotemarketV2 *VotemarketV2Caller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VotemarketV2 *VotemarketV2Session) FeeCollector() (common.Address, error) {
	return _VotemarketV2.Contract.FeeCollector(&_VotemarketV2.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VotemarketV2 *VotemarketV2CallerSession) FeeCollector() (common.Address, error) {
	return _VotemarketV2.Contract.FeeCollector(&_VotemarketV2.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV2 *VotemarketV2Caller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "gaugeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV2 *VotemarketV2Session) GaugeController() (common.Address, error) {
	return _VotemarketV2.Contract.GaugeController(&_VotemarketV2.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0x99eecb3b.
//
// Solidity: function gaugeController() view returns(address)
func (_VotemarketV2 *VotemarketV2CallerSession) GaugeController() (common.Address, error) {
	return _VotemarketV2.Contract.GaugeController(&_VotemarketV2.CallOpts)
}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV2 *VotemarketV2Caller) GetActivePeriod(opts *bind.CallOpts, bribeId *big.Int) (PlatformPeriod, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getActivePeriod", bribeId)

	if err != nil {
		return *new(PlatformPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformPeriod)).(*PlatformPeriod)

	return out0, err

}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV2 *VotemarketV2Session) GetActivePeriod(bribeId *big.Int) (PlatformPeriod, error) {
	return _VotemarketV2.Contract.GetActivePeriod(&_VotemarketV2.CallOpts, bribeId)
}

// GetActivePeriod is a free data retrieval call binding the contract method 0x228c076c.
//
// Solidity: function getActivePeriod(uint256 bribeId) view returns((uint8,uint256,uint256))
func (_VotemarketV2 *VotemarketV2CallerSession) GetActivePeriod(bribeId *big.Int) (PlatformPeriod, error) {
	return _VotemarketV2.Contract.GetActivePeriod(&_VotemarketV2.CallOpts, bribeId)
}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV2 *VotemarketV2Caller) GetActivePeriodPerBribe(opts *bind.CallOpts, bribeId *big.Int) (uint8, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getActivePeriodPerBribe", bribeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV2 *VotemarketV2Session) GetActivePeriodPerBribe(bribeId *big.Int) (uint8, error) {
	return _VotemarketV2.Contract.GetActivePeriodPerBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetActivePeriodPerBribe is a free data retrieval call binding the contract method 0x7d8e3769.
//
// Solidity: function getActivePeriodPerBribe(uint256 bribeId) view returns(uint8)
func (_VotemarketV2 *VotemarketV2CallerSession) GetActivePeriodPerBribe(bribeId *big.Int) (uint8, error) {
	return _VotemarketV2.Contract.GetActivePeriodPerBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV2 *VotemarketV2Caller) GetBlacklistedAddressesForBribe(opts *bind.CallOpts, bribeId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getBlacklistedAddressesForBribe", bribeId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV2 *VotemarketV2Session) GetBlacklistedAddressesForBribe(bribeId *big.Int) ([]common.Address, error) {
	return _VotemarketV2.Contract.GetBlacklistedAddressesForBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetBlacklistedAddressesForBribe is a free data retrieval call binding the contract method 0x54d97ed5.
//
// Solidity: function getBlacklistedAddressesForBribe(uint256 bribeId) view returns(address[])
func (_VotemarketV2 *VotemarketV2CallerSession) GetBlacklistedAddressesForBribe(bribeId *big.Int) ([]common.Address, error) {
	return _VotemarketV2.Contract.GetBlacklistedAddressesForBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2Caller) GetBribe(opts *bind.CallOpts, bribeId *big.Int) (PlatformBribe, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getBribe", bribeId)

	if err != nil {
		return *new(PlatformBribe), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformBribe)).(*PlatformBribe)

	return out0, err

}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2Session) GetBribe(bribeId *big.Int) (PlatformBribe, error) {
	return _VotemarketV2.Contract.GetBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetBribe is a free data retrieval call binding the contract method 0x06b6165f.
//
// Solidity: function getBribe(uint256 bribeId) view returns((address,address,address,uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2CallerSession) GetBribe(bribeId *big.Int) (PlatformBribe, error) {
	return _VotemarketV2.Contract.GetBribe(&_VotemarketV2.CallOpts, bribeId)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) GetCurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getCurrentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) GetCurrentPeriod() (*big.Int, error) {
	return _VotemarketV2.Contract.GetCurrentPeriod(&_VotemarketV2.CallOpts)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) GetCurrentPeriod() (*big.Int, error) {
	return _VotemarketV2.Contract.GetCurrentPeriod(&_VotemarketV2.CallOpts)
}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV2 *VotemarketV2Caller) GetPeriodsLeft(opts *bind.CallOpts, bribeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getPeriodsLeft", bribeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV2 *VotemarketV2Session) GetPeriodsLeft(bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.GetPeriodsLeft(&_VotemarketV2.CallOpts, bribeId)
}

// GetPeriodsLeft is a free data retrieval call binding the contract method 0x6b5646aa.
//
// Solidity: function getPeriodsLeft(uint256 bribeId) view returns(uint256 periodsLeft)
func (_VotemarketV2 *VotemarketV2CallerSession) GetPeriodsLeft(bribeId *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.GetPeriodsLeft(&_VotemarketV2.CallOpts, bribeId)
}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2Caller) GetUpgradedBribeQueued(opts *bind.CallOpts, bribeId *big.Int) (PlatformUpgrade, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "getUpgradedBribeQueued", bribeId)

	if err != nil {
		return *new(PlatformUpgrade), err
	}

	out0 := *abi.ConvertType(out[0], new(PlatformUpgrade)).(*PlatformUpgrade)

	return out0, err

}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2Session) GetUpgradedBribeQueued(bribeId *big.Int) (PlatformUpgrade, error) {
	return _VotemarketV2.Contract.GetUpgradedBribeQueued(&_VotemarketV2.CallOpts, bribeId)
}

// GetUpgradedBribeQueued is a free data retrieval call binding the contract method 0xc29b6677.
//
// Solidity: function getUpgradedBribeQueued(uint256 bribeId) view returns((uint8,uint256,uint256,uint256,address[]))
func (_VotemarketV2 *VotemarketV2CallerSession) GetUpgradedBribeQueued(bribeId *big.Int) (PlatformUpgrade, error) {
	return _VotemarketV2.Contract.GetUpgradedBribeQueued(&_VotemarketV2.CallOpts, bribeId)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV2 *VotemarketV2Caller) IsBlacklisted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "isBlacklisted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV2 *VotemarketV2Session) IsBlacklisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotemarketV2.Contract.IsBlacklisted(&_VotemarketV2.CallOpts, arg0, arg1)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x4ae001d8.
//
// Solidity: function isBlacklisted(uint256 , address ) view returns(bool)
func (_VotemarketV2 *VotemarketV2CallerSession) IsBlacklisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotemarketV2.Contract.IsBlacklisted(&_VotemarketV2.CallOpts, arg0, arg1)
}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV2 *VotemarketV2Caller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "isKilled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV2 *VotemarketV2Session) IsKilled() (bool, error) {
	return _VotemarketV2.Contract.IsKilled(&_VotemarketV2.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x8fe8a101.
//
// Solidity: function isKilled() view returns(bool)
func (_VotemarketV2 *VotemarketV2CallerSession) IsKilled() (bool, error) {
	return _VotemarketV2.Contract.IsKilled(&_VotemarketV2.CallOpts)
}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV2 *VotemarketV2Caller) IsUpgradeable(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "isUpgradeable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV2 *VotemarketV2Session) IsUpgradeable(arg0 *big.Int) (bool, error) {
	return _VotemarketV2.Contract.IsUpgradeable(&_VotemarketV2.CallOpts, arg0)
}

// IsUpgradeable is a free data retrieval call binding the contract method 0xaca47b7d.
//
// Solidity: function isUpgradeable(uint256 ) view returns(bool)
func (_VotemarketV2 *VotemarketV2CallerSession) IsUpgradeable(arg0 *big.Int) (bool, error) {
	return _VotemarketV2.Contract.IsUpgradeable(&_VotemarketV2.CallOpts, arg0)
}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) LastUserClaim(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "lastUserClaim", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) LastUserClaim(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.LastUserClaim(&_VotemarketV2.CallOpts, arg0, arg1)
}

// LastUserClaim is a free data retrieval call binding the contract method 0x4fcf04f9.
//
// Solidity: function lastUserClaim(address , uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) LastUserClaim(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.LastUserClaim(&_VotemarketV2.CallOpts, arg0, arg1)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) NextID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "nextID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) NextID() (*big.Int, error) {
	return _VotemarketV2.Contract.NextID(&_VotemarketV2.CallOpts)
}

// NextID is a free data retrieval call binding the contract method 0x1e96917d.
//
// Solidity: function nextID() view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) NextID() (*big.Int, error) {
	return _VotemarketV2.Contract.NextID(&_VotemarketV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotemarketV2 *VotemarketV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotemarketV2 *VotemarketV2Session) Owner() (common.Address, error) {
	return _VotemarketV2.Contract.Owner(&_VotemarketV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotemarketV2 *VotemarketV2CallerSession) Owner() (common.Address, error) {
	return _VotemarketV2.Contract.Owner(&_VotemarketV2.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0xb3651eea.
//
// Solidity: function recipient(address ) view returns(address)
func (_VotemarketV2 *VotemarketV2Caller) Recipient(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "recipient", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0xb3651eea.
//
// Solidity: function recipient(address ) view returns(address)
func (_VotemarketV2 *VotemarketV2Session) Recipient(arg0 common.Address) (common.Address, error) {
	return _VotemarketV2.Contract.Recipient(&_VotemarketV2.CallOpts, arg0)
}

// Recipient is a free data retrieval call binding the contract method 0xb3651eea.
//
// Solidity: function recipient(address ) view returns(address)
func (_VotemarketV2 *VotemarketV2CallerSession) Recipient(arg0 common.Address) (common.Address, error) {
	return _VotemarketV2.Contract.Recipient(&_VotemarketV2.CallOpts, arg0)
}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Caller) RewardPerVote(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "rewardPerVote", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) RewardPerVote(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.RewardPerVote(&_VotemarketV2.CallOpts, arg0)
}

// RewardPerVote is a free data retrieval call binding the contract method 0x3ad86d72.
//
// Solidity: function rewardPerVote(uint256 ) view returns(uint256)
func (_VotemarketV2 *VotemarketV2CallerSession) RewardPerVote(arg0 *big.Int) (*big.Int, error) {
	return _VotemarketV2.Contract.RewardPerVote(&_VotemarketV2.CallOpts, arg0)
}

// UpgradeBribeQueue is a free data retrieval call binding the contract method 0x83287273.
//
// Solidity: function upgradeBribeQueue(uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VotemarketV2 *VotemarketV2Caller) UpgradeBribeQueue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	var out []interface{}
	err := _VotemarketV2.contract.Call(opts, &out, "upgradeBribeQueue", arg0)

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
func (_VotemarketV2 *VotemarketV2Session) UpgradeBribeQueue(arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VotemarketV2.Contract.UpgradeBribeQueue(&_VotemarketV2.CallOpts, arg0)
}

// UpgradeBribeQueue is a free data retrieval call binding the contract method 0x83287273.
//
// Solidity: function upgradeBribeQueue(uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VotemarketV2 *VotemarketV2CallerSession) UpgradeBribeQueue(arg0 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VotemarketV2.Contract.UpgradeBribeQueue(&_VotemarketV2.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2Transactor) Claim(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claim", bribeId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) Claim(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.Claim(&_VotemarketV2.TransactOpts, bribeId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2TransactorSession) Claim(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.Claim(&_VotemarketV2.TransactOpts, bribeId)
}

// Claim0 is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 bribeId, address _recipient) returns(uint256)
func (_VotemarketV2 *VotemarketV2Transactor) Claim0(opts *bind.TransactOpts, bribeId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claim0", bribeId, _recipient)
}

// Claim0 is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 bribeId, address _recipient) returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) Claim0(bribeId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.Claim0(&_VotemarketV2.TransactOpts, bribeId, _recipient)
}

// Claim0 is a paid mutator transaction binding the contract method 0xddd5e1b2.
//
// Solidity: function claim(uint256 bribeId, address _recipient) returns(uint256)
func (_VotemarketV2 *VotemarketV2TransactorSession) Claim0(bribeId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.Claim0(&_VotemarketV2.TransactOpts, bribeId, _recipient)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Transactor) ClaimAll(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claimAll", ids)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Session) ClaimAll(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAll(&_VotemarketV2.TransactOpts, ids)
}

// ClaimAll is a paid mutator transaction binding the contract method 0x28c77820.
//
// Solidity: function claimAll(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) ClaimAll(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAll(&_VotemarketV2.TransactOpts, ids)
}

// ClaimAll0 is a paid mutator transaction binding the contract method 0xdd47a77f.
//
// Solidity: function claimAll(uint256[] ids, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Transactor) ClaimAll0(opts *bind.TransactOpts, ids []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claimAll0", ids, _recipient)
}

// ClaimAll0 is a paid mutator transaction binding the contract method 0xdd47a77f.
//
// Solidity: function claimAll(uint256[] ids, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Session) ClaimAll0(ids []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAll0(&_VotemarketV2.TransactOpts, ids, _recipient)
}

// ClaimAll0 is a paid mutator transaction binding the contract method 0xdd47a77f.
//
// Solidity: function claimAll(uint256[] ids, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) ClaimAll0(ids []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAll0(&_VotemarketV2.TransactOpts, ids, _recipient)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address _user, uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Transactor) ClaimAllFor(opts *bind.TransactOpts, _user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claimAllFor", _user, ids)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address _user, uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Session) ClaimAllFor(_user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAllFor(&_VotemarketV2.TransactOpts, _user, ids)
}

// ClaimAllFor is a paid mutator transaction binding the contract method 0xde4aaaf4.
//
// Solidity: function claimAllFor(address _user, uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) ClaimAllFor(_user common.Address, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimAllFor(&_VotemarketV2.TransactOpts, _user, ids)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x725acc6b.
//
// Solidity: function claimFees(address[] rewardTokens) returns()
func (_VotemarketV2 *VotemarketV2Transactor) ClaimFees(opts *bind.TransactOpts, rewardTokens []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claimFees", rewardTokens)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x725acc6b.
//
// Solidity: function claimFees(address[] rewardTokens) returns()
func (_VotemarketV2 *VotemarketV2Session) ClaimFees(rewardTokens []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimFees(&_VotemarketV2.TransactOpts, rewardTokens)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x725acc6b.
//
// Solidity: function claimFees(address[] rewardTokens) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) ClaimFees(rewardTokens []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimFees(&_VotemarketV2.TransactOpts, rewardTokens)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2Transactor) ClaimFor(opts *bind.TransactOpts, user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "claimFor", user, bribeId)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2Session) ClaimFor(user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimFor(&_VotemarketV2.TransactOpts, user, bribeId)
}

// ClaimFor is a paid mutator transaction binding the contract method 0x0de05659.
//
// Solidity: function claimFor(address user, uint256 bribeId) returns(uint256)
func (_VotemarketV2 *VotemarketV2TransactorSession) ClaimFor(user common.Address, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.ClaimFor(&_VotemarketV2.TransactOpts, user, bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2Transactor) CloseBribe(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "closeBribe", bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2Session) CloseBribe(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.CloseBribe(&_VotemarketV2.TransactOpts, bribeId)
}

// CloseBribe is a paid mutator transaction binding the contract method 0x28718374.
//
// Solidity: function closeBribe(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) CloseBribe(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.CloseBribe(&_VotemarketV2.TransactOpts, bribeId)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV2 *VotemarketV2Transactor) CreateBribe(opts *bind.TransactOpts, gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "createBribe", gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV2 *VotemarketV2Session) CreateBribe(gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV2.Contract.CreateBribe(&_VotemarketV2.TransactOpts, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// CreateBribe is a paid mutator transaction binding the contract method 0x60debfd3.
//
// Solidity: function createBribe(address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] blacklist, bool upgradeable) returns(uint256 newBribeID)
func (_VotemarketV2 *VotemarketV2TransactorSession) CreateBribe(gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, blacklist []common.Address, upgradeable bool) (*types.Transaction, error) {
	return _VotemarketV2.Contract.CreateBribe(&_VotemarketV2.TransactOpts, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, blacklist, upgradeable)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV2 *VotemarketV2Transactor) IncreaseBribeDuration(opts *bind.TransactOpts, _bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "increaseBribeDuration", _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV2 *VotemarketV2Session) IncreaseBribeDuration(_bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.IncreaseBribeDuration(&_VotemarketV2.TransactOpts, _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// IncreaseBribeDuration is a paid mutator transaction binding the contract method 0xf578c45b.
//
// Solidity: function increaseBribeDuration(uint256 _bribeId, uint8 _additionnalPeriods, uint256 _increasedAmount, uint256 _newMaxPricePerVote, address[] _addressesBlacklisted) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) IncreaseBribeDuration(_bribeId *big.Int, _additionnalPeriods uint8, _increasedAmount *big.Int, _newMaxPricePerVote *big.Int, _addressesBlacklisted []common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.IncreaseBribeDuration(&_VotemarketV2.TransactOpts, _bribeId, _additionnalPeriods, _increasedAmount, _newMaxPricePerVote, _addressesBlacklisted)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV2 *VotemarketV2Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV2 *VotemarketV2Session) Kill() (*types.Transaction, error) {
	return _VotemarketV2.Contract.Kill(&_VotemarketV2.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) Kill() (*types.Transaction, error) {
	return _VotemarketV2.Contract.Kill(&_VotemarketV2.TransactOpts)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VotemarketV2 *VotemarketV2Transactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VotemarketV2 *VotemarketV2Session) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetFeeCollector(&_VotemarketV2.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetFeeCollector(&_VotemarketV2.TransactOpts, _feeCollector)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _platformFee) returns()
func (_VotemarketV2 *VotemarketV2Transactor) SetPlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "setPlatformFee", _platformFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _platformFee) returns()
func (_VotemarketV2 *VotemarketV2Session) SetPlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetPlatformFee(&_VotemarketV2.TransactOpts, _platformFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _platformFee) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) SetPlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetPlatformFee(&_VotemarketV2.TransactOpts, _platformFee)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Transactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Session) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetRecipient(&_VotemarketV2.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetRecipient(&_VotemarketV2.TransactOpts, _recipient)
}

// SetRecipientFor is a paid mutator transaction binding the contract method 0xb6b6e3c8.
//
// Solidity: function setRecipientFor(address _for, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Transactor) SetRecipientFor(opts *bind.TransactOpts, _for common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "setRecipientFor", _for, _recipient)
}

// SetRecipientFor is a paid mutator transaction binding the contract method 0xb6b6e3c8.
//
// Solidity: function setRecipientFor(address _for, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2Session) SetRecipientFor(_for common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetRecipientFor(&_VotemarketV2.TransactOpts, _for, _recipient)
}

// SetRecipientFor is a paid mutator transaction binding the contract method 0xb6b6e3c8.
//
// Solidity: function setRecipientFor(address _for, address _recipient) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) SetRecipientFor(_for common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.SetRecipientFor(&_VotemarketV2.TransactOpts, _for, _recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotemarketV2 *VotemarketV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotemarketV2 *VotemarketV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.TransferOwnership(&_VotemarketV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.TransferOwnership(&_VotemarketV2.TransactOpts, newOwner)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2Transactor) UpdateBribePeriod(opts *bind.TransactOpts, bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "updateBribePeriod", bribeId)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2Session) UpdateBribePeriod(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateBribePeriod(&_VotemarketV2.TransactOpts, bribeId)
}

// UpdateBribePeriod is a paid mutator transaction binding the contract method 0xf9967382.
//
// Solidity: function updateBribePeriod(uint256 bribeId) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) UpdateBribePeriod(bribeId *big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateBribePeriod(&_VotemarketV2.TransactOpts, bribeId)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Transactor) UpdateBribePeriods(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "updateBribePeriods", ids)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2Session) UpdateBribePeriods(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateBribePeriods(&_VotemarketV2.TransactOpts, ids)
}

// UpdateBribePeriods is a paid mutator transaction binding the contract method 0xf8c07f2f.
//
// Solidity: function updateBribePeriods(uint256[] ids) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) UpdateBribePeriods(ids []*big.Int) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateBribePeriods(&_VotemarketV2.TransactOpts, ids)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV2 *VotemarketV2Transactor) UpdateManager(opts *bind.TransactOpts, bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV2.contract.Transact(opts, "updateManager", bribeId, newManager)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV2 *VotemarketV2Session) UpdateManager(bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateManager(&_VotemarketV2.TransactOpts, bribeId, newManager)
}

// UpdateManager is a paid mutator transaction binding the contract method 0xef2c4082.
//
// Solidity: function updateManager(uint256 bribeId, address newManager) returns()
func (_VotemarketV2 *VotemarketV2TransactorSession) UpdateManager(bribeId *big.Int, newManager common.Address) (*types.Transaction, error) {
	return _VotemarketV2.Contract.UpdateManager(&_VotemarketV2.TransactOpts, bribeId, newManager)
}

// VotemarketV2BribeClosedIterator is returned from FilterBribeClosed and is used to iterate over the raw logs and unpacked data for BribeClosed events raised by the VotemarketV2 contract.
type VotemarketV2BribeClosedIterator struct {
	Event *VotemarketV2BribeClosed // Event containing the contract specifics and raw log

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
func (it *VotemarketV2BribeClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2BribeClosed)
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
		it.Event = new(VotemarketV2BribeClosed)
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
func (it *VotemarketV2BribeClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2BribeClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2BribeClosed represents a BribeClosed event raised by the VotemarketV2 contract.
type VotemarketV2BribeClosed struct {
	Id              *big.Int
	RemainingReward *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBribeClosed is a free log retrieval operation binding the contract event 0x046a3d8b38f161b27b53792783d95179ee019fafaf2aedf362807cb52fc2ab46.
//
// Solidity: event BribeClosed(uint256 id, uint256 remainingReward)
func (_VotemarketV2 *VotemarketV2Filterer) FilterBribeClosed(opts *bind.FilterOpts) (*VotemarketV2BribeClosedIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "BribeClosed")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2BribeClosedIterator{contract: _VotemarketV2.contract, event: "BribeClosed", logs: logs, sub: sub}, nil
}

// WatchBribeClosed is a free log subscription operation binding the contract event 0x046a3d8b38f161b27b53792783d95179ee019fafaf2aedf362807cb52fc2ab46.
//
// Solidity: event BribeClosed(uint256 id, uint256 remainingReward)
func (_VotemarketV2 *VotemarketV2Filterer) WatchBribeClosed(opts *bind.WatchOpts, sink chan<- *VotemarketV2BribeClosed) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "BribeClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2BribeClosed)
				if err := _VotemarketV2.contract.UnpackLog(event, "BribeClosed", log); err != nil {
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
func (_VotemarketV2 *VotemarketV2Filterer) ParseBribeClosed(log types.Log) (*VotemarketV2BribeClosed, error) {
	event := new(VotemarketV2BribeClosed)
	if err := _VotemarketV2.contract.UnpackLog(event, "BribeClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2BribeCreatedIterator is returned from FilterBribeCreated and is used to iterate over the raw logs and unpacked data for BribeCreated events raised by the VotemarketV2 contract.
type VotemarketV2BribeCreatedIterator struct {
	Event *VotemarketV2BribeCreated // Event containing the contract specifics and raw log

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
func (it *VotemarketV2BribeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2BribeCreated)
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
		it.Event = new(VotemarketV2BribeCreated)
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
func (it *VotemarketV2BribeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2BribeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2BribeCreated represents a BribeCreated event raised by the VotemarketV2 contract.
type VotemarketV2BribeCreated struct {
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
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV2 *VotemarketV2Filterer) FilterBribeCreated(opts *bind.FilterOpts, id []*big.Int, gauge []common.Address) (*VotemarketV2BribeCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "BribeCreated", idRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2BribeCreatedIterator{contract: _VotemarketV2.contract, event: "BribeCreated", logs: logs, sub: sub}, nil
}

// WatchBribeCreated is a free log subscription operation binding the contract event 0x4ed4160f5ef12a0abd9d6134687dd7da5b8274bf240f997a1f377a76c52ccaf5.
//
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV2 *VotemarketV2Filterer) WatchBribeCreated(opts *bind.WatchOpts, sink chan<- *VotemarketV2BribeCreated, id []*big.Int, gauge []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var gaugeRule []interface{}
	for _, gaugeItem := range gauge {
		gaugeRule = append(gaugeRule, gaugeItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "BribeCreated", idRule, gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2BribeCreated)
				if err := _VotemarketV2.contract.UnpackLog(event, "BribeCreated", log); err != nil {
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
// Solidity: event BribeCreated(uint256 indexed id, address indexed gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 rewardPerPeriod, uint256 totalRewardAmount, bool isUpgradeable)
func (_VotemarketV2 *VotemarketV2Filterer) ParseBribeCreated(log types.Log) (*VotemarketV2BribeCreated, error) {
	event := new(VotemarketV2BribeCreated)
	if err := _VotemarketV2.contract.UnpackLog(event, "BribeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2BribeDurationIncreaseQueuedIterator is returned from FilterBribeDurationIncreaseQueued and is used to iterate over the raw logs and unpacked data for BribeDurationIncreaseQueued events raised by the VotemarketV2 contract.
type VotemarketV2BribeDurationIncreaseQueuedIterator struct {
	Event *VotemarketV2BribeDurationIncreaseQueued // Event containing the contract specifics and raw log

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
func (it *VotemarketV2BribeDurationIncreaseQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2BribeDurationIncreaseQueued)
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
		it.Event = new(VotemarketV2BribeDurationIncreaseQueued)
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
func (it *VotemarketV2BribeDurationIncreaseQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2BribeDurationIncreaseQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2BribeDurationIncreaseQueued represents a BribeDurationIncreaseQueued event raised by the VotemarketV2 contract.
type VotemarketV2BribeDurationIncreaseQueued struct {
	Id                *big.Int
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBribeDurationIncreaseQueued is a free log retrieval operation binding the contract event 0x0c841045cbcf87e9cc7521ce9e85cf523d731f87fd5b45feea376cff34067263.
//
// Solidity: event BribeDurationIncreaseQueued(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV2 *VotemarketV2Filterer) FilterBribeDurationIncreaseQueued(opts *bind.FilterOpts) (*VotemarketV2BribeDurationIncreaseQueuedIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "BribeDurationIncreaseQueued")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2BribeDurationIncreaseQueuedIterator{contract: _VotemarketV2.contract, event: "BribeDurationIncreaseQueued", logs: logs, sub: sub}, nil
}

// WatchBribeDurationIncreaseQueued is a free log subscription operation binding the contract event 0x0c841045cbcf87e9cc7521ce9e85cf523d731f87fd5b45feea376cff34067263.
//
// Solidity: event BribeDurationIncreaseQueued(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV2 *VotemarketV2Filterer) WatchBribeDurationIncreaseQueued(opts *bind.WatchOpts, sink chan<- *VotemarketV2BribeDurationIncreaseQueued) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "BribeDurationIncreaseQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2BribeDurationIncreaseQueued)
				if err := _VotemarketV2.contract.UnpackLog(event, "BribeDurationIncreaseQueued", log); err != nil {
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
func (_VotemarketV2 *VotemarketV2Filterer) ParseBribeDurationIncreaseQueued(log types.Log) (*VotemarketV2BribeDurationIncreaseQueued, error) {
	event := new(VotemarketV2BribeDurationIncreaseQueued)
	if err := _VotemarketV2.contract.UnpackLog(event, "BribeDurationIncreaseQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2BribeDurationIncreasedIterator is returned from FilterBribeDurationIncreased and is used to iterate over the raw logs and unpacked data for BribeDurationIncreased events raised by the VotemarketV2 contract.
type VotemarketV2BribeDurationIncreasedIterator struct {
	Event *VotemarketV2BribeDurationIncreased // Event containing the contract specifics and raw log

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
func (it *VotemarketV2BribeDurationIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2BribeDurationIncreased)
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
		it.Event = new(VotemarketV2BribeDurationIncreased)
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
func (it *VotemarketV2BribeDurationIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2BribeDurationIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2BribeDurationIncreased represents a BribeDurationIncreased event raised by the VotemarketV2 contract.
type VotemarketV2BribeDurationIncreased struct {
	Id                *big.Int
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBribeDurationIncreased is a free log retrieval operation binding the contract event 0xe90b0f7fffa9942eb28c4453083b14e929e4f6d39de19dd5f8cef36148b9c63a.
//
// Solidity: event BribeDurationIncreased(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV2 *VotemarketV2Filterer) FilterBribeDurationIncreased(opts *bind.FilterOpts) (*VotemarketV2BribeDurationIncreasedIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "BribeDurationIncreased")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2BribeDurationIncreasedIterator{contract: _VotemarketV2.contract, event: "BribeDurationIncreased", logs: logs, sub: sub}, nil
}

// WatchBribeDurationIncreased is a free log subscription operation binding the contract event 0xe90b0f7fffa9942eb28c4453083b14e929e4f6d39de19dd5f8cef36148b9c63a.
//
// Solidity: event BribeDurationIncreased(uint256 id, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote)
func (_VotemarketV2 *VotemarketV2Filterer) WatchBribeDurationIncreased(opts *bind.WatchOpts, sink chan<- *VotemarketV2BribeDurationIncreased) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "BribeDurationIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2BribeDurationIncreased)
				if err := _VotemarketV2.contract.UnpackLog(event, "BribeDurationIncreased", log); err != nil {
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
func (_VotemarketV2 *VotemarketV2Filterer) ParseBribeDurationIncreased(log types.Log) (*VotemarketV2BribeDurationIncreased, error) {
	event := new(VotemarketV2BribeDurationIncreased)
	if err := _VotemarketV2.contract.UnpackLog(event, "BribeDurationIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2ClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the VotemarketV2 contract.
type VotemarketV2ClaimedIterator struct {
	Event *VotemarketV2Claimed // Event containing the contract specifics and raw log

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
func (it *VotemarketV2ClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2Claimed)
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
		it.Event = new(VotemarketV2Claimed)
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
func (it *VotemarketV2ClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2ClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2Claimed represents a Claimed event raised by the VotemarketV2 contract.
type VotemarketV2Claimed struct {
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
// Solidity: event Claimed(address indexed user, address rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV2 *VotemarketV2Filterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address, bribeId []*big.Int) (*VotemarketV2ClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var bribeIdRule []interface{}
	for _, bribeIdItem := range bribeId {
		bribeIdRule = append(bribeIdRule, bribeIdItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "Claimed", userRule, bribeIdRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2ClaimedIterator{contract: _VotemarketV2.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76.
//
// Solidity: event Claimed(address indexed user, address rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV2 *VotemarketV2Filterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *VotemarketV2Claimed, user []common.Address, bribeId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var bribeIdRule []interface{}
	for _, bribeIdItem := range bribeId {
		bribeIdRule = append(bribeIdRule, bribeIdItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "Claimed", userRule, bribeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2Claimed)
				if err := _VotemarketV2.contract.UnpackLog(event, "Claimed", log); err != nil {
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
// Solidity: event Claimed(address indexed user, address rewardToken, uint256 indexed bribeId, uint256 amount, uint256 protocolFees, uint256 period)
func (_VotemarketV2 *VotemarketV2Filterer) ParseClaimed(log types.Log) (*VotemarketV2Claimed, error) {
	event := new(VotemarketV2Claimed)
	if err := _VotemarketV2.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2FeeCollectorUpdatedIterator is returned from FilterFeeCollectorUpdated and is used to iterate over the raw logs and unpacked data for FeeCollectorUpdated events raised by the VotemarketV2 contract.
type VotemarketV2FeeCollectorUpdatedIterator struct {
	Event *VotemarketV2FeeCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *VotemarketV2FeeCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2FeeCollectorUpdated)
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
		it.Event = new(VotemarketV2FeeCollectorUpdated)
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
func (it *VotemarketV2FeeCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2FeeCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2FeeCollectorUpdated represents a FeeCollectorUpdated event raised by the VotemarketV2 contract.
type VotemarketV2FeeCollectorUpdated struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorUpdated is a free log retrieval operation binding the contract event 0xe5693914d19c789bdee50a362998c0bc8d035a835f9871da5d51152f0582c34f.
//
// Solidity: event FeeCollectorUpdated(address feeCollector)
func (_VotemarketV2 *VotemarketV2Filterer) FilterFeeCollectorUpdated(opts *bind.FilterOpts) (*VotemarketV2FeeCollectorUpdatedIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2FeeCollectorUpdatedIterator{contract: _VotemarketV2.contract, event: "FeeCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorUpdated is a free log subscription operation binding the contract event 0xe5693914d19c789bdee50a362998c0bc8d035a835f9871da5d51152f0582c34f.
//
// Solidity: event FeeCollectorUpdated(address feeCollector)
func (_VotemarketV2 *VotemarketV2Filterer) WatchFeeCollectorUpdated(opts *bind.WatchOpts, sink chan<- *VotemarketV2FeeCollectorUpdated) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "FeeCollectorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2FeeCollectorUpdated)
				if err := _VotemarketV2.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
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

// ParseFeeCollectorUpdated is a log parse operation binding the contract event 0xe5693914d19c789bdee50a362998c0bc8d035a835f9871da5d51152f0582c34f.
//
// Solidity: event FeeCollectorUpdated(address feeCollector)
func (_VotemarketV2 *VotemarketV2Filterer) ParseFeeCollectorUpdated(log types.Log) (*VotemarketV2FeeCollectorUpdated, error) {
	event := new(VotemarketV2FeeCollectorUpdated)
	if err := _VotemarketV2.contract.UnpackLog(event, "FeeCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2FeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the VotemarketV2 contract.
type VotemarketV2FeeUpdatedIterator struct {
	Event *VotemarketV2FeeUpdated // Event containing the contract specifics and raw log

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
func (it *VotemarketV2FeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2FeeUpdated)
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
		it.Event = new(VotemarketV2FeeUpdated)
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
func (it *VotemarketV2FeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2FeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2FeeUpdated represents a FeeUpdated event raised by the VotemarketV2 contract.
type VotemarketV2FeeUpdated struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_VotemarketV2 *VotemarketV2Filterer) FilterFeeUpdated(opts *bind.FilterOpts) (*VotemarketV2FeeUpdatedIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2FeeUpdatedIterator{contract: _VotemarketV2.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_VotemarketV2 *VotemarketV2Filterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *VotemarketV2FeeUpdated) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2FeeUpdated)
				if err := _VotemarketV2.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 fee)
func (_VotemarketV2 *VotemarketV2Filterer) ParseFeeUpdated(log types.Log) (*VotemarketV2FeeUpdated, error) {
	event := new(VotemarketV2FeeUpdated)
	if err := _VotemarketV2.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2FeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the VotemarketV2 contract.
type VotemarketV2FeesCollectedIterator struct {
	Event *VotemarketV2FeesCollected // Event containing the contract specifics and raw log

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
func (it *VotemarketV2FeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2FeesCollected)
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
		it.Event = new(VotemarketV2FeesCollected)
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
func (it *VotemarketV2FeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2FeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2FeesCollected represents a FeesCollected event raised by the VotemarketV2 contract.
type VotemarketV2FeesCollected struct {
	RewardToken common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed rewardToken, uint256 amount)
func (_VotemarketV2 *VotemarketV2Filterer) FilterFeesCollected(opts *bind.FilterOpts, rewardToken []common.Address) (*VotemarketV2FeesCollectedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "FeesCollected", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2FeesCollectedIterator{contract: _VotemarketV2.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed rewardToken, uint256 amount)
func (_VotemarketV2 *VotemarketV2Filterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *VotemarketV2FeesCollected, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "FeesCollected", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2FeesCollected)
				if err := _VotemarketV2.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed rewardToken, uint256 amount)
func (_VotemarketV2 *VotemarketV2Filterer) ParseFeesCollected(log types.Log) (*VotemarketV2FeesCollected, error) {
	event := new(VotemarketV2FeesCollected)
	if err := _VotemarketV2.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2ManagerUpdatedIterator is returned from FilterManagerUpdated and is used to iterate over the raw logs and unpacked data for ManagerUpdated events raised by the VotemarketV2 contract.
type VotemarketV2ManagerUpdatedIterator struct {
	Event *VotemarketV2ManagerUpdated // Event containing the contract specifics and raw log

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
func (it *VotemarketV2ManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2ManagerUpdated)
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
		it.Event = new(VotemarketV2ManagerUpdated)
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
func (it *VotemarketV2ManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2ManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2ManagerUpdated represents a ManagerUpdated event raised by the VotemarketV2 contract.
type VotemarketV2ManagerUpdated struct {
	Id      *big.Int
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerUpdated is a free log retrieval operation binding the contract event 0x5baaf19ee8739d1720c1401554ff4abe5682505ec43d3f2eb61b9dfc0abd9745.
//
// Solidity: event ManagerUpdated(uint256 id, address indexed manager)
func (_VotemarketV2 *VotemarketV2Filterer) FilterManagerUpdated(opts *bind.FilterOpts, manager []common.Address) (*VotemarketV2ManagerUpdatedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2ManagerUpdatedIterator{contract: _VotemarketV2.contract, event: "ManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchManagerUpdated is a free log subscription operation binding the contract event 0x5baaf19ee8739d1720c1401554ff4abe5682505ec43d3f2eb61b9dfc0abd9745.
//
// Solidity: event ManagerUpdated(uint256 id, address indexed manager)
func (_VotemarketV2 *VotemarketV2Filterer) WatchManagerUpdated(opts *bind.WatchOpts, sink chan<- *VotemarketV2ManagerUpdated, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "ManagerUpdated", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2ManagerUpdated)
				if err := _VotemarketV2.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
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
func (_VotemarketV2 *VotemarketV2Filterer) ParseManagerUpdated(log types.Log) (*VotemarketV2ManagerUpdated, error) {
	event := new(VotemarketV2ManagerUpdated)
	if err := _VotemarketV2.contract.UnpackLog(event, "ManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotemarketV2 contract.
type VotemarketV2OwnershipTransferredIterator struct {
	Event *VotemarketV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotemarketV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2OwnershipTransferred)
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
		it.Event = new(VotemarketV2OwnershipTransferred)
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
func (it *VotemarketV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2OwnershipTransferred represents a OwnershipTransferred event raised by the VotemarketV2 contract.
type VotemarketV2OwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_VotemarketV2 *VotemarketV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*VotemarketV2OwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2OwnershipTransferredIterator{contract: _VotemarketV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_VotemarketV2 *VotemarketV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotemarketV2OwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2OwnershipTransferred)
				if err := _VotemarketV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_VotemarketV2 *VotemarketV2Filterer) ParseOwnershipTransferred(log types.Log) (*VotemarketV2OwnershipTransferred, error) {
	event := new(VotemarketV2OwnershipTransferred)
	if err := _VotemarketV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2PeriodRolledOverIterator is returned from FilterPeriodRolledOver and is used to iterate over the raw logs and unpacked data for PeriodRolledOver events raised by the VotemarketV2 contract.
type VotemarketV2PeriodRolledOverIterator struct {
	Event *VotemarketV2PeriodRolledOver // Event containing the contract specifics and raw log

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
func (it *VotemarketV2PeriodRolledOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2PeriodRolledOver)
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
		it.Event = new(VotemarketV2PeriodRolledOver)
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
func (it *VotemarketV2PeriodRolledOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2PeriodRolledOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2PeriodRolledOver represents a PeriodRolledOver event raised by the VotemarketV2 contract.
type VotemarketV2PeriodRolledOver struct {
	Id              *big.Int
	PeriodId        *big.Int
	Timestamp       *big.Int
	RewardPerPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPeriodRolledOver is a free log retrieval operation binding the contract event 0xb77c22cd311809931524bcc8d4a33a61a392e2304c8a7d476b64018e11ed6cb1.
//
// Solidity: event PeriodRolledOver(uint256 id, uint256 periodId, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV2 *VotemarketV2Filterer) FilterPeriodRolledOver(opts *bind.FilterOpts) (*VotemarketV2PeriodRolledOverIterator, error) {

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "PeriodRolledOver")
	if err != nil {
		return nil, err
	}
	return &VotemarketV2PeriodRolledOverIterator{contract: _VotemarketV2.contract, event: "PeriodRolledOver", logs: logs, sub: sub}, nil
}

// WatchPeriodRolledOver is a free log subscription operation binding the contract event 0xb77c22cd311809931524bcc8d4a33a61a392e2304c8a7d476b64018e11ed6cb1.
//
// Solidity: event PeriodRolledOver(uint256 id, uint256 periodId, uint256 timestamp, uint256 rewardPerPeriod)
func (_VotemarketV2 *VotemarketV2Filterer) WatchPeriodRolledOver(opts *bind.WatchOpts, sink chan<- *VotemarketV2PeriodRolledOver) (event.Subscription, error) {

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "PeriodRolledOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2PeriodRolledOver)
				if err := _VotemarketV2.contract.UnpackLog(event, "PeriodRolledOver", log); err != nil {
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
func (_VotemarketV2 *VotemarketV2Filterer) ParsePeriodRolledOver(log types.Log) (*VotemarketV2PeriodRolledOver, error) {
	event := new(VotemarketV2PeriodRolledOver)
	if err := _VotemarketV2.contract.UnpackLog(event, "PeriodRolledOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemarketV2RecipientSetIterator is returned from FilterRecipientSet and is used to iterate over the raw logs and unpacked data for RecipientSet events raised by the VotemarketV2 contract.
type VotemarketV2RecipientSetIterator struct {
	Event *VotemarketV2RecipientSet // Event containing the contract specifics and raw log

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
func (it *VotemarketV2RecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemarketV2RecipientSet)
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
		it.Event = new(VotemarketV2RecipientSet)
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
func (it *VotemarketV2RecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemarketV2RecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemarketV2RecipientSet represents a RecipientSet event raised by the VotemarketV2 contract.
type VotemarketV2RecipientSet struct {
	Sender    common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecipientSet is a free log retrieval operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed sender, address indexed recipient)
func (_VotemarketV2 *VotemarketV2Filterer) FilterRecipientSet(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*VotemarketV2RecipientSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _VotemarketV2.contract.FilterLogs(opts, "RecipientSet", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &VotemarketV2RecipientSetIterator{contract: _VotemarketV2.contract, event: "RecipientSet", logs: logs, sub: sub}, nil
}

// WatchRecipientSet is a free log subscription operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed sender, address indexed recipient)
func (_VotemarketV2 *VotemarketV2Filterer) WatchRecipientSet(opts *bind.WatchOpts, sink chan<- *VotemarketV2RecipientSet, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _VotemarketV2.contract.WatchLogs(opts, "RecipientSet", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemarketV2RecipientSet)
				if err := _VotemarketV2.contract.UnpackLog(event, "RecipientSet", log); err != nil {
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

// ParseRecipientSet is a log parse operation binding the contract event 0xc1416b5cdab50a9fbc872236e1aa54566c6deb40024e63a4b1737ecacf09d6f9.
//
// Solidity: event RecipientSet(address indexed sender, address indexed recipient)
func (_VotemarketV2 *VotemarketV2Filterer) ParseRecipientSet(log types.Log) (*VotemarketV2RecipientSet, error) {
	event := new(VotemarketV2RecipientSet)
	if err := _VotemarketV2.contract.UnpackLog(event, "RecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
