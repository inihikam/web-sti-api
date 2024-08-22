package config

import (
	"os"
)

type Config struct {
	APIServerURL string
	LogLevel     string
}

func LoadConfig() *Config {
	return &Config{
		APIServerURL: getEnv("API_SERVER_URL", "http://localhost:8080"),
		LogLevel:     getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	
	return value
}