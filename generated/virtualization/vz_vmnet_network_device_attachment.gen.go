// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVmnetNetworkDeviceAttachment] class.
var (
	VZVmnetNetworkDeviceAttachmentClass     _VZVmnetNetworkDeviceAttachmentClass
	VZVmnetNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZVmnetNetworkDeviceAttachmentClass() _VZVmnetNetworkDeviceAttachmentClass {
	VZVmnetNetworkDeviceAttachmentClassOnce.Do(func() {
		VZVmnetNetworkDeviceAttachmentClass = _VZVmnetNetworkDeviceAttachmentClass{objc.GetClass("VZVmnetNetworkDeviceAttachment")}
	})
	return VZVmnetNetworkDeviceAttachmentClass
}

type _VZVmnetNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZVmnetNetworkDeviceAttachment] class.
type IVZVmnetNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
	// properties:
	Network() unsafe.Pointer
	// methods:
}

// A network device attachment that allows a custom network topology.
//
// The Virtualization framework backs this attachment by a logical network which the client creates and customizes through the framework APIs to allow custom network topology which allows multiple virtual machines to appear on the same network and connect with each other.


// A network device attachment that allows a custom network topology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment
type VZVmnetNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZVmnetNetworkDeviceAttachmentFrom constructs a [VZVmnetNetworkDeviceAttachment] from an unsafe.Pointer.
//
// A network device attachment that allows a custom network topology.
func VZVmnetNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZVmnetNetworkDeviceAttachment {
	return VZVmnetNetworkDeviceAttachment{
		VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVmnetNetworkDeviceAttachmentClass) Alloc() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVmnetNetworkDeviceAttachmentClass) New() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVmnetNetworkDeviceAttachment) Init() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVmnetNetworkDeviceAttachment) Autorelease() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVmnetNetworkDeviceAttachment creates a new VZVmnetNetworkDeviceAttachment instance.
func NewVZVmnetNetworkDeviceAttachment() VZVmnetNetworkDeviceAttachment {
	return getVZVmnetNetworkDeviceAttachmentClass().New()
}



// Creates the attachment and configures it with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/init(network:)
func NewVZVmnetNetworkDeviceAttachmentWithNetwork(network unsafe.Pointer) VZVmnetNetworkDeviceAttachment {
	instance := getVZVmnetNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](instance.ID, objc.Sel("initWithNetwork:"), network)
	rv.Autorelease()
	return rv
}



// The network object that the you initialize the attachment with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/network
func (v_ VZVmnetNetworkDeviceAttachment) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("network"))
	return rv
}


