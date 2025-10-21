// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMediaPlaylistCreationMetadata

// ExampleNewMediaPlaylistCreationMetadataWithName demonstrates how to create a MediaPlaylistCreationMetadata instance using NewMediaPlaylistCreationMetadataWithName.
// Creates a new playlist metadata object with the designated name.
func ExampleNewMediaPlaylistCreationMetadataWithName() {
	_ = mediaplayer.NewMediaPlaylistCreationMetadataWithName(
		"name", // name string
	)
	// Output:
}
