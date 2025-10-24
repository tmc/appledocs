
// Code generated from Apple documentation for ThreadNetwork. DO NOT EDIT.

// Package threadnetwork provides Go bindings for the ThreadNetwork framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ThreadNetwork without requiring cgo.
package threadnetwork

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ThreadNetwork.framework/ThreadNetwork"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

