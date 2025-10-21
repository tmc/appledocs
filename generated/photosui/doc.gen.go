// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

// Package photosui provides Go bindings for the PhotosUI framework.
//
// Present a person’s photo library using a picker interface, display Live Photos, or extend the Photos app with custom functionality. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PhotosUI without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI
package photosui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PhotosUI.framework/PhotosUI"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


