
// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

// Package accessibility provides Go bindings for the Accessibility framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to Accessibility without requiring cgo.
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

