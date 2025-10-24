// Code generated from Apple documentation for MapKit. DO NOT EDIT.

// Package mapkit provides Go bindings for the MapKit framework.
//
// Display map or satellite imagery within your app, call out points of interest, and determine placemark information for map coordinates.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MapKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/MapKit
package mapkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MapKit.framework/MapKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

