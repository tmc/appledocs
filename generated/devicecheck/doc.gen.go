// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

// Package devicecheck provides Go bindings for the DeviceCheck framework.
//
// Reduce fraudulent use of your services by managing device state and asserting app
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DeviceCheck without requiring cgo.

// Reduce fraudulent use of your services by managing device state and asserting app
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck

package devicecheck

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DeviceCheck.framework/DeviceCheck"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

