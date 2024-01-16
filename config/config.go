package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
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
	dir := os.Getenv("ABS_PATH")

	env := os.Getenv("ENV")
	var cfgFilePath string

	switch env {
	case "":
		cfgFilePath = filepath.Join(dir, "./config/.local.env")
	case "test":
		cfgFilePath = filepath.Join(dir, "./config/.test.env")
	}

	fmt.Print("cfg file path: ", cfgFilePath, "\n")

	err := godotenv.Load(cfgFilePath)
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
