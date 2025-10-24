// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioConsoleDeviceConfiguration] class.
var (
	VZVirtioConsoleDeviceConfigurationClass     _VZVirtioConsoleDeviceConfigurationClass
	VZVirtioConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceConfigurationClass() _VZVirtioConsoleDeviceConfigurationClass {
	VZVirtioConsoleDeviceConfigurationClassOnce.Do(func() {
		VZVirtioConsoleDeviceConfigurationClass = _VZVirtioConsoleDeviceConfigurationClass{objc.GetClass("VZVirtioConsoleDeviceConfiguration")}
	})
	return VZVirtioConsoleDeviceConfigurationClass
}

type _VZVirtioConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsoleDeviceConfiguration] class.
type IVZVirtioConsoleDeviceConfiguration interface {
	IVZConsoleDeviceConfiguration
	// properties:
	Ports() IVZVirtioConsolePortConfigurationArray
	ConsoleDevices() IVZConsoleDeviceConfiguration
	SetConsoleDevices(value IVZConsoleDeviceConfiguration)
	// methods:
}

// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
//
// A object enables serial communication between the guest-operating system and host computer through the Virtio interface. The device sets up one or more ports through on the Virtio console device.


// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type VZVirtioConsoleDeviceConfiguration struct {
	VZConsoleDeviceConfiguration
}

// VZVirtioConsoleDeviceConfigurationFrom constructs a [VZVirtioConsoleDeviceConfiguration] from an unsafe.Pointer.
//
// A console device that enables communication between the host and the guest using console ports through a Virtio interface.
func VZVirtioConsoleDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsoleDeviceConfiguration {
	return VZVirtioConsoleDeviceConfiguration{
		VZConsoleDeviceConfiguration: VZConsoleDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceConfigurationClass) Alloc() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsoleDeviceConfigurationClass) New() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsoleDeviceConfiguration) Init() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsoleDeviceConfiguration) Autorelease() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceConfiguration creates a new VZVirtioConsoleDeviceConfiguration instance.
func NewVZVirtioConsoleDeviceConfiguration() VZVirtioConsoleDeviceConfiguration {
	return getVZVirtioConsoleDeviceConfigurationClass().New()
}




// The list of Virtio port configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration/ports
func (v_ VZVirtioConsoleDeviceConfiguration) Ports() IVZVirtioConsolePortConfigurationArray {
	rv := objc.Send[VZVirtioConsolePortConfigurationArray](v_.ID, objc.Sel("ports"))
	return rv
}


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsoleDeviceConfiguration) ConsoleDevices() IVZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](v_.ID, objc.Sel("consoleDevices"))
	return rv
}


// The array of console devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v_ VZVirtioConsoleDeviceConfiguration) SetConsoleDevices(value IVZConsoleDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setConsoleDevices:"), value)
}


