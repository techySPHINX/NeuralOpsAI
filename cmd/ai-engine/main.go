package main

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
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

	logger.Info("Starting AI Engine...")

	// Start gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", cfg.AIEngineAddr)
		if err != nil {
			logger.Fatal("failed to listen", zap.Error(err))
		}
		grpcServer := grpc.NewServer()
		ai_enginev1.RegisterAIEngineServiceServer(grpcServer, NewAIEngineGRPCServer(logger))
		logger.Info("gRPC server listening on", zap.String("addr", cfg.AIEngineAddr))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("failed to serve gRPC", zap.Error(err))
		}
	}()

	// Start HTTP server for health checks
	server := NewServer(logger)
	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("HTTP server listening on", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, server); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
