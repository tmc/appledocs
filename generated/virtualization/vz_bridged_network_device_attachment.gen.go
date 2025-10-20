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
}

// A network device that interacts directly with a physical network interface on the host computer.
//
// A object represents a physical interface on the host computer. Use this object when configuring a network interface for your virtual machine. A bridged network device sends and receives packets on the same physical interface as the host computer, but does so using a different network layer. To configure a network device with a bridged network interface: Obtain a reference to one of the host’s physical network interfaces from the property of . Create the object using the network interface. Assign the attachment object to the property of a object. Add the object to the property of your .
//
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




