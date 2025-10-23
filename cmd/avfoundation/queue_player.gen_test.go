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
