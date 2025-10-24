
// Code generated from Apple documentation for AppIntents. DO NOT EDIT.

// Package appintents provides Go bindings for the AppIntents framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AppIntents without requiring cgo.
package appintents

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppIntents.framework/AppIntents"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

