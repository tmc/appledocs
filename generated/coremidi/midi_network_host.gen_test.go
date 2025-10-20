// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)

// Suppress unused import errors
var _ = coremidi.NewMIDINetworkHost


// ExampleNewMIDINetworkHostWithNameNetServiceNameNetServiceDomain demonstrates how to create a MIDINetworkHost instance using NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain.
// Creates a host with the specified name, net service name, and domain.
func ExampleNewMIDINetworkHostWithNameNetServiceNameNetServiceDomain() {
	_ = coremidi.NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain(
		"name", // name string
		"netServiceName", // netServiceName string
		"netServiceDomain", // netServiceDomain string
	)
	// Output:
}


