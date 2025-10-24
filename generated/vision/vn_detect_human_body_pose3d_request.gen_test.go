// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewDetectHumanBodyPose3DRequest

// ExampleNewDetectHumanBodyPose3DRequest demonstrates how to create a DetectHumanBodyPose3DRequest instance.
// Creates a new request with no completion handler.
func ExampleNewDetectHumanBodyPose3DRequest() {
	_ = vision.NewDetectHumanBodyPose3DRequest()
	// Output:
}
// ExampleNewDetectHumanBodyPose3DRequestWithCompletionHandler demonstrates how to create a DetectHumanBodyPose3DRequest instance using NewDetectHumanBodyPose3DRequestWithCompletionHandler.
// Creates a new 3D body pose request with a completion handler.
func ExampleNewDetectHumanBodyPose3DRequestWithCompletionHandler() {
	_ = vision.NewDetectHumanBodyPose3DRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
