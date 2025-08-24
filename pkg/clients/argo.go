package clients

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-workflows/v3/pkg/apiclient/workflow"
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ArgoClient struct {
	client workflow.WorkflowServiceClient
}

func NewArgoClient(ctx context.Context, addr string) (*ArgoClient, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to argo server: %w", err)
	}

	return &ArgoClient{
		client: workflow.NewWorkflowServiceClient(conn),
	}, nil
}

func (c *ArgoClient) SubmitWorkflow(ctx context.Context, wf *v1alpha1.Workflow) (*v1alpha1.Workflow, error) {
	return c.client.CreateWorkflow(ctx, &workflow.WorkflowCreateRequest{
		Namespace: "default", // Or from config
		Workflow:  wf,
	})
}
