// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewMutableMovie

// ExampleMutableMovie_UnusedTrackID demonstrates using UnusedTrackID on a MutableMovie instance.
// Returns an identifier that no other tracks in the asset use.
func ExampleMutableMovie_UnusedTrackID() {
	obj := avfoundation.NewMutableMovie()
	_ = obj.UnusedTrackID()
	// Output:
	}

