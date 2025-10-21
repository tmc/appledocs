// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

// Package exceptionhandling provides Go bindings for the ExceptionHandling framework.
//
// Monitor and debug exceptional conditions in code. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExceptionHandling without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling
package exceptionhandling

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExceptionHandling.framework/ExceptionHandling"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


