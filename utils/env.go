package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	godotenv.Load()

	return os.Getenv(key)
}
