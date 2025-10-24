// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

// Package coremedia provides Go bindings for the CoreMedia framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreMedia without requiring cgo.
package coremedia

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreMedia.framework/CoreMedia"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

