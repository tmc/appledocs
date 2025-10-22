// Code generated from Apple documentation for EventKit. DO NOT EDIT.

// Package eventkit provides Go bindings for the EventKit framework.
//
// Create, view, and edit calendar and reminder events.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to EventKit without requiring cgo.

// Create, view, and edit calendar and reminder events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit

package eventkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/EventKit.framework/EventKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

