// Code generated from Apple documentation for AudioUnit. DO NOT EDIT.

// Package audiounit provides Go bindings for the AudioUnit framework.
//
// Add sophisticated audio manipulation and processing capabilities to your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AudioUnit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioUnit
package audiounit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AudioUnit.framework/AudioUnit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


