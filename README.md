# NeuralOpsAI

NeuralOpsAI is a powerful and flexible MLOps platform designed to streamline the entire machine learning lifecycle, from data ingestion and processing to model training, deployment, and monitoring. It provides a unified and automated framework for managing complex AI/ML workflows, enabling teams to build, scale, and operate machine learning systems with greater efficiency and reliability.

## 🎉 New Features (v2.0)

### 1. **Web Dashboard** 🖥️

A modern, responsive web interface for managing your entire MLOps platform:

- Real-time pipeline monitoring
- Visual pipeline builder with natural language support
- Interactive workflow DAG visualization
- System health metrics and analytics

### 2. **Model Registry** 📦

Comprehensive model versioning and lifecycle management:

- Version control for ML models
- Stage-based promotion (dev → staging → production)
- Metrics and parameter tracking
- Framework-agnostic design

### 3. **AutoML Integration** 🤖

Automated machine learning with intelligent optimization:

- Automated model selection
- Hyperparameter tuning (Bayesian, Grid, Random search)
- Multi-algorithm comparison
- Best model recommendation

### 4. **Enhanced Data Versioning** 📊

Advanced data lineage and versioning capabilities:

- Iceberg-based snapshot versioning
- Schema evolution tracking
- Data lineage visualization
- Version comparison and time travel

### 5. **Workflow Visualization** 🔄

Real-time workflow execution monitoring:

- Interactive DAG visualization
- Node-level status tracking
- Live log streaming
- Performance metrics

See [IMPROVEMENTS.md](IMPROVEMENTS.md) for detailed documentation.

## Features

- **Natural Language-based Pipeline Creation:** Define and create complex data and machine learning pipelines using simple natural language queries.
- **Automated Workflow Orchestration:** Leverages Argo Workflows and Temporal to provide robust and scalable orchestration of ML pipelines.
- **Microservices Architecture:** A modular and scalable architecture based on gRPC and RESTful APIs, with distinct services for different functionalities.
- **Cloud-Native and Kubernetes-Native:** Designed to run seamlessly on Kubernetes, enabling easy scaling and management of resources.
- **Extensible and Pluggable:** Easily integrate with various data sources, machine learning frameworks, and tools.
- **API-First Design:** A comprehensive API for programmatic control and integration with other systems.

## Architecture Overview

NeuralOpsAI follows a microservices-based architecture, leveraging a suite of powerful open-source technologies to deliver a robust and scalable MLOps platform. The key components and technologies are:

- **Web Dashboard:** Modern React-based UI for visual pipeline creation, monitoring, and management.
- **API Gateway:** The single entry point for all API requests, routing them to the appropriate backend services.
- **Orchestrator:** The core component responsible for parsing natural language queries, compiling them into executable workflows, and managing the overall pipeline lifecycle.
- **AI Engine:** Handles the machine learning-specific tasks, such as model training, evaluation, and serving.
- **Iceberg Manager:** Manages the data lake, powered by Apache Iceberg, for efficient and reliable data storage and access with versioning.
- **Model Registry:** Manages ML model versions, metadata, and lifecycle across environments.
- **AutoML Service:** Provides automated model selection and hyperparameter tuning capabilities.
- **Incident Responder:** Monitors the system for anomalies and incidents, and triggers automated responses.
- **Resource Optimizer:** Optimizes the allocation and utilization of computing resources for ML workloads.
- **Workflow Runner:** Executes the individual tasks within a pipeline, using Argo Workflows as the underlying execution engine.

### Core Technologies

- **[Apache Iceberg](https://iceberg.apache.org/):** Used as the open table format for the data lake, providing fast and reliable access to large datasets.
- **[Argo Workflows](https://argoproj.github.io/argo-workflows/):** The container-native workflow engine for orchestrating parallel jobs on Kubernetes.
- **[Minikube](https://minikube.sigs.k8s.io/docs/start/):** Used for running a local Kubernetes cluster for development and testing.
- **[Temporal](https://temporal.io/):** Provides a durable and reliable platform for orchestrating and executing asynchronous tasks and workflows.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.24 or later)
- [Node.js](https://nodejs.org/) (version 18 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [Buf](https://buf.build/docs/installation)

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

3.  **Generate protobuf files:**

    ```sh
    make proto
    ```

## Building and Running

### Quick Start with Docker Compose

The easiest way to get started:

```sh
make up
```

This will start all services including the dashboard. Access it at `http://localhost:8080`

### Build Everything

To build all services and the dashboard:

```sh
make build
```

### Build Individual Components

**Build Go services only:**

```sh
make build-services
```

**Build dashboard only:**

```sh
make dashboard
```

### Development Mode

**Run API Gateway locally:**

```sh
make run-gateway
```

**Run Dashboard in development mode:**

```sh
cd web/dashboard
npm install
npm run dev
```

**Run other services:**

```sh
# Model Registry
go run ./cmd/model-registry

# AutoML Service
go run ./cmd/automl-service

# Iceberg Manager
go run ./cmd/iceberg-manager
```

## API

The NeuralOps API is defined using OpenAPI 3.0. The API specification can be found in `api/openapi/gateway.yaml`.

### Core Endpoints

**Pipelines:**

- `POST /v1/pipelines:nl` - Create a pipeline from natural language
- `POST /v1/pipelines/{id}:run` - Run a pipeline
- `GET /v1/runs/{id}` - Get run status
- `GET /v1/pipelines` - List all pipelines

**Model Registry:**

- `POST /v1/models` - Register a new model
- `GET /v1/models` - List all models
- `POST /v1/models/{id}/versions` - Create model version
- `POST /v1/models/{id}/versions/{versionId}:promote` - Promote model

**AutoML:**

- `POST /v1/automl/experiments` - Create experiment
- `POST /v1/automl/experiments/{id}:start` - Start tuning
- `GET /v1/automl/experiments/{id}/status` - Get status
- `GET /v1/automl/experiments/{id}/best-model` - Get best model

**Data Catalog:**

- `GET /v1/data/catalog` - List datasets
- `GET /v1/data/catalog/{name}/versions` - List versions
- `GET /v1/data/catalog/{name}/versions/{version}/lineage` - Get lineage

**Health:**

- `GET /healthz` - Health check

## Project Structure

```
neuralops/
├── api/
│   ├── openapi/           # OpenAPI specifications
│   └── proto/             # Protocol Buffer definitions
│       ├── ai_engine/
│       ├── automl/        # ✨ New: AutoML service
│       ├── iceberg/
│       ├── model_registry/ # ✨ New: Model registry
│       └── orchestrator/
├── cmd/                   # Service entry points
│   ├── api-gateway/
│   ├── automl-service/    # ✨ New: AutoML service
│   ├── model-registry/    # ✨ New: Model registry
│   └── ...
├── internal/
│   ├── domain/            # Domain models
│   └── services/          # Business logic
├── web/
│   └── dashboard/         # ✨ New: React dashboard
│       ├── src/
│       │   ├── components/
│       │   ├── pages/
│       │   └── api/
│       └── package.json
├── docker-compose.enhanced.yaml  # ✨ New: Enhanced compose
├── Dockerfile.multi       # ✨ New: Multi-stage build
├── IMPROVEMENTS.md        # ✨ Detailed improvements doc
└── Makefile
```

## Usage Examples

### Create Pipeline via Natural Language

```bash
curl -X POST http://localhost:8080/v1/pipelines:nl \
  -H "Content-Type: application/json" \
  -d '{
    "query": "Create a pipeline that ingests customer data from S3, performs feature engineering, trains a random forest classifier, and deploys the model"
  }'
```

### Register a Model

```bash
curl -X POST http://localhost:8080/v1/models \
  -H "Content-Type: application/json" \
  -d '{
    "name": "customer-churn-v1",
    "framework": "scikit-learn",
    "artifact_uri": "s3://models/churn-v1.pkl",
    "metrics": {"accuracy": 0.94, "f1_score": 0.92}
  }'
```

### Start AutoML Experiment

```bash
curl -X POST http://localhost:8080/v1/automl/experiments \
  -H "Content-Type: application/json" \
  -d '{
    "name": "churn-prediction",
    "dataset_path": "s3://data/customers.csv",
    "target_column": "churned",
    "task_type": "classification"
  }'
```

## Stopping Services

```sh
make down
```

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

- `POST /v1/pipelines:nl`: Create a pipeline from a natural language query.
- `POST /v1/pipelines/{id}:run`: Run a pipeline.
- `GET /v1/runs/{id}`: Get the status of a run.
- `GET /healthz`: Health check.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
