// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

// Package coremotion provides Go bindings for the CoreMotion framework.
//
// Process accelerometer, gyroscope, pedometer, and environment-related events.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreMotion without requiring cgo.

// Process accelerometer, gyroscope, pedometer, and environment-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion
package coremotion

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreMotion.framework/CoreMotion"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

