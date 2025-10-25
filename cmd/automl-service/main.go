package main

import (
	"fmt"
	"net"
	"net/http"

	"neuralops/pkg/config"
	"neuralops/pkg/logging"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	logger.Info("Starting AutoML Service...")

	// Initialize AutoML engine
	// TODO: Initialize Ray Tune or Optuna for hyperparameter tuning
	
	// Start gRPC server in a goroutine
	go func() {
		grpcAddr := "localhost:8085"
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Fatal("failed to listen", zap.Error(err))
		}
		grpcServer := grpc.NewServer()
		// TODO: Register AutoML service
		// automlv1.RegisterAutoMLServiceServer(grpcServer, NewAutoMLGRPCServer(logger, engine))
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

type Server struct {
	logger *zap.Logger
}

func NewServer(logger *zap.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}
	http.NotFound(w, r)
}
