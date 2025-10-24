// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioConsoleDevice */


/* debug [class_header]: Header for VZVirtioConsoleDevice */
// The class instance for the [VZVirtioConsoleDevice] class.
var (
	VZVirtioConsoleDeviceClass     _VZVirtioConsoleDeviceClass
	VZVirtioConsoleDeviceClassOnce sync.Once
)

func getVZVirtioConsoleDeviceClass() _VZVirtioConsoleDeviceClass {
	VZVirtioConsoleDeviceClassOnce.Do(func() {
		VZVirtioConsoleDeviceClass = _VZVirtioConsoleDeviceClass{objc.GetClass("VZVirtioConsoleDevice")}
	})
	return VZVirtioConsoleDeviceClass
}

type _VZVirtioConsoleDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsoleDevice */
// An interface definition for the [VZVirtioConsoleDevice] class.
type IVZVirtioConsoleDevice interface {
	IVZConsoleDevice
	
/* debug [class_interface_properties]: Properties for VZVirtioConsoleDevice */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Ports() IVZVirtioConsolePortArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsoleDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsoleDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceClass) Alloc() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsoleDeviceClass) New() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsoleDevice) Init() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsoleDevice) Autorelease() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDevice creates a new VZVirtioConsoleDevice instance.
func NewVZVirtioConsoleDevice() VZVirtioConsoleDevice {
	return getVZVirtioConsoleDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsoleDevice */
// A class that represents a Virtio console device in a virtual machine.


// A class that represents a Virtio console device in a virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice
type VZVirtioConsoleDevice struct {
	VZConsoleDevice
}

// VZVirtioConsoleDeviceFrom constructs a [VZVirtioConsoleDevice] from an unsafe.Pointer.
//
// A class that represents a Virtio console device in a virtual machine.
func VZVirtioConsoleDeviceFrom(ptr unsafe.Pointer) VZVirtioConsoleDevice {
	return VZVirtioConsoleDevice{
		VZConsoleDevice: VZConsoleDeviceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsoleDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsoleDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsoleDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsoleDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsoleDevice */

// The delegate object for the console device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/delegate
func (v_ VZVirtioConsoleDevice) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the console device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/delegate
func (v_ VZVirtioConsoleDevice) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The array of console ports that a specific device uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/ports
func (v_ VZVirtioConsoleDevice) Ports() IVZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("ports"))
	return rv
}/* debug [instance_properties/getter]: ports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsoleDevice */



