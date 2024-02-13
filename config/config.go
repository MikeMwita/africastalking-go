package config

import (
	"os"
)

type Config struct {
	Username string
	APIKey   string
	Env      string
	Port     string
}

func LoadConfig() *Config {
	return &Config{
		Username: "sandbox", // Hardcoded for testing
		APIKey:   os.Getenv("API_KEY"),
		Env:      os.Getenv("ENV"),
		Port:     os.Getenv("PORT"),
	}
}
