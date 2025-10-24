
// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

// Package storekittest provides Go bindings for the StoreKitTest framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to StoreKitTest without requiring cgo.
package storekittest

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/StoreKitTest.framework/StoreKitTest"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

