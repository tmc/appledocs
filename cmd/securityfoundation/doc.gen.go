// Code generated from Apple documentation for SecurityFoundation. DO NOT EDIT.

// Package securityfoundation provides Go bindings for the SecurityFoundation framework.
//
// Restrict a user’s access to particular features in your Mac app or daemon.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SecurityFoundation without requiring cgo.

// Restrict a user’s access to particular features in your Mac app or daemon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SecurityFoundation
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

