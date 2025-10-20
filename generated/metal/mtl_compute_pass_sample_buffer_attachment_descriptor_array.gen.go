// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
var (
	ComputePassSampleBufferAttachmentDescriptorArrayClass     _ComputePassSampleBufferAttachmentDescriptorArrayClass
	ComputePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getComputePassSampleBufferAttachmentDescriptorArrayClass() _ComputePassSampleBufferAttachmentDescriptorArrayClass {
	ComputePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		ComputePassSampleBufferAttachmentDescriptorArrayClass = _ComputePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptorArray")}
	})
	return ComputePassSampleBufferAttachmentDescriptorArrayClass
}

type _ComputePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [ComputePassSampleBufferAttachmentDescriptorArray] class.
type IComputePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer
}

// A container that stores an array of sample buffer attachments for a compute pass.
//
// The number of elements in the array is at least the number of elements in an instance’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray
type ComputePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// ComputePassSampleBufferAttachmentDescriptorArrayFrom constructs a [ComputePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// A container that stores an array of sample buffer attachments for a compute pass.
func ComputePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ComputePassSampleBufferAttachmentDescriptorArray {
	return ComputePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComputePassSampleBufferAttachmentDescriptorArrayClass) New() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) Init() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) Autorelease() ComputePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ComputePassSampleBufferAttachmentDescriptorArray](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComputePassSampleBufferAttachmentDescriptorArray creates a new ComputePassSampleBufferAttachmentDescriptorArray instance.
func NewComputePassSampleBufferAttachmentDescriptorArray() ComputePassSampleBufferAttachmentDescriptorArray {
	return getComputePassSampleBufferAttachmentDescriptorArrayClass().New()
}


// Returns the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (c_ ComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}



