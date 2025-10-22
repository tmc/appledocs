// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZNATNetworkDeviceAttachment] class.
var (
	VZNATNetworkDeviceAttachmentClass     _VZNATNetworkDeviceAttachmentClass
	VZNATNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNATNetworkDeviceAttachmentClass() _VZNATNetworkDeviceAttachmentClass {
	VZNATNetworkDeviceAttachmentClassOnce.Do(func() {
		VZNATNetworkDeviceAttachmentClass = _VZNATNetworkDeviceAttachmentClass{objc.GetClass("VZNATNetworkDeviceAttachment")}
	})
	return VZNATNetworkDeviceAttachmentClass
}

type _VZNATNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZNATNetworkDeviceAttachment] class.
type IVZNATNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	Attachment() VZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() VZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}

// A device that routes network requests through the host computer and performs network address translation on the resulting packets.
//
// A works with the host computer to perform network address translation (NAT) on the guest system’s network packets, and then route those packets to outside networks. Use this attachment to give the guest system indirect access to external networks, instead of direct access through a shared physical network interface. To configure a network device with a NAT attachment: Create the object. Assign the attachment object to the property of a object. Add the object to the property of your . This attachment doesn’t require your app to have the entitlement.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment
type VZNATNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZNATNetworkDeviceAttachmentFrom constructs a [VZNATNetworkDeviceAttachment] from an unsafe.Pointer.
//
// A device that routes network requests through the host computer and performs network address translation on the resulting packets.
func VZNATNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZNATNetworkDeviceAttachment {
	return VZNATNetworkDeviceAttachment{
		VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNATNetworkDeviceAttachmentClass) Alloc() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZNATNetworkDeviceAttachmentClass) New() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNATNetworkDeviceAttachment) Init() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNATNetworkDeviceAttachment) Autorelease() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNATNetworkDeviceAttachment creates a new VZNATNetworkDeviceAttachment instance.
func NewVZNATNetworkDeviceAttachment() VZNATNetworkDeviceAttachment {
	return getVZNATNetworkDeviceAttachmentClass().New()
}


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZNATNetworkDeviceAttachment) Attachment() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The object that defines how the virtual network device communicates with the host system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZNATNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZNATNetworkDeviceAttachment) NetworkDevices() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// SetNetworkDevices sets the value of the networkDevices property.
// The array of network devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZNATNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}



