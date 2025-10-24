
// Code generated from Apple documentation for Accelerate. DO NOT EDIT.

// Package accelerate provides Go bindings for the Accelerate framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Accelerate without requiring cgo.
package accelerate

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Accelerate.framework/Accelerate"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

