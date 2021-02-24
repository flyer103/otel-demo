package otlp

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
	"go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

// NewController inits an otlp controller.
func NewController(ctx context.Context, endpoint string) (*controller.Controller, error) {
	klog.V(1).InfoS("Init driver", "type", "otlp")
	driver := otlpgrpc.NewDriver(
		otlpgrpc.WithInsecure(),
		otlpgrpc.WithEndpoint(endpoint),
		otlpgrpc.WithDialOption(grpc.WithBlock()),
	)
	klog.V(1).InfoS("Init exporter", "type", "otlp")
	exporter, err := otlp.NewExporter(ctx, driver)
	if err != nil {
		klog.ErrorS(err, "Failed to create exporter")
		return nil, err
	}

	klog.V(1).InfoS("New controller", "type", "otlp")
	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithPusher(exporter),
		controller.WithCollectPeriod(2*time.Second),
	)

	klog.V(1).InfoS("Set global instances", "type", "otlp")
	global.SetMeterProvider(pusher.MeterProvider())

	klog.V(1).InfoS("Start controller", "type", "otlp")
	err = pusher.Start(ctx)
	if err != nil {
		klog.ErrorS(err, "Failed to start pusher")
		return nil, err
	}

	return pusher, nil
}
