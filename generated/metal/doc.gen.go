
// Code generated from Apple documentation for Metal. DO NOT EDIT.

// Package metal provides Go bindings for the Metal framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Metal without requiring cgo.
package metal

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Metal.framework/Metal"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

