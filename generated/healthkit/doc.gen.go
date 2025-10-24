
// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

// Package healthkit provides Go bindings for the HealthKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to HealthKit without requiring cgo.
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

