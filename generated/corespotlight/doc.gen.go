
// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

// Package corespotlight provides Go bindings for the CoreSpotlight framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreSpotlight without requiring cgo.
package corespotlight

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreSpotlight.framework/CoreSpotlight"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

