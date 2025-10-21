// Code generated from Apple documentation for AppTrackingTransparency. DO NOT EDIT.

// Package apptrackingtransparency provides Go bindings for the AppTrackingTransparency framework.
//
// Request user authorization to access app-related data for tracking the user or the [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppTrackingTransparency without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency
package apptrackingtransparency

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppTrackingTransparency.framework/AppTrackingTransparency"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

