package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	DbName string
	DbUser string
	DbPass string
	DbPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		Host:   os.Getenv("HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbPort: os.Getenv("DB_PORT"),
	}

	return config
}
