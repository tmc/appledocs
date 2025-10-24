// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZBridgedNetworkDeviceAttachment */


/* debug [class_header]: Header for VZBridgedNetworkDeviceAttachment */
// The class instance for the [VZBridgedNetworkDeviceAttachment] class.
var (
	VZBridgedNetworkDeviceAttachmentClass     _VZBridgedNetworkDeviceAttachmentClass
	VZBridgedNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZBridgedNetworkDeviceAttachmentClass() _VZBridgedNetworkDeviceAttachmentClass {
	VZBridgedNetworkDeviceAttachmentClassOnce.Do(func() {
		VZBridgedNetworkDeviceAttachmentClass = _VZBridgedNetworkDeviceAttachmentClass{objc.GetClass("VZBridgedNetworkDeviceAttachment")}
	})
	return VZBridgedNetworkDeviceAttachmentClass
}

type _VZBridgedNetworkDeviceAttachmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZBridgedNetworkDeviceAttachment */
// An interface definition for the [VZBridgedNetworkDeviceAttachment] class.
type IVZBridgedNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	
/* debug [class_interface_properties]: Properties for VZBridgedNetworkDeviceAttachment */
	// properties:
	Interface() IVZBridgedNetworkInterface
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZBridgedNetworkDeviceAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZBridgedNetworkDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZBridgedNetworkDeviceAttachmentClass) Alloc() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZBridgedNetworkDeviceAttachmentClass) New() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZBridgedNetworkDeviceAttachment) Init() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZBridgedNetworkDeviceAttachment) Autorelease() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkDeviceAttachment creates a new VZBridgedNetworkDeviceAttachment instance.
func NewVZBridgedNetworkDeviceAttachment() VZBridgedNetworkDeviceAttachment {
	return getVZBridgedNetworkDeviceAttachmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZBridgedNetworkDeviceAttachment */
// A network device that interacts directly with a physical network interface on the host computer.
//
// A object represents a physical interface on the host computer. Use this object when configuring a network interface for your virtual machine. A bridged network device sends and receives packets on the same physical interface as the host computer, but does so using a different network layer. To configure a network device with a bridged network interface: Obtain a reference to one of the host’s physical network interfaces from the property of . Create the object using the network interface. Assign the attachment object to the property of a object. Add the object to the property of your .


// A network device that interacts directly with a physical network interface on the host computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment
type VZBridgedNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZBridgedNetworkDeviceAttachmentFrom constructs a [VZBridgedNetworkDeviceAttachment] from an unsafe.Pointer.
//
// A network device that interacts directly with a physical network interface on the host computer.
func VZBridgedNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZBridgedNetworkDeviceAttachment {
	return VZBridgedNetworkDeviceAttachment{
		VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZBridgedNetworkDeviceAttachment */

// Creates the attachment from a bridged network interface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/init(interface:)
func NewVZBridgedNetworkDeviceAttachmentWithInterface(interface_ IVZBridgedNetworkInterface) VZBridgedNetworkDeviceAttachment {
	instance := getVZBridgedNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](instance.ID, objc.Sel("initWithInterface:"), interface_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZBridgedNetworkDeviceAttachmentWithInterface */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZBridgedNetworkDeviceAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZBridgedNetworkDeviceAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZBridgedNetworkDeviceAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZBridgedNetworkDeviceAttachment */

// The network interface assigned to this attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/interface
func (v_ VZBridgedNetworkDeviceAttachment) Interface() IVZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](v_.ID, objc.Sel("interface"))
	return rv
}/* debug [instance_properties/getter]: interface */


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZBridgedNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZBridgedNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZBridgedNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}/* debug [instance_properties/getter]: networkDevices */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZBridgedNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}/* debug [instance_properties/setter]: networkDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZBridgedNetworkDeviceAttachment */


