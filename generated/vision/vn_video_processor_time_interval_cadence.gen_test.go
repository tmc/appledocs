// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewVideoProcessorTimeIntervalCadence

// ExampleNewVideoProcessorTimeIntervalCadenceWithTimeInterval demonstrates how to create a VideoProcessorTimeIntervalCadence instance using NewVideoProcessorTimeIntervalCadenceWithTimeInterval.
// Creates a new time-based cadence with a time interval.
func ExampleNewVideoProcessorTimeIntervalCadenceWithTimeInterval() {
	_ = vision.NewVideoProcessorTimeIntervalCadenceWithTimeInterval(
		0.0, // timeInterval float64
	)
	// Output:
}


