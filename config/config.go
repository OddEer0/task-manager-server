package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Env              string
		Host             string
		StoragePath      string
		ApiKey           string
		AccessTokenTime  string
		RefreshTokenTime string
		AbsPath          string
	}
)

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
		env = "local"
	case "test":
		cfgFilePath = filepath.Join(dir, "./config/.test.env")
		env = "local"
	}

	err := godotenv.Load(cfgFilePath)
	if err != nil {
		return nil, err
	}

	instance = &Config{
		Env:              env,
		AbsPath:          dir,
		Host:             os.Getenv("HOST"),
		StoragePath:      os.Getenv("STORAGE_PATH"),
		ApiKey:           os.Getenv("API_KEY"),
		AccessTokenTime:  os.Getenv("ACCESS_TOKEN"),
		RefreshTokenTime: os.Getenv("REFRESH_TOKEN"),
	}

	return instance, nil
}
