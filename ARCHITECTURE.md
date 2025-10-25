# NeuralOps AI - Architecture Diagram v2.0

## System Architecture

```
┌────────────────────────────────────────────────────────────────────────────┐
│                                                                            │
│                         🌐 Web Browser (User)                              │
│                                                                            │
└──────────────────────────────────┬─────────────────────────────────────────┘
                                   │
                                   │ HTTP/HTTPS
                                   │
┌──────────────────────────────────▼─────────────────────────────────────────┐
│                                                                            │
│                   📱 React Dashboard (TypeScript)                          │
│                        Port: 3000 (dev) / 8080 (prod)                      │
│                                                                            │
│  ┌──────────┐ ┌──────────┐ ┌─────────┐ ┌────────┐ ┌──────────┐          │
│  │Dashboard │ │Pipelines │ │ Models  │ │ AutoML │ │Workflows │          │
│  └──────────┘ └──────────┘ └─────────┘ └────────┘ └──────────┘          │
│                                                                            │
└──────────────────────────────────┬─────────────────────────────────────────┘
                                   │
                                   │ REST API
                                   │
┌──────────────────────────────────▼─────────────────────────────────────────┐
│                                                                            │
│                        🚪 API Gateway (Go)                                 │
│                              Port: 8080                                    │
│                                                                            │
│  • Route management          • Authentication (future)                    │
│  • Load balancing            • Rate limiting (future)                     │
│  • Request validation        • API versioning                             │
│                                                                            │
└───┬────────────┬────────────┬────────────┬────────────┬───────────────────┘
    │            │            │            │            │
    │ gRPC       │ gRPC       │ gRPC       │ gRPC       │ gRPC
    │            │            │            │            │
┌───▼────┐  ┌───▼─────┐  ┌───▼─────┐  ┌──▼──────┐  ┌──▼─────────┐
│Orchestr│  │AI Engine│  │ Iceberg │  │  Model  │  │   AutoML   │
│ator    │  │         │  │ Manager │  │Registry │  │  Service   │
│:8081   │  │:8082    │  │:8083    │  │:8084    │  │  :8085     │
└───┬────┘  └───┬─────┘  └───┬─────┘  └──┬──────┘  └──┬─────────┘
    │            │            │            │            │
    │            │            │            │            │
┌───▼────────────▼────────────▼────────────▼────────────▼───────────────────┐
│                                                                            │
│                    💾 Data & Storage Layer                                 │
│                                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  PostgreSQL  │  │    MinIO     │  │    Nessie    │  │     Argo     │ │
│  │              │  │              │  │              │  │  Workflows   │ │
│  │  Metadata    │  │ Object Store │  │   Catalog    │  │              │ │
│  │   :5432      │  │    :9000     │  │   :19120     │  │    :2746     │ │
│  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘ │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘
```

---

## Component Interactions

### 1. Pipeline Creation Flow

```
User → Dashboard → API Gateway → Orchestrator → Argo Workflows
                                      ↓
                                  AI Engine
                                      ↓
                                Iceberg Manager → MinIO + Nessie
```

### 2. Model Registry Flow

```
User → Dashboard → API Gateway → Model Registry → PostgreSQL
                                        ↓
                                     MinIO (artifacts)
```

### 3. AutoML Experiment Flow

```
User → Dashboard → API Gateway → AutoML Service → PostgreSQL (metadata)
                                        ↓
                                   AI Engine (training)
                                        ↓
                                  Model Registry
```

### 4. Data Versioning Flow

```
Data Change → Iceberg Manager → Nessie (catalog)
                   ↓
              MinIO (storage)
                   ↓
            PostgreSQL (metadata)
```

### 5. Workflow Monitoring Flow

```
Argo Workflows → Orchestrator → API Gateway → Dashboard
                                      ↑
                                  (polling/websocket)
```

---

## Data Flow Diagram

```
┌─────────────┐
│   Raw Data  │
│  (S3/Local) │
└──────┬──────┘
       │
       │ Ingestion
       ▼
┌─────────────────┐
│ Iceberg Manager │ ────┐
│  (Data Lake)    │     │
└─────────┬───────┘     │
          │             │ Version
          │             │ Metadata
          ▼             ▼
    ┌─────────┐   ┌──────────┐
    │  MinIO  │   │  Nessie  │
    │(Storage)│   │(Catalog) │
    └─────────┘   └──────────┘
          │
          │ Read Data
          ▼
    ┌──────────┐
    │AI Engine │
    │(Training)│
    └────┬─────┘
         │
         │ Trained Model
         ▼
  ┌────────────────┐
  │ Model Registry │
  │  (Versioning)  │
  └────────┬───────┘
           │
           │ Deploy
           ▼
     ┌──────────┐
     │Production│
     └──────────┘
```

---

## Service Responsibilities

### 🌐 Web Dashboard

- **Tech**: React 18, TypeScript, Material-UI
- **Purpose**: User interface for all operations
- **Features**:
  - Pipeline builder (visual + NL)
  - Model management
  - AutoML experiments
  - Workflow monitoring
  - Data catalog browser

### 🚪 API Gateway

- **Tech**: Go, Chi Router
- **Purpose**: Single entry point for all requests
- **Responsibilities**:
  - Route HTTP to gRPC services
  - Request validation
  - (Future) Authentication & authorization
  - (Future) Rate limiting

### 🎯 Orchestrator

- **Tech**: Go, gRPC
- **Purpose**: Pipeline orchestration
- **Responsibilities**:
  - Parse natural language queries
  - Compile to Argo workflows
  - Manage pipeline lifecycle
  - Track execution status

### 🤖 AI Engine

- **Tech**: Go, Python (workers), gRPC
- **Purpose**: ML operations
- **Responsibilities**:
  - Model training
  - Model evaluation
  - Inference serving
  - Metrics collection

### 📦 Iceberg Manager

- **Tech**: Go, Apache Iceberg, gRPC
- **Purpose**: Data lake management
- **Responsibilities**:
  - Data versioning (snapshots)
  - Schema evolution
  - Data lineage tracking
  - Query optimization

### 📚 Model Registry (NEW)

- **Tech**: Go, gRPC
- **Purpose**: ML model lifecycle management
- **Responsibilities**:
  - Model versioning
  - Metadata tracking
  - Stage promotion (dev→staging→prod)
  - Artifact storage

### 🧠 AutoML Service (NEW)

- **Tech**: Go, Python (Ray Tune/Optuna), gRPC
- **Purpose**: Automated machine learning
- **Responsibilities**:
  - Hyperparameter tuning
  - Model selection
  - Experiment tracking
  - Best model recommendation

---

## Storage Layer

### PostgreSQL

- **Purpose**: Relational metadata storage
- **Stores**:
  - Model metadata & versions
  - AutoML experiments & trials
  - User data & permissions
  - System configuration

### MinIO

- **Purpose**: Object storage
- **Stores**:
  - Data files (Parquet, CSV, etc.)
  - Model artifacts (.pkl, .h5, .pt)
  - Training datasets
  - Logs and results

### Nessie

- **Purpose**: Data catalog with Git-like versioning
- **Stores**:
  - Iceberg table metadata
  - Schema versions
  - Branch/tag information
  - Commit history

### Argo Workflows

- **Purpose**: Workflow execution engine
- **Manages**:
  - Pipeline DAGs
  - Task execution
  - Dependency resolution
  - Resource allocation

---

## Network Ports Reference

| Component       | Port(s)    | Protocol   | Purpose            |
| --------------- | ---------- | ---------- | ------------------ |
| Dashboard       | 3000       | HTTP       | Development server |
| API Gateway     | 8080       | HTTP/gRPC  | Public API         |
| Orchestrator    | 8081       | gRPC       | Internal service   |
| AI Engine       | 8082       | gRPC       | Internal service   |
| Iceberg Manager | 8083       | gRPC       | Internal service   |
| Model Registry  | 8084       | gRPC       | Internal service   |
| AutoML Service  | 8085       | gRPC       | Internal service   |
| MinIO           | 9000, 9001 | HTTP       | Storage + Console  |
| PostgreSQL      | 5432       | PostgreSQL | Database           |
| Nessie          | 19120      | HTTP       | Catalog API        |
| Argo            | 2746       | HTTP       | Workflow UI        |
| Temporal        | 7233       | gRPC       | Temporal server    |

---

## Security Architecture (Future)

```
┌─────────────────────────────────────────┐
│          Authentication Layer           │
│                                         │
│  ┌──────────┐  ┌──────────┐            │
│  │  OAuth2  │  │   JWT    │            │
│  │ Provider │  │  Tokens  │            │
│  └──────────┘  └──────────┘            │
└─────────────────┬───────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────┐
│         Authorization Layer             │
│                                         │
│  ┌──────────┐  ┌──────────┐            │
│  │   RBAC   │  │   ACL    │            │
│  │  Roles   │  │  Rules   │            │
│  └──────────┘  └──────────┘            │
└─────────────────────────────────────────┘
```

---

## Deployment Architecture

### Kubernetes Deployment

```
┌────────────────────────────────────────────┐
│         Kubernetes Cluster                 │
│                                            │
│  ┌──────────────────────────────────┐     │
│  │        Namespace: neuralops      │     │
│  │                                  │     │
│  │  ┌────────┐  ┌────────┐         │     │
│  │  │ Pod    │  │ Pod    │         │     │
│  │  │Gateway │  │Registry│  ...    │     │
│  │  └────────┘  └────────┘         │     │
│  │                                  │     │
│  │  ┌────────────────────────┐     │     │
│  │  │   StatefulSet          │     │     │
│  │  │   PostgreSQL, MinIO    │     │     │
│  │  └────────────────────────┘     │     │
│  └──────────────────────────────────┘     │
│                                            │
│  ┌──────────────────────────────────┐     │
│  │     Ingress Controller           │     │
│  │  (nginx/traefik)                 │     │
│  └──────────────────────────────────┘     │
└────────────────────────────────────────────┘
```

---

## Scalability Strategy

### Horizontal Scaling

- API Gateway: Multiple replicas behind load balancer
- Worker services: Auto-scale based on queue depth
- Database: Read replicas for query workloads

### Vertical Scaling

- AI Engine: GPU nodes for training
- AutoML Service: High-memory nodes for experiments

### Data Partitioning

- Iceberg: Partition by date/category
- MinIO: Distributed across nodes
- PostgreSQL: Sharding by tenant (future)

---

This architecture enables:
✅ **Scalability** - Horizontal and vertical scaling  
✅ **Reliability** - Service isolation and redundancy  
✅ **Maintainability** - Modular microservices  
✅ **Extensibility** - Plugin architecture  
✅ **Performance** - Distributed storage and processing
