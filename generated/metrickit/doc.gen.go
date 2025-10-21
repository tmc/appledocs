// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

// Package metrickit provides Go bindings for the MetricKit framework.
//
// Aggregate and analyze per-device reports on exception and crash diagnostics and on power and performance metrics. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetricKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit
package metrickit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MetricKit.framework/MetricKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


