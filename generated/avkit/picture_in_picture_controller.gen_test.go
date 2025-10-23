// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit_test

import (
	"github.com/tmc/appledocs/generated/avkit"
)

// Suppress unused import errors
var _ = avkit.NewPictureInPictureController

// ExampleNewPictureInPictureControllerWithContentSource demonstrates how to create a PictureInPictureController instance using NewPictureInPictureControllerWithContentSource.
// Creates a Picture in Picture controller with a content source.
func ExampleNewPictureInPictureControllerWithContentSource() {
	_ = avkit.NewPictureInPictureControllerWithContentSource(
		avkit.AVPictureInPictureControllerContentSource{}, // contentSource AVPictureInPictureControllerContentSource
	)
	// Output:
}
// ExampleNewPictureInPictureControllerWithPlayerLayer demonstrates how to create a PictureInPictureController instance using NewPictureInPictureControllerWithPlayerLayer.
// Creates a Picture in Picture controller with a player layer.
func ExampleNewPictureInPictureControllerWithPlayerLayer() {
	_ = avkit.NewPictureInPictureControllerWithPlayerLayer(
		avkit.PlayerLayer{}, // playerLayer PlayerLayer
	)
	// Output:
}
