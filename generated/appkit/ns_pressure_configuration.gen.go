// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPressureConfiguration */


/* debug [class_header]: Header for NSPressureConfiguration */
// The class instance for the [PressureConfiguration] class.
var (
	PressureConfigurationClass     _PressureConfigurationClass
	PressureConfigurationClassOnce sync.Once
)

func getPressureConfigurationClass() _PressureConfigurationClass {
	PressureConfigurationClassOnce.Do(func() {
		PressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}
	})
	return PressureConfigurationClass
}

type _PressureConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PressureConfiguration */
// An interface definition for the [PressureConfiguration] class.
type IPressureConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PressureConfiguration */
	// properties:
	PressureBehavior() PressureBehavior
	SetPressureBehavior(value PressureBehavior)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PressureConfiguration */
	// methods:
	Set()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PressureConfiguration */
// Alloc allocates a new instance without initialization.
func (pc _PressureConfigurationClass) Alloc() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PressureConfigurationClass) New() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PressureConfiguration) Init() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PressureConfiguration) Autorelease() PressureConfiguration {
	rv := objc.Send[PressureConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPressureConfiguration creates a new PressureConfiguration instance.
func NewPressureConfiguration() PressureConfiguration {
	return getPressureConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PressureConfiguration */
// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events.
//
// Use an object to configure the behavior and progression of a Force Touch trackpad when it responds to a mouse drag or pressure event sequence. Pressure configurations are assigned to views ( ) and gesture recognizers ( ).


// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration
type PressureConfiguration struct {
	objectivec.Object
}

// PressureConfigurationFrom constructs a [PressureConfiguration] from an unsafe.Pointer.
//
// An encapsulation of the behavior and progression of a Force Touch trackpad as it responds to specific events.
func PressureConfigurationFrom(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PressureConfiguration */

// Initializes a pressure configuration object with a specified pressure behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/init(pressureBehavior:)
func NewPressureConfigurationWithPressureBehavior(pressureBehavior PressureBehavior) PressureConfiguration {
	instance := getPressureConfigurationClass().Alloc()
	rv := objc.Send[PressureConfiguration](instance.ID, objc.Sel("initWithPressureBehavior:"), pressureBehavior)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPressureConfigurationWithPressureBehavior */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PressureConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PressureConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PressureConfiguration */

// Changes the pressure configuration of the trackpad to the initialized pressure configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/set()
func (p_ PressureConfiguration) Set() {
	objc.Send[objc.ID](p_.ID, objc.Sel("set"))
}/* debug [instance_methods/method]: Set */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PressureConfiguration */

// The pressure behavior of the pressure configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration/pressurebehavior
func (p_ PressureConfiguration) PressureBehavior() PressureBehavior {
	rv := objc.Send[PressureBehavior](p_.ID, objc.Sel("pressureBehavior"))
	return rv
}/* debug [instance_properties/getter]: pressureBehavior */


// The pressure behavior of the pressure configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspressureconfiguration/pressurebehavior
func (p_ PressureConfiguration) SetPressureBehavior(value PressureBehavior) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPressureBehavior:"), value)
}/* debug [instance_properties/setter]: pressureBehavior */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPressureConfiguration */


