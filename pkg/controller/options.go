package controller

import (
	"github.com/spf13/pflag"
)

// Options is the configuration of controllers.
type Options struct {
	// Use stdout exporter.
	UseStdoutExporter bool

	// Use Collector.
	UseCollector bool
	OTLPEndpoint string
}

// NewOptions news configurations of controllers.
func NewOptions() *Options {
	return &Options{}
}

// AddFlags adds flag options.
func (op *Options) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&op.UseStdoutExporter, "use-stdout-exporter", true, "Use stdout exporter to receive/process/export signal?")

	fs.BoolVar(&op.UseCollector, "use-collector", false, "Use Collector to receive/process/export signal?")
	fs.StringVar(&op.OTLPEndpoint, "collector-otlp-endpoint", "", "endpoint of otlp")
}
