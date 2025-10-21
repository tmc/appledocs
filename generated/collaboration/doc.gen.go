// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

// Package collaboration provides Go bindings for the Collaboration framework.
//
// Find and access identities, that is, users and groups. Display the Identity Picker, which lets users create and select identities. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Collaboration without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration
package collaboration

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Collaboration.framework/Collaboration"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

