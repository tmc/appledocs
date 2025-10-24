
// Code generated from Apple documentation for PreferencePanes. DO NOT EDIT.

// Package preferencepanes provides Go bindings for the PreferencePanes framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PreferencePanes without requiring cgo.
package preferencepanes

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PreferencePanes.framework/PreferencePanes"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

