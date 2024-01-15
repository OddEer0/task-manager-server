package config

import (
	"github.com/joho/godotenv"
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

func NewConfig() (*Config, error) {
	if instance != nil {
		return instance, nil
	}

	err := godotenv.Load("./config/.local.env")
	if err != nil {
		return nil, err
	}

	instance = &Config{
		Host:             os.Getenv("HOST"),
		StoragePath:      os.Getenv("STORAGE_PATH"),
		ApiKey:           os.Getenv("API_KEY"),
		AccessTokenTime:  os.Getenv("ACCESS_TOKEN"),
		RefreshTokenTime: os.Getenv("REFRESH_TOKEN"),
	}

	return instance, nil
}

func NewConfigTest() *Config {
	return &Config{
		Host:             "localhost:5000",
		StoragePath:      "./storage.db",
		ApiKey:           "supper-secret-key",
		AccessTokenTime:  "3m",
		RefreshTokenTime: "3m",
	}
}
