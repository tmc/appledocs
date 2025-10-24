// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

// Package mediaplayer provides Go bindings for the MediaPlayer framework.
//
// Find and play songs, audio podcasts, audio books, and more from within your app.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to MediaPlayer without requiring cgo.
//
// See: https://developer.apple.com/documentation/MediaPlayer
package mediaplayer

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/MediaPlayer.framework/MediaPlayer"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

