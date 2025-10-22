// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

// Package accessibility provides Go bindings for the Accessibility framework.
//
// Make your apps accessible to everyone who uses Apple devices.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Accessibility without requiring cgo.

// Make your apps accessible to everyone who uses Apple devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility

package accessibility

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Accessibility.framework/Accessibility"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

