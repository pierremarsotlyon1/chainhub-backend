// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vmV2

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

// Campaign is an auto generated low-level Go binding around an user-defined struct.
type Campaign struct {
	ChainId           *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	TotalDistributed  *big.Int
	StartTimestamp    *big.Int
	EndTimestamp      *big.Int
	Hook              common.Address
}

// CampaignUpgrade is an auto generated low-level Go binding around an user-defined struct.
type CampaignUpgrade struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}

// Period is an auto generated low-level Go binding around an user-defined struct.
type Period struct {
	RewardPerPeriod *big.Int
	RewardPerVote   *big.Int
	Leftover        *big.Int
	Updated         bool
}

// VmV2MetaData contains all meta data concerning the VmV2 contract.
var VmV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochLength\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minimumPeriods\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AUTH_BLACKLISTED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AUTH_GOVERNANCE_ONLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER_ONLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AUTH_WHITELIST_ONLY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CAMPAIGN_ENDED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CAMPAIGN_NOT_ENDED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CLAIM_AMOUNT_EXCEEDS_REWARD_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EPOCH_NOT_VALID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_INPUT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PROTECTED_ACCOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"STATE_MISSING\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_INPUT\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"CampaignClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"}],\"name\":\"CampaignCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"CampaignUpgradeQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"CampaignUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIM_WINDOW_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOSE_WINDOW_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ADDRESSES_PER_CAMPAIGN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_PERIODS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressesByCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_spacer\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDistributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignUpgradeById\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"closeCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"name\":\"createCampaign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customFeeByManager\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"getAddressesByCampaign\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"getCampaign\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDistributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"internalType\":\"structCampaign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getCampaignUpgrade\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCampaignUpgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getPeriodPerCampaign\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"updated\",\"type\":\"bool\"}],\"internalType\":\"structPeriod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getRemainingPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodsLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"}],\"name\":\"increaseTotalRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isClosedCampaign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProtected\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"name\":\"manageCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"periodByCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"updated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setCustomFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isProtected\",\"type\":\"bool\"}],\"name\":\"setIsProtected\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remote\",\"type\":\"address\"}],\"name\":\"setRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalClaimedByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalClaimedByCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_futureGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"updateEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistOnly\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VmV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VmV2MetaData.ABI instead.
var VmV2ABI = VmV2MetaData.ABI

// VmV2 is an auto generated Go binding around an Ethereum contract.
type VmV2 struct {
	VmV2Caller     // Read-only binding to the contract
	VmV2Transactor // Write-only binding to the contract
	VmV2Filterer   // Log filterer for contract events
}

// VmV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VmV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VmV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmV2Session struct {
	Contract     *VmV2             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmV2CallerSession struct {
	Contract *VmV2Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmV2TransactorSession struct {
	Contract     *VmV2Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VmV2Raw struct {
	Contract *VmV2 // Generic contract binding to access the raw methods on
}

// VmV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmV2CallerRaw struct {
	Contract *VmV2Caller // Generic read-only contract binding to access the raw methods on
}

// VmV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmV2TransactorRaw struct {
	Contract *VmV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVmV2 creates a new instance of VmV2, bound to a specific deployed contract.
func NewVmV2(address common.Address, backend bind.ContractBackend) (*VmV2, error) {
	contract, err := bindVmV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VmV2{VmV2Caller: VmV2Caller{contract: contract}, VmV2Transactor: VmV2Transactor{contract: contract}, VmV2Filterer: VmV2Filterer{contract: contract}}, nil
}

// NewVmV2Caller creates a new read-only instance of VmV2, bound to a specific deployed contract.
func NewVmV2Caller(address common.Address, caller bind.ContractCaller) (*VmV2Caller, error) {
	contract, err := bindVmV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmV2Caller{contract: contract}, nil
}

// NewVmV2Transactor creates a new write-only instance of VmV2, bound to a specific deployed contract.
func NewVmV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VmV2Transactor, error) {
	contract, err := bindVmV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmV2Transactor{contract: contract}, nil
}

// NewVmV2Filterer creates a new log filterer instance of VmV2, bound to a specific deployed contract.
func NewVmV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VmV2Filterer, error) {
	contract, err := bindVmV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmV2Filterer{contract: contract}, nil
}

// bindVmV2 binds a generic wrapper to an already deployed contract.
func bindVmV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VmV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmV2 *VmV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmV2.Contract.VmV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmV2 *VmV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmV2.Contract.VmV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmV2 *VmV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmV2.Contract.VmV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmV2 *VmV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmV2 *VmV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmV2 *VmV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmV2.Contract.contract.Transact(opts, method, params...)
}

// CLAIMWINDOWLENGTH is a free data retrieval call binding the contract method 0x1ac5c6ee.
//
// Solidity: function CLAIM_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Caller) CLAIMWINDOWLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "CLAIM_WINDOW_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLAIMWINDOWLENGTH is a free data retrieval call binding the contract method 0x1ac5c6ee.
//
// Solidity: function CLAIM_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Session) CLAIMWINDOWLENGTH() (*big.Int, error) {
	return _VmV2.Contract.CLAIMWINDOWLENGTH(&_VmV2.CallOpts)
}

// CLAIMWINDOWLENGTH is a free data retrieval call binding the contract method 0x1ac5c6ee.
//
// Solidity: function CLAIM_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2CallerSession) CLAIMWINDOWLENGTH() (*big.Int, error) {
	return _VmV2.Contract.CLAIMWINDOWLENGTH(&_VmV2.CallOpts)
}

// CLOSEWINDOWLENGTH is a free data retrieval call binding the contract method 0x2a0621d7.
//
// Solidity: function CLOSE_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Caller) CLOSEWINDOWLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "CLOSE_WINDOW_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLOSEWINDOWLENGTH is a free data retrieval call binding the contract method 0x2a0621d7.
//
// Solidity: function CLOSE_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Session) CLOSEWINDOWLENGTH() (*big.Int, error) {
	return _VmV2.Contract.CLOSEWINDOWLENGTH(&_VmV2.CallOpts)
}

// CLOSEWINDOWLENGTH is a free data retrieval call binding the contract method 0x2a0621d7.
//
// Solidity: function CLOSE_WINDOW_LENGTH() view returns(uint256)
func (_VmV2 *VmV2CallerSession) CLOSEWINDOWLENGTH() (*big.Int, error) {
	return _VmV2.Contract.CLOSEWINDOWLENGTH(&_VmV2.CallOpts)
}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Caller) EPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "EPOCH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() view returns(uint256)
func (_VmV2 *VmV2Session) EPOCHLENGTH() (*big.Int, error) {
	return _VmV2.Contract.EPOCHLENGTH(&_VmV2.CallOpts)
}

// EPOCHLENGTH is a free data retrieval call binding the contract method 0xac4746ab.
//
// Solidity: function EPOCH_LENGTH() view returns(uint256)
func (_VmV2 *VmV2CallerSession) EPOCHLENGTH() (*big.Int, error) {
	return _VmV2.Contract.EPOCHLENGTH(&_VmV2.CallOpts)
}

// MAXADDRESSESPERCAMPAIGN is a free data retrieval call binding the contract method 0x0f62accf.
//
// Solidity: function MAX_ADDRESSES_PER_CAMPAIGN() view returns(uint256)
func (_VmV2 *VmV2Caller) MAXADDRESSESPERCAMPAIGN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "MAX_ADDRESSES_PER_CAMPAIGN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADDRESSESPERCAMPAIGN is a free data retrieval call binding the contract method 0x0f62accf.
//
// Solidity: function MAX_ADDRESSES_PER_CAMPAIGN() view returns(uint256)
func (_VmV2 *VmV2Session) MAXADDRESSESPERCAMPAIGN() (*big.Int, error) {
	return _VmV2.Contract.MAXADDRESSESPERCAMPAIGN(&_VmV2.CallOpts)
}

// MAXADDRESSESPERCAMPAIGN is a free data retrieval call binding the contract method 0x0f62accf.
//
// Solidity: function MAX_ADDRESSES_PER_CAMPAIGN() view returns(uint256)
func (_VmV2 *VmV2CallerSession) MAXADDRESSESPERCAMPAIGN() (*big.Int, error) {
	return _VmV2.Contract.MAXADDRESSESPERCAMPAIGN(&_VmV2.CallOpts)
}

// MINIMUMPERIODS is a free data retrieval call binding the contract method 0x8e3c282b.
//
// Solidity: function MINIMUM_PERIODS() view returns(uint8)
func (_VmV2 *VmV2Caller) MINIMUMPERIODS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "MINIMUM_PERIODS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINIMUMPERIODS is a free data retrieval call binding the contract method 0x8e3c282b.
//
// Solidity: function MINIMUM_PERIODS() view returns(uint8)
func (_VmV2 *VmV2Session) MINIMUMPERIODS() (uint8, error) {
	return _VmV2.Contract.MINIMUMPERIODS(&_VmV2.CallOpts)
}

// MINIMUMPERIODS is a free data retrieval call binding the contract method 0x8e3c282b.
//
// Solidity: function MINIMUM_PERIODS() view returns(uint8)
func (_VmV2 *VmV2CallerSession) MINIMUMPERIODS() (uint8, error) {
	return _VmV2.Contract.MINIMUMPERIODS(&_VmV2.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_VmV2 *VmV2Caller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_VmV2 *VmV2Session) ORACLE() (common.Address, error) {
	return _VmV2.Contract.ORACLE(&_VmV2.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_VmV2 *VmV2CallerSession) ORACLE() (common.Address, error) {
	return _VmV2.Contract.ORACLE(&_VmV2.CallOpts)
}

// AddressesByCampaignId is a free data retrieval call binding the contract method 0x8e57daa7.
//
// Solidity: function addressesByCampaignId(uint256 ) view returns(uint256 _spacer)
func (_VmV2 *VmV2Caller) AddressesByCampaignId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "addressesByCampaignId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressesByCampaignId is a free data retrieval call binding the contract method 0x8e57daa7.
//
// Solidity: function addressesByCampaignId(uint256 ) view returns(uint256 _spacer)
func (_VmV2 *VmV2Session) AddressesByCampaignId(arg0 *big.Int) (*big.Int, error) {
	return _VmV2.Contract.AddressesByCampaignId(&_VmV2.CallOpts, arg0)
}

// AddressesByCampaignId is a free data retrieval call binding the contract method 0x8e57daa7.
//
// Solidity: function addressesByCampaignId(uint256 ) view returns(uint256 _spacer)
func (_VmV2 *VmV2CallerSession) AddressesByCampaignId(arg0 *big.Int) (*big.Int, error) {
	return _VmV2.Contract.AddressesByCampaignId(&_VmV2.CallOpts, arg0)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, uint256 totalDistributed, uint256 startTimestamp, uint256 endTimestamp, address hook)
func (_VmV2 *VmV2Caller) CampaignById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainId           *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	TotalDistributed  *big.Int
	StartTimestamp    *big.Int
	EndTimestamp      *big.Int
	Hook              common.Address
}, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "campaignById", arg0)

	outstruct := new(struct {
		ChainId           *big.Int
		Gauge             common.Address
		Manager           common.Address
		RewardToken       common.Address
		NumberOfPeriods   uint8
		MaxRewardPerVote  *big.Int
		TotalRewardAmount *big.Int
		TotalDistributed  *big.Int
		StartTimestamp    *big.Int
		EndTimestamp      *big.Int
		Hook              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gauge = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Manager = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.RewardToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.NumberOfPeriods = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MaxRewardPerVote = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalRewardAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalDistributed = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.EndTimestamp = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Hook = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, uint256 totalDistributed, uint256 startTimestamp, uint256 endTimestamp, address hook)
func (_VmV2 *VmV2Session) CampaignById(arg0 *big.Int) (struct {
	ChainId           *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	TotalDistributed  *big.Int
	StartTimestamp    *big.Int
	EndTimestamp      *big.Int
	Hook              common.Address
}, error) {
	return _VmV2.Contract.CampaignById(&_VmV2.CallOpts, arg0)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, uint256 totalDistributed, uint256 startTimestamp, uint256 endTimestamp, address hook)
func (_VmV2 *VmV2CallerSession) CampaignById(arg0 *big.Int) (struct {
	ChainId           *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	TotalDistributed  *big.Int
	StartTimestamp    *big.Int
	EndTimestamp      *big.Int
	Hook              common.Address
}, error) {
	return _VmV2.Contract.CampaignById(&_VmV2.CallOpts, arg0)
}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_VmV2 *VmV2Caller) CampaignCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "campaignCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_VmV2 *VmV2Session) CampaignCount() (*big.Int, error) {
	return _VmV2.Contract.CampaignCount(&_VmV2.CallOpts)
}

// CampaignCount is a free data retrieval call binding the contract method 0x7274e30d.
//
// Solidity: function campaignCount() view returns(uint256)
func (_VmV2 *VmV2CallerSession) CampaignCount() (*big.Int, error) {
	return _VmV2.Contract.CampaignCount(&_VmV2.CallOpts)
}

// CampaignUpgradeById is a free data retrieval call binding the contract method 0x3ce9772e.
//
// Solidity: function campaignUpgradeById(uint256 , uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VmV2 *VmV2Caller) CampaignUpgradeById(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "campaignUpgradeById", arg0, arg1)

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

// CampaignUpgradeById is a free data retrieval call binding the contract method 0x3ce9772e.
//
// Solidity: function campaignUpgradeById(uint256 , uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VmV2 *VmV2Session) CampaignUpgradeById(arg0 *big.Int, arg1 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VmV2.Contract.CampaignUpgradeById(&_VmV2.CallOpts, arg0, arg1)
}

// CampaignUpgradeById is a free data retrieval call binding the contract method 0x3ce9772e.
//
// Solidity: function campaignUpgradeById(uint256 , uint256 ) view returns(uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote, uint256 endTimestamp)
func (_VmV2 *VmV2CallerSession) CampaignUpgradeById(arg0 *big.Int, arg1 *big.Int) (struct {
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
	EndTimestamp      *big.Int
}, error) {
	return _VmV2.Contract.CampaignUpgradeById(&_VmV2.CallOpts, arg0, arg1)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VmV2 *VmV2Caller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VmV2 *VmV2Session) CurrentEpoch() (*big.Int, error) {
	return _VmV2.Contract.CurrentEpoch(&_VmV2.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VmV2 *VmV2CallerSession) CurrentEpoch() (*big.Int, error) {
	return _VmV2.Contract.CurrentEpoch(&_VmV2.CallOpts)
}

// CustomFeeByManager is a free data retrieval call binding the contract method 0x03f17e56.
//
// Solidity: function customFeeByManager(address ) view returns(uint256)
func (_VmV2 *VmV2Caller) CustomFeeByManager(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "customFeeByManager", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustomFeeByManager is a free data retrieval call binding the contract method 0x03f17e56.
//
// Solidity: function customFeeByManager(address ) view returns(uint256)
func (_VmV2 *VmV2Session) CustomFeeByManager(arg0 common.Address) (*big.Int, error) {
	return _VmV2.Contract.CustomFeeByManager(&_VmV2.CallOpts, arg0)
}

// CustomFeeByManager is a free data retrieval call binding the contract method 0x03f17e56.
//
// Solidity: function customFeeByManager(address ) view returns(uint256)
func (_VmV2 *VmV2CallerSession) CustomFeeByManager(arg0 common.Address) (*big.Int, error) {
	return _VmV2.Contract.CustomFeeByManager(&_VmV2.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VmV2 *VmV2Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VmV2 *VmV2Session) Fee() (*big.Int, error) {
	return _VmV2.Contract.Fee(&_VmV2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_VmV2 *VmV2CallerSession) Fee() (*big.Int, error) {
	return _VmV2.Contract.Fee(&_VmV2.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VmV2 *VmV2Caller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VmV2 *VmV2Session) FeeCollector() (common.Address, error) {
	return _VmV2.Contract.FeeCollector(&_VmV2.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_VmV2 *VmV2CallerSession) FeeCollector() (common.Address, error) {
	return _VmV2.Contract.FeeCollector(&_VmV2.CallOpts)
}

// FutureGovernance is a free data retrieval call binding the contract method 0x8070c503.
//
// Solidity: function futureGovernance() view returns(address)
func (_VmV2 *VmV2Caller) FutureGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "futureGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureGovernance is a free data retrieval call binding the contract method 0x8070c503.
//
// Solidity: function futureGovernance() view returns(address)
func (_VmV2 *VmV2Session) FutureGovernance() (common.Address, error) {
	return _VmV2.Contract.FutureGovernance(&_VmV2.CallOpts)
}

// FutureGovernance is a free data retrieval call binding the contract method 0x8070c503.
//
// Solidity: function futureGovernance() view returns(address)
func (_VmV2 *VmV2CallerSession) FutureGovernance() (common.Address, error) {
	return _VmV2.Contract.FutureGovernance(&_VmV2.CallOpts)
}

// GetAddressesByCampaign is a free data retrieval call binding the contract method 0x6fb08df2.
//
// Solidity: function getAddressesByCampaign(uint256 campaignId) view returns(address[])
func (_VmV2 *VmV2Caller) GetAddressesByCampaign(opts *bind.CallOpts, campaignId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "getAddressesByCampaign", campaignId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddressesByCampaign is a free data retrieval call binding the contract method 0x6fb08df2.
//
// Solidity: function getAddressesByCampaign(uint256 campaignId) view returns(address[])
func (_VmV2 *VmV2Session) GetAddressesByCampaign(campaignId *big.Int) ([]common.Address, error) {
	return _VmV2.Contract.GetAddressesByCampaign(&_VmV2.CallOpts, campaignId)
}

// GetAddressesByCampaign is a free data retrieval call binding the contract method 0x6fb08df2.
//
// Solidity: function getAddressesByCampaign(uint256 campaignId) view returns(address[])
func (_VmV2 *VmV2CallerSession) GetAddressesByCampaign(campaignId *big.Int) ([]common.Address, error) {
	return _VmV2.Contract.GetAddressesByCampaign(&_VmV2.CallOpts, campaignId)
}

// GetCampaign is a free data retrieval call binding the contract method 0x5598f8cc.
//
// Solidity: function getCampaign(uint256 campaignId) view returns((uint256,address,address,address,uint8,uint256,uint256,uint256,uint256,uint256,address))
func (_VmV2 *VmV2Caller) GetCampaign(opts *bind.CallOpts, campaignId *big.Int) (Campaign, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "getCampaign", campaignId)

	if err != nil {
		return *new(Campaign), err
	}

	out0 := *abi.ConvertType(out[0], new(Campaign)).(*Campaign)

	return out0, err

}

// GetCampaign is a free data retrieval call binding the contract method 0x5598f8cc.
//
// Solidity: function getCampaign(uint256 campaignId) view returns((uint256,address,address,address,uint8,uint256,uint256,uint256,uint256,uint256,address))
func (_VmV2 *VmV2Session) GetCampaign(campaignId *big.Int) (Campaign, error) {
	return _VmV2.Contract.GetCampaign(&_VmV2.CallOpts, campaignId)
}

// GetCampaign is a free data retrieval call binding the contract method 0x5598f8cc.
//
// Solidity: function getCampaign(uint256 campaignId) view returns((uint256,address,address,address,uint8,uint256,uint256,uint256,uint256,uint256,address))
func (_VmV2 *VmV2CallerSession) GetCampaign(campaignId *big.Int) (Campaign, error) {
	return _VmV2.Contract.GetCampaign(&_VmV2.CallOpts, campaignId)
}

// GetCampaignUpgrade is a free data retrieval call binding the contract method 0x17f3a722.
//
// Solidity: function getCampaignUpgrade(uint256 campaignId, uint256 epoch) view returns((uint8,uint256,uint256,uint256))
func (_VmV2 *VmV2Caller) GetCampaignUpgrade(opts *bind.CallOpts, campaignId *big.Int, epoch *big.Int) (CampaignUpgrade, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "getCampaignUpgrade", campaignId, epoch)

	if err != nil {
		return *new(CampaignUpgrade), err
	}

	out0 := *abi.ConvertType(out[0], new(CampaignUpgrade)).(*CampaignUpgrade)

	return out0, err

}

// GetCampaignUpgrade is a free data retrieval call binding the contract method 0x17f3a722.
//
// Solidity: function getCampaignUpgrade(uint256 campaignId, uint256 epoch) view returns((uint8,uint256,uint256,uint256))
func (_VmV2 *VmV2Session) GetCampaignUpgrade(campaignId *big.Int, epoch *big.Int) (CampaignUpgrade, error) {
	return _VmV2.Contract.GetCampaignUpgrade(&_VmV2.CallOpts, campaignId, epoch)
}

// GetCampaignUpgrade is a free data retrieval call binding the contract method 0x17f3a722.
//
// Solidity: function getCampaignUpgrade(uint256 campaignId, uint256 epoch) view returns((uint8,uint256,uint256,uint256))
func (_VmV2 *VmV2CallerSession) GetCampaignUpgrade(campaignId *big.Int, epoch *big.Int) (CampaignUpgrade, error) {
	return _VmV2.Contract.GetCampaignUpgrade(&_VmV2.CallOpts, campaignId, epoch)
}

// GetPeriodPerCampaign is a free data retrieval call binding the contract method 0x52aed578.
//
// Solidity: function getPeriodPerCampaign(uint256 campaignId, uint256 epoch) view returns((uint256,uint256,uint256,bool))
func (_VmV2 *VmV2Caller) GetPeriodPerCampaign(opts *bind.CallOpts, campaignId *big.Int, epoch *big.Int) (Period, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "getPeriodPerCampaign", campaignId, epoch)

	if err != nil {
		return *new(Period), err
	}

	out0 := *abi.ConvertType(out[0], new(Period)).(*Period)

	return out0, err

}

// GetPeriodPerCampaign is a free data retrieval call binding the contract method 0x52aed578.
//
// Solidity: function getPeriodPerCampaign(uint256 campaignId, uint256 epoch) view returns((uint256,uint256,uint256,bool))
func (_VmV2 *VmV2Session) GetPeriodPerCampaign(campaignId *big.Int, epoch *big.Int) (Period, error) {
	return _VmV2.Contract.GetPeriodPerCampaign(&_VmV2.CallOpts, campaignId, epoch)
}

// GetPeriodPerCampaign is a free data retrieval call binding the contract method 0x52aed578.
//
// Solidity: function getPeriodPerCampaign(uint256 campaignId, uint256 epoch) view returns((uint256,uint256,uint256,bool))
func (_VmV2 *VmV2CallerSession) GetPeriodPerCampaign(campaignId *big.Int, epoch *big.Int) (Period, error) {
	return _VmV2.Contract.GetPeriodPerCampaign(&_VmV2.CallOpts, campaignId, epoch)
}

// GetRemainingPeriods is a free data retrieval call binding the contract method 0x9dd53597.
//
// Solidity: function getRemainingPeriods(uint256 campaignId, uint256 epoch) view returns(uint256 periodsLeft)
func (_VmV2 *VmV2Caller) GetRemainingPeriods(opts *bind.CallOpts, campaignId *big.Int, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "getRemainingPeriods", campaignId, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingPeriods is a free data retrieval call binding the contract method 0x9dd53597.
//
// Solidity: function getRemainingPeriods(uint256 campaignId, uint256 epoch) view returns(uint256 periodsLeft)
func (_VmV2 *VmV2Session) GetRemainingPeriods(campaignId *big.Int, epoch *big.Int) (*big.Int, error) {
	return _VmV2.Contract.GetRemainingPeriods(&_VmV2.CallOpts, campaignId, epoch)
}

// GetRemainingPeriods is a free data retrieval call binding the contract method 0x9dd53597.
//
// Solidity: function getRemainingPeriods(uint256 campaignId, uint256 epoch) view returns(uint256 periodsLeft)
func (_VmV2 *VmV2CallerSession) GetRemainingPeriods(campaignId *big.Int, epoch *big.Int) (*big.Int, error) {
	return _VmV2.Contract.GetRemainingPeriods(&_VmV2.CallOpts, campaignId, epoch)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VmV2 *VmV2Caller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VmV2 *VmV2Session) Governance() (common.Address, error) {
	return _VmV2.Contract.Governance(&_VmV2.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_VmV2 *VmV2CallerSession) Governance() (common.Address, error) {
	return _VmV2.Contract.Governance(&_VmV2.CallOpts)
}

// IsClosedCampaign is a free data retrieval call binding the contract method 0x55b88ff6.
//
// Solidity: function isClosedCampaign(uint256 ) view returns(bool)
func (_VmV2 *VmV2Caller) IsClosedCampaign(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "isClosedCampaign", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClosedCampaign is a free data retrieval call binding the contract method 0x55b88ff6.
//
// Solidity: function isClosedCampaign(uint256 ) view returns(bool)
func (_VmV2 *VmV2Session) IsClosedCampaign(arg0 *big.Int) (bool, error) {
	return _VmV2.Contract.IsClosedCampaign(&_VmV2.CallOpts, arg0)
}

// IsClosedCampaign is a free data retrieval call binding the contract method 0x55b88ff6.
//
// Solidity: function isClosedCampaign(uint256 ) view returns(bool)
func (_VmV2 *VmV2CallerSession) IsClosedCampaign(arg0 *big.Int) (bool, error) {
	return _VmV2.Contract.IsClosedCampaign(&_VmV2.CallOpts, arg0)
}

// IsProtected is a free data retrieval call binding the contract method 0xce35de95.
//
// Solidity: function isProtected(address ) view returns(bool)
func (_VmV2 *VmV2Caller) IsProtected(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "isProtected", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProtected is a free data retrieval call binding the contract method 0xce35de95.
//
// Solidity: function isProtected(address ) view returns(bool)
func (_VmV2 *VmV2Session) IsProtected(arg0 common.Address) (bool, error) {
	return _VmV2.Contract.IsProtected(&_VmV2.CallOpts, arg0)
}

// IsProtected is a free data retrieval call binding the contract method 0xce35de95.
//
// Solidity: function isProtected(address ) view returns(bool)
func (_VmV2 *VmV2CallerSession) IsProtected(arg0 common.Address) (bool, error) {
	return _VmV2.Contract.IsProtected(&_VmV2.CallOpts, arg0)
}

// PeriodByCampaignId is a free data retrieval call binding the contract method 0x99e9d29e.
//
// Solidity: function periodByCampaignId(uint256 , uint256 ) view returns(uint256 rewardPerPeriod, uint256 rewardPerVote, uint256 leftover, bool updated)
func (_VmV2 *VmV2Caller) PeriodByCampaignId(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	RewardPerPeriod *big.Int
	RewardPerVote   *big.Int
	Leftover        *big.Int
	Updated         bool
}, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "periodByCampaignId", arg0, arg1)

	outstruct := new(struct {
		RewardPerPeriod *big.Int
		RewardPerVote   *big.Int
		Leftover        *big.Int
		Updated         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardPerPeriod = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPerVote = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Leftover = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Updated = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// PeriodByCampaignId is a free data retrieval call binding the contract method 0x99e9d29e.
//
// Solidity: function periodByCampaignId(uint256 , uint256 ) view returns(uint256 rewardPerPeriod, uint256 rewardPerVote, uint256 leftover, bool updated)
func (_VmV2 *VmV2Session) PeriodByCampaignId(arg0 *big.Int, arg1 *big.Int) (struct {
	RewardPerPeriod *big.Int
	RewardPerVote   *big.Int
	Leftover        *big.Int
	Updated         bool
}, error) {
	return _VmV2.Contract.PeriodByCampaignId(&_VmV2.CallOpts, arg0, arg1)
}

// PeriodByCampaignId is a free data retrieval call binding the contract method 0x99e9d29e.
//
// Solidity: function periodByCampaignId(uint256 , uint256 ) view returns(uint256 rewardPerPeriod, uint256 rewardPerVote, uint256 leftover, bool updated)
func (_VmV2 *VmV2CallerSession) PeriodByCampaignId(arg0 *big.Int, arg1 *big.Int) (struct {
	RewardPerPeriod *big.Int
	RewardPerVote   *big.Int
	Leftover        *big.Int
	Updated         bool
}, error) {
	return _VmV2.Contract.PeriodByCampaignId(&_VmV2.CallOpts, arg0, arg1)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_VmV2 *VmV2Caller) Recipients(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_VmV2 *VmV2Session) Recipients(arg0 common.Address) (common.Address, error) {
	return _VmV2.Contract.Recipients(&_VmV2.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_VmV2 *VmV2CallerSession) Recipients(arg0 common.Address) (common.Address, error) {
	return _VmV2.Contract.Recipients(&_VmV2.CallOpts, arg0)
}

// Remote is a free data retrieval call binding the contract method 0xd598215b.
//
// Solidity: function remote() view returns(address)
func (_VmV2 *VmV2Caller) Remote(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "remote")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Remote is a free data retrieval call binding the contract method 0xd598215b.
//
// Solidity: function remote() view returns(address)
func (_VmV2 *VmV2Session) Remote() (common.Address, error) {
	return _VmV2.Contract.Remote(&_VmV2.CallOpts)
}

// Remote is a free data retrieval call binding the contract method 0xd598215b.
//
// Solidity: function remote() view returns(address)
func (_VmV2 *VmV2CallerSession) Remote() (common.Address, error) {
	return _VmV2.Contract.Remote(&_VmV2.CallOpts)
}

// TotalClaimedByAccount is a free data retrieval call binding the contract method 0x66629f12.
//
// Solidity: function totalClaimedByAccount(uint256 , uint256 , address ) view returns(uint256)
func (_VmV2 *VmV2Caller) TotalClaimedByAccount(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "totalClaimedByAccount", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedByAccount is a free data retrieval call binding the contract method 0x66629f12.
//
// Solidity: function totalClaimedByAccount(uint256 , uint256 , address ) view returns(uint256)
func (_VmV2 *VmV2Session) TotalClaimedByAccount(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _VmV2.Contract.TotalClaimedByAccount(&_VmV2.CallOpts, arg0, arg1, arg2)
}

// TotalClaimedByAccount is a free data retrieval call binding the contract method 0x66629f12.
//
// Solidity: function totalClaimedByAccount(uint256 , uint256 , address ) view returns(uint256)
func (_VmV2 *VmV2CallerSession) TotalClaimedByAccount(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _VmV2.Contract.TotalClaimedByAccount(&_VmV2.CallOpts, arg0, arg1, arg2)
}

// TotalClaimedByCampaignId is a free data retrieval call binding the contract method 0x5dda2d63.
//
// Solidity: function totalClaimedByCampaignId(uint256 ) view returns(uint256)
func (_VmV2 *VmV2Caller) TotalClaimedByCampaignId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "totalClaimedByCampaignId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedByCampaignId is a free data retrieval call binding the contract method 0x5dda2d63.
//
// Solidity: function totalClaimedByCampaignId(uint256 ) view returns(uint256)
func (_VmV2 *VmV2Session) TotalClaimedByCampaignId(arg0 *big.Int) (*big.Int, error) {
	return _VmV2.Contract.TotalClaimedByCampaignId(&_VmV2.CallOpts, arg0)
}

// TotalClaimedByCampaignId is a free data retrieval call binding the contract method 0x5dda2d63.
//
// Solidity: function totalClaimedByCampaignId(uint256 ) view returns(uint256)
func (_VmV2 *VmV2CallerSession) TotalClaimedByCampaignId(arg0 *big.Int) (*big.Int, error) {
	return _VmV2.Contract.TotalClaimedByCampaignId(&_VmV2.CallOpts, arg0)
}

// WhitelistOnly is a free data retrieval call binding the contract method 0xdf4c3d10.
//
// Solidity: function whitelistOnly(uint256 ) view returns(bool)
func (_VmV2 *VmV2Caller) WhitelistOnly(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VmV2.contract.Call(opts, &out, "whitelistOnly", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistOnly is a free data retrieval call binding the contract method 0xdf4c3d10.
//
// Solidity: function whitelistOnly(uint256 ) view returns(bool)
func (_VmV2 *VmV2Session) WhitelistOnly(arg0 *big.Int) (bool, error) {
	return _VmV2.Contract.WhitelistOnly(&_VmV2.CallOpts, arg0)
}

// WhitelistOnly is a free data retrieval call binding the contract method 0xdf4c3d10.
//
// Solidity: function whitelistOnly(uint256 ) view returns(bool)
func (_VmV2 *VmV2CallerSession) WhitelistOnly(arg0 *big.Int) (bool, error) {
	return _VmV2.Contract.WhitelistOnly(&_VmV2.CallOpts, arg0)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_VmV2 *VmV2Transactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_VmV2 *VmV2Session) AcceptGovernance() (*types.Transaction, error) {
	return _VmV2.Contract.AcceptGovernance(&_VmV2.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_VmV2 *VmV2TransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _VmV2.Contract.AcceptGovernance(&_VmV2.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x09066d7c.
//
// Solidity: function claim(uint256 campaignId, uint256 epoch, bytes hookData, address receiver) returns(uint256 claimed)
func (_VmV2 *VmV2Transactor) Claim(opts *bind.TransactOpts, campaignId *big.Int, epoch *big.Int, hookData []byte, receiver common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "claim", campaignId, epoch, hookData, receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x09066d7c.
//
// Solidity: function claim(uint256 campaignId, uint256 epoch, bytes hookData, address receiver) returns(uint256 claimed)
func (_VmV2 *VmV2Session) Claim(campaignId *big.Int, epoch *big.Int, hookData []byte, receiver common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.Claim(&_VmV2.TransactOpts, campaignId, epoch, hookData, receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x09066d7c.
//
// Solidity: function claim(uint256 campaignId, uint256 epoch, bytes hookData, address receiver) returns(uint256 claimed)
func (_VmV2 *VmV2TransactorSession) Claim(campaignId *big.Int, epoch *big.Int, hookData []byte, receiver common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.Claim(&_VmV2.TransactOpts, campaignId, epoch, hookData, receiver)
}

// Claim0 is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 campaignId, address account, uint256 epoch, bytes hookData) returns(uint256 claimed)
func (_VmV2 *VmV2Transactor) Claim0(opts *bind.TransactOpts, campaignId *big.Int, account common.Address, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "claim0", campaignId, account, epoch, hookData)
}

// Claim0 is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 campaignId, address account, uint256 epoch, bytes hookData) returns(uint256 claimed)
func (_VmV2 *VmV2Session) Claim0(campaignId *big.Int, account common.Address, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.Contract.Claim0(&_VmV2.TransactOpts, campaignId, account, epoch, hookData)
}

// Claim0 is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 campaignId, address account, uint256 epoch, bytes hookData) returns(uint256 claimed)
func (_VmV2 *VmV2TransactorSession) Claim0(campaignId *big.Int, account common.Address, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.Contract.Claim0(&_VmV2.TransactOpts, campaignId, account, epoch, hookData)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xb0e1c1e1.
//
// Solidity: function closeCampaign(uint256 campaignId) returns()
func (_VmV2 *VmV2Transactor) CloseCampaign(opts *bind.TransactOpts, campaignId *big.Int) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "closeCampaign", campaignId)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xb0e1c1e1.
//
// Solidity: function closeCampaign(uint256 campaignId) returns()
func (_VmV2 *VmV2Session) CloseCampaign(campaignId *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.CloseCampaign(&_VmV2.TransactOpts, campaignId)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xb0e1c1e1.
//
// Solidity: function closeCampaign(uint256 campaignId) returns()
func (_VmV2 *VmV2TransactorSession) CloseCampaign(campaignId *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.CloseCampaign(&_VmV2.TransactOpts, campaignId)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xb956358c.
//
// Solidity: function createCampaign(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] addresses, address hook, bool isWhitelist) returns(uint256 campaignId)
func (_VmV2 *VmV2Transactor) CreateCampaign(opts *bind.TransactOpts, chainId *big.Int, gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, addresses []common.Address, hook common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "createCampaign", chainId, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, addresses, hook, isWhitelist)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xb956358c.
//
// Solidity: function createCampaign(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] addresses, address hook, bool isWhitelist) returns(uint256 campaignId)
func (_VmV2 *VmV2Session) CreateCampaign(chainId *big.Int, gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, addresses []common.Address, hook common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _VmV2.Contract.CreateCampaign(&_VmV2.TransactOpts, chainId, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, addresses, hook, isWhitelist)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xb956358c.
//
// Solidity: function createCampaign(uint256 chainId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount, address[] addresses, address hook, bool isWhitelist) returns(uint256 campaignId)
func (_VmV2 *VmV2TransactorSession) CreateCampaign(chainId *big.Int, gauge common.Address, manager common.Address, rewardToken common.Address, numberOfPeriods uint8, maxRewardPerVote *big.Int, totalRewardAmount *big.Int, addresses []common.Address, hook common.Address, isWhitelist bool) (*types.Transaction, error) {
	return _VmV2.Contract.CreateCampaign(&_VmV2.TransactOpts, chainId, gauge, manager, rewardToken, numberOfPeriods, maxRewardPerVote, totalRewardAmount, addresses, hook, isWhitelist)
}

// IncreaseTotalRewardAmount is a paid mutator transaction binding the contract method 0xfffa86fc.
//
// Solidity: function increaseTotalRewardAmount(uint256 campaignId, uint256 totalRewardAmount) returns()
func (_VmV2 *VmV2Transactor) IncreaseTotalRewardAmount(opts *bind.TransactOpts, campaignId *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "increaseTotalRewardAmount", campaignId, totalRewardAmount)
}

// IncreaseTotalRewardAmount is a paid mutator transaction binding the contract method 0xfffa86fc.
//
// Solidity: function increaseTotalRewardAmount(uint256 campaignId, uint256 totalRewardAmount) returns()
func (_VmV2 *VmV2Session) IncreaseTotalRewardAmount(campaignId *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.IncreaseTotalRewardAmount(&_VmV2.TransactOpts, campaignId, totalRewardAmount)
}

// IncreaseTotalRewardAmount is a paid mutator transaction binding the contract method 0xfffa86fc.
//
// Solidity: function increaseTotalRewardAmount(uint256 campaignId, uint256 totalRewardAmount) returns()
func (_VmV2 *VmV2TransactorSession) IncreaseTotalRewardAmount(campaignId *big.Int, totalRewardAmount *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.IncreaseTotalRewardAmount(&_VmV2.TransactOpts, campaignId, totalRewardAmount)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0x6dbb5011.
//
// Solidity: function manageCampaign(uint256 campaignId, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote) returns()
func (_VmV2 *VmV2Transactor) ManageCampaign(opts *bind.TransactOpts, campaignId *big.Int, numberOfPeriods uint8, totalRewardAmount *big.Int, maxRewardPerVote *big.Int) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "manageCampaign", campaignId, numberOfPeriods, totalRewardAmount, maxRewardPerVote)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0x6dbb5011.
//
// Solidity: function manageCampaign(uint256 campaignId, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote) returns()
func (_VmV2 *VmV2Session) ManageCampaign(campaignId *big.Int, numberOfPeriods uint8, totalRewardAmount *big.Int, maxRewardPerVote *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.ManageCampaign(&_VmV2.TransactOpts, campaignId, numberOfPeriods, totalRewardAmount, maxRewardPerVote)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0x6dbb5011.
//
// Solidity: function manageCampaign(uint256 campaignId, uint8 numberOfPeriods, uint256 totalRewardAmount, uint256 maxRewardPerVote) returns()
func (_VmV2 *VmV2TransactorSession) ManageCampaign(campaignId *big.Int, numberOfPeriods uint8, totalRewardAmount *big.Int, maxRewardPerVote *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.ManageCampaign(&_VmV2.TransactOpts, campaignId, numberOfPeriods, totalRewardAmount, maxRewardPerVote)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _account, uint256 _fee) returns()
func (_VmV2 *VmV2Transactor) SetCustomFee(opts *bind.TransactOpts, _account common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setCustomFee", _account, _fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _account, uint256 _fee) returns()
func (_VmV2 *VmV2Session) SetCustomFee(_account common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.SetCustomFee(&_VmV2.TransactOpts, _account, _fee)
}

// SetCustomFee is a paid mutator transaction binding the contract method 0xd49466a8.
//
// Solidity: function setCustomFee(address _account, uint256 _fee) returns()
func (_VmV2 *VmV2TransactorSession) SetCustomFee(_account common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.SetCustomFee(&_VmV2.TransactOpts, _account, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_VmV2 *VmV2Transactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_VmV2 *VmV2Session) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.SetFee(&_VmV2.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_VmV2 *VmV2TransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _VmV2.Contract.SetFee(&_VmV2.TransactOpts, _fee)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VmV2 *VmV2Transactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VmV2 *VmV2Session) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetFeeCollector(&_VmV2.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_VmV2 *VmV2TransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetFeeCollector(&_VmV2.TransactOpts, _feeCollector)
}

// SetIsProtected is a paid mutator transaction binding the contract method 0x1ec72172.
//
// Solidity: function setIsProtected(address _account, bool _isProtected) returns()
func (_VmV2 *VmV2Transactor) SetIsProtected(opts *bind.TransactOpts, _account common.Address, _isProtected bool) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setIsProtected", _account, _isProtected)
}

// SetIsProtected is a paid mutator transaction binding the contract method 0x1ec72172.
//
// Solidity: function setIsProtected(address _account, bool _isProtected) returns()
func (_VmV2 *VmV2Session) SetIsProtected(_account common.Address, _isProtected bool) (*types.Transaction, error) {
	return _VmV2.Contract.SetIsProtected(&_VmV2.TransactOpts, _account, _isProtected)
}

// SetIsProtected is a paid mutator transaction binding the contract method 0x1ec72172.
//
// Solidity: function setIsProtected(address _account, bool _isProtected) returns()
func (_VmV2 *VmV2TransactorSession) SetIsProtected(_account common.Address, _isProtected bool) (*types.Transaction, error) {
	return _VmV2.Contract.SetIsProtected(&_VmV2.TransactOpts, _account, _isProtected)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VmV2 *VmV2Transactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VmV2 *VmV2Session) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRecipient(&_VmV2.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_VmV2 *VmV2TransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRecipient(&_VmV2.TransactOpts, _recipient)
}

// SetRecipient0 is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address _account, address _recipient) returns()
func (_VmV2 *VmV2Transactor) SetRecipient0(opts *bind.TransactOpts, _account common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setRecipient0", _account, _recipient)
}

// SetRecipient0 is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address _account, address _recipient) returns()
func (_VmV2 *VmV2Session) SetRecipient0(_account common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRecipient0(&_VmV2.TransactOpts, _account, _recipient)
}

// SetRecipient0 is a paid mutator transaction binding the contract method 0x8bc8407a.
//
// Solidity: function setRecipient(address _account, address _recipient) returns()
func (_VmV2 *VmV2TransactorSession) SetRecipient0(_account common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRecipient0(&_VmV2.TransactOpts, _account, _recipient)
}

// SetRemote is a paid mutator transaction binding the contract method 0xf1efe24c.
//
// Solidity: function setRemote(address _remote) returns()
func (_VmV2 *VmV2Transactor) SetRemote(opts *bind.TransactOpts, _remote common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "setRemote", _remote)
}

// SetRemote is a paid mutator transaction binding the contract method 0xf1efe24c.
//
// Solidity: function setRemote(address _remote) returns()
func (_VmV2 *VmV2Session) SetRemote(_remote common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRemote(&_VmV2.TransactOpts, _remote)
}

// SetRemote is a paid mutator transaction binding the contract method 0xf1efe24c.
//
// Solidity: function setRemote(address _remote) returns()
func (_VmV2 *VmV2TransactorSession) SetRemote(_remote common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.SetRemote(&_VmV2.TransactOpts, _remote)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _futureGovernance) returns()
func (_VmV2 *VmV2Transactor) TransferGovernance(opts *bind.TransactOpts, _futureGovernance common.Address) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "transferGovernance", _futureGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _futureGovernance) returns()
func (_VmV2 *VmV2Session) TransferGovernance(_futureGovernance common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.TransferGovernance(&_VmV2.TransactOpts, _futureGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _futureGovernance) returns()
func (_VmV2 *VmV2TransactorSession) TransferGovernance(_futureGovernance common.Address) (*types.Transaction, error) {
	return _VmV2.Contract.TransferGovernance(&_VmV2.TransactOpts, _futureGovernance)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xa4bb0946.
//
// Solidity: function updateEpoch(uint256 campaignId, uint256 epoch, bytes hookData) returns(uint256 epoch_)
func (_VmV2 *VmV2Transactor) UpdateEpoch(opts *bind.TransactOpts, campaignId *big.Int, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.contract.Transact(opts, "updateEpoch", campaignId, epoch, hookData)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xa4bb0946.
//
// Solidity: function updateEpoch(uint256 campaignId, uint256 epoch, bytes hookData) returns(uint256 epoch_)
func (_VmV2 *VmV2Session) UpdateEpoch(campaignId *big.Int, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.Contract.UpdateEpoch(&_VmV2.TransactOpts, campaignId, epoch, hookData)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xa4bb0946.
//
// Solidity: function updateEpoch(uint256 campaignId, uint256 epoch, bytes hookData) returns(uint256 epoch_)
func (_VmV2 *VmV2TransactorSession) UpdateEpoch(campaignId *big.Int, epoch *big.Int, hookData []byte) (*types.Transaction, error) {
	return _VmV2.Contract.UpdateEpoch(&_VmV2.TransactOpts, campaignId, epoch, hookData)
}

// VmV2CampaignClosedIterator is returned from FilterCampaignClosed and is used to iterate over the raw logs and unpacked data for CampaignClosed events raised by the VmV2 contract.
type VmV2CampaignClosedIterator struct {
	Event *VmV2CampaignClosed // Event containing the contract specifics and raw log

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
func (it *VmV2CampaignClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmV2CampaignClosed)
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
		it.Event = new(VmV2CampaignClosed)
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
func (it *VmV2CampaignClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmV2CampaignClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmV2CampaignClosed represents a CampaignClosed event raised by the VmV2 contract.
type VmV2CampaignClosed struct {
	CampaignId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignClosed is a free log retrieval operation binding the contract event 0x5e6eb33a418de5dbbc17f989f7ae362cdfbb1748c5d603137c767027a354edbc.
//
// Solidity: event CampaignClosed(uint256 campaignId)
func (_VmV2 *VmV2Filterer) FilterCampaignClosed(opts *bind.FilterOpts) (*VmV2CampaignClosedIterator, error) {

	logs, sub, err := _VmV2.contract.FilterLogs(opts, "CampaignClosed")
	if err != nil {
		return nil, err
	}
	return &VmV2CampaignClosedIterator{contract: _VmV2.contract, event: "CampaignClosed", logs: logs, sub: sub}, nil
}

// WatchCampaignClosed is a free log subscription operation binding the contract event 0x5e6eb33a418de5dbbc17f989f7ae362cdfbb1748c5d603137c767027a354edbc.
//
// Solidity: event CampaignClosed(uint256 campaignId)
func (_VmV2 *VmV2Filterer) WatchCampaignClosed(opts *bind.WatchOpts, sink chan<- *VmV2CampaignClosed) (event.Subscription, error) {

	logs, sub, err := _VmV2.contract.WatchLogs(opts, "CampaignClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmV2CampaignClosed)
				if err := _VmV2.contract.UnpackLog(event, "CampaignClosed", log); err != nil {
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

// ParseCampaignClosed is a log parse operation binding the contract event 0x5e6eb33a418de5dbbc17f989f7ae362cdfbb1748c5d603137c767027a354edbc.
//
// Solidity: event CampaignClosed(uint256 campaignId)
func (_VmV2 *VmV2Filterer) ParseCampaignClosed(log types.Log) (*VmV2CampaignClosed, error) {
	event := new(VmV2CampaignClosed)
	if err := _VmV2.contract.UnpackLog(event, "CampaignClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmV2CampaignCreatedIterator is returned from FilterCampaignCreated and is used to iterate over the raw logs and unpacked data for CampaignCreated events raised by the VmV2 contract.
type VmV2CampaignCreatedIterator struct {
	Event *VmV2CampaignCreated // Event containing the contract specifics and raw log

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
func (it *VmV2CampaignCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmV2CampaignCreated)
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
		it.Event = new(VmV2CampaignCreated)
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
func (it *VmV2CampaignCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmV2CampaignCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmV2CampaignCreated represents a CampaignCreated event raised by the VmV2 contract.
type VmV2CampaignCreated struct {
	CampaignId        *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreated is a free log retrieval operation binding the contract event 0x0e291713fc4cb1bcf9276bc9ae54317736576d5353a44151e2e31c191b1ee62a.
//
// Solidity: event CampaignCreated(uint256 campaignId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VmV2 *VmV2Filterer) FilterCampaignCreated(opts *bind.FilterOpts) (*VmV2CampaignCreatedIterator, error) {

	logs, sub, err := _VmV2.contract.FilterLogs(opts, "CampaignCreated")
	if err != nil {
		return nil, err
	}
	return &VmV2CampaignCreatedIterator{contract: _VmV2.contract, event: "CampaignCreated", logs: logs, sub: sub}, nil
}

// WatchCampaignCreated is a free log subscription operation binding the contract event 0x0e291713fc4cb1bcf9276bc9ae54317736576d5353a44151e2e31c191b1ee62a.
//
// Solidity: event CampaignCreated(uint256 campaignId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VmV2 *VmV2Filterer) WatchCampaignCreated(opts *bind.WatchOpts, sink chan<- *VmV2CampaignCreated) (event.Subscription, error) {

	logs, sub, err := _VmV2.contract.WatchLogs(opts, "CampaignCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmV2CampaignCreated)
				if err := _VmV2.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
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

// ParseCampaignCreated is a log parse operation binding the contract event 0x0e291713fc4cb1bcf9276bc9ae54317736576d5353a44151e2e31c191b1ee62a.
//
// Solidity: event CampaignCreated(uint256 campaignId, address gauge, address manager, address rewardToken, uint8 numberOfPeriods, uint256 maxRewardPerVote, uint256 totalRewardAmount)
func (_VmV2 *VmV2Filterer) ParseCampaignCreated(log types.Log) (*VmV2CampaignCreated, error) {
	event := new(VmV2CampaignCreated)
	if err := _VmV2.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmV2CampaignUpgradeQueuedIterator is returned from FilterCampaignUpgradeQueued and is used to iterate over the raw logs and unpacked data for CampaignUpgradeQueued events raised by the VmV2 contract.
type VmV2CampaignUpgradeQueuedIterator struct {
	Event *VmV2CampaignUpgradeQueued // Event containing the contract specifics and raw log

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
func (it *VmV2CampaignUpgradeQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmV2CampaignUpgradeQueued)
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
		it.Event = new(VmV2CampaignUpgradeQueued)
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
func (it *VmV2CampaignUpgradeQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmV2CampaignUpgradeQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmV2CampaignUpgradeQueued represents a CampaignUpgradeQueued event raised by the VmV2 contract.
type VmV2CampaignUpgradeQueued struct {
	CampaignId *big.Int
	Epoch      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignUpgradeQueued is a free log retrieval operation binding the contract event 0xa8bf14cbd1773eaebffd457f5e6a2badb14741d5e87b85eb41e4adb86080ac0c.
//
// Solidity: event CampaignUpgradeQueued(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) FilterCampaignUpgradeQueued(opts *bind.FilterOpts) (*VmV2CampaignUpgradeQueuedIterator, error) {

	logs, sub, err := _VmV2.contract.FilterLogs(opts, "CampaignUpgradeQueued")
	if err != nil {
		return nil, err
	}
	return &VmV2CampaignUpgradeQueuedIterator{contract: _VmV2.contract, event: "CampaignUpgradeQueued", logs: logs, sub: sub}, nil
}

// WatchCampaignUpgradeQueued is a free log subscription operation binding the contract event 0xa8bf14cbd1773eaebffd457f5e6a2badb14741d5e87b85eb41e4adb86080ac0c.
//
// Solidity: event CampaignUpgradeQueued(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) WatchCampaignUpgradeQueued(opts *bind.WatchOpts, sink chan<- *VmV2CampaignUpgradeQueued) (event.Subscription, error) {

	logs, sub, err := _VmV2.contract.WatchLogs(opts, "CampaignUpgradeQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmV2CampaignUpgradeQueued)
				if err := _VmV2.contract.UnpackLog(event, "CampaignUpgradeQueued", log); err != nil {
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

// ParseCampaignUpgradeQueued is a log parse operation binding the contract event 0xa8bf14cbd1773eaebffd457f5e6a2badb14741d5e87b85eb41e4adb86080ac0c.
//
// Solidity: event CampaignUpgradeQueued(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) ParseCampaignUpgradeQueued(log types.Log) (*VmV2CampaignUpgradeQueued, error) {
	event := new(VmV2CampaignUpgradeQueued)
	if err := _VmV2.contract.UnpackLog(event, "CampaignUpgradeQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmV2CampaignUpgradedIterator is returned from FilterCampaignUpgraded and is used to iterate over the raw logs and unpacked data for CampaignUpgraded events raised by the VmV2 contract.
type VmV2CampaignUpgradedIterator struct {
	Event *VmV2CampaignUpgraded // Event containing the contract specifics and raw log

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
func (it *VmV2CampaignUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmV2CampaignUpgraded)
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
		it.Event = new(VmV2CampaignUpgraded)
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
func (it *VmV2CampaignUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmV2CampaignUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmV2CampaignUpgraded represents a CampaignUpgraded event raised by the VmV2 contract.
type VmV2CampaignUpgraded struct {
	CampaignId *big.Int
	Epoch      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCampaignUpgraded is a free log retrieval operation binding the contract event 0xa416776236c15a5c5398bdb783dfe811a1723512b5fb8914aeb21873defa1872.
//
// Solidity: event CampaignUpgraded(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) FilterCampaignUpgraded(opts *bind.FilterOpts) (*VmV2CampaignUpgradedIterator, error) {

	logs, sub, err := _VmV2.contract.FilterLogs(opts, "CampaignUpgraded")
	if err != nil {
		return nil, err
	}
	return &VmV2CampaignUpgradedIterator{contract: _VmV2.contract, event: "CampaignUpgraded", logs: logs, sub: sub}, nil
}

// WatchCampaignUpgraded is a free log subscription operation binding the contract event 0xa416776236c15a5c5398bdb783dfe811a1723512b5fb8914aeb21873defa1872.
//
// Solidity: event CampaignUpgraded(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) WatchCampaignUpgraded(opts *bind.WatchOpts, sink chan<- *VmV2CampaignUpgraded) (event.Subscription, error) {

	logs, sub, err := _VmV2.contract.WatchLogs(opts, "CampaignUpgraded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmV2CampaignUpgraded)
				if err := _VmV2.contract.UnpackLog(event, "CampaignUpgraded", log); err != nil {
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

// ParseCampaignUpgraded is a log parse operation binding the contract event 0xa416776236c15a5c5398bdb783dfe811a1723512b5fb8914aeb21873defa1872.
//
// Solidity: event CampaignUpgraded(uint256 campaignId, uint256 epoch)
func (_VmV2 *VmV2Filterer) ParseCampaignUpgraded(log types.Log) (*VmV2CampaignUpgraded, error) {
	event := new(VmV2CampaignUpgraded)
	if err := _VmV2.contract.UnpackLog(event, "CampaignUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmV2ClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the VmV2 contract.
type VmV2ClaimIterator struct {
	Event *VmV2Claim // Event containing the contract specifics and raw log

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
func (it *VmV2ClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmV2Claim)
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
		it.Event = new(VmV2Claim)
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
func (it *VmV2ClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmV2ClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmV2Claim represents a Claim event raised by the VmV2 contract.
type VmV2Claim struct {
	CampaignId *big.Int
	Account    common.Address
	Amount     *big.Int
	Fee        *big.Int
	Epoch      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x318e0a24a7fc05b12e358902d9d58475434a01768f87a4319fb35dc5b533e986.
//
// Solidity: event Claim(uint256 indexed campaignId, address indexed account, uint256 amount, uint256 fee, uint256 epoch)
func (_VmV2 *VmV2Filterer) FilterClaim(opts *bind.FilterOpts, campaignId []*big.Int, account []common.Address) (*VmV2ClaimIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VmV2.contract.FilterLogs(opts, "Claim", campaignIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &VmV2ClaimIterator{contract: _VmV2.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x318e0a24a7fc05b12e358902d9d58475434a01768f87a4319fb35dc5b533e986.
//
// Solidity: event Claim(uint256 indexed campaignId, address indexed account, uint256 amount, uint256 fee, uint256 epoch)
func (_VmV2 *VmV2Filterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VmV2Claim, campaignId []*big.Int, account []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VmV2.contract.WatchLogs(opts, "Claim", campaignIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmV2Claim)
				if err := _VmV2.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x318e0a24a7fc05b12e358902d9d58475434a01768f87a4319fb35dc5b533e986.
//
// Solidity: event Claim(uint256 indexed campaignId, address indexed account, uint256 amount, uint256 fee, uint256 epoch)
func (_VmV2 *VmV2Filterer) ParseClaim(log types.Log) (*VmV2Claim, error) {
	event := new(VmV2Claim)
	if err := _VmV2.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
