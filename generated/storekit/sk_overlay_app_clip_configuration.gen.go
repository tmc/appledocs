// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CampaignToken() objc.IObject /* cross-framework: NSString */
	SetCampaignToken(value objc.IObject /* cross-framework: NSString */)
	CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */
	SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */)
	Configuration() objc.IObject /* cross-framework: OverlayConfiguration */
	SetConfiguration(value objc.IObject /* cross-framework: OverlayConfiguration */)
	// methods:
}

// An object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding full app.


// An object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding full app.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/init(position:)
func NewOverlayAppClipConfigurationWithPosition(position OverlayPosition) OverlayAppClipConfiguration {
	instance := getOverlayAppClipConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppClipConfiguration](instance.ID, objc.Sel("initWithPosition:"), position)
	rv.Autorelease()
	return rv
}



// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appclipconfiguration/campaigntoken
func (o_ OverlayAppClipConfiguration) CampaignToken() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("campaignToken"))
	return rv
}


// A token you use to represent an ad campaign and measure its effectiveness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appclipconfiguration/campaigntoken
func (o_ OverlayAppClipConfiguration) SetCampaignToken(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCampaignToken:"), value)
}


// An identifier for a parent app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appclipconfiguration/customproductpageidentifier
func (o_ OverlayAppClipConfiguration) CustomProductPageIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("customProductPageIdentifier"))
	return rv
}


// An identifier for a parent app’s custom product page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appclipconfiguration/customproductpageidentifier
func (o_ OverlayAppClipConfiguration) SetCustomProductPageIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCustomProductPageIdentifier:"), value)
}


// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) Configuration() objc.IObject /* cross-framework: OverlayConfiguration */ {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}


// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) SetConfiguration(value objc.IObject /* cross-framework: OverlayConfiguration */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
}


