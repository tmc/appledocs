// Code generated from Apple documentation for MatterSupport. DO NOT EDIT.

// Package mattersupport provides Go bindings for the MatterSupport framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MatterSupport without requiring cgo.
package mattersupport

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MatterSupport.framework/MatterSupport"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
