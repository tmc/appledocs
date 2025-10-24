// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZConsoleDeviceConfiguration */


/* debug [class_header]: Header for VZConsoleDeviceConfiguration */
// The class instance for the [VZConsoleDeviceConfiguration] class.
var (
	VZConsoleDeviceConfigurationClass     _VZConsoleDeviceConfigurationClass
	VZConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZConsoleDeviceConfigurationClass() _VZConsoleDeviceConfigurationClass {
	VZConsoleDeviceConfigurationClassOnce.Do(func() {
		VZConsoleDeviceConfigurationClass = _VZConsoleDeviceConfigurationClass{objc.GetClass("VZConsoleDeviceConfiguration")}
	})
	return VZConsoleDeviceConfigurationClass
}

type _VZConsoleDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZConsoleDeviceConfiguration */
// An interface definition for the [VZConsoleDeviceConfiguration] class.
type IVZConsoleDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZConsoleDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZConsoleDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZConsoleDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZConsoleDeviceConfigurationClass) Alloc() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZConsoleDeviceConfigurationClass) New() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsoleDeviceConfiguration) Init() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsoleDeviceConfiguration) Autorelease() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDeviceConfiguration creates a new VZConsoleDeviceConfiguration instance.
func NewVZConsoleDeviceConfiguration() VZConsoleDeviceConfiguration {
	return getVZConsoleDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZConsoleDeviceConfiguration */
// The base class for a console device configuration.
//
// Don’t instantiate VZConsoleDeviceConfiguration directly, instead use one of its subclasses like instead.


// The base class for a console device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type VZConsoleDeviceConfiguration struct {
	objectivec.Object
}

// VZConsoleDeviceConfigurationFrom constructs a [VZConsoleDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a console device configuration.
func VZConsoleDeviceConfigurationFrom(ptr unsafe.Pointer) VZConsoleDeviceConfiguration {
	return VZConsoleDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZConsoleDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZConsoleDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZConsoleDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZConsoleDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZConsoleDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZConsoleDeviceConfiguration */



