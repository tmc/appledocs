// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDevice

// ExampleCaptureDevice_UnlockForConfiguration demonstrates using UnlockForConfiguration on a CaptureDevice instance.
// Releases exclusive control over device hardware properties.
func ExampleCaptureDevice_UnlockForConfiguration() {
	obj := avfoundation.NewCaptureDevice()
	obj.UnlockForConfiguration()
	// Output:
	}

