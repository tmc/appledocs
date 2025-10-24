// Code generated from Apple documentation for CoreHID. DO NOT EDIT.

// Package corehid provides Go bindings for the CoreHID framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreHID without requiring cgo.
package corehid

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreHID.framework/CoreHID"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
