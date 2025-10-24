// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioPCMBuffer

// ExampleNewAudioPCMBufferWithPCMFormatFrameCapacity demonstrates how to create a AudioPCMBuffer instance using NewAudioPCMBufferWithPCMFormatFrameCapacity.
// Creates a PCM audio buffer instance for PCM audio data.
func ExampleNewAudioPCMBufferWithPCMFormatFrameCapacity() {
	_ = avfaudio.NewAudioPCMBufferWithPCMFormatFrameCapacity(
		avfaudio.AVAudioFormat{},   // format AVAudioFormat
		avfaudio.AudioFrameCount{}, // frameCapacity AudioFrameCount
	)
	// Output:
}
