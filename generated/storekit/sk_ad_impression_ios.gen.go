//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AdImpression


// iOS-only properties

// A number that represents the advertising network’s campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adCampaignIdentifier
func (a_ AdImpression) AdCampaignIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("adCampaignIdentifier"))
	return rv
}
func (a_ AdImpression) SetAdCampaignIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	a_.ID.Send(objc.RegisterName("setAdCampaignIdentifier:"), value)
}

// A human-readable description of the ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adDescription
func (a_ AdImpression) AdDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adDescription"))
	return rv
}
func (a_ AdImpression) SetAdDescription(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setAdDescription:"), value)
}

// A random value to use for added security.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adImpressionIdentifier
func (a_ AdImpression) AdImpressionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adImpressionIdentifier"))
	return rv
}
func (a_ AdImpression) SetAdImpressionIdentifier(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setAdImpressionIdentifier:"), value)
}

// A string that represents the advertising network’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adNetworkIdentifier
func (a_ AdImpression) AdNetworkIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adNetworkIdentifier"))
	return rv
}
func (a_ AdImpression) SetAdNetworkIdentifier(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setAdNetworkIdentifier:"), value)
}

// The name of the entity that purchased the ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adPurchaserName
func (a_ AdImpression) AdPurchaserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adPurchaserName"))
	return rv
}
func (a_ AdImpression) SetAdPurchaserName(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setAdPurchaserName:"), value)
}

// The type of the ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/adType
func (a_ AdImpression) AdType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adType"))
	return rv
}
func (a_ AdImpression) SetAdType(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setAdType:"), value)
}

// The App Store ID of the app that the ad impression advertises.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/advertisedAppStoreItemIdentifier
func (a_ AdImpression) AdvertisedAppStoreItemIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("advertisedAppStoreItemIdentifier"))
	return rv
}
func (a_ AdImpression) SetAdvertisedAppStoreItemIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	a_.ID.Send(objc.RegisterName("setAdvertisedAppStoreItemIdentifier:"), value)
}

// The advertising network’s cryptographic signature for the ad impression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/signature
func (a_ AdImpression) Signature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("signature"))
	return rv
}
func (a_ AdImpression) SetSignature(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setSignature:"), value)
}

// The App Store ID of the app that displays the ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceAppStoreItemIdentifier
func (a_ AdImpression) SourceAppStoreItemIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("sourceAppStoreItemIdentifier"))
	return rv
}
func (a_ AdImpression) SetSourceAppStoreItemIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	a_.ID.Send(objc.RegisterName("setSourceAppStoreItemIdentifier:"), value)
}

// A four-digit integer that ad networks define to represent the ad campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/sourceIdentifier
func (a_ AdImpression) SourceIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("sourceIdentifier"))
	return rv
}
func (a_ AdImpression) SetSourceIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	a_.ID.Send(objc.RegisterName("setSourceIdentifier:"), value)
}

// A number that represents the UNIX time, in milliseconds, of the ad impression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/timestamp
func (a_ AdImpression) Timestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("timestamp"))
	return rv
}
func (a_ AdImpression) SetTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	a_.ID.Send(objc.RegisterName("setTimestamp:"), value)
}

// The version of the SKAdNetwork API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/version
func (a_ AdImpression) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("version"))
	return rv
}
func (a_ AdImpression) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	a_.ID.Send(objc.RegisterName("setVersion:"), value)
}




