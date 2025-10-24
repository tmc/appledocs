// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewGeneratePersonSegmentationRequest

// ExampleNewGeneratePersonSegmentationRequest demonstrates how to create a GeneratePersonSegmentationRequest instance.
// Creates a generate person segmentation request.
func ExampleNewGeneratePersonSegmentationRequest() {
	_ = vision.NewGeneratePersonSegmentationRequest()
	// Output:
}
// ExampleNewGeneratePersonSegmentationRequestWithCompletionHandler demonstrates how to create a GeneratePersonSegmentationRequest instance using NewGeneratePersonSegmentationRequestWithCompletionHandler.
// Creates a generate person segmentation request with a completion handler.
func ExampleNewGeneratePersonSegmentationRequestWithCompletionHandler() {
	_ = vision.NewGeneratePersonSegmentationRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
