// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureDeviceRotationCoordinator

// ExampleNewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer demonstrates how to create a CaptureDeviceRotationCoordinator instance using NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer.
// Creates a coordinator that provides separate compensation angles for content your app takes with a capture device, and for your app’s camera preview.
func ExampleNewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer() {
	_ = avfoundation.NewCaptureDeviceRotationCoordinatorWithDevicePreviewLayer(
		avfoundation.AVCaptureDevice{}, // device AVCaptureDevice
		avfoundation.CaptureVideoPreviewLayer{}, // previewLayer CaptureVideoPreviewLayer
	)
	// Output:
}
