
// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

// Package pencilkit provides Go bindings for the PencilKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PencilKit without requiring cgo.
package pencilkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PencilKit.framework/PencilKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

