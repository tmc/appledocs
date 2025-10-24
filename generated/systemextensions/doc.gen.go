// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

// Package systemextensions provides Go bindings for the SystemExtensions framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SystemExtensions without requiring cgo.
package systemextensions

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SystemExtensions.framework/SystemExtensions"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
