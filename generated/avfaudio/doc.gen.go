// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

// Package avfaudio provides Go bindings for the AVFAudio framework.
//
// Play, record, and process audio; configure your app’s system audio behavior. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AVFAudio without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio
package avfaudio

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVFAudio.framework/AVFAudio"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

