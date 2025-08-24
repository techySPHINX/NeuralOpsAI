# NeuralOpsAI

NeuralOpsAI is a powerful and flexible MLOps platform designed to streamline the entire machine learning lifecycle, from data ingestion and processing to model training, deployment, and monitoring. It provides a unified and automated framework for managing complex AI/ML workflows, enabling teams to build, scale, and operate machine learning systems with greater efficiency and reliability.

## Features

*   **Natural Language-based Pipeline Creation:** Define and create complex data and machine learning pipelines using simple natural language queries.
*   **Automated Workflow Orchestration:** Leverages Argo Workflows and Temporal to provide robust and scalable orchestration of ML pipelines.
*   **Microservices Architecture:** A modular and scalable architecture based on gRPC and RESTful APIs, with distinct services for different functionalities.
*   **Cloud-Native and Kubernetes-Native:** Designed to run seamlessly on Kubernetes, enabling easy scaling and management of resources.
*   **Extensible and Pluggable:** Easily integrate with various data sources, machine learning frameworks, and tools.
*   **API-First Design:** A comprehensive API for programmatic control and integration with other systems.

## Architecture Overview

NeuralOpsAI follows a microservices-based architecture, leveraging a suite of powerful open-source technologies to deliver a robust and scalable MLOps platform. The key components and technologies are:

*   **API Gateway:** The single entry point for all API requests, routing them to the appropriate backend services.
*   **Orchestrator:** The core component responsible for parsing natural language queries, compiling them into executable workflows, and managing the overall pipeline lifecycle.
*   **AI Engine:** Handles the machine learning-specific tasks, such as model training, evaluation, and serving.
*   **Iceberg Manager:** Manages the data lake, powered by Apache Iceberg, for efficient and reliable data storage and access.
*   **Incident Responder:** Monitors the system for anomalies and incidents, and triggers automated responses.
*   **Resource Optimizer:** Optimizes the allocation and utilization of computing resources for ML workloads.
*   **Workflow Runner:** Executes the individual tasks within a pipeline, using Argo Workflows as the underlying execution engine.

### Core Technologies

*   **[Apache Iceberg](https://iceberg.apache.org/):** Used as the open table format for the data lake, providing fast and reliable access to large datasets.
*   **[Argo Workflows](https://argoproj.github.io/argo-workflows/):** The container-native workflow engine for orchestrating parallel jobs on Kubernetes.
*   **[Minikube](https://minikube.sigs.k8s.io/docs/start/):** Used for running a local Kubernetes cluster for development and testing.
*   **[Temporal](https://temporal.io/):** Provides a durable and reliable platform for orchestrating and executing asynchronous tasks and workflows.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.21 or later)
*   [Docker](https://docs.docker.com/get-docker/)
*   [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
*   [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
*   [Buf](https://buf.build/docs/installation)

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/techySPHINX/NeuralOpsAI
    cd neuralops
    ```

2.  **Install development tools:**

    ```sh
    make dev-tools
    ```

3.  **Start a local Kubernetes cluster:**

    ```sh
    make kind-up
    ```

4.  **Start the Temporal server:**

    ```sh
    make temporal-up
    ```

## Building and Running

### Build the binaries

To build the Go binaries for all the services, run:

```sh
make build
```

### Build the Docker images

To build the Docker images for all the services, run:

```sh
make docker
```

### Run the API Gateway

To run the API Gateway locally, run:

```sh
make run-gateway
```

## API

The NeuralOps API is defined using OpenAPI 3.0. The API specification can be found in `api/openapi/gateway.yaml`.

### Endpoints

*   `POST /v1/pipelines:nl`: Create a pipeline from a natural language query.
*   `POST /v1/pipelines/{id}:run`: Run a pipeline.
*   `GET /v1/runs/{id}`: Get the status of a run.
*   `GET /healthz`: Health check.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
