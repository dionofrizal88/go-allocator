package agentallocation

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/go-redis/redis/v8"
)

// Controller struct is used when get configuration.
type Controller struct {
	config config.Configuration
	redis  *redis.Client
}

// NewController will initialize Controller.
func NewController(config config.Configuration, redis *redis.Client) *Controller {
	return &Controller{
		config: config,
		redis:  redis,
	}
}
