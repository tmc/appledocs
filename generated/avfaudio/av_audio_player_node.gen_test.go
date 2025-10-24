// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioPlayerNode

// ExampleNewAudioPlayerNode demonstrates how to create a AudioPlayerNode instance.
// Creates an initialized audio player node.
func ExampleNewAudioPlayerNode() {
	_ = avfaudio.NewAudioPlayerNode()
	// Output:
}

// ExampleAudioPlayerNode_Pause demonstrates using Pause on a AudioPlayerNode instance.
// Pauses the node’s playback.
func ExampleAudioPlayerNode_Pause() {
	obj := avfaudio.NewAudioPlayerNode()
	obj.Pause()
	// Output:
}

// ExampleAudioPlayerNode_Play demonstrates using Play on a AudioPlayerNode instance.
// Starts or resumes playback immediately.
func ExampleAudioPlayerNode_Play() {
	obj := avfaudio.NewAudioPlayerNode()
	obj.Play()
	// Output:
}

// ExampleAudioPlayerNode_Stop demonstrates using Stop on a AudioPlayerNode instance.
// Clears all of the node’s events you schedule and stops playback.
func ExampleAudioPlayerNode_Stop() {
	obj := avfaudio.NewAudioPlayerNode()
	obj.Stop()
	// Output:
}
