// Code generated from Apple documentation for LocalAuthenticationEmbeddedUI. DO NOT EDIT.

// Package localauthenticationembeddedui provides Go bindings for the LocalAuthenticationEmbeddedUI framework.
//
// Present a standard local authentication view icon in a custom authentication view.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to LocalAuthenticationEmbeddedUI without requiring cgo.
//
// See: https://developer.apple.com/documentation/LocalAuthenticationEmbeddedUI
package localauthenticationembeddedui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/LocalAuthenticationEmbeddedUI.framework/LocalAuthenticationEmbeddedUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

