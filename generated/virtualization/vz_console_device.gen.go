// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZConsoleDevice */


/* debug [class_header]: Header for VZConsoleDevice */
// The class instance for the [VZConsoleDevice] class.
var (
	VZConsoleDeviceClass     _VZConsoleDeviceClass
	VZConsoleDeviceClassOnce sync.Once
)

func getVZConsoleDeviceClass() _VZConsoleDeviceClass {
	VZConsoleDeviceClassOnce.Do(func() {
		VZConsoleDeviceClass = _VZConsoleDeviceClass{objc.GetClass("VZConsoleDevice")}
	})
	return VZConsoleDeviceClass
}

type _VZConsoleDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZConsoleDevice */
// An interface definition for the [VZConsoleDevice] class.
type IVZConsoleDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZConsoleDevice */
	// properties:
	ConsoleDevices() IVZConsoleDevice
	SetConsoleDevices(value IVZConsoleDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZConsoleDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZConsoleDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZConsoleDeviceClass) Alloc() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZConsoleDeviceClass) New() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZConsoleDevice) Init() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZConsoleDevice) Autorelease() VZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDevice creates a new VZConsoleDevice instance.
func NewVZConsoleDevice() VZConsoleDevice {
	return getVZConsoleDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZConsoleDevice */
// A class that represents a console device in a VM.
//
// Don’t instantiate a directly: You first configure console devices on the through a subclass of . After you create from the configuration, the console devices are available through the property. The actual type of corresponds to the type that the configuration uses. For example, a is a device of type .


// A class that represents a console device in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZConsoleDevice
type VZConsoleDevice struct {
	objectivec.Object
}

// VZConsoleDeviceFrom constructs a [VZConsoleDevice] from an unsafe.Pointer.
//
// A class that represents a console device in a VM.
func VZConsoleDeviceFrom(ptr unsafe.Pointer) VZConsoleDevice {
	return VZConsoleDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZConsoleDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZConsoleDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZConsoleDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZConsoleDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZConsoleDevice */

// The list of configured console devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/consoledevices
func (v_ VZConsoleDevice) ConsoleDevices() IVZConsoleDevice {
	rv := objc.Send[VZConsoleDevice](v_.ID, objc.Sel("consoleDevices"))
	return rv
}/* debug [instance_properties/getter]: consoleDevices */


// The list of configured console devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/consoledevices
func (v_ VZConsoleDevice) SetConsoleDevices(value IVZConsoleDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}/* debug [instance_properties/setter]: consoleDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZConsoleDevice */



