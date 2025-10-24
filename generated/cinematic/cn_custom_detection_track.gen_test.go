// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic_test

import (
	"github.com/tmc/appledocs/generated/cinematic"
)

// Suppress unused import errors
var _ = cinematic.NewCNCustomDetectionTrack

// ExampleNewCNCustomDetectionTrackWithDetectionsSmooth demonstrates how to create a CNCustomDetectionTrack instance using NewCNCustomDetectionTrackWithDetectionsSmooth.
// Initializes a custom detection track with an array of detections, optionally applying smoothing.
func ExampleNewCNCustomDetectionTrackWithDetectionsSmooth() {
	_ = cinematic.NewCNCustomDetectionTrackWithDetectionsSmooth(
		[]cinematic.ICNDetection{}, // detections []ICNDetection
		false, // applySmoothing bool
	)
	// Output:
}
