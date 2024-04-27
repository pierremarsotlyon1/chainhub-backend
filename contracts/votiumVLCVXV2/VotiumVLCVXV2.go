// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votiumVLCVXV2

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

// VotiumIncentive is an auto generated low-level Go binding around an user-defined struct.
type VotiumIncentive struct {
	Token       common.Address
	Amount      *big.Int
	MaxPerVote  *big.Int
	Distributed *big.Int
	Recycled    *big.Int
	Depositor   common.Address
	Excluded    []common.Address
}

// VotiumVLCVXV2MetaData contains all meta data concerning the VotiumVLCVXV2 contract.
var VotiumVLCVXV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approved2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_requireAllowlist\",\"type\":\"bool\"}],\"name\":\"AllowlistRequirement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_increase\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"}],\"name\":\"IncreasedIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxExclusions\",\"type\":\"uint256\"}],\"name\":\"MaxExclusions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"ModifiedTeam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_recycled\",\"type\":\"bool\"}],\"name\":\"NewIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"TokenAllow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"UpdatedDistributor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"UpdatedFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawUnprocessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"allowTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTeam\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"depositIncentiveSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositSplitGauges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numRounds\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositSplitGaugesRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numRounds\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositSplitRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositUnevenSplitGauges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numRounds\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"}],\"name\":\"depositUnevenSplitGaugesRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"depositUnevenSplitGaugesSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_batch\",\"type\":\"uint256\"}],\"name\":\"endRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"excludedVotesReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batch\",\"type\":\"uint256\"}],\"name\":\"finalizeRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"gaugesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inRoundGauges\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"incentives\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recycled\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"incentivesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incentive\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerVote\",\"type\":\"uint256\"}],\"name\":\"increaseIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incentive\",\"type\":\"uint256\"}],\"name\":\"invalidateIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRoundProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExclusions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approval\",\"type\":\"bool\"}],\"name\":\"modifyTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incentive\",\"type\":\"uint256\"}],\"name\":\"recycleUnprocessed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireAllowlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundGauges\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_requireAllowlist\",\"type\":\"bool\"}],\"name\":\"setAllowlistRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxExclusions\",\"type\":\"uint256\"}],\"name\":\"setMaxExclusions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_excluded\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totals\",\"type\":\"uint256[]\"}],\"name\":\"submitExcludedTotals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_gauges\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totals\",\"type\":\"uint256[]\"}],\"name\":\"submitVoteTotals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"updateDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"}],\"name\":\"updateFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"updateFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incentive\",\"type\":\"uint256\"}],\"name\":\"viewIncentive\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recycled\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"excluded\",\"type\":\"address[]\"}],\"internalType\":\"structVotium.Incentive\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"virtualBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votesReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incentive\",\"type\":\"uint256\"}],\"name\":\"withdrawUnprocessed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VotiumVLCVXV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VotiumVLCVXV2MetaData.ABI instead.
var VotiumVLCVXV2ABI = VotiumVLCVXV2MetaData.ABI

// VotiumVLCVXV2 is an auto generated Go binding around an Ethereum contract.
type VotiumVLCVXV2 struct {
	VotiumVLCVXV2Caller     // Read-only binding to the contract
	VotiumVLCVXV2Transactor // Write-only binding to the contract
	VotiumVLCVXV2Filterer   // Log filterer for contract events
}

// VotiumVLCVXV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VotiumVLCVXV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VotiumVLCVXV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotiumVLCVXV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotiumVLCVXV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotiumVLCVXV2Session struct {
	Contract     *VotiumVLCVXV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotiumVLCVXV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotiumVLCVXV2CallerSession struct {
	Contract *VotiumVLCVXV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VotiumVLCVXV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotiumVLCVXV2TransactorSession struct {
	Contract     *VotiumVLCVXV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VotiumVLCVXV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VotiumVLCVXV2Raw struct {
	Contract *VotiumVLCVXV2 // Generic contract binding to access the raw methods on
}

// VotiumVLCVXV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotiumVLCVXV2CallerRaw struct {
	Contract *VotiumVLCVXV2Caller // Generic read-only contract binding to access the raw methods on
}

// VotiumVLCVXV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotiumVLCVXV2TransactorRaw struct {
	Contract *VotiumVLCVXV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVotiumVLCVXV2 creates a new instance of VotiumVLCVXV2, bound to a specific deployed contract.
func NewVotiumVLCVXV2(address common.Address, backend bind.ContractBackend) (*VotiumVLCVXV2, error) {
	contract, err := bindVotiumVLCVXV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2{VotiumVLCVXV2Caller: VotiumVLCVXV2Caller{contract: contract}, VotiumVLCVXV2Transactor: VotiumVLCVXV2Transactor{contract: contract}, VotiumVLCVXV2Filterer: VotiumVLCVXV2Filterer{contract: contract}}, nil
}

// NewVotiumVLCVXV2Caller creates a new read-only instance of VotiumVLCVXV2, bound to a specific deployed contract.
func NewVotiumVLCVXV2Caller(address common.Address, caller bind.ContractCaller) (*VotiumVLCVXV2Caller, error) {
	contract, err := bindVotiumVLCVXV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2Caller{contract: contract}, nil
}

// NewVotiumVLCVXV2Transactor creates a new write-only instance of VotiumVLCVXV2, bound to a specific deployed contract.
func NewVotiumVLCVXV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VotiumVLCVXV2Transactor, error) {
	contract, err := bindVotiumVLCVXV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2Transactor{contract: contract}, nil
}

// NewVotiumVLCVXV2Filterer creates a new log filterer instance of VotiumVLCVXV2, bound to a specific deployed contract.
func NewVotiumVLCVXV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VotiumVLCVXV2Filterer, error) {
	contract, err := bindVotiumVLCVXV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2Filterer{contract: contract}, nil
}

// bindVotiumVLCVXV2 binds a generic wrapper to an already deployed contract.
func bindVotiumVLCVXV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotiumVLCVXV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVLCVXV2 *VotiumVLCVXV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVLCVXV2.Contract.VotiumVLCVXV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVLCVXV2 *VotiumVLCVXV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.VotiumVLCVXV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVLCVXV2 *VotiumVLCVXV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.VotiumVLCVXV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotiumVLCVXV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.contract.Transact(opts, method, params...)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DENOMINATOR() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.DENOMINATOR(&_VotiumVLCVXV2.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) DENOMINATOR() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.DENOMINATOR(&_VotiumVLCVXV2.CallOpts)
}

// ActiveRound is a free data retrieval call binding the contract method 0x11f0281e.
//
// Solidity: function activeRound() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) ActiveRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "activeRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveRound is a free data retrieval call binding the contract method 0x11f0281e.
//
// Solidity: function activeRound() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) ActiveRound() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.ActiveRound(&_VotiumVLCVXV2.CallOpts)
}

// ActiveRound is a free data retrieval call binding the contract method 0x11f0281e.
//
// Solidity: function activeRound() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) ActiveRound() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.ActiveRound(&_VotiumVLCVXV2.CallOpts)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) ApprovedTeam(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "approvedTeam", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.ApprovedTeam(&_VotiumVLCVXV2.CallOpts, arg0)
}

// ApprovedTeam is a free data retrieval call binding the contract method 0x8f3b80b2.
//
// Solidity: function approvedTeam(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) ApprovedTeam(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.ApprovedTeam(&_VotiumVLCVXV2.CallOpts, arg0)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) CurrentEpoch() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.CurrentEpoch(&_VotiumVLCVXV2.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) CurrentEpoch() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.CurrentEpoch(&_VotiumVLCVXV2.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) Distributor() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.Distributor(&_VotiumVLCVXV2.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) Distributor() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.Distributor(&_VotiumVLCVXV2.CallOpts)
}

// ExcludedVotesReceived is a free data retrieval call binding the contract method 0x578ddfd2.
//
// Solidity: function excludedVotesReceived(uint256 , address , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) ExcludedVotesReceived(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "excludedVotesReceived", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExcludedVotesReceived is a free data retrieval call binding the contract method 0x578ddfd2.
//
// Solidity: function excludedVotesReceived(uint256 , address , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) ExcludedVotesReceived(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.ExcludedVotesReceived(&_VotiumVLCVXV2.CallOpts, arg0, arg1, arg2)
}

// ExcludedVotesReceived is a free data retrieval call binding the contract method 0x578ddfd2.
//
// Solidity: function excludedVotesReceived(uint256 , address , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) ExcludedVotesReceived(arg0 *big.Int, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.ExcludedVotesReceived(&_VotiumVLCVXV2.CallOpts, arg0, arg1, arg2)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) FeeAddress() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.FeeAddress(&_VotiumVLCVXV2.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) FeeAddress() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.FeeAddress(&_VotiumVLCVXV2.CallOpts)
}

// GaugesLength is a free data retrieval call binding the contract method 0x23281dfb.
//
// Solidity: function gaugesLength(uint256 _round) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) GaugesLength(opts *bind.CallOpts, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "gaugesLength", _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaugesLength is a free data retrieval call binding the contract method 0x23281dfb.
//
// Solidity: function gaugesLength(uint256 _round) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) GaugesLength(_round *big.Int) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.GaugesLength(&_VotiumVLCVXV2.CallOpts, _round)
}

// GaugesLength is a free data retrieval call binding the contract method 0x23281dfb.
//
// Solidity: function gaugesLength(uint256 _round) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) GaugesLength(_round *big.Int) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.GaugesLength(&_VotiumVLCVXV2.CallOpts, _round)
}

// InRoundGauges is a free data retrieval call binding the contract method 0x00598ef6.
//
// Solidity: function inRoundGauges(uint256 , address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) InRoundGauges(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "inRoundGauges", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRoundGauges is a free data retrieval call binding the contract method 0x00598ef6.
//
// Solidity: function inRoundGauges(uint256 , address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) InRoundGauges(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.InRoundGauges(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// InRoundGauges is a free data retrieval call binding the contract method 0x00598ef6.
//
// Solidity: function inRoundGauges(uint256 , address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) InRoundGauges(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.InRoundGauges(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// Incentives is a free data retrieval call binding the contract method 0x70dfab6c.
//
// Solidity: function incentives(uint256 , address , uint256 ) view returns(address token, uint256 amount, uint256 maxPerVote, uint256 distributed, uint256 recycled, address depositor)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) Incentives(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Token       common.Address
	Amount      *big.Int
	MaxPerVote  *big.Int
	Distributed *big.Int
	Recycled    *big.Int
	Depositor   common.Address
}, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "incentives", arg0, arg1, arg2)

	outstruct := new(struct {
		Token       common.Address
		Amount      *big.Int
		MaxPerVote  *big.Int
		Distributed *big.Int
		Recycled    *big.Int
		Depositor   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxPerVote = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Distributed = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Recycled = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Depositor = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Incentives is a free data retrieval call binding the contract method 0x70dfab6c.
//
// Solidity: function incentives(uint256 , address , uint256 ) view returns(address token, uint256 amount, uint256 maxPerVote, uint256 distributed, uint256 recycled, address depositor)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) Incentives(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Token       common.Address
	Amount      *big.Int
	MaxPerVote  *big.Int
	Distributed *big.Int
	Recycled    *big.Int
	Depositor   common.Address
}, error) {
	return _VotiumVLCVXV2.Contract.Incentives(&_VotiumVLCVXV2.CallOpts, arg0, arg1, arg2)
}

// Incentives is a free data retrieval call binding the contract method 0x70dfab6c.
//
// Solidity: function incentives(uint256 , address , uint256 ) view returns(address token, uint256 amount, uint256 maxPerVote, uint256 distributed, uint256 recycled, address depositor)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) Incentives(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Token       common.Address
	Amount      *big.Int
	MaxPerVote  *big.Int
	Distributed *big.Int
	Recycled    *big.Int
	Depositor   common.Address
}, error) {
	return _VotiumVLCVXV2.Contract.Incentives(&_VotiumVLCVXV2.CallOpts, arg0, arg1, arg2)
}

// IncentivesLength is a free data retrieval call binding the contract method 0xdd07b719.
//
// Solidity: function incentivesLength(uint256 _round, address _gauge) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) IncentivesLength(opts *bind.CallOpts, _round *big.Int, _gauge common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "incentivesLength", _round, _gauge)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncentivesLength is a free data retrieval call binding the contract method 0xdd07b719.
//
// Solidity: function incentivesLength(uint256 _round, address _gauge) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) IncentivesLength(_round *big.Int, _gauge common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.IncentivesLength(&_VotiumVLCVXV2.CallOpts, _round, _gauge)
}

// IncentivesLength is a free data retrieval call binding the contract method 0xdd07b719.
//
// Solidity: function incentivesLength(uint256 _round, address _gauge) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) IncentivesLength(_round *big.Int, _gauge common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.IncentivesLength(&_VotiumVLCVXV2.CallOpts, _round, _gauge)
}

// LastRoundProcessed is a free data retrieval call binding the contract method 0x69e64c11.
//
// Solidity: function lastRoundProcessed() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) LastRoundProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "lastRoundProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRoundProcessed is a free data retrieval call binding the contract method 0x69e64c11.
//
// Solidity: function lastRoundProcessed() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) LastRoundProcessed() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.LastRoundProcessed(&_VotiumVLCVXV2.CallOpts)
}

// LastRoundProcessed is a free data retrieval call binding the contract method 0x69e64c11.
//
// Solidity: function lastRoundProcessed() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) LastRoundProcessed() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.LastRoundProcessed(&_VotiumVLCVXV2.CallOpts)
}

// MaxExclusions is a free data retrieval call binding the contract method 0xac6dbcf1.
//
// Solidity: function maxExclusions() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) MaxExclusions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "maxExclusions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExclusions is a free data retrieval call binding the contract method 0xac6dbcf1.
//
// Solidity: function maxExclusions() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) MaxExclusions() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.MaxExclusions(&_VotiumVLCVXV2.CallOpts)
}

// MaxExclusions is a free data retrieval call binding the contract method 0xac6dbcf1.
//
// Solidity: function maxExclusions() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) MaxExclusions() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.MaxExclusions(&_VotiumVLCVXV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) Owner() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.Owner(&_VotiumVLCVXV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) Owner() (common.Address, error) {
	return _VotiumVLCVXV2.Contract.Owner(&_VotiumVLCVXV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) PlatformFee() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.PlatformFee(&_VotiumVLCVXV2.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) PlatformFee() (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.PlatformFee(&_VotiumVLCVXV2.CallOpts)
}

// RequireAllowlist is a free data retrieval call binding the contract method 0x0c9a94d2.
//
// Solidity: function requireAllowlist() view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) RequireAllowlist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "requireAllowlist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireAllowlist is a free data retrieval call binding the contract method 0x0c9a94d2.
//
// Solidity: function requireAllowlist() view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) RequireAllowlist() (bool, error) {
	return _VotiumVLCVXV2.Contract.RequireAllowlist(&_VotiumVLCVXV2.CallOpts)
}

// RequireAllowlist is a free data retrieval call binding the contract method 0x0c9a94d2.
//
// Solidity: function requireAllowlist() view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) RequireAllowlist() (bool, error) {
	return _VotiumVLCVXV2.Contract.RequireAllowlist(&_VotiumVLCVXV2.CallOpts)
}

// RoundGauges is a free data retrieval call binding the contract method 0x0033ad8b.
//
// Solidity: function roundGauges(uint256 , uint256 ) view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) RoundGauges(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "roundGauges", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoundGauges is a free data retrieval call binding the contract method 0x0033ad8b.
//
// Solidity: function roundGauges(uint256 , uint256 ) view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) RoundGauges(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _VotiumVLCVXV2.Contract.RoundGauges(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// RoundGauges is a free data retrieval call binding the contract method 0x0033ad8b.
//
// Solidity: function roundGauges(uint256 , uint256 ) view returns(address)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) RoundGauges(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _VotiumVLCVXV2.Contract.RoundGauges(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// TokenAllowed is a free data retrieval call binding the contract method 0x55219d5a.
//
// Solidity: function tokenAllowed(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) TokenAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "tokenAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenAllowed is a free data retrieval call binding the contract method 0x55219d5a.
//
// Solidity: function tokenAllowed(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) TokenAllowed(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.TokenAllowed(&_VotiumVLCVXV2.CallOpts, arg0)
}

// TokenAllowed is a free data retrieval call binding the contract method 0x55219d5a.
//
// Solidity: function tokenAllowed(address ) view returns(bool)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) TokenAllowed(arg0 common.Address) (bool, error) {
	return _VotiumVLCVXV2.Contract.TokenAllowed(&_VotiumVLCVXV2.CallOpts, arg0)
}

// ViewIncentive is a free data retrieval call binding the contract method 0xce029b9e.
//
// Solidity: function viewIncentive(uint256 _round, address _gauge, uint256 _incentive) view returns((address,uint256,uint256,uint256,uint256,address,address[]))
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) ViewIncentive(opts *bind.CallOpts, _round *big.Int, _gauge common.Address, _incentive *big.Int) (VotiumIncentive, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "viewIncentive", _round, _gauge, _incentive)

	if err != nil {
		return *new(VotiumIncentive), err
	}

	out0 := *abi.ConvertType(out[0], new(VotiumIncentive)).(*VotiumIncentive)

	return out0, err

}

// ViewIncentive is a free data retrieval call binding the contract method 0xce029b9e.
//
// Solidity: function viewIncentive(uint256 _round, address _gauge, uint256 _incentive) view returns((address,uint256,uint256,uint256,uint256,address,address[]))
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) ViewIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int) (VotiumIncentive, error) {
	return _VotiumVLCVXV2.Contract.ViewIncentive(&_VotiumVLCVXV2.CallOpts, _round, _gauge, _incentive)
}

// ViewIncentive is a free data retrieval call binding the contract method 0xce029b9e.
//
// Solidity: function viewIncentive(uint256 _round, address _gauge, uint256 _incentive) view returns((address,uint256,uint256,uint256,uint256,address,address[]))
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) ViewIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int) (VotiumIncentive, error) {
	return _VotiumVLCVXV2.Contract.ViewIncentive(&_VotiumVLCVXV2.CallOpts, _round, _gauge, _incentive)
}

// VirtualBalance is a free data retrieval call binding the contract method 0xf5c46e41.
//
// Solidity: function virtualBalance(address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) VirtualBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "virtualBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualBalance is a free data retrieval call binding the contract method 0xf5c46e41.
//
// Solidity: function virtualBalance(address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) VirtualBalance(arg0 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.VirtualBalance(&_VotiumVLCVXV2.CallOpts, arg0)
}

// VirtualBalance is a free data retrieval call binding the contract method 0xf5c46e41.
//
// Solidity: function virtualBalance(address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) VirtualBalance(arg0 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.VirtualBalance(&_VotiumVLCVXV2.CallOpts, arg0)
}

// VotesReceived is a free data retrieval call binding the contract method 0x33dd6bed.
//
// Solidity: function votesReceived(uint256 , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Caller) VotesReceived(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotiumVLCVXV2.contract.Call(opts, &out, "votesReceived", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesReceived is a free data retrieval call binding the contract method 0x33dd6bed.
//
// Solidity: function votesReceived(uint256 , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) VotesReceived(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.VotesReceived(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// VotesReceived is a free data retrieval call binding the contract method 0x33dd6bed.
//
// Solidity: function votesReceived(uint256 , address ) view returns(uint256)
func (_VotiumVLCVXV2 *VotiumVLCVXV2CallerSession) VotesReceived(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _VotiumVLCVXV2.Contract.VotesReceived(&_VotiumVLCVXV2.CallOpts, arg0, arg1)
}

// AllowToken is a paid mutator transaction binding the contract method 0x5bc35ae8.
//
// Solidity: function allowToken(address _token, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) AllowToken(opts *bind.TransactOpts, _token common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "allowToken", _token, _allow)
}

// AllowToken is a paid mutator transaction binding the contract method 0x5bc35ae8.
//
// Solidity: function allowToken(address _token, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) AllowToken(_token common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.AllowToken(&_VotiumVLCVXV2.TransactOpts, _token, _allow)
}

// AllowToken is a paid mutator transaction binding the contract method 0x5bc35ae8.
//
// Solidity: function allowToken(address _token, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) AllowToken(_token common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.AllowToken(&_VotiumVLCVXV2.TransactOpts, _token, _allow)
}

// AllowTokens is a paid mutator transaction binding the contract method 0xd7dc8295.
//
// Solidity: function allowTokens(address[] _tokens, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) AllowTokens(opts *bind.TransactOpts, _tokens []common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "allowTokens", _tokens, _allow)
}

// AllowTokens is a paid mutator transaction binding the contract method 0xd7dc8295.
//
// Solidity: function allowTokens(address[] _tokens, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) AllowTokens(_tokens []common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.AllowTokens(&_VotiumVLCVXV2.TransactOpts, _tokens, _allow)
}

// AllowTokens is a paid mutator transaction binding the contract method 0xd7dc8295.
//
// Solidity: function allowTokens(address[] _tokens, bool _allow) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) AllowTokens(_tokens []common.Address, _allow bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.AllowTokens(&_VotiumVLCVXV2.TransactOpts, _tokens, _allow)
}

// DepositIncentive is a paid mutator transaction binding the contract method 0xad2585cd.
//
// Solidity: function depositIncentive(address _token, uint256 _amount, uint256 _round, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositIncentive(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _round *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositIncentive", _token, _amount, _round, _gauge, _maxPerVote, _excluded)
}

// DepositIncentive is a paid mutator transaction binding the contract method 0xad2585cd.
//
// Solidity: function depositIncentive(address _token, uint256 _amount, uint256 _round, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositIncentive(_token common.Address, _amount *big.Int, _round *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositIncentive(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _round, _gauge, _maxPerVote, _excluded)
}

// DepositIncentive is a paid mutator transaction binding the contract method 0xad2585cd.
//
// Solidity: function depositIncentive(address _token, uint256 _amount, uint256 _round, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositIncentive(_token common.Address, _amount *big.Int, _round *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositIncentive(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _round, _gauge, _maxPerVote, _excluded)
}

// DepositIncentiveSimple is a paid mutator transaction binding the contract method 0x355d8e44.
//
// Solidity: function depositIncentiveSimple(address _token, uint256 _amount, address _gauge) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositIncentiveSimple(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositIncentiveSimple", _token, _amount, _gauge)
}

// DepositIncentiveSimple is a paid mutator transaction binding the contract method 0x355d8e44.
//
// Solidity: function depositIncentiveSimple(address _token, uint256 _amount, address _gauge) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositIncentiveSimple(_token common.Address, _amount *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositIncentiveSimple(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _gauge)
}

// DepositIncentiveSimple is a paid mutator transaction binding the contract method 0x355d8e44.
//
// Solidity: function depositIncentiveSimple(address _token, uint256 _amount, address _gauge) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositIncentiveSimple(_token common.Address, _amount *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositIncentiveSimple(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _gauge)
}

// DepositSplitGauges is a paid mutator transaction binding the contract method 0xe1e67fb1.
//
// Solidity: function depositSplitGauges(address _token, uint256 _amount, uint256 _round, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositSplitGauges(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _round *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositSplitGauges", _token, _amount, _round, _gauges, _maxPerVote, _excluded)
}

// DepositSplitGauges is a paid mutator transaction binding the contract method 0xe1e67fb1.
//
// Solidity: function depositSplitGauges(address _token, uint256 _amount, uint256 _round, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositSplitGauges(_token common.Address, _amount *big.Int, _round *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitGauges(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _round, _gauges, _maxPerVote, _excluded)
}

// DepositSplitGauges is a paid mutator transaction binding the contract method 0xe1e67fb1.
//
// Solidity: function depositSplitGauges(address _token, uint256 _amount, uint256 _round, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositSplitGauges(_token common.Address, _amount *big.Int, _round *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitGauges(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _round, _gauges, _maxPerVote, _excluded)
}

// DepositSplitGaugesRounds is a paid mutator transaction binding the contract method 0x999de471.
//
// Solidity: function depositSplitGaugesRounds(address _token, uint256 _amount, uint256 _numRounds, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositSplitGaugesRounds(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _numRounds *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositSplitGaugesRounds", _token, _amount, _numRounds, _gauges, _maxPerVote, _excluded)
}

// DepositSplitGaugesRounds is a paid mutator transaction binding the contract method 0x999de471.
//
// Solidity: function depositSplitGaugesRounds(address _token, uint256 _amount, uint256 _numRounds, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositSplitGaugesRounds(_token common.Address, _amount *big.Int, _numRounds *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitGaugesRounds(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _numRounds, _gauges, _maxPerVote, _excluded)
}

// DepositSplitGaugesRounds is a paid mutator transaction binding the contract method 0x999de471.
//
// Solidity: function depositSplitGaugesRounds(address _token, uint256 _amount, uint256 _numRounds, address[] _gauges, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositSplitGaugesRounds(_token common.Address, _amount *big.Int, _numRounds *big.Int, _gauges []common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitGaugesRounds(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _numRounds, _gauges, _maxPerVote, _excluded)
}

// DepositSplitRounds is a paid mutator transaction binding the contract method 0x0ab0e3ed.
//
// Solidity: function depositSplitRounds(address _token, uint256 _amount, uint256 _numRounds, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositSplitRounds(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _numRounds *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositSplitRounds", _token, _amount, _numRounds, _gauge, _maxPerVote, _excluded)
}

// DepositSplitRounds is a paid mutator transaction binding the contract method 0x0ab0e3ed.
//
// Solidity: function depositSplitRounds(address _token, uint256 _amount, uint256 _numRounds, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositSplitRounds(_token common.Address, _amount *big.Int, _numRounds *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitRounds(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _numRounds, _gauge, _maxPerVote, _excluded)
}

// DepositSplitRounds is a paid mutator transaction binding the contract method 0x0ab0e3ed.
//
// Solidity: function depositSplitRounds(address _token, uint256 _amount, uint256 _numRounds, address _gauge, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositSplitRounds(_token common.Address, _amount *big.Int, _numRounds *big.Int, _gauge common.Address, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositSplitRounds(&_VotiumVLCVXV2.TransactOpts, _token, _amount, _numRounds, _gauge, _maxPerVote, _excluded)
}

// DepositUnevenSplitGauges is a paid mutator transaction binding the contract method 0xdef09779.
//
// Solidity: function depositUnevenSplitGauges(address _token, uint256 _round, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositUnevenSplitGauges(opts *bind.TransactOpts, _token common.Address, _round *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositUnevenSplitGauges", _token, _round, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGauges is a paid mutator transaction binding the contract method 0xdef09779.
//
// Solidity: function depositUnevenSplitGauges(address _token, uint256 _round, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositUnevenSplitGauges(_token common.Address, _round *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGauges(&_VotiumVLCVXV2.TransactOpts, _token, _round, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGauges is a paid mutator transaction binding the contract method 0xdef09779.
//
// Solidity: function depositUnevenSplitGauges(address _token, uint256 _round, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositUnevenSplitGauges(_token common.Address, _round *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGauges(&_VotiumVLCVXV2.TransactOpts, _token, _round, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGaugesRounds is a paid mutator transaction binding the contract method 0x9c355857.
//
// Solidity: function depositUnevenSplitGaugesRounds(address _token, uint256 _numRounds, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositUnevenSplitGaugesRounds(opts *bind.TransactOpts, _token common.Address, _numRounds *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositUnevenSplitGaugesRounds", _token, _numRounds, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGaugesRounds is a paid mutator transaction binding the contract method 0x9c355857.
//
// Solidity: function depositUnevenSplitGaugesRounds(address _token, uint256 _numRounds, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositUnevenSplitGaugesRounds(_token common.Address, _numRounds *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGaugesRounds(&_VotiumVLCVXV2.TransactOpts, _token, _numRounds, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGaugesRounds is a paid mutator transaction binding the contract method 0x9c355857.
//
// Solidity: function depositUnevenSplitGaugesRounds(address _token, uint256 _numRounds, address[] _gauges, uint256[] _amounts, uint256 _maxPerVote, address[] _excluded) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositUnevenSplitGaugesRounds(_token common.Address, _numRounds *big.Int, _gauges []common.Address, _amounts []*big.Int, _maxPerVote *big.Int, _excluded []common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGaugesRounds(&_VotiumVLCVXV2.TransactOpts, _token, _numRounds, _gauges, _amounts, _maxPerVote, _excluded)
}

// DepositUnevenSplitGaugesSimple is a paid mutator transaction binding the contract method 0x68406732.
//
// Solidity: function depositUnevenSplitGaugesSimple(address _token, address[] _gauges, uint256[] _amounts) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) DepositUnevenSplitGaugesSimple(opts *bind.TransactOpts, _token common.Address, _gauges []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "depositUnevenSplitGaugesSimple", _token, _gauges, _amounts)
}

// DepositUnevenSplitGaugesSimple is a paid mutator transaction binding the contract method 0x68406732.
//
// Solidity: function depositUnevenSplitGaugesSimple(address _token, address[] _gauges, uint256[] _amounts) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) DepositUnevenSplitGaugesSimple(_token common.Address, _gauges []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGaugesSimple(&_VotiumVLCVXV2.TransactOpts, _token, _gauges, _amounts)
}

// DepositUnevenSplitGaugesSimple is a paid mutator transaction binding the contract method 0x68406732.
//
// Solidity: function depositUnevenSplitGaugesSimple(address _token, address[] _gauges, uint256[] _amounts) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) DepositUnevenSplitGaugesSimple(_token common.Address, _gauges []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.DepositUnevenSplitGaugesSimple(&_VotiumVLCVXV2.TransactOpts, _token, _gauges, _amounts)
}

// EndRound is a paid mutator transaction binding the contract method 0x89483c95.
//
// Solidity: function endRound(uint256 _round, address[] _gauges, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) EndRound(opts *bind.TransactOpts, _round *big.Int, _gauges []common.Address, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "endRound", _round, _gauges, _batch)
}

// EndRound is a paid mutator transaction binding the contract method 0x89483c95.
//
// Solidity: function endRound(uint256 _round, address[] _gauges, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) EndRound(_round *big.Int, _gauges []common.Address, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.EndRound(&_VotiumVLCVXV2.TransactOpts, _round, _gauges, _batch)
}

// EndRound is a paid mutator transaction binding the contract method 0x89483c95.
//
// Solidity: function endRound(uint256 _round, address[] _gauges, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) EndRound(_round *big.Int, _gauges []common.Address, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.EndRound(&_VotiumVLCVXV2.TransactOpts, _round, _gauges, _batch)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) Execute(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "execute", _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.Execute(&_VotiumVLCVXV2.TransactOpts, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _to, uint256 _value, bytes _data) returns(bool, bytes)
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) Execute(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.Execute(&_VotiumVLCVXV2.TransactOpts, _to, _value, _data)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0xb2073a69.
//
// Solidity: function finalizeRound(uint256 _round, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) FinalizeRound(opts *bind.TransactOpts, _round *big.Int, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "finalizeRound", _round, _batch)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0xb2073a69.
//
// Solidity: function finalizeRound(uint256 _round, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) FinalizeRound(_round *big.Int, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.FinalizeRound(&_VotiumVLCVXV2.TransactOpts, _round, _batch)
}

// FinalizeRound is a paid mutator transaction binding the contract method 0xb2073a69.
//
// Solidity: function finalizeRound(uint256 _round, uint256 _batch) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) FinalizeRound(_round *big.Int, _batch *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.FinalizeRound(&_VotiumVLCVXV2.TransactOpts, _round, _batch)
}

// IncreaseIncentive is a paid mutator transaction binding the contract method 0x67822334.
//
// Solidity: function increaseIncentive(uint256 _round, address _gauge, uint256 _incentive, uint256 _increase, uint256 _maxPerVote) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) IncreaseIncentive(opts *bind.TransactOpts, _round *big.Int, _gauge common.Address, _incentive *big.Int, _increase *big.Int, _maxPerVote *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "increaseIncentive", _round, _gauge, _incentive, _increase, _maxPerVote)
}

// IncreaseIncentive is a paid mutator transaction binding the contract method 0x67822334.
//
// Solidity: function increaseIncentive(uint256 _round, address _gauge, uint256 _incentive, uint256 _increase, uint256 _maxPerVote) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) IncreaseIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int, _increase *big.Int, _maxPerVote *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.IncreaseIncentive(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive, _increase, _maxPerVote)
}

// IncreaseIncentive is a paid mutator transaction binding the contract method 0x67822334.
//
// Solidity: function increaseIncentive(uint256 _round, address _gauge, uint256 _incentive, uint256 _increase, uint256 _maxPerVote) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) IncreaseIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int, _increase *big.Int, _maxPerVote *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.IncreaseIncentive(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive, _increase, _maxPerVote)
}

// InvalidateIncentive is a paid mutator transaction binding the contract method 0xc540f4d7.
//
// Solidity: function invalidateIncentive(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) InvalidateIncentive(opts *bind.TransactOpts, _round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "invalidateIncentive", _round, _gauge, _incentive)
}

// InvalidateIncentive is a paid mutator transaction binding the contract method 0xc540f4d7.
//
// Solidity: function invalidateIncentive(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) InvalidateIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.InvalidateIncentive(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// InvalidateIncentive is a paid mutator transaction binding the contract method 0xc540f4d7.
//
// Solidity: function invalidateIncentive(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) InvalidateIncentive(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.InvalidateIncentive(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) ModifyTeam(opts *bind.TransactOpts, _member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "modifyTeam", _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.ModifyTeam(&_VotiumVLCVXV2.TransactOpts, _member, _approval)
}

// ModifyTeam is a paid mutator transaction binding the contract method 0xac0679a6.
//
// Solidity: function modifyTeam(address _member, bool _approval) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) ModifyTeam(_member common.Address, _approval bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.ModifyTeam(&_VotiumVLCVXV2.TransactOpts, _member, _approval)
}

// RecycleUnprocessed is a paid mutator transaction binding the contract method 0xa9219380.
//
// Solidity: function recycleUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) RecycleUnprocessed(opts *bind.TransactOpts, _round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "recycleUnprocessed", _round, _gauge, _incentive)
}

// RecycleUnprocessed is a paid mutator transaction binding the contract method 0xa9219380.
//
// Solidity: function recycleUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) RecycleUnprocessed(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.RecycleUnprocessed(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// RecycleUnprocessed is a paid mutator transaction binding the contract method 0xa9219380.
//
// Solidity: function recycleUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) RecycleUnprocessed(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.RecycleUnprocessed(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.RenounceOwnership(&_VotiumVLCVXV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.RenounceOwnership(&_VotiumVLCVXV2.TransactOpts)
}

// SetAllowlistRequired is a paid mutator transaction binding the contract method 0x3cf17931.
//
// Solidity: function setAllowlistRequired(bool _requireAllowlist) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) SetAllowlistRequired(opts *bind.TransactOpts, _requireAllowlist bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "setAllowlistRequired", _requireAllowlist)
}

// SetAllowlistRequired is a paid mutator transaction binding the contract method 0x3cf17931.
//
// Solidity: function setAllowlistRequired(bool _requireAllowlist) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) SetAllowlistRequired(_requireAllowlist bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SetAllowlistRequired(&_VotiumVLCVXV2.TransactOpts, _requireAllowlist)
}

// SetAllowlistRequired is a paid mutator transaction binding the contract method 0x3cf17931.
//
// Solidity: function setAllowlistRequired(bool _requireAllowlist) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) SetAllowlistRequired(_requireAllowlist bool) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SetAllowlistRequired(&_VotiumVLCVXV2.TransactOpts, _requireAllowlist)
}

// SetMaxExclusions is a paid mutator transaction binding the contract method 0x092d6082.
//
// Solidity: function setMaxExclusions(uint256 _maxExclusions) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) SetMaxExclusions(opts *bind.TransactOpts, _maxExclusions *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "setMaxExclusions", _maxExclusions)
}

// SetMaxExclusions is a paid mutator transaction binding the contract method 0x092d6082.
//
// Solidity: function setMaxExclusions(uint256 _maxExclusions) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) SetMaxExclusions(_maxExclusions *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SetMaxExclusions(&_VotiumVLCVXV2.TransactOpts, _maxExclusions)
}

// SetMaxExclusions is a paid mutator transaction binding the contract method 0x092d6082.
//
// Solidity: function setMaxExclusions(uint256 _maxExclusions) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) SetMaxExclusions(_maxExclusions *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SetMaxExclusions(&_VotiumVLCVXV2.TransactOpts, _maxExclusions)
}

// SubmitExcludedTotals is a paid mutator transaction binding the contract method 0x67131a65.
//
// Solidity: function submitExcludedTotals(uint256 _round, address _gauge, address[] _excluded, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) SubmitExcludedTotals(opts *bind.TransactOpts, _round *big.Int, _gauge common.Address, _excluded []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "submitExcludedTotals", _round, _gauge, _excluded, _totals)
}

// SubmitExcludedTotals is a paid mutator transaction binding the contract method 0x67131a65.
//
// Solidity: function submitExcludedTotals(uint256 _round, address _gauge, address[] _excluded, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) SubmitExcludedTotals(_round *big.Int, _gauge common.Address, _excluded []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SubmitExcludedTotals(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _excluded, _totals)
}

// SubmitExcludedTotals is a paid mutator transaction binding the contract method 0x67131a65.
//
// Solidity: function submitExcludedTotals(uint256 _round, address _gauge, address[] _excluded, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) SubmitExcludedTotals(_round *big.Int, _gauge common.Address, _excluded []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SubmitExcludedTotals(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _excluded, _totals)
}

// SubmitVoteTotals is a paid mutator transaction binding the contract method 0x4bbbce6c.
//
// Solidity: function submitVoteTotals(uint256 _round, address[] _gauges, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) SubmitVoteTotals(opts *bind.TransactOpts, _round *big.Int, _gauges []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "submitVoteTotals", _round, _gauges, _totals)
}

// SubmitVoteTotals is a paid mutator transaction binding the contract method 0x4bbbce6c.
//
// Solidity: function submitVoteTotals(uint256 _round, address[] _gauges, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) SubmitVoteTotals(_round *big.Int, _gauges []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SubmitVoteTotals(&_VotiumVLCVXV2.TransactOpts, _round, _gauges, _totals)
}

// SubmitVoteTotals is a paid mutator transaction binding the contract method 0x4bbbce6c.
//
// Solidity: function submitVoteTotals(uint256 _round, address[] _gauges, uint256[] _totals) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) SubmitVoteTotals(_round *big.Int, _gauges []common.Address, _totals []*big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.SubmitVoteTotals(&_VotiumVLCVXV2.TransactOpts, _round, _gauges, _totals)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.TransferOwnership(&_VotiumVLCVXV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.TransferOwnership(&_VotiumVLCVXV2.TransactOpts, newOwner)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) UpdateDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "updateDistributor", _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) UpdateDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateDistributor(&_VotiumVLCVXV2.TransactOpts, _distributor)
}

// UpdateDistributor is a paid mutator transaction binding the contract method 0xbc30a618.
//
// Solidity: function updateDistributor(address _distributor) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) UpdateDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateDistributor(&_VotiumVLCVXV2.TransactOpts, _distributor)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) UpdateFeeAddress(opts *bind.TransactOpts, _feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "updateFeeAddress", _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateFeeAddress(&_VotiumVLCVXV2.TransactOpts, _feeAddress)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address _feeAddress) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) UpdateFeeAddress(_feeAddress common.Address) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateFeeAddress(&_VotiumVLCVXV2.TransactOpts, _feeAddress)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) UpdateFeeAmount(opts *bind.TransactOpts, _feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "updateFeeAmount", _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateFeeAmount(&_VotiumVLCVXV2.TransactOpts, _feeAmount)
}

// UpdateFeeAmount is a paid mutator transaction binding the contract method 0x9ea55bb0.
//
// Solidity: function updateFeeAmount(uint256 _feeAmount) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) UpdateFeeAmount(_feeAmount *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.UpdateFeeAmount(&_VotiumVLCVXV2.TransactOpts, _feeAmount)
}

// WithdrawUnprocessed is a paid mutator transaction binding the contract method 0x94a6492f.
//
// Solidity: function withdrawUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Transactor) WithdrawUnprocessed(opts *bind.TransactOpts, _round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.contract.Transact(opts, "withdrawUnprocessed", _round, _gauge, _incentive)
}

// WithdrawUnprocessed is a paid mutator transaction binding the contract method 0x94a6492f.
//
// Solidity: function withdrawUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2Session) WithdrawUnprocessed(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.WithdrawUnprocessed(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// WithdrawUnprocessed is a paid mutator transaction binding the contract method 0x94a6492f.
//
// Solidity: function withdrawUnprocessed(uint256 _round, address _gauge, uint256 _incentive) returns()
func (_VotiumVLCVXV2 *VotiumVLCVXV2TransactorSession) WithdrawUnprocessed(_round *big.Int, _gauge common.Address, _incentive *big.Int) (*types.Transaction, error) {
	return _VotiumVLCVXV2.Contract.WithdrawUnprocessed(&_VotiumVLCVXV2.TransactOpts, _round, _gauge, _incentive)
}

// VotiumVLCVXV2AllowlistRequirementIterator is returned from FilterAllowlistRequirement and is used to iterate over the raw logs and unpacked data for AllowlistRequirement events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2AllowlistRequirementIterator struct {
	Event *VotiumVLCVXV2AllowlistRequirement // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2AllowlistRequirementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2AllowlistRequirement)
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
		it.Event = new(VotiumVLCVXV2AllowlistRequirement)
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
func (it *VotiumVLCVXV2AllowlistRequirementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2AllowlistRequirementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2AllowlistRequirement represents a AllowlistRequirement event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2AllowlistRequirement struct {
	RequireAllowlist bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAllowlistRequirement is a free log retrieval operation binding the contract event 0x5f58852311824c54c2949edc2466d7565000ce83581d96485eba68871f48c4b3.
//
// Solidity: event AllowlistRequirement(bool _requireAllowlist)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterAllowlistRequirement(opts *bind.FilterOpts) (*VotiumVLCVXV2AllowlistRequirementIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "AllowlistRequirement")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2AllowlistRequirementIterator{contract: _VotiumVLCVXV2.contract, event: "AllowlistRequirement", logs: logs, sub: sub}, nil
}

// WatchAllowlistRequirement is a free log subscription operation binding the contract event 0x5f58852311824c54c2949edc2466d7565000ce83581d96485eba68871f48c4b3.
//
// Solidity: event AllowlistRequirement(bool _requireAllowlist)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchAllowlistRequirement(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2AllowlistRequirement) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "AllowlistRequirement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2AllowlistRequirement)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "AllowlistRequirement", log); err != nil {
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

// ParseAllowlistRequirement is a log parse operation binding the contract event 0x5f58852311824c54c2949edc2466d7565000ce83581d96485eba68871f48c4b3.
//
// Solidity: event AllowlistRequirement(bool _requireAllowlist)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseAllowlistRequirement(log types.Log) (*VotiumVLCVXV2AllowlistRequirement, error) {
	event := new(VotiumVLCVXV2AllowlistRequirement)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "AllowlistRequirement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2IncreasedIncentiveIterator is returned from FilterIncreasedIncentive and is used to iterate over the raw logs and unpacked data for IncreasedIncentive events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2IncreasedIncentiveIterator struct {
	Event *VotiumVLCVXV2IncreasedIncentive // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2IncreasedIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2IncreasedIncentive)
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
		it.Event = new(VotiumVLCVXV2IncreasedIncentive)
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
func (it *VotiumVLCVXV2IncreasedIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2IncreasedIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2IncreasedIncentive represents a IncreasedIncentive event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2IncreasedIncentive struct {
	Index      *big.Int
	Token      common.Address
	Total      *big.Int
	Increase   *big.Int
	Round      *big.Int
	Gauge      common.Address
	MaxPerVote *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIncreasedIncentive is a free log retrieval operation binding the contract event 0xf7992095989ff6fb8484c32670972fb7c67e943107c3f158013e2b7fb96d1971.
//
// Solidity: event IncreasedIncentive(uint256 _index, address _token, uint256 _total, uint256 _increase, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterIncreasedIncentive(opts *bind.FilterOpts, _round []*big.Int, _gauge []common.Address) (*VotiumVLCVXV2IncreasedIncentiveIterator, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "IncreasedIncentive", _roundRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2IncreasedIncentiveIterator{contract: _VotiumVLCVXV2.contract, event: "IncreasedIncentive", logs: logs, sub: sub}, nil
}

// WatchIncreasedIncentive is a free log subscription operation binding the contract event 0xf7992095989ff6fb8484c32670972fb7c67e943107c3f158013e2b7fb96d1971.
//
// Solidity: event IncreasedIncentive(uint256 _index, address _token, uint256 _total, uint256 _increase, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchIncreasedIncentive(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2IncreasedIncentive, _round []*big.Int, _gauge []common.Address) (event.Subscription, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "IncreasedIncentive", _roundRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2IncreasedIncentive)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "IncreasedIncentive", log); err != nil {
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

// ParseIncreasedIncentive is a log parse operation binding the contract event 0xf7992095989ff6fb8484c32670972fb7c67e943107c3f158013e2b7fb96d1971.
//
// Solidity: event IncreasedIncentive(uint256 _index, address _token, uint256 _total, uint256 _increase, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseIncreasedIncentive(log types.Log) (*VotiumVLCVXV2IncreasedIncentive, error) {
	event := new(VotiumVLCVXV2IncreasedIncentive)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "IncreasedIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2MaxExclusionsIterator is returned from FilterMaxExclusions and is used to iterate over the raw logs and unpacked data for MaxExclusions events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2MaxExclusionsIterator struct {
	Event *VotiumVLCVXV2MaxExclusions // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2MaxExclusionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2MaxExclusions)
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
		it.Event = new(VotiumVLCVXV2MaxExclusions)
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
func (it *VotiumVLCVXV2MaxExclusionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2MaxExclusionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2MaxExclusions represents a MaxExclusions event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2MaxExclusions struct {
	MaxExclusions *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxExclusions is a free log retrieval operation binding the contract event 0x55af27e81963f03ec952ceca1a8ab02e80218c9db851cfdfade468d98eb12c78.
//
// Solidity: event MaxExclusions(uint256 _maxExclusions)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterMaxExclusions(opts *bind.FilterOpts) (*VotiumVLCVXV2MaxExclusionsIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "MaxExclusions")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2MaxExclusionsIterator{contract: _VotiumVLCVXV2.contract, event: "MaxExclusions", logs: logs, sub: sub}, nil
}

// WatchMaxExclusions is a free log subscription operation binding the contract event 0x55af27e81963f03ec952ceca1a8ab02e80218c9db851cfdfade468d98eb12c78.
//
// Solidity: event MaxExclusions(uint256 _maxExclusions)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchMaxExclusions(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2MaxExclusions) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "MaxExclusions")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2MaxExclusions)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "MaxExclusions", log); err != nil {
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

// ParseMaxExclusions is a log parse operation binding the contract event 0x55af27e81963f03ec952ceca1a8ab02e80218c9db851cfdfade468d98eb12c78.
//
// Solidity: event MaxExclusions(uint256 _maxExclusions)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseMaxExclusions(log types.Log) (*VotiumVLCVXV2MaxExclusions, error) {
	event := new(VotiumVLCVXV2MaxExclusions)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "MaxExclusions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2ModifiedTeamIterator is returned from FilterModifiedTeam and is used to iterate over the raw logs and unpacked data for ModifiedTeam events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2ModifiedTeamIterator struct {
	Event *VotiumVLCVXV2ModifiedTeam // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2ModifiedTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2ModifiedTeam)
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
		it.Event = new(VotiumVLCVXV2ModifiedTeam)
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
func (it *VotiumVLCVXV2ModifiedTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2ModifiedTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2ModifiedTeam represents a ModifiedTeam event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2ModifiedTeam struct {
	Member   common.Address
	Approval bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifiedTeam is a free log retrieval operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterModifiedTeam(opts *bind.FilterOpts) (*VotiumVLCVXV2ModifiedTeamIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2ModifiedTeamIterator{contract: _VotiumVLCVXV2.contract, event: "ModifiedTeam", logs: logs, sub: sub}, nil
}

// WatchModifiedTeam is a free log subscription operation binding the contract event 0xae0b47afe292832708082f706affd0f37c4a556bd6dbeafe9b6a8562251c5a40.
//
// Solidity: event ModifiedTeam(address _member, bool _approval)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchModifiedTeam(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2ModifiedTeam) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "ModifiedTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2ModifiedTeam)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
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
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseModifiedTeam(log types.Log) (*VotiumVLCVXV2ModifiedTeam, error) {
	event := new(VotiumVLCVXV2ModifiedTeam)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "ModifiedTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2NewIncentiveIterator is returned from FilterNewIncentive and is used to iterate over the raw logs and unpacked data for NewIncentive events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2NewIncentiveIterator struct {
	Event *VotiumVLCVXV2NewIncentive // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2NewIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2NewIncentive)
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
		it.Event = new(VotiumVLCVXV2NewIncentive)
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
func (it *VotiumVLCVXV2NewIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2NewIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2NewIncentive represents a NewIncentive event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2NewIncentive struct {
	Index      *big.Int
	Token      common.Address
	Amount     *big.Int
	Round      *big.Int
	Gauge      common.Address
	MaxPerVote *big.Int
	Excluded   []common.Address
	Depositor  common.Address
	Recycled   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewIncentive is a free log retrieval operation binding the contract event 0x7c0c0ef7f1ccead819631ed9c10b0728e76274ee5572b53716ea96e7ec735ffa.
//
// Solidity: event NewIncentive(uint256 _index, address _token, uint256 _amount, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote, address[] _excluded, address indexed _depositor, bool _recycled)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterNewIncentive(opts *bind.FilterOpts, _round []*big.Int, _gauge []common.Address, _depositor []common.Address) (*VotiumVLCVXV2NewIncentiveIterator, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "NewIncentive", _roundRule, _gaugeRule, _depositorRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2NewIncentiveIterator{contract: _VotiumVLCVXV2.contract, event: "NewIncentive", logs: logs, sub: sub}, nil
}

// WatchNewIncentive is a free log subscription operation binding the contract event 0x7c0c0ef7f1ccead819631ed9c10b0728e76274ee5572b53716ea96e7ec735ffa.
//
// Solidity: event NewIncentive(uint256 _index, address _token, uint256 _amount, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote, address[] _excluded, address indexed _depositor, bool _recycled)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchNewIncentive(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2NewIncentive, _round []*big.Int, _gauge []common.Address, _depositor []common.Address) (event.Subscription, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "NewIncentive", _roundRule, _gaugeRule, _depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2NewIncentive)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "NewIncentive", log); err != nil {
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

// ParseNewIncentive is a log parse operation binding the contract event 0x7c0c0ef7f1ccead819631ed9c10b0728e76274ee5572b53716ea96e7ec735ffa.
//
// Solidity: event NewIncentive(uint256 _index, address _token, uint256 _amount, uint256 indexed _round, address indexed _gauge, uint256 _maxPerVote, address[] _excluded, address indexed _depositor, bool _recycled)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseNewIncentive(log types.Log) (*VotiumVLCVXV2NewIncentive, error) {
	event := new(VotiumVLCVXV2NewIncentive)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "NewIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2OwnershipTransferredIterator struct {
	Event *VotiumVLCVXV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2OwnershipTransferred)
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
		it.Event = new(VotiumVLCVXV2OwnershipTransferred)
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
func (it *VotiumVLCVXV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2OwnershipTransferred represents a OwnershipTransferred event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotiumVLCVXV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2OwnershipTransferredIterator{contract: _VotiumVLCVXV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2OwnershipTransferred)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseOwnershipTransferred(log types.Log) (*VotiumVLCVXV2OwnershipTransferred, error) {
	event := new(VotiumVLCVXV2OwnershipTransferred)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2TokenAllowIterator is returned from FilterTokenAllow and is used to iterate over the raw logs and unpacked data for TokenAllow events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2TokenAllowIterator struct {
	Event *VotiumVLCVXV2TokenAllow // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2TokenAllowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2TokenAllow)
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
		it.Event = new(VotiumVLCVXV2TokenAllow)
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
func (it *VotiumVLCVXV2TokenAllowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2TokenAllowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2TokenAllow represents a TokenAllow event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2TokenAllow struct {
	Token common.Address
	Allow bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAllow is a free log retrieval operation binding the contract event 0xa270ddd235c544f3a45444196aa37996d1fbb15fce12c0a63d06b8a984c49961.
//
// Solidity: event TokenAllow(address _token, bool _allow)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterTokenAllow(opts *bind.FilterOpts) (*VotiumVLCVXV2TokenAllowIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "TokenAllow")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2TokenAllowIterator{contract: _VotiumVLCVXV2.contract, event: "TokenAllow", logs: logs, sub: sub}, nil
}

// WatchTokenAllow is a free log subscription operation binding the contract event 0xa270ddd235c544f3a45444196aa37996d1fbb15fce12c0a63d06b8a984c49961.
//
// Solidity: event TokenAllow(address _token, bool _allow)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchTokenAllow(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2TokenAllow) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "TokenAllow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2TokenAllow)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "TokenAllow", log); err != nil {
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

// ParseTokenAllow is a log parse operation binding the contract event 0xa270ddd235c544f3a45444196aa37996d1fbb15fce12c0a63d06b8a984c49961.
//
// Solidity: event TokenAllow(address _token, bool _allow)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseTokenAllow(log types.Log) (*VotiumVLCVXV2TokenAllow, error) {
	event := new(VotiumVLCVXV2TokenAllow)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "TokenAllow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2UpdatedDistributorIterator is returned from FilterUpdatedDistributor and is used to iterate over the raw logs and unpacked data for UpdatedDistributor events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2UpdatedDistributorIterator struct {
	Event *VotiumVLCVXV2UpdatedDistributor // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2UpdatedDistributorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2UpdatedDistributor)
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
		it.Event = new(VotiumVLCVXV2UpdatedDistributor)
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
func (it *VotiumVLCVXV2UpdatedDistributorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2UpdatedDistributorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2UpdatedDistributor represents a UpdatedDistributor event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2UpdatedDistributor struct {
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDistributor is a free log retrieval operation binding the contract event 0x9941240ae138e71bb3f963c4f6fd96e99fb386a8e63563343f9294902bdbf744.
//
// Solidity: event UpdatedDistributor(address _distributor)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterUpdatedDistributor(opts *bind.FilterOpts) (*VotiumVLCVXV2UpdatedDistributorIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "UpdatedDistributor")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2UpdatedDistributorIterator{contract: _VotiumVLCVXV2.contract, event: "UpdatedDistributor", logs: logs, sub: sub}, nil
}

// WatchUpdatedDistributor is a free log subscription operation binding the contract event 0x9941240ae138e71bb3f963c4f6fd96e99fb386a8e63563343f9294902bdbf744.
//
// Solidity: event UpdatedDistributor(address _distributor)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchUpdatedDistributor(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2UpdatedDistributor) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "UpdatedDistributor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2UpdatedDistributor)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
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
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseUpdatedDistributor(log types.Log) (*VotiumVLCVXV2UpdatedDistributor, error) {
	event := new(VotiumVLCVXV2UpdatedDistributor)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "UpdatedDistributor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2UpdatedFeeIterator is returned from FilterUpdatedFee and is used to iterate over the raw logs and unpacked data for UpdatedFee events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2UpdatedFeeIterator struct {
	Event *VotiumVLCVXV2UpdatedFee // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2UpdatedFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2UpdatedFee)
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
		it.Event = new(VotiumVLCVXV2UpdatedFee)
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
func (it *VotiumVLCVXV2UpdatedFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2UpdatedFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2UpdatedFee represents a UpdatedFee event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2UpdatedFee struct {
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFee is a free log retrieval operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterUpdatedFee(opts *bind.FilterOpts) (*VotiumVLCVXV2UpdatedFeeIterator, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2UpdatedFeeIterator{contract: _VotiumVLCVXV2.contract, event: "UpdatedFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedFee is a free log subscription operation binding the contract event 0x545f541f608330014315921189568bf5b2266925f757d74e5ad89ae1d2d6438c.
//
// Solidity: event UpdatedFee(uint256 _feeAmount)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchUpdatedFee(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2UpdatedFee) (event.Subscription, error) {

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "UpdatedFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2UpdatedFee)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
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
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseUpdatedFee(log types.Log) (*VotiumVLCVXV2UpdatedFee, error) {
	event := new(VotiumVLCVXV2UpdatedFee)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "UpdatedFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotiumVLCVXV2WithdrawUnprocessedIterator is returned from FilterWithdrawUnprocessed and is used to iterate over the raw logs and unpacked data for WithdrawUnprocessed events raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2WithdrawUnprocessedIterator struct {
	Event *VotiumVLCVXV2WithdrawUnprocessed // Event containing the contract specifics and raw log

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
func (it *VotiumVLCVXV2WithdrawUnprocessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotiumVLCVXV2WithdrawUnprocessed)
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
		it.Event = new(VotiumVLCVXV2WithdrawUnprocessed)
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
func (it *VotiumVLCVXV2WithdrawUnprocessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotiumVLCVXV2WithdrawUnprocessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotiumVLCVXV2WithdrawUnprocessed represents a WithdrawUnprocessed event raised by the VotiumVLCVXV2 contract.
type VotiumVLCVXV2WithdrawUnprocessed struct {
	Index  *big.Int
	Round  *big.Int
	Gauge  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUnprocessed is a free log retrieval operation binding the contract event 0xa999dfda6398b45ee80518d4d83252ab0f6eaccfe3a24df823fd71df9fb14cc5.
//
// Solidity: event WithdrawUnprocessed(uint256 _index, uint256 indexed _round, address indexed _gauge, uint256 _amount)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) FilterWithdrawUnprocessed(opts *bind.FilterOpts, _round []*big.Int, _gauge []common.Address) (*VotiumVLCVXV2WithdrawUnprocessedIterator, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.FilterLogs(opts, "WithdrawUnprocessed", _roundRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return &VotiumVLCVXV2WithdrawUnprocessedIterator{contract: _VotiumVLCVXV2.contract, event: "WithdrawUnprocessed", logs: logs, sub: sub}, nil
}

// WatchWithdrawUnprocessed is a free log subscription operation binding the contract event 0xa999dfda6398b45ee80518d4d83252ab0f6eaccfe3a24df823fd71df9fb14cc5.
//
// Solidity: event WithdrawUnprocessed(uint256 _index, uint256 indexed _round, address indexed _gauge, uint256 _amount)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) WatchWithdrawUnprocessed(opts *bind.WatchOpts, sink chan<- *VotiumVLCVXV2WithdrawUnprocessed, _round []*big.Int, _gauge []common.Address) (event.Subscription, error) {

	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}
	var _gaugeRule []interface{}
	for _, _gaugeItem := range _gauge {
		_gaugeRule = append(_gaugeRule, _gaugeItem)
	}

	logs, sub, err := _VotiumVLCVXV2.contract.WatchLogs(opts, "WithdrawUnprocessed", _roundRule, _gaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotiumVLCVXV2WithdrawUnprocessed)
				if err := _VotiumVLCVXV2.contract.UnpackLog(event, "WithdrawUnprocessed", log); err != nil {
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

// ParseWithdrawUnprocessed is a log parse operation binding the contract event 0xa999dfda6398b45ee80518d4d83252ab0f6eaccfe3a24df823fd71df9fb14cc5.
//
// Solidity: event WithdrawUnprocessed(uint256 _index, uint256 indexed _round, address indexed _gauge, uint256 _amount)
func (_VotiumVLCVXV2 *VotiumVLCVXV2Filterer) ParseWithdrawUnprocessed(log types.Log) (*VotiumVLCVXV2WithdrawUnprocessed, error) {
	event := new(VotiumVLCVXV2WithdrawUnprocessed)
	if err := _VotiumVLCVXV2.contract.UnpackLog(event, "WithdrawUnprocessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
