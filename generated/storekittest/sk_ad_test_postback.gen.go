// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AdTestPostback] class.
var (
	AdTestPostbackClass     _AdTestPostbackClass
	AdTestPostbackClassOnce sync.Once
)

func getAdTestPostbackClass() _AdTestPostbackClass {
	AdTestPostbackClassOnce.Do(func() {
		AdTestPostbackClass = _AdTestPostbackClass{objc.GetClass("SKAdTestPostback")}
	})
	return AdTestPostbackClass
}

type _AdTestPostbackClass struct {
	class objc.Class
}

// An interface definition for the [AdTestPostback] class.
type IAdTestPostback interface {
	objectivec.IObject
}

// A test postback that contains ad conversion information in the testing environment.
//
// Use this class to create test postbacks to use for unit testing. In the production environment, the system creates a postback after a user installs an advertised app. The advertised app is responsible for registering the installation and may update the conversion value. The system sends the postback after a timer expires. In the testing environment, you can mimic a postback by creating it directly. You control the property values within the postback. Use it to test your app’s ability to register the app installation and update conversion values, and to test your server’s ability to receive postbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback
type AdTestPostback struct {
	objectivec.Object
}

// AdTestPostbackFrom constructs a [AdTestPostback] from an unsafe.Pointer.
//
// A test postback that contains ad conversion information in the testing environment.
func AdTestPostbackFrom(ptr unsafe.Pointer) AdTestPostback {
	return AdTestPostback{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdTestPostbackClass) Alloc() AdTestPostback {
	rv := objc.Send[AdTestPostback](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdTestPostbackClass) New() AdTestPostback {
	rv := objc.Send[AdTestPostback](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestPostback) Init() AdTestPostback {
	rv := objc.Send[AdTestPostback](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestPostback) Autorelease() AdTestPostback {
	rv := objc.Send[AdTestPostback](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestPostback creates a new AdTestPostback instance.
func NewAdTestPostback() AdTestPostback {
	return getAdTestPostbackClass().New()
}


// Creates a test postback for an in-app ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/init(version:adNetworkIdentifier:adCampaignIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:conversionValue:fidelityType:isRedownload:didWin:postbackURL:)
func NewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL(version unsafe.Pointer, adNetworkIdentifier string, adCampaignIdentifier int, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, conversionValue int, fidelityType int, isRedownload bool, didWin bool, postbackURL string) AdTestPostback {
	instance := getAdTestPostbackClass().Alloc()
	rv := objc.Send[AdTestPostback](instance.ID, objc.Sel("initWithVersion:adNetworkIdentifier:adCampaignIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:conversionValue:fidelityType:isRedownload:didWin:postbackURL:"), version, objc.String(adNetworkIdentifier), adCampaignIdentifier, appStoreItemIdentifier, sourceAppStoreItemIdentifier, conversionValue, fidelityType, isRedownload, didWin, objc.String(postbackURL))
	rv.Autorelease()
	return rv
}

// Creates a test postback for a web ad or an in-app ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/init(version:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:didWin:postbackURL:)
func NewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL(version unsafe.Pointer, adNetworkIdentifier string, sourceIdentifier string, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, sourceDomain string, fidelityType int, isRedownload bool, didWin bool, postbackURL string) AdTestPostback {
	instance := getAdTestPostbackClass().Alloc()
	rv := objc.Send[AdTestPostback](instance.ID, objc.Sel("initWithVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:didWin:postbackURL:"), version, objc.String(adNetworkIdentifier), objc.String(sourceIdentifier), appStoreItemIdentifier, sourceAppStoreItemIdentifier, objc.String(sourceDomain), fidelityType, isRedownload, didWin, objc.String(postbackURL))
	rv.Autorelease()
	return rv
}


// Creates a sequence of test postbacks for an in-app or web ad impression.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/winningPostbacks(withVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:postbackURL:)
func (ac _AdTestPostbackClass) WinningPostbacksWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadPostbackURL(version unsafe.Pointer, adNetworkIdentifier string, sourceIdentifier string, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, sourceDomain string, fidelityType int, isRedownload bool, postbackURL string) []AdTestPostback {
	rv := objc.Send[[]AdTestPostback](objc.ID(ac.class), objc.Sel("winningPostbacksWithVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:postbackURL:"), version, objc.String(adNetworkIdentifier), objc.String(sourceIdentifier), appStoreItemIdentifier, sourceAppStoreItemIdentifier, objc.String(sourceDomain), fidelityType, isRedownload, objc.String(postbackURL))
	return rv
}

// A number that represents the advertising network’s campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/adCampaignIdentifier
func (a_ AdTestPostback) AdCampaignIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("adCampaignIdentifier"))
	return rv
}

// A string that represents the advertising network’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/adNetworkIdentifier
func (a_ AdTestPostback) AdNetworkIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("adNetworkIdentifier"))
	return rv
}

// The item identifier of the app that this ad impression advertises.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/appStoreItemIdentifier
func (a_ AdTestPostback) AppStoreItemIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("appStoreItemIdentifier"))
	return rv
}

// A value that indicates a high, medium, or low conversion value for an ad postback.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/coarseConversionValue
func (a_ AdTestPostback) CoarseConversionValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("coarseConversionValue"))
	return rv
}

// An unsigned 6-bit value that the app or ad network controls.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/conversionValue
func (a_ AdTestPostback) ConversionValue() int {
	rv := objc.Send[int](a_.ID, objc.Sel("conversionValue"))
	return rv
}

// A Boolean value that indicates whether the postback won the attribution.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/didWin
func (a_ AdTestPostback) DidWin() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didWin"))
	return rv
}

// An integer that indicates the type of ad impression, StoreKit-rendered or view-through.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/fidelityType
func (a_ AdTestPostback) FidelityType() int {
	rv := objc.Send[int](a_.ID, objc.Sel("fidelityType"))
	return rv
}


// SetFidelityType sets the value of the fidelityType property.
// An integer that indicates the type of ad impression, StoreKit-rendered or view-through.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/fidelityType
func (a_ AdTestPostback) SetFidelityType(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFidelityType:"), value)
}
// The specific conversion value of an ad postback.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/fineConversionValue
func (a_ AdTestPostback) FineConversionValue() int {
	rv := objc.Send[int](a_.ID, objc.Sel("fineConversionValue"))
	return rv
}

// A Boolean value that indicates whether the user redownloaded and reinstalled the app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/isRedownload
func (a_ AdTestPostback) IsRedownload() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRedownload"))
	return rv
}

// A Boolean value that indicates whether the postback is registered in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/isRegistered
func (a_ AdTestPostback) IsRegistered() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRegistered"))
	return rv
}

// The position of this postback among all postbacks for an ad impression.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/postbackSequenceIndex
func (a_ AdTestPostback) PostbackSequenceIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("postbackSequenceIndex"))
	return rv
}

// A URL on your server where the testing environment sends test postbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/postbackURL
func (a_ AdTestPostback) PostbackURL() string {
	rv := objc.Send[string](a_.ID, objc.Sel("postbackURL"))
	return rv
}

// The item identifier of the app that displays the ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceAppStoreItemIdentifier
func (a_ AdTestPostback) SourceAppStoreItemIdentifier() int {
	rv := objc.Send[int](a_.ID, objc.Sel("sourceAppStoreItemIdentifier"))
	return rv
}

// The source of a web ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceDomain
func (a_ AdTestPostback) SourceDomain() string {
	rv := objc.Send[string](a_.ID, objc.Sel("sourceDomain"))
	return rv
}

// A string that identifies an ad campaign.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/sourceIdentifier
func (a_ AdTestPostback) SourceIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("sourceIdentifier"))
	return rv
}

// A unique transaction identifier that the system generates.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/transactionIdentifier
func (a_ AdTestPostback) TransactionIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("transactionIdentifier"))
	return rv
}

// The SKAdNetwork version that the postback uses.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/version
func (a_ AdTestPostback) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("version"))
	return rv
}


