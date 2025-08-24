package main

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"neuralops/api/proto/ai_engine/v1"
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

	logger.Info("Starting API Gateway...")

	// Set up gRPC client connection to AI Engine
	conn, err := grpc.Dial(cfg.AIEngineAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("failed to connect to AI engine", zap.Error(err))
	}
	defer conn.Close()
	aiClient := ai_enginev1.NewAIEngineServiceClient(conn)

	server := NewServer(logger, aiClient)

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("Server listening on", zap.String("addr", addr))

	if err := http.ListenAndServe(addr, server); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
