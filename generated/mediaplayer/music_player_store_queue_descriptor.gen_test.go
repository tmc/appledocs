// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMusicPlayerStoreQueueDescriptor

// ExampleNewMusicPlayerStoreQueueDescriptorWithStoreIDs demonstrates how to create a MusicPlayerStoreQueueDescriptor instance using NewMusicPlayerStoreQueueDescriptorWithStoreIDs.
// Creates a new queue descriptor using the designated store identifiers.
func ExampleNewMusicPlayerStoreQueueDescriptorWithStoreIDs() {
	_ = mediaplayer.NewMusicPlayerStoreQueueDescriptorWithStoreIDs(
		[]mediaplayer.string{}, // storeIDs []string
	)
	// Output:
}
