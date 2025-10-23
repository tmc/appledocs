// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

// Package hypervisor provides Go bindings for the Hypervisor framework.
//
// Build virtualization solutions on top of a lightweight hypervisor, without third-party kernel extensions.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Hypervisor without requiring cgo.

// Build virtualization solutions on top of a lightweight hypervisor, without third-party kernel extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Hypervisor
package hypervisor

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Hypervisor.framework/Hypervisor"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

