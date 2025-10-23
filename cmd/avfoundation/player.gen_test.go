// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayer

// ExampleNewPlayer demonstrates how to create a Player instance.
// Creates a player object.
func ExampleNewPlayer() {
	_ = avfoundation.NewPlayer()
	// Output:
}
// ExampleNewPlayerWithPlayerItem demonstrates how to create a Player instance using NewPlayerWithPlayerItem.
// Creates a new player to play the specified player item.
func ExampleNewPlayerWithPlayerItem() {
	_ = avfoundation.NewPlayerWithPlayerItem(
		avfoundation.AVPlayerItem{}, // item AVPlayerItem
	)
	// Output:
}
