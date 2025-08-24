package compiler

import (
	"neuralops/api/proto/ai_engine/v1"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CompileToArgo(plan *ai_enginev1.PipelinePlan) (*wfv1.Workflow, error) {
	tasks := make([]wfv1.DAGTask, len(plan.Tasks))
	for i, task := range plan.Tasks {
		tasks[i] = wfv1.DAGTask{
			Name:         task.Name,
			Template:     "main-container", // A generic template for now
			Dependencies: task.DependsOn,
			Arguments: wfv1.Arguments{
				Parameters: []wfv1.Parameter{
					{Name: "message", Value: wfv1.AnyStringPtr(task.Description)},
				},
			},
		}
	}

	wf := &wfv1.Workflow{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: plan.Id + "-",
		},
		Spec: wfv1.WorkflowSpec{
			Entrypoint: "main",
			Templates: []wfv1.Template{
				{
					Name: "main",
					DAG: &wfv1.DAGTemplate{
						Tasks: tasks,
					},
				},
				{
					Name: "main-container",
					Inputs: wfv1.Inputs{
						Parameters: []wfv1.Parameter{
							{Name: "message"},
						},
					},
					Container: &wfv1.Container{
						Image:   "docker/whalesay:latest",
						Command: []string{"cowsay"},
						Args:    []string{"{{inputs.parameters.message}}"},
					},
				},
			},
		},
	}

	return wf, nil
}
