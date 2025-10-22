// Code generated from Apple documentation for DeveloperToolsSupport. DO NOT EDIT.

// Package developertoolssupport provides Go bindings for the DeveloperToolsSupport framework.
//
// Expose custom views and modifiers in the Xcode library.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to DeveloperToolsSupport without requiring cgo.

// Expose custom views and modifiers in the Xcode library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeveloperToolsSupport

package developertoolssupport

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DeveloperToolsSupport.framework/DeveloperToolsSupport"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

