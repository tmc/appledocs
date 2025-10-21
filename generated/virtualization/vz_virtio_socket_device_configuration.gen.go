// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioSocketDeviceConfiguration] class.
var (
	VZVirtioSocketDeviceConfigurationClass     _VZVirtioSocketDeviceConfigurationClass
	VZVirtioSocketDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSocketDeviceConfigurationClass() _VZVirtioSocketDeviceConfigurationClass {
	VZVirtioSocketDeviceConfigurationClassOnce.Do(func() {
		VZVirtioSocketDeviceConfigurationClass = _VZVirtioSocketDeviceConfigurationClass{objc.GetClass("VZVirtioSocketDeviceConfiguration")}
	})
	return VZVirtioSocketDeviceConfigurationClass
}

type _VZVirtioSocketDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioSocketDeviceConfiguration] class.
type IVZVirtioSocketDeviceConfiguration interface {
	IVZSocketDeviceConfiguration
}

// A configuration object that requests the creation of a socket device to communicate with the guest system.
//
// Use a object to implement port-based communication between the guest operating system and the host computer. When you add this object to the property of your , the virtual machine provides a corresponding object to use to configure the ports. Add only one to your virtual machine’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type VZVirtioSocketDeviceConfiguration struct {
	VZSocketDeviceConfiguration
}

// VZVirtioSocketDeviceConfigurationFrom constructs a [VZVirtioSocketDeviceConfiguration] from an unsafe.Pointer.
//
// A configuration object that requests the creation of a socket device to communicate with the guest system.
func VZVirtioSocketDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioSocketDeviceConfiguration {
	return VZVirtioSocketDeviceConfiguration{
		VZSocketDeviceConfiguration: VZSocketDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketDeviceConfigurationClass) Alloc() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioSocketDeviceConfigurationClass) New() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSocketDeviceConfiguration) Init() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSocketDeviceConfiguration) Autorelease() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDeviceConfiguration creates a new VZVirtioSocketDeviceConfiguration instance.
func NewVZVirtioSocketDeviceConfiguration() VZVirtioSocketDeviceConfiguration {
	return getVZVirtioSocketDeviceConfigurationClass().New()
}



// The socket device that you use to implement port-based communication with the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/socketdevices
func (v_ VZVirtioSocketDeviceConfiguration) SocketDevices() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](v_.ID, objc.Sel("socketDevices"))
	return rv
}


// SetSocketDevices sets the value of the socketDevices property.
// The socket device that you use to implement port-based communication with the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/socketdevices
func (v_ VZVirtioSocketDeviceConfiguration) SetSocketDevices(value IVZSocketDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSocketDevices:"), value)
}


