// Code generated from Apple documentation for PhotoKit. DO NOT EDIT.

// Package photokit provides Go bindings for the PhotoKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PhotoKit without requiring cgo.
package photokit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PhotoKit.framework/PhotoKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
