
// Code generated from Apple documentation for CallKit. DO NOT EDIT.

// Package callkit provides Go bindings for the CallKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CallKit without requiring cgo.
package callkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CallKit.framework/CallKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

