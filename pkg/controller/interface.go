package controller

import "context"

// ControllersInterface defines methods that controllers must implement.
type ControllersInterface interface {
	// Start starts all of the controllers.
	Start(ctx context.Context) error

	// Stop stops all of the controllers.
	Stop(ctx context.Context) error
}
