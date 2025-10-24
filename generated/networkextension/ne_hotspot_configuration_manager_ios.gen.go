//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotConfigurationManager


// Adds or updates a Wi-Fi network configuration after prompting the user for permission, and then attempts to join the network under certain conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/apply(_:completionHandler:)
func (n_ NEHotspotConfigurationManager) ApplyConfigurationCompletionHandler(configuration INEHotspotConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("applyConfiguration:completionHandler:"), configuration, completionHandler)
}

// Submits a completion handler the configuration manager calls to send your app the names of the SSIDs or Wi-Fi hotspot domains in the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/getConfiguredSSIDs(completionHandler:)
func (n_ NEHotspotConfigurationManager) GetConfiguredSSIDsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("getConfiguredSSIDsWithCompletionHandler:"), completionHandler)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/joinAccessoryHotspot(_:passphrase:completionHandler:)
func (n_ NEHotspotConfigurationManager) JoinAccessoryHotspotPassphraseCompletionHandler(accessory Accessory /* not a class type */, passphrase objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("joinAccessoryHotspot:passphrase:completionHandler:"), accessory, passphrase, completionHandler)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/joinAccessoryHotspotWithoutSecurity(_:completionHandler:)
func (n_ NEHotspotConfigurationManager) JoinAccessoryHotspotWithoutSecurityCompletionHandler(accessory Accessory /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("joinAccessoryHotspotWithoutSecurity:completionHandler:"), accessory, completionHandler)
}

// Removes a Wi-Fi hotspot configuration, identified by a Hotspot 2.0 domain name, that your app previously added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/removeConfiguration(forHS20DomainName:)
func (n_ NEHotspotConfigurationManager) RemoveConfigurationForHS20DomainName(domainName objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeConfigurationForHS20DomainName:"), domainName)
}

// Removes a Wi-Fi configuration, identified by an SSID, that your app previously added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfigurationManager/removeConfiguration(forSSID:)
func (n_ NEHotspotConfigurationManager) RemoveConfigurationForSSID(SSID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeConfigurationForSSID:"), SSID)
}

// iOS-only properties





