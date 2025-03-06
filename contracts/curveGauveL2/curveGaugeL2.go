// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveGaugeL2

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
	Rate       *big.Int
	FinishTime *big.Int
}

// CurveGaugeL2MetaData contains all meta data concerning the CurveGaugeL2 contract.
var CurveGaugeL2MetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_crv_token\",\"type\":\"address\"},{\"name\":\"_gauge_controller\",\"type\":\"address\"},{\"name\":\"_minter\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transmit_emissions\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"integrate_fraction\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"user_checkpoint\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_killed\",\"inputs\":[{\"name\":\"_is_killed\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_bridger\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_child_gauge\",\"inputs\":[{\"name\":\"_child\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_bridger\",\"type\":\"address\"},{\"name\":\"_chain_id\",\"type\":\"uint256\"},{\"name\":\"_child\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"chain_id\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bridger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"child_gauge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"inflation_params\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"finish_time\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_period\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"total_emissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]}]",
}

// CurveGaugeL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGaugeL2MetaData.ABI instead.
var CurveGaugeL2ABI = CurveGaugeL2MetaData.ABI

// CurveGaugeL2 is an auto generated Go binding around an Ethereum contract.
type CurveGaugeL2 struct {
	CurveGaugeL2Caller     // Read-only binding to the contract
	CurveGaugeL2Transactor // Write-only binding to the contract
	CurveGaugeL2Filterer   // Log filterer for contract events
}

// CurveGaugeL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGaugeL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGaugeL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGaugeL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGaugeL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGaugeL2Session struct {
	Contract     *CurveGaugeL2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveGaugeL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGaugeL2CallerSession struct {
	Contract *CurveGaugeL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CurveGaugeL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGaugeL2TransactorSession struct {
	Contract     *CurveGaugeL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CurveGaugeL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGaugeL2Raw struct {
	Contract *CurveGaugeL2 // Generic contract binding to access the raw methods on
}

// CurveGaugeL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGaugeL2CallerRaw struct {
	Contract *CurveGaugeL2Caller // Generic read-only contract binding to access the raw methods on
}

// CurveGaugeL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGaugeL2TransactorRaw struct {
	Contract *CurveGaugeL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGaugeL2 creates a new instance of CurveGaugeL2, bound to a specific deployed contract.
func NewCurveGaugeL2(address common.Address, backend bind.ContractBackend) (*CurveGaugeL2, error) {
	contract, err := bindCurveGaugeL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeL2{CurveGaugeL2Caller: CurveGaugeL2Caller{contract: contract}, CurveGaugeL2Transactor: CurveGaugeL2Transactor{contract: contract}, CurveGaugeL2Filterer: CurveGaugeL2Filterer{contract: contract}}, nil
}

// NewCurveGaugeL2Caller creates a new read-only instance of CurveGaugeL2, bound to a specific deployed contract.
func NewCurveGaugeL2Caller(address common.Address, caller bind.ContractCaller) (*CurveGaugeL2Caller, error) {
	contract, err := bindCurveGaugeL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeL2Caller{contract: contract}, nil
}

// NewCurveGaugeL2Transactor creates a new write-only instance of CurveGaugeL2, bound to a specific deployed contract.
func NewCurveGaugeL2Transactor(address common.Address, transactor bind.ContractTransactor) (*CurveGaugeL2Transactor, error) {
	contract, err := bindCurveGaugeL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeL2Transactor{contract: contract}, nil
}

// NewCurveGaugeL2Filterer creates a new log filterer instance of CurveGaugeL2, bound to a specific deployed contract.
func NewCurveGaugeL2Filterer(address common.Address, filterer bind.ContractFilterer) (*CurveGaugeL2Filterer, error) {
	contract, err := bindCurveGaugeL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGaugeL2Filterer{contract: contract}, nil
}

// bindCurveGaugeL2 binds a generic wrapper to an already deployed contract.
func bindCurveGaugeL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveGaugeL2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGaugeL2 *CurveGaugeL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGaugeL2.Contract.CurveGaugeL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGaugeL2 *CurveGaugeL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.CurveGaugeL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGaugeL2 *CurveGaugeL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.CurveGaugeL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGaugeL2 *CurveGaugeL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGaugeL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGaugeL2 *CurveGaugeL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGaugeL2 *CurveGaugeL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.contract.Transact(opts, method, params...)
}

// Bridger is a free data retrieval call binding the contract method 0x4d47fc85.
//
// Solidity: function bridger() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Caller) Bridger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "bridger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridger is a free data retrieval call binding the contract method 0x4d47fc85.
//
// Solidity: function bridger() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Session) Bridger() (common.Address, error) {
	return _CurveGaugeL2.Contract.Bridger(&_CurveGaugeL2.CallOpts)
}

// Bridger is a free data retrieval call binding the contract method 0x4d47fc85.
//
// Solidity: function bridger() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) Bridger() (common.Address, error) {
	return _CurveGaugeL2.Contract.Bridger(&_CurveGaugeL2.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Caller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "chain_id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Session) ChainId() (*big.Int, error) {
	return _CurveGaugeL2.Contract.ChainId(&_CurveGaugeL2.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x3af973b1.
//
// Solidity: function chain_id() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) ChainId() (*big.Int, error) {
	return _CurveGaugeL2.Contract.ChainId(&_CurveGaugeL2.CallOpts)
}

// ChildGauge is a free data retrieval call binding the contract method 0x9dbb1517.
//
// Solidity: function child_gauge() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Caller) ChildGauge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "child_gauge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildGauge is a free data retrieval call binding the contract method 0x9dbb1517.
//
// Solidity: function child_gauge() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Session) ChildGauge() (common.Address, error) {
	return _CurveGaugeL2.Contract.ChildGauge(&_CurveGaugeL2.CallOpts)
}

// ChildGauge is a free data retrieval call binding the contract method 0x9dbb1517.
//
// Solidity: function child_gauge() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) ChildGauge() (common.Address, error) {
	return _CurveGaugeL2.Contract.ChildGauge(&_CurveGaugeL2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2Session) Factory() (common.Address, error) {
	return _CurveGaugeL2.Contract.Factory(&_CurveGaugeL2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) Factory() (common.Address, error) {
	return _CurveGaugeL2.Contract.Factory(&_CurveGaugeL2.CallOpts)
}

// InflationParams is a free data retrieval call binding the contract method 0x78de221b.
//
// Solidity: function inflation_params() view returns((uint256,uint256))
func (_CurveGaugeL2 *CurveGaugeL2Caller) InflationParams(opts *bind.CallOpts) (Struct0, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "inflation_params")

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// InflationParams is a free data retrieval call binding the contract method 0x78de221b.
//
// Solidity: function inflation_params() view returns((uint256,uint256))
func (_CurveGaugeL2 *CurveGaugeL2Session) InflationParams() (Struct0, error) {
	return _CurveGaugeL2.Contract.InflationParams(&_CurveGaugeL2.CallOpts)
}

// InflationParams is a free data retrieval call binding the contract method 0x78de221b.
//
// Solidity: function inflation_params() view returns((uint256,uint256))
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) InflationParams() (Struct0, error) {
	return _CurveGaugeL2.Contract.InflationParams(&_CurveGaugeL2.CallOpts)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address _user) view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Caller) IntegrateFraction(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "integrate_fraction", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address _user) view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Session) IntegrateFraction(_user common.Address) (*big.Int, error) {
	return _CurveGaugeL2.Contract.IntegrateFraction(&_CurveGaugeL2.CallOpts, _user)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address _user) view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) IntegrateFraction(_user common.Address) (*big.Int, error) {
	return _CurveGaugeL2.Contract.IntegrateFraction(&_CurveGaugeL2.CallOpts, _user)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2Caller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2Session) IsKilled() (bool, error) {
	return _CurveGaugeL2.Contract.IsKilled(&_CurveGaugeL2.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) IsKilled() (bool, error) {
	return _CurveGaugeL2.Contract.IsKilled(&_CurveGaugeL2.CallOpts)
}

// LastPeriod is a free data retrieval call binding the contract method 0x6f471ed1.
//
// Solidity: function last_period() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Caller) LastPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "last_period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPeriod is a free data retrieval call binding the contract method 0x6f471ed1.
//
// Solidity: function last_period() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Session) LastPeriod() (*big.Int, error) {
	return _CurveGaugeL2.Contract.LastPeriod(&_CurveGaugeL2.CallOpts)
}

// LastPeriod is a free data retrieval call binding the contract method 0x6f471ed1.
//
// Solidity: function last_period() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) LastPeriod() (*big.Int, error) {
	return _CurveGaugeL2.Contract.LastPeriod(&_CurveGaugeL2.CallOpts)
}

// TotalEmissions is a free data retrieval call binding the contract method 0xdf664d97.
//
// Solidity: function total_emissions() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Caller) TotalEmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "total_emissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEmissions is a free data retrieval call binding the contract method 0xdf664d97.
//
// Solidity: function total_emissions() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2Session) TotalEmissions() (*big.Int, error) {
	return _CurveGaugeL2.Contract.TotalEmissions(&_CurveGaugeL2.CallOpts)
}

// TotalEmissions is a free data retrieval call binding the contract method 0xdf664d97.
//
// Solidity: function total_emissions() view returns(uint256)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) TotalEmissions() (*big.Int, error) {
	return _CurveGaugeL2.Contract.TotalEmissions(&_CurveGaugeL2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveGaugeL2 *CurveGaugeL2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveGaugeL2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveGaugeL2 *CurveGaugeL2Session) Version() (string, error) {
	return _CurveGaugeL2.Contract.Version(&_CurveGaugeL2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveGaugeL2 *CurveGaugeL2CallerSession) Version() (string, error) {
	return _CurveGaugeL2.Contract.Version(&_CurveGaugeL2.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _bridger, uint256 _chain_id, address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) Initialize(opts *bind.TransactOpts, _bridger common.Address, _chain_id *big.Int, _child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "initialize", _bridger, _chain_id, _child)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _bridger, uint256 _chain_id, address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) Initialize(_bridger common.Address, _chain_id *big.Int, _child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.Initialize(&_CurveGaugeL2.TransactOpts, _bridger, _chain_id, _child)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address _bridger, uint256 _chain_id, address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) Initialize(_bridger common.Address, _chain_id *big.Int, _child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.Initialize(&_CurveGaugeL2.TransactOpts, _bridger, _chain_id, _child)
}

// SetChildGauge is a paid mutator transaction binding the contract method 0x729c415d.
//
// Solidity: function set_child_gauge(address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) SetChildGauge(opts *bind.TransactOpts, _child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "set_child_gauge", _child)
}

// SetChildGauge is a paid mutator transaction binding the contract method 0x729c415d.
//
// Solidity: function set_child_gauge(address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) SetChildGauge(_child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.SetChildGauge(&_CurveGaugeL2.TransactOpts, _child)
}

// SetChildGauge is a paid mutator transaction binding the contract method 0x729c415d.
//
// Solidity: function set_child_gauge(address _child) returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) SetChildGauge(_child common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.SetChildGauge(&_CurveGaugeL2.TransactOpts, _child)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) SetKilled(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "set_killed", _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.SetKilled(&_CurveGaugeL2.TransactOpts, _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.SetKilled(&_CurveGaugeL2.TransactOpts, _is_killed)
}

// TransmitEmissions is a paid mutator transaction binding the contract method 0x0c33d05a.
//
// Solidity: function transmit_emissions() returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) TransmitEmissions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "transmit_emissions")
}

// TransmitEmissions is a paid mutator transaction binding the contract method 0x0c33d05a.
//
// Solidity: function transmit_emissions() returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) TransmitEmissions() (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.TransmitEmissions(&_CurveGaugeL2.TransactOpts)
}

// TransmitEmissions is a paid mutator transaction binding the contract method 0x0c33d05a.
//
// Solidity: function transmit_emissions() returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) TransmitEmissions() (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.TransmitEmissions(&_CurveGaugeL2.TransactOpts)
}

// UpdateBridger is a paid mutator transaction binding the contract method 0x3ccc3da7.
//
// Solidity: function update_bridger() returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) UpdateBridger(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "update_bridger")
}

// UpdateBridger is a paid mutator transaction binding the contract method 0x3ccc3da7.
//
// Solidity: function update_bridger() returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) UpdateBridger() (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.UpdateBridger(&_CurveGaugeL2.TransactOpts)
}

// UpdateBridger is a paid mutator transaction binding the contract method 0x3ccc3da7.
//
// Solidity: function update_bridger() returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) UpdateBridger() (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.UpdateBridger(&_CurveGaugeL2.TransactOpts)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address _user) returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2Transactor) UserCheckpoint(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.Transact(opts, "user_checkpoint", _user)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address _user) returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2Session) UserCheckpoint(_user common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.UserCheckpoint(&_CurveGaugeL2.TransactOpts, _user)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address _user) returns(bool)
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) UserCheckpoint(_user common.Address) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.UserCheckpoint(&_CurveGaugeL2.TransactOpts, _user)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveGaugeL2 *CurveGaugeL2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CurveGaugeL2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveGaugeL2 *CurveGaugeL2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.Fallback(&_CurveGaugeL2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveGaugeL2 *CurveGaugeL2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurveGaugeL2.Contract.Fallback(&_CurveGaugeL2.TransactOpts, calldata)
}
