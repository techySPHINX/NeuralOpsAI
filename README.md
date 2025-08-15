# NeuralOps

NeuralOps is a backend-first, cloud-native orchestration platform that automates data ingestion, transformation, and deployment workflows inside Kubernetes.

## Project Structure

- `cmd/api`: Contains the main Go application for the API Gateway.
- `Dockerfile`: Defines how to build the Docker image for the API Gateway.
- `k8s`: Contains Kubernetes manifests for deploying NeuralOps components.

## Getting Started (API Gateway)

1.  **Build the Docker Image:**
    ```bash
    docker build -t neuralops-api-gateway:latest .
    ```

2.  **Deploy to Kubernetes:**
    ```bash
    kubectl apply -f k8s/api-deployment.yaml
    kubectl apply -f k8s/api-service.yaml
    ```

3.  **Access the API Gateway:**
    If using `minikube` or `kind`:
    ```bash
    minikube service neuralops-api-gateway --url
    # Or for kind, use port-forwarding:
    # kubectl port-forward service/neuralops-api-gateway 8080:80
    ```
    Then open the URL in your browser or use `curl`.
