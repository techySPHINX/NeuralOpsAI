package argo

import (
	"context"
	"fmt"
	"log"

	workflowpkg "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	wfClient versioned.Interface
	namespace string
}

func NewClient(namespace string) (*Client, error) {
	// Try to get in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to kubeconfig for local development
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		config, err = kubeconfig.ClientConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get Kubernetes config: %w", err)
		}
	}

	wfClient, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Argo Workflows client: %w", err)
	}

	return &Client{
		wfClient: wfClient,
		namespace: namespace,
	}, nil
}

func (c *Client) SubmitWorkflow(ctx context.Context, workflow *workflowpkg.Workflow) (*workflowpkg.Workflow, error) {
	log.Printf("Submitting workflow %s to namespace %s", workflow.Name, c.namespace)
	submittedWf, err := c.wfClient.ArgoprojV1alpha1().Workflows(c.namespace).Create(ctx, workflow, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to submit workflow: %w", err)
	}
	log.Printf("Workflow %s submitted successfully. Name: %s", workflow.Name, submittedWf.Name)
	return submittedWf, nil
}

// Placeholder for loading workflow from YAML
func LoadWorkflowFromYAML(yamlPath string) (*workflowpkg.Workflow, error) {
	// In a real scenario, you would read the YAML file and unmarshal it into a Workflow object.
	// For now, we'll return a dummy workflow.
	log.Printf("Loading workflow from YAML: %s (placeholder)", yamlPath)
	return &workflowpkg.Workflow{
		// Minimal dummy workflow for compilation
		TypeMeta: metav1.TypeMeta{
			APIVersion: "argoproj.io/v1alpha1",
			Kind:       "Workflow",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "dummy-workflow-",
		},
		Spec: workflowpkg.WorkflowSpec{
			Entrypoint: "hello",
			Templates: []workflowpkg.Template{
				{
					Name: "hello",
					Container: &workflowpkg.Container{
						Image:   "alpine/git",
						Command: []string{"echo"},
						Args:    []string{"Hello from dummy workflow!"},
					},
				},
			},
		},
	},
	 nil
}
