// Code generated from Apple documentation for AVKit. DO NOT EDIT.

// Package avkit provides Go bindings for the AVKit framework.
//
// Create user interfaces for media playback, complete with transport controls, chapter navigation, picture-in-picture support, and display of subtitles and closed captions. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to AVKit without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit
package avkit

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVKit.framework/AVKit"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

