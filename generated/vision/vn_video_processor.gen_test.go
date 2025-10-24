// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewVideoProcessor

// ExampleVideoProcessor_Cancel demonstrates using Cancel on a VideoProcessor instance.
// Cancels the video processing.
func ExampleVideoProcessor_Cancel() {
	obj := vision.NewVideoProcessor()
	obj.Cancel()
	// Output:
	}

