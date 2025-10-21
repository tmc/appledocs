// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

// Package coremidi provides Go bindings for the CoreMIDI framework.
//
// Communicate with MIDI devices such as hardware keyboards and synthesizers. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreMIDI without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI
package coremidi

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreMIDI.framework/CoreMIDI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

