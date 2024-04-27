package utils

import "os"

func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}

	return true
}
