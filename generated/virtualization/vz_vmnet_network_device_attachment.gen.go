// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZVmnetNetworkDeviceAttachment */

/* debug [class_header]: Header for VZVmnetNetworkDeviceAttachment */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVmnetNetworkDeviceAttachment */
// An interface definition for the [VZVmnetNetworkDeviceAttachment] class.
type IVZVmnetNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	/* debug [class_interface_properties]: Properties for VZVmnetNetworkDeviceAttachment */
	// properties:
	Network() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVmnetNetworkDeviceAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVmnetNetworkDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZVmnetNetworkDeviceAttachmentClass) Alloc() VZVmnetNetworkDeviceAttachment {
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVmnetNetworkDeviceAttachment */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVmnetNetworkDeviceAttachment */

// Creates the attachment and configures it with the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/init(network:)
func NewVZVmnetNetworkDeviceAttachmentWithNetwork(network unsafe.Pointer) VZVmnetNetworkDeviceAttachment {
	instance := getVZVmnetNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZVmnetNetworkDeviceAttachment](instance.ID, objc.Sel("initWithNetwork:"), network)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZVmnetNetworkDeviceAttachmentWithNetwork */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVmnetNetworkDeviceAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVmnetNetworkDeviceAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVmnetNetworkDeviceAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVmnetNetworkDeviceAttachment */

// The network object that the you initialize the attachment with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVmnetNetworkDeviceAttachment/network
func (v_ VZVmnetNetworkDeviceAttachment) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("network"))
	return rv
} /* debug [instance_properties/getter]: network */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVmnetNetworkDeviceAttachment */
