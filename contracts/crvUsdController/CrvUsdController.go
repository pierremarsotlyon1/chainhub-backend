// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crvUsdController

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	User   common.Address
	X      *big.Int
	Y      *big.Int
	Debt   *big.Int
	Health *big.Int
}

// CrvUsdControllerMetaData contains all meta data concerning the CrvUsdController contract.
var CrvUsdControllerMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"UserState\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"n1\",\"type\":\"int256\",\"indexed\":false},{\"name\":\"n2\",\"type\":\"int256\",\"indexed\":false},{\"name\":\"liquidation_discount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Borrow\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_increase\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loan_increase\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Repay\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_decrease\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loan_decrease\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveCollateral\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_decrease\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Liquidate\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_received\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"stablecoin_received\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetMonetaryPolicy\",\"inputs\":[{\"name\":\"monetary_policy\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetBorrowingDiscounts\",\"inputs\":[{\"name\":\"loan_discount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"liquidation_discount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CollectFees\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"amm\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"collateral_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debt\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_exists\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_debt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_borrowable\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_borrowable\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_collateral\",\"inputs\":[{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calculate_debt_n1\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"create_loan\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"create_loan_extended\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"},{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"borrow_more\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"max_active_band\",\"type\":\"int256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"max_active_band\",\"type\":\"int256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay_extended\",\"inputs\":[{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health_calculator\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"d_collateral\",\"type\":\"int256\"},{\"name\":\"d_debt\",\"type\":\"int256\"},{\"name\":\"full\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health_calculator\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"d_collateral\",\"type\":\"int256\"},{\"name\":\"d_debt\",\"type\":\"int256\"},{\"name\":\"full\",\"type\":\"bool\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate_extended\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"},{\"name\":\"frac\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens_to_liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens_to_liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"frac\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"full\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_prices\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_state\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_amm_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_amm_admin_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_monetary_policy\",\"inputs\":[{\"name\":\"monetary_policy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_borrowing_discounts\",\"inputs\":[{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_callback\",\"inputs\":[{\"name\":\"cb\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"collect_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"liquidation_discounts\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loans\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_ix\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"n_loans\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"redeemed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"monetary_policy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"liquidation_discount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_discount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// CrvUsdControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use CrvUsdControllerMetaData.ABI instead.
var CrvUsdControllerABI = CrvUsdControllerMetaData.ABI

// CrvUsdController is an auto generated Go binding around an Ethereum contract.
type CrvUsdController struct {
	CrvUsdControllerCaller     // Read-only binding to the contract
	CrvUsdControllerTransactor // Write-only binding to the contract
	CrvUsdControllerFilterer   // Log filterer for contract events
}

// CrvUsdControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrvUsdControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrvUsdControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrvUsdControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrvUsdControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrvUsdControllerSession struct {
	Contract     *CrvUsdController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrvUsdControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrvUsdControllerCallerSession struct {
	Contract *CrvUsdControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CrvUsdControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrvUsdControllerTransactorSession struct {
	Contract     *CrvUsdControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CrvUsdControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrvUsdControllerRaw struct {
	Contract *CrvUsdController // Generic contract binding to access the raw methods on
}

// CrvUsdControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrvUsdControllerCallerRaw struct {
	Contract *CrvUsdControllerCaller // Generic read-only contract binding to access the raw methods on
}

// CrvUsdControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrvUsdControllerTransactorRaw struct {
	Contract *CrvUsdControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrvUsdController creates a new instance of CrvUsdController, bound to a specific deployed contract.
func NewCrvUsdController(address common.Address, backend bind.ContractBackend) (*CrvUsdController, error) {
	contract, err := bindCrvUsdController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrvUsdController{CrvUsdControllerCaller: CrvUsdControllerCaller{contract: contract}, CrvUsdControllerTransactor: CrvUsdControllerTransactor{contract: contract}, CrvUsdControllerFilterer: CrvUsdControllerFilterer{contract: contract}}, nil
}

// NewCrvUsdControllerCaller creates a new read-only instance of CrvUsdController, bound to a specific deployed contract.
func NewCrvUsdControllerCaller(address common.Address, caller bind.ContractCaller) (*CrvUsdControllerCaller, error) {
	contract, err := bindCrvUsdController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerCaller{contract: contract}, nil
}

// NewCrvUsdControllerTransactor creates a new write-only instance of CrvUsdController, bound to a specific deployed contract.
func NewCrvUsdControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*CrvUsdControllerTransactor, error) {
	contract, err := bindCrvUsdController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerTransactor{contract: contract}, nil
}

// NewCrvUsdControllerFilterer creates a new log filterer instance of CrvUsdController, bound to a specific deployed contract.
func NewCrvUsdControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*CrvUsdControllerFilterer, error) {
	contract, err := bindCrvUsdController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerFilterer{contract: contract}, nil
}

// bindCrvUsdController binds a generic wrapper to an already deployed contract.
func bindCrvUsdController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrvUsdControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdController *CrvUsdControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdController.Contract.CrvUsdControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdController *CrvUsdControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CrvUsdControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdController *CrvUsdControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CrvUsdControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrvUsdController *CrvUsdControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrvUsdController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrvUsdController *CrvUsdControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrvUsdController *CrvUsdControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrvUsdController.Contract.contract.Transact(opts, method, params...)
}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) AdminFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "admin_fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) AdminFees() (*big.Int, error) {
	return _CrvUsdController.Contract.AdminFees(&_CrvUsdController.CallOpts)
}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) AdminFees() (*big.Int, error) {
	return _CrvUsdController.Contract.AdminFees(&_CrvUsdController.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_CrvUsdController *CrvUsdControllerCaller) Amm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "amm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_CrvUsdController *CrvUsdControllerSession) Amm() (common.Address, error) {
	return _CrvUsdController.Contract.Amm(&_CrvUsdController.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() view returns(address)
func (_CrvUsdController *CrvUsdControllerCallerSession) Amm() (common.Address, error) {
	return _CrvUsdController.Contract.Amm(&_CrvUsdController.CallOpts)
}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) AmmPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "amm_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) AmmPrice() (*big.Int, error) {
	return _CrvUsdController.Contract.AmmPrice(&_CrvUsdController.CallOpts)
}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) AmmPrice() (*big.Int, error) {
	return _CrvUsdController.Contract.AmmPrice(&_CrvUsdController.CallOpts)
}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCaller) CalculateDebtN1(opts *bind.CallOpts, collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "calculate_debt_n1", collateral, debt, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerSession) CalculateDebtN1(collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.CalculateDebtN1(&_CrvUsdController.CallOpts, collateral, debt, N)
}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCallerSession) CalculateDebtN1(collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.CalculateDebtN1(&_CrvUsdController.CallOpts, collateral, debt, N)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_CrvUsdController *CrvUsdControllerCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "collateral_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_CrvUsdController *CrvUsdControllerSession) CollateralToken() (common.Address, error) {
	return _CrvUsdController.Contract.CollateralToken(&_CrvUsdController.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() view returns(address)
func (_CrvUsdController *CrvUsdControllerCallerSession) CollateralToken() (common.Address, error) {
	return _CrvUsdController.Contract.CollateralToken(&_CrvUsdController.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) Debt(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "debt", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) Debt(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.Debt(&_CrvUsdController.CallOpts, user)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) Debt(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.Debt(&_CrvUsdController.CallOpts, user)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrvUsdController *CrvUsdControllerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrvUsdController *CrvUsdControllerSession) Factory() (common.Address, error) {
	return _CrvUsdController.Contract.Factory(&_CrvUsdController.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrvUsdController *CrvUsdControllerCallerSession) Factory() (common.Address, error) {
	return _CrvUsdController.Contract.Factory(&_CrvUsdController.CallOpts)
}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCaller) Health(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "health", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_CrvUsdController *CrvUsdControllerSession) Health(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.Health(&_CrvUsdController.CallOpts, user)
}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCallerSession) Health(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.Health(&_CrvUsdController.CallOpts, user)
}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCaller) Health0(opts *bind.CallOpts, user common.Address, full bool) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "health0", user, full)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerSession) Health0(user common.Address, full bool) (*big.Int, error) {
	return _CrvUsdController.Contract.Health0(&_CrvUsdController.CallOpts, user, full)
}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCallerSession) Health0(user common.Address, full bool) (*big.Int, error) {
	return _CrvUsdController.Contract.Health0(&_CrvUsdController.CallOpts, user, full)
}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCaller) HealthCalculator(opts *bind.CallOpts, user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "health_calculator", user, d_collateral, d_debt, full)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerSession) HealthCalculator(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	return _CrvUsdController.Contract.HealthCalculator(&_CrvUsdController.CallOpts, user, d_collateral, d_debt, full)
}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCallerSession) HealthCalculator(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	return _CrvUsdController.Contract.HealthCalculator(&_CrvUsdController.CallOpts, user, d_collateral, d_debt, full)
}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCaller) HealthCalculator0(opts *bind.CallOpts, user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "health_calculator0", user, d_collateral, d_debt, full, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerSession) HealthCalculator0(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.HealthCalculator0(&_CrvUsdController.CallOpts, user, d_collateral, d_debt, full, N)
}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_CrvUsdController *CrvUsdControllerCallerSession) HealthCalculator0(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.HealthCalculator0(&_CrvUsdController.CallOpts, user, d_collateral, d_debt, full, N)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "liquidation_discount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) LiquidationDiscount() (*big.Int, error) {
	return _CrvUsdController.Contract.LiquidationDiscount(&_CrvUsdController.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) LiquidationDiscount() (*big.Int, error) {
	return _CrvUsdController.Contract.LiquidationDiscount(&_CrvUsdController.CallOpts)
}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) LiquidationDiscounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "liquidation_discounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) LiquidationDiscounts(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.LiquidationDiscounts(&_CrvUsdController.CallOpts, arg0)
}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) LiquidationDiscounts(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.LiquidationDiscounts(&_CrvUsdController.CallOpts, arg0)
}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) LoanDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "loan_discount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) LoanDiscount() (*big.Int, error) {
	return _CrvUsdController.Contract.LoanDiscount(&_CrvUsdController.CallOpts)
}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) LoanDiscount() (*big.Int, error) {
	return _CrvUsdController.Contract.LoanDiscount(&_CrvUsdController.CallOpts)
}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_CrvUsdController *CrvUsdControllerCaller) LoanExists(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "loan_exists", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_CrvUsdController *CrvUsdControllerSession) LoanExists(user common.Address) (bool, error) {
	return _CrvUsdController.Contract.LoanExists(&_CrvUsdController.CallOpts, user)
}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_CrvUsdController *CrvUsdControllerCallerSession) LoanExists(user common.Address) (bool, error) {
	return _CrvUsdController.Contract.LoanExists(&_CrvUsdController.CallOpts, user)
}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) LoanIx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "loan_ix", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) LoanIx(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.LoanIx(&_CrvUsdController.CallOpts, arg0)
}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) LoanIx(arg0 common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.LoanIx(&_CrvUsdController.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_CrvUsdController *CrvUsdControllerCaller) Loans(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "loans", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_CrvUsdController *CrvUsdControllerSession) Loans(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdController.Contract.Loans(&_CrvUsdController.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_CrvUsdController *CrvUsdControllerCallerSession) Loans(arg0 *big.Int) (common.Address, error) {
	return _CrvUsdController.Contract.Loans(&_CrvUsdController.CallOpts, arg0)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) MaxBorrowable(opts *bind.CallOpts, collateral *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "max_borrowable", collateral, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) MaxBorrowable(collateral *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MaxBorrowable(&_CrvUsdController.CallOpts, collateral, N)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) MaxBorrowable(collateral *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MaxBorrowable(&_CrvUsdController.CallOpts, collateral, N)
}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) MaxBorrowable0(opts *bind.CallOpts, collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "max_borrowable0", collateral, N, current_debt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) MaxBorrowable0(collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MaxBorrowable0(&_CrvUsdController.CallOpts, collateral, N, current_debt)
}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) MaxBorrowable0(collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MaxBorrowable0(&_CrvUsdController.CallOpts, collateral, N, current_debt)
}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) MinCollateral(opts *bind.CallOpts, debt *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "min_collateral", debt, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) MinCollateral(debt *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MinCollateral(&_CrvUsdController.CallOpts, debt, N)
}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) MinCollateral(debt *big.Int, N *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.MinCollateral(&_CrvUsdController.CallOpts, debt, N)
}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) Minted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "minted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) Minted() (*big.Int, error) {
	return _CrvUsdController.Contract.Minted(&_CrvUsdController.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) Minted() (*big.Int, error) {
	return _CrvUsdController.Contract.Minted(&_CrvUsdController.CallOpts)
}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_CrvUsdController *CrvUsdControllerCaller) MonetaryPolicy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "monetary_policy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_CrvUsdController *CrvUsdControllerSession) MonetaryPolicy() (common.Address, error) {
	return _CrvUsdController.Contract.MonetaryPolicy(&_CrvUsdController.CallOpts)
}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_CrvUsdController *CrvUsdControllerCallerSession) MonetaryPolicy() (common.Address, error) {
	return _CrvUsdController.Contract.MonetaryPolicy(&_CrvUsdController.CallOpts)
}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) NLoans(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "n_loans")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) NLoans() (*big.Int, error) {
	return _CrvUsdController.Contract.NLoans(&_CrvUsdController.CallOpts)
}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) NLoans() (*big.Int, error) {
	return _CrvUsdController.Contract.NLoans(&_CrvUsdController.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) Redeemed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "redeemed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) Redeemed() (*big.Int, error) {
	return _CrvUsdController.Contract.Redeemed(&_CrvUsdController.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) Redeemed() (*big.Int, error) {
	return _CrvUsdController.Contract.Redeemed(&_CrvUsdController.CallOpts)
}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) TokensToLiquidate(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "tokens_to_liquidate", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) TokensToLiquidate(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.TokensToLiquidate(&_CrvUsdController.CallOpts, user)
}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) TokensToLiquidate(user common.Address) (*big.Int, error) {
	return _CrvUsdController.Contract.TokensToLiquidate(&_CrvUsdController.CallOpts, user)
}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) TokensToLiquidate0(opts *bind.CallOpts, user common.Address, frac *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "tokens_to_liquidate0", user, frac)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) TokensToLiquidate0(user common.Address, frac *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.TokensToLiquidate0(&_CrvUsdController.CallOpts, user, frac)
}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) TokensToLiquidate0(user common.Address, frac *big.Int) (*big.Int, error) {
	return _CrvUsdController.Contract.TokensToLiquidate0(&_CrvUsdController.CallOpts, user, frac)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "total_debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) TotalDebt() (*big.Int, error) {
	return _CrvUsdController.Contract.TotalDebt(&_CrvUsdController.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_CrvUsdController *CrvUsdControllerCallerSession) TotalDebt() (*big.Int, error) {
	return _CrvUsdController.Contract.TotalDebt(&_CrvUsdController.CallOpts)
}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_CrvUsdController *CrvUsdControllerCaller) UserPrices(opts *bind.CallOpts, user common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "user_prices", user)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_CrvUsdController *CrvUsdControllerSession) UserPrices(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdController.Contract.UserPrices(&_CrvUsdController.CallOpts, user)
}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_CrvUsdController *CrvUsdControllerCallerSession) UserPrices(user common.Address) ([2]*big.Int, error) {
	return _CrvUsdController.Contract.UserPrices(&_CrvUsdController.CallOpts, user)
}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_CrvUsdController *CrvUsdControllerCaller) UserState(opts *bind.CallOpts, user common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "user_state", user)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_CrvUsdController *CrvUsdControllerSession) UserState(user common.Address) ([4]*big.Int, error) {
	return _CrvUsdController.Contract.UserState(&_CrvUsdController.CallOpts, user)
}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_CrvUsdController *CrvUsdControllerCallerSession) UserState(user common.Address) ([4]*big.Int, error) {
	return _CrvUsdController.Contract.UserState(&_CrvUsdController.CallOpts, user)
}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCaller) UsersToLiquidate(opts *bind.CallOpts) ([]Struct0, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "users_to_liquidate")

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerSession) UsersToLiquidate() ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate(&_CrvUsdController.CallOpts)
}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCallerSession) UsersToLiquidate() ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate(&_CrvUsdController.CallOpts)
}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCaller) UsersToLiquidate0(opts *bind.CallOpts, _from *big.Int) ([]Struct0, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "users_to_liquidate0", _from)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerSession) UsersToLiquidate0(_from *big.Int) ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate0(&_CrvUsdController.CallOpts, _from)
}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCallerSession) UsersToLiquidate0(_from *big.Int) ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate0(&_CrvUsdController.CallOpts, _from)
}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCaller) UsersToLiquidate1(opts *bind.CallOpts, _from *big.Int, _limit *big.Int) ([]Struct0, error) {
	var out []interface{}
	err := _CrvUsdController.contract.Call(opts, &out, "users_to_liquidate1", _from, _limit)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerSession) UsersToLiquidate1(_from *big.Int, _limit *big.Int) ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate1(&_CrvUsdController.CallOpts, _from, _limit)
}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_CrvUsdController *CrvUsdControllerCallerSession) UsersToLiquidate1(_from *big.Int, _limit *big.Int) ([]Struct0, error) {
	return _CrvUsdController.Contract.UsersToLiquidate1(&_CrvUsdController.CallOpts, _from, _limit)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) AddCollateral(opts *bind.TransactOpts, collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "add_collateral", collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) payable returns()
func (_CrvUsdController *CrvUsdControllerSession) AddCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.AddCollateral(&_CrvUsdController.TransactOpts, collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) AddCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.AddCollateral(&_CrvUsdController.TransactOpts, collateral)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) AddCollateral0(opts *bind.TransactOpts, collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "add_collateral0", collateral, _for)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) payable returns()
func (_CrvUsdController *CrvUsdControllerSession) AddCollateral0(collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.AddCollateral0(&_CrvUsdController.TransactOpts, collateral, _for)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) AddCollateral0(collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.AddCollateral0(&_CrvUsdController.TransactOpts, collateral, _for)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) BorrowMore(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "borrow_more", collateral, debt)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) payable returns()
func (_CrvUsdController *CrvUsdControllerSession) BorrowMore(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.BorrowMore(&_CrvUsdController.TransactOpts, collateral, debt)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) BorrowMore(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.BorrowMore(&_CrvUsdController.TransactOpts, collateral, debt)
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_CrvUsdController *CrvUsdControllerTransactor) CollectFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "collect_fees")
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_CrvUsdController *CrvUsdControllerSession) CollectFees() (*types.Transaction, error) {
	return _CrvUsdController.Contract.CollectFees(&_CrvUsdController.TransactOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_CrvUsdController *CrvUsdControllerTransactorSession) CollectFees() (*types.Transaction, error) {
	return _CrvUsdController.Contract.CollectFees(&_CrvUsdController.TransactOpts)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) CreateLoan(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "create_loan", collateral, debt, N)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) payable returns()
func (_CrvUsdController *CrvUsdControllerSession) CreateLoan(collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CreateLoan(&_CrvUsdController.TransactOpts, collateral, debt, N)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) CreateLoan(collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CreateLoan(&_CrvUsdController.TransactOpts, collateral, debt, N)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) CreateLoanExtended(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "create_loan_extended", collateral, debt, N, callbacker, callback_args)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) payable returns()
func (_CrvUsdController *CrvUsdControllerSession) CreateLoanExtended(collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CreateLoanExtended(&_CrvUsdController.TransactOpts, collateral, debt, N, callbacker, callback_args)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) CreateLoanExtended(collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.CreateLoanExtended(&_CrvUsdController.TransactOpts, collateral, debt, N, callbacker, callback_args)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Liquidate(opts *bind.TransactOpts, user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "liquidate", user, min_x)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_CrvUsdController *CrvUsdControllerSession) Liquidate(user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Liquidate(&_CrvUsdController.TransactOpts, user, min_x)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Liquidate(user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Liquidate(&_CrvUsdController.TransactOpts, user, min_x)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Liquidate0(opts *bind.TransactOpts, user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "liquidate0", user, min_x, use_eth)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerSession) Liquidate0(user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Liquidate0(&_CrvUsdController.TransactOpts, user, min_x, use_eth)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Liquidate0(user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Liquidate0(&_CrvUsdController.TransactOpts, user, min_x, use_eth)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) LiquidateExtended(opts *bind.TransactOpts, user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "liquidate_extended", user, min_x, frac, use_eth, callbacker, callback_args)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerSession) LiquidateExtended(user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.LiquidateExtended(&_CrvUsdController.TransactOpts, user, min_x, frac, use_eth, callbacker, callback_args)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) LiquidateExtended(user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.LiquidateExtended(&_CrvUsdController.TransactOpts, user, min_x, frac, use_eth, callbacker, callback_args)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) RemoveCollateral(opts *bind.TransactOpts, collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "remove_collateral", collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_CrvUsdController *CrvUsdControllerSession) RemoveCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RemoveCollateral(&_CrvUsdController.TransactOpts, collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) RemoveCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RemoveCollateral(&_CrvUsdController.TransactOpts, collateral)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) RemoveCollateral0(opts *bind.TransactOpts, collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "remove_collateral0", collateral, use_eth)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerSession) RemoveCollateral0(collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RemoveCollateral0(&_CrvUsdController.TransactOpts, collateral, use_eth)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) RemoveCollateral0(collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RemoveCollateral0(&_CrvUsdController.TransactOpts, collateral, use_eth)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Repay(opts *bind.TransactOpts, _d_debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "repay", _d_debt)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_CrvUsdController *CrvUsdControllerSession) Repay(_d_debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay(&_CrvUsdController.TransactOpts, _d_debt)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Repay(_d_debt *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay(&_CrvUsdController.TransactOpts, _d_debt)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Repay0(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "repay0", _d_debt, _for)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_CrvUsdController *CrvUsdControllerSession) Repay0(_d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay0(&_CrvUsdController.TransactOpts, _d_debt, _for)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Repay0(_d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay0(&_CrvUsdController.TransactOpts, _d_debt, _for)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Repay1(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "repay1", _d_debt, _for, max_active_band)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_CrvUsdController *CrvUsdControllerSession) Repay1(_d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay1(&_CrvUsdController.TransactOpts, _d_debt, _for, max_active_band)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Repay1(_d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay1(&_CrvUsdController.TransactOpts, _d_debt, _for, max_active_band)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Repay2(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "repay2", _d_debt, _for, max_active_band, use_eth)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerSession) Repay2(_d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay2(&_CrvUsdController.TransactOpts, _d_debt, _for, max_active_band, use_eth)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Repay2(_d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Repay2(&_CrvUsdController.TransactOpts, _d_debt, _for, max_active_band, use_eth)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) RepayExtended(opts *bind.TransactOpts, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "repay_extended", callbacker, callback_args)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerSession) RepayExtended(callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RepayExtended(&_CrvUsdController.TransactOpts, callbacker, callback_args)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) RepayExtended(callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.RepayExtended(&_CrvUsdController.TransactOpts, callbacker, callback_args)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) SetAmmAdminFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "set_amm_admin_fee", fee)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerSession) SetAmmAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetAmmAdminFee(&_CrvUsdController.TransactOpts, fee)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) SetAmmAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetAmmAdminFee(&_CrvUsdController.TransactOpts, fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) SetAmmFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "set_amm_fee", fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerSession) SetAmmFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetAmmFee(&_CrvUsdController.TransactOpts, fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) SetAmmFee(fee *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetAmmFee(&_CrvUsdController.TransactOpts, fee)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) SetBorrowingDiscounts(opts *bind.TransactOpts, loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "set_borrowing_discounts", loan_discount, liquidation_discount)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_CrvUsdController *CrvUsdControllerSession) SetBorrowingDiscounts(loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetBorrowingDiscounts(&_CrvUsdController.TransactOpts, loan_discount, liquidation_discount)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) SetBorrowingDiscounts(loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetBorrowingDiscounts(&_CrvUsdController.TransactOpts, loan_discount, liquidation_discount)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) SetCallback(opts *bind.TransactOpts, cb common.Address) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "set_callback", cb)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_CrvUsdController *CrvUsdControllerSession) SetCallback(cb common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetCallback(&_CrvUsdController.TransactOpts, cb)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) SetCallback(cb common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetCallback(&_CrvUsdController.TransactOpts, cb)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_CrvUsdController *CrvUsdControllerTransactor) SetMonetaryPolicy(opts *bind.TransactOpts, monetary_policy common.Address) (*types.Transaction, error) {
	return _CrvUsdController.contract.Transact(opts, "set_monetary_policy", monetary_policy)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_CrvUsdController *CrvUsdControllerSession) SetMonetaryPolicy(monetary_policy common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetMonetaryPolicy(&_CrvUsdController.TransactOpts, monetary_policy)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) SetMonetaryPolicy(monetary_policy common.Address) (*types.Transaction, error) {
	return _CrvUsdController.Contract.SetMonetaryPolicy(&_CrvUsdController.TransactOpts, monetary_policy)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrvUsdController *CrvUsdControllerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CrvUsdController.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrvUsdController *CrvUsdControllerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Fallback(&_CrvUsdController.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrvUsdController *CrvUsdControllerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrvUsdController.Contract.Fallback(&_CrvUsdController.TransactOpts, calldata)
}

// CrvUsdControllerBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the CrvUsdController contract.
type CrvUsdControllerBorrowIterator struct {
	Event *CrvUsdControllerBorrow // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerBorrow)
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
		it.Event = new(CrvUsdControllerBorrow)
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
func (it *CrvUsdControllerBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerBorrow represents a Borrow event raised by the CrvUsdController contract.
type CrvUsdControllerBorrow struct {
	User               common.Address
	CollateralIncrease *big.Int
	LoanIncrease       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xe1979fe4c35e0cef342fef5668e2c8e7a7e9f5d5d1ca8fee0ac6c427fa4153af.
//
// Solidity: event Borrow(address indexed user, uint256 collateral_increase, uint256 loan_increase)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterBorrow(opts *bind.FilterOpts, user []common.Address) (*CrvUsdControllerBorrowIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerBorrowIterator{contract: _CrvUsdController.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xe1979fe4c35e0cef342fef5668e2c8e7a7e9f5d5d1ca8fee0ac6c427fa4153af.
//
// Solidity: event Borrow(address indexed user, uint256 collateral_increase, uint256 loan_increase)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerBorrow, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerBorrow)
				if err := _CrvUsdController.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xe1979fe4c35e0cef342fef5668e2c8e7a7e9f5d5d1ca8fee0ac6c427fa4153af.
//
// Solidity: event Borrow(address indexed user, uint256 collateral_increase, uint256 loan_increase)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseBorrow(log types.Log) (*CrvUsdControllerBorrow, error) {
	event := new(CrvUsdControllerBorrow)
	if err := _CrvUsdController.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the CrvUsdController contract.
type CrvUsdControllerCollectFeesIterator struct {
	Event *CrvUsdControllerCollectFees // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerCollectFees)
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
		it.Event = new(CrvUsdControllerCollectFees)
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
func (it *CrvUsdControllerCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerCollectFees represents a CollectFees event raised by the CrvUsdController contract.
type CrvUsdControllerCollectFees struct {
	Amount    *big.Int
	NewSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f.
//
// Solidity: event CollectFees(uint256 amount, uint256 new_supply)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterCollectFees(opts *bind.FilterOpts) (*CrvUsdControllerCollectFeesIterator, error) {

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerCollectFeesIterator{contract: _CrvUsdController.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f.
//
// Solidity: event CollectFees(uint256 amount, uint256 new_supply)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerCollectFees) (event.Subscription, error) {

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerCollectFees)
				if err := _CrvUsdController.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f.
//
// Solidity: event CollectFees(uint256 amount, uint256 new_supply)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseCollectFees(log types.Log) (*CrvUsdControllerCollectFees, error) {
	event := new(CrvUsdControllerCollectFees)
	if err := _CrvUsdController.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the CrvUsdController contract.
type CrvUsdControllerLiquidateIterator struct {
	Event *CrvUsdControllerLiquidate // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerLiquidate)
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
		it.Event = new(CrvUsdControllerLiquidate)
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
func (it *CrvUsdControllerLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerLiquidate represents a Liquidate event raised by the CrvUsdController contract.
type CrvUsdControllerLiquidate struct {
	Liquidator         common.Address
	User               common.Address
	CollateralReceived *big.Int
	StablecoinReceived *big.Int
	Debt               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed user, uint256 collateral_received, uint256 stablecoin_received, uint256 debt)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, user []common.Address) (*CrvUsdControllerLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "Liquidate", liquidatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerLiquidateIterator{contract: _CrvUsdController.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed user, uint256 collateral_received, uint256 stablecoin_received, uint256 debt)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerLiquidate, liquidator []common.Address, user []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "Liquidate", liquidatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerLiquidate)
				if err := _CrvUsdController.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed user, uint256 collateral_received, uint256 stablecoin_received, uint256 debt)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseLiquidate(log types.Log) (*CrvUsdControllerLiquidate, error) {
	event := new(CrvUsdControllerLiquidate)
	if err := _CrvUsdController.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerRemoveCollateralIterator is returned from FilterRemoveCollateral and is used to iterate over the raw logs and unpacked data for RemoveCollateral events raised by the CrvUsdController contract.
type CrvUsdControllerRemoveCollateralIterator struct {
	Event *CrvUsdControllerRemoveCollateral // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerRemoveCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerRemoveCollateral)
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
		it.Event = new(CrvUsdControllerRemoveCollateral)
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
func (it *CrvUsdControllerRemoveCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerRemoveCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerRemoveCollateral represents a RemoveCollateral event raised by the CrvUsdController contract.
type CrvUsdControllerRemoveCollateral struct {
	User               common.Address
	CollateralDecrease *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRemoveCollateral is a free log retrieval operation binding the contract event 0xe25410a4059619c9594dc6f022fe231b02aaea733f689e7ab0cd21b3d4d0eb54.
//
// Solidity: event RemoveCollateral(address indexed user, uint256 collateral_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterRemoveCollateral(opts *bind.FilterOpts, user []common.Address) (*CrvUsdControllerRemoveCollateralIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "RemoveCollateral", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerRemoveCollateralIterator{contract: _CrvUsdController.contract, event: "RemoveCollateral", logs: logs, sub: sub}, nil
}

// WatchRemoveCollateral is a free log subscription operation binding the contract event 0xe25410a4059619c9594dc6f022fe231b02aaea733f689e7ab0cd21b3d4d0eb54.
//
// Solidity: event RemoveCollateral(address indexed user, uint256 collateral_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchRemoveCollateral(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerRemoveCollateral, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "RemoveCollateral", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerRemoveCollateral)
				if err := _CrvUsdController.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
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

// ParseRemoveCollateral is a log parse operation binding the contract event 0xe25410a4059619c9594dc6f022fe231b02aaea733f689e7ab0cd21b3d4d0eb54.
//
// Solidity: event RemoveCollateral(address indexed user, uint256 collateral_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseRemoveCollateral(log types.Log) (*CrvUsdControllerRemoveCollateral, error) {
	event := new(CrvUsdControllerRemoveCollateral)
	if err := _CrvUsdController.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the CrvUsdController contract.
type CrvUsdControllerRepayIterator struct {
	Event *CrvUsdControllerRepay // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerRepay)
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
		it.Event = new(CrvUsdControllerRepay)
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
func (it *CrvUsdControllerRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerRepay represents a Repay event raised by the CrvUsdController contract.
type CrvUsdControllerRepay struct {
	User               common.Address
	CollateralDecrease *big.Int
	LoanDecrease       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 collateral_decrease, uint256 loan_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterRepay(opts *bind.FilterOpts, user []common.Address) (*CrvUsdControllerRepayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "Repay", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerRepayIterator{contract: _CrvUsdController.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 collateral_decrease, uint256 loan_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerRepay, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "Repay", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerRepay)
				if err := _CrvUsdController.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 collateral_decrease, uint256 loan_decrease)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseRepay(log types.Log) (*CrvUsdControllerRepay, error) {
	event := new(CrvUsdControllerRepay)
	if err := _CrvUsdController.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerSetBorrowingDiscountsIterator is returned from FilterSetBorrowingDiscounts and is used to iterate over the raw logs and unpacked data for SetBorrowingDiscounts events raised by the CrvUsdController contract.
type CrvUsdControllerSetBorrowingDiscountsIterator struct {
	Event *CrvUsdControllerSetBorrowingDiscounts // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerSetBorrowingDiscountsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerSetBorrowingDiscounts)
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
		it.Event = new(CrvUsdControllerSetBorrowingDiscounts)
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
func (it *CrvUsdControllerSetBorrowingDiscountsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerSetBorrowingDiscountsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerSetBorrowingDiscounts represents a SetBorrowingDiscounts event raised by the CrvUsdController contract.
type CrvUsdControllerSetBorrowingDiscounts struct {
	LoanDiscount        *big.Int
	LiquidationDiscount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowingDiscounts is a free log retrieval operation binding the contract event 0xe2750bf9a7458977fcc01c1a0b615d12162f63b18cad78441bd64c590b337eca.
//
// Solidity: event SetBorrowingDiscounts(uint256 loan_discount, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterSetBorrowingDiscounts(opts *bind.FilterOpts) (*CrvUsdControllerSetBorrowingDiscountsIterator, error) {

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "SetBorrowingDiscounts")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerSetBorrowingDiscountsIterator{contract: _CrvUsdController.contract, event: "SetBorrowingDiscounts", logs: logs, sub: sub}, nil
}

// WatchSetBorrowingDiscounts is a free log subscription operation binding the contract event 0xe2750bf9a7458977fcc01c1a0b615d12162f63b18cad78441bd64c590b337eca.
//
// Solidity: event SetBorrowingDiscounts(uint256 loan_discount, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchSetBorrowingDiscounts(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerSetBorrowingDiscounts) (event.Subscription, error) {

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "SetBorrowingDiscounts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerSetBorrowingDiscounts)
				if err := _CrvUsdController.contract.UnpackLog(event, "SetBorrowingDiscounts", log); err != nil {
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

// ParseSetBorrowingDiscounts is a log parse operation binding the contract event 0xe2750bf9a7458977fcc01c1a0b615d12162f63b18cad78441bd64c590b337eca.
//
// Solidity: event SetBorrowingDiscounts(uint256 loan_discount, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseSetBorrowingDiscounts(log types.Log) (*CrvUsdControllerSetBorrowingDiscounts, error) {
	event := new(CrvUsdControllerSetBorrowingDiscounts)
	if err := _CrvUsdController.contract.UnpackLog(event, "SetBorrowingDiscounts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerSetMonetaryPolicyIterator is returned from FilterSetMonetaryPolicy and is used to iterate over the raw logs and unpacked data for SetMonetaryPolicy events raised by the CrvUsdController contract.
type CrvUsdControllerSetMonetaryPolicyIterator struct {
	Event *CrvUsdControllerSetMonetaryPolicy // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerSetMonetaryPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerSetMonetaryPolicy)
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
		it.Event = new(CrvUsdControllerSetMonetaryPolicy)
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
func (it *CrvUsdControllerSetMonetaryPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerSetMonetaryPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerSetMonetaryPolicy represents a SetMonetaryPolicy event raised by the CrvUsdController contract.
type CrvUsdControllerSetMonetaryPolicy struct {
	MonetaryPolicy common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetMonetaryPolicy is a free log retrieval operation binding the contract event 0x51fabb88f7860c9dbcc2a5a9b69a8b9476d63b87124591f97254e29f0e8daaeb.
//
// Solidity: event SetMonetaryPolicy(address monetary_policy)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterSetMonetaryPolicy(opts *bind.FilterOpts) (*CrvUsdControllerSetMonetaryPolicyIterator, error) {

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "SetMonetaryPolicy")
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerSetMonetaryPolicyIterator{contract: _CrvUsdController.contract, event: "SetMonetaryPolicy", logs: logs, sub: sub}, nil
}

// WatchSetMonetaryPolicy is a free log subscription operation binding the contract event 0x51fabb88f7860c9dbcc2a5a9b69a8b9476d63b87124591f97254e29f0e8daaeb.
//
// Solidity: event SetMonetaryPolicy(address monetary_policy)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchSetMonetaryPolicy(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerSetMonetaryPolicy) (event.Subscription, error) {

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "SetMonetaryPolicy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerSetMonetaryPolicy)
				if err := _CrvUsdController.contract.UnpackLog(event, "SetMonetaryPolicy", log); err != nil {
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

// ParseSetMonetaryPolicy is a log parse operation binding the contract event 0x51fabb88f7860c9dbcc2a5a9b69a8b9476d63b87124591f97254e29f0e8daaeb.
//
// Solidity: event SetMonetaryPolicy(address monetary_policy)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseSetMonetaryPolicy(log types.Log) (*CrvUsdControllerSetMonetaryPolicy, error) {
	event := new(CrvUsdControllerSetMonetaryPolicy)
	if err := _CrvUsdController.contract.UnpackLog(event, "SetMonetaryPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrvUsdControllerUserStateIterator is returned from FilterUserState and is used to iterate over the raw logs and unpacked data for UserState events raised by the CrvUsdController contract.
type CrvUsdControllerUserStateIterator struct {
	Event *CrvUsdControllerUserState // Event containing the contract specifics and raw log

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
func (it *CrvUsdControllerUserStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrvUsdControllerUserState)
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
		it.Event = new(CrvUsdControllerUserState)
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
func (it *CrvUsdControllerUserStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrvUsdControllerUserStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrvUsdControllerUserState represents a UserState event raised by the CrvUsdController contract.
type CrvUsdControllerUserState struct {
	User                common.Address
	Collateral          *big.Int
	Debt                *big.Int
	N1                  *big.Int
	N2                  *big.Int
	LiquidationDiscount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUserState is a free log retrieval operation binding the contract event 0xeec6b7095a637e006c79c1819d696e353a8f703db2c49fc0219e17a8fd04f7f2.
//
// Solidity: event UserState(address indexed user, uint256 collateral, uint256 debt, int256 n1, int256 n2, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) FilterUserState(opts *bind.FilterOpts, user []common.Address) (*CrvUsdControllerUserStateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.FilterLogs(opts, "UserState", userRule)
	if err != nil {
		return nil, err
	}
	return &CrvUsdControllerUserStateIterator{contract: _CrvUsdController.contract, event: "UserState", logs: logs, sub: sub}, nil
}

// WatchUserState is a free log subscription operation binding the contract event 0xeec6b7095a637e006c79c1819d696e353a8f703db2c49fc0219e17a8fd04f7f2.
//
// Solidity: event UserState(address indexed user, uint256 collateral, uint256 debt, int256 n1, int256 n2, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) WatchUserState(opts *bind.WatchOpts, sink chan<- *CrvUsdControllerUserState, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CrvUsdController.contract.WatchLogs(opts, "UserState", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrvUsdControllerUserState)
				if err := _CrvUsdController.contract.UnpackLog(event, "UserState", log); err != nil {
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

// ParseUserState is a log parse operation binding the contract event 0xeec6b7095a637e006c79c1819d696e353a8f703db2c49fc0219e17a8fd04f7f2.
//
// Solidity: event UserState(address indexed user, uint256 collateral, uint256 debt, int256 n1, int256 n2, uint256 liquidation_discount)
func (_CrvUsdController *CrvUsdControllerFilterer) ParseUserState(log types.Log) (*CrvUsdControllerUserState, error) {
	event := new(CrvUsdControllerUserState)
	if err := _CrvUsdController.contract.UnpackLog(event, "UserState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
