package config

import (
	"os"
	"strconv"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	LogLevel         string
	Port             int
	AIEngineAddr     string
	OrchestratorAddr string
	OpenAIAPIKey     string
	OpenAIEndpoint   string
	ArgoServerAddr   string
	MinIOEndpoint    string
	MinIOAccessKey   string
	MinIOSecretKey   string
	NessieEndpoint   string
	PrometheusAddr   string
}

// Load loads the configuration from environment variables.
func Load() (*AppConfig, error) {
	logLevel := getEnv("LOG_LEVEL", "info")
	portStr := getEnv("PORT", "8080")
	aiEngineAddr := getEnv("AI_ENGINE_ADDR", "localhost:8081")
	orchestratorAddr := getEnv("ORCHESTRATOR_ADDR", "localhost:8082")
	openAIAPIKey := getEnv("OPENAI_API_KEY", "")
	openAIEndpoint := getEnv("OPENAI_ENDPOINT", "https://api.openai.com/v1/chat/completions")
	argoServerAddr := getEnv("ARGO_SERVER_ADDR", "localhost:2746")
	minioEndpoint := getEnv("MINIO_ENDPOINT", "localhost:9000")
	minioAccessKey := getEnv("MINIO_ACCESS_KEY", "minioadmin")
	minioSecretKey := getEnv("MINIO_SECRET_KEY", "minioadmin")
	nessieEndpoint := getEnv("NESSIE_ENDPOINT", "http://localhost:19120/api/v1")
	prometheusAddr := getEnv("PROMETHEUS_ADDR", "http://localhost:9090") // Default Prometheus address

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		LogLevel:         logLevel,
		Port:             port,
		AIEngineAddr:     aiEngineAddr,
		OrchestratorAddr: orchestratorAddr,
		OpenAIAPIKey:     openAIAPIKey,
		OpenAIEndpoint:   openAIEndpoint,
		ArgoServerAddr:   argoServerAddr,
		MinIOEndpoint:    minioEndpoint,
		MinIOAccessKey:   minioAccessKey,
		MinIOSecretKey:   minioSecretKey,
		NessieEndpoint:   nessieEndpoint,
		PrometheusAddr:   prometheusAddr,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
