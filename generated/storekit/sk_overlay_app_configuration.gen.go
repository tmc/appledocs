// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AdditionalValueForKey(key appkit.string) objc.ID
	SetAdImpression(impression ISKAdImpression)
	SetAdditionalValueForKey(value objectivec.IObject, key appkit.string)
}

// An object that represents the attributes of an overlay you use to recommend another app on the App Store.
//
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




// Creates an object that represents the attributes of an overlay you use to recommend another app on the App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/init(appIdentifier:position:)
func NewOverlayAppConfigurationWithAppIdentifierPosition(appIdentifier appkit.string, position OverlayPosition) OverlayAppConfiguration {
	instance := getOverlayAppConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppConfiguration](instance.ID, objc.Sel("initWithAppIdentifier:position:"), appIdentifier, position)
	rv.Autorelease()
	return rv
}


// Returns the object associated with the key.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/additionalValue(forKey:)
func (o_ OverlayAppConfiguration) AdditionalValueForKey(key appkit.string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("additionalValueForKey:"), key)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdImpression(_:)
func (o_ OverlayAppConfiguration) SetAdImpression(impression ISKAdImpression) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdImpression:"), impression)
}

// Sets an additional value for a key; for example, a value for measuring the effectiveness of an ad campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdditionalValue(_:forKey:)
func (o_ OverlayAppConfiguration) SetAdditionalValueForKey(value objectivec.IObject, key appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdditionalValue:forKey:"), value, key)
}

// The iTunes identifier of the recommended app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/appIdentifier
func (o_ OverlayAppConfiguration) AppIdentifier() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("appIdentifier"))
	return rv
}


// SetAppIdentifier sets the value of the appIdentifier property.
// The iTunes identifier of the recommended app.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/appIdentifier
func (o_ OverlayAppConfiguration) SetAppIdentifier(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppIdentifier:"), value)
}

// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/campaignToken
func (o_ OverlayAppConfiguration) CampaignToken() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("campaignToken"))
	return rv
}


// SetCampaignToken sets the value of the campaignToken property.
// A token you use to represent an ad campaign and measure its effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/campaignToken
func (o_ OverlayAppConfiguration) SetCampaignToken(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCampaignToken:"), value)
}

// An optional identifier for an app’s custom product page.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/customProductPageIdentifier
func (o_ OverlayAppConfiguration) CustomProductPageIdentifier() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}


// SetCustomProductPageIdentifier sets the value of the customProductPageIdentifier property.
// An optional identifier for an app’s custom product page.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/customProductPageIdentifier
func (o_ OverlayAppConfiguration) SetCustomProductPageIdentifier(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomProductPageIdentifier:"), value)
}

// The release ID of the latest version of your app as displayed in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/latestReleaseID
func (o_ OverlayAppConfiguration) LatestReleaseID() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}


// SetLatestReleaseID sets the value of the latestReleaseID property.
// The release ID of the latest version of your app as displayed in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/latestReleaseID
func (o_ OverlayAppConfiguration) SetLatestReleaseID(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLatestReleaseID:"), value)
}

// The position of the overlay on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/position
func (o_ OverlayAppConfiguration) Position() OverlayPosition {
	rv := objc.Send[OverlayPosition](o_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the overlay on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/position
func (o_ OverlayAppConfiguration) SetPosition(value OverlayPosition) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPosition:"), value)
}

// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/providerToken
func (o_ OverlayAppConfiguration) ProviderToken() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("providerToken"))
	return rv
}


// SetProviderToken sets the value of the providerToken property.
// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/providerToken
func (o_ OverlayAppConfiguration) SetProviderToken(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setProviderToken:"), value)
}

// A Boolean value that indicates whether the user can dismiss the overlay.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/userDismissible
func (o_ OverlayAppConfiguration) UserDismissible() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("userDismissible"))
	return rv
}


// SetUserDismissible sets the value of the userDismissible property.
// A Boolean value that indicates whether the user can dismiss the overlay.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/userDismissible
func (o_ OverlayAppConfiguration) SetUserDismissible(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUserDismissible:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) AdAttributionReengagementURL() foundation.URL {
	rv := objc.Send[foundation.URL](o_.ID, objc.Sel("adAttributionReengagementURL"))
	return rv
}


// SetAdAttributionReengagementURL sets the value of the adAttributionReengagementURL property.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) SetAdAttributionReengagementURL(value foundation.IURL) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdAttributionReengagementURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) AppImpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("appImpression"))
	return rv
}


// SetAppImpression sets the value of the appImpression property.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) SetAppImpression(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppImpression:"), value)
}

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) Configuration() SKOverlayConfiguration {
	rv := objc.Send[SKOverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// An overlay’s attributes; for example, its position on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) SetConfiguration(value ISKOverlayConfiguration) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
}


