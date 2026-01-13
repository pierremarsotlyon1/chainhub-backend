package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/erc20"
	"main/contracts/llamalendController"
	"main/contracts/llamalendFactory"
	"main/contracts/llamalendVault"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	LLAMALEND_PER_DAY_PATH    = "./data/llamalend/llamalend-per-day.json"
	LLAMALEND_CONFIGS_MAINNET = "./data/configs/llamalend-mainnet.json"
	LLAMALEND_CONFIGS_ARB     = "./data/configs/llamalend-arb.json"

	BUCKET_LLAMALEND_DIR          = "data/llamalend"
	BUCKET_LLAMALEND_PER_DAY_FILE = BUCKET_LLAMALEND_DIR + "/llamalend-per-day.json"

	// Parallelism settings - reduced to avoid RPC rate limits
	MAX_CONCURRENT_LIQUIDATION_QUERIES = 5
	MAX_CONCURRENT_MARKET_FETCHES      = 5
	MAX_CONCURRENT_DAYS                = 5
	BLOCK_RANGE_SIZE                   = 10000
	RPC_TIMEOUT                        = 60 * time.Second
)

var LLAMALEND_FACTORIES = []interfaces.LlamalendConfig{
	{
		FactoryAddress:  common.HexToAddress("0xeA6876DDE9e3467564acBeE1Ed5bac88783205E0"),
		ChainId:         1,
		MarketName:      "ethereum",
		TimestampDeploy: 1710294146,
		BlockDeploy:     19422660,
		RpcUrl:          utils.GetPublicRpcUrl("mainnet"),
		ConfigPath:      LLAMALEND_CONFIGS_MAINNET,
	},
	{
		FactoryAddress:  common.HexToAddress("0xcaEC110C784c9DF37240a8Ce096D352A75922DeA"),
		ChainId:         42161,
		MarketName:      "arbitrum",
		TimestampDeploy: 1711266146,
		BlockDeploy:     193652535,
		RpcUrl:          utils.GetPublicRpcUrl("arbitrum"),
		ConfigPath:      LLAMALEND_CONFIGS_ARB,
	},
}

// ============================================================================
// LOGGER
// ============================================================================

type Logger struct {
	prefix string
}

func newLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("ℹ️  [%s] %s", l.prefix, msg)
}

func (l *Logger) Success(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("✅ [%s] %s", l.prefix, msg)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("⚠️  [%s] %s", l.prefix, msg)
}

func (l *Logger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("❌ [%s] %s", l.prefix, msg)
}

func (l *Logger) Progress(current, total int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	percent := float64(current) / float64(total) * 100
	log.Printf("🔄 [%s] [%d/%d %.1f%%] %s", l.prefix, current, total, percent, msg)
}

func (l *Logger) Start(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("🚀 [%s] %s", l.prefix, msg)
}

func (l *Logger) Done(duration time.Duration, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("🏁 [%s] %s (took %s)", l.prefix, msg, duration.Round(time.Millisecond))
}

func (l *Logger) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("🔍 [%s] %s", l.prefix, msg)
}

func (l *Logger) Separator() {
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// ============================================================================
// CACHES
// ============================================================================

type tokenCache struct {
	sync.RWMutex
	symbols  map[common.Address]string
	decimals map[common.Address]uint8
}

// Cache for block numbers by timestamp
type blockNumberCache struct {
	sync.RWMutex
	cache map[string]uint64
}

// Cache for market addresses (controller, vault, amm) - these never change
type marketAddressCache struct {
	sync.RWMutex
	controllers map[string]common.Address // key: "factoryAddr-marketId"
	vaults      map[string]common.Address
	amms        map[string]common.Address
}

// Cache for controller token addresses - these never change
type controllerTokenCache struct {
	sync.RWMutex
	collateralTokens map[common.Address]common.Address // controller -> collateral token
	borrowedTokens   map[common.Address]common.Address // controller -> borrowed token
}

var cache = &tokenCache{
	symbols:  make(map[common.Address]string),
	decimals: make(map[common.Address]uint8),
}

var blockCache = &blockNumberCache{
	cache: make(map[string]uint64),
}

var marketAddrCache = &marketAddressCache{
	controllers: make(map[string]common.Address),
	vaults:      make(map[string]common.Address),
	amms:        make(map[string]common.Address),
}

var ctrlTokenCache = &controllerTokenCache{
	collateralTokens: make(map[common.Address]common.Address),
	borrowedTokens:   make(map[common.Address]common.Address),
}

// Block number cache methods
func (c *blockNumberCache) get(chainName string, timestamp uint64) (uint64, bool) {
	c.RLock()
	defer c.RUnlock()
	key := fmt.Sprintf("%s-%d", chainName, timestamp)
	block, ok := c.cache[key]
	return block, ok
}

func (c *blockNumberCache) set(chainName string, timestamp uint64, blockNumber uint64) {
	c.Lock()
	defer c.Unlock()
	key := fmt.Sprintf("%s-%d", chainName, timestamp)
	c.cache[key] = blockNumber
}

func getBlockNumberCached(chainName string, timestamp uint64) uint64 {
	if block, ok := blockCache.get(chainName, timestamp); ok {
		return block
	}
	block := utils.GetBlockNumberByTimestamp(chainName, timestamp)
	blockCache.set(chainName, timestamp, block)
	return block
}

// Market address cache methods
func (c *marketAddressCache) makeKey(factory common.Address, marketId int) string {
	return fmt.Sprintf("%s-%d", strings.ToLower(factory.Hex()), marketId)
}

func (c *marketAddressCache) get(factory common.Address, marketId int) (controller, vault, amm common.Address, found bool) {
	c.RLock()
	defer c.RUnlock()
	key := c.makeKey(factory, marketId)
	controller, ok1 := c.controllers[key]
	vault, ok2 := c.vaults[key]
	amm, ok3 := c.amms[key]
	return controller, vault, amm, ok1 && ok2 && ok3
}

func (c *marketAddressCache) set(factory common.Address, marketId int, controller, vault, amm common.Address) {
	c.Lock()
	defer c.Unlock()
	key := c.makeKey(factory, marketId)
	c.controllers[key] = controller
	c.vaults[key] = vault
	c.amms[key] = amm
}

// Controller token cache methods
func (c *controllerTokenCache) get(controller common.Address) (collateral, borrowed common.Address, found bool) {
	c.RLock()
	defer c.RUnlock()
	collateral, ok1 := c.collateralTokens[controller]
	borrowed, ok2 := c.borrowedTokens[controller]
	return collateral, borrowed, ok1 && ok2
}

func (c *controllerTokenCache) set(controller common.Address, collateral, borrowed common.Address) {
	c.Lock()
	defer c.Unlock()
	c.collateralTokens[controller] = collateral
	c.borrowedTokens[controller] = borrowed
}

func (c *tokenCache) getSymbol(addr common.Address) (string, bool) {
	c.RLock()
	defer c.RUnlock()
	s, ok := c.symbols[addr]
	return s, ok
}

func (c *tokenCache) setSymbol(addr common.Address, symbol string) {
	c.Lock()
	defer c.Unlock()
	c.symbols[addr] = symbol
}

func (c *tokenCache) getDecimals(addr common.Address) (uint8, bool) {
	c.RLock()
	defer c.RUnlock()
	d, ok := c.decimals[addr]
	return d, ok
}

func (c *tokenCache) setDecimals(addr common.Address, decimals uint8) {
	c.Lock()
	defer c.Unlock()
	c.decimals[addr] = decimals
}

func (c *tokenCache) stats() (int, int) {
	c.RLock()
	defer c.RUnlock()
	return len(c.symbols), len(c.decimals)
}

// ============================================================================
// MARKET INDEX
// ============================================================================

type MarketIndex struct {
	sync.RWMutex
	index map[string]int
}

func newMarketIndex(markets []interfaces.LlamalendMarket) *MarketIndex {
	idx := &MarketIndex{
		index: make(map[string]int, len(markets)),
	}
	for i, m := range markets {
		key := idx.makeKey(m.FactoryAddress, m.ControllerAddress, m.Id, m.Timestamp)
		idx.index[key] = i
	}
	return idx
}

func (idx *MarketIndex) makeKey(factory, controller common.Address, id int, timestamp uint64) string {
	return fmt.Sprintf("%s-%s-%d-%d",
		strings.ToLower(factory.Hex()),
		strings.ToLower(controller.Hex()),
		id,
		timestamp)
}

func (idx *MarketIndex) get(factory, controller common.Address, id int, timestamp uint64) (int, bool) {
	idx.RLock()
	defer idx.RUnlock()
	i, ok := idx.index[idx.makeKey(factory, controller, id, timestamp)]
	return i, ok
}

func (idx *MarketIndex) set(factory, controller common.Address, id int, timestamp uint64, position int) {
	idx.Lock()
	defer idx.Unlock()
	idx.index[idx.makeKey(factory, controller, id, timestamp)] = position
}

// ============================================================================
// MAIN FUNCTIONS
// ============================================================================

// LlamalendFixSort reads existing data and re-saves it sorted
func LlamalendFixSort() error {
	logger := newLogger("FIX-SORT")
	logger.Start("Reading and sorting existing llamalend data...")

	markets := readMarkets()
	if len(markets) == 0 {
		logger.Warning("No markets found to sort")
		return nil
	}

	logger.Info("Loaded %d markets, saving sorted...", len(markets))
	writeMarkets(markets) // writeMarkets already sorts

	logger.Success("Done! Data is now sorted by date")
	return nil
}

// Llamalend fetches data from the last saved block/timestamp
func Llamalend() error {
	return llamalendFetch(false)
}

// LlamalendHistorical fetches all historical data
func LlamalendHistorical() error {
	return llamalendFetch(true)
}

func llamalendFetch(historical bool) error {
	mainLog := newLogger("MAIN")
	totalStart := time.Now()

	mainLog.Separator()
	mode := "INCREMENTAL"
	if historical {
		mode = "HISTORICAL"
	}
	mainLog.Start("Starting Llamalend fetch - Mode: %s", mode)
	mainLog.Info("Timestamp: %s", time.Now().Format("2006-01-02 15:04:05"))
	mainLog.Separator()

	// Fetch curve pools
	mainLog.Info("Fetching Curve pools...")
	poolsStart := time.Now()
	curvePools, err := utils.GetAllCurvePools()
	if err != nil {
		mainLog.Error("Failed to get curve pools: %v", err)
		return fmt.Errorf("failed to get curve pools: %w", err)
	}
	mainLog.Done(time.Since(poolsStart), "Loaded %d Curve pools", len(curvePools))

	// Load existing markets
	mainLog.Info("Loading existing markets from storage...")
	loadStart := time.Now()
	currentMarkets := readMarkets()
	marketIndex := newMarketIndex(currentMarkets)
	mainLog.Done(time.Since(loadStart), "Loaded %d existing markets (index size: %d)", len(currentMarkets), len(marketIndex.index))

	now := uint64(time.Now().Unix())

	// Stats
	totalMarketsAdded := 0
	totalMarketsUpdated := 0
	totalLiquidationsFound := int64(0)

	for factoryIdx, factory := range LLAMALEND_FACTORIES {
		factoryLog := newLogger(strings.ToUpper(factory.MarketName))
		factoryStart := time.Now()

		factoryLog.Separator()
		factoryLog.Start("Processing factory %d/%d: %s",
			factoryIdx+1, len(LLAMALEND_FACTORIES),
			factory.MarketName)
		factoryLog.Info("Factory address: %s", factory.FactoryAddress.Hex())
		factoryLog.Info("Chain ID: %d", factory.ChainId)

		// Connect to RPC
		factoryLog.Info("Connecting to RPC...")
		rpcStart := time.Now()
		client, err := ethclient.Dial(factory.RpcUrl)
		if err != nil {
			factoryLog.Error("Failed to connect: %v", err)
			return fmt.Errorf("failed to connect to %s: %w", factory.MarketName, err)
		}
		defer client.Close()
		factoryLog.Done(time.Since(rpcStart), "Connected to RPC")

		config := utils.ReadConfig(factory.ConfigPath)
		factoryLog.Info("Config loaded - Last processed block: %d", config.LastBlock)

		currentBlock, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			factoryLog.Error("Failed to get current block: %v", err)
			return fmt.Errorf("failed to get current block: %w", err)
		}
		factoryLog.Info("Current block: %d | Block timestamp: %s",
			currentBlock.Number.Uint64(),
			time.Unix(int64(currentBlock.Time), 0).Format("2006-01-02 15:04:05"))

		// Determine start timestamp based on config or deploy time
		var startTimestamp uint64
		if config.LastBlock > 0 && config.LastBlock > factory.BlockDeploy {
			// Resume from last processed block - get its timestamp
			lastBlockHeader, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(config.LastBlock)))
			if err != nil {
				factoryLog.Warning("Failed to get last block timestamp, falling back to deploy time: %v", err)
				startTimestamp = factory.TimestampDeploy
			} else {
				// Start from the day after the last processed block
				startTimestamp = uint64(utils.GetStartOfDay(lastBlockHeader.Time))
				factoryLog.Info("Resuming from last processed block %d (timestamp: %s)",
					config.LastBlock,
					time.Unix(int64(startTimestamp), 0).Format("2006-01-02"))
			}
		} else {
			startTimestamp = factory.TimestampDeploy
			factoryLog.Info("Starting from deploy time: %s",
				time.Unix(int64(startTimestamp), 0).Format("2006-01-02"))
		}

		// Determine timestamps to process
		var timestamps []uint64
		if historical {
			for ts := startTimestamp + utils.DAY_TO_SEC; ts < currentBlock.Time; ts += utils.DAY_TO_SEC {
				timestamps = append(timestamps, ts)
			}
			if len(timestamps) == 0 {
				factoryLog.Success("Already up to date! No new days to process")
			} else {
				factoryLog.Info("Historical mode: %d days to process (from %s to %s)",
					len(timestamps),
					time.Unix(int64(startTimestamp), 0).Format("2006-01-02"),
					time.Unix(int64(currentBlock.Time), 0).Format("2006-01-02"))
			}
		} else {
			timestamps = []uint64{now}
			factoryLog.Info("Incremental mode: processing current day only")
		}

		// Collect all markets for this factory
		factoryMarkets := make([]interfaces.LlamalendMarket, 0)
		var marketsUpdated int64
		var marketsAdded int64
		var marketsFailed int64

		factoryLog.Info("Starting market data fetch...")
		fetchStart := time.Now()

		if len(timestamps) > 1 {
			// Process multiple days in parallel
			type dayResult struct {
				tsIdx   int
				markets []interfaces.LlamalendMarket
				err     error
			}

			// Pre-fetch all block numbers sequentially (API calls may be rate-limited)
			factoryLog.Info("Pre-fetching block numbers for %d days...", len(timestamps))
			blockFetchStart := time.Now()
			blockNumbers := make(map[uint64]uint64) // timestamp -> blockNumber
			for i, ts := range timestamps {
				blockNumber := getBlockNumberCached(factory.MarketName, ts)
				blockNumbers[ts] = blockNumber
				if (i+1)%20 == 0 || i == len(timestamps)-1 {
					factoryLog.Progress(i+1, len(timestamps), "Fetched block numbers")
				}
			}
			factoryLog.Done(time.Since(blockFetchStart), "Block numbers fetched")

			factoryLog.Info("Starting parallel fetch: %d days with %d concurrent workers", len(timestamps), MAX_CONCURRENT_DAYS)

			results := make(chan dayResult, len(timestamps))
			sem := make(chan struct{}, MAX_CONCURRENT_DAYS)
			var wg sync.WaitGroup

			// Track active goroutines
			var activeWorkers int32

			for tsIdx, blockTimestamp := range timestamps {
				wg.Add(1)
				go func(idx int, ts uint64) {
					defer wg.Done()

					sem <- struct{}{} // Acquire semaphore FIRST
					currentActive := atomic.AddInt32(&activeWorkers, 1)
					dayTs := utils.GetStartOfDay(ts)
					dayStr := time.Unix(int64(dayTs), 0).Format("2006-01-02")

					// Log when worker starts (only first few)
					if idx < 5 {
						factoryLog.Debug("Worker %d started: %s (active: %d)", idx, dayStr, currentActive)
					}

					defer func() {
						atomic.AddInt32(&activeWorkers, -1)
						<-sem // Release semaphore
					}()

					blockNumber := blockNumbers[ts] // Use pre-fetched block number

					opts := &bind.CallOpts{
						BlockNumber: big.NewInt(int64(blockNumber)),
					}

					markets, err := fetchLlamalendMarkets(opts, factory, client, ts, dayTs, curvePools, factoryLog, false)
					results <- dayResult{tsIdx: idx, markets: markets, err: err}
				}(tsIdx, blockTimestamp)
			}

			factoryLog.Info("All %d goroutines launched, processing...", len(timestamps))

			// Close channel when done
			go func() {
				wg.Wait()
				close(results)
			}()

			// Collect results and show progress
			processedDays := 0
			lastLogTime := time.Now()
			for res := range results {
				processedDays++
				dayTs := utils.GetStartOfDay(timestamps[res.tsIdx])
				dayStr := time.Unix(int64(dayTs), 0).Format("2006-01-02")

				if res.err != nil {
					atomic.AddInt64(&marketsFailed, 1)
					factoryLog.Warning("Day %s failed: %v", dayStr, res.err)
					continue
				}

				// Merge markets using index for O(1) lookup
				for _, market := range res.markets {
					if pos, found := marketIndex.get(market.FactoryAddress, market.ControllerAddress, market.Id, market.Timestamp); found {
						updateMarket(&currentMarkets[pos], &market)
						atomic.AddInt64(&marketsUpdated, 1)
					} else {
						pos := len(currentMarkets)
						currentMarkets = append(currentMarkets, market)
						marketIndex.set(market.FactoryAddress, market.ControllerAddress, market.Id, market.Timestamp, pos)
						atomic.AddInt64(&marketsAdded, 1)
					}
				}

				factoryMarkets = append(factoryMarkets, res.markets...)

				// Progress log every 2 days or every second or at the end
				if processedDays%2 == 0 || time.Since(lastLogTime) > 1*time.Second || processedDays == len(timestamps) {
					elapsed := time.Since(fetchStart)
					rate := float64(processedDays) / elapsed.Seconds()
					remaining := len(timestamps) - processedDays
					eta := time.Duration(float64(remaining)/rate) * time.Second
					factoryLog.Progress(processedDays, len(timestamps),
						"%s | %.1f days/s | ETA: %s | Active: %d",
						dayStr, rate, eta.Round(time.Second), atomic.LoadInt32(&activeWorkers))
					lastLogTime = time.Now()
				}
			}
		} else {
			// Single day - process normally
			for tsIdx, blockTimestamp := range timestamps {
				blockNumber := getBlockNumberCached(factory.MarketName, blockTimestamp)
				dayTimestamp := utils.GetStartOfDay(blockTimestamp)
				dayStr := time.Unix(int64(dayTimestamp), 0).Format("2006-01-02")

				factoryLog.Info("Fetching data for %s (block %d)", dayStr, blockNumber)

				opts := &bind.CallOpts{
					BlockNumber: big.NewInt(int64(blockNumber)),
				}

				markets, err := fetchLlamalendMarkets(opts, factory, client, blockTimestamp, dayTimestamp, curvePools, factoryLog, true)
				if err != nil {
					factoryLog.Warning("Failed to fetch markets for %s: %v", dayStr, err)
					marketsFailed++
					continue
				}

				// Merge markets using index for O(1) lookup
				for _, market := range markets {
					if pos, found := marketIndex.get(market.FactoryAddress, market.ControllerAddress, market.Id, market.Timestamp); found {
						updateMarket(&currentMarkets[pos], &market)
						marketsUpdated++
					} else {
						pos := len(currentMarkets)
						currentMarkets = append(currentMarkets, market)
						marketIndex.set(market.FactoryAddress, market.ControllerAddress, market.Id, market.Timestamp, pos)
						marketsAdded++
					}
				}

				factoryMarkets = append(factoryMarkets, markets...)

				if len(timestamps) > 1 {
					factoryLog.Progress(tsIdx+1, len(timestamps), "Processing %s", dayStr)
				}
			}
		}

		factoryLog.Done(time.Since(fetchStart), "Market fetch completed")
		factoryLog.Success("Markets - Added: %d | Updated: %d | Failed days: %d",
			marketsAdded, marketsUpdated, marketsFailed)

		totalMarketsAdded += int(marketsAdded)
		totalMarketsUpdated += int(marketsUpdated)

		// Fetch hard liquidations
		// Use config.LastBlock if available, otherwise use deploy block
		blockFrom := config.LastBlock
		if blockFrom == 0 || blockFrom < factory.BlockDeploy {
			blockFrom = factory.BlockDeploy
		}

		blocksToProcess := currentBlock.Number.Int64() - int64(blockFrom)
		factoryLog.Info("Fetching hard liquidations...")
		factoryLog.Info("Block range: %d → %d (%d blocks, ~%d ranges)",
			blockFrom, currentBlock.Number.Uint64(),
			blocksToProcess, blocksToProcess/BLOCK_RANGE_SIZE+1)

		liqStart := time.Now()

		liquidationCounts, totalLiqs, err := fetchHardLiquidationsParallel(
			client,
			factoryMarkets,
			int64(blockFrom),
			currentBlock.Number.Int64(),
			factoryLog,
		)
		if err != nil {
			factoryLog.Warning("Some liquidation fetches had errors: %v", err)
		}

		factoryLog.Done(time.Since(liqStart), "Liquidation fetch completed")
		factoryLog.Success("Found %d liquidations across %d controllers", totalLiqs, len(liquidationCounts))

		totalLiquidationsFound += totalLiqs

		// Apply liquidation counts
		liqApplied := 0
		controllersWithLiqs := 0
		for controllerAddr, dayLiqs := range liquidationCounts {
			if len(dayLiqs) > 0 {
				controllersWithLiqs++
			}
			for dayTs, count := range dayLiqs {
				for i := range currentMarkets {
					if strings.EqualFold(currentMarkets[i].ControllerAddress.Hex(), controllerAddr.Hex()) &&
						currentMarkets[i].Timestamp == uint64(dayTs) {
						currentMarkets[i].HardLiquidation += count
						liqApplied += int(count)
						break
					}
				}
			}
		}
		factoryLog.Info("Applied %d liquidation events to %d controllers", liqApplied, controllersWithLiqs)

		// Save config with new block number
		utils.WriteConfig(config, currentBlock.Number.Uint64(), factory.ConfigPath)
		factoryLog.Success("Config updated - New last block: %d", currentBlock.Number.Uint64())

		// Cache stats
		symCount, decCount := cache.stats()
		factoryLog.Debug("Token cache stats - Symbols: %d | Decimals: %d", symCount, decCount)

		factoryLog.Done(time.Since(factoryStart), "Factory %s completed", factory.MarketName)
	}

	// Save all markets
	mainLog.Separator()
	mainLog.Info("Saving all markets to storage...")
	saveStart := time.Now()
	writeMarkets(currentMarkets)
	mainLog.Done(time.Since(saveStart), "Markets saved")

	// Final summary
	mainLog.Separator()
	mainLog.Success("═══ FINAL SUMMARY ═══")
	mainLog.Info("Total markets in database: %d", len(currentMarkets))
	mainLog.Info("Markets added this run: %d", totalMarketsAdded)
	mainLog.Info("Markets updated this run: %d", totalMarketsUpdated)
	mainLog.Info("Liquidations found: %d", totalLiquidationsFound)

	// Cache stats
	symCount, decCount := cache.stats()
	mainLog.Info("Cache stats - Tokens: %d symbols, %d decimals", symCount, decCount)
	mainLog.Info("Cache stats - Block numbers: %d | Market addresses: %d | Controller tokens: %d",
		len(blockCache.cache), len(marketAddrCache.controllers), len(ctrlTokenCache.collateralTokens))

	mainLog.Done(time.Since(totalStart), "Llamalend fetch completed successfully")
	mainLog.Separator()

	return nil
}

func updateMarket(existing, new *interfaces.LlamalendMarket) {
	existing.TotalDebt = new.TotalDebt
	existing.TotalDebtUSD = new.TotalDebtUSD
	existing.Supplied = new.Supplied
	existing.SuppliedUSD = new.SuppliedUSD
	existing.Collateral = new.Collateral
	existing.CollateralUSD = new.CollateralUSD
	existing.BorrowedTokenAddress = new.BorrowedTokenAddress
	existing.CollateralTokenAddress = new.CollateralTokenAddress
	existing.ChainId = new.ChainId
	existing.Loans = new.Loans
	existing.BorrowApr = new.BorrowApr
	existing.SupplyApr = new.SupplyApr
}

// ============================================================================
// MARKET FETCHING
// ============================================================================

func fetchLlamalendMarkets(opts *bind.CallOpts, factory interfaces.LlamalendConfig, client *ethclient.Client, blockTimestamp uint64, dayTimestamp int64, curvePools []interfaces.CurvePool, logger *Logger, verbose bool) ([]interfaces.LlamalendMarket, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), RPC_TIMEOUT)
	defer cancel()

	factoryContract, err := llamalendFactory.NewLlamalendFactory(factory.FactoryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory contract: %w", err)
	}

	optsWithCtx := &bind.CallOpts{
		BlockNumber: opts.BlockNumber,
		Context:     ctx,
	}

	marketCount, err := factoryContract.MarketCount(optsWithCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to get market count (block %d): %w", opts.BlockNumber, err)
	}

	totalMarkets := int(marketCount.Int64())
	if verbose {
		logger.Info("Factory has %d markets - fetching in parallel...", totalMarkets)
	}

	// Fetch all markets in parallel but with limited concurrency
	type marketResult struct {
		index  int
		market interfaces.LlamalendMarket
		err    error
	}

	results := make(chan marketResult, totalMarkets)
	sem := make(chan struct{}, MAX_CONCURRENT_MARKET_FETCHES)
	var wg sync.WaitGroup

	for i := 0; i < totalMarkets; i++ {
		wg.Add(1)
		go func(marketId int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			market, err := fetchSingleMarket(optsWithCtx, factoryContract, factory, client, marketId, blockTimestamp, dayTimestamp, curvePools)
			results <- marketResult{index: marketId, market: market, err: err}
		}(i)
	}

	// Close results channel when all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	markets := make([]interfaces.LlamalendMarket, 0, totalMarkets)
	var successCount, errorCount int

	for res := range results {
		if res.err != nil {
			errorCount++
			if verbose {
				logger.Warning("Market %d failed: %v", res.index, res.err)
			}
			continue
		}
		markets = append(markets, res.market)
		successCount++

		if verbose {
			logger.Debug("Market %d: %s | Debt: $%.2f | TVL: $%.2f | Loans: %d",
				res.index, res.market.MarketName, res.market.TotalDebtUSD, res.market.SuppliedUSD, res.market.Loans)
		}
	}

	if verbose {
		logger.Success("Fetched %d/%d markets (errors: %d)", successCount, totalMarkets, errorCount)
	}

	return markets, nil
}

func fetchSingleMarket(opts *bind.CallOpts, factoryContract *llamalendFactory.LlamalendFactory, factory interfaces.LlamalendConfig, client *ethclient.Client, marketId int, blockTimestamp uint64, dayTimestamp int64, curvePools []interfaces.CurvePool) (interfaces.LlamalendMarket, error) {
	// Use the context from opts (already has timeout set by caller)
	optsWithCtx := opts

	// Try to get market addresses from cache first
	controllerAddress, vaultAddress, ammAddress, found := marketAddrCache.get(factory.FactoryAddress, marketId)

	if !found {
		// Fetch from RPC in parallel
		idx := big.NewInt(int64(marketId))
		var factoryErr error
		var factoryErrMu sync.Mutex
		var factoryWg sync.WaitGroup

		setFactoryErr := func(err error) {
			factoryErrMu.Lock()
			if factoryErr == nil {
				factoryErr = err
			}
			factoryErrMu.Unlock()
		}

		factoryWg.Add(3)
		go func() {
			defer factoryWg.Done()
			var err error
			controllerAddress, err = factoryContract.Controllers(optsWithCtx, idx)
			if err != nil {
				setFactoryErr(fmt.Errorf("controllers: %w", err))
			}
		}()
		go func() {
			defer factoryWg.Done()
			var err error
			vaultAddress, err = factoryContract.Vaults(optsWithCtx, idx)
			if err != nil {
				setFactoryErr(fmt.Errorf("vaults: %w", err))
			}
		}()
		go func() {
			defer factoryWg.Done()
			var err error
			ammAddress, err = factoryContract.Amms(optsWithCtx, idx)
			if err != nil {
				setFactoryErr(fmt.Errorf("amms: %w", err))
			}
		}()
		factoryWg.Wait()

		if factoryErr != nil {
			return interfaces.LlamalendMarket{}, factoryErr
		}

		// Cache the addresses
		marketAddrCache.set(factory.FactoryAddress, marketId, controllerAddress, vaultAddress, ammAddress)
	}

	controllerContract, err := llamalendController.NewLlamalendController(controllerAddress, client)
	if err != nil {
		return interfaces.LlamalendMarket{}, err
	}

	vaultContract, err := llamalendVault.NewLlamalendVault(vaultAddress, client)
	if err != nil {
		return interfaces.LlamalendMarket{}, err
	}

	// Try to get token addresses from cache
	collateralTokenAddress, borrowedTokenAddress, tokensFound := ctrlTokenCache.get(controllerAddress)

	// Fetch contract data in parallel
	var wg sync.WaitGroup
	var borrowAprBN, supplyBN, totalDebtBN *big.Int
	var nbLoans *big.Int
	var fetchErr error
	var errMu sync.Mutex

	setErr := func(err error) {
		errMu.Lock()
		if fetchErr == nil {
			fetchErr = err
		}
		errMu.Unlock()
	}

	// Always need to fetch: borrowApr, supplyApr, nLoans, totalDebt (these change)
	wg.Add(4)

	go func() {
		defer wg.Done()
		var err error
		borrowAprBN, err = vaultContract.BorrowApr(optsWithCtx)
		if err != nil {
			setErr(fmt.Errorf("borrowApr: %w", err))
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		supplyBN, err = vaultContract.LendApr(optsWithCtx)
		if err != nil {
			setErr(fmt.Errorf("lendApr: %w", err))
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		nbLoans, err = controllerContract.NLoans(optsWithCtx)
		if err != nil {
			setErr(fmt.Errorf("nLoans: %w", err))
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		totalDebtBN, err = controllerContract.TotalDebt(optsWithCtx)
		if err != nil {
			setErr(fmt.Errorf("totalDebt: %w", err))
		}
	}()

	// Only fetch token addresses if not cached
	if !tokensFound {
		wg.Add(2)
		go func() {
			defer wg.Done()
			var err error
			collateralTokenAddress, err = controllerContract.CollateralToken(optsWithCtx)
			if err != nil {
				setErr(fmt.Errorf("collateralToken: %w", err))
			}
		}()

		go func() {
			defer wg.Done()
			var err error
			borrowedTokenAddress, err = controllerContract.BorrowedToken(optsWithCtx)
			if err != nil {
				setErr(fmt.Errorf("borrowedToken: %w", err))
			}
		}()
	}

	wg.Wait()

	if fetchErr != nil {
		return interfaces.LlamalendMarket{}, fetchErr
	}

	// Cache token addresses if we just fetched them
	if !tokensFound {
		ctrlTokenCache.set(controllerAddress, collateralTokenAddress, borrowedTokenAddress)
	}

	// Fetch token info and balances in parallel
	var collateralTokenSymbol, borrowTokenSymbol string
	var collateralTokenDecimals, borrowTokenDecimals uint8
	var collateralBN, availableToBorrowBN *big.Int
	var tokenErr error
	var tokenErrMu sync.Mutex

	setTokenErr := func(err error) {
		tokenErrMu.Lock()
		if tokenErr == nil {
			tokenErr = err
		}
		tokenErrMu.Unlock()
	}

	var tokenWg sync.WaitGroup
	tokenWg.Add(4)

	go func() {
		defer tokenWg.Done()
		var err error
		collateralTokenSymbol, collateralTokenDecimals, err = getTokenInfo(optsWithCtx, collateralTokenAddress, client)
		if err != nil {
			setTokenErr(fmt.Errorf("collateral token info: %w", err))
		}
	}()

	go func() {
		defer tokenWg.Done()
		var err error
		borrowTokenSymbol, borrowTokenDecimals, err = getTokenInfo(optsWithCtx, borrowedTokenAddress, client)
		if err != nil {
			setTokenErr(fmt.Errorf("borrow token info: %w", err))
		}
	}()

	go func() {
		defer tokenWg.Done()
		collateralTokenContract, err := erc20.NewErc20(collateralTokenAddress, client)
		if err != nil {
			setTokenErr(fmt.Errorf("collateral contract: %w", err))
			return
		}
		collateralBN, err = collateralTokenContract.BalanceOf(optsWithCtx, ammAddress)
		if err != nil {
			setTokenErr(fmt.Errorf("collateral balance: %w", err))
		}
	}()

	go func() {
		defer tokenWg.Done()
		borrowTokenContract, err := erc20.NewErc20(borrowedTokenAddress, client)
		if err != nil {
			setTokenErr(fmt.Errorf("borrow contract: %w", err))
			return
		}
		availableToBorrowBN, err = borrowTokenContract.BalanceOf(optsWithCtx, controllerAddress)
		if err != nil {
			setTokenErr(fmt.Errorf("available balance: %w", err))
		}
	}()

	tokenWg.Wait()

	if tokenErr != nil {
		return interfaces.LlamalendMarket{}, tokenErr
	}

	// Calculate prices and values
	borrowTokenPrice := getTokenPrice(borrowedTokenAddress, factory.MarketName, blockTimestamp, curvePools)
	collateralTokenPrice := getTokenPrice(collateralTokenAddress, factory.MarketName, blockTimestamp, curvePools)

	totalDebt := utils.Quo(totalDebtBN, uint64(borrowTokenDecimals))
	totalDebtUSD := totalDebt * borrowTokenPrice

	supplied := totalDebt + utils.Quo(availableToBorrowBN, uint64(borrowTokenDecimals))
	suppliedUSD := supplied * borrowTokenPrice

	collateral := utils.Quo(collateralBN, uint64(collateralTokenDecimals))
	collateralUSD := collateral * collateralTokenPrice

	return interfaces.LlamalendMarket{
		Id:                     marketId,
		ControllerAddress:      controllerAddress,
		FactoryAddress:         factory.FactoryAddress,
		MarketName:             collateralTokenSymbol + "/" + borrowTokenSymbol,
		Timestamp:              uint64(dayTimestamp),
		TotalDebt:              totalDebt,
		TotalDebtUSD:           totalDebtUSD,
		Supplied:               supplied,
		SuppliedUSD:            suppliedUSD,
		Collateral:             collateral,
		CollateralUSD:          collateralUSD,
		BorrowedTokenAddress:   borrowedTokenAddress,
		CollateralTokenAddress: collateralTokenAddress,
		ChainId:                factory.ChainId,
		Loans:                  nbLoans.Int64(),
		BorrowApr:              utils.Quo(borrowAprBN, 18),
		SupplyApr:              utils.Quo(supplyBN, 18),
	}, nil
}

func getTokenInfo(opts *bind.CallOpts, tokenAddress common.Address, client *ethclient.Client) (string, uint8, error) {
	symbol, hasSymbol := cache.getSymbol(tokenAddress)
	decimals, hasDecimals := cache.getDecimals(tokenAddress)

	if hasSymbol && hasDecimals {
		return symbol, decimals, nil
	}

	tokenContract, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		return "", 0, err
	}

	if !hasSymbol {
		symbol, err = tokenContract.Symbol(opts)
		if err != nil {
			return "", 0, err
		}
		cache.setSymbol(tokenAddress, symbol)
	}

	if !hasDecimals {
		decimals, err = tokenContract.Decimals(opts)
		if err != nil {
			return "", 0, err
		}
		cache.setDecimals(tokenAddress, decimals)
	}

	return symbol, decimals, nil
}

// ============================================================================
// LIQUIDATION FETCHING
// ============================================================================

func fetchHardLiquidationsParallel(client *ethclient.Client, markets []interfaces.LlamalendMarket, blockFrom, blockTo int64, logger *Logger) (map[common.Address]map[int64]uint64, int64, error) {
	// Get unique controllers
	controllers := make(map[common.Address]bool)
	for _, m := range markets {
		controllers[m.ControllerAddress] = true
	}

	logger.Info("Processing %d unique controllers", len(controllers))

	// Results map: controller -> dayTimestamp -> count
	results := make(map[common.Address]map[int64]uint64)
	var resultsMu sync.Mutex
	var totalLiquidations int64

	// Progress tracking
	var processed int32
	totalControllers := int32(len(controllers))
	startTime := time.Now()

	// Use semaphore for rate limiting
	sem := make(chan struct{}, MAX_CONCURRENT_LIQUIDATION_QUERIES)
	var wg sync.WaitGroup
	var errCount int32

	for controller := range controllers {
		wg.Add(1)
		go func(ctrl common.Address) {
			defer wg.Done()

			sem <- struct{}{}        // Acquire
			defer func() { <-sem }() // Release

			liquidations, err := fetchHardLiquidationFast(client, ctrl, blockFrom, blockTo)
			current := atomic.AddInt32(&processed, 1)

			if err != nil {
				atomic.AddInt32(&errCount, 1)
				if current%20 == 0 || current == totalControllers {
					logger.Progress(int(current), int(totalControllers),
						"Processing... (errors: %d)", atomic.LoadInt32(&errCount))
				}
				return
			}

			// Group by day
			dayMap := make(map[int64]uint64)
			for _, liq := range liquidations {
				dayTs := utils.GetStartOfDay(liq.Timestamp)
				dayMap[dayTs]++
			}

			resultsMu.Lock()
			results[ctrl] = dayMap
			atomic.AddInt64(&totalLiquidations, int64(len(liquidations)))
			resultsMu.Unlock()

			// Log progress
			if len(liquidations) > 0 {
				logger.Progress(int(current), int(totalControllers),
					"Controller %s: %d liquidations found",
					ctrl.Hex()[:10], len(liquidations))
			} else if current%20 == 0 || current == totalControllers {
				elapsed := time.Since(startTime)
				rate := float64(current) / elapsed.Seconds()
				eta := time.Duration(float64(totalControllers-current)/rate) * time.Second
				logger.Progress(int(current), int(totalControllers),
					"Processing... (%.1f/s, ETA: %s)", rate, eta.Round(time.Second))
			}
		}(controller)
	}

	wg.Wait()

	errCountFinal := atomic.LoadInt32(&errCount)
	if errCountFinal > 0 {
		logger.Warning("%d/%d controllers had errors", errCountFinal, totalControllers)
		return results, totalLiquidations, fmt.Errorf("%d liquidation fetches failed", errCountFinal)
	}

	return results, totalLiquidations, nil
}

func fetchHardLiquidationFast(client *ethclient.Client, controllerAddress common.Address, blockFrom, blockTo int64) ([]interfaces.LlamalendLiquidate, error) {
	liquidateEventTopic := common.HexToHash("0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0")

	// Calculate block ranges
	type blockRange struct {
		from, to int64
	}
	var ranges []blockRange
	for start := blockFrom; start <= blockTo; start += BLOCK_RANGE_SIZE {
		end := start + BLOCK_RANGE_SIZE - 1
		if end > blockTo {
			end = blockTo
		}
		ranges = append(ranges, blockRange{start, end})
	}

	// Process ranges in parallel with limited concurrency
	type rangeResult struct {
		liquidates []interfaces.LlamalendLiquidate
		err        error
	}
	results := make(chan rangeResult, len(ranges))

	sem := make(chan struct{}, 5) // Limit concurrent RPC calls per controller
	var wg sync.WaitGroup

	for _, r := range ranges {
		wg.Add(1)
		go func(br blockRange) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(br.from),
				ToBlock:   big.NewInt(br.to),
				Addresses: []common.Address{controllerAddress},
				Topics:    [][]common.Hash{{liquidateEventTopic}},
			}

			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				results <- rangeResult{err: err}
				return
			}

			if len(logs) == 0 {
				results <- rangeResult{}
				return
			}

			controllerContract, err := llamalendController.NewLlamalendController(controllerAddress, client)
			if err != nil {
				results <- rangeResult{err: err}
				return
			}

			var liquidates []interfaces.LlamalendLiquidate
			for _, vLog := range logs {
				event, err := controllerContract.ParseLiquidate(vLog)
				if err != nil {
					continue
				}

				// Get block timestamp
				block, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					continue
				}

				liquidates = append(liquidates, interfaces.LlamalendLiquidate{
					Timestamp: block.Time,
					User:      event.User,
				})
			}
			results <- rangeResult{liquidates: liquidates}
		}(r)
	}

	// Wait and close
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	var allLiquidates []interfaces.LlamalendLiquidate
	var lastErr error
	for res := range results {
		if res.err != nil {
			lastErr = res.err
			continue
		}
		allLiquidates = append(allLiquidates, res.liquidates...)
	}

	if lastErr != nil && len(allLiquidates) == 0 {
		return nil, lastErr
	}

	return allLiquidates, nil
}

// ============================================================================
// HELPERS
// ============================================================================

func getTokenPrice(tokenAddress common.Address, chainName string, timestamp uint64, curvePools []interfaces.CurvePool) float64 {
	tokenPrice := utils.GetHistoricalPriceTokenPrice(tokenAddress, chainName, timestamp)
	if tokenPrice == 0 {
		for _, pool := range curvePools {
			for _, coin := range pool.Coins {
				if strings.EqualFold(coin.Address, tokenAddress.Hex()) && coin.UsdPrice != nil {
					return coin.UsdPrice.(float64)
				}
			}
		}
	}
	return tokenPrice
}

// ============================================================================
// STORAGE
// ============================================================================

func readMarkets() []interfaces.LlamalendMarket {
	logger := newLogger("STORAGE")
	markets := make([]interfaces.LlamalendMarket, 0)

	logger.Info("Attempting to load from bucket: %s", BUCKET_LLAMALEND_PER_DAY_FILE)
	b, err := utils.ReadBucketFile(BUCKET_LLAMALEND_PER_DAY_FILE)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &markets); err != nil {
			logger.Warning("Failed to parse bucket file: %v", err)
		} else {
			logger.Success("Loaded %d markets from bucket (%.2f KB)", len(markets), float64(len(b))/1024)
			return markets
		}
	} else {
		logger.Info("Bucket file not available, trying local file...")
	}

	if !utils.FileExists(LLAMALEND_PER_DAY_PATH) {
		logger.Info("No existing data found - starting fresh")
		return markets
	}

	logger.Info("Loading from local file: %s", LLAMALEND_PER_DAY_PATH)
	file, err := os.ReadFile(LLAMALEND_PER_DAY_PATH)
	if err != nil {
		logger.Warning("Failed to read local file: %v", err)
		return markets
	}

	if err := json.Unmarshal(file, &markets); err != nil {
		logger.Warning("Failed to parse local file: %v", err)
	} else {
		logger.Success("Loaded %d markets from local file (%.2f KB)", len(markets), float64(len(file))/1024)
	}

	return markets
}

func writeMarkets(markets []interfaces.LlamalendMarket) {
	logger := newLogger("STORAGE")

	// Sort markets by timestamp (ascending), then by chainId, then by market id
	logger.Info("Sorting %d markets by date...", len(markets))
	sort.Slice(markets, func(i, j int) bool {
		if markets[i].Timestamp != markets[j].Timestamp {
			return markets[i].Timestamp < markets[j].Timestamp
		}
		if markets[i].ChainId != markets[j].ChainId {
			return markets[i].ChainId < markets[j].ChainId
		}
		return markets[i].Id < markets[j].Id
	})

	logger.Info("Marshaling %d markets...", len(markets))
	file, err := json.Marshal(markets)
	if err != nil {
		logger.Error("Failed to marshal markets: %v", err)
		return
	}
	logger.Info("Data size: %.2f KB", float64(len(file))/1024)

	logger.Info("Writing to local file: %s", LLAMALEND_PER_DAY_PATH)
	if err := os.WriteFile(LLAMALEND_PER_DAY_PATH, file, 0644); err != nil {
		logger.Error("Failed to write local file: %v", err)
	} else {
		logger.Success("Local file saved successfully")
	}

	logger.Info("Writing to bucket: %s", BUCKET_LLAMALEND_PER_DAY_FILE)
	if err := utils.WriteBucketFile(BUCKET_LLAMALEND_PER_DAY_FILE, markets); err != nil {
		logger.Warning("Failed to write bucket file: %v", err)
	} else {
		logger.Success("Bucket file saved successfully")
	}
}
