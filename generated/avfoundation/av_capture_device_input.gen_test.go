// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDeviceInput

// ExampleCaptureDeviceInput_UnfollowExternalSyncDevice demonstrates using UnfollowExternalSyncDevice on a CaptureDeviceInput instance.
// Discontinues external sync.
func ExampleCaptureDeviceInput_UnfollowExternalSyncDevice() {
	obj := avfoundation.NewCaptureDeviceInput()
	obj.UnfollowExternalSyncDevice()
	// Output:
	}

