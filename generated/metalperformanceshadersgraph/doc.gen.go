// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

// Package metalperformanceshadersgraph provides Go bindings for the MetalPerformanceShadersGraph framework.
//
// Build, compile, and execute compute graphs utilizing all the different compute devices on the platform, including GPU, CPU, and Neural Engine. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetalPerformanceShadersGraph without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph
package metalperformanceshadersgraph

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MetalPerformanceShadersGraph.framework/MetalPerformanceShadersGraph"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

