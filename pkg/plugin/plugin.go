package plugin

import "context"

// Runner is the interface that all plugins must implement.
type Runner interface {
	Apply(ctx context.Context, input []byte) ([]byte, error)
}
