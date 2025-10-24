// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit_test

import (
	"github.com/tmc/appledocs/generated/avkit"
)

// Suppress unused import errors
var _ = avkit.NewPictureInPictureController

// ExamplePictureInPictureController_InvalidatePlaybackState demonstrates using InvalidatePlaybackState on a PictureInPictureController instance.
// Invalidates the controller’s current playback state and fetches the updated state from the sample buffer playback delegate object.
func ExamplePictureInPictureController_InvalidatePlaybackState() {
	obj := avkit.NewPictureInPictureController()
	obj.InvalidatePlaybackState()
	// Output:
	}

// ExamplePictureInPictureController_StartPictureInPicture demonstrates using StartPictureInPicture on a PictureInPictureController instance.
// Starts Picture in Picture, if possible.
func ExamplePictureInPictureController_StartPictureInPicture() {
	obj := avkit.NewPictureInPictureController()
	obj.StartPictureInPicture()
	// Output:
	}

// ExamplePictureInPictureController_StopPictureInPicture demonstrates using StopPictureInPicture on a PictureInPictureController instance.
// Stops Picture in Picture, if active.
func ExamplePictureInPictureController_StopPictureInPicture() {
	obj := avkit.NewPictureInPictureController()
	obj.StopPictureInPicture()
	// Output:
	}

