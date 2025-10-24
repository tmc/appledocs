// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewRequest

// ExampleNewRequest demonstrates how to create a Request instance.
// Creates a new Vision request with no completion handler.
func ExampleNewRequest() {
	_ = vision.NewRequest()
	// Output:
}
// ExampleNewRequestWithCompletionHandler demonstrates how to create a Request instance using NewRequestWithCompletionHandler.
// Creates a new Vision request with an optional completion handler.
func ExampleNewRequestWithCompletionHandler() {
	_ = vision.NewRequestWithCompletionHandler(
		vision.RequestCompletionHandler /* not a class type */{}, // completionHandler RequestCompletionHandler /* not a class type */
	)
	// Output:
}
// ExampleRequest_Cancel demonstrates using Cancel on a Request instance.
// Cancels the request before it can finish executing.
func ExampleRequest_Cancel() {
	obj := vision.NewRequest()
	obj.Cancel()
	// Output:
	}

