// Code generated from Apple documentation for Photos. DO NOT EDIT.

// Package photos provides Go bindings for the Photos framework.
//
// Work with image and video assets that the Photos app manages, including those from iCloud Photos and Live Photos.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Photos without requiring cgo.
//
// See: https://developer.apple.com/documentation/Photos
package photos

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Photos.framework/Photos"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

