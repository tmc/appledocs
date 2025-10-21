// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEEthernetTunnelNetworkSettings


// ExampleNewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu demonstrates how to create a NEEthernetTunnelNetworkSettings instance using NewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu.
// Creates a settings object with a given tunnel remote address and MAC address.
func ExampleNewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu() {
	_ = networkextension.NewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu(
		"address", // address string
		"ethernetAddress", // ethernetAddress string
		0, // mtu int
	)
	// Output:
}


