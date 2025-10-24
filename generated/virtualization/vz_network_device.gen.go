// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZNetworkDevice */


/* debug [class_header]: Header for VZNetworkDevice */
// The class instance for the [VZNetworkDevice] class.
var (
	VZNetworkDeviceClass     _VZNetworkDeviceClass
	VZNetworkDeviceClassOnce sync.Once
)

func getVZNetworkDeviceClass() _VZNetworkDeviceClass {
	VZNetworkDeviceClassOnce.Do(func() {
		VZNetworkDeviceClass = _VZNetworkDeviceClass{objc.GetClass("VZNetworkDevice")}
	})
	return VZNetworkDeviceClass
}

type _VZNetworkDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZNetworkDevice */
// An interface definition for the [VZNetworkDevice] class.
type IVZNetworkDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZNetworkDevice */
	// properties:
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDevice
	SetNetworkDevices(value IVZNetworkDevice)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZNetworkDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZNetworkDevice */
// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceClass) Alloc() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZNetworkDeviceClass) New() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNetworkDevice) Init() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNetworkDevice) Autorelease() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDevice creates a new VZNetworkDevice instance.
func NewVZNetworkDevice() VZNetworkDevice {
	return getVZNetworkDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZNetworkDevice */
// A base class that represents a network device in a virtual machine.
//
// Don’t instantiate a   directly. When you create a   instance with a   the system creates the number of network devices based on the number of   objects you specify in the VM configuration. Before initializing the virtual machine (VM), validate the configuration using   to ensure the user’s computer supports the number of network and other devices you’ve specified.   For many purposes, a single network that uses a Network Address Translation (NAT) attachment and connects the VM to the host computer’s network is sufficient. You can use additional network interfaces for purposes of your own design, such as: Bridging several physical interfaces to connect to multiple networks. Using the file descriptor attachment to create specialized connections for different purposes. You access the network devices through the  . property. The network devices map to their respective configurations in a one to one relationship, where index   of   corresponds to the network device configuration at index   set on  . .


// A base class that represents a network device in a virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type VZNetworkDevice struct {
	objectivec.Object
}

// VZNetworkDeviceFrom constructs a [VZNetworkDevice] from an unsafe.Pointer.
//
// A base class that represents a network device in a virtual machine.
func VZNetworkDeviceFrom(ptr unsafe.Pointer) VZNetworkDevice {
	return VZNetworkDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZNetworkDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZNetworkDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZNetworkDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZNetworkDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZNetworkDevice */

// The network attachment that’s connected to this network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice/attachment
func (v_ VZNetworkDevice) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The network attachment that’s connected to this network device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice/attachment
func (v_ VZNetworkDevice) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */


// The list of configured network devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/networkdevices
func (v_ VZNetworkDevice) NetworkDevices() IVZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("networkDevices"))
	return rv
}/* debug [instance_properties/getter]: networkDevices */


// The list of configured network devices on the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/networkdevices
func (v_ VZNetworkDevice) SetNetworkDevices(value IVZNetworkDevice) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}/* debug [instance_properties/setter]: networkDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZNetworkDevice */



