
// Code generated from Apple documentation for ExtensionKit. DO NOT EDIT.

// Package extensionkit provides Go bindings for the ExtensionKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExtensionKit without requiring cgo.
package extensionkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExtensionKit.framework/ExtensionKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

