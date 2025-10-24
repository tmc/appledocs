// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZNetworkDeviceConfiguration */

/* debug [class_header]: Header for VZNetworkDeviceConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZNetworkDeviceConfiguration */
// An interface definition for the [VZNetworkDeviceConfiguration] class.
type IVZNetworkDeviceConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZNetworkDeviceConfiguration */
	// properties:
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	MACAddress() IVZMACAddress
	SetMACAddress(value IVZMACAddress)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZNetworkDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZNetworkDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceConfigurationClass) Alloc() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZNetworkDeviceConfiguration */
// The common configuration traits for network devices.
//
// Don’t instantiate the class directly. Instead, instantiate one of its subclasses, such as . Then use the properties of this class to configure the network device.

// The common configuration traits for network devices.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZNetworkDeviceConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZNetworkDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZNetworkDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZNetworkDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZNetworkDeviceConfiguration */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/attachment
func (v_ VZNetworkDeviceConfiguration) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
} /* debug [instance_properties/getter]: attachment */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/attachment
func (v_ VZNetworkDeviceConfiguration) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
} /* debug [instance_properties/setter]: attachment */

// The media access control (MAC) address to assign to the network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/macAddress
func (v_ VZNetworkDeviceConfiguration) MACAddress() IVZMACAddress {
	rv := objc.Send[VZMACAddress](v_.ID, objc.Sel("MACAddress"))
	return rv
} /* debug [instance_properties/getter]: MACAddress */

// The media access control (MAC) address to assign to the network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/macAddress
func (v_ VZNetworkDeviceConfiguration) SetMACAddress(value IVZMACAddress) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMACAddress:"), value)
} /* debug [instance_properties/setter]: MACAddress */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZNetworkDeviceConfiguration */
