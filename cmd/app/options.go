package app

import (
	"github.com/spf13/pflag"

	"github.com/flyer103/otel-demo/pkg/controller"
)

// Options is the configuration of demo.
type Options struct {
	ControllerOptions *controller.Options

	PrintVersion bool
}

// NewOptions news configurations of demo.
func NewOptions() *Options {
	controllerOptions := controller.NewOptions()

	return &Options{
		ControllerOptions: controllerOptions,
	}
}

// AddFlags adds flag options.
func (op *Options) AddFlags(fs *pflag.FlagSet) {
	op.ControllerOptions.AddFlags(fs)

	fs.BoolVar(&op.PrintVersion, "version", false, "Print version info and quit")
}
