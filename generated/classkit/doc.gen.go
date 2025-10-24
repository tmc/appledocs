
// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

// Package classkit provides Go bindings for the ClassKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ClassKit without requiring cgo.
package classkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ClassKit.framework/ClassKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

