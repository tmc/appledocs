// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

// Package iobluetooth provides Go bindings for the IOBluetooth framework.
//
// Gain user-space access to Bluetooth devices.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IOBluetooth without requiring cgo.
//
// See: https://developer.apple.com/documentation/IOBluetooth
package iobluetooth

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IOBluetooth.framework/IOBluetooth"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

