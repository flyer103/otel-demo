package controller

import (
	"context"

	otelController "go.opentelemetry.io/otel/sdk/metric/controller/basic"

	"github.com/flyer103/otel-demo/pkg/controller/otlp"
	"github.com/flyer103/otel-demo/pkg/controller/stdout"
)

type controllers struct {
	Controllers []*otelController.Controller
}

// NewControllers inits controllers.
func NewControllers(ctx context.Context, op *Options) (ControllersInterface, error) {
	c := []*otelController.Controller{}

	if op.UseStdoutExporter {
		stdoutController, err := stdout.NewController()
		if err != nil {
			return nil, err
		}
		c = append(c, stdoutController)
	}

	if op.UseCollector {
		otlpController, err := otlp.NewController(ctx, op.OTLPEndpoint)
		if err != nil {
			return nil, err
		}
		c = append(c, otlpController)
	}

	return &controllers{
		Controllers: c,
	}, nil
}

// Note: Controllers has started when initialization and we don't need to start explicitly.
func (c *controllers) Start(ctx context.Context) error {

	return nil
}

func (c *controllers) Stop(ctx context.Context) error {
	for _, controller := range c.Controllers {
		err := controller.Stop(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
