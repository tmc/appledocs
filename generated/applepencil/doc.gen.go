// Code generated from Apple documentation for ApplePencil. DO NOT EDIT.

// Package applepencil provides Go bindings for the ApplePencil framework.
//
// Enhance your iPad app’s user experience by supporting drawing, handwriting, and other features of Apple Pencil.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ApplePencil without requiring cgo.
//
// See: https://developer.apple.com/documentation/ApplePencil
package applepencil

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ApplePencil.framework/ApplePencil"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

