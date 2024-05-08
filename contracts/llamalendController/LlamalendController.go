// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package llamalendController

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

// LlamalendControllerMetaData contains all meta data concerning the LlamalendController contract.
var LlamalendControllerMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"UserState\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"n1\",\"type\":\"int256\",\"indexed\":false},{\"name\":\"n2\",\"type\":\"int256\",\"indexed\":false},{\"name\":\"liquidation_discount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Borrow\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_increase\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loan_increase\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Repay\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_decrease\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loan_decrease\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveCollateral\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_decrease\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Liquidate\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"name\":\"collateral_received\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"stablecoin_received\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetMonetaryPolicy\",\"inputs\":[{\"name\":\"monetary_policy\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetBorrowingDiscounts\",\"inputs\":[{\"name\":\"loan_discount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"liquidation_discount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CollectFees\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"collateral_token\",\"type\":\"address\"},{\"name\":\"monetary_policy\",\"type\":\"address\"},{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"},{\"name\":\"amm\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"amm\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"collateral_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"borrowed_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"save_rate\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"debt\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_exists\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_debt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_borrowable\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"max_borrowable\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_collateral\",\"inputs\":[{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calculate_debt_n1\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_loan\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_loan_extended\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"N\",\"type\":\"uint256\"},{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_collateral\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"borrow_more\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"borrow_more_extended\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"max_active_band\",\"type\":\"int256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay\",\"inputs\":[{\"name\":\"_d_debt\",\"type\":\"uint256\"},{\"name\":\"_for\",\"type\":\"address\"},{\"name\":\"max_active_band\",\"type\":\"int256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"repay_extended\",\"inputs\":[{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health_calculator\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"d_collateral\",\"type\":\"int256\"},{\"name\":\"d_debt\",\"type\":\"int256\"},{\"name\":\"full\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health_calculator\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"d_collateral\",\"type\":\"int256\"},{\"name\":\"d_debt\",\"type\":\"int256\"},{\"name\":\"full\",\"type\":\"bool\"},{\"name\":\"N\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"liquidate_extended\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"min_x\",\"type\":\"uint256\"},{\"name\":\"frac\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"callbacker\",\"type\":\"address\"},{\"name\":\"callback_args\",\"type\":\"uint256[]\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens_to_liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokens_to_liquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"frac\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"health\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"full\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"users_to_liquidate\",\"inputs\":[{\"name\":\"_from\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\"},{\"name\":\"health\",\"type\":\"int256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"amm_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_prices\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_state\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_amm_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_amm_admin_fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_monetary_policy\",\"inputs\":[{\"name\":\"monetary_policy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_borrowing_discounts\",\"inputs\":[{\"name\":\"loan_discount\",\"type\":\"uint256\"},{\"name\":\"liquidation_discount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_callback\",\"inputs\":[{\"name\":\"cb\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"collect_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"check_lock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"liquidation_discounts\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loans\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_ix\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"n_loans\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"redeemed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"monetary_policy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"liquidation_discount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"loan_discount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// LlamalendControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use LlamalendControllerMetaData.ABI instead.
var LlamalendControllerABI = LlamalendControllerMetaData.ABI

// LlamalendController is an auto generated Go binding around an Ethereum contract.
type LlamalendController struct {
	LlamalendControllerCaller     // Read-only binding to the contract
	LlamalendControllerTransactor // Write-only binding to the contract
	LlamalendControllerFilterer   // Log filterer for contract events
}

// LlamalendControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LlamalendControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LlamalendControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LlamalendControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LlamalendControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LlamalendControllerSession struct {
	Contract     *LlamalendController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LlamalendControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LlamalendControllerCallerSession struct {
	Contract *LlamalendControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LlamalendControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LlamalendControllerTransactorSession struct {
	Contract     *LlamalendControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LlamalendControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LlamalendControllerRaw struct {
	Contract *LlamalendController // Generic contract binding to access the raw methods on
}

// LlamalendControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LlamalendControllerCallerRaw struct {
	Contract *LlamalendControllerCaller // Generic read-only contract binding to access the raw methods on
}

// LlamalendControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LlamalendControllerTransactorRaw struct {
	Contract *LlamalendControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLlamalendController creates a new instance of LlamalendController, bound to a specific deployed contract.
func NewLlamalendController(address common.Address, backend bind.ContractBackend) (*LlamalendController, error) {
	contract, err := bindLlamalendController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LlamalendController{LlamalendControllerCaller: LlamalendControllerCaller{contract: contract}, LlamalendControllerTransactor: LlamalendControllerTransactor{contract: contract}, LlamalendControllerFilterer: LlamalendControllerFilterer{contract: contract}}, nil
}

// NewLlamalendControllerCaller creates a new read-only instance of LlamalendController, bound to a specific deployed contract.
func NewLlamalendControllerCaller(address common.Address, caller bind.ContractCaller) (*LlamalendControllerCaller, error) {
	contract, err := bindLlamalendController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerCaller{contract: contract}, nil
}

// NewLlamalendControllerTransactor creates a new write-only instance of LlamalendController, bound to a specific deployed contract.
func NewLlamalendControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*LlamalendControllerTransactor, error) {
	contract, err := bindLlamalendController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerTransactor{contract: contract}, nil
}

// NewLlamalendControllerFilterer creates a new log filterer instance of LlamalendController, bound to a specific deployed contract.
func NewLlamalendControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*LlamalendControllerFilterer, error) {
	contract, err := bindLlamalendController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerFilterer{contract: contract}, nil
}

// bindLlamalendController binds a generic wrapper to an already deployed contract.
func bindLlamalendController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LlamalendControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendController *LlamalendControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendController.Contract.LlamalendControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendController *LlamalendControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendController.Contract.LlamalendControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendController *LlamalendControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendController.Contract.LlamalendControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LlamalendController *LlamalendControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LlamalendController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LlamalendController *LlamalendControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LlamalendController *LlamalendControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LlamalendController.Contract.contract.Transact(opts, method, params...)
}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) AdminFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "admin_fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) AdminFees() (*big.Int, error) {
	return _LlamalendController.Contract.AdminFees(&_LlamalendController.CallOpts)
}

// AdminFees is a free data retrieval call binding the contract method 0x1b1800e3.
//
// Solidity: function admin_fees() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) AdminFees() (*big.Int, error) {
	return _LlamalendController.Contract.AdminFees(&_LlamalendController.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() pure returns(address)
func (_LlamalendController *LlamalendControllerCaller) Amm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "amm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() pure returns(address)
func (_LlamalendController *LlamalendControllerSession) Amm() (common.Address, error) {
	return _LlamalendController.Contract.Amm(&_LlamalendController.CallOpts)
}

// Amm is a free data retrieval call binding the contract method 0x2a943945.
//
// Solidity: function amm() pure returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) Amm() (common.Address, error) {
	return _LlamalendController.Contract.Amm(&_LlamalendController.CallOpts)
}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) AmmPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "amm_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) AmmPrice() (*big.Int, error) {
	return _LlamalendController.Contract.AmmPrice(&_LlamalendController.CallOpts)
}

// AmmPrice is a free data retrieval call binding the contract method 0xd9f11a64.
//
// Solidity: function amm_price() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) AmmPrice() (*big.Int, error) {
	return _LlamalendController.Contract.AmmPrice(&_LlamalendController.CallOpts)
}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() pure returns(address)
func (_LlamalendController *LlamalendControllerCaller) BorrowedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "borrowed_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() pure returns(address)
func (_LlamalendController *LlamalendControllerSession) BorrowedToken() (common.Address, error) {
	return _LlamalendController.Contract.BorrowedToken(&_LlamalendController.CallOpts)
}

// BorrowedToken is a free data retrieval call binding the contract method 0x765337b6.
//
// Solidity: function borrowed_token() pure returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) BorrowedToken() (common.Address, error) {
	return _LlamalendController.Contract.BorrowedToken(&_LlamalendController.CallOpts)
}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerCaller) CalculateDebtN1(opts *bind.CallOpts, collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "calculate_debt_n1", collateral, debt, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerSession) CalculateDebtN1(collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.CalculateDebtN1(&_LlamalendController.CallOpts, collateral, debt, N)
}

// CalculateDebtN1 is a free data retrieval call binding the contract method 0x720fb254.
//
// Solidity: function calculate_debt_n1(uint256 collateral, uint256 debt, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerCallerSession) CalculateDebtN1(collateral *big.Int, debt *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.CalculateDebtN1(&_LlamalendController.CallOpts, collateral, debt, N)
}

// CheckLock is a free data retrieval call binding the contract method 0xd0c581bf.
//
// Solidity: function check_lock() view returns(bool)
func (_LlamalendController *LlamalendControllerCaller) CheckLock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "check_lock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckLock is a free data retrieval call binding the contract method 0xd0c581bf.
//
// Solidity: function check_lock() view returns(bool)
func (_LlamalendController *LlamalendControllerSession) CheckLock() (bool, error) {
	return _LlamalendController.Contract.CheckLock(&_LlamalendController.CallOpts)
}

// CheckLock is a free data retrieval call binding the contract method 0xd0c581bf.
//
// Solidity: function check_lock() view returns(bool)
func (_LlamalendController *LlamalendControllerCallerSession) CheckLock() (bool, error) {
	return _LlamalendController.Contract.CheckLock(&_LlamalendController.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() pure returns(address)
func (_LlamalendController *LlamalendControllerCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "collateral_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() pure returns(address)
func (_LlamalendController *LlamalendControllerSession) CollateralToken() (common.Address, error) {
	return _LlamalendController.Contract.CollateralToken(&_LlamalendController.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0x2621db2f.
//
// Solidity: function collateral_token() pure returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) CollateralToken() (common.Address, error) {
	return _LlamalendController.Contract.CollateralToken(&_LlamalendController.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) Debt(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "debt", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) Debt(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.Debt(&_LlamalendController.CallOpts, user)
}

// Debt is a free data retrieval call binding the contract method 0x9b6c56ec.
//
// Solidity: function debt(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) Debt(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.Debt(&_LlamalendController.CallOpts, user)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_LlamalendController *LlamalendControllerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_LlamalendController *LlamalendControllerSession) Factory() (common.Address, error) {
	return _LlamalendController.Contract.Factory(&_LlamalendController.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) Factory() (common.Address, error) {
	return _LlamalendController.Contract.Factory(&_LlamalendController.CallOpts)
}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_LlamalendController *LlamalendControllerCaller) Health(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "health", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_LlamalendController *LlamalendControllerSession) Health(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.Health(&_LlamalendController.CallOpts, user)
}

// Health is a free data retrieval call binding the contract method 0xe2d8ebee.
//
// Solidity: function health(address user) view returns(int256)
func (_LlamalendController *LlamalendControllerCallerSession) Health(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.Health(&_LlamalendController.CallOpts, user)
}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerCaller) Health0(opts *bind.CallOpts, user common.Address, full bool) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "health0", user, full)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerSession) Health0(user common.Address, full bool) (*big.Int, error) {
	return _LlamalendController.Contract.Health0(&_LlamalendController.CallOpts, user, full)
}

// Health0 is a free data retrieval call binding the contract method 0x8908ea82.
//
// Solidity: function health(address user, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerCallerSession) Health0(user common.Address, full bool) (*big.Int, error) {
	return _LlamalendController.Contract.Health0(&_LlamalendController.CallOpts, user, full)
}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerCaller) HealthCalculator(opts *bind.CallOpts, user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "health_calculator", user, d_collateral, d_debt, full)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerSession) HealthCalculator(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	return _LlamalendController.Contract.HealthCalculator(&_LlamalendController.CallOpts, user, d_collateral, d_debt, full)
}

// HealthCalculator is a free data retrieval call binding the contract method 0x0b8db681.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full) view returns(int256)
func (_LlamalendController *LlamalendControllerCallerSession) HealthCalculator(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool) (*big.Int, error) {
	return _LlamalendController.Contract.HealthCalculator(&_LlamalendController.CallOpts, user, d_collateral, d_debt, full)
}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerCaller) HealthCalculator0(opts *bind.CallOpts, user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "health_calculator0", user, d_collateral, d_debt, full, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerSession) HealthCalculator0(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.HealthCalculator0(&_LlamalendController.CallOpts, user, d_collateral, d_debt, full, N)
}

// HealthCalculator0 is a free data retrieval call binding the contract method 0x22c71453.
//
// Solidity: function health_calculator(address user, int256 d_collateral, int256 d_debt, bool full, uint256 N) view returns(int256)
func (_LlamalendController *LlamalendControllerCallerSession) HealthCalculator0(user common.Address, d_collateral *big.Int, d_debt *big.Int, full bool, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.HealthCalculator0(&_LlamalendController.CallOpts, user, d_collateral, d_debt, full, N)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) LiquidationDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "liquidation_discount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) LiquidationDiscount() (*big.Int, error) {
	return _LlamalendController.Contract.LiquidationDiscount(&_LlamalendController.CallOpts)
}

// LiquidationDiscount is a free data retrieval call binding the contract method 0x627d2b83.
//
// Solidity: function liquidation_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) LiquidationDiscount() (*big.Int, error) {
	return _LlamalendController.Contract.LiquidationDiscount(&_LlamalendController.CallOpts)
}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) LiquidationDiscounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "liquidation_discounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) LiquidationDiscounts(arg0 common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.LiquidationDiscounts(&_LlamalendController.CallOpts, arg0)
}

// LiquidationDiscounts is a free data retrieval call binding the contract method 0x5457ff7b.
//
// Solidity: function liquidation_discounts(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) LiquidationDiscounts(arg0 common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.LiquidationDiscounts(&_LlamalendController.CallOpts, arg0)
}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) LoanDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "loan_discount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) LoanDiscount() (*big.Int, error) {
	return _LlamalendController.Contract.LoanDiscount(&_LlamalendController.CallOpts)
}

// LoanDiscount is a free data retrieval call binding the contract method 0x5449b9cb.
//
// Solidity: function loan_discount() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) LoanDiscount() (*big.Int, error) {
	return _LlamalendController.Contract.LoanDiscount(&_LlamalendController.CallOpts)
}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_LlamalendController *LlamalendControllerCaller) LoanExists(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "loan_exists", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_LlamalendController *LlamalendControllerSession) LoanExists(user common.Address) (bool, error) {
	return _LlamalendController.Contract.LoanExists(&_LlamalendController.CallOpts, user)
}

// LoanExists is a free data retrieval call binding the contract method 0xa21adb9e.
//
// Solidity: function loan_exists(address user) view returns(bool)
func (_LlamalendController *LlamalendControllerCallerSession) LoanExists(user common.Address) (bool, error) {
	return _LlamalendController.Contract.LoanExists(&_LlamalendController.CallOpts, user)
}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) LoanIx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "loan_ix", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) LoanIx(arg0 common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.LoanIx(&_LlamalendController.CallOpts, arg0)
}

// LoanIx is a free data retrieval call binding the contract method 0x7128f3b8.
//
// Solidity: function loan_ix(address arg0) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) LoanIx(arg0 common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.LoanIx(&_LlamalendController.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_LlamalendController *LlamalendControllerCaller) Loans(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "loans", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_LlamalendController *LlamalendControllerSession) Loans(arg0 *big.Int) (common.Address, error) {
	return _LlamalendController.Contract.Loans(&_LlamalendController.CallOpts, arg0)
}

// Loans is a free data retrieval call binding the contract method 0xe1ec3c68.
//
// Solidity: function loans(uint256 arg0) view returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) Loans(arg0 *big.Int) (common.Address, error) {
	return _LlamalendController.Contract.Loans(&_LlamalendController.CallOpts, arg0)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) MaxBorrowable(opts *bind.CallOpts, collateral *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "max_borrowable", collateral, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) MaxBorrowable(collateral *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MaxBorrowable(&_LlamalendController.CallOpts, collateral, N)
}

// MaxBorrowable is a free data retrieval call binding the contract method 0x9a497196.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) MaxBorrowable(collateral *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MaxBorrowable(&_LlamalendController.CallOpts, collateral, N)
}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) MaxBorrowable0(opts *bind.CallOpts, collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "max_borrowable0", collateral, N, current_debt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) MaxBorrowable0(collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MaxBorrowable0(&_LlamalendController.CallOpts, collateral, N, current_debt)
}

// MaxBorrowable0 is a free data retrieval call binding the contract method 0x1cf1f947.
//
// Solidity: function max_borrowable(uint256 collateral, uint256 N, uint256 current_debt) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) MaxBorrowable0(collateral *big.Int, N *big.Int, current_debt *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MaxBorrowable0(&_LlamalendController.CallOpts, collateral, N, current_debt)
}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) MinCollateral(opts *bind.CallOpts, debt *big.Int, N *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "min_collateral", debt, N)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) MinCollateral(debt *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MinCollateral(&_LlamalendController.CallOpts, debt, N)
}

// MinCollateral is a free data retrieval call binding the contract method 0xa7573206.
//
// Solidity: function min_collateral(uint256 debt, uint256 N) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) MinCollateral(debt *big.Int, N *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.MinCollateral(&_LlamalendController.CallOpts, debt, N)
}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) Minted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "minted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) Minted() (*big.Int, error) {
	return _LlamalendController.Contract.Minted(&_LlamalendController.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) Minted() (*big.Int, error) {
	return _LlamalendController.Contract.Minted(&_LlamalendController.CallOpts)
}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_LlamalendController *LlamalendControllerCaller) MonetaryPolicy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "monetary_policy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_LlamalendController *LlamalendControllerSession) MonetaryPolicy() (common.Address, error) {
	return _LlamalendController.Contract.MonetaryPolicy(&_LlamalendController.CallOpts)
}

// MonetaryPolicy is a free data retrieval call binding the contract method 0xadfae4ce.
//
// Solidity: function monetary_policy() view returns(address)
func (_LlamalendController *LlamalendControllerCallerSession) MonetaryPolicy() (common.Address, error) {
	return _LlamalendController.Contract.MonetaryPolicy(&_LlamalendController.CallOpts)
}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) NLoans(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "n_loans")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) NLoans() (*big.Int, error) {
	return _LlamalendController.Contract.NLoans(&_LlamalendController.CallOpts)
}

// NLoans is a free data retrieval call binding the contract method 0x6cce39be.
//
// Solidity: function n_loans() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) NLoans() (*big.Int, error) {
	return _LlamalendController.Contract.NLoans(&_LlamalendController.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) Redeemed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "redeemed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) Redeemed() (*big.Int, error) {
	return _LlamalendController.Contract.Redeemed(&_LlamalendController.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0xe231bff0.
//
// Solidity: function redeemed() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) Redeemed() (*big.Int, error) {
	return _LlamalendController.Contract.Redeemed(&_LlamalendController.CallOpts)
}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) TokensToLiquidate(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "tokens_to_liquidate", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) TokensToLiquidate(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.TokensToLiquidate(&_LlamalendController.CallOpts, user)
}

// TokensToLiquidate is a free data retrieval call binding the contract method 0x1b25cdaf.
//
// Solidity: function tokens_to_liquidate(address user) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) TokensToLiquidate(user common.Address) (*big.Int, error) {
	return _LlamalendController.Contract.TokensToLiquidate(&_LlamalendController.CallOpts, user)
}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) TokensToLiquidate0(opts *bind.CallOpts, user common.Address, frac *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "tokens_to_liquidate0", user, frac)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) TokensToLiquidate0(user common.Address, frac *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.TokensToLiquidate0(&_LlamalendController.CallOpts, user, frac)
}

// TokensToLiquidate0 is a free data retrieval call binding the contract method 0x546e040d.
//
// Solidity: function tokens_to_liquidate(address user, uint256 frac) view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) TokensToLiquidate0(user common.Address, frac *big.Int) (*big.Int, error) {
	return _LlamalendController.Contract.TokensToLiquidate0(&_LlamalendController.CallOpts, user, frac)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_LlamalendController *LlamalendControllerCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "total_debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_LlamalendController *LlamalendControllerSession) TotalDebt() (*big.Int, error) {
	return _LlamalendController.Contract.TotalDebt(&_LlamalendController.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0x31dc3ca8.
//
// Solidity: function total_debt() view returns(uint256)
func (_LlamalendController *LlamalendControllerCallerSession) TotalDebt() (*big.Int, error) {
	return _LlamalendController.Contract.TotalDebt(&_LlamalendController.CallOpts)
}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_LlamalendController *LlamalendControllerCaller) UserPrices(opts *bind.CallOpts, user common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "user_prices", user)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_LlamalendController *LlamalendControllerSession) UserPrices(user common.Address) ([2]*big.Int, error) {
	return _LlamalendController.Contract.UserPrices(&_LlamalendController.CallOpts, user)
}

// UserPrices is a free data retrieval call binding the contract method 0x2c5089c3.
//
// Solidity: function user_prices(address user) view returns(uint256[2])
func (_LlamalendController *LlamalendControllerCallerSession) UserPrices(user common.Address) ([2]*big.Int, error) {
	return _LlamalendController.Contract.UserPrices(&_LlamalendController.CallOpts, user)
}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_LlamalendController *LlamalendControllerCaller) UserState(opts *bind.CallOpts, user common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "user_state", user)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_LlamalendController *LlamalendControllerSession) UserState(user common.Address) ([4]*big.Int, error) {
	return _LlamalendController.Contract.UserState(&_LlamalendController.CallOpts, user)
}

// UserState is a free data retrieval call binding the contract method 0xec74d0a8.
//
// Solidity: function user_state(address user) view returns(uint256[4])
func (_LlamalendController *LlamalendControllerCallerSession) UserState(user common.Address) ([4]*big.Int, error) {
	return _LlamalendController.Contract.UserState(&_LlamalendController.CallOpts, user)
}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCaller) UsersToLiquidate(opts *bind.CallOpts) ([]Struct0, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "users_to_liquidate")

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerSession) UsersToLiquidate() ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate(&_LlamalendController.CallOpts)
}

// UsersToLiquidate is a free data retrieval call binding the contract method 0x007c98ab.
//
// Solidity: function users_to_liquidate() view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCallerSession) UsersToLiquidate() ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate(&_LlamalendController.CallOpts)
}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCaller) UsersToLiquidate0(opts *bind.CallOpts, _from *big.Int) ([]Struct0, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "users_to_liquidate0", _from)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerSession) UsersToLiquidate0(_from *big.Int) ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate0(&_LlamalendController.CallOpts, _from)
}

// UsersToLiquidate0 is a free data retrieval call binding the contract method 0x80e8f6ec.
//
// Solidity: function users_to_liquidate(uint256 _from) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCallerSession) UsersToLiquidate0(_from *big.Int) ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate0(&_LlamalendController.CallOpts, _from)
}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCaller) UsersToLiquidate1(opts *bind.CallOpts, _from *big.Int, _limit *big.Int) ([]Struct0, error) {
	var out []interface{}
	err := _LlamalendController.contract.Call(opts, &out, "users_to_liquidate1", _from, _limit)

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerSession) UsersToLiquidate1(_from *big.Int, _limit *big.Int) ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate1(&_LlamalendController.CallOpts, _from, _limit)
}

// UsersToLiquidate1 is a free data retrieval call binding the contract method 0x90f8667d.
//
// Solidity: function users_to_liquidate(uint256 _from, uint256 _limit) view returns((address,uint256,uint256,uint256,int256)[])
func (_LlamalendController *LlamalendControllerCallerSession) UsersToLiquidate1(_from *big.Int, _limit *big.Int) ([]Struct0, error) {
	return _LlamalendController.Contract.UsersToLiquidate1(&_LlamalendController.CallOpts, _from, _limit)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerTransactor) AddCollateral(opts *bind.TransactOpts, collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "add_collateral", collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerSession) AddCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.AddCollateral(&_LlamalendController.TransactOpts, collateral)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6f972f12.
//
// Solidity: function add_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) AddCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.AddCollateral(&_LlamalendController.TransactOpts, collateral)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) returns()
func (_LlamalendController *LlamalendControllerTransactor) AddCollateral0(opts *bind.TransactOpts, collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "add_collateral0", collateral, _for)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) returns()
func (_LlamalendController *LlamalendControllerSession) AddCollateral0(collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.AddCollateral0(&_LlamalendController.TransactOpts, collateral, _for)
}

// AddCollateral0 is a paid mutator transaction binding the contract method 0x24049e57.
//
// Solidity: function add_collateral(uint256 collateral, address _for) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) AddCollateral0(collateral *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.AddCollateral0(&_LlamalendController.TransactOpts, collateral, _for)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) returns()
func (_LlamalendController *LlamalendControllerTransactor) BorrowMore(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "borrow_more", collateral, debt)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) returns()
func (_LlamalendController *LlamalendControllerSession) BorrowMore(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.BorrowMore(&_LlamalendController.TransactOpts, collateral, debt)
}

// BorrowMore is a paid mutator transaction binding the contract method 0xdd171e7c.
//
// Solidity: function borrow_more(uint256 collateral, uint256 debt) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) BorrowMore(collateral *big.Int, debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.BorrowMore(&_LlamalendController.TransactOpts, collateral, debt)
}

// BorrowMoreExtended is a paid mutator transaction binding the contract method 0x36b7dbb7.
//
// Solidity: function borrow_more_extended(uint256 collateral, uint256 debt, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactor) BorrowMoreExtended(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "borrow_more_extended", collateral, debt, callbacker, callback_args)
}

// BorrowMoreExtended is a paid mutator transaction binding the contract method 0x36b7dbb7.
//
// Solidity: function borrow_more_extended(uint256 collateral, uint256 debt, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerSession) BorrowMoreExtended(collateral *big.Int, debt *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.BorrowMoreExtended(&_LlamalendController.TransactOpts, collateral, debt, callbacker, callback_args)
}

// BorrowMoreExtended is a paid mutator transaction binding the contract method 0x36b7dbb7.
//
// Solidity: function borrow_more_extended(uint256 collateral, uint256 debt, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) BorrowMoreExtended(collateral *big.Int, debt *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.BorrowMoreExtended(&_LlamalendController.TransactOpts, collateral, debt, callbacker, callback_args)
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_LlamalendController *LlamalendControllerTransactor) CollectFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "collect_fees")
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_LlamalendController *LlamalendControllerSession) CollectFees() (*types.Transaction, error) {
	return _LlamalendController.Contract.CollectFees(&_LlamalendController.TransactOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0x1e0cfcef.
//
// Solidity: function collect_fees() returns(uint256)
func (_LlamalendController *LlamalendControllerTransactorSession) CollectFees() (*types.Transaction, error) {
	return _LlamalendController.Contract.CollectFees(&_LlamalendController.TransactOpts)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) returns()
func (_LlamalendController *LlamalendControllerTransactor) CreateLoan(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "create_loan", collateral, debt, N)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) returns()
func (_LlamalendController *LlamalendControllerSession) CreateLoan(collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.CreateLoan(&_LlamalendController.TransactOpts, collateral, debt, N)
}

// CreateLoan is a paid mutator transaction binding the contract method 0x23cfed03.
//
// Solidity: function create_loan(uint256 collateral, uint256 debt, uint256 N) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) CreateLoan(collateral *big.Int, debt *big.Int, N *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.CreateLoan(&_LlamalendController.TransactOpts, collateral, debt, N)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactor) CreateLoanExtended(opts *bind.TransactOpts, collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "create_loan_extended", collateral, debt, N, callbacker, callback_args)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerSession) CreateLoanExtended(collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.CreateLoanExtended(&_LlamalendController.TransactOpts, collateral, debt, N, callbacker, callback_args)
}

// CreateLoanExtended is a paid mutator transaction binding the contract method 0xbc61ea23.
//
// Solidity: function create_loan_extended(uint256 collateral, uint256 debt, uint256 N, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) CreateLoanExtended(collateral *big.Int, debt *big.Int, N *big.Int, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.CreateLoanExtended(&_LlamalendController.TransactOpts, collateral, debt, N, callbacker, callback_args)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_LlamalendController *LlamalendControllerTransactor) Liquidate(opts *bind.TransactOpts, user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "liquidate", user, min_x)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_LlamalendController *LlamalendControllerSession) Liquidate(user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Liquidate(&_LlamalendController.TransactOpts, user, min_x)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address user, uint256 min_x) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Liquidate(user common.Address, min_x *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Liquidate(&_LlamalendController.TransactOpts, user, min_x)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactor) Liquidate0(opts *bind.TransactOpts, user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "liquidate0", user, min_x, use_eth)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerSession) Liquidate0(user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.Liquidate0(&_LlamalendController.TransactOpts, user, min_x, use_eth)
}

// Liquidate0 is a paid mutator transaction binding the contract method 0x3ecdb828.
//
// Solidity: function liquidate(address user, uint256 min_x, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Liquidate0(user common.Address, min_x *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.Liquidate0(&_LlamalendController.TransactOpts, user, min_x, use_eth)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactor) LiquidateExtended(opts *bind.TransactOpts, user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "liquidate_extended", user, min_x, frac, use_eth, callbacker, callback_args)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerSession) LiquidateExtended(user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.LiquidateExtended(&_LlamalendController.TransactOpts, user, min_x, frac, use_eth, callbacker, callback_args)
}

// LiquidateExtended is a paid mutator transaction binding the contract method 0x036aed88.
//
// Solidity: function liquidate_extended(address user, uint256 min_x, uint256 frac, bool use_eth, address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) LiquidateExtended(user common.Address, min_x *big.Int, frac *big.Int, use_eth bool, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.LiquidateExtended(&_LlamalendController.TransactOpts, user, min_x, frac, use_eth, callbacker, callback_args)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerTransactor) RemoveCollateral(opts *bind.TransactOpts, collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "remove_collateral", collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerSession) RemoveCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.RemoveCollateral(&_LlamalendController.TransactOpts, collateral)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0xd14ff5b6.
//
// Solidity: function remove_collateral(uint256 collateral) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) RemoveCollateral(collateral *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.RemoveCollateral(&_LlamalendController.TransactOpts, collateral)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactor) RemoveCollateral0(opts *bind.TransactOpts, collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "remove_collateral0", collateral, use_eth)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerSession) RemoveCollateral0(collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.RemoveCollateral0(&_LlamalendController.TransactOpts, collateral, use_eth)
}

// RemoveCollateral0 is a paid mutator transaction binding the contract method 0x2e4af52a.
//
// Solidity: function remove_collateral(uint256 collateral, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) RemoveCollateral0(collateral *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.RemoveCollateral0(&_LlamalendController.TransactOpts, collateral, use_eth)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_LlamalendController *LlamalendControllerTransactor) Repay(opts *bind.TransactOpts, _d_debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "repay", _d_debt)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_LlamalendController *LlamalendControllerSession) Repay(_d_debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay(&_LlamalendController.TransactOpts, _d_debt)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _d_debt) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Repay(_d_debt *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay(&_LlamalendController.TransactOpts, _d_debt)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_LlamalendController *LlamalendControllerTransactor) Repay0(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "repay0", _d_debt, _for)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_LlamalendController *LlamalendControllerSession) Repay0(_d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay0(&_LlamalendController.TransactOpts, _d_debt, _for)
}

// Repay0 is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 _d_debt, address _for) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Repay0(_d_debt *big.Int, _for common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay0(&_LlamalendController.TransactOpts, _d_debt, _for)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_LlamalendController *LlamalendControllerTransactor) Repay1(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "repay1", _d_debt, _for, max_active_band)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_LlamalendController *LlamalendControllerSession) Repay1(_d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay1(&_LlamalendController.TransactOpts, _d_debt, _for, max_active_band)
}

// Repay1 is a paid mutator transaction binding the contract method 0xb4440df4.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Repay1(_d_debt *big.Int, _for common.Address, max_active_band *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay1(&_LlamalendController.TransactOpts, _d_debt, _for, max_active_band)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactor) Repay2(opts *bind.TransactOpts, _d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "repay2", _d_debt, _for, max_active_band, use_eth)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerSession) Repay2(_d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay2(&_LlamalendController.TransactOpts, _d_debt, _for, max_active_band, use_eth)
}

// Repay2 is a paid mutator transaction binding the contract method 0x37671f93.
//
// Solidity: function repay(uint256 _d_debt, address _for, int256 max_active_band, bool use_eth) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) Repay2(_d_debt *big.Int, _for common.Address, max_active_band *big.Int, use_eth bool) (*types.Transaction, error) {
	return _LlamalendController.Contract.Repay2(&_LlamalendController.TransactOpts, _d_debt, _for, max_active_band, use_eth)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactor) RepayExtended(opts *bind.TransactOpts, callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "repay_extended", callbacker, callback_args)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerSession) RepayExtended(callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.RepayExtended(&_LlamalendController.TransactOpts, callbacker, callback_args)
}

// RepayExtended is a paid mutator transaction binding the contract method 0x152f65cb.
//
// Solidity: function repay_extended(address callbacker, uint256[] callback_args) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) RepayExtended(callbacker common.Address, callback_args []*big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.RepayExtended(&_LlamalendController.TransactOpts, callbacker, callback_args)
}

// SaveRate is a paid mutator transaction binding the contract method 0x69c6804e.
//
// Solidity: function save_rate() returns()
func (_LlamalendController *LlamalendControllerTransactor) SaveRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "save_rate")
}

// SaveRate is a paid mutator transaction binding the contract method 0x69c6804e.
//
// Solidity: function save_rate() returns()
func (_LlamalendController *LlamalendControllerSession) SaveRate() (*types.Transaction, error) {
	return _LlamalendController.Contract.SaveRate(&_LlamalendController.TransactOpts)
}

// SaveRate is a paid mutator transaction binding the contract method 0x69c6804e.
//
// Solidity: function save_rate() returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SaveRate() (*types.Transaction, error) {
	return _LlamalendController.Contract.SaveRate(&_LlamalendController.TransactOpts)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerTransactor) SetAmmAdminFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "set_amm_admin_fee", fee)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerSession) SetAmmAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetAmmAdminFee(&_LlamalendController.TransactOpts, fee)
}

// SetAmmAdminFee is a paid mutator transaction binding the contract method 0xa5b4804a.
//
// Solidity: function set_amm_admin_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SetAmmAdminFee(fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetAmmAdminFee(&_LlamalendController.TransactOpts, fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerTransactor) SetAmmFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "set_amm_fee", fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerSession) SetAmmFee(fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetAmmFee(&_LlamalendController.TransactOpts, fee)
}

// SetAmmFee is a paid mutator transaction binding the contract method 0x4189617d.
//
// Solidity: function set_amm_fee(uint256 fee) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SetAmmFee(fee *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetAmmFee(&_LlamalendController.TransactOpts, fee)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_LlamalendController *LlamalendControllerTransactor) SetBorrowingDiscounts(opts *bind.TransactOpts, loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "set_borrowing_discounts", loan_discount, liquidation_discount)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_LlamalendController *LlamalendControllerSession) SetBorrowingDiscounts(loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetBorrowingDiscounts(&_LlamalendController.TransactOpts, loan_discount, liquidation_discount)
}

// SetBorrowingDiscounts is a paid mutator transaction binding the contract method 0x2a0c3586.
//
// Solidity: function set_borrowing_discounts(uint256 loan_discount, uint256 liquidation_discount) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SetBorrowingDiscounts(loan_discount *big.Int, liquidation_discount *big.Int) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetBorrowingDiscounts(&_LlamalendController.TransactOpts, loan_discount, liquidation_discount)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_LlamalendController *LlamalendControllerTransactor) SetCallback(opts *bind.TransactOpts, cb common.Address) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "set_callback", cb)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_LlamalendController *LlamalendControllerSession) SetCallback(cb common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetCallback(&_LlamalendController.TransactOpts, cb)
}

// SetCallback is a paid mutator transaction binding the contract method 0xcc1891c7.
//
// Solidity: function set_callback(address cb) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SetCallback(cb common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetCallback(&_LlamalendController.TransactOpts, cb)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_LlamalendController *LlamalendControllerTransactor) SetMonetaryPolicy(opts *bind.TransactOpts, monetary_policy common.Address) (*types.Transaction, error) {
	return _LlamalendController.contract.Transact(opts, "set_monetary_policy", monetary_policy)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_LlamalendController *LlamalendControllerSession) SetMonetaryPolicy(monetary_policy common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetMonetaryPolicy(&_LlamalendController.TransactOpts, monetary_policy)
}

// SetMonetaryPolicy is a paid mutator transaction binding the contract method 0x81d2f1b7.
//
// Solidity: function set_monetary_policy(address monetary_policy) returns()
func (_LlamalendController *LlamalendControllerTransactorSession) SetMonetaryPolicy(monetary_policy common.Address) (*types.Transaction, error) {
	return _LlamalendController.Contract.SetMonetaryPolicy(&_LlamalendController.TransactOpts, monetary_policy)
}

// LlamalendControllerBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LlamalendController contract.
type LlamalendControllerBorrowIterator struct {
	Event *LlamalendControllerBorrow // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerBorrow)
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
		it.Event = new(LlamalendControllerBorrow)
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
func (it *LlamalendControllerBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerBorrow represents a Borrow event raised by the LlamalendController contract.
type LlamalendControllerBorrow struct {
	User               common.Address
	CollateralIncrease *big.Int
	LoanIncrease       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xe1979fe4c35e0cef342fef5668e2c8e7a7e9f5d5d1ca8fee0ac6c427fa4153af.
//
// Solidity: event Borrow(address indexed user, uint256 collateral_increase, uint256 loan_increase)
func (_LlamalendController *LlamalendControllerFilterer) FilterBorrow(opts *bind.FilterOpts, user []common.Address) (*LlamalendControllerBorrowIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerBorrowIterator{contract: _LlamalendController.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xe1979fe4c35e0cef342fef5668e2c8e7a7e9f5d5d1ca8fee0ac6c427fa4153af.
//
// Solidity: event Borrow(address indexed user, uint256 collateral_increase, uint256 loan_increase)
func (_LlamalendController *LlamalendControllerFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LlamalendControllerBorrow, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerBorrow)
				if err := _LlamalendController.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseBorrow(log types.Log) (*LlamalendControllerBorrow, error) {
	event := new(LlamalendControllerBorrow)
	if err := _LlamalendController.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the LlamalendController contract.
type LlamalendControllerCollectFeesIterator struct {
	Event *LlamalendControllerCollectFees // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerCollectFees)
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
		it.Event = new(LlamalendControllerCollectFees)
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
func (it *LlamalendControllerCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerCollectFees represents a CollectFees event raised by the LlamalendController contract.
type LlamalendControllerCollectFees struct {
	Amount    *big.Int
	NewSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f.
//
// Solidity: event CollectFees(uint256 amount, uint256 new_supply)
func (_LlamalendController *LlamalendControllerFilterer) FilterCollectFees(opts *bind.FilterOpts) (*LlamalendControllerCollectFeesIterator, error) {

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerCollectFeesIterator{contract: _LlamalendController.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f.
//
// Solidity: event CollectFees(uint256 amount, uint256 new_supply)
func (_LlamalendController *LlamalendControllerFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *LlamalendControllerCollectFees) (event.Subscription, error) {

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerCollectFees)
				if err := _LlamalendController.contract.UnpackLog(event, "CollectFees", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseCollectFees(log types.Log) (*LlamalendControllerCollectFees, error) {
	event := new(LlamalendControllerCollectFees)
	if err := _LlamalendController.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the LlamalendController contract.
type LlamalendControllerLiquidateIterator struct {
	Event *LlamalendControllerLiquidate // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerLiquidate)
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
		it.Event = new(LlamalendControllerLiquidate)
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
func (it *LlamalendControllerLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerLiquidate represents a Liquidate event raised by the LlamalendController contract.
type LlamalendControllerLiquidate struct {
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
func (_LlamalendController *LlamalendControllerFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, user []common.Address) (*LlamalendControllerLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "Liquidate", liquidatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerLiquidateIterator{contract: _LlamalendController.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed user, uint256 collateral_received, uint256 stablecoin_received, uint256 debt)
func (_LlamalendController *LlamalendControllerFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *LlamalendControllerLiquidate, liquidator []common.Address, user []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "Liquidate", liquidatorRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerLiquidate)
				if err := _LlamalendController.contract.UnpackLog(event, "Liquidate", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseLiquidate(log types.Log) (*LlamalendControllerLiquidate, error) {
	event := new(LlamalendControllerLiquidate)
	if err := _LlamalendController.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerRemoveCollateralIterator is returned from FilterRemoveCollateral and is used to iterate over the raw logs and unpacked data for RemoveCollateral events raised by the LlamalendController contract.
type LlamalendControllerRemoveCollateralIterator struct {
	Event *LlamalendControllerRemoveCollateral // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerRemoveCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerRemoveCollateral)
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
		it.Event = new(LlamalendControllerRemoveCollateral)
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
func (it *LlamalendControllerRemoveCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerRemoveCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerRemoveCollateral represents a RemoveCollateral event raised by the LlamalendController contract.
type LlamalendControllerRemoveCollateral struct {
	User               common.Address
	CollateralDecrease *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRemoveCollateral is a free log retrieval operation binding the contract event 0xe25410a4059619c9594dc6f022fe231b02aaea733f689e7ab0cd21b3d4d0eb54.
//
// Solidity: event RemoveCollateral(address indexed user, uint256 collateral_decrease)
func (_LlamalendController *LlamalendControllerFilterer) FilterRemoveCollateral(opts *bind.FilterOpts, user []common.Address) (*LlamalendControllerRemoveCollateralIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "RemoveCollateral", userRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerRemoveCollateralIterator{contract: _LlamalendController.contract, event: "RemoveCollateral", logs: logs, sub: sub}, nil
}

// WatchRemoveCollateral is a free log subscription operation binding the contract event 0xe25410a4059619c9594dc6f022fe231b02aaea733f689e7ab0cd21b3d4d0eb54.
//
// Solidity: event RemoveCollateral(address indexed user, uint256 collateral_decrease)
func (_LlamalendController *LlamalendControllerFilterer) WatchRemoveCollateral(opts *bind.WatchOpts, sink chan<- *LlamalendControllerRemoveCollateral, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "RemoveCollateral", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerRemoveCollateral)
				if err := _LlamalendController.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseRemoveCollateral(log types.Log) (*LlamalendControllerRemoveCollateral, error) {
	event := new(LlamalendControllerRemoveCollateral)
	if err := _LlamalendController.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the LlamalendController contract.
type LlamalendControllerRepayIterator struct {
	Event *LlamalendControllerRepay // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerRepay)
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
		it.Event = new(LlamalendControllerRepay)
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
func (it *LlamalendControllerRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerRepay represents a Repay event raised by the LlamalendController contract.
type LlamalendControllerRepay struct {
	User               common.Address
	CollateralDecrease *big.Int
	LoanDecrease       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 collateral_decrease, uint256 loan_decrease)
func (_LlamalendController *LlamalendControllerFilterer) FilterRepay(opts *bind.FilterOpts, user []common.Address) (*LlamalendControllerRepayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "Repay", userRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerRepayIterator{contract: _LlamalendController.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 collateral_decrease, uint256 loan_decrease)
func (_LlamalendController *LlamalendControllerFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LlamalendControllerRepay, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "Repay", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerRepay)
				if err := _LlamalendController.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseRepay(log types.Log) (*LlamalendControllerRepay, error) {
	event := new(LlamalendControllerRepay)
	if err := _LlamalendController.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerSetBorrowingDiscountsIterator is returned from FilterSetBorrowingDiscounts and is used to iterate over the raw logs and unpacked data for SetBorrowingDiscounts events raised by the LlamalendController contract.
type LlamalendControllerSetBorrowingDiscountsIterator struct {
	Event *LlamalendControllerSetBorrowingDiscounts // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerSetBorrowingDiscountsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerSetBorrowingDiscounts)
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
		it.Event = new(LlamalendControllerSetBorrowingDiscounts)
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
func (it *LlamalendControllerSetBorrowingDiscountsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerSetBorrowingDiscountsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerSetBorrowingDiscounts represents a SetBorrowingDiscounts event raised by the LlamalendController contract.
type LlamalendControllerSetBorrowingDiscounts struct {
	LoanDiscount        *big.Int
	LiquidationDiscount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowingDiscounts is a free log retrieval operation binding the contract event 0xe2750bf9a7458977fcc01c1a0b615d12162f63b18cad78441bd64c590b337eca.
//
// Solidity: event SetBorrowingDiscounts(uint256 loan_discount, uint256 liquidation_discount)
func (_LlamalendController *LlamalendControllerFilterer) FilterSetBorrowingDiscounts(opts *bind.FilterOpts) (*LlamalendControllerSetBorrowingDiscountsIterator, error) {

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "SetBorrowingDiscounts")
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerSetBorrowingDiscountsIterator{contract: _LlamalendController.contract, event: "SetBorrowingDiscounts", logs: logs, sub: sub}, nil
}

// WatchSetBorrowingDiscounts is a free log subscription operation binding the contract event 0xe2750bf9a7458977fcc01c1a0b615d12162f63b18cad78441bd64c590b337eca.
//
// Solidity: event SetBorrowingDiscounts(uint256 loan_discount, uint256 liquidation_discount)
func (_LlamalendController *LlamalendControllerFilterer) WatchSetBorrowingDiscounts(opts *bind.WatchOpts, sink chan<- *LlamalendControllerSetBorrowingDiscounts) (event.Subscription, error) {

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "SetBorrowingDiscounts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerSetBorrowingDiscounts)
				if err := _LlamalendController.contract.UnpackLog(event, "SetBorrowingDiscounts", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseSetBorrowingDiscounts(log types.Log) (*LlamalendControllerSetBorrowingDiscounts, error) {
	event := new(LlamalendControllerSetBorrowingDiscounts)
	if err := _LlamalendController.contract.UnpackLog(event, "SetBorrowingDiscounts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerSetMonetaryPolicyIterator is returned from FilterSetMonetaryPolicy and is used to iterate over the raw logs and unpacked data for SetMonetaryPolicy events raised by the LlamalendController contract.
type LlamalendControllerSetMonetaryPolicyIterator struct {
	Event *LlamalendControllerSetMonetaryPolicy // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerSetMonetaryPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerSetMonetaryPolicy)
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
		it.Event = new(LlamalendControllerSetMonetaryPolicy)
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
func (it *LlamalendControllerSetMonetaryPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerSetMonetaryPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerSetMonetaryPolicy represents a SetMonetaryPolicy event raised by the LlamalendController contract.
type LlamalendControllerSetMonetaryPolicy struct {
	MonetaryPolicy common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetMonetaryPolicy is a free log retrieval operation binding the contract event 0x51fabb88f7860c9dbcc2a5a9b69a8b9476d63b87124591f97254e29f0e8daaeb.
//
// Solidity: event SetMonetaryPolicy(address monetary_policy)
func (_LlamalendController *LlamalendControllerFilterer) FilterSetMonetaryPolicy(opts *bind.FilterOpts) (*LlamalendControllerSetMonetaryPolicyIterator, error) {

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "SetMonetaryPolicy")
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerSetMonetaryPolicyIterator{contract: _LlamalendController.contract, event: "SetMonetaryPolicy", logs: logs, sub: sub}, nil
}

// WatchSetMonetaryPolicy is a free log subscription operation binding the contract event 0x51fabb88f7860c9dbcc2a5a9b69a8b9476d63b87124591f97254e29f0e8daaeb.
//
// Solidity: event SetMonetaryPolicy(address monetary_policy)
func (_LlamalendController *LlamalendControllerFilterer) WatchSetMonetaryPolicy(opts *bind.WatchOpts, sink chan<- *LlamalendControllerSetMonetaryPolicy) (event.Subscription, error) {

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "SetMonetaryPolicy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerSetMonetaryPolicy)
				if err := _LlamalendController.contract.UnpackLog(event, "SetMonetaryPolicy", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseSetMonetaryPolicy(log types.Log) (*LlamalendControllerSetMonetaryPolicy, error) {
	event := new(LlamalendControllerSetMonetaryPolicy)
	if err := _LlamalendController.contract.UnpackLog(event, "SetMonetaryPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LlamalendControllerUserStateIterator is returned from FilterUserState and is used to iterate over the raw logs and unpacked data for UserState events raised by the LlamalendController contract.
type LlamalendControllerUserStateIterator struct {
	Event *LlamalendControllerUserState // Event containing the contract specifics and raw log

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
func (it *LlamalendControllerUserStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LlamalendControllerUserState)
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
		it.Event = new(LlamalendControllerUserState)
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
func (it *LlamalendControllerUserStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LlamalendControllerUserStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LlamalendControllerUserState represents a UserState event raised by the LlamalendController contract.
type LlamalendControllerUserState struct {
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
func (_LlamalendController *LlamalendControllerFilterer) FilterUserState(opts *bind.FilterOpts, user []common.Address) (*LlamalendControllerUserStateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.FilterLogs(opts, "UserState", userRule)
	if err != nil {
		return nil, err
	}
	return &LlamalendControllerUserStateIterator{contract: _LlamalendController.contract, event: "UserState", logs: logs, sub: sub}, nil
}

// WatchUserState is a free log subscription operation binding the contract event 0xeec6b7095a637e006c79c1819d696e353a8f703db2c49fc0219e17a8fd04f7f2.
//
// Solidity: event UserState(address indexed user, uint256 collateral, uint256 debt, int256 n1, int256 n2, uint256 liquidation_discount)
func (_LlamalendController *LlamalendControllerFilterer) WatchUserState(opts *bind.WatchOpts, sink chan<- *LlamalendControllerUserState, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LlamalendController.contract.WatchLogs(opts, "UserState", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LlamalendControllerUserState)
				if err := _LlamalendController.contract.UnpackLog(event, "UserState", log); err != nil {
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
func (_LlamalendController *LlamalendControllerFilterer) ParseUserState(log types.Log) (*LlamalendControllerUserState, error) {
	event := new(LlamalendControllerUserState)
	if err := _LlamalendController.contract.UnpackLog(event, "UserState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
