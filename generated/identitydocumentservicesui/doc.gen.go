// Code generated from Apple documentation for IdentityDocumentServicesUI. DO NOT EDIT.

// Package identitydocumentservicesui provides Go bindings for the IdentityDocumentServicesUI framework.
//
// Provide an interface so people can present mobile documents.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IdentityDocumentServicesUI without requiring cgo.

// Provide an interface so people can present mobile documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IdentityDocumentServicesUI
package identitydocumentservicesui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IdentityDocumentServicesUI.framework/IdentityDocumentServicesUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

