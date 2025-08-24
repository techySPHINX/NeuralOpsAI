package main

import (
	"context"
	"fmt"
	"time"

	"neuralops/api/proto/optimizer/v1"
	"neuralops/pkg/clients"
	"neuralops/pkg/logging"
	"github.com/prometheus/common/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OptimizerGRPCServer struct {
	optimizerv1.UnimplementedOptimizerServiceServer
	logger          *logging.Logger
	prometheusClient *clients.PrometheusClient
	kubernetesClient *clients.KubernetesClient
	helmClient       *clients.HelmClient
}

func NewOptimizerGRPCServer(logger *logging.Logger, promClient *clients.PrometheusClient, kubeClient *clients.KubernetesClient, helmClient *clients.HelmClient) *OptimizerGRPCServer {
	return &OptimizerGRPCServer{
		logger:          logger,
		prometheusClient: promClient,
		kubernetesClient: kubeClient,
		helmClient:       helmClient,
	}
}

func (s *OptimizerGRPCServer) Optimize(ctx context.Context, req *optimizerv1.OptimizeRequest) (*optimizerv1.OptimizeResponse, error) {
	s.logger.Info("Received Optimize request", "deployment_name", req.DeploymentName, "namespace", req.Namespace)

	// 1. Scrape Prometheus for CPU/memory usage of the deployment.
	// Example query for CPU usage:
	cpuQuery := fmt.Sprintf(`sum(rate(container_cpu_usage_seconds_total{namespace="%s", pod=~"%s-.*", container!=""}[5m]))`, req.Namespace, req.DeploymentName)
	cpuResult, warnings, err := s.prometheusClient.Query(ctx, cpuQuery, time.Now())
	if err != nil {
		s.logger.Error("failed to query prometheus for CPU", "error", err)
		return &optimizerv1.OptimizeResponse{Success: false, Message: "failed to query prometheus for CPU"}, err
	}
	if len(warnings) > 0 {
		s.logger.Warn("Prometheus CPU query warnings", "warnings", warnings)
	}

	// Example query for Memory usage:
	memQuery := fmt.Sprintf(`sum(container_memory_usage_bytes{namespace="%s", pod=~"%s-.*", container!=""})`, req.Namespace, req.DeploymentName)
	memResult, warnings, err := s.prometheusClient.Query(ctx, memQuery, time.Now())
	if err != nil {
		s.logger.Error("failed to query prometheus for Memory", "error", err)
		return &optimizerv1.OptimizeResponse{Success: false, Message: "failed to query prometheus for Memory"}, err
	}
	if len(warnings) > 0 {
		s.logger.Warn("Prometheus Memory query warnings", "warnings", warnings)
	}

	// 2. Apply heuristics to recommend new CPU/memory limits/requests.
	// This is a very simplified heuristic. In a real scenario, you'd use more sophisticated logic.
	recommendedCPU := make(map[string]string)
	recommendedMemory := make(map[string]string)

	if cpuVector, ok := cpuResult.(model.Vector); ok && len(cpuVector) > 0 {
		// Assuming a single result for simplicity
		cpuUsage := cpuVector[0].Value
		// If CPU usage is less than 50% of current request, recommend reducing by 20%
		// This is a placeholder heuristic
		s.logger.Info("Current CPU usage", "value", cpuUsage)
		recommendedCPU["cpu_request"] = "200m" // Example
		recommendedCPU["cpu_limit"] = "400m"   // Example
	}

	if memVector, ok := memResult.(model.Vector); ok && len(memVector) > 0 {
		// Assuming a single result for simplicity
		memUsage := memVector[0].Value
		s.logger.Info("Current Memory usage", "value", memUsage)
		recommendedMemory["memory_request"] = "256Mi" // Example
		recommendedMemory["memory_limit"] = "512Mi"   // Example
	}

	// 3. Use Helm SDK to patch the deployment's values and perform a helm upgrade.
	// This assumes the deployment was installed via Helm and we know its release name and chart path.
	// For this example, we'll just log the recommended values.
	s.logger.Info("Recommended CPU", "values", recommendedCPU)
	s.logger.Info("Recommended Memory", "values", recommendedMemory)

	// Example of how you might construct values for Helm upgrade
	helmValues := map[string]interface{}{
		"resources": map[string]interface{}{
			"requests": map[string]interface{}{
				"cpu":    recommendedCPU["cpu_request"],
				"memory": recommendedMemory["memory_request"],
			},
			"limits": map[string]interface{}{
				"cpu":    recommendedCPU["cpu_limit"],
				"memory": recommendedMemory["memory_limit"],
			},
		},
	}

	// This part would require knowing the Helm release name and chart path.
	// For now, it's commented out.
	/*
	releaseName := req.DeploymentName // Assuming deployment name is also the Helm release name
	chartPath := "/path/to/your/chart" // This needs to be dynamic or configured

	_, err = s.helmClient.Upgrade(ctx, releaseName, chartPath, helmValues)
	if err != nil {
		s.logger.Error("failed to perform helm upgrade", "error", err)
		return &optimizerv1.OptimizeResponse{Success: false, Message: "failed to apply optimization"}, err
	}
	*/

	return &optimizerv1.OptimizeResponse{
		Success:           true,
		Message:           fmt.Sprintf("Optimization request for %s/%s processed. Recommended resources logged.", req.Namespace, req.DeploymentName),
		RecommendedCpu:    recommendedCPU,
		RecommendedMemory: recommendedMemory,
	}, nil
}
