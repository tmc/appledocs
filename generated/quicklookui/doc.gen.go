
// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

// Package quicklookui provides Go bindings for the QuickLookUI framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to QuickLookUI without requiring cgo.
package quicklookui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/QuickLookUI.framework/QuickLookUI"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

