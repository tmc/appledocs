
// Code generated from Apple documentation for AVKit. DO NOT EDIT.

// Package avkit provides Go bindings for the AVKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AVKit without requiring cgo.
package avkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVKit.framework/AVKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

