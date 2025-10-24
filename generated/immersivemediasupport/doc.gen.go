// Code generated from Apple documentation for ImmersiveMediaSupport. DO NOT EDIT.

// Package immersivemediasupport provides Go bindings for the ImmersiveMediaSupport framework.
//
// Read and write essential Apple Immersive Video metadata.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ImmersiveMediaSupport without requiring cgo.
//
// See: https://developer.apple.com/documentation/ImmersiveMediaSupport
package immersivemediasupport

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ImmersiveMediaSupport.framework/ImmersiveMediaSupport"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

