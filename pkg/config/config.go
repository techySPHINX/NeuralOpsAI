package config

import (
	"os"
	"strconv"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	LogLevel      string
	Port          int
	AIEngineAddr string
}

// Load loads the configuration from environment variables.
func Load() (*AppConfig, error) {
	logLevel := getEnv("LOG_LEVEL", "info")
	portStr := getEnv("PORT", "8080")
	aiEngineAddr := getEnv("AI_ENGINE_ADDR", "localhost:8081") // Default for local dev

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		LogLevel:      logLevel,
		Port:          port,
		AIEngineAddr: aiEngineAddr,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
