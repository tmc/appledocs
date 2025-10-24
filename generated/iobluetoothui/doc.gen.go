// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

// Package iobluetoothui provides Go bindings for the IOBluetoothUI framework.
//
// Present an interface through which users can pair their devices with other Bluetooth devices.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IOBluetoothUI without requiring cgo.
//
// See: https://developer.apple.com/documentation/IOBluetoothUI
package iobluetoothui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IOBluetoothUI.framework/IOBluetoothUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

