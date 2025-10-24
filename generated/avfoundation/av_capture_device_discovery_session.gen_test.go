// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDeviceDiscoverySession

// ExampleNewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition demonstrates how to create a CaptureDeviceDiscoverySession instance using NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition.
// Creates a discovery session that finds devices that match the specified criteria.
func ExampleNewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition() {
	_ = avfoundation.NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition(
		[]avfoundation.string{}, // deviceTypes []string
		avfoundation.MediaType /* typedef */{}, // mediaType MediaType /* typedef */
		avfoundation.CaptureDevicePosition{}, // position CaptureDevicePosition
	)
	// Output:
}
