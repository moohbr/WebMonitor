package infrastructure

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

// LoadEnv loads the environment variables
func LoadEnv() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv gets an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
