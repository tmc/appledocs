// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKAdImpression */

/* debug [class_header]: Header for SKAdImpression */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for AdImpression */
// An interface definition for the [AdImpression] class.
type IAdImpression interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for AdImpression */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for AdImpression */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for AdImpression */
// Alloc allocates a new instance without initialization.
func (ac _AdImpressionClass) Alloc() AdImpression {
	rv := objc.Send[AdImpression](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for AdImpression */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for AdImpression */

// Creates an ad impression object using the supplied values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdImpression/init(sourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:)
func NewAdImpressionWithSourceAppStoreItemIdentifierAdvertisedAppStoreItemIdentifierAdNetworkIdentifierAdCampaignIdentifierAdImpressionIdentifierTimestampSignatureVersion(sourceAppStoreItemIdentifier objc.IObject /* cross-framework: NSNumber */, advertisedAppStoreItemIdentifier objc.IObject /* cross-framework: NSNumber */, adNetworkIdentifier objc.IObject /* cross-framework: NSString */, adCampaignIdentifier objc.IObject /* cross-framework: NSNumber */, adImpressionIdentifier objc.IObject /* cross-framework: NSString */, timestamp objc.IObject /* cross-framework: NSNumber */, signature objc.IObject /* cross-framework: NSString */, version objc.IObject /* cross-framework: NSString */) AdImpression {
	instance := getAdImpressionClass().Alloc()
	rv := objc.Send[AdImpression](instance.ID, objc.Sel("initWithSourceAppStoreItemIdentifier:advertisedAppStoreItemIdentifier:adNetworkIdentifier:adCampaignIdentifier:adImpressionIdentifier:timestamp:signature:version:"), sourceAppStoreItemIdentifier, advertisedAppStoreItemIdentifier, adNetworkIdentifier, adCampaignIdentifier, adImpressionIdentifier, timestamp, signature, version)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewAdImpressionWithSourceAppStoreItemIdentifierAdvertisedAppStoreItemIdentifierAdNetworkIdentifierAdCampaignIdentifierAdImpressionIdentifierTimestampSignatureVersion */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for AdImpression */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for AdImpression */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for AdImpression */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for AdImpression */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKAdImpression */
