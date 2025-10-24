// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewTrackTranslationalImageRegistrationRequest

// ExampleNewTrackTranslationalImageRegistrationRequest demonstrates how to create a TrackTranslationalImageRegistrationRequest instance.
// Creates a new request that tracks the translational registration of two images.
func ExampleNewTrackTranslationalImageRegistrationRequest() {
	_ = vision.NewTrackTranslationalImageRegistrationRequest()
	// Output:
}
// ExampleNewTrackTranslationalImageRegistrationRequestWithCompletionHandler demonstrates how to create a TrackTranslationalImageRegistrationRequest instance using NewTrackTranslationalImageRegistrationRequestWithCompletionHandler.
// Creates a new request that tracks the translational registration of two images, with a system callback on completion.
func ExampleNewTrackTranslationalImageRegistrationRequestWithCompletionHandler() {
	_ = vision.NewTrackTranslationalImageRegistrationRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
