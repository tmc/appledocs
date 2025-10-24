// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZDiskImageStorageDeviceAttachment */

/* debug [class_header]: Header for VZDiskImageStorageDeviceAttachment */
// The class instance for the [VZDiskImageStorageDeviceAttachment] class.
var (
	VZDiskImageStorageDeviceAttachmentClass     _VZDiskImageStorageDeviceAttachmentClass
	VZDiskImageStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskImageStorageDeviceAttachmentClass() _VZDiskImageStorageDeviceAttachmentClass {
	VZDiskImageStorageDeviceAttachmentClassOnce.Do(func() {
		VZDiskImageStorageDeviceAttachmentClass = _VZDiskImageStorageDeviceAttachmentClass{objc.GetClass("VZDiskImageStorageDeviceAttachment")}
	})
	return VZDiskImageStorageDeviceAttachmentClass
}

type _VZDiskImageStorageDeviceAttachmentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZDiskImageStorageDeviceAttachment */
// An interface definition for the [VZDiskImageStorageDeviceAttachment] class.
type IVZDiskImageStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	/* debug [class_interface_properties]: Properties for VZDiskImageStorageDeviceAttachment */
	// properties:
	CachingMode() VZDiskImageCachingMode
	ReadOnly() bool
	SynchronizationMode() VZDiskImageSynchronizationMode
	URL() objc.IObject /* cross-framework: NSURL */
	IsReadOnly() bool
	SetIsReadOnly(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZDiskImageStorageDeviceAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZDiskImageStorageDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZDiskImageStorageDeviceAttachmentClass) Alloc() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZDiskImageStorageDeviceAttachmentClass) New() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDiskImageStorageDeviceAttachment) Init() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDiskImageStorageDeviceAttachment) Autorelease() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImageStorageDeviceAttachment creates a new VZDiskImageStorageDeviceAttachment instance.
func NewVZDiskImageStorageDeviceAttachment() VZDiskImageStorageDeviceAttachment {
	return getVZDiskImageStorageDeviceAttachmentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZDiskImageStorageDeviceAttachment */
// A device that stores content in a disk image.
//
// Use a object to manage the storage for a disk in a virtual machine (VM). The guest operating system sees the storage as a disk, and when the guest operating system writes files to the disk, the virtual machine stores the files in the disk image you provide. The virtualization framework supports two disk image formats:

// A device that stores content in a disk image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment
type VZDiskImageStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskImageStorageDeviceAttachmentFrom constructs a [VZDiskImageStorageDeviceAttachment] from an unsafe.Pointer.
//
// A device that stores content in a disk image.
func VZDiskImageStorageDeviceAttachmentFrom(ptr unsafe.Pointer) VZDiskImageStorageDeviceAttachment {
	return VZDiskImageStorageDeviceAttachment{
		VZStorageDeviceAttachment: VZStorageDeviceAttachmentFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZDiskImageStorageDeviceAttachment */

// Initialize the attachment from a local file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:cachingMode:synchronizationMode:)
func NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyCachingModeSynchronizationModeError(url objc.IObject /* cross-framework: NSURL */, readOnly bool, cachingMode VZDiskImageCachingMode, synchronizationMode VZDiskImageSynchronizationMode, error_ unsafe.Pointer) VZDiskImageStorageDeviceAttachment {
	instance := getVZDiskImageStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](instance.ID, objc.Sel("initWithURL:readOnly:cachingMode:synchronizationMode:error:"), url, readOnly, cachingMode, synchronizationMode, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyCachingModeSynchronizationModeError */

// Creates the attachment object from the specified disk image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:)
func NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError(url objc.IObject /* cross-framework: NSURL */, readOnly bool, error_ unsafe.Pointer) VZDiskImageStorageDeviceAttachment {
	instance := getVZDiskImageStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](instance.ID, objc.Sel("initWithURL:readOnly:error:"), url, readOnly, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZDiskImageStorageDeviceAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZDiskImageStorageDeviceAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZDiskImageStorageDeviceAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZDiskImageStorageDeviceAttachment */

// The current cacheing mode for the virtual disk image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/cachingMode
func (v_ VZDiskImageStorageDeviceAttachment) CachingMode() VZDiskImageCachingMode {
	rv := objc.Send[VZDiskImageCachingMode](v_.ID, objc.Sel("cachingMode"))
	return rv
} /* debug [instance_properties/getter]: cachingMode */

// A Boolean value that indicates whether the underlying disk image is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/isReadOnly
func (v_ VZDiskImageStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("readOnly"))
	return rv
} /* debug [instance_properties/getter]: readOnly */

// The mode in which the disk image synchronizes data with the underlying storage device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/synchronizationMode
func (v_ VZDiskImageStorageDeviceAttachment) SynchronizationMode() VZDiskImageSynchronizationMode {
	rv := objc.Send[VZDiskImageSynchronizationMode](v_.ID, objc.Sel("synchronizationMode"))
	return rv
} /* debug [instance_properties/getter]: synchronizationMode */

// The URL of the underlying disk image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/url
func (v_ VZDiskImageStorageDeviceAttachment) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

// A Boolean value that indicates whether the underlying disk image is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/isreadonly
func (v_ VZDiskImageStorageDeviceAttachment) IsReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isReadOnly"))
	return rv
} /* debug [instance_properties/getter]: isReadOnly */

// A Boolean value that indicates whether the underlying disk image is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/isreadonly
func (v_ VZDiskImageStorageDeviceAttachment) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsReadOnly:"), value)
} /* debug [instance_properties/setter]: isReadOnly */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZDiskImageStorageDeviceAttachment */
