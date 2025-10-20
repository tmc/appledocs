// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

// Package uniformtypeidentifiers provides Go bindings for the UniformTypeIdentifiers framework.
//
// Provide uniform type identifiers that describe file types for storage or transfer. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to UniformTypeIdentifiers without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/UniformTypeIdentifiers
package uniformtypeidentifiers

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/UniformTypeIdentifiers.framework/UniformTypeIdentifiers"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


