
// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

// Package securityfoundation provides Go bindings for the SecurityFoundation framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SecurityFoundation without requiring cgo.
package securityfoundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SecurityFoundation.framework/SecurityFoundation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

