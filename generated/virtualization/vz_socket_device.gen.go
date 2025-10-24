// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSocketDevice */


/* debug [class_header]: Header for VZSocketDevice */
// The class instance for the [VZSocketDevice] class.
var (
	VZSocketDeviceClass     _VZSocketDeviceClass
	VZSocketDeviceClassOnce sync.Once
)

func getVZSocketDeviceClass() _VZSocketDeviceClass {
	VZSocketDeviceClassOnce.Do(func() {
		VZSocketDeviceClass = _VZSocketDeviceClass{objc.GetClass("VZSocketDevice")}
	})
	return VZSocketDeviceClass
}

type _VZSocketDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZSocketDevice */
// An interface definition for the [VZSocketDevice] class.
type IVZSocketDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZSocketDevice */
	// properties:
	SocketDevices() IVZSocketDevice
	SetSocketDevices(value IVZSocketDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZSocketDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZSocketDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZSocketDeviceClass) Alloc() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZSocketDeviceClass) New() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSocketDevice) Init() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSocketDevice) Autorelease() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDevice creates a new VZSocketDevice instance.
func NewVZSocketDevice() VZSocketDevice {
	return getVZSocketDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZSocketDevice */
// The common behavior of socket devices.
//
// Don’t create or use a object directly. If your virtual machine’s configuration includes a object, the virtual machine returns a object in its property. Use that object to configure the port-based communications for your virtual machine.


// The common behavior of socket devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSocketDevice
type VZSocketDevice struct {
	objectivec.Object
}

// VZSocketDeviceFrom constructs a [VZSocketDevice] from an unsafe.Pointer.
//
// The common behavior of socket devices.
func VZSocketDeviceFrom(ptr unsafe.Pointer) VZSocketDevice {
	return VZSocketDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZSocketDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZSocketDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZSocketDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZSocketDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZSocketDevice */

// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZSocketDevice) SocketDevices() IVZSocketDevice {
	rv := objc.Send[VZSocketDevice](v_.ID, objc.Sel("socketDevices"))
	return rv
}/* debug [instance_properties/getter]: socketDevices */


// The array of socket devices that the VM configures for use ports in the guest VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (v_ VZSocketDevice) SetSocketDevices(value IVZSocketDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}/* debug [instance_properties/setter]: socketDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZSocketDevice */



