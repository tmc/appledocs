// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZBridgedNetworkDeviceAttachment] class.
type IVZBridgedNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	// properties:
	Interface() IVZBridgedNetworkInterface
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZBridgedNetworkDeviceAttachmentClass) Alloc() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates the attachment from a bridged network interface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/init(interface:)
func NewVZBridgedNetworkDeviceAttachmentWithInterface(interface_ IVZBridgedNetworkInterface) VZBridgedNetworkDeviceAttachment {
	instance := getVZBridgedNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](instance.ID, objc.Sel("initWithInterface:"), interface_)
	rv.Autorelease()
	return rv
}



// The network interface assigned to this attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/interface
func (v_ VZBridgedNetworkDeviceAttachment) Interface() IVZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](v_.ID, objc.Sel("interface"))
	return rv
}


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZBridgedNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZBridgedNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZBridgedNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZBridgedNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}


