// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKOverlayAppClipConfiguration */


/* debug [class_header]: Header for SKOverlayAppClipConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OverlayAppClipConfiguration */
// An interface definition for the [OverlayAppClipConfiguration] class.
type IOverlayAppClipConfiguration interface {
	IOverlayConfiguration
	
/* debug [class_interface_properties]: Properties for OverlayAppClipConfiguration */
	// properties:
	Configuration() ISKOverlayConfiguration
	SetConfiguration(value ISKOverlayConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OverlayAppClipConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OverlayAppClipConfiguration */
// Alloc allocates a new instance without initialization.
func (oc _OverlayAppClipConfigurationClass) Alloc() OverlayAppClipConfiguration {
	rv := objc.Send[OverlayAppClipConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OverlayAppClipConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OverlayAppClipConfiguration */

// Creates an object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/AppClipConfiguration/init(position:)
func NewOverlayAppClipConfigurationWithPosition(position OverlayPosition) OverlayAppClipConfiguration {
	instance := getOverlayAppClipConfigurationClass().Alloc()
	rv := objc.Send[OverlayAppClipConfiguration](instance.ID, objc.Sel("initWithPosition:"), position)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOverlayAppClipConfigurationWithPosition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OverlayAppClipConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OverlayAppClipConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OverlayAppClipConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OverlayAppClipConfiguration */

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) Configuration() ISKOverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayAppClipConfiguration) SetConfiguration(value ISKOverlayConfiguration) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKOverlayAppClipConfiguration */


