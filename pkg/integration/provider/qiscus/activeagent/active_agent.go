package activeagent

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider"
)

// ActiveAgent is a concrete that contain input to perform know active agent.
type ActiveAgent struct {
	AppConfig        config.Configuration
	InputActiveAgent InputActiveAgent
	Authorization    string
	RequestID        string
}

var _ provider.Interface = &ActiveAgent{}
