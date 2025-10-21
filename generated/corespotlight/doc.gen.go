// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

// Package corespotlight provides Go bindings for the CoreSpotlight framework.
//
// Add search capabilities to your app, and index your content so [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreSpotlight without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight
package corespotlight

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreSpotlight.framework/CoreSpotlight"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


