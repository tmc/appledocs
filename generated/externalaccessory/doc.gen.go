// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

// Package externalaccessory provides Go bindings for the ExternalAccessory framework.
//
// Communicate with accessories that connect to a device with the Apple Lightning connector, or with Bluetooth wireless technology. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExternalAccessory without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory
package externalaccessory

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExternalAccessory.framework/ExternalAccessory"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

