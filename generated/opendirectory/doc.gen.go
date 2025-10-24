// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

// Package opendirectory provides Go bindings for the OpenDirectory framework.
//
// Authenticate users, and search for contact information in Open Directory and LDAP directories.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to OpenDirectory without requiring cgo.
//
// See: https://developer.apple.com/documentation/OpenDirectory
package opendirectory

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/OpenDirectory.framework/OpenDirectory"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

