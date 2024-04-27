// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curvePool

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

// CurvePoolMetaData contains all meta data concerning the CurvePool contract.
var CurvePoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"approx_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"packed_price_scale\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[3]\"},{\"name\":\"_math\",\"type\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\"},{\"name\":\"_salt\",\"type\":\"bytes32\"},{\"name\":\"packed_precisions\",\"type\":\"uint256\"},{\"name\":\"packed_A_gamma\",\"type\":\"uint256\"},{\"name\":\"packed_fee_params\",\"type\":\"uint256\"},{\"name\":\"packed_rebalancing_params\",\"type\":\"uint256\"},{\"name\":\"packed_prices\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_extended\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"cb\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"claim_admin_fees\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_add_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_sub_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_fee\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"precisions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_calc\",\"inputs\":[{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"WETH20\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MATH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"packed_rebalancing_params\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"packed_fee_params\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ADMIN_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
}

// CurvePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolMetaData.ABI instead.
var CurvePoolABI = CurvePoolMetaData.ABI

// CurvePool is an auto generated Go binding around an Ethereum contract.
type CurvePool struct {
	CurvePoolCaller     // Read-only binding to the contract
	CurvePoolTransactor // Write-only binding to the contract
	CurvePoolFilterer   // Log filterer for contract events
}

// CurvePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolSession struct {
	Contract     *CurvePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolCallerSession struct {
	Contract *CurvePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurvePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolTransactorSession struct {
	Contract     *CurvePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurvePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolRaw struct {
	Contract *CurvePool // Generic contract binding to access the raw methods on
}

// CurvePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolCallerRaw struct {
	Contract *CurvePoolCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolTransactorRaw struct {
	Contract *CurvePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePool creates a new instance of CurvePool, bound to a specific deployed contract.
func NewCurvePool(address common.Address, backend bind.ContractBackend) (*CurvePool, error) {
	contract, err := bindCurvePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePool{CurvePoolCaller: CurvePoolCaller{contract: contract}, CurvePoolTransactor: CurvePoolTransactor{contract: contract}, CurvePoolFilterer: CurvePoolFilterer{contract: contract}}, nil
}

// NewCurvePoolCaller creates a new read-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolCaller, error) {
	contract, err := bindCurvePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCaller{contract: contract}, nil
}

// NewCurvePoolTransactor creates a new write-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolTransactor, error) {
	contract, err := bindCurvePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTransactor{contract: contract}, nil
}

// NewCurvePoolFilterer creates a new log filterer instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolFilterer, error) {
	contract, err := bindCurvePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFilterer{contract: contract}, nil
}

// bindCurvePool binds a generic wrapper to an already deployed contract.
func bindCurvePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurvePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.CurvePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolSession) A() (*big.Int, error) {
	return _CurvePool.Contract.A(&_CurvePool.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) A() (*big.Int, error) {
	return _CurvePool.Contract.A(&_CurvePool.CallOpts)
}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_CurvePool *CurvePoolCaller) ADMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "ADMIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_CurvePool *CurvePoolSession) ADMINFEE() (*big.Int, error) {
	return _CurvePool.Contract.ADMINFEE(&_CurvePool.CallOpts)
}

// ADMINFEE is a free data retrieval call binding the contract method 0x4469ed14.
//
// Solidity: function ADMIN_FEE() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) ADMINFEE() (*big.Int, error) {
	return _CurvePool.Contract.ADMINFEE(&_CurvePool.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurvePool *CurvePoolCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurvePool *CurvePoolSession) D() (*big.Int, error) {
	return _CurvePool.Contract.D(&_CurvePool.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) D() (*big.Int, error) {
	return _CurvePool.Contract.D(&_CurvePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePool *CurvePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePool *CurvePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePool.Contract.DOMAINSEPARATOR(&_CurvePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_CurvePool *CurvePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _CurvePool.Contract.DOMAINSEPARATOR(&_CurvePool.CallOpts)
}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_CurvePool *CurvePoolCaller) MATH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "MATH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_CurvePool *CurvePoolSession) MATH() (common.Address, error) {
	return _CurvePool.Contract.MATH(&_CurvePool.CallOpts)
}

// MATH is a free data retrieval call binding the contract method 0xed6c1546.
//
// Solidity: function MATH() view returns(address)
func (_CurvePool *CurvePoolCallerSession) MATH() (common.Address, error) {
	return _CurvePool.Contract.MATH(&_CurvePool.CallOpts)
}

// WETH20 is a free data retrieval call binding the contract method 0x17e26cd1.
//
// Solidity: function WETH20() view returns(address)
func (_CurvePool *CurvePoolCaller) WETH20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "WETH20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH20 is a free data retrieval call binding the contract method 0x17e26cd1.
//
// Solidity: function WETH20() view returns(address)
func (_CurvePool *CurvePoolSession) WETH20() (common.Address, error) {
	return _CurvePool.Contract.WETH20(&_CurvePool.CallOpts)
}

// WETH20 is a free data retrieval call binding the contract method 0x17e26cd1.
//
// Solidity: function WETH20() view returns(address)
func (_CurvePool *CurvePoolCallerSession) WETH20() (common.Address, error) {
	return _CurvePool.Contract.WETH20(&_CurvePool.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurvePool *CurvePoolCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurvePool *CurvePoolSession) AdjustmentStep() (*big.Int, error) {
	return _CurvePool.Contract.AdjustmentStep(&_CurvePool.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AdjustmentStep() (*big.Int, error) {
	return _CurvePool.Contract.AdjustmentStep(&_CurvePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurvePool.Contract.AdminActionsDeadline(&_CurvePool.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurvePool.Contract.AdminActionsDeadline(&_CurvePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePool *CurvePoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePool *CurvePoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.Allowance(&_CurvePool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.Allowance(&_CurvePool.CallOpts, arg0, arg1)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurvePool *CurvePoolCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurvePool *CurvePoolSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurvePool.Contract.AllowedExtraProfit(&_CurvePool.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurvePool.Contract.AllowedExtraProfit(&_CurvePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.BalanceOf(&_CurvePool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.BalanceOf(&_CurvePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenAmount(&_CurvePool.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenAmount(&_CurvePool.CallOpts, amounts, deposit)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolCaller) CalcTokenFee(opts *bind.CallOpts, amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "calc_token_fee", amounts, xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenFee(&_CurvePool.CallOpts, amounts, xp)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcTokenFee(&_CurvePool.CallOpts, amounts, xp)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcWithdrawOneCoin(&_CurvePool.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.CalcWithdrawOneCoin(&_CurvePool.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePool *CurvePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePool *CurvePoolSession) Decimals() (uint8, error) {
	return _CurvePool.Contract.Decimals(&_CurvePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CurvePool *CurvePoolCallerSession) Decimals() (uint8, error) {
	return _CurvePool.Contract.Decimals(&_CurvePool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurvePool *CurvePoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurvePool *CurvePoolSession) Factory() (common.Address, error) {
	return _CurvePool.Contract.Factory(&_CurvePool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CurvePool *CurvePoolCallerSession) Factory() (common.Address, error) {
	return _CurvePool.Contract.Factory(&_CurvePool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) Fee() (*big.Int, error) {
	return _CurvePool.Contract.Fee(&_CurvePool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Fee() (*big.Int, error) {
	return _CurvePool.Contract.Fee(&_CurvePool.CallOpts)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolCaller) FeeCalc(opts *bind.CallOpts, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "fee_calc", xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurvePool.Contract.FeeCalc(&_CurvePool.CallOpts, xp)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurvePool.Contract.FeeCalc(&_CurvePool.CallOpts, xp)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurvePool *CurvePoolSession) FeeGamma() (*big.Int, error) {
	return _CurvePool.Contract.FeeGamma(&_CurvePool.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FeeGamma() (*big.Int, error) {
	return _CurvePool.Contract.FeeGamma(&_CurvePool.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePool *CurvePoolCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePool *CurvePoolSession) FeeReceiver() (common.Address, error) {
	return _CurvePool.Contract.FeeReceiver(&_CurvePool.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurvePool *CurvePoolCallerSession) FeeReceiver() (common.Address, error) {
	return _CurvePool.Contract.FeeReceiver(&_CurvePool.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureAGamma() (*big.Int, error) {
	return _CurvePool.Contract.FutureAGamma(&_CurvePool.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureAGamma() (*big.Int, error) {
	return _CurvePool.Contract.FutureAGamma(&_CurvePool.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolSession) FutureAGammaTime() (*big.Int, error) {
	return _CurvePool.Contract.FutureAGammaTime(&_CurvePool.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _CurvePool.Contract.FutureAGammaTime(&_CurvePool.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurvePool *CurvePoolCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurvePool *CurvePoolSession) Gamma() (*big.Int, error) {
	return _CurvePool.Contract.Gamma(&_CurvePool.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Gamma() (*big.Int, error) {
	return _CurvePool.Contract.Gamma(&_CurvePool.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_CurvePool *CurvePoolSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDx(&_CurvePool.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x37ed3a7a.
//
// Solidity: function get_dx(uint256 i, uint256 j, uint256 dy) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDx(&_CurvePool.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.GetVirtualPrice(&_CurvePool.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.GetVirtualPrice(&_CurvePool.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolSession) InitialAGamma() (*big.Int, error) {
	return _CurvePool.Contract.InitialAGamma(&_CurvePool.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) InitialAGamma() (*big.Int, error) {
	return _CurvePool.Contract.InitialAGamma(&_CurvePool.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolSession) InitialAGammaTime() (*big.Int, error) {
	return _CurvePool.Contract.InitialAGammaTime(&_CurvePool.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _CurvePool.Contract.InitialAGammaTime(&_CurvePool.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCaller) LastPrices(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "last_prices", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.LastPrices(&_CurvePool.CallOpts, k)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.LastPrices(&_CurvePool.CallOpts, k)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurvePool *CurvePoolCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurvePool *CurvePoolSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurvePool.Contract.LastPricesTimestamp(&_CurvePool.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurvePool.Contract.LastPricesTimestamp(&_CurvePool.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_CurvePool *CurvePoolCaller) LpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "lp_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_CurvePool *CurvePoolSession) LpPrice() (*big.Int, error) {
	return _CurvePool.Contract.LpPrice(&_CurvePool.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) LpPrice() (*big.Int, error) {
	return _CurvePool.Contract.LpPrice(&_CurvePool.CallOpts)
}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_CurvePool *CurvePoolCaller) MaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_CurvePool *CurvePoolSession) MaTime() (*big.Int, error) {
	return _CurvePool.Contract.MaTime(&_CurvePool.CallOpts)
}

// MaTime is a free data retrieval call binding the contract method 0x09c3da6a.
//
// Solidity: function ma_time() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) MaTime() (*big.Int, error) {
	return _CurvePool.Contract.MaTime(&_CurvePool.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) MidFee() (*big.Int, error) {
	return _CurvePool.Contract.MidFee(&_CurvePool.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) MidFee() (*big.Int, error) {
	return _CurvePool.Contract.MidFee(&_CurvePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePool *CurvePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePool *CurvePoolSession) Name() (string, error) {
	return _CurvePool.Contract.Name(&_CurvePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurvePool *CurvePoolCallerSession) Name() (string, error) {
	return _CurvePool.Contract.Name(&_CurvePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.Nonces(&_CurvePool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _CurvePool.Contract.Nonces(&_CurvePool.CallOpts, arg0)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurvePool *CurvePoolCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurvePool *CurvePoolSession) OutFee() (*big.Int, error) {
	return _CurvePool.Contract.OutFee(&_CurvePool.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) OutFee() (*big.Int, error) {
	return _CurvePool.Contract.OutFee(&_CurvePool.CallOpts)
}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_CurvePool *CurvePoolCaller) PackedFeeParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "packed_fee_params")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_CurvePool *CurvePoolSession) PackedFeeParams() (*big.Int, error) {
	return _CurvePool.Contract.PackedFeeParams(&_CurvePool.CallOpts)
}

// PackedFeeParams is a free data retrieval call binding the contract method 0xe3616405.
//
// Solidity: function packed_fee_params() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) PackedFeeParams() (*big.Int, error) {
	return _CurvePool.Contract.PackedFeeParams(&_CurvePool.CallOpts)
}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_CurvePool *CurvePoolCaller) PackedRebalancingParams(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "packed_rebalancing_params")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_CurvePool *CurvePoolSession) PackedRebalancingParams() (*big.Int, error) {
	return _CurvePool.Contract.PackedRebalancingParams(&_CurvePool.CallOpts)
}

// PackedRebalancingParams is a free data retrieval call binding the contract method 0x3dd65478.
//
// Solidity: function packed_rebalancing_params() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) PackedRebalancingParams() (*big.Int, error) {
	return _CurvePool.Contract.PackedRebalancingParams(&_CurvePool.CallOpts)
}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[3])
func (_CurvePool *CurvePoolCaller) Precisions(opts *bind.CallOpts) ([3]*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "precisions")

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[3])
func (_CurvePool *CurvePoolSession) Precisions() ([3]*big.Int, error) {
	return _CurvePool.Contract.Precisions(&_CurvePool.CallOpts)
}

// Precisions is a free data retrieval call binding the contract method 0x3620604b.
//
// Solidity: function precisions() view returns(uint256[3])
func (_CurvePool *CurvePoolCallerSession) Precisions() ([3]*big.Int, error) {
	return _CurvePool.Contract.Precisions(&_CurvePool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCaller) PriceOracle(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "price_oracle", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.PriceOracle(&_CurvePool.CallOpts, k)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.PriceOracle(&_CurvePool.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCaller) PriceScale(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "price_scale", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.PriceScale(&_CurvePool.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.PriceScale(&_CurvePool.CallOpts, k)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurvePool *CurvePoolCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurvePool *CurvePoolSession) Salt() ([32]byte, error) {
	return _CurvePool.Contract.Salt(&_CurvePool.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_CurvePool *CurvePoolCallerSession) Salt() ([32]byte, error) {
	return _CurvePool.Contract.Salt(&_CurvePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePool *CurvePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePool *CurvePoolSession) Symbol() (string, error) {
	return _CurvePool.Contract.Symbol(&_CurvePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurvePool *CurvePoolCallerSession) Symbol() (string, error) {
	return _CurvePool.Contract.Symbol(&_CurvePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePool *CurvePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePool *CurvePoolSession) TotalSupply() (*big.Int, error) {
	return _CurvePool.Contract.TotalSupply(&_CurvePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _CurvePool.Contract.TotalSupply(&_CurvePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePool *CurvePoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePool *CurvePoolSession) Version() (string, error) {
	return _CurvePool.Contract.Version(&_CurvePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurvePool *CurvePoolCallerSession) Version() (string, error) {
	return _CurvePool.Contract.Version(&_CurvePool.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolSession) VirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.VirtualPrice(&_CurvePool.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) VirtualPrice() (*big.Int, error) {
	return _CurvePool.Contract.VirtualPrice(&_CurvePool.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurvePool *CurvePoolCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurvePool *CurvePoolSession) XcpProfit() (*big.Int, error) {
	return _CurvePool.Contract.XcpProfit(&_CurvePool.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) XcpProfit() (*big.Int, error) {
	return _CurvePool.Contract.XcpProfit(&_CurvePool.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurvePool *CurvePoolCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurvePool *CurvePoolSession) XcpProfitA() (*big.Int, error) {
	return _CurvePool.Contract.XcpProfitA(&_CurvePool.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) XcpProfitA() (*big.Int, error) {
	return _CurvePool.Contract.XcpProfitA(&_CurvePool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurvePool *CurvePoolSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity(&_CurvePool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity(&_CurvePool.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolSession) AddLiquidity0(amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity0(&_CurvePool.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x2b6e993a.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) AddLiquidity0(amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity0(&_CurvePool.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x5cecb5f7.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) AddLiquidity1(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "add_liquidity1", amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x5cecb5f7.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolSession) AddLiquidity1(amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity1(&_CurvePool.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x5cecb5f7.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) AddLiquidity1(amounts [3]*big.Int, min_mint_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.AddLiquidity1(&_CurvePool.TransactOpts, amounts, min_mint_amount, use_eth, receiver)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurvePool *CurvePoolTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurvePool *CurvePoolSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyNewParameters(&_CurvePool.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurvePool *CurvePoolTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.ApplyNewParameters(&_CurvePool.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Approve(&_CurvePool.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Approve(&_CurvePool.TransactOpts, _spender, _value)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurvePool *CurvePoolTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurvePool *CurvePoolSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.ClaimAdminFees(&_CurvePool.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurvePool *CurvePoolTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurvePool.Contract.ClaimAdminFees(&_CurvePool.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x4711a4f8.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_CurvePool *CurvePoolTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x4711a4f8.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_CurvePool *CurvePoolSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitNewParameters(&_CurvePool.TransactOpts, _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0x4711a4f8.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_time) returns()
func (_CurvePool *CurvePoolTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.CommitNewParameters(&_CurvePool.TransactOpts, _new_mid_fee, _new_out_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_time)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_CurvePool *CurvePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "decreaseAllowance", _spender, _sub_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_CurvePool *CurvePoolSession) DecreaseAllowance(_spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.DecreaseAllowance(&_CurvePool.TransactOpts, _spender, _sub_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _sub_value) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) DecreaseAllowance(_spender common.Address, _sub_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.DecreaseAllowance(&_CurvePool.TransactOpts, _spender, _sub_value)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange0(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange0(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) Exchange1(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange1", i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolSession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange1(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// Exchange1 is a paid mutator transaction binding the contract method 0xce7d6503.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) Exchange1(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.Exchange1(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth, receiver)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) returns(uint256)
func (_CurvePool *CurvePoolTransactor) ExchangeExtended(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange_extended", i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) returns(uint256)
func (_CurvePool *CurvePoolSession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeExtended(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeExtended is a paid mutator transaction binding the contract method 0xdd96994f.
//
// Solidity: function exchange_extended(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth, address sender, address receiver, bytes32 cb) returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) ExchangeExtended(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool, sender common.Address, receiver common.Address, cb [32]byte) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeExtended(&_CurvePool.TransactOpts, i, j, dx, min_dy, use_eth, sender, receiver, cb)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeUnderlying(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeUnderlying(&_CurvePool.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactor) ExchangeUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "exchange_underlying0", i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolSession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeUnderlying0(&_CurvePool.TransactOpts, i, j, dx, min_dy, receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy, address receiver) payable returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) ExchangeUnderlying0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.ExchangeUnderlying0(&_CurvePool.TransactOpts, i, j, dx, min_dy, receiver)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_CurvePool *CurvePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "increaseAllowance", _spender, _add_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_CurvePool *CurvePoolSession) IncreaseAllowance(_spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.IncreaseAllowance(&_CurvePool.TransactOpts, _spender, _add_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _add_value) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) IncreaseAllowance(_spender common.Address, _add_value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.IncreaseAllowance(&_CurvePool.TransactOpts, _spender, _add_value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePool *CurvePoolTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePool *CurvePoolSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePool.Contract.Permit(&_CurvePool.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _CurvePool.Contract.Permit(&_CurvePool.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurvePool *CurvePoolTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurvePool *CurvePoolSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RampAGamma(&_CurvePool.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurvePool *CurvePoolTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RampAGamma(&_CurvePool.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns(uint256[3])
func (_CurvePool *CurvePoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns(uint256[3])
func (_CurvePool *CurvePoolSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity(&_CurvePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns(uint256[3])
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity(&_CurvePool.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth) returns(uint256[3])
func (_CurvePool *CurvePoolTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity0", _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth) returns(uint256[3])
func (_CurvePool *CurvePoolSession) RemoveLiquidity0(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity0(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xfce64736.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth) returns(uint256[3])
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidity0(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity0(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1da3d238.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver) returns(uint256[3])
func (_CurvePool *CurvePoolTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity1", _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1da3d238.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver) returns(uint256[3])
func (_CurvePool *CurvePoolSession) RemoveLiquidity1(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity1(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x1da3d238.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver) returns(uint256[3])
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidity1(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity1(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth, receiver)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5cd34780.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver, bool claim_admin_fees) returns(uint256[3])
func (_CurvePool *CurvePoolTransactor) RemoveLiquidity2(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address, claim_admin_fees bool) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity2", _amount, min_amounts, use_eth, receiver, claim_admin_fees)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5cd34780.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver, bool claim_admin_fees) returns(uint256[3])
func (_CurvePool *CurvePoolSession) RemoveLiquidity2(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address, claim_admin_fees bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity2(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth, receiver, claim_admin_fees)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5cd34780.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts, bool use_eth, address receiver, bool claim_admin_fees) returns(uint256[3])
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidity2(_amount *big.Int, min_amounts [3]*big.Int, use_eth bool, receiver common.Address, claim_admin_fees bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidity2(&_CurvePool.TransactOpts, _amount, min_amounts, use_eth, receiver, claim_admin_fees)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_CurvePool *CurvePoolTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_CurvePool *CurvePoolSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin(&_CurvePool.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin(&_CurvePool.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_CurvePool *CurvePoolTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity_one_coin0", token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_CurvePool *CurvePoolSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin0(&_CurvePool.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin0(&_CurvePool.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_CurvePool *CurvePoolTransactor) RemoveLiquidityOneCoin1(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "remove_liquidity_one_coin1", token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_CurvePool *CurvePoolSession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin1(&_CurvePool.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x07329bcd.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth, address receiver) returns(uint256)
func (_CurvePool *CurvePoolTransactorSession) RemoveLiquidityOneCoin1(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool, receiver common.Address) (*types.Transaction, error) {
	return _CurvePool.Contract.RemoveLiquidityOneCoin1(&_CurvePool.TransactOpts, token_amount, i, min_amount, use_eth, receiver)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertNewParameters(&_CurvePool.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurvePool *CurvePoolTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurvePool.Contract.RevertNewParameters(&_CurvePool.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurvePool *CurvePoolTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurvePool *CurvePoolSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurvePool.Contract.StopRampAGamma(&_CurvePool.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurvePool *CurvePoolTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurvePool.Contract.StopRampAGamma(&_CurvePool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Transfer(&_CurvePool.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.Transfer(&_CurvePool.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.TransferFrom(&_CurvePool.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurvePool *CurvePoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurvePool.Contract.TransferFrom(&_CurvePool.TransactOpts, _from, _to, _value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurvePool *CurvePoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CurvePool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurvePool *CurvePoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurvePool.Contract.Fallback(&_CurvePool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurvePool *CurvePoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurvePool.Contract.Fallback(&_CurvePool.TransactOpts, calldata)
}

// CurvePoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurvePool contract.
type CurvePoolAddLiquidityIterator struct {
	Event *CurvePoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolAddLiquidity)
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
		it.Event = new(CurvePoolAddLiquidity)
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
func (it *CurvePoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolAddLiquidity represents a AddLiquidity event raised by the CurvePool contract.
type CurvePoolAddLiquidity struct {
	Provider         common.Address
	TokenAmounts     [3]*big.Int
	Fee              *big.Int
	TokenSupply      *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xe1b60455bd9e33720b547f60e4e0cfbf1252d0f2ee0147d53029945f39fe3c1a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolAddLiquidityIterator{contract: _CurvePool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xe1b60455bd9e33720b547f60e4e0cfbf1252d0f2ee0147d53029945f39fe3c1a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolAddLiquidity)
				if err := _CurvePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xe1b60455bd9e33720b547f60e4e0cfbf1252d0f2ee0147d53029945f39fe3c1a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) ParseAddLiquidity(log types.Log) (*CurvePoolAddLiquidity, error) {
	event := new(CurvePoolAddLiquidity)
	if err := _CurvePool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurvePool contract.
type CurvePoolApprovalIterator struct {
	Event *CurvePoolApproval // Event containing the contract specifics and raw log

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
func (it *CurvePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolApproval)
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
		it.Event = new(CurvePoolApproval)
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
func (it *CurvePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolApproval represents a Approval event raised by the CurvePool contract.
type CurvePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePool *CurvePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CurvePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolApprovalIterator{contract: _CurvePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CurvePool *CurvePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurvePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolApproval)
				if err := _CurvePool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CurvePool *CurvePoolFilterer) ParseApproval(log types.Log) (*CurvePoolApproval, error) {
	event := new(CurvePoolApproval)
	if err := _CurvePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the CurvePool contract.
type CurvePoolClaimAdminFeeIterator struct {
	Event *CurvePoolClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *CurvePoolClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolClaimAdminFee)
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
		it.Event = new(CurvePoolClaimAdminFee)
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
func (it *CurvePoolClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolClaimAdminFee represents a ClaimAdminFee event raised by the CurvePool contract.
type CurvePoolClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurvePool *CurvePoolFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*CurvePoolClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolClaimAdminFeeIterator{contract: _CurvePool.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurvePool *CurvePoolFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *CurvePoolClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolClaimAdminFee)
				if err := _CurvePool.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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

// ParseClaimAdminFee is a log parse operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurvePool *CurvePoolFilterer) ParseClaimAdminFee(log types.Log) (*CurvePoolClaimAdminFee, error) {
	event := new(CurvePoolClaimAdminFee)
	if err := _CurvePool.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the CurvePool contract.
type CurvePoolCommitNewParametersIterator struct {
	Event *CurvePoolCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvePoolCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolCommitNewParameters)
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
		it.Event = new(CurvePoolCommitNewParameters)
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
func (it *CurvePoolCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolCommitNewParameters represents a CommitNewParameters event raised by the CurvePool contract.
type CurvePoolCommitNewParameters struct {
	Deadline           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaTime             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0xec36b92a482408f90e07357ca20c8cfaca85affe765903cb242e377fafb166af.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*CurvePoolCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCommitNewParametersIterator{contract: _CurvePool.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0xec36b92a482408f90e07357ca20c8cfaca85affe765903cb242e377fafb166af.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *CurvePoolCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolCommitNewParameters)
				if err := _CurvePool.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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

// ParseCommitNewParameters is a log parse operation binding the contract event 0xec36b92a482408f90e07357ca20c8cfaca85affe765903cb242e377fafb166af.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) ParseCommitNewParameters(log types.Log) (*CurvePoolCommitNewParameters, error) {
	event := new(CurvePoolCommitNewParameters)
	if err := _CurvePool.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the CurvePool contract.
type CurvePoolNewParametersIterator struct {
	Event *CurvePoolNewParameters // Event containing the contract specifics and raw log

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
func (it *CurvePoolNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolNewParameters)
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
		it.Event = new(CurvePoolNewParameters)
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
func (it *CurvePoolNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolNewParameters represents a NewParameters event raised by the CurvePool contract.
type CurvePoolNewParameters struct {
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaTime             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) FilterNewParameters(opts *bind.FilterOpts) (*CurvePoolNewParametersIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &CurvePoolNewParametersIterator{contract: _CurvePool.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *CurvePoolNewParameters) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolNewParameters)
				if err := _CurvePool.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0xa32137411fc7c20db359079cd84af0e2cad58cd7a182a8a5e23e08e554e88bf0.
//
// Solidity: event NewParameters(uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_time)
func (_CurvePool *CurvePoolFilterer) ParseNewParameters(log types.Log) (*CurvePoolNewParameters, error) {
	event := new(CurvePoolNewParameters)
	if err := _CurvePool.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the CurvePool contract.
type CurvePoolRampAgammaIterator struct {
	Event *CurvePoolRampAgamma // Event containing the contract specifics and raw log

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
func (it *CurvePoolRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRampAgamma)
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
		it.Event = new(CurvePoolRampAgamma)
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
func (it *CurvePoolRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRampAgamma represents a RampAgamma event raised by the CurvePool contract.
type CurvePoolRampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*CurvePoolRampAgammaIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &CurvePoolRampAgammaIterator{contract: _CurvePool.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *CurvePoolRampAgamma) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRampAgamma)
				if err := _CurvePool.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurvePool *CurvePoolFilterer) ParseRampAgamma(log types.Log) (*CurvePoolRampAgamma, error) {
	event := new(CurvePoolRampAgamma)
	if err := _CurvePool.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurvePool contract.
type CurvePoolRemoveLiquidityIterator struct {
	Event *CurvePoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoveLiquidity)
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
		it.Event = new(CurvePoolRemoveLiquidity)
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
func (it *CurvePoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoveLiquidity represents a RemoveLiquidity event raised by the CurvePool contract.
type CurvePoolRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemoveLiquidityIterator{contract: _CurvePool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoveLiquidity)
				if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurvePool *CurvePoolFilterer) ParseRemoveLiquidity(log types.Log) (*CurvePoolRemoveLiquidity, error) {
	event := new(CurvePoolRemoveLiquidity)
	if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurvePool contract.
type CurvePoolRemoveLiquidityOneIterator struct {
	Event *CurvePoolRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoveLiquidityOne)
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
		it.Event = new(CurvePoolRemoveLiquidityOne)
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
func (it *CurvePoolRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurvePool contract.
type CurvePoolRemoveLiquidityOne struct {
	Provider         common.Address
	TokenAmount      *big.Int
	CoinIndex        *big.Int
	CoinAmount       *big.Int
	ApproxFee        *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurvePoolRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemoveLiquidityOneIterator{contract: _CurvePool.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoveLiquidityOne)
				if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0xe200e24d4a4c7cd367dd9befe394dc8a14e6d58c88ff5e2f512d65a9e0aa9c5c.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount, uint256 approx_fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurvePoolRemoveLiquidityOne, error) {
	event := new(CurvePoolRemoveLiquidityOne)
	if err := _CurvePool.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurvePool contract.
type CurvePoolStopRampAIterator struct {
	Event *CurvePoolStopRampA // Event containing the contract specifics and raw log

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
func (it *CurvePoolStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolStopRampA)
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
		it.Event = new(CurvePoolStopRampA)
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
func (it *CurvePoolStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolStopRampA represents a StopRampA event raised by the CurvePool contract.
type CurvePoolStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurvePool *CurvePoolFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurvePoolStopRampAIterator, error) {

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurvePoolStopRampAIterator{contract: _CurvePool.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurvePool *CurvePoolFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurvePoolStopRampA) (event.Subscription, error) {

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolStopRampA)
				if err := _CurvePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurvePool *CurvePoolFilterer) ParseStopRampA(log types.Log) (*CurvePoolStopRampA, error) {
	event := new(CurvePoolStopRampA)
	if err := _CurvePool.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurvePool contract.
type CurvePoolTokenExchangeIterator struct {
	Event *CurvePoolTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurvePoolTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolTokenExchange)
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
		it.Event = new(CurvePoolTokenExchange)
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
func (it *CurvePoolTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolTokenExchange represents a TokenExchange event raised by the CurvePool contract.
type CurvePoolTokenExchange struct {
	Buyer            common.Address
	SoldId           *big.Int
	TokensSold       *big.Int
	BoughtId         *big.Int
	TokensBought     *big.Int
	Fee              *big.Int
	PackedPriceScale *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurvePoolTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTokenExchangeIterator{contract: _CurvePool.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurvePoolTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolTokenExchange)
				if err := _CurvePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought, uint256 fee, uint256 packed_price_scale)
func (_CurvePool *CurvePoolFilterer) ParseTokenExchange(log types.Log) (*CurvePoolTokenExchange, error) {
	event := new(CurvePoolTokenExchange)
	if err := _CurvePool.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurvePool contract.
type CurvePoolTransferIterator struct {
	Event *CurvePoolTransfer // Event containing the contract specifics and raw log

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
func (it *CurvePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolTransfer)
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
		it.Event = new(CurvePoolTransfer)
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
func (it *CurvePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolTransfer represents a Transfer event raised by the CurvePool contract.
type CurvePoolTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePool *CurvePoolFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*CurvePoolTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePool.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTransferIterator{contract: _CurvePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_CurvePool *CurvePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurvePoolTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CurvePool.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolTransfer)
				if err := _CurvePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CurvePool *CurvePoolFilterer) ParseTransfer(log types.Log) (*CurvePoolTransfer, error) {
	event := new(CurvePoolTransfer)
	if err := _CurvePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
