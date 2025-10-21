// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

// Package safetykit provides Go bindings for the SafetyKit framework.
//
// Detect and respond to car crash events in your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SafetyKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SafetyKit
package safetykit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SafetyKit.framework/SafetyKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


