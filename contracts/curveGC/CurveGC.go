// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveGC

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

// CurveGCMetaData contains all meta data concerning the CurveGC contract.
var CurveGCMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"}]",
}

// CurveGCABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveGCMetaData.ABI instead.
var CurveGCABI = CurveGCMetaData.ABI

// CurveGC is an auto generated Go binding around an Ethereum contract.
type CurveGC struct {
	CurveGCCaller     // Read-only binding to the contract
	CurveGCTransactor // Write-only binding to the contract
	CurveGCFilterer   // Log filterer for contract events
}

// CurveGCCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveGCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveGCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveGCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveGCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveGCSession struct {
	Contract     *CurveGC          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveGCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveGCCallerSession struct {
	Contract *CurveGCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CurveGCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveGCTransactorSession struct {
	Contract     *CurveGCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CurveGCRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveGCRaw struct {
	Contract *CurveGC // Generic contract binding to access the raw methods on
}

// CurveGCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveGCCallerRaw struct {
	Contract *CurveGCCaller // Generic read-only contract binding to access the raw methods on
}

// CurveGCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveGCTransactorRaw struct {
	Contract *CurveGCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveGC creates a new instance of CurveGC, bound to a specific deployed contract.
func NewCurveGC(address common.Address, backend bind.ContractBackend) (*CurveGC, error) {
	contract, err := bindCurveGC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveGC{CurveGCCaller: CurveGCCaller{contract: contract}, CurveGCTransactor: CurveGCTransactor{contract: contract}, CurveGCFilterer: CurveGCFilterer{contract: contract}}, nil
}

// NewCurveGCCaller creates a new read-only instance of CurveGC, bound to a specific deployed contract.
func NewCurveGCCaller(address common.Address, caller bind.ContractCaller) (*CurveGCCaller, error) {
	contract, err := bindCurveGC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGCCaller{contract: contract}, nil
}

// NewCurveGCTransactor creates a new write-only instance of CurveGC, bound to a specific deployed contract.
func NewCurveGCTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveGCTransactor, error) {
	contract, err := bindCurveGC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveGCTransactor{contract: contract}, nil
}

// NewCurveGCFilterer creates a new log filterer instance of CurveGC, bound to a specific deployed contract.
func NewCurveGCFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveGCFilterer, error) {
	contract, err := bindCurveGC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveGCFilterer{contract: contract}, nil
}

// bindCurveGC binds a generic wrapper to an already deployed contract.
func bindCurveGC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveGCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGC *CurveGCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGC.Contract.CurveGCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGC *CurveGCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGC.Contract.CurveGCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGC *CurveGCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGC.Contract.CurveGCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveGC *CurveGCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveGC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveGC *CurveGCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveGC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveGC *CurveGCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveGC.Contract.contract.Transact(opts, method, params...)
}

// CurveGCDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CurveGC contract.
type CurveGCDepositIterator struct {
	Event *CurveGCDeposit // Event containing the contract specifics and raw log

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
func (it *CurveGCDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveGCDeposit)
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
		it.Event = new(CurveGCDeposit)
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
func (it *CurveGCDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveGCDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveGCDeposit represents a Deposit event raised by the CurveGC contract.
type CurveGCDeposit struct {
	Provider common.Address
	Value    *big.Int
	Locktime *big.Int
	Arg3     *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveGC *CurveGCFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*CurveGCDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _CurveGC.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &CurveGCDepositIterator{contract: _CurveGC.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveGC *CurveGCFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CurveGCDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _CurveGC.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveGCDeposit)
				if err := _CurveGC.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveGC *CurveGCFilterer) ParseDeposit(log types.Log) (*CurveGCDeposit, error) {
	event := new(CurveGCDeposit)
	if err := _CurveGC.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
