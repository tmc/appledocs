// Code generated from Apple documentation for GeoToolbox. DO NOT EDIT.

// Package geotoolbox provides Go bindings for the GeoToolbox framework.
//
// Determine place descriptor information for map coordinates.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to GeoToolbox without requiring cgo.

// Determine place descriptor information for map coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GeoToolbox

package geotoolbox

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/GeoToolbox.framework/GeoToolbox"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

