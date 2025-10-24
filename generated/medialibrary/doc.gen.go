// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

// Package medialibrary provides Go bindings for the MediaLibrary framework.
//
// Access read-only collections of the user’s multimedia content.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MediaLibrary without requiring cgo.
//
// See: https://developer.apple.com/documentation/MediaLibrary
package medialibrary

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MediaLibrary.framework/MediaLibrary"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

