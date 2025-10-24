
// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

// Package audiotoolbox provides Go bindings for the AudioToolbox framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AudioToolbox without requiring cgo.
package audiotoolbox

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AudioToolbox.framework/AudioToolbox"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

