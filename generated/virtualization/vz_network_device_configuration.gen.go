// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZNetworkDeviceConfiguration] class.
var (
	VZNetworkDeviceConfigurationClass     _VZNetworkDeviceConfigurationClass
	VZNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZNetworkDeviceConfigurationClass() _VZNetworkDeviceConfigurationClass {
	VZNetworkDeviceConfigurationClassOnce.Do(func() {
		VZNetworkDeviceConfigurationClass = _VZNetworkDeviceConfigurationClass{objc.GetClass("VZNetworkDeviceConfiguration")}
	})
	return VZNetworkDeviceConfigurationClass
}

type _VZNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZNetworkDeviceConfiguration] class.
type IVZNetworkDeviceConfiguration interface {
	objectivec.IObject
}

// The common configuration traits for network devices.
//
// Don’t instantiate the class directly. Instead, instantiate one of its subclasses, such as . Then use the properties of this class to configure the network device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration
type VZNetworkDeviceConfiguration struct {
	objectivec.Object
}

// VZNetworkDeviceConfigurationFrom constructs a [VZNetworkDeviceConfiguration] from an unsafe.Pointer.
//
// The common configuration traits for network devices.
func VZNetworkDeviceConfigurationFrom(ptr unsafe.Pointer) VZNetworkDeviceConfiguration {
	return VZNetworkDeviceConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceConfigurationClass) Alloc() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZNetworkDeviceConfigurationClass) New() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNetworkDeviceConfiguration) Init() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNetworkDeviceConfiguration) Autorelease() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceConfiguration creates a new VZNetworkDeviceConfiguration instance.
func NewVZNetworkDeviceConfiguration() VZNetworkDeviceConfiguration {
	return getVZNetworkDeviceConfigurationClass().New()
}


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/attachment
func (v_ VZNetworkDeviceConfiguration) Attachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The object that defines how the virtual network device communicates with the host system.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/attachment
func (v_ VZNetworkDeviceConfiguration) SetAttachment(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}

// The media access control (MAC) address to assign to the network device.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/macAddress
func (v_ VZNetworkDeviceConfiguration) MACAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("MACAddress"))
	return rv
}


// SetMACAddress sets the value of the MACAddress property.
// The media access control (MAC) address to assign to the network device.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/macAddress
func (v_ VZNetworkDeviceConfiguration) SetMACAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMACAddress:"), value)
}



