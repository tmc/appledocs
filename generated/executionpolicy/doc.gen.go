// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

// Package executionpolicy provides Go bindings for the ExecutionPolicy framework.
//
// Provide functionality so developer tools can manage execution policy exceptions.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExecutionPolicy without requiring cgo.

// Provide functionality so developer tools can manage execution policy exceptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy
package executionpolicy

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExecutionPolicy.framework/ExecutionPolicy"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

