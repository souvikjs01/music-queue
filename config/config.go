package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load env file ", err)
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
