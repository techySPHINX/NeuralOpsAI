package config

import (
	"os"
	"strconv"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	LogLevel string
	Port     int
}

// Load loads the configuration from environment variables.
func Load() (*AppConfig, error) {
	logLevel := getEnv("LOG_LEVEL", "info")
	portStr := getEnv("PORT", "8080")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		LogLevel: logLevel,
		Port:     port,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
