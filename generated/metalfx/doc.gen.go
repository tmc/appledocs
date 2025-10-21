// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

// Package metalfx provides Go bindings for the MetalFX framework.
//
// Boost your Metal app’s performance by upscaling lower-resolution content to save GPU time. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MetalFX without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX
package metalfx

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MetalFX.framework/MetalFX"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

