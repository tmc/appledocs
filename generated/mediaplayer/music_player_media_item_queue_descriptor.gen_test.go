// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMusicPlayerMediaItemQueueDescriptor

// ExampleNewMusicPlayerMediaItemQueueDescriptorWithItemCollection demonstrates how to create a MusicPlayerMediaItemQueueDescriptor instance using NewMusicPlayerMediaItemQueueDescriptorWithItemCollection.
// Creates a new queue descriptor using the designated collection.
func ExampleNewMusicPlayerMediaItemQueueDescriptorWithItemCollection() {
	_ = mediaplayer.NewMusicPlayerMediaItemQueueDescriptorWithItemCollection(
		mediaplayer.MPMediaItemCollection{}, // itemCollection MPMediaItemCollection
	)
	// Output:
}
// ExampleNewMusicPlayerMediaItemQueueDescriptorWithQuery demonstrates how to create a MusicPlayerMediaItemQueueDescriptor instance using NewMusicPlayerMediaItemQueueDescriptorWithQuery.
// Creates a new queue descriptor using the designated query.
func ExampleNewMusicPlayerMediaItemQueueDescriptorWithQuery() {
	_ = mediaplayer.NewMusicPlayerMediaItemQueueDescriptorWithQuery(
		mediaplayer.MPMediaQuery{}, // query MPMediaQuery
	)
	// Output:
}
