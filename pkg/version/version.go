package version

import (
	"fmt"
	"runtime"
)

// Package is the package of the app.
var Package = "github.com/flyer103/otel-demo"

// Version is the version of the app.
var Version = "unknown"

// GoVersion is the version of go that builds the app.
var GoVersion = runtime.Version()

// PrintVersion prints the app's version.
func PrintVersion() {
	fmt.Println(Package)
	fmt.Println(Version, GoVersion)
}
