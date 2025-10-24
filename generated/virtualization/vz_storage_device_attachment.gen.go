// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZStorageDeviceAttachment */

/* debug [class_header]: Header for VZStorageDeviceAttachment */
// The class instance for the [VZStorageDeviceAttachment] class.
var (
	VZStorageDeviceAttachmentClass     _VZStorageDeviceAttachmentClass
	VZStorageDeviceAttachmentClassOnce sync.Once
)

func getVZStorageDeviceAttachmentClass() _VZStorageDeviceAttachmentClass {
	VZStorageDeviceAttachmentClassOnce.Do(func() {
		VZStorageDeviceAttachmentClass = _VZStorageDeviceAttachmentClass{objc.GetClass("VZStorageDeviceAttachment")}
	})
	return VZStorageDeviceAttachmentClass
}

type _VZStorageDeviceAttachmentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZStorageDeviceAttachment */
// An interface definition for the [VZStorageDeviceAttachment] class.
type IVZStorageDeviceAttachment interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZStorageDeviceAttachment */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZStorageDeviceAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZStorageDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceAttachmentClass) Alloc() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZStorageDeviceAttachmentClass) New() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZStorageDeviceAttachment) Init() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZStorageDeviceAttachment) Autorelease() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceAttachment creates a new VZStorageDeviceAttachment instance.
func NewVZStorageDeviceAttachment() VZStorageDeviceAttachment {
	return getVZStorageDeviceAttachmentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZStorageDeviceAttachment */
// The common behaviors for storage devices in the guest system.
//
// A object defines the implementation of a storage interface in a virtual machine. You use the attachment object to specify the source of the storage on the host computer. Don’t create objects directly. Instead, instantiate an appropriate subclass such as , which provides storage using a disk image.

// The common behaviors for storage devices in the guest system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment
type VZStorageDeviceAttachment struct {
	objectivec.Object
}

// VZStorageDeviceAttachmentFrom constructs a [VZStorageDeviceAttachment] from an unsafe.Pointer.
//
// The common behaviors for storage devices in the guest system.
func VZStorageDeviceAttachmentFrom(ptr unsafe.Pointer) VZStorageDeviceAttachment {
	return VZStorageDeviceAttachment{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZStorageDeviceAttachment */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZStorageDeviceAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZStorageDeviceAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZStorageDeviceAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZStorageDeviceAttachment */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZStorageDeviceAttachment */
