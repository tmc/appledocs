// Code generated from Apple documentation for MusicKit. DO NOT EDIT.

// Package musickit provides Go bindings for the MusicKit framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MusicKit without requiring cgo.
package musickit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MusicKit.framework/MusicKit"

func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}
