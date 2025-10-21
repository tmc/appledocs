// Code generated from Apple documentation for EnterpriseProgramAPI. DO NOT EDIT.

// Package enterpriseprogramapi provides Go bindings for the EnterpriseProgramAPI framework.
//
// Automate the tasks you perform on the Apple Developer website. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to EnterpriseProgramAPI without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/EnterpriseProgramAPI
package enterpriseprogramapi

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/EnterpriseProgramAPI.framework/EnterpriseProgramAPI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

