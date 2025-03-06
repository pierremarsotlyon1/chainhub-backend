// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vmRemoteManager

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

// CampaignRemoteManagerCampaignClosingParams is an auto generated low-level Go binding around an user-defined struct.
type CampaignRemoteManagerCampaignClosingParams struct {
	CampaignId *big.Int
}

// CampaignRemoteManagerCampaignCreationParams is an auto generated low-level Go binding around an user-defined struct.
type CampaignRemoteManagerCampaignCreationParams struct {
	ChainId           *big.Int
	Gauge             common.Address
	Manager           common.Address
	RewardToken       common.Address
	NumberOfPeriods   uint8
	MaxRewardPerVote  *big.Int
	TotalRewardAmount *big.Int
	Addresses         []common.Address
	Hook              common.Address
	IsWhitelist       bool
}

// CampaignRemoteManagerCampaignManagementParams is an auto generated low-level Go binding around an user-defined struct.
type CampaignRemoteManagerCampaignManagementParams struct {
	CampaignId        *big.Int
	RewardToken       common.Address
	NumberOfPeriods   uint8
	TotalRewardAmount *big.Int
	MaxRewardPerVote  *big.Int
}

// VmRemoteManagerMetaData contains all meta data concerning the VmRemoteManager contract.
var VmRemoteManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_laPoste\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidActionType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCampaignManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLaPoste\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"indexed\":true,\"internalType\":\"structCampaignRemoteManager.CampaignClosingParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"CampaignClosingPayloadSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"indexed\":true,\"internalType\":\"structCampaignRemoteManager.CampaignCreationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"CampaignCreationPayloadSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"indexed\":true,\"internalType\":\"structCampaignRemoteManager.CampaignManagementParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"CampaignManagementPayloadSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"PlatformWhitelistUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LA_POSTE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"}],\"internalType\":\"structCampaignRemoteManager.CampaignClosingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votemarket\",\"type\":\"address\"}],\"name\":\"closeCampaign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhitelist\",\"type\":\"bool\"}],\"internalType\":\"structCampaignRemoteManager.CampaignCreationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votemarket\",\"type\":\"address\"}],\"name\":\"createCampaign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"numberOfPeriods\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRewardPerVote\",\"type\":\"uint256\"}],\"internalType\":\"structCampaignRemoteManager.CampaignManagementParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votemarket\",\"type\":\"address\"}],\"name\":\"manageCampaign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setPlatformWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedPlatforms\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VmRemoteManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VmRemoteManagerMetaData.ABI instead.
var VmRemoteManagerABI = VmRemoteManagerMetaData.ABI

// VmRemoteManager is an auto generated Go binding around an Ethereum contract.
type VmRemoteManager struct {
	VmRemoteManagerCaller     // Read-only binding to the contract
	VmRemoteManagerTransactor // Write-only binding to the contract
	VmRemoteManagerFilterer   // Log filterer for contract events
}

// VmRemoteManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmRemoteManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmRemoteManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmRemoteManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmRemoteManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmRemoteManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmRemoteManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmRemoteManagerSession struct {
	Contract     *VmRemoteManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmRemoteManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmRemoteManagerCallerSession struct {
	Contract *VmRemoteManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VmRemoteManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmRemoteManagerTransactorSession struct {
	Contract     *VmRemoteManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VmRemoteManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmRemoteManagerRaw struct {
	Contract *VmRemoteManager // Generic contract binding to access the raw methods on
}

// VmRemoteManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmRemoteManagerCallerRaw struct {
	Contract *VmRemoteManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VmRemoteManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmRemoteManagerTransactorRaw struct {
	Contract *VmRemoteManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVmRemoteManager creates a new instance of VmRemoteManager, bound to a specific deployed contract.
func NewVmRemoteManager(address common.Address, backend bind.ContractBackend) (*VmRemoteManager, error) {
	contract, err := bindVmRemoteManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManager{VmRemoteManagerCaller: VmRemoteManagerCaller{contract: contract}, VmRemoteManagerTransactor: VmRemoteManagerTransactor{contract: contract}, VmRemoteManagerFilterer: VmRemoteManagerFilterer{contract: contract}}, nil
}

// NewVmRemoteManagerCaller creates a new read-only instance of VmRemoteManager, bound to a specific deployed contract.
func NewVmRemoteManagerCaller(address common.Address, caller bind.ContractCaller) (*VmRemoteManagerCaller, error) {
	contract, err := bindVmRemoteManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerCaller{contract: contract}, nil
}

// NewVmRemoteManagerTransactor creates a new write-only instance of VmRemoteManager, bound to a specific deployed contract.
func NewVmRemoteManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VmRemoteManagerTransactor, error) {
	contract, err := bindVmRemoteManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerTransactor{contract: contract}, nil
}

// NewVmRemoteManagerFilterer creates a new log filterer instance of VmRemoteManager, bound to a specific deployed contract.
func NewVmRemoteManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VmRemoteManagerFilterer, error) {
	contract, err := bindVmRemoteManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerFilterer{contract: contract}, nil
}

// bindVmRemoteManager binds a generic wrapper to an already deployed contract.
func bindVmRemoteManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VmRemoteManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmRemoteManager *VmRemoteManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmRemoteManager.Contract.VmRemoteManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmRemoteManager *VmRemoteManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.VmRemoteManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmRemoteManager *VmRemoteManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.VmRemoteManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmRemoteManager *VmRemoteManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmRemoteManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmRemoteManager *VmRemoteManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmRemoteManager *VmRemoteManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.contract.Transact(opts, method, params...)
}

// LAPOSTE is a free data retrieval call binding the contract method 0xb8232554.
//
// Solidity: function LA_POSTE() view returns(address)
func (_VmRemoteManager *VmRemoteManagerCaller) LAPOSTE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmRemoteManager.contract.Call(opts, &out, "LA_POSTE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LAPOSTE is a free data retrieval call binding the contract method 0xb8232554.
//
// Solidity: function LA_POSTE() view returns(address)
func (_VmRemoteManager *VmRemoteManagerSession) LAPOSTE() (common.Address, error) {
	return _VmRemoteManager.Contract.LAPOSTE(&_VmRemoteManager.CallOpts)
}

// LAPOSTE is a free data retrieval call binding the contract method 0xb8232554.
//
// Solidity: function LA_POSTE() view returns(address)
func (_VmRemoteManager *VmRemoteManagerCallerSession) LAPOSTE() (common.Address, error) {
	return _VmRemoteManager.Contract.LAPOSTE(&_VmRemoteManager.CallOpts)
}

// TOKENFACTORY is a free data retrieval call binding the contract method 0x5a58fe4b.
//
// Solidity: function TOKEN_FACTORY() view returns(address)
func (_VmRemoteManager *VmRemoteManagerCaller) TOKENFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmRemoteManager.contract.Call(opts, &out, "TOKEN_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENFACTORY is a free data retrieval call binding the contract method 0x5a58fe4b.
//
// Solidity: function TOKEN_FACTORY() view returns(address)
func (_VmRemoteManager *VmRemoteManagerSession) TOKENFACTORY() (common.Address, error) {
	return _VmRemoteManager.Contract.TOKENFACTORY(&_VmRemoteManager.CallOpts)
}

// TOKENFACTORY is a free data retrieval call binding the contract method 0x5a58fe4b.
//
// Solidity: function TOKEN_FACTORY() view returns(address)
func (_VmRemoteManager *VmRemoteManagerCallerSession) TOKENFACTORY() (common.Address, error) {
	return _VmRemoteManager.Contract.TOKENFACTORY(&_VmRemoteManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_VmRemoteManager *VmRemoteManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmRemoteManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_VmRemoteManager *VmRemoteManagerSession) Owner() (common.Address, error) {
	return _VmRemoteManager.Contract.Owner(&_VmRemoteManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_VmRemoteManager *VmRemoteManagerCallerSession) Owner() (common.Address, error) {
	return _VmRemoteManager.Contract.Owner(&_VmRemoteManager.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_VmRemoteManager *VmRemoteManagerCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VmRemoteManager.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_VmRemoteManager *VmRemoteManagerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _VmRemoteManager.Contract.OwnershipHandoverExpiresAt(&_VmRemoteManager.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_VmRemoteManager *VmRemoteManagerCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _VmRemoteManager.Contract.OwnershipHandoverExpiresAt(&_VmRemoteManager.CallOpts, pendingOwner)
}

// WhitelistedPlatforms is a free data retrieval call binding the contract method 0x8924ec96.
//
// Solidity: function whitelistedPlatforms(address ) view returns(bool)
func (_VmRemoteManager *VmRemoteManagerCaller) WhitelistedPlatforms(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VmRemoteManager.contract.Call(opts, &out, "whitelistedPlatforms", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedPlatforms is a free data retrieval call binding the contract method 0x8924ec96.
//
// Solidity: function whitelistedPlatforms(address ) view returns(bool)
func (_VmRemoteManager *VmRemoteManagerSession) WhitelistedPlatforms(arg0 common.Address) (bool, error) {
	return _VmRemoteManager.Contract.WhitelistedPlatforms(&_VmRemoteManager.CallOpts, arg0)
}

// WhitelistedPlatforms is a free data retrieval call binding the contract method 0x8924ec96.
//
// Solidity: function whitelistedPlatforms(address ) view returns(bool)
func (_VmRemoteManager *VmRemoteManagerCallerSession) WhitelistedPlatforms(arg0 common.Address) (bool, error) {
	return _VmRemoteManager.Contract.WhitelistedPlatforms(&_VmRemoteManager.CallOpts, arg0)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CancelOwnershipHandover(&_VmRemoteManager.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CancelOwnershipHandover(&_VmRemoteManager.TransactOpts)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xc533f305.
//
// Solidity: function closeCampaign((uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) CloseCampaign(opts *bind.TransactOpts, params CampaignRemoteManagerCampaignClosingParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "closeCampaign", params, destinationChainId, additionalGasLimit, votemarket)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xc533f305.
//
// Solidity: function closeCampaign((uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) CloseCampaign(params CampaignRemoteManagerCampaignClosingParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CloseCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// CloseCampaign is a paid mutator transaction binding the contract method 0xc533f305.
//
// Solidity: function closeCampaign((uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) CloseCampaign(params CampaignRemoteManagerCampaignClosingParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CloseCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CompleteOwnershipHandover(&_VmRemoteManager.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CompleteOwnershipHandover(&_VmRemoteManager.TransactOpts, pendingOwner)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x6509497d.
//
// Solidity: function createCampaign((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) CreateCampaign(opts *bind.TransactOpts, params CampaignRemoteManagerCampaignCreationParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "createCampaign", params, destinationChainId, additionalGasLimit, votemarket)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x6509497d.
//
// Solidity: function createCampaign((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) CreateCampaign(params CampaignRemoteManagerCampaignCreationParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CreateCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x6509497d.
//
// Solidity: function createCampaign((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) CreateCampaign(params CampaignRemoteManagerCampaignCreationParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.CreateCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0xd717b354.
//
// Solidity: function manageCampaign((uint256,address,uint8,uint256,uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) ManageCampaign(opts *bind.TransactOpts, params CampaignRemoteManagerCampaignManagementParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "manageCampaign", params, destinationChainId, additionalGasLimit, votemarket)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0xd717b354.
//
// Solidity: function manageCampaign((uint256,address,uint8,uint256,uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) ManageCampaign(params CampaignRemoteManagerCampaignManagementParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.ManageCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// ManageCampaign is a paid mutator transaction binding the contract method 0xd717b354.
//
// Solidity: function manageCampaign((uint256,address,uint8,uint256,uint256) params, uint256 destinationChainId, uint256 additionalGasLimit, address votemarket) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) ManageCampaign(params CampaignRemoteManagerCampaignManagementParams, destinationChainId *big.Int, additionalGasLimit *big.Int, votemarket common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.ManageCampaign(&_VmRemoteManager.TransactOpts, params, destinationChainId, additionalGasLimit, votemarket)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1885c2f9.
//
// Solidity: function receiveMessage(uint256 chainId, address sender, bytes payload) returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) ReceiveMessage(opts *bind.TransactOpts, chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "receiveMessage", chainId, sender, payload)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1885c2f9.
//
// Solidity: function receiveMessage(uint256 chainId, address sender, bytes payload) returns()
func (_VmRemoteManager *VmRemoteManagerSession) ReceiveMessage(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.ReceiveMessage(&_VmRemoteManager.TransactOpts, chainId, sender, payload)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1885c2f9.
//
// Solidity: function receiveMessage(uint256 chainId, address sender, bytes payload) returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) ReceiveMessage(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.ReceiveMessage(&_VmRemoteManager.TransactOpts, chainId, sender, payload)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_VmRemoteManager *VmRemoteManagerSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RecoverERC20(&_VmRemoteManager.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RecoverERC20(&_VmRemoteManager.TransactOpts, token, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RenounceOwnership(&_VmRemoteManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RenounceOwnership(&_VmRemoteManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RequestOwnershipHandover(&_VmRemoteManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _VmRemoteManager.Contract.RequestOwnershipHandover(&_VmRemoteManager.TransactOpts)
}

// SetPlatformWhitelist is a paid mutator transaction binding the contract method 0xc3ef672b.
//
// Solidity: function setPlatformWhitelist(address platform, bool whitelisted) returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) SetPlatformWhitelist(opts *bind.TransactOpts, platform common.Address, whitelisted bool) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "setPlatformWhitelist", platform, whitelisted)
}

// SetPlatformWhitelist is a paid mutator transaction binding the contract method 0xc3ef672b.
//
// Solidity: function setPlatformWhitelist(address platform, bool whitelisted) returns()
func (_VmRemoteManager *VmRemoteManagerSession) SetPlatformWhitelist(platform common.Address, whitelisted bool) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.SetPlatformWhitelist(&_VmRemoteManager.TransactOpts, platform, whitelisted)
}

// SetPlatformWhitelist is a paid mutator transaction binding the contract method 0xc3ef672b.
//
// Solidity: function setPlatformWhitelist(address platform, bool whitelisted) returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) SetPlatformWhitelist(platform common.Address, whitelisted bool) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.SetPlatformWhitelist(&_VmRemoteManager.TransactOpts, platform, whitelisted)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.TransferOwnership(&_VmRemoteManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_VmRemoteManager *VmRemoteManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmRemoteManager.Contract.TransferOwnership(&_VmRemoteManager.TransactOpts, newOwner)
}

// VmRemoteManagerCampaignClosingPayloadSentIterator is returned from FilterCampaignClosingPayloadSent and is used to iterate over the raw logs and unpacked data for CampaignClosingPayloadSent events raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignClosingPayloadSentIterator struct {
	Event *VmRemoteManagerCampaignClosingPayloadSent // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerCampaignClosingPayloadSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerCampaignClosingPayloadSent)
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
		it.Event = new(VmRemoteManagerCampaignClosingPayloadSent)
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
func (it *VmRemoteManagerCampaignClosingPayloadSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerCampaignClosingPayloadSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerCampaignClosingPayloadSent represents a CampaignClosingPayloadSent event raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignClosingPayloadSent struct {
	Params CampaignRemoteManagerCampaignClosingParams
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCampaignClosingPayloadSent is a free log retrieval operation binding the contract event 0xc2777135af21b673acfb520535a74f921ac9528406c8519680841d38a4a86afb.
//
// Solidity: event CampaignClosingPayloadSent((uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterCampaignClosingPayloadSent(opts *bind.FilterOpts, params []CampaignRemoteManagerCampaignClosingParams) (*VmRemoteManagerCampaignClosingPayloadSentIterator, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "CampaignClosingPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerCampaignClosingPayloadSentIterator{contract: _VmRemoteManager.contract, event: "CampaignClosingPayloadSent", logs: logs, sub: sub}, nil
}

// WatchCampaignClosingPayloadSent is a free log subscription operation binding the contract event 0xc2777135af21b673acfb520535a74f921ac9528406c8519680841d38a4a86afb.
//
// Solidity: event CampaignClosingPayloadSent((uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchCampaignClosingPayloadSent(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerCampaignClosingPayloadSent, params []CampaignRemoteManagerCampaignClosingParams) (event.Subscription, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "CampaignClosingPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerCampaignClosingPayloadSent)
				if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignClosingPayloadSent", log); err != nil {
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

// ParseCampaignClosingPayloadSent is a log parse operation binding the contract event 0xc2777135af21b673acfb520535a74f921ac9528406c8519680841d38a4a86afb.
//
// Solidity: event CampaignClosingPayloadSent((uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseCampaignClosingPayloadSent(log types.Log) (*VmRemoteManagerCampaignClosingPayloadSent, error) {
	event := new(VmRemoteManagerCampaignClosingPayloadSent)
	if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignClosingPayloadSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerCampaignCreationPayloadSentIterator is returned from FilterCampaignCreationPayloadSent and is used to iterate over the raw logs and unpacked data for CampaignCreationPayloadSent events raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignCreationPayloadSentIterator struct {
	Event *VmRemoteManagerCampaignCreationPayloadSent // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerCampaignCreationPayloadSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerCampaignCreationPayloadSent)
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
		it.Event = new(VmRemoteManagerCampaignCreationPayloadSent)
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
func (it *VmRemoteManagerCampaignCreationPayloadSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerCampaignCreationPayloadSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerCampaignCreationPayloadSent represents a CampaignCreationPayloadSent event raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignCreationPayloadSent struct {
	Params CampaignRemoteManagerCampaignCreationParams
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreationPayloadSent is a free log retrieval operation binding the contract event 0x4aae98d337a0a5163a588f53c6f71e48434770a50edfaed755ab46d6b78b0695.
//
// Solidity: event CampaignCreationPayloadSent((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterCampaignCreationPayloadSent(opts *bind.FilterOpts, params []CampaignRemoteManagerCampaignCreationParams) (*VmRemoteManagerCampaignCreationPayloadSentIterator, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "CampaignCreationPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerCampaignCreationPayloadSentIterator{contract: _VmRemoteManager.contract, event: "CampaignCreationPayloadSent", logs: logs, sub: sub}, nil
}

// WatchCampaignCreationPayloadSent is a free log subscription operation binding the contract event 0x4aae98d337a0a5163a588f53c6f71e48434770a50edfaed755ab46d6b78b0695.
//
// Solidity: event CampaignCreationPayloadSent((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchCampaignCreationPayloadSent(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerCampaignCreationPayloadSent, params []CampaignRemoteManagerCampaignCreationParams) (event.Subscription, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "CampaignCreationPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerCampaignCreationPayloadSent)
				if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignCreationPayloadSent", log); err != nil {
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

// ParseCampaignCreationPayloadSent is a log parse operation binding the contract event 0x4aae98d337a0a5163a588f53c6f71e48434770a50edfaed755ab46d6b78b0695.
//
// Solidity: event CampaignCreationPayloadSent((uint256,address,address,address,uint8,uint256,uint256,address[],address,bool) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseCampaignCreationPayloadSent(log types.Log) (*VmRemoteManagerCampaignCreationPayloadSent, error) {
	event := new(VmRemoteManagerCampaignCreationPayloadSent)
	if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignCreationPayloadSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerCampaignManagementPayloadSentIterator is returned from FilterCampaignManagementPayloadSent and is used to iterate over the raw logs and unpacked data for CampaignManagementPayloadSent events raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignManagementPayloadSentIterator struct {
	Event *VmRemoteManagerCampaignManagementPayloadSent // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerCampaignManagementPayloadSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerCampaignManagementPayloadSent)
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
		it.Event = new(VmRemoteManagerCampaignManagementPayloadSent)
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
func (it *VmRemoteManagerCampaignManagementPayloadSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerCampaignManagementPayloadSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerCampaignManagementPayloadSent represents a CampaignManagementPayloadSent event raised by the VmRemoteManager contract.
type VmRemoteManagerCampaignManagementPayloadSent struct {
	Params CampaignRemoteManagerCampaignManagementParams
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCampaignManagementPayloadSent is a free log retrieval operation binding the contract event 0x45f4b0b01168bf869ef12edb3bbadd7b3f87625c391622287175b57bdea9df6b.
//
// Solidity: event CampaignManagementPayloadSent((uint256,address,uint8,uint256,uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterCampaignManagementPayloadSent(opts *bind.FilterOpts, params []CampaignRemoteManagerCampaignManagementParams) (*VmRemoteManagerCampaignManagementPayloadSentIterator, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "CampaignManagementPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerCampaignManagementPayloadSentIterator{contract: _VmRemoteManager.contract, event: "CampaignManagementPayloadSent", logs: logs, sub: sub}, nil
}

// WatchCampaignManagementPayloadSent is a free log subscription operation binding the contract event 0x45f4b0b01168bf869ef12edb3bbadd7b3f87625c391622287175b57bdea9df6b.
//
// Solidity: event CampaignManagementPayloadSent((uint256,address,uint8,uint256,uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchCampaignManagementPayloadSent(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerCampaignManagementPayloadSent, params []CampaignRemoteManagerCampaignManagementParams) (event.Subscription, error) {

	var paramsRule []interface{}
	for _, paramsItem := range params {
		paramsRule = append(paramsRule, paramsItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "CampaignManagementPayloadSent", paramsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerCampaignManagementPayloadSent)
				if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignManagementPayloadSent", log); err != nil {
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

// ParseCampaignManagementPayloadSent is a log parse operation binding the contract event 0x45f4b0b01168bf869ef12edb3bbadd7b3f87625c391622287175b57bdea9df6b.
//
// Solidity: event CampaignManagementPayloadSent((uint256,address,uint8,uint256,uint256) indexed params)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseCampaignManagementPayloadSent(log types.Log) (*VmRemoteManagerCampaignManagementPayloadSent, error) {
	event := new(VmRemoteManagerCampaignManagementPayloadSent)
	if err := _VmRemoteManager.contract.UnpackLog(event, "CampaignManagementPayloadSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipHandoverCanceledIterator struct {
	Event *VmRemoteManagerOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerOwnershipHandoverCanceled)
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
		it.Event = new(VmRemoteManagerOwnershipHandoverCanceled)
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
func (it *VmRemoteManagerOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*VmRemoteManagerOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerOwnershipHandoverCanceledIterator{contract: _VmRemoteManager.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerOwnershipHandoverCanceled)
				if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*VmRemoteManagerOwnershipHandoverCanceled, error) {
	event := new(VmRemoteManagerOwnershipHandoverCanceled)
	if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipHandoverRequestedIterator struct {
	Event *VmRemoteManagerOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerOwnershipHandoverRequested)
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
		it.Event = new(VmRemoteManagerOwnershipHandoverRequested)
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
func (it *VmRemoteManagerOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*VmRemoteManagerOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerOwnershipHandoverRequestedIterator{contract: _VmRemoteManager.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerOwnershipHandoverRequested)
				if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseOwnershipHandoverRequested(log types.Log) (*VmRemoteManagerOwnershipHandoverRequested, error) {
	event := new(VmRemoteManagerOwnershipHandoverRequested)
	if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipTransferredIterator struct {
	Event *VmRemoteManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerOwnershipTransferred)
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
		it.Event = new(VmRemoteManagerOwnershipTransferred)
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
func (it *VmRemoteManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerOwnershipTransferred represents a OwnershipTransferred event raised by the VmRemoteManager contract.
type VmRemoteManagerOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*VmRemoteManagerOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerOwnershipTransferredIterator{contract: _VmRemoteManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerOwnershipTransferred)
				if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParseOwnershipTransferred(log types.Log) (*VmRemoteManagerOwnershipTransferred, error) {
	event := new(VmRemoteManagerOwnershipTransferred)
	if err := _VmRemoteManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmRemoteManagerPlatformWhitelistUpdatedIterator is returned from FilterPlatformWhitelistUpdated and is used to iterate over the raw logs and unpacked data for PlatformWhitelistUpdated events raised by the VmRemoteManager contract.
type VmRemoteManagerPlatformWhitelistUpdatedIterator struct {
	Event *VmRemoteManagerPlatformWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *VmRemoteManagerPlatformWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmRemoteManagerPlatformWhitelistUpdated)
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
		it.Event = new(VmRemoteManagerPlatformWhitelistUpdated)
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
func (it *VmRemoteManagerPlatformWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmRemoteManagerPlatformWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmRemoteManagerPlatformWhitelistUpdated represents a PlatformWhitelistUpdated event raised by the VmRemoteManager contract.
type VmRemoteManagerPlatformWhitelistUpdated struct {
	Platform    common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPlatformWhitelistUpdated is a free log retrieval operation binding the contract event 0x10e044199752f36cd954b8ea3850de8317639444d605b34fca03e119fc18e4f2.
//
// Solidity: event PlatformWhitelistUpdated(address indexed platform, bool whitelisted)
func (_VmRemoteManager *VmRemoteManagerFilterer) FilterPlatformWhitelistUpdated(opts *bind.FilterOpts, platform []common.Address) (*VmRemoteManagerPlatformWhitelistUpdatedIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _VmRemoteManager.contract.FilterLogs(opts, "PlatformWhitelistUpdated", platformRule)
	if err != nil {
		return nil, err
	}
	return &VmRemoteManagerPlatformWhitelistUpdatedIterator{contract: _VmRemoteManager.contract, event: "PlatformWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformWhitelistUpdated is a free log subscription operation binding the contract event 0x10e044199752f36cd954b8ea3850de8317639444d605b34fca03e119fc18e4f2.
//
// Solidity: event PlatformWhitelistUpdated(address indexed platform, bool whitelisted)
func (_VmRemoteManager *VmRemoteManagerFilterer) WatchPlatformWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *VmRemoteManagerPlatformWhitelistUpdated, platform []common.Address) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _VmRemoteManager.contract.WatchLogs(opts, "PlatformWhitelistUpdated", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmRemoteManagerPlatformWhitelistUpdated)
				if err := _VmRemoteManager.contract.UnpackLog(event, "PlatformWhitelistUpdated", log); err != nil {
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

// ParsePlatformWhitelistUpdated is a log parse operation binding the contract event 0x10e044199752f36cd954b8ea3850de8317639444d605b34fca03e119fc18e4f2.
//
// Solidity: event PlatformWhitelistUpdated(address indexed platform, bool whitelisted)
func (_VmRemoteManager *VmRemoteManagerFilterer) ParsePlatformWhitelistUpdated(log types.Log) (*VmRemoteManagerPlatformWhitelistUpdated, error) {
	event := new(VmRemoteManagerPlatformWhitelistUpdated)
	if err := _VmRemoteManager.contract.UnpackLog(event, "PlatformWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
