package adapters

import (
	"context"
	"neuralops/api/proto/ai_engine/v1"
)

// Adapter is the interface for a large language model adapter.
type Adapter interface {
	GeneratePlan(ctx context.Context, query string) (*ai_enginev1.PipelinePlan, error)
}
