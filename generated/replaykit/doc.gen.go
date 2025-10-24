
// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

// Package replaykit provides Go bindings for the ReplayKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ReplayKit without requiring cgo.
package replaykit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ReplayKit.framework/ReplayKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

