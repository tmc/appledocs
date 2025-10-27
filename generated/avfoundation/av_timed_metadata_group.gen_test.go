// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewTimedMetadataGroup

// ExampleNewTimedMetadataGroupWithSampleBuffer demonstrates how to create a TimedMetadataGroup instance using NewTimedMetadataGroupWithSampleBuffer.
func ExampleNewTimedMetadataGroupWithSampleBuffer() {
	_ = avfoundation.NewTimedMetadataGroupWithSampleBuffer(
		avfoundation.SampleBufferRef /* not a class type */{}, // sampleBuffer SampleBufferRef /* not a class type */
	)
	// Output:
}
// ExampleTimedMetadataGroup_CopyFormatDescription demonstrates using CopyFormatDescription on a TimedMetadataGroup instance.
// Creates a format description based on the receiver’s items.
func ExampleTimedMetadataGroup_CopyFormatDescription() {
	obj := avfoundation.NewTimedMetadataGroup()
	_ = obj.CopyFormatDescription()
	// Output:
	}

