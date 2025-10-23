// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)

// Suppress unused import errors
var _ = coremidi.NewMIDINetworkHost

// ExampleNewMIDINetworkHostWithNameAddressPort demonstrates how to create a MIDINetworkHost instance using NewMIDINetworkHostWithNameAddressPort.
// Creates a host with the specified name, adress, and port.
func ExampleNewMIDINetworkHostWithNameAddressPort() {
	_ = coremidi.NewMIDINetworkHostWithNameAddressPort(
		"name", // name string
		"address", // address string
		0, // port uint
	)
	// Output:
}
// ExampleNewMIDINetworkHostWithNameNetService demonstrates how to create a MIDINetworkHost instance using NewMIDINetworkHostWithNameNetService.
// Creates a host with the specified name and net service.
func ExampleNewMIDINetworkHostWithNameNetService() {
	_ = coremidi.NewMIDINetworkHostWithNameNetService(
		"name", // name string
		coremidi.NetService{}, // netService NetService
	)
	// Output:
}
