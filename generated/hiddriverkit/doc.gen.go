
// Code generated from Apple documentation for HIDDriverKit. DO NOT EDIT.

// Package hiddriverkit provides Go bindings for the HIDDriverKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to HIDDriverKit without requiring cgo.
package hiddriverkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/HIDDriverKit.framework/HIDDriverKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

