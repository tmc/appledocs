// Code generated from Apple documentation for Compression. DO NOT EDIT.

// Package compression provides Go bindings for the Compression framework.
//
// Leverage common compression algorithms for lossless data compression.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Compression without requiring cgo.
//
// See: https://developer.apple.com/documentation/Compression
package compression

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Compression.framework/Compression"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

