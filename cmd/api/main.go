package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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
	// Placeholder for Argo Workflow triggering logic
	fmt.Fprintf(w, "Triggering Argo Workflow (placeholder)...")
}
