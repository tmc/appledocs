
// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

// Package coregraphics provides Go bindings for the CoreGraphics framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreGraphics without requiring cgo.
package coregraphics

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

