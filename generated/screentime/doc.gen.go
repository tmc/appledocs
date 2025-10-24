
// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

// Package screentime provides Go bindings for the ScreenTime framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ScreenTime without requiring cgo.
package screentime

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ScreenTime.framework/ScreenTime"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

