// Code generated from Apple documentation for MediaToolbox. DO NOT EDIT.

// Package mediatoolbox provides Go bindings for the MediaToolbox framework.
//
// Enable support for media format readers; tap and process audio from an audio mix.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MediaToolbox without requiring cgo.
//
// See: https://developer.apple.com/documentation/MediaToolbox
package mediatoolbox

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MediaToolbox.framework/MediaToolbox"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

