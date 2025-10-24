// Code generated from Apple documentation for RealityKit. DO NOT EDIT.

// Package realitykit provides Go bindings for the RealityKit framework.
//
// Simulate and render 3D content for use in your augmented reality apps.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to RealityKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/RealityKit
package realitykit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/RealityKit.framework/RealityKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

