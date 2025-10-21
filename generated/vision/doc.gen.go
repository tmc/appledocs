// Code generated from Apple documentation for Vision. DO NOT EDIT.

// Package vision provides Go bindings for the Vision framework.
//
// Apply computer vision algorithms to perform a variety of tasks on input images and videos. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Vision without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision
package vision

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Vision.framework/Vision"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

