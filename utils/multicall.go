package utils

import (
	"context"
	"fmt"
	"main/contracts/multicall"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MulticallResponse struct {
	Success    bool
	ReturnData []byte
}

func Multicall(client *ethclient.Client, calls []multicall.Multicall3Call, multicallAddress common.Address, blockNumber *big.Int) []MulticallResponse {

	parsedABI, err := multicall.MulticallMetaData.GetAbi()
	if err != nil {
		fmt.Println("Multicall parsedABI", err)
		return nil
	}

	chunks := ChunkMulticall(calls, 20)

	responses := make([]MulticallResponse, 0)

	for _, chunk := range chunks {
		packedData, err := parsedABI.Pack("tryAggregate", false, chunk)
		if err != nil {
			fmt.Println("aggregate", err)
			return nil
		}

		msg := ethereum.CallMsg{
			To:   &multicallAddress,
			Data: packedData,
			Gas:  200_000_000,
		}

		result, err := client.CallContract(context.Background(), msg, blockNumber)
		if err != nil {
			fmt.Println("CallContract", err)
			return nil
		}

		// Dépactage des résultats
		var response []MulticallResponse

		err = parsedABI.UnpackIntoInterface(&response, "tryAggregate", result)
		if err != nil {
			fmt.Println("aggregate unpack", err)
			return nil
		}

		responses = append(responses, response...)
	}

	return responses
}

func ChunkMulticall(items []multicall.Multicall3Call, chunkSize int) (chunks [][]multicall.Multicall3Call) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
