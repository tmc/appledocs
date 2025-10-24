// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCapturePhotoSettings

// ExampleNewCapturePhotoSettingsWithRawPixelFormatType demonstrates how to create a CapturePhotoSettings instance using NewCapturePhotoSettingsWithRawPixelFormatType.
// Creates a photo settings object for RAW-format-only capture with the specified pixel format.
func ExampleNewCapturePhotoSettingsWithRawPixelFormatType() {
	_ = avfoundation.NewCapturePhotoSettingsWithRawPixelFormatType(
		avfoundation.uint32 /* not a class type */{}, // rawPixelFormatType uint32 /* not a class type */
	)
	// Output:
}
