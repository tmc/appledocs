
// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

// Package colorsync provides Go bindings for the ColorSync framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ColorSync without requiring cgo.
package colorsync

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ColorSync.framework/ColorSync"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

