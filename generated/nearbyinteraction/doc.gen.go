// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

// Package nearbyinteraction provides Go bindings for the NearbyInteraction framework.
//
// Locate and interact with nearby devices using identifiers, distance, and direction. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to NearbyInteraction without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction
package nearbyinteraction

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/NearbyInteraction.framework/NearbyInteraction"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


