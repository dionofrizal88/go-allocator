package auth

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/integration/provider"
)

// Auth is a concrete that contain input to perform know auth.
type Auth struct {
	AppConfig config.Configuration
	InputAuth InputAuth
	RequestID string
}

var _ provider.Interface = &Auth{}
