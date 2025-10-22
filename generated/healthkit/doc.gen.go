// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

// Package healthkit provides Go bindings for the HealthKit framework.
//
// Access and share health and fitness data while maintaining the user’s privacy and control.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to HealthKit without requiring cgo.

// Access and share health and fitness data while maintaining the user’s privacy and control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit

package healthkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/HealthKit.framework/HealthKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

