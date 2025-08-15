package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"neuralops/internal/argo"
)

var argoClient *argo.Client

func main() {
	var err error
	argoClient, err = argo.NewClient("argo") // Assuming Argo Workflows is in 'argo' namespace
	if err != nil {
		log.Fatalf("Failed to create Argo client: %v", err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/trigger-workflow", triggerWorkflowHandler)
	port := "8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from NeuralOps API Gateway!")
}

func triggerWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	// In a real scenario, you would load the workflow from a YAML file.
	// For now, we are using the dummy workflow from internal/argo/client.go
	workflow, err := argo.LoadWorkflowFromYAML("workflows/hello-world-workflow.yaml") // Placeholder path
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load workflow: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = argoClient.SubmitWorkflow(context.Background(), workflow)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to submit workflow: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Argo Workflow submitted successfully!")
}
