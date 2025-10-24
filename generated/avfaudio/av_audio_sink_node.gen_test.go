// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioSinkNode

// ExampleNewAudioSinkNodeWithReceiverBlock demonstrates how to create a AudioSinkNode instance using NewAudioSinkNodeWithReceiverBlock.
// Creates an audio sink node with a block that receives audio data.
func ExampleNewAudioSinkNodeWithReceiverBlock() {
	_ = avfaudio.NewAudioSinkNodeWithReceiverBlock(
		avfaudio.AudioSinkNodeReceiverBlock /* not a class type */{}, // block AudioSinkNodeReceiverBlock /* not a class type */
	)
	// Output:
}
