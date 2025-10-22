// Code generated from Apple documentation for ExtensionKit. DO NOT EDIT.

// Package extensionkit provides Go bindings for the ExtensionKit framework.
//
// Make custom UI from an app extension available in a host app, and manage the list of enabled
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExtensionKit without requiring cgo.

// Make custom UI from an app extension available in a host app, and manage the list of enabled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExtensionKit

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

