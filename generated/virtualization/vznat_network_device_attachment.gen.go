// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZNATNetworkDeviceAttachment */


/* debug [class_header]: Header for VZNATNetworkDeviceAttachment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZNATNetworkDeviceAttachment */
// An interface definition for the [VZNATNetworkDeviceAttachment] class.
type IVZNATNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	
/* debug [class_interface_properties]: Properties for VZNATNetworkDeviceAttachment */
	// properties:
	Attachment() IVZNetworkDeviceAttachment
	SetAttachment(value IVZNetworkDeviceAttachment)
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZNATNetworkDeviceAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZNATNetworkDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZNATNetworkDeviceAttachmentClass) Alloc() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZNATNetworkDeviceAttachment */
// A device that routes network requests through the host computer and performs network address translation on the resulting packets.
//
// A works with the host computer to perform network address translation (NAT) on the guest system’s network packets, and then route those packets to outside networks. Use this attachment to give the guest system indirect access to external networks, instead of direct access through a shared physical network interface. To configure a network device with a NAT attachment: Create the object. Assign the attachment object to the property of a object. Add the object to the property of your . This attachment doesn’t require your app to have the entitlement.


// A device that routes network requests through the host computer and performs network address translation on the resulting packets.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZNATNetworkDeviceAttachment */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZNATNetworkDeviceAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZNATNetworkDeviceAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZNATNetworkDeviceAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZNATNetworkDeviceAttachment */

// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZNATNetworkDeviceAttachment) Attachment() IVZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// The object that defines how the virtual network device communicates with the host system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vznetworkdeviceconfiguration/attachment
func (v_ VZNATNetworkDeviceAttachment) SetAttachment(value IVZNetworkDeviceAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZNATNetworkDeviceAttachment) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v_.ID, objc.Sel("networkDevices"))
	return rv
}/* debug [instance_properties/getter]: networkDevices */


// The array of network devices that you expose to the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v_ VZNATNetworkDeviceAttachment) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNetworkDevices:"), value)
}/* debug [instance_properties/setter]: networkDevices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZNATNetworkDeviceAttachment */


