// Code generated from Apple documentation for GLKit. DO NOT EDIT.

// Package glkit provides Go bindings for the GLKit framework.
//
// Speed up OpenGL ES or OpenGL app development. Use math libraries, background texture loading, pre-created shader effects, and a standard view and view controller to implement your rendering loop. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GLKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit
package glkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GLKit.framework/GLKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

