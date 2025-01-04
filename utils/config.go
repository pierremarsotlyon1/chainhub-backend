package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"main/interfaces"
	"os"
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
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &config); err != nil {
			log.Fatal(err)
		}
		return config
	}

	// From GitHub
	if !FileExists(path) {
		return interfaces.Config{
			LastBlock:  0,
			LastUpdate: 0,
		}
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func WriteConfig(config interfaces.Config, currentBlock uint64, path string) {
	config.LastBlock = currentBlock
	config.LastUpdate = uint64(time.Now().Unix())

	// Write in GitHub
	file, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Write in Google Cloud
	bucketPath := path
	if bucketPath[0] == '.' {
		bucketPath = bucketPath[2:]
	}
	if err := WriteBucketFile(bucketPath, config); err != nil {
		fmt.Println(err)
	}
}
