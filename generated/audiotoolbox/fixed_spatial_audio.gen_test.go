// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox_test

import (
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// Suppress unused import errors
var _ = audiotoolbox.NewFixedSpatialAudio

// ExampleNewFixedSpatialAudioWithSoundStageSize demonstrates how to create a FixedSpatialAudio instance using NewFixedSpatialAudioWithSoundStageSize.
func ExampleNewFixedSpatialAudioWithSoundStageSize() {
	_ = audiotoolbox.NewFixedSpatialAudioWithSoundStageSize(
		audiotoolbox.CASoundStageSize{}, // soundStageSize CASoundStageSize
	)
	// Output:
}
