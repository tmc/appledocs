// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNEHotspotConfiguration

// ExampleNewNEHotspotConfigurationWithHS20SettingsEapSettings demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithHS20SettingsEapSettings.
// Creates a new hotspot configuration, identified by a domain name, for a Hotspot 2.0 Wi-Fi network with HS 2.0 and EAP settings.
func ExampleNewNEHotspotConfigurationWithHS20SettingsEapSettings() {
	_ = networkextension.NewNEHotspotConfigurationWithHS20SettingsEapSettings(
		networkextension.NEHotspotHS20Settings{}, // hs20Settings NEHotspotHS20Settings
		networkextension.NEHotspotEAPSettings{}, // eapSettings NEHotspotEAPSettings
	)
	// Output:
}
