// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZFileHandleNetworkDeviceAttachment] class.
var (
	VZFileHandleNetworkDeviceAttachmentClass     _VZFileHandleNetworkDeviceAttachmentClass
	VZFileHandleNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZFileHandleNetworkDeviceAttachmentClass() _VZFileHandleNetworkDeviceAttachmentClass {
	VZFileHandleNetworkDeviceAttachmentClassOnce.Do(func() {
		VZFileHandleNetworkDeviceAttachmentClass = _VZFileHandleNetworkDeviceAttachmentClass{objc.GetClass("VZFileHandleNetworkDeviceAttachment")}
	})
	return VZFileHandleNetworkDeviceAttachmentClass
}

type _VZFileHandleNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZFileHandleNetworkDeviceAttachment] class.
type IVZFileHandleNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	MaximumTransmissionUnit() int
	SetMaximumTransmissionUnit(value int)
	FileHandle() foundation.FileHandle
	SetFileHandle(value foundation.IFileHandle)
	Attachment() VZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() VZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}

// A network device that transmits raw network packets and frames using a datagram socket.
//
// A object maps a network interface to a connected datagram socket. This attachment transmits data at the data link layer. You configure and manage the socket in your app, and manage the corresponding data transfers. To configure a network device with a socket-based file handle: Create a socket with the type in your app. Create a from the socket’s file descriptor. Create the object using the file handle. Assign the attachment object to the property of a object. Add the object to the property of your . This attachment doesn’t require your app to have the entitlement.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment
type VZFileHandleNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZFileHandleNetworkDeviceAttachmentFrom constructs a [VZFileHandleNetworkDeviceAttachment] from an unsafe.Pointer.
//
// A network device that transmits raw network packets and frames using a datagram socket.
func VZFileHandleNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZFileHandleNetworkDeviceAttachment {
	return VZFileHandleNetworkDeviceAttachment{
		VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZFileHandleNetworkDeviceAttachmentClass) Alloc() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZFileHandleNetworkDeviceAttachmentClass) New() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZFileHandleNetworkDeviceAttachment) Init() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZFileHandleNetworkDeviceAttachment) Autorelease() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleNetworkDeviceAttachment creates a new VZFileHandleNetworkDeviceAttachment instance.
func NewVZFileHandleNetworkDeviceAttachment() VZFileHandleNetworkDeviceAttachment {
	return getVZFileHandleNetworkDeviceAttachmentClass().New()
}




// Creates the attachment from a file handle that contains a connected datagram socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/init(fileHandle:)
func NewVZFileHandleNetworkDeviceAttachmentWithFileHandle(fileHandle foundation.IFileHandle) VZFileHandleNetworkDeviceAttachment {
	instance := getVZFileHandleNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](instance.ID, objc.Sel("initWithFileHandle:"), fileHandle)
	rv.Autorelease()
	return rv
}


// An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/maximumTransmissionUnit
func (v_ VZFileHandleNetworkDeviceAttachment) MaximumTransmissionUnit() int {
	rv := objc.Send[int](v_.ID, objc.Sel("maximumTransmissionUnit"))
	return rv
}


// SetMaximumTransmissionUnit sets the value of the maximumTransmissionUnit property.
// An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/maximumTransmissionUnit
func (v_ VZFileHandleNetworkDeviceAttachment) SetMaximumTransmissionUnit(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumTransmissionUnit:"), value)
}

// The file handle assigned to this attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfilehandlenetworkdeviceattachment/filehandle
func (v_ VZFileHandleNetworkDeviceAttachment) FileHandle() foundation.FileHandle {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandle"))
	return rv
}


// SetFileHandle sets the value of the fileHandle property.
// The file handle assigned to this attachment.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfilehandlenetworkdeviceattachment/filehandle
func (v_ VZFileHandleNetworkDeviceAttachment) SetFileHandle(value foundation.IFileHandle) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFileHandle:"), value)
}

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZFileHandleNetworkDeviceAttachment) Attachment() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The object that defines how the virtual network device communicates with the host system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZFileHandleNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZFileHandleNetworkDeviceAttachment) NetworkDevices() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}


// SetNetworkDevices sets the value of the networkDevices property.
// The array of network devices that you expose to the guest operating system.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZFileHandleNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}


