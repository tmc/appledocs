// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioFormat

// ExampleNewAudioFormatStandardFormatWithSampleRateChannels demonstrates how to create a AudioFormat instance using NewAudioFormatStandardFormatWithSampleRateChannels.
// Creates an audio format instance with the specified sample rate and channel count.
func ExampleNewAudioFormatStandardFormatWithSampleRateChannels() {
	_ = avfaudio.NewAudioFormatStandardFormatWithSampleRateChannels(
		0.0, // sampleRate float64
		avfaudio.AudioChannelCount /* typedef */{}, // channels AudioChannelCount /* typedef */
	)
	// Output:
}
// ExampleNewAudioFormatWithCMAudioFormatDescription demonstrates how to create a AudioFormat instance using NewAudioFormatWithCMAudioFormatDescription.
// Creates an audio format instance from a Core Media audio format description.
func ExampleNewAudioFormatWithCMAudioFormatDescription() {
	_ = avfaudio.NewAudioFormatWithCMAudioFormatDescription(
		avfaudio.AudioFormatDescriptionRef /* not a class type */{}, // formatDescription AudioFormatDescriptionRef /* not a class type */
	)
	// Output:
}
// ExampleNewAudioFormatWithCommonFormatSampleRateChannelsInterleaved demonstrates how to create a AudioFormat instance using NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved.
// Creates an audio format instance.
func ExampleNewAudioFormatWithCommonFormatSampleRateChannelsInterleaved() {
	_ = avfaudio.NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved(
		avfaudio.AudioCommonFormat{}, // format AudioCommonFormat
		0.0, // sampleRate float64
		avfaudio.AudioChannelCount /* typedef */{}, // channels AudioChannelCount /* typedef */
		false, // interleaved bool
	)
	// Output:
}
