// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewQueuePlayer

// ExampleNewQueuePlayerWithItems demonstrates how to create a QueuePlayer instance using NewQueuePlayerWithItems.
// Creates an object that plays a queue of items.
func ExampleNewQueuePlayerWithItems() {
	_ = avfoundation.NewQueuePlayerWithItems(
		[]avfoundation.PlayerItem{}, // items []PlayerItem
	)
	// Output:
}
// ExampleQueuePlayer_AdvanceToNextItem demonstrates using AdvanceToNextItem on a QueuePlayer instance.
// Ends playback of the current item and starts playback of the next item in the player’s queue.
func ExampleQueuePlayer_AdvanceToNextItem() {
	obj := avfoundation.NewQueuePlayer()
	obj.AdvanceToNextItem()
	// Output:
	}

// ExampleQueuePlayer_Items demonstrates using Items on a QueuePlayer instance.
// Returns an array of the currently enqueued items.
func ExampleQueuePlayer_Items() {
	obj := avfoundation.NewQueuePlayer()
	_ = obj.Items()
	// Output:
	}

// ExampleQueuePlayer_RemoveAllItems demonstrates using RemoveAllItems on a QueuePlayer instance.
// Removes all player items from the queue.
func ExampleQueuePlayer_RemoveAllItems() {
	obj := avfoundation.NewQueuePlayer()
	obj.RemoveAllItems()
	// Output:
	}

