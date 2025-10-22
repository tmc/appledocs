// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit_test

import (
	"github.com/tmc/appledocs/generated/eventkit"
)

// Suppress unused import errors
var _ = eventkit.NewEKVirtualConferenceDescriptor

// ExampleNewEKVirtualConferenceDescriptorWithTitleURLDescriptorsConferenceDetails demonstrates how to create a EKVirtualConferenceDescriptor instance using NewEKVirtualConferenceDescriptorWithTitleURLDescriptorsConferenceDetails.
// Creates an object that describes a virtual conference, including a name and URL to join the conference.
func ExampleNewEKVirtualConferenceDescriptorWithTitleURLDescriptorsConferenceDetails() {
	_ = eventkit.NewEKVirtualConferenceDescriptorWithTitleURLDescriptorsConferenceDetails(
		"title", // title string
		[]eventkit.EKVirtualConferenceURLDescriptor{}, // URLDescriptors []EKVirtualConferenceURLDescriptor
		"conferenceDetails", // conferenceDetails string
	)
	// Output:
}
