package main

import (
	"context"

	"neuralops/api/proto/ai_engine/v1"
	"neuralops/llm/adapters"
	"neuralops/pkg/logging"
)

type AIEngineGRPCServer struct {
	ai_enginev1.UnimplementedAIEngineServiceServer
	logger  *logging.Logger
	adapter adapters.Adapter
}

func NewAIEngineGRPCServer(logger *logging.Logger, adapter adapters.Adapter) *AIEngineGRPCServer {
	return &AIEngineGRPCServer{
		logger:  logger,
		adapter: adapter,
	}
}

func (s *AIEngineGRPCServer) Plan(ctx context.Context, req *ai_enginev1.PlanRequest) (*ai_enginev1.PlanResponse, error) {
	s.logger.Info("Received Plan request", "query", req.Query)

	plan, err := s.adapter.GeneratePlan(ctx, req.Query)
	if err != nil {
		s.logger.Error("failed to generate plan", "error", err)
		return nil, err
	}

	return &ai_enginev1.PlanResponse{Plan: plan}, nil
}
