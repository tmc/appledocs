// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	// methods:
}

// A class that defines an ad impression for a view-through ad.
//
// Create a instance when you’re preparing to present a view-through ad. In the instance, you set: Values known to you, including your ad network ID, the App Store IDs of the source app and the advertised app, and the version. A value you determine – the campaign ID. Values you generate, including the timestamp, a nonce (ad-impression identifier), and the cryptographic signature. For information about generating the cryptographic signature, see . Use your instance when you call to begin presenting your view-through ad. Use the same instance when you call to end the ad presentation.


// A class that defines an ad impression for a view-through ad.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/init(sourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:)
func NewAdImpressionWithSourceAppStoreItemIdentifierAdvertisedAppStoreItemIdentifierAdNetworkIdentifierAdCampaignIdentifierAdImpressionIdentifierTimestampSignatureVersion(sourceAppStoreItemIdentifier objc.IObject /* cross-framework: NSNumber */, advertisedAppStoreItemIdentifier objc.IObject /* cross-framework: NSNumber */, adNetworkIdentifier objc.IObject /* cross-framework: NSString */, adCampaignIdentifier objc.IObject /* cross-framework: NSNumber */, adImpressionIdentifier objc.IObject /* cross-framework: NSString */, timestamp objc.IObject /* cross-framework: NSNumber */, signature objc.IObject /* cross-framework: NSString */, version objc.IObject /* cross-framework: NSString */) AdImpression {
	instance := getAdImpressionClass().Alloc()
	rv := objc.Send[AdImpression](instance.ID, objc.Sel("initWithSourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:"), sourceAppStoreItemIdentifier, advertisedAppStoreItemIdentifier, adNetworkIdentifier, adCampaignIdentifier, adImpressionIdentifier, timestamp, signature, version)
	rv.Autorelease()
	return rv
}



