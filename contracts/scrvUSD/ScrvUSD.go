// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scrvUSD

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
	Activation  *big.Int
	LastReport  *big.Int
	CurrentDebt *big.Int
	MaxDebt     *big.Int
}

// ScrvUSDMetaData contains all meta data concerning the ScrvUSD contract.
var ScrvUSDMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyChanged\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"change_type\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StrategyReported\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"gain\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"loss\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"protocol_fees\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"total_fees\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"total_refunds\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DebtUpdated\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"current_debt\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RoleSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"name\":\"role\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateFutureRoleManager\",\"inputs\":[{\"name\":\"future_role_manager\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateRoleManager\",\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateAccountant\",\"inputs\":[{\"name\":\"accountant\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimitModule\",\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateWithdrawLimitModule\",\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDefaultQueue\",\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateUseDefaultQueue\",\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateAutoAllocate\",\"inputs\":[{\"name\":\"auto_allocate\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatedMaxDebtForStrategy\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"new_debt\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateDepositLimit\",\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateMinimumTotalIdle\",\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateProfitMaxUnlockTime\",\"inputs\":[{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"DebtPurchased\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Shutdown\",\"inputs\":[],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_accountant\",\"inputs\":[{\"name\":\"new_accountant\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_default_queue\",\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_use_default_queue\",\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_auto_allocate\",\"inputs\":[{\"name\":\"auto_allocate\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit\",\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit\",\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"},{\"name\":\"override\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit_module\",\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_deposit_limit_module\",\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"},{\"name\":\"override\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_withdraw_limit_module\",\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_minimum_total_idle\",\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setProfitMaxUnlockTime\",\"inputs\":[{\"name\":\"new_profit_max_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_role\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer_role_manager\",\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_role_manager\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isShutdown\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"unlockedShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pricePerShare\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_default_queue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"process_report\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"buy_debt\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_strategy\",\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_strategy\",\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"},{\"name\":\"add_to_queue\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revoke_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"force_revoke_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_max_debt_for_strategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"new_max_debt\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_debt\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_debt\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"shutdown_vault\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalIdle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalDebt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"apiVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"assess_share_of_unrealised_losses\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"assets_needed\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"profitMaxUnlockTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fullProfitUnlockDate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"profitUnlockingRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lastProfitUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"strategies\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"last_report\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"},{\"name\":\"max_debt\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"default_queue\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"use_default_queue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"auto_allocate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minimum_total_idle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"deposit_limit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"accountant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"deposit_limit_module\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"withdraw_limit_module\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"roles\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"role_manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_role_manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// ScrvUSDABI is the input ABI used to generate the binding from.
// Deprecated: Use ScrvUSDMetaData.ABI instead.
var ScrvUSDABI = ScrvUSDMetaData.ABI

// ScrvUSD is an auto generated Go binding around an Ethereum contract.
type ScrvUSD struct {
	ScrvUSDCaller     // Read-only binding to the contract
	ScrvUSDTransactor // Write-only binding to the contract
	ScrvUSDFilterer   // Log filterer for contract events
}

// ScrvUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScrvUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrvUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScrvUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrvUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScrvUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrvUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScrvUSDSession struct {
	Contract     *ScrvUSD          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScrvUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScrvUSDCallerSession struct {
	Contract *ScrvUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ScrvUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScrvUSDTransactorSession struct {
	Contract     *ScrvUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ScrvUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScrvUSDRaw struct {
	Contract *ScrvUSD // Generic contract binding to access the raw methods on
}

// ScrvUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScrvUSDCallerRaw struct {
	Contract *ScrvUSDCaller // Generic read-only contract binding to access the raw methods on
}

// ScrvUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScrvUSDTransactorRaw struct {
	Contract *ScrvUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScrvUSD creates a new instance of ScrvUSD, bound to a specific deployed contract.
func NewScrvUSD(address common.Address, backend bind.ContractBackend) (*ScrvUSD, error) {
	contract, err := bindScrvUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrvUSD{ScrvUSDCaller: ScrvUSDCaller{contract: contract}, ScrvUSDTransactor: ScrvUSDTransactor{contract: contract}, ScrvUSDFilterer: ScrvUSDFilterer{contract: contract}}, nil
}

// NewScrvUSDCaller creates a new read-only instance of ScrvUSD, bound to a specific deployed contract.
func NewScrvUSDCaller(address common.Address, caller bind.ContractCaller) (*ScrvUSDCaller, error) {
	contract, err := bindScrvUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDCaller{contract: contract}, nil
}

// NewScrvUSDTransactor creates a new write-only instance of ScrvUSD, bound to a specific deployed contract.
func NewScrvUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrvUSDTransactor, error) {
	contract, err := bindScrvUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDTransactor{contract: contract}, nil
}

// NewScrvUSDFilterer creates a new log filterer instance of ScrvUSD, bound to a specific deployed contract.
func NewScrvUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrvUSDFilterer, error) {
	contract, err := bindScrvUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDFilterer{contract: contract}, nil
}

// bindScrvUSD binds a generic wrapper to an already deployed contract.
func bindScrvUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScrvUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrvUSD *ScrvUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrvUSD.Contract.ScrvUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrvUSD *ScrvUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ScrvUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrvUSD *ScrvUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ScrvUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrvUSD *ScrvUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrvUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrvUSD *ScrvUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrvUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrvUSD *ScrvUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrvUSD.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ScrvUSD *ScrvUSDCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ScrvUSD *ScrvUSDSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ScrvUSD.Contract.DOMAINSEPARATOR(&_ScrvUSD.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ScrvUSD *ScrvUSDCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ScrvUSD.Contract.DOMAINSEPARATOR(&_ScrvUSD.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_ScrvUSD *ScrvUSDSession) FACTORY() (common.Address, error) {
	return _ScrvUSD.Contract.FACTORY(&_ScrvUSD.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) FACTORY() (common.Address, error) {
	return _ScrvUSD.Contract.FACTORY(&_ScrvUSD.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "accountant")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_ScrvUSD *ScrvUSDSession) Accountant() (common.Address, error) {
	return _ScrvUSD.Contract.Accountant(&_ScrvUSD.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) Accountant() (common.Address, error) {
	return _ScrvUSD.Contract.Accountant(&_ScrvUSD.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Allowance(&_ScrvUSD.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Allowance(&_ScrvUSD.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ScrvUSD *ScrvUSDCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ScrvUSD *ScrvUSDSession) ApiVersion() (string, error) {
	return _ScrvUSD.Contract.ApiVersion(&_ScrvUSD.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() pure returns(string)
func (_ScrvUSD *ScrvUSDCallerSession) ApiVersion() (string, error) {
	return _ScrvUSD.Contract.ApiVersion(&_ScrvUSD.CallOpts)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) AssessShareOfUnrealisedLosses(opts *bind.CallOpts, strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "assess_share_of_unrealised_losses", strategy, assets_needed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.AssessShareOfUnrealisedLosses(&_ScrvUSD.CallOpts, strategy, assets_needed)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.AssessShareOfUnrealisedLosses(&_ScrvUSD.CallOpts, strategy, assets_needed)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ScrvUSD *ScrvUSDSession) Asset() (common.Address, error) {
	return _ScrvUSD.Contract.Asset(&_ScrvUSD.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) Asset() (common.Address, error) {
	return _ScrvUSD.Contract.Asset(&_ScrvUSD.CallOpts)
}

// AutoAllocate is a free data retrieval call binding the contract method 0xdd47db90.
//
// Solidity: function auto_allocate() view returns(bool)
func (_ScrvUSD *ScrvUSDCaller) AutoAllocate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "auto_allocate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutoAllocate is a free data retrieval call binding the contract method 0xdd47db90.
//
// Solidity: function auto_allocate() view returns(bool)
func (_ScrvUSD *ScrvUSDSession) AutoAllocate() (bool, error) {
	return _ScrvUSD.Contract.AutoAllocate(&_ScrvUSD.CallOpts)
}

// AutoAllocate is a free data retrieval call binding the contract method 0xdd47db90.
//
// Solidity: function auto_allocate() view returns(bool)
func (_ScrvUSD *ScrvUSDCallerSession) AutoAllocate() (bool, error) {
	return _ScrvUSD.Contract.AutoAllocate(&_ScrvUSD.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.BalanceOf(&_ScrvUSD.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.BalanceOf(&_ScrvUSD.CallOpts, addr)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.ConvertToAssets(&_ScrvUSD.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.ConvertToAssets(&_ScrvUSD.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.ConvertToShares(&_ScrvUSD.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.ConvertToShares(&_ScrvUSD.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ScrvUSD *ScrvUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ScrvUSD *ScrvUSDSession) Decimals() (uint8, error) {
	return _ScrvUSD.Contract.Decimals(&_ScrvUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ScrvUSD *ScrvUSDCallerSession) Decimals() (uint8, error) {
	return _ScrvUSD.Contract.Decimals(&_ScrvUSD.CallOpts)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_ScrvUSD *ScrvUSDCaller) DefaultQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "default_queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_ScrvUSD *ScrvUSDSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _ScrvUSD.Contract.DefaultQueue(&_ScrvUSD.CallOpts, arg0)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _ScrvUSD.Contract.DefaultQueue(&_ScrvUSD.CallOpts, arg0)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "deposit_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) DepositLimit() (*big.Int, error) {
	return _ScrvUSD.Contract.DepositLimit(&_ScrvUSD.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) DepositLimit() (*big.Int, error) {
	return _ScrvUSD.Contract.DepositLimit(&_ScrvUSD.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) DepositLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "deposit_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDSession) DepositLimitModule() (common.Address, error) {
	return _ScrvUSD.Contract.DepositLimitModule(&_ScrvUSD.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) DepositLimitModule() (common.Address, error) {
	return _ScrvUSD.Contract.DepositLimitModule(&_ScrvUSD.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) FullProfitUnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "fullProfitUnlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) FullProfitUnlockDate() (*big.Int, error) {
	return _ScrvUSD.Contract.FullProfitUnlockDate(&_ScrvUSD.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) FullProfitUnlockDate() (*big.Int, error) {
	return _ScrvUSD.Contract.FullProfitUnlockDate(&_ScrvUSD.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) FutureRoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "future_role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDSession) FutureRoleManager() (common.Address, error) {
	return _ScrvUSD.Contract.FutureRoleManager(&_ScrvUSD.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) FutureRoleManager() (common.Address, error) {
	return _ScrvUSD.Contract.FutureRoleManager(&_ScrvUSD.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_ScrvUSD *ScrvUSDCaller) GetDefaultQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "get_default_queue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_ScrvUSD *ScrvUSDSession) GetDefaultQueue() ([]common.Address, error) {
	return _ScrvUSD.Contract.GetDefaultQueue(&_ScrvUSD.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_ScrvUSD *ScrvUSDCallerSession) GetDefaultQueue() ([]common.Address, error) {
	return _ScrvUSD.Contract.GetDefaultQueue(&_ScrvUSD.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ScrvUSD *ScrvUSDCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ScrvUSD *ScrvUSDSession) IsShutdown() (bool, error) {
	return _ScrvUSD.Contract.IsShutdown(&_ScrvUSD.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ScrvUSD *ScrvUSDCallerSession) IsShutdown() (bool, error) {
	return _ScrvUSD.Contract.IsShutdown(&_ScrvUSD.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) LastProfitUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "lastProfitUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) LastProfitUpdate() (*big.Int, error) {
	return _ScrvUSD.Contract.LastProfitUpdate(&_ScrvUSD.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) LastProfitUpdate() (*big.Int, error) {
	return _ScrvUSD.Contract.LastProfitUpdate(&_ScrvUSD.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxDeposit(&_ScrvUSD.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxDeposit(&_ScrvUSD.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxMint(&_ScrvUSD.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxMint(&_ScrvUSD.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem(&_ScrvUSD.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem(&_ScrvUSD.CallOpts, owner)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxRedeem0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxRedeem0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem0(&_ScrvUSD.CallOpts, owner, max_loss)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem0(&_ScrvUSD.CallOpts, owner, max_loss)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxRedeem1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxRedeem1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem1(&_ScrvUSD.CallOpts, owner, max_loss, strategies)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxRedeem1(&_ScrvUSD.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw(&_ScrvUSD.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw(&_ScrvUSD.CallOpts, owner)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxWithdraw0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxWithdraw0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw0(&_ScrvUSD.CallOpts, owner, max_loss)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw0(&_ScrvUSD.CallOpts, owner, max_loss)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MaxWithdraw1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "maxWithdraw1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw1(&_ScrvUSD.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.MaxWithdraw1(&_ScrvUSD.CallOpts, owner, max_loss, strategies)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) MinimumTotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "minimum_total_idle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) MinimumTotalIdle() (*big.Int, error) {
	return _ScrvUSD.Contract.MinimumTotalIdle(&_ScrvUSD.CallOpts)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) MinimumTotalIdle() (*big.Int, error) {
	return _ScrvUSD.Contract.MinimumTotalIdle(&_ScrvUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScrvUSD *ScrvUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScrvUSD *ScrvUSDSession) Name() (string, error) {
	return _ScrvUSD.Contract.Name(&_ScrvUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ScrvUSD *ScrvUSDCallerSession) Name() (string, error) {
	return _ScrvUSD.Contract.Name(&_ScrvUSD.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Nonces(&_ScrvUSD.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Nonces(&_ScrvUSD.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewDeposit(&_ScrvUSD.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewDeposit(&_ScrvUSD.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewMint(&_ScrvUSD.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewMint(&_ScrvUSD.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewRedeem(&_ScrvUSD.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewRedeem(&_ScrvUSD.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewWithdraw(&_ScrvUSD.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ScrvUSD.Contract.PreviewWithdraw(&_ScrvUSD.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) PricePerShare() (*big.Int, error) {
	return _ScrvUSD.Contract.PricePerShare(&_ScrvUSD.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) PricePerShare() (*big.Int, error) {
	return _ScrvUSD.Contract.PricePerShare(&_ScrvUSD.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _ScrvUSD.Contract.ProfitMaxUnlockTime(&_ScrvUSD.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _ScrvUSD.Contract.ProfitMaxUnlockTime(&_ScrvUSD.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) ProfitUnlockingRate() (*big.Int, error) {
	return _ScrvUSD.Contract.ProfitUnlockingRate(&_ScrvUSD.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _ScrvUSD.Contract.ProfitUnlockingRate(&_ScrvUSD.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDSession) RoleManager() (common.Address, error) {
	return _ScrvUSD.Contract.RoleManager(&_ScrvUSD.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) RoleManager() (common.Address, error) {
	return _ScrvUSD.Contract.RoleManager(&_ScrvUSD.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) Roles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Roles(&_ScrvUSD.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _ScrvUSD.Contract.Roles(&_ScrvUSD.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_ScrvUSD *ScrvUSDCaller) Strategies(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_ScrvUSD *ScrvUSDSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _ScrvUSD.Contract.Strategies(&_ScrvUSD.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_ScrvUSD *ScrvUSDCallerSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _ScrvUSD.Contract.Strategies(&_ScrvUSD.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScrvUSD *ScrvUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScrvUSD *ScrvUSDSession) Symbol() (string, error) {
	return _ScrvUSD.Contract.Symbol(&_ScrvUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ScrvUSD *ScrvUSDCallerSession) Symbol() (string, error) {
	return _ScrvUSD.Contract.Symbol(&_ScrvUSD.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) TotalAssets() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalAssets(&_ScrvUSD.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) TotalAssets() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalAssets(&_ScrvUSD.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) TotalDebt() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalDebt(&_ScrvUSD.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) TotalDebt() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalDebt(&_ScrvUSD.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) TotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "totalIdle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) TotalIdle() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalIdle(&_ScrvUSD.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) TotalIdle() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalIdle(&_ScrvUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) TotalSupply() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalSupply(&_ScrvUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _ScrvUSD.Contract.TotalSupply(&_ScrvUSD.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_ScrvUSD *ScrvUSDCaller) UnlockedShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "unlockedShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_ScrvUSD *ScrvUSDSession) UnlockedShares() (*big.Int, error) {
	return _ScrvUSD.Contract.UnlockedShares(&_ScrvUSD.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_ScrvUSD *ScrvUSDCallerSession) UnlockedShares() (*big.Int, error) {
	return _ScrvUSD.Contract.UnlockedShares(&_ScrvUSD.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_ScrvUSD *ScrvUSDCaller) UseDefaultQueue(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "use_default_queue")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_ScrvUSD *ScrvUSDSession) UseDefaultQueue() (bool, error) {
	return _ScrvUSD.Contract.UseDefaultQueue(&_ScrvUSD.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_ScrvUSD *ScrvUSDCallerSession) UseDefaultQueue() (bool, error) {
	return _ScrvUSD.Contract.UseDefaultQueue(&_ScrvUSD.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDCaller) WithdrawLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrvUSD.contract.Call(opts, &out, "withdraw_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDSession) WithdrawLimitModule() (common.Address, error) {
	return _ScrvUSD.Contract.WithdrawLimitModule(&_ScrvUSD.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_ScrvUSD *ScrvUSDCallerSession) WithdrawLimitModule() (common.Address, error) {
	return _ScrvUSD.Contract.WithdrawLimitModule(&_ScrvUSD.CallOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_ScrvUSD *ScrvUSDTransactor) AcceptRoleManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "accept_role_manager")
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_ScrvUSD *ScrvUSDSession) AcceptRoleManager() (*types.Transaction, error) {
	return _ScrvUSD.Contract.AcceptRoleManager(&_ScrvUSD.TransactOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_ScrvUSD *ScrvUSDTransactorSession) AcceptRoleManager() (*types.Transaction, error) {
	return _ScrvUSD.Contract.AcceptRoleManager(&_ScrvUSD.TransactOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactor) AddRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "add_role", account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddRole(&_ScrvUSD.TransactOpts, account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddRole(&_ScrvUSD.TransactOpts, account, role)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_ScrvUSD *ScrvUSDTransactor) AddStrategy(opts *bind.TransactOpts, new_strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "add_strategy", new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_ScrvUSD *ScrvUSDSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddStrategy(&_ScrvUSD.TransactOpts, new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddStrategy(&_ScrvUSD.TransactOpts, new_strategy)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_ScrvUSD *ScrvUSDTransactor) AddStrategy0(opts *bind.TransactOpts, new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "add_strategy0", new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_ScrvUSD *ScrvUSDSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddStrategy0(&_ScrvUSD.TransactOpts, new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.AddStrategy0(&_ScrvUSD.TransactOpts, new_strategy, add_to_queue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Approve(&_ScrvUSD.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Approve(&_ScrvUSD.TransactOpts, spender, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_ScrvUSD *ScrvUSDTransactor) BuyDebt(opts *bind.TransactOpts, strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "buy_debt", strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_ScrvUSD *ScrvUSDSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.BuyDebt(&_ScrvUSD.TransactOpts, strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.BuyDebt(&_ScrvUSD.TransactOpts, strategy, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Deposit(&_ScrvUSD.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Deposit(&_ScrvUSD.TransactOpts, assets, receiver)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDTransactor) ForceRevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "force_revoke_strategy", strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ForceRevokeStrategy(&_ScrvUSD.TransactOpts, strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ForceRevokeStrategy(&_ScrvUSD.TransactOpts, strategy)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDTransactor) Initialize(opts *bind.TransactOpts, asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "initialize", asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Initialize(&_ScrvUSD.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Initialize(&_ScrvUSD.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Mint(&_ScrvUSD.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Mint(&_ScrvUSD.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_ScrvUSD *ScrvUSDTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_ScrvUSD *ScrvUSDSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Permit(&_ScrvUSD.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_ScrvUSD *ScrvUSDTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Permit(&_ScrvUSD.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_ScrvUSD *ScrvUSDTransactor) ProcessReport(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "process_report", strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_ScrvUSD *ScrvUSDSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ProcessReport(&_ScrvUSD.TransactOpts, strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.ProcessReport(&_ScrvUSD.TransactOpts, strategy)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem(&_ScrvUSD.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem(&_ScrvUSD.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "redeem0", shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem0(&_ScrvUSD.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem0(&_ScrvUSD.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Redeem1(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "redeem1", shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem1(&_ScrvUSD.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Redeem1(&_ScrvUSD.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactor) RemoveRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "remove_role", account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.RemoveRole(&_ScrvUSD.TransactOpts, account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.RemoveRole(&_ScrvUSD.TransactOpts, account, role)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDTransactor) RevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "revoke_strategy", strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.RevokeStrategy(&_ScrvUSD.TransactOpts, strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.RevokeStrategy(&_ScrvUSD.TransactOpts, strategy)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_ScrvUSD *ScrvUSDSession) SetName(name string) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetName(&_ScrvUSD.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetName(&_ScrvUSD.TransactOpts, name)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "setProfitMaxUnlockTime", new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetProfitMaxUnlockTime(&_ScrvUSD.TransactOpts, new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetProfitMaxUnlockTime(&_ScrvUSD.TransactOpts, new_profit_max_unlock_time)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetSymbol(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "setSymbol", symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_ScrvUSD *ScrvUSDSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetSymbol(&_ScrvUSD.TransactOpts, symbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetSymbol(symbol string) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetSymbol(&_ScrvUSD.TransactOpts, symbol)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetAccountant(opts *bind.TransactOpts, new_accountant common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_accountant", new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_ScrvUSD *ScrvUSDSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetAccountant(&_ScrvUSD.TransactOpts, new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetAccountant(&_ScrvUSD.TransactOpts, new_accountant)
}

// SetAutoAllocate is a paid mutator transaction binding the contract method 0x6f71047d.
//
// Solidity: function set_auto_allocate(bool auto_allocate) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetAutoAllocate(opts *bind.TransactOpts, auto_allocate bool) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_auto_allocate", auto_allocate)
}

// SetAutoAllocate is a paid mutator transaction binding the contract method 0x6f71047d.
//
// Solidity: function set_auto_allocate(bool auto_allocate) returns()
func (_ScrvUSD *ScrvUSDSession) SetAutoAllocate(auto_allocate bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetAutoAllocate(&_ScrvUSD.TransactOpts, auto_allocate)
}

// SetAutoAllocate is a paid mutator transaction binding the contract method 0x6f71047d.
//
// Solidity: function set_auto_allocate(bool auto_allocate) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetAutoAllocate(auto_allocate bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetAutoAllocate(&_ScrvUSD.TransactOpts, auto_allocate)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetDefaultQueue(opts *bind.TransactOpts, new_default_queue []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_default_queue", new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_ScrvUSD *ScrvUSDSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDefaultQueue(&_ScrvUSD.TransactOpts, new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDefaultQueue(&_ScrvUSD.TransactOpts, new_default_queue)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetDepositLimit(opts *bind.TransactOpts, deposit_limit *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_deposit_limit", deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_ScrvUSD *ScrvUSDSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimit(&_ScrvUSD.TransactOpts, deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimit(&_ScrvUSD.TransactOpts, deposit_limit)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetDepositLimit0(opts *bind.TransactOpts, deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_deposit_limit0", deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_ScrvUSD *ScrvUSDSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimit0(&_ScrvUSD.TransactOpts, deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimit0(&_ScrvUSD.TransactOpts, deposit_limit, override)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetDepositLimitModule(opts *bind.TransactOpts, deposit_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_deposit_limit_module", deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_ScrvUSD *ScrvUSDSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimitModule(&_ScrvUSD.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimitModule(&_ScrvUSD.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetDepositLimitModule0(opts *bind.TransactOpts, deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_deposit_limit_module0", deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_ScrvUSD *ScrvUSDSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimitModule0(&_ScrvUSD.TransactOpts, deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetDepositLimitModule0(&_ScrvUSD.TransactOpts, deposit_limit_module, override)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetMinimumTotalIdle(opts *bind.TransactOpts, minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_minimum_total_idle", minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_ScrvUSD *ScrvUSDSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetMinimumTotalIdle(&_ScrvUSD.TransactOpts, minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetMinimumTotalIdle(&_ScrvUSD.TransactOpts, minimum_total_idle)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_role", account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetRole(&_ScrvUSD.TransactOpts, account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetRole(&_ScrvUSD.TransactOpts, account, role)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetUseDefaultQueue(opts *bind.TransactOpts, use_default_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_use_default_queue", use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_ScrvUSD *ScrvUSDSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetUseDefaultQueue(&_ScrvUSD.TransactOpts, use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetUseDefaultQueue(&_ScrvUSD.TransactOpts, use_default_queue)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_ScrvUSD *ScrvUSDTransactor) SetWithdrawLimitModule(opts *bind.TransactOpts, withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "set_withdraw_limit_module", withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_ScrvUSD *ScrvUSDSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetWithdrawLimitModule(&_ScrvUSD.TransactOpts, withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.SetWithdrawLimitModule(&_ScrvUSD.TransactOpts, withdraw_limit_module)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_ScrvUSD *ScrvUSDTransactor) ShutdownVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "shutdown_vault")
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_ScrvUSD *ScrvUSDSession) ShutdownVault() (*types.Transaction, error) {
	return _ScrvUSD.Contract.ShutdownVault(&_ScrvUSD.TransactOpts)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_ScrvUSD *ScrvUSDTransactorSession) ShutdownVault() (*types.Transaction, error) {
	return _ScrvUSD.Contract.ShutdownVault(&_ScrvUSD.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Transfer(&_ScrvUSD.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Transfer(&_ScrvUSD.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.TransferFrom(&_ScrvUSD.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_ScrvUSD *ScrvUSDTransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.TransferFrom(&_ScrvUSD.TransactOpts, sender, receiver, amount)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_ScrvUSD *ScrvUSDTransactor) TransferRoleManager(opts *bind.TransactOpts, role_manager common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "transfer_role_manager", role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_ScrvUSD *ScrvUSDSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.TransferRoleManager(&_ScrvUSD.TransactOpts, role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.TransferRoleManager(&_ScrvUSD.TransactOpts, role_manager)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) UpdateDebt(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "update_debt", strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateDebt(&_ScrvUSD.TransactOpts, strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateDebt(&_ScrvUSD.TransactOpts, strategy, target_debt)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) UpdateDebt0(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "update_debt0", strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateDebt0(&_ScrvUSD.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateDebt0(&_ScrvUSD.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_ScrvUSD *ScrvUSDTransactor) UpdateMaxDebtForStrategy(opts *bind.TransactOpts, strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "update_max_debt_for_strategy", strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_ScrvUSD *ScrvUSDSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateMaxDebtForStrategy(&_ScrvUSD.TransactOpts, strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_ScrvUSD *ScrvUSDTransactorSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.UpdateMaxDebtForStrategy(&_ScrvUSD.TransactOpts, strategy, new_max_debt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw(&_ScrvUSD.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw(&_ScrvUSD.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "withdraw0", assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw0(&_ScrvUSD.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw0(&_ScrvUSD.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactor) Withdraw1(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.contract.Transact(opts, "withdraw1", assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw1(&_ScrvUSD.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_ScrvUSD *ScrvUSDTransactorSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _ScrvUSD.Contract.Withdraw1(&_ScrvUSD.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// ScrvUSDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ScrvUSD contract.
type ScrvUSDApprovalIterator struct {
	Event *ScrvUSDApproval // Event containing the contract specifics and raw log

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
func (it *ScrvUSDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDApproval)
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
		it.Event = new(ScrvUSDApproval)
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
func (it *ScrvUSDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDApproval represents a Approval event raised by the ScrvUSD contract.
type ScrvUSDApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ScrvUSDApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDApprovalIterator{contract: _ScrvUSD.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ScrvUSDApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDApproval)
				if err := _ScrvUSD.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) ParseApproval(log types.Log) (*ScrvUSDApproval, error) {
	event := new(ScrvUSDApproval)
	if err := _ScrvUSD.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDDebtPurchasedIterator is returned from FilterDebtPurchased and is used to iterate over the raw logs and unpacked data for DebtPurchased events raised by the ScrvUSD contract.
type ScrvUSDDebtPurchasedIterator struct {
	Event *ScrvUSDDebtPurchased // Event containing the contract specifics and raw log

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
func (it *ScrvUSDDebtPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDDebtPurchased)
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
		it.Event = new(ScrvUSDDebtPurchased)
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
func (it *ScrvUSDDebtPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDDebtPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDDebtPurchased represents a DebtPurchased event raised by the ScrvUSD contract.
type ScrvUSDDebtPurchased struct {
	Strategy common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebtPurchased is a free log retrieval operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_ScrvUSD *ScrvUSDFilterer) FilterDebtPurchased(opts *bind.FilterOpts, strategy []common.Address) (*ScrvUSDDebtPurchasedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDDebtPurchasedIterator{contract: _ScrvUSD.contract, event: "DebtPurchased", logs: logs, sub: sub}, nil
}

// WatchDebtPurchased is a free log subscription operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_ScrvUSD *ScrvUSDFilterer) WatchDebtPurchased(opts *bind.WatchOpts, sink chan<- *ScrvUSDDebtPurchased, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDDebtPurchased)
				if err := _ScrvUSD.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
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

// ParseDebtPurchased is a log parse operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_ScrvUSD *ScrvUSDFilterer) ParseDebtPurchased(log types.Log) (*ScrvUSDDebtPurchased, error) {
	event := new(ScrvUSDDebtPurchased)
	if err := _ScrvUSD.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDDebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the ScrvUSD contract.
type ScrvUSDDebtUpdatedIterator struct {
	Event *ScrvUSDDebtUpdated // Event containing the contract specifics and raw log

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
func (it *ScrvUSDDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDDebtUpdated)
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
		it.Event = new(ScrvUSDDebtUpdated)
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
func (it *ScrvUSDDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDDebtUpdated represents a DebtUpdated event raised by the ScrvUSD contract.
type ScrvUSDDebtUpdated struct {
	Strategy    common.Address
	CurrentDebt *big.Int
	NewDebt     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) FilterDebtUpdated(opts *bind.FilterOpts, strategy []common.Address) (*ScrvUSDDebtUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDDebtUpdatedIterator{contract: _ScrvUSD.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *ScrvUSDDebtUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDDebtUpdated)
				if err := _ScrvUSD.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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

// ParseDebtUpdated is a log parse operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) ParseDebtUpdated(log types.Log) (*ScrvUSDDebtUpdated, error) {
	event := new(ScrvUSDDebtUpdated)
	if err := _ScrvUSD.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ScrvUSD contract.
type ScrvUSDDepositIterator struct {
	Event *ScrvUSDDeposit // Event containing the contract specifics and raw log

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
func (it *ScrvUSDDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDDeposit)
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
		it.Event = new(ScrvUSDDeposit)
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
func (it *ScrvUSDDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDDeposit represents a Deposit event raised by the ScrvUSD contract.
type ScrvUSDDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ScrvUSDDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDDepositIterator{contract: _ScrvUSD.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ScrvUSDDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDDeposit)
				if err := _ScrvUSD.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) ParseDeposit(log types.Log) (*ScrvUSDDeposit, error) {
	event := new(ScrvUSDDeposit)
	if err := _ScrvUSD.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the ScrvUSD contract.
type ScrvUSDRoleSetIterator struct {
	Event *ScrvUSDRoleSet // Event containing the contract specifics and raw log

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
func (it *ScrvUSDRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDRoleSet)
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
		it.Event = new(ScrvUSDRoleSet)
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
func (it *ScrvUSDRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDRoleSet represents a RoleSet event raised by the ScrvUSD contract.
type ScrvUSDRoleSet struct {
	Account common.Address
	Role    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_ScrvUSD *ScrvUSDFilterer) FilterRoleSet(opts *bind.FilterOpts, account []common.Address, role []*big.Int) (*ScrvUSDRoleSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDRoleSetIterator{contract: _ScrvUSD.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_ScrvUSD *ScrvUSDFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *ScrvUSDRoleSet, account []common.Address, role []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDRoleSet)
				if err := _ScrvUSD.contract.UnpackLog(event, "RoleSet", log); err != nil {
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

// ParseRoleSet is a log parse operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_ScrvUSD *ScrvUSDFilterer) ParseRoleSet(log types.Log) (*ScrvUSDRoleSet, error) {
	event := new(ScrvUSDRoleSet)
	if err := _ScrvUSD.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the ScrvUSD contract.
type ScrvUSDShutdownIterator struct {
	Event *ScrvUSDShutdown // Event containing the contract specifics and raw log

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
func (it *ScrvUSDShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDShutdown)
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
		it.Event = new(ScrvUSDShutdown)
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
func (it *ScrvUSDShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDShutdown represents a Shutdown event raised by the ScrvUSD contract.
type ScrvUSDShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_ScrvUSD *ScrvUSDFilterer) FilterShutdown(opts *bind.FilterOpts) (*ScrvUSDShutdownIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDShutdownIterator{contract: _ScrvUSD.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_ScrvUSD *ScrvUSDFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *ScrvUSDShutdown) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDShutdown)
				if err := _ScrvUSD.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_ScrvUSD *ScrvUSDFilterer) ParseShutdown(log types.Log) (*ScrvUSDShutdown, error) {
	event := new(ScrvUSDShutdown)
	if err := _ScrvUSD.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDStrategyChangedIterator is returned from FilterStrategyChanged and is used to iterate over the raw logs and unpacked data for StrategyChanged events raised by the ScrvUSD contract.
type ScrvUSDStrategyChangedIterator struct {
	Event *ScrvUSDStrategyChanged // Event containing the contract specifics and raw log

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
func (it *ScrvUSDStrategyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDStrategyChanged)
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
		it.Event = new(ScrvUSDStrategyChanged)
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
func (it *ScrvUSDStrategyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDStrategyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDStrategyChanged represents a StrategyChanged event raised by the ScrvUSD contract.
type ScrvUSDStrategyChanged struct {
	Strategy   common.Address
	ChangeType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyChanged is a free log retrieval operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_ScrvUSD *ScrvUSDFilterer) FilterStrategyChanged(opts *bind.FilterOpts, strategy []common.Address, change_type []*big.Int) (*ScrvUSDStrategyChangedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDStrategyChangedIterator{contract: _ScrvUSD.contract, event: "StrategyChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyChanged is a free log subscription operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_ScrvUSD *ScrvUSDFilterer) WatchStrategyChanged(opts *bind.WatchOpts, sink chan<- *ScrvUSDStrategyChanged, strategy []common.Address, change_type []*big.Int) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDStrategyChanged)
				if err := _ScrvUSD.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
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

// ParseStrategyChanged is a log parse operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_ScrvUSD *ScrvUSDFilterer) ParseStrategyChanged(log types.Log) (*ScrvUSDStrategyChanged, error) {
	event := new(ScrvUSDStrategyChanged)
	if err := _ScrvUSD.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDStrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the ScrvUSD contract.
type ScrvUSDStrategyReportedIterator struct {
	Event *ScrvUSDStrategyReported // Event containing the contract specifics and raw log

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
func (it *ScrvUSDStrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDStrategyReported)
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
		it.Event = new(ScrvUSDStrategyReported)
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
func (it *ScrvUSDStrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDStrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDStrategyReported represents a StrategyReported event raised by the ScrvUSD contract.
type ScrvUSDStrategyReported struct {
	Strategy     common.Address
	Gain         *big.Int
	Loss         *big.Int
	CurrentDebt  *big.Int
	ProtocolFees *big.Int
	TotalFees    *big.Int
	TotalRefunds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_ScrvUSD *ScrvUSDFilterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*ScrvUSDStrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDStrategyReportedIterator{contract: _ScrvUSD.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_ScrvUSD *ScrvUSDFilterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *ScrvUSDStrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDStrategyReported)
				if err := _ScrvUSD.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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

// ParseStrategyReported is a log parse operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_ScrvUSD *ScrvUSDFilterer) ParseStrategyReported(log types.Log) (*ScrvUSDStrategyReported, error) {
	event := new(ScrvUSDStrategyReported)
	if err := _ScrvUSD.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ScrvUSD contract.
type ScrvUSDTransferIterator struct {
	Event *ScrvUSDTransfer // Event containing the contract specifics and raw log

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
func (it *ScrvUSDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDTransfer)
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
		it.Event = new(ScrvUSDTransfer)
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
func (it *ScrvUSDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDTransfer represents a Transfer event raised by the ScrvUSD contract.
type ScrvUSDTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ScrvUSDTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDTransferIterator{contract: _ScrvUSD.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ScrvUSDTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDTransfer)
				if err := _ScrvUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_ScrvUSD *ScrvUSDFilterer) ParseTransfer(log types.Log) (*ScrvUSDTransfer, error) {
	event := new(ScrvUSDTransfer)
	if err := _ScrvUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateAccountantIterator is returned from FilterUpdateAccountant and is used to iterate over the raw logs and unpacked data for UpdateAccountant events raised by the ScrvUSD contract.
type ScrvUSDUpdateAccountantIterator struct {
	Event *ScrvUSDUpdateAccountant // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateAccountant)
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
		it.Event = new(ScrvUSDUpdateAccountant)
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
func (it *ScrvUSDUpdateAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateAccountant represents a UpdateAccountant event raised by the ScrvUSD contract.
type ScrvUSDUpdateAccountant struct {
	Accountant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccountant is a free log retrieval operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateAccountant(opts *bind.FilterOpts, accountant []common.Address) (*ScrvUSDUpdateAccountantIterator, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateAccountantIterator{contract: _ScrvUSD.contract, event: "UpdateAccountant", logs: logs, sub: sub}, nil
}

// WatchUpdateAccountant is a free log subscription operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateAccountant(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateAccountant, accountant []common.Address) (event.Subscription, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateAccountant)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
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

// ParseUpdateAccountant is a log parse operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateAccountant(log types.Log) (*ScrvUSDUpdateAccountant, error) {
	event := new(ScrvUSDUpdateAccountant)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateAutoAllocateIterator is returned from FilterUpdateAutoAllocate and is used to iterate over the raw logs and unpacked data for UpdateAutoAllocate events raised by the ScrvUSD contract.
type ScrvUSDUpdateAutoAllocateIterator struct {
	Event *ScrvUSDUpdateAutoAllocate // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateAutoAllocateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateAutoAllocate)
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
		it.Event = new(ScrvUSDUpdateAutoAllocate)
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
func (it *ScrvUSDUpdateAutoAllocateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateAutoAllocateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateAutoAllocate represents a UpdateAutoAllocate event raised by the ScrvUSD contract.
type ScrvUSDUpdateAutoAllocate struct {
	AutoAllocate bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAutoAllocate is a free log retrieval operation binding the contract event 0x68d2a9e78bb2679fce38ced936827f483b51ff64f0776fc33fc39b2881f237f8.
//
// Solidity: event UpdateAutoAllocate(bool auto_allocate)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateAutoAllocate(opts *bind.FilterOpts) (*ScrvUSDUpdateAutoAllocateIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateAutoAllocate")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateAutoAllocateIterator{contract: _ScrvUSD.contract, event: "UpdateAutoAllocate", logs: logs, sub: sub}, nil
}

// WatchUpdateAutoAllocate is a free log subscription operation binding the contract event 0x68d2a9e78bb2679fce38ced936827f483b51ff64f0776fc33fc39b2881f237f8.
//
// Solidity: event UpdateAutoAllocate(bool auto_allocate)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateAutoAllocate(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateAutoAllocate) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateAutoAllocate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateAutoAllocate)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateAutoAllocate", log); err != nil {
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

// ParseUpdateAutoAllocate is a log parse operation binding the contract event 0x68d2a9e78bb2679fce38ced936827f483b51ff64f0776fc33fc39b2881f237f8.
//
// Solidity: event UpdateAutoAllocate(bool auto_allocate)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateAutoAllocate(log types.Log) (*ScrvUSDUpdateAutoAllocate, error) {
	event := new(ScrvUSDUpdateAutoAllocate)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateAutoAllocate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateDefaultQueueIterator is returned from FilterUpdateDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateDefaultQueue events raised by the ScrvUSD contract.
type ScrvUSDUpdateDefaultQueueIterator struct {
	Event *ScrvUSDUpdateDefaultQueue // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateDefaultQueue)
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
		it.Event = new(ScrvUSDUpdateDefaultQueue)
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
func (it *ScrvUSDUpdateDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateDefaultQueue represents a UpdateDefaultQueue event raised by the ScrvUSD contract.
type ScrvUSDUpdateDefaultQueue struct {
	NewDefaultQueue []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultQueue is a free log retrieval operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateDefaultQueue(opts *bind.FilterOpts) (*ScrvUSDUpdateDefaultQueueIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateDefaultQueueIterator{contract: _ScrvUSD.contract, event: "UpdateDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultQueue is a free log subscription operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateDefaultQueue(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateDefaultQueue)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
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

// ParseUpdateDefaultQueue is a log parse operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateDefaultQueue(log types.Log) (*ScrvUSDUpdateDefaultQueue, error) {
	event := new(ScrvUSDUpdateDefaultQueue)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the ScrvUSD contract.
type ScrvUSDUpdateDepositLimitIterator struct {
	Event *ScrvUSDUpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateDepositLimit)
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
		it.Event = new(ScrvUSDUpdateDepositLimit)
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
func (it *ScrvUSDUpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateDepositLimit represents a UpdateDepositLimit event raised by the ScrvUSD contract.
type ScrvUSDUpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*ScrvUSDUpdateDepositLimitIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateDepositLimitIterator{contract: _ScrvUSD.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateDepositLimit)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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

// ParseUpdateDepositLimit is a log parse operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateDepositLimit(log types.Log) (*ScrvUSDUpdateDepositLimit, error) {
	event := new(ScrvUSDUpdateDepositLimit)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateDepositLimitModuleIterator is returned from FilterUpdateDepositLimitModule and is used to iterate over the raw logs and unpacked data for UpdateDepositLimitModule events raised by the ScrvUSD contract.
type ScrvUSDUpdateDepositLimitModuleIterator struct {
	Event *ScrvUSDUpdateDepositLimitModule // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateDepositLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateDepositLimitModule)
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
		it.Event = new(ScrvUSDUpdateDepositLimitModule)
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
func (it *ScrvUSDUpdateDepositLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateDepositLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateDepositLimitModule represents a UpdateDepositLimitModule event raised by the ScrvUSD contract.
type ScrvUSDUpdateDepositLimitModule struct {
	DepositLimitModule common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimitModule is a free log retrieval operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateDepositLimitModule(opts *bind.FilterOpts, deposit_limit_module []common.Address) (*ScrvUSDUpdateDepositLimitModuleIterator, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateDepositLimitModuleIterator{contract: _ScrvUSD.contract, event: "UpdateDepositLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimitModule is a free log subscription operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateDepositLimitModule(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateDepositLimitModule, deposit_limit_module []common.Address) (event.Subscription, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateDepositLimitModule)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
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

// ParseUpdateDepositLimitModule is a log parse operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateDepositLimitModule(log types.Log) (*ScrvUSDUpdateDepositLimitModule, error) {
	event := new(ScrvUSDUpdateDepositLimitModule)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateFutureRoleManagerIterator is returned from FilterUpdateFutureRoleManager and is used to iterate over the raw logs and unpacked data for UpdateFutureRoleManager events raised by the ScrvUSD contract.
type ScrvUSDUpdateFutureRoleManagerIterator struct {
	Event *ScrvUSDUpdateFutureRoleManager // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateFutureRoleManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateFutureRoleManager)
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
		it.Event = new(ScrvUSDUpdateFutureRoleManager)
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
func (it *ScrvUSDUpdateFutureRoleManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateFutureRoleManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateFutureRoleManager represents a UpdateFutureRoleManager event raised by the ScrvUSD contract.
type ScrvUSDUpdateFutureRoleManager struct {
	FutureRoleManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateFutureRoleManager is a free log retrieval operation binding the contract event 0x364d2f5fdafcc5778c92c031ccf715352535731482158c029793bd034198f571.
//
// Solidity: event UpdateFutureRoleManager(address indexed future_role_manager)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateFutureRoleManager(opts *bind.FilterOpts, future_role_manager []common.Address) (*ScrvUSDUpdateFutureRoleManagerIterator, error) {

	var future_role_managerRule []interface{}
	for _, future_role_managerItem := range future_role_manager {
		future_role_managerRule = append(future_role_managerRule, future_role_managerItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateFutureRoleManager", future_role_managerRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateFutureRoleManagerIterator{contract: _ScrvUSD.contract, event: "UpdateFutureRoleManager", logs: logs, sub: sub}, nil
}

// WatchUpdateFutureRoleManager is a free log subscription operation binding the contract event 0x364d2f5fdafcc5778c92c031ccf715352535731482158c029793bd034198f571.
//
// Solidity: event UpdateFutureRoleManager(address indexed future_role_manager)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateFutureRoleManager(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateFutureRoleManager, future_role_manager []common.Address) (event.Subscription, error) {

	var future_role_managerRule []interface{}
	for _, future_role_managerItem := range future_role_manager {
		future_role_managerRule = append(future_role_managerRule, future_role_managerItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateFutureRoleManager", future_role_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateFutureRoleManager)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateFutureRoleManager", log); err != nil {
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

// ParseUpdateFutureRoleManager is a log parse operation binding the contract event 0x364d2f5fdafcc5778c92c031ccf715352535731482158c029793bd034198f571.
//
// Solidity: event UpdateFutureRoleManager(address indexed future_role_manager)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateFutureRoleManager(log types.Log) (*ScrvUSDUpdateFutureRoleManager, error) {
	event := new(ScrvUSDUpdateFutureRoleManager)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateFutureRoleManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateMinimumTotalIdleIterator is returned from FilterUpdateMinimumTotalIdle and is used to iterate over the raw logs and unpacked data for UpdateMinimumTotalIdle events raised by the ScrvUSD contract.
type ScrvUSDUpdateMinimumTotalIdleIterator struct {
	Event *ScrvUSDUpdateMinimumTotalIdle // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateMinimumTotalIdleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateMinimumTotalIdle)
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
		it.Event = new(ScrvUSDUpdateMinimumTotalIdle)
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
func (it *ScrvUSDUpdateMinimumTotalIdleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateMinimumTotalIdleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateMinimumTotalIdle represents a UpdateMinimumTotalIdle event raised by the ScrvUSD contract.
type ScrvUSDUpdateMinimumTotalIdle struct {
	MinimumTotalIdle *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumTotalIdle is a free log retrieval operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateMinimumTotalIdle(opts *bind.FilterOpts) (*ScrvUSDUpdateMinimumTotalIdleIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateMinimumTotalIdleIterator{contract: _ScrvUSD.contract, event: "UpdateMinimumTotalIdle", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumTotalIdle is a free log subscription operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateMinimumTotalIdle(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateMinimumTotalIdle) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateMinimumTotalIdle)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
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

// ParseUpdateMinimumTotalIdle is a log parse operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateMinimumTotalIdle(log types.Log) (*ScrvUSDUpdateMinimumTotalIdle, error) {
	event := new(ScrvUSDUpdateMinimumTotalIdle)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateProfitMaxUnlockTimeIterator is returned from FilterUpdateProfitMaxUnlockTime and is used to iterate over the raw logs and unpacked data for UpdateProfitMaxUnlockTime events raised by the ScrvUSD contract.
type ScrvUSDUpdateProfitMaxUnlockTimeIterator struct {
	Event *ScrvUSDUpdateProfitMaxUnlockTime // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateProfitMaxUnlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateProfitMaxUnlockTime)
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
		it.Event = new(ScrvUSDUpdateProfitMaxUnlockTime)
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
func (it *ScrvUSDUpdateProfitMaxUnlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateProfitMaxUnlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateProfitMaxUnlockTime represents a UpdateProfitMaxUnlockTime event raised by the ScrvUSD contract.
type ScrvUSDUpdateProfitMaxUnlockTime struct {
	ProfitMaxUnlockTime *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfitMaxUnlockTime is a free log retrieval operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateProfitMaxUnlockTime(opts *bind.FilterOpts) (*ScrvUSDUpdateProfitMaxUnlockTimeIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateProfitMaxUnlockTimeIterator{contract: _ScrvUSD.contract, event: "UpdateProfitMaxUnlockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateProfitMaxUnlockTime is a free log subscription operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateProfitMaxUnlockTime(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateProfitMaxUnlockTime) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateProfitMaxUnlockTime)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
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

// ParseUpdateProfitMaxUnlockTime is a log parse operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateProfitMaxUnlockTime(log types.Log) (*ScrvUSDUpdateProfitMaxUnlockTime, error) {
	event := new(ScrvUSDUpdateProfitMaxUnlockTime)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateRoleManagerIterator is returned from FilterUpdateRoleManager and is used to iterate over the raw logs and unpacked data for UpdateRoleManager events raised by the ScrvUSD contract.
type ScrvUSDUpdateRoleManagerIterator struct {
	Event *ScrvUSDUpdateRoleManager // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateRoleManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateRoleManager)
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
		it.Event = new(ScrvUSDUpdateRoleManager)
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
func (it *ScrvUSDUpdateRoleManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateRoleManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateRoleManager represents a UpdateRoleManager event raised by the ScrvUSD contract.
type ScrvUSDUpdateRoleManager struct {
	RoleManager common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRoleManager is a free log retrieval operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateRoleManager(opts *bind.FilterOpts, role_manager []common.Address) (*ScrvUSDUpdateRoleManagerIterator, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateRoleManagerIterator{contract: _ScrvUSD.contract, event: "UpdateRoleManager", logs: logs, sub: sub}, nil
}

// WatchUpdateRoleManager is a free log subscription operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateRoleManager(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateRoleManager, role_manager []common.Address) (event.Subscription, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateRoleManager)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
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

// ParseUpdateRoleManager is a log parse operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateRoleManager(log types.Log) (*ScrvUSDUpdateRoleManager, error) {
	event := new(ScrvUSDUpdateRoleManager)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateUseDefaultQueueIterator is returned from FilterUpdateUseDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultQueue events raised by the ScrvUSD contract.
type ScrvUSDUpdateUseDefaultQueueIterator struct {
	Event *ScrvUSDUpdateUseDefaultQueue // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateUseDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateUseDefaultQueue)
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
		it.Event = new(ScrvUSDUpdateUseDefaultQueue)
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
func (it *ScrvUSDUpdateUseDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateUseDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateUseDefaultQueue represents a UpdateUseDefaultQueue event raised by the ScrvUSD contract.
type ScrvUSDUpdateUseDefaultQueue struct {
	UseDefaultQueue bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultQueue is a free log retrieval operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateUseDefaultQueue(opts *bind.FilterOpts) (*ScrvUSDUpdateUseDefaultQueueIterator, error) {

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateUseDefaultQueueIterator{contract: _ScrvUSD.contract, event: "UpdateUseDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultQueue is a free log subscription operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateUseDefaultQueue(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateUseDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateUseDefaultQueue)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
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

// ParseUpdateUseDefaultQueue is a log parse operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateUseDefaultQueue(log types.Log) (*ScrvUSDUpdateUseDefaultQueue, error) {
	event := new(ScrvUSDUpdateUseDefaultQueue)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdateWithdrawLimitModuleIterator is returned from FilterUpdateWithdrawLimitModule and is used to iterate over the raw logs and unpacked data for UpdateWithdrawLimitModule events raised by the ScrvUSD contract.
type ScrvUSDUpdateWithdrawLimitModuleIterator struct {
	Event *ScrvUSDUpdateWithdrawLimitModule // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdateWithdrawLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdateWithdrawLimitModule)
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
		it.Event = new(ScrvUSDUpdateWithdrawLimitModule)
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
func (it *ScrvUSDUpdateWithdrawLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdateWithdrawLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdateWithdrawLimitModule represents a UpdateWithdrawLimitModule event raised by the ScrvUSD contract.
type ScrvUSDUpdateWithdrawLimitModule struct {
	WithdrawLimitModule common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawLimitModule is a free log retrieval operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdateWithdrawLimitModule(opts *bind.FilterOpts, withdraw_limit_module []common.Address) (*ScrvUSDUpdateWithdrawLimitModuleIterator, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdateWithdrawLimitModuleIterator{contract: _ScrvUSD.contract, event: "UpdateWithdrawLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawLimitModule is a free log subscription operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdateWithdrawLimitModule(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdateWithdrawLimitModule, withdraw_limit_module []common.Address) (event.Subscription, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdateWithdrawLimitModule)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
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

// ParseUpdateWithdrawLimitModule is a log parse operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdateWithdrawLimitModule(log types.Log) (*ScrvUSDUpdateWithdrawLimitModule, error) {
	event := new(ScrvUSDUpdateWithdrawLimitModule)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDUpdatedMaxDebtForStrategyIterator is returned from FilterUpdatedMaxDebtForStrategy and is used to iterate over the raw logs and unpacked data for UpdatedMaxDebtForStrategy events raised by the ScrvUSD contract.
type ScrvUSDUpdatedMaxDebtForStrategyIterator struct {
	Event *ScrvUSDUpdatedMaxDebtForStrategy // Event containing the contract specifics and raw log

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
func (it *ScrvUSDUpdatedMaxDebtForStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDUpdatedMaxDebtForStrategy)
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
		it.Event = new(ScrvUSDUpdatedMaxDebtForStrategy)
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
func (it *ScrvUSDUpdatedMaxDebtForStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDUpdatedMaxDebtForStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDUpdatedMaxDebtForStrategy represents a UpdatedMaxDebtForStrategy event raised by the ScrvUSD contract.
type ScrvUSDUpdatedMaxDebtForStrategy struct {
	Sender   common.Address
	Strategy common.Address
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxDebtForStrategy is a free log retrieval operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) FilterUpdatedMaxDebtForStrategy(opts *bind.FilterOpts, sender []common.Address, strategy []common.Address) (*ScrvUSDUpdatedMaxDebtForStrategyIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDUpdatedMaxDebtForStrategyIterator{contract: _ScrvUSD.contract, event: "UpdatedMaxDebtForStrategy", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxDebtForStrategy is a free log subscription operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) WatchUpdatedMaxDebtForStrategy(opts *bind.WatchOpts, sink chan<- *ScrvUSDUpdatedMaxDebtForStrategy, sender []common.Address, strategy []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDUpdatedMaxDebtForStrategy)
				if err := _ScrvUSD.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
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

// ParseUpdatedMaxDebtForStrategy is a log parse operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_ScrvUSD *ScrvUSDFilterer) ParseUpdatedMaxDebtForStrategy(log types.Log) (*ScrvUSDUpdatedMaxDebtForStrategy, error) {
	event := new(ScrvUSDUpdatedMaxDebtForStrategy)
	if err := _ScrvUSD.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrvUSDWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ScrvUSD contract.
type ScrvUSDWithdrawIterator struct {
	Event *ScrvUSDWithdraw // Event containing the contract specifics and raw log

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
func (it *ScrvUSDWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrvUSDWithdraw)
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
		it.Event = new(ScrvUSDWithdraw)
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
func (it *ScrvUSDWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrvUSDWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrvUSDWithdraw represents a Withdraw event raised by the ScrvUSD contract.
type ScrvUSDWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ScrvUSDWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ScrvUSD.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ScrvUSDWithdrawIterator{contract: _ScrvUSD.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ScrvUSDWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ScrvUSD.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrvUSDWithdraw)
				if err := _ScrvUSD.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ScrvUSD *ScrvUSDFilterer) ParseWithdraw(log types.Log) (*ScrvUSDWithdraw, error) {
	event := new(ScrvUSDWithdraw)
	if err := _ScrvUSD.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
