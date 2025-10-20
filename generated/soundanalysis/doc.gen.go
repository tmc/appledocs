// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

// Package soundanalysis provides Go bindings for the SoundAnalysis framework.
//
// Classify various sounds by analyzing audio files or streams. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to SoundAnalysis without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis
package soundanalysis

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/SoundAnalysis.framework/SoundAnalysis"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


