// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
var (
	VZDiskBlockDeviceStorageDeviceAttachmentClass     _VZDiskBlockDeviceStorageDeviceAttachmentClass
	VZDiskBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskBlockDeviceStorageDeviceAttachmentClass() _VZDiskBlockDeviceStorageDeviceAttachmentClass {
	VZDiskBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		VZDiskBlockDeviceStorageDeviceAttachmentClass = _VZDiskBlockDeviceStorageDeviceAttachmentClass{objc.GetClass("VZDiskBlockDeviceStorageDeviceAttachment")}
	})
	return VZDiskBlockDeviceStorageDeviceAttachmentClass
}

type _VZDiskBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
type IVZDiskBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment
}

// A storage device attachment that uses a disk to store data.
//
// The disk block device implements a storage attachment by using an actual disk rather than a disk image on a file system. In the following example, a disk device at executes the I/O operations directly on that disk rather than through a file system: By default, only the user can access the disk file handle. Running virtual machines as isn’t recommended. The best practice is to open the file in a separate process that has privileges, then pass the open file descriptor using XPC or a Unix socket to a non- process running Virtualization. For more information about Unix sockets, see ; for more information on XPC services, see the framework documentation.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment
type VZDiskBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskBlockDeviceStorageDeviceAttachmentFrom constructs a [VZDiskBlockDeviceStorageDeviceAttachment] from an unsafe.Pointer.
//
// A storage device attachment that uses a disk to store data.
func VZDiskBlockDeviceStorageDeviceAttachmentFrom(ptr unsafe.Pointer) VZDiskBlockDeviceStorageDeviceAttachment {
	return VZDiskBlockDeviceStorageDeviceAttachment{
		VZStorageDeviceAttachment: VZStorageDeviceAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZDiskBlockDeviceStorageDeviceAttachmentClass) Alloc() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZDiskBlockDeviceStorageDeviceAttachmentClass) New() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) Init() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) Autorelease() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskBlockDeviceStorageDeviceAttachment creates a new VZDiskBlockDeviceStorageDeviceAttachment instance.
func NewVZDiskBlockDeviceStorageDeviceAttachment() VZDiskBlockDeviceStorageDeviceAttachment {
	return getVZDiskBlockDeviceStorageDeviceAttachmentClass().New()
}




