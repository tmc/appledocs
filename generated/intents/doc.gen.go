// Code generated from Apple documentation for Intents. DO NOT EDIT.

// Package intents provides Go bindings for the Intents framework.
//
// Empower people to customize interactions for your app on their device. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Intents without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents
package intents

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Intents.framework/Intents"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

