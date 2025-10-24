// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacGraphicsDeviceConfiguration */


/* debug [class_header]: Header for VZMacGraphicsDeviceConfiguration */
// The class instance for the [VZMacGraphicsDeviceConfiguration] class.
var (
	VZMacGraphicsDeviceConfigurationClass     _VZMacGraphicsDeviceConfigurationClass
	VZMacGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDeviceConfigurationClass() _VZMacGraphicsDeviceConfigurationClass {
	VZMacGraphicsDeviceConfigurationClassOnce.Do(func() {
		VZMacGraphicsDeviceConfigurationClass = _VZMacGraphicsDeviceConfigurationClass{objc.GetClass("VZMacGraphicsDeviceConfiguration")}
	})
	return VZMacGraphicsDeviceConfigurationClass
}

type _VZMacGraphicsDeviceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacGraphicsDeviceConfiguration */
// An interface definition for the [VZMacGraphicsDeviceConfiguration] class.
type IVZMacGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration
	
/* debug [class_interface_properties]: Properties for VZMacGraphicsDeviceConfiguration */
	// properties:
	Displays() []VZMacGraphicsDisplayConfiguration
	SetDisplays(value []VZMacGraphicsDisplayConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacGraphicsDeviceConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacGraphicsDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDeviceConfigurationClass) Alloc() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacGraphicsDeviceConfigurationClass) New() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDeviceConfiguration) Init() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDeviceConfiguration) Autorelease() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDeviceConfiguration creates a new VZMacGraphicsDeviceConfiguration instance.
func NewVZMacGraphicsDeviceConfiguration() VZMacGraphicsDeviceConfiguration {
	return getVZMacGraphicsDeviceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacGraphicsDeviceConfiguration */
// Configuration for a display attached to a Mac graphics device.
//
// Use this device to attach a display that’s shown in a .


// Configuration for a display attached to a Mac graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type VZMacGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZMacGraphicsDeviceConfigurationFrom constructs a [VZMacGraphicsDeviceConfiguration] from an unsafe.Pointer.
//
// Configuration for a display attached to a Mac graphics device.
func VZMacGraphicsDeviceConfigurationFrom(ptr unsafe.Pointer) VZMacGraphicsDeviceConfiguration {
	return VZMacGraphicsDeviceConfiguration{
		VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacGraphicsDeviceConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacGraphicsDeviceConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacGraphicsDeviceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacGraphicsDeviceConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacGraphicsDeviceConfiguration */

// The displays associated with this graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/displays
func (v_ VZMacGraphicsDeviceConfiguration) Displays() []VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[[]VZMacGraphicsDisplayConfiguration](v_.ID, objc.Sel("displays"))
	return rv
}/* debug [instance_properties/getter]: displays */


// The displays associated with this graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/displays
func (v_ VZMacGraphicsDeviceConfiguration) SetDisplays(value []VZMacGraphicsDisplayConfiguration) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisplays:"), nsArray)
}/* debug [instance_properties/setter]: displays */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacGraphicsDeviceConfiguration */


