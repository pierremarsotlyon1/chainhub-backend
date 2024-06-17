package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/escrow"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	LOCKS_PATH      = "./data/locks.json"
	STATS_LOCK_PATH = "./data/stats-locks.json"
	locks_config    = "./data/configs/locks-config.json"
)

func FetchLocks(client *ethclient.Client, currentBlock uint64) {

	fmt.Println("Fetching locks")

	config := utils.ReadConfig(locks_config)

	locks := readLocks()
	locks = append(locks, fetchVeCRVLocks(client, currentBlock, config)...)

	writeLocks(locks)
	computeLocksStats(locks)

	// Write config
	utils.WriteConfig(config, currentBlock, locks_config)
}

func fetchVeCRVLocks(client *ethclient.Client, currentBlock uint64, config interfaces.Config) []interfaces.Lock {

	from := config.LastBlock
	if from == 0 {
		from = 10647812
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{utils.CURVE_ESCROW_ADDRESS},
		Topics:    [][]common.Hash{{common.HexToHash("0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	locks := make([]interfaces.Lock, 0)

	abi, err := escrow.EscrowMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	for _, vLog := range logs {
		receivedMap := map[string]interface{}{}
		if err := abi.UnpackIntoMap(receivedMap, "Deposit", vLog.Data); err != nil {
			fmt.Println(err)
			continue
		}

		value := receivedMap["value"].(*big.Int)
		if value.Cmp(big.NewInt(0)) != 1 {
			continue
		}

		user := receivedMap["provider"].(common.Address)

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		timestamp := block.Time()

		locks = append(locks, interfaces.Lock{
			Tx:        vLog.TxHash,
			Timestamp: timestamp,
			Value:     value,
			User:      user,
		})
	}

	return locks
}

func historicalLocks() {
	file, err := os.ReadFile("./locks.json")
	if err != nil {
		log.Fatal(err)
	}

	data := make([][]string, 0)
	if err := json.Unmarshal([]byte(file), &data); err != nil {
		log.Fatal(err)
	}

	locks := data

	responses := make([]interfaces.Lock, 0)
	for _, lock := range locks {

		timestamp, _ := strconv.Atoi(lock[0])
		value, _ := big.NewInt(0).SetString(lock[2], 10)

		responses = append(responses, interfaces.Lock{
			Timestamp: uint64(timestamp),
			Value:     value,
			User:      common.HexToAddress(lock[1]),
		})
	}

	writeLocks(responses)
}

func computeLocksStats(locks []interfaces.Lock) {

	locksPerPeriod := make(map[uint64]*big.Int)
	wrappersLocks := make(map[uint64]map[common.Address]*big.Int)

	for _, lock := range locks {
		currentPeriod := uint64(lock.Timestamp/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC

		value, exists := locksPerPeriod[currentPeriod]
		if !exists {
			locksPerPeriod[currentPeriod] = big.NewInt(0)
			value = locksPerPeriod[currentPeriod]
		}

		// Init wrapper data
		_, exists = wrappersLocks[currentPeriod]
		if !exists {
			wrappersLocks[currentPeriod] = make(map[common.Address]*big.Int)
			for _, wrapper := range WRAPPERS {
				wrappersLocks[currentPeriod][wrapper.Address] = big.NewInt(0)
			}
		}

		locksPerPeriod[currentPeriod] = new(big.Int).Add(value, lock.Value)

		for _, wrapper := range WRAPPERS {
			if strings.EqualFold(wrapper.Address.Hex(), lock.User.Hex()) {
				wrappersLocks[currentPeriod][wrapper.Address] = new(big.Int).Add(wrappersLocks[currentPeriod][wrapper.Address], lock.Value)
				break
			}
		}
	}

	statsLock := make([]interfaces.StatsLock, 0)
	for period, lockAmount := range locksPerPeriod {
		statLock := interfaces.StatsLock{
			Amount:    utils.Quo(lockAmount, 18),
			Timestamp: period,
		}

		for _, wrapper := range WRAPPERS {
			switch wrapper.Address.Hex() {
			case common.HexToAddress("0x52f541764E6e90eeBc5c21Ff570De0e2D63766B6").Hex():
				statLock.StakeDAO = utils.Quo(wrappersLocks[period][wrapper.Address], 18)
			case common.HexToAddress("0xF147b8125d2ef93FB6965Db97D6746952a133934").Hex():
				statLock.Yearn = utils.Quo(wrappersLocks[period][wrapper.Address], 18)
			case common.HexToAddress("0x989AEb4d175e16225E39E87d0D97A3360524AD80").Hex():
				statLock.Convex = utils.Quo(wrappersLocks[period][wrapper.Address], 18)
			}
		}
		statsLock = append(statsLock, statLock)
	}

	sort.Slice(statsLock, func(i, j int) bool { return statsLock[i].Timestamp < statsLock[j].Timestamp })

	writeStatsLocks(statsLock)
}

func writeStatsLocks(statsLock []interfaces.StatsLock) {
	file, err := json.Marshal(statsLock)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(STATS_LOCK_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readLocks() []interfaces.Lock {

	if !utils.FileExists(LOCKS_PATH) {
		return make([]interfaces.Lock, 0)
	}

	file, err := os.ReadFile(LOCKS_PATH)
	if err != nil {
		log.Fatal(err)
	}

	locks := make([]interfaces.Lock, 0)
	if err := json.Unmarshal([]byte(file), &locks); err != nil {
		log.Fatal(err)
	}

	return locks
}

func writeLocks(locks []interfaces.Lock) {
	file, err := json.Marshal(locks)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(LOCKS_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}
