// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer_test

import (
	"github.com/tmc/appledocs/generated/mediaplayer"
)

// Suppress unused import errors
var _ = mediaplayer.NewMediaPickerController

// ExampleNewMediaPickerControllerWithMediaTypes demonstrates how to create a MediaPickerController instance using NewMediaPickerControllerWithMediaTypes.
// Initializes a media item picker for specified media types.
func ExampleNewMediaPickerControllerWithMediaTypes() {
	_ = mediaplayer.NewMediaPickerControllerWithMediaTypes(
		mediaplayer.MediaType{}, // mediaTypes MediaType
	)
	// Output:
}
