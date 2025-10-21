// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

// Package corelocation provides Go bindings for the CoreLocation framework.
//
// Obtain the geographic location and orientation of a device. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreLocation without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation
package corelocation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreLocation.framework/CoreLocation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

