// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic_test

import (
	"github.com/tmc/appledocs/generated/cinematic"
)

// Suppress unused import errors
var _ = cinematic.NewCNRenderingSessionFrameAttributes

// ExampleNewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes demonstrates how to create a CNRenderingSessionFrameAttributes instance using NewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes.
// Initializes the rendering frame attributes from a sample buffer read from a Cinematic metadata track.
func ExampleNewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes() {
	_ = cinematic.NewCNRenderingSessionFrameAttributesWithSampleBufferSessionAttributes(
		cinematic.SampleBufferRef{}, // sampleBuffer SampleBufferRef
		cinematic.CNRenderingSessionAttributes{}, // sessionAttributes CNRenderingSessionAttributes
	)
	// Output:
}
// ExampleNewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes demonstrates how to create a CNRenderingSessionFrameAttributes instance using NewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes.
// Initializes the rendering frame attributes from a timed metadata group read from a Cinematic metadata track.
func ExampleNewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes() {
	_ = cinematic.NewCNRenderingSessionFrameAttributesWithTimedMetadataGroupSessionAttributes(
		cinematic.TimedMetadataGroup{}, // metadataGroup TimedMetadataGroup
		cinematic.CNRenderingSessionAttributes{}, // sessionAttributes CNRenderingSessionAttributes
	)
	// Output:
}
