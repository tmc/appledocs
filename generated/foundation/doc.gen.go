// Code generated from Apple documentation for Foundation. DO NOT EDIT.

// Package foundation provides Go bindings for the Foundation framework.
//
// Access essential data types, collections, and operating-system services to define the base layer of functionality for your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Foundation without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation
package foundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Foundation.framework/Foundation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

