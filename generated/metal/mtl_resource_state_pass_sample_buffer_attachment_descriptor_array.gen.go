// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
var (
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClass     _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce sync.Once
)

func getResourceStatePassSampleBufferAttachmentDescriptorArrayClass() _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass {
	ResourceStatePassSampleBufferAttachmentDescriptorArrayClassOnce.Do(func() {
		ResourceStatePassSampleBufferAttachmentDescriptorArrayClass = _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptorArray")}
	})
	return ResourceStatePassSampleBufferAttachmentDescriptorArrayClass
}

type _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass struct {
	class objc.Class
}

// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptorArray] class.
type IResourceStatePassSampleBufferAttachmentDescriptorArray interface {
	objectivec.IObject
	ObjectAtIndexedSubscript(attachmentIndex uint) ResourceStatePassSampleBufferAttachmentDescriptor
}

// An array of sample buffer attachments for a resource state pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray
type ResourceStatePassSampleBufferAttachmentDescriptorArray struct {
	objectivec.Object
}

// ResourceStatePassSampleBufferAttachmentDescriptorArrayFrom constructs a [ResourceStatePassSampleBufferAttachmentDescriptorArray] from an unsafe.Pointer.
//
// An array of sample buffer attachments for a resource state pass.
func ResourceStatePassSampleBufferAttachmentDescriptorArrayFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return ResourceStatePassSampleBufferAttachmentDescriptorArray{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorArrayClass) New() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) Init() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) Autorelease() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptorArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceStatePassSampleBufferAttachmentDescriptorArray creates a new ResourceStatePassSampleBufferAttachmentDescriptorArray instance.
func NewResourceStatePassSampleBufferAttachmentDescriptorArray() ResourceStatePassSampleBufferAttachmentDescriptorArray {
	return getResourceStatePassSampleBufferAttachmentDescriptorArrayClass().New()
}


// Returns the descriptor object for the specified sample buffer attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptorArray/subscript(_:)
func (r_ ResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript(attachmentIndex uint) ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("objectAtIndexedSubscript:"), attachmentIndex)
	return rv
}



