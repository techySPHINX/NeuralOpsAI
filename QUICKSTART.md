# 🚀 Quick Start Guide - NeuralOps AI v2.0

## 5-Minute Setup

### Prerequisites Check

```bash
# Check Go version (need 1.24+)
go version

# Check Node.js (need 18+)
node --version

# Check Docker
docker --version
```

### Step 1: Clone & Setup

```bash
git clone https://github.com/techySPHINX/NeuralOpsAI
cd neuralops
```

### Step 2: Install Dependencies

```bash
# Install Go tools
make dev-tools

# Install Node dependencies
cd web/dashboard
npm install
cd ../..
```

### Step 3: Start Services (Easiest Method)

```bash
# Start everything with Docker Compose
make up

# Wait 30 seconds for services to start
```

### Step 4: Access the Dashboard

Open your browser to:

```
http://localhost:8080
```

You should see the NeuralOps AI Dashboard! 🎉

---

## What You Can Do Now

### 1. Explore the Dashboard

- **Dashboard**: View system metrics
- **Pipelines**: Create and manage ML pipelines
- **Models**: Register and version models
- **AutoML**: Run hyperparameter experiments
- **Workflows**: Monitor pipeline executions
- **Data Catalog**: Browse datasets

### 2. Create Your First Pipeline

#### Via Natural Language:

```bash
curl -X POST http://localhost:8080/v1/pipelines:nl \
  -H "Content-Type: application/json" \
  -d '{
    "query": "Create a pipeline that loads iris dataset, trains a classifier, and evaluates it"
  }'
```

#### Via Dashboard:

1. Go to "Pipelines" page
2. Click "Create Pipeline"
3. Enter natural language description
4. Click "Generate"
5. Save and run!

### 3. Register a Model

```bash
curl -X POST http://localhost:8080/v1/models \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-first-model",
    "framework": "scikit-learn",
    "description": "My first ML model"
  }'
```

Or use the Models page in the dashboard.

### 4. Run AutoML Experiment

```bash
curl -X POST http://localhost:8080/v1/automl/experiments \
  -H "Content-Type: application/json" \
  -d '{
    "name": "classification-experiment",
    "dataset_path": "/data/iris.csv",
    "target_column": "species",
    "task_type": "TASK_TYPE_CLASSIFICATION"
  }'
```

Or use the AutoML page in the dashboard.

---

## Development Mode (For Developers)

### Terminal 1: Backend Services

```bash
# Model Registry
go run ./cmd/model-registry
```

### Terminal 2: AutoML Service

```bash
go run ./cmd/automl-service
```

### Terminal 3: Dashboard (Hot Reload)

```bash
cd web/dashboard
npm run dev
# Opens at http://localhost:3000
```

### Terminal 4: API Gateway

```bash
go run ./cmd/api-gateway
```

---

## Service Ports

| Service         | Port  | URL                    |
| --------------- | ----- | ---------------------- |
| Dashboard       | 3000  | http://localhost:3000  |
| API Gateway     | 8080  | http://localhost:8080  |
| Orchestrator    | 8081  | localhost:8081         |
| AI Engine       | 8082  | localhost:8082         |
| Iceberg Manager | 8083  | localhost:8083         |
| Model Registry  | 8084  | localhost:8084         |
| AutoML Service  | 8085  | localhost:8085         |
| MinIO           | 9000  | http://localhost:9000  |
| MinIO Console   | 9001  | http://localhost:9001  |
| PostgreSQL      | 5432  | localhost:5432         |
| Nessie          | 19120 | http://localhost:19120 |

---

## Stopping Everything

```bash
# Stop Docker Compose services
make down

# Or manually stop dev services with Ctrl+C
```

---

## Troubleshooting

### Services won't start

```bash
# Check if ports are in use
netstat -an | findstr "8080"

# Clean and restart
make down
make clean
make up
```

### Dashboard shows connection error

```bash
# Check if API Gateway is running
curl http://localhost:8080/healthz

# Should return "OK"
```

### Can't access MinIO

```bash
# Check MinIO is running
docker ps | findstr minio

# Access MinIO Console
# URL: http://localhost:9001
# User: minioadmin
# Pass: minioadmin
```

### Database connection fails

```bash
# Check PostgreSQL
docker ps | findstr postgres

# Test connection
psql -h localhost -U neuralops -d neuralops
# Password: neuralops
```

---

## Next Steps

1. **Read the Docs:**

   - `IMPROVEMENTS.md` - Detailed feature documentation
   - `web/dashboard/README.md` - Dashboard guide
   - `README.md` - Main documentation

2. **Try the Features:**

   - Create a pipeline using natural language
   - Register a model with versions
   - Run an AutoML experiment
   - Visualize a workflow DAG

3. **Explore the Code:**

   - `api/proto/` - gRPC definitions
   - `cmd/` - Service implementations
   - `web/dashboard/src/` - React components

4. **Join the Community:**
   - Report issues on GitHub
   - Contribute improvements
   - Share your use cases

---

## Common Commands

```bash
# Generate protobuf code
make proto

# Build all services
make build

# Run tests
make test

# Clean build artifacts
make clean

# Start all services
make up

# Stop all services
make down

# Build Docker images
make docker
```

---

## Sample Workflows

### Complete ML Pipeline

1. Upload data to MinIO
2. Create pipeline via Dashboard
3. Monitor execution in Workflows
4. Register trained model
5. Promote to production
6. Track data lineage

### AutoML Flow

1. Prepare dataset
2. Create AutoML experiment
3. Configure hyperparameters
4. Start tuning
5. View results
6. Deploy best model

### Model Versioning

1. Register initial model
2. Track metrics
3. Create new versions
4. Compare performance
5. Promote best version
6. Archive old versions

---

## Getting Help

- **Documentation**: `IMPROVEMENTS.md`
- **API Reference**: `api/openapi/gateway.yaml`
- **GitHub Issues**: [Report a bug](https://github.com/techySPHINX/NeuralOpsAI/issues)
- **Discussions**: [Ask questions](https://github.com/techySPHINX/NeuralOpsAI/discussions)

---

## What's New in v2.0

✅ Web Dashboard with 7 pages  
✅ Model Registry for versioning  
✅ AutoML with hyperparameter tuning  
✅ Enhanced data versioning  
✅ Interactive workflow visualization

See `IMPROVEMENTS.md` for details!

---

**Happy MLOps! 🚀🤖**
