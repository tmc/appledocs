// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZDiskBlockDeviceStorageDeviceAttachment */


/* debug [class_header]: Header for VZDiskBlockDeviceStorageDeviceAttachment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZDiskBlockDeviceStorageDeviceAttachment */
// An interface definition for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
type IVZDiskBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment
	
/* debug [class_interface_properties]: Properties for VZDiskBlockDeviceStorageDeviceAttachment */
	// properties:
	FileHandle() foundation.FileHandle
	ReadOnly() bool
	SynchronizationMode() VZDiskSynchronizationMode
	IsReadOnly() bool
	SetIsReadOnly(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZDiskBlockDeviceStorageDeviceAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZDiskBlockDeviceStorageDeviceAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZDiskBlockDeviceStorageDeviceAttachmentClass) Alloc() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZDiskBlockDeviceStorageDeviceAttachment */
// A storage device attachment that uses a disk to store data.
//
// The disk block device implements a storage attachment by using an actual disk rather than a disk image on a file system. In the following example, a disk device at executes the I/O operations directly on that disk rather than through a file system: By default, only the user can access the disk file handle. Running virtual machines as isn’t recommended. The best practice is to open the file in a separate process that has privileges, then pass the open file descriptor using XPC or a Unix socket to a non- process running Virtualization. For more information about Unix sockets, see ; for more information on XPC services, see the framework documentation.


// A storage device attachment that uses a disk to store data.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZDiskBlockDeviceStorageDeviceAttachment */

// Creates a new block storage device attachment from a file handle and with the specified access mode, synchronization mode, and error object that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/init(fileHandle:readOnly:synchronizationMode:)
func NewVZDiskBlockDeviceStorageDeviceAttachmentWithFileHandleReadOnlySynchronizationModeError(fileHandle foundation.FileHandle, readOnly bool, synchronizationMode VZDiskSynchronizationMode, error_ objectivec.IObject) VZDiskBlockDeviceStorageDeviceAttachment {
	instance := getVZDiskBlockDeviceStorageDeviceAttachmentClass().Alloc()
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](instance.ID, objc.Sel("initWithFileHandle:readOnly:synchronizationMode:error:"), fileHandle, readOnly, synchronizationMode, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZDiskBlockDeviceStorageDeviceAttachmentWithFileHandleReadOnlySynchronizationModeError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZDiskBlockDeviceStorageDeviceAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZDiskBlockDeviceStorageDeviceAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZDiskBlockDeviceStorageDeviceAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZDiskBlockDeviceStorageDeviceAttachment */

// A file handle to a block device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/fileHandle
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) FileHandle() foundation.FileHandle {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandle"))
	return rv
}/* debug [instance_properties/getter]: fileHandle */


// A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/isReadOnly
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("readOnly"))
	return rv
}/* debug [instance_properties/getter]: readOnly */


// The value that defines how the disk synchronizes with the underlying storage when the guest operating system flushes data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/synchronizationMode
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) SynchronizationMode() VZDiskSynchronizationMode {
	rv := objc.Send[VZDiskSynchronizationMode](v_.ID, objc.Sel("synchronizationMode"))
	return rv
}/* debug [instance_properties/getter]: synchronizationMode */


// A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskblockdevicestoragedeviceattachment/isreadonly
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) IsReadOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isReadOnly"))
	return rv
}/* debug [instance_properties/getter]: isReadOnly */


// A Boolean value that indicates whether this disk attachment is read-only; otherwise, if the file handle allows writes, the device can write data into it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzdiskblockdevicestoragedeviceattachment/isreadonly
func (v_ VZDiskBlockDeviceStorageDeviceAttachment) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsReadOnly:"), value)
}/* debug [instance_properties/setter]: isReadOnly */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZDiskBlockDeviceStorageDeviceAttachment */


