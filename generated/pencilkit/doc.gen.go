// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

// Package pencilkit provides Go bindings for the PencilKit framework.
//
// Capture touch and Apple Pencil input as a drawing, and display that content from your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PencilKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit
package pencilkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PencilKit.framework/PencilKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


