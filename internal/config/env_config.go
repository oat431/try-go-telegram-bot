package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Loading env file")
}

func LoadEnvConfig() {
	failedToLoadEnv := godotenv.Load(".env")
	if failedToLoadEnv != nil {
		log.Printf("Failed to load env file: %v", failedToLoadEnv)
	}
}
