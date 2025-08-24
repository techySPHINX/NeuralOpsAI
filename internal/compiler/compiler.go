package compiler

import (
	"neuralops/api/proto/orchestrator/v1"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CompileToArgo(req *orchestratorv1.SubmitPipelineRequest) (*wfv1.Workflow, error) {
	plan := req.Plan
	taskCode := req.TaskCode

	tasks := make([]wfv1.DAGTask, len(plan.Tasks))
	for i, task := range plan.Tasks {
		if code, ok := taskCode[task.Name]; ok {
			// This task has generated code, use the workflow-runner
			tasks[i] = wfv1.DAGTask{
				Name:         task.Name,
				Template:     "workflow-runner",
				Dependencies: task.DependsOn,
				Arguments: wfv1.Arguments{
					Artifacts: []wfv1.Artifact{
						{
							Name: "plugin-code",
							Path: "/app/plugin.go",
							ArtifactLocation: wfv1.ArtifactLocation{
								Raw: &wfv1.RawArtifact{
									Data: code,
								},
							},
						},
					},
				},
			}
		} else {
			// This task does not have generated code, use the old template
			tasks[i] = wfv1.DAGTask{
				Name:         task.Name,
				Template:     "main-container",
				Dependencies: task.DependsOn,
				Arguments: wfv1.Arguments{
					Parameters: []wfv1.Parameter{
						{Name: "message", Value: wfv1.AnyStringPtr(task.Description)},
					},
				},
			}
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
					Name: "workflow-runner",
					Inputs: wfv1.Inputs{
						Artifacts: []wfv1.Artifact{
							{
								Name: "plugin-code",
								Path: "/app/plugin.go",
							},
						},
					},
					Container: &wfv1.Container{
						Image: "neuralops-workflow-runner:latest",
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
