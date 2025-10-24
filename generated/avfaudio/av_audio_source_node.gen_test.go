// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioSourceNode

// ExampleNewAudioSourceNodeWithRenderBlock demonstrates how to create a AudioSourceNode instance using NewAudioSourceNodeWithRenderBlock.
// Creates an audio source node with a block that supplies audio data.
func ExampleNewAudioSourceNodeWithRenderBlock() {
	_ = avfaudio.NewAudioSourceNodeWithRenderBlock(
		avfaudio.AudioSourceNodeRenderBlock /* not a class type */{}, // block AudioSourceNodeRenderBlock /* not a class type */
	)
	// Output:
}
