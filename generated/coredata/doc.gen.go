// Code generated from Apple documentation for CoreData. DO NOT EDIT.

// Package coredata provides Go bindings for the CoreData framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreData without requiring cgo.
package coredata

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreData.framework/CoreData"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

