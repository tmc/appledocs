// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [VZDiskImageStorageDeviceAttachment] class.
type IVZDiskImageStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment
}

// A device that stores content in a disk image.
//
// Use a object to manage the storage for a disk in a virtual machine (VM). The guest operating system sees the storage as a disk, and when the guest operating system writes files to the disk, the virtual machine stores the files in the disk image you provide. The virtualization framework supports two disk image formats:
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZDiskImageStorageDeviceAttachmentClass) Alloc() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

// NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError initializes a disk image storage device attachment with a URL, read-only flag, and error parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageStorageDeviceAttachment/init(url:readOnly:error:)
func NewVZDiskImageStorageDeviceAttachmentWithURLReadOnlyError(
	diskImageURL foundation.URL,
	readOnly bool,
	error_ unsafe.Pointer,
) VZDiskImageStorageDeviceAttachment {
	class := getVZDiskImageStorageDeviceAttachmentClass()
	alloc := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(class.class), objc.Sel("alloc"))
	inst := objc.Send[VZDiskImageStorageDeviceAttachment](
		alloc.ID,
		objc.Sel("initWithURL:readOnly:error:"),
		unsafe.Pointer(diskImageURL.ID),
		readOnly,
		error_,
	)
	return inst
}




