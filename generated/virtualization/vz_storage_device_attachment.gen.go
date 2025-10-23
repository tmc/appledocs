// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZStorageDeviceAttachment] class.
type IVZStorageDeviceAttachment interface {
	objectivec.IObject
}

// A parent class referenced by other Virtualization classes.


// A parent class referenced by other Virtualization classes. [Full Topic]
type VZStorageDeviceAttachment struct {
	objectivec.Object
}

// VZStorageDeviceAttachmentFrom constructs a [VZStorageDeviceAttachment] from an unsafe.Pointer.
//
// A parent class referenced by other Virtualization classes.
func VZStorageDeviceAttachmentFrom(ptr unsafe.Pointer) VZStorageDeviceAttachment {
	return VZStorageDeviceAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZStorageDeviceAttachmentClass) Alloc() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




