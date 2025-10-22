// Code generated from Apple documentation for CoreText. DO NOT EDIT.

// Package coretext provides Go bindings for the CoreText framework.
//
// Create text layouts, optimize font handling, and access font metrics and glyph data.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreText without requiring cgo.

// Create text layouts, optimize font handling, and access font metrics and glyph data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText

package coretext

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreText.framework/CoreText"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

