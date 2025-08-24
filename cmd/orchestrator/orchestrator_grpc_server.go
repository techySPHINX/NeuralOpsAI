package main

import (
	"context"

	"neuralops/api/proto/orchestrator/v1"
	"neuralops/internal/compiler"
	"neuralops/pkg/clients"
	"neuralops/pkg/logging"
)

type OrchestratorGRPCServer struct {
	orchestratorv1.UnimplementedOrchestratorServiceServer
	logger     *logging.Logger
	argoClient *clients.ArgoClient
}

func NewOrchestratorGRPCServer(logger *logging.Logger, argoClient *clients.ArgoClient) *OrchestratorGRPCServer {
	return &OrchestratorGRPCServer{
		logger:     logger,
		argoClient: argoClient,
	}
}

func (s *OrchestratorGRPCServer) SubmitPipeline(ctx context.Context, req *orchestratorv1.SubmitPipelineRequest) (*orchestratorv1.SubmitPipelineResponse, error) {
	s.logger.Info("Received SubmitPipeline request", "plan_id", req.Plan.Id)

	argoWF, err := compiler.CompileToArgo(req) // Pass the whole request
	if err != nil {
		s.logger.Error("failed to compile plan to Argo workflow", "error", err)
		return nil, err
	}

	submittedWF, err := s.argoClient.SubmitWorkflow(ctx, argoWF)
	if err != nil {
		s.logger.Error("failed to submit workflow to Argo", "error", err)
		return nil, err
	}

	s.logger.Info("Successfully submitted workflow to Argo", "workflow_name", submittedWF.Name)

	return &orchestratorv1.SubmitPipelineResponse{RunId: submittedWF.Name}, nil
}
