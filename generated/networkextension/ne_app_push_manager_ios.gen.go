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

// iOS-only properties

// A delegate that receives incoming call information from the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushManager/delegate
func (n_ NEAppPushManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}
func (n_ NEAppPushManager) SetDelegate(value objc.ID) {
	n_.ID.Send(objc.RegisterName("setDelegate:"), value)
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





