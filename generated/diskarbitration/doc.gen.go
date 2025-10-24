// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

// Package diskarbitration provides Go bindings for the DiskArbitration framework.
//
// Provides mechanisms to register and block disk mount or unmount events.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DiskArbitration without requiring cgo.
//
// See: https://developer.apple.com/documentation/DiskArbitration
package diskarbitration

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DiskArbitration.framework/DiskArbitration"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

