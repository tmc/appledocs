// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

// Package endpointsecurity provides Go bindings for the EndpointSecurity framework.
//
// Develop system extensions that enhance user security.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to EndpointSecurity without requiring cgo.

// Develop system extensions that enhance user security.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity
package endpointsecurity

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/EndpointSecurity.framework/EndpointSecurity"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

