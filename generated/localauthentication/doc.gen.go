// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

// Package localauthentication provides Go bindings for the LocalAuthentication framework.
//
// Authenticate users biometrically or with a passphrase they already know. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to LocalAuthentication without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication
package localauthentication

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/LocalAuthentication.framework/LocalAuthentication"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


