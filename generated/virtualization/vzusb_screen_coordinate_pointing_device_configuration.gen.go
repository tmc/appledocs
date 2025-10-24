// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZUSBScreenCoordinatePointingDeviceConfiguration */


/* debug [class_header]: Header for VZUSBScreenCoordinatePointingDeviceConfiguration */
// The class instance for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
var (
	VZUSBScreenCoordinatePointingDeviceConfigurationClass     _VZUSBScreenCoordinatePointingDeviceConfigurationClass
	VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce sync.Once
)

func getVZUSBScreenCoordinatePointingDeviceConfigurationClass() _VZUSBScreenCoordinatePointingDeviceConfigurationClass {
	VZUSBScreenCoordinatePointingDeviceConfigurationClassOnce.Do(func() {
		VZUSBScreenCoordinatePointingDeviceConfigurationClass = _VZUSBScreenCoordinatePointingDeviceConfigurationClass{objc.GetClass("VZUSBScreenCoordinatePointingDeviceConfiguration")}
	})
	return VZUSBScreenCoordinatePointingDeviceConfigurationClass
}

type _VZUSBScreenCoordinatePointingDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZUSBScreenCoordinatePointingDeviceConfiguration */
// An interface definition for the [VZUSBScreenCoordinatePointingDeviceConfiguration] class.
type IVZUSBScreenCoordinatePointingDeviceConfiguration interface {
	IVZPointingDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZUSBScreenCoordinatePointingDeviceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZUSBScreenCoordinatePointingDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZUSBScreenCoordinatePointingDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBScreenCoordinatePointingDeviceConfigurationClass) Alloc() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZUSBScreenCoordinatePointingDeviceConfigurationClass) New() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBScreenCoordinatePointingDeviceConfiguration) Init() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBScreenCoordinatePointingDeviceConfiguration) Autorelease() VZUSBScreenCoordinatePointingDeviceConfiguration {
	rv := objc.Send[VZUSBScreenCoordinatePointingDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBScreenCoordinatePointingDeviceConfiguration creates a new VZUSBScreenCoordinatePointingDeviceConfiguration instance.
func NewVZUSBScreenCoordinatePointingDeviceConfiguration() VZUSBScreenCoordinatePointingDeviceConfiguration {
	return getVZUSBScreenCoordinatePointingDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZUSBScreenCoordinatePointingDeviceConfiguration */
// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
//
// A can use this device to send pointer events to the VM.


// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBScreenCoordinatePointingDeviceConfiguration
type VZUSBScreenCoordinatePointingDeviceConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZUSBScreenCoordinatePointingDeviceConfigurationFrom constructs a [VZUSBScreenCoordinatePointingDeviceConfiguration] from an unsafe.Pointer.
//
// An object that defines the configuration for a USB pointing device that reports absolute coordinates.
func VZUSBScreenCoordinatePointingDeviceConfigurationFrom(ptr unsafe.Pointer) VZUSBScreenCoordinatePointingDeviceConfiguration {
	return VZUSBScreenCoordinatePointingDeviceConfiguration{
		VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZUSBScreenCoordinatePointingDeviceConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZUSBScreenCoordinatePointingDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZUSBScreenCoordinatePointingDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZUSBScreenCoordinatePointingDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZUSBScreenCoordinatePointingDeviceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZUSBScreenCoordinatePointingDeviceConfiguration */


