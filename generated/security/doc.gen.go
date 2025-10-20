// Code generated from Apple documentation for Security. DO NOT EDIT.

// Package security provides Go bindings for the Security framework.
//
// Secure the data your app manages, and control access to your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Security without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Security
package security

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Security.framework/Security"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


