package utils

import (
	"slices"

	"github.com/ethereum/go-ethereum/common"
)

func ChunkAddressSlice(slice []common.Address, chunkSize int) [][]common.Address {
	var chunks [][]common.Address
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	// Ajoute les éléments restants s'ils existent
	if len(slice) > 0 {
		chunks = append(chunks, slice)
	}
	return chunks
}

func ArrayContains(s []string, str string) bool {
	return slices.Contains(s, str)
}
