// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit_test

import (
	"github.com/tmc/appledocs/generated/avkit"
)

// Suppress unused import errors
var _ = avkit.NewPlaybackSpeed

// ExampleNewPlaybackSpeedWithRateLocalizedName demonstrates how to create a PlaybackSpeed instance using NewPlaybackSpeedWithRateLocalizedName.
// Creates a playback speed with a rate and localized name.
func ExampleNewPlaybackSpeedWithRateLocalizedName() {
	_ = avkit.NewPlaybackSpeedWithRateLocalizedName(
		0.0,             // rate float32
		"localizedName", // localizedName string
	)
	// Output:
}
