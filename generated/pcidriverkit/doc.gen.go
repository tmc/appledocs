// Code generated from Apple documentation for PCIDriverKit. DO NOT EDIT.

// Package pcidriverkit provides Go bindings for the PCIDriverKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PCIDriverKit without requiring cgo.
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
