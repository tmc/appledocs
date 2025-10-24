
// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

// Package coreimage provides Go bindings for the CoreImage framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreImage without requiring cgo.
package coreimage

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreImage.framework/CoreImage"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

