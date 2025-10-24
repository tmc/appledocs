// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZUSBMassStorageDevice */

/* debug [class_header]: Header for VZUSBMassStorageDevice */
// The class instance for the [VZUSBMassStorageDevice] class.
var (
	VZUSBMassStorageDeviceClass     _VZUSBMassStorageDeviceClass
	VZUSBMassStorageDeviceClassOnce sync.Once
)

func getVZUSBMassStorageDeviceClass() _VZUSBMassStorageDeviceClass {
	VZUSBMassStorageDeviceClassOnce.Do(func() {
		VZUSBMassStorageDeviceClass = _VZUSBMassStorageDeviceClass{objc.GetClass("VZUSBMassStorageDevice")}
	})
	return VZUSBMassStorageDeviceClass
}

type _VZUSBMassStorageDeviceClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZUSBMassStorageDevice */
// An interface definition for the [VZUSBMassStorageDevice] class.
type IVZUSBMassStorageDevice interface {
	IVZStorageDevice

	/* debug [class_interface_properties]: Properties for VZUSBMassStorageDevice */
	// properties:
	UsbDevices() unsafe.Pointer
	SetUsbDevices(value unsafe.Pointer)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZUSBMassStorageDevice */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZUSBMassStorageDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZUSBMassStorageDeviceClass) Alloc() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZUSBMassStorageDeviceClass) New() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZUSBMassStorageDevice) Init() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZUSBMassStorageDevice) Autorelease() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDevice creates a new VZUSBMassStorageDevice instance.
func NewVZUSBMassStorageDevice() VZUSBMassStorageDevice {
	return getVZUSBMassStorageDeviceClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZUSBMassStorageDevice */
// A class that represents a hot-pluggable USB mass storage device.
//
// Create this device either by instantiating it directly and passing to its initializer, or instantiating a in a . Direct instantiation creates an object that you can pass to . Instantiation through makes the device available in the property.

// A class that represents a hot-pluggable USB mass storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice
type VZUSBMassStorageDevice struct {
	VZStorageDevice
}

// VZUSBMassStorageDeviceFrom constructs a [VZUSBMassStorageDevice] from an unsafe.Pointer.
//
// A class that represents a hot-pluggable USB mass storage device.
func VZUSBMassStorageDeviceFrom(ptr unsafe.Pointer) VZUSBMassStorageDevice {
	return VZUSBMassStorageDevice{
		VZStorageDevice: VZStorageDeviceFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZUSBMassStorageDevice */

// Creates a USB mass storage device with the provided configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/init(configuration:)
func NewVZUSBMassStorageDeviceWithConfiguration(configuration IVZUSBMassStorageDeviceConfiguration) VZUSBMassStorageDevice {
	instance := getVZUSBMassStorageDeviceClass().Alloc()
	rv := objc.Send[VZUSBMassStorageDevice](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZUSBMassStorageDeviceWithConfiguration */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZUSBMassStorageDevice */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZUSBMassStorageDevice */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZUSBMassStorageDevice */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZUSBMassStorageDevice */

// The list of attached USB devices for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzusbcontroller/usbdevices
func (v_ VZUSBMassStorageDevice) UsbDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("usbDevices"))
	return rv
} /* debug [instance_properties/getter]: usbDevices */

// The list of attached USB devices for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzusbcontroller/usbdevices
func (v_ VZUSBMassStorageDevice) SetUsbDevices(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUsbDevices:"), value)
} /* debug [instance_properties/setter]: usbDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZUSBMassStorageDevice */
