// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZNetworkDeviceAttachment] class.
type IVZNetworkDeviceAttachment interface {
	objectivec.IObject
}

// A parent class referenced by other Virtualization classes.


// A parent class referenced by other Virtualization classes. [Full Topic]
type VZNetworkDeviceAttachment struct {
	objectivec.Object
}

// VZNetworkDeviceAttachmentFrom constructs a [VZNetworkDeviceAttachment] from an unsafe.Pointer.
//
// A parent class referenced by other Virtualization classes.
func VZNetworkDeviceAttachmentFrom(ptr unsafe.Pointer) VZNetworkDeviceAttachment {
	return VZNetworkDeviceAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceAttachmentClass) Alloc() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




