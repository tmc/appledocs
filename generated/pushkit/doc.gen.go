// Code generated from Apple documentation for PushKit. DO NOT EDIT.

// Package pushkit provides Go bindings for the PushKit framework.
//
// Respond to push notifications related to your app’s complications, file providers, and VoIP services.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PushKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/PushKit
package pushkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PushKit.framework/PushKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

