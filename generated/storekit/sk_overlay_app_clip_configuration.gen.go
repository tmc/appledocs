// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OverlayAppClipConfiguration] class.
var (
	OverlayAppClipConfigurationClass     _OverlayAppClipConfigurationClass
	OverlayAppClipConfigurationClassOnce sync.Once
)

func getOverlayAppClipConfigurationClass() _OverlayAppClipConfigurationClass {
	OverlayAppClipConfigurationClassOnce.Do(func() {
		OverlayAppClipConfigurationClass = _OverlayAppClipConfigurationClass{objc.GetClass("SKOverlayAppClipConfiguration")}
	})
	return OverlayAppClipConfigurationClass
}

type _OverlayAppClipConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [OverlayAppClipConfiguration] class.
type IOverlayAppClipConfiguration interface {
	IOverlayConfiguration
	AdditionalValueForKey(key string) objc.ID
	SetAdditionalValueForKey(value objectivec.IObject, key string)
	CampaignToken() string
	SetCampaignToken(value string)
	CustomProductPageIdentifier() string
	SetCustomProductPageIdentifier(value string)
	LatestReleaseID() string
	SetLatestReleaseID(value string)
	Position() OverlayPosition
	SetPosition(value OverlayPosition)
	ProviderToken() string
	SetProviderToken(value string)
	Configuration() SKOverlayConfiguration
	SetConfiguration(value ISKOverlayConfiguration)
}

// An object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding full app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration
type OverlayAppClipConfiguration struct {
	OverlayConfiguration
}

// OverlayAppClipConfigurationFrom constructs a [OverlayAppClipConfiguration] from an unsafe.Pointer.
//
// An object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding full app.
func OverlayAppClipConfigurationFrom(ptr unsafe.Pointer) OverlayAppClipConfiguration {
	return OverlayAppClipConfiguration{
		OverlayConfiguration: OverlayConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OverlayAppClipConfigurationClass) Alloc() OverlayAppClipConfiguration {
	rv := objc.Send[OverlayAppClipConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OverlayAppClipConfigurationClass) New() OverlayAppClipConfiguration {
	rv := objc.Send[OverlayAppClipConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayAppClipConfiguration) Init() OverlayAppClipConfiguration {
	rv := objc.Send[OverlayAppClipConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayAppClipConfiguration) Autorelease() OverlayAppClipConfiguration {
	rv := objc.Send[OverlayAppClipConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayAppClipConfiguration creates a new OverlayAppClipConfiguration instance.
func NewOverlayAppClipConfiguration() OverlayAppClipConfiguration {
	return getOverlayAppClipConfigurationClass().New()
}




// Creates an object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/init(position:)
func NewOverlayAppClipConfigurationWithPosition(position OverlayPosition) OverlayAppClipConfiguration {
	instance := getOverlayAppClipConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppClipConfiguration](instance.ID, objc.Sel("initWithPosition:"), position)
	rv.Autorelease()
	return rv
}


// Returns the object associated with the key.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/additionalValue(forKey:)
func (o_ OverlayAppClipConfiguration) AdditionalValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("additionalValueForKey:"), objc.String(key))
	return rv
}

// Sets an additional value for a key, such as a value for measuring the effectiveness of an ad campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/setAdditionalValue(_:forKey:)
func (o_ OverlayAppClipConfiguration) SetAdditionalValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdditionalValue:forKey:"), value, objc.String(key))
}

// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/campaignToken
func (o_ OverlayAppClipConfiguration) CampaignToken() string {
	rv := objc.Send[string](o_.ID, objc.Sel("campaignToken"))
	return rv
}


// SetCampaignToken sets the value of the campaignToken property.
// A token you use to represent an ad campaign and measure its effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/campaignToken
func (o_ OverlayAppClipConfiguration) SetCampaignToken(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCampaignToken:"), objc.String(value))
}

// An identifier for a parent app’s custom product page.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/customProductPageIdentifier
func (o_ OverlayAppClipConfiguration) CustomProductPageIdentifier() string {
	rv := objc.Send[string](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}


// SetCustomProductPageIdentifier sets the value of the customProductPageIdentifier property.
// An identifier for a parent app’s custom product page.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/customProductPageIdentifier
func (o_ OverlayAppClipConfiguration) SetCustomProductPageIdentifier(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomProductPageIdentifier:"), objc.String(value))
}

// The release ID of the latest version of your parent app as displayed in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/latestReleaseID
func (o_ OverlayAppClipConfiguration) LatestReleaseID() string {
	rv := objc.Send[string](o_.ID, objc.Sel("latestReleaseID"))
	return rv
}


// SetLatestReleaseID sets the value of the latestReleaseID property.
// The release ID of the latest version of your parent app as displayed in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/latestReleaseID
func (o_ OverlayAppClipConfiguration) SetLatestReleaseID(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setLatestReleaseID:"), objc.String(value))
}

// The position of the overlay on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/position
func (o_ OverlayAppClipConfiguration) Position() OverlayPosition {
	rv := objc.Send[OverlayPosition](o_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the overlay on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/position
func (o_ OverlayAppClipConfiguration) SetPosition(value OverlayPosition) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPosition:"), value)
}

// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/providerToken
func (o_ OverlayAppClipConfiguration) ProviderToken() string {
	rv := objc.Send[string](o_.ID, objc.Sel("providerToken"))
	return rv
}


// SetProviderToken sets the value of the providerToken property.
// A token that represents the provider of an app promotion campaign, and that you use to measure the campaign’s effectiveness.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/providerToken
func (o_ OverlayAppClipConfiguration) SetProviderToken(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setProviderToken:"), objc.String(value))
}

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) Configuration() SKOverlayConfiguration {
	rv := objc.Send[SKOverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// An overlay’s attributes; for example, its position on the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) SetConfiguration(value ISKOverlayConfiguration) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
}


