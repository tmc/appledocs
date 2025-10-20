// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
var (
	BlitPassSampleBufferAttachmentDescriptorArrayClass     _BlitPassSampleBufferAttachmentDescriptorArrayClass
	BlitPassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getBlitPassSampleBufferAttachmentDescriptorArrayClass() _BlitPassSampleBufferAttachmentDescriptorArrayClass {
	BlitPassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		BlitPassSampleBufferAttachmentDescriptorArrayClass = _BlitPassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptorArray")}
	})
	return BlitPassSampleBufferAttachmentDescriptorArrayClass
}

type _BlitPassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [BlitPassSampleBufferAttachmentDescriptorArray] class.
type IBlitPassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
}

// A container that stores an array of sample buffer attachments for a blit pass.
//
// The number of elements in the array is at least the number of elements in an instance’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptorArray
type BlitPassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// BlitPassSampleBufferAttachmentDescriptorArrayFrom constructs a [BlitPassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// A container that stores an array of sample buffer attachments for a blit pass.
func BlitPassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) BlitPassSampleBufferAttachmentDescriptorArray {
	return BlitPassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) Alloc() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BlitPassSampleBufferAttachmentDescriptorArrayClass) New() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) Init() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlitPassSampleBufferAttachmentDescriptorArray) Autorelease() BlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlitPassSampleBufferAttachmentDescriptorArray creates a new BlitPassSampleBufferAttachmentDescriptorArray instance.
func NewBlitPassSampleBufferAttachmentDescriptorArray() BlitPassSampleBufferAttachmentDescriptorArray {
	return getBlitPassSampleBufferAttachmentDescriptorArrayClass().New()
}




