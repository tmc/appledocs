// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

// Package avrouting provides Go bindings for the AVRouting framework.
//
// Display custom destinations to stream media in the system route picker. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AVRouting without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AVRouting
package avrouting

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVRouting.framework/AVRouting"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

