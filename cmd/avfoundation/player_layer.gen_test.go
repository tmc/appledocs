// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerLayer

// ExampleNewPlayerLayerWithPlayer demonstrates how to create a PlayerLayer instance using NewPlayerLayerWithPlayer.
// Creates a layer object to present the visual contents of a player’s current item.
func ExampleNewPlayerLayerWithPlayer() {
	_ = avfoundation.NewPlayerLayerWithPlayer(
		avfoundation.AVPlayer{}, // player AVPlayer
	)
	// Output:
}
