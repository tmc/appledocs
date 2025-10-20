// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

// Package cfnetwork provides Go bindings for the CFNetwork framework.
//
// Access network services and handle changes in network configurations. Build on abstractions of network protocols to simplify tasks such as working with BSD sockets, administering HTTP and FTP servers, and managing Bonjour services. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CFNetwork without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork
package cfnetwork

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CFNetwork.framework/CFNetwork"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


