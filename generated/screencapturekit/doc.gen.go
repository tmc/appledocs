// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

// Package screencapturekit provides Go bindings for the ScreenCaptureKit framework.
//
// Filter and select screen content and stream it to your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScreenCaptureKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit
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

