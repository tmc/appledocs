// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
var (
	ResourceStatePassSampleBufferAttachmentDescriptorClass     _ResourceStatePassSampleBufferAttachmentDescriptorClass
	ResourceStatePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getResourceStatePassSampleBufferAttachmentDescriptorClass() _ResourceStatePassSampleBufferAttachmentDescriptorClass {
	ResourceStatePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		ResourceStatePassSampleBufferAttachmentDescriptorClass = _ResourceStatePassSampleBufferAttachmentDescriptorClass{objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptor")}
	})
	return ResourceStatePassSampleBufferAttachmentDescriptorClass
}

type _ResourceStatePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ResourceStatePassSampleBufferAttachmentDescriptor] class.
type IResourceStatePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject
	EndOfEncoderSampleIndex() int
	SetEndOfEncoderSampleIndex(value int)
	SampleBuffer() unsafe.Pointer
	SetSampleBuffer(value unsafe.Pointer)
	StartOfEncoderSampleIndex() int
	SetStartOfEncoderSampleIndex(value int)
}

// A description of where to store GPU counter information at the start and end of a resource state pass.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor
type ResourceStatePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// ResourceStatePassSampleBufferAttachmentDescriptorFrom constructs a [ResourceStatePassSampleBufferAttachmentDescriptor] from an unsafe.Pointer.
//
// A description of where to store GPU counter information at the start and end of a resource state pass.
func ResourceStatePassSampleBufferAttachmentDescriptorFrom(ptr unsafe.Pointer) ResourceStatePassSampleBufferAttachmentDescriptor {
	return ResourceStatePassSampleBufferAttachmentDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) Alloc() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResourceStatePassSampleBufferAttachmentDescriptorClass) New() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) Init() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) Autorelease() ResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[ResourceStatePassSampleBufferAttachmentDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResourceStatePassSampleBufferAttachmentDescriptor creates a new ResourceStatePassSampleBufferAttachmentDescriptor instance.
func NewResourceStatePassSampleBufferAttachmentDescriptor() ResourceStatePassSampleBufferAttachmentDescriptor {
	return getResourceStatePassSampleBufferAttachmentDescriptorClass().New()
}


// The index the Metal device object should use to store GPU counters when ending the resource state pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() int {
	rv := objc.Send[int](r_.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}


// SetEndOfEncoderSampleIndex sets the value of the endOfEncoderSampleIndex property.
// The index the Metal device object should use to store GPU counters when ending the resource state pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/endofencodersampleindex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/samplebuffer
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("sampleBuffer"))
	return rv
}


// SetSampleBuffer sets the value of the sampleBuffer property.
// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/samplebuffer
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleBuffer:"), value)
}

// The index the Metal device object should use to store GPU counters when starting the resource state pass.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/startofencodersampleindex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() int {
	rv := objc.Send[int](r_.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}


// SetStartOfEncoderSampleIndex sets the value of the startOfEncoderSampleIndex property.
// The index the Metal device object should use to store GPU counters when starting the resource state pass.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatepasssamplebufferattachmentdescriptor/startofencodersampleindex
func (r_ ResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}



