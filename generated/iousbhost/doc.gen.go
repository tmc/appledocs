// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

// Package iousbhost provides Go bindings for the IOUSBHost framework.
//
// Create host-mode user space drivers for USB devices.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IOUSBHost without requiring cgo.

// Create host-mode user space drivers for USB devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost
package iousbhost

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IOUSBHost.framework/IOUSBHost"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

