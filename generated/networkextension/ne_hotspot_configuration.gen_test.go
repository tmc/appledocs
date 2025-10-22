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
// ExampleNewNEHotspotConfigurationWithSSID demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithSSID.
// Creates a new hotspot configuration, identified by an SSID, for an open Wi-Fi network.
func ExampleNewNEHotspotConfigurationWithSSID() {
	_ = networkextension.NewNEHotspotConfigurationWithSSID(
		"SSID", // SSID string
	)
	// Output:
}
// ExampleNewNEHotspotConfigurationWithSSIDEapSettings demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithSSIDEapSettings.
// Creates a new hotspot configuration, identified by an SSID, for a WPA/WPA2 enterprise Wi-Fi network with EAP settings.
func ExampleNewNEHotspotConfigurationWithSSIDEapSettings() {
	_ = networkextension.NewNEHotspotConfigurationWithSSIDEapSettings(
		"SSID", // SSID string
		networkextension.NEHotspotEAPSettings{}, // eapSettings NEHotspotEAPSettings
	)
	// Output:
}
// ExampleNewNEHotspotConfigurationWithSSIDPassphraseIsWEP demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithSSIDPassphraseIsWEP.
// Creates a new hotspot configuration, identified by an SSID, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
func ExampleNewNEHotspotConfigurationWithSSIDPassphraseIsWEP() {
	_ = networkextension.NewNEHotspotConfigurationWithSSIDPassphraseIsWEP(
		"SSID", // SSID string
		"passphrase", // passphrase string
		false, // isWEP bool
	)
	// Output:
}
// ExampleNewNEHotspotConfigurationWithSSIDPrefix demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithSSIDPrefix.
// Creates a new hotspot configuration, identified by an SSID prefix string, for an open Wi-Fi network.
func ExampleNewNEHotspotConfigurationWithSSIDPrefix() {
	_ = networkextension.NewNEHotspotConfigurationWithSSIDPrefix(
		"SSIDPrefix", // SSIDPrefix string
	)
	// Output:
}
// ExampleNewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP demonstrates how to create a NEHotspotConfiguration instance using NewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP.
// Creates a new hotspot configuration, identified by an SSID prefix string, for a protected WEP or WPA/WPA2 personal Wi-Fi network.
func ExampleNewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP() {
	_ = networkextension.NewNEHotspotConfigurationWithSSIDPrefixPassphraseIsWEP(
		"SSIDPrefix", // SSIDPrefix string
		"passphrase", // passphrase string
		false, // isWEP bool
	)
	// Output:
}
