
// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

// Package mediaaccessibility provides Go bindings for the MediaAccessibility framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MediaAccessibility without requiring cgo.
package mediaaccessibility

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MediaAccessibility.framework/MediaAccessibility"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

