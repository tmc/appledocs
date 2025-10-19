// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

// Package coreaudio provides Go bindings for the CoreAudio framework.
//
// Use the Core Audio framework to interact with device’s audio hardware. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreAudio without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio
package coreaudio

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreAudio.framework/CoreAudio"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


