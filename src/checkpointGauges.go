package src

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"log"
	"main/contracts/curveGC"
	"main/contracts/multicall"
	"main/utils"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const batchSize = 100

var (
	chainsSupported = []string{"arbitrum", "base", "optimism", "fraxtal", "sonic", "taiko"}

	// Gauge minimal ABI
	gaugeABIJSON = `[{
		"type":"function",
		"name":"user_checkpoint",
		"stateMutability":"nonpayable",
		"inputs":[{"name":"user","type":"address"}],
		"outputs":[]
	}]`

	// Multicall3 ABI (only aggregate3 needed)
	multicallABIJSON = `[
		{"inputs":[
			{"components":[
				{"internalType":"address","name":"target","type":"address"},
				{"internalType":"bool","name":"allowFailure","type":"bool"},
				{"internalType":"bytes","name":"callData","type":"bytes"}],
			"internalType":"struct Multicall3.Call3[]","name":"calls","type":"tuple[]"}],
		"name":"aggregate3",
		"outputs":[{"components":[
			{"internalType":"bool","name":"success","type":"bool"},
			{"internalType":"bytes","name":"returnData","type":"bytes"}],
		"internalType":"struct Multicall3.Result[]","name":"returnData","type":"tuple[]"}],
		"stateMutability":"payable","type":"function"}]`

	dryRun = true
)

func mustABI(jsonStr string) abi.ABI {
	a, err := abi.JSON(strings.NewReader(jsonStr))
	if err != nil {
		log.Fatalf("ABI parse error: %v", err)
	}
	return a
}

func loadKey() (*ecdsa.PrivateKey, common.Address, error) {
	pk := os.Getenv("PRIVATE_KEY")
	if pk == "" {
		return nil, common.Address{}, errors.New("PRIVATE_KEY missing")
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(pk, "0x"))
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("invalid PRIVATE_KEY: %w", err)
	}
	return key, crypto.PubkeyToAddress(key.PublicKey), nil
}

func chunk[T any](in []T, n int) [][]T {
	if n <= 0 {
		return [][]T{in}
	}
	var out [][]T
	for i := 0; i < len(in); i += n {
		j := i + n
		if j > len(in) {
			j = len(in)
		}
		out = append(out, in[i:j])
	}
	return out
}

func CheckpointGauges() {

	key, from, err := loadKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sender: %s", from.Hex())

	defaultDryRun := os.Getenv("GITHUB_ACTIONS") == ""
	flag.BoolVar(&dryRun, "dry-run", defaultDryRun, "If true, do not broadcast")

	gauges := utils.GetAllGauges()

	// Fetch the gauge inflation onchain
	rpcMainnetURL := utils.GetPublicRpcUrl("mainnet")

	mainnetClient, err := ethclient.DialContext(context.Background(), rpcMainnetURL)
	if err != nil {
		log.Printf("[%s] dial: %v", "mainnet", err)
		return
	}
	defer mainnetClient.Close()

	currentBlockNumber, err := mainnetClient.BlockNumber(context.Background())
	if err != nil {
		log.Printf("[%s] dial: %v", "mainnet", err)
		return
	}

	gcAbi, err := curveGC.CurveGCMetaData.GetAbi()
	if err != nil {
		log.Printf("[%s] dial: %v", "mainnet", err)
		return
	}

	calls := make([]multicall.Multicall3Call, 0)
	for _, gauge := range gauges {
		chain := strings.ToLower(strings.TrimSpace(gauge.BlockchainId))
		if !gauge.IsPool || !common.IsHexAddress(gauge.Gauge) || !utils.ArrayContains(chainsSupported, chain) {
			continue
		}

		gaugeAddress := gauge.RootGauge
		if len(gaugeAddress) == 0 {
			gaugeAddress = gauge.Gauge
		}

		get_gauge_weight, err := gcAbi.Pack("get_gauge_weight", common.HexToAddress(gaugeAddress))
		if err != nil {
			log.Printf("[%s] dial: %v", "mainnet", err)
			return
		}

		calls = append(calls, multicall.Multicall3Call{
			Target:   utils.CURVE_GC_ADDRESS,
			CallData: get_gauge_weight,
		})
	}

	multicallResponses := utils.Multicall(mainnetClient, calls, MULTICALL_CHAIN["ethereum"], big.NewInt(int64(currentBlockNumber)))

	// Group by chain
	byChain := map[string][]common.Address{}
	for _, gauge := range gauges {
		chain := strings.ToLower(strings.TrimSpace(gauge.BlockchainId))
		if !gauge.IsPool || !common.IsHexAddress(gauge.Gauge) || !utils.ArrayContains(chainsSupported, chain) {
			continue
		}

		gaugeWeightRes := multicallResponses[0]
		multicallResponses = multicallResponses[1:]

		if !gaugeWeightRes.Success {
			log.Printf("Error when fetching gauge weight for %s", gauge.Gauge)
			return
		}

		gaugeWeight := big.NewInt(0).SetBytes(gaugeWeightRes.ReturnData)
		if gaugeWeight.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		byChain[chain] = append(byChain[chain], common.HexToAddress(gauge.Gauge))
	}

	gaugeABI := mustABI(gaugeABIJSON)
	mcABI := mustABI(multicallABIJSON)
	ctx := context.Background()

	for chain, gauges := range byChain {
		rpcURL := utils.GetPublicRpcUrl(chain)
		if len(rpcURL) == 0 {
			log.Printf("[%s]", chain)
			continue
		}

		client, err := ethclient.DialContext(ctx, rpcURL)
		if err != nil {
			log.Printf("[%s] dial: %v", chain, err)
			continue
		}
		defer client.Close()

		// chainID & base tx opts
		chainID, err := client.NetworkID(ctx)
		if err != nil {
			log.Printf("[%s] network id: %v", chain, err)
			continue
		}
		nonce, err := client.PendingNonceAt(ctx, from)
		if err != nil {
			log.Printf("[%s] nonce: %v", chain, err)
			continue
		}
		log.Printf("= %s (chainID=%s) | %d gauges", strings.ToUpper(chain), chainID.String(), len(gauges))

		// Prepare batches of up to batchSize gauges
		for bi, batch := range chunk(gauges, batchSize) {
			// Build Call3[] for aggregate3
			type Call3 struct {
				Target       common.Address
				AllowFailure bool
				CallData     []byte
			}
			call3s := make([]Call3, 0, len(batch))
			for _, g := range batch {
				cd, err := gaugeABI.Pack("user_checkpoint", utils.MULTICALL_MAINNET)
				if err != nil {
					log.Fatalf("[%s] abi pack user_checkpoint: %v", chain, err)
				}
				call3s = append(call3s, Call3{
					Target:       g,
					AllowFailure: true,
					CallData:     cd,
				})
			}

			// Encode aggregate3(call3s)
			data, err := mcABI.Pack("aggregate3", call3s)
			if err != nil {
				log.Fatalf("[%s] pack aggregate3: %v", chain, err)
			}

			// Gas estimation
			call := ethereum.CallMsg{
				From: from, To: &utils.MULTICALL_MAINNET, Value: big.NewInt(0), Data: data,
			}
			gas, err := client.EstimateGas(ctx, call)
			if err != nil {
				// On L2 certains RPC refusent estimate sur gros calldata -> fallback conservateur
				log.Printf("[%s] estimate failed (%v), fallback gas", chain, err)
				continue
			}

			// Fees (EIP-1559)
			tip, err := client.SuggestGasTipCap(ctx)
			if err != nil {
				log.Fatalf("[%s] suggest tip: %v", chain, err)
			}
			h, err := client.HeaderByNumber(ctx, nil)
			if err != nil {
				log.Fatalf("[%s] header: %v", chain, err)
			}
			base := big.NewInt(0)
			if h.BaseFee != nil {
				base = new(big.Int).Set(h.BaseFee)
			}
			maxFee := new(big.Int).Mul(base, big.NewInt(2))
			maxFee.Add(maxFee, tip)

			// Build tx
			tx := types.NewTx(&types.DynamicFeeTx{
				ChainID:   chainID,
				Nonce:     nonce,
				To:        &utils.MULTICALL_MAINNET,
				Gas:       gas,
				GasTipCap: tip,
				GasFeeCap: maxFee,
				Value:     big.NewInt(0), // aggregate3 is payable but we send 0
				Data:      data,
			})

			if dryRun {
				log.Printf("[dry-run] %s batch %d/%d: %d calls  nonce=%d gas≈%d maxFee=%s tip=%s",
					strings.ToUpper(chain), bi+1, (len(gauges)+batchSize-1)/batchSize, len(batch),
					nonce, gas, maxFee.String(), tip.String())
			} else {
				signed, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), key)
				if err != nil {
					log.Printf("[%s] sign tx: %v", chain, err)
					continue
				}
				if err := client.SendTransaction(ctx, signed); err != nil {
					log.Printf("[%s] send tx: %v", chain, err)
					continue
				}
				log.Printf("sent %s batch %d/%d -> %s  nonce=%d gas=%d",
					strings.ToUpper(chain), bi+1, (len(gauges)+batchSize-1)/batchSize, signed.Hash().Hex(), nonce, gas)
			}

			nonce++
			time.Sleep(300 * time.Millisecond)
		}
		log.Printf("Done %s ✅", strings.ToUpper(chain))
	}
}
