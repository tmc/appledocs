// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewVideoCompositionCoreAnimationTool

// ExampleNewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID demonstrates how to create a VideoCompositionCoreAnimationTool instance using NewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID.
// Adds a Core Animation layer to the video composition.
func ExampleNewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID() {
	_ = avfoundation.NewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID(
		avfoundation.Layer{}, // layer Layer
		avfoundation.PersistentTrackID /* not a class type */{}, // trackID PersistentTrackID /* not a class type */
	)
	// Output:
}
// ExampleNewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer demonstrates how to create a VideoCompositionCoreAnimationTool instance using NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer.
// Composes the composited video frame with a Core Animation layer.
func ExampleNewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer() {
	_ = avfoundation.NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayerInLayer(
		avfoundation.Layer{}, // videoLayer Layer
		avfoundation.Layer{}, // animationLayer Layer
	)
	// Output:
}
// ExampleNewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer demonstrates how to create a VideoCompositionCoreAnimationTool instance using NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer.
// Composes the composited video frames with the Core Animation layer.
func ExampleNewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer() {
	_ = avfoundation.NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer(
		[]avfoundation.Layer{}, // videoLayers []Layer
		avfoundation.Layer{}, // animationLayer Layer
	)
	// Output:
}
