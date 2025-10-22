// Code generated from Apple documentation for ExtensionFoundation. DO NOT EDIT.

// Package extensionfoundation provides Go bindings for the ExtensionFoundation framework.
//
// Create executable bundles to extend the functionality of other apps.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ExtensionFoundation without requiring cgo.

// Create executable bundles to extend the functionality of other apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExtensionFoundation

package extensionfoundation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ExtensionFoundation.framework/ExtensionFoundation"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

