package dutchfees

import (
	"context"
	"errors"
	"maps"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"main/interfaces"
)

// These receipts are incremental transfers into the collector, measured at the
// same Ethereum block as its inventory. They exclude balances already there.
type ethereumDashboard struct {
	protocol         protocol
	price            float64
	controllerCrvUSD float64
	pegKeeperUSD     float64
	partial          bool
	beforeBuffer     float64
}

func readDashboardReceipts(ctx context.Context, r stockReader, cfg chainConfig, head snapshot, p protocol, prices map[common.Address]tokenMeta) (*ethereumDashboard, error) {
	price := prices[cfg.Target].PriceUSD
	if !p.RouteVerified || price <= 0 || !finite(price) {
		return nil, errors.New("cannot value the Ethereum veCRV route")
	}
	if p.Installed == cfg.LegacyBurner && (p.GlobalMask|p.TargetMask)&8 != 0 {
		return &ethereumDashboard{protocol: p, price: price, partial: true}, nil
	}
	block := new(big.Int).SetUint64(head.Number)
	controller, partial, err := controllerReceipts(ctx, r, cfg, block)
	if err != nil {
		return nil, err
	}
	fees := &ethereumDashboard{protocol: p, price: price, controllerCrvUSD: units(controller), partial: partial}
	// PegKeeper profit is LP inventory. A fresh LP receipt must still be
	// collectible and exchangeable before it can enter this week's proceeds.
	if beforeCollectCloses(p) && p.GlobalMask&6 == 0 {
		profit, incomplete, err := pegKeeperReceiptsUSD(ctx, r, cfg, block, prices)
		if err != nil {
			return nil, err
		}
		fees.pegKeeperUSD = profit
		fees.partial = fees.partial || incomplete
	}
	return fees, nil
}

type Distribution struct {
	Migrated          map[string]float64
	EthereumBlockHash common.Hash
	ethereum          *ethereumDashboard
	partial           bool
}

// Dashboard retains the historical per-network calculation until each burner's
// migration. All Ethereum components are reported in USD attributable to veCRV.
func (d Distribution) Dashboard(earned map[string]float64, missingNetwork bool) interfaces.VeCrvFees {
	components := maps.Clone(earned)
	eth := d.ethereum
	weight := float64(eth.protocol.Weight) / 10000
	if _, installed := d.Migrated["ethereum"]; installed {
		// The Ethereum entry already contains the veCRV share of downstream balances.
		// Apply the buffer and keeper rate to the combined pipeline just once.
		before := eth.beforeBuffer
		withControllers := before + eth.controllerCrvUSD
		withKeepers := withControllers + eth.pegKeeperUSD/eth.price
		toUSD := weight * eth.price
		components["crvusd"] = (eth.protocol.forwardable(withControllers) - eth.protocol.forwardable(before)) * toUSD
		components["pegKeepers"] = (eth.protocol.forwardable(withKeepers) - eth.protocol.forwardable(withControllers)) * toUSD
	} else if (eth.protocol.GlobalMask|eth.protocol.TargetMask)&8 != 0 {
		// Exclude unavailable legacy projections from the veCRV snapshot only.
		delete(components, "ethereum")
		delete(components, "crvusd")
		return interfaces.VeCrvFees{Components: components, Partial: true}
	} else {
		components["ethereum"] *= weight
		components["crvusd"] = (components["crvusd"] + eth.controllerCrvUSD) * eth.price * weight
		components["pegKeepers"] = eth.pegKeeperUSD * weight
	}

	return interfaces.VeCrvFees{Components: components, Partial: missingNetwork || d.partial || eth.partial}
}
