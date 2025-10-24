
// Code generated from Apple documentation for CoreText. DO NOT EDIT.

// Package coretext provides Go bindings for the CoreText framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreText without requiring cgo.
package coretext

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreText.framework/CoreText"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

