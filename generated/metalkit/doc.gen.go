// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

// Package metalkit provides Go bindings for the MetalKit framework.
//
// Build Metal apps quicker and easier using a common set of utility classes. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetalKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit
package metalkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MetalKit.framework/MetalKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

