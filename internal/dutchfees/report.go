package dutchfees

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const weekSeconds uint64 = 7 * 24 * 60 * 60

type chainConfig struct {
	Name                                                                  string
	ID                                                                    uint64
	Collector, Burner, LegacyBurner, Target, Allocator, Bridger, Treasury common.Address
}

var address = common.HexToAddress
var treasury = address("0x6508ef65b0bd57eabd0f1d52685a70433b2d290b")
var networks = []chainConfig{
	{Name: "ethereum", LegacyBurner: address("0xc0fc3ddfec95ca45a0d2393f518d3ea1ccf44f8b"), ID: 1, Collector: address("0xa2bcd1a4efbd04b63cd03f5aff2561106ebcce00"), Burner: address("0xc014a92c76709d6c7f7c7841823b7164d6d2abd4"), Target: address("0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E"), Allocator: address("0x22530D384cd9915e096ead2Db7F82EE81f8eB468")},
	{Name: "arbitrum", LegacyBurner: address("0xd48503022b2fbba4adf5a78f56b988be57d84c91"), ID: 42161, Collector: address("0x68e7f1e30e0bc195b19c0441cdb34c0ed513f8dd"), Burner: address("0x66a81a7f703c8306b425c397e8e85965cdf71698"), Target: address("0x498Bf2B1e120FeD3ad3D42EA2165E9b73f99C1e5"), Bridger: address("0x8A5a5299f35614Ac558AA290C2d5856EDeC1B5Ad"), Treasury: treasury},
	{Name: "base", LegacyBurner: address("0x9c0f9f476b50cf07e704c070366c1adbf39f264e"), ID: 8453, Collector: address("0x56ba5c600db7524c5ec253567f8821aaadfcdbea"), Burner: address("0x275cabb174d5962b5e532e49bcc7e546da570a31"), Target: address("0x417Ac0e078398C154EdFadD9Ef675d30Be60Af93"), Bridger: address("0x2Fb9ef510f1c52fdfDf84dfa8EEa72E431B7651D"), Treasury: treasury},
}

type pending struct {
	CollectorCrvUSD    string   `json:"collectorCrvUSD,omitempty"`
	BurnerCrvUSD       string   `json:"installedBurnerCrvUSD,omitempty"`
	HookerCrvUSD       string   `json:"hookerCrvUSD,omitempty"`
	AllocatorCrvUSD    string   `json:"allocatorCrvUSD,omitempty"`
	BufferCrvUSD       string   `json:"reservedBufferCrvUSD,omitempty"`
	CashCrvUSD         string   `json:"cashBeforeBufferAndKeeperCrvUSD,omitempty"`
	RemainingMarketUSD float64  `json:"pricedRemainingToConvertUSD"`
	TargetPriceUSD     *float64 `json:"crvUSDMarketPriceUSD,omitempty"`
	CurrentLotQuote    string   `json:"fullRemainingAuctionQuoteCrvUSD,omitempty"`
	EstimateCrvUSD     *float64 `json:"estimatedPendingCrvUSD"`
	VeEstimate         *float64 `json:"estimatedPendingVeCRVCrvUSD,omitempty"`
	BridgeMinimum      string   `json:"bridgeMinimumCrvUSD,omitempty"`
	BridgeMinimumMet   *bool    `json:"estimatedBridgeMinimumMet,omitempty"`
	WeightBps          *uint64  `json:"veCRVWeightBps,omitempty"`
	Partial            bool     `json:"partial"`
}

type chainReport struct {
	Chain            string             `json:"chain"`
	Block            snapshot           `json:"block"`
	Phase            string             `json:"phase,omitempty"`
	Status           string             `json:"burnerStatus"`
	ProposalBurner   common.Address     `json:"proposalBurner"`
	InstalledBurner  common.Address     `json:"installedBurner"`
	RouteVerified    bool               `json:"routeVerified,omitempty"`
	GlobalMask       *uint64            `json:"globalKillMask,omitempty"`
	TargetMask       *uint64            `json:"targetKillMask,omitempty"`
	Pending          *pending           `json:"pending,omitempty"`
	LegacyComponents map[string]float64 `json:"legacyComponents,omitempty"`
	WeekFlows        *flows             `json:"weekFlows,omitempty"`
	StageSource      string             `json:"stageSource,omitempty"`
	InventoryScope   string             `json:"inventoryScope,omitempty"`
	Inventory        *stock             `json:"inventory,omitempty"`
	Pools            *poolEstimate      `json:"pools,omitempty"`
	RPCRequests      uint64             `json:"rpcRequests"`
}

func beforeCollectCloses(p protocol) bool { return p.Phase == "sleep" || p.Phase == "collect" }

func units(n *big.Int) float64 {
	v, _ := new(big.Rat).SetFrac(n, big.NewInt(1e18)).Float64()
	return v
}

func makePending(cfg chainConfig, p protocol, s stock, pools *poolEstimate, targetPrice float64) pending {
	cash := new(big.Int).Add(s.collectorCash, s.burnerCash)
	allCash := new(big.Int).Add(cash, p.HookCash)
	allCash.Add(allCash, p.AllocatorCash)
	out := pending{CollectorCrvUSD: amountText(s.collectorCash, 18), BurnerCrvUSD: amountText(s.burnerCash, 18), HookerCrvUSD: amountText(p.HookCash, 18), BufferCrvUSD: amountText(p.CollectorBuffer, 18), CashCrvUSD: amountText(allCash, 18)}
	if targetPrice > 0 {
		out.TargetPriceUSD = &targetPrice
	}
	if cfg.Allocator != (common.Address{}) {
		out.AllocatorCrvUSD = amountText(p.AllocatorCash, 18)
		out.WeightBps = &p.Weight
	} else if p.RouteVerified {
		out.BridgeMinimum = amountText(p.BridgeMinimum, 18)
	}
	if p.Installed != cfg.Burner || !p.RouteVerified || (p.GlobalMask|p.TargetMask)&8 != 0 {
		out.Partial = true
		return out
	}
	// Existing crvUSD can still be forwarded when collection or trading is paused.
	if beforeCollectCloses(p) {
		canStage := p.GlobalMask&6 == 0
		if canStage {
			out.RemainingMarketUSD = s.EligibleUSD
			out.Partial = !s.EligibleComplete
		} else if p.GlobalMask&4 == 0 {
			out.RemainingMarketUSD = s.StagedRemainingUSD
			out.Partial = !s.StagedRemainingComplete
		}
		if pools == nil {
			out.Partial = true
		} else if canStage {
			out.RemainingMarketUSD += pools.EligibleNonTargetUSD
		}
	} else if p.Phase == "exchange" && p.GlobalMask&4 == 0 {
		out.RemainingMarketUSD = s.StagedRemainingUSD
		out.Partial = !s.StagedRemainingComplete
	}

	// crvUSD received directly from a pool needs no auction. It can still be
	// claimed during EXCHANGE or FORWARD if that pool survey was requested.
	if pools != nil {
		cash.Add(cash, pools.TargetRaw)
		out.Partial = out.Partial || pools.Partial
	}
	if out.RemainingMarketUSD > 0 && targetPrice <= 0 {
		out.Partial = true
		return out
	}
	potential := units(cash)
	if out.RemainingMarketUSD > 0 {
		potential += out.RemainingMarketUSD / targetPrice
	}
	forwardable := max(0, potential-units(p.CollectorBuffer))
	// Before FORWARD the eventual keeper timestamp is unknown. During FORWARD use its rate at the snapshot.
	if p.Phase == "forward" {
		forwardable *= 1 - units(p.ForwardRate)
	}
	total := forwardable + units(p.HookCash) + units(p.AllocatorCash)
	out.EstimateCrvUSD = &total
	if cfg.Allocator != (common.Address{}) {
		share := total * float64(p.Weight) / 10000
		out.VeEstimate = &share
	} else {
		meets := total >= units(p.BridgeMinimum)
		out.BridgeMinimumMet = &meets
	}
	return out
}
