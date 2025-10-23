// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

// Package dispatch provides Go bindings for the Dispatch framework.
//
// Execute code concurrently on multicore hardware by submitting work to dispatch queues managed by the system.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Dispatch without requiring cgo.

// Execute code concurrently on multicore hardware by submitting work to dispatch queues managed by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch
package dispatch

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Dispatch.framework/Dispatch"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

