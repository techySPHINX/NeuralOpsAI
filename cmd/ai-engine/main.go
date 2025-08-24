package main

import (
	"fmt"
	"net/http"

	"neuralops/pkg/config"
	"neuralops/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	logger, err := logging.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %v", err))
	}

	logger.Info("Starting AI Engine...")

	server := NewServer(logger)

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("Server listening on", zap.String("addr", addr))

	if err := http.ListenAndServe(addr, server); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
