// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioTime

// ExampleNewAudioTimeWithHostTime demonstrates how to create a AudioTime instance using NewAudioTimeWithHostTime.
// Creates an audio time object with the specified host time.
func ExampleNewAudioTimeWithHostTime() {
	_ = avfaudio.NewAudioTimeWithHostTime(
		0, // hostTime uint64
	)
	// Output:
}
