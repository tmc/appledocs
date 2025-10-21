// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

// Package quartzcore provides Go bindings for the QuartzCore framework.
//
// Render, compose, and animate visual elements. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to QuartzCore without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore
package quartzcore

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/QuartzCore.framework/QuartzCore"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

