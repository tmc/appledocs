// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioGraphicsDevice */


/* debug [class_header]: Header for VZVirtioGraphicsDevice */
// The class instance for the [VZVirtioGraphicsDevice] class.
var (
	VZVirtioGraphicsDeviceClass     _VZVirtioGraphicsDeviceClass
	VZVirtioGraphicsDeviceClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceClass() _VZVirtioGraphicsDeviceClass {
	VZVirtioGraphicsDeviceClassOnce.Do(func() {
		VZVirtioGraphicsDeviceClass = _VZVirtioGraphicsDeviceClass{objc.GetClass("VZVirtioGraphicsDevice")}
	})
	return VZVirtioGraphicsDeviceClass
}

type _VZVirtioGraphicsDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioGraphicsDevice */
// An interface definition for the [VZVirtioGraphicsDevice] class.
type IVZVirtioGraphicsDevice interface {
	IVZGraphicsDevice
	
/* debug [class_interface_properties]: Properties for VZVirtioGraphicsDevice */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioGraphicsDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioGraphicsDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioGraphicsDeviceClass) Alloc() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioGraphicsDeviceClass) New() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioGraphicsDevice) Init() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioGraphicsDevice) Autorelease() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDevice creates a new VZVirtioGraphicsDevice instance.
func NewVZVirtioGraphicsDevice() VZVirtioGraphicsDevice {
	return getVZVirtioGraphicsDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioGraphicsDevice */
// A Virtio graphics device.


// A Virtio graphics device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDevice
type VZVirtioGraphicsDevice struct {
	VZGraphicsDevice
}

// VZVirtioGraphicsDeviceFrom constructs a [VZVirtioGraphicsDevice] from an unsafe.Pointer.
//
// A Virtio graphics device.
func VZVirtioGraphicsDeviceFrom(ptr unsafe.Pointer) VZVirtioGraphicsDevice {
	return VZVirtioGraphicsDevice{
		VZGraphicsDevice: VZGraphicsDeviceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioGraphicsDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioGraphicsDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioGraphicsDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioGraphicsDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioGraphicsDevice */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioGraphicsDevice */



