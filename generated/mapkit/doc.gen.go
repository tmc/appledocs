
// Code generated from Apple documentation for MapKit. DO NOT EDIT.

// Package mapkit provides Go bindings for the MapKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MapKit without requiring cgo.
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

