// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

// Package devicediscoveryextension provides Go bindings for the DeviceDiscoveryExtension framework.
//
// Stream media to a third-party device that a user selects in a system menu. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DeviceDiscoveryExtension without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension
package devicediscoveryextension

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DeviceDiscoveryExtension.framework/DeviceDiscoveryExtension"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


