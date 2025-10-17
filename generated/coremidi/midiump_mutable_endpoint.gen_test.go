// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)


// ExampleNewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback demonstrates how to create a MIDIUMPMutableEndpoint instance using NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback.
func ExampleNewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback() {
	_ = coremidi.NewMIDIUMPMutableEndpointWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback(
		"name", // name string
		nil, // deviceInfo unsafe.Pointer
		"productInstanceID", // productInstanceID string
		nil, // MIDIProtocol unsafe.Pointer
		nil, // destinationCallback unsafe.Pointer
	)
	// Output:
}


