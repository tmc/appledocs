// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

// Package multipeerconnectivity provides Go bindings for the MultipeerConnectivity framework.
//
// Support peer-to-peer connectivity and the discovery of nearby devices. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MultipeerConnectivity without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity
package multipeerconnectivity

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MultipeerConnectivity.framework/MultipeerConnectivity"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


