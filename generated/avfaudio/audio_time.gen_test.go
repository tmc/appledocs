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

// ExampleNewAudioTimeWithHostTimeSampleTimeAtRate demonstrates how to create a AudioTime instance using NewAudioTimeWithHostTimeSampleTimeAtRate.
// Creates an audio time object with the specified host time, sample time, and sample rate.
func ExampleNewAudioTimeWithHostTimeSampleTimeAtRate() {
	_ = avfaudio.NewAudioTimeWithHostTimeSampleTimeAtRate(
		0,                             // hostTime uint64
		avfaudio.AudioFramePosition{}, // sampleTime AudioFramePosition
		0.0,                           // sampleRate float64
	)
	// Output:
}

// ExampleNewAudioTimeWithSampleTimeAtRate demonstrates how to create a AudioTime instance using NewAudioTimeWithSampleTimeAtRate.
// Creates an audio time object with the specified timestamp and sample rate.
func ExampleNewAudioTimeWithSampleTimeAtRate() {
	_ = avfaudio.NewAudioTimeWithSampleTimeAtRate(
		avfaudio.AudioFramePosition{}, // sampleTime AudioFramePosition
		0.0,                           // sampleRate float64
	)
	// Output:
}
