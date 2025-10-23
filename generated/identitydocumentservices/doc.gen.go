// Code generated from Apple documentation for IdentityDocumentServices. DO NOT EDIT.

// Package identitydocumentservices provides Go bindings for the IdentityDocumentServices framework.
//
// Share mobile documents using the Digital Credentials API.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IdentityDocumentServices without requiring cgo.

// Share mobile documents using the Digital Credentials API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityDocumentServices
package identitydocumentservices

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IdentityDocumentServices.framework/IdentityDocumentServices"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

