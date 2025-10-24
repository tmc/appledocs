// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCaptionGroup

// ExampleNewCaptionGroupWithCaptionsTimeRange demonstrates how to create a CaptionGroup instance using NewCaptionGroupWithCaptionsTimeRange.
// Creates a caption group with captions and a time range.
func ExampleNewCaptionGroupWithCaptionsTimeRange() {
	_ = avfoundation.NewCaptionGroupWithCaptionsTimeRange(
		[]avfoundation.Caption{}, // captions []Caption
		avfoundation.TimeRange /* not a class type */{}, // timeRange TimeRange /* not a class type */
	)
	// Output:
}
// ExampleNewCaptionGroupWithTimeRange demonstrates how to create a CaptionGroup instance using NewCaptionGroupWithTimeRange.
// Creates a caption group with a time range.
func ExampleNewCaptionGroupWithTimeRange() {
	_ = avfoundation.NewCaptionGroupWithTimeRange(
		avfoundation.TimeRange /* not a class type */{}, // timeRange TimeRange /* not a class type */
	)
	// Output:
}
