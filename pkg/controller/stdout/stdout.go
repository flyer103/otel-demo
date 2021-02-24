package stdout

import (
	"go.opentelemetry.io/otel/exporters/stdout"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
)

// NewController inits a stdout controller.
func NewController() (*controller.Controller, error) {
	return stdout.InstallNewPipeline([]stdout.Option{
		stdout.WithPrettyPrint(),
	}, nil)
}
