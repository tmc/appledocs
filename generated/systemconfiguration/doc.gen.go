// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

// Package systemconfiguration provides Go bindings for the SystemConfiguration framework.
//
// Allow applications to access a device’s network configuration settings. Determine the reachability of the device, such as whether Wi-Fi or cell connectivity are active. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SystemConfiguration without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration
package systemconfiguration

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SystemConfiguration.framework/SystemConfiguration"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


