
// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

// Package cinematic provides Go bindings for the Cinematic framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Cinematic without requiring cgo.
package cinematic

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Cinematic.framework/Cinematic"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

