// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewTrackHomographicImageRegistrationRequest

// ExampleNewTrackHomographicImageRegistrationRequest demonstrates how to create a TrackHomographicImageRegistrationRequest instance.
// Creates a new request that tracks the homographic transformation of two images.
func ExampleNewTrackHomographicImageRegistrationRequest() {
	_ = vision.NewTrackHomographicImageRegistrationRequest()
	// Output:
}
// ExampleNewTrackHomographicImageRegistrationRequestWithCompletionHandler demonstrates how to create a TrackHomographicImageRegistrationRequest instance using NewTrackHomographicImageRegistrationRequestWithCompletionHandler.
// Creates a new request that tracks the homographic transformation of two images, with a system callback on completion.
func ExampleNewTrackHomographicImageRegistrationRequestWithCompletionHandler() {
	_ = vision.NewTrackHomographicImageRegistrationRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
