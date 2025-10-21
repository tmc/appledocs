// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

// Package corefoundation provides Go bindings for the CoreFoundation framework.
//
// Access low-level functions, primitive data types, and various collection types that are bridged seamlessly with the Foundation framework. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreFoundation without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation
package corefoundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


