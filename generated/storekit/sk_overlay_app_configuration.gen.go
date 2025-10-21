// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	AdditionalValueForKey(key string) objc.ID
	SetAdImpression(impression unsafe.Pointer)
	SetAdditionalValueForKey(value objc.ID, key string)
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
func NewOverlayAppConfigurationWithAppIdentifierPosition(appIdentifier string, position unsafe.Pointer) OverlayAppConfiguration {
	instance := getOverlayAppConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppConfiguration](instance.ID, objc.Sel("initWithAppIdentifier:position:"), objc.String(appIdentifier), position)
	rv.Autorelease()
	return rv
}


// Returns the object associated with the key.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/additionalValue(forKey:)
func (o_ OverlayAppConfiguration) AdditionalValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("additionalValueForKey:"), objc.String(key))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdImpression(_:)
func (o_ OverlayAppConfiguration) SetAdImpression(impression unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdImpression:"), impression)
}

// Sets an additional value for a key; for example, a value for measuring the effectiveness of an ad campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/setAdditionalValue(_:forKey:)
func (o_ OverlayAppConfiguration) SetAdditionalValueForKey(value objc.ID, key string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdditionalValue:forKey:"), value, objc.String(key))
}

// The iTunes identifier of the recommended app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/appIdentifier
func (o_ OverlayAppConfiguration) AppIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("appIdentifier"))
	return rv
}


// SetAppIdentifier sets the value of the appIdentifier property.
// The iTunes identifier of the recommended app.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/appIdentifier
func (o_ OverlayAppConfiguration) SetAppIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppIdentifier:"), value)
}
// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/campaignToken
func (o_ OverlayAppConfiguration) CampaignToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("campaignToken"))
	return rv
}


// SetCampaignToken sets the value of the campaignToken property.
// A token you use to represent an ad campaign and measure its effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/campaignToken
func (o_ OverlayAppConfiguration) SetCampaignToken(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCampaignToken:"), value)
}
// An optional identifier for an app’s custom product page.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/customProductPageIdentifier
func (o_ OverlayAppConfiguration) CustomProductPageIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}


// SetCustomProductPageIdentifier sets the value of the customProductPageIdentifier property.
// An optional identifier for an app’s custom product page.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/customProductPageIdentifier
func (o_ OverlayAppConfiguration) SetCustomProductPageIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomProductPageIdentifier:"), value)
}
// The release ID of the latest version of your app as displayed in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/latestReleaseID
func (o_ OverlayAppConfiguration) LatestReleaseID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}


// SetLatestReleaseID sets the value of the latestReleaseID property.
// The release ID of the latest version of your app as displayed in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/latestReleaseID
func (o_ OverlayAppConfiguration) SetLatestReleaseID(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLatestReleaseID:"), value)
}
// The position of the overlay on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/position
func (o_ OverlayAppConfiguration) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the overlay on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/position
func (o_ OverlayAppConfiguration) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPosition:"), value)
}
// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/providerToken
func (o_ OverlayAppConfiguration) ProviderToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("providerToken"))
	return rv
}


// SetProviderToken sets the value of the providerToken property.
// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/providerToken
func (o_ OverlayAppConfiguration) SetProviderToken(value unsafe.Pointer) {
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

