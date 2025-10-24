// Code generated from Apple documentation for Distributed. DO NOT EDIT.

// Package distributed provides Go bindings for the Distributed framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Distributed without requiring cgo.
package distributed

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Distributed.framework/Distributed"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
