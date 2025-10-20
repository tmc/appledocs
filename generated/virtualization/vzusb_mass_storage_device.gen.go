// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZUSBMassStorageDevice] class.
type IVZUSBMassStorageDevice interface {
	IVZStorageDevice
}

// A class that represents a hot-pluggable USB mass storage device.
//
// Create this device either by instantiating it directly and passing to its initializer, or instantiating a in a . Direct instantiation creates an object that you can pass to . Instantiation through makes the device available in the property.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZUSBMassStorageDeviceClass) Alloc() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




