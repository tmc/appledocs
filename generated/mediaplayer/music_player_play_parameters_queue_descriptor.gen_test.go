// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMusicPlayerPlayParametersQueueDescriptor

// ExampleNewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue demonstrates how to create a MusicPlayerPlayParametersQueueDescriptor instance using NewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue.
// Creates a new queue descriptor using the designated queue of play parameters.
func ExampleNewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue() {
	_ = mediaplayer.NewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue(
		[]mediaplayer.MusicPlayerPlayParameters{}, // playParametersQueue []MusicPlayerPlayParameters
	)
	// Output:
}
