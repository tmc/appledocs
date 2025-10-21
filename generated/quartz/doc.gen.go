// Code generated from Apple documentation for Quartz. DO NOT EDIT.

// Package quartz provides Go bindings for the Quartz framework.
//
// Allow users to browse, edit, and save images, using slideshows and Core Image filters. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Quartz without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz
package quartz

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Quartz.framework/Quartz"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

