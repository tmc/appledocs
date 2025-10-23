// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [USBHostDevice] class.
var (
	USBHostDeviceClass     _USBHostDeviceClass
	USBHostDeviceClassOnce sync.Once
)

func getUSBHostDeviceClass() _USBHostDeviceClass {
	USBHostDeviceClassOnce.Do(func() {
		USBHostDeviceClass = _USBHostDeviceClass{objc.GetClass("IOUSBHostDevice")}
	})
	return USBHostDeviceClass
}

type _USBHostDeviceClass struct {
	class objc.Class
}

// An interface definition for the [USBHostDevice] class.
type IUSBHostDevice interface {
	IUSBHostObject
	// properties:
	ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */
	SetConfigurationDescriptor(value USBConfigurationDescriptor /* not a class type */)
	// methods:
	ResetWithError(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
}

// The class that claims and configures devices, retrieves descriptors, and sends device requests.
//
// This class enables management of the device state, including sending control requests to the default endpoint 0, configuring the device, and resetting the device. The interest handler also allows monitoring of the device state. The client creates the class and initializes it with .


// The class that claims and configures devices, retrieves descriptors, and sends device requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice
type USBHostDevice struct {
	USBHostObject
}

// USBHostDeviceFrom constructs a [USBHostDevice] from an unsafe.Pointer.
//
// The class that claims and configures devices, retrieves descriptors, and sends device requests.
func USBHostDeviceFrom(ptr unsafe.Pointer) USBHostDevice {
	return USBHostDevice{
		USBHostObject: USBHostObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostDeviceClass) Alloc() USBHostDevice {
	rv := objc.Send[USBHostDevice](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _USBHostDeviceClass) New() USBHostDevice {
	rv := objc.Send[USBHostDevice](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostDevice) Init() USBHostDevice {
	rv := objc.Send[USBHostDevice](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostDevice) Autorelease() USBHostDevice {
	rv := objc.Send[USBHostDevice](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostDevice creates a new USBHostDevice instance.
func NewUSBHostDevice() USBHostDevice {
	return getUSBHostDeviceClass().New()
}



// Terminates the device and attempts to re-enumerate it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDevice/reset()
func (u_ USBHostDevice) ResetWithError(error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("resetWithError:"), error_)
	return rv
}


// The currently selected configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdevice/configurationdescriptor
func (u_ USBHostDevice) ConfigurationDescriptor() USBConfigurationDescriptor /* not a class type */ {
	rv := objc.Send[USBConfigurationDescriptor](u_.ID, objc.Sel("configurationDescriptor"))
	return rv
}


// The currently selected configuration descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdevice/configurationdescriptor
func (u_ USBHostDevice) SetConfigurationDescriptor(value USBConfigurationDescriptor /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfigurationDescriptor:"), value)
}



