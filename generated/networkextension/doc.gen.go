// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

// Package networkextension provides Go bindings for the NetworkExtension framework.
//
// Customize and extend core networking features. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to NetworkExtension without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension
package networkextension

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/NetworkExtension.framework/NetworkExtension"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


