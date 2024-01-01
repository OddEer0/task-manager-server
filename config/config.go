package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host             string
	StoragePath      string
	ApiKey           string
	AccessTokenTime  string
	RefreshTokenTime string
}

var instance *Config = nil

func MustLoad() *Config {
	if instance != nil {
		return instance
	}

	err := godotenv.Load("./config/.local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	instance = &Config{
		Host:             os.Getenv("HOST"),
		StoragePath:      os.Getenv("STORAGE_PATH"),
		ApiKey:           os.Getenv("API_KEY"),
		AccessTokenTime:  os.Getenv("ACCESS_TOKEN"),
		RefreshTokenTime: os.Getenv("REFRESH_TOKEN"),
	}

	return instance
}
