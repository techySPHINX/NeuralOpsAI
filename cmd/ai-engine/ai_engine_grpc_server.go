package main

import (
	"context"

	"neuralops/api/proto/ai_engine/v1"
	"neuralops/pkg/logging"
)

type AIEngineGRPCServer struct {
	ai_enginev1.UnimplementedAIEngineServiceServer
	logger *logging.Logger
}

func NewAIEngineGRPCServer(logger *logging.Logger) *AIEngineGRPCServer {
	return &AIEngineGRPCServer{logger: logger}
}

func (s *AIEngineGRPCServer) Plan(ctx context.Context, req *ai_enginev1.PlanRequest) (*ai_enginev1.PlanResponse, error) {
	s.logger.Info("Received Plan request", "query", req.Query)

	// Dummy plan generation
	plan := &ai_enginev1.PipelinePlan{
		Id:          "plan-123",
		Description: "A dummy pipeline plan for query: " + req.Query,
		Tasks: []*ai_enginev1.Task{
			{
				Name:        "ingest-data",
				Description: "Ingest data from a source",
				Type:        "ingest",
				Config: map[string]string{
					"source": "s3://my-bucket/data.csv",
				},
			},
			{
				Name:        "transform-data",
				Description: "Transform the data",
				Type:        "transform",
				DependsOn:   []string{"ingest-data"},
				Config: map[string]string{
					"script": "SELECT * FROM input",
				},
			},
		},
	}

	return &ai_enginev1.PlanResponse{Plan: plan}, nil
}
