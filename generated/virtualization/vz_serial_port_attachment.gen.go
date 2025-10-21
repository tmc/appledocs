// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZSerialPortAttachment] class.
var (
	VZSerialPortAttachmentClass     _VZSerialPortAttachmentClass
	VZSerialPortAttachmentClassOnce sync.Once
)

func getVZSerialPortAttachmentClass() _VZSerialPortAttachmentClass {
	VZSerialPortAttachmentClassOnce.Do(func() {
		VZSerialPortAttachmentClass = _VZSerialPortAttachmentClass{objc.GetClass("VZSerialPortAttachment")}
	})
	return VZSerialPortAttachmentClass
}

type _VZSerialPortAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZSerialPortAttachment] class.
type IVZSerialPortAttachment interface {
	objectivec.IObject
}

// The common behaviors for the serial attachment points of your virtual machine.
//
// Don’t create a object directly. Instead, instantiate a concrete subclass such as to configure how the virtual machine’s serial port connects with the host computer.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type VZSerialPortAttachment struct {
	objectivec.Object
}

// VZSerialPortAttachmentFrom constructs a [VZSerialPortAttachment] from an unsafe.Pointer.
//
// The common behaviors for the serial attachment points of your virtual machine.
func VZSerialPortAttachmentFrom(ptr unsafe.Pointer) VZSerialPortAttachment {
	return VZSerialPortAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZSerialPortAttachmentClass) Alloc() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZSerialPortAttachmentClass) New() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSerialPortAttachment) Init() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSerialPortAttachment) Autorelease() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortAttachment creates a new VZSerialPortAttachment instance.
func NewVZSerialPortAttachment() VZSerialPortAttachment {
	return getVZSerialPortAttachmentClass().New()
}




