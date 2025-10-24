// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureScreenInput

// ExampleNewCaptureScreenInput demonstrates how to create a CaptureScreenInput instance.
// Initializes a capture screen input that provides media data from the main screen.
func ExampleNewCaptureScreenInput() {
	_ = avfoundation.NewCaptureScreenInput()
	// Output:
}
// ExampleNewCaptureScreenInputWithDisplayID demonstrates how to create a CaptureScreenInput instance using NewCaptureScreenInputWithDisplayID.
// Initializes a capture screen input that provides media data from the specified display.
func ExampleNewCaptureScreenInputWithDisplayID() {
	_ = avfoundation.NewCaptureScreenInputWithDisplayID(
		avfoundation.DirectDisplayID /* not a class type */{}, // displayID DirectDisplayID /* not a class type */
	)
	// Output:
}
