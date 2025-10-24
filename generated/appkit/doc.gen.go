// Code generated from Apple documentation for AppKit. DO NOT EDIT.

// Package appkit provides Go bindings for the AppKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppKit without requiring cgo.
package appkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppKit.framework/AppKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

