// Code generated from Apple documentation for AdServices. DO NOT EDIT.

// Package adservices provides Go bindings for the AdServices framework.
//
// Attribute app-download campaigns that originate from the App Store, Apple News, or [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AdServices without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AdServices
package adservices

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AdServices.framework/AdServices"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

