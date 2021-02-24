package host

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/host"
)

// Run runs host meter.
func Run(ctx context.Context) error {
	return host.Start()
}
