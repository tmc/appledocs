// Code generated from Apple documentation for ImagePlayground. DO NOT EDIT.

// Package imageplayground provides Go bindings for the ImagePlayground framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ImagePlayground without requiring cgo.
package imageplayground

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ImagePlayground.framework/ImagePlayground"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
