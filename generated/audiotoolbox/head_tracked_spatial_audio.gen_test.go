// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox_test

import (
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// Suppress unused import errors
var _ = audiotoolbox.NewHeadTrackedSpatialAudio

// ExampleNewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy demonstrates how to create a HeadTrackedSpatialAudio instance using NewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy.
func ExampleNewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy() {
	_ = audiotoolbox.NewHeadTrackedSpatialAudioWithSoundStageSizeAnchoringStrategy(
		audiotoolbox.CASoundStageSize{}, // soundStageSize CASoundStageSize
		audiotoolbox.CAAnchoringStrategy{}, // anchoringStrategy CAAnchoringStrategy
	)
	// Output:
}
