// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZPointingDeviceConfiguration */


/* debug [class_header]: Header for VZPointingDeviceConfiguration */
// The class instance for the [VZPointingDeviceConfiguration] class.
var (
	VZPointingDeviceConfigurationClass     _VZPointingDeviceConfigurationClass
	VZPointingDeviceConfigurationClassOnce sync.Once
)

func getVZPointingDeviceConfigurationClass() _VZPointingDeviceConfigurationClass {
	VZPointingDeviceConfigurationClassOnce.Do(func() {
		VZPointingDeviceConfigurationClass = _VZPointingDeviceConfigurationClass{objc.GetClass("VZPointingDeviceConfiguration")}
	})
	return VZPointingDeviceConfigurationClass
}

type _VZPointingDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZPointingDeviceConfiguration */
// An interface definition for the [VZPointingDeviceConfiguration] class.
type IVZPointingDeviceConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZPointingDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZPointingDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZPointingDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZPointingDeviceConfigurationClass) Alloc() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZPointingDeviceConfigurationClass) New() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZPointingDeviceConfiguration) Init() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZPointingDeviceConfiguration) Autorelease() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPointingDeviceConfiguration creates a new VZPointingDeviceConfiguration instance.
func NewVZPointingDeviceConfiguration() VZPointingDeviceConfiguration {
	return getVZPointingDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZPointingDeviceConfiguration */
// The base class for a pointing device configuration.
//
// Don’t instantiate a directly, use one of its subclasses like instead.


// The base class for a pointing device configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type VZPointingDeviceConfiguration struct {
	objectivec.Object
}

// VZPointingDeviceConfigurationFrom constructs a [VZPointingDeviceConfiguration] from an unsafe.Pointer.
//
// The base class for a pointing device configuration.
func VZPointingDeviceConfigurationFrom(ptr unsafe.Pointer) VZPointingDeviceConfiguration {
	return VZPointingDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZPointingDeviceConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZPointingDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZPointingDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZPointingDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZPointingDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZPointingDeviceConfiguration */



