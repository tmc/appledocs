// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioConverter

// ExampleAudioConverter_Reset demonstrates using Reset on a AudioConverter instance.
// Resets the converter so you can convert a new audio stream.
func ExampleAudioConverter_Reset() {
	obj := avfaudio.NewAudioConverter()
	obj.Reset()
	// Output:
	}

