// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic_test

import (
	"github.com/tmc/appledocs/generated/cinematic"
)

// Suppress unused import errors
var _ = cinematic.NewCNFixedDetectionTrack

// ExampleNewCNFixedDetectionTrackWithOriginalDetection demonstrates how to create a CNFixedDetectionTrack instance using NewCNFixedDetectionTrackWithOriginalDetection.
// Creates a detection track with fixed focus at the disparity of an existing detection.
func ExampleNewCNFixedDetectionTrackWithOriginalDetection() {
	_ = cinematic.NewCNFixedDetectionTrackWithOriginalDetection(
		cinematic.CNDetection{}, // originalDetection CNDetection
	)
	// Output:
}
