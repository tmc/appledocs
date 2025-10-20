// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

// Package authenticationservices provides Go bindings for the AuthenticationServices framework.
//
// Make it easy for users to log into apps and services. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AuthenticationServices without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices
package authenticationservices

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AuthenticationServices.framework/AuthenticationServices"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


