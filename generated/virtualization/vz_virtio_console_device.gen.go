// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZVirtioConsoleDevice] class.
type IVZVirtioConsoleDevice interface {
	IVZConsoleDevice
}

// A class that represents a Virtio console device in a virtual machine.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceClass) Alloc() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The delegate object for the console device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/delegate
func (v_ VZVirtioConsoleDevice) Delegate() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the console device.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/delegate
func (v_ VZVirtioConsoleDevice) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}
// The array of console ports that a specific device uses.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/ports
func (v_ VZVirtioConsoleDevice) Ports() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("ports"))
	return rv
}



