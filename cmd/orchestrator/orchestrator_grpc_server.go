package main

import (
	"context"

	"github.com/google/uuid"
	"neuralops/api/proto/orchestrator/v1"
	"neuralops/pkg/logging"
)

type OrchestratorGRPCServer struct {
	orchestratorv1.UnimplementedOrchestratorServiceServer
	logger *logging.Logger
}

func NewOrchestratorGRPCServer(logger *logging.Logger) *OrchestratorGRPCServer {
	return &OrchestratorGRPCServer{logger: logger}
}

func (s *OrchestratorGRPCServer) SubmitPipeline(ctx context.Context, req *orchestratorv1.SubmitPipelineRequest) (*orchestratorv1.SubmitPipelineResponse, error) {
	s.logger.Info("Received SubmitPipeline request", "plan_id", req.Plan.Id)

	// In a real implementation, we would compile the plan and submit it.
	runID := uuid.New().String()

	return &orchestratorv1.SubmitPipelineResponse{RunId: runID}, nil
}
