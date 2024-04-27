package utils

import (
	"encoding/json"
	"log"
	"main/interfaces"
	"os"
	"time"
)

func ReadConfig(path string) interfaces.Config {

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

	var config interfaces.Config
	if err := json.Unmarshal([]byte(file), &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func WriteConfig(config interfaces.Config, currentBlock uint64, path string) {
	config.LastBlock = currentBlock
	config.LastUpdate = uint64(time.Now().Unix())

	file, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}
}
