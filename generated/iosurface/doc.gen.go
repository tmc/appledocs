
// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

// Package iosurface provides Go bindings for the IOSurface framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IOSurface without requiring cgo.
package iosurface

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IOSurface.framework/IOSurface"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

