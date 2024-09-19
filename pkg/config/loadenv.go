package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() error {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err) // Log a fatal error if .env loading fails
	}
	return nil
}
