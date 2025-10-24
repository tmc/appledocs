// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SKOverlayAppConfiguration */

/* debug [class_header]: Header for SKOverlayAppConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OverlayAppConfiguration */
// An interface definition for the [OverlayAppConfiguration] class.
type IOverlayAppConfiguration interface {
	IOverlayConfiguration

	/* debug [class_interface_properties]: Properties for OverlayAppConfiguration */
	// properties:
	AdAttributionReengagementURL() foundation.URL
	SetAdAttributionReengagementURL(value foundation.URL)
	AppImpression() unsafe.Pointer
	SetAppImpression(value unsafe.Pointer)
	Configuration() ISKOverlayConfiguration
	SetConfiguration(value ISKOverlayConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OverlayAppConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OverlayAppConfiguration */
// Alloc allocates a new instance without initialization.
func (oc _OverlayAppConfigurationClass) Alloc() OverlayAppConfiguration {
	rv := objc.Send[OverlayAppConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OverlayAppConfiguration */
// An object that represents the attributes of an overlay you use to recommend another app on the App Store.

// An object that represents the attributes of an overlay you use to recommend another app on the App Store.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OverlayAppConfiguration */

// Creates an object that represents the attributes of an overlay you use to recommend another app on the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppConfiguration/init(appIdentifier:position:)
func NewOverlayAppConfigurationWithAppIdentifierPosition(appIdentifier objc.IObject /* cross-framework: NSString */, position OverlayPosition) OverlayAppConfiguration {
	instance := getOverlayAppConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppConfiguration](instance.ID, objc.Sel("initWithAppIdentifier:position:"), appIdentifier, position)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewOverlayAppConfigurationWithAppIdentifierPosition */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OverlayAppConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OverlayAppConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OverlayAppConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OverlayAppConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) AdAttributionReengagementURL() foundation.URL {
	rv := objc.Send[foundation.URL](o_.ID, objc.Sel("adAttributionReengagementURL"))
	return rv
} /* debug [instance_properties/getter]: adAttributionReengagementURL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/adattributionreengagementurl
func (o_ OverlayAppConfiguration) SetAdAttributionReengagementURL(value foundation.URL) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAdAttributionReengagementURL:"), value)
} /* debug [instance_properties/setter]: adAttributionReengagementURL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) AppImpression() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("appImpression"))
	return rv
} /* debug [instance_properties/getter]: appImpression */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/appconfiguration/appimpression
func (o_ OverlayAppConfiguration) SetAppImpression(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAppImpression:"), value)
} /* debug [instance_properties/setter]: appImpression */

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) Configuration() ISKOverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
} /* debug [instance_properties/getter]: configuration */

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppConfiguration) SetConfiguration(value ISKOverlayConfiguration) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
} /* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKOverlayAppConfiguration */
