// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

// Package objectivec provides Go bindings for the ObjectiveC framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ObjectiveC without requiring cgo.
package objectivec

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/usr/lib/libobjc.dylib"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

