
// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

// Package driverkit provides Go bindings for the DriverKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DriverKit without requiring cgo.
package driverkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DriverKit.framework/DriverKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

