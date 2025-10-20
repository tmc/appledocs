// Code generated from Apple documentation for AppKit. DO NOT EDIT.

// Package appkit provides Go bindings for the AppKit framework.
//
// Construct and manage a graphical, event-driven user interface for your macOS app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit
package appkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppKit.framework/AppKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
