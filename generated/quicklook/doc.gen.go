
// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

// Package quicklook provides Go bindings for the QuickLook framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to QuickLook without requiring cgo.
package quicklook

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/QuickLook.framework/QuickLook"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

