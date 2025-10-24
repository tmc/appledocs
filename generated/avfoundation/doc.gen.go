
// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

// Package avfoundation provides Go bindings for the AVFoundation framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AVFoundation without requiring cgo.
package avfoundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVFoundation.framework/AVFoundation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

