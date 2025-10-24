
// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

// Package cfnetwork provides Go bindings for the CFNetwork framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CFNetwork without requiring cgo.
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

