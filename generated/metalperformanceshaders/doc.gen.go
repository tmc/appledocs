// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

// Package metalperformanceshaders provides Go bindings for the MetalPerformanceShaders framework.
//
// Optimize graphics and compute performance with kernels that are fine-tuned for the unique characteristics of each Metal GPU family. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetalPerformanceShaders without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders
package metalperformanceshaders

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MetalPerformanceShaders.framework/MetalPerformanceShaders"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

