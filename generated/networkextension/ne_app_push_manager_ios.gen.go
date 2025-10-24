//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEAppPushManager


// Loads the manager’s saved configuration from the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/loadFromPreferences(completionHandler:)
func (n_ NEAppPushManager) LoadFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("loadFromPreferencesWithCompletionHandler:"), completionHandler)
}

// Removes the manager’s configuration from the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/removeFromPreferences(completionHandler:)
func (n_ NEAppPushManager) RemoveFromPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("removeFromPreferencesWithCompletionHandler:"), completionHandler)
}

// Saves the manager’s configuration in the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/saveToPreferences(completionHandler:)
func (n_ NEAppPushManager) SaveToPreferencesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("saveToPreferencesWithCompletionHandler:"), completionHandler)
}

// iOS-only properties

// A delegate that receives incoming call information from the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/delegate
func (n_ NEAppPushManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}
func (n_ NEAppPushManager) SetDelegate(value unsafe.Pointer) {
	n_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// A Boolean value that indicates whether a configuration is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/isActive
func (n_ NEAppPushManager) Active() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("active"))
	return rv
}

// A property you use to toggle enabling the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/isEnabled
func (n_ NEAppPushManager) Enabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("enabled"))
	return rv
}
func (n_ NEAppPushManager) SetEnabled(value bool) {
	n_.ID.Send(objc.RegisterName("setEnabled:"), value)
}

// A string that contains the localized description of the app push manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/localizedDescription
func (n_ NEAppPushManager) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("localizedDescription"))
	return rv
}
func (n_ NEAppPushManager) SetLocalizedDescription(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setLocalizedDescription:"), value)
}

// A property that indicates Ethernet support for Local Push Connectivity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/matchEthernet
func (n_ NEAppPushManager) MatchEthernet() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("matchEthernet"))
	return rv
}
func (n_ NEAppPushManager) SetMatchEthernet(value bool) {
	n_.ID.Send(objc.RegisterName("setMatchEthernet:"), value)
}

// An array of private LTE networks that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/matchPrivateLTENetworks
func (n_ NEAppPushManager) MatchPrivateLTENetworks() []NEPrivateLTENetwork {
	rv := objc.Send[[]NEPrivateLTENetwork](n_.ID, objc.Sel("matchPrivateLTENetworks"))
	return rv
}
func (n_ NEAppPushManager) SetMatchPrivateLTENetworks(value []NEPrivateLTENetwork) {
	n_.ID.Send(objc.RegisterName("setMatchPrivateLTENetworks:"), value)
}

// An array of Wi-Fi SSID strings that the system matches for local push activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/matchSSIDs
func (n_ NEAppPushManager) MatchSSIDs() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchSSIDs"))
	return rv
}
func (n_ NEAppPushManager) SetMatchSSIDs(value []string) {
	n_.ID.Send(objc.RegisterName("setMatchSSIDs:"), value)
}

// A string that contains the bundle identifier of the push provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/providerBundleIdentifier
func (n_ NEAppPushManager) ProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("providerBundleIdentifier"))
	return rv
}
func (n_ NEAppPushManager) SetProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	n_.ID.Send(objc.RegisterName("setProviderBundleIdentifier:"), value)
}

// A dictionary that contains vendor-specific key-value pairs, that you use to configure a provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/providerConfiguration
func (n_ NEAppPushManager) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}
func (n_ NEAppPushManager) SetProviderConfiguration(value foundation.IDictionary) {
	n_.ID.Send(objc.RegisterName("setProviderConfiguration:"), value)
}





