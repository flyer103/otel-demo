module github.com/flyer103/otel-demo

go 1.16

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter v0.20.0
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/collector v0.20.0
	go.opentelemetry.io/contrib/instrumentation/host v0.17.0
	go.opentelemetry.io/otel v0.17.0
	go.opentelemetry.io/otel/exporters/otlp v0.17.0
	go.opentelemetry.io/otel/exporters/stdout v0.17.0
	go.opentelemetry.io/otel/metric v0.17.0
	go.opentelemetry.io/otel/sdk/metric v0.17.0
	go.opentelemetry.io/otel/trace v0.17.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	k8s.io/component-base v0.20.4
	k8s.io/klog/v2 v2.5.0
)
