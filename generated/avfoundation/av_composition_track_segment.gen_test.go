// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewCompositionTrackSegment

// ExampleNewCompositionTrackSegmentWithTimeRange demonstrates how to create a CompositionTrackSegment instance using NewCompositionTrackSegmentWithTimeRange.
// Creates an object that presents an empty composition track segment.
func ExampleNewCompositionTrackSegmentWithTimeRange() {
	_ = avfoundation.NewCompositionTrackSegmentWithTimeRange(
		avfoundation.TimeRange /* not a class type */{}, // timeRange TimeRange /* not a class type */
	)
	// Output:
}
