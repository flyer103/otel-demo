package custom

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/klog/v2"
)

// Run runs custom example.
func Run(ctx context.Context) error {
	klog.V(1).InfoS("Record trace manually", "stage", "run")
	tracer := otel.Tracer("ex.com/webserver")
	anotherKey := label.Key("ex.com/another")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var span trace.Span
	for {
		select {
		case <-ctx.Done():
			klog.V(1).InfoS("Stop record trace", "stage", "stop")
			break
		case t := <-ticker.C:
			ctx, span = tracer.Start(ctx, "operation")

			event := fmt.Sprintf("Now: %s", t.Format("2006.01.02 15:04:05"))
			span.AddEvent(event, trace.WithAttributes(label.Int("bogons", 100)))
			span.SetAttributes(anotherKey.String("yes"))

			if t.Unix()%2 == 0 {
				span.End()
			}
		}
	}

	return nil
}
