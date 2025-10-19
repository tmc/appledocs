// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
var mTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass = _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray")}

type _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] class.
type IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom constructs a [MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
func MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) Alloc() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass) New() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) Init() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) Autorelease() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray creates a new MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray instance.
func NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	return mTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayClass.New()
}




