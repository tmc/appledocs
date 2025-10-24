// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVirtioNetworkDeviceConfiguration */

/* debug [class_header]: Header for VZVirtioNetworkDeviceConfiguration */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioNetworkDeviceConfiguration */
// An interface definition for the [VZVirtioNetworkDeviceConfiguration] class.
type IVZVirtioNetworkDeviceConfiguration interface {
	IVZNetworkDeviceConfiguration

	/* debug [class_interface_properties]: Properties for VZVirtioNetworkDeviceConfiguration */
	// properties:
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioNetworkDeviceConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioNetworkDeviceConfiguration */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioNetworkDeviceConfigurationClass) Alloc() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioNetworkDeviceConfiguration */
// A configuration object that requests the creation of a network device for the guest system.
//
// Use a object to configure one network interface of your virtual machine. After creating this object, assign an appropriate value to its inherited property to define the type of network interface you want. You can also assign a specific MAC address, or let the system generate a random address for you. After creating and configuring a object, assign it to the property of your virtual machine’s configuration.

// A configuration object that requests the creation of a network device for the guest system.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioNetworkDeviceConfiguration */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioNetworkDeviceConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioNetworkDeviceConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioNetworkDeviceConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioNetworkDeviceConfiguration */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZVirtioNetworkDeviceConfiguration) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
} /* debug [instance_properties/getter]: attachment */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZVirtioNetworkDeviceConfiguration) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
} /* debug [instance_properties/setter]: attachment */

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZVirtioNetworkDeviceConfiguration) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
} /* debug [instance_properties/getter]: networkDevices */

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZVirtioNetworkDeviceConfiguration) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
} /* debug [instance_properties/setter]: networkDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioNetworkDeviceConfiguration */
