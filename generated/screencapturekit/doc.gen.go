
// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

// Package screencapturekit provides Go bindings for the ScreenCaptureKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScreenCaptureKit without requiring cgo.
package screencapturekit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

