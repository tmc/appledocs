
// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

// Package backgroundassets provides Go bindings for the BackgroundAssets framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to BackgroundAssets without requiring cgo.
package backgroundassets

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/BackgroundAssets.framework/BackgroundAssets"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

