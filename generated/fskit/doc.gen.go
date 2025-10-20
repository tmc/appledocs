// Code generated from Apple documentation for FSKit. DO NOT EDIT.

// Package fskit provides Go bindings for the FSKit framework.
//
// Implement a file system that runs in user space. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to FSKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit
package fskit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/FSKit.framework/FSKit"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


