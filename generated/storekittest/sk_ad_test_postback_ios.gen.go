//go:build darwin && ios

// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AdTestPostback


// iOS-only properties

// A number that represents the advertising network’s campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/adCampaignIdentifier
func (a_ AdTestPostback) AdCampaignIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("adCampaignIdentifier"))
	return rv
}

// A string that represents the advertising network’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/adNetworkIdentifier
func (a_ AdTestPostback) AdNetworkIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("adNetworkIdentifier"))
	return rv
}

// The item identifier of the app that this ad impression advertises.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/appStoreItemIdentifier
func (a_ AdTestPostback) AppStoreItemIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("appStoreItemIdentifier"))
	return rv
}

// A value that indicates a high, medium, or low conversion value for an ad postback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/coarseConversionValue
func (a_ AdTestPostback) CoarseConversionValue() AdNetworkCoarseConversionValue /* not a class type */ {
	rv := objc.Send[AdNetworkCoarseConversionValue](a_.ID, objc.Sel("coarseConversionValue"))
	return rv
}

// An unsigned 6-bit value that the app or ad network controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/conversionValue
func (a_ AdTestPostback) ConversionValue() int {
	rv := objc.Send[int](a_.ID, objc.Sel("conversionValue"))
	return rv
}

// A Boolean value that indicates whether the postback won the attribution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/didWin
func (a_ AdTestPostback) DidWin() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didWin"))
	return rv
}

// An integer that indicates the type of ad impression, StoreKit-rendered or view-through.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/fidelityType
func (a_ AdTestPostback) FidelityType() int {
	rv := objc.Send[int](a_.ID, objc.Sel("fidelityType"))
	return rv
}
func (a_ AdTestPostback) SetFidelityType(value int) {
	a_.ID.Send(objc.RegisterName("setFidelityType:"), value)
}

// The specific conversion value of an ad postback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/fineConversionValue
func (a_ AdTestPostback) FineConversionValue() int {
	rv := objc.Send[int](a_.ID, objc.Sel("fineConversionValue"))
	return rv
}

// A Boolean value that indicates whether the user redownloaded and reinstalled the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/isRedownload
func (a_ AdTestPostback) IsRedownload() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRedownload"))
	return rv
}

// A Boolean value that indicates whether the postback is registered in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/isRegistered
func (a_ AdTestPostback) IsRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRegistered"))
	return rv
}

// The position of this postback among all postbacks for an ad impression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/postbackSequenceIndex
func (a_ AdTestPostback) PostbackSequenceIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("postbackSequenceIndex"))
	return rv
}

// A URL on your server where the testing environment sends test postbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/postbackURL
func (a_ AdTestPostback) PostbackURL() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("postbackURL"))
	return rv
}

// The item identifier of the app that displays the ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceAppStoreItemIdentifier
func (a_ AdTestPostback) SourceAppStoreItemIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("sourceAppStoreItemIdentifier"))
	return rv
}

// The source of a web ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceDomain
func (a_ AdTestPostback) SourceDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("sourceDomain"))
	return rv
}

// A string that identifies an ad campaign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceIdentifier
func (a_ AdTestPostback) SourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("sourceIdentifier"))
	return rv
}

// A unique transaction identifier that the system generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/transactionIdentifier
func (a_ AdTestPostback) TransactionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("transactionIdentifier"))
	return rv
}

// The SKAdNetwork version that the postback uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/version
func (a_ AdTestPostback) Version() objc.IObject /* cross-framework: AdTestPostbackVersion */ {
	rv := objc.Send[AdTestPostbackVersion](a_.ID, objc.Sel("version"))
	return rv
}




