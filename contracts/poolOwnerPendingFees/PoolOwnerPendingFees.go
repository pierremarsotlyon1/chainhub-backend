// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolOwnerPendingFees

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

// PoolOwnerPendingFeesMetaData contains all meta data concerning the PoolOwnerPendingFees contract.
var PoolOwnerPendingFeesMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitAdmins\",\"inputs\":[{\"type\":\"address\",\"name\":\"ownership_admin\",\"indexed\":false},{\"type\":\"address\",\"name\":\"parameter_admin\",\"indexed\":false},{\"type\":\"address\",\"name\":\"emergency_admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyAdmins\",\"inputs\":[{\"type\":\"address\",\"name\":\"ownership_admin\",\"indexed\":false},{\"type\":\"address\",\"name\":\"parameter_admin\",\"indexed\":false},{\"type\":\"address\",\"name\":\"emergency_admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddBurner\",\"inputs\":[{\"type\":\"address\",\"name\":\"burner\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_ownership_admin\"},{\"type\":\"address\",\"name\":\"_parameter_admin\"},{\"type\":\"address\",\"name\":\"_emergency_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"name\":\"commit_set_admins\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_o_admin\"},{\"type\":\"address\",\"name\":\"_p_admin\"},{\"type\":\"address\",\"name\":\"_e_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":109078},{\"name\":\"apply_set_admins\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":111238},{\"name\":\"set_burner\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_coin\"},{\"type\":\"address\",\"name\":\"_burner\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":103723},{\"name\":\"set_many_burners\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"_coins\"},{\"type\":\"address[20]\",\"name\":\"_burners\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":989368},{\"name\":\"withdraw_admin_fees\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":58106},{\"name\":\"withdraw_many\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"_pools\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":93116},{\"name\":\"burn\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_coin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":93478},{\"name\":\"burn_many\",\"outputs\":[],\"inputs\":[{\"type\":\"address[20]\",\"name\":\"_coins\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":780568},{\"name\":\"kill_me\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":59139},{\"name\":\"unkill_me\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60016},{\"name\":\"set_burner_kill\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"_is_killed\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37514},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"new_owner\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":59341},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":58346},{\"name\":\"accept_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":58376},{\"name\":\"revert_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60546},{\"name\":\"commit_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"amplification\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"},{\"type\":\"uint256\",\"name\":\"min_asymmetry\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":94512},{\"name\":\"apply_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":76224},{\"name\":\"revert_new_parameters\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":61604},{\"name\":\"commit_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"new_fee\"},{\"type\":\"uint256\",\"name\":\"new_admin_fee\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":59463},{\"name\":\"apply_new_fee\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":58556},{\"name\":\"ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"_future_A\"},{\"type\":\"uint256\",\"name\":\"_future_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":59523},{\"name\":\"stop_ramp_A\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60756},{\"name\":\"set_aave_referral\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"referral_code\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":59571},{\"name\":\"set_donate_approval\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_caller\"},{\"type\":\"bool\",\"name\":\"_is_approved\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37487},{\"name\":\"donate_admin_fees\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":60697},{\"name\":\"ownership_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"parameter_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"emergency_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1991},{\"name\":\"future_ownership_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2021},{\"name\":\"future_parameter_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2051},{\"name\":\"future_emergency_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2081},{\"name\":\"min_asymmetries\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2326},{\"name\":\"burners\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2356},{\"name\":\"burner_kill\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"donate_approval\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2631}]",
}

// PoolOwnerPendingFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolOwnerPendingFeesMetaData.ABI instead.
var PoolOwnerPendingFeesABI = PoolOwnerPendingFeesMetaData.ABI

// PoolOwnerPendingFees is an auto generated Go binding around an Ethereum contract.
type PoolOwnerPendingFees struct {
	PoolOwnerPendingFeesCaller     // Read-only binding to the contract
	PoolOwnerPendingFeesTransactor // Write-only binding to the contract
	PoolOwnerPendingFeesFilterer   // Log filterer for contract events
}

// PoolOwnerPendingFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolOwnerPendingFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolOwnerPendingFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolOwnerPendingFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolOwnerPendingFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolOwnerPendingFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolOwnerPendingFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolOwnerPendingFeesSession struct {
	Contract     *PoolOwnerPendingFees // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PoolOwnerPendingFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolOwnerPendingFeesCallerSession struct {
	Contract *PoolOwnerPendingFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PoolOwnerPendingFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolOwnerPendingFeesTransactorSession struct {
	Contract     *PoolOwnerPendingFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PoolOwnerPendingFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolOwnerPendingFeesRaw struct {
	Contract *PoolOwnerPendingFees // Generic contract binding to access the raw methods on
}

// PoolOwnerPendingFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolOwnerPendingFeesCallerRaw struct {
	Contract *PoolOwnerPendingFeesCaller // Generic read-only contract binding to access the raw methods on
}

// PoolOwnerPendingFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolOwnerPendingFeesTransactorRaw struct {
	Contract *PoolOwnerPendingFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolOwnerPendingFees creates a new instance of PoolOwnerPendingFees, bound to a specific deployed contract.
func NewPoolOwnerPendingFees(address common.Address, backend bind.ContractBackend) (*PoolOwnerPendingFees, error) {
	contract, err := bindPoolOwnerPendingFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFees{PoolOwnerPendingFeesCaller: PoolOwnerPendingFeesCaller{contract: contract}, PoolOwnerPendingFeesTransactor: PoolOwnerPendingFeesTransactor{contract: contract}, PoolOwnerPendingFeesFilterer: PoolOwnerPendingFeesFilterer{contract: contract}}, nil
}

// NewPoolOwnerPendingFeesCaller creates a new read-only instance of PoolOwnerPendingFees, bound to a specific deployed contract.
func NewPoolOwnerPendingFeesCaller(address common.Address, caller bind.ContractCaller) (*PoolOwnerPendingFeesCaller, error) {
	contract, err := bindPoolOwnerPendingFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesCaller{contract: contract}, nil
}

// NewPoolOwnerPendingFeesTransactor creates a new write-only instance of PoolOwnerPendingFees, bound to a specific deployed contract.
func NewPoolOwnerPendingFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolOwnerPendingFeesTransactor, error) {
	contract, err := bindPoolOwnerPendingFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesTransactor{contract: contract}, nil
}

// NewPoolOwnerPendingFeesFilterer creates a new log filterer instance of PoolOwnerPendingFees, bound to a specific deployed contract.
func NewPoolOwnerPendingFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolOwnerPendingFeesFilterer, error) {
	contract, err := bindPoolOwnerPendingFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesFilterer{contract: contract}, nil
}

// bindPoolOwnerPendingFees binds a generic wrapper to an already deployed contract.
func bindPoolOwnerPendingFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolOwnerPendingFeesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolOwnerPendingFees.Contract.PoolOwnerPendingFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.PoolOwnerPendingFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.PoolOwnerPendingFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolOwnerPendingFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.contract.Transact(opts, method, params...)
}

// BurnerKill is a free data retrieval call binding the contract method 0x49dd3788.
//
// Solidity: function burner_kill() view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) BurnerKill(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "burner_kill")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnerKill is a free data retrieval call binding the contract method 0x49dd3788.
//
// Solidity: function burner_kill() view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) BurnerKill() (bool, error) {
	return _PoolOwnerPendingFees.Contract.BurnerKill(&_PoolOwnerPendingFees.CallOpts)
}

// BurnerKill is a free data retrieval call binding the contract method 0x49dd3788.
//
// Solidity: function burner_kill() view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) BurnerKill() (bool, error) {
	return _PoolOwnerPendingFees.Contract.BurnerKill(&_PoolOwnerPendingFees.CallOpts)
}

// Burners is a free data retrieval call binding the contract method 0x03d41e0e.
//
// Solidity: function burners(address arg0) view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) Burners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "burners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Burners is a free data retrieval call binding the contract method 0x03d41e0e.
//
// Solidity: function burners(address arg0) view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) Burners(arg0 common.Address) (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.Burners(&_PoolOwnerPendingFees.CallOpts, arg0)
}

// Burners is a free data retrieval call binding the contract method 0x03d41e0e.
//
// Solidity: function burners(address arg0) view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) Burners(arg0 common.Address) (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.Burners(&_PoolOwnerPendingFees.CallOpts, arg0)
}

// DonateApproval is a free data retrieval call binding the contract method 0x87dcefb7.
//
// Solidity: function donate_approval(address arg0, address arg1) view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) DonateApproval(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "donate_approval", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DonateApproval is a free data retrieval call binding the contract method 0x87dcefb7.
//
// Solidity: function donate_approval(address arg0, address arg1) view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) DonateApproval(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _PoolOwnerPendingFees.Contract.DonateApproval(&_PoolOwnerPendingFees.CallOpts, arg0, arg1)
}

// DonateApproval is a free data retrieval call binding the contract method 0x87dcefb7.
//
// Solidity: function donate_approval(address arg0, address arg1) view returns(bool)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) DonateApproval(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _PoolOwnerPendingFees.Contract.DonateApproval(&_PoolOwnerPendingFees.CallOpts, arg0, arg1)
}

// EmergencyAdmin is a free data retrieval call binding the contract method 0x680c7783.
//
// Solidity: function emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) EmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "emergency_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyAdmin is a free data retrieval call binding the contract method 0x680c7783.
//
// Solidity: function emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) EmergencyAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.EmergencyAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// EmergencyAdmin is a free data retrieval call binding the contract method 0x680c7783.
//
// Solidity: function emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) EmergencyAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.EmergencyAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureEmergencyAdmin is a free data retrieval call binding the contract method 0x5866507a.
//
// Solidity: function future_emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) FutureEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "future_emergency_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureEmergencyAdmin is a free data retrieval call binding the contract method 0x5866507a.
//
// Solidity: function future_emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) FutureEmergencyAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureEmergencyAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureEmergencyAdmin is a free data retrieval call binding the contract method 0x5866507a.
//
// Solidity: function future_emergency_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) FutureEmergencyAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureEmergencyAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureOwnershipAdmin is a free data retrieval call binding the contract method 0x3c2fcbf4.
//
// Solidity: function future_ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) FutureOwnershipAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "future_ownership_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwnershipAdmin is a free data retrieval call binding the contract method 0x3c2fcbf4.
//
// Solidity: function future_ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) FutureOwnershipAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureOwnershipAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureOwnershipAdmin is a free data retrieval call binding the contract method 0x3c2fcbf4.
//
// Solidity: function future_ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) FutureOwnershipAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureOwnershipAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureParameterAdmin is a free data retrieval call binding the contract method 0x824b5085.
//
// Solidity: function future_parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) FutureParameterAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "future_parameter_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureParameterAdmin is a free data retrieval call binding the contract method 0x824b5085.
//
// Solidity: function future_parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) FutureParameterAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureParameterAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// FutureParameterAdmin is a free data retrieval call binding the contract method 0x824b5085.
//
// Solidity: function future_parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) FutureParameterAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.FutureParameterAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// MinAsymmetries is a free data retrieval call binding the contract method 0xdd165f91.
//
// Solidity: function min_asymmetries(address arg0) view returns(uint256)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) MinAsymmetries(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "min_asymmetries", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAsymmetries is a free data retrieval call binding the contract method 0xdd165f91.
//
// Solidity: function min_asymmetries(address arg0) view returns(uint256)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) MinAsymmetries(arg0 common.Address) (*big.Int, error) {
	return _PoolOwnerPendingFees.Contract.MinAsymmetries(&_PoolOwnerPendingFees.CallOpts, arg0)
}

// MinAsymmetries is a free data retrieval call binding the contract method 0xdd165f91.
//
// Solidity: function min_asymmetries(address arg0) view returns(uint256)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) MinAsymmetries(arg0 common.Address) (*big.Int, error) {
	return _PoolOwnerPendingFees.Contract.MinAsymmetries(&_PoolOwnerPendingFees.CallOpts, arg0)
}

// OwnershipAdmin is a free data retrieval call binding the contract method 0x47c8715f.
//
// Solidity: function ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) OwnershipAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "ownership_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnershipAdmin is a free data retrieval call binding the contract method 0x47c8715f.
//
// Solidity: function ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) OwnershipAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.OwnershipAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// OwnershipAdmin is a free data retrieval call binding the contract method 0x47c8715f.
//
// Solidity: function ownership_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) OwnershipAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.OwnershipAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// ParameterAdmin is a free data retrieval call binding the contract method 0xa5b0b7e9.
//
// Solidity: function parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCaller) ParameterAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolOwnerPendingFees.contract.Call(opts, &out, "parameter_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParameterAdmin is a free data retrieval call binding the contract method 0xa5b0b7e9.
//
// Solidity: function parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) ParameterAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.ParameterAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// ParameterAdmin is a free data retrieval call binding the contract method 0xa5b0b7e9.
//
// Solidity: function parameter_admin() view returns(address)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesCallerSession) ParameterAdmin() (common.Address, error) {
	return _PoolOwnerPendingFees.Contract.ParameterAdmin(&_PoolOwnerPendingFees.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xa8f22cf1.
//
// Solidity: function accept_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) AcceptTransferOwnership(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "accept_transfer_ownership", _pool)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xa8f22cf1.
//
// Solidity: function accept_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) AcceptTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.AcceptTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xa8f22cf1.
//
// Solidity: function accept_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) AcceptTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.AcceptTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0xe8d64d6c.
//
// Solidity: function apply_new_fee(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) ApplyNewFee(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "apply_new_fee", _pool)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0xe8d64d6c.
//
// Solidity: function apply_new_fee(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) ApplyNewFee(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyNewFee(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplyNewFee is a paid mutator transaction binding the contract method 0xe8d64d6c.
//
// Solidity: function apply_new_fee(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) ApplyNewFee(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyNewFee(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0xcf56a4d8.
//
// Solidity: function apply_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) ApplyNewParameters(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "apply_new_parameters", _pool)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0xcf56a4d8.
//
// Solidity: function apply_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) ApplyNewParameters(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0xcf56a4d8.
//
// Solidity: function apply_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) ApplyNewParameters(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplySetAdmins is a paid mutator transaction binding the contract method 0x61893921.
//
// Solidity: function apply_set_admins() returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) ApplySetAdmins(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "apply_set_admins")
}

// ApplySetAdmins is a paid mutator transaction binding the contract method 0x61893921.
//
// Solidity: function apply_set_admins() returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) ApplySetAdmins() (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplySetAdmins(&_PoolOwnerPendingFees.TransactOpts)
}

// ApplySetAdmins is a paid mutator transaction binding the contract method 0x61893921.
//
// Solidity: function apply_set_admins() returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) ApplySetAdmins() (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplySetAdmins(&_PoolOwnerPendingFees.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x5f608d1e.
//
// Solidity: function apply_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) ApplyTransferOwnership(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "apply_transfer_ownership", _pool)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x5f608d1e.
//
// Solidity: function apply_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) ApplyTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x5f608d1e.
//
// Solidity: function apply_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) ApplyTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.ApplyTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) Burn(opts *bind.TransactOpts, _coin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "burn", _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.Burn(&_PoolOwnerPendingFees.TransactOpts, _coin)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address _coin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) Burn(_coin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.Burn(&_PoolOwnerPendingFees.TransactOpts, _coin)
}

// BurnMany is a paid mutator transaction binding the contract method 0x910a8aec.
//
// Solidity: function burn_many(address[20] _coins) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) BurnMany(opts *bind.TransactOpts, _coins [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "burn_many", _coins)
}

// BurnMany is a paid mutator transaction binding the contract method 0x910a8aec.
//
// Solidity: function burn_many(address[20] _coins) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) BurnMany(_coins [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.BurnMany(&_PoolOwnerPendingFees.TransactOpts, _coins)
}

// BurnMany is a paid mutator transaction binding the contract method 0x910a8aec.
//
// Solidity: function burn_many(address[20] _coins) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) BurnMany(_coins [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.BurnMany(&_PoolOwnerPendingFees.TransactOpts, _coins)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xcfca0bdb.
//
// Solidity: function commit_new_fee(address _pool, uint256 new_fee, uint256 new_admin_fee) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) CommitNewFee(opts *bind.TransactOpts, _pool common.Address, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "commit_new_fee", _pool, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xcfca0bdb.
//
// Solidity: function commit_new_fee(address _pool, uint256 new_fee, uint256 new_admin_fee) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) CommitNewFee(_pool common.Address, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitNewFee(&_PoolOwnerPendingFees.TransactOpts, _pool, new_fee, new_admin_fee)
}

// CommitNewFee is a paid mutator transaction binding the contract method 0xcfca0bdb.
//
// Solidity: function commit_new_fee(address _pool, uint256 new_fee, uint256 new_admin_fee) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) CommitNewFee(_pool common.Address, new_fee *big.Int, new_admin_fee *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitNewFee(&_PoolOwnerPendingFees.TransactOpts, _pool, new_fee, new_admin_fee)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x6f331ff3.
//
// Solidity: function commit_new_parameters(address _pool, uint256 amplification, uint256 new_fee, uint256 new_admin_fee, uint256 min_asymmetry) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) CommitNewParameters(opts *bind.TransactOpts, _pool common.Address, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int, min_asymmetry *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "commit_new_parameters", _pool, amplification, new_fee, new_admin_fee, min_asymmetry)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x6f331ff3.
//
// Solidity: function commit_new_parameters(address _pool, uint256 amplification, uint256 new_fee, uint256 new_admin_fee, uint256 min_asymmetry) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) CommitNewParameters(_pool common.Address, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int, min_asymmetry *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool, amplification, new_fee, new_admin_fee, min_asymmetry)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x6f331ff3.
//
// Solidity: function commit_new_parameters(address _pool, uint256 amplification, uint256 new_fee, uint256 new_admin_fee, uint256 min_asymmetry) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) CommitNewParameters(_pool common.Address, amplification *big.Int, new_fee *big.Int, new_admin_fee *big.Int, min_asymmetry *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool, amplification, new_fee, new_admin_fee, min_asymmetry)
}

// CommitSetAdmins is a paid mutator transaction binding the contract method 0x8cb16c8a.
//
// Solidity: function commit_set_admins(address _o_admin, address _p_admin, address _e_admin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) CommitSetAdmins(opts *bind.TransactOpts, _o_admin common.Address, _p_admin common.Address, _e_admin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "commit_set_admins", _o_admin, _p_admin, _e_admin)
}

// CommitSetAdmins is a paid mutator transaction binding the contract method 0x8cb16c8a.
//
// Solidity: function commit_set_admins(address _o_admin, address _p_admin, address _e_admin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) CommitSetAdmins(_o_admin common.Address, _p_admin common.Address, _e_admin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitSetAdmins(&_PoolOwnerPendingFees.TransactOpts, _o_admin, _p_admin, _e_admin)
}

// CommitSetAdmins is a paid mutator transaction binding the contract method 0x8cb16c8a.
//
// Solidity: function commit_set_admins(address _o_admin, address _p_admin, address _e_admin) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) CommitSetAdmins(_o_admin common.Address, _p_admin common.Address, _e_admin common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitSetAdmins(&_PoolOwnerPendingFees.TransactOpts, _o_admin, _p_admin, _e_admin)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x3ea1c6f4.
//
// Solidity: function commit_transfer_ownership(address _pool, address new_owner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _pool common.Address, new_owner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "commit_transfer_ownership", _pool, new_owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x3ea1c6f4.
//
// Solidity: function commit_transfer_ownership(address _pool, address new_owner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) CommitTransferOwnership(_pool common.Address, new_owner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool, new_owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x3ea1c6f4.
//
// Solidity: function commit_transfer_ownership(address _pool, address new_owner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) CommitTransferOwnership(_pool common.Address, new_owner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.CommitTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool, new_owner)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0xdb0a8406.
//
// Solidity: function donate_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) DonateAdminFees(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "donate_admin_fees", _pool)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0xdb0a8406.
//
// Solidity: function donate_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) DonateAdminFees(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.DonateAdminFees(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// DonateAdminFees is a paid mutator transaction binding the contract method 0xdb0a8406.
//
// Solidity: function donate_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) DonateAdminFees(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.DonateAdminFees(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// KillMe is a paid mutator transaction binding the contract method 0xb01f275f.
//
// Solidity: function kill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) KillMe(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "kill_me", _pool)
}

// KillMe is a paid mutator transaction binding the contract method 0xb01f275f.
//
// Solidity: function kill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) KillMe(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.KillMe(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// KillMe is a paid mutator transaction binding the contract method 0xb01f275f.
//
// Solidity: function kill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) KillMe(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.KillMe(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// RampA is a paid mutator transaction binding the contract method 0x9d4a4380.
//
// Solidity: function ramp_A(address _pool, uint256 _future_A, uint256 _future_time) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) RampA(opts *bind.TransactOpts, _pool common.Address, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "ramp_A", _pool, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x9d4a4380.
//
// Solidity: function ramp_A(address _pool, uint256 _future_A, uint256 _future_time) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) RampA(_pool common.Address, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RampA(&_PoolOwnerPendingFees.TransactOpts, _pool, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x9d4a4380.
//
// Solidity: function ramp_A(address _pool, uint256 _future_A, uint256 _future_time) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) RampA(_pool common.Address, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RampA(&_PoolOwnerPendingFees.TransactOpts, _pool, _future_A, _future_time)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x5082b389.
//
// Solidity: function revert_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) RevertNewParameters(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "revert_new_parameters", _pool)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x5082b389.
//
// Solidity: function revert_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) RevertNewParameters(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RevertNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x5082b389.
//
// Solidity: function revert_new_parameters(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) RevertNewParameters(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RevertNewParameters(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0xa352c2eb.
//
// Solidity: function revert_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) RevertTransferOwnership(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "revert_transfer_ownership", _pool)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0xa352c2eb.
//
// Solidity: function revert_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) RevertTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RevertTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0xa352c2eb.
//
// Solidity: function revert_transfer_ownership(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) RevertTransferOwnership(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.RevertTransferOwnership(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xdda3c543.
//
// Solidity: function set_aave_referral(address _pool, uint256 referral_code) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) SetAaveReferral(opts *bind.TransactOpts, _pool common.Address, referral_code *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "set_aave_referral", _pool, referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xdda3c543.
//
// Solidity: function set_aave_referral(address _pool, uint256 referral_code) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) SetAaveReferral(_pool common.Address, referral_code *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetAaveReferral(&_PoolOwnerPendingFees.TransactOpts, _pool, referral_code)
}

// SetAaveReferral is a paid mutator transaction binding the contract method 0xdda3c543.
//
// Solidity: function set_aave_referral(address _pool, uint256 referral_code) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) SetAaveReferral(_pool common.Address, referral_code *big.Int) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetAaveReferral(&_PoolOwnerPendingFees.TransactOpts, _pool, referral_code)
}

// SetBurner is a paid mutator transaction binding the contract method 0x1198c785.
//
// Solidity: function set_burner(address _coin, address _burner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) SetBurner(opts *bind.TransactOpts, _coin common.Address, _burner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "set_burner", _coin, _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0x1198c785.
//
// Solidity: function set_burner(address _coin, address _burner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) SetBurner(_coin common.Address, _burner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetBurner(&_PoolOwnerPendingFees.TransactOpts, _coin, _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0x1198c785.
//
// Solidity: function set_burner(address _coin, address _burner) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) SetBurner(_coin common.Address, _burner common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetBurner(&_PoolOwnerPendingFees.TransactOpts, _coin, _burner)
}

// SetBurnerKill is a paid mutator transaction binding the contract method 0xf132f2a5.
//
// Solidity: function set_burner_kill(bool _is_killed) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) SetBurnerKill(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "set_burner_kill", _is_killed)
}

// SetBurnerKill is a paid mutator transaction binding the contract method 0xf132f2a5.
//
// Solidity: function set_burner_kill(bool _is_killed) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) SetBurnerKill(_is_killed bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetBurnerKill(&_PoolOwnerPendingFees.TransactOpts, _is_killed)
}

// SetBurnerKill is a paid mutator transaction binding the contract method 0xf132f2a5.
//
// Solidity: function set_burner_kill(bool _is_killed) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) SetBurnerKill(_is_killed bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetBurnerKill(&_PoolOwnerPendingFees.TransactOpts, _is_killed)
}

// SetDonateApproval is a paid mutator transaction binding the contract method 0xf7709539.
//
// Solidity: function set_donate_approval(address _pool, address _caller, bool _is_approved) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) SetDonateApproval(opts *bind.TransactOpts, _pool common.Address, _caller common.Address, _is_approved bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "set_donate_approval", _pool, _caller, _is_approved)
}

// SetDonateApproval is a paid mutator transaction binding the contract method 0xf7709539.
//
// Solidity: function set_donate_approval(address _pool, address _caller, bool _is_approved) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) SetDonateApproval(_pool common.Address, _caller common.Address, _is_approved bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetDonateApproval(&_PoolOwnerPendingFees.TransactOpts, _pool, _caller, _is_approved)
}

// SetDonateApproval is a paid mutator transaction binding the contract method 0xf7709539.
//
// Solidity: function set_donate_approval(address _pool, address _caller, bool _is_approved) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) SetDonateApproval(_pool common.Address, _caller common.Address, _is_approved bool) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetDonateApproval(&_PoolOwnerPendingFees.TransactOpts, _pool, _caller, _is_approved)
}

// SetManyBurners is a paid mutator transaction binding the contract method 0x0ab74d6f.
//
// Solidity: function set_many_burners(address[20] _coins, address[20] _burners) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) SetManyBurners(opts *bind.TransactOpts, _coins [20]common.Address, _burners [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "set_many_burners", _coins, _burners)
}

// SetManyBurners is a paid mutator transaction binding the contract method 0x0ab74d6f.
//
// Solidity: function set_many_burners(address[20] _coins, address[20] _burners) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) SetManyBurners(_coins [20]common.Address, _burners [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetManyBurners(&_PoolOwnerPendingFees.TransactOpts, _coins, _burners)
}

// SetManyBurners is a paid mutator transaction binding the contract method 0x0ab74d6f.
//
// Solidity: function set_many_burners(address[20] _coins, address[20] _burners) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) SetManyBurners(_coins [20]common.Address, _burners [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.SetManyBurners(&_PoolOwnerPendingFees.TransactOpts, _coins, _burners)
}

// StopRampA is a paid mutator transaction binding the contract method 0x53f79b2b.
//
// Solidity: function stop_ramp_A(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) StopRampA(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "stop_ramp_A", _pool)
}

// StopRampA is a paid mutator transaction binding the contract method 0x53f79b2b.
//
// Solidity: function stop_ramp_A(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) StopRampA(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.StopRampA(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// StopRampA is a paid mutator transaction binding the contract method 0x53f79b2b.
//
// Solidity: function stop_ramp_A(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) StopRampA(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.StopRampA(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x1cfbc236.
//
// Solidity: function unkill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) UnkillMe(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "unkill_me", _pool)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x1cfbc236.
//
// Solidity: function unkill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) UnkillMe(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.UnkillMe(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x1cfbc236.
//
// Solidity: function unkill_me(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) UnkillMe(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.UnkillMe(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0xe4e67c0f.
//
// Solidity: function withdraw_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) WithdrawAdminFees(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "withdraw_admin_fees", _pool)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0xe4e67c0f.
//
// Solidity: function withdraw_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) WithdrawAdminFees(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.WithdrawAdminFees(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0xe4e67c0f.
//
// Solidity: function withdraw_admin_fees(address _pool) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) WithdrawAdminFees(_pool common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.WithdrawAdminFees(&_PoolOwnerPendingFees.TransactOpts, _pool)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0xd7f649fd.
//
// Solidity: function withdraw_many(address[20] _pools) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) WithdrawMany(opts *bind.TransactOpts, _pools [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.Transact(opts, "withdraw_many", _pools)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0xd7f649fd.
//
// Solidity: function withdraw_many(address[20] _pools) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) WithdrawMany(_pools [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.WithdrawMany(&_PoolOwnerPendingFees.TransactOpts, _pools)
}

// WithdrawMany is a paid mutator transaction binding the contract method 0xd7f649fd.
//
// Solidity: function withdraw_many(address[20] _pools) returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) WithdrawMany(_pools [20]common.Address) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.WithdrawMany(&_PoolOwnerPendingFees.TransactOpts, _pools)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.Fallback(&_PoolOwnerPendingFees.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PoolOwnerPendingFees.Contract.Fallback(&_PoolOwnerPendingFees.TransactOpts, calldata)
}

// PoolOwnerPendingFeesAddBurnerIterator is returned from FilterAddBurner and is used to iterate over the raw logs and unpacked data for AddBurner events raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesAddBurnerIterator struct {
	Event *PoolOwnerPendingFeesAddBurner // Event containing the contract specifics and raw log

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
func (it *PoolOwnerPendingFeesAddBurnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnerPendingFeesAddBurner)
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
		it.Event = new(PoolOwnerPendingFeesAddBurner)
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
func (it *PoolOwnerPendingFeesAddBurnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnerPendingFeesAddBurnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnerPendingFeesAddBurner represents a AddBurner event raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesAddBurner struct {
	Burner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddBurner is a free log retrieval operation binding the contract event 0x2a85edc5fabdd9bbaa6d309617215d5b6905e0ed8a48d656d86fc9863e3c4b77.
//
// Solidity: event AddBurner(address burner)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) FilterAddBurner(opts *bind.FilterOpts) (*PoolOwnerPendingFeesAddBurnerIterator, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.FilterLogs(opts, "AddBurner")
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesAddBurnerIterator{contract: _PoolOwnerPendingFees.contract, event: "AddBurner", logs: logs, sub: sub}, nil
}

// WatchAddBurner is a free log subscription operation binding the contract event 0x2a85edc5fabdd9bbaa6d309617215d5b6905e0ed8a48d656d86fc9863e3c4b77.
//
// Solidity: event AddBurner(address burner)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) WatchAddBurner(opts *bind.WatchOpts, sink chan<- *PoolOwnerPendingFeesAddBurner) (event.Subscription, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.WatchLogs(opts, "AddBurner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnerPendingFeesAddBurner)
				if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "AddBurner", log); err != nil {
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

// ParseAddBurner is a log parse operation binding the contract event 0x2a85edc5fabdd9bbaa6d309617215d5b6905e0ed8a48d656d86fc9863e3c4b77.
//
// Solidity: event AddBurner(address burner)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) ParseAddBurner(log types.Log) (*PoolOwnerPendingFeesAddBurner, error) {
	event := new(PoolOwnerPendingFeesAddBurner)
	if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "AddBurner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnerPendingFeesApplyAdminsIterator is returned from FilterApplyAdmins and is used to iterate over the raw logs and unpacked data for ApplyAdmins events raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesApplyAdminsIterator struct {
	Event *PoolOwnerPendingFeesApplyAdmins // Event containing the contract specifics and raw log

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
func (it *PoolOwnerPendingFeesApplyAdminsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnerPendingFeesApplyAdmins)
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
		it.Event = new(PoolOwnerPendingFeesApplyAdmins)
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
func (it *PoolOwnerPendingFeesApplyAdminsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnerPendingFeesApplyAdminsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnerPendingFeesApplyAdmins represents a ApplyAdmins event raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesApplyAdmins struct {
	OwnershipAdmin common.Address
	ParameterAdmin common.Address
	EmergencyAdmin common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApplyAdmins is a free log retrieval operation binding the contract event 0xd61a16912efb9a1c5bd5361dff238b95f452672ded751a425c11db5e4f588176.
//
// Solidity: event ApplyAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) FilterApplyAdmins(opts *bind.FilterOpts) (*PoolOwnerPendingFeesApplyAdminsIterator, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.FilterLogs(opts, "ApplyAdmins")
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesApplyAdminsIterator{contract: _PoolOwnerPendingFees.contract, event: "ApplyAdmins", logs: logs, sub: sub}, nil
}

// WatchApplyAdmins is a free log subscription operation binding the contract event 0xd61a16912efb9a1c5bd5361dff238b95f452672ded751a425c11db5e4f588176.
//
// Solidity: event ApplyAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) WatchApplyAdmins(opts *bind.WatchOpts, sink chan<- *PoolOwnerPendingFeesApplyAdmins) (event.Subscription, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.WatchLogs(opts, "ApplyAdmins")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnerPendingFeesApplyAdmins)
				if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "ApplyAdmins", log); err != nil {
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

// ParseApplyAdmins is a log parse operation binding the contract event 0xd61a16912efb9a1c5bd5361dff238b95f452672ded751a425c11db5e4f588176.
//
// Solidity: event ApplyAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) ParseApplyAdmins(log types.Log) (*PoolOwnerPendingFeesApplyAdmins, error) {
	event := new(PoolOwnerPendingFeesApplyAdmins)
	if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "ApplyAdmins", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnerPendingFeesCommitAdminsIterator is returned from FilterCommitAdmins and is used to iterate over the raw logs and unpacked data for CommitAdmins events raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesCommitAdminsIterator struct {
	Event *PoolOwnerPendingFeesCommitAdmins // Event containing the contract specifics and raw log

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
func (it *PoolOwnerPendingFeesCommitAdminsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnerPendingFeesCommitAdmins)
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
		it.Event = new(PoolOwnerPendingFeesCommitAdmins)
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
func (it *PoolOwnerPendingFeesCommitAdminsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnerPendingFeesCommitAdminsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnerPendingFeesCommitAdmins represents a CommitAdmins event raised by the PoolOwnerPendingFees contract.
type PoolOwnerPendingFeesCommitAdmins struct {
	OwnershipAdmin common.Address
	ParameterAdmin common.Address
	EmergencyAdmin common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCommitAdmins is a free log retrieval operation binding the contract event 0x78572131fd8b9a2e345c48a6afbf55bc1219e393553feac694f89889903d2704.
//
// Solidity: event CommitAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) FilterCommitAdmins(opts *bind.FilterOpts) (*PoolOwnerPendingFeesCommitAdminsIterator, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.FilterLogs(opts, "CommitAdmins")
	if err != nil {
		return nil, err
	}
	return &PoolOwnerPendingFeesCommitAdminsIterator{contract: _PoolOwnerPendingFees.contract, event: "CommitAdmins", logs: logs, sub: sub}, nil
}

// WatchCommitAdmins is a free log subscription operation binding the contract event 0x78572131fd8b9a2e345c48a6afbf55bc1219e393553feac694f89889903d2704.
//
// Solidity: event CommitAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) WatchCommitAdmins(opts *bind.WatchOpts, sink chan<- *PoolOwnerPendingFeesCommitAdmins) (event.Subscription, error) {

	logs, sub, err := _PoolOwnerPendingFees.contract.WatchLogs(opts, "CommitAdmins")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnerPendingFeesCommitAdmins)
				if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "CommitAdmins", log); err != nil {
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

// ParseCommitAdmins is a log parse operation binding the contract event 0x78572131fd8b9a2e345c48a6afbf55bc1219e393553feac694f89889903d2704.
//
// Solidity: event CommitAdmins(address ownership_admin, address parameter_admin, address emergency_admin)
func (_PoolOwnerPendingFees *PoolOwnerPendingFeesFilterer) ParseCommitAdmins(log types.Log) (*PoolOwnerPendingFeesCommitAdmins, error) {
	event := new(PoolOwnerPendingFeesCommitAdmins)
	if err := _PoolOwnerPendingFees.contract.UnpackLog(event, "CommitAdmins", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
