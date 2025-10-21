// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

// Package coreaudiokit provides Go bindings for the CoreAudioKit framework.
//
// Add user interfaces to audio units. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreAudioKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit
package coreaudiokit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreAudioKit.framework/CoreAudioKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

