package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/klog/v2"

	"github.com/flyer103/otel-demo/pkg/controller"
	metricsCustom "github.com/flyer103/otel-demo/pkg/signal/metrics/custom"
	"github.com/flyer103/otel-demo/pkg/signal/metrics/host"
	traceCustom "github.com/flyer103/otel-demo/pkg/signal/trace/custom"
)

// Run launches the app.
func Run(op *Options) error {
	ctx := context.Background()

	// Start Controllers.
	klog.InfoS("Init controllers", "type", "controller", "stage", "init")
	cs, err := controller.NewControllers(ctx, op.ControllerOptions)
	if err != nil {
		klog.ErrorS(err, "Failed to init controllers", "type", "controller")
		return err
	}

	klog.InfoS("Start controllers", "type", "controller", "stage", "init")
	err = cs.Start(ctx)
	if err != nil {
		klog.ErrorS(err, "Failed to start controllers", "type", "controller", "stage", "init")
		return err
	}
	klog.InfoS("Controllers has started", "type", "controller", "stage", "run")
	defer func(ctx context.Context) {
		err := cs.Stop(ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to stop controllers", "type", "controller", "stage", "stop")
		}
	}(ctx)

	// Use metrics.
	klog.InfoS("Run host metrics", "type", "metrics")
	go func(ctx context.Context) {
		err := host.Run(ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to run host example", "type", "metrics")
		}
	}(ctx)

	klog.InfoS("Run custom metrics", "type", "metrics")
	go func(ctx context.Context) {
		err := metricsCustom.Run(ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to run custom example", "type", "metrics")
		}
	}(ctx)

	// Use trace.
	klog.InfoS("Run custom trace", "type", "trace")
	go func(ctx context.Context) {
		err := traceCustom.Run(ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to run custom example", "type", "trace")
		}
	}(ctx)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)
	<-stopCh

	return nil
}
