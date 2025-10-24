// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZMacGraphicsDevice */


/* debug [class_header]: Header for VZMacGraphicsDevice */
// The class instance for the [VZMacGraphicsDevice] class.
var (
	VZMacGraphicsDeviceClass     _VZMacGraphicsDeviceClass
	VZMacGraphicsDeviceClassOnce sync.Once
)

func getVZMacGraphicsDeviceClass() _VZMacGraphicsDeviceClass {
	VZMacGraphicsDeviceClassOnce.Do(func() {
		VZMacGraphicsDeviceClass = _VZMacGraphicsDeviceClass{objc.GetClass("VZMacGraphicsDevice")}
	})
	return VZMacGraphicsDeviceClass
}

type _VZMacGraphicsDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMacGraphicsDevice */
// An interface definition for the [VZMacGraphicsDevice] class.
type IVZMacGraphicsDevice interface {
	IVZGraphicsDevice
	
/* debug [class_interface_properties]: Properties for VZMacGraphicsDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMacGraphicsDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMacGraphicsDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZMacGraphicsDeviceClass) Alloc() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacGraphicsDeviceClass) New() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacGraphicsDevice) Init() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacGraphicsDevice) Autorelease() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDevice creates a new VZMacGraphicsDevice instance.
func NewVZMacGraphicsDevice() VZMacGraphicsDevice {
	return getVZMacGraphicsDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMacGraphicsDevice */
// An object that represents a Mac graphics device.
//
// You don’t instantiate a   directly. Graphics devices are first configured on the through a subclass of  .  When the framework creates a VZVirtualMachine from the configuration, the graphics devices are available through the property.


// An object that represents a Mac graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type VZMacGraphicsDevice struct {
	VZGraphicsDevice
}

// VZMacGraphicsDeviceFrom constructs a [VZMacGraphicsDevice] from an unsafe.Pointer.
//
// An object that represents a Mac graphics device.
func VZMacGraphicsDeviceFrom(ptr unsafe.Pointer) VZMacGraphicsDevice {
	return VZMacGraphicsDevice{
		VZGraphicsDevice: VZGraphicsDeviceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMacGraphicsDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMacGraphicsDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMacGraphicsDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMacGraphicsDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMacGraphicsDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMacGraphicsDevice */



