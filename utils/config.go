package utils

import (
	"encoding/json"
	"log"
	"main/interfaces"
	"time"
)

func ReadConfig(path string) interfaces.Config {

	var config interfaces.Config

	// From Google Cloud
	bucketFile := path
	if bucketFile[0] == '.' {
		bucketFile = bucketFile[2:]
	}
	b, err := ReadBucketFile(bucketFile)
	if err != nil {
		return interfaces.Config{
			LastBlock:  0,
			LastUpdate: 0,
		}
	}

	if err := json.Unmarshal(b, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func WriteConfig(config interfaces.Config, currentBlock uint64, path string) {
	config.LastBlock = currentBlock
	config.LastUpdate = uint64(time.Now().Unix())

	// Write in Google Cloud
	bucketPath := path
	if bucketPath[0] == '.' {
		bucketPath = bucketPath[2:]
	}
	if err := WriteBucketFile(bucketPath, config); err != nil {
		log.Fatal(err)
	}
}
