// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioUnitEQ

// ExampleNewAudioUnitEQWithNumberOfBands demonstrates how to create a AudioUnitEQ instance using NewAudioUnitEQWithNumberOfBands.
// Creates an audio unit equalizer object with the specified number of bands.
func ExampleNewAudioUnitEQWithNumberOfBands() {
	_ = avfaudio.NewAudioUnitEQWithNumberOfBands(
		0, // numberOfBands uint
	)
	// Output:
}
