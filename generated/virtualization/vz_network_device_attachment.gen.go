// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZNetworkDeviceAttachment */


/* debug [class_header]: Header for VZNetworkDeviceAttachment */
// The class instance for the [VZNetworkDeviceAttachment] class.
var (
	VZNetworkDeviceAttachmentClass     _VZNetworkDeviceAttachmentClass
	VZNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkDeviceAttachmentClass() _VZNetworkDeviceAttachmentClass {
	VZNetworkDeviceAttachmentClassOnce.Do(func() {
		VZNetworkDeviceAttachmentClass = _VZNetworkDeviceAttachmentClass{objc.GetClass("VZNetworkDeviceAttachment")}
	})
	return VZNetworkDeviceAttachmentClass
}

type _VZNetworkDeviceAttachmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZNetworkDeviceAttachment */
// An interface definition for the [VZNetworkDeviceAttachment] class.
type IVZNetworkDeviceAttachment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZNetworkDeviceAttachment */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZNetworkDeviceAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZNetworkDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceAttachmentClass) Alloc() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZNetworkDeviceAttachmentClass) New() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNetworkDeviceAttachment) Init() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNetworkDeviceAttachment) Autorelease() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceAttachment creates a new VZNetworkDeviceAttachment instance.
func NewVZNetworkDeviceAttachment() VZNetworkDeviceAttachment {
	return getVZNetworkDeviceAttachmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZNetworkDeviceAttachment */
// The common behaviors for the network attachment points of your virtual machine.
//
// Don’t create a object directly. Instead, instantiate one of its concrete subclasses and use that object to configure your network devices. Each concrete subclass represents a specific type of network interface on the host computer.


// The common behaviors for the network attachment points of your virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment
type VZNetworkDeviceAttachment struct {
	objectivec.Object
}

// VZNetworkDeviceAttachmentFrom constructs a [VZNetworkDeviceAttachment] from an unsafe.Pointer.
//
// The common behaviors for the network attachment points of your virtual machine.
func VZNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZNetworkDeviceAttachment {
	return VZNetworkDeviceAttachment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZNetworkDeviceAttachment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZNetworkDeviceAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZNetworkDeviceAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZNetworkDeviceAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZNetworkDeviceAttachment */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZNetworkDeviceAttachment */



