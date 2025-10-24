// Code generated from Apple documentation for IdentityLookupUI. DO NOT EDIT.

// Package identitylookupui provides Go bindings for the IdentityLookupUI framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IdentityLookupUI without requiring cgo.
//
// See: https://developer.apple.com/documentation/IdentityLookupUI
package identitylookupui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IdentityLookupUI.framework/IdentityLookupUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

