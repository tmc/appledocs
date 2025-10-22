// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKVirtualConferenceRoomTypeDescriptor

// ExampleNewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier demonstrates how to create a EKVirtualConferenceRoomTypeDescriptor instance using NewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier.
// Creates an object that describes a location where a virtual conference takes place.
func ExampleNewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier() {
	_ = eventkit.NewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier(
		"title", // title string
		eventkit.EKVirtualConferenceRoomTypeIdentifier{}, // identifier EKVirtualConferenceRoomTypeIdentifier
	)
	// Output:
}
