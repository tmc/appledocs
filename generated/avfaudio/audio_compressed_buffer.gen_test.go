// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioCompressedBuffer

// ExampleNewAudioCompressedBufferWithFormatPacketCapacity demonstrates how to create a AudioCompressedBuffer instance using NewAudioCompressedBufferWithFormatPacketCapacity.
// Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
func ExampleNewAudioCompressedBufferWithFormatPacketCapacity() {
	_ = avfaudio.NewAudioCompressedBufferWithFormatPacketCapacity(
		avfaudio.AVAudioFormat{}, // format AVAudioFormat
		avfaudio.AudioPacketCount{}, // packetCapacity AudioPacketCount
	)
	// Output:
}
// ExampleNewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize demonstrates how to create a AudioCompressedBuffer instance using NewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize.
// Creates a buffer that contains audio data in a compressed state.
func ExampleNewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize() {
	_ = avfaudio.NewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize(
		avfaudio.AVAudioFormat{}, // format AVAudioFormat
		avfaudio.AudioPacketCount{}, // packetCapacity AudioPacketCount
		0, // maximumPacketSize int
	)
	// Output:
}
