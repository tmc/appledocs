// Code generated from Apple documentation for Network. DO NOT EDIT.

// Package network provides Go bindings for the Network framework.
//
// Create network connections to send and receive data using transport and security protocols. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Network without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Network
package network

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Network.framework/Network"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

