// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal_test

import (
	"github.com/tmc/appledocs/generated/metal"
)

// Suppress unused import errors
var _ = metal.NewCaptureManager

// ExampleNewCaptureManager demonstrates how to create a CaptureManager instance.
func ExampleNewCaptureManager() {
	_ = metal.NewCaptureManager()
	// Output:
}
// ExampleCaptureManager_StopCapture demonstrates using StopCapture on a CaptureManager instance.
// Stops capturing Metal commands.
func ExampleCaptureManager_StopCapture() {
	obj := metal.NewCaptureManager()
	obj.StopCapture()
	// Output:
	}

