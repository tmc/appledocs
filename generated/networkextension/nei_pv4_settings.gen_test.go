// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEIPv4Settings

// ExampleNewNEIPv4SettingsWithAddressesSubnetMasks demonstrates how to create a NEIPv4Settings instance using NewNEIPv4SettingsWithAddressesSubnetMasks.
// Initializes an IPv4 settings object.
func ExampleNewNEIPv4SettingsWithAddressesSubnetMasks() {
	_ = networkextension.NewNEIPv4SettingsWithAddressesSubnetMasks(
		[]networkextension.string{}, // addresses []string
		[]networkextension.string{}, // subnetMasks []string
	)
	// Output:
}
