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

// The class instance for the [AdImpression] class.
var (
	AdImpressionClass     _AdImpressionClass
	AdImpressionClassOnce sync.Once
)

func getAdImpressionClass() _AdImpressionClass {
	AdImpressionClassOnce.Do(func() {
		AdImpressionClass = _AdImpressionClass{objc.GetClass("SKAdImpression")}
	})
	return AdImpressionClass
}

type _AdImpressionClass struct {
	class objc.Class
}

// An interface definition for the [AdImpression] class.
type IAdImpression interface {
	objectivec.IObject
}

// A class that defines an ad impression for a view-through ad.
//
// Create a instance when you’re preparing to present a view-through ad. In the instance, you set: Values known to you, including your ad network ID, the App Store IDs of the source app and the advertised app, and the version. A value you determine – the campaign ID. Values you generate, including the timestamp, a nonce (ad-impression identifier), and the cryptographic signature. For information about generating the cryptographic signature, see . Use your instance when you call to begin presenting your view-through ad. Use the same instance when you call to end the ad presentation.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression
type AdImpression struct {
	objectivec.Object
}

// AdImpressionFrom constructs a [AdImpression] from an unsafe.Pointer.
//
// A class that defines an ad impression for a view-through ad.
func AdImpressionFrom(ptr unsafe.Pointer) AdImpression {
	return AdImpression{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdImpressionClass) Alloc() AdImpression {
	rv := objc.Send[AdImpression](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdImpressionClass) New() AdImpression {
	rv := objc.Send[AdImpression](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdImpression) Init() AdImpression {
	rv := objc.Send[AdImpression](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdImpression) Autorelease() AdImpression {
	rv := objc.Send[AdImpression](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdImpression creates a new AdImpression instance.
func NewAdImpression() AdImpression {
	return getAdImpressionClass().New()
}




// Creates an ad impression object using the supplied values.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/init(sourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:)
func NewAdImpressionWithSourceAppStoreItemIdentifierAdvertisedAppStoreItemIdentifierAdNetworkIdentifierAdCampaignIdentifierAdImpressionIdentifierTimestampSignatureVersion(sourceAppStoreItemIdentifier foundation.INumber, advertisedAppStoreItemIdentifier foundation.INumber, adNetworkIdentifier appkit.string, adCampaignIdentifier foundation.INumber, adImpressionIdentifier appkit.string, timestamp foundation.INumber, signature appkit.string, version appkit.string) AdImpression {
	instance := getAdImpressionClass().Alloc()
	rv := objc.Send[AdImpression](instance.ID, objc.Sel("initWithSourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:"), sourceAppStoreItemIdentifier, advertisedAppStoreItemIdentifier, adNetworkIdentifier, adCampaignIdentifier, adImpressionIdentifier, timestamp, signature, version)
	rv.Autorelease()
	return rv
}


// A number that represents the advertising network’s campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adCampaignIdentifier
func (a_ AdImpression) AdCampaignIdentifier() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("adCampaignIdentifier"))
	return rv
}


// SetAdCampaignIdentifier sets the value of the adCampaignIdentifier property.
// A number that represents the advertising network’s campaign.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adCampaignIdentifier
func (a_ AdImpression) SetAdCampaignIdentifier(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdCampaignIdentifier:"), value)
}

// A human-readable description of the ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adDescription
func (a_ AdImpression) AdDescription() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("adDescription"))
	return rv
}


// SetAdDescription sets the value of the adDescription property.
// A human-readable description of the ad.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adDescription
func (a_ AdImpression) SetAdDescription(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdDescription:"), value)
}

// A random value to use for added security.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adImpressionIdentifier
func (a_ AdImpression) AdImpressionIdentifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("adImpressionIdentifier"))
	return rv
}


// SetAdImpressionIdentifier sets the value of the adImpressionIdentifier property.
// A random value to use for added security.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adImpressionIdentifier
func (a_ AdImpression) SetAdImpressionIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdImpressionIdentifier:"), value)
}

// A string that represents the advertising network’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adNetworkIdentifier
func (a_ AdImpression) AdNetworkIdentifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("adNetworkIdentifier"))
	return rv
}


// SetAdNetworkIdentifier sets the value of the adNetworkIdentifier property.
// A string that represents the advertising network’s unique identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adNetworkIdentifier
func (a_ AdImpression) SetAdNetworkIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdNetworkIdentifier:"), value)
}

// The name of the entity that purchased the ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adPurchaserName
func (a_ AdImpression) AdPurchaserName() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("adPurchaserName"))
	return rv
}


// SetAdPurchaserName sets the value of the adPurchaserName property.
// The name of the entity that purchased the ad.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adPurchaserName
func (a_ AdImpression) SetAdPurchaserName(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdPurchaserName:"), value)
}

// The type of the ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adType
func (a_ AdImpression) AdType() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("adType"))
	return rv
}


// SetAdType sets the value of the adType property.
// The type of the ad.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adType
func (a_ AdImpression) SetAdType(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdType:"), value)
}

// The App Store ID of the app that the ad impression advertises.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/advertisedAppStoreItemIdentifier
func (a_ AdImpression) AdvertisedAppStoreItemIdentifier() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("advertisedAppStoreItemIdentifier"))
	return rv
}


// SetAdvertisedAppStoreItemIdentifier sets the value of the advertisedAppStoreItemIdentifier property.
// The App Store ID of the app that the ad impression advertises.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/advertisedAppStoreItemIdentifier
func (a_ AdImpression) SetAdvertisedAppStoreItemIdentifier(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAdvertisedAppStoreItemIdentifier:"), value)
}

// The advertising network’s cryptographic signature for the ad impression.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/signature
func (a_ AdImpression) Signature() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("signature"))
	return rv
}


// SetSignature sets the value of the signature property.
// The advertising network’s cryptographic signature for the ad impression.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/signature
func (a_ AdImpression) SetSignature(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSignature:"), value)
}

// The App Store ID of the app that displays the ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceAppStoreItemIdentifier
func (a_ AdImpression) SourceAppStoreItemIdentifier() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("sourceAppStoreItemIdentifier"))
	return rv
}


// SetSourceAppStoreItemIdentifier sets the value of the sourceAppStoreItemIdentifier property.
// The App Store ID of the app that displays the ad.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceAppStoreItemIdentifier
func (a_ AdImpression) SetSourceAppStoreItemIdentifier(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSourceAppStoreItemIdentifier:"), value)
}

// A four-digit integer that ad networks define to represent the ad campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceIdentifier
func (a_ AdImpression) SourceIdentifier() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("sourceIdentifier"))
	return rv
}


// SetSourceIdentifier sets the value of the sourceIdentifier property.
// A four-digit integer that ad networks define to represent the ad campaign.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceIdentifier
func (a_ AdImpression) SetSourceIdentifier(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSourceIdentifier:"), value)
}

// A number that represents the UNIX time, in milliseconds, of the ad impression.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/timestamp
func (a_ AdImpression) Timestamp() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// A number that represents the UNIX time, in milliseconds, of the ad impression.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/timestamp
func (a_ AdImpression) SetTimestamp(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimestamp:"), value)
}

// The version of the SKAdNetwork API.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/version
func (a_ AdImpression) Version() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The version of the SKAdNetwork API.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/version
func (a_ AdImpression) SetVersion(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersion:"), value)
}


