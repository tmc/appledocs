// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

// Package virtualization provides Go bindings for the Virtualization framework.
//
// Create virtual machines and run macOS and Linux-based operating systems. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Virtualization without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization
package virtualization

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Virtualization.framework/Virtualization"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


