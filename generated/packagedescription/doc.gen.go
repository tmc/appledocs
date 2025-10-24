// Code generated from Apple documentation for PackageDescription. DO NOT EDIT.

// Package packagedescription provides Go bindings for the PackageDescription framework.
//
// Create reusable code, organize it in a lightweight way, and share it across your projects and with other developers.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PackageDescription without requiring cgo.
//
// See: https://developer.apple.com/documentation/PackageDescription
package packagedescription

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PackageDescription.framework/PackageDescription"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

