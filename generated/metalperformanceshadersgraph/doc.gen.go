
// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

// Package metalperformanceshadersgraph provides Go bindings for the MetalPerformanceShadersGraph framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetalPerformanceShadersGraph without requiring cgo.
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

