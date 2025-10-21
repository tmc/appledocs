// Code generated from Apple documentation for LightweightCodeRequirements. DO NOT EDIT.

// Package lightweightcoderequirements provides Go bindings for the LightweightCodeRequirements framework.
//
// Test the identity of executable code on disk and in running processes. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to LightweightCodeRequirements without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/LightweightCodeRequirements
package lightweightcoderequirements

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/LightweightCodeRequirements.framework/LightweightCodeRequirements"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

