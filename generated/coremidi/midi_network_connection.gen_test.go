// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)

// Suppress unused import errors
var _ = coremidi.NewMIDINetworkConnection

// ExampleNewMIDINetworkConnectionWithHost demonstrates how to create a MIDINetworkConnection instance using NewMIDINetworkConnectionWithHost.
// Creates a connection to the specified host.
func ExampleNewMIDINetworkConnectionWithHost() {
	_ = coremidi.NewMIDINetworkConnectionWithHost(
		coremidi.MIDINetworkHost{}, // host MIDINetworkHost
	)
	// Output:
}
