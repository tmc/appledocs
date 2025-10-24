
// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

// Package corevideo provides Go bindings for the CoreVideo framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreVideo without requiring cgo.
package corevideo

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreVideo.framework/CoreVideo"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

