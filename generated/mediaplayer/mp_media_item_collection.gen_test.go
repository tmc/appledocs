// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMediaItemCollection

// ExampleNewMediaItemCollectionWithItems demonstrates how to create a MediaItemCollection instance using NewMediaItemCollectionWithItems.
// Initializes a media item collection with an array of media items.
func ExampleNewMediaItemCollectionWithItems() {
	_ = mediaplayer.NewMediaItemCollectionWithItems(
		[]mediaplayer.IMediaItem{}, // items []IMediaItem
	)
	// Output:
}
