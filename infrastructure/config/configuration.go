package config

import (
	"log"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

func LoadEnv() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func ConvertToInt(value string) int {
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return valueInt
}

func ConvertToBool(value string) bool {
	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatal(err)
	}
	return valueBool
}
