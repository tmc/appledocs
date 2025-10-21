// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


// The current cacheing mode for the virtual disk image.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/cachingmode
func (v_ VZDiskImageStorageDeviceAttachment) CachingMode() VZDiskImageCachingMode {
	rv := objc.Send[VZDiskImageCachingMode](v_.ID, objc.Sel("cachingMode"))
	return rv
}


// SetCachingMode sets the value of the cachingMode property.
// The current cacheing mode for the virtual disk image.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/cachingmode
func (v_ VZDiskImageStorageDeviceAttachment) SetCachingMode(value VZDiskImageCachingMode) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCachingMode:"), value)
}

// A Boolean value that indicates whether the underlying disk image is read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/isreadonly
func (v_ VZDiskImageStorageDeviceAttachment) IsReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isReadOnly"))
	return rv
}


// SetIsReadOnly sets the value of the isReadOnly property.
// A Boolean value that indicates whether the underlying disk image is read-only.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/isreadonly
func (v_ VZDiskImageStorageDeviceAttachment) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsReadOnly:"), value)
}

// The mode in which the disk image synchronizes data with the underlying storage device.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/synchronizationmode
func (v_ VZDiskImageStorageDeviceAttachment) SynchronizationMode() VZDiskImageSynchronizationMode {
	rv := objc.Send[VZDiskImageSynchronizationMode](v_.ID, objc.Sel("synchronizationMode"))
	return rv
}


// SetSynchronizationMode sets the value of the synchronizationMode property.
// The mode in which the disk image synchronizes data with the underlying storage device.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/synchronizationmode
func (v_ VZDiskImageStorageDeviceAttachment) SetSynchronizationMode(value VZDiskImageSynchronizationMode) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSynchronizationMode:"), value)
}

// The URL of the underlying disk image.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/url
func (v_ VZDiskImageStorageDeviceAttachment) Url() foundation.URL {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL of the underlying disk image.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskimagestoragedeviceattachment/url
func (v_ VZDiskImageStorageDeviceAttachment) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUrl:"), value)
}



