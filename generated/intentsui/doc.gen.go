// Code generated from Apple documentation for IntentsUI. DO NOT EDIT.

// Package intentsui provides Go bindings for the IntentsUI framework.
//
// Customize content in the interface for Siri and Maps. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to IntentsUI without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/IntentsUI
package intentsui

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/IntentsUI.framework/IntentsUI"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


