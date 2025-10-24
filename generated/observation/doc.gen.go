// Code generated from Apple documentation for Observation. DO NOT EDIT.

// Package observation provides Go bindings for the Observation framework.
//
// Make responsive apps that update the presentation when underlying data changes.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Observation without requiring cgo.
//
// See: https://developer.apple.com/documentation/Observation
package observation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Observation.framework/Observation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

