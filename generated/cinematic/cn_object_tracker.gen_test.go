// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic_test

import (
	"github.com/tmc/appledocs/generated/cinematic"
)

// Suppress unused import errors
var _ = cinematic.NewCNObjectTracker

// ExampleCNObjectTracker_FinishDetectionTrack demonstrates using FinishDetectionTrack on a CNObjectTracker instance.
// Finish constructing the detection track and return it.
func ExampleCNObjectTracker_FinishDetectionTrack() {
	obj := cinematic.NewCNObjectTracker()
	_ = obj.FinishDetectionTrack()
	// Output:
	}

// ExampleCNObjectTracker_ResetDetectionTrack demonstrates using ResetDetectionTrack on a CNObjectTracker instance.
// Resets the builder to construct a new detection track.
func ExampleCNObjectTracker_ResetDetectionTrack() {
	obj := cinematic.NewCNObjectTracker()
	obj.ResetDetectionTrack()
	// Output:
	}

