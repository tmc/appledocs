// Code generated from Apple documentation for PCIDriverKit. DO NOT EDIT.

// Package pcidriverkit provides Go bindings for the PCIDriverKit framework.
//
// Develop device drivers for Peripheral Component Interconnect (PCI) accessories. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PCIDriverKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit
package pcidriverkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PCIDriverKit.framework/PCIDriverKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


