package custom

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"k8s.io/klog/v2"
)

// Run runs custom example.
func Run(ctx context.Context) error {
	klog.V(1).InfoS("Prepare metric", "stage", "init")
	lemonKey := label.Key("ex.com/lemons")
	commonLabels := []label.KeyValue{
		lemonKey.Int(10),
		label.String("A", "1"),
		label.String("B", "2"),
		label.String("C", "3"),
	}
	meter := global.Meter("ex.com/webserver")

	klog.V(1).InfoS("Run observer to record specified metric automatically", "stage", "run")
	observerCallback := func(_ context.Context, result metric.Float64ObserverResult) {
		result.Observe(float64(time.Now().Unix()), commonLabels...)
	}
	_ = metric.Must(meter).NewFloat64ValueObserver("request", observerCallback,
		metric.WithDescription("Has observed some requests"),
	)

	klog.V(1).InfoS("Record metrics manually", "stage", "run")
	valueRecorder := metric.Must(meter).NewFloat64ValueRecorder("latency")
	boundRecorder := metric.Must(meter).NewFloat64ValueRecorder("throughput").Bind(commonLabels...)
	defer boundRecorder.Unbind()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			klog.V(1).InfoS("Stop record metrics", "stage", "stop")
			break
		case t := <-ticker.C:
			val := float64(t.Unix())

			meter.RecordBatch(ctx, commonLabels, valueRecorder.Measurement(val))
			boundRecorder.Record(ctx, val)
			klog.V(1).InfoS("Record 1 metric", "stage", "run")
		}
	}

	return nil
}
