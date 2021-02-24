package main

import (
	"os"

	"github.com/spf13/pflag"
	"k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"

	"github.com/flyer103/otel-demo/cmd/app"
	"github.com/flyer103/otel-demo/pkg/version"
)

func main() {
	op := app.NewOptions()
	op.AddFlags(pflag.CommandLine)

	flag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if op.PrintVersion {
		version.PrintVersion()
		os.Exit(0)
	}

	if err := app.Run(op); err != nil {
		klog.ErrorS(err, "Failed to run app")
		os.Exit(1)
	}
}
