// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptureMetadataInput

// ExampleNewCaptureMetadataInputWithFormatDescriptionClock demonstrates how to create a CaptureMetadataInput instance using NewCaptureMetadataInputWithFormatDescriptionClock.
// Creates capture metadata input to provide timed groups to a capture session.
func ExampleNewCaptureMetadataInputWithFormatDescriptionClock() {
	_ = avfoundation.NewCaptureMetadataInputWithFormatDescriptionClock(
		avfoundation.MetadataFormatDescriptionRef /* not a class type */{}, // desc MetadataFormatDescriptionRef /* not a class type */
		avfoundation.ClockRef /* not a class type */{}, // clock ClockRef /* not a class type */
	)
	// Output:
}
