
// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

// Package mediaextension provides Go bindings for the MediaExtension framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MediaExtension without requiring cgo.
package mediaextension

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MediaExtension.framework/MediaExtension"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

