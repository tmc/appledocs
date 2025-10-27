// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

// Package coreaudiotypes provides Go bindings for the CoreAudioTypes framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to CoreAudioTypes without requiring cgo.
package coreaudiotypes

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreAudioTypes.framework/CoreAudioTypes"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

