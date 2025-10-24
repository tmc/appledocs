// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewVideoProcessorFrameRateCadence

// ExampleNewVideoProcessorFrameRateCadenceWithFrameRate demonstrates how to create a VideoProcessorFrameRateCadence instance using NewVideoProcessorFrameRateCadenceWithFrameRate.
// Creates a new frame-based cadence with a frame rate.
func ExampleNewVideoProcessorFrameRateCadenceWithFrameRate() {
	_ = vision.NewVideoProcessorFrameRateCadenceWithFrameRate(
		0, // frameRate int
	)
	// Output:
}
