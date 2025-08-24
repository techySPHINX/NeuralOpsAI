package clients

import (
	"context"
	"fmt"
	"io"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

type HelmClient struct {
	settings *cli.EnvSettings
	actionConfig *action.Configuration
}

func NewHelmClient(namespace string) (*HelmClient, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	// Init with a dummy logger for now
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Printf(format, v...)
	}); err != nil {
		return nil, fmt.Errorf("failed to initialize helm action config: %w", err)
	}

	return &HelmClient{
		settings: settings,
		actionConfig: actionConfig,
	}, nil
}

func (c *HelmClient) Upgrade(ctx context.Context, releaseName, chartPath string, values map[string]interface{}) (*release.Release, error) {
	client := action.NewUpgrade(c.actionConfig)
	client.Namespace = c.actionConfig.Namespace
	client.Install = true // Install if not present
	client.Atomic = true // Rollback on failure

	ch, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load chart: %w", err)
	}

	return client.Run(releaseName, ch, values)
}
