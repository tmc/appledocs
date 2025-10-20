// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

// Package securityinterface provides Go bindings for the SecurityInterface framework.
//
// Provide user interface elements for security features such as authorization, access to digital certificates, and access to items in keychains. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SecurityInterface without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface
package securityinterface

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SecurityInterface.framework/SecurityInterface"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


