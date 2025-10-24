// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKOverlayConfiguration */

/* debug [class_header]: Header for SKOverlayConfiguration */
// The class instance for the [OverlayConfiguration] class.
var (
	OverlayConfigurationClass     _OverlayConfigurationClass
	OverlayConfigurationClassOnce sync.Once
)

func getOverlayConfigurationClass() _OverlayConfigurationClass {
	OverlayConfigurationClassOnce.Do(func() {
		OverlayConfigurationClass = _OverlayConfigurationClass{objc.GetClass("SKOverlayConfiguration")}
	})
	return OverlayConfigurationClass
}

type _OverlayConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OverlayConfiguration */
// An interface definition for the [OverlayConfiguration] class.
type IOverlayConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for OverlayConfiguration */
	// properties:
	Configuration() ISKOverlayConfiguration
	SetConfiguration(value ISKOverlayConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OverlayConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OverlayConfiguration */
// Alloc allocates a new instance without initialization.
func (oc _OverlayConfigurationClass) Alloc() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OverlayConfigurationClass) New() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayConfiguration) Init() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayConfiguration) Autorelease() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayConfiguration creates a new OverlayConfiguration instance.
func NewOverlayConfiguration() OverlayConfiguration {
	return getOverlayConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OverlayConfiguration */
// The abstract superclass for all classes that represent an overlay’s attributes.

// The abstract superclass for all classes that represent an overlay’s attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Configuration-swift.class
type OverlayConfiguration struct {
	objectivec.Object
}

// OverlayConfigurationFrom constructs a [OverlayConfiguration] from an unsafe.Pointer.
//
// The abstract superclass for all classes that represent an overlay’s attributes.
func OverlayConfigurationFrom(ptr unsafe.Pointer) OverlayConfiguration {
	return OverlayConfiguration{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OverlayConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OverlayConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OverlayConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OverlayConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OverlayConfiguration */

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayConfiguration) Configuration() ISKOverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
} /* debug [instance_properties/getter]: configuration */

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/configuration-swift.property
func (o_ OverlayConfiguration) SetConfiguration(value ISKOverlayConfiguration) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConfiguration:"), value)
} /* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKOverlayConfiguration */
