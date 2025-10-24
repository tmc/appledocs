// Code generated from Apple documentation for Combine. DO NOT EDIT.

// Package combine provides Go bindings for the Combine framework.
//
// Customize handling of asynchronous events by combining event-processing operators.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Combine without requiring cgo.
//
// See: https://developer.apple.com/documentation/Combine
package combine

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Combine.framework/Combine"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

