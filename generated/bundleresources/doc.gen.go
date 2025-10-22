// Code generated from Apple documentation for BundleResources. DO NOT EDIT.

// Package bundleresources provides Go bindings for the BundleResources framework.
//
// Resources located in an app, framework, or plugin bundle.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to BundleResources without requiring cgo.

// Resources located in an app, framework, or plugin bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BundleResources

package bundleresources

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/BundleResources.framework/BundleResources"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

