// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

// Package linkpresentation provides Go bindings for the LinkPresentation framework.
//
// Fetch, provide, and present rich links in your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to LinkPresentation without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation
package linkpresentation

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/LinkPresentation.framework/LinkPresentation"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


