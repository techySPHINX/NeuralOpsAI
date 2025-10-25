# NeuralOps AI Dashboard

A modern web dashboard for managing ML pipelines, models, and workflows in the NeuralOps AI platform.

## Features Implemented

### 1. **UI Dashboard** ✅

- Modern React-based interface with Material-UI
- Responsive design with dark theme
- Real-time pipeline monitoring
- System health metrics visualization
- Interactive charts and statistics

### 2. **Model Registry** ✅

- Model versioning and management
- Framework-agnostic model storage
- Model status tracking (production, staging, archived)
- Version comparison and rollback capabilities
- Metadata and metrics tracking

### 3. **AutoML Integration** ✅

- Automated model selection
- Hyperparameter tuning experiments
- Multiple algorithm comparison
- Best model recommendation
- Experiment tracking and visualization

### 4. **Data Versioning** ✅

- Iceberg-based data versioning
- Schema evolution tracking
- Data lineage visualization
- Version comparison and diff
- Snapshot management

### 5. **Workflow Visualization** ✅

- Interactive DAG visualization
- Real-time execution status
- Node-level monitoring
- Log streaming
- Task dependencies visualization

## Setup

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
cd web/dashboard
npm install
```

### Development

```bash
npm run dev
```

The dashboard will be available at `http://localhost:3000`

### Build for Production

```bash
npm run build
```

## Project Structure

```
web/dashboard/
├── src/
│   ├── api/              # API client utilities
│   ├── components/       # Reusable components
│   │   ├── Layout.tsx
│   │   └── WorkflowVisualizer.tsx
│   ├── pages/           # Main pages
│   │   ├── Dashboard.tsx
│   │   ├── Pipelines.tsx
│   │   ├── PipelineBuilder.tsx
│   │   ├── Models.tsx
│   │   ├── AutoML.tsx
│   │   ├── Workflows.tsx
│   │   └── DataCatalog.tsx
│   ├── App.tsx
│   └── main.tsx
├── package.json
└── vite.config.ts
```

## Environment Variables

Create a `.env` file:

```env
VITE_API_URL=http://localhost:8080/api
```

## Features Overview

### Pipeline Builder

- Natural language pipeline creation
- Visual drag-and-drop interface
- Real-time validation
- Template library

### Model Registry

- Register models with metadata
- Track model versions
- Promote models across environments
- Download model artifacts

### AutoML Experiments

- Configure dataset and target
- Select task type (classification, regression, etc.)
- Automatic algorithm selection
- Hyperparameter optimization
- Results comparison

### Data Catalog

- Browse available datasets
- View dataset schemas
- Track data lineage
- Compare versions
- Access metadata

### Workflow Monitor

- Real-time workflow execution
- Interactive DAG visualization
- Node status tracking
- Log access
- Performance metrics

## API Integration

The dashboard integrates with the following NeuralOps services:

- **API Gateway** (`:8080`) - Main entry point
- **Orchestrator** (`:8081`) - Pipeline orchestration
- **AI Engine** (`:8082`) - ML operations
- **Iceberg Manager** (`:8083`) - Data management
- **Model Registry** (`:8084`) - Model versioning
- **AutoML Service** (`:8085`) - Automated ML

## Technologies Used

- **React 18** - UI framework
- **TypeScript** - Type safety
- **Material-UI** - Component library
- **React Flow** - Workflow visualization
- **Recharts** - Data visualization
- **Axios** - HTTP client
- **React Query** - Data fetching
- **Vite** - Build tool
