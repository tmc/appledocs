// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

// Package modelio provides Go bindings for the ModelIO framework.
//
// Import, export, and manipulate 3D models using a common infrastructure that integrates MetalKit, GLKit, and SceneKit. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ModelIO without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO
package modelio

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ModelIO.framework/ModelIO"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


