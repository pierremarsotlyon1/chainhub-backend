// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aragon

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
	_ = abi.ConvertType
)

// AragonMetaData contains all meta data concerning the Aragon contract.
var AragonMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"ADD_PROTECTED_TOKEN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC1271_INTERFACE_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC1271_RETURN_INVALID_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProtectedTokensLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RUN_SCRIPT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SAFE_EXECUTE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REMOVE_PROTECTED_TOKEN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDepositable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"presignHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DESIGNATE_SIGNER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeProtectedToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXECUTE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"addProtectedToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"protectedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERC1271_RETURN_VALID_SIGNATURE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_designatedSigner\",\"type\":\"address\"}],\"name\":\"setDesignatedSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"designatedSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeExecute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROTECTED_TOKENS_CAP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADD_PRESIGNED_HASH_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isPresigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_ethValue\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_evmScript\",\"type\":\"bytes\"}],\"name\":\"canForward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evmScript\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForwarder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SafeExecute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ethValue\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Execute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AddProtectedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RemoveProtectedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"PresignHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"SetDesignatedSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ReceiveERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"VaultTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"VaultDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"}]",
}

// AragonABI is the input ABI used to generate the binding from.
// Deprecated: Use AragonMetaData.ABI instead.
var AragonABI = AragonMetaData.ABI

// Aragon is an auto generated Go binding around an Ethereum contract.
type Aragon struct {
	AragonCaller     // Read-only binding to the contract
	AragonTransactor // Write-only binding to the contract
	AragonFilterer   // Log filterer for contract events
}

// AragonCaller is an auto generated read-only Go binding around an Ethereum contract.
type AragonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AragonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AragonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AragonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AragonSession struct {
	Contract     *Aragon           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AragonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AragonCallerSession struct {
	Contract *AragonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AragonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AragonTransactorSession struct {
	Contract     *AragonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AragonRaw is an auto generated low-level Go binding around an Ethereum contract.
type AragonRaw struct {
	Contract *Aragon // Generic contract binding to access the raw methods on
}

// AragonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AragonCallerRaw struct {
	Contract *AragonCaller // Generic read-only contract binding to access the raw methods on
}

// AragonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AragonTransactorRaw struct {
	Contract *AragonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAragon creates a new instance of Aragon, bound to a specific deployed contract.
func NewAragon(address common.Address, backend bind.ContractBackend) (*Aragon, error) {
	contract, err := bindAragon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aragon{AragonCaller: AragonCaller{contract: contract}, AragonTransactor: AragonTransactor{contract: contract}, AragonFilterer: AragonFilterer{contract: contract}}, nil
}

// NewAragonCaller creates a new read-only instance of Aragon, bound to a specific deployed contract.
func NewAragonCaller(address common.Address, caller bind.ContractCaller) (*AragonCaller, error) {
	contract, err := bindAragon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AragonCaller{contract: contract}, nil
}

// NewAragonTransactor creates a new write-only instance of Aragon, bound to a specific deployed contract.
func NewAragonTransactor(address common.Address, transactor bind.ContractTransactor) (*AragonTransactor, error) {
	contract, err := bindAragon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AragonTransactor{contract: contract}, nil
}

// NewAragonFilterer creates a new log filterer instance of Aragon, bound to a specific deployed contract.
func NewAragonFilterer(address common.Address, filterer bind.ContractFilterer) (*AragonFilterer, error) {
	contract, err := bindAragon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AragonFilterer{contract: contract}, nil
}

// bindAragon binds a generic wrapper to an already deployed contract.
func bindAragon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AragonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aragon *AragonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aragon.Contract.AragonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aragon *AragonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aragon.Contract.AragonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aragon *AragonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aragon.Contract.AragonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aragon *AragonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aragon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aragon *AragonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aragon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aragon *AragonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aragon.Contract.contract.Transact(opts, method, params...)
}

// ADDPRESIGNEDHASHROLE is a free data retrieval call binding the contract method 0xb06c4244.
//
// Solidity: function ADD_PRESIGNED_HASH_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) ADDPRESIGNEDHASHROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "ADD_PRESIGNED_HASH_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDPRESIGNEDHASHROLE is a free data retrieval call binding the contract method 0xb06c4244.
//
// Solidity: function ADD_PRESIGNED_HASH_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) ADDPRESIGNEDHASHROLE() ([32]byte, error) {
	return _Aragon.Contract.ADDPRESIGNEDHASHROLE(&_Aragon.CallOpts)
}

// ADDPRESIGNEDHASHROLE is a free data retrieval call binding the contract method 0xb06c4244.
//
// Solidity: function ADD_PRESIGNED_HASH_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) ADDPRESIGNEDHASHROLE() ([32]byte, error) {
	return _Aragon.Contract.ADDPRESIGNEDHASHROLE(&_Aragon.CallOpts)
}

// ADDPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x007bb003.
//
// Solidity: function ADD_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) ADDPROTECTEDTOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "ADD_PROTECTED_TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x007bb003.
//
// Solidity: function ADD_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) ADDPROTECTEDTOKENROLE() ([32]byte, error) {
	return _Aragon.Contract.ADDPROTECTEDTOKENROLE(&_Aragon.CallOpts)
}

// ADDPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x007bb003.
//
// Solidity: function ADD_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) ADDPROTECTEDTOKENROLE() ([32]byte, error) {
	return _Aragon.Contract.ADDPROTECTEDTOKENROLE(&_Aragon.CallOpts)
}

// DESIGNATESIGNERROLE is a free data retrieval call binding the contract method 0x54842f14.
//
// Solidity: function DESIGNATE_SIGNER_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) DESIGNATESIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "DESIGNATE_SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DESIGNATESIGNERROLE is a free data retrieval call binding the contract method 0x54842f14.
//
// Solidity: function DESIGNATE_SIGNER_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) DESIGNATESIGNERROLE() ([32]byte, error) {
	return _Aragon.Contract.DESIGNATESIGNERROLE(&_Aragon.CallOpts)
}

// DESIGNATESIGNERROLE is a free data retrieval call binding the contract method 0x54842f14.
//
// Solidity: function DESIGNATE_SIGNER_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) DESIGNATESIGNERROLE() ([32]byte, error) {
	return _Aragon.Contract.DESIGNATESIGNERROLE(&_Aragon.CallOpts)
}

// ERC1271INTERFACEID is a free data retrieval call binding the contract method 0x11a5e409.
//
// Solidity: function ERC1271_INTERFACE_ID() view returns(bytes4)
func (_Aragon *AragonCaller) ERC1271INTERFACEID(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "ERC1271_INTERFACE_ID")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ERC1271INTERFACEID is a free data retrieval call binding the contract method 0x11a5e409.
//
// Solidity: function ERC1271_INTERFACE_ID() view returns(bytes4)
func (_Aragon *AragonSession) ERC1271INTERFACEID() ([4]byte, error) {
	return _Aragon.Contract.ERC1271INTERFACEID(&_Aragon.CallOpts)
}

// ERC1271INTERFACEID is a free data retrieval call binding the contract method 0x11a5e409.
//
// Solidity: function ERC1271_INTERFACE_ID() view returns(bytes4)
func (_Aragon *AragonCallerSession) ERC1271INTERFACEID() ([4]byte, error) {
	return _Aragon.Contract.ERC1271INTERFACEID(&_Aragon.CallOpts)
}

// ERC1271RETURNINVALIDSIGNATURE is a free data retrieval call binding the contract method 0x1ce30181.
//
// Solidity: function ERC1271_RETURN_INVALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonCaller) ERC1271RETURNINVALIDSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "ERC1271_RETURN_INVALID_SIGNATURE")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ERC1271RETURNINVALIDSIGNATURE is a free data retrieval call binding the contract method 0x1ce30181.
//
// Solidity: function ERC1271_RETURN_INVALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonSession) ERC1271RETURNINVALIDSIGNATURE() ([4]byte, error) {
	return _Aragon.Contract.ERC1271RETURNINVALIDSIGNATURE(&_Aragon.CallOpts)
}

// ERC1271RETURNINVALIDSIGNATURE is a free data retrieval call binding the contract method 0x1ce30181.
//
// Solidity: function ERC1271_RETURN_INVALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonCallerSession) ERC1271RETURNINVALIDSIGNATURE() ([4]byte, error) {
	return _Aragon.Contract.ERC1271RETURNINVALIDSIGNATURE(&_Aragon.CallOpts)
}

// ERC1271RETURNVALIDSIGNATURE is a free data retrieval call binding the contract method 0x9890cdca.
//
// Solidity: function ERC1271_RETURN_VALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonCaller) ERC1271RETURNVALIDSIGNATURE(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "ERC1271_RETURN_VALID_SIGNATURE")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ERC1271RETURNVALIDSIGNATURE is a free data retrieval call binding the contract method 0x9890cdca.
//
// Solidity: function ERC1271_RETURN_VALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonSession) ERC1271RETURNVALIDSIGNATURE() ([4]byte, error) {
	return _Aragon.Contract.ERC1271RETURNVALIDSIGNATURE(&_Aragon.CallOpts)
}

// ERC1271RETURNVALIDSIGNATURE is a free data retrieval call binding the contract method 0x9890cdca.
//
// Solidity: function ERC1271_RETURN_VALID_SIGNATURE() view returns(bytes4)
func (_Aragon *AragonCallerSession) ERC1271RETURNVALIDSIGNATURE() ([4]byte, error) {
	return _Aragon.Contract.ERC1271RETURNVALIDSIGNATURE(&_Aragon.CallOpts)
}

// EXECUTEROLE is a free data retrieval call binding the contract method 0x5fa5e4e6.
//
// Solidity: function EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) EXECUTEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "EXECUTE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTEROLE is a free data retrieval call binding the contract method 0x5fa5e4e6.
//
// Solidity: function EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) EXECUTEROLE() ([32]byte, error) {
	return _Aragon.Contract.EXECUTEROLE(&_Aragon.CallOpts)
}

// EXECUTEROLE is a free data retrieval call binding the contract method 0x5fa5e4e6.
//
// Solidity: function EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) EXECUTEROLE() ([32]byte, error) {
	return _Aragon.Contract.EXECUTEROLE(&_Aragon.CallOpts)
}

// PROTECTEDTOKENSCAP is a free data retrieval call binding the contract method 0xb03bdb04.
//
// Solidity: function PROTECTED_TOKENS_CAP() view returns(uint256)
func (_Aragon *AragonCaller) PROTECTEDTOKENSCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "PROTECTED_TOKENS_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTECTEDTOKENSCAP is a free data retrieval call binding the contract method 0xb03bdb04.
//
// Solidity: function PROTECTED_TOKENS_CAP() view returns(uint256)
func (_Aragon *AragonSession) PROTECTEDTOKENSCAP() (*big.Int, error) {
	return _Aragon.Contract.PROTECTEDTOKENSCAP(&_Aragon.CallOpts)
}

// PROTECTEDTOKENSCAP is a free data retrieval call binding the contract method 0xb03bdb04.
//
// Solidity: function PROTECTED_TOKENS_CAP() view returns(uint256)
func (_Aragon *AragonCallerSession) PROTECTEDTOKENSCAP() (*big.Int, error) {
	return _Aragon.Contract.PROTECTEDTOKENSCAP(&_Aragon.CallOpts)
}

// REMOVEPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x42b2d066.
//
// Solidity: function REMOVE_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) REMOVEPROTECTEDTOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "REMOVE_PROTECTED_TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REMOVEPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x42b2d066.
//
// Solidity: function REMOVE_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) REMOVEPROTECTEDTOKENROLE() ([32]byte, error) {
	return _Aragon.Contract.REMOVEPROTECTEDTOKENROLE(&_Aragon.CallOpts)
}

// REMOVEPROTECTEDTOKENROLE is a free data retrieval call binding the contract method 0x42b2d066.
//
// Solidity: function REMOVE_PROTECTED_TOKEN_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) REMOVEPROTECTEDTOKENROLE() ([32]byte, error) {
	return _Aragon.Contract.REMOVEPROTECTEDTOKENROLE(&_Aragon.CallOpts)
}

// RUNSCRIPTROLE is a free data retrieval call binding the contract method 0x368c3c34.
//
// Solidity: function RUN_SCRIPT_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) RUNSCRIPTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "RUN_SCRIPT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RUNSCRIPTROLE is a free data retrieval call binding the contract method 0x368c3c34.
//
// Solidity: function RUN_SCRIPT_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) RUNSCRIPTROLE() ([32]byte, error) {
	return _Aragon.Contract.RUNSCRIPTROLE(&_Aragon.CallOpts)
}

// RUNSCRIPTROLE is a free data retrieval call binding the contract method 0x368c3c34.
//
// Solidity: function RUN_SCRIPT_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) RUNSCRIPTROLE() ([32]byte, error) {
	return _Aragon.Contract.RUNSCRIPTROLE(&_Aragon.CallOpts)
}

// SAFEEXECUTEROLE is a free data retrieval call binding the contract method 0x3e4eb756.
//
// Solidity: function SAFE_EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) SAFEEXECUTEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "SAFE_EXECUTE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SAFEEXECUTEROLE is a free data retrieval call binding the contract method 0x3e4eb756.
//
// Solidity: function SAFE_EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) SAFEEXECUTEROLE() ([32]byte, error) {
	return _Aragon.Contract.SAFEEXECUTEROLE(&_Aragon.CallOpts)
}

// SAFEEXECUTEROLE is a free data retrieval call binding the contract method 0x3e4eb756.
//
// Solidity: function SAFE_EXECUTE_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) SAFEEXECUTEROLE() ([32]byte, error) {
	return _Aragon.Contract.SAFEEXECUTEROLE(&_Aragon.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Aragon *AragonCaller) TRANSFERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "TRANSFER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Aragon *AragonSession) TRANSFERROLE() ([32]byte, error) {
	return _Aragon.Contract.TRANSFERROLE(&_Aragon.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Aragon *AragonCallerSession) TRANSFERROLE() ([32]byte, error) {
	return _Aragon.Contract.TRANSFERROLE(&_Aragon.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address ) view returns(bool)
func (_Aragon *AragonCaller) AllowRecoverability(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "allowRecoverability", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address ) view returns(bool)
func (_Aragon *AragonSession) AllowRecoverability(arg0 common.Address) (bool, error) {
	return _Aragon.Contract.AllowRecoverability(&_Aragon.CallOpts, arg0)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address ) view returns(bool)
func (_Aragon *AragonCallerSession) AllowRecoverability(arg0 common.Address) (bool, error) {
	return _Aragon.Contract.AllowRecoverability(&_Aragon.CallOpts, arg0)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Aragon *AragonCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Aragon *AragonSession) AppId() ([32]byte, error) {
	return _Aragon.Contract.AppId(&_Aragon.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Aragon *AragonCallerSession) AppId() ([32]byte, error) {
	return _Aragon.Contract.AppId(&_Aragon.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _token) view returns(uint256)
func (_Aragon *AragonCaller) Balance(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "balance", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _token) view returns(uint256)
func (_Aragon *AragonSession) Balance(_token common.Address) (*big.Int, error) {
	return _Aragon.Contract.Balance(&_Aragon.CallOpts, _token)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address _token) view returns(uint256)
func (_Aragon *AragonCallerSession) Balance(_token common.Address) (*big.Int, error) {
	return _Aragon.Contract.Balance(&_Aragon.CallOpts, _token)
}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes _evmScript) view returns(bool)
func (_Aragon *AragonCaller) CanForward(opts *bind.CallOpts, _sender common.Address, _evmScript []byte) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "canForward", _sender, _evmScript)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes _evmScript) view returns(bool)
func (_Aragon *AragonSession) CanForward(_sender common.Address, _evmScript []byte) (bool, error) {
	return _Aragon.Contract.CanForward(&_Aragon.CallOpts, _sender, _evmScript)
}

// CanForward is a free data retrieval call binding the contract method 0xc0774df3.
//
// Solidity: function canForward(address _sender, bytes _evmScript) view returns(bool)
func (_Aragon *AragonCallerSession) CanForward(_sender common.Address, _evmScript []byte) (bool, error) {
	return _Aragon.Contract.CanForward(&_Aragon.CallOpts, _sender, _evmScript)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Aragon *AragonCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Aragon *AragonSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Aragon.Contract.CanPerform(&_Aragon.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Aragon *AragonCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Aragon.Contract.CanPerform(&_Aragon.CallOpts, _sender, _role, _params)
}

// DesignatedSigner is a free data retrieval call binding the contract method 0xaae25051.
//
// Solidity: function designatedSigner() view returns(address)
func (_Aragon *AragonCaller) DesignatedSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "designatedSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DesignatedSigner is a free data retrieval call binding the contract method 0xaae25051.
//
// Solidity: function designatedSigner() view returns(address)
func (_Aragon *AragonSession) DesignatedSigner() (common.Address, error) {
	return _Aragon.Contract.DesignatedSigner(&_Aragon.CallOpts)
}

// DesignatedSigner is a free data retrieval call binding the contract method 0xaae25051.
//
// Solidity: function designatedSigner() view returns(address)
func (_Aragon *AragonCallerSession) DesignatedSigner() (common.Address, error) {
	return _Aragon.Contract.DesignatedSigner(&_Aragon.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Aragon *AragonCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Aragon *AragonSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Aragon.Contract.GetEVMScriptExecutor(&_Aragon.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Aragon *AragonCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Aragon.Contract.GetEVMScriptExecutor(&_Aragon.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Aragon *AragonCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Aragon *AragonSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Aragon.Contract.GetEVMScriptRegistry(&_Aragon.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Aragon *AragonCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Aragon.Contract.GetEVMScriptRegistry(&_Aragon.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Aragon *AragonCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Aragon *AragonSession) GetInitializationBlock() (*big.Int, error) {
	return _Aragon.Contract.GetInitializationBlock(&_Aragon.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Aragon *AragonCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _Aragon.Contract.GetInitializationBlock(&_Aragon.CallOpts)
}

// GetProtectedTokensLength is a free data retrieval call binding the contract method 0x26f06d24.
//
// Solidity: function getProtectedTokensLength() view returns(uint256)
func (_Aragon *AragonCaller) GetProtectedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "getProtectedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtectedTokensLength is a free data retrieval call binding the contract method 0x26f06d24.
//
// Solidity: function getProtectedTokensLength() view returns(uint256)
func (_Aragon *AragonSession) GetProtectedTokensLength() (*big.Int, error) {
	return _Aragon.Contract.GetProtectedTokensLength(&_Aragon.CallOpts)
}

// GetProtectedTokensLength is a free data retrieval call binding the contract method 0x26f06d24.
//
// Solidity: function getProtectedTokensLength() view returns(uint256)
func (_Aragon *AragonCallerSession) GetProtectedTokensLength() (*big.Int, error) {
	return _Aragon.Contract.GetProtectedTokensLength(&_Aragon.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Aragon *AragonCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Aragon *AragonSession) GetRecoveryVault() (common.Address, error) {
	return _Aragon.Contract.GetRecoveryVault(&_Aragon.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Aragon *AragonCallerSession) GetRecoveryVault() (common.Address, error) {
	return _Aragon.Contract.GetRecoveryVault(&_Aragon.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Aragon *AragonCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Aragon *AragonSession) HasInitialized() (bool, error) {
	return _Aragon.Contract.HasInitialized(&_Aragon.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Aragon *AragonCallerSession) HasInitialized() (bool, error) {
	return _Aragon.Contract.HasInitialized(&_Aragon.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_Aragon *AragonCaller) IsDepositable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isDepositable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_Aragon *AragonSession) IsDepositable() (bool, error) {
	return _Aragon.Contract.IsDepositable(&_Aragon.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_Aragon *AragonCallerSession) IsDepositable() (bool, error) {
	return _Aragon.Contract.IsDepositable(&_Aragon.CallOpts)
}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Aragon *AragonCaller) IsForwarder(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isForwarder")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Aragon *AragonSession) IsForwarder() (bool, error) {
	return _Aragon.Contract.IsForwarder(&_Aragon.CallOpts)
}

// IsForwarder is a free data retrieval call binding the contract method 0xfd64eccb.
//
// Solidity: function isForwarder() pure returns(bool)
func (_Aragon *AragonCallerSession) IsForwarder() (bool, error) {
	return _Aragon.Contract.IsForwarder(&_Aragon.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Aragon *AragonCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Aragon *AragonSession) IsPetrified() (bool, error) {
	return _Aragon.Contract.IsPetrified(&_Aragon.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Aragon *AragonCallerSession) IsPetrified() (bool, error) {
	return _Aragon.Contract.IsPetrified(&_Aragon.CallOpts)
}

// IsPresigned is a free data retrieval call binding the contract method 0xb4fa653c.
//
// Solidity: function isPresigned(bytes32 ) view returns(bool)
func (_Aragon *AragonCaller) IsPresigned(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isPresigned", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPresigned is a free data retrieval call binding the contract method 0xb4fa653c.
//
// Solidity: function isPresigned(bytes32 ) view returns(bool)
func (_Aragon *AragonSession) IsPresigned(arg0 [32]byte) (bool, error) {
	return _Aragon.Contract.IsPresigned(&_Aragon.CallOpts, arg0)
}

// IsPresigned is a free data retrieval call binding the contract method 0xb4fa653c.
//
// Solidity: function isPresigned(bytes32 ) view returns(bool)
func (_Aragon *AragonCallerSession) IsPresigned(arg0 [32]byte) (bool, error) {
	return _Aragon.Contract.IsPresigned(&_Aragon.CallOpts, arg0)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Aragon *AragonCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Aragon *AragonSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Aragon.Contract.IsValidSignature(&_Aragon.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Aragon *AragonCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Aragon.Contract.IsValidSignature(&_Aragon.CallOpts, _hash, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Aragon *AragonCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "isValidSignature0", _data, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Aragon *AragonSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Aragon.Contract.IsValidSignature0(&_Aragon.CallOpts, _data, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Aragon *AragonCallerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Aragon.Contract.IsValidSignature0(&_Aragon.CallOpts, _data, _signature)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Aragon *AragonCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Aragon *AragonSession) Kernel() (common.Address, error) {
	return _Aragon.Contract.Kernel(&_Aragon.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Aragon *AragonCallerSession) Kernel() (common.Address, error) {
	return _Aragon.Contract.Kernel(&_Aragon.CallOpts)
}

// ProtectedTokens is a free data retrieval call binding the contract method 0x851a3790.
//
// Solidity: function protectedTokens(uint256 ) view returns(address)
func (_Aragon *AragonCaller) ProtectedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "protectedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtectedTokens is a free data retrieval call binding the contract method 0x851a3790.
//
// Solidity: function protectedTokens(uint256 ) view returns(address)
func (_Aragon *AragonSession) ProtectedTokens(arg0 *big.Int) (common.Address, error) {
	return _Aragon.Contract.ProtectedTokens(&_Aragon.CallOpts, arg0)
}

// ProtectedTokens is a free data retrieval call binding the contract method 0x851a3790.
//
// Solidity: function protectedTokens(uint256 ) view returns(address)
func (_Aragon *AragonCallerSession) ProtectedTokens(arg0 *big.Int) (common.Address, error) {
	return _Aragon.Contract.ProtectedTokens(&_Aragon.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_Aragon *AragonCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Aragon.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_Aragon *AragonSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Aragon.Contract.SupportsInterface(&_Aragon.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_Aragon *AragonCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Aragon.Contract.SupportsInterface(&_Aragon.CallOpts, _interfaceId)
}

// AddProtectedToken is a paid mutator transaction binding the contract method 0x6298e902.
//
// Solidity: function addProtectedToken(address _token) returns()
func (_Aragon *AragonTransactor) AddProtectedToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "addProtectedToken", _token)
}

// AddProtectedToken is a paid mutator transaction binding the contract method 0x6298e902.
//
// Solidity: function addProtectedToken(address _token) returns()
func (_Aragon *AragonSession) AddProtectedToken(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.AddProtectedToken(&_Aragon.TransactOpts, _token)
}

// AddProtectedToken is a paid mutator transaction binding the contract method 0x6298e902.
//
// Solidity: function addProtectedToken(address _token) returns()
func (_Aragon *AragonTransactorSession) AddProtectedToken(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.AddProtectedToken(&_Aragon.TransactOpts, _token)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _value) payable returns()
func (_Aragon *AragonTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "deposit", _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _value) payable returns()
func (_Aragon *AragonSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.Contract.Deposit(&_Aragon.TransactOpts, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _value) payable returns()
func (_Aragon *AragonTransactorSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.Contract.Deposit(&_Aragon.TransactOpts, _token, _value)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _target, uint256 _ethValue, bytes _data) returns()
func (_Aragon *AragonTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _ethValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "execute", _target, _ethValue, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _target, uint256 _ethValue, bytes _data) returns()
func (_Aragon *AragonSession) Execute(_target common.Address, _ethValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Execute(&_Aragon.TransactOpts, _target, _ethValue, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address _target, uint256 _ethValue, bytes _data) returns()
func (_Aragon *AragonTransactorSession) Execute(_target common.Address, _ethValue *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Execute(&_Aragon.TransactOpts, _target, _ethValue, _data)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Aragon *AragonTransactor) Forward(opts *bind.TransactOpts, _evmScript []byte) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "forward", _evmScript)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Aragon *AragonSession) Forward(_evmScript []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Forward(&_Aragon.TransactOpts, _evmScript)
}

// Forward is a paid mutator transaction binding the contract method 0xd948d468.
//
// Solidity: function forward(bytes _evmScript) returns()
func (_Aragon *AragonTransactorSession) Forward(_evmScript []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Forward(&_Aragon.TransactOpts, _evmScript)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Aragon *AragonTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Aragon *AragonSession) Initialize() (*types.Transaction, error) {
	return _Aragon.Contract.Initialize(&_Aragon.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Aragon *AragonTransactorSession) Initialize() (*types.Transaction, error) {
	return _Aragon.Contract.Initialize(&_Aragon.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Aragon *AragonTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Aragon *AragonSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.OnERC721Received(&_Aragon.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Aragon *AragonTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.OnERC721Received(&_Aragon.TransactOpts, _operator, _from, _tokenId, _data)
}

// PresignHash is a paid mutator transaction binding the contract method 0x4c7ec0b0.
//
// Solidity: function presignHash(bytes32 _hash) returns()
func (_Aragon *AragonTransactor) PresignHash(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "presignHash", _hash)
}

// PresignHash is a paid mutator transaction binding the contract method 0x4c7ec0b0.
//
// Solidity: function presignHash(bytes32 _hash) returns()
func (_Aragon *AragonSession) PresignHash(_hash [32]byte) (*types.Transaction, error) {
	return _Aragon.Contract.PresignHash(&_Aragon.TransactOpts, _hash)
}

// PresignHash is a paid mutator transaction binding the contract method 0x4c7ec0b0.
//
// Solidity: function presignHash(bytes32 _hash) returns()
func (_Aragon *AragonTransactorSession) PresignHash(_hash [32]byte) (*types.Transaction, error) {
	return _Aragon.Contract.PresignHash(&_Aragon.TransactOpts, _hash)
}

// RemoveProtectedToken is a paid mutator transaction binding the contract method 0x578eb50b.
//
// Solidity: function removeProtectedToken(address _token) returns()
func (_Aragon *AragonTransactor) RemoveProtectedToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "removeProtectedToken", _token)
}

// RemoveProtectedToken is a paid mutator transaction binding the contract method 0x578eb50b.
//
// Solidity: function removeProtectedToken(address _token) returns()
func (_Aragon *AragonSession) RemoveProtectedToken(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.RemoveProtectedToken(&_Aragon.TransactOpts, _token)
}

// RemoveProtectedToken is a paid mutator transaction binding the contract method 0x578eb50b.
//
// Solidity: function removeProtectedToken(address _token) returns()
func (_Aragon *AragonTransactorSession) RemoveProtectedToken(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.RemoveProtectedToken(&_Aragon.TransactOpts, _token)
}

// SafeExecute is a paid mutator transaction binding the contract method 0xab23c345.
//
// Solidity: function safeExecute(address _target, bytes _data) returns()
func (_Aragon *AragonTransactor) SafeExecute(opts *bind.TransactOpts, _target common.Address, _data []byte) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "safeExecute", _target, _data)
}

// SafeExecute is a paid mutator transaction binding the contract method 0xab23c345.
//
// Solidity: function safeExecute(address _target, bytes _data) returns()
func (_Aragon *AragonSession) SafeExecute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.SafeExecute(&_Aragon.TransactOpts, _target, _data)
}

// SafeExecute is a paid mutator transaction binding the contract method 0xab23c345.
//
// Solidity: function safeExecute(address _target, bytes _data) returns()
func (_Aragon *AragonTransactorSession) SafeExecute(_target common.Address, _data []byte) (*types.Transaction, error) {
	return _Aragon.Contract.SafeExecute(&_Aragon.TransactOpts, _target, _data)
}

// SetDesignatedSigner is a paid mutator transaction binding the contract method 0xa83e52b4.
//
// Solidity: function setDesignatedSigner(address _designatedSigner) returns()
func (_Aragon *AragonTransactor) SetDesignatedSigner(opts *bind.TransactOpts, _designatedSigner common.Address) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "setDesignatedSigner", _designatedSigner)
}

// SetDesignatedSigner is a paid mutator transaction binding the contract method 0xa83e52b4.
//
// Solidity: function setDesignatedSigner(address _designatedSigner) returns()
func (_Aragon *AragonSession) SetDesignatedSigner(_designatedSigner common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.SetDesignatedSigner(&_Aragon.TransactOpts, _designatedSigner)
}

// SetDesignatedSigner is a paid mutator transaction binding the contract method 0xa83e52b4.
//
// Solidity: function setDesignatedSigner(address _designatedSigner) returns()
func (_Aragon *AragonTransactorSession) SetDesignatedSigner(_designatedSigner common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.SetDesignatedSigner(&_Aragon.TransactOpts, _designatedSigner)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _to, uint256 _value) returns()
func (_Aragon *AragonTransactor) Transfer(opts *bind.TransactOpts, _token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "transfer", _token, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _to, uint256 _value) returns()
func (_Aragon *AragonSession) Transfer(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.Contract.Transfer(&_Aragon.TransactOpts, _token, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _token, address _to, uint256 _value) returns()
func (_Aragon *AragonTransactorSession) Transfer(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Aragon.Contract.Transfer(&_Aragon.TransactOpts, _token, _to, _value)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Aragon *AragonTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Aragon.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Aragon *AragonSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.TransferToVault(&_Aragon.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Aragon *AragonTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Aragon.Contract.TransferToVault(&_Aragon.TransactOpts, _token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aragon *AragonTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Aragon.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aragon *AragonSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Fallback(&_Aragon.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aragon *AragonTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Aragon.Contract.Fallback(&_Aragon.TransactOpts, calldata)
}

// AragonAddProtectedTokenIterator is returned from FilterAddProtectedToken and is used to iterate over the raw logs and unpacked data for AddProtectedToken events raised by the Aragon contract.
type AragonAddProtectedTokenIterator struct {
	Event *AragonAddProtectedToken // Event containing the contract specifics and raw log

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
func (it *AragonAddProtectedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonAddProtectedToken)
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
		it.Event = new(AragonAddProtectedToken)
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
func (it *AragonAddProtectedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonAddProtectedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonAddProtectedToken represents a AddProtectedToken event raised by the Aragon contract.
type AragonAddProtectedToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProtectedToken is a free log retrieval operation binding the contract event 0xf70a5123a7f334e5dac1d9aa3a6aafbc316712bf2519ffe0d3aa4f7cba52767e.
//
// Solidity: event AddProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) FilterAddProtectedToken(opts *bind.FilterOpts, token []common.Address) (*AragonAddProtectedTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "AddProtectedToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &AragonAddProtectedTokenIterator{contract: _Aragon.contract, event: "AddProtectedToken", logs: logs, sub: sub}, nil
}

// WatchAddProtectedToken is a free log subscription operation binding the contract event 0xf70a5123a7f334e5dac1d9aa3a6aafbc316712bf2519ffe0d3aa4f7cba52767e.
//
// Solidity: event AddProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) WatchAddProtectedToken(opts *bind.WatchOpts, sink chan<- *AragonAddProtectedToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "AddProtectedToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonAddProtectedToken)
				if err := _Aragon.contract.UnpackLog(event, "AddProtectedToken", log); err != nil {
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

// ParseAddProtectedToken is a log parse operation binding the contract event 0xf70a5123a7f334e5dac1d9aa3a6aafbc316712bf2519ffe0d3aa4f7cba52767e.
//
// Solidity: event AddProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) ParseAddProtectedToken(log types.Log) (*AragonAddProtectedToken, error) {
	event := new(AragonAddProtectedToken)
	if err := _Aragon.contract.UnpackLog(event, "AddProtectedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonExecuteIterator is returned from FilterExecute and is used to iterate over the raw logs and unpacked data for Execute events raised by the Aragon contract.
type AragonExecuteIterator struct {
	Event *AragonExecute // Event containing the contract specifics and raw log

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
func (it *AragonExecuteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonExecute)
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
		it.Event = new(AragonExecute)
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
func (it *AragonExecuteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonExecuteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonExecute represents a Execute event raised by the Aragon contract.
type AragonExecute struct {
	Sender   common.Address
	Target   common.Address
	EthValue *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecute is a free log retrieval operation binding the contract event 0xc25cfed0b22da6a56f0e5ff784979a0b8623eddf2aee4acd33c2adefb09cbab6.
//
// Solidity: event Execute(address indexed sender, address indexed target, uint256 ethValue, bytes data)
func (_Aragon *AragonFilterer) FilterExecute(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*AragonExecuteIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "Execute", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AragonExecuteIterator{contract: _Aragon.contract, event: "Execute", logs: logs, sub: sub}, nil
}

// WatchExecute is a free log subscription operation binding the contract event 0xc25cfed0b22da6a56f0e5ff784979a0b8623eddf2aee4acd33c2adefb09cbab6.
//
// Solidity: event Execute(address indexed sender, address indexed target, uint256 ethValue, bytes data)
func (_Aragon *AragonFilterer) WatchExecute(opts *bind.WatchOpts, sink chan<- *AragonExecute, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "Execute", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonExecute)
				if err := _Aragon.contract.UnpackLog(event, "Execute", log); err != nil {
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

// ParseExecute is a log parse operation binding the contract event 0xc25cfed0b22da6a56f0e5ff784979a0b8623eddf2aee4acd33c2adefb09cbab6.
//
// Solidity: event Execute(address indexed sender, address indexed target, uint256 ethValue, bytes data)
func (_Aragon *AragonFilterer) ParseExecute(log types.Log) (*AragonExecute, error) {
	event := new(AragonExecute)
	if err := _Aragon.contract.UnpackLog(event, "Execute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonPresignHashIterator is returned from FilterPresignHash and is used to iterate over the raw logs and unpacked data for PresignHash events raised by the Aragon contract.
type AragonPresignHashIterator struct {
	Event *AragonPresignHash // Event containing the contract specifics and raw log

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
func (it *AragonPresignHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonPresignHash)
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
		it.Event = new(AragonPresignHash)
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
func (it *AragonPresignHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonPresignHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonPresignHash represents a PresignHash event raised by the Aragon contract.
type AragonPresignHash struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPresignHash is a free log retrieval operation binding the contract event 0xb150212e573caa2a0216cf8f273ef996af143fd7b2d35abd92c8105536cac160.
//
// Solidity: event PresignHash(address indexed sender, bytes32 indexed hash)
func (_Aragon *AragonFilterer) FilterPresignHash(opts *bind.FilterOpts, sender []common.Address, hash [][32]byte) (*AragonPresignHashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "PresignHash", senderRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &AragonPresignHashIterator{contract: _Aragon.contract, event: "PresignHash", logs: logs, sub: sub}, nil
}

// WatchPresignHash is a free log subscription operation binding the contract event 0xb150212e573caa2a0216cf8f273ef996af143fd7b2d35abd92c8105536cac160.
//
// Solidity: event PresignHash(address indexed sender, bytes32 indexed hash)
func (_Aragon *AragonFilterer) WatchPresignHash(opts *bind.WatchOpts, sink chan<- *AragonPresignHash, sender []common.Address, hash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "PresignHash", senderRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonPresignHash)
				if err := _Aragon.contract.UnpackLog(event, "PresignHash", log); err != nil {
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

// ParsePresignHash is a log parse operation binding the contract event 0xb150212e573caa2a0216cf8f273ef996af143fd7b2d35abd92c8105536cac160.
//
// Solidity: event PresignHash(address indexed sender, bytes32 indexed hash)
func (_Aragon *AragonFilterer) ParsePresignHash(log types.Log) (*AragonPresignHash, error) {
	event := new(AragonPresignHash)
	if err := _Aragon.contract.UnpackLog(event, "PresignHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonReceiveERC721Iterator is returned from FilterReceiveERC721 and is used to iterate over the raw logs and unpacked data for ReceiveERC721 events raised by the Aragon contract.
type AragonReceiveERC721Iterator struct {
	Event *AragonReceiveERC721 // Event containing the contract specifics and raw log

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
func (it *AragonReceiveERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonReceiveERC721)
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
		it.Event = new(AragonReceiveERC721)
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
func (it *AragonReceiveERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonReceiveERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonReceiveERC721 represents a ReceiveERC721 event raised by the Aragon contract.
type AragonReceiveERC721 struct {
	Token    common.Address
	Operator common.Address
	From     common.Address
	TokenId  *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveERC721 is a free log retrieval operation binding the contract event 0xe0d2ab3bf6896c073bb33b920a9b60f3c8207b3cc3b7561c3101cb081a8f0883.
//
// Solidity: event ReceiveERC721(address indexed token, address indexed operator, address indexed from, uint256 tokenId, bytes data)
func (_Aragon *AragonFilterer) FilterReceiveERC721(opts *bind.FilterOpts, token []common.Address, operator []common.Address, from []common.Address) (*AragonReceiveERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "ReceiveERC721", tokenRule, operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &AragonReceiveERC721Iterator{contract: _Aragon.contract, event: "ReceiveERC721", logs: logs, sub: sub}, nil
}

// WatchReceiveERC721 is a free log subscription operation binding the contract event 0xe0d2ab3bf6896c073bb33b920a9b60f3c8207b3cc3b7561c3101cb081a8f0883.
//
// Solidity: event ReceiveERC721(address indexed token, address indexed operator, address indexed from, uint256 tokenId, bytes data)
func (_Aragon *AragonFilterer) WatchReceiveERC721(opts *bind.WatchOpts, sink chan<- *AragonReceiveERC721, token []common.Address, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "ReceiveERC721", tokenRule, operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonReceiveERC721)
				if err := _Aragon.contract.UnpackLog(event, "ReceiveERC721", log); err != nil {
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

// ParseReceiveERC721 is a log parse operation binding the contract event 0xe0d2ab3bf6896c073bb33b920a9b60f3c8207b3cc3b7561c3101cb081a8f0883.
//
// Solidity: event ReceiveERC721(address indexed token, address indexed operator, address indexed from, uint256 tokenId, bytes data)
func (_Aragon *AragonFilterer) ParseReceiveERC721(log types.Log) (*AragonReceiveERC721, error) {
	event := new(AragonReceiveERC721)
	if err := _Aragon.contract.UnpackLog(event, "ReceiveERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the Aragon contract.
type AragonRecoverToVaultIterator struct {
	Event *AragonRecoverToVault // Event containing the contract specifics and raw log

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
func (it *AragonRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonRecoverToVault)
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
		it.Event = new(AragonRecoverToVault)
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
func (it *AragonRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonRecoverToVault represents a RecoverToVault event raised by the Aragon contract.
type AragonRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Aragon *AragonFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*AragonRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AragonRecoverToVaultIterator{contract: _Aragon.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Aragon *AragonFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *AragonRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonRecoverToVault)
				if err := _Aragon.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Aragon *AragonFilterer) ParseRecoverToVault(log types.Log) (*AragonRecoverToVault, error) {
	event := new(AragonRecoverToVault)
	if err := _Aragon.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonRemoveProtectedTokenIterator is returned from FilterRemoveProtectedToken and is used to iterate over the raw logs and unpacked data for RemoveProtectedToken events raised by the Aragon contract.
type AragonRemoveProtectedTokenIterator struct {
	Event *AragonRemoveProtectedToken // Event containing the contract specifics and raw log

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
func (it *AragonRemoveProtectedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonRemoveProtectedToken)
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
		it.Event = new(AragonRemoveProtectedToken)
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
func (it *AragonRemoveProtectedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonRemoveProtectedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonRemoveProtectedToken represents a RemoveProtectedToken event raised by the Aragon contract.
type AragonRemoveProtectedToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveProtectedToken is a free log retrieval operation binding the contract event 0x3da25279c93c5b22b359bebff8b5ddbfd9b0506be8344b93c9a7dc999459fe04.
//
// Solidity: event RemoveProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) FilterRemoveProtectedToken(opts *bind.FilterOpts, token []common.Address) (*AragonRemoveProtectedTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "RemoveProtectedToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &AragonRemoveProtectedTokenIterator{contract: _Aragon.contract, event: "RemoveProtectedToken", logs: logs, sub: sub}, nil
}

// WatchRemoveProtectedToken is a free log subscription operation binding the contract event 0x3da25279c93c5b22b359bebff8b5ddbfd9b0506be8344b93c9a7dc999459fe04.
//
// Solidity: event RemoveProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) WatchRemoveProtectedToken(opts *bind.WatchOpts, sink chan<- *AragonRemoveProtectedToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "RemoveProtectedToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonRemoveProtectedToken)
				if err := _Aragon.contract.UnpackLog(event, "RemoveProtectedToken", log); err != nil {
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

// ParseRemoveProtectedToken is a log parse operation binding the contract event 0x3da25279c93c5b22b359bebff8b5ddbfd9b0506be8344b93c9a7dc999459fe04.
//
// Solidity: event RemoveProtectedToken(address indexed token)
func (_Aragon *AragonFilterer) ParseRemoveProtectedToken(log types.Log) (*AragonRemoveProtectedToken, error) {
	event := new(AragonRemoveProtectedToken)
	if err := _Aragon.contract.UnpackLog(event, "RemoveProtectedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonSafeExecuteIterator is returned from FilterSafeExecute and is used to iterate over the raw logs and unpacked data for SafeExecute events raised by the Aragon contract.
type AragonSafeExecuteIterator struct {
	Event *AragonSafeExecute // Event containing the contract specifics and raw log

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
func (it *AragonSafeExecuteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonSafeExecute)
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
		it.Event = new(AragonSafeExecute)
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
func (it *AragonSafeExecuteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonSafeExecuteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonSafeExecute represents a SafeExecute event raised by the Aragon contract.
type AragonSafeExecute struct {
	Sender common.Address
	Target common.Address
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeExecute is a free log retrieval operation binding the contract event 0x47f4287d5285559a03d8affeadc53ef3fca238a63c256c08a22bfd30fd5d33ce.
//
// Solidity: event SafeExecute(address indexed sender, address indexed target, bytes data)
func (_Aragon *AragonFilterer) FilterSafeExecute(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*AragonSafeExecuteIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "SafeExecute", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AragonSafeExecuteIterator{contract: _Aragon.contract, event: "SafeExecute", logs: logs, sub: sub}, nil
}

// WatchSafeExecute is a free log subscription operation binding the contract event 0x47f4287d5285559a03d8affeadc53ef3fca238a63c256c08a22bfd30fd5d33ce.
//
// Solidity: event SafeExecute(address indexed sender, address indexed target, bytes data)
func (_Aragon *AragonFilterer) WatchSafeExecute(opts *bind.WatchOpts, sink chan<- *AragonSafeExecute, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "SafeExecute", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonSafeExecute)
				if err := _Aragon.contract.UnpackLog(event, "SafeExecute", log); err != nil {
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

// ParseSafeExecute is a log parse operation binding the contract event 0x47f4287d5285559a03d8affeadc53ef3fca238a63c256c08a22bfd30fd5d33ce.
//
// Solidity: event SafeExecute(address indexed sender, address indexed target, bytes data)
func (_Aragon *AragonFilterer) ParseSafeExecute(log types.Log) (*AragonSafeExecute, error) {
	event := new(AragonSafeExecute)
	if err := _Aragon.contract.UnpackLog(event, "SafeExecute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the Aragon contract.
type AragonScriptResultIterator struct {
	Event *AragonScriptResult // Event containing the contract specifics and raw log

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
func (it *AragonScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonScriptResult)
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
		it.Event = new(AragonScriptResult)
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
func (it *AragonScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonScriptResult represents a ScriptResult event raised by the Aragon contract.
type AragonScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Aragon *AragonFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*AragonScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &AragonScriptResultIterator{contract: _Aragon.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Aragon *AragonFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *AragonScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonScriptResult)
				if err := _Aragon.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Aragon *AragonFilterer) ParseScriptResult(log types.Log) (*AragonScriptResult, error) {
	event := new(AragonScriptResult)
	if err := _Aragon.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonSetDesignatedSignerIterator is returned from FilterSetDesignatedSigner and is used to iterate over the raw logs and unpacked data for SetDesignatedSigner events raised by the Aragon contract.
type AragonSetDesignatedSignerIterator struct {
	Event *AragonSetDesignatedSigner // Event containing the contract specifics and raw log

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
func (it *AragonSetDesignatedSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonSetDesignatedSigner)
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
		it.Event = new(AragonSetDesignatedSigner)
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
func (it *AragonSetDesignatedSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonSetDesignatedSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonSetDesignatedSigner represents a SetDesignatedSigner event raised by the Aragon contract.
type AragonSetDesignatedSigner struct {
	Sender    common.Address
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetDesignatedSigner is a free log retrieval operation binding the contract event 0x3ecf54ed9acd859c5ee7f080794267b8f08b65d2446d816cef1efccd6d00d735.
//
// Solidity: event SetDesignatedSigner(address indexed sender, address indexed oldSigner, address indexed newSigner)
func (_Aragon *AragonFilterer) FilterSetDesignatedSigner(opts *bind.FilterOpts, sender []common.Address, oldSigner []common.Address, newSigner []common.Address) (*AragonSetDesignatedSignerIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "SetDesignatedSigner", senderRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &AragonSetDesignatedSignerIterator{contract: _Aragon.contract, event: "SetDesignatedSigner", logs: logs, sub: sub}, nil
}

// WatchSetDesignatedSigner is a free log subscription operation binding the contract event 0x3ecf54ed9acd859c5ee7f080794267b8f08b65d2446d816cef1efccd6d00d735.
//
// Solidity: event SetDesignatedSigner(address indexed sender, address indexed oldSigner, address indexed newSigner)
func (_Aragon *AragonFilterer) WatchSetDesignatedSigner(opts *bind.WatchOpts, sink chan<- *AragonSetDesignatedSigner, sender []common.Address, oldSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var oldSignerRule []interface{}
	for _, oldSignerItem := range oldSigner {
		oldSignerRule = append(oldSignerRule, oldSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "SetDesignatedSigner", senderRule, oldSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonSetDesignatedSigner)
				if err := _Aragon.contract.UnpackLog(event, "SetDesignatedSigner", log); err != nil {
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

// ParseSetDesignatedSigner is a log parse operation binding the contract event 0x3ecf54ed9acd859c5ee7f080794267b8f08b65d2446d816cef1efccd6d00d735.
//
// Solidity: event SetDesignatedSigner(address indexed sender, address indexed oldSigner, address indexed newSigner)
func (_Aragon *AragonFilterer) ParseSetDesignatedSigner(log types.Log) (*AragonSetDesignatedSigner, error) {
	event := new(AragonSetDesignatedSigner)
	if err := _Aragon.contract.UnpackLog(event, "SetDesignatedSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonVaultDepositIterator is returned from FilterVaultDeposit and is used to iterate over the raw logs and unpacked data for VaultDeposit events raised by the Aragon contract.
type AragonVaultDepositIterator struct {
	Event *AragonVaultDeposit // Event containing the contract specifics and raw log

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
func (it *AragonVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonVaultDeposit)
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
		it.Event = new(AragonVaultDeposit)
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
func (it *AragonVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonVaultDeposit represents a VaultDeposit event raised by the Aragon contract.
type AragonVaultDeposit struct {
	Token  common.Address
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultDeposit is a free log retrieval operation binding the contract event 0x2790b90165fd3973ad7edde4eca71b4f8808dd4857a2a3a3e8ae5642a5cb196e.
//
// Solidity: event VaultDeposit(address indexed token, address indexed sender, uint256 amount)
func (_Aragon *AragonFilterer) FilterVaultDeposit(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*AragonVaultDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "VaultDeposit", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AragonVaultDepositIterator{contract: _Aragon.contract, event: "VaultDeposit", logs: logs, sub: sub}, nil
}

// WatchVaultDeposit is a free log subscription operation binding the contract event 0x2790b90165fd3973ad7edde4eca71b4f8808dd4857a2a3a3e8ae5642a5cb196e.
//
// Solidity: event VaultDeposit(address indexed token, address indexed sender, uint256 amount)
func (_Aragon *AragonFilterer) WatchVaultDeposit(opts *bind.WatchOpts, sink chan<- *AragonVaultDeposit, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "VaultDeposit", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonVaultDeposit)
				if err := _Aragon.contract.UnpackLog(event, "VaultDeposit", log); err != nil {
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

// ParseVaultDeposit is a log parse operation binding the contract event 0x2790b90165fd3973ad7edde4eca71b4f8808dd4857a2a3a3e8ae5642a5cb196e.
//
// Solidity: event VaultDeposit(address indexed token, address indexed sender, uint256 amount)
func (_Aragon *AragonFilterer) ParseVaultDeposit(log types.Log) (*AragonVaultDeposit, error) {
	event := new(AragonVaultDeposit)
	if err := _Aragon.contract.UnpackLog(event, "VaultDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AragonVaultTransferIterator is returned from FilterVaultTransfer and is used to iterate over the raw logs and unpacked data for VaultTransfer events raised by the Aragon contract.
type AragonVaultTransferIterator struct {
	Event *AragonVaultTransfer // Event containing the contract specifics and raw log

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
func (it *AragonVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AragonVaultTransfer)
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
		it.Event = new(AragonVaultTransfer)
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
func (it *AragonVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AragonVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AragonVaultTransfer represents a VaultTransfer event raised by the Aragon contract.
type AragonVaultTransfer struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultTransfer is a free log retrieval operation binding the contract event 0x239e7f6cdac8fb35a788a46b431b54da87de90b82448a2c294be5e92a6e579af.
//
// Solidity: event VaultTransfer(address indexed token, address indexed to, uint256 amount)
func (_Aragon *AragonFilterer) FilterVaultTransfer(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*AragonVaultTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Aragon.contract.FilterLogs(opts, "VaultTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AragonVaultTransferIterator{contract: _Aragon.contract, event: "VaultTransfer", logs: logs, sub: sub}, nil
}

// WatchVaultTransfer is a free log subscription operation binding the contract event 0x239e7f6cdac8fb35a788a46b431b54da87de90b82448a2c294be5e92a6e579af.
//
// Solidity: event VaultTransfer(address indexed token, address indexed to, uint256 amount)
func (_Aragon *AragonFilterer) WatchVaultTransfer(opts *bind.WatchOpts, sink chan<- *AragonVaultTransfer, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Aragon.contract.WatchLogs(opts, "VaultTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AragonVaultTransfer)
				if err := _Aragon.contract.UnpackLog(event, "VaultTransfer", log); err != nil {
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

// ParseVaultTransfer is a log parse operation binding the contract event 0x239e7f6cdac8fb35a788a46b431b54da87de90b82448a2c294be5e92a6e579af.
//
// Solidity: event VaultTransfer(address indexed token, address indexed to, uint256 amount)
func (_Aragon *AragonFilterer) ParseVaultTransfer(log types.Log) (*AragonVaultTransfer, error) {
	event := new(AragonVaultTransfer)
	if err := _Aragon.contract.UnpackLog(event, "VaultTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
