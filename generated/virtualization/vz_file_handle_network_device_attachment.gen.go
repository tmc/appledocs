// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZFileHandleNetworkDeviceAttachment */

/* debug [class_header]: Header for VZFileHandleNetworkDeviceAttachment */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZFileHandleNetworkDeviceAttachment */
// An interface definition for the [VZFileHandleNetworkDeviceAttachment] class.
type IVZFileHandleNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	/* debug [class_interface_properties]: Properties for VZFileHandleNetworkDeviceAttachment */
	// properties:
	FileHandle() foundation.FileHandle
	MaximumTransmissionUnit() int
	SetMaximumTransmissionUnit(value int)
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZFileHandleNetworkDeviceAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZFileHandleNetworkDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZFileHandleNetworkDeviceAttachmentClass) Alloc() VZFileHandleNetworkDeviceAttachment {
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZFileHandleNetworkDeviceAttachment */
// A network device that transmits raw network packets and frames using a datagram socket.
//
// A object maps a network interface to a connected datagram socket. This attachment transmits data at the data link layer. You configure and manage the socket in your app, and manage the corresponding data transfers. To configure a network device with a socket-based file handle: Create a socket with the type in your app. Create a from the socket’s file descriptor. Create the object using the file handle. Assign the attachment object to the property of a object. Add the object to the property of your . This attachment doesn’t require your app to have the entitlement.

// A network device that transmits raw network packets and frames using a datagram socket.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZFileHandleNetworkDeviceAttachment */

// Creates the attachment from a file handle that contains a connected datagram socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/init(fileHandle:)
func NewVZFileHandleNetworkDeviceAttachmentWithFileHandle(fileHandle foundation.FileHandle) VZFileHandleNetworkDeviceAttachment {
	instance := getVZFileHandleNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZFileHandleNetworkDeviceAttachment](instance.ID, objc.Sel("initWithFileHandle:"), fileHandle)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZFileHandleNetworkDeviceAttachmentWithFileHandle */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZFileHandleNetworkDeviceAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZFileHandleNetworkDeviceAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZFileHandleNetworkDeviceAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZFileHandleNetworkDeviceAttachment */

// The file handle assigned to this attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/fileHandle
func (v_ VZFileHandleNetworkDeviceAttachment) FileHandle() foundation.FileHandle {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandle"))
	return rv
} /* debug [instance_properties/getter]: fileHandle */

// An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/maximumTransmissionUnit
func (v_ VZFileHandleNetworkDeviceAttachment) MaximumTransmissionUnit() int {
	rv := objc.Send[int](v_.ID, objc.Sel("maximumTransmissionUnit"))
	return rv
} /* debug [instance_properties/getter]: maximumTransmissionUnit */

// An integer value that indicates the maximum transmission unit (MTU) associated with this attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleNetworkDeviceAttachment/maximumTransmissionUnit
func (v_ VZFileHandleNetworkDeviceAttachment) SetMaximumTransmissionUnit(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumTransmissionUnit:"), value)
} /* debug [instance_properties/setter]: maximumTransmissionUnit */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZFileHandleNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
} /* debug [instance_properties/getter]: attachment */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZFileHandleNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
} /* debug [instance_properties/setter]: attachment */

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZFileHandleNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
} /* debug [instance_properties/getter]: networkDevices */

// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZFileHandleNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
} /* debug [instance_properties/setter]: networkDevices */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZFileHandleNetworkDeviceAttachment */
