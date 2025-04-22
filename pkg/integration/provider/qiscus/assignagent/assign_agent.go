package assignagent

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider"
)

// AssignAgent is a concrete that contain input to perform know assign agent.
type AssignAgent struct {
	AppConfig        config.Configuration
	InputAssignAgent InputAssignAgent
	RequestID        string
}

var _ provider.Interface = &AssignAgent{}
