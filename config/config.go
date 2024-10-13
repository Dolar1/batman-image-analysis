package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var appConfig Config

type Config struct {
	DBUrl      string
	ServerPort string
}

// LoadConfig loads the configuration from the .env file.
func Load() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		DBUrl:      os.Getenv("DB_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	// Check for missing configurations
	if config.DBUrl == "" {
		log.Fatal("DB_URL is not set in .env file")
	}
	if config.ServerPort == "" {
		log.Fatal("SERVER_PORT is not set in .env file")
	}

	appConfig = *config
	return config, nil
}

func Port() string {
	return appConfig.ServerPort
}

func DBUrl() string {
	return appConfig.DBUrl
}
