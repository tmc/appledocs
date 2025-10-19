// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)


// ExampleNewAVPlayerWithPlayerItem demonstrates how to create a AVPlayer instance using NewAVPlayerWithPlayerItem.
// Creates a new player to play the specified player item.
func ExampleNewAVPlayerWithPlayerItem() {
	_ = avfoundation.NewAVPlayerWithPlayerItem(
		nil, // item unsafe.Pointer
	)
	// Output:
}

// ExampleNewAVPlayerWithURL demonstrates how to create a AVPlayer instance using NewAVPlayerWithURL.
// Creates a new player to play a single audiovisual resource referenced by a given URL.
func ExampleNewAVPlayerWithURL() {
	_ = avfoundation.NewAVPlayerWithURL(
		nil, // URL unsafe.Pointer
	)
	// Output:
}

// ExampleNewAVPlayer demonstrates how to create a AVPlayer instance.
// Creates a player object.
func ExampleNewAVPlayer() {
	_ = avfoundation.NewAVPlayer()
	// Output:
}


