// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioFile

// ExampleNewAudioFile demonstrates how to create a AudioFile instance.
func ExampleNewAudioFile() {
	_ = avfaudio.NewAudioFile()
	// Output:
}
// ExampleAudioFile_Close demonstrates using Close on a AudioFile instance.
// Closes the audio file.
func ExampleAudioFile_Close() {
	obj := avfaudio.NewAudioFile()
	obj.Close()
	// Output:
	}

