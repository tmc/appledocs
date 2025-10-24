// Code generated from Apple documentation for DeclaredAgeRange. DO NOT EDIT.

// Package declaredagerange provides Go bindings for the DeclaredAgeRange framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DeclaredAgeRange without requiring cgo.
package declaredagerange

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DeclaredAgeRange.framework/DeclaredAgeRange"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
