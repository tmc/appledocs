
// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

// Package mlcompute provides Go bindings for the MLCompute framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MLCompute without requiring cgo.
package mlcompute

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MLCompute.framework/MLCompute"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

