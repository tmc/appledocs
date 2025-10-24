// Code generated from Apple documentation for DeviceManagement. DO NOT EDIT.

// Package devicemanagement provides Go bindings for the DeviceManagement framework.
//
// Manage your organization’s devices remotely.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DeviceManagement without requiring cgo.
//
// See: https://developer.apple.com/documentation/DeviceManagement
package devicemanagement

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DeviceManagement.framework/DeviceManagement"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

