import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Pipeline APIs
export const pipelineAPI = {
  create: (data: any) => apiClient.post('/v1/pipelines:nl', data),
  list: () => apiClient.get('/v1/pipelines'),
  get: (id: string) => apiClient.get(`/v1/pipelines/${id}`),
  run: (id: string) => apiClient.post(`/v1/pipelines/${id}:run`),
  delete: (id: string) => apiClient.delete(`/v1/pipelines/${id}`),
};

// Model Registry APIs
export const modelAPI = {
  register: (data: any) => apiClient.post('/v1/models', data),
  list: () => apiClient.get('/v1/models'),
  get: (id: string) => apiClient.get(`/v1/models/${id}`),
  updateStatus: (id: string, status: string) =>
    apiClient.patch(`/v1/models/${id}/status`, { status }),
  delete: (id: string) => apiClient.delete(`/v1/models/${id}`),
  createVersion: (modelId: string, data: any) =>
    apiClient.post(`/v1/models/${modelId}/versions`, data),
  listVersions: (modelId: string) =>
    apiClient.get(`/v1/models/${modelId}/versions`),
  promote: (modelId: string, versionId: string, stage: string) =>
    apiClient.post(`/v1/models/${modelId}/versions/${versionId}:promote`, { stage }),
};

// AutoML APIs
export const automlAPI = {
  createExperiment: (data: any) => apiClient.post('/v1/automl/experiments', data),
  listExperiments: () => apiClient.get('/v1/automl/experiments'),
  getExperiment: (id: string) => apiClient.get(`/v1/automl/experiments/${id}`),
  startTuning: (experimentId: string) =>
    apiClient.post(`/v1/automl/experiments/${experimentId}:start`),
  getStatus: (experimentId: string) =>
    apiClient.get(`/v1/automl/experiments/${experimentId}/status`),
  getBestModel: (experimentId: string) =>
    apiClient.get(`/v1/automl/experiments/${experimentId}/best-model`),
};

// Data Catalog APIs
export const dataCatalogAPI = {
  list: () => apiClient.get('/v1/data/catalog'),
  get: (name: string) => apiClient.get(`/v1/data/catalog/${name}`),
  listVersions: (datasetName: string) =>
    apiClient.get(`/v1/data/catalog/${datasetName}/versions`),
  getVersion: (datasetName: string, version: string) =>
    apiClient.get(`/v1/data/catalog/${datasetName}/versions/${version}`),
  getLineage: (datasetName: string, version: string) =>
    apiClient.get(`/v1/data/catalog/${datasetName}/versions/${version}/lineage`),
  compareVersions: (datasetName: string, v1: string, v2: string) =>
    apiClient.get(`/v1/data/catalog/${datasetName}/compare`, { params: { v1, v2 } }),
};

// Workflow APIs
export const workflowAPI = {
  list: () => apiClient.get('/v1/workflows'),
  get: (id: string) => apiClient.get(`/v1/workflows/${id}`),
  getStatus: (runId: string) => apiClient.get(`/v1/runs/${runId}`),
  getLogs: (runId: string, nodeId: string) =>
    apiClient.get(`/v1/runs/${runId}/nodes/${nodeId}/logs`),
};

export default apiClient;
