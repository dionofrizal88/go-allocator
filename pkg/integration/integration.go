package integration

import (
	"context"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider"
)

// Integration contain provider that used for request or get info.
type Integration struct {
	provider provider.Interface
}

// NewIntegration is a constructor.
func NewIntegration(provider provider.Interface) *Integration {
	return &Integration{
		provider: provider,
	}
}

// Info is for get info of demography.
func (c *Integration) Info(ctx context.Context) (interface{}, error) {
	return c.provider.Info(ctx)
}
