// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEDNSSettings

// ExampleNewNEDNSSettingsWithServers demonstrates how to create a NEDNSSettings instance using NewNEDNSSettingsWithServers.
// Initialize the   object.
func ExampleNewNEDNSSettingsWithServers() {
	_ = networkextension.NewNEDNSSettingsWithServers(
		[]networkextension.string{}, // servers []string
	)
	// Output:
}
