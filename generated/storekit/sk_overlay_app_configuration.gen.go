// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [OverlayAppConfiguration] class.
var (
	OverlayAppConfigurationClass     _OverlayAppConfigurationClass
	OverlayAppConfigurationClassOnce sync.Once
)

func getOverlayAppConfigurationClass() _OverlayAppConfigurationClass {
	OverlayAppConfigurationClassOnce.Do(func() {
		OverlayAppConfigurationClass = _OverlayAppConfigurationClass{objc.GetClass("SKOverlayAppConfiguration")}
	})
	return OverlayAppConfigurationClass
}

type _OverlayAppConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [OverlayAppConfiguration] class.
type IOverlayAppConfiguration interface {
	IOverlayConfiguration
	// properties:
	AdAttributionReengagementURL() objc.IObject /* cross-framework: URL */
	SetAdAttributionReengagementURL(value objc.IObject /* cross-framework: URL */)
	AppIdentifier() objc.IObject /* cross-framework: NSString */
	SetAppIdentifier(value objc.IObject /* cross-framework: NSString */)
	AppImpression() unsafe.Pointer
	SetAppImpression(value unsafe.Pointer)
	CampaignToken() objc.IObject /* cross-framework: NSString */
	SetCampaignToken(value objc.IObject /* cross-framework: NSString */)
	CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */
	SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */)
	LatestReleaseID() objc.IObject /* cross-framework: NSString */
	SetLatestReleaseID(value objc.IObject /* cross-framework: NSString */)
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	ProviderToken() objc.IObject /* cross-framework: NSString */
	SetProviderToken(value objc.IObject /* cross-framework: NSString */)
	UserDismissible() bool
	SetUserDismissible(value bool)
	Configuration() objc.IObject /* cross-framework: OverlayConfiguration */
	SetConfiguration(value objc.IObject /* cross-framework: OverlayConfiguration */)
	// methods:
}

// An object that represents the attributes of an overlay you use to recommend another app on the App Store.


// An object that represents the attributes of an overlay you use to recommend another app on the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration
type OverlayAppConfiguration struct {
	OverlayConfiguration
}

// OverlayAppConfigurationFrom constructs a [OverlayAppConfiguration] from an unsafe.Pointer.
//
// An object that represents the attributes of an overlay you use to recommend another app on the App Store.
func OverlayAppConfigurationFrom(ptr unsafe.Pointer) OverlayAppConfiguration {
	return OverlayAppConfiguration{
		OverlayConfiguration: OverlayConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OverlayAppConfigurationClass) Alloc() OverlayAppConfiguration {
	rv := objc.Send[OverlayAppConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OverlayAppConfigurationClass) New() OverlayAppConfiguration {
	rv := objc.Send[OverlayAppConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayAppConfiguration) Init() OverlayAppConfiguration {
	rv := objc.Send[OverlayAppConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayAppConfiguration) Autorelease() OverlayAppConfiguration {
	rv := objc.Send[OverlayAppConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayAppConfiguration creates a new OverlayAppConfiguration instance.
func NewOverlayAppConfiguration() OverlayAppConfiguration {
	return getOverlayAppConfigurationClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) AdAttributionReengagementURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](o_.ID, objc.Sel("adAttributionReengagementURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) SetAdAttributionReengagementURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdAttributionReengagementURL:"), value)
}


// The iTunes identifier of the recommended app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appidentifier
func (o_ OverlayAppConfiguration) AppIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("appIdentifier"))
	return rv
}


// The iTunes identifier of the recommended app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appidentifier
func (o_ OverlayAppConfiguration) SetAppIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) AppImpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("appImpression"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) SetAppImpression(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppImpression:"), value)
}


// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/campaigntoken
func (o_ OverlayAppConfiguration) CampaignToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("campaignToken"))
	return rv
}


// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/campaigntoken
func (o_ OverlayAppConfiguration) SetCampaignToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCampaignToken:"), value)
}


// An optional identifier for an app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/customproductpageidentifier
func (o_ OverlayAppConfiguration) CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}


// An optional identifier for an app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/customproductpageidentifier
func (o_ OverlayAppConfiguration) SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomProductPageIdentifier:"), value)
}


// The release ID of the latest version of your app as displayed in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/latestreleaseid
func (o_ OverlayAppConfiguration) LatestReleaseID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}


// The release ID of the latest version of your app as displayed in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/latestreleaseid
func (o_ OverlayAppConfiguration) SetLatestReleaseID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLatestReleaseID:"), value)
}


// The position of the overlay on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/position
func (o_ OverlayAppConfiguration) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("position"))
	return rv
}


// The position of the overlay on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/position
func (o_ OverlayAppConfiguration) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPosition:"), value)
}


// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/providertoken
func (o_ OverlayAppConfiguration) ProviderToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("providerToken"))
	return rv
}


// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/providertoken
func (o_ OverlayAppConfiguration) SetProviderToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setProviderToken:"), value)
}


// A Boolean value that indicates whether the user can dismiss the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/userdismissible
func (o_ OverlayAppConfiguration) UserDismissible() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("userDismissible"))
	return rv
}


// A Boolean value that indicates whether the user can dismiss the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/userdismissible
func (o_ OverlayAppConfiguration) SetUserDismissible(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUserDismissible:"), value)
}


// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) Configuration() objc.IObject /* cross-framework: OverlayConfiguration */ {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}


// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) SetConfiguration(value objc.IObject /* cross-framework: OverlayConfiguration */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
}



