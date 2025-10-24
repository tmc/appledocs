// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerInterstitialEventController

// ExamplePlayerInterstitialEventController_SkipCurrentEvent demonstrates using SkipCurrentEvent on a PlayerInterstitialEventController instance.
// Causes the playback of the currently playing interstital event to be abandoned.
func ExamplePlayerInterstitialEventController_SkipCurrentEvent() {
	obj := avfoundation.NewPlayerInterstitialEventController()
	obj.SkipCurrentEvent()
	// Output:
	}

