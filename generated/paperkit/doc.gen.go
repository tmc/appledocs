// Code generated from Apple documentation for PaperKit. DO NOT EDIT.

// Package paperkit provides Go bindings for the PaperKit framework.
//
// Add drawings, shapes, and a consistent markup experience to your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to PaperKit without requiring cgo.
//
// See: https://developer.apple.com/documentation/PaperKit
package paperkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/PaperKit.framework/PaperKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

