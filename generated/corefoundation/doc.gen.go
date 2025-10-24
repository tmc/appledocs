// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

// Package corefoundation provides Go bindings for the CoreFoundation framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreFoundation without requiring cgo.
package corefoundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

