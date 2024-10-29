package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadConfig() Config {
	err := godotenv.Load()

	var config Config

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	databaseConfig := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	config.Database = databaseConfig

	return config
}
