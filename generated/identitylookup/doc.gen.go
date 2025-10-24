
// Code generated from Apple documentation for IdentityLookup. DO NOT EDIT.

// Package identitylookup provides Go bindings for the IdentityLookup framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IdentityLookup without requiring cgo.
package identitylookup

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IdentityLookup.framework/IdentityLookup"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

