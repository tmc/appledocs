// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic_test

import (
	"github.com/tmc/appledocs/generated/cinematic"
)

// Suppress unused import errors
var _ = cinematic.NewCNFixedDetectionTrack

// ExampleNewCNFixedDetectionTrackWithFocusDisparity demonstrates how to create a CNFixedDetectionTrack instance using NewCNFixedDetectionTrackWithFocusDisparity.
// Creates a detection track with fixed focus at the given disparity.
func ExampleNewCNFixedDetectionTrackWithFocusDisparity() {
	_ = cinematic.NewCNFixedDetectionTrackWithFocusDisparity(
		0.0, // focusDisparity float32
	)
	// Output:
}
// ExampleNewCNFixedDetectionTrackWithOriginalDetection demonstrates how to create a CNFixedDetectionTrack instance using NewCNFixedDetectionTrackWithOriginalDetection.
// Creates a detection track with fixed focus at the disparity of an existing detection.
func ExampleNewCNFixedDetectionTrackWithOriginalDetection() {
	_ = cinematic.NewCNFixedDetectionTrackWithOriginalDetection(
		cinematic.CNDetection{}, // originalDetection CNDetection
	)
	// Output:
}
