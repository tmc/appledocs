// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

// Package coremediaio provides Go bindings for the CoreMediaIO framework.
//
// Securely support custom camera devices in macOS. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreMediaIO without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO
package coremediaio

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreMediaIO.framework/CoreMediaIO"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


