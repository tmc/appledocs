// Code generated from Apple documentation for PHASE. DO NOT EDIT.

// Package phase provides Go bindings for the PHASE framework.
//
// Create dynamic audio experiences in your game or app that react to events and cues in the environment.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PHASE without requiring cgo.
//
// See: https://developer.apple.com/documentation/PHASE
package phase

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PHASE.framework/PHASE"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

