//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for OverlayAppConfiguration


// Returns the object associated with the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/additionalValue(forKey:)
func (o_ OverlayAppConfiguration) AdditionalValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("additionalValueForKey:"), key)
	return rv
}

// Sets an additional value for a key; for example, a value for measuring the effectiveness of an ad campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdditionalValue(_:forKey:)
func (o_ OverlayAppConfiguration) SetAdditionalValueForKey(value objc.IObject, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdditionalValue:forKey:"), value, key)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdImpression(_:)
func (o_ OverlayAppConfiguration) SetAdImpression(impression ISKAdImpression) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdImpression:"), impression)
}

// iOS-only properties

// The iTunes identifier of the recommended app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/appIdentifier
func (o_ OverlayAppConfiguration) AppIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("appIdentifier"))
	return rv
}
func (o_ OverlayAppConfiguration) SetAppIdentifier(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setAppIdentifier:"), value)
}

// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/campaignToken
func (o_ OverlayAppConfiguration) CampaignToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("campaignToken"))
	return rv
}
func (o_ OverlayAppConfiguration) SetCampaignToken(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setCampaignToken:"), value)
}

// An optional identifier for an app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/customProductPageIdentifier
func (o_ OverlayAppConfiguration) CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}
func (o_ OverlayAppConfiguration) SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setCustomProductPageIdentifier:"), value)
}

// The release ID of the latest version of your app as displayed in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/latestReleaseID
func (o_ OverlayAppConfiguration) LatestReleaseID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}
func (o_ OverlayAppConfiguration) SetLatestReleaseID(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setLatestReleaseID:"), value)
}

// The position of the overlay on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/position
func (o_ OverlayAppConfiguration) Position() OverlayPosition {
	rv := objc.Send[OverlayPosition](o_.ID, objc.Sel("position"))
	return rv
}
func (o_ OverlayAppConfiguration) SetPosition(value OverlayPosition) {
	o_.ID.Send(objc.RegisterName("setPosition:"), value)
}

// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/providerToken
func (o_ OverlayAppConfiguration) ProviderToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("providerToken"))
	return rv
}
func (o_ OverlayAppConfiguration) SetProviderToken(value objc.IObject /* cross-framework: NSString */) {
	o_.ID.Send(objc.RegisterName("setProviderToken:"), value)
}

// A Boolean value that indicates whether the user can dismiss the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/userDismissible
func (o_ OverlayAppConfiguration) UserDismissible() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("userDismissible"))
	return rv
}
func (o_ OverlayAppConfiguration) SetUserDismissible(value bool) {
	o_.ID.Send(objc.RegisterName("setUserDismissible:"), value)
}




