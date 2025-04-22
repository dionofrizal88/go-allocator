package provider

import "context"

// Interface is an interface that implemented by each provider.
type Interface interface {
	// SetInput is for set input. Input maybe different each provider.
	SetInput(interface{}) error
	// Info return struct maybe different each provider.
	Info(context.Context) (interface{}, error)
}
