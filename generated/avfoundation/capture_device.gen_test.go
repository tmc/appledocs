// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDevice

// ExampleNewCaptureDeviceWithUniqueID demonstrates how to create a CaptureDevice instance using NewCaptureDeviceWithUniqueID.
// Creates an object that represents a device with the specified identifier.
func ExampleNewCaptureDeviceWithUniqueID() {
	_ = avfoundation.NewCaptureDeviceWithUniqueID(
		"deviceUniqueID", // deviceUniqueID string
	)
	// Output:
}
