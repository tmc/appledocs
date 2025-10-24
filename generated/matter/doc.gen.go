
// Code generated from Apple documentation for Matter. DO NOT EDIT.

// Package matter provides Go bindings for the Matter framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Matter without requiring cgo.
package matter

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Matter.framework/Matter"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

