// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

// Package backgroundassets provides Go bindings for the BackgroundAssets framework.
//
// Improve or eliminate the time people wait while your app downloads assets.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to BackgroundAssets without requiring cgo.

// Improve or eliminate the time people wait while your app downloads assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets
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

