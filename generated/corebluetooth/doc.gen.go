// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

// Package corebluetooth provides Go bindings for the CoreBluetooth framework.
//
// Communicate with Bluetooth low energy and BR/EDR (“Classic”) Devices. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreBluetooth without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth
package corebluetooth

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreBluetooth.framework/CoreBluetooth"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


