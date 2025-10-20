// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioNetworkDeviceConfiguration] class.
var (
	VZVirtioNetworkDeviceConfigurationClass     _VZVirtioNetworkDeviceConfigurationClass
	VZVirtioNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioNetworkDeviceConfigurationClass() _VZVirtioNetworkDeviceConfigurationClass {
	VZVirtioNetworkDeviceConfigurationClassOnce.Do(func() {
		VZVirtioNetworkDeviceConfigurationClass = _VZVirtioNetworkDeviceConfigurationClass{objc.GetClass("VZVirtioNetworkDeviceConfiguration")}
	})
	return VZVirtioNetworkDeviceConfigurationClass
}

type _VZVirtioNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioNetworkDeviceConfiguration] class.
type IVZVirtioNetworkDeviceConfiguration interface {
	IVZNetworkDeviceConfiguration
}

// A configuration object that requests the creation of a network device for the guest system.
//
// Use a object to configure one network interface of your virtual machine. After creating this object, assign an appropriate value to its inherited property to define the type of network interface you want. You can also assign a specific MAC address, or let the system generate a random address for you. After creating and configuring a object, assign it to the property of your virtual machine’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioNetworkDeviceConfiguration
type VZVirtioNetworkDeviceConfiguration struct {
	VZNetworkDeviceConfiguration
}

// VZVirtioNetworkDeviceConfigurationFrom constructs a [VZVirtioNetworkDeviceConfiguration] from an unsafe.Pointer.
//
// A configuration object that requests the creation of a network device for the guest system.
func VZVirtioNetworkDeviceConfigurationFrom(ptr unsafe.Pointer) VZVirtioNetworkDeviceConfiguration {
	return VZVirtioNetworkDeviceConfiguration{
		VZNetworkDeviceConfiguration: VZNetworkDeviceConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioNetworkDeviceConfigurationClass) Alloc() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioNetworkDeviceConfigurationClass) New() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioNetworkDeviceConfiguration) Init() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioNetworkDeviceConfiguration) Autorelease() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioNetworkDeviceConfiguration creates a new VZVirtioNetworkDeviceConfiguration instance.
func NewVZVirtioNetworkDeviceConfiguration() VZVirtioNetworkDeviceConfiguration {
	return getVZVirtioNetworkDeviceConfigurationClass().New()
}




