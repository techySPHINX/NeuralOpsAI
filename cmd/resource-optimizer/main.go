package main

import (
	"fmt"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"neuralops/api/proto/optimizer/v1"
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

	logger.Info("Starting Resource Optimizer...")

	// TODO: Initialize Prometheus client, Helm client, K8s client

	// Start gRPC server in a goroutine
	go func() {
		// Note: Using a different port for gRPC, e.g., 8084
		grpcAddr := "localhost:8084"
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Fatal("failed to listen", zap.Error(err))
		}
		grpcServer := grpc.NewServer()
		optimizerv1.RegisterOptimizerServiceServer(grpcServer, NewOptimizerGRPCServer(logger))
		logger.Info("gRPC server listening on", zap.String("addr", grpcAddr))
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
