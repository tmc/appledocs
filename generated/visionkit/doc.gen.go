// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

// Package visionkit provides Go bindings for the VisionKit framework.
//
// Identify and extract information in the environment using the device’s camera, or in images that your app displays. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to VisionKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/VisionKit
package visionkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/VisionKit.framework/VisionKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

