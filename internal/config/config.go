package config

import (
	"os"
)

type Config struct {
	AppPort string
	AppEnv  string
}

func NewConfig() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
	}
}

// Helper to read environment variable or fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
