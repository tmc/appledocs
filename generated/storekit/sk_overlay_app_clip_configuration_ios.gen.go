//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for OverlayAppClipConfiguration


// Returns the object associated with the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/additionalValue(forKey:)
func (o_ OverlayAppClipConfiguration) AdditionalValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("additionalValueForKey:"), key)
	return rv
}

// Sets an additional value for a key, such as a value for measuring the effectiveness of an ad campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/setAdditionalValue(_:forKey:)
func (o_ OverlayAppClipConfiguration) SetAdditionalValueForKey(value objc.IObject, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdditionalValue:forKey:"), value, key)
}

// iOS-only properties

// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/campaignToken
func (o_ OverlayAppClipConfiguration) CampaignToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("campaignToken"))
	return rv
}
func (o_ OverlayAppClipConfiguration) SetCampaignToken(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setCampaignToken:"), value)
}

// An identifier for a parent app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/customProductPageIdentifier
func (o_ OverlayAppClipConfiguration) CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}
func (o_ OverlayAppClipConfiguration) SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setCustomProductPageIdentifier:"), value)
}

// The release ID of the latest version of your parent app as displayed in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/latestReleaseID
func (o_ OverlayAppClipConfiguration) LatestReleaseID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}
func (o_ OverlayAppClipConfiguration) SetLatestReleaseID(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setLatestReleaseID:"), value)
}

// The position of the overlay on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/position
func (o_ OverlayAppClipConfiguration) Position() OverlayPosition {
	rv := objc.Send[OverlayPosition](o_.ID, objc.Sel("position"))
	return rv
}
func (o_ OverlayAppClipConfiguration) SetPosition(value OverlayPosition) {
	o_.ID.Send(objc.RegisterName("setPosition:"), value)
}

// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/providerToken
func (o_ OverlayAppClipConfiguration) ProviderToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("providerToken"))
	return rv
}
func (o_ OverlayAppClipConfiguration) SetProviderToken(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setProviderToken:"), value)
}




