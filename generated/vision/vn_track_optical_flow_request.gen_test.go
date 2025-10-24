// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewTrackOpticalFlowRequest

// ExampleNewTrackOpticalFlowRequest demonstrates how to create a TrackOpticalFlowRequest instance.
// Creates a new request that tracks the optical from one image to another.
func ExampleNewTrackOpticalFlowRequest() {
	_ = vision.NewTrackOpticalFlowRequest()
	// Output:
}
// ExampleNewTrackOpticalFlowRequestWithCompletionHandler demonstrates how to create a TrackOpticalFlowRequest instance using NewTrackOpticalFlowRequestWithCompletionHandler.
// Creates a new request that tracks the optical from one image to another, with a system callback on completion.
func ExampleNewTrackOpticalFlowRequestWithCompletionHandler() {
	_ = vision.NewTrackOpticalFlowRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
